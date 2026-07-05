# \IamProvidersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddProvider**](IamProvidersAPI.md#IamApiControllerAddProvider) | **Post** /v1/iam/providers | Api Controller Add Provider
[**IamApiControllerDeleteProvider**](IamProvidersAPI.md#IamApiControllerDeleteProvider) | **Delete** /v1/iam/providers/{id} | Api Controller Delete Provider
[**IamApiControllerGetGlobalProviders**](IamProvidersAPI.md#IamApiControllerGetGlobalProviders) | **Get** /v1/iam/global-providers | Api Controller Get Global Providers
[**IamApiControllerGetProvider**](IamProvidersAPI.md#IamApiControllerGetProvider) | **Get** /v1/iam/providers/{id} | Api Controller Get Provider
[**IamApiControllerGetProviders**](IamProvidersAPI.md#IamApiControllerGetProviders) | **Get** /v1/iam/providers | Api Controller Get Providers
[**IamApiControllerUpdateProvider**](IamProvidersAPI.md#IamApiControllerUpdateProvider) | **Put** /v1/iam/providers/{id} | Api Controller Update Provider



## IamApiControllerAddProvider

> IamControllersResponse IamApiControllerAddProvider(ctx).IamObjectProvider(iamObjectProvider).Execute()

Api Controller Add Provider



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
	iamObjectProvider := *openapiclient.NewIamObjectProvider() // IamObjectProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamProvidersAPI.IamApiControllerAddProvider(context.Background()).IamObjectProvider(iamObjectProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamProvidersAPI.IamApiControllerAddProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddProvider`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamProvidersAPI.IamApiControllerAddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectProvider** | [**IamObjectProvider**](IamObjectProvider.md) | The details of the provider | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteProvider

> IamControllersResponse IamApiControllerDeleteProvider(ctx, id).IamObjectProvider(iamObjectProvider).Execute()

Api Controller Delete Provider



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectProvider := *openapiclient.NewIamObjectProvider() // IamObjectProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamProvidersAPI.IamApiControllerDeleteProvider(context.Background(), id).IamObjectProvider(iamObjectProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamProvidersAPI.IamApiControllerDeleteProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteProvider`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamProvidersAPI.IamApiControllerDeleteProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectProvider** | [**IamObjectProvider**](IamObjectProvider.md) | The details of the provider | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetGlobalProviders

> []IamObjectProvider IamApiControllerGetGlobalProviders(ctx).Execute()

Api Controller Get Global Providers



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
	resp, r, err := apiClient.IamProvidersAPI.IamApiControllerGetGlobalProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamProvidersAPI.IamApiControllerGetGlobalProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGlobalProviders`: []IamObjectProvider
	fmt.Fprintf(os.Stdout, "Response from `IamProvidersAPI.IamApiControllerGetGlobalProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetGlobalProvidersRequest struct via the builder pattern


### Return type

[**[]IamObjectProvider**](IamObjectProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetProvider

> IamObjectProvider IamApiControllerGetProvider(ctx, id).Execute()

Api Controller Get Provider



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
	id := "id_example" // string | The id ( owner/name ) of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamProvidersAPI.IamApiControllerGetProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamProvidersAPI.IamApiControllerGetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetProvider`: IamObjectProvider
	fmt.Fprintf(os.Stdout, "Response from `IamProvidersAPI.IamApiControllerGetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectProvider**](IamObjectProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetProviders

> []IamObjectProvider IamApiControllerGetProviders(ctx).Owner(owner).Execute()

Api Controller Get Providers



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
	owner := "owner_example" // string | The owner of providers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamProvidersAPI.IamApiControllerGetProviders(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamProvidersAPI.IamApiControllerGetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetProviders`: []IamObjectProvider
	fmt.Fprintf(os.Stdout, "Response from `IamProvidersAPI.IamApiControllerGetProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of providers | 

### Return type

[**[]IamObjectProvider**](IamObjectProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateProvider

> IamControllersResponse IamApiControllerUpdateProvider(ctx, id).IamObjectProvider(iamObjectProvider).Execute()

Api Controller Update Provider



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
	id := "id_example" // string | The id ( owner/name ) of the provider
	iamObjectProvider := *openapiclient.NewIamObjectProvider() // IamObjectProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamProvidersAPI.IamApiControllerUpdateProvider(context.Background(), id).IamObjectProvider(iamObjectProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamProvidersAPI.IamApiControllerUpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateProvider`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamProvidersAPI.IamApiControllerUpdateProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectProvider** | [**IamObjectProvider**](IamObjectProvider.md) | The details of the provider | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

