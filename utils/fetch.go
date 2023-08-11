package utils

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

var (
	jar, _ = cookiejar.New(nil)

	client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: jar,
	}
)

func fetchWithRequest(req *http.Request) (*http.Response, error) {
	return client.Do(req)
}

func generateRequest(method, url string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(
		method,
		url,
		body,
	)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func fetch(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := generateRequest(method, url, headers, body)
	if err != nil {
		return nil, err
	}
	return fetchWithRequest(req)
}

func Fetch(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	return fetch(method, url, headers, body)
}
