# \PubsubAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePubsubJetstreamStreamsByStream**](PubsubAPI.md#DeletePubsubJetstreamStreamsByStream) | **Delete** /v1/pubsub/jetstream/streams/{stream} | Removes one stream of the caller&#39;s org — its retained messages and its consumers with it — and answers 204 with no body.
[**DeletePubsubJetstreamStreamsByStreamConsumersByName**](PubsubAPI.md#DeletePubsubJetstreamStreamsByStreamConsumersByName) | **Delete** /v1/pubsub/jetstream/streams/{stream}/consumers/{name} | Removes one consumer — its cursor, not the stream&#39;s messages — and answers 204 with no body.
[**DeletePubsubKvByBucket**](PubsubAPI.md#DeletePubsubKvByBucket) | **Delete** /v1/pubsub/kv/{bucket} | Removes one bucket of the caller&#39;s org — every key and every revision with it — and answers 204 with no body.
[**DeletePubsubKvByBucketByKey**](PubsubAPI.md#DeletePubsubKvByBucketByKey) | **Delete** /v1/pubsub/kv/{bucket}/{key} | Delete removes one key — a delete marker in the key&#39;s history, so watchers see it and Get answers 404 — and answers 204 with no body.
[**GetPubsubJetstreamStreams**](PubsubAPI.md#GetPubsubJetstreamStreams) | **Get** /v1/pubsub/jetstream/streams | Returns the org&#39;s streams, sorted by name.
[**GetPubsubJetstreamStreamsByStream**](PubsubAPI.md#GetPubsubJetstreamStreamsByStream) | **Get** /v1/pubsub/jetstream/streams/{stream} | Returns one stream of the caller&#39;s org — its configuration and its live state (messages, bytes, sequence range, consumer count).
[**GetPubsubJetstreamStreamsByStreamConsumers**](PubsubAPI.md#GetPubsubJetstreamStreamsByStreamConsumers) | **Get** /v1/pubsub/jetstream/streams/{stream}/consumers | Returns one stream&#39;s consumers, sorted by name.
[**GetPubsubJetstreamStreamsByStreamConsumersByName**](PubsubAPI.md#GetPubsubJetstreamStreamsByStreamConsumersByName) | **Get** /v1/pubsub/jetstream/streams/{stream}/consumers/{name} | Returns one consumer of one org stream — its configuration and its cursor: delivered and acknowledged sequences, pending and redelivered counts.
[**GetPubsubKvByBucketByKey**](PubsubAPI.md#GetPubsubKvByBucketByKey) | **Get** /v1/pubsub/kv/{bucket}/{key} | Get returns one key&#39;s current value and revision.
[**GetPubsubKvByBucketByKeyHistory**](PubsubAPI.md#GetPubsubKvByBucketByKeyHistory) | **Get** /v1/pubsub/kv/{bucket}/{key}/history | History returns one key&#39;s retained revisions, oldest first — every put and every delete marker up to the bucket&#39;s History depth.
[**PostPubsubJetstreamStreams**](PubsubAPI.md#PostPubsubJetstreamStreams) | **Post** /v1/pubsub/jetstream/streams | Creates a durable stream capturing the given subjects and returns it.
[**PostPubsubJetstreamStreamsByStreamConsumers**](PubsubAPI.md#PostPubsubJetstreamStreamsByStreamConsumers) | **Post** /v1/pubsub/jetstream/streams/{stream}/consumers | Creates a durable consumer on one stream and returns it.
[**PostPubsubJetstreamStreamsByStreamConsumersByNameNext**](PubsubAPI.md#PostPubsubJetstreamStreamsByStreamConsumersByNameNext) | **Post** /v1/pubsub/jetstream/streams/{stream}/consumers/{name}/next | Fetch pulls the next batch from a consumer and acknowledges it — the request/response way to consume a stream.
[**PostPubsubKvByBucket**](PubsubAPI.md#PostPubsubKvByBucket) | **Post** /v1/pubsub/kv/{bucket} | Creates a KV bucket and returns it.
[**PostPubsubPublish**](PubsubAPI.md#PostPubsubPublish) | **Post** /v1/pubsub/publish | Publish puts one message on the org&#39;s bus.
[**PostPubsubRequest**](PubsubAPI.md#PostPubsubRequest) | **Post** /v1/pubsub/request | Request sends one request on the org&#39;s bus and waits for one reply — the synchronous half of pub/sub, for callers speaking to a responder subscribed on the NATS port.
[**PutPubsubJetstreamStreamsByStream**](PubsubAPI.md#PutPubsubJetstreamStreamsByStream) | **Put** /v1/pubsub/jetstream/streams/{stream} | Rewrites a stream&#39;s configuration — subjects, limits, discard — and returns the updated stream.
[**PutPubsubKvByBucketByKey**](PubsubAPI.md#PutPubsubKvByBucketByKey) | **Put** /v1/pubsub/kv/{bucket}/{key} | Put sets one key to one value and returns the revision the write created.



## DeletePubsubJetstreamStreamsByStream

> DeletePubsubJetstreamStreamsByStream(ctx, stream).Execute()

Removes one stream of the caller's org — its retained messages and its consumers with it — and answers 204 with no body.



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
	stream := "stream_example" // string | Stream is the stream's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubsubAPI.DeletePubsubJetstreamStreamsByStream(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.DeletePubsubJetstreamStreamsByStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePubsubJetstreamStreamsByStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePubsubJetstreamStreamsByStreamConsumersByName

> DeletePubsubJetstreamStreamsByStreamConsumersByName(ctx, stream, name).Execute()

Removes one consumer — its cursor, not the stream's messages — and answers 204 with no body.



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
	stream := "stream_example" // string | Stream is the stream, from the path.
	name := "name_example" // string | Name is the consumer, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubsubAPI.DeletePubsubJetstreamStreamsByStreamConsumersByName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.DeletePubsubJetstreamStreamsByStreamConsumersByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream, from the path. | 
**name** | **string** | Name is the consumer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePubsubJetstreamStreamsByStreamConsumersByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePubsubKvByBucket

> DeletePubsubKvByBucket(ctx, bucket).Execute()

Removes one bucket of the caller's org — every key and every revision with it — and answers 204 with no body.



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
	bucket := "bucket_example" // string | Bucket is the bucket's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubsubAPI.DeletePubsubKvByBucket(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.DeletePubsubKvByBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePubsubKvByBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePubsubKvByBucketByKey

> DeletePubsubKvByBucketByKey(ctx, bucket, key).Execute()

Delete removes one key — a delete marker in the key's history, so watchers see it and Get answers 404 — and answers 204 with no body.



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
	bucket := "bucket_example" // string | Bucket is the bucket, from the path.
	key := "key_example" // string | Key is the key, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubsubAPI.DeletePubsubKvByBucketByKey(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.DeletePubsubKvByBucketByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePubsubKvByBucketByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPubsubJetstreamStreams

> StreamPage GetPubsubJetstreamStreams(ctx).Execute()

Returns the org's streams, sorted by name.



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
	resp, r, err := apiClient.PubsubAPI.GetPubsubJetstreamStreams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.GetPubsubJetstreamStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPubsubJetstreamStreams`: StreamPage
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.GetPubsubJetstreamStreams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPubsubJetstreamStreamsRequest struct via the builder pattern


### Return type

[**StreamPage**](StreamPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPubsubJetstreamStreamsByStream

> StreamRecord GetPubsubJetstreamStreamsByStream(ctx, stream).Execute()

Returns one stream of the caller's org — its configuration and its live state (messages, bytes, sequence range, consumer count).



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
	stream := "stream_example" // string | Stream is the stream's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.GetPubsubJetstreamStreamsByStream(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.GetPubsubJetstreamStreamsByStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPubsubJetstreamStreamsByStream`: StreamRecord
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.GetPubsubJetstreamStreamsByStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPubsubJetstreamStreamsByStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamRecord**](StreamRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPubsubJetstreamStreamsByStreamConsumers

> ConsumerPage GetPubsubJetstreamStreamsByStreamConsumers(ctx, stream).Execute()

Returns one stream's consumers, sorted by name.



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
	stream := "stream_example" // string | Stream is the stream's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.GetPubsubJetstreamStreamsByStreamConsumers(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.GetPubsubJetstreamStreamsByStreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPubsubJetstreamStreamsByStreamConsumers`: ConsumerPage
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.GetPubsubJetstreamStreamsByStreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPubsubJetstreamStreamsByStreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumerPage**](ConsumerPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPubsubJetstreamStreamsByStreamConsumersByName

> ConsumerRecord GetPubsubJetstreamStreamsByStreamConsumersByName(ctx, stream, name).Execute()

Returns one consumer of one org stream — its configuration and its cursor: delivered and acknowledged sequences, pending and redelivered counts.



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
	stream := "stream_example" // string | Stream is the stream, from the path.
	name := "name_example" // string | Name is the consumer, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.GetPubsubJetstreamStreamsByStreamConsumersByName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.GetPubsubJetstreamStreamsByStreamConsumersByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPubsubJetstreamStreamsByStreamConsumersByName`: ConsumerRecord
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.GetPubsubJetstreamStreamsByStreamConsumersByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream, from the path. | 
**name** | **string** | Name is the consumer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPubsubJetstreamStreamsByStreamConsumersByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsumerRecord**](ConsumerRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPubsubKvByBucketByKey

> KvEntry GetPubsubKvByBucketByKey(ctx, bucket, key).Execute()

Get returns one key's current value and revision.



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
	bucket := "bucket_example" // string | Bucket is the bucket, from the path.
	key := "key_example" // string | Key is the key, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.GetPubsubKvByBucketByKey(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.GetPubsubKvByBucketByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPubsubKvByBucketByKey`: KvEntry
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.GetPubsubKvByBucketByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPubsubKvByBucketByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KvEntry**](KvEntry.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPubsubKvByBucketByKeyHistory

> KvPage GetPubsubKvByBucketByKeyHistory(ctx, bucket, key).Execute()

History returns one key's retained revisions, oldest first — every put and every delete marker up to the bucket's History depth.



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
	bucket := "bucket_example" // string | Bucket is the bucket, from the path.
	key := "key_example" // string | Key is the key, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.GetPubsubKvByBucketByKeyHistory(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.GetPubsubKvByBucketByKeyHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPubsubKvByBucketByKeyHistory`: KvPage
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.GetPubsubKvByBucketByKeyHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPubsubKvByBucketByKeyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KvPage**](KvPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPubsubJetstreamStreams

> StreamRecord PostPubsubJetstreamStreams(ctx).StreamWrite(streamWrite).Execute()

Creates a durable stream capturing the given subjects and returns it.



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
	streamWrite := *openapiclient.NewStreamWrite() // StreamWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PostPubsubJetstreamStreams(context.Background()).StreamWrite(streamWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PostPubsubJetstreamStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPubsubJetstreamStreams`: StreamRecord
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PostPubsubJetstreamStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPubsubJetstreamStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamWrite** | [**StreamWrite**](StreamWrite.md) |  | 

### Return type

[**StreamRecord**](StreamRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPubsubJetstreamStreamsByStreamConsumers

> ConsumerRecord PostPubsubJetstreamStreamsByStreamConsumers(ctx, stream).ConsumerWrite(consumerWrite).Execute()

Creates a durable consumer on one stream and returns it.



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
	stream := "stream_example" // string | Stream is the stream to consume, from the path.
	consumerWrite := *openapiclient.NewConsumerWrite() // ConsumerWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PostPubsubJetstreamStreamsByStreamConsumers(context.Background(), stream).ConsumerWrite(consumerWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PostPubsubJetstreamStreamsByStreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPubsubJetstreamStreamsByStreamConsumers`: ConsumerRecord
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PostPubsubJetstreamStreamsByStreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream to consume, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPubsubJetstreamStreamsByStreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumerWrite** | [**ConsumerWrite**](ConsumerWrite.md) |  | 

### Return type

[**ConsumerRecord**](ConsumerRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPubsubJetstreamStreamsByStreamConsumersByNameNext

> MessagePage PostPubsubJetstreamStreamsByStreamConsumersByNameNext(ctx, stream, name).FetchQuery(fetchQuery).Execute()

Fetch pulls the next batch from a consumer and acknowledges it — the request/response way to consume a stream.



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
	stream := "stream_example" // string | Stream is the stream, from the path.
	name := "name_example" // string | Name is the consumer, from the path.
	fetchQuery := *openapiclient.NewFetchQuery() // FetchQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PostPubsubJetstreamStreamsByStreamConsumersByNameNext(context.Background(), stream, name).FetchQuery(fetchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PostPubsubJetstreamStreamsByStreamConsumersByNameNext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPubsubJetstreamStreamsByStreamConsumersByNameNext`: MessagePage
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PostPubsubJetstreamStreamsByStreamConsumersByNameNext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream, from the path. | 
**name** | **string** | Name is the consumer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPubsubJetstreamStreamsByStreamConsumersByNameNextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fetchQuery** | [**FetchQuery**](FetchQuery.md) |  | 

### Return type

[**MessagePage**](MessagePage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPubsubKvByBucket

> BucketRecord PostPubsubKvByBucket(ctx, bucket).BucketWrite(bucketWrite).Execute()

Creates a KV bucket and returns it.



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
	bucket := "bucket_example" // string | Bucket is the bucket's name within the org, from the path: 1–64 of [A-Za-z0-9_], no dash.
	bucketWrite := *openapiclient.NewBucketWrite() // BucketWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PostPubsubKvByBucket(context.Background(), bucket).BucketWrite(bucketWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PostPubsubKvByBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPubsubKvByBucket`: BucketRecord
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PostPubsubKvByBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket&#39;s name within the org, from the path: 1–64 of [A-Za-z0-9_], no dash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPubsubKvByBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketWrite** | [**BucketWrite**](BucketWrite.md) |  | 

### Return type

[**BucketRecord**](BucketRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPubsubPublish

> BusAck PostPubsubPublish(ctx).BusPublish(busPublish).Execute()

Publish puts one message on the org's bus.



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
	busPublish := *openapiclient.NewBusPublish() // BusPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PostPubsubPublish(context.Background()).BusPublish(busPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PostPubsubPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPubsubPublish`: BusAck
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PostPubsubPublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPubsubPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **busPublish** | [**BusPublish**](BusPublish.md) |  | 

### Return type

[**BusAck**](BusAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPubsubRequest

> BusMessage PostPubsubRequest(ctx).BusRequest(busRequest).Execute()

Request sends one request on the org's bus and waits for one reply — the synchronous half of pub/sub, for callers speaking to a responder subscribed on the NATS port.



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
	busRequest := *openapiclient.NewBusRequest() // BusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PostPubsubRequest(context.Background()).BusRequest(busRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PostPubsubRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPubsubRequest`: BusMessage
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PostPubsubRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPubsubRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **busRequest** | [**BusRequest**](BusRequest.md) |  | 

### Return type

[**BusMessage**](BusMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPubsubJetstreamStreamsByStream

> StreamRecord PutPubsubJetstreamStreamsByStream(ctx, stream).StreamUpdate(streamUpdate).Execute()

Rewrites a stream's configuration — subjects, limits, discard — and returns the updated stream.



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
	stream := "stream_example" // string | Stream is the stream to update, from the path.
	streamUpdate := *openapiclient.NewStreamUpdate() // StreamUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PutPubsubJetstreamStreamsByStream(context.Background(), stream).StreamUpdate(streamUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PutPubsubJetstreamStreamsByStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPubsubJetstreamStreamsByStream`: StreamRecord
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PutPubsubJetstreamStreamsByStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPubsubJetstreamStreamsByStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamUpdate** | [**StreamUpdate**](StreamUpdate.md) |  | 

### Return type

[**StreamRecord**](StreamRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPubsubKvByBucketByKey

> KvAck PutPubsubKvByBucketByKey(ctx, bucket, key).KvWrite(kvWrite).Execute()

Put sets one key to one value and returns the revision the write created.



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
	bucket := "bucket_example" // string | Bucket is the bucket, from the path.
	key := "key_example" // string | Key is the key, from the path.
	kvWrite := *openapiclient.NewKvWrite() // KvWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PutPubsubKvByBucketByKey(context.Background(), bucket, key).KvWrite(kvWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PutPubsubKvByBucketByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPubsubKvByBucketByKey`: KvAck
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PutPubsubKvByBucketByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPubsubKvByBucketByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvWrite** | [**KvWrite**](KvWrite.md) |  | 

### Return type

[**KvAck**](KvAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

