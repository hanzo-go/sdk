# \ErrorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Errors**](ErrorsAPI.md#CloudGetV1Errors) | **Get** /v1/errors | Errors returns the caller org&#39;s most recently captured errors, newest first.



## CloudGetV1Errors

> CloudErrorList CloudGetV1Errors(ctx).Limit(limit).Execute()

Errors returns the caller org's most recently captured errors, newest first.



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
	limit := int32(100) // int32 | Limit is how many rows to return, newest first. Default 50, maximum 200; a value at or below zero, or one that is not a number, takes the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ErrorsAPI.CloudGetV1Errors(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrorsAPI.CloudGetV1Errors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Errors`: CloudErrorList
	fmt.Fprintf(os.Stdout, "Response from `ErrorsAPI.CloudGetV1Errors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit is how many rows to return, newest first. Default 50, maximum 200; a value at or below zero, or one that is not a number, takes the default. | 

### Return type

[**CloudErrorList**](CloudErrorList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

