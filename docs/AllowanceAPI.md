# \AllowanceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllowance**](AllowanceAPI.md#GetAllowance) | **Get** /v1/allowance | Answers what the CALLER has left of their plan&#39;s free-call allowance this period, and the instant the count starts again.



## GetAllowance

> Allowance GetAllowance(ctx).Execute()

Answers what the CALLER has left of their plan's free-call allowance this period, and the instant the count starts again.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowanceAPI.GetAllowance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowanceAPI.GetAllowance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllowance`: Allowance
	fmt.Fprintf(os.Stdout, "Response from `AllowanceAPI.GetAllowance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowanceRequest struct via the builder pattern


### Return type

[**Allowance**](Allowance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

