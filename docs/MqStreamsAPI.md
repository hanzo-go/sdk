# \MqStreamsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqCreateStream**](MqStreamsAPI.md#MqCreateStream) | **Post** /v1/mq/streams | Create stream
[**MqDeleteStream**](MqStreamsAPI.md#MqDeleteStream) | **Delete** /v1/mq/streams/{name} | Delete stream
[**MqDeleteStreamMessage**](MqStreamsAPI.md#MqDeleteStreamMessage) | **Delete** /v1/mq/streams/{name}/messages/{seq} | Delete specific message
[**MqGetStream**](MqStreamsAPI.md#MqGetStream) | **Get** /v1/mq/streams/{name} | Get stream info
[**MqGetStreamMessages**](MqStreamsAPI.md#MqGetStreamMessages) | **Get** /v1/mq/streams/{name}/messages | Get stream messages
[**MqListStreams**](MqStreamsAPI.md#MqListStreams) | **Get** /v1/mq/streams | List streams
[**MqPurgeStream**](MqStreamsAPI.md#MqPurgeStream) | **Post** /v1/mq/streams/{name}/purge | Purge stream messages
[**MqUpdateStream**](MqStreamsAPI.md#MqUpdateStream) | **Put** /v1/mq/streams/{name} | Update stream config



## MqCreateStream

> MqStream MqCreateStream(ctx).MqStreamConfig(mqStreamConfig).Execute()

Create stream



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
	mqStreamConfig := *openapiclient.NewMqStreamConfig("Name_example", []string{"Subjects_example"}) // MqStreamConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqStreamsAPI.MqCreateStream(context.Background()).MqStreamConfig(mqStreamConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqStreamsAPI.MqCreateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqCreateStream`: MqStream
	fmt.Fprintf(os.Stdout, "Response from `MqStreamsAPI.MqCreateStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqCreateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mqStreamConfig** | [**MqStreamConfig**](MqStreamConfig.md) |  | 

### Return type

[**MqStream**](MqStream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqDeleteStream

> MqDeleteStream(ctx, name).Execute()

Delete stream



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
	name := "name_example" // string | Stream name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqStreamsAPI.MqDeleteStream(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqStreamsAPI.MqDeleteStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Stream name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqDeleteStreamRequest struct via the builder pattern


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


## MqDeleteStreamMessage

> MqDeleteStreamMessage(ctx, name, seq).Execute()

Delete specific message



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
	name := "name_example" // string | Stream name.
	seq := int32(56) // int32 | Message sequence number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqStreamsAPI.MqDeleteStreamMessage(context.Background(), name, seq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqStreamsAPI.MqDeleteStreamMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Stream name. | 
**seq** | **int32** | Message sequence number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqDeleteStreamMessageRequest struct via the builder pattern


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


## MqGetStream

> MqStream MqGetStream(ctx, name).Execute()

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
	name := "name_example" // string | Stream name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqStreamsAPI.MqGetStream(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqStreamsAPI.MqGetStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetStream`: MqStream
	fmt.Fprintf(os.Stdout, "Response from `MqStreamsAPI.MqGetStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Stream name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MqStream**](MqStream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqGetStreamMessages

> MqGetStreamMessages200Response MqGetStreamMessages(ctx, name).Seq(seq).LastBySubject(lastBySubject).NextBySubject(nextBySubject).Limit(limit).Execute()

Get stream messages



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
	name := "name_example" // string | Stream name.
	seq := int32(56) // int32 | Get message at this sequence number. (optional)
	lastBySubject := "lastBySubject_example" // string | Get last message for the given subject. (optional)
	nextBySubject := "nextBySubject_example" // string | Get next message for the given subject (requires seq param as starting point).  (optional)
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqStreamsAPI.MqGetStreamMessages(context.Background(), name).Seq(seq).LastBySubject(lastBySubject).NextBySubject(nextBySubject).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqStreamsAPI.MqGetStreamMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetStreamMessages`: MqGetStreamMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `MqStreamsAPI.MqGetStreamMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Stream name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetStreamMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seq** | **int32** | Get message at this sequence number. | 
 **lastBySubject** | **string** | Get last message for the given subject. | 
 **nextBySubject** | **string** | Get next message for the given subject (requires seq param as starting point).  | 
 **limit** | **int32** | Maximum number of items to return. | [default to 100]

### Return type

[**MqGetStreamMessages200Response**](MqGetStreamMessages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListStreams

> MqListStreams200Response MqListStreams(ctx).Limit(limit).Offset(offset).Execute()

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
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqStreamsAPI.MqListStreams(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqStreamsAPI.MqListStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListStreams`: MqListStreams200Response
	fmt.Fprintf(os.Stdout, "Response from `MqStreamsAPI.MqListStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqListStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]

### Return type

[**MqListStreams200Response**](MqListStreams200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqPurgeStream

> MqPurgeStream200Response MqPurgeStream(ctx, name).MqPurgeStreamRequest(mqPurgeStreamRequest).Execute()

Purge stream messages



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
	name := "name_example" // string | Stream name.
	mqPurgeStreamRequest := *openapiclient.NewMqPurgeStreamRequest() // MqPurgeStreamRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqStreamsAPI.MqPurgeStream(context.Background(), name).MqPurgeStreamRequest(mqPurgeStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqStreamsAPI.MqPurgeStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqPurgeStream`: MqPurgeStream200Response
	fmt.Fprintf(os.Stdout, "Response from `MqStreamsAPI.MqPurgeStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Stream name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqPurgeStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mqPurgeStreamRequest** | [**MqPurgeStreamRequest**](MqPurgeStreamRequest.md) |  | 

### Return type

[**MqPurgeStream200Response**](MqPurgeStream200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqUpdateStream

> MqStream MqUpdateStream(ctx, name).MqStreamConfig(mqStreamConfig).Execute()

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
	name := "name_example" // string | Stream name.
	mqStreamConfig := *openapiclient.NewMqStreamConfig("Name_example", []string{"Subjects_example"}) // MqStreamConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqStreamsAPI.MqUpdateStream(context.Background(), name).MqStreamConfig(mqStreamConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqStreamsAPI.MqUpdateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqUpdateStream`: MqStream
	fmt.Fprintf(os.Stdout, "Response from `MqStreamsAPI.MqUpdateStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Stream name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqUpdateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mqStreamConfig** | [**MqStreamConfig**](MqStreamConfig.md) |  | 

### Return type

[**MqStream**](MqStream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

