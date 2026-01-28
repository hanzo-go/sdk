// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"slices"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ChatCompletionService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatCompletionService] method instead.
type ChatCompletionService struct {
	Options []option.RequestOption
}

// NewChatCompletionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatCompletionService(opts ...option.RequestOption) (r *ChatCompletionService) {
	r = &ChatCompletionService{}
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
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *ChatCompletionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChatCompletionNewResponse = interface{}

type ChatCompletionNewParams struct {
	Messages                  param.Field[[]ChatCompletionNewParamsMessageUnion]    `json:"messages,required"`
	Model                     param.Field[string]                                   `json:"model,required"`
	Caching                   param.Field[bool]                                     `json:"caching"`
	ContextWindowFallbackDict param.Field[map[string]string]                        `json:"context_window_fallback_dict"`
	Fallbacks                 param.Field[[]string]                                 `json:"fallbacks"`
	FrequencyPenalty          param.Field[float64]                                  `json:"frequency_penalty"`
	FunctionCall              param.Field[ChatCompletionNewParamsFunctionCallUnion] `json:"function_call"`
	Functions                 param.Field[[]map[string]interface{}]                 `json:"functions"`
	Guardrails                param.Field[[]string]                                 `json:"guardrails"`
	LogitBias                 param.Field[map[string]float64]                       `json:"logit_bias"`
	Logprobs                  param.Field[bool]                                     `json:"logprobs"`
	MaxTokens                 param.Field[int64]                                    `json:"max_tokens"`
	Metadata                  param.Field[map[string]interface{}]                   `json:"metadata"`
	N                         param.Field[int64]                                    `json:"n"`
	NumRetries                param.Field[int64]                                    `json:"num_retries"`
	ParallelToolCalls         param.Field[bool]                                     `json:"parallel_tool_calls"`
	PresencePenalty           param.Field[float64]                                  `json:"presence_penalty"`
	ResponseFormat            param.Field[map[string]interface{}]                   `json:"response_format"`
	Seed                      param.Field[int64]                                    `json:"seed"`
	ServiceTier               param.Field[string]                                   `json:"service_tier"`
	Stop                      param.Field[ChatCompletionNewParamsStopUnion]         `json:"stop"`
	Stream                    param.Field[bool]                                     `json:"stream"`
	StreamOptions             param.Field[map[string]interface{}]                   `json:"stream_options"`
	Temperature               param.Field[float64]                                  `json:"temperature"`
	ToolChoice                param.Field[ChatCompletionNewParamsToolChoiceUnion]   `json:"tool_choice"`
	Tools                     param.Field[[]map[string]interface{}]                 `json:"tools"`
	TopLogprobs               param.Field[int64]                                    `json:"top_logprobs"`
	TopP                      param.Field[float64]                                  `json:"top_p"`
	User                      param.Field[string]                                   `json:"user"`
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessage struct {
	Role             param.Field[ChatCompletionNewParamsMessagesRole] `json:"role,required"`
	CacheControl     param.Field[interface{}]                         `json:"cache_control"`
	Content          param.Field[interface{}]                         `json:"content"`
	FunctionCall     param.Field[interface{}]                         `json:"function_call"`
	Name             param.Field[string]                              `json:"name"`
	ReasoningContent param.Field[string]                              `json:"reasoning_content"`
	ThinkingBlocks   param.Field[interface{}]                         `json:"thinking_blocks"`
	ToolCallID       param.Field[string]                              `json:"tool_call_id"`
	ToolCalls        param.Field[interface{}]                         `json:"tool_calls"`
}

func (r ChatCompletionNewParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessage) implementsChatCompletionNewParamsMessageUnion() {}

// Satisfied by [ChatCompletionNewParamsMessagesChatCompletionUserMessage],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessage],
// [ChatCompletionNewParamsMessagesChatCompletionToolMessage],
// [ChatCompletionNewParamsMessagesChatCompletionSystemMessage],
// [ChatCompletionNewParamsMessagesChatCompletionFunctionMessage],
// [ChatCompletionNewParamsMessagesChatCompletionDeveloperMessage],
// [ChatCompletionNewParamsMessage].
type ChatCompletionNewParamsMessageUnion interface {
	implementsChatCompletionNewParamsMessageUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessage struct {
	Content      param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentUnion] `json:"content,required"`
	Role         param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageRole]         `json:"role,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControl] `json:"cache_control"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessage) implementsChatCompletionNewParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArray].
