// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// OpenAIDeploymentChatService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIDeploymentChatService] method instead.
type OpenAIDeploymentChatService struct {
	Options []option.RequestOption
}

// NewOpenAIDeploymentChatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIDeploymentChatService(opts ...option.RequestOption) (r *OpenAIDeploymentChatService) {
	r = &OpenAIDeploymentChatService{}
	r.Options = opts
	return
}

// Follows the exact same API spec as
// `OpenAI's Chat API https://platform.openai.com/docs/api-reference/chat`
//
// ```bash
// curl -X POST http://localhost:4000/v1/chat/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer sk-1234"
//
//	-d '{
//	    "model": "gpt-4o",
//	    "messages": [
//	        {
//	            "role": "user",
//	            "content": "Hello!"
//	        }
//	    ]
//	}'
//
// ```
func (r *OpenAIDeploymentChatService) Complete(ctx context.Context, model string, body OpenAIDeploymentChatCompleteParams, opts ...option.RequestOption) (res *OpenAIDeploymentChatCompleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if model == "" {
		err = errors.New("missing required model parameter")
		return
	}
	path := fmt.Sprintf("openai/deployments/%s/chat/completions", model)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OpenAIDeploymentChatCompleteResponse = interface{}

type OpenAIDeploymentChatCompleteParams struct {
	Messages                  param.Field[[]OpenAIDeploymentChatCompleteParamsMessageUnion]    `json:"messages,required"`
	Model                     param.Field[string]                                              `json:"model,required"`
	Caching                   param.Field[bool]                                                `json:"caching"`
	ContextWindowFallbackDict param.Field[map[string]string]                                   `json:"context_window_fallback_dict"`
	Fallbacks                 param.Field[[]string]                                            `json:"fallbacks"`
	FrequencyPenalty          param.Field[float64]                                             `json:"frequency_penalty"`
	FunctionCall              param.Field[OpenAIDeploymentChatCompleteParamsFunctionCallUnion] `json:"function_call"`
	Functions                 param.Field[[]map[string]interface{}]                            `json:"functions"`
	Guardrails                param.Field[[]string]                                            `json:"guardrails"`
	LogitBias                 param.Field[map[string]float64]                                  `json:"logit_bias"`
	Logprobs                  param.Field[bool]                                                `json:"logprobs"`
	MaxTokens                 param.Field[int64]                                               `json:"max_tokens"`
	Metadata                  param.Field[map[string]interface{}]                              `json:"metadata"`
	N                         param.Field[int64]                                               `json:"n"`
	NumRetries                param.Field[int64]                                               `json:"num_retries"`
	ParallelToolCalls         param.Field[bool]                                                `json:"parallel_tool_calls"`
	PresencePenalty           param.Field[float64]                                             `json:"presence_penalty"`
	ResponseFormat            param.Field[map[string]interface{}]                              `json:"response_format"`
	Seed                      param.Field[int64]                                               `json:"seed"`
	ServiceTier               param.Field[string]                                              `json:"service_tier"`
	Stop                      param.Field[OpenAIDeploymentChatCompleteParamsStopUnion]         `json:"stop"`
	Stream                    param.Field[bool]                                                `json:"stream"`
	StreamOptions             param.Field[map[string]interface{}]                              `json:"stream_options"`
	Temperature               param.Field[float64]                                             `json:"temperature"`
	ToolChoice                param.Field[OpenAIDeploymentChatCompleteParamsToolChoiceUnion]   `json:"tool_choice"`
	Tools                     param.Field[[]map[string]interface{}]                            `json:"tools"`
	TopLogprobs               param.Field[int64]                                               `json:"top_logprobs"`
	TopP                      param.Field[float64]                                             `json:"top_p"`
	User                      param.Field[string]                                              `json:"user"`
}

func (r OpenAIDeploymentChatCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessage struct {
	Role             param.Field[OpenAIDeploymentChatCompleteParamsMessagesRole] `json:"role,required"`
	CacheControl     param.Field[interface{}]                                    `json:"cache_control"`
	Content          param.Field[interface{}]                                    `json:"content"`
	FunctionCall     param.Field[interface{}]                                    `json:"function_call"`
	Name             param.Field[string]                                         `json:"name"`
	ReasoningContent param.Field[string]                                         `json:"reasoning_content"`
	ThinkingBlocks   param.Field[interface{}]                                    `json:"thinking_blocks"`
	ToolCallID       param.Field[string]                                         `json:"tool_call_id"`
	ToolCalls        param.Field[interface{}]                                    `json:"tool_calls"`
}

func (r OpenAIDeploymentChatCompleteParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessage) implementsOpenAIDeploymentChatCompleteParamsMessageUnion() {
}

// Satisfied by
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessage],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessage],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessage],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessage],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessage],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessage],
// [OpenAIDeploymentChatCompleteParamsMessage].
type OpenAIDeploymentChatCompleteParamsMessageUnion interface {
	implementsOpenAIDeploymentChatCompleteParamsMessageUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessage struct {
	Content      param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentUnion] `json:"content,required"`
	Role         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageRole]         `json:"role,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControl] `json:"cache_control"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessage) implementsOpenAIDeploymentChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArray].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArray []OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArray) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItem struct {
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                         `json:"cache_control"`
	Citations    param.Field[interface{}]                                                                         `json:"citations"`
	Context      param.Field[string]                                                                              `json:"context"`
	File         param.Field[interface{}]                                                                         `json:"file"`
	ImageURL     param.Field[interface{}]                                                                         `json:"image_url"`
	InputAudio   param.Field[interface{}]                                                                         `json:"input_audio"`
	Source       param.Field[interface{}]                                                                         `json:"source"`
	Text         param.Field[string]                                                                              `json:"text"`
	Title        param.Field[string]                                                                              `json:"title"`
	VideoURL     param.Field[interface{}]                                                                         `json:"video_url"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItem) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

// Satisfied by
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItem].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion interface {
	implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject struct {
	Text         param.Field[string]                                                                                                              `json:"text,required"`
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType]         `json:"type,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl] `json:"cache_control"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectTypeText OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType = "text"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectTypeText:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject struct {
	ImageURL param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion] `json:"image_url,required"`
	Type     param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType]          `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject struct {
	URL    param.Field[string] `json:"url,required"`
	Detail param.Field[string] `json:"detail"`
	Format param.Field[string] `json:"format"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectTypeImageURL OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType = "image_url"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectTypeImageURL:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject struct {
	InputAudio param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio] `json:"input_audio,required"`
	Type       param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType]       `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio struct {
	Data   param.Field[string]                                                                                                                   `json:"data,required"`
	Format param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat] `json:"format,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatWav OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat = "wav"
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatMP3 OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat = "mp3"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatWav, OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatMP3:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectTypeInputAudio OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType = "input_audio"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectTypeInputAudio:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject struct {
	Citations param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations] `json:"citations,required"`
	Context   param.Field[string]                                                                                                               `json:"context,required"`
	Source    param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource]    `json:"source,required"`
	Title     param.Field[string]                                                                                                               `json:"title,required"`
	Type      param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType]      `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations struct {
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource struct {
	Data      param.Field[string]                                                                                                                `json:"data,required"`
	MediaType param.Field[string]                                                                                                                `json:"media_type,required"`
	Type      param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceTypeText OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType = "text"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceTypeText:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectTypeDocument OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType = "document"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectTypeDocument:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject struct {
	Type     param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType]          `json:"type,required"`
	VideoURL param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion] `json:"video_url,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectTypeVideoURL OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType = "video_url"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectTypeVideoURL:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject struct {
	URL    param.Field[string] `json:"url,required"`
	Detail param.Field[string] `json:"detail"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject struct {
	File param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile] `json:"file,required"`
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile struct {
	FileData param.Field[string] `json:"file_data"`
	FileID   param.Field[string] `json:"file_id"`
	Filename param.Field[string] `json:"filename"`
	Format   param.Field[string] `json:"format"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectTypeFile OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType = "file"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectTypeFile:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeText       OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "text"
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeImageURL   OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "image_url"
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeInputAudio OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "input_audio"
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeDocument   OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "document"
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeVideoURL   OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "video_url"
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeFile       OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "file"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeText, OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeImageURL, OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeInputAudio, OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeDocument, OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeVideoURL, OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeFile:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageRole string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageRoleUser OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageRole = "user"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageRole) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageRoleUser:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControl struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControlType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControlType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControlTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControlType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControlType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessage struct {
	Role             param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageRole]                 `json:"role,required"`
	CacheControl     param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControl]         `json:"cache_control"`
	Content          param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion]         `json:"content"`
	FunctionCall     param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageFunctionCall]         `json:"function_call"`
	Name             param.Field[string]                                                                                       `json:"name"`
	ReasoningContent param.Field[string]                                                                                       `json:"reasoning_content"`
	ThinkingBlocks   param.Field[[]OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion] `json:"thinking_blocks"`
	ToolCalls        param.Field[[]OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCall]           `json:"tool_calls"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessage) implementsOpenAIDeploymentChatCompleteParamsMessageUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageRole string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageRoleAssistant OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageRole = "assistant"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageRole) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageRoleAssistant:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControl struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArray].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArray []OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArray) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItem struct {
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                              `json:"cache_control"`
	Signature    param.Field[string]                                                                                   `json:"signature"`
	Text         param.Field[string]                                                                                   `json:"text"`
	Thinking     param.Field[string]                                                                                   `json:"thinking"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItem) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

