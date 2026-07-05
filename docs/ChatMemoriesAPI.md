# \ChatMemoriesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteMemoriesBykey**](ChatMemoriesAPI.md#ChatDeleteMemoriesBykey) | **Delete** /v1/chat/memories/{key} | Delete a memory
[**ChatGetMemories**](ChatMemoriesAPI.md#ChatGetMemories) | **Get** /v1/chat/memories | Get all user memories
[**ChatPatchMemoriesBykey**](ChatMemoriesAPI.md#ChatPatchMemoriesBykey) | **Patch** /v1/chat/memories/{key} | Update a memory
[**ChatPatchMemoriesPreferences**](ChatMemoriesAPI.md#ChatPatchMemoriesPreferences) | **Patch** /v1/chat/memories/preferences | Update memory preferences
[**ChatPostMemories**](ChatMemoriesAPI.md#ChatPostMemories) | **Post** /v1/chat/memories | Create a memory



## ChatDeleteMemoriesBykey

> map[string]interface{} ChatDeleteMemoriesBykey(ctx, key).Execute()

Delete a memory

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMemoriesAPI.ChatDeleteMemoriesBykey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMemoriesAPI.ChatDeleteMemoriesBykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteMemoriesBykey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMemoriesAPI.ChatDeleteMemoriesBykey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteMemoriesBykeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMemories

> ChatGetMemories200Response ChatGetMemories(ctx).Execute()

Get all user memories

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
	resp, r, err := apiClient.ChatMemoriesAPI.ChatGetMemories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMemoriesAPI.ChatGetMemories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMemories`: ChatGetMemories200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMemoriesAPI.ChatGetMemories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMemoriesRequest struct via the builder pattern


### Return type

[**ChatGetMemories200Response**](ChatGetMemories200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPatchMemoriesBykey

> map[string]interface{} ChatPatchMemoriesBykey(ctx, key).ChatPatchMemoriesBykeyRequest(chatPatchMemoriesBykeyRequest).Execute()

Update a memory

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
	key := "key_example" // string | 
	chatPatchMemoriesBykeyRequest := *openapiclient.NewChatPatchMemoriesBykeyRequest("Value_example") // ChatPatchMemoriesBykeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMemoriesAPI.ChatPatchMemoriesBykey(context.Background(), key).ChatPatchMemoriesBykeyRequest(chatPatchMemoriesBykeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMemoriesAPI.ChatPatchMemoriesBykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchMemoriesBykey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMemoriesAPI.ChatPatchMemoriesBykey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchMemoriesBykeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatPatchMemoriesBykeyRequest** | [**ChatPatchMemoriesBykeyRequest**](ChatPatchMemoriesBykeyRequest.md) |  | 

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


## ChatPatchMemoriesPreferences

> map[string]interface{} ChatPatchMemoriesPreferences(ctx).ChatPatchMemoriesPreferencesRequest(chatPatchMemoriesPreferencesRequest).Execute()

Update memory preferences

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
	chatPatchMemoriesPreferencesRequest := *openapiclient.NewChatPatchMemoriesPreferencesRequest(false) // ChatPatchMemoriesPreferencesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMemoriesAPI.ChatPatchMemoriesPreferences(context.Background()).ChatPatchMemoriesPreferencesRequest(chatPatchMemoriesPreferencesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMemoriesAPI.ChatPatchMemoriesPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchMemoriesPreferences`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMemoriesAPI.ChatPatchMemoriesPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchMemoriesPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPatchMemoriesPreferencesRequest** | [**ChatPatchMemoriesPreferencesRequest**](ChatPatchMemoriesPreferencesRequest.md) |  | 

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


## ChatPostMemories

> ChatPostMemories201Response ChatPostMemories(ctx).ChatPostMemoriesRequest(chatPostMemoriesRequest).Execute()

Create a memory

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
	chatPostMemoriesRequest := *openapiclient.NewChatPostMemoriesRequest("Key_example", "Value_example") // ChatPostMemoriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMemoriesAPI.ChatPostMemories(context.Background()).ChatPostMemoriesRequest(chatPostMemoriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMemoriesAPI.ChatPostMemories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostMemories`: ChatPostMemories201Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMemoriesAPI.ChatPostMemories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostMemoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostMemoriesRequest** | [**ChatPostMemoriesRequest**](ChatPostMemoriesRequest.md) |  | 

### Return type

[**ChatPostMemories201Response**](ChatPostMemories201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

