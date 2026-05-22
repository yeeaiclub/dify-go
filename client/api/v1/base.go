package v1

import "github.com/yeeaiclub/dify-go/internal/handler"

// BaseClient the base client of dify
type BaseClient struct {
	client  *handler.Client // HTTP client for making API requests
	apiKey  string          // API key for authentication
	baseURL string          // Base URL of the API server
}
