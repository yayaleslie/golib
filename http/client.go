package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	client *http.Client
}

func (c *Client) Get(ctx context.Context, url string, header map[string]string, body []byte) (int, []byte, error) {
	return c.Do(ctx, http.MethodGet, url, header, body, true)
}

func (c *Client) Post(ctx context.Context, url string, header map[string]string, body []byte) (int, []byte, error) {
	return c.Do(ctx, http.MethodPost, url, header, body, true)
}

func (c *Client) Do(ctx context.Context, method string, url string, header map[string]string, body []byte, checkStatus bool) (respCode int, respBody []byte, err error) {
	var req *http.Request

	if len(body) > 0 {
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
	}

	if err != nil {
		return
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respCode = resp.StatusCode

	if checkStatus && respCode != 200 {
		err = fmt.Errorf("error http code %d", respCode)
		return
	}

	respBody, err = io.ReadAll(resp.Body)
	return
}

func (c *Client) DoWithReader(ctx context.Context, method string, url string, header map[string]string, body io.Reader, checkStatus bool) (respCode int, respBody []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respCode = resp.StatusCode
	if checkStatus && respCode != 200 {
		err = fmt.Errorf("error http code %d", respCode)
		return
	}

	respBody, err = io.ReadAll(resp.Body)
	return
}

func NewClient(timeout time.Duration, clients ...*http.Client) *Client {
	cookie, _ := cookiejar.New(nil)
	client := &Client{client: &http.Client{Jar: cookie, Timeout: timeout}}
	if len(clients) > 0 && clients[0] != nil {
		if clients[0].Transport != nil {
			client.client.Transport = clients[0].Transport
		}
		if clients[0].CheckRedirect != nil {
			client.client.CheckRedirect = clients[0].CheckRedirect
		}
		if clients[0].Jar != nil {
			client.client.Jar = clients[0].Jar
		}
		if clients[0].Timeout > 0 {
			client.client.Timeout = clients[0].Timeout
		}
	}
	return client
}
