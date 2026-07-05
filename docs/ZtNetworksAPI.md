# \ZtNetworksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZtGetNetwork**](ZtNetworksAPI.md#ZtGetNetwork) | **Get** /v1/networks/{id} | Get one overlay network by id
[**ZtListNetworks**](ZtNetworksAPI.md#ZtListNetworks) | **Get** /v1/networks | List the org&#39;s ZT overlay network(s)



## ZtGetNetwork

> ZtNetworkView ZtGetNetwork(ctx, id).Execute()

Get one overlay network by id

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
	id := "id_example" // string | The org-derived network id (org-<org>)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZtNetworksAPI.ZtGetNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZtNetworksAPI.ZtGetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZtGetNetwork`: ZtNetworkView
	fmt.Fprintf(os.Stdout, "Response from `ZtNetworksAPI.ZtGetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The org-derived network id (org-&lt;org&gt;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiZtGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ZtNetworkView**](ZtNetworkView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZtListNetworks

> ZtListNetworks200Response ZtListNetworks(ctx).Execute()

List the org's ZT overlay network(s)

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
	resp, r, err := apiClient.ZtNetworksAPI.ZtListNetworks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZtNetworksAPI.ZtListNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZtListNetworks`: ZtListNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `ZtNetworksAPI.ZtListNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiZtListNetworksRequest struct via the builder pattern


### Return type

[**ZtListNetworks200Response**](ZtListNetworks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

