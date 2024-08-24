package pkg

import (
	"io"
	"net/http"
)

type Http struct {
	http   *http.Client
	Header http.Header
}

func NewHttp() *Http {
	return &Http{
		http: &http.Client{},
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}
}

func (c *Http) Get(url string) ([]byte, int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return c.sendRequest(req)
}

func (c *Http) Post(url string, data io.Reader) ([]byte, int, error) {
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return c.sendRequest(req)
}

func (c *Http) Put(url string, data io.Reader) ([]byte, int, error) {
	req, err := http.NewRequest("PUT", url, data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return c.sendRequest(req)
}

func (c *Http) Delete(url string, data io.Reader) ([]byte, int, error) {
	req, err := http.NewRequest("DELETE", url, data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return c.sendRequest(req)
}

func (c *Http) sendRequest(req *http.Request) ([]byte, int, error) {
	req.Header = c.Header
	res, err := c.http.Do(req)
	if err != nil {
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}
	return body, res.StatusCode, nil
}
