// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package handler

import (
	"io"
	"net/http"
)

// Response holds the response data for an API request.
type Response struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

// Event is an http server-sent event
type Event struct {
	Type string
	Data io.Reader
	Done bool
}
