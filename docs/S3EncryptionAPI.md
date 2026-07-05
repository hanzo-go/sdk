# \S3EncryptionAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3GetBucketEncryption**](S3EncryptionAPI.md#S3GetBucketEncryption) | **Get** /v1/s3/{bucket}?encryption | Get encryption configuration
[**S3PutBucketEncryption**](S3EncryptionAPI.md#S3PutBucketEncryption) | **Put** /v1/s3/{bucket}?encryption | Set encryption configuration



## S3GetBucketEncryption

> S3EncryptionConfig S3GetBucketEncryption(ctx, bucket).Execute()

Get encryption configuration

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
	resp, r, err := apiClient.S3EncryptionAPI.S3GetBucketEncryption(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3EncryptionAPI.S3GetBucketEncryption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3GetBucketEncryption`: S3EncryptionConfig
	fmt.Fprintf(os.Stdout, "Response from `S3EncryptionAPI.S3GetBucketEncryption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3GetBucketEncryptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3EncryptionConfig**](S3EncryptionConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3PutBucketEncryption

> S3PutBucketEncryption(ctx, bucket).S3EncryptionConfig(s3EncryptionConfig).Execute()

Set encryption configuration

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
	s3EncryptionConfig := *openapiclient.NewS3EncryptionConfig() // S3EncryptionConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3EncryptionAPI.S3PutBucketEncryption(context.Background(), bucket).S3EncryptionConfig(s3EncryptionConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3EncryptionAPI.S3PutBucketEncryption``: %v\n", err)
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

Other parameters are passed through a pointer to a apiS3PutBucketEncryptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3EncryptionConfig** | [**S3EncryptionConfig**](S3EncryptionConfig.md) |  | 

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

