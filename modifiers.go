package requests

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type RequestModifier func(r *http.Request)

func DefaultRequest() RequestModifier {
	return func(r *http.Request) {
		r.Header.Add("Cache-Control", "no-cache")
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/json")
	}
}

func CustomHeaderValue(key string, value string) RequestModifier {
	return func(r *http.Request) {
		r.Header.Add(key, value)
	}
}

func SetBasicAuth(username, password string) RequestModifier {
	return func(r *http.Request) {
		auth := fmt.Sprintf("%s:%s", username, password)
		r.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth))))
	}
}

func CustomHeaders(headers map[string]string) RequestModifier {
	return func(r *http.Request) {
		for k, v := range headers {
			r.Header.Add(k, v)
		}
	}
}