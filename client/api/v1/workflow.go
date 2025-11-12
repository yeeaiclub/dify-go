// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/yeeaiclub/dify-go/internal/logger"

	"github.com/yeeaiclub/dify-go/internal/handler"
	"github.com/yeeaiclub/dify-go/schema"
)

const (
	defaultChannelBufferSize = 10
)

// Workflow represents a client for interacting with the workflow API endpoints.
type Workflow struct {
	client  *handler.Client // HTTP client for making API requests
	apiKey  string          // API key for authentication
	baseURL string          // Base URL of the API server
}

// NewWorkflow creates a new Workflow client instance.
func NewWorkflow(baseURL, apiKey string) *Workflow {
	return &Workflow{
		client:  handler.NewClient(),
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// RunStream executes a workflow in streaming mode. Cannot execute if there is no published workflow.
func (w *Workflow) RunStream(
	ctx context.Context,
	req schema.RunWorkflowRequest,
) (chan schema.RunWorkflowResponse, error) {
	if req.ResponseMode != "stream" {
		return nil, errors.New("response mode must be stream")
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
	ch := make(chan schema.RunWorkflowResponse, defaultChannelBufferSize)
	eventCh, err := w.client.SendStream(ctx, r)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				log.Debugf("context canceled: %v", ctx.Err())
				return
			case value, ok := <-eventCh:
				if !ok {
					return
				}
				var resp schema.RunWorkflowResponse
				err := json.Unmarshal(value.Body, &resp)
				if err != nil {
					log.Errorf("failed to unmarshal json: %v", err)
					continue
				}
				select {
				case ch <- resp:
					// Message sent successfully
				case <-ctx.Done():
					// Context canceled while sending
					return
				}
			}
		}
	}()
	return ch, nil
}

// Run executes a workflow in blocking mode, Cannot execute if there is no published workflow.
func (w *Workflow) Run(
	ctx context.Context,
	req schema.RunWorkflowRequest,
) (schema.RunWorkflowResponse, error) {
	if req.ResponseMode != "blocking" {
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
