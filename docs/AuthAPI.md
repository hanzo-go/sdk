# \AuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsGetTelemetryScript**](AuthAPI.md#AnalyticsGetTelemetryScript) | **Get** /v1/analytics/scripts/telemetry | Get telemetry pixel script
[**AnalyticsHeartbeat**](AuthAPI.md#AnalyticsHeartbeat) | **Get** /v1/analytics/heartbeat | Health check
[**AnalyticsLogin**](AuthAPI.md#AnalyticsLogin) | **Post** /v1/analytics/auth/login | Log in with username and password
[**AnalyticsLogout**](AuthAPI.md#AnalyticsLogout) | **Post** /v1/analytics/auth/logout | Log out and invalidate token
[**AnalyticsSsoAuth**](AuthAPI.md#AnalyticsSsoAuth) | **Post** /v1/analytics/auth/sso | Exchange SSO credentials for a session token
[**AnalyticsVerifyAuth**](AuthAPI.md#AnalyticsVerifyAuth) | **Post** /v1/analytics/auth/verify | Verify current auth token
[**BotAuthCallback**](AuthAPI.md#BotAuthCallback) | **Get** /v1/bot/auth/callback | OAuth callback - exchange code for session
[**BotAuthLogin**](AuthAPI.md#BotAuthLogin) | **Get** /v1/bot/auth/login | Initiate OAuth login via Hanzo IAM
[**BotAuthLogout**](AuthAPI.md#BotAuthLogout) | **Post** /v1/bot/auth/logout | Invalidate current session
[**BotAuthMe**](AuthAPI.md#BotAuthMe) | **Get** /v1/bot/auth/me | Get current authenticated user
[**BotWhoami**](AuthAPI.md#BotWhoami) | **Get** /v1/bot/whoami | CLI alias for /v1/bot/auth/me
[**CommerceAuthenticate**](AuthAPI.md#CommerceAuthenticate) | **Post** /v1/commerce/auth | Authenticate user (OAuth2)



## AnalyticsGetTelemetryScript

> string AnalyticsGetTelemetryScript(ctx).Execute()

Get telemetry pixel script

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
	resp, r, err := apiClient.AuthAPI.AnalyticsGetTelemetryScript(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AnalyticsGetTelemetryScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTelemetryScript`: string
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AnalyticsGetTelemetryScript`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetTelemetryScriptRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsHeartbeat

> AnalyticsHeartbeat200Response AnalyticsHeartbeat(ctx).Execute()

Health check

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
	resp, r, err := apiClient.AuthAPI.AnalyticsHeartbeat(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AnalyticsHeartbeat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsHeartbeat`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AnalyticsHeartbeat`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsHeartbeatRequest struct via the builder pattern


### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsLogin

> AnalyticsLogin200Response AnalyticsLogin(ctx).AnalyticsLoginRequest(analyticsLoginRequest).Execute()

Log in with username and password

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
	analyticsLoginRequest := *openapiclient.NewAnalyticsLoginRequest("Username_example", "Password_example") // AnalyticsLoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AnalyticsLogin(context.Background()).AnalyticsLoginRequest(analyticsLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AnalyticsLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsLogin`: AnalyticsLogin200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AnalyticsLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsLoginRequest** | [**AnalyticsLoginRequest**](AnalyticsLoginRequest.md) |  | 

### Return type

[**AnalyticsLogin200Response**](AnalyticsLogin200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsLogout

> map[string]interface{} AnalyticsLogout(ctx).Execute()

Log out and invalidate token

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
	resp, r, err := apiClient.AuthAPI.AnalyticsLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AnalyticsLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsLogout`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AnalyticsLogout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsLogoutRequest struct via the builder pattern


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


## AnalyticsSsoAuth

> AnalyticsSsoAuth200Response AnalyticsSsoAuth(ctx).Execute()

Exchange SSO credentials for a session token

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
	resp, r, err := apiClient.AuthAPI.AnalyticsSsoAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AnalyticsSsoAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsSsoAuth`: AnalyticsSsoAuth200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AnalyticsSsoAuth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsSsoAuthRequest struct via the builder pattern


### Return type

[**AnalyticsSsoAuth200Response**](AnalyticsSsoAuth200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsVerifyAuth

> AnalyticsUser AnalyticsVerifyAuth(ctx).Execute()

Verify current auth token

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
	resp, r, err := apiClient.AuthAPI.AnalyticsVerifyAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AnalyticsVerifyAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsVerifyAuth`: AnalyticsUser
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AnalyticsVerifyAuth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsVerifyAuthRequest struct via the builder pattern


### Return type

[**AnalyticsUser**](AnalyticsUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	r, err := apiClient.AuthAPI.BotAuthCallback(context.Background()).Code(code).State(state).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.BotAuthCallback``: %v\n", err)
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
	r, err := apiClient.AuthAPI.BotAuthLogin(context.Background()).RedirectUri(redirectUri).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.BotAuthLogin``: %v\n", err)
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
	resp, r, err := apiClient.AuthAPI.BotAuthLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.BotAuthLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotAuthLogout`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.BotAuthLogout`: %v\n", resp)
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
	resp, r, err := apiClient.AuthAPI.BotAuthMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.BotAuthMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotAuthMe`: BotUser
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.BotAuthMe`: %v\n", resp)
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
	resp, r, err := apiClient.AuthAPI.BotWhoami(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.BotWhoami``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotWhoami`: BotUser
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.BotWhoami`: %v\n", resp)
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
	resp, r, err := apiClient.AuthAPI.CommerceAuthenticate(context.Background()).CommerceOAuthRequest(commerceOAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.CommerceAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceAuthenticate`: CommerceOAuthResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.CommerceAuthenticate`: %v\n", resp)
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

