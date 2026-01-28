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

// EngineChatService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEngineChatService] method instead.
type EngineChatService struct {
	Options []option.RequestOption
}

// NewEngineChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEngineChatService(opts ...option.RequestOption) (r *EngineChatService) {
	r = &EngineChatService{}
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
func (r *EngineChatService) Complete(ctx context.Context, model string, body EngineChatCompleteParams, opts ...option.RequestOption) (res *EngineChatCompleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if model == "" {
		err = errors.New("missing required model parameter")
		return
	}
	path := fmt.Sprintf("engines/%s/chat/completions", model)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EngineChatCompleteResponse = interface{}

type EngineChatCompleteParams struct {
	Messages                  param.Field[[]EngineChatCompleteParamsMessageUnion]    `json:"messages,required"`
	Model                     param.Field[string]                                    `json:"model,required"`
	Caching                   param.Field[bool]                                      `json:"caching"`
	ContextWindowFallbackDict param.Field[map[string]string]                         `json:"context_window_fallback_dict"`
	Fallbacks                 param.Field[[]string]                                  `json:"fallbacks"`
	FrequencyPenalty          param.Field[float64]                                   `json:"frequency_penalty"`
	FunctionCall              param.Field[EngineChatCompleteParamsFunctionCallUnion] `json:"function_call"`
	Functions                 param.Field[[]map[string]interface{}]                  `json:"functions"`
	Guardrails                param.Field[[]string]                                  `json:"guardrails"`
	LogitBias                 param.Field[map[string]float64]                        `json:"logit_bias"`
	Logprobs                  param.Field[bool]                                      `json:"logprobs"`
	MaxTokens                 param.Field[int64]                                     `json:"max_tokens"`
	Metadata                  param.Field[map[string]interface{}]                    `json:"metadata"`
	N                         param.Field[int64]                                     `json:"n"`
	NumRetries                param.Field[int64]                                     `json:"num_retries"`
	ParallelToolCalls         param.Field[bool]                                      `json:"parallel_tool_calls"`
	PresencePenalty           param.Field[float64]                                   `json:"presence_penalty"`
	ResponseFormat            param.Field[map[string]interface{}]                    `json:"response_format"`
	Seed                      param.Field[int64]                                     `json:"seed"`
	ServiceTier               param.Field[string]                                    `json:"service_tier"`
	Stop                      param.Field[EngineChatCompleteParamsStopUnion]         `json:"stop"`
	Stream                    param.Field[bool]                                      `json:"stream"`
	StreamOptions             param.Field[map[string]interface{}]                    `json:"stream_options"`
	Temperature               param.Field[float64]                                   `json:"temperature"`
	ToolChoice                param.Field[EngineChatCompleteParamsToolChoiceUnion]   `json:"tool_choice"`
	Tools                     param.Field[[]map[string]interface{}]                  `json:"tools"`
	TopLogprobs               param.Field[int64]                                     `json:"top_logprobs"`
	TopP                      param.Field[float64]                                   `json:"top_p"`
	User                      param.Field[string]                                    `json:"user"`
}

func (r EngineChatCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessage struct {
	Role             param.Field[EngineChatCompleteParamsMessagesRole] `json:"role,required"`
	CacheControl     param.Field[interface{}]                          `json:"cache_control"`
	Content          param.Field[interface{}]                          `json:"content"`
	FunctionCall     param.Field[interface{}]                          `json:"function_call"`
	Name             param.Field[string]                               `json:"name"`
	ReasoningContent param.Field[string]                               `json:"reasoning_content"`
	ThinkingBlocks   param.Field[interface{}]                          `json:"thinking_blocks"`
	ToolCallID       param.Field[string]                               `json:"tool_call_id"`
	ToolCalls        param.Field[interface{}]                          `json:"tool_calls"`
}

func (r EngineChatCompleteParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessage) implementsEngineChatCompleteParamsMessageUnion() {}

// Satisfied by [EngineChatCompleteParamsMessagesChatCompletionUserMessage],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessage],
// [EngineChatCompleteParamsMessagesChatCompletionToolMessage],
// [EngineChatCompleteParamsMessagesChatCompletionSystemMessage],
// [EngineChatCompleteParamsMessagesChatCompletionFunctionMessage],
// [EngineChatCompleteParamsMessagesChatCompletionDeveloperMessage],
// [EngineChatCompleteParamsMessage].
type EngineChatCompleteParamsMessageUnion interface {
	implementsEngineChatCompleteParamsMessageUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessage struct {
	Content      param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentUnion] `json:"content,required"`
	Role         param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageRole]         `json:"role,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControl] `json:"cache_control"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessage) implementsEngineChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArray].
