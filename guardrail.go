// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// GuardrailService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGuardrailService] method instead.
type GuardrailService struct {
	Options []option.RequestOption
}

// NewGuardrailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGuardrailService(opts ...option.RequestOption) (r *GuardrailService) {
	r = &GuardrailService{}
	r.Options = opts
	return
}

// List the guardrails that are available on the proxy server
//
// 👉 [Guardrail docs](https://docs.litellm.ai/docs/proxy/guardrails/quick_start)
//
// Example Request:
//
// ```bash
// curl -X GET "http://localhost:4000/guardrails/list" -H "Authorization: Bearer <your_api_key>"
// ```
//
// Example Response:
//
// ```json
//
//	{
//	  "guardrails": [
//	    {
//	      "guardrail_name": "bedrock-pre-guard",
//	      "guardrail_info": {
//	        "params": [
//	          {
//	            "name": "toxicity_score",
//	            "type": "float",
//	            "description": "Score between 0-1 indicating content toxicity level"
//	          },
//	          {
//	            "name": "pii_detection",
//	            "type": "boolean"
//	          }
//	        ]
//	      }
//	    }
//	  ]
//	}
//
// ```
func (r *GuardrailService) List(ctx context.Context, opts ...option.RequestOption) (res *GuardrailListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "guardrails/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GuardrailListResponse struct {
	Guardrails []GuardrailListResponseGuardrail `json:"guardrails,required"`
	JSON       guardrailListResponseJSON        `json:"-"`
}

// guardrailListResponseJSON contains the JSON metadata for the struct
// [GuardrailListResponse]
type guardrailListResponseJSON struct {
	Guardrails  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GuardrailListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guardrailListResponseJSON) RawJSON() string {
	return r.raw
}

type GuardrailListResponseGuardrail struct {
	GuardrailName               string                                                     `json:"guardrail_name,required"`
	CreatedAt                   time.Time                                                  `json:"created_at,nullable" format:"date-time"`
	GuardrailDefinitionLocation GuardrailListResponseGuardrailsGuardrailDefinitionLocation `json:"guardrail_definition_location"`
	GuardrailID                 string                                                     `json:"guardrail_id,nullable"`
	GuardrailInfo               map[string]interface{}                                     `json:"guardrail_info,nullable"`
	LitellmParams               GuardrailListResponseGuardrailsLitellmParams               `json:"litellm_params,nullable"`
	UpdatedAt                   time.Time                                                  `json:"updated_at,nullable" format:"date-time"`
	JSON                        guardrailListResponseGuardrailJSON                         `json:"-"`
}

// guardrailListResponseGuardrailJSON contains the JSON metadata for the struct
// [GuardrailListResponseGuardrail]
type guardrailListResponseGuardrailJSON struct {
	GuardrailName               apijson.Field
	CreatedAt                   apijson.Field
	GuardrailDefinitionLocation apijson.Field
	GuardrailID                 apijson.Field
	GuardrailInfo               apijson.Field
	LitellmParams               apijson.Field
	UpdatedAt                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *GuardrailListResponseGuardrail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guardrailListResponseGuardrailJSON) RawJSON() string {
	return r.raw
}

type GuardrailListResponseGuardrailsGuardrailDefinitionLocation string

const (
	GuardrailListResponseGuardrailsGuardrailDefinitionLocationDB     GuardrailListResponseGuardrailsGuardrailDefinitionLocation = "db"
	GuardrailListResponseGuardrailsGuardrailDefinitionLocationConfig GuardrailListResponseGuardrailsGuardrailDefinitionLocation = "config"
)

func (r GuardrailListResponseGuardrailsGuardrailDefinitionLocation) IsKnown() bool {
	switch r {
	case GuardrailListResponseGuardrailsGuardrailDefinitionLocationDB, GuardrailListResponseGuardrailsGuardrailDefinitionLocationConfig:
		return true
	}
	return false
}

