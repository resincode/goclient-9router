# goclient-9router

Small Go client for controlling a local 9router instance from another app.

## Scope

- Start a local 9router process with no browser, tray, host, and port options.
- Check /api/health, auth status, password login, and logout.
- Manage 9router API keys through /api/keys.
- Manage provider nodes through /api/provider-nodes.
- Manage provider connections through /api/providers.
- Manage model aliases, custom models, and disabled models.
- Apply or remove Codex CLI settings through /api/cli-tools/codex-settings.
- Call OpenAI-compatible proxy endpoints under /v1.

The internal dashboard APIs are based on 9router 0.4.x route behavior. Sensitive fields returned by 9router are not printed by this library, but callers should still treat API key responses as secrets.

## Example

~~~go
package main

import (
	"context"
	"fmt"

	ninerouter "github.com/resincode/goclient-9router"
)

func main() {
	ctx := context.Background()
	client, err := ninerouter.New(ninerouter.DefaultBaseURL)
	if err != nil {
		panic(err)
	}

	health, err := client.Health(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("healthy:", health.OK)

	key, err := client.CreateKey(ctx, "web-ui")
	if err != nil {
		panic(err)
	}
	fmt.Println("created key id:", key["id"])
}
~~~

## Notes

- If dashboard login is required, call Login(ctx, password) first. The client stores returned cookies in memory for later calls.
- For local CLI-style dashboard API access without a password, construct the client with WithAutoCLIToken().
- For /v1 calls, construct the client with WithAPIKey(key).
- Start only starts the process. Use Health with retries in the caller to wait until the service is ready.
- Go is managed with mise on this host.
