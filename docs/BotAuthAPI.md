# \BotAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotAuthCallback**](BotAuthAPI.md#BotAuthCallback) | **Get** /v1/bot/auth/callback | OAuth callback - exchange code for session
[**BotAuthLogin**](BotAuthAPI.md#BotAuthLogin) | **Get** /v1/bot/auth/login | Initiate OAuth login via Hanzo IAM
[**BotAuthLogout**](BotAuthAPI.md#BotAuthLogout) | **Post** /v1/bot/auth/logout | Invalidate current session
[**BotAuthMe**](BotAuthAPI.md#BotAuthMe) | **Get** /v1/bot/auth/me | Get current authenticated user
[**BotWhoami**](BotAuthAPI.md#BotWhoami) | **Get** /v1/bot/whoami | CLI alias for /v1/bot/auth/me



## BotAuthCallback

> BotAuthCallback(ctx).Code(code).State(state).Error_(error_).Execute()

OAuth callback - exchange code for session

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
	code := "code_example" // string | Authorization code from IAM
	state := "state_example" // string |  (optional)
	error_ := "error__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAuthAPI.BotAuthCallback(context.Background()).Code(code).State(state).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAuthAPI.BotAuthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotAuthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Authorization code from IAM | 
 **state** | **string** |  | 
 **error_** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotAuthLogin

> BotAuthLogin(ctx).RedirectUri(redirectUri).State(state).Execute()

Initiate OAuth login via Hanzo IAM

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
	redirectUri := "redirectUri_example" // string | Override callback URL (optional)
	state := "state_example" // string | Opaque state for CSRF protection (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAuthAPI.BotAuthLogin(context.Background()).RedirectUri(redirectUri).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAuthAPI.BotAuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectUri** | **string** | Override callback URL | 
 **state** | **string** | Opaque state for CSRF protection | 

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


## BotAuthLogout

> AnalyticsHeartbeat200Response BotAuthLogout(ctx).Execute()

Invalidate current session

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
	resp, r, err := apiClient.BotAuthAPI.BotAuthLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAuthAPI.BotAuthLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotAuthLogout`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `BotAuthAPI.BotAuthLogout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBotAuthLogoutRequest struct via the builder pattern


### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotAuthMe

> BotUser BotAuthMe(ctx).Execute()

Get current authenticated user

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
	resp, r, err := apiClient.BotAuthAPI.BotAuthMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAuthAPI.BotAuthMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotAuthMe`: BotUser
	fmt.Fprintf(os.Stdout, "Response from `BotAuthAPI.BotAuthMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBotAuthMeRequest struct via the builder pattern


### Return type

[**BotUser**](BotUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotWhoami

> BotUser BotWhoami(ctx).Execute()

CLI alias for /v1/bot/auth/me

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
	resp, r, err := apiClient.BotAuthAPI.BotWhoami(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAuthAPI.BotWhoami``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotWhoami`: BotUser
	fmt.Fprintf(os.Stdout, "Response from `BotAuthAPI.BotWhoami`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBotWhoamiRequest struct via the builder pattern


### Return type

[**BotUser**](BotUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

