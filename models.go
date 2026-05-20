package ninerouter

import (
	"context"
	"net/http"
	"net/url"
)

func (c *Client) ListModelAliases(ctx context.Context) (map[string]string, error) {
	var out struct {
		Aliases map[string]string `json:"aliases"`
	}
	err := c.doJSON(ctx, http.MethodGet, "/api/models/alias", nil, &out)
	return out.Aliases, err
}

func (c *Client) SetModelAlias(ctx context.Context, alias, model string) error {
	return c.doJSON(ctx, http.MethodPut, "/api/models/alias", map[string]string{
		"alias": alias,
		"model": model,
	}, nil)
}

func (c *Client) DeleteModelAlias(ctx context.Context, alias string) error {
	return c.doJSON(ctx, http.MethodDelete, "/api/models/alias?alias="+url.QueryEscape(alias), nil, nil)
}

func (c *Client) ListCustomModels(ctx context.Context) ([]CustomModel, error) {
	var out struct {
		Models []CustomModel `json:"models"`
	}
	err := c.doJSON(ctx, http.MethodGet, "/api/models/custom", nil, &out)
	return out.Models, err
}

func (c *Client) AddCustomModel(ctx context.Context, providerAlias, id, modelType, name string) (map[string]any, error) {
	body := map[string]any{
		"providerAlias": providerAlias,
		"id":            id,
	}
	if modelType != "" {
		body["type"] = modelType
	}
	if name != "" {
		body["name"] = name
	}
	var out map[string]any
	err := c.doJSON(ctx, http.MethodPost, "/api/models/custom", body, &out)
	return out, err
}

func (c *Client) DeleteCustomModel(ctx context.Context, providerAlias, id, modelType string) error {
	q := url.Values{}
	q.Set("providerAlias", providerAlias)
	q.Set("id", id)
	if modelType != "" {
		q.Set("type", modelType)
	}
	return c.doJSON(ctx, http.MethodDelete, "/api/models/custom?"+q.Encode(), nil, nil)
}

func (c *Client) DisabledModels(ctx context.Context, providerAlias string) (map[string]any, error) {
	path := "/api/models/disabled"
	if providerAlias != "" {
		path += "?providerAlias=" + url.QueryEscape(providerAlias)
	}
	var out map[string]any
	err := c.doJSON(ctx, http.MethodGet, path, nil, &out)
	return out, err
}

func (c *Client) DisableModels(ctx context.Context, providerAlias string, ids []string) error {
	return c.doJSON(ctx, http.MethodPost, "/api/models/disabled", map[string]any{
		"providerAlias": providerAlias,
		"ids":           ids,
	}, nil)
}

func (c *Client) EnableModels(ctx context.Context, providerAlias string, ids ...string) error {
	q := url.Values{}
	q.Set("providerAlias", providerAlias)
	if len(ids) > 0 {
		q.Set("id", ids[0])
	}
	return c.doJSON(ctx, http.MethodDelete, "/api/models/disabled?"+q.Encode(), nil, nil)
}
