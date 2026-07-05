# \PubsubConsumersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PubsubCreateConsumer**](PubsubConsumersAPI.md#PubsubCreateConsumer) | **Post** /v1/pubsub/jetstream/streams/{stream}/consumers | Create a consumer
[**PubsubDeleteConsumer**](PubsubConsumersAPI.md#PubsubDeleteConsumer) | **Delete** /v1/pubsub/jetstream/streams/{stream}/consumers/{consumer} | Delete a consumer
[**PubsubGetConsumer**](PubsubConsumersAPI.md#PubsubGetConsumer) | **Get** /v1/pubsub/jetstream/streams/{stream}/consumers/{consumer} | Get consumer info
[**PubsubListConsumers**](PubsubConsumersAPI.md#PubsubListConsumers) | **Get** /v1/pubsub/jetstream/streams/{stream}/consumers | List consumers



## PubsubCreateConsumer

> PubsubConsumerInfo PubsubCreateConsumer(ctx, stream).PubsubConsumerConfig(pubsubConsumerConfig).Execute()

Create a consumer



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
	stream := "stream_example" // string | 
	pubsubConsumerConfig := *openapiclient.NewPubsubConsumerConfig("order-processor") // PubsubConsumerConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubConsumersAPI.PubsubCreateConsumer(context.Background(), stream).PubsubConsumerConfig(pubsubConsumerConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubConsumersAPI.PubsubCreateConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubCreateConsumer`: PubsubConsumerInfo
	fmt.Fprintf(os.Stdout, "Response from `PubsubConsumersAPI.PubsubCreateConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubCreateConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pubsubConsumerConfig** | [**PubsubConsumerConfig**](PubsubConsumerConfig.md) |  | 

### Return type

[**PubsubConsumerInfo**](PubsubConsumerInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubDeleteConsumer

> PubsubDeleteConsumer(ctx, stream, consumer).Execute()

Delete a consumer

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
	stream := "stream_example" // string | 
	consumer := "consumer_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubsubConsumersAPI.PubsubDeleteConsumer(context.Background(), stream, consumer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubConsumersAPI.PubsubDeleteConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** |  | 
**consumer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubDeleteConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubGetConsumer

> PubsubConsumerInfo PubsubGetConsumer(ctx, stream, consumer).Execute()

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
	stream := "stream_example" // string | 
	consumer := "consumer_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubConsumersAPI.PubsubGetConsumer(context.Background(), stream, consumer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubConsumersAPI.PubsubGetConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetConsumer`: PubsubConsumerInfo
	fmt.Fprintf(os.Stdout, "Response from `PubsubConsumersAPI.PubsubGetConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** |  | 
**consumer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PubsubConsumerInfo**](PubsubConsumerInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubListConsumers

> PubsubListConsumers200Response PubsubListConsumers(ctx, stream).Execute()

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
	stream := "stream_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubConsumersAPI.PubsubListConsumers(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubConsumersAPI.PubsubListConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubListConsumers`: PubsubListConsumers200Response
	fmt.Fprintf(os.Stdout, "Response from `PubsubConsumersAPI.PubsubListConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubListConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PubsubListConsumers200Response**](PubsubListConsumers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

