# \CommerceAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceAuthenticate**](CommerceAuthAPI.md#CommerceAuthenticate) | **Post** /v1/commerce/auth | Authenticate user (OAuth2)



## CommerceAuthenticate

> CommerceOAuthResponse CommerceAuthenticate(ctx).CommerceOAuthRequest(commerceOAuthRequest).Execute()

Authenticate user (OAuth2)



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
	commerceOAuthRequest := *openapiclient.NewCommerceOAuthRequest() // CommerceOAuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAuthAPI.CommerceAuthenticate(context.Background()).CommerceOAuthRequest(commerceOAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAuthAPI.CommerceAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceAuthenticate`: CommerceOAuthResponse
	fmt.Fprintf(os.Stdout, "Response from `CommerceAuthAPI.CommerceAuthenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceOAuthRequest** | [**CommerceOAuthRequest**](CommerceOAuthRequest.md) |  | 

### Return type

[**CommerceOAuthResponse**](CommerceOAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

