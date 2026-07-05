# \AutoAiProvidersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoDeleteAiProvider**](AutoAiProvidersAPI.md#AutoDeleteAiProvider) | **Delete** /v1/auto/ai-providers/{provider} | Delete an AI provider
[**AutoListAiProviders**](AutoAiProvidersAPI.md#AutoListAiProviders) | **Get** /v1/auto/ai-providers | List configured AI providers
[**AutoUpsertAiProvider**](AutoAiProvidersAPI.md#AutoUpsertAiProvider) | **Post** /v1/auto/ai-providers | Add or update an AI provider



## AutoDeleteAiProvider

> AutoDeleteAiProvider(ctx, provider).Execute()

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
	r, err := apiClient.AutoAiProvidersAPI.AutoDeleteAiProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAiProvidersAPI.AutoDeleteAiProvider``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAutoDeleteAiProviderRequest struct via the builder pattern


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


## AutoListAiProviders

> map[string]interface{} AutoListAiProviders(ctx).Execute()

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
	resp, r, err := apiClient.AutoAiProvidersAPI.AutoListAiProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAiProvidersAPI.AutoListAiProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListAiProviders`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAiProvidersAPI.AutoListAiProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutoListAiProvidersRequest struct via the builder pattern


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


## AutoUpsertAiProvider

> map[string]interface{} AutoUpsertAiProvider(ctx).AutoUpsertAiProviderRequest(autoUpsertAiProviderRequest).Execute()

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
	resp, r, err := apiClient.AutoAiProvidersAPI.AutoUpsertAiProvider(context.Background()).AutoUpsertAiProviderRequest(autoUpsertAiProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAiProvidersAPI.AutoUpsertAiProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoUpsertAiProvider`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAiProvidersAPI.AutoUpsertAiProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoUpsertAiProviderRequest struct via the builder pattern


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

