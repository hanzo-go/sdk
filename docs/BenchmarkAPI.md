# \BenchmarkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBenchmarkCatalog**](BenchmarkAPI.md#GetBenchmarkCatalog) | **Get** /v1/benchmark/catalog | Is the canonical public benchmarks this arena runs — the id, title, axis, item count and upstream source of each, with native marking the ones the standardized harness runs today; the rest are registered and adapter-pending.
[**GetBenchmarkCompare**](BenchmarkAPI.md#GetBenchmarkCompare) | **Get** /v1/benchmark/compare | Is the ONLY valid arm-vs-arm test: it pairs the two models on the items BOTH completed, and answers rescue and damage counts with an exact-McNemar p.
[**GetBenchmarkLeaderboard**](BenchmarkAPI.md#GetBenchmarkLeaderboard) | **Get** /v1/benchmark/leaderboard | Answers one row per model for the benchmark named — what our own harness measured, beside what the vendor claims, and the gap between them.
[**GetBenchmarkPresets**](BenchmarkAPI.md#GetBenchmarkPresets) | **Get** /v1/benchmark/presets | Are the router blends available to compose from — a named set of model arms, the rank they escalate through and the panel width that bounds fan-out — each served by the model layer as enso-&lt;name&gt;.
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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

