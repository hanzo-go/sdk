# \PubsubStreamsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PubsubCreateStream**](PubsubStreamsAPI.md#PubsubCreateStream) | **Post** /v1/pubsub/jetstream/streams | Create a stream
[**PubsubDeleteStream**](PubsubStreamsAPI.md#PubsubDeleteStream) | **Delete** /v1/pubsub/jetstream/streams/{stream} | Delete a stream
[**PubsubGetStream**](PubsubStreamsAPI.md#PubsubGetStream) | **Get** /v1/pubsub/jetstream/streams/{stream} | Get stream info
[**PubsubListStreams**](PubsubStreamsAPI.md#PubsubListStreams) | **Get** /v1/pubsub/jetstream/streams | List streams
[**PubsubUpdateStream**](PubsubStreamsAPI.md#PubsubUpdateStream) | **Put** /v1/pubsub/jetstream/streams/{stream} | Update stream config



## PubsubCreateStream

> PubsubStreamInfo PubsubCreateStream(ctx).PubsubStreamConfig(pubsubStreamConfig).Execute()

Create a stream



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
	pubsubStreamConfig := *openapiclient.NewPubsubStreamConfig("ORDERS", []string{"Subjects_example"}) // PubsubStreamConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubStreamsAPI.PubsubCreateStream(context.Background()).PubsubStreamConfig(pubsubStreamConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubStreamsAPI.PubsubCreateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubCreateStream`: PubsubStreamInfo
	fmt.Fprintf(os.Stdout, "Response from `PubsubStreamsAPI.PubsubCreateStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPubsubCreateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubsubStreamConfig** | [**PubsubStreamConfig**](PubsubStreamConfig.md) |  | 

### Return type

[**PubsubStreamInfo**](PubsubStreamInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubDeleteStream

> PubsubDeleteStream(ctx, stream).Execute()

Delete a stream

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
	r, err := apiClient.PubsubStreamsAPI.PubsubDeleteStream(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubStreamsAPI.PubsubDeleteStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubDeleteStreamRequest struct via the builder pattern


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


## PubsubGetStream

> PubsubStreamInfo PubsubGetStream(ctx, stream).Execute()

Get stream info

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
	resp, r, err := apiClient.PubsubStreamsAPI.PubsubGetStream(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubStreamsAPI.PubsubGetStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetStream`: PubsubStreamInfo
	fmt.Fprintf(os.Stdout, "Response from `PubsubStreamsAPI.PubsubGetStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PubsubStreamInfo**](PubsubStreamInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubListStreams

> PubsubListStreams200Response PubsubListStreams(ctx).Execute()

List streams



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
	resp, r, err := apiClient.PubsubStreamsAPI.PubsubListStreams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubStreamsAPI.PubsubListStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubListStreams`: PubsubListStreams200Response
	fmt.Fprintf(os.Stdout, "Response from `PubsubStreamsAPI.PubsubListStreams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubListStreamsRequest struct via the builder pattern


### Return type

[**PubsubListStreams200Response**](PubsubListStreams200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubUpdateStream

> PubsubStreamInfo PubsubUpdateStream(ctx, stream).PubsubStreamConfig(pubsubStreamConfig).Execute()

Update stream config

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
	pubsubStreamConfig := *openapiclient.NewPubsubStreamConfig("ORDERS", []string{"Subjects_example"}) // PubsubStreamConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubStreamsAPI.PubsubUpdateStream(context.Background(), stream).PubsubStreamConfig(pubsubStreamConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubStreamsAPI.PubsubUpdateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubUpdateStream`: PubsubStreamInfo
	fmt.Fprintf(os.Stdout, "Response from `PubsubStreamsAPI.PubsubUpdateStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubUpdateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pubsubStreamConfig** | [**PubsubStreamConfig**](PubsubStreamConfig.md) |  | 

### Return type

[**PubsubStreamInfo**](PubsubStreamInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

