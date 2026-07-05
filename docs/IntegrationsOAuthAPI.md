# \IntegrationsOAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsConnectProvider**](IntegrationsOAuthAPI.md#IntegrationsConnectProvider) | **Post** /v1/integrations/{provider}/connect | Begin an OAuth flow (returns the provider authorize URL)
[**IntegrationsDisconnectProvider**](IntegrationsOAuthAPI.md#IntegrationsDisconnectProvider) | **Post** /v1/integrations/{provider}/disconnect | Revoke and forget an org&#39;s connection (idempotent)
[**IntegrationsProviderCallback**](IntegrationsOAuthAPI.md#IntegrationsProviderCallback) | **Get** /v1/integrations/{provider}/callback | OAuth return — seal tokens and redirect to console



## IntegrationsConnectProvider

> IntegrationsConnectProvider200Response IntegrationsConnectProvider(ctx, provider).Execute()

Begin an OAuth flow (returns the provider authorize URL)

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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsOAuthAPI.IntegrationsConnectProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsOAuthAPI.IntegrationsConnectProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsConnectProvider`: IntegrationsConnectProvider200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsOAuthAPI.IntegrationsConnectProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsConnectProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationsConnectProvider200Response**](IntegrationsConnectProvider200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsDisconnectProvider

> IntegrationsDisconnectProvider200Response IntegrationsDisconnectProvider(ctx, provider).Execute()

Revoke and forget an org's connection (idempotent)

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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsOAuthAPI.IntegrationsDisconnectProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsOAuthAPI.IntegrationsDisconnectProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsDisconnectProvider`: IntegrationsDisconnectProvider200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsOAuthAPI.IntegrationsDisconnectProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsDisconnectProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationsDisconnectProvider200Response**](IntegrationsDisconnectProvider200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsProviderCallback

> IntegrationsProviderCallback(ctx, provider).State(state).Code(code).Error_(error_).Execute()

OAuth return — seal tokens and redirect to console



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
	state := "state_example" // string | 
	provider := "provider_example" // string | 
	code := "code_example" // string |  (optional)
	error_ := "error__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsOAuthAPI.IntegrationsProviderCallback(context.Background(), provider).State(state).Code(code).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsOAuthAPI.IntegrationsProviderCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsProviderCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** |  | 

 **code** | **string** |  | 
 **error_** | **string** |  | 

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