type GuardrailListResponseGuardrailsLitellmParams struct {
	// Additional provider-specific parameters for generic guardrail APIs
	AdditionalProviderSpecificParams map[string]interface{} `json:"additional_provider_specific_params,nullable"`
	// Base URL for the guardrail service API
	APIBase string `json:"api_base,nullable"`
	// Optional custom API endpoint for Model Armor
	APIEndpoint string `json:"api_endpoint,nullable"`
	// API key for the guardrail service
	APIKey string `json:"api_key,nullable"`
	// Threshold configuration for Lakera guardrail categories
	CategoryThresholds GuardrailListResponseGuardrailsLitellmParamsCategoryThresholds `json:"category_thresholds,nullable"`
	// Path to Google Cloud credentials JSON file or JSON string
	Credentials string `json:"credentials,nullable"`
	// Whether the guardrail is enabled by default
	DefaultOn bool `json:"default_on,nullable"`
	// Configuration for detect-secrets guardrail
	DetectSecretsConfig map[string]interface{} `json:"detect_secrets_config,nullable"`
	// When True, guardrails only receive the latest message for the relevant role
	// (e.g., newest user input pre-call, newest assistant output post-call)
	ExperimentalUseLatestRoleMessageOnly bool `json:"experimental_use_latest_role_message_only,nullable"`
	// Whether to fail the request if Model Armor encounters an error
	FailOnError bool `json:"fail_on_error,nullable"`
	// Name of the guardrail in guardrails.ai
	GuardName string `json:"guard_name,nullable"`
	// Google Cloud location/region (e.g., us-central1)
	Location string `json:"location,nullable"`
	// Will mask request content if guardrail makes any changes
	MaskRequestContent bool `json:"mask_request_content,nullable"`
	// Will mask response content if guardrail makes any changes
	MaskResponseContent bool `json:"mask_response_content,nullable"`
	// Optional field if guardrail requires a 'model' parameter
	Model string `json:"model,nullable"`
	// Recipe for input (LLM request)
	PangeaInputRecipe string `json:"pangea_input_recipe,nullable"`
	// Recipe for output (LLM response)
	PangeaOutputRecipe string `json:"pangea_output_recipe,nullable"`
	// The ID of your Model Armor template
	TemplateID string `json:"template_id,nullable"`
	// Custom message when a guardrail blocks an action. Supports placeholders like
	// {tool_name}, {rule_id}, and {default_message}.
	ViolationMessageTemplate string                                           `json:"violation_message_template,nullable"`
	ExtraFields              map[string]interface{}                           `json:"-,extras"`
	JSON                     guardrailListResponseGuardrailsLitellmParamsJSON `json:"-"`
}

// guardrailListResponseGuardrailsLitellmParamsJSON contains the JSON metadata for
// the struct [GuardrailListResponseGuardrailsLitellmParams]
type guardrailListResponseGuardrailsLitellmParamsJSON struct {
	AdditionalProviderSpecificParams     apijson.Field
	APIBase                              apijson.Field
	APIEndpoint                          apijson.Field
	APIKey                               apijson.Field
	CategoryThresholds                   apijson.Field
	Credentials                          apijson.Field
	DefaultOn                            apijson.Field
	DetectSecretsConfig                  apijson.Field
	ExperimentalUseLatestRoleMessageOnly apijson.Field
	FailOnError                          apijson.Field
	GuardName                            apijson.Field
	Location                             apijson.Field
	MaskRequestContent                   apijson.Field
	MaskResponseContent                  apijson.Field
	Model                                apijson.Field
	PangeaInputRecipe                    apijson.Field
	PangeaOutputRecipe                   apijson.Field
	TemplateID                           apijson.Field
	ViolationMessageTemplate             apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *GuardrailListResponseGuardrailsLitellmParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guardrailListResponseGuardrailsLitellmParamsJSON) RawJSON() string {
	return r.raw
}

// Threshold configuration for Lakera guardrail categories
type GuardrailListResponseGuardrailsLitellmParamsCategoryThresholds struct {
	Jailbreak       float64                                                            `json:"jailbreak"`
	PromptInjection float64                                                            `json:"prompt_injection"`
	ExtraFields     map[string]interface{}                                             `json:"-,extras"`
	JSON            guardrailListResponseGuardrailsLitellmParamsCategoryThresholdsJSON `json:"-"`
}

// guardrailListResponseGuardrailsLitellmParamsCategoryThresholdsJSON contains the
// JSON metadata for the struct
// [GuardrailListResponseGuardrailsLitellmParamsCategoryThresholds]
type guardrailListResponseGuardrailsLitellmParamsCategoryThresholdsJSON struct {
	Jailbreak       apijson.Field
	PromptInjection apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *GuardrailListResponseGuardrailsLitellmParamsCategoryThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guardrailListResponseGuardrailsLitellmParamsCategoryThresholdsJSON) RawJSON() string {
	return r.raw
}
