# \ChatAgentsOpenResponsesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetAgentsV1ResponsesByid**](ChatAgentsOpenResponsesAPI.md#ChatGetAgentsV1ResponsesByid) | **Get** /v1/chat/agents/v1/responses/{id} | Get a stored response
[**ChatGetAgentsV1ResponsesModels**](ChatAgentsOpenResponsesAPI.md#ChatGetAgentsV1ResponsesModels) | **Get** /v1/chat/agents/v1/responses/models | List agents as models
[**ChatPostAgentsV1Responses**](ChatAgentsOpenResponsesAPI.md#ChatPostAgentsV1Responses) | **Post** /v1/chat/agents/v1/responses | Create a response



## ChatGetAgentsV1ResponsesByid

> ChatResponseObject ChatGetAgentsV1ResponsesByid(ctx, id).Execute()

Get a stored response

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsOpenResponsesAPI.ChatGetAgentsV1ResponsesByid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsOpenResponsesAPI.ChatGetAgentsV1ResponsesByid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsV1ResponsesByid`: ChatResponseObject
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsOpenResponsesAPI.ChatGetAgentsV1ResponsesByid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsV1ResponsesByidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatResponseObject**](ChatResponseObject.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsV1ResponsesModels

> map[string]interface{} ChatGetAgentsV1ResponsesModels(ctx).Execute()

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
	resp, r, err := apiClient.ChatAgentsOpenResponsesAPI.ChatGetAgentsV1ResponsesModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsOpenResponsesAPI.ChatGetAgentsV1ResponsesModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsV1ResponsesModels`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsOpenResponsesAPI.ChatGetAgentsV1ResponsesModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsV1ResponsesModelsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgentsV1Responses

> ChatResponseObject ChatPostAgentsV1Responses(ctx).ChatPostAgentsV1ResponsesRequest(chatPostAgentsV1ResponsesRequest).Execute()

Create a response



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
	chatPostAgentsV1ResponsesRequest := *openapiclient.NewChatPostAgentsV1ResponsesRequest("Model_example", openapiclient.chat_postAgentsV1Responses_request_input{ArrayOfMapmapOfStringAny: new([]map[string]interface{})}) // ChatPostAgentsV1ResponsesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsOpenResponsesAPI.ChatPostAgentsV1Responses(context.Background()).ChatPostAgentsV1ResponsesRequest(chatPostAgentsV1ResponsesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsOpenResponsesAPI.ChatPostAgentsV1Responses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsV1Responses`: ChatResponseObject
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsOpenResponsesAPI.ChatPostAgentsV1Responses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsV1ResponsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAgentsV1ResponsesRequest** | [**ChatPostAgentsV1ResponsesRequest**](ChatPostAgentsV1ResponsesRequest.md) |  | 

### Return type

[**ChatResponseObject**](ChatResponseObject.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

