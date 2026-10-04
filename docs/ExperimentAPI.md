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

> ExperimentExperimentList GetExperiment(ctx).Execute()

Is every experiment in the caller's org, with its variants, status and decision, ordered by project then id.



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
	resp, r, err := apiClient.ExperimentAPI.GetExperiment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.GetExperiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperiment`: ExperimentExperimentList
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.GetExperiment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentRequest struct via the builder pattern


### Return type

[**ExperimentExperimentList**](ExperimentExperimentList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentById

> ExperimentTrial GetExperimentById(ctx, id).Execute()

Is one experiment's definition and lifecycle: variants, weights, control arm, status and winner.



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
	id := "id_example" // string | ID is the experiment the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentAPI.GetExperimentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.GetExperimentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentById`: ExperimentTrial
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

[**ExperimentTrial**](ExperimentTrial.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentByIdAssign

> ExperimentAssignment GetExperimentByIdAssign(ctx, id).Subject(subject).Props(props).Execute()

Is the variant one subject is bucketed into, and the payload that variant carries.



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
	// response from `GetExperimentByIdAssign`: ExperimentAssignment
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

[**ExperimentAssignment**](ExperimentAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentHealth

> ExperimentHealth GetExperimentHealth(ctx).Execute()

Is whether the experiments subsystem is mounted and serving in this process.



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
	resp, r, err := apiClient.ExperimentAPI.GetExperimentHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.GetExperimentHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentHealth`: ExperimentHealth
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.GetExperimentHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentHealthRequest struct via the builder pattern


### Return type

[**ExperimentHealth**](ExperimentHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExperiment

> ExperimentTrial PostExperiment(ctx).ExperimentCreateBody(experimentCreateBody).Execute()

Registers a controlled experiment AND puts its assignment flag live, in that order, so the arms start bucketing subjects the moment this returns 201 — the flag is created active at 100% rollout, with each variant weighted as declared.



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
	experimentCreateBody := *openapiclient.NewExperimentCreateBody("Id_example", "MetricEvent_example", []openapiclient.ExperimentArm{*openapiclient.NewExperimentArm()}) // ExperimentCreateBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentAPI.PostExperiment(context.Background()).ExperimentCreateBody(experimentCreateBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.PostExperiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperiment`: ExperimentTrial
	fmt.Fprintf(os.Stdout, "Response from `ExperimentAPI.PostExperiment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **experimentCreateBody** | [**ExperimentCreateBody**](ExperimentCreateBody.md) |  | 

### Return type

[**ExperimentTrial**](ExperimentTrial.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExperimentByIdAnalyze

> ExperimentAnalysis PostExperimentByIdAnalyze(ctx, id).ExperimentAnalyzeQuery(experimentAnalyzeQuery).Execute()

Is per-variant conversion, lift and statistical significance against the control arm.



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
	id := "id_example" // string | ID is the experiment the URL names.
	experimentAnalyzeQuery := *openapiclient.NewExperimentAnalyzeQuery() // ExperimentAnalyzeQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentAPI.PostExperimentByIdAnalyze(context.Background(), id).ExperimentAnalyzeQuery(experimentAnalyzeQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.PostExperimentByIdAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperimentByIdAnalyze`: ExperimentAnalysis
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

 **experimentAnalyzeQuery** | [**ExperimentAnalyzeQuery**](ExperimentAnalyzeQuery.md) |  | 

### Return type

[**ExperimentAnalysis**](ExperimentAnalysis.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExperimentByIdDecide

> ExperimentTrial PostExperimentByIdDecide(ctx, id).ExperimentDecideBody(experimentDecideBody).Execute()

Promotes one variant to the whole rollout and records who decided.



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
	experimentDecideBody := *openapiclient.NewExperimentDecideBody("Winner_example") // ExperimentDecideBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentAPI.PostExperimentByIdDecide(context.Background(), id).ExperimentDecideBody(experimentDecideBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentAPI.PostExperimentByIdDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExperimentByIdDecide`: ExperimentTrial
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

 **experimentDecideBody** | [**ExperimentDecideBody**](ExperimentDecideBody.md) |  | 

### Return type

[**ExperimentTrial**](ExperimentTrial.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

