# \S3API

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteS3BucketsByBucket**](S3API.md#DeleteS3BucketsByBucket) | **Delete** /v1/s3/buckets/{bucket} | Removes an EMPTY bucket and answers 204.
[**GetS3Buckets**](S3API.md#GetS3Buckets) | **Get** /v1/s3/buckets | Lists the caller org&#39;s own buckets.
[**GetS3BucketsByBucketObjects**](S3API.md#GetS3BucketsByBucketObjects) | **Get** /v1/s3/buckets/{bucket}/objects | Lists one folder level of a bucket.
[**GetS3Health**](S3API.md#GetS3Health) | **Get** /v1/s3/health | Health reports whether this deployment can serve object storage.
[**PostS3Buckets**](S3API.md#PostS3Buckets) | **Post** /v1/s3/buckets | Makes a new bucket for the caller&#39;s org and answers 201 with it.
[**PostS3BucketsByBucketObjects**](S3API.md#PostS3BucketsByBucketObjects) | **Post** /v1/s3/buckets/{bucket}/objects | Mints a presigned PUT URL the caller uploads to DIRECTLY.



## DeleteS3BucketsByBucket

> DeleteS3BucketsByBucket(ctx, bucket).Execute()

Removes an EMPTY bucket and answers 204.



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
	bucket := "bucket_example" // string | Bucket is the bucket's friendly name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3API.DeleteS3BucketsByBucket(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.DeleteS3BucketsByBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket&#39;s friendly name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteS3BucketsByBucketRequest struct via the builder pattern


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


## GetS3Buckets

> BucketList GetS3Buckets(ctx).Execute()

Lists the caller org's own buckets.



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
	resp, r, err := apiClient.S3API.GetS3Buckets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetS3Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3Buckets`: BucketList
	fmt.Fprintf(os.Stdout, "Response from `S3API.GetS3Buckets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3BucketsRequest struct via the builder pattern


### Return type

[**BucketList**](BucketList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3BucketsByBucketObjects

> ObjectList GetS3BucketsByBucketObjects(ctx, bucket).Prefix(prefix).Recursive(recursive).Execute()

Lists one folder level of a bucket.



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
	bucket := "bucket_example" // string | Bucket is the bucket to list, from the path.
	prefix := "prefix_example" // string |  (optional)
	recursive := "recursive_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.GetS3BucketsByBucketObjects(context.Background(), bucket).Prefix(prefix).Recursive(recursive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetS3BucketsByBucketObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3BucketsByBucketObjects`: ObjectList
	fmt.Fprintf(os.Stdout, "Response from `S3API.GetS3BucketsByBucketObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3BucketsByBucketObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** |  | 
 **recursive** | **string** |  | 

### Return type

[**ObjectList**](ObjectList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3Health

> S3Health GetS3Health(ctx).Execute()

Health reports whether this deployment can serve object storage.



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
	resp, r, err := apiClient.S3API.GetS3Health(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetS3Health``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3Health`: S3Health
	fmt.Fprintf(os.Stdout, "Response from `S3API.GetS3Health`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3HealthRequest struct via the builder pattern


### Return type

[**S3Health**](S3Health.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3Buckets

> BucketItem PostS3Buckets(ctx).BucketIn(bucketIn).Execute()

Makes a new bucket for the caller's org and answers 201 with it.



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
	bucketIn := *openapiclient.NewBucketIn() // BucketIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.PostS3Buckets(context.Background()).BucketIn(bucketIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.PostS3Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3Buckets`: BucketItem
	fmt.Fprintf(os.Stdout, "Response from `S3API.PostS3Buckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostS3BucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketIn** | [**BucketIn**](BucketIn.md) |  | 

### Return type

[**BucketItem**](BucketItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3BucketsByBucketObjects

> PresignResponse PostS3BucketsByBucketObjects(ctx, bucket).UploadIn(uploadIn).Execute()

Mints a presigned PUT URL the caller uploads to DIRECTLY.



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
	bucket := "bucket_example" // string | Bucket is the bucket to upload into, from the path.
	uploadIn := *openapiclient.NewUploadIn() // UploadIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.PostS3BucketsByBucketObjects(context.Background(), bucket).UploadIn(uploadIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.PostS3BucketsByBucketObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3BucketsByBucketObjects`: PresignResponse
	fmt.Fprintf(os.Stdout, "Response from `S3API.PostS3BucketsByBucketObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket to upload into, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostS3BucketsByBucketObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadIn** | [**UploadIn**](UploadIn.md) |  | 

### Return type

[**PresignResponse**](PresignResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

