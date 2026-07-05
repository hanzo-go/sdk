# \AiOpenAICompatibleAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiCreateChatCompletion**](AiOpenAICompatibleAPI.md#AiCreateChatCompletion) | **Post** /v1/chat/completions | Create chat completion
[**AiCreateCompletion**](AiOpenAICompatibleAPI.md#AiCreateCompletion) | **Post** /v1/completions | Create completion (legacy text)
[**AiCreateEmbeddings**](AiOpenAICompatibleAPI.md#AiCreateEmbeddings) | **Post** /v1/embeddings | Create embeddings
[**AiCreateImage**](AiOpenAICompatibleAPI.md#AiCreateImage) | **Post** /v1/images/generations | Create image
[**AiCreateSpeech**](AiOpenAICompatibleAPI.md#AiCreateSpeech) | **Post** /v1/audio/speech | Create speech (text-to-speech)
[**AiCreateTranscription**](AiOpenAICompatibleAPI.md#AiCreateTranscription) | **Post** /v1/audio/transcriptions | Create transcription (speech-to-text)
[**AiRerank**](AiOpenAICompatibleAPI.md#AiRerank) | **Post** /v1/rerank | Rerank documents against a query



## AiCreateChatCompletion

> AiChatCompletionResponse AiCreateChatCompletion(ctx).AiChatCompletionRequest(aiChatCompletionRequest).Execute()

Create chat completion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	aiChatCompletionRequest := *openapiclient.NewAiChatCompletionRequest("Model_example", []openapiclient.AiChatMessage{*openapiclient.NewAiChatMessage("Role_example")}) // AiChatCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiOpenAICompatibleAPI.AiCreateChatCompletion(context.Background()).AiChatCompletionRequest(aiChatCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiOpenAICompatibleAPI.AiCreateChatCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCreateChatCompletion`: AiChatCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `AiOpenAICompatibleAPI.AiCreateChatCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCreateChatCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiChatCompletionRequest** | [**AiChatCompletionRequest**](AiChatCompletionRequest.md) |  | 

### Return type

[**AiChatCompletionResponse**](AiChatCompletionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiCreateCompletion

> map[string]interface{} AiCreateCompletion(ctx).AiCompletionRequest(aiCompletionRequest).Execute()

Create completion (legacy text)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	aiCompletionRequest := *openapiclient.NewAiCompletionRequest("Model_example", interface{}(123)) // AiCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiOpenAICompatibleAPI.AiCreateCompletion(context.Background()).AiCompletionRequest(aiCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiOpenAICompatibleAPI.AiCreateCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCreateCompletion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AiOpenAICompatibleAPI.AiCreateCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCreateCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiCompletionRequest** | [**AiCompletionRequest**](AiCompletionRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiCreateEmbeddings

> AiEmbeddingResponse AiCreateEmbeddings(ctx).AiEmbeddingRequest(aiEmbeddingRequest).Execute()

Create embeddings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	aiEmbeddingRequest := *openapiclient.NewAiEmbeddingRequest("Model_example", interface{}(123)) // AiEmbeddingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiOpenAICompatibleAPI.AiCreateEmbeddings(context.Background()).AiEmbeddingRequest(aiEmbeddingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiOpenAICompatibleAPI.AiCreateEmbeddings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCreateEmbeddings`: AiEmbeddingResponse
	fmt.Fprintf(os.Stdout, "Response from `AiOpenAICompatibleAPI.AiCreateEmbeddings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCreateEmbeddingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiEmbeddingRequest** | [**AiEmbeddingRequest**](AiEmbeddingRequest.md) |  | 

### Return type

[**AiEmbeddingResponse**](AiEmbeddingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiCreateImage

> AiImageResponse AiCreateImage(ctx).AiImageGenerationRequest(aiImageGenerationRequest).Execute()

Create image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	aiImageGenerationRequest := *openapiclient.NewAiImageGenerationRequest("Prompt_example") // AiImageGenerationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiOpenAICompatibleAPI.AiCreateImage(context.Background()).AiImageGenerationRequest(aiImageGenerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiOpenAICompatibleAPI.AiCreateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCreateImage`: AiImageResponse
	fmt.Fprintf(os.Stdout, "Response from `AiOpenAICompatibleAPI.AiCreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiImageGenerationRequest** | [**AiImageGenerationRequest**](AiImageGenerationRequest.md) |  | 

### Return type

[**AiImageResponse**](AiImageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiCreateSpeech

> *os.File AiCreateSpeech(ctx).AiSpeechRequest(aiSpeechRequest).Execute()

Create speech (text-to-speech)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	aiSpeechRequest := *openapiclient.NewAiSpeechRequest("Model_example", "Input_example", "Voice_example") // AiSpeechRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiOpenAICompatibleAPI.AiCreateSpeech(context.Background()).AiSpeechRequest(aiSpeechRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiOpenAICompatibleAPI.AiCreateSpeech``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCreateSpeech`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AiOpenAICompatibleAPI.AiCreateSpeech`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCreateSpeechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSpeechRequest** | [**AiSpeechRequest**](AiSpeechRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiCreateTranscription

> AiTranscriptionResponse AiCreateTranscription(ctx).File(file).Model(model).Language(language).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).Execute()

Create transcription (speech-to-text)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | 
	model := "model_example" // string | 
	language := "language_example" // string |  (optional)
	prompt := "prompt_example" // string |  (optional)
	responseFormat := "responseFormat_example" // string |  (optional)
	temperature := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiOpenAICompatibleAPI.AiCreateTranscription(context.Background()).File(file).Model(model).Language(language).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiOpenAICompatibleAPI.AiCreateTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCreateTranscription`: AiTranscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `AiOpenAICompatibleAPI.AiCreateTranscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCreateTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **model** | **string** |  | 
 **language** | **string** |  | 
 **prompt** | **string** |  | 
 **responseFormat** | **string** |  | 
 **temperature** | **float32** |  | 

### Return type

[**AiTranscriptionResponse**](AiTranscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiRerank

> map[string]interface{} AiRerank(ctx).AiRerankRequest(aiRerankRequest).Execute()

Rerank documents against a query

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	aiRerankRequest := *openapiclient.NewAiRerankRequest("Model_example", "Query_example", []string{"Documents_example"}) // AiRerankRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiOpenAICompatibleAPI.AiRerank(context.Background()).AiRerankRequest(aiRerankRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiOpenAICompatibleAPI.AiRerank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiRerank`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AiOpenAICompatibleAPI.AiRerank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiRerankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiRerankRequest** | [**AiRerankRequest**](AiRerankRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

