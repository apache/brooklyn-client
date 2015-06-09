package net

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

func (net *Network) SendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.Status != "200 OK" {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", string(body))
	}
	return body, err
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

func (net *Network) SendPostFileRequest(url, filePath string) ([]byte, error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	req := net.NewPostRequest(url, file)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body, err := net.SendRequest(req)
	return body, err
}
