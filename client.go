package dtrobot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type headerHandle func(headers Headers)

type Headers map[string]string

type Response map[string]interface{}

func Post(url string, data []byte, handles... headerHandle) (Response, error) {
	headers := make(Headers)
	for _, handle := range handles {
		handle(headers)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%s", resp.Status))
	}


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := make(Response)

	json.Unmarshal(body, &response)

	return response, nil
}

func WithHeader(key, value string) headerHandle {
	return func(headers Headers) {
		headers[key] = value
	}
}
