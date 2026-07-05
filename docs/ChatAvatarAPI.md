# \ChatAvatarAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatPostAssistantsV2AvatarByassistantId**](ChatAvatarAPI.md#ChatPostAssistantsV2AvatarByassistantId) | **Post** /v1/chat/assistants/v2/avatar/{assistant_id} | Upload assistant avatar (v2)
[**ChatPostFilesImagesAgentsByagentIdAvatar**](ChatAvatarAPI.md#ChatPostFilesImagesAgentsByagentIdAvatar) | **Post** /v1/chat/files/images/agents/{agent_id}/avatar | Upload agent avatar
[**ChatPostFilesImagesAssistantsByassistantIdAvatar**](ChatAvatarAPI.md#ChatPostFilesImagesAssistantsByassistantIdAvatar) | **Post** /v1/chat/files/images/assistants/{assistant_id}/avatar | Upload assistant avatar (v1)
[**ChatPostFilesImagesAvatar**](ChatAvatarAPI.md#ChatPostFilesImagesAvatar) | **Post** /v1/chat/files/images/avatar | Upload user avatar



## ChatPostAssistantsV2AvatarByassistantId

> map[string]interface{} ChatPostAssistantsV2AvatarByassistantId(ctx, assistantId).File(file).Metadata(metadata).Execute()

Upload assistant avatar (v2)

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
	assistantId := "assistantId_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	metadata := "metadata_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAvatarAPI.ChatPostAssistantsV2AvatarByassistantId(context.Background(), assistantId).File(file).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAvatarAPI.ChatPostAssistantsV2AvatarByassistantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV2AvatarByassistantId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAvatarAPI.ChatPostAssistantsV2AvatarByassistantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV2AvatarByassistantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **metadata** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostFilesImagesAgentsByagentIdAvatar

> map[string]interface{} ChatPostFilesImagesAgentsByagentIdAvatar(ctx, agentId).File(file).Metadata(metadata).Execute()

Upload agent avatar

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
	agentId := "agentId_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 
	metadata := "metadata_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAvatarAPI.ChatPostFilesImagesAgentsByagentIdAvatar(context.Background(), agentId).File(file).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAvatarAPI.ChatPostFilesImagesAgentsByagentIdAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostFilesImagesAgentsByagentIdAvatar`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAvatarAPI.ChatPostFilesImagesAgentsByagentIdAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostFilesImagesAgentsByagentIdAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **metadata** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostFilesImagesAssistantsByassistantIdAvatar

> map[string]interface{} ChatPostFilesImagesAssistantsByassistantIdAvatar(ctx, assistantId).File(file).Metadata(metadata).Execute()

Upload assistant avatar (v1)

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
	assistantId := "assistantId_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 
	metadata := "metadata_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAvatarAPI.ChatPostFilesImagesAssistantsByassistantIdAvatar(context.Background(), assistantId).File(file).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAvatarAPI.ChatPostFilesImagesAssistantsByassistantIdAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostFilesImagesAssistantsByassistantIdAvatar`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAvatarAPI.ChatPostFilesImagesAssistantsByassistantIdAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostFilesImagesAssistantsByassistantIdAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **metadata** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostFilesImagesAvatar

> ChatPostFilesImagesAvatar200Response ChatPostFilesImagesAvatar(ctx).File(file).Manual(manual).Execute()

Upload user avatar

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
	manual := "manual_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAvatarAPI.ChatPostFilesImagesAvatar(context.Background()).File(file).Manual(manual).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAvatarAPI.ChatPostFilesImagesAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostFilesImagesAvatar`: ChatPostFilesImagesAvatar200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAvatarAPI.ChatPostFilesImagesAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostFilesImagesAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **manual** | **string** |  | 

### Return type

[**ChatPostFilesImagesAvatar200Response**](ChatPostFilesImagesAvatar200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

