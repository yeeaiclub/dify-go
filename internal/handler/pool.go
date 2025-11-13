// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package handler

import "sync"

// GetRequestBuilder gets a RequestBuilder from the provided pool.
func GetRequestBuilder(pool *sync.Pool) *RequestBuilder {
	return pool.Get().(*RequestBuilder)
}

// PutRequestBuilder puts a RequestBuilder back into the pool.
func PutRequestBuilder(pool *sync.Pool, v *RequestBuilder) {
	v.request = Request{}
	pool.Put(v)
}
