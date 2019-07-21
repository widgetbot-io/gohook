package utils

import (
	"encoding/json"
	"git.deploys.io/disweb/gohook/structs"
	"io"
	"io/ioutil"
	"net/http"
)

func Parse(r *http.Request) (structs.BaseDetection, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	payload, _ := ioutil.ReadAll(r.Body)

	var base structs.BaseDetection
	err := json.Unmarshal([]byte(payload), &base)

	if err != nil {
		panic(err)
	}

	return base, nil
}
