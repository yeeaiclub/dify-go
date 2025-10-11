// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T) {
	t.Run("build POST request with body", func(t *testing.T) {
		type Req struct {
			Name string `json:"name"`
		}
		req, err := NewRequestBuilder().
			BaseURL("example.com").
			Path("users").
			Method(http.MethodPost).
			Token("token").
			Body(Req{Name: "wyz"}).
			Build()
		require.NoError(t, err)
		assert.Equal(t, req.BaseURL, "example.com")
		assert.Equal(t, req.Path, "users")
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.AuthToken, "token")
		assert.Equal(t, req.Body, Req{Name: "wyz"})
	})

	t.Run("build GET request with query parameters", func(t *testing.T) {
		type Query struct {
			Name string `url:"name"`
		}

		req, err := NewRequestBuilder().
			BaseURL("example.com").
			Path("users").
			Method(http.MethodGet).
			Token("token").
			Query(Query{Name: "wyz"}).
			Build()

		require.NoError(t, err)
		assert.Equal(t, req.BaseURL, "example.com")
		assert.Equal(t, req.Path, "users")
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.AuthToken, "token")
		assert.ElementsMatch(t, req.Query, []any{Query{Name: "wyz"}})
	})

	t.Run("build request with path parameters", func(t *testing.T) {
		req, err := NewRequestBuilder().
			BaseURL("example.com").
			Path("users").
			PathParm("123").
			Method(http.MethodGet).
			Token("token").
			Build()

		require.NoError(t, err)
		assert.Equal(t, req.BaseURL, "example.com")
		assert.Equal(t, req.Path, "users/123")
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.AuthToken, "token")
	})
}