// Satisfied by
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItem].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion interface {
	implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject struct {
	Text         param.Field[string]                                                                                                                   `json:"text,required"`
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType]         `json:"type,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl] `json:"cache_control"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectTypeText OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType = "text"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectTypeText:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock struct {
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion] `json:"cache_control"`
	Signature    param.Field[string]                                                                                                                           `json:"signature"`
	Thinking     param.Field[string]                                                                                                                           `json:"thinking"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockTypeThinking OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType = "thinking"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockTypeThinking:
		return true
	}
	return false
}

// Satisfied by
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion interface {
	implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap map[string]interface{}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayTypeText     OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType = "text"
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayTypeThinking OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType = "thinking"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayTypeText, OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayTypeThinking:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageFunctionCall struct {
	Arguments              param.Field[string]                 `json:"arguments"`
	Name                   param.Field[string]                 `json:"name"`
	ProviderSpecificFields param.Field[map[string]interface{}] `json:"provider_specific_fields"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlock struct {
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                                `json:"cache_control"`
	Data         param.Field[string]                                                                                     `json:"data"`
	Signature    param.Field[string]                                                                                     `json:"signature"`
	Thinking     param.Field[string]                                                                                     `json:"thinking"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlock) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

// Satisfied by
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlock].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion interface {
	implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock struct {
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion] `json:"cache_control"`
	Signature    param.Field[string]                                                                                                                             `json:"signature"`
	Thinking     param.Field[string]                                                                                                                             `json:"thinking"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType = "thinking"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking:
		return true
	}
	return false
}

