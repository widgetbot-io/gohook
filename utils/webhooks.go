package utils

import (
	"io"
	"io/ioutil"
	"net/http"
)

func Parse(r *http.Request) ([]byte, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	payload, _ := ioutil.ReadAll(r.Body)

	return payload, nil
}
