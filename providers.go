package ninerouter

import (
	"context"
	"net/http"
)

type CreateProviderConnectionRequest struct {
	Provider               string
	APIKey                 string
	Name                   string
	DisplayName            string
	Priority               int
	GlobalPriority         int
	DefaultModel           string
	TestStatus             string
	ProviderSpecificData   map[string]any
	ProxyPoolID            string
	ConnectionProxyEnabled bool
	ConnectionProxyURL     string
	ConnectionNoProxy      string
}

type UpdateProviderConnectionRequest map[string]any

func (r CreateProviderConnectionRequest) body() map[string]any {
	body := map[string]any{
		"provider": r.Provider,
	}
	if r.APIKey != "" {
		body["apiKey"] = r.APIKey
	}
	if r.Name != "" {
		body["name"] = r.Name
	}
	if r.DisplayName != "" {
		body["displayName"] = r.DisplayName
	}
	if r.Priority != 0 {
		body["priority"] = r.Priority
	}
	if r.GlobalPriority != 0 {
		body["globalPriority"] = r.GlobalPriority
	}
	if r.DefaultModel != "" {
		body["defaultModel"] = r.DefaultModel
	}
	if r.TestStatus != "" {
		body["testStatus"] = r.TestStatus
	}
	if r.ProviderSpecificData != nil {
		body["providerSpecificData"] = r.ProviderSpecificData
	}
	if r.ProxyPoolID != "" {
		body["proxyPoolId"] = r.ProxyPoolID
	}
	if r.ConnectionProxyEnabled {
		body["connectionProxyEnabled"] = true
		body["connectionProxyUrl"] = r.ConnectionProxyURL
	}
	if r.ConnectionNoProxy != "" {
		body["connectionNoProxy"] = r.ConnectionNoProxy
	}
	return body
}

func (c *Client) ListProviderConnections(ctx context.Context) ([]ProviderConnection, error) {
	var out struct {
		Connections []ProviderConnection `json:"connections"`
	}
	err := c.doJSON(ctx, http.MethodGet, "/api/providers", nil, &out)
	return out.Connections, err
}

func (c *Client) CreateProviderConnection(ctx context.Context, req CreateProviderConnectionRequest) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodPost, "/api/providers", req.body(), &out)
	return out, err
}

func (c *Client) GetProviderConnection(ctx context.Context, id string) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodGet, "/api/providers/"+id, nil, &out)
	return out, err
}

func (c *Client) UpdateProviderConnection(ctx context.Context, id string, req UpdateProviderConnectionRequest) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodPut, "/api/providers/"+id, map[string]any(req), &out)
	return out, err
}

func (c *Client) DeleteProviderConnection(ctx context.Context, id string) error {
	return c.doJSON(ctx, http.MethodDelete, "/api/providers/"+id, nil, nil)
}

func (c *Client) ListProviderNodes(ctx context.Context) ([]ProviderNode, error) {
	var out struct {
		Nodes []ProviderNode `json:"nodes"`
	}
	err := c.doJSON(ctx, http.MethodGet, "/api/provider-nodes", nil, &out)
	return out.Nodes, err
}

func (c *Client) CreateProviderNode(ctx context.Context, body map[string]any) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodPost, "/api/provider-nodes", body, &out)
	return out, err
}

func (c *Client) UpdateProviderNode(ctx context.Context, id string, body map[string]any) (map[string]any, error) {
	var out map[string]any
	err := c.doJSON(ctx, http.MethodPut, "/api/provider-nodes/"+id, body, &out)
	return out, err
}

func (c *Client) DeleteProviderNode(ctx context.Context, id string) error {
	return c.doJSON(ctx, http.MethodDelete, "/api/provider-nodes/"+id, nil, nil)
}
