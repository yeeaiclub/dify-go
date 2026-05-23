package v1

import (
	"time"

	"github.com/yeeaiclub/dify-go/internal/handler"
)

const (
	// defaultMaxIdleConns is the default maximum number of idle connections.
	defaultMaxIdleConns = 100
	// defaultMaxIdleConnsPerHost is the default maximum number of idle connections per host.
	defaultMaxIdleConnsPerHost = 10
	// defaultIdleConnTimeoutSec is the default idle connection timeout in seconds.
	defaultIdleConnTimeoutSec = 90
)

// PoolConfig defines connection pool configuration for the HTTP client.
type PoolConfig struct {
	// MaxIdleConns is the maximum number of idle connections across all hosts.
	// Default: 100
	MaxIdleConns int
	// MaxIdleConnsPerHost is the maximum number of idle connections per host.
	// Default: 10
	MaxIdleConnsPerHost int
	// MaxConnsPerHost is the maximum number of connections per host (0 = unlimited).
	// Default: 0 (unlimited)
	MaxConnsPerHost int
	// IdleConnTimeout is the maximum amount of time an idle connection will remain idle.
	// Default: 90 seconds
	IdleConnTimeout time.Duration
}

// BaseClient the base client of dify
type BaseClient struct {
	client  *handler.Client // HTTP client for making API requests
	apiKey  string          // API key for authentication
	baseURL string          // Base URL of the API server
}

// NewBaseClient creates a new BaseClient with the given configuration.
func NewBaseClient(baseURL, apiKey string, poolConfig *PoolConfig) *BaseClient {
	var opts []handler.ClientOption

	if poolConfig != nil {
		if poolConfig.MaxIdleConns > 0 {
			opts = append(opts, handler.WithMaxIdleConns(poolConfig.MaxIdleConns))
		}
		if poolConfig.MaxIdleConnsPerHost > 0 {
			opts = append(opts, handler.WithMaxIdleConnsPerHost(poolConfig.MaxIdleConnsPerHost))
		}
		if poolConfig.MaxConnsPerHost > 0 {
			opts = append(opts, handler.WithMaxConnsPerHost(poolConfig.MaxConnsPerHost))
		}
		if poolConfig.IdleConnTimeout > 0 {
			opts = append(opts, handler.WithIdleConnTimeout(poolConfig.IdleConnTimeout))
		}
	}

	return &BaseClient{
		client:  handler.NewClient(opts...),
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// DefaultPoolConfig returns a PoolConfig with sensible defaults.
func DefaultPoolConfig() *PoolConfig {
	return &PoolConfig{
		MaxIdleConns:        defaultMaxIdleConns,
		MaxIdleConnsPerHost: defaultMaxIdleConnsPerHost,
		IdleConnTimeout:     defaultIdleConnTimeoutSec * time.Second,
	}
}
