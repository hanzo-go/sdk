# \GuardSanitizeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GuardSanitizeInput**](GuardSanitizeAPI.md#GuardSanitizeInput) | **Post** /v1/guard/sanitize/input | Sanitize user input
[**GuardSanitizeOutput**](GuardSanitizeAPI.md#GuardSanitizeOutput) | **Post** /v1/guard/sanitize/output | Sanitize LLM output



## GuardSanitizeInput

> GuardSanitizeResult GuardSanitizeInput(ctx).GuardSanitizeRequest(guardSanitizeRequest).Execute()

Sanitize user input



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
	guardSanitizeRequest := *openapiclient.NewGuardSanitizeRequest("My SSN is 123-45-6789 and email is ceo@company.com") // GuardSanitizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuardSanitizeAPI.GuardSanitizeInput(context.Background()).GuardSanitizeRequest(guardSanitizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardSanitizeAPI.GuardSanitizeInput``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GuardSanitizeInput`: GuardSanitizeResult
	fmt.Fprintf(os.Stdout, "Response from `GuardSanitizeAPI.GuardSanitizeInput`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardSanitizeInputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guardSanitizeRequest** | [**GuardSanitizeRequest**](GuardSanitizeRequest.md) |  | 

### Return type

[**GuardSanitizeResult**](GuardSanitizeResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GuardSanitizeOutput

> GuardSanitizeResult GuardSanitizeOutput(ctx).GuardSanitizeRequest(guardSanitizeRequest).Execute()

Sanitize LLM output



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
	guardSanitizeRequest := *openapiclient.NewGuardSanitizeRequest("My SSN is 123-45-6789 and email is ceo@company.com") // GuardSanitizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuardSanitizeAPI.GuardSanitizeOutput(context.Background()).GuardSanitizeRequest(guardSanitizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardSanitizeAPI.GuardSanitizeOutput``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GuardSanitizeOutput`: GuardSanitizeResult
	fmt.Fprintf(os.Stdout, "Response from `GuardSanitizeAPI.GuardSanitizeOutput`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardSanitizeOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guardSanitizeRequest** | [**GuardSanitizeRequest**](GuardSanitizeRequest.md) |  | 

### Return type

[**GuardSanitizeResult**](GuardSanitizeResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

