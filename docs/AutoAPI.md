# \AutoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutoFlowsByFlow**](AutoAPI.md#DeleteAutoFlowsByFlow) | **Delete** /v1/auto/flows/{flow} | Deletes one of the caller&#39;s flows.
[**GetAutoFlows**](AutoAPI.md#GetAutoFlows) | **Get** /v1/auto/flows | Flows lists the caller&#39;s flows, newest first.
[**GetAutoFlowsByFlow**](AutoAPI.md#GetAutoFlowsByFlow) | **Get** /v1/auto/flows/{flow} | Flow reads one of the caller&#39;s flows — the full record, graph included.
[**GetAutoPieces**](AutoAPI.md#GetAutoPieces) | **Get** /v1/auto/pieces | Pieces lists the product&#39;s built-in piece catalog: the trigger and action types a flow&#39;s nodes can use (webhook, schedule, http, set, branch), each with its input descriptors.
[**GetAutoRuns**](AutoAPI.md#GetAutoRuns) | **Get** /v1/auto/runs | Runs lists the caller&#39;s run records, newest first — optionally one flow&#39;s.
[**GetAutoRunsByRun**](AutoAPI.md#GetAutoRunsByRun) | **Get** /v1/auto/runs/{run} | Run reads one run record: status, input, output (each executed node&#39;s result keyed by node id once completed), error detail if it failed, and timestamps.
[**GetAutoStatus**](AutoAPI.md#GetAutoStatus) | **Get** /v1/auto/status | Status reports whether the auto service is reachable — its own health endpoint as an honest lens for \&quot;is the automation plane up\&quot;.
[**PatchAutoFlowsByFlow**](AutoAPI.md#PatchAutoFlowsByFlow) | **Patch** /v1/auto/flows/{flow} | Patches one of the caller&#39;s flows: the name, the graph, or both — only the stated fields move.
[**PostAutoFlows**](AutoAPI.md#PostAutoFlows) | **Post** /v1/auto/flows | Creates a flow in the caller&#39;s org.
[**PostAutoFlowsByFlowPublish**](AutoAPI.md#PostAutoFlowsByFlowPublish) | **Post** /v1/auto/flows/{flow}/publish | Publish snapshots the flow&#39;s current graph as its next immutable version and arms the flow&#39;s triggers.
[**PostAutoRuns**](AutoAPI.md#PostAutoRuns) | **Post** /v1/auto/runs | Start begins one asynchronous run of a flow: the product dispatches the graph to its durable execution engine (the hanzo tasks plane) and answers immediately with the run record in status running.



## DeleteAutoFlowsByFlow

> interface{} DeleteAutoFlowsByFlow(ctx, flow).Execute()

Deletes one of the caller's flows.



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
	flow := "flow_example" // string | Flow is the flow's id, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.DeleteAutoFlowsByFlow(context.Background(), flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.DeleteAutoFlowsByFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAutoFlowsByFlow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.DeleteAutoFlowsByFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flow** | **string** | Flow is the flow&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoFlowsByFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoFlows

> interface{} GetAutoFlows(ctx).Execute()

Flows lists the caller's flows, newest first.



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
	resp, r, err := apiClient.AutoAPI.GetAutoFlows(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoFlows`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoFlows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoFlowsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoFlowsByFlow

> interface{} GetAutoFlowsByFlow(ctx, flow).Execute()

Flow reads one of the caller's flows — the full record, graph included.



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
	flow := "flow_example" // string | Flow is the flow's id, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoFlowsByFlow(context.Background(), flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoFlowsByFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoFlowsByFlow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoFlowsByFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flow** | **string** | Flow is the flow&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoFlowsByFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoPieces

> interface{} GetAutoPieces(ctx).Execute()

Pieces lists the product's built-in piece catalog: the trigger and action types a flow's nodes can use (webhook, schedule, http, set, branch), each with its input descriptors.



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
	resp, r, err := apiClient.AutoAPI.GetAutoPieces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoPieces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoPieces`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoPieces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoPiecesRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoRuns

> interface{} GetAutoRuns(ctx).Flow(flow).Execute()

Runs lists the caller's run records, newest first — optionally one flow's.



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
	flow := "flow_example" // string | Flow narrows the list to one flow's runs when present. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoRuns(context.Background()).Flow(flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoRuns`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | Flow narrows the list to one flow&#39;s runs when present. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoRunsByRun

> interface{} GetAutoRunsByRun(ctx, run).Execute()

Run reads one run record: status, input, output (each executed node's result keyed by node id once completed), error detail if it failed, and timestamps.



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
	run := "run_example" // string | Run is the run's id, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoRunsByRun(context.Background(), run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoRunsByRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoRunsByRun`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoRunsByRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**run** | **string** | Run is the run&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoRunsByRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoStatus

> AutoStatus GetAutoStatus(ctx).Execute()

Status reports whether the auto service is reachable — its own health endpoint as an honest lens for \"is the automation plane up\".



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
	resp, r, err := apiClient.AutoAPI.GetAutoStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoStatus`: AutoStatus
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoStatusRequest struct via the builder pattern


### Return type

[**AutoStatus**](AutoStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutoFlowsByFlow

> interface{} PatchAutoFlowsByFlow(ctx, flow).AutoUpdate(autoUpdate).Execute()

Patches one of the caller's flows: the name, the graph, or both — only the stated fields move.



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
	flow := "flow_example" // string | Flow is the flow's id, taken from the path.
	autoUpdate := *openapiclient.NewAutoUpdate() // AutoUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PatchAutoFlowsByFlow(context.Background(), flow).AutoUpdate(autoUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PatchAutoFlowsByFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutoFlowsByFlow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PatchAutoFlowsByFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flow** | **string** | Flow is the flow&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutoFlowsByFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoUpdate** | [**AutoUpdate**](AutoUpdate.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlows

> interface{} PostAutoFlows(ctx).AutoCreate(autoCreate).Execute()

Creates a flow in the caller's org.



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
	autoCreate := *openapiclient.NewAutoCreate() // AutoCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlows(context.Background()).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlows`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreate** | [**AutoCreate**](AutoCreate.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlowsByFlowPublish

> interface{} PostAutoFlowsByFlowPublish(ctx, flow).Execute()

Publish snapshots the flow's current graph as its next immutable version and arms the flow's triggers.



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
	flow := "flow_example" // string | Flow is the flow's id, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlowsByFlowPublish(context.Background(), flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlowsByFlowPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlowsByFlowPublish`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoFlowsByFlowPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flow** | **string** | Flow is the flow&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoFlowsByFlowPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoRuns

> interface{} PostAutoRuns(ctx).AutoStart(autoStart).Execute()

Start begins one asynchronous run of a flow: the product dispatches the graph to its durable execution engine (the hanzo tasks plane) and answers immediately with the run record in status running.



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
	autoStart := *openapiclient.NewAutoStart() // AutoStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoRuns(context.Background()).AutoStart(autoStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoRuns`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoStart** | [**AutoStart**](AutoStart.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

