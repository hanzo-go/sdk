# \PubSubAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1PubsubJetstreamStreamsStream**](PubSubAPI.md#CloudDeleteV1PubsubJetstreamStreamsStream) | **Delete** /v1/pubsub/jetstream/streams/{stream} | DeleteStream removes one stream of the caller&#39;s org — its retained messages and its consumers with it — and answers 204 with no body.
[**CloudDeleteV1PubsubJetstreamStreamsStreamConsumersName**](PubSubAPI.md#CloudDeleteV1PubsubJetstreamStreamsStreamConsumersName) | **Delete** /v1/pubsub/jetstream/streams/{stream}/consumers/{name} | DeleteConsumer removes one consumer — its cursor, not the stream&#39;s messages — and answers 204 with no body.
[**CloudDeleteV1PubsubKvBucket**](PubSubAPI.md#CloudDeleteV1PubsubKvBucket) | **Delete** /v1/pubsub/kv/{bucket} | DeleteBucket removes one bucket of the caller&#39;s org — every key and every revision with it — and answers 204 with no body.
[**CloudDeleteV1PubsubKvBucketKey**](PubSubAPI.md#CloudDeleteV1PubsubKvBucketKey) | **Delete** /v1/pubsub/kv/{bucket}/{key} | Delete removes one key — a delete marker in the key&#39;s history, so watchers see it and Get answers 404 — and answers 204 with no body.
[**CloudGetV1PubsubJetstreamStreams**](PubSubAPI.md#CloudGetV1PubsubJetstreamStreams) | **Get** /v1/pubsub/jetstream/streams | ListStreams returns the org&#39;s streams, sorted by name.
[**CloudGetV1PubsubJetstreamStreamsStream**](PubSubAPI.md#CloudGetV1PubsubJetstreamStreamsStream) | **Get** /v1/pubsub/jetstream/streams/{stream} | GetStream returns one stream of the caller&#39;s org — its configuration and its live state (messages, bytes, sequence range, consumer count).
[**CloudGetV1PubsubJetstreamStreamsStreamConsumers**](PubSubAPI.md#CloudGetV1PubsubJetstreamStreamsStreamConsumers) | **Get** /v1/pubsub/jetstream/streams/{stream}/consumers | ListConsumers returns one stream&#39;s consumers, sorted by name.
[**CloudGetV1PubsubJetstreamStreamsStreamConsumersName**](PubSubAPI.md#CloudGetV1PubsubJetstreamStreamsStreamConsumersName) | **Get** /v1/pubsub/jetstream/streams/{stream}/consumers/{name} | GetConsumer returns one consumer of one org stream — its configuration and its cursor: delivered and acknowledged sequences, pending and redelivered counts.
[**CloudGetV1PubsubKvBucketKey**](PubSubAPI.md#CloudGetV1PubsubKvBucketKey) | **Get** /v1/pubsub/kv/{bucket}/{key} | Get returns one key&#39;s current value and revision.
[**CloudGetV1PubsubKvBucketKeyHistory**](PubSubAPI.md#CloudGetV1PubsubKvBucketKeyHistory) | **Get** /v1/pubsub/kv/{bucket}/{key}/history | History returns one key&#39;s retained revisions, oldest first — every put and every delete marker up to the bucket&#39;s History depth.
[**CloudPostV1PubsubJetstreamStreams**](PubSubAPI.md#CloudPostV1PubsubJetstreamStreams) | **Post** /v1/pubsub/jetstream/streams | CreateStream creates a durable stream capturing the given subjects and returns it.
[**CloudPostV1PubsubJetstreamStreamsStreamConsumers**](PubSubAPI.md#CloudPostV1PubsubJetstreamStreamsStreamConsumers) | **Post** /v1/pubsub/jetstream/streams/{stream}/consumers | CreateConsumer creates a durable consumer on one stream and returns it.
[**CloudPostV1PubsubJetstreamStreamsStreamConsumersNameNext**](PubSubAPI.md#CloudPostV1PubsubJetstreamStreamsStreamConsumersNameNext) | **Post** /v1/pubsub/jetstream/streams/{stream}/consumers/{name}/next | Fetch pulls the next batch from a consumer and acknowledges it — the request/response way to consume a stream.
[**CloudPostV1PubsubKvBucket**](PubSubAPI.md#CloudPostV1PubsubKvBucket) | **Post** /v1/pubsub/kv/{bucket} | CreateBucket creates a KV bucket and returns it.
[**CloudPostV1PubsubPublish**](PubSubAPI.md#CloudPostV1PubsubPublish) | **Post** /v1/pubsub/publish | Publish puts one message on the org&#39;s bus.
[**CloudPostV1PubsubRequest**](PubSubAPI.md#CloudPostV1PubsubRequest) | **Post** /v1/pubsub/request | Request sends one request on the org&#39;s bus and waits for one reply — the synchronous half of pub/sub, for callers speaking to a responder subscribed on the NATS port.
[**CloudPutV1PubsubJetstreamStreamsStream**](PubSubAPI.md#CloudPutV1PubsubJetstreamStreamsStream) | **Put** /v1/pubsub/jetstream/streams/{stream} | UpdateStream rewrites a stream&#39;s configuration — subjects, limits, discard — and returns the updated stream.
[**CloudPutV1PubsubKvBucketKey**](PubSubAPI.md#CloudPutV1PubsubKvBucketKey) | **Put** /v1/pubsub/kv/{bucket}/{key} | Put sets one key to one value and returns the revision the write created.
[**KvListChannels**](PubSubAPI.md#KvListChannels) | **Get** /v1/kv/pubsub/channels | List active channels
[**KvPublish**](PubSubAPI.md#KvPublish) | **Post** /v1/kv/pubsub/publish | Publish message to channel
[**KvSubscribe**](PubSubAPI.md#KvSubscribe) | **Get** /v1/kv/pubsub/subscribe | Subscribe to channels (SSE)



## CloudDeleteV1PubsubJetstreamStreamsStream

> CloudDeleteV1PubsubJetstreamStreamsStream(ctx, stream).Execute()

DeleteStream removes one stream of the caller's org — its retained messages and its consumers with it — and answers 204 with no body.



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
	r, err := apiClient.PubSubAPI.CloudDeleteV1PubsubJetstreamStreamsStream(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudDeleteV1PubsubJetstreamStreamsStream``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PubsubJetstreamStreamsStreamRequest struct via the builder pattern


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


## CloudDeleteV1PubsubJetstreamStreamsStreamConsumersName

> CloudDeleteV1PubsubJetstreamStreamsStreamConsumersName(ctx, stream, name).Execute()

DeleteConsumer removes one consumer — its cursor, not the stream's messages — and answers 204 with no body.



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
	r, err := apiClient.PubSubAPI.CloudDeleteV1PubsubJetstreamStreamsStreamConsumersName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudDeleteV1PubsubJetstreamStreamsStreamConsumersName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PubsubJetstreamStreamsStreamConsumersNameRequest struct via the builder pattern


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


## CloudDeleteV1PubsubKvBucket

> CloudDeleteV1PubsubKvBucket(ctx, bucket).Execute()

DeleteBucket removes one bucket of the caller's org — every key and every revision with it — and answers 204 with no body.



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
	r, err := apiClient.PubSubAPI.CloudDeleteV1PubsubKvBucket(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudDeleteV1PubsubKvBucket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PubsubKvBucketRequest struct via the builder pattern


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


## CloudDeleteV1PubsubKvBucketKey

> CloudDeleteV1PubsubKvBucketKey(ctx, bucket, key).Execute()

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
	r, err := apiClient.PubSubAPI.CloudDeleteV1PubsubKvBucketKey(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudDeleteV1PubsubKvBucketKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PubsubKvBucketKeyRequest struct via the builder pattern


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


## CloudGetV1PubsubJetstreamStreams

> CloudStreamPage CloudGetV1PubsubJetstreamStreams(ctx).Execute()

ListStreams returns the org's streams, sorted by name.



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
	resp, r, err := apiClient.PubSubAPI.CloudGetV1PubsubJetstreamStreams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudGetV1PubsubJetstreamStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PubsubJetstreamStreams`: CloudStreamPage
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudGetV1PubsubJetstreamStreams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PubsubJetstreamStreamsRequest struct via the builder pattern


### Return type

[**CloudStreamPage**](CloudStreamPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PubsubJetstreamStreamsStream

> CloudStreamRecord CloudGetV1PubsubJetstreamStreamsStream(ctx, stream).Execute()

GetStream returns one stream of the caller's org — its configuration and its live state (messages, bytes, sequence range, consumer count).



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
	resp, r, err := apiClient.PubSubAPI.CloudGetV1PubsubJetstreamStreamsStream(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudGetV1PubsubJetstreamStreamsStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PubsubJetstreamStreamsStream`: CloudStreamRecord
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudGetV1PubsubJetstreamStreamsStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PubsubJetstreamStreamsStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudStreamRecord**](CloudStreamRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PubsubJetstreamStreamsStreamConsumers

> CloudConsumerPage CloudGetV1PubsubJetstreamStreamsStreamConsumers(ctx, stream).Execute()

ListConsumers returns one stream's consumers, sorted by name.



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
	resp, r, err := apiClient.PubSubAPI.CloudGetV1PubsubJetstreamStreamsStreamConsumers(context.Background(), stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudGetV1PubsubJetstreamStreamsStreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PubsubJetstreamStreamsStreamConsumers`: CloudConsumerPage
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudGetV1PubsubJetstreamStreamsStreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PubsubJetstreamStreamsStreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudConsumerPage**](CloudConsumerPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PubsubJetstreamStreamsStreamConsumersName

> CloudConsumerRecord CloudGetV1PubsubJetstreamStreamsStreamConsumersName(ctx, stream, name).Execute()

GetConsumer returns one consumer of one org stream — its configuration and its cursor: delivered and acknowledged sequences, pending and redelivered counts.



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
	resp, r, err := apiClient.PubSubAPI.CloudGetV1PubsubJetstreamStreamsStreamConsumersName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudGetV1PubsubJetstreamStreamsStreamConsumersName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PubsubJetstreamStreamsStreamConsumersName`: CloudConsumerRecord
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudGetV1PubsubJetstreamStreamsStreamConsumersName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream, from the path. | 
**name** | **string** | Name is the consumer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PubsubJetstreamStreamsStreamConsumersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudConsumerRecord**](CloudConsumerRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PubsubKvBucketKey

> CloudKvEntry CloudGetV1PubsubKvBucketKey(ctx, bucket, key).Execute()

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
	resp, r, err := apiClient.PubSubAPI.CloudGetV1PubsubKvBucketKey(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudGetV1PubsubKvBucketKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PubsubKvBucketKey`: CloudKvEntry
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudGetV1PubsubKvBucketKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PubsubKvBucketKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudKvEntry**](CloudKvEntry.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PubsubKvBucketKeyHistory

> CloudKvPage CloudGetV1PubsubKvBucketKeyHistory(ctx, bucket, key).Execute()

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
	resp, r, err := apiClient.PubSubAPI.CloudGetV1PubsubKvBucketKeyHistory(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudGetV1PubsubKvBucketKeyHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PubsubKvBucketKeyHistory`: CloudKvPage
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudGetV1PubsubKvBucketKeyHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PubsubKvBucketKeyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudKvPage**](CloudKvPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PubsubJetstreamStreams

> CloudStreamRecord CloudPostV1PubsubJetstreamStreams(ctx).CloudStreamWrite(cloudStreamWrite).Execute()

CreateStream creates a durable stream capturing the given subjects and returns it.



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
	cloudStreamWrite := *openapiclient.NewCloudStreamWrite() // CloudStreamWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.CloudPostV1PubsubJetstreamStreams(context.Background()).CloudStreamWrite(cloudStreamWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudPostV1PubsubJetstreamStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PubsubJetstreamStreams`: CloudStreamRecord
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudPostV1PubsubJetstreamStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PubsubJetstreamStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudStreamWrite** | [**CloudStreamWrite**](CloudStreamWrite.md) |  | 

### Return type

[**CloudStreamRecord**](CloudStreamRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PubsubJetstreamStreamsStreamConsumers

> CloudConsumerRecord CloudPostV1PubsubJetstreamStreamsStreamConsumers(ctx, stream).CloudConsumerWrite(cloudConsumerWrite).Execute()

CreateConsumer creates a durable consumer on one stream and returns it.



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
	cloudConsumerWrite := *openapiclient.NewCloudConsumerWrite() // CloudConsumerWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.CloudPostV1PubsubJetstreamStreamsStreamConsumers(context.Background(), stream).CloudConsumerWrite(cloudConsumerWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudPostV1PubsubJetstreamStreamsStreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PubsubJetstreamStreamsStreamConsumers`: CloudConsumerRecord
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudPostV1PubsubJetstreamStreamsStreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream to consume, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PubsubJetstreamStreamsStreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudConsumerWrite** | [**CloudConsumerWrite**](CloudConsumerWrite.md) |  | 

### Return type

[**CloudConsumerRecord**](CloudConsumerRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PubsubJetstreamStreamsStreamConsumersNameNext

> CloudMessagePage CloudPostV1PubsubJetstreamStreamsStreamConsumersNameNext(ctx, stream, name).CloudFetchQuery(cloudFetchQuery).Execute()

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
	cloudFetchQuery := *openapiclient.NewCloudFetchQuery() // CloudFetchQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.CloudPostV1PubsubJetstreamStreamsStreamConsumersNameNext(context.Background(), stream, name).CloudFetchQuery(cloudFetchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudPostV1PubsubJetstreamStreamsStreamConsumersNameNext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PubsubJetstreamStreamsStreamConsumersNameNext`: CloudMessagePage
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudPostV1PubsubJetstreamStreamsStreamConsumersNameNext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream, from the path. | 
**name** | **string** | Name is the consumer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PubsubJetstreamStreamsStreamConsumersNameNextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudFetchQuery** | [**CloudFetchQuery**](CloudFetchQuery.md) |  | 

### Return type

[**CloudMessagePage**](CloudMessagePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PubsubKvBucket

> CloudBucketRecord CloudPostV1PubsubKvBucket(ctx, bucket).CloudBucketWrite(cloudBucketWrite).Execute()

CreateBucket creates a KV bucket and returns it.



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
	cloudBucketWrite := *openapiclient.NewCloudBucketWrite() // CloudBucketWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.CloudPostV1PubsubKvBucket(context.Background(), bucket).CloudBucketWrite(cloudBucketWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudPostV1PubsubKvBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PubsubKvBucket`: CloudBucketRecord
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudPostV1PubsubKvBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket&#39;s name within the org, from the path: 1–64 of [A-Za-z0-9_], no dash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PubsubKvBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudBucketWrite** | [**CloudBucketWrite**](CloudBucketWrite.md) |  | 

### Return type

[**CloudBucketRecord**](CloudBucketRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PubsubPublish

> CloudBusAck CloudPostV1PubsubPublish(ctx).CloudBusPublish(cloudBusPublish).Execute()

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
	cloudBusPublish := *openapiclient.NewCloudBusPublish() // CloudBusPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.CloudPostV1PubsubPublish(context.Background()).CloudBusPublish(cloudBusPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudPostV1PubsubPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PubsubPublish`: CloudBusAck
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudPostV1PubsubPublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PubsubPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudBusPublish** | [**CloudBusPublish**](CloudBusPublish.md) |  | 

### Return type

[**CloudBusAck**](CloudBusAck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PubsubRequest

> CloudBusMessage CloudPostV1PubsubRequest(ctx).CloudBusRequest(cloudBusRequest).Execute()

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
	cloudBusRequest := *openapiclient.NewCloudBusRequest() // CloudBusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.CloudPostV1PubsubRequest(context.Background()).CloudBusRequest(cloudBusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudPostV1PubsubRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PubsubRequest`: CloudBusMessage
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudPostV1PubsubRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PubsubRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudBusRequest** | [**CloudBusRequest**](CloudBusRequest.md) |  | 

### Return type

[**CloudBusMessage**](CloudBusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1PubsubJetstreamStreamsStream

> CloudStreamRecord CloudPutV1PubsubJetstreamStreamsStream(ctx, stream).CloudStreamUpdate(cloudStreamUpdate).Execute()

UpdateStream rewrites a stream's configuration — subjects, limits, discard — and returns the updated stream.



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
	cloudStreamUpdate := *openapiclient.NewCloudStreamUpdate() // CloudStreamUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.CloudPutV1PubsubJetstreamStreamsStream(context.Background(), stream).CloudStreamUpdate(cloudStreamUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudPutV1PubsubJetstreamStreamsStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1PubsubJetstreamStreamsStream`: CloudStreamRecord
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudPutV1PubsubJetstreamStreamsStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1PubsubJetstreamStreamsStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudStreamUpdate** | [**CloudStreamUpdate**](CloudStreamUpdate.md) |  | 

### Return type

[**CloudStreamRecord**](CloudStreamRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1PubsubKvBucketKey

> CloudKvAck CloudPutV1PubsubKvBucketKey(ctx, bucket, key).CloudKvWrite(cloudKvWrite).Execute()

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
	cloudKvWrite := *openapiclient.NewCloudKvWrite() // CloudKvWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.CloudPutV1PubsubKvBucketKey(context.Background(), bucket, key).CloudKvWrite(cloudKvWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.CloudPutV1PubsubKvBucketKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1PubsubKvBucketKey`: CloudKvAck
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.CloudPutV1PubsubKvBucketKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1PubsubKvBucketKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudKvWrite** | [**CloudKvWrite**](CloudKvWrite.md) |  | 

### Return type

[**CloudKvAck**](CloudKvAck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvListChannels

> KvListChannels200Response KvListChannels(ctx).Pattern(pattern).Execute()

List active channels

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
	pattern := "pattern_example" // string |  (optional) (default to "*")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.KvListChannels(context.Background()).Pattern(pattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.KvListChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvListChannels`: KvListChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.KvListChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvListChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pattern** | **string** |  | [default to &quot;*&quot;]

### Return type

[**KvListChannels200Response**](KvListChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvPublish

> KvPublish200Response KvPublish(ctx).KvPublishRequest(kvPublishRequest).Execute()

Publish message to channel

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
	kvPublishRequest := *openapiclient.NewKvPublishRequest("Channel_example", "Message_example") // KvPublishRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.KvPublish(context.Background()).KvPublishRequest(kvPublishRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.KvPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvPublish`: KvPublish200Response
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.KvPublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvPublishRequest** | [**KvPublishRequest**](KvPublishRequest.md) |  | 

### Return type

[**KvPublish200Response**](KvPublish200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvSubscribe

> string KvSubscribe(ctx).Channels(channels).Pattern(pattern).Execute()

Subscribe to channels (SSE)

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
	channels := []string{"Inner_example"} // []string | Channels to subscribe to
	pattern := "pattern_example" // string | Pattern to subscribe to (e.g. user:*) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.KvSubscribe(context.Background()).Channels(channels).Pattern(pattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.KvSubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvSubscribe`: string
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.KvSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channels** | **[]string** | Channels to subscribe to | 
 **pattern** | **string** | Pattern to subscribe to (e.g. user:*) | 

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

