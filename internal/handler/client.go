// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package handler

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	log "github.com/yeeaiclub/dify-go/internal/logger"

	goquery "github.com/google/go-querystring/query"
)

const (
	// defaultTimeout is the default timeout duration for the http client (in seconds).
	defaultTimeout = 30
)

// ClientOptions defines config options for the client.
type ClientOptions struct {
	Timeout time.Duration
}

// ClientOption defines a functional option for configuring the client.
type ClientOption func(options *ClientOptions)

// WithTimeout sets the timeout duration for the http client.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(options *ClientOptions) {
		options.Timeout = timeout
	}
}

// Client is a http client that execute requests.
type Client struct {
	client *http.Client
}

// NewClient returns a client to execute requests.
func NewClient(opts ...ClientOption) *Client {
	opt := &ClientOptions{
		Timeout: defaultTimeout * time.Second,
	}

	for _, option := range opts {
		option(opt)
	}

	client := &http.Client{Timeout: opt.Timeout}
	return &Client{client: client}
}

// Send sends a HTTP request and returns response.
func (c *Client) Send(ctx context.Context, req Request) (*Response, error) {
	httpReq, err := c.buildRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	return c.doRequest(httpReq)
}

// SendStream sends an HTTP request and returns streaming responses via the provided channel.
func (c *Client) SendStream(ctx context.Context, req Request, evh chan *Event) error {
	httpReq, err := c.buildRequest(ctx, req)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Accept", "text/event-stream")
	httpReq.Header.Set("Cache-Control", "no-cache")
	httpReq.Header.Set("Connection", "keep-alive")
	return c.doStreamRequest(ctx, httpReq, evh)
}

// marshalBody serializes the request body to JSON.
func (c *Client) marshalBody(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	return bytes.NewBuffer(jsonData), nil
}

// buildRequest builds a new HTTP request with the given parameters.
func (c *Client) buildRequest(ctx context.Context, req Request) (*http.Request, error) {
	// Build the request URL
	u, err := buildURL(req.BaseURL, req.Path, req.Query)
	if err != nil {
		return nil, err
	}

	reqBody, err := c.marshalBody(req.Body)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.Method, u, reqBody)
	if err != nil {
		return nil, err
	}

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+req.AuthToken)

	return httpReq, nil
}

// doRequest executes the HTTP request and returns the processed response.
func (c *Client) doRequest(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req) //nolint:nolintlint
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Errorf("failed to close the http body: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to doStreamRequest response body: %w", err)
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
	}, nil
}

// doStreamRequest executes the HTTP request and returns streaming responses via the provided channel.
func (c *Client) doStreamRequest(ctx context.Context, req *http.Request, evCh chan *Event) error {
	ctxReq := req.WithContext(ctx)
	resp, err := c.client.Do(ctxReq) //nolint:bodyclose // ignore the error, it will be handled in the next step
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error resp code: %d", resp.StatusCode)
	}

	go func() {
		defer func() {
			err = resp.Body.Close()
			if err != nil {
				log.Errorf("failed to close the http body: %v", err)
			}
		}()
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					evCh <- &Event{Done: true}
					return
				}
				log.Errorf("failed to read: %v", err)
				return
			}
			line = bytes.TrimRight(line, "\r\n")
			if bytes.IndexByte(line, ':') == -1 {
				continue
			}
			field := line
			value := []byte{}
			if i := bytes.IndexByte(line, ':'); i >= 0 {
				field = line[:i]
				if i+1 < len(line) {
					value = line[i+1:]
					if len(value) > 0 && value[0] == ' ' {
						value = value[1:]
					}
				}
			}

			if string(field) == "data" {
				evCh <- &Event{Data: bytes.NewBuffer(value), Type: string(field)}
			}
		}
	}()
	return nil
}

// buildURL constructs a complete URL from base URL, path, and query parameters.
func buildURL(baseURL string, path string, queryStructs []any) (string, error) {
	fullURL, err := url.JoinPath(baseURL, path)
	if err != nil {
		return "", err
	}
	u, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}
	err = queryParams(u, queryStructs)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// queryParams adds query parameters to the URL from the provided structs.
func queryParams(reqURL *url.URL, queryStructs []any) error {
	urlValues, err := url.ParseQuery(reqURL.RawQuery)
	if err != nil {
		return err
	}

	for _, queryStruct := range queryStructs {
		queryValues, err := goquery.Values(queryStruct)
		if err != nil {
			return err
		}
		for key, values := range queryValues {
			for _, value := range values {
				urlValues.Add(key, value)
			}
		}
	}
	reqURL.RawQuery = urlValues.Encode()
	return nil
}
