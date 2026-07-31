# \TrainAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1TrainExperimentsName**](TrainAPI.md#CloudDeleteV1TrainExperimentsName) | **Delete** /v1/train/experiments/{name} | DeleteExperiment deletes a hyperparameter-tuning experiment.
[**CloudDeleteV1TrainJobsName**](TrainAPI.md#CloudDeleteV1TrainJobsName) | **Delete** /v1/train/jobs/{name} | DeleteJob deletes a training job.
[**CloudGetV1TrainExperiments**](TrainAPI.md#CloudGetV1TrainExperiments) | **Get** /v1/train/experiments | ListExperiments lists the caller org&#39;s tuning experiments.
[**CloudGetV1TrainExperimentsName**](TrainAPI.md#CloudGetV1TrainExperimentsName) | **Get** /v1/train/experiments/{name} | GetExperiment returns one hyperparameter-tuning experiment.
[**CloudGetV1TrainExperimentsNameTrials**](TrainAPI.md#CloudGetV1TrainExperimentsNameTrials) | **Get** /v1/train/experiments/{name}/trials | ListTrials lists the katib Trials one experiment owns.
[**CloudGetV1TrainHealth**](TrainAPI.md#CloudGetV1TrainHealth) | **Get** /v1/train/health | 
[**CloudGetV1TrainJobs**](TrainAPI.md#CloudGetV1TrainJobs) | **Get** /v1/train/jobs | ListJobs lists the training jobs in the caller&#39;s org.
[**CloudGetV1TrainJobsName**](TrainAPI.md#CloudGetV1TrainJobsName) | **Get** /v1/train/jobs/{name} | GetJob returns one training job.
[**CloudPostV1TrainExperiments**](TrainAPI.md#CloudPostV1TrainExperiments) | **Post** /v1/train/experiments | 
[**CloudPostV1TrainJobs**](TrainAPI.md#CloudPostV1TrainJobs) | **Post** /v1/train/jobs | 



## CloudDeleteV1TrainExperimentsName

> CloudDeleteV1TrainExperimentsName(ctx, name).Execute()

DeleteExperiment deletes a hyperparameter-tuning experiment.



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
	name := "sweep-1" // string | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource's metadata.name must be.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrainAPI.CloudDeleteV1TrainExperimentsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudDeleteV1TrainExperimentsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource&#39;s metadata.name must be. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1TrainExperimentsNameRequest struct via the builder pattern


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


## CloudDeleteV1TrainJobsName

> CloudDeleteV1TrainJobsName(ctx, name).Execute()

DeleteJob deletes a training job.



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
	name := "finetune-1" // string | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource's metadata.name must be.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrainAPI.CloudDeleteV1TrainJobsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudDeleteV1TrainJobsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource&#39;s metadata.name must be. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1TrainJobsNameRequest struct via the builder pattern


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


## CloudGetV1TrainExperiments

> CloudMlResourceList CloudGetV1TrainExperiments(ctx).Execute()

ListExperiments lists the caller org's tuning experiments.



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
	resp, r, err := apiClient.TrainAPI.CloudGetV1TrainExperiments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudGetV1TrainExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrainExperiments`: CloudMlResourceList
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CloudGetV1TrainExperiments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrainExperimentsRequest struct via the builder pattern


### Return type

[**CloudMlResourceList**](CloudMlResourceList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TrainExperimentsName

> CloudMlResource CloudGetV1TrainExperimentsName(ctx, name).Execute()

GetExperiment returns one hyperparameter-tuning experiment.



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
	name := "sweep-1" // string | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource's metadata.name must be.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CloudGetV1TrainExperimentsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudGetV1TrainExperimentsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrainExperimentsName`: CloudMlResource
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CloudGetV1TrainExperimentsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource&#39;s metadata.name must be. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrainExperimentsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudMlResource**](CloudMlResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TrainExperimentsNameTrials

> CloudMlTrials CloudGetV1TrainExperimentsNameTrials(ctx, name).Execute()

ListTrials lists the katib Trials one experiment owns.



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
	name := "sweep-1" // string | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource's metadata.name must be.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CloudGetV1TrainExperimentsNameTrials(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudGetV1TrainExperimentsNameTrials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrainExperimentsNameTrials`: CloudMlTrials
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CloudGetV1TrainExperimentsNameTrials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource&#39;s metadata.name must be. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrainExperimentsNameTrialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudMlTrials**](CloudMlTrials.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TrainHealth

> CloudGetV1TrainHealth(ctx).Execute()



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
	r, err := apiClient.TrainAPI.CloudGetV1TrainHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudGetV1TrainHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrainHealthRequest struct via the builder pattern


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


## CloudGetV1TrainJobs

> CloudMlResourceList CloudGetV1TrainJobs(ctx).Execute()

ListJobs lists the training jobs in the caller's org.



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
	resp, r, err := apiClient.TrainAPI.CloudGetV1TrainJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudGetV1TrainJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrainJobs`: CloudMlResourceList
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CloudGetV1TrainJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrainJobsRequest struct via the builder pattern


### Return type

[**CloudMlResourceList**](CloudMlResourceList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TrainJobsName

> CloudMlResource CloudGetV1TrainJobsName(ctx, name).Execute()

GetJob returns one training job.



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
	name := "finetune-1" // string | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource's metadata.name must be.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CloudGetV1TrainJobsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudGetV1TrainJobsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrainJobsName`: CloudMlResource
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CloudGetV1TrainJobsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource to act on, taken from the path. Lower-cased and trimmed to the DNS-1123 label a CustomResource&#39;s metadata.name must be. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrainJobsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudMlResource**](CloudMlResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1TrainExperiments

> CloudPostV1TrainExperiments(ctx).Execute()



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
	r, err := apiClient.TrainAPI.CloudPostV1TrainExperiments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudPostV1TrainExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TrainExperimentsRequest struct via the builder pattern


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


## CloudPostV1TrainJobs

> CloudPostV1TrainJobs(ctx).Execute()



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
	r, err := apiClient.TrainAPI.CloudPostV1TrainJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CloudPostV1TrainJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TrainJobsRequest struct via the builder pattern


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

