# \NetworkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetwork**](NetworkAPI.md#GetNetwork) | **Get** /v1/network | Returns the caller&#39;s org overlay network on the Zero Trust fabric.
[**GetNetworkById**](NetworkAPI.md#GetNetworkById) | **Get** /v1/network/{id} | Returns one overlay network by id, scoped to the caller&#39;s org.
[**GetNetworkRouters**](NetworkAPI.md#GetNetworkRouters) | **Get** /v1/network/routers | Returns the Zero Trust routers the caller&#39;s org owns.
[**GetNetworkServices**](NetworkAPI.md#GetNetworkServices) | **Get** /v1/network/services | Returns the Zero Trust edge services the caller&#39;s org owns.



## GetNetwork

> NetworkList GetNetwork(ctx).Execute()

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
	resp, r, err := apiClient.NetworkAPI.GetNetwork(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetwork`: NetworkList
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetwork`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


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


## GetNetworkById

> NetworkView GetNetworkById(ctx, id).Execute()

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
	resp, r, err := apiClient.NetworkAPI.GetNetworkById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkById`: NetworkView
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the network id from the path. The URL is the addressing authority, so it binds from there whatever else the request carries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkByIdRequest struct via the builder pattern


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


## GetNetworkRouters

> RouterList GetNetworkRouters(ctx).Execute()

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
	resp, r, err := apiClient.NetworkAPI.GetNetworkRouters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkRouters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouters`: RouterList
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkRouters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersRequest struct via the builder pattern


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


## GetNetworkServices

> MeshServiceList GetNetworkServices(ctx).Execute()

Returns the Zero Trust edge services the caller's org owns.



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
	resp, r, err := apiClient.NetworkAPI.GetNetworkServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkServices`: MeshServiceList
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkServicesRequest struct via the builder pattern


### Return type

[**MeshServiceList**](MeshServiceList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

