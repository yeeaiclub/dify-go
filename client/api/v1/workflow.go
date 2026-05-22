// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"encoding/json"
	"errors"
	"iter"
	"net/http"

	"github.com/yeeaiclub/dify-go/internal/handler"
	"github.com/yeeaiclub/dify-go/schema"
)

const (
	// StreamMode represents the streaming response mode for workflow execution
	StreamMode = "streaming"
	// BlockingMode represents the blocking response mode for workflow execution
	BlockingMode = "blocking"
)

// WorkflowService represents a client for interacting with the workflow API endpoints.
type WorkflowService struct {
	client  *handler.Client // HTTP client for making API requests
	apiKey  string          // API key for authentication
	baseURL string          // Base URL of the API server
}

// NewWorkflowService creates a new Workflow client instance.
func NewWorkflowService(baseURL, apiKey string) *WorkflowService {
	return &WorkflowService{
		client:  handler.NewClient(),
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// RunStream executes a workflow in streaming mode. Cannot execute if there is no published workflow.
func (w *WorkflowService) RunStream(
	ctx context.Context,
	req schema.RunWorkflowRequest,
) (iter.Seq2[[]byte, error], error) {
	if req.ResponseMode != StreamMode {
		return nil, errors.New("invalid response mode")
	}

	r, err := handler.NewRequestBuilder().
		BaseURL(w.baseURL).
		Token(w.apiKey).
		Path("v1/workflows/run").
		Method(http.MethodPost).
		Body(req).
		Build()
	if err != nil {
		return nil, err
	}

	return w.client.SendStream(ctx, r)
}

// Run executes a workflow in blocking mode, Cannot execute if there is no published workflow.
func (w *WorkflowService) Run(
	ctx context.Context,
	req schema.RunWorkflowRequest,
) (schema.RunWorkflowResponse, error) {
	if req.ResponseMode != BlockingMode {
		return schema.RunWorkflowResponse{}, errors.New("response mode must be blocking")
	}

	r, err := handler.NewRequestBuilder().
		BaseURL(w.baseURL).
		Token(w.apiKey).
		Path("v1/workflows/run").
		Method(http.MethodPost).
		Body(req).
		Build()
	if err != nil {
		return schema.RunWorkflowResponse{}, err
	}
	resp, err := w.client.Send(ctx, r)
	if err != nil {
		return schema.RunWorkflowResponse{}, err
	}
	var respData schema.RunWorkflowResponse
	err = json.Unmarshal(resp.Body, &respData)
	if err != nil {
		return schema.RunWorkflowResponse{}, err
	}
	return respData, nil
}
