package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
}

type Response struct {
	Code    int
	Content []byte
}

func NewClient() *Client {
	return &Client{
		HttpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (client *Client) Execute(req *http.Request) *Response {
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return &Response{Code: resp.StatusCode, Content: buf}
}
