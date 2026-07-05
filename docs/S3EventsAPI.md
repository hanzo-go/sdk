# \S3EventsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3GetBucketNotification**](S3EventsAPI.md#S3GetBucketNotification) | **Get** /v1/s3/{bucket}?notification | Get event notification config
[**S3PutBucketNotification**](S3EventsAPI.md#S3PutBucketNotification) | **Put** /v1/s3/{bucket}?notification | Set event notification config



## S3GetBucketNotification

> S3GetBucketNotification200Response S3GetBucketNotification(ctx, bucket).Execute()

Get event notification config

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
	resp, r, err := apiClient.S3EventsAPI.S3GetBucketNotification(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3EventsAPI.S3GetBucketNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3GetBucketNotification`: S3GetBucketNotification200Response
	fmt.Fprintf(os.Stdout, "Response from `S3EventsAPI.S3GetBucketNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3GetBucketNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3GetBucketNotification200Response**](S3GetBucketNotification200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3PutBucketNotification

> S3PutBucketNotification(ctx, bucket).S3GetBucketNotification200Response(s3GetBucketNotification200Response).Execute()

Set event notification config

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
	s3GetBucketNotification200Response := *openapiclient.NewS3GetBucketNotification200Response() // S3GetBucketNotification200Response | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3EventsAPI.S3PutBucketNotification(context.Background(), bucket).S3GetBucketNotification200Response(s3GetBucketNotification200Response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3EventsAPI.S3PutBucketNotification``: %v\n", err)
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

Other parameters are passed through a pointer to a apiS3PutBucketNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3GetBucketNotification200Response** | [**S3GetBucketNotification200Response**](S3GetBucketNotification200Response.md) |  | 

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

