package hanzoai

import (
	"testing"

	"github.com/hanzoai/go-sdk/option"
)

func TestNewModelService(t *testing.T) {
	service := NewModelService()
	if service == nil {
		t.Fatal("NewModelService() returned nil")
	}
	if service.Options == nil {
		t.Error("ModelService.Options is nil")
	}
	if service.Info == nil {
		t.Error("ModelService.Info is nil")
	}
	if service.Update == nil {
		t.Error("ModelService.Update is nil")
	}
}

func TestNewModelServiceWithOptions(t *testing.T) {
	opts := []option.RequestOption{
		option.WithAPIKey("test-key"),
	}
	service := NewModelService(opts...)
	if service == nil {
		t.Fatal("NewModelService() returned nil")
	}
	if len(service.Options) != len(opts) {
		t.Errorf("ModelService.Options length = %d, want %d", len(service.Options), len(opts))
	}
}

func TestNewDeleteService(t *testing.T) {
	service := NewDeleteService()
	if service == nil {
		t.Fatal("NewDeleteService() returned nil")
	}
	if service.Options == nil {
		t.Error("DeleteService.Options is nil")
	}
}

func TestNewModelInfoService(t *testing.T) {
	service := NewModelInfoService()
	if service == nil {
		t.Fatal("NewModelInfoService() returned nil")
	}
	if service.Options == nil {
		t.Error("ModelInfoService.Options is nil")
	}
}

func TestNewModelUpdateService(t *testing.T) {
	service := NewModelUpdateService()
	if service == nil {
		t.Fatal("NewModelUpdateService() returned nil")
	}
	if service.Options == nil {
		t.Error("ModelUpdateService.Options is nil")
	}
}

func TestModelInfoParam_MarshalJSON(t *testing.T) {
	param := ModelInfoParam{
		ID: F("test-id"),
	}

	data, err := param.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON() error = %v", err)
	}
	if data == nil {
		t.Error("MarshalJSON() returned nil data")
	}
}

func TestModelInfoTier_IsKnown(t *testing.T) {
	tests := []struct {
		tier    ModelInfoTier
		isKnown bool
	}{
		{ModelInfoTierFree, true},
		{ModelInfoTierPaid, true},
		{ModelInfoTier("unknown"), false},
	}

	for _, tt := range tests {
		t.Run(string(tt.tier), func(t *testing.T) {
			if got := tt.tier.IsKnown(); got != tt.isKnown {
				t.Errorf("IsKnown() = %v, want %v", got, tt.isKnown)
			}
		})
	}
}