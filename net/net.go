package net

import(
	"net/http"
	"io/ioutil"
	"io"
	"fmt"
)

func NewRequest(method, path string, body io.Reader) *http.Request{
	req, _ := http.NewRequest(method, path, body)
	req.SetBasicAuth("brooklyn", "Sns4Hh9j7l")
	return req
}

func NewGetRequest(url string) *http.Request{
	return NewRequest("GET", url, nil)
}

func NewPostRequest(url string, body io.Reader) *http.Request{
	return NewRequest("POST", url, body)
}

func SendRequest(req *http.Request) ([]byte, error) {
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