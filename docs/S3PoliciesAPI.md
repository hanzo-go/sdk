# \S3PoliciesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3GetBucketPolicy**](S3PoliciesAPI.md#S3GetBucketPolicy) | **Get** /v1/s3/{bucket}?policy | Get bucket policy
[**S3PutBucketPolicy**](S3PoliciesAPI.md#S3PutBucketPolicy) | **Put** /v1/s3/{bucket}?policy | Set bucket policy



## S3GetBucketPolicy

> S3BucketPolicy S3GetBucketPolicy(ctx, bucket).Execute()

Get bucket policy

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
	resp, r, err := apiClient.S3PoliciesAPI.S3GetBucketPolicy(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3PoliciesAPI.S3GetBucketPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3GetBucketPolicy`: S3BucketPolicy
	fmt.Fprintf(os.Stdout, "Response from `S3PoliciesAPI.S3GetBucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3GetBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3BucketPolicy**](S3BucketPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3PutBucketPolicy

> S3PutBucketPolicy(ctx, bucket).S3BucketPolicy(s3BucketPolicy).Execute()

Set bucket policy

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
	s3BucketPolicy := *openapiclient.NewS3BucketPolicy() // S3BucketPolicy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3PoliciesAPI.S3PutBucketPolicy(context.Background(), bucket).S3BucketPolicy(s3BucketPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3PoliciesAPI.S3PutBucketPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiS3PutBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3BucketPolicy** | [**S3BucketPolicy**](S3BucketPolicy.md) |  | 

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

