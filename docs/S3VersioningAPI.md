# \S3VersioningAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3GetBucketVersioning**](S3VersioningAPI.md#S3GetBucketVersioning) | **Get** /v1/s3/{bucket}?versioning | Get versioning status
[**S3PutBucketVersioning**](S3VersioningAPI.md#S3PutBucketVersioning) | **Put** /v1/s3/{bucket}?versioning | Set versioning status



## S3GetBucketVersioning

> S3VersioningConfig S3GetBucketVersioning(ctx, bucket).Execute()

Get versioning status

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3VersioningAPI.S3GetBucketVersioning(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3VersioningAPI.S3GetBucketVersioning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3GetBucketVersioning`: S3VersioningConfig
	fmt.Fprintf(os.Stdout, "Response from `S3VersioningAPI.S3GetBucketVersioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3GetBucketVersioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3VersioningConfig**](S3VersioningConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3PutBucketVersioning

> S3PutBucketVersioning(ctx, bucket).S3VersioningConfig(s3VersioningConfig).Execute()

Set versioning status

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
	s3VersioningConfig := *openapiclient.NewS3VersioningConfig() // S3VersioningConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3VersioningAPI.S3PutBucketVersioning(context.Background(), bucket).S3VersioningConfig(s3VersioningConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3VersioningAPI.S3PutBucketVersioning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3PutBucketVersioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3VersioningConfig** | [**S3VersioningConfig**](S3VersioningConfig.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

