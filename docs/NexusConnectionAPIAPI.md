# \NexusConnectionAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddConnection**](NexusConnectionAPIAPI.md#NexusAddConnection) | **Post** /v1/nexus/add-connection | add Connection
[**NexusAddNodeTunnel**](NexusConnectionAPIAPI.md#NexusAddNodeTunnel) | **Get** /v1/nexus/add-node-tunnel | add Node Tunnel
[**NexusDeleteConnection**](NexusConnectionAPIAPI.md#NexusDeleteConnection) | **Post** /v1/nexus/delete-connection | delete Connection
[**NexusGetConnection**](NexusConnectionAPIAPI.md#NexusGetConnection) | **Get** /v1/nexus/get-connection | get Connection
[**NexusGetConnections**](NexusConnectionAPIAPI.md#NexusGetConnections) | **Get** /v1/nexus/get-connections | get Connections
[**NexusGetNodeTunnel**](NexusConnectionAPIAPI.md#NexusGetNodeTunnel) | **Get** /v1/nexus/get-node-tunnel | get Node Tunnel
[**NexusStartConnection**](NexusConnectionAPIAPI.md#NexusStartConnection) | **Post** /v1/nexus/start-connection | start Connection
[**NexusStopConnection**](NexusConnectionAPIAPI.md#NexusStopConnection) | **Post** /v1/nexus/stop-connection | stop Connection
[**NexusUpdateConnection**](NexusConnectionAPIAPI.md#NexusUpdateConnection) | **Post** /v1/nexus/update-connection | update Connection



## NexusAddConnection

> map[string]interface{} NexusAddConnection(ctx).NexusConnection(nexusConnection).Execute()

add Connection



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
	nexusConnection := *openapiclient.NewNexusConnection() // NexusConnection | The connection object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusAddConnection(context.Background()).NexusConnection(nexusConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusAddConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusAddConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusConnection** | [**NexusConnection**](NexusConnection.md) | The connection object | 

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


## NexusAddNodeTunnel

> map[string]interface{} NexusAddNodeTunnel(ctx).NodeId(nodeId).Execute()

add Node Tunnel



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
	nodeId := "nodeId_example" // string | The id of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusAddNodeTunnel(context.Background()).NodeId(nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusAddNodeTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddNodeTunnel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusAddNodeTunnel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddNodeTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **string** | The id of the node | 

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


## NexusDeleteConnection

> map[string]interface{} NexusDeleteConnection(ctx).Id(id).Execute()

delete Connection



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
	id := "id_example" // string | The id of the connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusDeleteConnection(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusDeleteConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusDeleteConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the connection | 

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


## NexusGetConnection

> NexusConnection NexusGetConnection(ctx).Id(id).Execute()

get Connection



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
	id := "id_example" // string | The id of the connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusGetConnection(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusGetConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetConnection`: NexusConnection
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusGetConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the connection | 

### Return type

[**NexusConnection**](NexusConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetConnections

> NexusConnection NexusGetConnections(ctx).PageSize(pageSize).P(p).Execute()

get Connections



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
	p := "p_example" // string | The page number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusGetConnections(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusGetConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetConnections`: NexusConnection
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusGetConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The page number | 

### Return type

[**NexusConnection**](NexusConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetNodeTunnel

> map[string]interface{} NexusGetNodeTunnel(ctx).Width(width).Height(height).Dpi(dpi).ConnectionId(connectionId).Username(username).Password(password).Execute()

get Node Tunnel



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
	connectionId := "connectionId_example" // string | The id of the connection
	username := "username_example" // string | The username for the tunnel
	password := "password_example" // string | The password for the tunnel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusGetNodeTunnel(context.Background()).Width(width).Height(height).Dpi(dpi).ConnectionId(connectionId).Username(username).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusGetNodeTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetNodeTunnel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusGetNodeTunnel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetNodeTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **width** | **string** | The width of the tunnel | 
 **height** | **string** | The height of the tunnel | 
 **dpi** | **string** | The dpi of the tunnel | 
 **connectionId** | **string** | The id of the connection | 
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


## NexusStartConnection

> map[string]interface{} NexusStartConnection(ctx).Id(id).Execute()

start Connection



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
	id := "id_example" // string | The id of the connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusStartConnection(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusStartConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusStartConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusStartConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusStartConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the connection | 

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


## NexusStopConnection

> map[string]interface{} NexusStopConnection(ctx).Id(id).Execute()

stop Connection



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
	id := "id_example" // string | The id of the connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusStopConnection(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusStopConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusStopConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusStopConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusStopConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the connection | 

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


## NexusUpdateConnection

> map[string]interface{} NexusUpdateConnection(ctx).Id(id).NexusConnection(nexusConnection).Execute()

update Connection



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
	id := "id_example" // string | The id of the connection
	nexusConnection := *openapiclient.NewNexusConnection() // NexusConnection | The connection object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusConnectionAPIAPI.NexusUpdateConnection(context.Background()).Id(id).NexusConnection(nexusConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusConnectionAPIAPI.NexusUpdateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusConnectionAPIAPI.NexusUpdateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the connection | 
 **nexusConnection** | [**NexusConnection**](NexusConnection.md) | The connection object | 

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

