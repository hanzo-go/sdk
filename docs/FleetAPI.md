# \FleetAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFleetJob**](FleetAPI.md#CancelFleetJob) | **Post** /v1/fleet/jobs/{id}/cancel | Cancels a queued or running render in the caller&#39;s org.
[**ListFleet**](FleetAPI.md#ListFleet) | **Get** /v1/fleet | Returns every compute unit the caller&#39;s org has, from every source, each carrying its latest utilization: agent run-targets, the BYO machines that dialed in, attached BYO clusters and Visor-provisioned machines.
[**ListFleetJobs**](FleetAPI.md#ListFleetJobs) | **Get** /v1/fleet/jobs | Returns the caller org&#39;s gpu-jobs render queue, each row tagged with the GPU it targets (empty &#x3D; the shared any-GPU lane) and the node claiming it, optionally narrowed to one GPU&#39;s queue and/or one status.
[**ListFleetSamples**](FleetAPI.md#ListFleetSamples) | **Get** /v1/fleet/samples | Returns the caller org&#39;s utilization series, oldest first.
[**ListFleetWorkers**](FleetAPI.md#ListFleetWorkers) | **Get** /v1/fleet/workers | Returns the caller org&#39;s BYO machines — the ones that dialed in via &#x60;hanzo link&#x60; — with everything each host reported about itself.
[**RecordFleetSample**](FleetAPI.md#RecordFleetSample) | **Post** /v1/fleet/samples | Records a BYO worker&#39;s live GPU utilization into the SAME series the fleet board overlays.



## CancelFleetJob

> JobCanceled CancelFleetJob(ctx, id).JobCancel(jobCancel).Execute()

Cancels a queued or running render in the caller's org.



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
	id := "id_example" // string | ID is the job (activity) id, from the URL path.
	jobCancel := *openapiclient.NewJobCancel() // JobCancel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CancelFleetJob(context.Background(), id).JobCancel(jobCancel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CancelFleetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelFleetJob`: JobCanceled
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CancelFleetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job (activity) id, from the URL path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelFleetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobCancel** | [**JobCancel**](JobCancel.md) |  | 

### Return type

[**JobCanceled**](JobCanceled.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleet

> FleetBoard ListFleet(ctx).Execute()

Returns every compute unit the caller's org has, from every source, each carrying its latest utilization: agent run-targets, the BYO machines that dialed in, attached BYO clusters and Visor-provisioned machines.



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
	resp, r, err := apiClient.FleetAPI.ListFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.ListFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleet`: FleetBoard
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.ListFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFleetRequest struct via the builder pattern


### Return type

[**FleetBoard**](FleetBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleetJobs

> JobList ListFleetJobs(ctx).Gpu(gpu).Status(status).Execute()

Returns the caller org's gpu-jobs render queue, each row tagged with the GPU it targets (empty = the shared any-GPU lane) and the node claiming it, optionally narrowed to one GPU's queue and/or one status.



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
	gpu := "gpu_example" // string | GPU selects one node's lane: jobs TARGETED at it (gpu:<node>) or CLAIMED by it. The literal \"shared\" selects the any-GPU lane — no target, no claimant. Matched case-insensitively. (optional)
	status := "status_example" // string | Status selects one lifecycle state: queued, running, stalled, completed, failed or canceled. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.ListFleetJobs(context.Background()).Gpu(gpu).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.ListFleetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleetJobs`: JobList
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.ListFleetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFleetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpu** | **string** | GPU selects one node&#39;s lane: jobs TARGETED at it (gpu:&lt;node&gt;) or CLAIMED by it. The literal \&quot;shared\&quot; selects the any-GPU lane — no target, no claimant. Matched case-insensitively. | 
 **status** | **string** | Status selects one lifecycle state: queued, running, stalled, completed, failed or canceled. | 

### Return type

[**JobList**](JobList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleetSamples

> SampleList ListFleetSamples(ctx).Unit(unit).Source(source).Range_(range_).Execute()

Returns the caller org's utilization series, oldest first.



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
	unit := "unit_example" // string | Unit selects one compute unit's series by its source-local id. (optional)
	source := "source_example" // string | Source selects one plane: \"agent\", \"byo\" or \"visor\". (optional)
	range_ := "range__example" // string | Range is the lookback window (e.g. \"1h\", \"24h\", \"7d\"); empty takes the warehouse default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.ListFleetSamples(context.Background()).Unit(unit).Source(source).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.ListFleetSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleetSamples`: SampleList
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.ListFleetSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFleetSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unit** | **string** | Unit selects one compute unit&#39;s series by its source-local id. | 
 **source** | **string** | Source selects one plane: \&quot;agent\&quot;, \&quot;byo\&quot; or \&quot;visor\&quot;. | 
 **range_** | **string** | Range is the lookback window (e.g. \&quot;1h\&quot;, \&quot;24h\&quot;, \&quot;7d\&quot;); empty takes the warehouse default. | 

### Return type

[**SampleList**](SampleList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleetWorkers

> WorkerList ListFleetWorkers(ctx).Execute()

Returns the caller org's BYO machines — the ones that dialed in via `hanzo link` — with everything each host reported about itself.



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
	resp, r, err := apiClient.FleetAPI.ListFleetWorkers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.ListFleetWorkers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleetWorkers`: WorkerList
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.ListFleetWorkers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFleetWorkersRequest struct via the builder pattern


### Return type

[**WorkerList**](WorkerList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordFleetSample

> SampleAccepted RecordFleetSample(ctx).SampleIngest(sampleIngest).Execute()

Records a BYO worker's live GPU utilization into the SAME series the fleet board overlays.



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
	sampleIngest := *openapiclient.NewSampleIngest() // SampleIngest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.RecordFleetSample(context.Background()).SampleIngest(sampleIngest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.RecordFleetSample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordFleetSample`: SampleAccepted
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.RecordFleetSample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordFleetSampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sampleIngest** | [**SampleIngest**](SampleIngest.md) |  | 

### Return type

[**SampleAccepted**](SampleAccepted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

