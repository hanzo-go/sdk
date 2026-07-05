# \NexusProviderAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddProvider**](NexusProviderAPIAPI.md#NexusAddProvider) | **Post** /v1/nexus/add-provider | add Provider
[**NexusDeleteProvider**](NexusProviderAPIAPI.md#NexusDeleteProvider) | **Post** /v1/nexus/delete-provider | delete Provider
[**NexusGetGlobalProviders**](NexusProviderAPIAPI.md#NexusGetGlobalProviders) | **Get** /v1/nexus/get-global-providers | get Global Providers
[**NexusGetProvider**](NexusProviderAPIAPI.md#NexusGetProvider) | **Get** /v1/nexus/get-provider | get Provider
[**NexusGetProviders**](NexusProviderAPIAPI.md#NexusGetProviders) | **Get** /v1/nexus/get-providers | get Providers
[**NexusRefreshMcpTools**](NexusProviderAPIAPI.md#NexusRefreshMcpTools) | **Post** /v1/nexus/refresh-mcp-tools | refresh Mcp Tools
[**NexusUpdateProvider**](NexusProviderAPIAPI.md#NexusUpdateProvider) | **Post** /v1/nexus/update-provider | update Provider



## NexusAddProvider

> NexusResponse NexusAddProvider(ctx).NexusProvider(nexusProvider).Execute()

add Provider



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
	nexusProvider := *openapiclient.NewNexusProvider() // NexusProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusProviderAPIAPI.NexusAddProvider(context.Background()).NexusProvider(nexusProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusProviderAPIAPI.NexusAddProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddProvider`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusProviderAPIAPI.NexusAddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusProvider** | [**NexusProvider**](NexusProvider.md) | The details of the provider | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteProvider

> NexusResponse NexusDeleteProvider(ctx).NexusProvider(nexusProvider).Execute()

delete Provider



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
	nexusProvider := *openapiclient.NewNexusProvider() // NexusProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusProviderAPIAPI.NexusDeleteProvider(context.Background()).NexusProvider(nexusProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusProviderAPIAPI.NexusDeleteProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteProvider`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusProviderAPIAPI.NexusDeleteProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusProvider** | [**NexusProvider**](NexusProvider.md) | The details of the provider | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalProviders

> []NexusProvider NexusGetGlobalProviders(ctx).Execute()

get Global Providers



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
	resp, r, err := apiClient.NexusProviderAPIAPI.NexusGetGlobalProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusProviderAPIAPI.NexusGetGlobalProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalProviders`: []NexusProvider
	fmt.Fprintf(os.Stdout, "Response from `NexusProviderAPIAPI.NexusGetGlobalProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalProvidersRequest struct via the builder pattern


### Return type

[**[]NexusProvider**](NexusProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetProvider

> NexusProvider NexusGetProvider(ctx).Id(id).Execute()

get Provider



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
	id := "id_example" // string | The id of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusProviderAPIAPI.NexusGetProvider(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusProviderAPIAPI.NexusGetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetProvider`: NexusProvider
	fmt.Fprintf(os.Stdout, "Response from `NexusProviderAPIAPI.NexusGetProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the provider | 

### Return type

[**NexusProvider**](NexusProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetProviders

> []NexusProvider NexusGetProviders(ctx).Execute()

get Providers



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
	resp, r, err := apiClient.NexusProviderAPIAPI.NexusGetProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusProviderAPIAPI.NexusGetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetProviders`: []NexusProvider
	fmt.Fprintf(os.Stdout, "Response from `NexusProviderAPIAPI.NexusGetProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetProvidersRequest struct via the builder pattern


### Return type

[**[]NexusProvider**](NexusProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusRefreshMcpTools

> NexusResponse NexusRefreshMcpTools(ctx).NexusProvider(nexusProvider).Execute()

refresh Mcp Tools



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
	nexusProvider := *openapiclient.NewNexusProvider() // NexusProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusProviderAPIAPI.NexusRefreshMcpTools(context.Background()).NexusProvider(nexusProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusProviderAPIAPI.NexusRefreshMcpTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusRefreshMcpTools`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusProviderAPIAPI.NexusRefreshMcpTools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusRefreshMcpToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusProvider** | [**NexusProvider**](NexusProvider.md) | The details of the provider | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateProvider

> NexusResponse NexusUpdateProvider(ctx).Id(id).NexusProvider(nexusProvider).Execute()

update Provider



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
	nexusProvider := *openapiclient.NewNexusProvider() // NexusProvider | The details of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusProviderAPIAPI.NexusUpdateProvider(context.Background()).Id(id).NexusProvider(nexusProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusProviderAPIAPI.NexusUpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateProvider`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusProviderAPIAPI.NexusUpdateProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the provider | 
 **nexusProvider** | [**NexusProvider**](NexusProvider.md) | The details of the provider | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

