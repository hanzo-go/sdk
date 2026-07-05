# \FlowAiProvidersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowDeleteAiProvider**](FlowAiProvidersAPI.md#FlowDeleteAiProvider) | **Delete** /v1/flow/ai-providers/{provider} | Delete an AI provider
[**FlowListAiProviders**](FlowAiProvidersAPI.md#FlowListAiProviders) | **Get** /v1/flow/ai-providers | List configured AI providers
[**FlowUpsertAiProvider**](FlowAiProvidersAPI.md#FlowUpsertAiProvider) | **Post** /v1/flow/ai-providers | Add or update an AI provider



## FlowDeleteAiProvider

> FlowDeleteAiProvider(ctx, provider).Execute()

Delete an AI provider

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
	r, err := apiClient.FlowAiProvidersAPI.FlowDeleteAiProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAiProvidersAPI.FlowDeleteAiProvider``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowDeleteAiProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowListAiProviders

> map[string]interface{} FlowListAiProviders(ctx).Execute()

List configured AI providers

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
	resp, r, err := apiClient.FlowAiProvidersAPI.FlowListAiProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAiProvidersAPI.FlowListAiProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListAiProviders`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAiProvidersAPI.FlowListAiProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListAiProvidersRequest struct via the builder pattern


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


## FlowUpsertAiProvider

> map[string]interface{} FlowUpsertAiProvider(ctx).AutoUpsertAiProviderRequest(autoUpsertAiProviderRequest).Execute()

Add or update an AI provider

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
	autoUpsertAiProviderRequest := *openapiclient.NewAutoUpsertAiProviderRequest("Provider_example", "BaseUrl_example") // AutoUpsertAiProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAiProvidersAPI.FlowUpsertAiProvider(context.Background()).AutoUpsertAiProviderRequest(autoUpsertAiProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAiProvidersAPI.FlowUpsertAiProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpsertAiProvider`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAiProvidersAPI.FlowUpsertAiProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpsertAiProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoUpsertAiProviderRequest** | [**AutoUpsertAiProviderRequest**](AutoUpsertAiProviderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

