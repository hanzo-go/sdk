# \AutomationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1AutomationsFlowsId**](AutomationsAPI.md#CloudDeleteV1AutomationsFlowsId) | **Delete** /v1/automations/flows/{id} | DeleteFlow deletes one automation, its versions and its run history.
[**CloudGetV1AutomationsConnectors**](AutomationsAPI.md#CloudGetV1AutomationsConnectors) | **Get** /v1/automations/connectors | Connectors returns the connector catalogue.
[**CloudGetV1AutomationsFlows**](AutomationsAPI.md#CloudGetV1AutomationsFlows) | **Get** /v1/automations/flows | ListFlows returns the caller org&#39;s automations, most-recently-updated first.
[**CloudGetV1AutomationsFlowsId**](AutomationsAPI.md#CloudGetV1AutomationsFlowsId) | **Get** /v1/automations/flows/{id} | GetFlow returns one automation and its latest version.
[**CloudGetV1AutomationsFlowsIdVersions**](AutomationsAPI.md#CloudGetV1AutomationsFlowsIdVersions) | **Get** /v1/automations/flows/{id}/versions | ListVersions returns one flow&#39;s versions, newest first.
[**CloudGetV1AutomationsPieces**](AutomationsAPI.md#CloudGetV1AutomationsPieces) | **Get** /v1/automations/pieces | Pieces is the retired-name alias of the connector catalogue.
[**CloudGetV1AutomationsRuns**](AutomationsAPI.md#CloudGetV1AutomationsRuns) | **Get** /v1/automations/runs | ListRuns returns the caller org&#39;s run history, newest first.
[**CloudGetV1AutomationsRunsId**](AutomationsAPI.md#CloudGetV1AutomationsRunsId) | **Get** /v1/automations/runs/{id} | GetRun returns one run.
[**CloudPatchV1AutomationsFlowsId**](AutomationsAPI.md#CloudPatchV1AutomationsFlowsId) | **Patch** /v1/automations/flows/{id} | UpdateFlow updates one automation&#39;s metadata in place.
[**CloudPostV1AutomationsConnectorsIdRun**](AutomationsAPI.md#CloudPostV1AutomationsConnectorsIdRun) | **Post** /v1/automations/connectors/{id}/run | Run executes one connector action in-process and answers the outcome.
[**CloudPostV1AutomationsFlows**](AutomationsAPI.md#CloudPostV1AutomationsFlows) | **Post** /v1/automations/flows | CreateFlow creates an automation and its initial DRAFT version in one call.
[**CloudPostV1AutomationsFlowsByIdOperations**](AutomationsAPI.md#CloudPostV1AutomationsFlowsByIdOperations) | **Post** /v1/automations/flows/{id}/operations | 
[**CloudPostV1AutomationsFlowsIdDisable**](AutomationsAPI.md#CloudPostV1AutomationsFlowsIdDisable) | **Post** /v1/automations/flows/{id}/disable | DisableFlow disarms a flow&#39;s trigger and marks it DISABLED.
[**CloudPostV1AutomationsFlowsIdEnable**](AutomationsAPI.md#CloudPostV1AutomationsFlowsIdEnable) | **Post** /v1/automations/flows/{id}/enable | EnableFlow arms a flow&#39;s trigger and marks it ENABLED.
[**CloudPostV1AutomationsFlowsIdRun**](AutomationsAPI.md#CloudPostV1AutomationsFlowsIdRun) | **Post** /v1/automations/flows/{id}/run | RunFlow starts one durable run of a flow now.
[**CloudPostV1AutomationsFlowsIdVersions**](AutomationsAPI.md#CloudPostV1AutomationsFlowsIdVersions) | **Post** /v1/automations/flows/{id}/versions | CreateVersion adds a new DRAFT version to a flow.
[**CloudPostV1AutomationsHooksBySourceByEvent**](AutomationsAPI.md#CloudPostV1AutomationsHooksBySourceByEvent) | **Post** /v1/automations/hooks/{source}/{event} | 
[**CloudPostV1AutomationsRunsByIdResume**](AutomationsAPI.md#CloudPostV1AutomationsRunsByIdResume) | **Post** /v1/automations/runs/{id}/resume | 



## CloudDeleteV1AutomationsFlowsId

> CloudDeleteV1AutomationsFlowsId(ctx, id).Execute()

DeleteFlow deletes one automation, its versions and its run history.



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
	r, err := apiClient.AutomationsAPI.CloudDeleteV1AutomationsFlowsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudDeleteV1AutomationsFlowsId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1AutomationsFlowsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutomationsConnectors

> CloudCatalog CloudGetV1AutomationsConnectors(ctx).Execute()

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
	resp, r, err := apiClient.AutomationsAPI.CloudGetV1AutomationsConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudGetV1AutomationsConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutomationsConnectors`: CloudCatalog
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudGetV1AutomationsConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutomationsConnectorsRequest struct via the builder pattern


### Return type

[**CloudCatalog**](CloudCatalog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutomationsFlows

> CloudFlowPage CloudGetV1AutomationsFlows(ctx).Limit(limit).Execute()

ListFlows returns the caller org's automations, most-recently-updated first.



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
	resp, r, err := apiClient.AutomationsAPI.CloudGetV1AutomationsFlows(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudGetV1AutomationsFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutomationsFlows`: CloudFlowPage
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudGetV1AutomationsFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutomationsFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**CloudFlowPage**](CloudFlowPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutomationsFlowsId

> CloudPopulatedFlow CloudGetV1AutomationsFlowsId(ctx, id).Execute()

GetFlow returns one automation and its latest version.



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
	resp, r, err := apiClient.AutomationsAPI.CloudGetV1AutomationsFlowsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudGetV1AutomationsFlowsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutomationsFlowsId`: CloudPopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudGetV1AutomationsFlowsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutomationsFlowsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudPopulatedFlow**](CloudPopulatedFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutomationsFlowsIdVersions

> CloudVersionPage CloudGetV1AutomationsFlowsIdVersions(ctx, id).Limit(limit).Execute()

ListVersions returns one flow's versions, newest first.



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
	resp, r, err := apiClient.AutomationsAPI.CloudGetV1AutomationsFlowsIdVersions(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudGetV1AutomationsFlowsIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutomationsFlowsIdVersions`: CloudVersionPage
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudGetV1AutomationsFlowsIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow whose versions to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutomationsFlowsIdVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**CloudVersionPage**](CloudVersionPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutomationsPieces

> CloudCatalog CloudGetV1AutomationsPieces(ctx).Execute()

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
	resp, r, err := apiClient.AutomationsAPI.CloudGetV1AutomationsPieces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudGetV1AutomationsPieces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutomationsPieces`: CloudCatalog
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudGetV1AutomationsPieces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutomationsPiecesRequest struct via the builder pattern


### Return type

[**CloudCatalog**](CloudCatalog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutomationsRuns

> CloudRunPage CloudGetV1AutomationsRuns(ctx).FlowId(flowId).Limit(limit).Execute()

ListRuns returns the caller org's run history, newest first.



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
	resp, r, err := apiClient.AutomationsAPI.CloudGetV1AutomationsRuns(context.Background()).FlowId(flowId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudGetV1AutomationsRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutomationsRuns`: CloudRunPage
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudGetV1AutomationsRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutomationsRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** | FlowID narrows the history to one flow. Omit it for the whole org&#39;s runs. | 
 **limit** | **int32** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**CloudRunPage**](CloudRunPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AutomationsRunsId

> CloudFlowRun CloudGetV1AutomationsRunsId(ctx, id).Execute()

GetRun returns one run.



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
	resp, r, err := apiClient.AutomationsAPI.CloudGetV1AutomationsRunsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudGetV1AutomationsRunsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AutomationsRunsId`: CloudFlowRun
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudGetV1AutomationsRunsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the run to read, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AutomationsRunsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudFlowRun**](CloudFlowRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1AutomationsFlowsId

> CloudFlow CloudPatchV1AutomationsFlowsId(ctx, id).CloudPatchFlowIn(cloudPatchFlowIn).Execute()

UpdateFlow updates one automation's metadata in place.



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
	cloudPatchFlowIn := *openapiclient.NewCloudPatchFlowIn() // CloudPatchFlowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.CloudPatchV1AutomationsFlowsId(context.Background(), id).CloudPatchFlowIn(cloudPatchFlowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPatchV1AutomationsFlowsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1AutomationsFlowsId`: CloudFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudPatchV1AutomationsFlowsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1AutomationsFlowsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPatchFlowIn** | [**CloudPatchFlowIn**](CloudPatchFlowIn.md) |  | 

### Return type

[**CloudFlow**](CloudFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsConnectorsIdRun

> CloudRunResp CloudPostV1AutomationsConnectorsIdRun(ctx, id).CloudRunIn(cloudRunIn).Execute()

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
	cloudRunIn := *openapiclient.NewCloudRunIn() // CloudRunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsConnectorsIdRun(context.Background(), id).CloudRunIn(cloudRunIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsConnectorsIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutomationsConnectorsIdRun`: CloudRunResp
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudPostV1AutomationsConnectorsIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector to run, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsConnectorsIdRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudRunIn** | [**CloudRunIn**](CloudRunIn.md) |  | 

### Return type

[**CloudRunResp**](CloudRunResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsFlows

> CloudPopulatedFlow CloudPostV1AutomationsFlows(ctx).CloudCreateFlowReq(cloudCreateFlowReq).Execute()

CreateFlow creates an automation and its initial DRAFT version in one call.



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
	cloudCreateFlowReq := *openapiclient.NewCloudCreateFlowReq() // CloudCreateFlowReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsFlows(context.Background()).CloudCreateFlowReq(cloudCreateFlowReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutomationsFlows`: CloudPopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudPostV1AutomationsFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateFlowReq** | [**CloudCreateFlowReq**](CloudCreateFlowReq.md) |  | 

### Return type

[**CloudPopulatedFlow**](CloudPopulatedFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsFlowsByIdOperations

> CloudPostV1AutomationsFlowsByIdOperations(ctx, id).Execute()



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
	r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsFlowsByIdOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsFlowsByIdOperations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsFlowsByIdOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsFlowsIdDisable

> CloudFlow CloudPostV1AutomationsFlowsIdDisable(ctx, id).Execute()

DisableFlow disarms a flow's trigger and marks it DISABLED.



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
	resp, r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsFlowsIdDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsFlowsIdDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutomationsFlowsIdDisable`: CloudFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudPostV1AutomationsFlowsIdDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsFlowsIdDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudFlow**](CloudFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsFlowsIdEnable

> CloudFlow CloudPostV1AutomationsFlowsIdEnable(ctx, id).Execute()

EnableFlow arms a flow's trigger and marks it ENABLED.



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
	resp, r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsFlowsIdEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsFlowsIdEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutomationsFlowsIdEnable`: CloudFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudPostV1AutomationsFlowsIdEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsFlowsIdEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudFlow**](CloudFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsFlowsIdRun

> CloudFlowRun CloudPostV1AutomationsFlowsIdRun(ctx, id).Execute()

RunFlow starts one durable run of a flow now.



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
	resp, r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsFlowsIdRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsFlowsIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutomationsFlowsIdRun`: CloudFlowRun
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudPostV1AutomationsFlowsIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsFlowsIdRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudFlowRun**](CloudFlowRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsFlowsIdVersions

> CloudFlowVersion CloudPostV1AutomationsFlowsIdVersions(ctx, id).CloudCreateVersionIn(cloudCreateVersionIn).Execute()

CreateVersion adds a new DRAFT version to a flow.



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
	cloudCreateVersionIn := *openapiclient.NewCloudCreateVersionIn() // CloudCreateVersionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsFlowsIdVersions(context.Background(), id).CloudCreateVersionIn(cloudCreateVersionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsFlowsIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AutomationsFlowsIdVersions`: CloudFlowVersion
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.CloudPostV1AutomationsFlowsIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the flow to add a version to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsFlowsIdVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCreateVersionIn** | [**CloudCreateVersionIn**](CloudCreateVersionIn.md) |  | 

### Return type

[**CloudFlowVersion**](CloudFlowVersion.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsHooksBySourceByEvent

> CloudPostV1AutomationsHooksBySourceByEvent(ctx, source, event).Execute()



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
	r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsHooksBySourceByEvent(context.Background(), source, event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsHooksBySourceByEvent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsHooksBySourceByEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AutomationsRunsByIdResume

> CloudPostV1AutomationsRunsByIdResume(ctx, id).Execute()



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
	r, err := apiClient.AutomationsAPI.CloudPostV1AutomationsRunsByIdResume(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.CloudPostV1AutomationsRunsByIdResume``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AutomationsRunsByIdResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