type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArray []ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArray) ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItem struct {
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                              `json:"cache_control"`
	Citations    param.Field[interface{}]                                                              `json:"citations"`
	Context      param.Field[string]                                                                   `json:"context"`
	File         param.Field[interface{}]                                                              `json:"file"`
	ImageURL     param.Field[interface{}]                                                              `json:"image_url"`
	InputAudio   param.Field[interface{}]                                                              `json:"input_audio"`
	Source       param.Field[interface{}]                                                              `json:"source"`
	Text         param.Field[string]                                                                   `json:"text"`
	Title        param.Field[string]                                                                   `json:"title"`
	VideoURL     param.Field[interface{}]                                                              `json:"video_url"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItem) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

// Satisfied by
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItem].
type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion interface {
	implementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject struct {
	Text         param.Field[string]                                                                                                   `json:"text,required"`
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType]         `json:"type,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl] `json:"cache_control"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObject) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectTypeText ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType = "text"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectTypeText:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject struct {
	ImageURL param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion] `json:"image_url,required"`
	Type     param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType]          `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObject) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject].
type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject struct {
	URL    param.Field[string] `json:"url,required"`
	Detail param.Field[string] `json:"detail"`
	Format param.Field[string] `json:"format"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLChatCompletionImageURLObject) ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectTypeImageURL ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType = "image_url"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectTypeImageURL:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject struct {
	InputAudio param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio] `json:"input_audio,required"`
	Type       param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType]       `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObject) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio struct {
	Data   param.Field[string]                                                                                                        `json:"data,required"`
	Format param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat] `json:"format,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatWav ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat = "wav"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatMP3 ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat = "mp3"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormat) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatWav, ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectInputAudioFormatMP3:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectTypeInputAudio ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType = "input_audio"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionAudioObjectTypeInputAudio:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject struct {
	Citations param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations] `json:"citations,required"`
	Context   param.Field[string]                                                                                                    `json:"context,required"`
	Source    param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource]    `json:"source,required"`
	Title     param.Field[string]                                                                                                    `json:"title,required"`
	Type      param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType]      `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObject) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations struct {
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource struct {
	Data      param.Field[string]                                                                                                     `json:"data,required"`
	MediaType param.Field[string]                                                                                                     `json:"media_type,required"`
	Type      param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceTypeText ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType = "text"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectSourceTypeText:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectTypeDocument ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType = "document"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionDocumentObjectTypeDocument:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject struct {
	Type     param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType]          `json:"type,required"`
	VideoURL param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion] `json:"video_url,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObject) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectTypeVideoURL ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType = "video_url"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectTypeVideoURL:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject].
type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject struct {
	URL    param.Field[string] `json:"url,required"`
	Detail param.Field[string] `json:"detail"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLChatCompletionVideoURLObject) ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject struct {
	File param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile] `json:"file,required"`
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObject) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile struct {
	FileData param.Field[string] `json:"file_data"`
	FileID   param.Field[string] `json:"file_id"`
	Filename param.Field[string] `json:"filename"`
	Format   param.Field[string] `json:"format"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectTypeFile ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType = "file"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionFileObjectTypeFile:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeText       ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType = "text"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeImageURL   ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType = "image_url"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeInputAudio ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType = "input_audio"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeDocument   ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType = "document"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeVideoURL   ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType = "video_url"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeFile       ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType = "file"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeText, ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeImageURL, ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeInputAudio, ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeDocument, ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeVideoURL, ChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayTypeFile:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageRoleUser ChatCompletionNewParamsMessagesChatCompletionUserMessageRole = "user"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageRoleUser:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControl struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControlType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControlType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControlTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControlType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControlType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessage struct {
	Role             param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageRole]                 `json:"role,required"`
	CacheControl     param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControl]         `json:"cache_control"`
	Content          param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentUnion]         `json:"content"`
	FunctionCall     param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageFunctionCall]         `json:"function_call"`
	Name             param.Field[string]                                                                            `json:"name"`
	ReasoningContent param.Field[string]                                                                            `json:"reasoning_content"`
	ThinkingBlocks   param.Field[[]ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion] `json:"thinking_blocks"`
	ToolCalls        param.Field[[]ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCall]           `json:"tool_calls"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessage) implementsChatCompletionNewParamsMessageUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageRoleAssistant ChatCompletionNewParamsMessagesChatCompletionAssistantMessageRole = "assistant"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageRoleAssistant:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControl struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControlType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControlType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControlTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControlType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControlType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArray].
type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArray []ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArray) ImplementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItem struct {
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                   `json:"cache_control"`
	Signature    param.Field[string]                                                                        `json:"signature"`
	Text         param.Field[string]                                                                        `json:"text"`
	Thinking     param.Field[string]                                                                        `json:"thinking"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItem) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

