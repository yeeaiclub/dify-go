// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yeeaiclub/dify-go/internal/handler"
	"github.com/yeeaiclub/dify-go/schema"
)

const (
	defaultChannelBufferSize = 10
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
	respCh chan schema.StreamEvent[schema.RunWorkflowResponse],
) error {
	if req.ResponseMode != StreamMode {
		return nil
	}

	r, err := handler.NewRequestBuilder().
		BaseURL(w.baseURL).
		Token(w.apiKey).
		Path("v1/workflows/run").
		Method(http.MethodPost).
		Body(req).
		Build()
	if err != nil {
		return err
	}

	go func() {
		evh := make(chan *handler.Event, defaultChannelBufferSize)
		defer close(respCh) // Ensure the response channel is closed when done

		// Send the streaming request
		err := w.client.SendStream(ctx, r, evh)
		if err != nil {
			respCh <- schema.StreamEvent[schema.RunWorkflowResponse]{
				Err: err.Error(), // Send the actual error, not ctx.Err()
			}
			return
		}

		// Process the stream events
		for {
			select {
			case <-ctx.Done():
				respCh <- schema.StreamEvent[schema.RunWorkflowResponse]{
					Err: ctx.Err().Error(),
				}
				return
			case ev, ok := <-evh:
				if !ok {
					return
				}
				if ev.Done {
					respCh <- schema.StreamEvent[schema.RunWorkflowResponse]{
						Done: true,
					}
					return
				}
				var data schema.RunWorkflowResponse
				err := json.NewDecoder(ev.Data).Decode(&data)
				if err != nil {
					respCh <- schema.StreamEvent[schema.RunWorkflowResponse]{
						Err: err.Error(),
					}
					return
				}
				respCh <- schema.StreamEvent[schema.RunWorkflowResponse]{
					Type: ev.Type,
					Data: data,
				}
			}
		}
	}()
	return nil
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
