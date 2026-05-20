package ninerouter

import (
	"context"
	"net/http"
)

func (c *Client) Health(ctx context.Context) (*HealthResponse, error) {
	var out HealthResponse
	err := c.doJSON(ctx, http.MethodGet, "/api/health", nil, &out)
	return &out, err
}

func (c *Client) Shutdown(ctx context.Context) error {
	return c.doJSON(ctx, http.MethodPost, "/api/shutdown", nil, nil)
}
