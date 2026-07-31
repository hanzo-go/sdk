# \ChatAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddChat**](ChatAPIAPI.md#CloudApiControllerAddChat) | **Post** /v1/cloud/add-chat | Api Controller Add Chat
[**CloudApiControllerDeleteChat**](ChatAPIAPI.md#CloudApiControllerDeleteChat) | **Post** /v1/cloud/delete-chat | Api Controller Delete Chat
[**CloudApiControllerGetChat**](ChatAPIAPI.md#CloudApiControllerGetChat) | **Get** /v1/cloud/get-chat | Api Controller Get Chat
[**CloudApiControllerGetChats**](ChatAPIAPI.md#CloudApiControllerGetChats) | **Get** /v1/cloud/get-chats | Api Controller Get Chats
[**CloudApiControllerGetGlobalChats**](ChatAPIAPI.md#CloudApiControllerGetGlobalChats) | **Get** /v1/cloud/get-global-chats | Api Controller Get Global Chats
[**CloudApiControllerUpdateChat**](ChatAPIAPI.md#CloudApiControllerUpdateChat) | **Post** /v1/cloud/update-chat | Api Controller Update Chat



## CloudApiControllerAddChat

> CloudControllersResponse CloudApiControllerAddChat(ctx).CloudObjectChat(cloudObjectChat).Execute()

Api Controller Add Chat



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
	cloudObjectChat := *openapiclient.NewCloudObjectChat() // CloudObjectChat | The details of the chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.CloudApiControllerAddChat(context.Background()).CloudObjectChat(cloudObjectChat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.CloudApiControllerAddChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddChat`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.CloudApiControllerAddChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectChat** | [**CloudObjectChat**](CloudObjectChat.md) | The details of the chat | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteChat

> CloudControllersResponse CloudApiControllerDeleteChat(ctx).CloudObjectChat(cloudObjectChat).Execute()

Api Controller Delete Chat



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
	cloudObjectChat := *openapiclient.NewCloudObjectChat() // CloudObjectChat | The details of the chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.CloudApiControllerDeleteChat(context.Background()).CloudObjectChat(cloudObjectChat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.CloudApiControllerDeleteChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteChat`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.CloudApiControllerDeleteChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectChat** | [**CloudObjectChat**](CloudObjectChat.md) | The details of the chat | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetChat

> CloudObjectChat CloudApiControllerGetChat(ctx).Id(id).Execute()

Api Controller Get Chat



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
	id := "id_example" // string | The id of chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.CloudApiControllerGetChat(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.CloudApiControllerGetChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetChat`: CloudObjectChat
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.CloudApiControllerGetChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of chat | 

### Return type

[**CloudObjectChat**](CloudObjectChat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetChats

> []CloudObjectChat CloudApiControllerGetChats(ctx).User(user).Field(field).Value(value).Execute()

Api Controller Get Chats



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
	user := "user_example" // string | The user of chat
	field := "field_example" // string | The field of chat
	value := "value_example" // string | The value of chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.CloudApiControllerGetChats(context.Background()).User(user).Field(field).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.CloudApiControllerGetChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetChats`: []CloudObjectChat
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.CloudApiControllerGetChats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | The user of chat | 
 **field** | **string** | The field of chat | 
 **value** | **string** | The value of chat | 

### Return type

[**[]CloudObjectChat**](CloudObjectChat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetGlobalChats

> []CloudObjectChat CloudApiControllerGetGlobalChats(ctx).Execute()

Api Controller Get Global Chats



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
	resp, r, err := apiClient.ChatAPIAPI.CloudApiControllerGetGlobalChats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.CloudApiControllerGetGlobalChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalChats`: []CloudObjectChat
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.CloudApiControllerGetGlobalChats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalChatsRequest struct via the builder pattern


### Return type

[**[]CloudObjectChat**](CloudObjectChat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateChat

> CloudControllersResponse CloudApiControllerUpdateChat(ctx).Id(id).CloudObjectChat(cloudObjectChat).Execute()

Api Controller Update Chat



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
	id := "id_example" // string | The id (owner/name) of the chat
	cloudObjectChat := *openapiclient.NewCloudObjectChat() // CloudObjectChat | The details of the chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIAPI.CloudApiControllerUpdateChat(context.Background()).Id(id).CloudObjectChat(cloudObjectChat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIAPI.CloudApiControllerUpdateChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateChat`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIAPI.CloudApiControllerUpdateChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the chat | 
 **cloudObjectChat** | [**CloudObjectChat**](CloudObjectChat.md) | The details of the chat | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

