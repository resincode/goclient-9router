package ninerouter

import (
	"context"
	"net/http"
)

func (c *Client) CodexSettings(ctx context.Context) (CodexSettings, error) {
	var out CodexSettings
	err := c.doJSON(ctx, http.MethodGet, "/api/cli-tools/codex-settings", nil, &out)
	return out, err
}

func (c *Client) ApplyCodexSettings(ctx context.Context, baseURL, apiKey, model, subagentModel string) error {
	body := map[string]any{
		"baseUrl": baseURL,
		"apiKey":  apiKey,
		"model":   model,
	}
	if subagentModel != "" {
		body["subagentModel"] = subagentModel
	}
	return c.doJSON(ctx, http.MethodPost, "/api/cli-tools/codex-settings", body, nil)
}

func (c *Client) RemoveCodexSettings(ctx context.Context) error {
	return c.doJSON(ctx, http.MethodDelete, "/api/cli-tools/codex-settings", nil, nil)
}