// Satisfied by
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItem].
type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion interface {
	implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject struct {
	Text         param.Field[string]                                                                                                        `json:"text,required"`
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType]         `json:"type,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl] `json:"cache_control"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObject) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectTypeText ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType = "text"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectTypeText:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionTextObjectCacheControlTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock struct {
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion] `json:"cache_control"`
	Signature    param.Field[string]                                                                                                                `json:"signature"`
	Thinking     param.Field[string]                                                                                                                `json:"thinking"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlock) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockTypeThinking ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType = "thinking"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockTypeThinking:
		return true
	}
	return false
}

// Satisfied by
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent].
type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion interface {
	implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap map[string]interface{}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlMap) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayTypeText     ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayType = "text"
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayTypeThinking ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayType = "thinking"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayTypeText, ChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentArrayTypeThinking:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageFunctionCall struct {
	Arguments              param.Field[string]                 `json:"arguments"`
	Name                   param.Field[string]                 `json:"name"`
	ProviderSpecificFields param.Field[map[string]interface{}] `json:"provider_specific_fields"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlock struct {
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                     `json:"cache_control"`
	Data         param.Field[string]                                                                          `json:"data"`
	Signature    param.Field[string]                                                                          `json:"signature"`
	Thinking     param.Field[string]                                                                          `json:"thinking"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlock) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

// Satisfied by
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlock].
type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion interface {
	implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock struct {
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion] `json:"cache_control"`
	Signature    param.Field[string]                                                                                                                  `json:"signature"`
	Thinking     param.Field[string]                                                                                                                  `json:"thinking"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlock) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType = "thinking"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockTypeThinking:
		return true
	}
	return false
}

// Satisfied by
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent].
type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion interface {
	implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap map[string]interface{}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlMap) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContent) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock struct {
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType]              `json:"type,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion] `json:"cache_control"`
	Data         param.Field[string]                                                                                                                          `json:"data"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlock) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlockUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType = "redacted_thinking"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

// Satisfied by
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent].
type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion interface {
	implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap map[string]interface{}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlMap) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContent) implementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksChatCompletionRedactedThinkingBlockCacheControlChatCompletionCachedContentTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeThinking         ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksType = "thinking"
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeRedactedThinking ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksType = "redacted_thinking"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeThinking, ChatCompletionNewParamsMessagesChatCompletionAssistantMessageThinkingBlocksTypeRedactedThinking:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCall struct {
	ID       param.Field[string]                                                                         `json:"id,required"`
	Function param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsFunction] `json:"function,required"`
	Type     param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsType]     `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsFunction struct {
	Arguments              param.Field[string]                 `json:"arguments"`
	Name                   param.Field[string]                 `json:"name"`
	ProviderSpecificFields param.Field[map[string]interface{}] `json:"provider_specific_fields"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsType string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsTypeFunction ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsType = "function"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageToolCallsTypeFunction:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessage struct {
	Content    param.Field[ChatCompletionNewParamsMessagesChatCompletionToolMessageContentUnion] `json:"content,required"`
	Role       param.Field[ChatCompletionNewParamsMessagesChatCompletionToolMessageRole]         `json:"role,required"`
	ToolCallID param.Field[string]                                                               `json:"tool_call_id,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessage) implementsChatCompletionNewParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArray].
type ChatCompletionNewParamsMessagesChatCompletionToolMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionToolMessageContentUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArray []ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayItem

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArray) ImplementsChatCompletionNewParamsMessagesChatCompletionToolMessageContentUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayItem struct {
	Text         param.Field[string]                                                                           `json:"text,required"`
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayType]         `json:"type,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControl] `json:"cache_control"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayType string

const (
	ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayTypeText ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayType = "text"
)

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayTypeText:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControl struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControlType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControlType string

const (
	ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControlTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControlType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControlType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionToolMessageContentArrayCacheControlTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessageRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionToolMessageRoleTool ChatCompletionNewParamsMessagesChatCompletionToolMessageRole = "tool"
)

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionToolMessageRoleTool:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionSystemMessage struct {
	Content      param.Field[ChatCompletionNewParamsMessagesChatCompletionSystemMessageContentUnion] `json:"content,required"`
	Role         param.Field[ChatCompletionNewParamsMessagesChatCompletionSystemMessageRole]         `json:"role,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControl] `json:"cache_control"`
	Name         param.Field[string]                                                                 `json:"name"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessage) implementsChatCompletionNewParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionSystemMessageContentArray].
