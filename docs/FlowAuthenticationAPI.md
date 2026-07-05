# \FlowAuthenticationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowFederatedClaim**](FlowAuthenticationAPI.md#FlowFederatedClaim) | **Post** /v1/flow/authn/federated/claim | Complete federated authentication (EE)
[**FlowFederatedLogin**](FlowAuthenticationAPI.md#FlowFederatedLogin) | **Get** /v1/flow/authn/federated/login | Initiate federated authentication (EE)
[**FlowSignIn**](FlowAuthenticationAPI.md#FlowSignIn) | **Post** /v1/flow/authentication/sign-in | Sign in with credentials
[**FlowSignUp**](FlowAuthenticationAPI.md#FlowSignUp) | **Post** /v1/flow/authentication/sign-up | Sign up a new user



## FlowFederatedClaim

> map[string]interface{} FlowFederatedClaim(ctx).Body(body).Execute()

Complete federated authentication (EE)

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAuthenticationAPI.FlowFederatedClaim(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAuthenticationAPI.FlowFederatedClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowFederatedClaim`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAuthenticationAPI.FlowFederatedClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowFederatedClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowFederatedLogin

> FlowFederatedLogin(ctx).Execute()

Initiate federated authentication (EE)

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
	r, err := apiClient.FlowAuthenticationAPI.FlowFederatedLogin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAuthenticationAPI.FlowFederatedLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowFederatedLoginRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowSignIn

> map[string]interface{} FlowSignIn(ctx).AutoSignInRequest(autoSignInRequest).Execute()

Sign in with credentials

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
	autoSignInRequest := *openapiclient.NewAutoSignInRequest("Email_example", "Password_example") // AutoSignInRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAuthenticationAPI.FlowSignIn(context.Background()).AutoSignInRequest(autoSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAuthenticationAPI.FlowSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowSignIn`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAuthenticationAPI.FlowSignIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoSignInRequest** | [**AutoSignInRequest**](AutoSignInRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowSignUp

> map[string]interface{} FlowSignUp(ctx).AutoSignUpRequest(autoSignUpRequest).Execute()

Sign up a new user

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
	autoSignUpRequest := *openapiclient.NewAutoSignUpRequest("Email_example", "Password_example", "FirstName_example", "LastName_example") // AutoSignUpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAuthenticationAPI.FlowSignUp(context.Background()).AutoSignUpRequest(autoSignUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAuthenticationAPI.FlowSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowSignUp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAuthenticationAPI.FlowSignUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoSignUpRequest** | [**AutoSignUpRequest**](AutoSignUpRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

