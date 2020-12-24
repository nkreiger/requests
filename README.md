# Requests Package
GoLang package responsible for handling all types of custom and general API Requests


# Installation Instructions

`go get github.com/nkreiger/requests`

# Usage

```go
package whatever

import (
	"github.com/nkreiger/requests"
	"io"
	"log"
)

// Request executes an API call via HTTP request
// basic request
// returns response body as a []byte
func Request(method, url string, payload io.Reader) ([]byte, error) {
	requestClient := requests.NewRequestClient(method, url, payload, requests.DefaultRequest())
	
	req, err := requestClient.Request()
	
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	
	res, err := requests.ExecuteRequest(req)
	
	return res, err
}

// request executes an API call via HTTP request
// custom request
// returns response body as a []byte
func RequestCustom(method, url string, payload io.Reader, headers map[string]string) ([]byte, error) {
	requestClient := requests.NewRequestClient(method, url, payload, requests.CustomHeaders(headers))

	req, err := requestClient.Request()

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	res, err := requests.ExecuteRequest(req)

	return res, err
}
```