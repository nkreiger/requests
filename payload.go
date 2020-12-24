package requests

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToReader(payload interface{}) io.Reader {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil
	}

	return bytes.NewReader(data)
}