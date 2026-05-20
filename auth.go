package ninerouter

import (
	"context"
	"net/http"
)

func (c *Client) AuthStatus(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodGet, "/api/auth/status", nil, &out)
	return out, err
}

func (c *Client) Login(ctx context.Context, password string) error {
	var out map[string]any
	return c.doJSON(ctx, http.MethodPost, "/api/auth/login", map[string]string{
		"password": password,
	}, &out)
}

func (c *Client) Logout(ctx context.Context) error {
	return c.doJSON(ctx, http.MethodPost, "/api/auth/logout", nil, nil)
}
