# \BenchmarkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBenchmarkCatalog**](BenchmarkAPI.md#GetBenchmarkCatalog) | **Get** /v1/benchmark/catalog | Is the canonical public benchmarks this arena runs — the id, title, axis, item count and upstream source of each, with native marking the ones the standardized harness runs today; the rest are registered and adapter-pending.
[**GetBenchmarkClaims**](BenchmarkAPI.md#GetBenchmarkClaims) | **Get** /v1/benchmark/claims | Lists the effective published claims: what the leaderboard will use for each (benchmark, model) after the seed, the import and any stored correction are layered.
[**GetBenchmarkCompare**](BenchmarkAPI.md#GetBenchmarkCompare) | **Get** /v1/benchmark/compare | Is the ONLY valid arm-vs-arm test: it pairs the two models on the items BOTH completed, and answers rescue and damage counts with an exact-McNemar p.
[**GetBenchmarkHistory**](BenchmarkAPI.md#GetBenchmarkHistory) | **Get** /v1/benchmark/history | Returns each model&#39;s measured score per run over time, oldest first, with the change between runs.
[**GetBenchmarkLeaderboard**](BenchmarkAPI.md#GetBenchmarkLeaderboard) | **Get** /v1/benchmark/leaderboard | Answers one row per model for the benchmark named — what our own harness measured, beside what the vendor claims, and the gap between them.
[**GetBenchmarkPresets**](BenchmarkAPI.md#GetBenchmarkPresets) | **Get** /v1/benchmark/presets | Are the router blends available to compose from — a named set of model arms, the rank they escalate through and the panel width that bounds fan-out — each served by the model layer as enso-&lt;name&gt;.
[**PostBenchmarkClaims**](BenchmarkAPI.md#PostBenchmarkClaims) | **Post** /v1/benchmark/claims | Records published claims: one to correct a number, many to import a leaderboard.
[**PostBenchmarkPresets**](BenchmarkAPI.md#PostBenchmarkPresets) | **Post** /v1/benchmark/presets | Validates a router blend — its name, its arms, the rank they escalate through and the panel fan-out width — and answers 202 with the preset and the enso-&lt;name&gt; it would be served as.
[**PostBenchmarkRuns**](BenchmarkAPI.md#PostBenchmarkRuns) | **Post** /v1/benchmark/runs | Admits and queues a benchmark run against a model or your own endpoint, and answers 202 with the receipt.



## GetBenchmarkCatalog

> BenchmarkCatalog GetBenchmarkCatalog(ctx).Execute()

Is the canonical public benchmarks this arena runs — the id, title, axis, item count and upstream source of each, with native marking the ones the standardized harness runs today; the rest are registered and adapter-pending.



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
	resp, r, err := apiClient.BenchmarkAPI.GetBenchmarkCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.GetBenchmarkCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBenchmarkCatalog`: BenchmarkCatalog
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.GetBenchmarkCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchmarkCatalogRequest struct via the builder pattern


### Return type

[**BenchmarkCatalog**](BenchmarkCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkClaims

> ClaimsOut GetBenchmarkClaims(ctx).Benchmark(benchmark).Model(model).Provider(provider).Source(source).Protocol(protocol).Execute()

Lists the effective published claims: what the leaderboard will use for each (benchmark, model) after the seed, the import and any stored correction are layered.



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
	benchmark := "benchmark_example" // string | Benchmark filters to one benchmark id. Empty returns every benchmark. (optional)
	model := "model_example" // string | Model filters to one model. Empty returns every model. (optional)
	provider := "provider_example" // string | Provider filters to one lab or leaderboard — the way to read what a single source claims across every model it covers. (optional)
	source := "source_example" // string | Source filters to one citation, which is the finest grain there is: a source is what makes two claims about one model independent rather than a restatement of each other. (optional)
	protocol := "protocol_example" // string | Protocol filters by HOW a claim was scored, so provider cards can be read apart from third parties running their own harness. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.GetBenchmarkClaims(context.Background()).Benchmark(benchmark).Model(model).Provider(provider).Source(source).Protocol(protocol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.GetBenchmarkClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBenchmarkClaims`: ClaimsOut
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.GetBenchmarkClaims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchmarkClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchmark** | **string** | Benchmark filters to one benchmark id. Empty returns every benchmark. | 
 **model** | **string** | Model filters to one model. Empty returns every model. | 
 **provider** | **string** | Provider filters to one lab or leaderboard — the way to read what a single source claims across every model it covers. | 
 **source** | **string** | Source filters to one citation, which is the finest grain there is: a source is what makes two claims about one model independent rather than a restatement of each other. | 
 **protocol** | **string** | Protocol filters by HOW a claim was scored, so provider cards can be read apart from third parties running their own harness. | 

### Return type

[**ClaimsOut**](ClaimsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkCompare

> Pairing GetBenchmarkCompare(ctx).A(a).B(b).Benchmark(benchmark).Execute()

Is the ONLY valid arm-vs-arm test: it pairs the two models on the items BOTH completed, and answers rescue and damage counts with an exact-McNemar p.



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
	a := "a_example" // string | A is the first model id. It is required.
	b := "b_example" // string | B is the second model id. It is required.
	benchmark := "benchmark_example" // string | Benchmark is the catalog id to compare on, defaulting to gpqa_diamond. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.GetBenchmarkCompare(context.Background()).A(a).B(b).Benchmark(benchmark).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.GetBenchmarkCompare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBenchmarkCompare`: Pairing
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.GetBenchmarkCompare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchmarkCompareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **a** | **string** | A is the first model id. It is required. | 
 **b** | **string** | B is the second model id. It is required. | 
 **benchmark** | **string** | Benchmark is the catalog id to compare on, defaulting to gpqa_diamond. | 

### Return type

[**Pairing**](Pairing.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkHistory

> HistoryOut GetBenchmarkHistory(ctx).Benchmark(benchmark).Model(model).Execute()

Returns each model's measured score per run over time, oldest first, with the change between runs.



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
	benchmark := "benchmark_example" // string | Benchmark is the catalog id to read, defaulting to gpqa_diamond. (optional)
	model := "model_example" // string | Model filters to one model. Empty returns every model measured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.GetBenchmarkHistory(context.Background()).Benchmark(benchmark).Model(model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.GetBenchmarkHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBenchmarkHistory`: HistoryOut
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.GetBenchmarkHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchmarkHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchmark** | **string** | Benchmark is the catalog id to read, defaulting to gpqa_diamond. | 
 **model** | **string** | Model filters to one model. Empty returns every model measured. | 

### Return type

[**HistoryOut**](HistoryOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkLeaderboard

> Leaderboard GetBenchmarkLeaderboard(ctx).Benchmark(benchmark).Execute()

Answers one row per model for the benchmark named — what our own harness measured, beside what the vendor claims, and the gap between them.



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
	benchmark := "benchmark_example" // string | Benchmark is the catalog id to read, defaulting to gpqa_diamond. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.GetBenchmarkLeaderboard(context.Background()).Benchmark(benchmark).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.GetBenchmarkLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBenchmarkLeaderboard`: Leaderboard
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.GetBenchmarkLeaderboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchmarkLeaderboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchmark** | **string** | Benchmark is the catalog id to read, defaulting to gpqa_diamond. | 

### Return type

[**Leaderboard**](Leaderboard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkPresets

> PresetList GetBenchmarkPresets(ctx).Execute()

Are the router blends available to compose from — a named set of model arms, the rank they escalate through and the panel width that bounds fan-out — each served by the model layer as enso-<name>.



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
	resp, r, err := apiClient.BenchmarkAPI.GetBenchmarkPresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.GetBenchmarkPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBenchmarkPresets`: PresetList
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.GetBenchmarkPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchmarkPresetsRequest struct via the builder pattern


### Return type

[**PresetList**](PresetList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBenchmarkClaims

> PutClaimsOut PostBenchmarkClaims(ctx).PutClaimsIn(putClaimsIn).Execute()

Records published claims: one to correct a number, many to import a leaderboard.



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
	putClaimsIn := *openapiclient.NewPutClaimsIn() // PutClaimsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.PostBenchmarkClaims(context.Background()).PutClaimsIn(putClaimsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.PostBenchmarkClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBenchmarkClaims`: PutClaimsOut
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.PostBenchmarkClaims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBenchmarkClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putClaimsIn** | [**PutClaimsIn**](PutClaimsIn.md) |  | 

### Return type

[**PutClaimsOut**](PutClaimsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBenchmarkPresets

> PresetAccepted PostBenchmarkPresets(ctx).Preset(preset).Execute()

Validates a router blend — its name, its arms, the rank they escalate through and the panel fan-out width — and answers 202 with the preset and the enso-<name> it would be served as.



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
	preset := *openapiclient.NewPreset() // Preset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.PostBenchmarkPresets(context.Background()).Preset(preset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.PostBenchmarkPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBenchmarkPresets`: PresetAccepted
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.PostBenchmarkPresets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBenchmarkPresetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preset** | [**Preset**](Preset.md) |  | 

### Return type

[**PresetAccepted**](PresetAccepted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBenchmarkRuns

> Admission PostBenchmarkRuns(ctx).Suite(suite).Execute()

Admits and queues a benchmark run against a model or your own endpoint, and answers 202 with the receipt.



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
	suite := *openapiclient.NewSuite([]string{"Benchmarks_example"}) // Suite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.PostBenchmarkRuns(context.Background()).Suite(suite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.PostBenchmarkRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBenchmarkRuns`: Admission
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.PostBenchmarkRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBenchmarkRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suite** | [**Suite**](Suite.md) |  | 

### Return type

[**Admission**](Admission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

