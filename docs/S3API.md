# \S3API

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteS3BucketsByBucket**](S3API.md#DeleteS3BucketsByBucket) | **Delete** /v1/s3/buckets/{bucket} | Removes an EMPTY bucket and answers 204.
[**DeleteS3BucketsByBucketUploadsByUpload**](S3API.md#DeleteS3BucketsByBucketUploadsByUpload) | **Delete** /v1/s3/buckets/{bucket}/uploads/{upload} | Aborts a multipart upload and deletes the parts it stored.
[**GetS3Buckets**](S3API.md#GetS3Buckets) | **Get** /v1/s3/buckets | Lists the caller org&#39;s own buckets.
[**GetS3BucketsByBucketObjects**](S3API.md#GetS3BucketsByBucketObjects) | **Get** /v1/s3/buckets/{bucket}/objects | Lists one folder level of a bucket.
[**GetS3BucketsByBucketUploadsByUpload**](S3API.md#GetS3BucketsByBucketUploadsByUpload) | **Get** /v1/s3/buckets/{bucket}/uploads/{upload} | Answers the parts a multipart upload already holds, in order: what a client resuming after a dropped connection reads to send only the rest.
[**GetS3Health**](S3API.md#GetS3Health) | **Get** /v1/s3/health | Reports whether this deployment can serve object storage.
[**PostS3Buckets**](S3API.md#PostS3Buckets) | **Post** /v1/s3/buckets | Makes a new bucket for the caller&#39;s org and answers 201 with it.
[**PostS3BucketsByBucketObjects**](S3API.md#PostS3BucketsByBucketObjects) | **Post** /v1/s3/buckets/{bucket}/objects | Mints a presigned PUT URL the caller uploads to DIRECTLY.
[**PostS3BucketsByBucketUploads**](S3API.md#PostS3BucketsByBucketUploads) | **Post** /v1/s3/buckets/{bucket}/uploads | Begins a multipart upload of a large file into one of the caller&#39;s org buckets.
[**PostS3BucketsByBucketUploadsByUploadComplete**](S3API.md#PostS3BucketsByBucketUploadsByUploadComplete) | **Post** /v1/s3/buckets/{bucket}/uploads/{upload}/complete | Assembles a multipart upload into its object from every part the store holds, in order.
[**PostS3BucketsByBucketUploadsByUploadParts**](S3API.md#PostS3BucketsByBucketUploadsByUploadParts) | **Post** /v1/s3/buckets/{bucket}/uploads/{upload}/parts | Answers presigned PUT URLs for some parts of a multipart upload, at most 64 a call, each valid for minutes: mint them as the upload reaches them.



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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteS3BucketsByBucketUploadsByUpload

> S3UploadGone DeleteS3BucketsByBucketUploadsByUpload(ctx, bucket, upload).Key(key).Execute()

Aborts a multipart upload and deletes the parts it stored.



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
	upload := "upload_example" // string | Upload is the upload's id, from the path.
	key := "key_example" // string | Key is the object key the upload was started on. In the query for a read or an abandon, in the body to complete. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.DeleteS3BucketsByBucketUploadsByUpload(context.Background(), bucket, upload).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.DeleteS3BucketsByBucketUploadsByUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteS3BucketsByBucketUploadsByUpload`: S3UploadGone
	fmt.Fprintf(os.Stdout, "Response from `S3API.DeleteS3BucketsByBucketUploadsByUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**upload** | **string** | Upload is the upload&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteS3BucketsByBucketUploadsByUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **string** | Key is the object key the upload was started on. In the query for a read or an abandon, in the body to complete. | 

### Return type

[**S3UploadGone**](S3UploadGone.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3Buckets

> S3BucketList GetS3Buckets(ctx).Execute()

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
	// response from `GetS3Buckets`: S3BucketList
	fmt.Fprintf(os.Stdout, "Response from `S3API.GetS3Buckets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3BucketsRequest struct via the builder pattern


### Return type

[**S3BucketList**](S3BucketList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3BucketsByBucketObjects

> S3ObjectList GetS3BucketsByBucketObjects(ctx, bucket).Prefix(prefix).Recursive(recursive).Execute()

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
	// response from `GetS3BucketsByBucketObjects`: S3ObjectList
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

[**S3ObjectList**](S3ObjectList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3BucketsByBucketUploadsByUpload

> S3StoredParts GetS3BucketsByBucketUploadsByUpload(ctx, bucket, upload).Key(key).Execute()

Answers the parts a multipart upload already holds, in order: what a client resuming after a dropped connection reads to send only the rest.



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
	upload := "upload_example" // string | Upload is the upload's id, from the path.
	key := "key_example" // string | Key is the object key the upload was started on. In the query for a read or an abandon, in the body to complete. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.GetS3BucketsByBucketUploadsByUpload(context.Background(), bucket, upload).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetS3BucketsByBucketUploadsByUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3BucketsByBucketUploadsByUpload`: S3StoredParts
	fmt.Fprintf(os.Stdout, "Response from `S3API.GetS3BucketsByBucketUploadsByUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**upload** | **string** | Upload is the upload&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3BucketsByBucketUploadsByUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **string** | Key is the object key the upload was started on. In the query for a read or an abandon, in the body to complete. | 

### Return type

[**S3StoredParts**](S3StoredParts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3Health

> S3S3Health GetS3Health(ctx).Execute()

Reports whether this deployment can serve object storage.



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
	// response from `GetS3Health`: S3S3Health
	fmt.Fprintf(os.Stdout, "Response from `S3API.GetS3Health`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3HealthRequest struct via the builder pattern


### Return type

[**S3S3Health**](S3S3Health.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3Buckets

> S3BucketItem PostS3Buckets(ctx).S3BucketIn(s3BucketIn).Execute()

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
	s3BucketIn := *openapiclient.NewS3BucketIn() // S3BucketIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.PostS3Buckets(context.Background()).S3BucketIn(s3BucketIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.PostS3Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3Buckets`: S3BucketItem
	fmt.Fprintf(os.Stdout, "Response from `S3API.PostS3Buckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostS3BucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s3BucketIn** | [**S3BucketIn**](S3BucketIn.md) |  | 

### Return type

[**S3BucketItem**](S3BucketItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3BucketsByBucketObjects

> S3PresignResponse PostS3BucketsByBucketObjects(ctx, bucket).S3UploadIn(s3UploadIn).Execute()

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
	s3UploadIn := *openapiclient.NewS3UploadIn() // S3UploadIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.PostS3BucketsByBucketObjects(context.Background(), bucket).S3UploadIn(s3UploadIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.PostS3BucketsByBucketObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3BucketsByBucketObjects`: S3PresignResponse
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

 **s3UploadIn** | [**S3UploadIn**](S3UploadIn.md) |  | 

### Return type

[**S3PresignResponse**](S3PresignResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3BucketsByBucketUploads

> S3UploadStarted PostS3BucketsByBucketUploads(ctx, bucket).S3UploadStart(s3UploadStart).Execute()

Begins a multipart upload of a large file into one of the caller's org buckets.



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
	s3UploadStart := *openapiclient.NewS3UploadStart() // S3UploadStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.PostS3BucketsByBucketUploads(context.Background(), bucket).S3UploadStart(s3UploadStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.PostS3BucketsByBucketUploads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3BucketsByBucketUploads`: S3UploadStarted
	fmt.Fprintf(os.Stdout, "Response from `S3API.PostS3BucketsByBucketUploads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket to upload into, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostS3BucketsByBucketUploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3UploadStart** | [**S3UploadStart**](S3UploadStart.md) |  | 

### Return type

[**S3UploadStarted**](S3UploadStarted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3BucketsByBucketUploadsByUploadComplete

> S3UploadDone PostS3BucketsByBucketUploadsByUploadComplete(ctx, bucket, upload).S3UploadRef(s3UploadRef).Execute()

Assembles a multipart upload into its object from every part the store holds, in order.



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
	upload := "upload_example" // string | Upload is the upload's id, from the path.
	s3UploadRef := *openapiclient.NewS3UploadRef() // S3UploadRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.PostS3BucketsByBucketUploadsByUploadComplete(context.Background(), bucket, upload).S3UploadRef(s3UploadRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.PostS3BucketsByBucketUploadsByUploadComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3BucketsByBucketUploadsByUploadComplete`: S3UploadDone
	fmt.Fprintf(os.Stdout, "Response from `S3API.PostS3BucketsByBucketUploadsByUploadComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**upload** | **string** | Upload is the upload&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostS3BucketsByBucketUploadsByUploadCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **s3UploadRef** | [**S3UploadRef**](S3UploadRef.md) |  | 

### Return type

[**S3UploadDone**](S3UploadDone.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3BucketsByBucketUploadsByUploadParts

> S3PartURLs PostS3BucketsByBucketUploadsByUploadParts(ctx, bucket, upload).S3UploadParts(s3UploadParts).Execute()

Answers presigned PUT URLs for some parts of a multipart upload, at most 64 a call, each valid for minutes: mint them as the upload reaches them.



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
	upload := "upload_example" // string | Upload is the upload's id, from the path.
	s3UploadParts := *openapiclient.NewS3UploadParts() // S3UploadParts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3API.PostS3BucketsByBucketUploadsByUploadParts(context.Background(), bucket, upload).S3UploadParts(s3UploadParts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3API.PostS3BucketsByBucketUploadsByUploadParts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3BucketsByBucketUploadsByUploadParts`: S3PartURLs
	fmt.Fprintf(os.Stdout, "Response from `S3API.PostS3BucketsByBucketUploadsByUploadParts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the bucket, from the path. | 
**upload** | **string** | Upload is the upload&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostS3BucketsByBucketUploadsByUploadPartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **s3UploadParts** | [**S3UploadParts**](S3UploadParts.md) |  | 

### Return type

[**S3PartURLs**](S3PartURLs.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

