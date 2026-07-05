# \PaasAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasGetCurrentUser**](PaasAuthAPI.md#PaasGetCurrentUser) | **Get** /v1/paas/auth/me | Get current user
[**PaasLogin**](PaasAuthAPI.md#PaasLogin) | **Post** /v1/paas/auth/login | Login via IAM



## PaasGetCurrentUser

> map[string]interface{} PaasGetCurrentUser(ctx).Execute()

Get current user

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
	resp, r, err := apiClient.PaasAuthAPI.PaasGetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasAuthAPI.PaasGetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetCurrentUser`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasAuthAPI.PaasGetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetCurrentUserRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasLogin

> PaasLogin200Response PaasLogin(ctx).PaasLoginRequest(paasLoginRequest).Execute()

Login via IAM

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
	paasLoginRequest := *openapiclient.NewPaasLoginRequest() // PaasLoginRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasAuthAPI.PaasLogin(context.Background()).PaasLoginRequest(paasLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasAuthAPI.PaasLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasLogin`: PaasLogin200Response
	fmt.Fprintf(os.Stdout, "Response from `PaasAuthAPI.PaasLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaasLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paasLoginRequest** | [**PaasLoginRequest**](PaasLoginRequest.md) |  | 

### Return type

[**PaasLogin200Response**](PaasLogin200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

