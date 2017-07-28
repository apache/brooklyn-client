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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"crypto/tls"
	"net"
	"time"
)

type Network struct {
	BrooklynUrl  string
	BrooklynUser string
	BrooklynPass string
	SkipSslChecks bool
	Verbosity    string
}

func NewNetwork(brooklynUrl, brooklynUser, brooklynPass string, skipSslChecks bool, verbose string) (net *Network) {
	net = new(Network)
	net.BrooklynUrl = brooklynUrl
	net.BrooklynUser = brooklynUser
	net.BrooklynPass = brooklynPass
	net.SkipSslChecks = skipSslChecks
	net.Verbosity = verbose
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
	return makeErrorBody(theError, body)
}

func makeSimpleError(code int, body []byte) error {
	theError := HttpError{
		Code:    code,
	}
	return makeErrorBody(theError, body)
}

func makeErrorBody(theError HttpError, body []byte) error {
	details := make(map[string]interface{})
	if err := json.Unmarshal(body, &details); nil == err {
		if message, ok := details["message"]; ok {
			theError.Body = message.(string)
			return theError
		}
	}
	theError.Body = string(body)
	return theError
}

func (net *Network) SendRequest(req *http.Request) ([]byte, error) {
	body, _, err := net.SendRequestGetStatusCode(req)
	return body, err
}

func (net *Network) makeClient() (*http.Client) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: net.SkipSslChecks},
	}
	client := &http.Client{Transport: tr}
	return client
}

func debug(verbosity string, supp func(b bool) ([]byte, error)) {
	writer := func(data []byte, err error) {
		if err == nil {
			fmt.Fprintf(os.Stderr, "%s", data)
			// include newline if data doesn't have one
			if data[len(data)-1] != '\n' {
				fmt.Fprintln(os.Stderr, "")
			}
		} else {
			log.Fatalf("%s\n", err)
		}
	}
	switch verbosity {
	case "verbose":
		writer(supp(false))
	case "vverbose":
		writer(supp(true))
	}
}

func (net *Network) SendRequestGetStatusCode(req *http.Request) ([]byte, int, error) {
	client := net.makeClient()
	debug(net.Verbosity, func (includeBody bool) ([]byte, error) {
		var authHeader = req.Header.Get("Authorization")
		if authHeader != "" {
			req.Header.Set("Authorization", "******")
		}
		data, err := httputil.DumpRequestOut(req, includeBody)
		if authHeader != "" {
			req.Header.Set("Authorization", authHeader)
		}
		return data, err
	})
	resp, err := client.Do(req)
	debug(net.Verbosity, func (includeBody bool) ([]byte, error) {
		return httputil.DumpResponse(resp, includeBody)
	})
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = nil
	}
	if failed := unsuccessful(resp.StatusCode); failed {
		return nil, resp.StatusCode, makeError(resp, resp.StatusCode, body)
	}
	return body, resp.StatusCode, err
}

const httpSuccessSeriesFrom = 200
const httpSuccessSeriesTo = 300

func unsuccessful(code int) (bool) {
	return code < httpSuccessSeriesFrom || httpSuccessSeriesTo <= code
}

func (net *Network) SendGetRequest(url string) ([]byte, error) {
	req := net.NewGetRequest(url)
	req.Header.Set("Accept", "application/json, text/plain")
	body, err := net.SendRequest(req)
	return body, err
}

func (net *Network) SendDeleteRequest(url string) ([]byte, error) {
	req := net.NewDeleteRequest(url)
	body, code, err := net.SendRequestGetStatusCode(req)
	if nil != err {
		return nil, err
	}
	if unsuccessful(code) {
		return nil, makeSimpleError(code, body)
	}
	return body, err
}

func (net *Network) SendEmptyPostRequest(url string) ([]byte, error) {
	req := net.NewPostRequest(url, nil)
	body, err := net.SendRequest(req)
	return body, err
}

func (net *Network) SendPostRequestWithContentType(urlStr string, data []byte, contentType string) ([]byte, error) {
	req := net.NewPostRequest(urlStr, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", contentType)
	body, err := net.SendRequest(req)
	return body, err
}

func (net *Network) SendPostRequest(urlStr string, data []byte) ([]byte, error) {
	return net.SendPostRequestWithContentType(urlStr, data, "application/json")
}

func (net *Network) SendPostResourceRequest(restUrl string, resourceUrl string, contentType string) ([]byte, error) {
	resource, err := net.openResource(resourceUrl)
	if err != nil {
		return nil, err
	}
	defer resource.Close()
	req := net.NewPostRequest(restUrl, resource)
	req.Header.Set("Content-Type", contentType)
	body, err := net.SendRequest(req)
	return body, err
}

func (net *Network) openResource(resourceUrl string) (io.ReadCloser, error) {
	u, err := url.Parse(resourceUrl)
	if err != nil {
		return nil, err
	}
	if "" == u.Scheme || "file" == u.Scheme {
		return net.openFileResource(u)

	} else if "http" == u.Scheme || "https" == u.Scheme {
		return net.openHttpResource(resourceUrl)

	} else {
		return nil, errors.New("Unrecognised protocol scheme: " + u.Scheme)
	}
}

func (net *Network) openFileResource(url *url.URL) (io.ReadCloser, error) {
	filePath := url.Path;
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (net *Network) openHttpResource(resourceUrl string) (io.ReadCloser, error) {
	client := net.makeClient()
	resp, err := client.Get(resourceUrl)
	if err != nil {
		return nil, err
	}
	if failed := unsuccessful(resp.StatusCode) ; failed {
		return nil, errors.New("Error retrieving " + resourceUrl + " (" + resp.Status + ")")
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
	_, err = net.DialTimeout("tcp", url.Host, time.Duration(30) * time.Second)
	if err != nil {
		return errors.New("Could not connect to " + url.Host)
	}
	return nil
}
