# \AutoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutoFlowsById**](AutoAPI.md#DeleteAutoFlowsById) | **Delete** /v1/auto/flows/{id} | Deletes one automation, its versions and its run history.
[**GetAutoConnectors**](AutoAPI.md#GetAutoConnectors) | **Get** /v1/auto/connectors | Connectors returns the connector catalogue.
[**GetAutoFlows**](AutoAPI.md#GetAutoFlows) | **Get** /v1/auto/flows | Returns the caller org&#39;s automations, most-recently-updated first.
[**GetAutoFlowsById**](AutoAPI.md#GetAutoFlowsById) | **Get** /v1/auto/flows/{id} | Returns one automation and its latest version.
[**GetAutoFlowsByIdVersions**](AutoAPI.md#GetAutoFlowsByIdVersions) | **Get** /v1/auto/flows/{id}/versions | Returns one flow&#39;s versions, newest first.
[**GetAutoRuns**](AutoAPI.md#GetAutoRuns) | **Get** /v1/auto/runs | Returns the caller org&#39;s run history, newest first.
[**GetAutoRunsById**](AutoAPI.md#GetAutoRunsById) | **Get** /v1/auto/runs/{id} | Returns one run.
[**PatchAutoFlowsById**](AutoAPI.md#PatchAutoFlowsById) | **Patch** /v1/auto/flows/{id} | Updates one automation&#39;s metadata in place.
[**PostAutoConnectorsByIdRun**](AutoAPI.md#PostAutoConnectorsByIdRun) | **Post** /v1/auto/connectors/{id}/run | Run executes one connector action in-process and answers the outcome.
[**PostAutoFlows**](AutoAPI.md#PostAutoFlows) | **Post** /v1/auto/flows | Creates an automation and its initial DRAFT version in one call.
[**PostAutoFlowsByIdDisable**](AutoAPI.md#PostAutoFlowsByIdDisable) | **Post** /v1/auto/flows/{id}/disable | Disarms a flow&#39;s trigger and marks it DISABLED.
[**PostAutoFlowsByIdEnable**](AutoAPI.md#PostAutoFlowsByIdEnable) | **Post** /v1/auto/flows/{id}/enable | Arms a flow&#39;s trigger and marks it ENABLED.
[**PostAutoFlowsByIdOperations**](AutoAPI.md#PostAutoFlowsByIdOperations) | **Post** /v1/auto/flows/{id}/operations | Edit a flow — rename it, retarget its trigger, or add, move and delete steps
[**PostAutoFlowsByIdRun**](AutoAPI.md#PostAutoFlowsByIdRun) | **Post** /v1/auto/flows/{id}/run | Starts one durable run of a flow now.
[**PostAutoFlowsByIdVersions**](AutoAPI.md#PostAutoFlowsByIdVersions) | **Post** /v1/auto/flows/{id}/versions | Adds a new DRAFT version to a flow.
[**PostAutoHooksBySourceByEvent**](AutoAPI.md#PostAutoHooksBySourceByEvent) | **Post** /v1/auto/hooks/{source}/{event} | Fire an event that starts every enabled flow subscribed to it
[**PostAutoRunsByIdResume**](AutoAPI.md#PostAutoRunsByIdResume) | **Post** /v1/auto/runs/{id}/resume | Release a run waiting at an approval step, with the approval payload



## DeleteAutoFlowsById

> DeleteAutoFlowsById(ctx, id).Execute()

Deletes one automation, its versions and its run history.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoAPI.DeleteAutoFlowsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.DeleteAutoFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoFlowsByIdRequest struct via the builder pattern


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


## GetAutoConnectors

> Catalog GetAutoConnectors(ctx).Execute()

Connectors returns the connector catalogue.



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
	resp, r, err := apiClient.AutoAPI.GetAutoConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoConnectors`: Catalog
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoConnectorsRequest struct via the builder pattern


### Return type

[**Catalog**](Catalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoFlows

> FlowPage GetAutoFlows(ctx).Limit(limit).Execute()

Returns the caller org's automations, most-recently-updated first.



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
	limit := int32(56) // int32 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoFlows(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoFlows`: FlowPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**FlowPage**](FlowPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoFlowsById

> PopulatedFlow GetAutoFlowsById(ctx, id).Execute()

Returns one automation and its latest version.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoFlowsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoFlowsById`: PopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoFlowsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoFlowsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PopulatedFlow**](PopulatedFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoFlowsByIdVersions

> VersionPage GetAutoFlowsByIdVersions(ctx, id).Limit(limit).Execute()

Returns one flow's versions, newest first.



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
	id := "flow_1" // string | ID is the flow whose versions to list, from the path.
	limit := int32(56) // int32 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoFlowsByIdVersions(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoFlowsByIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoFlowsByIdVersions`: VersionPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoFlowsByIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow whose versions to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoFlowsByIdVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**VersionPage**](VersionPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoRuns

> RunPage GetAutoRuns(ctx).FlowId(flowId).Limit(limit).Execute()

Returns the caller org's run history, newest first.



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
	flowId := "flowId_example" // string | FlowID narrows the history to one flow. Omit it for the whole org's runs. (optional)
	limit := int32(56) // int32 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoRuns(context.Background()).FlowId(flowId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoRuns`: RunPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** | FlowID narrows the history to one flow. Omit it for the whole org&#39;s runs. | 
 **limit** | **int32** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**RunPage**](RunPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoRunsById

> FlowRun GetAutoRunsById(ctx, id).Execute()

Returns one run.



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
	id := "run_1" // string | ID is the run to read, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoRunsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoRunsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoRunsById`: FlowRun
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoRunsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the run to read, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoRunsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowRun**](FlowRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutoFlowsById

> Flow PatchAutoFlowsById(ctx, id).PatchFlowIn(patchFlowIn).Execute()

Updates one automation's metadata in place.



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
	id := "flow_1" // string | ID is the flow to update, from the path.
	patchFlowIn := *openapiclient.NewPatchFlowIn() // PatchFlowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PatchAutoFlowsById(context.Background(), id).PatchFlowIn(patchFlowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PatchAutoFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutoFlowsById`: Flow
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PatchAutoFlowsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutoFlowsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchFlowIn** | [**PatchFlowIn**](PatchFlowIn.md) |  | 

### Return type

[**Flow**](Flow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoConnectorsByIdRun

> RunResp PostAutoConnectorsByIdRun(ctx, id).RunIn(runIn).Execute()

Run executes one connector action in-process and answers the outcome.



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
	id := "notion" // string | ID is the connector to run, from the path.
	runIn := *openapiclient.NewRunIn() // RunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoConnectorsByIdRun(context.Background(), id).RunIn(runIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoConnectorsByIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoConnectorsByIdRun`: RunResp
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoConnectorsByIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector to run, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoConnectorsByIdRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runIn** | [**RunIn**](RunIn.md) |  | 

### Return type

[**RunResp**](RunResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlows

> PopulatedFlow PostAutoFlows(ctx).CreateFlowReq(createFlowReq).Execute()

Creates an automation and its initial DRAFT version in one call.



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
	createFlowReq := *openapiclient.NewCreateFlowReq() // CreateFlowReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlows(context.Background()).CreateFlowReq(createFlowReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlows`: PopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFlowReq** | [**CreateFlowReq**](CreateFlowReq.md) |  | 

### Return type

[**PopulatedFlow**](PopulatedFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlowsByIdDisable

> Flow PostAutoFlowsByIdDisable(ctx, id).Execute()

Disarms a flow's trigger and marks it DISABLED.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlowsByIdDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlowsByIdDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlowsByIdDisable`: Flow
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoFlowsByIdDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoFlowsByIdDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Flow**](Flow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlowsByIdEnable

> Flow PostAutoFlowsByIdEnable(ctx, id).Execute()

Arms a flow's trigger and marks it ENABLED.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlowsByIdEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlowsByIdEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlowsByIdEnable`: Flow
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoFlowsByIdEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoFlowsByIdEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Flow**](Flow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlowsByIdOperations

> PostAutoFlowsByIdOperations(ctx, id).Execute()

Edit a flow — rename it, retarget its trigger, or add, move and delete steps



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoAPI.PostAutoFlowsByIdOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlowsByIdOperations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAutoFlowsByIdOperationsRequest struct via the builder pattern


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


## PostAutoFlowsByIdRun

> FlowRun PostAutoFlowsByIdRun(ctx, id).Execute()

Starts one durable run of a flow now.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlowsByIdRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlowsByIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlowsByIdRun`: FlowRun
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoFlowsByIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoFlowsByIdRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowRun**](FlowRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlowsByIdVersions

> FlowVersion PostAutoFlowsByIdVersions(ctx, id).CreateVersionIn(createVersionIn).Execute()

Adds a new DRAFT version to a flow.



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
	id := "flow_1" // string | ID is the flow to add a version to, from the path.
	createVersionIn := *openapiclient.NewCreateVersionIn() // CreateVersionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlowsByIdVersions(context.Background(), id).CreateVersionIn(createVersionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlowsByIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlowsByIdVersions`: FlowVersion
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoFlowsByIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to add a version to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoFlowsByIdVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVersionIn** | [**CreateVersionIn**](CreateVersionIn.md) |  | 

### Return type

[**FlowVersion**](FlowVersion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoHooksBySourceByEvent

> PostAutoHooksBySourceByEvent(ctx, source, event).Execute()

Fire an event that starts every enabled flow subscribed to it



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
	source := "source_example" // string | 
	event := "event_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoAPI.PostAutoHooksBySourceByEvent(context.Background(), source, event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoHooksBySourceByEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** |  | 
**event** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoHooksBySourceByEventRequest struct via the builder pattern


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


## PostAutoRunsByIdResume

> PostAutoRunsByIdResume(ctx, id).Execute()

Release a run waiting at an approval step, with the approval payload



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoAPI.PostAutoRunsByIdResume(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoRunsByIdResume``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAutoRunsByIdResumeRequest struct via the builder pattern


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

