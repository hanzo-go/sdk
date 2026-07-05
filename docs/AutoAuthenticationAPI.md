# \AutoAuthenticationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoSignIn**](AutoAuthenticationAPI.md#AutoSignIn) | **Post** /v1/auto/authentication/sign-in | Sign in with credentials
[**AutoSignUp**](AutoAuthenticationAPI.md#AutoSignUp) | **Post** /v1/auto/authentication/sign-up | Sign up a new user



## AutoSignIn

> map[string]interface{} AutoSignIn(ctx).AutoSignInRequest(autoSignInRequest).Execute()

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
	resp, r, err := apiClient.AutoAuthenticationAPI.AutoSignIn(context.Background()).AutoSignInRequest(autoSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAuthenticationAPI.AutoSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoSignIn`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAuthenticationAPI.AutoSignIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoSignInRequest struct via the builder pattern


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


## AutoSignUp

> map[string]interface{} AutoSignUp(ctx).AutoSignUpRequest(autoSignUpRequest).Execute()

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
	resp, r, err := apiClient.AutoAuthenticationAPI.AutoSignUp(context.Background()).AutoSignUpRequest(autoSignUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAuthenticationAPI.AutoSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoSignUp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAuthenticationAPI.AutoSignUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoSignUpRequest struct via the builder pattern


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

