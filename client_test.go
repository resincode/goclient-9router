package ninerouter

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/health" {
			t.Fatalf("unexpected path %s", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}))
	t.Cleanup(server.Close)

	client, err := New(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	health, err := client.Health(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if !health.OK {
		t.Fatal("expected healthy response")
	}
}

func TestCreateKeyPayload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/keys" {
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body["name"] != "web-ui" {
			t.Fatalf("unexpected name %q", body["name"])
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"id": "key-id", "name": "web-ui"})
	}))
	t.Cleanup(server.Close)

	client, err := New(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	key, err := client.CreateKey(context.Background(), "web-ui")
	if err != nil {
		t.Fatal(err)
	}
	if key["id"] != "key-id" {
		t.Fatalf("unexpected key id %v", key["id"])
	}
}

func TestBearerToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer test-key" {
			t.Fatalf("unexpected auth header %q", got)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"id": "resp"})
	}))
	t.Cleanup(server.Close)

	client, err := New(server.URL, WithAPIKey("test-key"))
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]any
	if err := client.Responses(context.Background(), map[string]string{"model": "x"}, &out); err != nil {
		t.Fatal(err)
	}
}

func TestCLITokenHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("x-9r-cli-token"); got != "cli-token" {
			t.Fatalf("unexpected cli token header %q", got)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"keys": []any{}})
	}))
	t.Cleanup(server.Close)

	client, err := New(server.URL, WithCLIToken("cli-token"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.ListKeys(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestDeriveCLITokenMatches9RouterCLI(t *testing.T) {
	machineID := "0123456789abcdef0123456789abcdef"
	machineHash := sha256Hex(machineID)
	got := sha256Hex(machineHash + cliTokenSalt)[:16]
	if got != "5ef8f7e16c5bede3" {
		t.Fatalf("unexpected cli token %q", got)
	}
}
