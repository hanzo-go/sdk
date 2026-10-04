# \KvAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKvByBucket**](KvAPI.md#DeleteKvByBucket) | **Delete** /v1/kv/{bucket} | Removes one bucket of the caller&#39;s org — every key and every revision with it — and answers 204 with no body.
[**DeleteKvByBucketByKey**](KvAPI.md#DeleteKvByBucketByKey) | **Delete** /v1/kv/{bucket}/{key} | Delete removes one key — a delete marker in the key&#39;s history, so watchers see it and Get answers 404 — and answers 204 with no body.
[**GetKvByBucketByKey**](KvAPI.md#GetKvByBucketByKey) | **Get** /v1/kv/{bucket}/{key} | Returns one key&#39;s current value and revision.
[**GetKvByBucketByKeyHistory**](KvAPI.md#GetKvByBucketByKeyHistory) | **Get** /v1/kv/{bucket}/{key}/history | Returns one key&#39;s retained revisions, oldest first — every put and every delete marker up to the bucket&#39;s History depth.
[**PostKvByBucket**](KvAPI.md#PostKvByBucket) | **Post** /v1/kv/{bucket} | Creates a KV bucket and returns it.
[**PutKvByBucketByKey**](KvAPI.md#PutKvByBucketByKey) | **Put** /v1/kv/{bucket}/{key} | Sets one key to one value and returns the revision the write created.



## DeleteKvByBucket

> DeleteKvByBucket(ctx, bucket).Execute()

Removes one bucket of the caller's org — every key and every revision with it — and answers 204 with no body.



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
	bucket := "bucket_example" // string | Bucket is the bucket's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KvAPI.DeleteKvByBucket(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.DeleteKvByBucket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKvByBucketRequest struct via the builder pattern


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


## DeleteKvByBucketByKey

> DeleteKvByBucketByKey(ctx, bucket, key).Execute()

Delete removes one key — a delete marker in the key's history, so watchers see it and Get answers 404 — and answers 204 with no body.



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
	bucket := "bucket_example" // string | Bucket is the bucket, from the path.
	key := "key_example" // string | Key is the key, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KvAPI.DeleteKvByBucketByKey(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.DeleteKvByBucketByKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKvByBucketByKeyRequest struct via the builder pattern


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


## GetKvByBucketByKey

> KvKvEntry GetKvByBucketByKey(ctx, bucket, key).Execute()

Returns one key's current value and revision.



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
	bucket := "bucket_example" // string | Bucket is the bucket, from the path.
	key := "key_example" // string | Key is the key, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvAPI.GetKvByBucketByKey(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.GetKvByBucketByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKvByBucketByKey`: KvKvEntry
	fmt.Fprintf(os.Stdout, "Response from `KvAPI.GetKvByBucketByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKvByBucketByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KvKvEntry**](KvKvEntry.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKvByBucketByKeyHistory

> KvKvPage GetKvByBucketByKeyHistory(ctx, bucket, key).Execute()

Returns one key's retained revisions, oldest first — every put and every delete marker up to the bucket's History depth.



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
	bucket := "bucket_example" // string | Bucket is the bucket, from the path.
	key := "key_example" // string | Key is the key, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvAPI.GetKvByBucketByKeyHistory(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.GetKvByBucketByKeyHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKvByBucketByKeyHistory`: KvKvPage
	fmt.Fprintf(os.Stdout, "Response from `KvAPI.GetKvByBucketByKeyHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKvByBucketByKeyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KvKvPage**](KvKvPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKvByBucket

> KvBucketRecord PostKvByBucket(ctx, bucket).KvBucketWrite(kvBucketWrite).Execute()

Creates a KV bucket and returns it.



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
	bucket := "bucket_example" // string | Bucket is the bucket's name within the org, from the path: 1–64 of [A-Za-z0-9_], no dash.
	kvBucketWrite := *openapiclient.NewKvBucketWrite() // KvBucketWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvAPI.PostKvByBucket(context.Background(), bucket).KvBucketWrite(kvBucketWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.PostKvByBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKvByBucket`: KvBucketRecord
	fmt.Fprintf(os.Stdout, "Response from `KvAPI.PostKvByBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket&#39;s name within the org, from the path: 1–64 of [A-Za-z0-9_], no dash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostKvByBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvBucketWrite** | [**KvBucketWrite**](KvBucketWrite.md) |  | 

### Return type

[**KvBucketRecord**](KvBucketRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutKvByBucketByKey

> KvKvAck PutKvByBucketByKey(ctx, bucket, key).KvKvWrite(kvKvWrite).Execute()

Sets one key to one value and returns the revision the write created.



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
	bucket := "bucket_example" // string | Bucket is the bucket, from the path.
	key := "key_example" // string | Key is the key, from the path.
	kvKvWrite := *openapiclient.NewKvKvWrite() // KvKvWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvAPI.PutKvByBucketByKey(context.Background(), bucket, key).KvKvWrite(kvKvWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.PutKvByBucketByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutKvByBucketByKey`: KvKvAck
	fmt.Fprintf(os.Stdout, "Response from `KvAPI.PutKvByBucketByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**key** | **string** | Key is the key, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutKvByBucketByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvKvWrite** | [**KvKvWrite**](KvKvWrite.md) |  | 

### Return type

[**KvKvAck**](KvKvAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

