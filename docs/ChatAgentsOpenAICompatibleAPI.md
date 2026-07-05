# \ChatAgentsOpenAICompatibleAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetAgentsV1Models**](ChatAgentsOpenAICompatibleAPI.md#ChatGetAgentsV1Models) | **Get** /v1/chat/agents/v1/models | List agents as models
[**ChatGetAgentsV1ModelsBymodel**](ChatAgentsOpenAICompatibleAPI.md#ChatGetAgentsV1ModelsBymodel) | **Get** /v1/chat/agents/v1/models/{model} | Get agent/model details
[**ChatPostAgentsV1ChatCompletions**](ChatAgentsOpenAICompatibleAPI.md#ChatPostAgentsV1ChatCompletions) | **Post** /v1/chat/agents/v1/chat/completions | OpenAI-compatible chat completions



## ChatGetAgentsV1Models

> ChatGetAgentsV1Models200Response ChatGetAgentsV1Models(ctx).Execute()

List agents as models

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsOpenAICompatibleAPI.ChatGetAgentsV1Models(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsOpenAICompatibleAPI.ChatGetAgentsV1Models``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsV1Models`: ChatGetAgentsV1Models200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsOpenAICompatibleAPI.ChatGetAgentsV1Models`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsV1ModelsRequest struct via the builder pattern


### Return type

[**ChatGetAgentsV1Models200Response**](ChatGetAgentsV1Models200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsV1ModelsBymodel

> ChatModelObject ChatGetAgentsV1ModelsBymodel(ctx, model).Execute()

Get agent/model details

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
	model := "model_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsOpenAICompatibleAPI.ChatGetAgentsV1ModelsBymodel(context.Background(), model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsOpenAICompatibleAPI.ChatGetAgentsV1ModelsBymodel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsV1ModelsBymodel`: ChatModelObject
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsOpenAICompatibleAPI.ChatGetAgentsV1ModelsBymodel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsV1ModelsBymodelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatModelObject**](ChatModelObject.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgentsV1ChatCompletions

> ChatChatCompletion ChatPostAgentsV1ChatCompletions(ctx).ChatPostAgentsV1ChatCompletionsRequest(chatPostAgentsV1ChatCompletionsRequest).Execute()

OpenAI-compatible chat completions



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
	chatPostAgentsV1ChatCompletionsRequest := *openapiclient.NewChatPostAgentsV1ChatCompletionsRequest("Model_example", []openapiclient.ChatChatMessage{*openapiclient.NewChatChatMessage("Role_example", "Content_example")}) // ChatPostAgentsV1ChatCompletionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsOpenAICompatibleAPI.ChatPostAgentsV1ChatCompletions(context.Background()).ChatPostAgentsV1ChatCompletionsRequest(chatPostAgentsV1ChatCompletionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsOpenAICompatibleAPI.ChatPostAgentsV1ChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsV1ChatCompletions`: ChatChatCompletion
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsOpenAICompatibleAPI.ChatPostAgentsV1ChatCompletions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsV1ChatCompletionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAgentsV1ChatCompletionsRequest** | [**ChatPostAgentsV1ChatCompletionsRequest**](ChatPostAgentsV1ChatCompletionsRequest.md) |  | 

### Return type

[**ChatChatCompletion**](ChatChatCompletion.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

