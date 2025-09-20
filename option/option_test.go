package option_test

import (
	"net/http"
	"testing"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

func TestWithAPIKey(t *testing.T) {
	key := "test-api-key"
	opt := option.WithAPIKey(key)
	if opt == nil {
		t.Fatal("WithAPIKey() returned nil")
	}

	// Create a request to apply the option to
	req, _ := http.NewRequest("GET", "https://api.test.com", nil)
	cfg := &requestconfig.RequestConfig{
		Request: req,
	}

	// Apply option
	err := opt.Apply(cfg)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
}

func TestWithBaseURL(t *testing.T) {
	url := "https://custom.api.com"
	opt := option.WithBaseURL(url)
	if opt == nil {
		t.Fatal("WithBaseURL() returned nil")
	}

	cfg := &requestconfig.RequestConfig{}
	err := opt.Apply(cfg)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}

	if cfg.BaseURL == nil || cfg.BaseURL.String() != url+"/" {
		t.Errorf("BaseURL = %v, want %v", cfg.BaseURL, url)
	}
}

func TestWithEnvironmentProduction(t *testing.T) {
	opt := option.WithEnvironmentProduction()
	if opt == nil {
		t.Fatal("WithEnvironmentProduction() returned nil")
	}

	cfg := &requestconfig.RequestConfig{}
	err := opt.Apply(cfg)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}

	expectedURL := "https://api.hanzo.ai/"
	if cfg.BaseURL == nil || cfg.BaseURL.String() != expectedURL {
		t.Errorf("BaseURL = %v, want %v", cfg.BaseURL, expectedURL)
	}
}

func TestWithHeader(t *testing.T) {
	opt := option.WithHeader("X-Custom-Header", "value")
	if opt == nil {
		t.Fatal("WithHeader() returned nil")
	}

	req, _ := http.NewRequest("GET", "https://api.test.com", nil)
	cfg := &requestconfig.RequestConfig{
		Request: req,
	}

	err := opt.Apply(cfg)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
}

func TestWithHTTPClient(t *testing.T) {
	client := &http.Client{}
	opt := option.WithHTTPClient(client)
	if opt == nil {
		t.Fatal("WithHTTPClient() returned nil")
	}

	cfg := &requestconfig.RequestConfig{}
	err := opt.Apply(cfg)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}

	if cfg.HTTPClient != client {
		t.Error("HTTPClient was not set correctly")
	}
}