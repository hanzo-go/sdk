# \ErrorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetErrors**](ErrorsAPI.md#GetErrors) | **Get** /v1/errors | Errors returns the caller org&#39;s most recently captured errors, newest first.



## GetErrors

> ErrorList GetErrors(ctx).Limit(limit).Execute()

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
	resp, r, err := apiClient.ErrorsAPI.GetErrors(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrorsAPI.GetErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetErrors`: ErrorList
	fmt.Fprintf(os.Stdout, "Response from `ErrorsAPI.GetErrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit is how many rows to return, newest first. Default 50, maximum 200; a value at or below zero, or one that is not a number, takes the default. | 

### Return type

[**ErrorList**](ErrorList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

