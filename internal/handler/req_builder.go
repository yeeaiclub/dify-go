// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package handler

import (
	"errors"
)

// Builder defines the interface for building an API request.
type Builder interface {
	// BaseURL sets the base url for the request
	BaseURL(baseURL string) Builder
	// Path sets the Path relative to the base URL (e.g., "/v1/users").
	Path(path string) Builder
	// PathParm add a path parameter to the request Path.
	PathParm(param string) Builder
	// Token sets the token used in the request headers
	Token(token string) Builder
	// Method sets the http Method
	Method(method string) Builder
	// Body sets the request Body payload
	Body(body any) Builder
	// Query add a Query struct to the request.
	Query(queryStruct any) Builder
	// Headers .
	Headers(headers map[string]string) Builder
	// Build return the Request
	Build() (Request, error)
}

// Request holds the config for a http request.
type Request struct {
	BaseURL   string
	AuthToken string
	Path      string
	Method    string
	Body      any
	Headers   map[string]string
	Query     []any
}

var _ Builder = (*RequestBuilder)(nil)

// RequestBuilder implements the Builder interface for constructing a Request.
type RequestBuilder struct {
	request Request
}

// NewRequestBuilder creates a new instance of RequestBuilder.
func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{request: Request{Query: make([]any, 0)}}
}

// BaseURL  sets the base URL for the request.
func (r *RequestBuilder) BaseURL(baseURL string) Builder {
	r.request.BaseURL = baseURL
	return r
}

// Path sets the request Path.
func (r *RequestBuilder) Path(path string) Builder {
	r.request.Path = path
	return r
}

// Token sets the authentication token.
func (r *RequestBuilder) Token(token string) Builder {
	r.request.AuthToken = token
	return r
}

// Method sets the HTTP Method.
func (r *RequestBuilder) Method(method string) Builder {
	r.request.Method = method
	return r
}

// Body sets the request Body.
func (r *RequestBuilder) Body(body any) Builder {
	r.request.Body = body
	return r
}

// Query add a Query struct to the request.
func (r *RequestBuilder) Query(queryStruct any) Builder {
	if r.request.Query != nil {
		r.request.Query = append(r.request.Query, queryStruct)
	}
	return r
}

// PathParm add a path parameter to the request Path.
func (r *RequestBuilder) PathParm(param string) Builder {
	r.request.Path = r.request.Path + "/" + param
	return r
}

// Headers add header to the request.
func (r *RequestBuilder) Headers(headers map[string]string) Builder {
	r.request.Headers = headers
	return r
}

// Build return a request.
func (r *RequestBuilder) Build() (Request, error) {
	var err error
	if r.request.BaseURL == "" {
		e := errors.New("failed to create the request, should have a BaseURL")
		err = errors.Join(err, e)
	}

	if r.request.Path == "" {
		e := errors.New("failed to create the request, should have a Path")
		err = errors.Join(err, e)
	}

	if r.request.Method == "" {
		e := errors.New("failed to create the request, should have a Method")
		err = errors.Join(err, e)
	}
	return r.request, err
}

// Reset the builder's request.
func (r *RequestBuilder) Reset() {
	r.request = Request{}
}
