# \MqAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1MqStreamsName**](MqAPI.md#CloudDeleteV1MqStreamsName) | **Delete** /v1/mq/streams/{name} | Removes a stream with all its messages and consumers.
[**CloudDeleteV1MqStreamsNameMessagesSeq**](MqAPI.md#CloudDeleteV1MqStreamsNameMessagesSeq) | **Delete** /v1/mq/streams/{name}/messages/{seq} | Erases one message by sequence; the sequence gap remains.
[**CloudDeleteV1MqStreamsStreamConsumersName**](MqAPI.md#CloudDeleteV1MqStreamsStreamConsumersName) | **Delete** /v1/mq/streams/{stream}/consumers/{name} | Removes a consumer and its delivery state; unacknowledged messages stay in the stream.
[**CloudGetV1MqHealth**](MqAPI.md#CloudGetV1MqHealth) | **Get** /v1/mq/health | Reports whether the message plane behind this surface answers.
[**CloudGetV1MqInfo**](MqAPI.md#CloudGetV1MqInfo) | **Get** /v1/mq/info | Returns the broker&#39;s identity and the org&#39;s stream count.
[**CloudGetV1MqStreams**](MqAPI.md#CloudGetV1MqStreams) | **Get** /v1/mq/streams | Returns the org&#39;s streams, name-ordered, with their live state.
[**CloudGetV1MqStreamsName**](MqAPI.md#CloudGetV1MqStreamsName) | **Get** /v1/mq/streams/{name} | Returns one stream&#39;s configuration and live state.
[**CloudGetV1MqStreamsNameMessages**](MqAPI.md#CloudGetV1MqStreamsNameMessages) | **Get** /v1/mq/streams/{name}/messages | Reads stored messages without a consumer: by sequence, by newest on a subject, or walking a subject forward from a sequence.
[**CloudGetV1MqStreamsStreamConsumers**](MqAPI.md#CloudGetV1MqStreamsStreamConsumers) | **Get** /v1/mq/streams/{stream}/consumers | Returns a stream&#39;s consumers, name-ordered, with delivery state.
[**CloudGetV1MqStreamsStreamConsumersName**](MqAPI.md#CloudGetV1MqStreamsStreamConsumersName) | **Get** /v1/mq/streams/{stream}/consumers/{name} | Returns one consumer&#39;s configuration and delivery state.
[**CloudPostV1MqStreams**](MqAPI.md#CloudPostV1MqStreams) | **Post** /v1/mq/streams | Creates a durable stream in the org&#39;s namespace and returns it.
[**CloudPostV1MqStreamsNamePurge**](MqAPI.md#CloudPostV1MqStreamsNamePurge) | **Post** /v1/mq/streams/{name}/purge | Removes messages from a stream, leaving its consumers in place.
[**CloudPostV1MqStreamsStreamConsumers**](MqAPI.md#CloudPostV1MqStreamsStreamConsumers) | **Post** /v1/mq/streams/{stream}/consumers | Creates a durable pull consumer on a stream and returns it.
[**CloudPostV1MqStreamsStreamConsumersNameNext**](MqAPI.md#CloudPostV1MqStreamsStreamConsumersNameNext) | **Post** /v1/mq/streams/{stream}/consumers/{name}/next | Pulls the consumer&#39;s next batch.
[**CloudPutV1MqStreamsName**](MqAPI.md#CloudPutV1MqStreamsName) | **Put** /v1/mq/streams/{name} | Reconfigures an existing stream; the path names the stream, and the immutable fields (storage, retention) must restate what they are.



## CloudDeleteV1MqStreamsName

> CloudDeleteV1MqStreamsName(ctx, name).Execute()

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
	r, err := apiClient.MqAPI.CloudDeleteV1MqStreamsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudDeleteV1MqStreamsName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1MqStreamsNameRequest struct via the builder pattern


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


## CloudDeleteV1MqStreamsNameMessagesSeq

> CloudDeleteV1MqStreamsNameMessagesSeq(ctx, name, seq).Execute()

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
	r, err := apiClient.MqAPI.CloudDeleteV1MqStreamsNameMessagesSeq(context.Background(), name, seq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudDeleteV1MqStreamsNameMessagesSeq``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1MqStreamsNameMessagesSeqRequest struct via the builder pattern


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


## CloudDeleteV1MqStreamsStreamConsumersName

> CloudDeleteV1MqStreamsStreamConsumersName(ctx, stream, name).Execute()

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
	r, err := apiClient.MqAPI.CloudDeleteV1MqStreamsStreamConsumersName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudDeleteV1MqStreamsStreamConsumersName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1MqStreamsStreamConsumersNameRequest struct via the builder pattern


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


## CloudGetV1MqHealth

> CloudHealth CloudGetV1MqHealth(ctx).Execute()

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
	resp, r, err := apiClient.MqAPI.CloudGetV1MqHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudGetV1MqHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MqHealth`: CloudHealth
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudGetV1MqHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MqHealthRequest struct via the builder pattern


### Return type

[**CloudHealth**](CloudHealth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MqInfo

> CloudInfoOut CloudGetV1MqInfo(ctx).Execute()

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
	resp, r, err := apiClient.MqAPI.CloudGetV1MqInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudGetV1MqInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MqInfo`: CloudInfoOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudGetV1MqInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MqInfoRequest struct via the builder pattern


### Return type

[**CloudInfoOut**](CloudInfoOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MqStreams

> CloudStreams CloudGetV1MqStreams(ctx).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.MqAPI.CloudGetV1MqStreams(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudGetV1MqStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MqStreams`: CloudStreams
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudGetV1MqStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MqStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the streams returned (1–1000, default 100). | 
 **offset** | **int32** | Offset skips that many streams, name-ordered. | 

### Return type

[**CloudStreams**](CloudStreams.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MqStreamsName

> CloudStream CloudGetV1MqStreamsName(ctx, name).Execute()

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
	resp, r, err := apiClient.MqAPI.CloudGetV1MqStreamsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudGetV1MqStreamsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MqStreamsName`: CloudStream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudGetV1MqStreamsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MqStreamsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudStream**](CloudStream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MqStreamsNameMessages

> CloudReadOut CloudGetV1MqStreamsNameMessages(ctx, name).Seq(seq).LastBySubject(lastBySubject).NextBySubject(nextBySubject).Limit(limit).Execute()

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
	resp, r, err := apiClient.MqAPI.CloudGetV1MqStreamsNameMessages(context.Background(), name).Seq(seq).LastBySubject(lastBySubject).NextBySubject(nextBySubject).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudGetV1MqStreamsNameMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MqStreamsNameMessages`: CloudReadOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudGetV1MqStreamsNameMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MqStreamsNameMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seq** | **int32** | Seq reads the message at this sequence (with next_by_subject: the walk&#39;s start). | 
 **lastBySubject** | **string** | LastBySubject reads the newest message on this org-relative subject. | 
 **nextBySubject** | **string** | NextBySubject walks forward from seq collecting messages on this org-relative subject (wildcards supported). | 
 **limit** | **int32** | Limit caps a next_by_subject walk (1–1000, default 100). | 

### Return type

[**CloudReadOut**](CloudReadOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MqStreamsStreamConsumers

> CloudPickOut CloudGetV1MqStreamsStreamConsumers(ctx, stream).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.MqAPI.CloudGetV1MqStreamsStreamConsumers(context.Background(), stream).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudGetV1MqStreamsStreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MqStreamsStreamConsumers`: CloudPickOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudGetV1MqStreamsStreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MqStreamsStreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps the consumers returned (1–1000, default 100). | 
 **offset** | **int32** | Offset skips that many consumers, name-ordered. | 

### Return type

[**CloudPickOut**](CloudPickOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MqStreamsStreamConsumersName

> CloudConsumer CloudGetV1MqStreamsStreamConsumersName(ctx, stream, name).Execute()

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
	resp, r, err := apiClient.MqAPI.CloudGetV1MqStreamsStreamConsumersName(context.Background(), stream, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudGetV1MqStreamsStreamConsumersName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MqStreamsStreamConsumersName`: CloudConsumer
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudGetV1MqStreamsStreamConsumersName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 
**name** | **string** | Name is the consumer name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MqStreamsStreamConsumersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudConsumer**](CloudConsumer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MqStreams

> CloudStream CloudPostV1MqStreams(ctx).CloudConfig(cloudConfig).Execute()

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
	cloudConfig := *openapiclient.NewCloudConfig() // CloudConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.CloudPostV1MqStreams(context.Background()).CloudConfig(cloudConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudPostV1MqStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MqStreams`: CloudStream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudPostV1MqStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MqStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudConfig** | [**CloudConfig**](CloudConfig.md) |  | 

### Return type

[**CloudStream**](CloudStream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MqStreamsNamePurge

> CloudPurgeOut CloudPostV1MqStreamsNamePurge(ctx, name).CloudPurge(cloudPurge).Execute()

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
	cloudPurge := *openapiclient.NewCloudPurge() // CloudPurge | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.CloudPostV1MqStreamsNamePurge(context.Background(), name).CloudPurge(cloudPurge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudPostV1MqStreamsNamePurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MqStreamsNamePurge`: CloudPurgeOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudPostV1MqStreamsNamePurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MqStreamsNamePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPurge** | [**CloudPurge**](CloudPurge.md) |  | 

### Return type

[**CloudPurgeOut**](CloudPurgeOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MqStreamsStreamConsumers

> CloudConsumer CloudPostV1MqStreamsStreamConsumers(ctx, stream).CloudMakeIn(cloudMakeIn).Execute()

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
	cloudMakeIn := *openapiclient.NewCloudMakeIn() // CloudMakeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.CloudPostV1MqStreamsStreamConsumers(context.Background(), stream).CloudMakeIn(cloudMakeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudPostV1MqStreamsStreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MqStreamsStreamConsumers`: CloudConsumer
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudPostV1MqStreamsStreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MqStreamsStreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudMakeIn** | [**CloudMakeIn**](CloudMakeIn.md) |  | 

### Return type

[**CloudConsumer**](CloudConsumer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MqStreamsStreamConsumersNameNext

> CloudReadOut CloudPostV1MqStreamsStreamConsumersNameNext(ctx, stream, name).CloudNextIn(cloudNextIn).Execute()

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
	cloudNextIn := *openapiclient.NewCloudNextIn() // CloudNextIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.CloudPostV1MqStreamsStreamConsumersNameNext(context.Background(), stream, name).CloudNextIn(cloudNextIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudPostV1MqStreamsStreamConsumersNameNext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MqStreamsStreamConsumersNameNext`: CloudReadOut
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudPostV1MqStreamsStreamConsumersNameNext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | Stream is the stream name, from the path. | 
**name** | **string** | Name is the consumer name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MqStreamsStreamConsumersNameNextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudNextIn** | [**CloudNextIn**](CloudNextIn.md) |  | 

### Return type

[**CloudReadOut**](CloudReadOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1MqStreamsName

> CloudStream CloudPutV1MqStreamsName(ctx, name).CloudConfig(cloudConfig).Execute()

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
	cloudConfig := *openapiclient.NewCloudConfig() // CloudConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAPI.CloudPutV1MqStreamsName(context.Background(), name).CloudConfig(cloudConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAPI.CloudPutV1MqStreamsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1MqStreamsName`: CloudStream
	fmt.Fprintf(os.Stdout, "Response from `MqAPI.CloudPutV1MqStreamsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the stream name, unique within the org (alphanumeric, hyphens, underscores). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1MqStreamsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudConfig** | [**CloudConfig**](CloudConfig.md) |  | 

### Return type

[**CloudStream**](CloudStream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

