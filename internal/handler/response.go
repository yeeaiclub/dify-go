// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package handler

import "net/http"

// Response holds the response data for an API request.
type Response struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}
