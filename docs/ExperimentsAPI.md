# \ExperimentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExperiments**](ExperimentsAPI.md#GetExperiments) | **Get** /v1/experiments | Is every experiment in the caller&#39;s org, with its variants, status and decision, ordered by project then id.
[**GetExperimentsById**](ExperimentsAPI.md#GetExperimentsById) | **Get** /v1/experiments/{id} | Is one experiment&#39;s definition and lifecycle: variants, weights, control arm, status and winner.
[**GetExperimentsByIdAssign**](ExperimentsAPI.md#GetExperimentsByIdAssign) | **Get** /v1/experiments/{id}/assign | Is the variant one subject is bucketed into, and the payload that variant carries.
[**GetExperimentsHealth**](ExperimentsAPI.md#GetExperimentsHealth) | **Get** /v1/experiments/health | Is whether the experiments subsystem is mounted and serving in this process.
[**PostExperiments**](ExperimentsAPI.md#PostExperiments) | **Post** /v1/experiments | Registers a controlled experiment AND puts its assignment flag live, in that order, so the arms start bucketing subjects the moment this returns 201 — the flag is created active at 100% rollout, with each variant weighted as declared.
[**PostExperimentsByIdAnalyze**](ExperimentsAPI.md#PostExperimentsByIdAnalyze) | **Post** /v1/experiments/{id}/analyze | Is per-variant conversion, lift and statistical significance against the control arm.
[**PostExperimentsByIdDecide**](ExperimentsAPI.md#PostExperimentsByIdDecide) | **Post** /v1/experiments/{id}/decide | Promotes one variant to the whole rollout and records who decided.



## GetExperiments

> ExperimentList GetExperiments(ctx).Execute()

Is every experiment in the caller's org, with its variants, status and decision, ordered by project then id.



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
	resp, r, err := apiClient.ExperimentsAPI.GetExperiments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.GetExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperiments`: ExperimentList
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.GetExperiments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentsRequest struct via the builder pattern


### Return type

[**ExperimentList**](ExperimentList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentsById

> Trial GetExperimentsById(ctx, id).Execute()

Is one experiment's definition and lifecycle: variants, weights, control arm, status and winner.



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
	id := "id_example" // string | ID is the experiment the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.GetExperimentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.GetExperimentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentsById`: Trial
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.GetExperimentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the experiment the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Trial**](Trial.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentsByIdAssign

> Assignment GetExperimentsByIdAssign(ctx, id).Subject(subject).Props(props).Execute()

Is the variant one subject is bucketed into, and the payload that variant carries.



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
	id := "id_example" // string | ID is the experiment the URL names.
	subject := "subject_example" // string | Subject is the unit to bucket — a user, org, session or audience key, matching the experiment's subjectKind.
	props := "props_example" // string | Props is a JSON object of person properties for targeting. A value that is not valid JSON is dropped rather than refused, so a malformed one changes the bucketing without saying so. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.GetExperimentsByIdAssign(context.Background(), id).Subject(subject).Props(props).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.GetExperimentsByIdAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentsByIdAssign`: Assignment
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.GetExperimentsByIdAssign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the experiment the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentsByIdAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** | Subject is the unit to bucket — a user, org, session or audience key, matching the experiment&#39;s subjectKind. | 
 **props** | **string** | Props is a JSON object of person properties for targeting. A value that is not valid JSON is dropped rather than refused, so a malformed one changes the bucketing without saying so. | 

### Return type

[**Assignment**](Assignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentsHealth

> Health GetExperimentsHealth(ctx).Execute()

Is whether the experiments subsystem is mounted and serving in this process.



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
	resp, r, err := apiClient.ExperimentsAPI.GetExperimentsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.GetExperimentsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentsHealth`: Health
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.GetExperimentsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentsHealthRequest struct via the builder pattern


### Return type

[**Health**](Health.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExperiments

> Trial PostExperiments(ctx).CreateBody(createBody).Execute()

Registers a controlled experiment AND puts its assignment flag live, in that order, so the arms start bucketing subjects the moment this returns 201 — the flag is created active at 100% rollout, with each variant weighted as declared.



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
	createBody := *openapiclient.NewCreateBody("Id_example", "MetricEvent_example", []openapiclient.Arm{*openapiclient.NewArm()}) // CreateBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.PostExperiments(context.Background()).CreateBody(createBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.PostExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperiments`: Trial
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.PostExperiments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExperimentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBody** | [**CreateBody**](CreateBody.md) |  | 

### Return type

[**Trial**](Trial.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExperimentsByIdAnalyze

> Analysis PostExperimentsByIdAnalyze(ctx, id).AnalyzeQuery(analyzeQuery).Execute()

Is per-variant conversion, lift and statistical significance against the control arm.



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
	id := "id_example" // string | ID is the experiment the URL names.
	analyzeQuery := *openapiclient.NewAnalyzeQuery() // AnalyzeQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.PostExperimentsByIdAnalyze(context.Background(), id).AnalyzeQuery(analyzeQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.PostExperimentsByIdAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperimentsByIdAnalyze`: Analysis
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.PostExperimentsByIdAnalyze`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the experiment the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExperimentsByIdAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyzeQuery** | [**AnalyzeQuery**](AnalyzeQuery.md) |  | 

### Return type

[**Analysis**](Analysis.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExperimentsByIdDecide

> Trial PostExperimentsByIdDecide(ctx, id).DecideBody(decideBody).Execute()

Promotes one variant to the whole rollout and records who decided.



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
	decideBody := *openapiclient.NewDecideBody("Winner_example") // DecideBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.PostExperimentsByIdDecide(context.Background(), id).DecideBody(decideBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.PostExperimentsByIdDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperimentsByIdDecide`: Trial
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.PostExperimentsByIdDecide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExperimentsByIdDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decideBody** | [**DecideBody**](DecideBody.md) |  | 

### Return type

[**Trial**](Trial.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

