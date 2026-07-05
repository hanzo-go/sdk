# \NexusChatAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddChat**](NexusChatAPIAPI.md#NexusAddChat) | **Post** /v1/nexus/add-chat | add Chat
[**NexusDeleteChat**](NexusChatAPIAPI.md#NexusDeleteChat) | **Post** /v1/nexus/delete-chat | delete Chat
[**NexusGetChat**](NexusChatAPIAPI.md#NexusGetChat) | **Get** /v1/nexus/get-chat | get Chat
[**NexusGetChats**](NexusChatAPIAPI.md#NexusGetChats) | **Get** /v1/nexus/get-chats | get Chats
[**NexusGetGlobalChats**](NexusChatAPIAPI.md#NexusGetGlobalChats) | **Get** /v1/nexus/get-global-chats | get Global Chats
[**NexusUpdateChat**](NexusChatAPIAPI.md#NexusUpdateChat) | **Post** /v1/nexus/update-chat | update Chat



## NexusAddChat

> NexusResponse NexusAddChat(ctx).NexusChat(nexusChat).Execute()

add Chat



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
	nexusChat := *openapiclient.NewNexusChat() // NexusChat | The details of the chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusChatAPIAPI.NexusAddChat(context.Background()).NexusChat(nexusChat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusChatAPIAPI.NexusAddChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddChat`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusChatAPIAPI.NexusAddChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusChat** | [**NexusChat**](NexusChat.md) | The details of the chat | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteChat

> NexusResponse NexusDeleteChat(ctx).NexusChat(nexusChat).Execute()

delete Chat



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
	nexusChat := *openapiclient.NewNexusChat() // NexusChat | The details of the chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusChatAPIAPI.NexusDeleteChat(context.Background()).NexusChat(nexusChat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusChatAPIAPI.NexusDeleteChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteChat`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusChatAPIAPI.NexusDeleteChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusChat** | [**NexusChat**](NexusChat.md) | The details of the chat | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetChat

> NexusChat NexusGetChat(ctx).Id(id).Execute()

get Chat



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
	id := "id_example" // string | The id of the chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusChatAPIAPI.NexusGetChat(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusChatAPIAPI.NexusGetChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetChat`: NexusChat
	fmt.Fprintf(os.Stdout, "Response from `NexusChatAPIAPI.NexusGetChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the chat | 

### Return type

[**NexusChat**](NexusChat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetChats

> []NexusChat NexusGetChats(ctx).User(user).Field(field).Value(value).Execute()

get Chats



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
	user := "user_example" // string | The user of the chats
	field := "field_example" // string | The field to filter by
	value := "value_example" // string | The value to filter by

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusChatAPIAPI.NexusGetChats(context.Background()).User(user).Field(field).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusChatAPIAPI.NexusGetChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetChats`: []NexusChat
	fmt.Fprintf(os.Stdout, "Response from `NexusChatAPIAPI.NexusGetChats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | The user of the chats | 
 **field** | **string** | The field to filter by | 
 **value** | **string** | The value to filter by | 

### Return type

[**[]NexusChat**](NexusChat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalChats

> []NexusChat NexusGetGlobalChats(ctx).Execute()

get Global Chats



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
	resp, r, err := apiClient.NexusChatAPIAPI.NexusGetGlobalChats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusChatAPIAPI.NexusGetGlobalChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalChats`: []NexusChat
	fmt.Fprintf(os.Stdout, "Response from `NexusChatAPIAPI.NexusGetGlobalChats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalChatsRequest struct via the builder pattern


### Return type

[**[]NexusChat**](NexusChat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateChat

> NexusResponse NexusUpdateChat(ctx).Id(id).NexusChat(nexusChat).Execute()

update Chat



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
	nexusChat := *openapiclient.NewNexusChat() // NexusChat | The details of the chat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusChatAPIAPI.NexusUpdateChat(context.Background()).Id(id).NexusChat(nexusChat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusChatAPIAPI.NexusUpdateChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateChat`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusChatAPIAPI.NexusUpdateChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the chat | 
 **nexusChat** | [**NexusChat**](NexusChat.md) | The details of the chat | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

