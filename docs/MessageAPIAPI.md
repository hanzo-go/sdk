# \MessageAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddMessage**](MessageAPIAPI.md#CloudApiControllerAddMessage) | **Post** /v1/cloud/add-message | Api Controller Add Message
[**CloudApiControllerDeleteMessage**](MessageAPIAPI.md#CloudApiControllerDeleteMessage) | **Post** /v1/cloud/delete-message | Api Controller Delete Message
[**CloudApiControllerGetAnswer**](MessageAPIAPI.md#CloudApiControllerGetAnswer) | **Get** /v1/cloud/get-answer | Api Controller Get Answer
[**CloudApiControllerGetGlobalMessages**](MessageAPIAPI.md#CloudApiControllerGetGlobalMessages) | **Get** /v1/cloud/get-global-messages | Api Controller Get Global Messages
[**CloudApiControllerGetMessage**](MessageAPIAPI.md#CloudApiControllerGetMessage) | **Get** /v1/cloud/get-message | Api Controller Get Message
[**CloudApiControllerGetMessageAnswer**](MessageAPIAPI.md#CloudApiControllerGetMessageAnswer) | **Get** /v1/cloud/get-message-answer | Api Controller Get Message Answer
[**CloudApiControllerGetMessages**](MessageAPIAPI.md#CloudApiControllerGetMessages) | **Get** /v1/cloud/get-Messages | Api Controller Get Messages
[**CloudApiControllerUpdateMessage**](MessageAPIAPI.md#CloudApiControllerUpdateMessage) | **Post** /v1/cloud/update-message | Api Controller Update Message



## CloudApiControllerAddMessage

> CloudObjectChat CloudApiControllerAddMessage(ctx).CloudObjectMessage(cloudObjectMessage).Execute()

Api Controller Add Message



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
	cloudObjectMessage := *openapiclient.NewCloudObjectMessage() // CloudObjectMessage | The details of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPIAPI.CloudApiControllerAddMessage(context.Background()).CloudObjectMessage(cloudObjectMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPIAPI.CloudApiControllerAddMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddMessage`: CloudObjectChat
	fmt.Fprintf(os.Stdout, "Response from `MessageAPIAPI.CloudApiControllerAddMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectMessage** | [**CloudObjectMessage**](CloudObjectMessage.md) | The details of the message | 

### Return type

[**CloudObjectChat**](CloudObjectChat.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteMessage

> CloudControllersResponse CloudApiControllerDeleteMessage(ctx).CloudObjectMessage(cloudObjectMessage).Execute()

Api Controller Delete Message



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
	cloudObjectMessage := *openapiclient.NewCloudObjectMessage() // CloudObjectMessage | The details of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPIAPI.CloudApiControllerDeleteMessage(context.Background()).CloudObjectMessage(cloudObjectMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPIAPI.CloudApiControllerDeleteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteMessage`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPIAPI.CloudApiControllerDeleteMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectMessage** | [**CloudObjectMessage**](CloudObjectMessage.md) | The details of the message | 

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


## CloudApiControllerGetAnswer

> map[string]interface{} CloudApiControllerGetAnswer(ctx).Provider(provider).Question(question).Framework(framework).Video(video).Execute()

Api Controller Get Answer



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
	question := "question_example" // string | The question of message
	framework := "framework_example" // string | The framework
	video := "video_example" // string | The video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPIAPI.CloudApiControllerGetAnswer(context.Background()).Provider(provider).Question(question).Framework(framework).Video(video).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPIAPI.CloudApiControllerGetAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetAnswer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MessageAPIAPI.CloudApiControllerGetAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | The provider | 
 **question** | **string** | The question of message | 
 **framework** | **string** | The framework | 
 **video** | **string** | The video | 

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


## CloudApiControllerGetGlobalMessages

> []CloudObjectMessage CloudApiControllerGetGlobalMessages(ctx).Execute()

Api Controller Get Global Messages



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
	resp, r, err := apiClient.MessageAPIAPI.CloudApiControllerGetGlobalMessages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPIAPI.CloudApiControllerGetGlobalMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalMessages`: []CloudObjectMessage
	fmt.Fprintf(os.Stdout, "Response from `MessageAPIAPI.CloudApiControllerGetGlobalMessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalMessagesRequest struct via the builder pattern


### Return type

[**[]CloudObjectMessage**](CloudObjectMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetMessage

> CloudObjectMessage CloudApiControllerGetMessage(ctx).Id(id).Execute()

Api Controller Get Message



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
	id := "id_example" // string | The id of message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPIAPI.CloudApiControllerGetMessage(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPIAPI.CloudApiControllerGetMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetMessage`: CloudObjectMessage
	fmt.Fprintf(os.Stdout, "Response from `MessageAPIAPI.CloudApiControllerGetMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of message | 

### Return type

[**CloudObjectMessage**](CloudObjectMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetMessageAnswer

> map[string]interface{} CloudApiControllerGetMessageAnswer(ctx).Id(id).Execute()

Api Controller Get Message Answer



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
	id := "id_example" // string | The id of message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPIAPI.CloudApiControllerGetMessageAnswer(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPIAPI.CloudApiControllerGetMessageAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetMessageAnswer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MessageAPIAPI.CloudApiControllerGetMessageAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetMessageAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of message | 

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


## CloudApiControllerGetMessages

> []CloudObjectMessage CloudApiControllerGetMessages(ctx).User(user).Chat(chat).Execute()

Api Controller Get Messages



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
	user := "user_example" // string | The user of message
	chat := "chat_example" // string | The chat of message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPIAPI.CloudApiControllerGetMessages(context.Background()).User(user).Chat(chat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPIAPI.CloudApiControllerGetMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetMessages`: []CloudObjectMessage
	fmt.Fprintf(os.Stdout, "Response from `MessageAPIAPI.CloudApiControllerGetMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | The user of message | 
 **chat** | **string** | The chat of message | 

### Return type

[**[]CloudObjectMessage**](CloudObjectMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateMessage

> CloudControllersResponse CloudApiControllerUpdateMessage(ctx).Id(id).CloudObjectMessage(cloudObjectMessage).Execute()

Api Controller Update Message



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
	cloudObjectMessage := *openapiclient.NewCloudObjectMessage() // CloudObjectMessage | The details of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPIAPI.CloudApiControllerUpdateMessage(context.Background()).Id(id).CloudObjectMessage(cloudObjectMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPIAPI.CloudApiControllerUpdateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateMessage`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPIAPI.CloudApiControllerUpdateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the message | 
 **cloudObjectMessage** | [**CloudObjectMessage**](CloudObjectMessage.md) | The details of the message | 

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

