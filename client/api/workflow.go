package api

import (
	"context"
	"github.com/yeeaiclub/dify-go/internal/handler"
	"github.com/yeeaiclub/dify-go/types"
)

type Workflow struct {
	client *handler.Client
}

func NewWorkflow(apiKey string) *Workflow {
	return &Workflow{}
}

func (w *Workflow) RunStream(
	ctx context.Context,
	req types.RunWorkflowRequest,
) (chan types.RunWorkflowResponse, error) {
	return nil, nil
}
