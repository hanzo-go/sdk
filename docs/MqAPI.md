# \MqAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMqStreamByName**](MqAPI.md#DeleteMqStreamByName) | **Delete** /v1/mq/stream/{name} | Removes a stream with all its messages and consumers.
[**DeleteMqStreamByNameMessageBySeq**](MqAPI.md#DeleteMqStreamByNameMessageBySeq) | **Delete** /v1/mq/stream/{name}/message/{seq} | Erases one message by sequence; the sequence gap remains.
[**DeleteMqStreamByStreamConsumerByName**](MqAPI.md#DeleteMqStreamByStreamConsumerByName) | **Delete** /v1/mq/stream/{stream}/consumer/{name} | Removes a consumer and its delivery state; unacknowledged messages stay in the stream.
[**GetMqHealth**](MqAPI.md#GetMqHealth) | **Get** /v1/mq/health | Reports whether the message plane behind this surface answers.
[**GetMqInfo**](MqAPI.md#GetMqInfo) | **Get** /v1/mq/info | Returns the broker&#39;s identity and the org&#39;s stream count.
[**GetMqStream**](MqAPI.md#GetMqStream) | **Get** /v1/mq/stream | Returns the org&#39;s streams, name-ordered, with their live state.
[**GetMqStreamByName**](MqAPI.md#GetMqStreamByName) | **Get** /v1/mq/stream/{name} | Returns one stream&#39;s configuration and live state.
[**GetMqStreamByNameMessage**](MqAPI.md#GetMqStreamByNameMessage) | **Get** /v1/mq/stream/{name}/message | Reads stored messages without a consumer: by sequence, by newest on a subject, or walking a subject forward from a sequence.
[**GetMqStreamByStreamConsumer**](MqAPI.md#GetMqStreamByStreamConsumer) | **Get** /v1/mq/stream/{stream}/consumer | Returns a stream&#39;s consumers, name-ordered, with delivery state.
[**GetMqStreamByStreamConsumerByName**](MqAPI.md#GetMqStreamByStreamConsumerByName) | **Get** /v1/mq/stream/{stream}/consumer/{name} | Returns one consumer&#39;s configuration and delivery state.
[**PostMqStream**](MqAPI.md#PostMqStream) | **Post** /v1/mq/stream | Creates a durable stream in the org&#39;s namespace and returns it.
[**PostMqStreamByNamePurge**](MqAPI.md#PostMqStreamByNamePurge) | **Post** /v1/mq/stream/{name}/purge | Removes messages from a stream, leaving its consumers in place.
[**PostMqStreamByStreamConsumer**](MqAPI.md#PostMqStreamByStreamConsumer) | **Post** /v1/mq/stream/{stream}/consumer | Creates a durable pull consumer on a stream and returns it.
[**PostMqStreamByStreamConsumerByNameNext**](MqAPI.md#PostMqStreamByStreamConsumerByNameNext) | **Post** /v1/mq/stream/{stream}/consumer/{name}/next | Pulls the consumer&#39;s next batch.
[**PutMqStreamByName**](MqAPI.md#PutMqStreamByName) | **Put** /v1/mq/stream/{name} | Reconfigures an existing stream; the path names the stream, and the immutable fields (storage, retention) must restate what they are.



## DeleteMqStreamByName

> DeleteMqStreamByName(ctx, name).Execute()

Removes a stream with all its messages and consumers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the stream name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqAPI.DeleteMqStreamByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.DeleteMqStreamByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMqStreamByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMqStreamByNameMessageBySeq

> DeleteMqStreamByNameMessageBySeq(ctx, name, seq).Execute()

Erases one message by sequence; the sequence gap remains.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the stream name, from the path.
	seq := int32(56) // int32 | Seq is the message's stream sequence, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqAPI.DeleteMqStreamByNameMessageBySeq(context.Background(), name, seq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.DeleteMqStreamByNameMessageBySeq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 
**seq** | **int32** | Seq is the message&#39;s stream sequence, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMqStreamByNameMessageBySeqRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMqStreamByStreamConsumerByName

> DeleteMqStreamByStreamConsumerByName(ctx, stream, name).Execute()

Removes a consumer and its delivery state; unacknowledged messages stay in the stream.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	stream := "stream_example" // string | Stream is the stream name, from the path.
	name := "name_example" // string | Name is the consumer name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqAPI.DeleteMqStreamByStreamConsumerByName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.DeleteMqStreamByStreamConsumerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 
**name** | **string** | Name is the consumer name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMqStreamByStreamConsumerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqHealth

> MqHealth GetMqHealth(ctx).Execute()

Reports whether the message plane behind this surface answers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqHealth`: MqHealth
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqHealthRequest struct via the builder pattern


### Return type

[**MqHealth**](MqHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqInfo

> MqInfoOut GetMqInfo(ctx).Execute()

Returns the broker's identity and the org's stream count.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqInfo`: MqInfoOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqInfoRequest struct via the builder pattern


### Return type

[**MqInfoOut**](MqInfoOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStream

> MqStreams GetMqStream(ctx).Limit(limit).Offset(offset).Execute()

Returns the org's streams, name-ordered, with their live state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	limit := int64(789) // int64 | Limit caps the streams returned (1–1000, default 100). (optional)
	offset := int64(789) // int64 | Offset skips that many streams, name-ordered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStream(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStream`: MqStreams
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the streams returned (1–1000, default 100). | 
 **offset** | **int64** | Offset skips that many streams, name-ordered. | 

### Return type

[**MqStreams**](MqStreams.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreamByName

> MqStream GetMqStreamByName(ctx, name).Execute()

Returns one stream's configuration and live state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the stream name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreamByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreamByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreamByName`: MqStream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreamByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MqStream**](MqStream.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreamByNameMessage

> MqReadOut GetMqStreamByNameMessage(ctx, name).Seq(seq).LastBySubject(lastBySubject).NextBySubject(nextBySubject).Limit(limit).Execute()

Reads stored messages without a consumer: by sequence, by newest on a subject, or walking a subject forward from a sequence.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the stream name, from the path.
	seq := int32(56) // int32 | Seq reads the message at this sequence (with next_by_subject: the walk's start). (optional)
	lastBySubject := "lastBySubject_example" // string | LastBySubject reads the newest message on this org-relative subject. (optional)
	nextBySubject := "nextBySubject_example" // string | NextBySubject walks forward from seq collecting messages on this org-relative subject (wildcards supported). (optional)
	limit := int64(789) // int64 | Limit caps a next_by_subject walk (1–1000, default 100). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreamByNameMessage(context.Background(), name).Seq(seq).LastBySubject(lastBySubject).NextBySubject(nextBySubject).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreamByNameMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreamByNameMessage`: MqReadOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreamByNameMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamByNameMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seq** | **int32** | Seq reads the message at this sequence (with next_by_subject: the walk&#39;s start). | 
 **lastBySubject** | **string** | LastBySubject reads the newest message on this org-relative subject. | 
 **nextBySubject** | **string** | NextBySubject walks forward from seq collecting messages on this org-relative subject (wildcards supported). | 
 **limit** | **int64** | Limit caps a next_by_subject walk (1–1000, default 100). | 

### Return type

[**MqReadOut**](MqReadOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreamByStreamConsumer

> MqPickOut GetMqStreamByStreamConsumer(ctx, stream).Limit(limit).Offset(offset).Execute()

Returns a stream's consumers, name-ordered, with delivery state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	stream := "stream_example" // string | Stream is the stream name, from the path.
	limit := int64(789) // int64 | Limit caps the consumers returned (1–1000, default 100). (optional)
	offset := int64(789) // int64 | Offset skips that many consumers, name-ordered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreamByStreamConsumer(context.Background(), stream).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreamByStreamConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreamByStreamConsumer`: MqPickOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreamByStreamConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamByStreamConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Limit caps the consumers returned (1–1000, default 100). | 
 **offset** | **int64** | Offset skips that many consumers, name-ordered. | 

### Return type

[**MqPickOut**](MqPickOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreamByStreamConsumerByName

> MqConsumer GetMqStreamByStreamConsumerByName(ctx, stream, name).Execute()

Returns one consumer's configuration and delivery state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	stream := "stream_example" // string | Stream is the stream name, from the path.
	name := "name_example" // string | Name is the consumer name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreamByStreamConsumerByName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreamByStreamConsumerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreamByStreamConsumerByName`: MqConsumer
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreamByStreamConsumerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 
**name** | **string** | Name is the consumer name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamByStreamConsumerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MqConsumer**](MqConsumer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMqStream

> MqStream PostMqStream(ctx).MqConfig(mqConfig).Execute()

Creates a durable stream in the org's namespace and returns it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	mqConfig := *openapiclient.NewMqConfig() // MqConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PostMqStream(context.Background()).MqConfig(mqConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PostMqStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMqStream`: MqStream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PostMqStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMqStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mqConfig** | [**MqConfig**](MqConfig.md) |  | 

### Return type

[**MqStream**](MqStream.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMqStreamByNamePurge

> MqPurgeOut PostMqStreamByNamePurge(ctx, name).MqPurge(mqPurge).Execute()

Removes messages from a stream, leaving its consumers in place.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the stream name, from the path.
	mqPurge := *openapiclient.NewMqPurge() // MqPurge | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PostMqStreamByNamePurge(context.Background(), name).MqPurge(mqPurge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PostMqStreamByNamePurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMqStreamByNamePurge`: MqPurgeOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PostMqStreamByNamePurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMqStreamByNamePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mqPurge** | [**MqPurge**](MqPurge.md) |  | 

### Return type

[**MqPurgeOut**](MqPurgeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMqStreamByStreamConsumer

> MqConsumer PostMqStreamByStreamConsumer(ctx, stream).MqMakeIn(mqMakeIn).Execute()

Creates a durable pull consumer on a stream and returns it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	stream := "stream_example" // string | Stream is the stream name, from the path.
	mqMakeIn := *openapiclient.NewMqMakeIn() // MqMakeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PostMqStreamByStreamConsumer(context.Background(), stream).MqMakeIn(mqMakeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PostMqStreamByStreamConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMqStreamByStreamConsumer`: MqConsumer
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PostMqStreamByStreamConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMqStreamByStreamConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mqMakeIn** | [**MqMakeIn**](MqMakeIn.md) |  | 

### Return type

[**MqConsumer**](MqConsumer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMqStreamByStreamConsumerByNameNext

> MqReadOut PostMqStreamByStreamConsumerByNameNext(ctx, stream, name).MqNextIn(mqNextIn).Execute()

Pulls the consumer's next batch.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	stream := "stream_example" // string | Stream is the stream name, from the path.
	name := "name_example" // string | Name is the consumer name, from the path.
	mqNextIn := *openapiclient.NewMqNextIn() // MqNextIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PostMqStreamByStreamConsumerByNameNext(context.Background(), stream, name).MqNextIn(mqNextIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PostMqStreamByStreamConsumerByNameNext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMqStreamByStreamConsumerByNameNext`: MqReadOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PostMqStreamByStreamConsumerByNameNext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 
**name** | **string** | Name is the consumer name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMqStreamByStreamConsumerByNameNextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mqNextIn** | [**MqNextIn**](MqNextIn.md) |  | 

### Return type

[**MqReadOut**](MqReadOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMqStreamByName

> MqStream PutMqStreamByName(ctx, name).MqConfig(mqConfig).Execute()

Reconfigures an existing stream; the path names the stream, and the immutable fields (storage, retention) must restate what they are.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the stream name, unique within the org (alphanumeric, hyphens, underscores).
	mqConfig := *openapiclient.NewMqConfig() // MqConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PutMqStreamByName(context.Background(), name).MqConfig(mqConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PutMqStreamByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMqStreamByName`: MqStream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PutMqStreamByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, unique within the org (alphanumeric, hyphens, underscores). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMqStreamByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mqConfig** | [**MqConfig**](MqConfig.md) |  | 

### Return type

[**MqStream**](MqStream.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

