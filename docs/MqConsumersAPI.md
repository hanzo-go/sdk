# \MqConsumersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqCreateConsumer**](MqConsumersAPI.md#MqCreateConsumer) | **Post** /v1/mq/streams/{stream}/consumers | Create consumer
[**MqDeleteConsumer**](MqConsumersAPI.md#MqDeleteConsumer) | **Delete** /v1/mq/streams/{stream}/consumers/{name} | Delete consumer
[**MqGetConsumer**](MqConsumersAPI.md#MqGetConsumer) | **Get** /v1/mq/streams/{stream}/consumers/{name} | Get consumer info
[**MqListConsumers**](MqConsumersAPI.md#MqListConsumers) | **Get** /v1/mq/streams/{stream}/consumers | List consumers
[**MqPullMessages**](MqConsumersAPI.md#MqPullMessages) | **Post** /v1/mq/streams/{stream}/consumers/{name}/next | Pull next message(s)



## MqCreateConsumer

> MqConsumer MqCreateConsumer(ctx, stream).MqConsumerConfig(mqConsumerConfig).Execute()

Create consumer



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
	stream := "stream_example" // string | Stream name.
	mqConsumerConfig := *openapiclient.NewMqConsumerConfig() // MqConsumerConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqConsumersAPI.MqCreateConsumer(context.Background(), stream).MqConsumerConfig(mqConsumerConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqConsumersAPI.MqCreateConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqCreateConsumer`: MqConsumer
	fmt.Fprintf(os.Stdout, "Response from `MqConsumersAPI.MqCreateConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqCreateConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mqConsumerConfig** | [**MqConsumerConfig**](MqConsumerConfig.md) |  | 

### Return type

[**MqConsumer**](MqConsumer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqDeleteConsumer

> MqDeleteConsumer(ctx, stream, name).Execute()

Delete consumer



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
	stream := "stream_example" // string | Stream name.
	name := "name_example" // string | Consumer name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqConsumersAPI.MqDeleteConsumer(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqConsumersAPI.MqDeleteConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream name. | 
**name** | **string** | Consumer name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqDeleteConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqGetConsumer

> MqConsumer MqGetConsumer(ctx, stream, name).Execute()

Get consumer info



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
	stream := "stream_example" // string | Stream name.
	name := "name_example" // string | Consumer name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqConsumersAPI.MqGetConsumer(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqConsumersAPI.MqGetConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetConsumer`: MqConsumer
	fmt.Fprintf(os.Stdout, "Response from `MqConsumersAPI.MqGetConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream name. | 
**name** | **string** | Consumer name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MqConsumer**](MqConsumer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListConsumers

> MqListConsumers200Response MqListConsumers(ctx, stream).Limit(limit).Offset(offset).Execute()

List consumers



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
	stream := "stream_example" // string | Stream name.
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqConsumersAPI.MqListConsumers(context.Background(), stream).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqConsumersAPI.MqListConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListConsumers`: MqListConsumers200Response
	fmt.Fprintf(os.Stdout, "Response from `MqConsumersAPI.MqListConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqListConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]

### Return type

[**MqListConsumers200Response**](MqListConsumers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqPullMessages

> MqGetStreamMessages200Response MqPullMessages(ctx, stream, name).MqPullMessagesRequest(mqPullMessagesRequest).Execute()

Pull next message(s)



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
	stream := "stream_example" // string | Stream name.
	name := "name_example" // string | Consumer name.
	mqPullMessagesRequest := *openapiclient.NewMqPullMessagesRequest() // MqPullMessagesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqConsumersAPI.MqPullMessages(context.Background(), stream, name).MqPullMessagesRequest(mqPullMessagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqConsumersAPI.MqPullMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqPullMessages`: MqGetStreamMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `MqConsumersAPI.MqPullMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream name. | 
**name** | **string** | Consumer name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqPullMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mqPullMessagesRequest** | [**MqPullMessagesRequest**](MqPullMessagesRequest.md) |  | 

### Return type

[**MqGetStreamMessages200Response**](MqGetStreamMessages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

