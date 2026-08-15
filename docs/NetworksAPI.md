# \NetworksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworks**](NetworksAPI.md#GetNetworks) | **Get** /v1/networks | Returns the caller&#39;s org overlay network on the Zero Trust fabric.
[**GetNetworksById**](NetworksAPI.md#GetNetworksById) | **Get** /v1/networks/{id} | Returns one overlay network by id, scoped to the caller&#39;s org.
[**GetNetworksRouters**](NetworksAPI.md#GetNetworksRouters) | **Get** /v1/networks/routers | Returns the Zero Trust routers the caller&#39;s org owns.



## GetNetworks

> NetworkList GetNetworks(ctx).Execute()

Returns the caller's org overlay network on the Zero Trust fabric.



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
	resp, r, err := apiClient.NetworksAPI.GetNetworks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworks`: NetworkList
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksRequest struct via the builder pattern


### Return type

[**NetworkList**](NetworkList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworksById

> NetworkView GetNetworksById(ctx, id).Execute()

Returns one overlay network by id, scoped to the caller's org.



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
	id := "id_example" // string | ID is the network id from the path. The URL is the addressing authority, so it binds from there whatever else the request carries.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworksById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworksById`: NetworkView
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the network id from the path. The URL is the addressing authority, so it binds from there whatever else the request carries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkView**](NetworkView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworksRouters

> RouterList GetNetworksRouters(ctx).Execute()

Returns the Zero Trust routers the caller's org owns.



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
	resp, r, err := apiClient.NetworksAPI.GetNetworksRouters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworksRouters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworksRouters`: RouterList
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworksRouters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksRoutersRequest struct via the builder pattern


### Return type

[**RouterList**](RouterList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

