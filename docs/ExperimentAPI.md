# \ExperimentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExperiment**](ExperimentAPI.md#GetExperiment) | **Get** /v1/experiment | Is every experiment in the caller&#39;s org, with its variants, status and decision, ordered by project then id.
[**GetExperimentById**](ExperimentAPI.md#GetExperimentById) | **Get** /v1/experiment/{id} | Is one experiment&#39;s definition and lifecycle: variants, weights, control arm, status and winner.
[**GetExperimentByIdAssign**](ExperimentAPI.md#GetExperimentByIdAssign) | **Get** /v1/experiment/{id}/assign | Is the variant one subject is bucketed into, and the payload that variant carries.
[**GetExperimentHealth**](ExperimentAPI.md#GetExperimentHealth) | **Get** /v1/experiment/health | Is whether the experiments subsystem is mounted and serving in this process.
[**PostExperiment**](ExperimentAPI.md#PostExperiment) | **Post** /v1/experiment | Registers a controlled experiment AND puts its assignment flag live, in that order, so the arms start bucketing subjects the moment this returns 201 — the flag is created active at 100% rollout, with each variant weighted as declared.
[**PostExperimentByIdAnalyze**](ExperimentAPI.md#PostExperimentByIdAnalyze) | **Post** /v1/experiment/{id}/analyze | Is per-variant conversion, lift and statistical significance against the control arm.
[**PostExperimentByIdDecide**](ExperimentAPI.md#PostExperimentByIdDecide) | **Post** /v1/experiment/{id}/decide | Promotes one variant to the whole rollout and records who decided.



## GetExperiment

> ExperimentList GetExperiment(ctx).Execute()

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
	resp, r, err := apiClient.ExperimentAPI.GetExperiment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.GetExperiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperiment`: ExperimentList
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.GetExperiment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentRequest struct via the builder pattern


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


## GetExperimentById

> Trial GetExperimentById(ctx, id).Execute()

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
	resp, r, err := apiClient.ExperimentAPI.GetExperimentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.GetExperimentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentById`: Trial
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.GetExperimentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the experiment the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentByIdRequest struct via the builder pattern


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


## GetExperimentByIdAssign

> Assignment GetExperimentByIdAssign(ctx, id).Subject(subject).Props(props).Execute()

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
	resp, r, err := apiClient.ExperimentAPI.GetExperimentByIdAssign(context.Background(), id).Subject(subject).Props(props).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.GetExperimentByIdAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentByIdAssign`: Assignment
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.GetExperimentByIdAssign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the experiment the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentByIdAssignRequest struct via the builder pattern


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


## GetExperimentHealth

> Health GetExperimentHealth(ctx).Execute()

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
	resp, r, err := apiClient.ExperimentAPI.GetExperimentHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.GetExperimentHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentHealth`: Health
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.GetExperimentHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentHealthRequest struct via the builder pattern


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


## PostExperiment

> Trial PostExperiment(ctx).CreateBody(createBody).Execute()

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
	resp, r, err := apiClient.ExperimentAPI.PostExperiment(context.Background()).CreateBody(createBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.PostExperiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperiment`: Trial
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.PostExperiment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExperimentRequest struct via the builder pattern


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


## PostExperimentByIdAnalyze

> Analysis PostExperimentByIdAnalyze(ctx, id).AnalyzeQuery(analyzeQuery).Execute()

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
	resp, r, err := apiClient.ExperimentAPI.PostExperimentByIdAnalyze(context.Background(), id).AnalyzeQuery(analyzeQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.PostExperimentByIdAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperimentByIdAnalyze`: Analysis
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.PostExperimentByIdAnalyze`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the experiment the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExperimentByIdAnalyzeRequest struct via the builder pattern


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


## PostExperimentByIdDecide

> Trial PostExperimentByIdDecide(ctx, id).DecideBody(decideBody).Execute()

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
	resp, r, err := apiClient.ExperimentAPI.PostExperimentByIdDecide(context.Background(), id).DecideBody(decideBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.PostExperimentByIdDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperimentByIdDecide`: Trial
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.PostExperimentByIdDecide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExperimentByIdDecideRequest struct via the builder pattern


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

