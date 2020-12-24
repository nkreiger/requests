package requests

import "net/http"

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

func CustomHeaders(headers map[string]string) RequestModifier {
	return func(r *http.Request) {
		for k, v := range headers {
			r.Header.Add(k, v)
		}
	}
}