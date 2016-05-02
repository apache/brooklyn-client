/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Network struct {
	BrooklynUrl  string
	BrooklynUser string
	BrooklynPass string
}

func NewNetwork(brooklynUrl, brooklynUser, brooklynPass string) (net *Network) {
	net = new(Network)
	net.BrooklynUrl = brooklynUrl
	net.BrooklynUser = brooklynUser
	net.BrooklynPass = brooklynPass
	return
}

func (net *Network) NewRequest(method, path string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, net.BrooklynUrl+path, body)
	req.SetBasicAuth(net.BrooklynUser, net.BrooklynPass)
	return req
}

func (net *Network) NewGetRequest(url string) *http.Request {
	return net.NewRequest("GET", url, nil)
}

func (net *Network) NewPostRequest(url string, body io.Reader) *http.Request {
	return net.NewRequest("POST", url, body)
}

func (net *Network) NewDeleteRequest(url string) *http.Request {
	return net.NewRequest("DELETE", url, nil)
}

type HttpError struct {
	Code    int
	Status  string
	Headers http.Header
	Body    string
}

func (err HttpError) Error() string {
	return err.Status
}

func makeError(resp *http.Response, code int, body []byte) error {
	theError := HttpError{
		Code:    code,
		Status:  resp.Status,
		Headers: resp.Header,
	}
	details := make(map[string]string)
	if err := json.Unmarshal(body, &details); nil == err {
		if message, ok := details["message"]; ok {
			theError.Body = message
			return theError
		}
	}
	theError.Body = string(body)
	return theError
}

func (net *Network) SendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if code, failed := unsuccessful(resp.Status); failed {
		return nil, makeError(resp, code, body)
	}
	return body, err
}

const httpSuccessSeriesFrom = 200
const httpSuccessSeriesTo = 300

func unsuccessful(status string) (int, bool) {
	tokens := strings.Split(status, " ")
	if 0 == len(tokens) {
		return -1, false
	}
	code, err := strconv.Atoi(tokens[0])
	if nil != err {
		return -1, false
	}
	return code, code < httpSuccessSeriesFrom || httpSuccessSeriesTo <= code
}

func (net *Network) SendGetRequest(url string) ([]byte, error) {
	req := net.NewGetRequest(url)
	req.Header.Set("Accept", "application/json, text/plain")
	body, err := net.SendRequest(req)
	return body, err
}

func (net *Network) SendDeleteRequest(url string) ([]byte, error) {
	req := net.NewDeleteRequest(url)
	body, err := net.SendRequest(req)
	return body, err
}

func (net *Network) SendEmptyPostRequest(url string) ([]byte, error) {
	req := net.NewPostRequest(url, nil)
	body, err := net.SendRequest(req)
	return body, err
}

func (net *Network) SendPostRequest(urlStr string, data []byte) ([]byte, error) {
	req := net.NewPostRequest(urlStr, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	body, err := net.SendRequest(req)
	return body, err
}

func (net *Network) SendPostResourceRequest(restUrl string, resourceUrl string, contentType string) ([]byte, error) {
	resource, err := openResource(resourceUrl)
	defer resource.Close()
	req := net.NewPostRequest(restUrl, resource)
	req.Header.Set("Content-Type", contentType)
	body, err := net.SendRequest(req)
	return body, err
}

func openResource(resourceUrl string) (io.ReadCloser, error) {
	u, err := url.Parse(resourceUrl)
	if err != nil {
		return nil, err
	}
	if "" == u.Scheme || "file" == u.Scheme {
		return openFileResource(u)

	} else if "http" == u.Scheme || "https" == u.Scheme {
		return openHttpResource(resourceUrl)

	} else {
		return nil, errors.New("Unrecognised protocol scheme: " + u.Scheme)
	}
}

func openFileResource(url *url.URL) (io.ReadCloser, error) {
	filePath := url.Path;
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func openHttpResource(resourceUrl string) (io.ReadCloser, error) {
	resp, err := http.Get(resourceUrl)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}


func VerifyLoginURL(network *Network) error {
	url, err := url.Parse(network.BrooklynUrl)
	if err != nil {
		return err
	}
	if url.Scheme != "http" && url.Scheme != "https" {
		return errors.New("Use login command to set Brooklyn URL with a scheme of \"http\" or \"https\"")
	}
	if url.Host == "" {
		return errors.New("Use login command to set Brooklyn URL with a valid host[:port]")
	}
	return nil
}
