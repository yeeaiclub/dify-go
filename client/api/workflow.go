package api

import (
	"context"
	"encoding/json"
	log "github.com/yeeaiclub/dify-go/internal/logger"
	"net/http"

	"github.com/yeeaiclub/dify-go/internal/handler"
	"github.com/yeeaiclub/dify-go/types"
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

// RunStream executes a workflow. Cannot execute if there is no published workflow.
func (w *Workflow) RunStream(
	ctx context.Context,
	req types.RunWorkflowRequest,
) (chan types.RunWorkflowResponse, error) {

	r, err := handler.NewRequestBuilder().
		BaseURL(w.baseURL).
		Token(w.apiKey).
		Path("workflows/run").
		Method(http.MethodPost).
		Body(req).
		Build()

	if err != nil {
		return nil, err
	}
	ch := make(chan types.RunWorkflowResponse, 10)
	eventCh, err := w.client.SendStream(ctx, r)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				log.Debugf("context cancelled: %v", ctx.Err())
				return
			case value, ok := <-eventCh:
				if !ok {
					return
				}
				var resp types.RunWorkflowResponse
				err := json.Unmarshal(value.Body, &resp)
				if err != nil {
					log.Errorf("failed to unmarshal json: %v", err)
					continue
				}
				select {
				case ch <- resp:
					// Message sent successfully
				case <-ctx.Done():
					// Context cancelled while sending
					return
				}
			}
		}
	}()
	return ch, nil
}
