# \LifecycleAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3GetBucketLifecycle**](LifecycleAPI.md#S3GetBucketLifecycle) | **Get** /v1/s3/{bucket}?lifecycle | Get lifecycle rules
[**S3PutBucketLifecycle**](LifecycleAPI.md#S3PutBucketLifecycle) | **Put** /v1/s3/{bucket}?lifecycle | Set lifecycle rules



## S3GetBucketLifecycle

> S3GetBucketLifecycle200Response S3GetBucketLifecycle(ctx, bucket).Execute()

Get lifecycle rules

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
	resp, r, err := apiClient.LifecycleAPI.S3GetBucketLifecycle(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.S3GetBucketLifecycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3GetBucketLifecycle`: S3GetBucketLifecycle200Response
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.S3GetBucketLifecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3GetBucketLifecycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3GetBucketLifecycle200Response**](S3GetBucketLifecycle200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3PutBucketLifecycle

> S3PutBucketLifecycle(ctx, bucket).S3GetBucketLifecycle200Response(s3GetBucketLifecycle200Response).Execute()

Set lifecycle rules

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
	s3GetBucketLifecycle200Response := *openapiclient.NewS3GetBucketLifecycle200Response() // S3GetBucketLifecycle200Response | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LifecycleAPI.S3PutBucketLifecycle(context.Background(), bucket).S3GetBucketLifecycle200Response(s3GetBucketLifecycle200Response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.S3PutBucketLifecycle``: %v\n", err)
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

Other parameters are passed through a pointer to a apiS3PutBucketLifecycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3GetBucketLifecycle200Response** | [**S3GetBucketLifecycle200Response**](S3GetBucketLifecycle200Response.md) |  | 

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

