# \NetworkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchGetNetwork**](NetworkAPI.md#SearchGetNetwork) | **Get** /v1/search/network | Get network/federation configuration
[**SearchUpdateNetwork**](NetworkAPI.md#SearchUpdateNetwork) | **Patch** /v1/search/network | Update network configuration



## SearchGetNetwork

> SearchNetwork SearchGetNetwork(ctx).Execute()

Get network/federation configuration

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
	resp, r, err := apiClient.NetworkAPI.SearchGetNetwork(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.SearchGetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetNetwork`: SearchNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.SearchGetNetwork`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetNetworkRequest struct via the builder pattern


### Return type

[**SearchNetwork**](SearchNetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateNetwork

> SearchNetwork SearchUpdateNetwork(ctx).SearchNetwork(searchNetwork).Execute()

Update network configuration

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
	searchNetwork := *openapiclient.NewSearchNetwork() // SearchNetwork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.SearchUpdateNetwork(context.Background()).SearchNetwork(searchNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.SearchUpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateNetwork`: SearchNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.SearchUpdateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchNetwork** | [**SearchNetwork**](SearchNetwork.md) |  | 

### Return type

[**SearchNetwork**](SearchNetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

