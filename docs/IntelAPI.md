# \IntelAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorldWorldModelChanges**](IntelAPI.md#WorldWorldModelChanges) | **Get** /v1/world/model/changes | World model — what changed since a timestamp (the inform-our-AI hook)
[**WorldWorldModelCountry**](IntelAPI.md#WorldWorldModelCountry) | **Get** /v1/world/model/country/{iso} | World model — one country&#39;s full state vector + recent deltas
[**WorldWorldModelState**](IntelAPI.md#WorldWorldModelState) | **Get** /v1/world/model/state | World model — full compact world-state snapshot (all entities, instability-ranked)
[**WorldWorldModelStream**](IntelAPI.md#WorldWorldModelStream) | **Get** /v1/world/model/stream | World model — SSE stream of state deltas as folds land
[**WorldWorldModelTop**](IntelAPI.md#WorldWorldModelTop) | **Get** /v1/world/model/top | World model — top entities by metric
[**WorldWorldPizzintDashboardData**](IntelAPI.md#WorldWorldPizzintDashboardData) | **Get** /v1/world/pizzint/dashboard-data | PIZZINT dashboard aggregate
[**WorldWorldPizzintGdeltBatch**](IntelAPI.md#WorldWorldPizzintGdeltBatch) | **Get** /v1/world/pizzint/gdelt/batch | PIZZINT GDELT batch
[**WorldWorldRiskScores**](IntelAPI.md#WorldWorldRiskScores) | **Get** /v1/world/risk-scores | Country risk scores
[**WorldWorldServiceStatus**](IntelAPI.md#WorldWorldServiceStatus) | **Get** /v1/world/service-status | Upstream data-source status
[**WorldWorldTemporalBaseline**](IntelAPI.md#WorldWorldTemporalBaseline) | **Get** /v1/world/temporal-baseline | Temporal baseline metrics
[**WorldWorldTheaterPosture**](IntelAPI.md#WorldWorldTheaterPosture) | **Get** /v1/world/theater-posture | Strategic theater posture



## WorldWorldModelChanges

> map[string]interface{} WorldWorldModelChanges(ctx).Since(since).Execute()

World model — what changed since a timestamp (the inform-our-AI hook)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	since := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelAPI.WorldWorldModelChanges(context.Background()).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldModelChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldModelChanges`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldModelChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldModelChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **time.Time** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldModelCountry

> map[string]interface{} WorldWorldModelCountry(ctx, iso).Execute()

World model — one country's full state vector + recent deltas

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
	iso := "iso_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelAPI.WorldWorldModelCountry(context.Background(), iso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldModelCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldModelCountry`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldModelCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iso** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldModelCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldModelState

> map[string]interface{} WorldWorldModelState(ctx).Kind(kind).Execute()

World model — full compact world-state snapshot (all entities, instability-ranked)

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
	kind := "kind_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelAPI.WorldWorldModelState(context.Background()).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldModelState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldModelState`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldModelState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldModelStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldModelStream

> string WorldWorldModelStream(ctx).Execute()

World model — SSE stream of state deltas as folds land

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
	resp, r, err := apiClient.IntelAPI.WorldWorldModelStream(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldModelStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldModelStream`: string
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldModelStream`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldModelStreamRequest struct via the builder pattern


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldModelTop

> map[string]interface{} WorldWorldModelTop(ctx).Metric(metric).Kind(kind).N(n).Execute()

World model — top entities by metric

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
	metric := "metric_example" // string |  (optional)
	kind := "kind_example" // string |  (optional)
	n := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelAPI.WorldWorldModelTop(context.Background()).Metric(metric).Kind(kind).N(n).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldModelTop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldModelTop`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldModelTop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldModelTopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metric** | **string** |  | 
 **kind** | **string** |  | 
 **n** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldPizzintDashboardData

> map[string]interface{} WorldWorldPizzintDashboardData(ctx).Execute()

PIZZINT dashboard aggregate

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
	resp, r, err := apiClient.IntelAPI.WorldWorldPizzintDashboardData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldPizzintDashboardData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldPizzintDashboardData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldPizzintDashboardData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldPizzintDashboardDataRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldPizzintGdeltBatch

> map[string]interface{} WorldWorldPizzintGdeltBatch(ctx).Execute()

PIZZINT GDELT batch

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
	resp, r, err := apiClient.IntelAPI.WorldWorldPizzintGdeltBatch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldPizzintGdeltBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldPizzintGdeltBatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldPizzintGdeltBatch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldPizzintGdeltBatchRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldRiskScores

> map[string]interface{} WorldWorldRiskScores(ctx).Execute()

Country risk scores

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
	resp, r, err := apiClient.IntelAPI.WorldWorldRiskScores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldRiskScores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldRiskScores`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldRiskScores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldRiskScoresRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldServiceStatus

> map[string]interface{} WorldWorldServiceStatus(ctx).Execute()

Upstream data-source status

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
	resp, r, err := apiClient.IntelAPI.WorldWorldServiceStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldServiceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldServiceStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldServiceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldServiceStatusRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldTemporalBaseline

> map[string]interface{} WorldWorldTemporalBaseline(ctx).Execute()

Temporal baseline metrics

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
	resp, r, err := apiClient.IntelAPI.WorldWorldTemporalBaseline(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldTemporalBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldTemporalBaseline`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldTemporalBaseline`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldTemporalBaselineRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldTheaterPosture

> map[string]interface{} WorldWorldTheaterPosture(ctx).Execute()

Strategic theater posture

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
	resp, r, err := apiClient.IntelAPI.WorldWorldTheaterPosture(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelAPI.WorldWorldTheaterPosture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldTheaterPosture`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntelAPI.WorldWorldTheaterPosture`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldTheaterPostureRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

