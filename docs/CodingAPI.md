# \CodingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCoding**](CodingAPI.md#PostCoding) | **Post** /v1/coding | Start one autonomous coding run against a repo in the caller&#39;s org



## PostCoding

> CodingStarted PostCoding(ctx).CodingStartIn(codingStartIn).Execute()

Start one autonomous coding run against a repo in the caller's org



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
	codingStartIn := *openapiclient.NewCodingStartIn() // CodingStartIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodingAPI.PostCoding(context.Background()).CodingStartIn(codingStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodingAPI.PostCoding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCoding`: CodingStarted
	fmt.Fprintf(os.Stdout, "Response from `CodingAPI.PostCoding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codingStartIn** | [**CodingStartIn**](CodingStartIn.md) |  | 

### Return type

[**CodingStarted**](CodingStarted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

