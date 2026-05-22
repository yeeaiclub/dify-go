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
	*BaseClient
}

// NewWorkflowService creates a new Workflow client instance.
func NewWorkflowService(baseURL, apiKey string) *WorkflowService {
	baseClient := &BaseClient{
		client:  handler.NewClient(),
		apiKey:  apiKey,
		baseURL: baseURL,
	}
	return &WorkflowService{baseClient}
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

// GetLogs retrieves workflow execution logs with optional filtering.
// Supports pagination and filtering by status, keyword, and time range.
func (w *WorkflowService) GetLogs(
	ctx context.Context,
	query schema.WorkflowRunLogQuery,
) (schema.WorkflowLogsResponse, error) {
	r, err := handler.NewRequestBuilder().
		BaseURL(w.baseURL).
		Token(w.apiKey).
		Path("v1/workflows/logs").
		Method(http.MethodGet).
		Query(query).
		Build()
	if err != nil {
		return schema.WorkflowLogsResponse{}, err
	}

	resp, err := w.client.Send(ctx, r)
	if err != nil {
		return schema.WorkflowLogsResponse{}, err
	}
	var respData schema.WorkflowLogsResponse
	err = json.Unmarshal(resp.Body, &respData)
	if err != nil {
		return schema.WorkflowLogsResponse{}, err
	}
	return respData, nil
}
