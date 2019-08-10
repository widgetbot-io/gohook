package utils

import (
	"io"
	"io/ioutil"
)

func Parse(body io.ReadCloser) ([]byte, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, body)
		_ = body.Close()
	}()

	payload, _ := ioutil.ReadAll(body)

	return payload, nil
}
