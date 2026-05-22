// Copyright The yeeaiclub Authors
// SPDX-License-Identifier: Apache-2.0

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
	Event         string                  `json:"event"`
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

type WorkflowRunLogQuery struct {
	Keyword                   string `url:"keyword"`
	Status                    string `url:"status"`
	Page                      int    `url:"page"`
	Limit                     int    `url:"limit"`
	CreatedAtBefore           string `url:"created_at__before"`
	CreatedAtAfter            string `url:"created_at__after"`
	CreatedByEndUserSessionID string `url:"created_by_end_user_session_id"`
	CreatedByAccount          string `url:"created_by_account"`
}

type WorkflowLogsResponse struct {
	Page    int                        `json:"page"`
	Limit   int                        `json:"limit"`
	Total   int                        `json:"total"`
	HasMore bool                       `json:"has_more"`
	Data    []WorkflowLogsResponseData `json:"data"`
}

type WorkflowLogsResponseData struct {
	ID                string            `json:"id"`
	WorkflowRunDetail WorkflowRunDetail `json:"workflow_run"`
	CreatedFrom       string            `json:"created_from"`
	CreatedByRole     string            `json:"created_by_role"`
	CreatedByAccount  *string           `json:"created_by_account,omitempty"`
	CreatedByEndUser  *CreatedByEndUser `json:"created_by_end_user,omitempty"`
	CreatedAt         int64             `json:"created_at"`
}

type WorkflowRunDetail struct {
	ID              string  `json:"id"`
	Version         string  `json:"version"`
	Status          string  `json:"status"`
	Error           *string `json:"error,omitempty"`
	ElapsedTime     float64 `json:"elapsed_time"`
	TotalTokens     int     `json:"total_tokens"`
	TotalSteps      int     `json:"total_steps"`
	CreatedAt       int64   `json:"created_at"`
	FinishedAt      int64   `json:"finished_at"`
	ExceptionsCount int     `json:"exceptions_count"`
	TriggeredFrom   string  `json:"triggered_from"`
}

type CreatedByEndUser struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	IsAnonymous bool   `json:"is_anonymous"`
	SessionID   string `json:"session_id"`
}
