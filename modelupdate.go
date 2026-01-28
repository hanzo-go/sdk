// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"slices"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ModelUpdateService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelUpdateService] method instead.
type ModelUpdateService struct {
	Options []option.RequestOption
}

// NewModelUpdateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelUpdateService(opts ...option.RequestOption) (r *ModelUpdateService) {
	r = &ModelUpdateService{}
	r.Options = opts
	return
}

// Edit existing model params
func (r *ModelUpdateService) Full(ctx context.Context, body ModelUpdateFullParams, opts ...option.RequestOption) (res *ModelUpdateFullResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "model/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PATCH Endpoint for partial model updates.
//
// Only updates the fields specified in the request while preserving other existing
// values. Follows proper PATCH semantics by only modifying provided fields.
//
// Args: model_id: The ID of the model to update patch_data: The fields to update
// and their new values user_api_key_dict: User authentication information
//
// Returns: Updated model information
//
// Raises: ProxyException: For various error conditions including authentication
// and database errors
func (r *ModelUpdateService) Partial(ctx context.Context, modelID string, body ModelUpdatePartialParams, opts ...option.RequestOption) (res *ModelUpdatePartialResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("model/%s/update", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// ModelInfoParam represents model information parameters
type ModelInfoParam struct {
	ID                  param.Field[string]        `json:"id,required"`
	BaseModel           param.Field[string]        `json:"base_model"`
	CreatedAt           param.Field[time.Time]     `json:"created_at" format:"date-time"`
	CreatedBy           param.Field[string]        `json:"created_by"`
	DBModel             param.Field[bool]          `json:"db_model"`
	TeamID              param.Field[string]        `json:"team_id"`
	TeamPublicModelName param.Field[string]        `json:"team_public_model_name"`
	Tier                param.Field[ModelInfoTier] `json:"tier"`
	UpdatedAt           param.Field[time.Time]     `json:"updated_at" format:"date-time"`
	UpdatedBy           param.Field[string]        `json:"updated_by"`
	ExtraFields         map[string]interface{}     `json:"-,extras"`
}

func (r ModelInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelInfoTier string

const (
	ModelInfoTierFree ModelInfoTier = "free"
	ModelInfoTierPaid ModelInfoTier = "paid"
)

func (r ModelInfoTier) IsKnown() bool {
	switch r {
	case ModelInfoTierFree, ModelInfoTierPaid:
		return true
	}
	return false
}

type UpdateDeploymentParam struct {
	LitellmParams param.Field[UpdateDeploymentLitellmParamsParam] `json:"litellm_params"`
	ModelInfo     param.Field[ModelInfoParam]                     `json:"model_info"`
	ModelName     param.Field[string]                             `json:"model_name"`
}

func (r UpdateDeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsParam struct {
	APIBase                                    param.Field[string]                                                                    `json:"api_base"`
	APIKey                                     param.Field[string]                                                                    `json:"api_key"`
	APIVersion                                 param.Field[string]                                                                    `json:"api_version"`
	AutoRouterConfig                           param.Field[string]                                                                    `json:"auto_router_config"`
	AutoRouterConfigPath                       param.Field[string]                                                                    `json:"auto_router_config_path"`
	AutoRouterDefaultModel                     param.Field[string]                                                                    `json:"auto_router_default_model"`
	AutoRouterEmbeddingModel                   param.Field[string]                                                                    `json:"auto_router_embedding_model"`
	AwsAccessKeyID                             param.Field[string]                                                                    `json:"aws_access_key_id"`
	AwsBedrockRuntimeEndpoint                  param.Field[string]                                                                    `json:"aws_bedrock_runtime_endpoint"`
	AwsRegionName                              param.Field[string]                                                                    `json:"aws_region_name"`
	AwsSecretAccessKey                         param.Field[string]                                                                    `json:"aws_secret_access_key"`
	BudgetDuration                             param.Field[string]                                                                    `json:"budget_duration"`
	CacheCreationInputAudioTokenCost           param.Field[float64]                                                                   `json:"cache_creation_input_audio_token_cost"`
	CacheCreationInputTokenCost                param.Field[float64]                                                                   `json:"cache_creation_input_token_cost"`
	CacheCreationInputTokenCostAbove1hr        param.Field[float64]                                                                   `json:"cache_creation_input_token_cost_above_1hr"`
	CacheCreationInputTokenCostAbove200kTokens param.Field[float64]                                                                   `json:"cache_creation_input_token_cost_above_200k_tokens"`
	CacheReadInputAudioTokenCost               param.Field[float64]                                                                   `json:"cache_read_input_audio_token_cost"`
	CacheReadInputTokenCost                    param.Field[float64]                                                                   `json:"cache_read_input_token_cost"`
	CacheReadInputTokenCostAbove200kTokens     param.Field[float64]                                                                   `json:"cache_read_input_token_cost_above_200k_tokens"`
	CacheReadInputTokenCostFlex                param.Field[float64]                                                                   `json:"cache_read_input_token_cost_flex"`
	CacheReadInputTokenCostPriority            param.Field[float64]                                                                   `json:"cache_read_input_token_cost_priority"`
	CitationCostPerToken                       param.Field[float64]                                                                   `json:"citation_cost_per_token"`
	ConfigurableClientsideAuthParams           param.Field[[]UpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsUnionParam] `json:"configurable_clientside_auth_params"`
	CustomLlmProvider                          param.Field[string]                                                                    `json:"custom_llm_provider"`
	GcsBucketName                              param.Field[string]                                                                    `json:"gcs_bucket_name"`
	InputCostPerAudioPerSecond                 param.Field[float64]                                                                   `json:"input_cost_per_audio_per_second"`
	InputCostPerAudioPerSecondAbove128kTokens  param.Field[float64]                                                                   `json:"input_cost_per_audio_per_second_above_128k_tokens"`
	InputCostPerAudioToken                     param.Field[float64]                                                                   `json:"input_cost_per_audio_token"`
	InputCostPerCharacter                      param.Field[float64]                                                                   `json:"input_cost_per_character"`
	InputCostPerCharacterAbove128kTokens       param.Field[float64]                                                                   `json:"input_cost_per_character_above_128k_tokens"`
	InputCostPerImage                          param.Field[float64]                                                                   `json:"input_cost_per_image"`
	InputCostPerImageAbove128kTokens           param.Field[float64]                                                                   `json:"input_cost_per_image_above_128k_tokens"`
	InputCostPerPixel                          param.Field[float64]                                                                   `json:"input_cost_per_pixel"`
	InputCostPerQuery                          param.Field[float64]                                                                   `json:"input_cost_per_query"`
	InputCostPerSecond                         param.Field[float64]                                                                   `json:"input_cost_per_second"`
	InputCostPerToken                          param.Field[float64]                                                                   `json:"input_cost_per_token"`
	InputCostPerTokenAbove128kTokens           param.Field[float64]                                                                   `json:"input_cost_per_token_above_128k_tokens"`
	InputCostPerTokenAbove200kTokens           param.Field[float64]                                                                   `json:"input_cost_per_token_above_200k_tokens"`
	InputCostPerTokenBatches                   param.Field[float64]                                                                   `json:"input_cost_per_token_batches"`
	InputCostPerTokenCacheHit                  param.Field[float64]                                                                   `json:"input_cost_per_token_cache_hit"`
	InputCostPerTokenFlex                      param.Field[float64]                                                                   `json:"input_cost_per_token_flex"`
	InputCostPerTokenPriority                  param.Field[float64]                                                                   `json:"input_cost_per_token_priority"`
	InputCostPerVideoPerSecond                 param.Field[float64]                                                                   `json:"input_cost_per_video_per_second"`
	InputCostPerVideoPerSecondAbove128kTokens  param.Field[float64]                                                                   `json:"input_cost_per_video_per_second_above_128k_tokens"`
	InputCostPerVideoPerSecondAbove15sInterval param.Field[float64]                                                                   `json:"input_cost_per_video_per_second_above_15s_interval"`
	InputCostPerVideoPerSecondAbove8sInterval  param.Field[float64]                                                                   `json:"input_cost_per_video_per_second_above_8s_interval"`
	LitellmCredentialName                      param.Field[string]                                                                    `json:"litellm_credential_name"`
	LitellmTraceID                             param.Field[string]                                                                    `json:"litellm_trace_id"`
	MaxBudget                                  param.Field[float64]                                                                   `json:"max_budget"`
	MaxFileSizeMB                              param.Field[float64]                                                                   `json:"max_file_size_mb"`
	MaxRetries                                 param.Field[int64]                                                                     `json:"max_retries"`
	MergeReasoningContentInChoices             param.Field[bool]                                                                      `json:"merge_reasoning_content_in_choices"`
	MilvusTextField                            param.Field[string]                                                                    `json:"milvus_text_field"`
	MockResponse                               param.Field[UpdateDeploymentLitellmParamsMockResponseUnionParam]                       `json:"mock_response"`
	Model                                      param.Field[string]                                                                    `json:"model"`
	ModelInfo                                  param.Field[map[string]interface{}]                                                    `json:"model_info"`
	Organization                               param.Field[string]                                                                    `json:"organization"`
	OutputCostPerAudioPerSecond                param.Field[float64]                                                                   `json:"output_cost_per_audio_per_second"`
	OutputCostPerAudioToken                    param.Field[float64]                                                                   `json:"output_cost_per_audio_token"`
	OutputCostPerCharacter                     param.Field[float64]                                                                   `json:"output_cost_per_character"`
	OutputCostPerCharacterAbove128kTokens      param.Field[float64]                                                                   `json:"output_cost_per_character_above_128k_tokens"`
	OutputCostPerImage                         param.Field[float64]                                                                   `json:"output_cost_per_image"`
	OutputCostPerImageToken                    param.Field[float64]                                                                   `json:"output_cost_per_image_token"`
	OutputCostPerPixel                         param.Field[float64]                                                                   `json:"output_cost_per_pixel"`
	OutputCostPerReasoningToken                param.Field[float64]                                                                   `json:"output_cost_per_reasoning_token"`
	OutputCostPerSecond                        param.Field[float64]                                                                   `json:"output_cost_per_second"`
	OutputCostPerToken                         param.Field[float64]                                                                   `json:"output_cost_per_token"`
	OutputCostPerTokenAbove128kTokens          param.Field[float64]                                                                   `json:"output_cost_per_token_above_128k_tokens"`
	OutputCostPerTokenAbove200kTokens          param.Field[float64]                                                                   `json:"output_cost_per_token_above_200k_tokens"`
	OutputCostPerTokenBatches                  param.Field[float64]                                                                   `json:"output_cost_per_token_batches"`
	OutputCostPerTokenFlex                     param.Field[float64]                                                                   `json:"output_cost_per_token_flex"`
	OutputCostPerTokenPriority                 param.Field[float64]                                                                   `json:"output_cost_per_token_priority"`
	OutputCostPerVideoPerSecond                param.Field[float64]                                                                   `json:"output_cost_per_video_per_second"`
	RegionName                                 param.Field[string]                                                                    `json:"region_name"`
	Rpm                                        param.Field[int64]                                                                     `json:"rpm"`
	S3BucketName                               param.Field[string]                                                                    `json:"s3_bucket_name"`
	S3EncryptionKeyID                          param.Field[string]                                                                    `json:"s3_encryption_key_id"`
	SearchContextCostPerQuery                  param.Field[map[string]interface{}]                                                    `json:"search_context_cost_per_query"`
	StreamTimeout                              param.Field[UpdateDeploymentLitellmParamsStreamTimeoutUnionParam]                      `json:"stream_timeout"`
	TieredPricing                              param.Field[[]map[string]interface{}]                                                  `json:"tiered_pricing"`
	Timeout                                    param.Field[UpdateDeploymentLitellmParamsTimeoutUnionParam]                            `json:"timeout"`
	Tpm                                        param.Field[int64]                                                                     `json:"tpm"`
	UseInPassThrough                           param.Field[bool]                                                                      `json:"use_in_pass_through"`
	UseLitellmProxy                            param.Field[bool]                                                                      `json:"use_litellm_proxy"`
	VectorStoreID                              param.Field[string]                                                                    `json:"vector_store_id"`
	VertexCredentials                          param.Field[UpdateDeploymentLitellmParamsVertexCredentialsUnionParam]                  `json:"vertex_credentials"`
	VertexLocation                             param.Field[string]                                                                    `json:"vertex_location"`
	VertexProject                              param.Field[string]                                                                    `json:"vertex_project"`
	WatsonxRegionName                          param.Field[string]                                                                    `json:"watsonx_region_name"`
	ExtraFields                                map[string]interface{}                                                                 `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [UpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsConfigurableClientsideParamsCustomAuthInputParam].
type UpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsUnionParam interface {
	ImplementsUpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsUnionParam()
}

type UpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsConfigurableClientsideParamsCustomAuthInputParam struct {
	APIBase     param.Field[string]    `json:"api_base,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsConfigurableClientsideParamsCustomAuthInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsConfigurableClientsideParamsCustomAuthInputParam) ImplementsUpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsUnionParam() {
}

// Satisfied by [shared.UnionString],
// [UpdateDeploymentLitellmParamsMockResponseModelResponseParam].
//
// Use [Raw()] to specify an arbitrary value for this param
type UpdateDeploymentLitellmParamsMockResponseUnionParam interface {
	ImplementsUpdateDeploymentLitellmParamsMockResponseUnionParam()
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseParam struct {
	ID                param.Field[string]                                                                    `json:"id,required"`
	Choices           param.Field[[]UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesUnionParam] `json:"choices,required"`
	Created           param.Field[int64]                                                                     `json:"created,required"`
	Object            param.Field[string]                                                                    `json:"object,required"`
	Model             param.Field[string]                                                                    `json:"model"`
	SystemFingerprint param.Field[string]                                                                    `json:"system_fingerprint"`
	ExtraFields       map[string]interface{}                                                                 `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseParam) ImplementsUpdateDeploymentLitellmParamsMockResponseUnionParam() {
}

// Satisfied by
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesParam],
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesStreamingChoicesParam].
type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesUnionParam interface {
	implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesUnionParam()
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesParam struct {
	FinishReason           param.Field[string]                                                                            `json:"finish_reason,required"`
	Index                  param.Field[int64]                                                                             `json:"index,required"`
	Message                param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageParam]  `json:"message,required"`
	Logprobs               param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsParam] `json:"logprobs"`
	ProviderSpecificFields param.Field[map[string]interface{}]                                                            `json:"provider_specific_fields"`
	ExtraFields            map[string]interface{}                                                                         `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesUnionParam() {
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageParam struct {
	Content                param.Field[string]                                                                                                `json:"content,required"`
	FunctionCall           param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageFunctionCallParam]          `json:"function_call,required"`
	Role                   param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole]                       `json:"role,required"`
	ToolCalls              param.Field[[]map[string]interface{}]                                                                              `json:"tool_calls,required"`
	Annotations            param.Field[[]UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationParam]          `json:"annotations"`
	Audio                  param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAudioParam]                 `json:"audio"`
	Images                 param.Field[[]UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImageParam]               `json:"images"`
	ProviderSpecificFields param.Field[map[string]interface{}]                                                                                `json:"provider_specific_fields"`
	ReasoningContent       param.Field[string]                                                                                                `json:"reasoning_content"`
	ThinkingBlocks         param.Field[[]UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksUnionParam] `json:"thinking_blocks"`
	ExtraFields            map[string]interface{}                                                                                             `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageFunctionCallParam struct {
	Arguments   param.Field[string]    `json:"arguments,required"`
	Name        param.Field[string]    `json:"name"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageFunctionCallParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole string

const (
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleAssistant UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "assistant"
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleUser      UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "user"
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleSystem    UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "system"
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleTool      UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "tool"
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleFunction  UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "function"
)

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole) IsKnown() bool {
	switch r {
	case UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleAssistant, UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleUser, UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleSystem, UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleTool, UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleFunction:
		return true
	}
	return false
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationParam struct {
	Type        param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsType]             `json:"type"`
	URLCitation param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsURLCitationParam] `json:"url_citation"`
	ExtraFields map[string]interface{}                                                                                              `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsType string

const (
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsTypeURLCitation UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsType = "url_citation"
)

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsType) IsKnown() bool {
	switch r {
	case UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsTypeURLCitation:
		return true
	}
	return false
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsURLCitationParam struct {
	EndIndex    param.Field[int64]     `json:"end_index"`
	StartIndex  param.Field[int64]     `json:"start_index"`
	Title       param.Field[string]    `json:"title"`
	URL         param.Field[string]    `json:"url"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsURLCitationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAudioParam struct {
	ID          param.Field[string]    `json:"id,required"`
	Data        param.Field[string]    `json:"data,required"`
	ExpiresAt   param.Field[int64]     `json:"expires_at,required"`
	Transcript  param.Field[string]    `json:"transcript,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageAudioParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImageParam struct {
	ImageURL    param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesImageURLParam] `json:"image_url,required"`
	Index       param.Field[int64]                                                                                          `json:"index,required"`
	Type        param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesType]          `json:"type,required"`
	ExtraFields map[string]interface{}                                                                                      `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesImageURLParam struct {
	URL         param.Field[string]    `json:"url,required"`
	Detail      param.Field[string]    `json:"detail"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesImageURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesType string

const (
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesTypeImageURL UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesType = "image_url"
)

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesType) IsKnown() bool {
	switch r {
	case UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesTypeImageURL:
		return true
	}
	return false
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockParam struct {
	Type         param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                                   `json:"cache_control"`
	Data         param.Field[string]                                                                                        `json:"data"`
	Signature    param.Field[string]                                                                                        `json:"signature"`
	Thinking     param.Field[string]                                                                                        `json:"thinking"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksUnionParam() {
}

// Satisfied by
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockParam],
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockParam],
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockParam].
type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksUnionParam interface {
	implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksUnionParam()
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockParam struct {
	Type         param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockType]                   `json:"type,required"`
	CacheControl param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnionParam] `json:"cache_control"`
	Signature    param.Field[string]                                                                                                                                     `json:"signature"`
	Thinking     param.Field[string]                                                                                                                                     `json:"thinking"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksUnionParam() {
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockType string

const (
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockType = "thinking"
)

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockType) IsKnown() bool {
	switch r {
	case UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking:
		return true
	}
	return false
}

// Satisfied by
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMapParam],
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentParam].
type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnionParam interface {
	implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnionParam()
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMapParam map[string]interface{}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMapParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnionParam() {
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentParam struct {
	Type param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnionParam() {
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockParam struct {
	Type         param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockType]                   `json:"type,required"`
	CacheControl param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnionParam] `json:"cache_control"`
	Data         param.Field[string]                                                                                                                                             `json:"data"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksUnionParam() {
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockType string

const (
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockType = "redacted_thinking"
)

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

// Satisfied by
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMapParam],
// [UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentParam].
type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnionParam interface {
	implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnionParam()
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMapParam map[string]interface{}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMapParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnionParam() {
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentParam struct {
	Type param.Field[UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnionParam() {
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType string

const (
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksTypeThinking         UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType = "thinking"
	UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksTypeRedactedThinking UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType = "redacted_thinking"
)

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType) IsKnown() bool {
	switch r {
	case UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksTypeThinking, UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksTypeRedactedThinking:
		return true
	}
	return false
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsParam struct {
	Content     param.Field[[]UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentParam] `json:"content"`
	ExtraFields map[string]interface{}                                                                                  `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentParam struct {
	Token       param.Field[string]                                                                                               `json:"token,required"`
	Logprob     param.Field[float64]                                                                                              `json:"logprob,required"`
	TopLogprobs param.Field[[]UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentTopLogprobParam] `json:"top_logprobs,required"`
	Bytes       param.Field[[]int64]                                                                                              `json:"bytes"`
	ExtraFields map[string]interface{}                                                                                            `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentTopLogprobParam struct {
	Token       param.Field[string]    `json:"token,required"`
	Logprob     param.Field[float64]   `json:"logprob,required"`
	Bytes       param.Field[[]int64]   `json:"bytes"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentTopLogprobParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesStreamingChoicesParam map[string]interface{}

func (r UpdateDeploymentLitellmParamsMockResponseModelResponseChoicesStreamingChoicesParam) implementsUpdateDeploymentLitellmParamsMockResponseModelResponseChoicesUnionParam() {
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type UpdateDeploymentLitellmParamsStreamTimeoutUnionParam interface {
	ImplementsUpdateDeploymentLitellmParamsStreamTimeoutUnionParam()
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type UpdateDeploymentLitellmParamsTimeoutUnionParam interface {
	ImplementsUpdateDeploymentLitellmParamsTimeoutUnionParam()
}

// Satisfied by [shared.UnionString],
// [UpdateDeploymentLitellmParamsVertexCredentialsMapParam].
type UpdateDeploymentLitellmParamsVertexCredentialsUnionParam interface {
	ImplementsUpdateDeploymentLitellmParamsVertexCredentialsUnionParam()
}

type UpdateDeploymentLitellmParamsVertexCredentialsMapParam map[string]interface{}

func (r UpdateDeploymentLitellmParamsVertexCredentialsMapParam) ImplementsUpdateDeploymentLitellmParamsVertexCredentialsUnionParam() {
}

type ModelUpdateFullResponse = interface{}

type ModelUpdatePartialResponse = interface{}

type ModelUpdateFullParams struct {
	UpdateDeployment UpdateDeploymentParam `json:"update_deployment,required"`
}

func (r ModelUpdateFullParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateDeployment)
}

type ModelUpdatePartialParams struct {
	UpdateDeployment UpdateDeploymentParam `json:"update_deployment,required"`
}

func (r ModelUpdatePartialParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateDeployment)
}
