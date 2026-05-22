package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/yeeaiclub/dify-go/internal/handler"
	"github.com/yeeaiclub/dify-go/schema"
)

// Application represents a client for interacting with the application API endpoints.
type Application struct {
	*BaseClient
}

// NewApplication creates a new Application client instance.
func NewApplication(baseURL, apiKey string) *Application {
	baseClient := &BaseClient{
		client:  handler.NewClient(),
		apiKey:  apiKey,
		baseURL: baseURL,
	}
	return &Application{baseClient}
}

// GetParameters retrieves the application's input form configuration.
func (app *Application) GetParameters(ctx context.Context) (schema.ApplicationParameters, error) {
	r, err := handler.NewRequestBuilder().
		Method(http.MethodGet).
		BaseURL(app.baseURL).
		Token(app.apiKey).
		Path("v1/parameters").
		Build()
	if err != nil {
		return schema.ApplicationParameters{}, err
	}
	resp, err := app.client.Send(ctx, r)
	if err != nil {
		return schema.ApplicationParameters{}, err
	}
	var respData schema.ApplicationParameters
	err = json.Unmarshal(resp.Body, &respData)
	if err != nil {
		return schema.ApplicationParameters{}, err
	}
	return respData, nil
}
