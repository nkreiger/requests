package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ExecuteRequest(request *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("attempted: %v \n error: %v", request.URL.String(), err)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Printf("Non 200 Code: %v", res.StatusCode)
		return nil, fmt.Errorf("non 200 Return Code: %v", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}