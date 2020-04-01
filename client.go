package hometax

import (
	"context"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	context.Context
	httpCli *http.Client
}

func NewClient(c context.Context) *Client {
	cookieJar, _ := cookiejar.New(nil)

	return &Client{
		c,
		&http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           cookieJar,
			Timeout:       0,
		},
	}
}
