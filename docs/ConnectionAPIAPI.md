# \ConnectionAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddConnection**](ConnectionAPIAPI.md#CloudApiControllerAddConnection) | **Post** /v1/cloud/add-connection | Api Controller Add Connection
[**CloudApiControllerAddNodeTunnel**](ConnectionAPIAPI.md#CloudApiControllerAddNodeTunnel) | **Get** /v1/cloud/add-node-tunnel | Api Controller Add Node Tunnel
[**CloudApiControllerDeleteConnection**](ConnectionAPIAPI.md#CloudApiControllerDeleteConnection) | **Post** /v1/cloud/delete-connection | Api Controller Delete Connection
[**CloudApiControllerGetConnection**](ConnectionAPIAPI.md#CloudApiControllerGetConnection) | **Get** /v1/cloud/get-connection | Api Controller Get Connection
[**CloudApiControllerGetConnections**](ConnectionAPIAPI.md#CloudApiControllerGetConnections) | **Get** /v1/cloud/get-connections | Api Controller Get Connections
[**CloudApiControllerGetNodeTunnel**](ConnectionAPIAPI.md#CloudApiControllerGetNodeTunnel) | **Get** /v1/cloud/get-node-tunnel | Api Controller Get Node Tunnel
[**CloudApiControllerStartConnection**](ConnectionAPIAPI.md#CloudApiControllerStartConnection) | **Post** /v1/cloud/start-connection | Api Controller Start Connection
[**CloudApiControllerStopConnection**](ConnectionAPIAPI.md#CloudApiControllerStopConnection) | **Post** /v1/cloud/stop-connection | Api Controller Stop Connection
[**CloudApiControllerUpdateConnection**](ConnectionAPIAPI.md#CloudApiControllerUpdateConnection) | **Post** /v1/cloud/update-connection | Api Controller Update Connection



## CloudApiControllerAddConnection

> map[string]interface{} CloudApiControllerAddConnection(ctx).CloudObjectConnection(cloudObjectConnection).Execute()

Api Controller Add Connection



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
	cloudObjectConnection := *openapiclient.NewCloudObjectConnection() // CloudObjectConnection | The connection object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerAddConnection(context.Background()).CloudObjectConnection(cloudObjectConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerAddConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerAddConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectConnection** | [**CloudObjectConnection**](CloudObjectConnection.md) | The connection object | 

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


## CloudApiControllerAddNodeTunnel

> map[string]interface{} CloudApiControllerAddNodeTunnel(ctx).NodeId(nodeId).Execute()

Api Controller Add Node Tunnel



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
	nodeId := "nodeId_example" // string | The id of node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerAddNodeTunnel(context.Background()).NodeId(nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerAddNodeTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddNodeTunnel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerAddNodeTunnel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddNodeTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **string** | The id of node | 

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


## CloudApiControllerDeleteConnection

> map[string]interface{} CloudApiControllerDeleteConnection(ctx).Id(id).CloudObjectConnection(cloudObjectConnection).Execute()

Api Controller Delete Connection



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
	id := "id_example" // string | The id of connection
	cloudObjectConnection := *openapiclient.NewCloudObjectConnection() // CloudObjectConnection | The connection to delete (handler binds object.Connection; owner and name identify it).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerDeleteConnection(context.Background()).Id(id).CloudObjectConnection(cloudObjectConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerDeleteConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerDeleteConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of connection | 
 **cloudObjectConnection** | [**CloudObjectConnection**](CloudObjectConnection.md) | The connection to delete (handler binds object.Connection; owner and name identify it). | 

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


## CloudApiControllerGetConnection

> CloudObjectConnection CloudApiControllerGetConnection(ctx).Id(id).Execute()

Api Controller Get Connection



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
	id := "id_example" // string | The id of connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerGetConnection(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerGetConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetConnection`: CloudObjectConnection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerGetConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of connection | 

### Return type

[**CloudObjectConnection**](CloudObjectConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetConnections

> CloudObjectConnection CloudApiControllerGetConnections(ctx).PageSize(pageSize).P(p).Execute()

Api Controller Get Connections



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
	pageSize := "pageSize_example" // string | The size of each page
	p := "p_example" // string | The number of the page

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerGetConnections(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerGetConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetConnections`: CloudObjectConnection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerGetConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

### Return type

[**CloudObjectConnection**](CloudObjectConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetNodeTunnel

> map[string]interface{} CloudApiControllerGetNodeTunnel(ctx).Width(width).Height(height).Dpi(dpi).ConnectionId(connectionId).Username(username).Password(password).Execute()

Api Controller Get Node Tunnel



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
	width := "width_example" // string | The width of the tunnel
	height := "height_example" // string | The height of the tunnel
	dpi := "dpi_example" // string | The dpi of the tunnel
	connectionId := "connectionId_example" // string | The id of the connectionId
	username := "username_example" // string | The username for the tunnel
	password := "password_example" // string | The password for the tunnel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerGetNodeTunnel(context.Background()).Width(width).Height(height).Dpi(dpi).ConnectionId(connectionId).Username(username).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerGetNodeTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetNodeTunnel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerGetNodeTunnel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetNodeTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **width** | **string** | The width of the tunnel | 
 **height** | **string** | The height of the tunnel | 
 **dpi** | **string** | The dpi of the tunnel | 
 **connectionId** | **string** | The id of the connectionId | 
 **username** | **string** | The username for the tunnel | 
 **password** | **string** | The password for the tunnel | 

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


## CloudApiControllerStartConnection

> map[string]interface{} CloudApiControllerStartConnection(ctx).Id(id).Execute()

Api Controller Start Connection



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
	id := "id_example" // string | The id of connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerStartConnection(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerStartConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerStartConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerStartConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerStartConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of connection | 

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


## CloudApiControllerStopConnection

> map[string]interface{} CloudApiControllerStopConnection(ctx).Id(id).Execute()

Api Controller Stop Connection



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
	id := "id_example" // string | The id of connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerStopConnection(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerStopConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerStopConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerStopConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerStopConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of connection | 

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


## CloudApiControllerUpdateConnection

> map[string]interface{} CloudApiControllerUpdateConnection(ctx).Id(id).CloudObjectConnection(cloudObjectConnection).Execute()

Api Controller Update Connection



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
	id := "id_example" // string | The id of connection
	cloudObjectConnection := *openapiclient.NewCloudObjectConnection() // CloudObjectConnection | The connection object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionAPIAPI.CloudApiControllerUpdateConnection(context.Background()).Id(id).CloudObjectConnection(cloudObjectConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionAPIAPI.CloudApiControllerUpdateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionAPIAPI.CloudApiControllerUpdateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of connection | 
 **cloudObjectConnection** | [**CloudObjectConnection**](CloudObjectConnection.md) | The connection object | 

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

