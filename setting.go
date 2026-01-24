// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"slices"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// SettingService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	return
}

// Returns a list of litellm level settings
//
// This is useful for debugging and ensuring the proxy server is configured
// correctly.
//
// Response schema:
//
// ```
//
//	{
//	    "alerting": _alerting,
//	    "litellm.callbacks": litellm_callbacks,
//	    "litellm.input_callback": litellm_input_callbacks,
//	    "litellm.failure_callback": litellm_failure_callbacks,
//	    "litellm.success_callback": litellm_success_callbacks,
//	    "litellm._async_success_callback": litellm_async_success_callbacks,
//	    "litellm._async_failure_callback": litellm_async_failure_callbacks,
//	    "litellm._async_input_callback": litellm_async_input_callbacks,
//	    "all_litellm_callbacks": all_litellm_callbacks,
//	    "num_callbacks": len(all_litellm_callbacks),
//	    "num_alerting": _num_alerting,
//	    "litellm.request_timeout": litellm.request_timeout,
//	}
//
// ```
func (r *SettingService) Get(ctx context.Context, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SettingGetResponse = interface{}
