# \BenchmarkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBenchmarkCatalog**](BenchmarkAPI.md#GetBenchmarkCatalog) | **Get** /v1/benchmark/catalog | Is the canonical public benchmarks this arena runs — the id, title, axis, item count and upstream source of each, with native marking the ones the standardized harness runs today; the rest are registered and adapter-pending.
[**GetBenchmarkClaims**](BenchmarkAPI.md#GetBenchmarkClaims) | **Get** /v1/benchmark/claims | Lists the effective claims the caller may read, each labelled with the org that made it and the user who recorded it.
[**GetBenchmarkCompare**](BenchmarkAPI.md#GetBenchmarkCompare) | **Get** /v1/benchmark/compare | Is the ONLY valid arm-vs-arm test: it pairs the two models on the items BOTH completed, and answers rescue and damage counts with an exact-McNemar p.
[**GetBenchmarkHistory**](BenchmarkAPI.md#GetBenchmarkHistory) | **Get** /v1/benchmark/history | Returns each model&#39;s measured score per run over time, oldest first, with the change between runs.
[**GetBenchmarkLeaderboard**](BenchmarkAPI.md#GetBenchmarkLeaderboard) | **Get** /v1/benchmark/leaderboard | Answers one row per model for the benchmark named — what our own harness measured, beside what the vendor claims, and the gap between them.
[**GetBenchmarkPresets**](BenchmarkAPI.md#GetBenchmarkPresets) | **Get** /v1/benchmark/presets | Are the router blends available to compose from — a named set of model arms, the rank they escalate through and the panel width that bounds fan-out — each served by the model layer as enso-&lt;name&gt;.
[**PostBenchmarkClaims**](BenchmarkAPI.md#PostBenchmarkClaims) | **Post** /v1/benchmark/claims | Records claims for the caller&#39;s org: one to correct a number, many to import a leaderboard.
[**PostBenchmarkPresets**](BenchmarkAPI.md#PostBenchmarkPresets) | **Post** /v1/benchmark/presets | Validates a router blend — its name, its arms, the rank they escalate through and the panel fan-out width — and answers 202 with the preset and the enso-&lt;name&gt; it would be served as.
[**PostBenchmarkRuns**](BenchmarkAPI.md#PostBenchmarkRuns) | **Post** /v1/benchmark/runs | Admits and queues a benchmark run against a model or your own endpoint, and answers 202 with the receipt.



## GetBenchmarkCatalog

> BenchmarkBenchmarkCatalog GetBenchmarkCatalog(ctx).Execute()

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
	// response from `GetBenchmarkCatalog`: BenchmarkBenchmarkCatalog
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.GetBenchmarkCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchmarkCatalogRequest struct via the builder pattern


### Return type

[**BenchmarkBenchmarkCatalog**](BenchmarkBenchmarkCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkClaims

> BenchmarkClaimsOut GetBenchmarkClaims(ctx).Benchmark(benchmark).Model(model).Provider(provider).Source(source).Protocol(protocol).Org(org).Execute()

Lists the effective claims the caller may read, each labelled with the org that made it and the user who recorded it.



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
	org := "org_example" // string | Org filters to the claims one org made; \"admin\" reads the platform's own. It narrows what the caller may already read and never widens it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.GetBenchmarkClaims(context.Background()).Benchmark(benchmark).Model(model).Provider(provider).Source(source).Protocol(protocol).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.GetBenchmarkClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBenchmarkClaims`: BenchmarkClaimsOut
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
 **org** | **string** | Org filters to the claims one org made; \&quot;admin\&quot; reads the platform&#39;s own. It narrows what the caller may already read and never widens it. | 

### Return type

[**BenchmarkClaimsOut**](BenchmarkClaimsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkCompare

> BenchmarkPairing GetBenchmarkCompare(ctx).A(a).B(b).Benchmark(benchmark).Execute()

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
	// response from `GetBenchmarkCompare`: BenchmarkPairing
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

[**BenchmarkPairing**](BenchmarkPairing.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkHistory

> BenchmarkHistoryOut GetBenchmarkHistory(ctx).Benchmark(benchmark).Model(model).Execute()

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
	// response from `GetBenchmarkHistory`: BenchmarkHistoryOut
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

[**BenchmarkHistoryOut**](BenchmarkHistoryOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkLeaderboard

> BenchmarkLeaderboard GetBenchmarkLeaderboard(ctx).Benchmark(benchmark).Execute()

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
	// response from `GetBenchmarkLeaderboard`: BenchmarkLeaderboard
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

[**BenchmarkLeaderboard**](BenchmarkLeaderboard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchmarkPresets

> BenchmarkPresetList GetBenchmarkPresets(ctx).Execute()

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
	// response from `GetBenchmarkPresets`: BenchmarkPresetList
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.GetBenchmarkPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchmarkPresetsRequest struct via the builder pattern


### Return type

[**BenchmarkPresetList**](BenchmarkPresetList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBenchmarkClaims

> BenchmarkPutClaimsOut PostBenchmarkClaims(ctx).BenchmarkPutClaimsIn(benchmarkPutClaimsIn).Execute()

Records claims for the caller's org: one to correct a number, many to import a leaderboard.



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
	benchmarkPutClaimsIn := *openapiclient.NewBenchmarkPutClaimsIn() // BenchmarkPutClaimsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.PostBenchmarkClaims(context.Background()).BenchmarkPutClaimsIn(benchmarkPutClaimsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.PostBenchmarkClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBenchmarkClaims`: BenchmarkPutClaimsOut
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.PostBenchmarkClaims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBenchmarkClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchmarkPutClaimsIn** | [**BenchmarkPutClaimsIn**](BenchmarkPutClaimsIn.md) |  | 

### Return type

[**BenchmarkPutClaimsOut**](BenchmarkPutClaimsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBenchmarkPresets

> BenchmarkPresetAccepted PostBenchmarkPresets(ctx).BenchmarkPreset(benchmarkPreset).Execute()

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
	benchmarkPreset := *openapiclient.NewBenchmarkPreset() // BenchmarkPreset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.PostBenchmarkPresets(context.Background()).BenchmarkPreset(benchmarkPreset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.PostBenchmarkPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBenchmarkPresets`: BenchmarkPresetAccepted
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.PostBenchmarkPresets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBenchmarkPresetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchmarkPreset** | [**BenchmarkPreset**](BenchmarkPreset.md) |  | 

### Return type

[**BenchmarkPresetAccepted**](BenchmarkPresetAccepted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBenchmarkRuns

> BenchmarkAdmission PostBenchmarkRuns(ctx).BenchmarkSuite(benchmarkSuite).Execute()

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
	benchmarkSuite := *openapiclient.NewBenchmarkSuite([]string{"Benchmarks_example"}) // BenchmarkSuite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BenchmarkAPI.PostBenchmarkRuns(context.Background()).BenchmarkSuite(benchmarkSuite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BenchmarkAPI.PostBenchmarkRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBenchmarkRuns`: BenchmarkAdmission
	fmt.Fprintf(os.Stdout, "Response from `BenchmarkAPI.PostBenchmarkRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBenchmarkRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchmarkSuite** | [**BenchmarkSuite**](BenchmarkSuite.md) |  | 

### Return type

[**BenchmarkAdmission**](BenchmarkAdmission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

