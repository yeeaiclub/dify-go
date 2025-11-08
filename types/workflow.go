package types

import "encoding/json"

type RunWorkflowRequestFile struct {
	Type           string `json:"type"`
	TransferMethod string `json:"transfer_method"`
	URL            string `json:"url"`
	UploadFileID   string `json:"upload_file_id"`
}

type RunWorkflowRequest struct {
	Inputs       json.RawMessage          `json:"inputs"`
	ResponseMode string                   `json:"response_mode"`
	User         string                   `json:"user"`
	Files        []RunWorkflowRequestFile `json:"files"`
}

type RunWorkflowResponse struct {
	WorkflowRunID string                  `json:"workflow_run_id"`
	TaskID        string                  `json:"task_id"`
	Data          RunWorkflowResponseData `json:"data"`
}

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

type StreamEvent[T any] struct {
	Err  string `json:"err"`
	Data T      `json:"data"`
	Done bool   `json:"done"`
}
