package net

import(
	"net/http"
	"io/ioutil"
	"io"
	"fmt"
)

type Network struct {
	BrooklynUrl string
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

func (net *Network) NewRequest(method, path string, body io.Reader) *http.Request{
	req, _ := http.NewRequest(method, net.BrooklynUrl + path, body)
	req.SetBasicAuth(net.BrooklynUser, net.BrooklynPass)
	return req
}

func (net *Network) NewGetRequest(url string) *http.Request{
	return net.NewRequest("GET", url, nil)
}

func (net *Network) NewPostRequest(url string, body io.Reader) *http.Request{
	return net.NewRequest("POST", url, body)
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