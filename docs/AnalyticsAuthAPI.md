# \AnalyticsAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsGetTelemetryScript**](AnalyticsAuthAPI.md#AnalyticsGetTelemetryScript) | **Get** /v1/analytics/scripts/telemetry | Get telemetry pixel script
[**AnalyticsHeartbeat**](AnalyticsAuthAPI.md#AnalyticsHeartbeat) | **Get** /v1/analytics/heartbeat | Health check
[**AnalyticsLogin**](AnalyticsAuthAPI.md#AnalyticsLogin) | **Post** /v1/analytics/auth/login | Log in with username and password
[**AnalyticsLogout**](AnalyticsAuthAPI.md#AnalyticsLogout) | **Post** /v1/analytics/auth/logout | Log out and invalidate token
[**AnalyticsSsoAuth**](AnalyticsAuthAPI.md#AnalyticsSsoAuth) | **Post** /v1/analytics/auth/sso | Exchange SSO credentials for a session token
[**AnalyticsVerifyAuth**](AnalyticsAuthAPI.md#AnalyticsVerifyAuth) | **Post** /v1/analytics/auth/verify | Verify current auth token



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
	resp, r, err := apiClient.AnalyticsAuthAPI.AnalyticsGetTelemetryScript(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAuthAPI.AnalyticsGetTelemetryScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTelemetryScript`: string
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAuthAPI.AnalyticsGetTelemetryScript`: %v\n", resp)
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
	resp, r, err := apiClient.AnalyticsAuthAPI.AnalyticsHeartbeat(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAuthAPI.AnalyticsHeartbeat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsHeartbeat`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAuthAPI.AnalyticsHeartbeat`: %v\n", resp)
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
	resp, r, err := apiClient.AnalyticsAuthAPI.AnalyticsLogin(context.Background()).AnalyticsLoginRequest(analyticsLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAuthAPI.AnalyticsLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsLogin`: AnalyticsLogin200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAuthAPI.AnalyticsLogin`: %v\n", resp)
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
	resp, r, err := apiClient.AnalyticsAuthAPI.AnalyticsLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAuthAPI.AnalyticsLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsLogout`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAuthAPI.AnalyticsLogout`: %v\n", resp)
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
	resp, r, err := apiClient.AnalyticsAuthAPI.AnalyticsSsoAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAuthAPI.AnalyticsSsoAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsSsoAuth`: AnalyticsSsoAuth200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAuthAPI.AnalyticsSsoAuth`: %v\n", resp)
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
	resp, r, err := apiClient.AnalyticsAuthAPI.AnalyticsVerifyAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAuthAPI.AnalyticsVerifyAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsVerifyAuth`: AnalyticsUser
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAuthAPI.AnalyticsVerifyAuth`: %v\n", resp)
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

