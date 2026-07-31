# \ProviderAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddProvider**](ProviderAPIAPI.md#CloudApiControllerAddProvider) | **Post** /v1/cloud/add-provider | Api Controller Add Provider
[**CloudApiControllerDeleteProvider**](ProviderAPIAPI.md#CloudApiControllerDeleteProvider) | **Post** /v1/cloud/delete-provider | Api Controller Delete Provider
[**CloudApiControllerGetGlobalProviders**](ProviderAPIAPI.md#CloudApiControllerGetGlobalProviders) | **Get** /v1/cloud/get-global-providers | Api Controller Get Global Providers
[**CloudApiControllerGetProvider**](ProviderAPIAPI.md#CloudApiControllerGetProvider) | **Get** /v1/cloud/get-provider | Api Controller Get Provider
[**CloudApiControllerGetProviders**](ProviderAPIAPI.md#CloudApiControllerGetProviders) | **Get** /v1/cloud/get-providers | Api Controller Get Providers
[**CloudApiControllerRefreshMcpTools**](ProviderAPIAPI.md#CloudApiControllerRefreshMcpTools) | **Post** /v1/cloud/refresh-mcp-tools | Api Controller Refresh Mcp Tools
[**CloudApiControllerUpdateProvider**](ProviderAPIAPI.md#CloudApiControllerUpdateProvider) | **Post** /v1/cloud/update-provider | Api Controller Update Provider



## CloudApiControllerAddProvider

> CloudControllersResponse CloudApiControllerAddProvider(ctx).CloudObjectProvider(cloudObjectProvider).Execute()

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
	cloudObjectProvider := *openapiclient.NewCloudObjectProvider() // CloudObjectProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPIAPI.CloudApiControllerAddProvider(context.Background()).CloudObjectProvider(cloudObjectProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPIAPI.CloudApiControllerAddProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddProvider`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPIAPI.CloudApiControllerAddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectProvider** | [**CloudObjectProvider**](CloudObjectProvider.md) | The details of the provider | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteProvider

> CloudControllersResponse CloudApiControllerDeleteProvider(ctx).CloudObjectProvider(cloudObjectProvider).Execute()

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
	cloudObjectProvider := *openapiclient.NewCloudObjectProvider() // CloudObjectProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPIAPI.CloudApiControllerDeleteProvider(context.Background()).CloudObjectProvider(cloudObjectProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPIAPI.CloudApiControllerDeleteProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteProvider`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPIAPI.CloudApiControllerDeleteProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectProvider** | [**CloudObjectProvider**](CloudObjectProvider.md) | The details of the provider | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetGlobalProviders

> []CloudObjectProvider CloudApiControllerGetGlobalProviders(ctx).Execute()

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
	resp, r, err := apiClient.ProviderAPIAPI.CloudApiControllerGetGlobalProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPIAPI.CloudApiControllerGetGlobalProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalProviders`: []CloudObjectProvider
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPIAPI.CloudApiControllerGetGlobalProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalProvidersRequest struct via the builder pattern


### Return type

[**[]CloudObjectProvider**](CloudObjectProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetProvider

> CloudObjectProvider CloudApiControllerGetProvider(ctx).Id(id).Execute()

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
	id := "id_example" // string | The id of provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPIAPI.CloudApiControllerGetProvider(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPIAPI.CloudApiControllerGetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetProvider`: CloudObjectProvider
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPIAPI.CloudApiControllerGetProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of provider | 

### Return type

[**CloudObjectProvider**](CloudObjectProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetProviders

> []CloudObjectProvider CloudApiControllerGetProviders(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPIAPI.CloudApiControllerGetProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPIAPI.CloudApiControllerGetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetProviders`: []CloudObjectProvider
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPIAPI.CloudApiControllerGetProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetProvidersRequest struct via the builder pattern


### Return type

[**[]CloudObjectProvider**](CloudObjectProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerRefreshMcpTools

> CloudControllersResponse CloudApiControllerRefreshMcpTools(ctx).CloudObjectProvider(cloudObjectProvider).Execute()

Api Controller Refresh Mcp Tools



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
	cloudObjectProvider := *openapiclient.NewCloudObjectProvider() // CloudObjectProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPIAPI.CloudApiControllerRefreshMcpTools(context.Background()).CloudObjectProvider(cloudObjectProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPIAPI.CloudApiControllerRefreshMcpTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerRefreshMcpTools`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPIAPI.CloudApiControllerRefreshMcpTools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerRefreshMcpToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectProvider** | [**CloudObjectProvider**](CloudObjectProvider.md) | The details of the provider | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateProvider

> CloudControllersResponse CloudApiControllerUpdateProvider(ctx).Id(id).CloudObjectProvider(cloudObjectProvider).Execute()

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
	id := "id_example" // string | The id (owner/name) of the provider
	cloudObjectProvider := *openapiclient.NewCloudObjectProvider() // CloudObjectProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPIAPI.CloudApiControllerUpdateProvider(context.Background()).Id(id).CloudObjectProvider(cloudObjectProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPIAPI.CloudApiControllerUpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateProvider`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPIAPI.CloudApiControllerUpdateProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the provider | 
 **cloudObjectProvider** | [**CloudObjectProvider**](CloudObjectProvider.md) | The details of the provider | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

