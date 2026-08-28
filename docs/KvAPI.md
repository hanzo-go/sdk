# \KvAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKvByBucket**](KvAPI.md#DeleteKvByBucket) | **Delete** /v1/kv/{bucket} | Removes one bucket of the caller&#39;s org — every key and every revision with it — and answers 204 with no body.
[**DeleteKvByBucketByKey**](KvAPI.md#DeleteKvByBucketByKey) | **Delete** /v1/kv/{bucket}/{key} | Delete removes one key — a delete marker in the key&#39;s history, so watchers see it and Get answers 404 — and answers 204 with no body.
[**GetKvByBucketByKey**](KvAPI.md#GetKvByBucketByKey) | **Get** /v1/kv/{bucket}/{key} | Get returns one key&#39;s current value and revision.
[**GetKvByBucketByKeyHistory**](KvAPI.md#GetKvByBucketByKeyHistory) | **Get** /v1/kv/{bucket}/{key}/history | History returns one key&#39;s retained revisions, oldest first — every put and every delete marker up to the bucket&#39;s History depth.
[**PostKvByBucket**](KvAPI.md#PostKvByBucket) | **Post** /v1/kv/{bucket} | Creates a KV bucket and returns it.
[**PutKvByBucketByKey**](KvAPI.md#PutKvByBucketByKey) | **Put** /v1/kv/{bucket}/{key} | Put sets one key to one value and returns the revision the write created.



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
- **Accept**: Not defined

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKvByBucketByKey

> KvEntry GetKvByBucketByKey(ctx, bucket, key).Execute()

Get returns one key's current value and revision.



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
	// response from `GetKvByBucketByKey`: KvEntry
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

[**KvEntry**](KvEntry.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKvByBucketByKeyHistory

> KvPage GetKvByBucketByKeyHistory(ctx, bucket, key).Execute()

History returns one key's retained revisions, oldest first — every put and every delete marker up to the bucket's History depth.



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
	// response from `GetKvByBucketByKeyHistory`: KvPage
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

[**KvPage**](KvPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKvByBucket

> BucketRecord PostKvByBucket(ctx, bucket).BucketWrite(bucketWrite).Execute()

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
	bucketWrite := *openapiclient.NewBucketWrite() // BucketWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvAPI.PostKvByBucket(context.Background(), bucket).BucketWrite(bucketWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.PostKvByBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKvByBucket`: BucketRecord
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


## PutKvByBucketByKey

> KvAck PutKvByBucketByKey(ctx, bucket, key).KvWrite(kvWrite).Execute()

Put sets one key to one value and returns the revision the write created.



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
	kvWrite := *openapiclient.NewKvWrite() // KvWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvAPI.PutKvByBucketByKey(context.Background(), bucket, key).KvWrite(kvWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.PutKvByBucketByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutKvByBucketByKey`: KvAck
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