type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentUnion interface {
	ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArray []EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArray) ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItem struct {
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                               `json:"cache_control"`
	Citations    param.Field[interface{}]                                                               `json:"citations"`
	Context      param.Field[string]                                                                    `json:"context"`
	File         param.Field[interface{}]                                                               `json:"file"`
	ImageURL     param.Field[interface{}]                                                               `json:"image_url"`
	InputAudio   param.Field[interface{}]                                                               `json:"input_audio"`
	Source       param.Field[interface{}]                                                               `json:"source"`
	Text         param.Field[string]                                                                    `json:"text"`
	Title        param.Field[string]                                                                    `json:"title"`
	VideoURL     param.Field[interface{}]                                                               `json:"video_url"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItem) implementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

// Satisfied by
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItem].
type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion interface {
	implementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject struct {
	Text         param.Field[string]                                                                                                    `json:"text,required"`
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType]         `json:"type,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl] `json:"cache_control"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject) implementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectTypeText EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType = "text"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectTypeText:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject struct {
	ImageURL param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion] `json:"image_url,required"`
	Type     param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType]          `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject) implementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

// Satisfied by [shared.UnionString],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject].
type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion interface {
	ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject struct {
	URL    param.Field[string] `json:"url,required"`
	Detail param.Field[string] `json:"detail"`
	Format param.Field[string] `json:"format"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject) ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectTypeImageURL EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType = "image_url"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectTypeImageURL:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject struct {
	InputAudio param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio] `json:"input_audio,required"`
	Type       param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType]       `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject) implementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio struct {
	Data   param.Field[string]                                                                                                         `json:"data,required"`
	Format param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat] `json:"format,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatWav EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat = "wav"
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatMP3 EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat = "mp3"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatWav, EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatMP3:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectTypeInputAudio EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType = "input_audio"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectTypeInputAudio:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject struct {
	Citations param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations] `json:"citations,required"`
	Context   param.Field[string]                                                                                                     `json:"context,required"`
	Source    param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource]    `json:"source,required"`
	Title     param.Field[string]                                                                                                     `json:"title,required"`
	Type      param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType]      `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject) implementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations struct {
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource struct {
	Data      param.Field[string]                                                                                                      `json:"data,required"`
	MediaType param.Field[string]                                                                                                      `json:"media_type,required"`
	Type      param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceTypeText EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType = "text"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceTypeText:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectTypeDocument EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType = "document"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectTypeDocument:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject struct {
	Type     param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType]          `json:"type,required"`
	VideoURL param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion] `json:"video_url,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject) implementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectTypeVideoURL EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType = "video_url"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectTypeVideoURL:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject].
type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion interface {
	ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject struct {
	URL    param.Field[string] `json:"url,required"`
	Detail param.Field[string] `json:"detail"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject) ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject struct {
	File param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile] `json:"file,required"`
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject) implementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile struct {
	FileData param.Field[string] `json:"file_data"`
	FileID   param.Field[string] `json:"file_id"`
	Filename param.Field[string] `json:"filename"`
	Format   param.Field[string] `json:"format"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectTypeFile EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType = "file"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectTypeFile:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeText       EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "text"
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeImageURL   EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "image_url"
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeInputAudio EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "input_audio"
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeDocument   EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "document"
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeVideoURL   EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "video_url"
	EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeFile       EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType = "file"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeText, EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeImageURL, EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeInputAudio, EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeDocument, EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeVideoURL, EngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayTypeFile:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageRole string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageRoleUser EngineChatCompleteParamsMessagesChatCompletionUserMessageRole = "user"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageRole) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageRoleUser:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControl struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControlType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControlType string

const (
	EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControlTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControlType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControlType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionUserMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessage struct {
	Role             param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageRole]                 `json:"role,required"`
	CacheControl     param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControl]         `json:"cache_control"`
	Content          param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion]         `json:"content"`
	FunctionCall     param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageFunctionCall]         `json:"function_call"`
	Name             param.Field[string]                                                                             `json:"name"`
	ReasoningContent param.Field[string]                                                                             `json:"reasoning_content"`
	ThinkingBlocks   param.Field[[]EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion] `json:"thinking_blocks"`
	ToolCalls        param.Field[[]EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCall]           `json:"tool_calls"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessage) implementsEngineChatCompleteParamsMessageUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageRole string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageRoleAssistant EngineChatCompleteParamsMessagesChatCompletionAssistantMessageRole = "assistant"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageRole) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageRoleAssistant:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControl struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArray].
type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion interface {
	ImplementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArray []EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArray) ImplementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItem struct {
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                    `json:"cache_control"`
	Signature    param.Field[string]                                                                         `json:"signature"`
	Text         param.Field[string]                                                                         `json:"text"`
	Thinking     param.Field[string]                                                                         `json:"thinking"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItem) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