// Satisfied by
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion interface {
	implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap map[string]interface{}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock struct {
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion] `json:"cache_control"`
	Data         param.Field[string]                                                                                                                                     `json:"data"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType = "redacted_thinking"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

// Satisfied by
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion interface {
	implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap map[string]interface{}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent) implementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeThinking         OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType = "thinking"
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeRedactedThinking OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType = "redacted_thinking"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeThinking, OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeRedactedThinking:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCall struct {
	ID       param.Field[string]                                                                                    `json:"id,required"`
	Function param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsFunction] `json:"function,required"`
	Type     param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsType]     `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsFunction struct {
	Arguments              param.Field[string]                 `json:"arguments"`
	Name                   param.Field[string]                 `json:"name"`
	ProviderSpecificFields param.Field[map[string]interface{}] `json:"provider_specific_fields"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsTypeFunction OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsType = "function"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsTypeFunction:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessage struct {
	Content    param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentUnion] `json:"content,required"`
	Role       param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageRole]         `json:"role,required"`
	ToolCallID param.Field[string]                                                                          `json:"tool_call_id,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessage) implementsOpenAIDeploymentChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArray].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArray []OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayItem

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArray) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayItem struct {
	Text         param.Field[string]                                                                                      `json:"text,required"`
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayType]         `json:"type,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControl] `json:"cache_control"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayTypeText OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayType = "text"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayTypeText:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControl struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageRole string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageRoleTool OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageRole = "tool"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageRole) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageRoleTool:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessage struct {
	Content      param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion] `json:"content,required"`
	Role         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageRole]         `json:"role,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControl] `json:"cache_control"`
	Name         param.Field[string]                                                                            `json:"name"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessage) implementsOpenAIDeploymentChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageContentArray].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageContentArray []interface{}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageContentArray) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageRole string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageRoleSystem OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageRole = "system"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageRole) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageRoleSystem:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControl struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessage struct {
	Content    param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion] `json:"content,required"`
	Name       param.Field[string]                                                                              `json:"name,required"`
	Role       param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageRole]         `json:"role,required"`
	ToolCallID param.Field[string]                                                                              `json:"tool_call_id,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessage) implementsOpenAIDeploymentChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArray].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArray []OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayItem

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArray) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayItem struct {
	Text         param.Field[string]                                                                                          `json:"text,required"`
	Type         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayType]         `json:"type,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl] `json:"cache_control"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayTypeText OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayType = "text"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayTypeText:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageRole string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageRoleFunction OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageRole = "function"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageRole) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageRoleFunction:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessage struct {
	Content      param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion] `json:"content,required"`
	Role         param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageRole]         `json:"role,required"`
	CacheControl param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControl] `json:"cache_control"`
	Name         param.Field[string]                                                                               `json:"name"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessage) implementsOpenAIDeploymentChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageContentArray].
type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion()
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageContentArray []interface{}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageContentArray) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion() {
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageRole string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageRoleDeveloper OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageRole = "developer"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageRole) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageRoleDeveloper:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControl struct {
	Type param.Field[OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlType] `json:"type,required"`
}

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlType string

const (
	OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlTypeEphemeral OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlType = "ephemeral"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlType) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type OpenAIDeploymentChatCompleteParamsMessagesRole string

const (
	OpenAIDeploymentChatCompleteParamsMessagesRoleUser      OpenAIDeploymentChatCompleteParamsMessagesRole = "user"
	OpenAIDeploymentChatCompleteParamsMessagesRoleAssistant OpenAIDeploymentChatCompleteParamsMessagesRole = "assistant"
	OpenAIDeploymentChatCompleteParamsMessagesRoleTool      OpenAIDeploymentChatCompleteParamsMessagesRole = "tool"
	OpenAIDeploymentChatCompleteParamsMessagesRoleSystem    OpenAIDeploymentChatCompleteParamsMessagesRole = "system"
	OpenAIDeploymentChatCompleteParamsMessagesRoleFunction  OpenAIDeploymentChatCompleteParamsMessagesRole = "function"
	OpenAIDeploymentChatCompleteParamsMessagesRoleDeveloper OpenAIDeploymentChatCompleteParamsMessagesRole = "developer"
)

func (r OpenAIDeploymentChatCompleteParamsMessagesRole) IsKnown() bool {
	switch r {
	case OpenAIDeploymentChatCompleteParamsMessagesRoleUser, OpenAIDeploymentChatCompleteParamsMessagesRoleAssistant, OpenAIDeploymentChatCompleteParamsMessagesRoleTool, OpenAIDeploymentChatCompleteParamsMessagesRoleSystem, OpenAIDeploymentChatCompleteParamsMessagesRoleFunction, OpenAIDeploymentChatCompleteParamsMessagesRoleDeveloper:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsFunctionCallMap].
type OpenAIDeploymentChatCompleteParamsFunctionCallUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsFunctionCallUnion()
}

type OpenAIDeploymentChatCompleteParamsFunctionCallMap map[string]interface{}

func (r OpenAIDeploymentChatCompleteParamsFunctionCallMap) ImplementsOpenAIDeploymentChatCompleteParamsFunctionCallUnion() {
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsStopArray].
type OpenAIDeploymentChatCompleteParamsStopUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsStopUnion()
}

type OpenAIDeploymentChatCompleteParamsStopArray []string

func (r OpenAIDeploymentChatCompleteParamsStopArray) ImplementsOpenAIDeploymentChatCompleteParamsStopUnion() {
}

// Satisfied by [shared.UnionString],
// [OpenAIDeploymentChatCompleteParamsToolChoiceMap].
type OpenAIDeploymentChatCompleteParamsToolChoiceUnion interface {
	ImplementsOpenAIDeploymentChatCompleteParamsToolChoiceUnion()
}

type OpenAIDeploymentChatCompleteParamsToolChoiceMap map[string]interface{}

func (r OpenAIDeploymentChatCompleteParamsToolChoiceMap) ImplementsOpenAIDeploymentChatCompleteParamsToolChoiceUnion() {
}
