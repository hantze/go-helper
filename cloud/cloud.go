package cloud

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPRequest ...
type HTTPRequest struct {
	host       string
	httpClient *http.Client
}

func (hr *HTTPRequest) getURL(path string) string {
	return fmt.Sprintf("%s/%s", hr.host, path)
}

// Get ...
func (hr *HTTPRequest) Get(path string, params *map[string]string) ([]byte, error) {
	url := hr.getURL(path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if params != nil && len(*params) != 0 {
		q := req.URL.Query()
		for key, value := range *params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}
	response, err := hr.httpClient.Do(req)
	return readResponse(response, err)
}

// GetWithHeader ...
func (hr *HTTPRequest) GetWithHeader(path string, params *map[string]string, headers *map[string]string) (int, []byte, error) {
	url := hr.getURL(path)
	req, err := http.NewRequest("GET", url, nil)
	if headers != nil && len(*headers) != 0 {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}
	if params != nil && len(*params) != 0 {
		q := req.URL.Query()
		for key, value := range *params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}
	if err != nil {
		return -1, nil, err
	}
	response, err := hr.httpClient.Do(req)
	r, e := readResponse(response, err)
	return response.StatusCode, r, e
}

// Post ...
func (hr *HTTPRequest) Post(path string, params *[]byte) ([]byte, error) {
	url := hr.getURL(path)
	var body io.Reader
	if body != nil {
		body = bytes.NewBuffer(*params)
	} else {
		body = nil
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	response, err := hr.httpClient.Do(req)
	return readResponse(response, err)
}

// PostJSON ...
func (hr *HTTPRequest) PostJSON(path string, params []byte, header string) (int, []byte, error) {
	url := hr.getURL(path)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	req.Header.Set("X-Custom-Header", header)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return 400, nil, err
	}
	response, err := hr.httpClient.Do(req)
	if err != nil {
		return 400, nil, err
	}
	resp, err := readResponse(response, err)
	return response.StatusCode, resp, err
}

func readResponse(response *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode < 200 && response.StatusCode > 299 {
		return nil, errors.New(string(body))
	}
	return body, err
}

// NewHTTPRequest ...
func NewHTTPRequest(host string, httpClient *http.Client) *HTTPRequest {
	if httpClient == nil {
		timeout := time.Duration(5 * time.Second)
		httpClient = &http.Client{
			Timeout: timeout,
		}
	}
	return &HTTPRequest{host: host, httpClient: httpClient}
}
