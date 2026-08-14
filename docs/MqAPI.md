# \MqAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMqStreamsByName**](MqAPI.md#DeleteMqStreamsByName) | **Delete** /v1/mq/streams/{name} | Removes a stream with all its messages and consumers.
[**DeleteMqStreamsByNameMessagesBySeq**](MqAPI.md#DeleteMqStreamsByNameMessagesBySeq) | **Delete** /v1/mq/streams/{name}/messages/{seq} | Erases one message by sequence; the sequence gap remains.
[**DeleteMqStreamsByStreamConsumersByName**](MqAPI.md#DeleteMqStreamsByStreamConsumersByName) | **Delete** /v1/mq/streams/{stream}/consumers/{name} | Removes a consumer and its delivery state; unacknowledged messages stay in the stream.
[**GetMqHealth**](MqAPI.md#GetMqHealth) | **Get** /v1/mq/health | Reports whether the message plane behind this surface answers.
[**GetMqInfo**](MqAPI.md#GetMqInfo) | **Get** /v1/mq/info | Returns the broker&#39;s identity and the org&#39;s stream count.
[**GetMqStreams**](MqAPI.md#GetMqStreams) | **Get** /v1/mq/streams | Returns the org&#39;s streams, name-ordered, with their live state.
[**GetMqStreamsByName**](MqAPI.md#GetMqStreamsByName) | **Get** /v1/mq/streams/{name} | Returns one stream&#39;s configuration and live state.
[**GetMqStreamsByNameMessages**](MqAPI.md#GetMqStreamsByNameMessages) | **Get** /v1/mq/streams/{name}/messages | Reads stored messages without a consumer: by sequence, by newest on a subject, or walking a subject forward from a sequence.
[**GetMqStreamsByStreamConsumers**](MqAPI.md#GetMqStreamsByStreamConsumers) | **Get** /v1/mq/streams/{stream}/consumers | Returns a stream&#39;s consumers, name-ordered, with delivery state.
[**GetMqStreamsByStreamConsumersByName**](MqAPI.md#GetMqStreamsByStreamConsumersByName) | **Get** /v1/mq/streams/{stream}/consumers/{name} | Returns one consumer&#39;s configuration and delivery state.
[**PostMqStreams**](MqAPI.md#PostMqStreams) | **Post** /v1/mq/streams | Creates a durable stream in the org&#39;s namespace and returns it.
[**PostMqStreamsByNamePurge**](MqAPI.md#PostMqStreamsByNamePurge) | **Post** /v1/mq/streams/{name}/purge | Removes messages from a stream, leaving its consumers in place.
[**PostMqStreamsByStreamConsumers**](MqAPI.md#PostMqStreamsByStreamConsumers) | **Post** /v1/mq/streams/{stream}/consumers | Creates a durable pull consumer on a stream and returns it.
[**PostMqStreamsByStreamConsumersByNameNext**](MqAPI.md#PostMqStreamsByStreamConsumersByNameNext) | **Post** /v1/mq/streams/{stream}/consumers/{name}/next | Pulls the consumer&#39;s next batch.
[**PutMqStreamsByName**](MqAPI.md#PutMqStreamsByName) | **Put** /v1/mq/streams/{name} | Reconfigures an existing stream; the path names the stream, and the immutable fields (storage, retention) must restate what they are.



## DeleteMqStreamsByName

> DeleteMqStreamsByName(ctx, name).Execute()

Removes a stream with all its messages and consumers.



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
	name := "name_example" // string | Name is the stream name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqAPI.DeleteMqStreamsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.DeleteMqStreamsByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMqStreamsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMqStreamsByNameMessagesBySeq

> DeleteMqStreamsByNameMessagesBySeq(ctx, name, seq).Execute()

Erases one message by sequence; the sequence gap remains.



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
	name := "name_example" // string | Name is the stream name, from the path.
	seq := int32(56) // int32 | Seq is the message's stream sequence, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqAPI.DeleteMqStreamsByNameMessagesBySeq(context.Background(), name, seq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.DeleteMqStreamsByNameMessagesBySeq``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMqStreamsByNameMessagesBySeqRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMqStreamsByStreamConsumersByName

> DeleteMqStreamsByStreamConsumersByName(ctx, stream, name).Execute()

Removes a consumer and its delivery state; unacknowledged messages stay in the stream.



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
	stream := "stream_example" // string | Stream is the stream name, from the path.
	name := "name_example" // string | Name is the consumer name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqAPI.DeleteMqStreamsByStreamConsumersByName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.DeleteMqStreamsByStreamConsumersByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMqStreamsByStreamConsumersByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqHealth

> Health GetMqHealth(ctx).Execute()

Reports whether the message plane behind this surface answers.



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
	resp, r, err := apiClient.MqAPI.GetMqHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqHealth`: Health
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqHealthRequest struct via the builder pattern


### Return type

[**Health**](Health.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqInfo

> InfoOut GetMqInfo(ctx).Execute()

Returns the broker's identity and the org's stream count.



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
	resp, r, err := apiClient.MqAPI.GetMqInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqInfo`: InfoOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqInfoRequest struct via the builder pattern


### Return type

[**InfoOut**](InfoOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreams

> Streams GetMqStreams(ctx).Limit(limit).Offset(offset).Execute()

Returns the org's streams, name-ordered, with their live state.



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
	limit := int32(56) // int32 | Limit caps the streams returned (1–1000, default 100). (optional)
	offset := int32(56) // int32 | Offset skips that many streams, name-ordered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreams(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreams`: Streams
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the streams returned (1–1000, default 100). | 
 **offset** | **int32** | Offset skips that many streams, name-ordered. | 

### Return type

[**Streams**](Streams.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreamsByName

> Stream GetMqStreamsByName(ctx, name).Execute()

Returns one stream's configuration and live state.



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
	name := "name_example" // string | Name is the stream name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreamsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreamsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreamsByName`: Stream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreamsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stream**](Stream.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreamsByNameMessages

> ReadOut GetMqStreamsByNameMessages(ctx, name).Seq(seq).LastBySubject(lastBySubject).NextBySubject(nextBySubject).Limit(limit).Execute()

Reads stored messages without a consumer: by sequence, by newest on a subject, or walking a subject forward from a sequence.



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
	name := "name_example" // string | Name is the stream name, from the path.
	seq := int32(56) // int32 | Seq reads the message at this sequence (with next_by_subject: the walk's start). (optional)
	lastBySubject := "lastBySubject_example" // string | LastBySubject reads the newest message on this org-relative subject. (optional)
	nextBySubject := "nextBySubject_example" // string | NextBySubject walks forward from seq collecting messages on this org-relative subject (wildcards supported). (optional)
	limit := int32(56) // int32 | Limit caps a next_by_subject walk (1–1000, default 100). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreamsByNameMessages(context.Background(), name).Seq(seq).LastBySubject(lastBySubject).NextBySubject(nextBySubject).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreamsByNameMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreamsByNameMessages`: ReadOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreamsByNameMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamsByNameMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seq** | **int32** | Seq reads the message at this sequence (with next_by_subject: the walk&#39;s start). | 
 **lastBySubject** | **string** | LastBySubject reads the newest message on this org-relative subject. | 
 **nextBySubject** | **string** | NextBySubject walks forward from seq collecting messages on this org-relative subject (wildcards supported). | 
 **limit** | **int32** | Limit caps a next_by_subject walk (1–1000, default 100). | 

### Return type

[**ReadOut**](ReadOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreamsByStreamConsumers

> PickOut GetMqStreamsByStreamConsumers(ctx, stream).Limit(limit).Offset(offset).Execute()

Returns a stream's consumers, name-ordered, with delivery state.



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
	stream := "stream_example" // string | Stream is the stream name, from the path.
	limit := int32(56) // int32 | Limit caps the consumers returned (1–1000, default 100). (optional)
	offset := int32(56) // int32 | Offset skips that many consumers, name-ordered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreamsByStreamConsumers(context.Background(), stream).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreamsByStreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreamsByStreamConsumers`: PickOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreamsByStreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamsByStreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps the consumers returned (1–1000, default 100). | 
 **offset** | **int32** | Offset skips that many consumers, name-ordered. | 

### Return type

[**PickOut**](PickOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMqStreamsByStreamConsumersByName

> Consumer GetMqStreamsByStreamConsumersByName(ctx, stream, name).Execute()

Returns one consumer's configuration and delivery state.



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
	stream := "stream_example" // string | Stream is the stream name, from the path.
	name := "name_example" // string | Name is the consumer name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.GetMqStreamsByStreamConsumersByName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.GetMqStreamsByStreamConsumersByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMqStreamsByStreamConsumersByName`: Consumer
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.GetMqStreamsByStreamConsumersByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 
**name** | **string** | Name is the consumer name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMqStreamsByStreamConsumersByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Consumer**](Consumer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMqStreams

> Stream PostMqStreams(ctx).StreamConfig(streamConfig).Execute()

Creates a durable stream in the org's namespace and returns it.



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
	streamConfig := *openapiclient.NewStreamConfig() // StreamConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PostMqStreams(context.Background()).StreamConfig(streamConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PostMqStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMqStreams`: Stream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PostMqStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMqStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamConfig** | [**StreamConfig**](StreamConfig.md) |  | 

### Return type

[**Stream**](Stream.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMqStreamsByNamePurge

> PurgeOut PostMqStreamsByNamePurge(ctx, name).Purge(purge).Execute()

Removes messages from a stream, leaving its consumers in place.



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
	name := "name_example" // string | Name is the stream name, from the path.
	purge := *openapiclient.NewPurge() // Purge | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PostMqStreamsByNamePurge(context.Background(), name).Purge(purge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PostMqStreamsByNamePurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMqStreamsByNamePurge`: PurgeOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PostMqStreamsByNamePurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMqStreamsByNamePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | [**Purge**](Purge.md) |  | 

### Return type

[**PurgeOut**](PurgeOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMqStreamsByStreamConsumers

> Consumer PostMqStreamsByStreamConsumers(ctx, stream).MakeIn(makeIn).Execute()

Creates a durable pull consumer on a stream and returns it.



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
	stream := "stream_example" // string | Stream is the stream name, from the path.
	makeIn := *openapiclient.NewMakeIn() // MakeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PostMqStreamsByStreamConsumers(context.Background(), stream).MakeIn(makeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PostMqStreamsByStreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMqStreamsByStreamConsumers`: Consumer
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PostMqStreamsByStreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMqStreamsByStreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **makeIn** | [**MakeIn**](MakeIn.md) |  | 

### Return type

[**Consumer**](Consumer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMqStreamsByStreamConsumersByNameNext

> ReadOut PostMqStreamsByStreamConsumersByNameNext(ctx, stream, name).NextIn(nextIn).Execute()

Pulls the consumer's next batch.



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
	stream := "stream_example" // string | Stream is the stream name, from the path.
	name := "name_example" // string | Name is the consumer name, from the path.
	nextIn := *openapiclient.NewNextIn() // NextIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PostMqStreamsByStreamConsumersByNameNext(context.Background(), stream, name).NextIn(nextIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PostMqStreamsByStreamConsumersByNameNext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMqStreamsByStreamConsumersByNameNext`: ReadOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PostMqStreamsByStreamConsumersByNameNext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 
**name** | **string** | Name is the consumer name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMqStreamsByStreamConsumersByNameNextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nextIn** | [**NextIn**](NextIn.md) |  | 

### Return type

[**ReadOut**](ReadOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMqStreamsByName

> Stream PutMqStreamsByName(ctx, name).StreamConfig(streamConfig).Execute()

Reconfigures an existing stream; the path names the stream, and the immutable fields (storage, retention) must restate what they are.



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
	name := "name_example" // string | Name is the stream name, unique within the org (alphanumeric, hyphens, underscores).
	streamConfig := *openapiclient.NewStreamConfig() // StreamConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.PutMqStreamsByName(context.Background(), name).StreamConfig(streamConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.PutMqStreamsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMqStreamsByName`: Stream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.PutMqStreamsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, unique within the org (alphanumeric, hyphens, underscores). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMqStreamsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamConfig** | [**StreamConfig**](StreamConfig.md) |  | 

### Return type

[**Stream**](Stream.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

