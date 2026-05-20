package ninerouter

type HealthResponse struct {
	OK bool `json:"ok"`
}

type APIKey struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Key       string `json:"key,omitempty"`
	MachineID string `json:"machineId,omitempty"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type ProviderConnection map[string]any

type ProviderNode map[string]any

type CustomModel map[string]any

type CodexSettings map[string]any
