package net

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"errors"
    "strings"
    "strconv"
    "encoding/json"
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


func makeError (resp *http.Response, code int, body []byte) error {
    theError := HttpError {
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
	if code, failed := unsuccessful(resp.Status) ; failed {
		return nil, makeError(resp, code, body)
	}
	return body, err
}


const httpSuccessSeriesFrom = 200;
const httpSuccessSeriesTo = 300;
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

func (net *Network) SendGetRequestWithHeaders(url string, headers map[string]string) ([]byte, error) {
	req := net.NewGetRequest(url)
	for header, value := range headers {
		req.Header.Set(header, value)
	}
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

func (net *Network) SendPostFileRequest(url, filePath string, contentType string) ([]byte, error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	req := net.NewPostRequest(url, file)
	req.Header.Set("Content-Type", contentType)
	body, err := net.SendRequest(req)
	return body, err
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
