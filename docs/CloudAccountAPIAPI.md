# \CloudAccountAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerGetAccount**](CloudAccountAPIAPI.md#CloudApiControllerGetAccount) | **Get** /v1/cloud/get-account | Api Controller Get Account
[**CloudApiControllerSignin**](CloudAccountAPIAPI.md#CloudApiControllerSignin) | **Post** /v1/cloud/signin | Api Controller Signin
[**CloudApiControllerSignout**](CloudAccountAPIAPI.md#CloudApiControllerSignout) | **Post** /v1/cloud/signout | Api Controller Signout



## CloudApiControllerGetAccount

> map[string]interface{} CloudApiControllerGetAccount(ctx).Execute()

Api Controller Get Account



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
	resp, r, err := apiClient.CloudAccountAPIAPI.CloudApiControllerGetAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAccountAPIAPI.CloudApiControllerGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAccountAPIAPI.CloudApiControllerGetAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetAccountRequest struct via the builder pattern


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


## CloudApiControllerSignin

> map[string]interface{} CloudApiControllerSignin(ctx).Code(code).State(state).Execute()

Api Controller Signin



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
	code := "code_example" // string | code of account
	state := "state_example" // string | state of account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAccountAPIAPI.CloudApiControllerSignin(context.Background()).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAccountAPIAPI.CloudApiControllerSignin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerSignin`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAccountAPIAPI.CloudApiControllerSignin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerSigninRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | code of account | 
 **state** | **string** | state of account | 

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


## CloudApiControllerSignout

> CloudControllersResponse CloudApiControllerSignout(ctx).Execute()

Api Controller Signout



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
	resp, r, err := apiClient.CloudAccountAPIAPI.CloudApiControllerSignout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAccountAPIAPI.CloudApiControllerSignout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerSignout`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAccountAPIAPI.CloudApiControllerSignout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerSignoutRequest struct via the builder pattern


### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

