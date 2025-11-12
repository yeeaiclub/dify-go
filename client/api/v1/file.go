// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/yeeaiclub/dify-go/internal/handler"
	"github.com/yeeaiclub/dify-go/schema"
)

type FileService struct {
	client  *handler.Client // HTTP client for making API requests
	apiKey  string          // API key for authentication
	baseURL string          // Base URL of the API server
}

func NewFileService(baseURL, apiKey string) *FileService {
	return &FileService{
		client:  handler.NewClient(),
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// Upload upload file to dify.
func (f *FileService) Upload(ctx context.Context, req schema.UploadFileRequest) (schema.UploadFileResponse, error) {
	r, err := handler.NewRequestBuilder().
		BaseURL(f.baseURL).
		Token(f.apiKey).
		Path("v1/files/upload").
		Method(http.MethodPost).
		Body(req).
		Build()
	if err != nil {
		return schema.UploadFileResponse{}, err
	}
	var respData schema.UploadFileResponse
	resp, err := f.client.Send(ctx, r)
	if err != nil {
		return schema.UploadFileResponse{}, err
	}
	err = json.Unmarshal(resp.Body, &respData)
	if err != nil {
		return schema.UploadFileResponse{}, err
	}
	return respData, nil
}
