package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"math/rand"
	"net/http"
)

// RandomString - random character by n number
func RandomString(n int) string {
	var characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]byte, n)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

// HTTPRequest - make a http request
func HTTPRequest(m string, URL string, d map[string]interface{}, h map[string]interface{}) (*http.Response, error) {
	jsonData, _ := json.Marshal(d)
	buffer := bytes.NewBuffer(jsonData)
	request, err := http.NewRequest(m, URL, buffer)
	if err != nil {
		return nil, err
	}
	for key, val := range h {
		request.Header.Add(key, val.(string))
	}
	// setting http client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
