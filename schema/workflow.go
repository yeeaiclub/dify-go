// Package schema defines the data structures used by the Flowise Go SDK.
// This package contains all the request and response schema for the Flowise API operations.
package schema

import "encoding/json"

// RunWorkflowRequestFile represents a file included in a workflow run request.
type RunWorkflowRequestFile struct {
	Type           string `json:"type"`
	TransferMethod string `json:"transfer_method"`
	URL            string `json:"url"`
	UploadFileID   string `json:"upload_file_id"`
}

// RunWorkflowRequest represents the request body for running a workflow.
type RunWorkflowRequest struct {
	Inputs       json.RawMessage          `json:"inputs"`
	ResponseMode string                   `json:"response_mode"`
	User         string                   `json:"user"`
	Files        []RunWorkflowRequestFile `json:"files"`
}

// RunWorkflowResponse represents the response from running a workflow.
type RunWorkflowResponse struct {
	WorkflowRunID string                  `json:"workflow_run_id"`
	TaskID        string                  `json:"task_id"`
	Data          RunWorkflowResponseData `json:"data"`
}

// RunWorkflowResponseData contains the detailed data from a workflow run response.
type RunWorkflowResponseData struct {
	ID          string         `json:"id"`
	WorkflowID  string         `json:"workflow_id"`
	Status      string         `json:"status"`
	Outputs     map[string]any `json:"outputs"`
	Error       string         `json:"error"`
	ElapsedTime float64        `json:"elapsed_time"`
	TotalToken  int            `json:"total_token"`
	TotalSteps  int            `json:"total_steps"`
	CreatedAt   int            `json:"created_at"`
	FinishedAt  int            `json:"finished_at"`
}

// StreamEvent represents an event in a streaming response.
// T is the type of data contained in the event.
type StreamEvent[T any] struct {
	Err  string `json:"err"`
	Data T      `json:"data"`
	Done bool   `json:"done"`
}
