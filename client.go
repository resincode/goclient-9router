package ninerouter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const DefaultBaseURL = "http://127.0.0.1:20128"

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	apiKey     string
	cliToken   string
	cookies    []*http.Cookie
	userAgent  string
}

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func WithCLIToken(cliToken string) Option {
	return func(c *Client) {
		c.cliToken = cliToken
	}
}

func WithAutoCLIToken() Option {
	return func(c *Client) {
		c.cliToken = deriveCLIToken()
	}
}

func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

func New(baseURL string, opts ...Option) (*Client, error) {
	if strings.TrimSpace(baseURL) == "" {
		baseURL = DefaultBaseURL
	}
	u, err := url.Parse(strings.TrimRight(baseURL, "/"))
	if err != nil {
		return nil, err
	}
	c := &Client{
		baseURL: u,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		userAgent: "9router-client-go",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) BaseURL() string {
	return c.baseURL.String()
}

func (c *Client) newRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	u := c.baseURL.ResolveReference(&url.URL{Path: joinPath(c.baseURL.Path, path)})
	if strings.Contains(path, "?") {
		parts := strings.SplitN(path, "?", 2)
		u.Path = joinPath(c.baseURL.Path, parts[0])
		u.RawQuery = parts[1]
	}

	var r io.Reader
	if body != nil {
		buf := bytes.NewBuffer(nil)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
		r = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	if c.cliToken != "" {
		req.Header.Set("x-9r-cli-token", c.cliToken)
	}
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}
	return req, nil
}

func (c *Client) doJSON(ctx context.Context, method, path string, body any, out any) error {
	req, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	c.storeCookies(resp.Cookies())

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return parseError(resp.StatusCode, data)
	}
	if out == nil || len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, out)
}

func (c *Client) DoRaw(ctx context.Context, method, path string, body any) (*http.Response, error) {
	req, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func (c *Client) storeCookies(cookies []*http.Cookie) {
	if len(cookies) == 0 {
		return
	}
	byName := make(map[string]*http.Cookie, len(c.cookies)+len(cookies))
	for _, cookie := range c.cookies {
		byName[cookie.Name] = cookie
	}
	for _, cookie := range cookies {
		byName[cookie.Name] = cookie
	}
	c.cookies = c.cookies[:0]
	for _, cookie := range byName {
		c.cookies = append(c.cookies, cookie)
	}
}

func joinPath(basePath, path string) string {
	basePath = strings.TrimRight(basePath, "/")
	path = "/" + strings.TrimLeft(path, "/")
	if basePath == "" {
		return path
	}
	return basePath + path
}

func parseError(status int, data []byte) error {
	var payload struct {
		Error any
	}
	if err := json.Unmarshal(data, &payload); err == nil && payload.Error != nil {
		return fmt.Errorf("9router: status %d: %v", status, payload.Error)
	}
	return fmt.Errorf("9router: status %d: %s", status, strings.TrimSpace(string(data)))
}
