package http_util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Post(url string, req, resp interface{}, options ...HttpOption) error {
	client := http.Client{}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-type", "application/json")
	for _, opts := range options {
		opts(httpReq)
	}

	ret, err := client.Do(httpReq)
	if err != nil {
		return err
	}

	defer ret.Body.Close()

	retBytes, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(retBytes, resp)
	if err != nil {
		return err
	}

	return nil
}

func Get(url string, resp interface{}, options ...HttpOption) error {
	client := http.Client{}

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-type", "application/json")
	for _, opt := range options {
		opt(httpReq)
	}

	ret, err := client.Do(httpReq)
	if err != nil {
		return err
	}

	defer ret.Body.Close()

	retBytes, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(retBytes, resp)
	if err != nil {
		return err
	}

	return nil
}

type HttpOption func(req *http.Request)

func WithBearer(token string) HttpOption {
	return func(req *http.Request) {
		bearer := "Bearer " + token

		req.Header.Set("Authorization", bearer)
	}
}

func WithAccept(accept string) HttpOption {
	return func(req *http.Request) {
		req.Header.Set("Accept", accept)
	}
}

func WithContentType(contentType string) HttpOption {
	return func(req *http.Request) {
		req.Header.Set("Content-type", contentType)
	}
}
