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

// EmbeddingService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddingService] method instead.
type EmbeddingService struct {
	Options []option.RequestOption
}

// NewEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddingService(opts ...option.RequestOption) (r *EmbeddingService) {
	r = &EmbeddingService{}
	r.Options = opts
	return
}

// Follows the exact same API spec as
// `OpenAI's Embeddings API https://platform.openai.com/docs/api-reference/embeddings`
//
// ```bash
// curl -X POST http://localhost:4000/v1/embeddings
// -H "Content-Type: application/json"
// -H "Authorization: Bearer sk-1234"
//
//	-d '{
//	    "model": "text-embedding-ada-002",
//	    "input": "The quick brown fox jumps over the lazy dog"
//	}'
//
// ```
func (r *EmbeddingService) New(ctx context.Context, body EmbeddingNewParams, opts ...option.RequestOption) (res *EmbeddingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EmbeddingNewResponse = interface{}

type EmbeddingNewParams struct {
	Model             param.Field[string]                                   `json:"model,required"`
	APIBase           param.Field[string]                                   `json:"api_base"`
	APIKey            param.Field[string]                                   `json:"api_key"`
	APIType           param.Field[string]                                   `json:"api_type"`
	APIVersion        param.Field[string]                                   `json:"api_version"`
	Caching           param.Field[bool]                                     `json:"caching"`
	CustomLlmProvider param.Field[EmbeddingNewParamsCustomLlmProviderUnion] `json:"custom_llm_provider"`
	Input             param.Field[[]string]                                 `json:"input"`
	LitellmCallID     param.Field[string]                                   `json:"litellm_call_id"`
	LitellmLoggingObj param.Field[map[string]interface{}]                   `json:"litellm_logging_obj"`
	LoggerFn          param.Field[string]                                   `json:"logger_fn"`
	Timeout           param.Field[int64]                                    `json:"timeout"`
	User              param.Field[string]                                   `json:"user"`
}

func (r EmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [EmbeddingNewParamsCustomLlmProviderMap].
type EmbeddingNewParamsCustomLlmProviderUnion interface {
	ImplementsEmbeddingNewParamsCustomLlmProviderUnion()
}

type EmbeddingNewParamsCustomLlmProviderMap map[string]interface{}

func (r EmbeddingNewParamsCustomLlmProviderMap) ImplementsEmbeddingNewParamsCustomLlmProviderUnion() {
}
