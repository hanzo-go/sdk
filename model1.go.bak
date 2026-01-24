// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ModelService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
	Info    *ModelInfoService
	Update  *ModelUpdateService
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	r.Info = NewModelInfoService(opts...)
	r.Update = NewModelUpdateService(opts...)
	return
}

// Allows adding new models to the model list in the config.yaml
func (r *ModelService) New(ctx context.Context, body ModelNewParams, opts ...option.RequestOption) (res *ModelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "model/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows deleting models in the model list in the config.yaml
func (r *ModelService) Delete(ctx context.Context, body ModelDeleteParams, opts ...option.RequestOption) (res *ModelDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "model/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

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

type ModelNewResponse = interface{}

type ModelDeleteResponse = interface{}

type ModelNewParams struct {
	// LiteLLM Params with 'model' requirement - used for completions
	LitellmParams param.Field[ModelNewParamsLitellmParams] `json:"litellm_params,required"`
	ModelInfo     param.Field[ModelInfoParam]              `json:"model_info,required"`
	ModelName     param.Field[string]                      `json:"model_name,required"`
}

func (r ModelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// LiteLLM Params with 'model' requirement - used for completions
type ModelNewParamsLitellmParams struct {
	Model                                      param.Field[string]                                                            `json:"model,required"`
	APIBase                                    param.Field[string]                                                            `json:"api_base"`
	APIKey                                     param.Field[string]                                                            `json:"api_key"`
	APIVersion                                 param.Field[string]                                                            `json:"api_version"`
	AutoRouterConfig                           param.Field[string]                                                            `json:"auto_router_config"`
	AutoRouterConfigPath                       param.Field[string]                                                            `json:"auto_router_config_path"`
	AutoRouterDefaultModel                     param.Field[string]                                                            `json:"auto_router_default_model"`
	AutoRouterEmbeddingModel                   param.Field[string]                                                            `json:"auto_router_embedding_model"`
	AwsAccessKeyID                             param.Field[string]                                                            `json:"aws_access_key_id"`
	AwsBedrockRuntimeEndpoint                  param.Field[string]                                                            `json:"aws_bedrock_runtime_endpoint"`
	AwsRegionName                              param.Field[string]                                                            `json:"aws_region_name"`
	AwsSecretAccessKey                         param.Field[string]                                                            `json:"aws_secret_access_key"`
	BudgetDuration                             param.Field[string]                                                            `json:"budget_duration"`
	CacheCreationInputAudioTokenCost           param.Field[float64]                                                           `json:"cache_creation_input_audio_token_cost"`
	CacheCreationInputTokenCost                param.Field[float64]                                                           `json:"cache_creation_input_token_cost"`
	CacheCreationInputTokenCostAbove1hr        param.Field[float64]                                                           `json:"cache_creation_input_token_cost_above_1hr"`
	CacheCreationInputTokenCostAbove200kTokens param.Field[float64]                                                           `json:"cache_creation_input_token_cost_above_200k_tokens"`
	CacheReadInputAudioTokenCost               param.Field[float64]                                                           `json:"cache_read_input_audio_token_cost"`
	CacheReadInputTokenCost                    param.Field[float64]                                                           `json:"cache_read_input_token_cost"`
	CacheReadInputTokenCostAbove200kTokens     param.Field[float64]                                                           `json:"cache_read_input_token_cost_above_200k_tokens"`
	CacheReadInputTokenCostFlex                param.Field[float64]                                                           `json:"cache_read_input_token_cost_flex"`
	CacheReadInputTokenCostPriority            param.Field[float64]                                                           `json:"cache_read_input_token_cost_priority"`
	CitationCostPerToken                       param.Field[float64]                                                           `json:"citation_cost_per_token"`
	ConfigurableClientsideAuthParams           param.Field[[]ModelNewParamsLitellmParamsConfigurableClientsideAuthParamUnion] `json:"configurable_clientside_auth_params"`
	CustomLlmProvider                          param.Field[string]                                                            `json:"custom_llm_provider"`
	GcsBucketName                              param.Field[string]                                                            `json:"gcs_bucket_name"`
	InputCostPerAudioPerSecond                 param.Field[float64]                                                           `json:"input_cost_per_audio_per_second"`
	InputCostPerAudioPerSecondAbove128kTokens  param.Field[float64]                                                           `json:"input_cost_per_audio_per_second_above_128k_tokens"`
	InputCostPerAudioToken                     param.Field[float64]                                                           `json:"input_cost_per_audio_token"`
	InputCostPerCharacter                      param.Field[float64]                                                           `json:"input_cost_per_character"`
	InputCostPerCharacterAbove128kTokens       param.Field[float64]                                                           `json:"input_cost_per_character_above_128k_tokens"`
	InputCostPerImage                          param.Field[float64]                                                           `json:"input_cost_per_image"`
	InputCostPerImageAbove128kTokens           param.Field[float64]                                                           `json:"input_cost_per_image_above_128k_tokens"`
	InputCostPerPixel                          param.Field[float64]                                                           `json:"input_cost_per_pixel"`
	InputCostPerQuery                          param.Field[float64]                                                           `json:"input_cost_per_query"`
	InputCostPerSecond                         param.Field[float64]                                                           `json:"input_cost_per_second"`
	InputCostPerToken                          param.Field[float64]                                                           `json:"input_cost_per_token"`
	InputCostPerTokenAbove128kTokens           param.Field[float64]                                                           `json:"input_cost_per_token_above_128k_tokens"`
	InputCostPerTokenAbove200kTokens           param.Field[float64]                                                           `json:"input_cost_per_token_above_200k_tokens"`
	InputCostPerTokenBatches                   param.Field[float64]                                                           `json:"input_cost_per_token_batches"`
	InputCostPerTokenCacheHit                  param.Field[float64]                                                           `json:"input_cost_per_token_cache_hit"`
	InputCostPerTokenFlex                      param.Field[float64]                                                           `json:"input_cost_per_token_flex"`
	InputCostPerTokenPriority                  param.Field[float64]                                                           `json:"input_cost_per_token_priority"`
	InputCostPerVideoPerSecond                 param.Field[float64]                                                           `json:"input_cost_per_video_per_second"`
	InputCostPerVideoPerSecondAbove128kTokens  param.Field[float64]                                                           `json:"input_cost_per_video_per_second_above_128k_tokens"`
	InputCostPerVideoPerSecondAbove15sInterval param.Field[float64]                                                           `json:"input_cost_per_video_per_second_above_15s_interval"`
	InputCostPerVideoPerSecondAbove8sInterval  param.Field[float64]                                                           `json:"input_cost_per_video_per_second_above_8s_interval"`
	LitellmCredentialName                      param.Field[string]                                                            `json:"litellm_credential_name"`
	LitellmTraceID                             param.Field[string]                                                            `json:"litellm_trace_id"`
	MaxBudget                                  param.Field[float64]                                                           `json:"max_budget"`
	MaxFileSizeMB                              param.Field[float64]                                                           `json:"max_file_size_mb"`
	MaxRetries                                 param.Field[int64]                                                             `json:"max_retries"`
	MergeReasoningContentInChoices             param.Field[bool]                                                              `json:"merge_reasoning_content_in_choices"`
	MilvusTextField                            param.Field[string]                                                            `json:"milvus_text_field"`
	MockResponse                               param.Field[ModelNewParamsLitellmParamsMockResponseUnion]                      `json:"mock_response"`
	ModelInfo                                  param.Field[map[string]interface{}]                                            `json:"model_info"`
	Organization                               param.Field[string]                                                            `json:"organization"`
	OutputCostPerAudioPerSecond                param.Field[float64]                                                           `json:"output_cost_per_audio_per_second"`
	OutputCostPerAudioToken                    param.Field[float64]                                                           `json:"output_cost_per_audio_token"`
	OutputCostPerCharacter                     param.Field[float64]                                                           `json:"output_cost_per_character"`
	OutputCostPerCharacterAbove128kTokens      param.Field[float64]                                                           `json:"output_cost_per_character_above_128k_tokens"`
	OutputCostPerImage                         param.Field[float64]                                                           `json:"output_cost_per_image"`
	OutputCostPerImageToken                    param.Field[float64]                                                           `json:"output_cost_per_image_token"`
	OutputCostPerPixel                         param.Field[float64]                                                           `json:"output_cost_per_pixel"`
	OutputCostPerReasoningToken                param.Field[float64]                                                           `json:"output_cost_per_reasoning_token"`
	OutputCostPerSecond                        param.Field[float64]                                                           `json:"output_cost_per_second"`
	OutputCostPerToken                         param.Field[float64]                                                           `json:"output_cost_per_token"`
	OutputCostPerTokenAbove128kTokens          param.Field[float64]                                                           `json:"output_cost_per_token_above_128k_tokens"`
	OutputCostPerTokenAbove200kTokens          param.Field[float64]                                                           `json:"output_cost_per_token_above_200k_tokens"`
	OutputCostPerTokenBatches                  param.Field[float64]                                                           `json:"output_cost_per_token_batches"`
	OutputCostPerTokenFlex                     param.Field[float64]                                                           `json:"output_cost_per_token_flex"`
	OutputCostPerTokenPriority                 param.Field[float64]                                                           `json:"output_cost_per_token_priority"`
	OutputCostPerVideoPerSecond                param.Field[float64]                                                           `json:"output_cost_per_video_per_second"`
	RegionName                                 param.Field[string]                                                            `json:"region_name"`
	Rpm                                        param.Field[int64]                                                             `json:"rpm"`
	S3BucketName                               param.Field[string]                                                            `json:"s3_bucket_name"`
	S3EncryptionKeyID                          param.Field[string]                                                            `json:"s3_encryption_key_id"`
	SearchContextCostPerQuery                  param.Field[map[string]interface{}]                                            `json:"search_context_cost_per_query"`
	StreamTimeout                              param.Field[ModelNewParamsLitellmParamsStreamTimeoutUnion]                     `json:"stream_timeout"`
	TieredPricing                              param.Field[[]map[string]interface{}]                                          `json:"tiered_pricing"`
	Timeout                                    param.Field[ModelNewParamsLitellmParamsTimeoutUnion]                           `json:"timeout"`
	Tpm                                        param.Field[int64]                                                             `json:"tpm"`
	UseInPassThrough                           param.Field[bool]                                                              `json:"use_in_pass_through"`
	UseLitellmProxy                            param.Field[bool]                                                              `json:"use_litellm_proxy"`
	VectorStoreID                              param.Field[string]                                                            `json:"vector_store_id"`
	VertexCredentials                          param.Field[ModelNewParamsLitellmParamsVertexCredentialsUnion]                 `json:"vertex_credentials"`
	VertexLocation                             param.Field[string]                                                            `json:"vertex_location"`
	VertexProject                              param.Field[string]                                                            `json:"vertex_project"`
	WatsonxRegionName                          param.Field[string]                                                            `json:"watsonx_region_name"`
	ExtraFields                                map[string]interface{}                                                         `json:"-,extras"`
}

func (r ModelNewParamsLitellmParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [ModelNewParamsLitellmParamsConfigurableClientsideAuthParamsConfigurableClientsideParamsCustomAuthInput].
type ModelNewParamsLitellmParamsConfigurableClientsideAuthParamUnion interface {
	ImplementsModelNewParamsLitellmParamsConfigurableClientsideAuthParamUnion()
}

type ModelNewParamsLitellmParamsConfigurableClientsideAuthParamsConfigurableClientsideParamsCustomAuthInput struct {
	APIBase     param.Field[string]    `json:"api_base,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsConfigurableClientsideAuthParamsConfigurableClientsideParamsCustomAuthInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ModelNewParamsLitellmParamsConfigurableClientsideAuthParamsConfigurableClientsideParamsCustomAuthInput) ImplementsModelNewParamsLitellmParamsConfigurableClientsideAuthParamUnion() {
}

// Satisfied by [shared.UnionString],
// [ModelNewParamsLitellmParamsMockResponseModelResponse].
//
// Use [Raw()] to specify an arbitrary value for this param
type ModelNewParamsLitellmParamsMockResponseUnion interface {
	ImplementsModelNewParamsLitellmParamsMockResponseUnion()
}

type ModelNewParamsLitellmParamsMockResponseModelResponse struct {
	ID                param.Field[string]                                                            `json:"id,required"`
	Choices           param.Field[[]ModelNewParamsLitellmParamsMockResponseModelResponseChoiceUnion] `json:"choices,required"`
	Created           param.Field[int64]                                                             `json:"created,required"`
	Object            param.Field[string]                                                            `json:"object,required"`
	Model             param.Field[string]                                                            `json:"model"`
	SystemFingerprint param.Field[string]                                                            `json:"system_fingerprint"`
	ExtraFields       map[string]interface{}                                                         `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponse) ImplementsModelNewParamsLitellmParamsMockResponseUnion() {
}

// Satisfied by
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoices],
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesStreamingChoices].
type ModelNewParamsLitellmParamsMockResponseModelResponseChoiceUnion interface {
	implementsModelNewParamsLitellmParamsMockResponseModelResponseChoiceUnion()
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoices struct {
	FinishReason           param.Field[string]                                                                     `json:"finish_reason,required"`
	Index                  param.Field[int64]                                                                      `json:"index,required"`
	Message                param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessage]  `json:"message,required"`
	Logprobs               param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobs] `json:"logprobs"`
	ProviderSpecificFields param.Field[map[string]interface{}]                                                     `json:"provider_specific_fields"`
	ExtraFields            map[string]interface{}                                                                  `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoices) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoiceUnion() {
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessage struct {
	Content                param.Field[string]                                                                                        `json:"content,required"`
	FunctionCall           param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageFunctionCall]         `json:"function_call,required"`
	Role                   param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole]                 `json:"role,required"`
	ToolCalls              param.Field[[]map[string]interface{}]                                                                      `json:"tool_calls,required"`
	Annotations            param.Field[[]ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotation]         `json:"annotations"`
	Audio                  param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAudio]                `json:"audio"`
	Images                 param.Field[[]ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImage]              `json:"images"`
	ProviderSpecificFields param.Field[map[string]interface{}]                                                                        `json:"provider_specific_fields"`
	ReasoningContent       param.Field[string]                                                                                        `json:"reasoning_content"`
	ThinkingBlocks         param.Field[[]ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockUnion] `json:"thinking_blocks"`
	ExtraFields            map[string]interface{}                                                                                     `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageFunctionCall struct {
	Arguments   param.Field[string]    `json:"arguments,required"`
	Name        param.Field[string]    `json:"name"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole string

const (
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleAssistant ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "assistant"
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleUser      ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "user"
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleSystem    ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "system"
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleTool      ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "tool"
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleFunction  ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole = "function"
)

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRole) IsKnown() bool {
	switch r {
	case ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleAssistant, ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleUser, ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleSystem, ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleTool, ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageRoleFunction:
		return true
	}
	return false
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotation struct {
	Type        param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsType]        `json:"type"`
	URLCitation param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsURLCitation] `json:"url_citation"`
	ExtraFields map[string]interface{}                                                                                       `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsType string

const (
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsTypeURLCitation ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsType = "url_citation"
)

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsType) IsKnown() bool {
	switch r {
	case ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsTypeURLCitation:
		return true
	}
	return false
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsURLCitation struct {
	EndIndex    param.Field[int64]     `json:"end_index"`
	StartIndex  param.Field[int64]     `json:"start_index"`
	Title       param.Field[string]    `json:"title"`
	URL         param.Field[string]    `json:"url"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAnnotationsURLCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAudio struct {
	ID          param.Field[string]    `json:"id,required"`
	Data        param.Field[string]    `json:"data,required"`
	ExpiresAt   param.Field[int64]     `json:"expires_at,required"`
	Transcript  param.Field[string]    `json:"transcript,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImage struct {
	ImageURL    param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesImageURL] `json:"image_url,required"`
	Index       param.Field[int64]                                                                                   `json:"index,required"`
	Type        param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesType]     `json:"type,required"`
	ExtraFields map[string]interface{}                                                                               `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesImageURL struct {
	URL         param.Field[string]    `json:"url,required"`
	Detail      param.Field[string]    `json:"detail"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesType string

const (
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesTypeImageURL ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesType = "image_url"
)

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesType) IsKnown() bool {
	switch r {
	case ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageImagesTypeImageURL:
		return true
	}
	return false
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlock struct {
	Type         param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                                 `json:"cache_control"`
	Data         param.Field[string]                                                                                      `json:"data"`
	Signature    param.Field[string]                                                                                      `json:"signature"`
	Thinking     param.Field[string]                                                                                      `json:"thinking"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlock) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockUnion() {
}

// Satisfied by
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlock],
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlock],
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlock].
type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockUnion interface {
	implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockUnion()
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlock struct {
	Type         param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion] `json:"cache_control"`
	Signature    param.Field[string]                                                                                                                              `json:"signature"`
	Thinking     param.Field[string]                                                                                                                              `json:"thinking"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlock) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockUnion() {
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockType string

const (
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockType = "thinking"
)

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockType) IsKnown() bool {
	switch r {
	case ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking:
		return true
	}
	return false
}

// Satisfied by
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap],
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent].
type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion interface {
	implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion()
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap map[string]interface{}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion() {
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion() {
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlock struct {
	Type         param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion] `json:"cache_control"`
	Data         param.Field[string]                                                                                                                                      `json:"data"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlock) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlockUnion() {
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockType string

const (
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockType = "redacted_thinking"
)

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

// Satisfied by
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap],
// [ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent].
type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion interface {
	implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion()
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap map[string]interface{}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion() {
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion() {
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType string

const (
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksTypeThinking         ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType = "thinking"
	ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksTypeRedactedThinking ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType = "redacted_thinking"
)

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksType) IsKnown() bool {
	switch r {
	case ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksTypeThinking, ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesMessageThinkingBlocksTypeRedactedThinking:
		return true
	}
	return false
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobs struct {
	Content     param.Field[[]ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContent] `json:"content"`
	ExtraFields map[string]interface{}                                                                           `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContent struct {
	Token       param.Field[string]                                                                                        `json:"token,required"`
	Logprob     param.Field[float64]                                                                                       `json:"logprob,required"`
	TopLogprobs param.Field[[]ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentTopLogprob] `json:"top_logprobs,required"`
	Bytes       param.Field[[]int64]                                                                                       `json:"bytes"`
	ExtraFields map[string]interface{}                                                                                     `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentTopLogprob struct {
	Token       param.Field[string]    `json:"token,required"`
	Logprob     param.Field[float64]   `json:"logprob,required"`
	Bytes       param.Field[[]int64]   `json:"bytes"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesChoicesLogprobsContentTopLogprob) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelNewParamsLitellmParamsMockResponseModelResponseChoicesStreamingChoices map[string]interface{}

func (r ModelNewParamsLitellmParamsMockResponseModelResponseChoicesStreamingChoices) implementsModelNewParamsLitellmParamsMockResponseModelResponseChoiceUnion() {
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type ModelNewParamsLitellmParamsStreamTimeoutUnion interface {
	ImplementsModelNewParamsLitellmParamsStreamTimeoutUnion()
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type ModelNewParamsLitellmParamsTimeoutUnion interface {
	ImplementsModelNewParamsLitellmParamsTimeoutUnion()
}

// Satisfied by [shared.UnionString],
// [ModelNewParamsLitellmParamsVertexCredentialsMap].
type ModelNewParamsLitellmParamsVertexCredentialsUnion interface {
	ImplementsModelNewParamsLitellmParamsVertexCredentialsUnion()
}

type ModelNewParamsLitellmParamsVertexCredentialsMap map[string]interface{}

func (r ModelNewParamsLitellmParamsVertexCredentialsMap) ImplementsModelNewParamsLitellmParamsVertexCredentialsUnion() {
}

type ModelDeleteParams struct {
	ID param.Field[string] `json:"id,required"`
}

func (r ModelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
