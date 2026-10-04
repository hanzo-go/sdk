# \EvalAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEvalDatasetsByName**](EvalAPI.md#DeleteEvalDatasetsByName) | **Delete** /v1/eval/datasets/{name} | Removes the named dataset of the caller&#39;s org AND all of its examples, in one transaction.
[**GetEvalDatasets**](EvalAPI.md#GetEvalDatasets) | **Get** /v1/eval/datasets | Is the datasets your org has, each with its name, description, metadata and timestamps.
[**GetEvalDatasetsByName**](EvalAPI.md#GetEvalDatasetsByName) | **Get** /v1/eval/datasets/{name} | Returns one dataset of the caller&#39;s org by name, together with its live item count — the one read that answers how big the set actually is.
[**GetEvalDatasetsByNameItems**](EvalAPI.md#GetEvalDatasetsByNameItems) | **Get** /v1/eval/datasets/{name}/items | Is the examples in one of your datasets — the set is named in the path, because this collection only exists inside one.
[**GetEvalEvaluators**](EvalAPI.md#GetEvalEvaluators) | **Get** /v1/eval/evaluators | Is the judges your org has defined, each with its judge model, criteria and the score name it writes under.
[**GetEvalMetrics**](EvalAPI.md#GetEvalMetrics) | **Get** /v1/eval/metrics | Is your org&#39;s AI overview board over a window: totals (generations, prompt and completion tokens, cost in cents, errors, success rate, distinct models and users), a gap-filled time series, a per-model breakdown with the long tail folded into \&quot;other\&quot;, and latency percentiles read from the GenAI spans.
[**GetEvalRubrics**](EvalAPI.md#GetEvalRubrics) | **Get** /v1/eval/rubrics | Is the score shapes your org has declared — each name&#39;s data type, its numeric bounds and its allowed categories.
[**GetEvalRuns**](EvalAPI.md#GetEvalRuns) | **Get** /v1/eval/runs | Is your past runs and how they scored — the dataset and model, the judge model, how many examples were attempted and how many scored, the average score, and when it happened.
[**GetEvalScores**](EvalAPI.md#GetEvalScores) | **Get** /v1/eval/scores | Is the score events your org has recorded, narrowed by any of name, runName and traceId.
[**GetEvalTraces**](EvalAPI.md#GetEvalTraces) | **Get** /v1/eval/traces | Is the traces behind your evaluations — one per model call an evaluation made, carrying its input, output, model and timing — narrowed by any of sessionId, runName and datasetName.
[**PostEvalDatasets**](EvalAPI.md#PostEvalDatasets) | **Post** /v1/eval/datasets | Writes a dataset — the named set of graded examples a run scores a model against — under the caller&#39;s org and answers 201 with it.
[**PostEvalDatasetsByNameItems**](EvalAPI.md#PostEvalDatasetsByNameItems) | **Post** /v1/eval/datasets/{name}/items | Writes one graded example — its input, its expected output, free-form metadata and a status — into the dataset named in the path, and answers 201 with it.
[**PostEvalEvaluators**](EvalAPI.md#PostEvalEvaluators) | **Post** /v1/eval/evaluators | Saves a reusable judge for the caller&#39;s org — the judge model and the written criteria it grades against — and answers 201 with it.
[**PostEvalRubrics**](EvalAPI.md#PostEvalRubrics) | **Post** /v1/eval/rubrics | Defines the shape of one score name for the caller&#39;s org and answers 201 with it.
[**PostEvalRuns**](EvalAPI.md#PostEvalRuns) | **Post** /v1/eval/runs | Runs a real evaluation and answers the summary when it is finished — this is synchronous work, not a job id.
[**PostEvalScores**](EvalAPI.md#PostEvalScores) | **Post** /v1/eval/scores | Files one score event for the caller&#39;s org and answers 201 with it.



## DeleteEvalDatasetsByName

> DeleteEvalDatasetsByName(ctx, name).Execute()

Removes the named dataset of the caller's org AND all of its examples, in one transaction.



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
	name := "name_example" // string | Name is the dataset the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EvalAPI.DeleteEvalDatasetsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.DeleteEvalDatasetsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEvalDatasetsByNameRequest struct via the builder pattern


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


## GetEvalDatasets

> EvalDatasetList GetEvalDatasets(ctx).Limit(limit).Execute()

Is the datasets your org has, each with its name, description, metadata and timestamps.



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
	limit := int64(789) // int64 | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalDatasets(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalDatasets`: EvalDatasetList
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalDatasets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalDatasetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. | 

### Return type

[**EvalDatasetList**](EvalDatasetList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalDatasetsByName

> EvalDatasetView GetEvalDatasetsByName(ctx, name).Execute()

Returns one dataset of the caller's org by name, together with its live item count — the one read that answers how big the set actually is.



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
	name := "name_example" // string | Name is the dataset the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalDatasetsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalDatasetsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalDatasetsByName`: EvalDatasetView
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalDatasetsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalDatasetsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EvalDatasetView**](EvalDatasetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalDatasetsByNameItems

> EvalItemList GetEvalDatasetsByNameItems(ctx, name).Limit(limit).Execute()

Is the examples in one of your datasets — the set is named in the path, because this collection only exists inside one.



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
	name := "name_example" // string | Dataset is the set to read, from the path — this collection only exists inside one.
	limit := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalDatasetsByNameItems(context.Background(), name).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalDatasetsByNameItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalDatasetsByNameItems`: EvalItemList
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalDatasetsByNameItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Dataset is the set to read, from the path — this collection only exists inside one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalDatasetsByNameItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** |  | 

### Return type

[**EvalItemList**](EvalItemList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalEvaluators

> EvalEvaluatorList GetEvalEvaluators(ctx).Limit(limit).Execute()

Is the judges your org has defined, each with its judge model, criteria and the score name it writes under.



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
	limit := int64(789) // int64 | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalEvaluators(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalEvaluators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalEvaluators`: EvalEvaluatorList
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalEvaluators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalEvaluatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. | 

### Return type

[**EvalEvaluatorList**](EvalEvaluatorList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalMetrics

> EvalBoard GetEvalMetrics(ctx).Range_(range_).Interval(interval).Execute()

Is your org's AI overview board over a window: totals (generations, prompt and completion tokens, cost in cents, errors, success rate, distinct models and users), a gap-filled time series, a per-model breakdown with the long tail folded into \"other\", and latency percentiles read from the GenAI spans.



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
	range_ := "range__example" // string | Range is 24h (the default), 7d or 30d. Anything else normalises to 24h rather than failing, so the board always has a valid window. (optional)
	interval := "interval_example" // string | Interval overrides the bucket the series is grouped into: \"hour\" or \"day\". Any other value leaves the range's own default in place. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalMetrics(context.Background()).Range_(range_).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalMetrics`: EvalBoard
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is 24h (the default), 7d or 30d. Anything else normalises to 24h rather than failing, so the board always has a valid window. | 
 **interval** | **string** | Interval overrides the bucket the series is grouped into: \&quot;hour\&quot; or \&quot;day\&quot;. Any other value leaves the range&#39;s own default in place. | 

### Return type

[**EvalBoard**](EvalBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalRubrics

> EvalScoreConfigList GetEvalRubrics(ctx).Limit(limit).Execute()

Is the score shapes your org has declared — each name's data type, its numeric bounds and its allowed categories.



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
	limit := int64(789) // int64 | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalRubrics(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalRubrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalRubrics`: EvalScoreConfigList
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalRubrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalRubricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. | 

### Return type

[**EvalScoreConfigList**](EvalScoreConfigList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalRuns

> EvalRuns GetEvalRuns(ctx).DatasetName(datasetName).Limit(limit).Execute()

Is your past runs and how they scored — the dataset and model, the judge model, how many examples were attempted and how many scored, the average score, and when it happened.



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
	datasetName := "datasetName_example" // string | Dataset narrows to the runs against one dataset. (optional)
	limit := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalRuns(context.Background()).DatasetName(datasetName).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalRuns`: EvalRuns
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasetName** | **string** | Dataset narrows to the runs against one dataset. | 
 **limit** | **int64** |  | 

### Return type

[**EvalRuns**](EvalRuns.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalScores

> EvalScoreList GetEvalScores(ctx).Name(name).RunName(runName).TraceId(traceId).Limit(limit).Execute()

Is the score events your org has recorded, narrowed by any of name, runName and traceId.



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
	name := "name_example" // string | Name narrows to one score name. (optional)
	runName := "runName_example" // string | RunName narrows to the scores of one run. (optional)
	traceId := "traceId_example" // string | TraceID narrows to the scores on one model call. (optional)
	limit := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalScores(context.Background()).Name(name).RunName(runName).TraceId(traceId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalScores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalScores`: EvalScoreList
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name narrows to one score name. | 
 **runName** | **string** | RunName narrows to the scores of one run. | 
 **traceId** | **string** | TraceID narrows to the scores on one model call. | 
 **limit** | **int64** |  | 

### Return type

[**EvalScoreList**](EvalScoreList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalTraces

> EvalTraceList GetEvalTraces(ctx).SessionId(sessionId).RunName(runName).DatasetName(datasetName).Limit(limit).Execute()

Is the traces behind your evaluations — one per model call an evaluation made, carrying its input, output, model and timing — narrowed by any of sessionId, runName and datasetName.



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
	sessionId := "sessionId_example" // string | SessionID narrows to one session, which for an evaluation is one run. (optional)
	runName := "runName_example" // string | RunName narrows to the calls one run made. (optional)
	datasetName := "datasetName_example" // string | Dataset narrows to the calls made against one dataset. (optional)
	limit := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.GetEvalTraces(context.Background()).SessionId(sessionId).RunName(runName).DatasetName(datasetName).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.GetEvalTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalTraces`: EvalTraceList
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.GetEvalTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionId** | **string** | SessionID narrows to one session, which for an evaluation is one run. | 
 **runName** | **string** | RunName narrows to the calls one run made. | 
 **datasetName** | **string** | Dataset narrows to the calls made against one dataset. | 
 **limit** | **int64** |  | 

### Return type

[**EvalTraceList**](EvalTraceList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalDatasets

> EvalDatasetView PostEvalDatasets(ctx).EvalDatasetReq(evalDatasetReq).Execute()

Writes a dataset — the named set of graded examples a run scores a model against — under the caller's org and answers 201 with it.



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
	evalDatasetReq := *openapiclient.NewEvalDatasetReq("Name_example") // EvalDatasetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.PostEvalDatasets(context.Background()).EvalDatasetReq(evalDatasetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.PostEvalDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalDatasets`: EvalDatasetView
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.PostEvalDatasets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalDatasetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalDatasetReq** | [**EvalDatasetReq**](EvalDatasetReq.md) |  | 

### Return type

[**EvalDatasetView**](EvalDatasetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalDatasetsByNameItems

> EvalItemView PostEvalDatasetsByNameItems(ctx, name).EvalItemReq(evalItemReq).Execute()

Writes one graded example — its input, its expected output, free-form metadata and a status — into the dataset named in the path, and answers 201 with it.



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
	name := "name_example" // string | 
	evalItemReq := *openapiclient.NewEvalItemReq() // EvalItemReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.PostEvalDatasetsByNameItems(context.Background(), name).EvalItemReq(evalItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.PostEvalDatasetsByNameItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalDatasetsByNameItems`: EvalItemView
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.PostEvalDatasetsByNameItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalDatasetsByNameItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evalItemReq** | [**EvalItemReq**](EvalItemReq.md) |  | 

### Return type

[**EvalItemView**](EvalItemView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalEvaluators

> EvalEvaluatorView PostEvalEvaluators(ctx).EvalEvaluatorReq(evalEvaluatorReq).Execute()

Saves a reusable judge for the caller's org — the judge model and the written criteria it grades against — and answers 201 with it.



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
	evalEvaluatorReq := *openapiclient.NewEvalEvaluatorReq("Name_example") // EvalEvaluatorReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.PostEvalEvaluators(context.Background()).EvalEvaluatorReq(evalEvaluatorReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.PostEvalEvaluators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalEvaluators`: EvalEvaluatorView
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.PostEvalEvaluators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalEvaluatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalEvaluatorReq** | [**EvalEvaluatorReq**](EvalEvaluatorReq.md) |  | 

### Return type

[**EvalEvaluatorView**](EvalEvaluatorView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalRubrics

> EvalScoreConfigView PostEvalRubrics(ctx).EvalScoreConfigReq(evalScoreConfigReq).Execute()

Defines the shape of one score name for the caller's org and answers 201 with it.



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
	evalScoreConfigReq := *openapiclient.NewEvalScoreConfigReq("Name_example") // EvalScoreConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.PostEvalRubrics(context.Background()).EvalScoreConfigReq(evalScoreConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.PostEvalRubrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalRubrics`: EvalScoreConfigView
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.PostEvalRubrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalRubricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalScoreConfigReq** | [**EvalScoreConfigReq**](EvalScoreConfigReq.md) |  | 

### Return type

[**EvalScoreConfigView**](EvalScoreConfigView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalRuns

> EvalRunSummary PostEvalRuns(ctx).EvalRunRequest(evalRunRequest).Authorization(authorization).Execute()

Runs a real evaluation and answers the summary when it is finished — this is synchronous work, not a job id.



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
	evalRunRequest := *openapiclient.NewEvalRunRequest("Dataset_example", "Model_example") // EvalRunRequest | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.PostEvalRuns(context.Background()).EvalRunRequest(evalRunRequest).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.PostEvalRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalRuns`: EvalRunSummary
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.PostEvalRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalRunRequest** | [**EvalRunRequest**](EvalRunRequest.md) |  | 
 **authorization** | **string** |  | 

### Return type

[**EvalRunSummary**](EvalRunSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalScores

> EvalScoreView PostEvalScores(ctx).EvalScoreReq(evalScoreReq).Execute()

Files one score event for the caller's org and answers 201 with it.



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
	evalScoreReq := *openapiclient.NewEvalScoreReq("Name_example") // EvalScoreReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalAPI.PostEvalScores(context.Background()).EvalScoreReq(evalScoreReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalAPI.PostEvalScores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalScores`: EvalScoreView
	fmt.Fprintf(os.Stdout, "Response from `EvalAPI.PostEvalScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalScoreReq** | [**EvalScoreReq**](EvalScoreReq.md) |  | 

### Return type

[**EvalScoreView**](EvalScoreView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

