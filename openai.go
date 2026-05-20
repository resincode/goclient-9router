package ninerouter

import (
	"context"
	"net/http"
)

func (c *Client) ChatCompletions(ctx context.Context, request any, response any) error {
	return c.doJSON(ctx, http.MethodPost, "/v1/chat/completions", request, response)
}

func (c *Client) Responses(ctx context.Context, request any, response any) error {
	return c.doJSON(ctx, http.MethodPost, "/v1/responses", request, response)
}

func (c *Client) Messages(ctx context.Context, request any, response any) error {
	return c.doJSON(ctx, http.MethodPost, "/v1/messages", request, response)
}
