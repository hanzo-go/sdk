# \AutoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutoAutomationsById**](AutoAPI.md#DeleteAutoAutomationsById) | **Delete** /v1/auto/automations/{id} | Deletes an automation, its schedule and its run history, and stops a Dev run it has going.
[**DeleteAutoFlowsById**](AutoAPI.md#DeleteAutoFlowsById) | **Delete** /v1/auto/flows/{id} | Deletes one automation, its versions and its run history.
[**GetAutoAutomations**](AutoAPI.md#GetAutoAutomations) | **Get** /v1/auto/automations | Returns your automations — another member&#39;s are never listed.
[**GetAutoAutomationsById**](AutoAPI.md#GetAutoAutomationsById) | **Get** /v1/auto/automations/{id} | Returns one automation.
[**GetAutoAutomationsByIdRuns**](AutoAPI.md#GetAutoAutomationsByIdRuns) | **Get** /v1/auto/automations/{id}/runs | Returns one automation&#39;s runs, newest first, to the person it runs as.
[**GetAutoEvents**](AutoAPI.md#GetAutoEvents) | **Get** /v1/auto/events | Stream your jobs and automations as they change, as Server-Sent Events
[**GetAutoFlows**](AutoAPI.md#GetAutoFlows) | **Get** /v1/auto/flows | Returns the caller org&#39;s automations, most-recently-updated first.
[**GetAutoFlowsById**](AutoAPI.md#GetAutoFlowsById) | **Get** /v1/auto/flows/{id} | Returns one automation and its latest version.
[**GetAutoFlowsByIdVersions**](AutoAPI.md#GetAutoFlowsByIdVersions) | **Get** /v1/auto/flows/{id}/versions | Returns one flow&#39;s versions, newest first.
[**GetAutoJobs**](AutoAPI.md#GetAutoJobs) | **Get** /v1/auto/jobs | Lists everything of yours that runs without you watching, across every product: each run of your automations and workflows, and every Studio render, long video or other background run a product recorded for you.
[**GetAutoJobsById**](AutoAPI.md#GetAutoJobsById) | **Get** /v1/auto/jobs/{id} | Reads one of your jobs.
[**GetAutoProvider**](AutoAPI.md#GetAutoProvider) | **Get** /v1/auto/provider | Returns the connector catalogue.
[**GetAutoRuns**](AutoAPI.md#GetAutoRuns) | **Get** /v1/auto/runs | Returns your run history, newest first: the runs of the flows you hold.
[**GetAutoRunsById**](AutoAPI.md#GetAutoRunsById) | **Get** /v1/auto/runs/{id} | Returns one run.
[**GetAutoTemplates**](AutoAPI.md#GetAutoTemplates) | **Get** /v1/auto/templates | Returns the starter automations: a name, what it does, the instructions it runs and the schedule it suggests, in the caller&#39;s own zone.
[**PatchAutoAutomationsById**](AutoAPI.md#PatchAutoAutomationsById) | **Patch** /v1/auto/automations/{id} | Changes an automation, for the person it runs as: any field it was created with, and &#x60;enabled&#x60;, which arms or disarms its schedule.
[**PatchAutoFlowsById**](AutoAPI.md#PatchAutoFlowsById) | **Patch** /v1/auto/flows/{id} | Updates one automation&#39;s metadata in place.
[**PostAutoAutomations**](AutoAPI.md#PostAutoAutomations) | **Post** /v1/auto/automations | Creates an automation and arms its schedule.
[**PostAutoAutomationsByIdRun**](AutoAPI.md#PostAutoAutomationsByIdRun) | **Post** /v1/auto/automations/{id}/run | Starts one run now, whether or not its schedule is armed.
[**PostAutoAutomationsByIdRunsByRunReview**](AutoAPI.md#PostAutoAutomationsByIdRunsByRunReview) | **Post** /v1/auto/automations/{id}/runs/{run}/review | Answers a run that holds a post for review, for the person the automation runs as: &#x60;post&#x60; true posts it to the automation&#39;s accounts and answers the run while it posts; false discards it.
[**PostAutoCron**](AutoAPI.md#PostAutoCron) | **Post** /v1/auto/cron | Reads a schedule said in words — \&quot;every 5 minutes\&quot;, \&quot;weekdays at 8:30\&quot;, \&quot;every monday at 9am\&quot;, \&quot;tomorrow at 7pm\&quot;, \&quot;in 20 minutes\&quot;, or a five-field cron expression — in your time zone, and answers the schedule an automation takes, said back in plain words, with when it would first run.
[**PostAutoFlows**](AutoAPI.md#PostAutoFlows) | **Post** /v1/auto/flows | Creates an automation and its initial DRAFT version in one call.
[**PostAutoFlowsByIdDisable**](AutoAPI.md#PostAutoFlowsByIdDisable) | **Post** /v1/auto/flows/{id}/disable | Disarms a flow&#39;s trigger and marks it DISABLED.
[**PostAutoFlowsByIdEnable**](AutoAPI.md#PostAutoFlowsByIdEnable) | **Post** /v1/auto/flows/{id}/enable | Arms a flow&#39;s trigger and marks it ENABLED.
[**PostAutoFlowsByIdOperations**](AutoAPI.md#PostAutoFlowsByIdOperations) | **Post** /v1/auto/flows/{id}/operations | Edit a flow — rename it, retarget its trigger, or add, move and delete steps
[**PostAutoFlowsByIdRun**](AutoAPI.md#PostAutoFlowsByIdRun) | **Post** /v1/auto/flows/{id}/run | Starts one durable run of a flow now.
[**PostAutoFlowsByIdVersions**](AutoAPI.md#PostAutoFlowsByIdVersions) | **Post** /v1/auto/flows/{id}/versions | Adds a new DRAFT version to a flow.
[**PostAutoHooksBySourceByEvent**](AutoAPI.md#PostAutoHooksBySourceByEvent) | **Post** /v1/auto/hooks/{source}/{event} | Fire an event that starts every enabled flow subscribed to it
[**PostAutoProviderByIdRun**](AutoAPI.md#PostAutoProviderByIdRun) | **Post** /v1/auto/provider/{id}/run | Executes one provider action in-process and answers the outcome.
[**PostAutoRunsByIdResume**](AutoAPI.md#PostAutoRunsByIdResume) | **Post** /v1/auto/runs/{id}/resume | Release a run waiting at an approval step, with the approval payload



## DeleteAutoAutomationsById

> DeleteAutoAutomationsById(ctx, id).Execute()

Deletes an automation, its schedule and its run history, and stops a Dev run it has going.



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
	id := "flow_1" // string | ID is the automation, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoAPI.DeleteAutoAutomationsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.DeleteAutoAutomationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the automation, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoAutomationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoAutomations

> AutoAutomationPage GetAutoAutomations(ctx).Q(q).Sort(sort).Source(source).Execute()

Returns your automations — another member's are never listed.



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
	q := "q_example" // string | Q keeps the automations whose name or instructions contain it, ignoring case. (optional)
	sort := "sort_example" // string | Sort is name, next or updated (the default, newest first). (optional)
	source := "source_example" // string | Source keeps the automations made in one product: chat, dev, studio or workflow. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoAutomations(context.Background()).Q(q).Sort(sort).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoAutomations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoAutomations`: AutoAutomationPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoAutomations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q keeps the automations whose name or instructions contain it, ignoring case. | 
 **sort** | **string** | Sort is name, next or updated (the default, newest first). | 
 **source** | **string** | Source keeps the automations made in one product: chat, dev, studio or workflow. | 

### Return type

[**AutoAutomationPage**](AutoAutomationPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoAutomationsById

> AutoAutomation GetAutoAutomationsById(ctx, id).Execute()

Returns one automation.



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
	id := "flow_1" // string | ID is the automation, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoAutomationsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoAutomationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoAutomationsById`: AutoAutomation
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoAutomationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the automation, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoAutomationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoAutomation**](AutoAutomation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoAutomationsByIdRuns

> AutoAutomationRunPage GetAutoAutomationsByIdRuns(ctx, id).Limit(limit).Execute()

Returns one automation's runs, newest first, to the person it runs as.



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
	id := "flow_1" // string | ID is the automation, from the path.
	limit := int64(789) // int64 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoAutomationsByIdRuns(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoAutomationsByIdRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoAutomationsByIdRuns`: AutoAutomationRunPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoAutomationsByIdRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the automation, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoAutomationsByIdRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**AutoAutomationRunPage**](AutoAutomationRunPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoEvents

> *os.File GetAutoEvents(ctx).Execute()

Stream your jobs and automations as they change, as Server-Sent Events



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
	resp, r, err := apiClient.AutoAPI.GetAutoEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoEvents`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoEventsRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoFlows

> AutoFlowPage GetAutoFlows(ctx).Limit(limit).Execute()

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
	limit := int64(789) // int64 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoFlows(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoFlows`: AutoFlowPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**AutoFlowPage**](AutoFlowPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoFlowsById

> AutoPopulatedFlow GetAutoFlowsById(ctx, id).Execute()

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
	// response from `GetAutoFlowsById`: AutoPopulatedFlow
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

[**AutoPopulatedFlow**](AutoPopulatedFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoFlowsByIdVersions

> AutoVersionPage GetAutoFlowsByIdVersions(ctx, id).Limit(limit).Execute()

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
	limit := int64(789) // int64 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoFlowsByIdVersions(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoFlowsByIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoFlowsByIdVersions`: AutoVersionPage
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

 **limit** | **int64** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**AutoVersionPage**](AutoVersionPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoJobs

> AutoJobPage GetAutoJobs(ctx).Product(product).State(state).Limit(limit).Execute()

Lists everything of yours that runs without you watching, across every product: each run of your automations and workflows, and every Studio render, long video or other background run a product recorded for you.



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
	product := "studio" // string | Product keeps one product's jobs: chat, dev, studio or workflow. (optional)
	state := "running" // string | State keeps one state: scheduled, running, succeeded, failed or cancelled. (optional)
	limit := int64(789) // int64 | Limit bounds the page (default 100, at most 500). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoJobs(context.Background()).Product(product).State(state).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoJobs`: AutoJobPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Product keeps one product&#39;s jobs: chat, dev, studio or workflow. | 
 **state** | **string** | State keeps one state: scheduled, running, succeeded, failed or cancelled. | 
 **limit** | **int64** | Limit bounds the page (default 100, at most 500). | 

### Return type

[**AutoJobPage**](AutoJobPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoJobsById

> AutoJob GetAutoJobsById(ctx, id).Execute()

Reads one of your jobs.



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
	id := "run_3f7a…" // string | ID is the job, from the path: a run_ or a job_ id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoJobsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoJobsById`: AutoJob
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoJobsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path: a run_ or a job_ id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoJobsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoJob**](AutoJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoProvider

> AutoCatalog GetAutoProvider(ctx).Execute()

Returns the connector catalogue.



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
	resp, r, err := apiClient.AutoAPI.GetAutoProvider(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoProvider`: AutoCatalog
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoProviderRequest struct via the builder pattern


### Return type

[**AutoCatalog**](AutoCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoRuns

> AutoRunPage GetAutoRuns(ctx).FlowId(flowId).Limit(limit).Execute()

Returns your run history, newest first: the runs of the flows you hold.



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
	limit := int64(789) // int64 | Limit bounds the page (default 200, maximum 1000). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.GetAutoRuns(context.Background()).FlowId(flowId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoRuns`: AutoRunPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** | FlowID narrows the history to one flow. Omit it for the whole org&#39;s runs. | 
 **limit** | **int64** | Limit bounds the page (default 200, maximum 1000). | 

### Return type

[**AutoRunPage**](AutoRunPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoRunsById

> AutoFlowRun GetAutoRunsById(ctx, id).Execute()

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
	// response from `GetAutoRunsById`: AutoFlowRun
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

[**AutoFlowRun**](AutoFlowRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTemplates

> AutoStarterPage GetAutoTemplates(ctx).Execute()

Returns the starter automations: a name, what it does, the instructions it runs and the schedule it suggests, in the caller's own zone.



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
	resp, r, err := apiClient.AutoAPI.GetAutoTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.GetAutoTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTemplates`: AutoStarterPage
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.GetAutoTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTemplatesRequest struct via the builder pattern


### Return type

[**AutoStarterPage**](AutoStarterPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutoAutomationsById

> AutoAutomation PatchAutoAutomationsById(ctx, id).AutoAutomationPatch(autoAutomationPatch).Execute()

Changes an automation, for the person it runs as: any field it was created with, and `enabled`, which arms or disarms its schedule.



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
	id := "flow_1" // string | ID is the automation, from the path.
	autoAutomationPatch := *openapiclient.NewAutoAutomationPatch() // AutoAutomationPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PatchAutoAutomationsById(context.Background(), id).AutoAutomationPatch(autoAutomationPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PatchAutoAutomationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutoAutomationsById`: AutoAutomation
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PatchAutoAutomationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the automation, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutoAutomationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoAutomationPatch** | [**AutoAutomationPatch**](AutoAutomationPatch.md) |  | 

### Return type

[**AutoAutomation**](AutoAutomation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutoFlowsById

> AutoFlow PatchAutoFlowsById(ctx, id).AutoPatchFlowIn(autoPatchFlowIn).Execute()

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
	autoPatchFlowIn := *openapiclient.NewAutoPatchFlowIn() // AutoPatchFlowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PatchAutoFlowsById(context.Background(), id).AutoPatchFlowIn(autoPatchFlowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PatchAutoFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutoFlowsById`: AutoFlow
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

 **autoPatchFlowIn** | [**AutoPatchFlowIn**](AutoPatchFlowIn.md) |  | 

### Return type

[**AutoFlow**](AutoFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoAutomations

> AutoAutomation PostAutoAutomations(ctx).AutoAutomationIn(autoAutomationIn).Execute()

Creates an automation and arms its schedule.



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
	autoAutomationIn := *openapiclient.NewAutoAutomationIn() // AutoAutomationIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoAutomations(context.Background()).AutoAutomationIn(autoAutomationIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoAutomations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoAutomations`: AutoAutomation
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoAutomations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoAutomationIn** | [**AutoAutomationIn**](AutoAutomationIn.md) |  | 

### Return type

[**AutoAutomation**](AutoAutomation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoAutomationsByIdRun

> AutoRunStarted PostAutoAutomationsByIdRun(ctx, id).Execute()

Starts one run now, whether or not its schedule is armed.



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
	id := "flow_1" // string | ID is the automation, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoAutomationsByIdRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoAutomationsByIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoAutomationsByIdRun`: AutoRunStarted
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoAutomationsByIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the automation, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoAutomationsByIdRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoRunStarted**](AutoRunStarted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoAutomationsByIdRunsByRunReview

> AutoAutomationRun PostAutoAutomationsByIdRunsByRunReview(ctx, id, run).AutoReviewIn(autoReviewIn).Execute()

Answers a run that holds a post for review, for the person the automation runs as: `post` true posts it to the automation's accounts and answers the run while it posts; false discards it.



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
	id := "flow_1" // string | ID is the automation, and Run its run in review, from the path.
	run := "run_1" // string | 
	autoReviewIn := *openapiclient.NewAutoReviewIn() // AutoReviewIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoAutomationsByIdRunsByRunReview(context.Background(), id, run).AutoReviewIn(autoReviewIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoAutomationsByIdRunsByRunReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoAutomationsByIdRunsByRunReview`: AutoAutomationRun
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoAutomationsByIdRunsByRunReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the automation, and Run its run in review, from the path. | 
**run** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoAutomationsByIdRunsByRunReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoReviewIn** | [**AutoReviewIn**](AutoReviewIn.md) |  | 

### Return type

[**AutoAutomationRun**](AutoAutomationRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoCron

> AutoWhenOut PostAutoCron(ctx).AutoWhenIn(autoWhenIn).Execute()

Reads a schedule said in words — \"every 5 minutes\", \"weekdays at 8:30\", \"every monday at 9am\", \"tomorrow at 7pm\", \"in 20 minutes\", or a five-field cron expression — in your time zone, and answers the schedule an automation takes, said back in plain words, with when it would first run.



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
	autoWhenIn := *openapiclient.NewAutoWhenIn() // AutoWhenIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoCron(context.Background()).AutoWhenIn(autoWhenIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoCron``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoCron`: AutoWhenOut
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoCron`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoCronRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoWhenIn** | [**AutoWhenIn**](AutoWhenIn.md) |  | 

### Return type

[**AutoWhenOut**](AutoWhenOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlows

> AutoPopulatedFlow PostAutoFlows(ctx).AutoCreateFlowReq(autoCreateFlowReq).Execute()

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
	autoCreateFlowReq := *openapiclient.NewAutoCreateFlowReq() // AutoCreateFlowReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlows(context.Background()).AutoCreateFlowReq(autoCreateFlowReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlows`: AutoPopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateFlowReq** | [**AutoCreateFlowReq**](AutoCreateFlowReq.md) |  | 

### Return type

[**AutoPopulatedFlow**](AutoPopulatedFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlowsByIdDisable

> AutoFlow PostAutoFlowsByIdDisable(ctx, id).Execute()

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
	// response from `PostAutoFlowsByIdDisable`: AutoFlow
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

[**AutoFlow**](AutoFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlowsByIdEnable

> AutoFlow PostAutoFlowsByIdEnable(ctx, id).Execute()

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
	// response from `PostAutoFlowsByIdEnable`: AutoFlow
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

[**AutoFlow**](AutoFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

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

> AutoFlowRun PostAutoFlowsByIdRun(ctx, id).Execute()

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
	// response from `PostAutoFlowsByIdRun`: AutoFlowRun
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

[**AutoFlowRun**](AutoFlowRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoFlowsByIdVersions

> AutoFlowVersion PostAutoFlowsByIdVersions(ctx, id).AutoCreateVersionIn(autoCreateVersionIn).Execute()

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
	autoCreateVersionIn := *openapiclient.NewAutoCreateVersionIn() // AutoCreateVersionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoFlowsByIdVersions(context.Background(), id).AutoCreateVersionIn(autoCreateVersionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoFlowsByIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoFlowsByIdVersions`: AutoFlowVersion
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

 **autoCreateVersionIn** | [**AutoCreateVersionIn**](AutoCreateVersionIn.md) |  | 

### Return type

[**AutoFlowVersion**](AutoFlowVersion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

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


## PostAutoProviderByIdRun

> AutoRunResp PostAutoProviderByIdRun(ctx, id).AutoRunIn(autoRunIn).Execute()

Executes one provider action in-process and answers the outcome.



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
	id := "notion" // string | ID is the provider to run, from the path.
	autoRunIn := *openapiclient.NewAutoRunIn() // AutoRunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAPI.PostAutoProviderByIdRun(context.Background(), id).AutoRunIn(autoRunIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAPI.PostAutoProviderByIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutoProviderByIdRun`: AutoRunResp
	fmt.Fprintf(os.Stdout, "Response from `AutoAPI.PostAutoProviderByIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the provider to run, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoProviderByIdRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoRunIn** | [**AutoRunIn**](AutoRunIn.md) |  | 

### Return type

[**AutoRunResp**](AutoRunResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

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

