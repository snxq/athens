package config

// HTTPConfig specifies the properties required to ues http file server as the storage.
type HTTPConfig struct {
	BaseURL string `validate:"required" envconfig:"ATHENS_HTTP_STORAGE_BASE"`
}