// Satisfied by
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItem].
type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion interface {
	implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject struct {
	Text         param.Field[string]                                                                                                         `json:"text,required"`
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType]         `json:"type,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl] `json:"cache_control"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectTypeText EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType = "text"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectTypeText:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock struct {
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion] `json:"cache_control"`
	Signature    param.Field[string]                                                                                                                 `json:"signature"`
	Thinking     param.Field[string]                                                                                                                 `json:"thinking"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockTypeThinking EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType = "thinking"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockTypeThinking:
		return true
	}
	return false
}

// Satisfied by
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent].
type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion interface {
	implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap map[string]interface{}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayTypeText     EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType = "text"
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayTypeThinking EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType = "thinking"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayTypeText, EngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentArrayTypeThinking:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageFunctionCall struct {
	Arguments              param.Field[string]                 `json:"arguments"`
	Name                   param.Field[string]                 `json:"name"`
	ProviderSpecificFields param.Field[map[string]interface{}] `json:"provider_specific_fields"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlock struct {
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                      `json:"cache_control"`
	Data         param.Field[string]                                                                           `json:"data"`
	Signature    param.Field[string]                                                                           `json:"signature"`
	Thinking     param.Field[string]                                                                           `json:"thinking"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlock) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

// Satisfied by
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlock].
type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion interface {
	implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock struct {
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion] `json:"cache_control"`
	Signature    param.Field[string]                                                                                                                   `json:"signature"`
	Thinking     param.Field[string]                                                                                                                   `json:"thinking"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType = "thinking"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking:
		return true
	}
	return false
}

