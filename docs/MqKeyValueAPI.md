# \MqKeyValueAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqCreateKVBucket**](MqKeyValueAPI.md#MqCreateKVBucket) | **Post** /v1/mq/kv | Create KV bucket
[**MqDeleteKVBucket**](MqKeyValueAPI.md#MqDeleteKVBucket) | **Delete** /v1/mq/kv/{bucket} | Delete bucket
[**MqDeleteKVEntry**](MqKeyValueAPI.md#MqDeleteKVEntry) | **Delete** /v1/mq/kv/{bucket}/{key} | Delete key
[**MqGetKVBucket**](MqKeyValueAPI.md#MqGetKVBucket) | **Get** /v1/mq/kv/{bucket} | Get bucket info
[**MqGetKVEntry**](MqKeyValueAPI.md#MqGetKVEntry) | **Get** /v1/mq/kv/{bucket}/{key} | Get value
[**MqGetKVHistory**](MqKeyValueAPI.md#MqGetKVHistory) | **Get** /v1/mq/kv/{bucket}/{key}/history | Get key history
[**MqListKVBuckets**](MqKeyValueAPI.md#MqListKVBuckets) | **Get** /v1/mq/kv | List KV buckets
[**MqListKVKeys**](MqKeyValueAPI.md#MqListKVKeys) | **Get** /v1/mq/kv/{bucket}/keys | List keys in bucket
[**MqPutKVEntry**](MqKeyValueAPI.md#MqPutKVEntry) | **Put** /v1/mq/kv/{bucket}/{key} | Put value
[**MqWatchKVBucket**](MqKeyValueAPI.md#MqWatchKVBucket) | **Get** /v1/mq/kv/{bucket}/watch | Watch bucket changes via SSE



## MqCreateKVBucket

> MqKVBucket MqCreateKVBucket(ctx).MqKVBucketConfig(mqKVBucketConfig).Execute()

Create KV bucket



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
	mqKVBucketConfig := *openapiclient.NewMqKVBucketConfig("Name_example") // MqKVBucketConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqKeyValueAPI.MqCreateKVBucket(context.Background()).MqKVBucketConfig(mqKVBucketConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqCreateKVBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqCreateKVBucket`: MqKVBucket
	fmt.Fprintf(os.Stdout, "Response from `MqKeyValueAPI.MqCreateKVBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqCreateKVBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mqKVBucketConfig** | [**MqKVBucketConfig**](MqKVBucketConfig.md) |  | 

### Return type

[**MqKVBucket**](MqKVBucket.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqDeleteKVBucket

> MqDeleteKVBucket(ctx, bucket).Execute()

Delete bucket



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
	bucket := "bucket_example" // string | KV bucket name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqKeyValueAPI.MqDeleteKVBucket(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqDeleteKVBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | KV bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqDeleteKVBucketRequest struct via the builder pattern


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


## MqDeleteKVEntry

> MqDeleteKVEntry(ctx, bucket, key).Execute()

Delete key



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
	bucket := "bucket_example" // string | KV bucket name.
	key := "key_example" // string | Key name (supports dotted hierarchy).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqKeyValueAPI.MqDeleteKVEntry(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqDeleteKVEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | KV bucket name. | 
**key** | **string** | Key name (supports dotted hierarchy). | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqDeleteKVEntryRequest struct via the builder pattern


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


## MqGetKVBucket

> MqKVBucket MqGetKVBucket(ctx, bucket).Execute()

Get bucket info



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
	bucket := "bucket_example" // string | KV bucket name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqKeyValueAPI.MqGetKVBucket(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqGetKVBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetKVBucket`: MqKVBucket
	fmt.Fprintf(os.Stdout, "Response from `MqKeyValueAPI.MqGetKVBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | KV bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetKVBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MqKVBucket**](MqKVBucket.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqGetKVEntry

> MqKVEntry MqGetKVEntry(ctx, bucket, key).Execute()

Get value



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
	bucket := "bucket_example" // string | KV bucket name.
	key := "key_example" // string | Key name (supports dotted hierarchy).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqKeyValueAPI.MqGetKVEntry(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqGetKVEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetKVEntry`: MqKVEntry
	fmt.Fprintf(os.Stdout, "Response from `MqKeyValueAPI.MqGetKVEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | KV bucket name. | 
**key** | **string** | Key name (supports dotted hierarchy). | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetKVEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MqKVEntry**](MqKVEntry.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqGetKVHistory

> MqGetKVHistory200Response MqGetKVHistory(ctx, bucket, key).Execute()

Get key history



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
	bucket := "bucket_example" // string | KV bucket name.
	key := "key_example" // string | Key name (supports dotted hierarchy).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqKeyValueAPI.MqGetKVHistory(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqGetKVHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetKVHistory`: MqGetKVHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `MqKeyValueAPI.MqGetKVHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | KV bucket name. | 
**key** | **string** | Key name (supports dotted hierarchy). | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetKVHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MqGetKVHistory200Response**](MqGetKVHistory200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListKVBuckets

> MqListKVBuckets200Response MqListKVBuckets(ctx).Limit(limit).Offset(offset).Execute()

List KV buckets



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
	resp, r, err := apiClient.MqKeyValueAPI.MqListKVBuckets(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqListKVBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListKVBuckets`: MqListKVBuckets200Response
	fmt.Fprintf(os.Stdout, "Response from `MqKeyValueAPI.MqListKVBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqListKVBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]

### Return type

[**MqListKVBuckets200Response**](MqListKVBuckets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListKVKeys

> MqListKVKeys200Response MqListKVKeys(ctx, bucket).Filter(filter).Limit(limit).Offset(offset).Execute()

List keys in bucket



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
	bucket := "bucket_example" // string | KV bucket name.
	filter := "filter_example" // string | Key filter pattern (supports wildcards). (optional)
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqKeyValueAPI.MqListKVKeys(context.Background(), bucket).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqListKVKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListKVKeys`: MqListKVKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `MqKeyValueAPI.MqListKVKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | KV bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqListKVKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Key filter pattern (supports wildcards). | 
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]

### Return type

[**MqListKVKeys200Response**](MqListKVKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqPutKVEntry

> MqPutKVEntry200Response MqPutKVEntry(ctx, bucket, key).MqPutKVEntryRequest(mqPutKVEntryRequest).XMQExpectedRevision(xMQExpectedRevision).Execute()

Put value



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
	bucket := "bucket_example" // string | KV bucket name.
	key := "key_example" // string | Key name (supports dotted hierarchy).
	mqPutKVEntryRequest := *openapiclient.NewMqPutKVEntryRequest("Value_example") // MqPutKVEntryRequest | 
	xMQExpectedRevision := int32(56) // int32 | Expected current revision for CAS. The put fails with 409 if the current revision does not match.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqKeyValueAPI.MqPutKVEntry(context.Background(), bucket, key).MqPutKVEntryRequest(mqPutKVEntryRequest).XMQExpectedRevision(xMQExpectedRevision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqPutKVEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqPutKVEntry`: MqPutKVEntry200Response
	fmt.Fprintf(os.Stdout, "Response from `MqKeyValueAPI.MqPutKVEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | KV bucket name. | 
**key** | **string** | Key name (supports dotted hierarchy). | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqPutKVEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mqPutKVEntryRequest** | [**MqPutKVEntryRequest**](MqPutKVEntryRequest.md) |  | 
 **xMQExpectedRevision** | **int32** | Expected current revision for CAS. The put fails with 409 if the current revision does not match.  | 

### Return type

[**MqPutKVEntry200Response**](MqPutKVEntry200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqWatchKVBucket

> MqKVEntry MqWatchKVBucket(ctx, bucket).Key(key).IncludeHistory(includeHistory).Execute()

Watch bucket changes via SSE



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
	bucket := "bucket_example" // string | KV bucket name.
	key := "key_example" // string | Key pattern filter (supports wildcards). (optional)
	includeHistory := true // bool | Include all historical revisions before live updates. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqKeyValueAPI.MqWatchKVBucket(context.Background(), bucket).Key(key).IncludeHistory(includeHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqKeyValueAPI.MqWatchKVBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqWatchKVBucket`: MqKVEntry
	fmt.Fprintf(os.Stdout, "Response from `MqKeyValueAPI.MqWatchKVBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | KV bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqWatchKVBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | Key pattern filter (supports wildcards). | 
 **includeHistory** | **bool** | Include all historical revisions before live updates. | [default to false]

### Return type

[**MqKVEntry**](MqKVEntry.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

