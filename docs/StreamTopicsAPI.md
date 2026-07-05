# \StreamTopicsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamCreateTopic**](StreamTopicsAPI.md#StreamCreateTopic) | **Post** /v1/stream/topics | Create a topic
[**StreamDeleteTopic**](StreamTopicsAPI.md#StreamDeleteTopic) | **Delete** /v1/stream/topics/{topic} | Delete a topic
[**StreamGetTopicMetadata**](StreamTopicsAPI.md#StreamGetTopicMetadata) | **Get** /v1/stream/topics/{topic} | Get topic metadata
[**StreamListTopics**](StreamTopicsAPI.md#StreamListTopics) | **Get** /v1/stream/topics | List topics



## StreamCreateTopic

> StreamTopic StreamCreateTopic(ctx).StreamCreateTopicRequest(streamCreateTopicRequest).Execute()

Create a topic



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
	streamCreateTopicRequest := *openapiclient.NewStreamCreateTopicRequest("events") // StreamCreateTopicRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamTopicsAPI.StreamCreateTopic(context.Background()).StreamCreateTopicRequest(streamCreateTopicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamTopicsAPI.StreamCreateTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamCreateTopic`: StreamTopic
	fmt.Fprintf(os.Stdout, "Response from `StreamTopicsAPI.StreamCreateTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStreamCreateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamCreateTopicRequest** | [**StreamCreateTopicRequest**](StreamCreateTopicRequest.md) |  | 

### Return type

[**StreamTopic**](StreamTopic.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamDeleteTopic

> StreamDeleteTopic(ctx, topic).Execute()

Delete a topic



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
	topic := "topic_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StreamTopicsAPI.StreamDeleteTopic(context.Background(), topic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamTopicsAPI.StreamDeleteTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamDeleteTopicRequest struct via the builder pattern


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


## StreamGetTopicMetadata

> StreamTopicMetadata StreamGetTopicMetadata(ctx, topic).Execute()

Get topic metadata



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
	topic := "topic_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamTopicsAPI.StreamGetTopicMetadata(context.Background(), topic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamTopicsAPI.StreamGetTopicMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamGetTopicMetadata`: StreamTopicMetadata
	fmt.Fprintf(os.Stdout, "Response from `StreamTopicsAPI.StreamGetTopicMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamGetTopicMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamTopicMetadata**](StreamTopicMetadata.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamListTopics

> StreamListTopics200Response StreamListTopics(ctx).Execute()

List topics



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
	resp, r, err := apiClient.StreamTopicsAPI.StreamListTopics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamTopicsAPI.StreamListTopics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamListTopics`: StreamListTopics200Response
	fmt.Fprintf(os.Stdout, "Response from `StreamTopicsAPI.StreamListTopics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStreamListTopicsRequest struct via the builder pattern


### Return type

[**StreamListTopics200Response**](StreamListTopics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

