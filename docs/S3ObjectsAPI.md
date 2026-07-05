# \S3ObjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3DeleteObject**](S3ObjectsAPI.md#S3DeleteObject) | **Delete** /v1/s3/{bucket}/{key} | Delete an object
[**S3GetObject**](S3ObjectsAPI.md#S3GetObject) | **Get** /v1/s3/{bucket}/{key} | Download an object
[**S3HeadObject**](S3ObjectsAPI.md#S3HeadObject) | **Head** /v1/s3/{bucket}/{key} | Get object metadata
[**S3ListObjectsV2**](S3ObjectsAPI.md#S3ListObjectsV2) | **Get** /v1/s3/{bucket} | List objects in bucket
[**S3PutObject**](S3ObjectsAPI.md#S3PutObject) | **Put** /v1/s3/{bucket}/{key} | Upload an object



## S3DeleteObject

> S3DeleteObject(ctx, bucket, key).VersionId(versionId).Execute()

Delete an object

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
	bucket := "bucket_example" // string | 
	key := "key_example" // string | 
	versionId := "versionId_example" // string | Specific version to delete (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3ObjectsAPI.S3DeleteObject(context.Background(), bucket, key).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ObjectsAPI.S3DeleteObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3DeleteObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionId** | **string** | Specific version to delete | 

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


## S3GetObject

> *os.File S3GetObject(ctx, bucket, key).VersionId(versionId).Range_(range_).Execute()

Download an object

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
	bucket := "bucket_example" // string | 
	key := "key_example" // string | 
	versionId := "versionId_example" // string | Specific version to retrieve (optional)
	range_ := "range__example" // string | Byte range (e.g. bytes=0-1023) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3ObjectsAPI.S3GetObject(context.Background(), bucket, key).VersionId(versionId).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ObjectsAPI.S3GetObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3GetObject`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `S3ObjectsAPI.S3GetObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3GetObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionId** | **string** | Specific version to retrieve | 
 **range_** | **string** | Byte range (e.g. bytes&#x3D;0-1023) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3HeadObject

> S3HeadObject(ctx, bucket, key).Execute()

Get object metadata

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
	bucket := "bucket_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3ObjectsAPI.S3HeadObject(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ObjectsAPI.S3HeadObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3HeadObjectRequest struct via the builder pattern


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


## S3ListObjectsV2

> S3ListObjectsV2200Response S3ListObjectsV2(ctx, bucket).Prefix(prefix).Delimiter(delimiter).MaxKeys(maxKeys).ContinuationToken(continuationToken).Execute()

List objects in bucket

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
	bucket := "bucket_example" // string | 
	prefix := "prefix_example" // string | Filter by key prefix (optional)
	delimiter := "delimiter_example" // string | Grouping delimiter (e.g. /) (optional)
	maxKeys := int32(56) // int32 |  (optional) (default to 1000)
	continuationToken := "continuationToken_example" // string | Pagination token from previous response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3ObjectsAPI.S3ListObjectsV2(context.Background(), bucket).Prefix(prefix).Delimiter(delimiter).MaxKeys(maxKeys).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ObjectsAPI.S3ListObjectsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3ListObjectsV2`: S3ListObjectsV2200Response
	fmt.Fprintf(os.Stdout, "Response from `S3ObjectsAPI.S3ListObjectsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3ListObjectsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Filter by key prefix | 
 **delimiter** | **string** | Grouping delimiter (e.g. /) | 
 **maxKeys** | **int32** |  | [default to 1000]
 **continuationToken** | **string** | Pagination token from previous response | 

### Return type

[**S3ListObjectsV2200Response**](S3ListObjectsV2200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3PutObject

> S3PutObject(ctx, bucket, key).Body(body).ContentType(contentType).XAmzServerSideEncryption(xAmzServerSideEncryption).Execute()

Upload an object

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
	bucket := "bucket_example" // string | 
	key := "key_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File | 
	contentType := "contentType_example" // string |  (optional) (default to "application/octet-stream")
	xAmzServerSideEncryption := "xAmzServerSideEncryption_example" // string | Server-side encryption algorithm (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3ObjectsAPI.S3PutObject(context.Background(), bucket, key).Body(body).ContentType(contentType).XAmzServerSideEncryption(xAmzServerSideEncryption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ObjectsAPI.S3PutObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3PutObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 
 **contentType** | **string** |  | [default to &quot;application/octet-stream&quot;]
 **xAmzServerSideEncryption** | **string** | Server-side encryption algorithm | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

