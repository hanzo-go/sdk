# \NetworkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkIdentitiesById**](NetworkAPI.md#DeleteNetworkIdentitiesById) | **Delete** /v1/network/identities/{id} | Removes one of the org&#39;s fabric identities.
[**GetNetwork**](NetworkAPI.md#GetNetwork) | **Get** /v1/network | Returns the caller&#39;s org overlay network on the Zero Trust fabric.
[**GetNetworkById**](NetworkAPI.md#GetNetworkById) | **Get** /v1/network/{id} | Returns one overlay network by id, scoped to the caller&#39;s org.
[**GetNetworkIdentities**](NetworkAPI.md#GetNetworkIdentities) | **Get** /v1/network/identities | Returns the fabric identities the caller&#39;s org owns.
[**GetNetworkRouters**](NetworkAPI.md#GetNetworkRouters) | **Get** /v1/network/routers | Returns the Zero Trust routers the caller&#39;s org owns.
[**GetNetworkServices**](NetworkAPI.md#GetNetworkServices) | **Get** /v1/network/services | Returns the Zero Trust edge services the caller&#39;s org owns.
[**PostNetworkIdentities**](NetworkAPI.md#PostNetworkIdentities) | **Post** /v1/network/identities | Mints a fabric identity for a device the caller&#39;s org brings.
[**PostNetworkServices**](NetworkAPI.md#PostNetworkServices) | **Post** /v1/network/services | Puts a name on the org&#39;s overlay: a fabric service forwarding to host:port on whichever of the org&#39;s devices carries the \&quot;&lt;name&gt;-host\&quot; role, dialable at \&quot;&lt;name&gt;.&lt;org&gt;.zt\&quot; by any of the org&#39;s identities — and by the cloud&#39;s own, which is what lets a BYO cluster&#39;s apiserver be attached to the fleet with a \&quot;.zt\&quot; kubeconfig.



## DeleteNetworkIdentitiesById

> DeleteNetworkIdentitiesById(ctx, id).Execute()

Removes one of the org's fabric identities.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the identity id from the path. The URL is the addressing authority, so it binds from there whatever else the request carries.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkAPI.DeleteNetworkIdentitiesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.DeleteNetworkIdentitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the identity id from the path. The URL is the addressing authority, so it binds from there whatever else the request carries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkIdentitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## GetNetworkIdentities

> IdentityList GetNetworkIdentities(ctx).Execute()

Returns the fabric identities the caller's org owns.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkIdentities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkIdentities`: IdentityList
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkIdentities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkIdentitiesRequest struct via the builder pattern


### Return type

[**IdentityList**](IdentityList.md)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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


## PostNetworkIdentities

> IdentityView PostNetworkIdentities(ctx).IdentityIn(identityIn).Execute()

Mints a fabric identity for a device the caller's org brings.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	identityIn := *openapiclient.NewIdentityIn() // IdentityIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.PostNetworkIdentities(context.Background()).IdentityIn(identityIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.PostNetworkIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworkIdentities`: IdentityView
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.PostNetworkIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworkIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityIn** | [**IdentityIn**](IdentityIn.md) |  | 

### Return type

[**IdentityView**](IdentityView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNetworkServices

> PublishedView PostNetworkServices(ctx).ServiceIn(serviceIn).Execute()

Puts a name on the org's overlay: a fabric service forwarding to host:port on whichever of the org's devices carries the \"<name>-host\" role, dialable at \"<name>.<org>.zt\" by any of the org's identities — and by the cloud's own, which is what lets a BYO cluster's apiserver be attached to the fleet with a \".zt\" kubeconfig.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	serviceIn := *openapiclient.NewServiceIn() // ServiceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.PostNetworkServices(context.Background()).ServiceIn(serviceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.PostNetworkServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworkServices`: PublishedView
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.PostNetworkServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworkServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceIn** | [**ServiceIn**](ServiceIn.md) |  | 

### Return type

[**PublishedView**](PublishedView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

