# \BotAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBotByWildcard1**](BotAPI.md#DeleteBotByWildcard1) | **Delete** /v1/bot/{wildcard1} | Relay one of the bot runtime&#39;s own operational paths
[**GetBotByWildcard1**](BotAPI.md#GetBotByWildcard1) | **Get** /v1/bot/{wildcard1} | Relay one of the bot runtime&#39;s own operational paths
[**GetBotConnect**](BotAPI.md#GetBotConnect) | **Get** /v1/bot/connect | The socket a bot node dials and holds open to become invokable.
[**GetBotNodes**](BotAPI.md#GetBotNodes) | **Get** /v1/bot/nodes | Returns the caller org&#39;s currently connected bot nodes: what each one calls itself, the platform it runs on, its agent version, when its socket was established, and the capabilities and commands it reported.
[**PatchBotByWildcard1**](BotAPI.md#PatchBotByWildcard1) | **Patch** /v1/bot/{wildcard1} | Relay one of the bot runtime&#39;s own operational paths
[**PostBotByWildcard1**](BotAPI.md#PostBotByWildcard1) | **Post** /v1/bot/{wildcard1} | Relay one of the bot runtime&#39;s own operational paths
[**PostBotNodesByIdInvoke**](BotAPI.md#PostBotNodesByIdInvoke) | **Post** /v1/bot/nodes/{id}/invoke | Ask one of your connected machines to run a command, and get its answer back.
[**PostBotPeerInvoke**](BotAPI.md#PostBotPeerInvoke) | **Post** /v1/bot/peer/invoke | Replica-to-replica forward of one invocation to the pod holding the node&#39;s socket.
[**PutBotByWildcard1**](BotAPI.md#PutBotByWildcard1) | **Put** /v1/bot/{wildcard1} | Relay one of the bot runtime&#39;s own operational paths



## DeleteBotByWildcard1

> DeleteBotByWildcard1(ctx, wildcard1).Execute()

Relay one of the bot runtime's own operational paths



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.DeleteBotByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.DeleteBotByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBotByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBotByWildcard1

> GetBotByWildcard1(ctx, wildcard1).Execute()

Relay one of the bot runtime's own operational paths



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.GetBotByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBotByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBotConnect

> GetBotConnect(ctx).Execute()

The socket a bot node dials and holds open to become invokable.



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
	r, err := apiClient.BotAPI.GetBotConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBotConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotConnectRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBotNodes

> NodesView GetBotNodes(ctx).Execute()

Returns the caller org's currently connected bot nodes: what each one calls itself, the platform it runs on, its agent version, when its socket was established, and the capabilities and commands it reported.



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
	resp, r, err := apiClient.BotAPI.GetBotNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBotNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBotNodes`: NodesView
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetBotNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotNodesRequest struct via the builder pattern


### Return type

[**NodesView**](NodesView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchBotByWildcard1

> PatchBotByWildcard1(ctx, wildcard1).Execute()

Relay one of the bot runtime's own operational paths



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.PatchBotByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.PatchBotByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBotByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBotByWildcard1

> PostBotByWildcard1(ctx, wildcard1).Execute()

Relay one of the bot runtime's own operational paths



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.PostBotByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.PostBotByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBotByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBotNodesByIdInvoke

> PostBotNodesByIdInvoke(ctx, id).Execute()

Ask one of your connected machines to run a command, and get its answer back.



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.PostBotNodesByIdInvoke(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.PostBotNodesByIdInvoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBotNodesByIdInvokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBotPeerInvoke

> PostBotPeerInvoke(ctx).Execute()

Replica-to-replica forward of one invocation to the pod holding the node's socket.



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
	r, err := apiClient.BotAPI.PostBotPeerInvoke(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.PostBotPeerInvoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBotPeerInvokeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBotByWildcard1

> PutBotByWildcard1(ctx, wildcard1).Execute()

Relay one of the bot runtime's own operational paths



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.PutBotByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.PutBotByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBotByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

