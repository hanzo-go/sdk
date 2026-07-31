# \EnforcementAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthzAuthzCheck**](EnforcementAPI.md#AuthzAuthzCheck) | **Post** /v1/authz/check | Check a permission



## AuthzAuthzCheck

> AuthzCheckResponse AuthzAuthzCheck(ctx).AuthzEnforceRequest(authzEnforceRequest).Execute()

Check a permission



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
	authzEnforceRequest := *openapiclient.NewAuthzEnforceRequest("Sub_example", "Obj_example", "Act_example") // AuthzEnforceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnforcementAPI.AuthzAuthzCheck(context.Background()).AuthzEnforceRequest(authzEnforceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnforcementAPI.AuthzAuthzCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthzAuthzCheck`: AuthzCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `EnforcementAPI.AuthzAuthzCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthzAuthzCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authzEnforceRequest** | [**AuthzEnforceRequest**](AuthzEnforceRequest.md) |  | 

### Return type

[**AuthzCheckResponse**](AuthzCheckResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

