# \EvalsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEvalsDatasetsByName**](EvalsAPI.md#DeleteEvalsDatasetsByName) | **Delete** /v1/evals/datasets/{name} | Removes the named dataset of the caller&#39;s org AND all of its examples, in one transaction.
[**GetEvalsDatasets**](EvalsAPI.md#GetEvalsDatasets) | **Get** /v1/evals/datasets | Is the datasets your org has, each with its name, description, metadata and timestamps.
[**GetEvalsDatasetsByName**](EvalsAPI.md#GetEvalsDatasetsByName) | **Get** /v1/evals/datasets/{name} | Returns one dataset of the caller&#39;s org by name, together with its live item count — the one read that answers how big the set actually is.
[**GetEvalsDatasetsByNameItems**](EvalsAPI.md#GetEvalsDatasetsByNameItems) | **Get** /v1/evals/datasets/{name}/items | Is the examples in one of your datasets — the set is named in the path, because this collection only exists inside one.
[**GetEvalsEvaluators**](EvalsAPI.md#GetEvalsEvaluators) | **Get** /v1/evals/evaluators | Is the judges your org has defined, each with its judge model, criteria and the score name it writes under.
[**GetEvalsMetrics**](EvalsAPI.md#GetEvalsMetrics) | **Get** /v1/evals/metrics | Is your org&#39;s AI overview board over a window: totals (generations, prompt and completion tokens, cost in cents, errors, success rate, distinct models and users), a gap-filled time series, a per-model breakdown with the long tail folded into \&quot;other\&quot;, and latency percentiles read from the GenAI spans.
[**GetEvalsRubrics**](EvalsAPI.md#GetEvalsRubrics) | **Get** /v1/evals/rubrics | Is the score shapes your org has declared — each name&#39;s data type, its numeric bounds and its allowed categories.
[**GetEvalsRuns**](EvalsAPI.md#GetEvalsRuns) | **Get** /v1/evals/runs | Is your past runs and how they scored — the dataset and model, the judge model, how many examples were attempted and how many scored, the average score, and when it happened.
[**GetEvalsScores**](EvalsAPI.md#GetEvalsScores) | **Get** /v1/evals/scores | Is the score events your org has recorded, narrowed by any of name, runName and traceId.
[**GetEvalsTraces**](EvalsAPI.md#GetEvalsTraces) | **Get** /v1/evals/traces | Is the traces behind your evaluations — one per model call an evaluation made, carrying its input, output, model and timing — narrowed by any of sessionId, runName and datasetName.
[**PostEvalsDatasets**](EvalsAPI.md#PostEvalsDatasets) | **Post** /v1/evals/datasets | Writes a dataset — the named set of graded examples a run scores a model against — under the caller&#39;s org and answers 201 with it.
[**PostEvalsDatasetsByNameItems**](EvalsAPI.md#PostEvalsDatasetsByNameItems) | **Post** /v1/evals/datasets/{name}/items | Writes one graded example — its input, its expected output, free-form metadata and a status — into the dataset named in the path, and answers 201 with it.
[**PostEvalsEvaluators**](EvalsAPI.md#PostEvalsEvaluators) | **Post** /v1/evals/evaluators | Saves a reusable judge for the caller&#39;s org — the judge model and the written criteria it grades against — and answers 201 with it.
[**PostEvalsRubrics**](EvalsAPI.md#PostEvalsRubrics) | **Post** /v1/evals/rubrics | Defines the shape of one score name for the caller&#39;s org and answers 201 with it.
[**PostEvalsRuns**](EvalsAPI.md#PostEvalsRuns) | **Post** /v1/evals/runs | Runs a real evaluation and answers the summary when it is finished — this is synchronous work, not a job id.
[**PostEvalsScores**](EvalsAPI.md#PostEvalsScores) | **Post** /v1/evals/scores | Files one score event for the caller&#39;s org and answers 201 with it.



## DeleteEvalsDatasetsByName

> map[string]interface{} DeleteEvalsDatasetsByName(ctx, name).Execute()

Removes the named dataset of the caller's org AND all of its examples, in one transaction.



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
	name := "name_example" // string | Name is the dataset the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.DeleteEvalsDatasetsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.DeleteEvalsDatasetsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEvalsDatasetsByName`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.DeleteEvalsDatasetsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEvalsDatasetsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsDatasets

> DatasetList GetEvalsDatasets(ctx).Limit(limit).Execute()

Is the datasets your org has, each with its name, description, metadata and timestamps.



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
	limit := int32(56) // int32 | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsDatasets(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsDatasets`: DatasetList
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsDatasets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsDatasetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. | 

### Return type

[**DatasetList**](DatasetList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsDatasetsByName

> DatasetView GetEvalsDatasetsByName(ctx, name).Execute()

Returns one dataset of the caller's org by name, together with its live item count — the one read that answers how big the set actually is.



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
	name := "name_example" // string | Name is the dataset the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsDatasetsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsDatasetsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsDatasetsByName`: DatasetView
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsDatasetsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsDatasetsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatasetView**](DatasetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsDatasetsByNameItems

> ItemList GetEvalsDatasetsByNameItems(ctx, name).Limit(limit).Execute()

Is the examples in one of your datasets — the set is named in the path, because this collection only exists inside one.



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
	name := "name_example" // string | Dataset is the set to read, from the path — this collection only exists inside one.
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsDatasetsByNameItems(context.Background(), name).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsDatasetsByNameItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsDatasetsByNameItems`: ItemList
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsDatasetsByNameItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Dataset is the set to read, from the path — this collection only exists inside one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsDatasetsByNameItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 

### Return type

[**ItemList**](ItemList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsEvaluators

> EvaluatorList GetEvalsEvaluators(ctx).Limit(limit).Execute()

Is the judges your org has defined, each with its judge model, criteria and the score name it writes under.



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
	limit := int32(56) // int32 | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsEvaluators(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsEvaluators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsEvaluators`: EvaluatorList
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsEvaluators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsEvaluatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. | 

### Return type

[**EvaluatorList**](EvaluatorList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsMetrics

> Board GetEvalsMetrics(ctx).Range_(range_).Interval(interval).Execute()

Is your org's AI overview board over a window: totals (generations, prompt and completion tokens, cost in cents, errors, success rate, distinct models and users), a gap-filled time series, a per-model breakdown with the long tail folded into \"other\", and latency percentiles read from the GenAI spans.



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
	range_ := "range__example" // string | Range is 24h (the default), 7d or 30d. Anything else normalises to 24h rather than failing, so the board always has a valid window. (optional)
	interval := "interval_example" // string | Interval overrides the bucket the series is grouped into: \"hour\" or \"day\". Any other value leaves the range's own default in place. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsMetrics(context.Background()).Range_(range_).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsMetrics`: Board
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is 24h (the default), 7d or 30d. Anything else normalises to 24h rather than failing, so the board always has a valid window. | 
 **interval** | **string** | Interval overrides the bucket the series is grouped into: \&quot;hour\&quot; or \&quot;day\&quot;. Any other value leaves the range&#39;s own default in place. | 

### Return type

[**Board**](Board.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsRubrics

> ScoreConfigList GetEvalsRubrics(ctx).Limit(limit).Execute()

Is the score shapes your org has declared — each name's data type, its numeric bounds and its allowed categories.



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
	limit := int32(56) // int32 | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsRubrics(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsRubrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsRubrics`: ScoreConfigList
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsRubrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsRubricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned. It defaults to 100 and is capped at 500; a non-positive or unparseable value falls back to the default rather than failing, because a typo about paging is not a reason to refuse a read. | 

### Return type

[**ScoreConfigList**](ScoreConfigList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsRuns

> Runs GetEvalsRuns(ctx).DatasetName(datasetName).Limit(limit).Execute()

Is your past runs and how they scored — the dataset and model, the judge model, how many examples were attempted and how many scored, the average score, and when it happened.



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
	datasetName := "datasetName_example" // string | Dataset narrows to the runs against one dataset. (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsRuns(context.Background()).DatasetName(datasetName).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsRuns`: Runs
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasetName** | **string** | Dataset narrows to the runs against one dataset. | 
 **limit** | **int32** |  | 

### Return type

[**Runs**](Runs.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsScores

> ScoreList GetEvalsScores(ctx).Name(name).RunName(runName).TraceId(traceId).Limit(limit).Execute()

Is the score events your org has recorded, narrowed by any of name, runName and traceId.



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
	name := "name_example" // string | Name narrows to one score name. (optional)
	runName := "runName_example" // string | RunName narrows to the scores of one run. (optional)
	traceId := "traceId_example" // string | TraceID narrows to the scores on one model call. (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsScores(context.Background()).Name(name).RunName(runName).TraceId(traceId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsScores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsScores`: ScoreList
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name narrows to one score name. | 
 **runName** | **string** | RunName narrows to the scores of one run. | 
 **traceId** | **string** | TraceID narrows to the scores on one model call. | 
 **limit** | **int32** |  | 

### Return type

[**ScoreList**](ScoreList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvalsTraces

> TraceList GetEvalsTraces(ctx).SessionId(sessionId).RunName(runName).DatasetName(datasetName).Limit(limit).Execute()

Is the traces behind your evaluations — one per model call an evaluation made, carrying its input, output, model and timing — narrowed by any of sessionId, runName and datasetName.



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
	sessionId := "sessionId_example" // string | SessionID narrows to one session, which for an evaluation is one run. (optional)
	runName := "runName_example" // string | RunName narrows to the calls one run made. (optional)
	datasetName := "datasetName_example" // string | Dataset narrows to the calls made against one dataset. (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.GetEvalsTraces(context.Background()).SessionId(sessionId).RunName(runName).DatasetName(datasetName).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.GetEvalsTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvalsTraces`: TraceList
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.GetEvalsTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEvalsTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionId** | **string** | SessionID narrows to one session, which for an evaluation is one run. | 
 **runName** | **string** | RunName narrows to the calls one run made. | 
 **datasetName** | **string** | Dataset narrows to the calls made against one dataset. | 
 **limit** | **int32** |  | 

### Return type

[**TraceList**](TraceList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalsDatasets

> DatasetView PostEvalsDatasets(ctx).DatasetReq(datasetReq).Execute()

Writes a dataset — the named set of graded examples a run scores a model against — under the caller's org and answers 201 with it.



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
	datasetReq := *openapiclient.NewDatasetReq("Name_example") // DatasetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.PostEvalsDatasets(context.Background()).DatasetReq(datasetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.PostEvalsDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalsDatasets`: DatasetView
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.PostEvalsDatasets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalsDatasetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasetReq** | [**DatasetReq**](DatasetReq.md) |  | 

### Return type

[**DatasetView**](DatasetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalsDatasetsByNameItems

> ItemView PostEvalsDatasetsByNameItems(ctx, name).ItemReq(itemReq).Execute()

Writes one graded example — its input, its expected output, free-form metadata and a status — into the dataset named in the path, and answers 201 with it.



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
	name := "name_example" // string | 
	itemReq := *openapiclient.NewItemReq() // ItemReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.PostEvalsDatasetsByNameItems(context.Background(), name).ItemReq(itemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.PostEvalsDatasetsByNameItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalsDatasetsByNameItems`: ItemView
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.PostEvalsDatasetsByNameItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalsDatasetsByNameItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemReq** | [**ItemReq**](ItemReq.md) |  | 

### Return type

[**ItemView**](ItemView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalsEvaluators

> EvaluatorView PostEvalsEvaluators(ctx).EvaluatorReq(evaluatorReq).Execute()

Saves a reusable judge for the caller's org — the judge model and the written criteria it grades against — and answers 201 with it.



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
	evaluatorReq := *openapiclient.NewEvaluatorReq("Name_example") // EvaluatorReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.PostEvalsEvaluators(context.Background()).EvaluatorReq(evaluatorReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.PostEvalsEvaluators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalsEvaluators`: EvaluatorView
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.PostEvalsEvaluators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalsEvaluatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluatorReq** | [**EvaluatorReq**](EvaluatorReq.md) |  | 

### Return type

[**EvaluatorView**](EvaluatorView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalsRubrics

> ScoreConfigView PostEvalsRubrics(ctx).ScoreConfigReq(scoreConfigReq).Execute()

Defines the shape of one score name for the caller's org and answers 201 with it.



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
	scoreConfigReq := *openapiclient.NewScoreConfigReq("Name_example") // ScoreConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.PostEvalsRubrics(context.Background()).ScoreConfigReq(scoreConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.PostEvalsRubrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalsRubrics`: ScoreConfigView
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.PostEvalsRubrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalsRubricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scoreConfigReq** | [**ScoreConfigReq**](ScoreConfigReq.md) |  | 

### Return type

[**ScoreConfigView**](ScoreConfigView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalsRuns

> RunSummary PostEvalsRuns(ctx).RunRequest(runRequest).Authorization(authorization).Execute()

Runs a real evaluation and answers the summary when it is finished — this is synchronous work, not a job id.



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
	runRequest := *openapiclient.NewRunRequest("Dataset_example", "Model_example") // RunRequest | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.PostEvalsRuns(context.Background()).RunRequest(runRequest).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.PostEvalsRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalsRuns`: RunSummary
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.PostEvalsRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalsRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runRequest** | [**RunRequest**](RunRequest.md) |  | 
 **authorization** | **string** |  | 

### Return type

[**RunSummary**](RunSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvalsScores

> ScoreView PostEvalsScores(ctx).ScoreReq(scoreReq).Execute()

Files one score event for the caller's org and answers 201 with it.



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
	scoreReq := *openapiclient.NewScoreReq("Name_example") // ScoreReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsAPI.PostEvalsScores(context.Background()).ScoreReq(scoreReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsAPI.PostEvalsScores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvalsScores`: ScoreView
	fmt.Fprintf(os.Stdout, "Response from `EvalsAPI.PostEvalsScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEvalsScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scoreReq** | [**ScoreReq**](ScoreReq.md) |  | 

### Return type

[**ScoreView**](ScoreView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

