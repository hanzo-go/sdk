// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ConfigPassThroughEndpointService contains methods and other services that help
// with interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigPassThroughEndpointService] method instead.
type ConfigPassThroughEndpointService struct {
	Options []option.RequestOption
}

// NewConfigPassThroughEndpointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConfigPassThroughEndpointService(opts ...option.RequestOption) (r *ConfigPassThroughEndpointService) {
	r = &ConfigPassThroughEndpointService{}
	r.Options = opts
	return
}

// Create new pass-through endpoint
func (r *ConfigPassThroughEndpointService) New(ctx context.Context, body ConfigPassThroughEndpointNewParams, opts ...option.RequestOption) (res *ConfigPassThroughEndpointNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config/pass_through_endpoint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a pass-through endpoint by ID.
func (r *ConfigPassThroughEndpointService) Update(ctx context.Context, endpointID string, body ConfigPassThroughEndpointUpdateParams, opts ...option.RequestOption) (res *ConfigPassThroughEndpointUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if endpointID == "" {
		err = errors.New("missing required endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("config/pass_through_endpoint/%s", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GET configured pass through endpoint.
//
// If no endpoint_id given, return all configured endpoints.
func (r *ConfigPassThroughEndpointService) List(ctx context.Context, query ConfigPassThroughEndpointListParams, opts ...option.RequestOption) (res *PassThroughEndpointResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config/pass_through_endpoint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a pass-through endpoint by ID.
//
// Returns - the deleted endpoint
func (r *ConfigPassThroughEndpointService) Delete(ctx context.Context, body ConfigPassThroughEndpointDeleteParams, opts ...option.RequestOption) (res *PassThroughEndpointResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config/pass_through_endpoint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type PassThroughEndpointResponse struct {
	Endpoints []PassThroughGenericEndpoint    `json:"endpoints,required"`
	JSON      passThroughEndpointResponseJSON `json:"-"`
}

// passThroughEndpointResponseJSON contains the JSON metadata for the struct
// [PassThroughEndpointResponse]
type passThroughEndpointResponseJSON struct {
	Endpoints   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PassThroughEndpointResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r passThroughEndpointResponseJSON) RawJSON() string {
	return r.raw
}

type PassThroughGenericEndpoint struct {
	// The route to be added to the LiteLLM Proxy Server.
	Path string `json:"path,required"`
	// The URL to which requests for this path should be forwarded.
	Target string `json:"target,required"`
	// Optional unique identifier for the pass-through endpoint. If not provided,
	// endpoints will be identified by path for backwards compatibility.
	ID string `json:"id,nullable"`
	// Whether authentication is required for the pass-through endpoint. If True,
	// requests to the endpoint will require a valid LiteLLM API key.
	Auth bool `json:"auth"`
	// The USD cost per request to the target endpoint. This is used to calculate the
	// cost of the request to the target endpoint.
	CostPerRequest float64 `json:"cost_per_request"`
	// Guardrails configuration for this passthrough endpoint. Dict keys are guardrail
	// names, values are optional settings for field targeting. When set, all
	// org/team/key level guardrails will also execute. Defaults to None (no guardrails
	// execute).
	Guardrails map[string]PassThroughGenericEndpointGuardrail `json:"guardrails,nullable"`
	// Key-value pairs of headers to be forwarded with the request. You can set any key
	// value pair here and it will be forwarded to your target endpoint
	Headers map[string]interface{} `json:"headers"`
	// If True, requests to subpaths of the path will be forwarded to the target
	// endpoint. For example, if the path is /bria and include_subpath is True,
	// requests to /bria/v1/text-to-image/base/2.3 will be forwarded to the target
	// endpoint.
	IncludeSubpath bool                           `json:"include_subpath"`
	JSON           passThroughGenericEndpointJSON `json:"-"`
}

// passThroughGenericEndpointJSON contains the JSON metadata for the struct
// [PassThroughGenericEndpoint]
type passThroughGenericEndpointJSON struct {
	Path           apijson.Field
	Target         apijson.Field
	ID             apijson.Field
	Auth           apijson.Field
	CostPerRequest apijson.Field
	Guardrails     apijson.Field
	Headers        apijson.Field
	IncludeSubpath apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PassThroughGenericEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r passThroughGenericEndpointJSON) RawJSON() string {
	return r.raw
}

// Settings for a specific guardrail on a passthrough endpoint.
//
// Allows field-level targeting for guardrail execution.
type PassThroughGenericEndpointGuardrail struct {
	// JSONPath expressions for input field targeting (pre_call). Examples: 'query',
	// 'documents[*].text', 'messages[*].content'. If not specified, guardrail runs on
	// entire request payload.
	RequestFields []string `json:"request_fields,nullable"`
	// JSONPath expressions for output field targeting (post_call). Examples:
	// 'results[*].text', 'output'. If not specified, guardrail runs on entire response
	// payload.
	ResponseFields []string                                `json:"response_fields,nullable"`
	JSON           passThroughGenericEndpointGuardrailJSON `json:"-"`
}

// passThroughGenericEndpointGuardrailJSON contains the JSON metadata for the
// struct [PassThroughGenericEndpointGuardrail]
type passThroughGenericEndpointGuardrailJSON struct {
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PassThroughGenericEndpointGuardrail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r passThroughGenericEndpointGuardrailJSON) RawJSON() string {
	return r.raw
}

type PassThroughGenericEndpointParam struct {
	// The route to be added to the LiteLLM Proxy Server.
	Path param.Field[string] `json:"path,required"`
	// The URL to which requests for this path should be forwarded.
	Target param.Field[string] `json:"target,required"`
	// Optional unique identifier for the pass-through endpoint. If not provided,
	// endpoints will be identified by path for backwards compatibility.
	ID param.Field[string] `json:"id"`
	// Whether authentication is required for the pass-through endpoint. If True,
	// requests to the endpoint will require a valid LiteLLM API key.
	Auth param.Field[bool] `json:"auth"`
	// The USD cost per request to the target endpoint. This is used to calculate the
	// cost of the request to the target endpoint.
	CostPerRequest param.Field[float64] `json:"cost_per_request"`
	// Guardrails configuration for this passthrough endpoint. Dict keys are guardrail
	// names, values are optional settings for field targeting. When set, all
	// org/team/key level guardrails will also execute. Defaults to None (no guardrails
	// execute).
	Guardrails param.Field[map[string]PassThroughGenericEndpointGuardrailParam] `json:"guardrails"`
	// Key-value pairs of headers to be forwarded with the request. You can set any key
	// value pair here and it will be forwarded to your target endpoint
	Headers param.Field[map[string]interface{}] `json:"headers"`
	// If True, requests to subpaths of the path will be forwarded to the target
	// endpoint. For example, if the path is /bria and include_subpath is True,
	// requests to /bria/v1/text-to-image/base/2.3 will be forwarded to the target
	// endpoint.
	IncludeSubpath param.Field[bool] `json:"include_subpath"`
}

func (r PassThroughGenericEndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for a specific guardrail on a passthrough endpoint.
//
// Allows field-level targeting for guardrail execution.
type PassThroughGenericEndpointGuardrailParam struct {
	// JSONPath expressions for input field targeting (pre_call). Examples: 'query',
	// 'documents[*].text', 'messages[*].content'. If not specified, guardrail runs on
	// entire request payload.
	RequestFields param.Field[[]string] `json:"request_fields"`
	// JSONPath expressions for output field targeting (post_call). Examples:
	// 'results[*].text', 'output'. If not specified, guardrail runs on entire response
	// payload.
	ResponseFields param.Field[[]string] `json:"response_fields"`
}

func (r PassThroughGenericEndpointGuardrailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigPassThroughEndpointNewResponse = interface{}

type ConfigPassThroughEndpointUpdateResponse = interface{}

type ConfigPassThroughEndpointNewParams struct {
	PassThroughGenericEndpoint PassThroughGenericEndpointParam `json:"pass_through_generic_endpoint,required"`
}

func (r ConfigPassThroughEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PassThroughGenericEndpoint)
}

type ConfigPassThroughEndpointUpdateParams struct {
	PassThroughGenericEndpoint PassThroughGenericEndpointParam `json:"pass_through_generic_endpoint,required"`
}

func (r ConfigPassThroughEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PassThroughGenericEndpoint)
}

type ConfigPassThroughEndpointListParams struct {
	EndpointID param.Field[string] `query:"endpoint_id"`
	TeamID     param.Field[string] `query:"team_id"`
}

// URLQuery serializes [ConfigPassThroughEndpointListParams]'s query parameters as
// `url.Values`.
func (r ConfigPassThroughEndpointListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigPassThroughEndpointDeleteParams struct {
	EndpointID param.Field[string] `query:"endpoint_id,required"`
}

// URLQuery serializes [ConfigPassThroughEndpointDeleteParams]'s query parameters
// as `url.Values`.
func (r ConfigPassThroughEndpointDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
