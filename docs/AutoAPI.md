# \AutoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1AutoFlowsFlow**](AutoAPI.md#CloudDeleteV1AutoFlowsFlow) | **Delete** /v1/auto/flows/{flow} | FlowDelete deletes one of the caller&#39;s flows.
[**CloudGetV1AutoFlows**](AutoAPI.md#CloudGetV1AutoFlows) | **Get** /v1/auto/flows | Flows lists the caller&#39;s flows, newest first.
[**CloudGetV1AutoFlowsFlow**](AutoAPI.md#CloudGetV1AutoFlowsFlow) | **Get** /v1/auto/flows/{flow} | Flow reads one of the caller&#39;s flows — the full record, graph included.
[**CloudGetV1AutoPieces**](AutoAPI.md#CloudGetV1AutoPieces) | **Get** /v1/auto/pieces | Pieces lists the product&#39;s built-in piece catalog: the trigger and action types a flow&#39;s nodes can use (webhook, schedule, http, set, branch), each with its input descriptors.
[**CloudGetV1AutoRuns**](AutoAPI.md#CloudGetV1AutoRuns) | **Get** /v1/auto/runs | Runs lists the caller&#39;s run records, newest first — optionally one flow&#39;s.
[**CloudGetV1AutoRunsRun**](AutoAPI.md#CloudGetV1AutoRunsRun) | **Get** /v1/auto/runs/{run} | Run reads one run record: status, input, output (each executed node&#39;s result keyed by node id once completed), error detail if it failed, and timestamps.
[**CloudGetV1AutoStatus**](AutoAPI.md#CloudGetV1AutoStatus) | **Get** /v1/auto/status | Status reports whether the auto service is reachable — its own health endpoint as an honest lens for \&quot;is the automation plane up\&quot;.
[**CloudPatchV1AutoFlowsFlow**](AutoAPI.md#CloudPatchV1AutoFlowsFlow) | **Patch** /v1/auto/flows/{flow} | FlowUpdate patches one of the caller&#39;s flows: the name, the graph, or both — only the stated fields move.
[**CloudPostV1AutoFlows**](AutoAPI.md#CloudPostV1AutoFlows) | **Post** /v1/auto/flows | FlowCreate creates a flow in the caller&#39;s org.
[**CloudPostV1AutoFlowsFlowPublish**](AutoAPI.md#CloudPostV1AutoFlowsFlowPublish) | **Post** /v1/auto/flows/{flow}/publish | Publish snapshots the flow&#39;s current graph as its next immutable version and arms the flow&#39;s triggers.
[**CloudPostV1AutoRuns**](AutoAPI.md#CloudPostV1AutoRuns) | **Post** /v1/auto/runs | Start begins one asynchronous run of a flow: the product dispatches the graph to its durable execution engine (the hanzo tasks plane) and answers immediately with the run record in status running.



## CloudDeleteV1AutoFlowsFlow

> interface{} CloudDeleteV1AutoFlowsFlow(ctx, flow).Execute()

FlowDelete deletes one of the caller's flows.



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
	resp, r, err := apiClient.AutoAPI.CloudDeleteV1AutoFlowsFlow(context.Background(), flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudDeleteV1AutoFlowsFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1AutoFlowsFlow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudDeleteV1AutoFlowsFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flow** | **string** | Flow is the flow&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1AutoFlowsFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutoFlows

> interface{} CloudGetV1AutoFlows(ctx).Execute()

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
	resp, r, err := apiClient.AutoAPI.CloudGetV1AutoFlows(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudGetV1AutoFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutoFlows`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudGetV1AutoFlows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutoFlowsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutoFlowsFlow

> interface{} CloudGetV1AutoFlowsFlow(ctx, flow).Execute()

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
	resp, r, err := apiClient.AutoAPI.CloudGetV1AutoFlowsFlow(context.Background(), flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudGetV1AutoFlowsFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutoFlowsFlow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudGetV1AutoFlowsFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flow** | **string** | Flow is the flow&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutoFlowsFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutoPieces

> interface{} CloudGetV1AutoPieces(ctx).Execute()

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
	resp, r, err := apiClient.AutoAPI.CloudGetV1AutoPieces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudGetV1AutoPieces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutoPieces`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudGetV1AutoPieces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutoPiecesRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutoRuns

> interface{} CloudGetV1AutoRuns(ctx).Flow(flow).Execute()

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
	resp, r, err := apiClient.AutoAPI.CloudGetV1AutoRuns(context.Background()).Flow(flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudGetV1AutoRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutoRuns`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudGetV1AutoRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutoRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | Flow narrows the list to one flow&#39;s runs when present. | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutoRunsRun

> interface{} CloudGetV1AutoRunsRun(ctx, run).Execute()

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
	resp, r, err := apiClient.AutoAPI.CloudGetV1AutoRunsRun(context.Background(), run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudGetV1AutoRunsRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutoRunsRun`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudGetV1AutoRunsRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**run** | **string** | Run is the run&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutoRunsRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutoStatus

> CloudAutoStatus CloudGetV1AutoStatus(ctx).Execute()

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
	resp, r, err := apiClient.AutoAPI.CloudGetV1AutoStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudGetV1AutoStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutoStatus`: CloudAutoStatus
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudGetV1AutoStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutoStatusRequest struct via the builder pattern


### Return type

[**CloudAutoStatus**](CloudAutoStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1AutoFlowsFlow

> interface{} CloudPatchV1AutoFlowsFlow(ctx, flow).CloudAutoUpdate(cloudAutoUpdate).Execute()

FlowUpdate patches one of the caller's flows: the name, the graph, or both — only the stated fields move.



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
	cloudAutoUpdate := *openapiclient.NewCloudAutoUpdate() // CloudAutoUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.CloudPatchV1AutoFlowsFlow(context.Background(), flow).CloudAutoUpdate(cloudAutoUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudPatchV1AutoFlowsFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1AutoFlowsFlow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudPatchV1AutoFlowsFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flow** | **string** | Flow is the flow&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1AutoFlowsFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAutoUpdate** | [**CloudAutoUpdate**](CloudAutoUpdate.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutoFlows

> interface{} CloudPostV1AutoFlows(ctx).CloudAutoCreate(cloudAutoCreate).Execute()

FlowCreate creates a flow in the caller's org.



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
	cloudAutoCreate := *openapiclient.NewCloudAutoCreate() // CloudAutoCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.CloudPostV1AutoFlows(context.Background()).CloudAutoCreate(cloudAutoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudPostV1AutoFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutoFlows`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudPostV1AutoFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutoFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAutoCreate** | [**CloudAutoCreate**](CloudAutoCreate.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutoFlowsFlowPublish

> interface{} CloudPostV1AutoFlowsFlowPublish(ctx, flow).Execute()

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
	resp, r, err := apiClient.AutoAPI.CloudPostV1AutoFlowsFlowPublish(context.Background(), flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudPostV1AutoFlowsFlowPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutoFlowsFlowPublish`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudPostV1AutoFlowsFlowPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flow** | **string** | Flow is the flow&#39;s id, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutoFlowsFlowPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutoRuns

> interface{} CloudPostV1AutoRuns(ctx).CloudAutoStart(cloudAutoStart).Execute()

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
	cloudAutoStart := *openapiclient.NewCloudAutoStart() // CloudAutoStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.CloudPostV1AutoRuns(context.Background()).CloudAutoStart(cloudAutoStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.CloudPostV1AutoRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutoRuns`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.CloudPostV1AutoRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutoRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAutoStart** | [**CloudAutoStart**](CloudAutoStart.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