// Satisfied by
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent].
type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion interface {
	implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap map[string]interface{}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock struct {
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion] `json:"cache_control"`
	Data         param.Field[string]                                                                                                                           `json:"data"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType = "redacted_thinking"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

// Satisfied by
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap],
// [EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent].
type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion interface {
	implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap map[string]interface{}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent) implementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeThinking         EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType = "thinking"
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeRedactedThinking EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType = "redacted_thinking"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeThinking, EngineChatCompleteParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeRedactedThinking:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCall struct {
	ID       param.Field[string]                                                                          `json:"id,required"`
	Function param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsFunction] `json:"function,required"`
	Type     param.Field[EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsType]     `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsFunction struct {
	Arguments              param.Field[string]                 `json:"arguments"`
	Name                   param.Field[string]                 `json:"name"`
	ProviderSpecificFields param.Field[map[string]interface{}] `json:"provider_specific_fields"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsType string

const (
	EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsTypeFunction EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsType = "function"
)

func (r EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionAssistantMessageToolCallsTypeFunction:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionToolMessage struct {
	Content    param.Field[EngineChatCompleteParamsMessagesChatCompletionToolMessageContentUnion] `json:"content,required"`
	Role       param.Field[EngineChatCompleteParamsMessagesChatCompletionToolMessageRole]         `json:"role,required"`
	ToolCallID param.Field[string]                                                                `json:"tool_call_id,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionToolMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionToolMessage) implementsEngineChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArray].
type EngineChatCompleteParamsMessagesChatCompletionToolMessageContentUnion interface {
	ImplementsEngineChatCompleteParamsMessagesChatCompletionToolMessageContentUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArray []EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayItem

func (r EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArray) ImplementsEngineChatCompleteParamsMessagesChatCompletionToolMessageContentUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayItem struct {
	Text         param.Field[string]                                                                            `json:"text,required"`
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayType]         `json:"type,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControl] `json:"cache_control"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayType string

const (
	EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayTypeText EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayType = "text"
)

func (r EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayTypeText:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControl struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlType string

const (
	EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionToolMessageContentArrayCacheControlTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionToolMessageRole string

const (
	EngineChatCompleteParamsMessagesChatCompletionToolMessageRoleTool EngineChatCompleteParamsMessagesChatCompletionToolMessageRole = "tool"
)

func (r EngineChatCompleteParamsMessagesChatCompletionToolMessageRole) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionToolMessageRoleTool:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionSystemMessage struct {
	Content      param.Field[EngineChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion] `json:"content,required"`
	Role         param.Field[EngineChatCompleteParamsMessagesChatCompletionSystemMessageRole]         `json:"role,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControl] `json:"cache_control"`
	Name         param.Field[string]                                                                  `json:"name"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionSystemMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionSystemMessage) implementsEngineChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [EngineChatCompleteParamsMessagesChatCompletionSystemMessageContentArray].
type EngineChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion interface {
	ImplementsEngineChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionSystemMessageContentArray []interface{}

func (r EngineChatCompleteParamsMessagesChatCompletionSystemMessageContentArray) ImplementsEngineChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionSystemMessageRole string

const (
	EngineChatCompleteParamsMessagesChatCompletionSystemMessageRoleSystem EngineChatCompleteParamsMessagesChatCompletionSystemMessageRole = "system"
)

func (r EngineChatCompleteParamsMessagesChatCompletionSystemMessageRole) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionSystemMessageRoleSystem:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControl struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlType string

const (
	EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionSystemMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionFunctionMessage struct {
	Content    param.Field[EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion] `json:"content,required"`
	Name       param.Field[string]                                                                    `json:"name,required"`
	Role       param.Field[EngineChatCompleteParamsMessagesChatCompletionFunctionMessageRole]         `json:"role,required"`
	ToolCallID param.Field[string]                                                                    `json:"tool_call_id,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionFunctionMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionFunctionMessage) implementsEngineChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArray].
type EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion interface {
	ImplementsEngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArray []EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayItem

func (r EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArray) ImplementsEngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayItem struct {
	Text         param.Field[string]                                                                                `json:"text,required"`
	Type         param.Field[EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayType]         `json:"type,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl] `json:"cache_control"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayType string

const (
	EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayTypeText EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayType = "text"
)

func (r EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayTypeText:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType string

const (
	EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionFunctionMessageRole string

const (
	EngineChatCompleteParamsMessagesChatCompletionFunctionMessageRoleFunction EngineChatCompleteParamsMessagesChatCompletionFunctionMessageRole = "function"
)

func (r EngineChatCompleteParamsMessagesChatCompletionFunctionMessageRole) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionFunctionMessageRoleFunction:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionDeveloperMessage struct {
	Content      param.Field[EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion] `json:"content,required"`
	Role         param.Field[EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageRole]         `json:"role,required"`
	CacheControl param.Field[EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControl] `json:"cache_control"`
	Name         param.Field[string]                                                                     `json:"name"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionDeveloperMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EngineChatCompleteParamsMessagesChatCompletionDeveloperMessage) implementsEngineChatCompleteParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageContentArray].
type EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion interface {
	ImplementsEngineChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion()
}

type EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageContentArray []interface{}

func (r EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageContentArray) ImplementsEngineChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion() {
}

type EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageRole string

const (
	EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageRoleDeveloper EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageRole = "developer"
)

func (r EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageRole) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageRoleDeveloper:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControl struct {
	Type param.Field[EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlType] `json:"type,required"`
}

func (r EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlType string

const (
	EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlTypeEphemeral EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlType = "ephemeral"
)

func (r EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlType) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesChatCompletionDeveloperMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type EngineChatCompleteParamsMessagesRole string

const (
	EngineChatCompleteParamsMessagesRoleUser      EngineChatCompleteParamsMessagesRole = "user"
	EngineChatCompleteParamsMessagesRoleAssistant EngineChatCompleteParamsMessagesRole = "assistant"
	EngineChatCompleteParamsMessagesRoleTool      EngineChatCompleteParamsMessagesRole = "tool"
	EngineChatCompleteParamsMessagesRoleSystem    EngineChatCompleteParamsMessagesRole = "system"
	EngineChatCompleteParamsMessagesRoleFunction  EngineChatCompleteParamsMessagesRole = "function"
	EngineChatCompleteParamsMessagesRoleDeveloper EngineChatCompleteParamsMessagesRole = "developer"
)

func (r EngineChatCompleteParamsMessagesRole) IsKnown() bool {
	switch r {
	case EngineChatCompleteParamsMessagesRoleUser, EngineChatCompleteParamsMessagesRoleAssistant, EngineChatCompleteParamsMessagesRoleTool, EngineChatCompleteParamsMessagesRoleSystem, EngineChatCompleteParamsMessagesRoleFunction, EngineChatCompleteParamsMessagesRoleDeveloper:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [EngineChatCompleteParamsFunctionCallMap].
type EngineChatCompleteParamsFunctionCallUnion interface {
	ImplementsEngineChatCompleteParamsFunctionCallUnion()
}

type EngineChatCompleteParamsFunctionCallMap map[string]interface{}

func (r EngineChatCompleteParamsFunctionCallMap) ImplementsEngineChatCompleteParamsFunctionCallUnion() {
}

// Satisfied by [shared.UnionString], [EngineChatCompleteParamsStopArray].
type EngineChatCompleteParamsStopUnion interface {
	ImplementsEngineChatCompleteParamsStopUnion()
}

type EngineChatCompleteParamsStopArray []string

func (r EngineChatCompleteParamsStopArray) ImplementsEngineChatCompleteParamsStopUnion() {}

// Satisfied by [shared.UnionString], [EngineChatCompleteParamsToolChoiceMap].
type EngineChatCompleteParamsToolChoiceUnion interface {
	ImplementsEngineChatCompleteParamsToolChoiceUnion()
}

type EngineChatCompleteParamsToolChoiceMap map[string]interface{}

func (r EngineChatCompleteParamsToolChoiceMap) ImplementsEngineChatCompleteParamsToolChoiceUnion() {}
