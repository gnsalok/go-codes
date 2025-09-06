package myapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"httpxclient/internal/httpx"
)

// API is the small surface your consumers depend on (Mockery will generate mocks for this).
// Keep this stable and minimal.
type API interface {
	GetUser(ctx context.Context, id string) (User, error)
	CreateUser(ctx context.Context, in CreateUserRequest) (User, error)
}

type Client struct {
	baseURL string
	http    httpx.Doer
	apiKey  string
}

type Option func(*Client)

func WithHTTPClient(h httpx.Doer) Option { return func(c *Client) { c.http = h } }
func WithAPIKey(key string) Option       { return func(c *Client) { c.apiKey = key } }
func WithBaseURL(url string) Option      { return func(c *Client) { c.baseURL = url } }

// Return a concrete type. Consumers can still depend on API interface.
func NewClient(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL: baseURL,
		http:    &http.Client{Timeout: 10 * time.Second},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (c *Client) GetUser(ctx context.Context, id string) (User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/users/"+id, nil)
	if err != nil {
		return User{}, err
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return User{}, fmt.Errorf("get user: %s: %s", resp.Status, b)
	}
	var out User
	return out, json.NewDecoder(resp.Body).Decode(&out)
}

func (c *Client) CreateUser(ctx context.Context, in CreateUserRequest) (User, error) {
	body, r, err := marshalJSON(ctx, http.MethodPost, c.baseURL+"/users", in)
	if err != nil {
		return User{}, err
	}
	if c.apiKey != "" {
		r.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return User{}, fmt.Errorf("create user: %s: %s", resp.Status, b)
	}
	var out User
	return out, json.NewDecoder(resp.Body).Decode(&out)
}

func marshalJSON(ctx context.Context, method, url string, v any) (io.Reader, *http.Request, error) {
	pr, pw := io.Pipe()
	go func() {
		enc := json.NewEncoder(pw)
		pw.CloseWithError(enc.Encode(v))
	}()
	req, err := http.NewRequestWithContext(ctx, method, url, pr)
	if err != nil {
		_ = pr
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return pr, req, nil
}
