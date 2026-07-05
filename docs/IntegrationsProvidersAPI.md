# \IntegrationsProvidersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsGetProvider**](IntegrationsProvidersAPI.md#IntegrationsGetProvider) | **Get** /v1/integrations/{provider} | Get one provider with this org&#39;s connection status
[**IntegrationsListProviders**](IntegrationsProvidersAPI.md#IntegrationsListProviders) | **Get** /v1/integrations | List providers with this org&#39;s connection status



## IntegrationsGetProvider

> IntegrationsProviderView IntegrationsGetProvider(ctx, provider).Execute()

Get one provider with this org's connection status

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
	provider := "provider_example" // string | Provider slug (e.g. slack, github)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsProvidersAPI.IntegrationsGetProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsProvidersAPI.IntegrationsGetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsGetProvider`: IntegrationsProviderView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsProvidersAPI.IntegrationsGetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider slug (e.g. slack, github) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationsProviderView**](IntegrationsProviderView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsListProviders

> IntegrationsListProviders200Response IntegrationsListProviders(ctx).Execute()

List providers with this org's connection status

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
	resp, r, err := apiClient.IntegrationsProvidersAPI.IntegrationsListProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsProvidersAPI.IntegrationsListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsListProviders`: IntegrationsListProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsProvidersAPI.IntegrationsListProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsListProvidersRequest struct via the builder pattern


### Return type

[**IntegrationsListProviders200Response**](IntegrationsListProviders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