type ChatCompletionNewParamsMessagesChatCompletionSystemMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionSystemMessageContentUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionSystemMessageContentArray []interface{}

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessageContentArray) ImplementsChatCompletionNewParamsMessagesChatCompletionSystemMessageContentUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionSystemMessageRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionSystemMessageRoleSystem ChatCompletionNewParamsMessagesChatCompletionSystemMessageRole = "system"
)

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionSystemMessageRoleSystem:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControl struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControlType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControlType string

const (
	ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControlTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControlType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControlType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionSystemMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionFunctionMessage struct {
	Content    param.Field[ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentUnion] `json:"content,required"`
	Name       param.Field[string]                                                                   `json:"name,required"`
	Role       param.Field[ChatCompletionNewParamsMessagesChatCompletionFunctionMessageRole]         `json:"role,required"`
	ToolCallID param.Field[string]                                                                   `json:"tool_call_id,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessage) implementsChatCompletionNewParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArray].
type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArray []ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayItem

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArray) ImplementsChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayItem struct {
	Text         param.Field[string]                                                                               `json:"text,required"`
	Type         param.Field[ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayType]         `json:"type,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl] `json:"cache_control"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayType string

const (
	ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayTypeText ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayType = "text"
)

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayTypeText:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType string

const (
	ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentArrayCacheControlTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionFunctionMessageRoleFunction ChatCompletionNewParamsMessagesChatCompletionFunctionMessageRole = "function"
)

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionFunctionMessageRoleFunction:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionDeveloperMessage struct {
	Content      param.Field[ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageContentUnion] `json:"content,required"`
	Role         param.Field[ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageRole]         `json:"role,required"`
	CacheControl param.Field[ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControl] `json:"cache_control"`
	Name         param.Field[string]                                                                    `json:"name"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionDeveloperMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionDeveloperMessage) implementsChatCompletionNewParamsMessageUnion() {
}

// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageContentArray].
type ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionDeveloperMessageContentUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageContentArray []interface{}

func (r ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageContentArray) ImplementsChatCompletionNewParamsMessagesChatCompletionDeveloperMessageContentUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageRoleDeveloper ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageRole = "developer"
)

func (r ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageRoleDeveloper:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControl struct {
	Type param.Field[ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControlType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControlType string

const (
	ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControlTypeEphemeral ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControlType = "ephemeral"
)

func (r ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControlType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionDeveloperMessageCacheControlTypeEphemeral:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesRole string

const (
	ChatCompletionNewParamsMessagesRoleUser      ChatCompletionNewParamsMessagesRole = "user"
	ChatCompletionNewParamsMessagesRoleAssistant ChatCompletionNewParamsMessagesRole = "assistant"
	ChatCompletionNewParamsMessagesRoleTool      ChatCompletionNewParamsMessagesRole = "tool"
	ChatCompletionNewParamsMessagesRoleSystem    ChatCompletionNewParamsMessagesRole = "system"
	ChatCompletionNewParamsMessagesRoleFunction  ChatCompletionNewParamsMessagesRole = "function"
	ChatCompletionNewParamsMessagesRoleDeveloper ChatCompletionNewParamsMessagesRole = "developer"
)

func (r ChatCompletionNewParamsMessagesRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesRoleUser, ChatCompletionNewParamsMessagesRoleAssistant, ChatCompletionNewParamsMessagesRoleTool, ChatCompletionNewParamsMessagesRoleSystem, ChatCompletionNewParamsMessagesRoleFunction, ChatCompletionNewParamsMessagesRoleDeveloper:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [ChatCompletionNewParamsFunctionCallMap].
type ChatCompletionNewParamsFunctionCallUnion interface {
	ImplementsChatCompletionNewParamsFunctionCallUnion()
}

type ChatCompletionNewParamsFunctionCallMap map[string]interface{}

func (r ChatCompletionNewParamsFunctionCallMap) ImplementsChatCompletionNewParamsFunctionCallUnion() {
}

// Satisfied by [shared.UnionString], [ChatCompletionNewParamsStopArray].
type ChatCompletionNewParamsStopUnion interface {
	ImplementsChatCompletionNewParamsStopUnion()
}

type ChatCompletionNewParamsStopArray []string

func (r ChatCompletionNewParamsStopArray) ImplementsChatCompletionNewParamsStopUnion() {}

// Satisfied by [shared.UnionString], [ChatCompletionNewParamsToolChoiceMap].
type ChatCompletionNewParamsToolChoiceUnion interface {
	ImplementsChatCompletionNewParamsToolChoiceUnion()
}

type ChatCompletionNewParamsToolChoiceMap map[string]interface{}

func (r ChatCompletionNewParamsToolChoiceMap) ImplementsChatCompletionNewParamsToolChoiceUnion() {}
