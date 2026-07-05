# \KmsAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsLogin1**](KmsAuthAPI.md#KmsLogin1) | **Post** /v1/kms/auth/login1 | Login step 1 - SRP init
[**KmsRenewAccessToken**](KmsAuthAPI.md#KmsRenewAccessToken) | **Post** /v1/kms/auth/token/renew | Renew access token



## KmsLogin1

> KmsLoginResponse KmsLogin1(ctx).KmsLoginRequest(kmsLoginRequest).Execute()

Login step 1 - SRP init

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
	kmsLoginRequest := *openapiclient.NewKmsLoginRequest("Email_example", "Password_example") // KmsLoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAuthAPI.KmsLogin1(context.Background()).KmsLoginRequest(kmsLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAuthAPI.KmsLogin1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsLogin1`: KmsLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `KmsAuthAPI.KmsLogin1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsLogin1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsLoginRequest** | [**KmsLoginRequest**](KmsLoginRequest.md) |  | 

### Return type

[**KmsLoginResponse**](KmsLoginResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsRenewAccessToken

> KmsTokenResponse KmsRenewAccessToken(ctx).Execute()

Renew access token

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
	resp, r, err := apiClient.KmsAuthAPI.KmsRenewAccessToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAuthAPI.KmsRenewAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsRenewAccessToken`: KmsTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `KmsAuthAPI.KmsRenewAccessToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsRenewAccessTokenRequest struct via the builder pattern


### Return type

[**KmsTokenResponse**](KmsTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

