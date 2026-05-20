package ninerouter

import (
	"context"
	"net/http"
)

func (c *Client) ListKeys(ctx context.Context) ([]APIKey, error) {
	var out struct {
		Keys []APIKey `json:"keys"`
	}
	err := c.doJSON(ctx, http.MethodGet, "/api/keys", nil, &out)
	return out.Keys, err
}

func (c *Client) CreateKey(ctx context.Context, name string) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodPost, "/api/keys", map[string]string{"name": name}, &out)
	return out, err
}

func (c *Client) GetKey(ctx context.Context, id string) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodGet, "/api/keys/"+id, nil, &out)
	return out, err
}

func (c *Client) SetKeyActive(ctx context.Context, id string, active bool) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodPut, "/api/keys/"+id, map[string]bool{"isActive": active}, &out)
	return out, err
}

func (c *Client) DeleteKey(ctx context.Context, id string) error {
	return c.doJSON(ctx, http.MethodDelete, "/api/keys/"+id, nil, nil)
}
