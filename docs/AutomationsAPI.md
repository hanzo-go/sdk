# \AutomationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutomationsFlowsById**](AutomationsAPI.md#DeleteAutomationsFlowsById) | **Delete** /v1/automations/flows/{id} | Deletes one automation, its versions and its run history.
[**GetAutomationsConnectors**](AutomationsAPI.md#GetAutomationsConnectors) | **Get** /v1/automations/connectors | Connectors returns the connector catalogue.
[**GetAutomationsFlows**](AutomationsAPI.md#GetAutomationsFlows) | **Get** /v1/automations/flows | Returns the caller org&#39;s automations, most-recently-updated first.
[**GetAutomationsFlowsById**](AutomationsAPI.md#GetAutomationsFlowsById) | **Get** /v1/automations/flows/{id} | Returns one automation and its latest version.
[**GetAutomationsFlowsByIdVersions**](AutomationsAPI.md#GetAutomationsFlowsByIdVersions) | **Get** /v1/automations/flows/{id}/versions | Returns one flow&#39;s versions, newest first.
[**GetAutomationsPieces**](AutomationsAPI.md#GetAutomationsPieces) | **Get** /v1/automations/pieces | Pieces is the retired-name alias of the connector catalogue.
[**GetAutomationsRuns**](AutomationsAPI.md#GetAutomationsRuns) | **Get** /v1/automations/runs | Returns the caller org&#39;s run history, newest first.
[**GetAutomationsRunsById**](AutomationsAPI.md#GetAutomationsRunsById) | **Get** /v1/automations/runs/{id} | Returns one run.
[**PatchAutomationsFlowsById**](AutomationsAPI.md#PatchAutomationsFlowsById) | **Patch** /v1/automations/flows/{id} | Updates one automation&#39;s metadata in place.
[**PostAutomationsConnectorsByIdRun**](AutomationsAPI.md#PostAutomationsConnectorsByIdRun) | **Post** /v1/automations/connectors/{id}/run | Run executes one connector action in-process and answers the outcome.
[**PostAutomationsFlows**](AutomationsAPI.md#PostAutomationsFlows) | **Post** /v1/automations/flows | Creates an automation and its initial DRAFT version in one call.
[**PostAutomationsFlowsByIdDisable**](AutomationsAPI.md#PostAutomationsFlowsByIdDisable) | **Post** /v1/automations/flows/{id}/disable | Disarms a flow&#39;s trigger and marks it DISABLED.
[**PostAutomationsFlowsByIdEnable**](AutomationsAPI.md#PostAutomationsFlowsByIdEnable) | **Post** /v1/automations/flows/{id}/enable | Arms a flow&#39;s trigger and marks it ENABLED.
[**PostAutomationsFlowsByIdOperations**](AutomationsAPI.md#PostAutomationsFlowsByIdOperations) | **Post** /v1/automations/flows/{id}/operations | Edit a flow — rename it, retarget its trigger, or add, move and delete steps
[**PostAutomationsFlowsByIdRun**](AutomationsAPI.md#PostAutomationsFlowsByIdRun) | **Post** /v1/automations/flows/{id}/run | Starts one durable run of a flow now.
[**PostAutomationsFlowsByIdVersions**](AutomationsAPI.md#PostAutomationsFlowsByIdVersions) | **Post** /v1/automations/flows/{id}/versions | Adds a new DRAFT version to a flow.
[**PostAutomationsHooksBySourceByEvent**](AutomationsAPI.md#PostAutomationsHooksBySourceByEvent) | **Post** /v1/automations/hooks/{source}/{event} | Fire an event that starts every enabled flow subscribed to it
[**PostAutomationsRunsByIdResume**](AutomationsAPI.md#PostAutomationsRunsByIdResume) | **Post** /v1/automations/runs/{id}/resume | Release a run waiting at an approval step, with the approval payload



## DeleteAutomationsFlowsById

> DeleteAutomationsFlowsById(ctx, id).Execute()

Deletes one automation, its versions and its run history.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationsAPI.DeleteAutomationsFlowsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.DeleteAutomationsFlowsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAutomationsFlowsByIdRequest struct via the builder pattern


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


## GetAutomationsConnectors

> Catalog GetAutomationsConnectors(ctx).Execute()

Connectors returns the connector catalogue.



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
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsConnectors`: Catalog
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsConnectorsRequest struct via the builder pattern


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


## GetAutomationsFlows

> FlowPage GetAutomationsFlows(ctx).Limit(limit).Execute()

Returns the caller org's automations, most-recently-updated first.



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
	limit := int32(56) // int32 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsFlows(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsFlows`: FlowPage
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsFlowsRequest struct via the builder pattern


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


## GetAutomationsFlowsById

> PopulatedFlow GetAutomationsFlowsById(ctx, id).Execute()

Returns one automation and its latest version.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsFlowsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsFlowsById`: PopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsFlowsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsFlowsByIdRequest struct via the builder pattern


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


## GetAutomationsFlowsByIdVersions

> VersionPage GetAutomationsFlowsByIdVersions(ctx, id).Limit(limit).Execute()

Returns one flow's versions, newest first.



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
	id := "flow_1" // string | ID is the flow whose versions to list, from the path.
	limit := int32(56) // int32 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsFlowsByIdVersions(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsFlowsByIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsFlowsByIdVersions`: VersionPage
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsFlowsByIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow whose versions to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsFlowsByIdVersionsRequest struct via the builder pattern


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


## GetAutomationsPieces

> Catalog GetAutomationsPieces(ctx).Execute()

Pieces is the retired-name alias of the connector catalogue.



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
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsPieces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsPieces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsPieces`: Catalog
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsPieces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsPiecesRequest struct via the builder pattern


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


## GetAutomationsRuns

> RunPage GetAutomationsRuns(ctx).FlowId(flowId).Limit(limit).Execute()

Returns the caller org's run history, newest first.



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
	flowId := "flowId_example" // string | FlowID narrows the history to one flow. Omit it for the whole org's runs. (optional)
	limit := int32(56) // int32 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsRuns(context.Background()).FlowId(flowId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsRuns`: RunPage
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsRunsRequest struct via the builder pattern


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


## GetAutomationsRunsById

> FlowRun GetAutomationsRunsById(ctx, id).Execute()

Returns one run.



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
	id := "run_1" // string | ID is the run to read, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsRunsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsRunsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsRunsById`: FlowRun
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsRunsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the run to read, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsRunsByIdRequest struct via the builder pattern


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


## PatchAutomationsFlowsById

> Flow PatchAutomationsFlowsById(ctx, id).PatchFlowIn(patchFlowIn).Execute()

Updates one automation's metadata in place.



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
	id := "flow_1" // string | ID is the flow to update, from the path.
	patchFlowIn := *openapiclient.NewPatchFlowIn() // PatchFlowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PatchAutomationsFlowsById(context.Background(), id).PatchFlowIn(patchFlowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PatchAutomationsFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutomationsFlowsById`: Flow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PatchAutomationsFlowsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutomationsFlowsByIdRequest struct via the builder pattern


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


## PostAutomationsConnectorsByIdRun

> RunResp PostAutomationsConnectorsByIdRun(ctx, id).RunIn(runIn).Execute()

Run executes one connector action in-process and answers the outcome.



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
	id := "notion" // string | ID is the connector to run, from the path.
	runIn := *openapiclient.NewRunIn() // RunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomationsConnectorsByIdRun(context.Background(), id).RunIn(runIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsConnectorsByIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationsConnectorsByIdRun`: RunResp
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomationsConnectorsByIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector to run, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsConnectorsByIdRunRequest struct via the builder pattern


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


## PostAutomationsFlows

> PopulatedFlow PostAutomationsFlows(ctx).CreateFlowReq(createFlowReq).Execute()

Creates an automation and its initial DRAFT version in one call.



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
	createFlowReq := *openapiclient.NewCreateFlowReq() // CreateFlowReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomationsFlows(context.Background()).CreateFlowReq(createFlowReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationsFlows`: PopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomationsFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsFlowsRequest struct via the builder pattern


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


## PostAutomationsFlowsByIdDisable

> Flow PostAutomationsFlowsByIdDisable(ctx, id).Execute()

Disarms a flow's trigger and marks it DISABLED.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomationsFlowsByIdDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsFlowsByIdDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationsFlowsByIdDisable`: Flow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomationsFlowsByIdDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsFlowsByIdDisableRequest struct via the builder pattern


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


## PostAutomationsFlowsByIdEnable

> Flow PostAutomationsFlowsByIdEnable(ctx, id).Execute()

Arms a flow's trigger and marks it ENABLED.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomationsFlowsByIdEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsFlowsByIdEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationsFlowsByIdEnable`: Flow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomationsFlowsByIdEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsFlowsByIdEnableRequest struct via the builder pattern


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


## PostAutomationsFlowsByIdOperations

> PostAutomationsFlowsByIdOperations(ctx, id).Execute()

Edit a flow — rename it, retarget its trigger, or add, move and delete steps



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
	r, err := apiClient.AutomationsAPI.PostAutomationsFlowsByIdOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsFlowsByIdOperations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAutomationsFlowsByIdOperationsRequest struct via the builder pattern


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


## PostAutomationsFlowsByIdRun

> FlowRun PostAutomationsFlowsByIdRun(ctx, id).Execute()

Starts one durable run of a flow now.



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
	id := "flow_1" // string | ID is the flow to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomationsFlowsByIdRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsFlowsByIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationsFlowsByIdRun`: FlowRun
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomationsFlowsByIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsFlowsByIdRunRequest struct via the builder pattern


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


## PostAutomationsFlowsByIdVersions

> FlowVersion PostAutomationsFlowsByIdVersions(ctx, id).CreateVersionIn(createVersionIn).Execute()

Adds a new DRAFT version to a flow.



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
	id := "flow_1" // string | ID is the flow to add a version to, from the path.
	createVersionIn := *openapiclient.NewCreateVersionIn() // CreateVersionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomationsFlowsByIdVersions(context.Background(), id).CreateVersionIn(createVersionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsFlowsByIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationsFlowsByIdVersions`: FlowVersion
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomationsFlowsByIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to add a version to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsFlowsByIdVersionsRequest struct via the builder pattern


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


## PostAutomationsHooksBySourceByEvent

> PostAutomationsHooksBySourceByEvent(ctx, source, event).Execute()

Fire an event that starts every enabled flow subscribed to it



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
	source := "source_example" // string | 
	event := "event_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationsAPI.PostAutomationsHooksBySourceByEvent(context.Background(), source, event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsHooksBySourceByEvent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAutomationsHooksBySourceByEventRequest struct via the builder pattern


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


## PostAutomationsRunsByIdResume

> PostAutomationsRunsByIdResume(ctx, id).Execute()

Release a run waiting at an approval step, with the approval payload



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
	r, err := apiClient.AutomationsAPI.PostAutomationsRunsByIdResume(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsRunsByIdResume``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAutomationsRunsByIdResumeRequest struct via the builder pattern


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

