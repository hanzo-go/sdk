# \NexusMessageAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddMessage**](NexusMessageAPIAPI.md#NexusAddMessage) | **Post** /v1/nexus/add-message | add Message
[**NexusDeleteMessage**](NexusMessageAPIAPI.md#NexusDeleteMessage) | **Post** /v1/nexus/delete-message | delete Message
[**NexusGetAnswer**](NexusMessageAPIAPI.md#NexusGetAnswer) | **Get** /v1/nexus/get-answer | get Answer
[**NexusGetGlobalMessages**](NexusMessageAPIAPI.md#NexusGetGlobalMessages) | **Get** /v1/nexus/get-global-messages | get Global Messages
[**NexusGetMessage**](NexusMessageAPIAPI.md#NexusGetMessage) | **Get** /v1/nexus/get-message | get Message
[**NexusGetMessageAnswer**](NexusMessageAPIAPI.md#NexusGetMessageAnswer) | **Get** /v1/nexus/get-message-answer | get Message Answer
[**NexusGetMessages**](NexusMessageAPIAPI.md#NexusGetMessages) | **Get** /v1/nexus/get-Messages | get Messages
[**NexusUpdateMessage**](NexusMessageAPIAPI.md#NexusUpdateMessage) | **Post** /v1/nexus/update-message | update Message



## NexusAddMessage

> NexusChat NexusAddMessage(ctx).NexusMessage(nexusMessage).Execute()

add Message



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
	nexusMessage := *openapiclient.NewNexusMessage() // NexusMessage | The details of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMessageAPIAPI.NexusAddMessage(context.Background()).NexusMessage(nexusMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMessageAPIAPI.NexusAddMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddMessage`: NexusChat
	fmt.Fprintf(os.Stdout, "Response from `NexusMessageAPIAPI.NexusAddMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusMessage** | [**NexusMessage**](NexusMessage.md) | The details of the message | 

### Return type

[**NexusChat**](NexusChat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteMessage

> NexusResponse NexusDeleteMessage(ctx).NexusMessage(nexusMessage).Execute()

delete Message



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
	nexusMessage := *openapiclient.NewNexusMessage() // NexusMessage | The details of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMessageAPIAPI.NexusDeleteMessage(context.Background()).NexusMessage(nexusMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMessageAPIAPI.NexusDeleteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteMessage`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusMessageAPIAPI.NexusDeleteMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusMessage** | [**NexusMessage**](NexusMessage.md) | The details of the message | 

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


## NexusGetAnswer

> string NexusGetAnswer(ctx).Provider(provider).Question(question).Framework(framework).Video(video).Execute()

get Answer



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
	provider := "provider_example" // string | The provider
	question := "question_example" // string | The question
	framework := "framework_example" // string | The framework
	video := "video_example" // string | The video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMessageAPIAPI.NexusGetAnswer(context.Background()).Provider(provider).Question(question).Framework(framework).Video(video).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMessageAPIAPI.NexusGetAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetAnswer`: string
	fmt.Fprintf(os.Stdout, "Response from `NexusMessageAPIAPI.NexusGetAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | The provider | 
 **question** | **string** | The question | 
 **framework** | **string** | The framework | 
 **video** | **string** | The video | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalMessages

> []NexusMessage NexusGetGlobalMessages(ctx).Execute()

get Global Messages



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
	resp, r, err := apiClient.NexusMessageAPIAPI.NexusGetGlobalMessages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMessageAPIAPI.NexusGetGlobalMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalMessages`: []NexusMessage
	fmt.Fprintf(os.Stdout, "Response from `NexusMessageAPIAPI.NexusGetGlobalMessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalMessagesRequest struct via the builder pattern


### Return type

[**[]NexusMessage**](NexusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetMessage

> NexusMessage NexusGetMessage(ctx).Id(id).Execute()

get Message



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
	id := "id_example" // string | The id of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMessageAPIAPI.NexusGetMessage(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMessageAPIAPI.NexusGetMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetMessage`: NexusMessage
	fmt.Fprintf(os.Stdout, "Response from `NexusMessageAPIAPI.NexusGetMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the message | 

### Return type

[**NexusMessage**](NexusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetMessageAnswer

> string NexusGetMessageAnswer(ctx).Id(id).Execute()

get Message Answer



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
	id := "id_example" // string | The id of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMessageAPIAPI.NexusGetMessageAnswer(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMessageAPIAPI.NexusGetMessageAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetMessageAnswer`: string
	fmt.Fprintf(os.Stdout, "Response from `NexusMessageAPIAPI.NexusGetMessageAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetMessageAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the message | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetMessages

> []NexusMessage NexusGetMessages(ctx).User(user).Chat(chat).Execute()

get Messages



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
	user := "user_example" // string | The user of the messages
	chat := "chat_example" // string | The chat of the messages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMessageAPIAPI.NexusGetMessages(context.Background()).User(user).Chat(chat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMessageAPIAPI.NexusGetMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetMessages`: []NexusMessage
	fmt.Fprintf(os.Stdout, "Response from `NexusMessageAPIAPI.NexusGetMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | The user of the messages | 
 **chat** | **string** | The chat of the messages | 

### Return type

[**[]NexusMessage**](NexusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateMessage

> NexusResponse NexusUpdateMessage(ctx).Id(id).NexusMessage(nexusMessage).Execute()

update Message



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
	id := "id_example" // string | The id (owner/name) of the message
	nexusMessage := *openapiclient.NewNexusMessage() // NexusMessage | The details of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMessageAPIAPI.NexusUpdateMessage(context.Background()).Id(id).NexusMessage(nexusMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMessageAPIAPI.NexusUpdateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateMessage`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusMessageAPIAPI.NexusUpdateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the message | 
 **nexusMessage** | [**NexusMessage**](NexusMessage.md) | The details of the message | 

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

