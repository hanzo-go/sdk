# \FleetAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudCancelFleetJob**](FleetAPI.md#CloudCancelFleetJob) | **Post** /v1/fleet/jobs/{id}/cancel | Cancels a queued or running render in the caller&#39;s org.
[**CloudListFleet**](FleetAPI.md#CloudListFleet) | **Get** /v1/fleet | Returns every compute unit the caller&#39;s org has, from every source, each carrying its latest utilization: agent run-targets, the BYO machines that dialed in, attached BYO clusters and Visor-provisioned machines.
[**CloudListFleetJobs**](FleetAPI.md#CloudListFleetJobs) | **Get** /v1/fleet/jobs | Returns the caller org&#39;s gpu-jobs render queue, each row tagged with the GPU it targets (empty &#x3D; the shared any-GPU lane) and the node claiming it, optionally narrowed to one GPU&#39;s queue and/or one status.
[**CloudListFleetSamples**](FleetAPI.md#CloudListFleetSamples) | **Get** /v1/fleet/samples | Returns the caller org&#39;s utilization series, oldest first.
[**CloudListFleetWorkers**](FleetAPI.md#CloudListFleetWorkers) | **Get** /v1/fleet/workers | Returns the caller org&#39;s BYO machines — the ones that dialed in via &#x60;hanzo link&#x60; — with everything each host reported about itself.
[**CloudRecordFleetSample**](FleetAPI.md#CloudRecordFleetSample) | **Post** /v1/fleet/samples | Records a BYO worker&#39;s live GPU utilization into the SAME series the fleet board overlays.



## CloudCancelFleetJob

> CloudJobCanceled CloudCancelFleetJob(ctx, id).CloudJobCancel(cloudJobCancel).Execute()

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
	cloudJobCancel := *openapiclient.NewCloudJobCancel() // CloudJobCancel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CloudCancelFleetJob(context.Background(), id).CloudJobCancel(cloudJobCancel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CloudCancelFleetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudCancelFleetJob`: CloudJobCanceled
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CloudCancelFleetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job (activity) id, from the URL path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudCancelFleetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudJobCancel** | [**CloudJobCancel**](CloudJobCancel.md) |  | 

### Return type

[**CloudJobCanceled**](CloudJobCanceled.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListFleet

> CloudFleetBoard CloudListFleet(ctx).Execute()

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
	resp, r, err := apiClient.FleetAPI.CloudListFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CloudListFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListFleet`: CloudFleetBoard
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CloudListFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListFleetRequest struct via the builder pattern


### Return type

[**CloudFleetBoard**](CloudFleetBoard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListFleetJobs

> CloudJobList CloudListFleetJobs(ctx).Gpu(gpu).Status(status).Execute()

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
	resp, r, err := apiClient.FleetAPI.CloudListFleetJobs(context.Background()).Gpu(gpu).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CloudListFleetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListFleetJobs`: CloudJobList
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CloudListFleetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudListFleetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpu** | **string** | GPU selects one node&#39;s lane: jobs TARGETED at it (gpu:&lt;node&gt;) or CLAIMED by it. The literal \&quot;shared\&quot; selects the any-GPU lane — no target, no claimant. Matched case-insensitively. | 
 **status** | **string** | Status selects one lifecycle state: queued, running, stalled, completed, failed or canceled. | 

### Return type

[**CloudJobList**](CloudJobList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListFleetSamples

> CloudSampleList CloudListFleetSamples(ctx).Unit(unit).Source(source).Range_(range_).Execute()

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
	resp, r, err := apiClient.FleetAPI.CloudListFleetSamples(context.Background()).Unit(unit).Source(source).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CloudListFleetSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListFleetSamples`: CloudSampleList
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CloudListFleetSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudListFleetSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unit** | **string** | Unit selects one compute unit&#39;s series by its source-local id. | 
 **source** | **string** | Source selects one plane: \&quot;agent\&quot;, \&quot;byo\&quot; or \&quot;visor\&quot;. | 
 **range_** | **string** | Range is the lookback window (e.g. \&quot;1h\&quot;, \&quot;24h\&quot;, \&quot;7d\&quot;); empty takes the warehouse default. | 

### Return type

[**CloudSampleList**](CloudSampleList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListFleetWorkers

> CloudWorkerList CloudListFleetWorkers(ctx).Execute()

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
	resp, r, err := apiClient.FleetAPI.CloudListFleetWorkers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CloudListFleetWorkers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListFleetWorkers`: CloudWorkerList
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CloudListFleetWorkers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListFleetWorkersRequest struct via the builder pattern


### Return type

[**CloudWorkerList**](CloudWorkerList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudRecordFleetSample

> CloudSampleAccepted CloudRecordFleetSample(ctx).CloudSampleIngest(cloudSampleIngest).Execute()

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
	cloudSampleIngest := *openapiclient.NewCloudSampleIngest() // CloudSampleIngest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CloudRecordFleetSample(context.Background()).CloudSampleIngest(cloudSampleIngest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CloudRecordFleetSample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudRecordFleetSample`: CloudSampleAccepted
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CloudRecordFleetSample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudRecordFleetSampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSampleIngest** | [**CloudSampleIngest**](CloudSampleIngest.md) |  | 

### Return type

[**CloudSampleAccepted**](CloudSampleAccepted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

