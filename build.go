package requests

import (
	"io"
	"net/http"
)

type RequestClient struct {
	URL	string
	Method string
	Body io.Reader
	modifiers []RequestModifier
}

func NewRequestClient(method, url string, body io.Reader, opts ...RequestModifier) *RequestClient {
	return &RequestClient{
		URL:       url,
		Method:    method,
		Body:      body,
		modifiers: opts,
	}
}

func (c *RequestClient) Request() (*http.Request, error) {
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		return nil, err
	}

	for _, mod := range c.modifiers {
		mod(req)
	}

	return req, nil
}