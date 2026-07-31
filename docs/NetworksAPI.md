# \NetworksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Networks**](NetworksAPI.md#CloudGetV1Networks) | **Get** /v1/networks | Returns the caller&#39;s org overlay network on the Zero Trust fabric.
[**CloudGetV1NetworksId**](NetworksAPI.md#CloudGetV1NetworksId) | **Get** /v1/networks/{id} | Returns one overlay network by id, scoped to the caller&#39;s org.



## CloudGetV1Networks

> CloudNetworkList CloudGetV1Networks(ctx).Execute()

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
	resp, r, err := apiClient.NetworksAPI.CloudGetV1Networks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CloudGetV1Networks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Networks`: CloudNetworkList
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CloudGetV1Networks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1NetworksRequest struct via the builder pattern


### Return type

[**CloudNetworkList**](CloudNetworkList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1NetworksId

> CloudNetworkView CloudGetV1NetworksId(ctx, id).Execute()

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
	resp, r, err := apiClient.NetworksAPI.CloudGetV1NetworksId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CloudGetV1NetworksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1NetworksId`: CloudNetworkView
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CloudGetV1NetworksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the network id from the path. The URL is the addressing authority, so it binds from there whatever else the request carries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1NetworksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudNetworkView**](CloudNetworkView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

