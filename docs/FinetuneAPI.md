# \FinetuneAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinetuneHfDatasets**](FinetuneAPI.md#GetFinetuneHfDatasets) | **Get** /v1/finetune/hf/datasets | Proxies a HuggingFace dataset search (dataset picker).
[**GetFinetuneHfModels**](FinetuneAPI.md#GetFinetuneHfModels) | **Get** /v1/finetune/hf/models | Proxies a HuggingFace model search (base-model picker).
[**GetFinetuneHfRepo**](FinetuneAPI.md#GetFinetuneHfRepo) | **Get** /v1/finetune/hf/repo | Returns a repo&#39;s detail (files, gated/private state).
[**GetFinetuneJob**](FinetuneAPI.md#GetFinetuneJob) | **Get** /v1/finetune/job | Returns one job with refreshed live status.
[**GetFinetuneJobs**](FinetuneAPI.md#GetFinetuneJobs) | **Get** /v1/finetune/jobs | Returns the org&#39;s jobs, refreshing live status for active ones.
[**GetFinetunePresets**](FinetuneAPI.md#GetFinetunePresets) | **Get** /v1/finetune/presets | Returns the new-job catalog plus, when a selection is passed (?baseModel&amp;method&amp;task&amp;preset[&amp;datasetExamples]), the recommended config so the console can render \&quot;Recommended\&quot; as a one-click, ready-to-run default.
[**PostFinetuneCancel**](FinetuneAPI.md#PostFinetuneCancel) | **Post** /v1/finetune/cancel | Deletes the TrainJob CR, meters the GPU-hours used so far, and marks the job cancelled.
[**PostFinetuneDeploy**](FinetuneAPI.md#PostFinetuneDeploy) | **Post** /v1/finetune/deploy | Serves a completed job&#39;s checkpoints and registers the result as a routable model on api.hanzo.ai.
[**PostFinetuneJobs**](FinetuneAPI.md#PostFinetuneJobs) | **Post** /v1/finetune/jobs | Validates the request, resolves efficient defaults, persists the job, and submits a real TrainJob CR.



## GetFinetuneHfDatasets

> GetFinetuneHfDatasets(ctx).Execute()

Proxies a HuggingFace dataset search (dataset picker).



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
	r, err := apiClient.FinetuneAPI.GetFinetuneHfDatasets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.GetFinetuneHfDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinetuneHfDatasetsRequest struct via the builder pattern


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


## GetFinetuneHfModels

> GetFinetuneHfModels(ctx).Execute()

Proxies a HuggingFace model search (base-model picker).



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
	r, err := apiClient.FinetuneAPI.GetFinetuneHfModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.GetFinetuneHfModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinetuneHfModelsRequest struct via the builder pattern


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


## GetFinetuneHfRepo

> GetFinetuneHfRepo(ctx).Execute()

Returns a repo's detail (files, gated/private state).



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
	r, err := apiClient.FinetuneAPI.GetFinetuneHfRepo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.GetFinetuneHfRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinetuneHfRepoRequest struct via the builder pattern


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


## GetFinetuneJob

> GetFinetuneJob(ctx).Execute()

Returns one job with refreshed live status.



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
	r, err := apiClient.FinetuneAPI.GetFinetuneJob(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.GetFinetuneJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinetuneJobRequest struct via the builder pattern


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


## GetFinetuneJobs

> GetFinetuneJobs(ctx).Execute()

Returns the org's jobs, refreshing live status for active ones.



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
	r, err := apiClient.FinetuneAPI.GetFinetuneJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.GetFinetuneJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinetuneJobsRequest struct via the builder pattern


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


## GetFinetunePresets

> GetFinetunePresets(ctx).Execute()

Returns the new-job catalog plus, when a selection is passed (?baseModel&method&task&preset[&datasetExamples]), the recommended config so the console can render \"Recommended\" as a one-click, ready-to-run default.



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
	r, err := apiClient.FinetuneAPI.GetFinetunePresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.GetFinetunePresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinetunePresetsRequest struct via the builder pattern


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


## PostFinetuneCancel

> PostFinetuneCancel(ctx).Execute()

Deletes the TrainJob CR, meters the GPU-hours used so far, and marks the job cancelled.



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
	r, err := apiClient.FinetuneAPI.PostFinetuneCancel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.PostFinetuneCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinetuneCancelRequest struct via the builder pattern


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


## PostFinetuneDeploy

> PostFinetuneDeploy(ctx).Execute()

Serves a completed job's checkpoints and registers the result as a routable model on api.hanzo.ai.



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
	r, err := apiClient.FinetuneAPI.PostFinetuneDeploy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.PostFinetuneDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinetuneDeployRequest struct via the builder pattern


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


## PostFinetuneJobs

> PostFinetuneJobs(ctx).Execute()

Validates the request, resolves efficient defaults, persists the job, and submits a real TrainJob CR.



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
	r, err := apiClient.FinetuneAPI.PostFinetuneJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinetuneAPI.PostFinetuneJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinetuneJobsRequest struct via the builder pattern


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

