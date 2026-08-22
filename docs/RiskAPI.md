# \RiskAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRiskHealth**](RiskAPI.md#GetRiskHealth) | **Get** /v1/risk/health | Whether the risk model plane can actually work right now
[**RiskAdoptModel**](RiskAPI.md#RiskAdoptModel) | **Put** /v1/risk/state/model | Put one of your organisation&#39;s own published model values in force
[**RiskFeatures**](RiskAPI.md#RiskFeatures) | **Get** /v1/risk/features | The feature catalogue: what the model reads, and what your surface carries
[**RiskLearn**](RiskAPI.md#RiskLearn) | **Post** /v1/risk/learn | Teach your organisation&#39;s own model from its own events
[**RiskPolicy**](RiskAPI.md#RiskPolicy) | **Get** /v1/risk/policy | Your organisation&#39;s decision-regime history, and which version is in force
[**RiskPublishModel**](RiskAPI.md#RiskPublishModel) | **Post** /v1/risk/state/model | Publish your organisation&#39;s model as a named, immutable value
[**RiskScore**](RiskAPI.md#RiskScore) | **Post** /v1/risk/score | Score one event against your organisation&#39;s own model
[**RiskSearch**](RiskAPI.md#RiskSearch) | **Post** /v1/risk/search | Search exhaustively for the model shape that fits your own history
[**RiskSearchResult**](RiskAPI.md#RiskSearchResult) | **Get** /v1/risk/search/{id} | Read back one exhaustive search
[**RiskSetPolicy**](RiskAPI.md#RiskSetPolicy) | **Put** /v1/risk/policy | State the decision regime: the appetite, the sample, and whether the model is live
[**RiskState**](RiskAPI.md#RiskState) | **Get** /v1/risk/state | Report your organisation&#39;s model: what it learned, and what it realised



## GetRiskHealth

> GetRiskHealth(ctx).Execute()

Whether the risk model plane can actually work right now



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
	r, err := apiClient.RiskAPI.GetRiskHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.GetRiskHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskHealthRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskAdoptModel

> RiskModelState RiskAdoptModel(ctx).RiskAdoptIn(riskAdoptIn).Execute()

Put one of your organisation's own published model values in force



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
	riskAdoptIn := *openapiclient.NewRiskAdoptIn() // RiskAdoptIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskAdoptModel(context.Background()).RiskAdoptIn(riskAdoptIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskAdoptModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAdoptModel`: RiskModelState
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskAdoptModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskAdoptModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskAdoptIn** | [**RiskAdoptIn**](RiskAdoptIn.md) |  | 

### Return type

[**RiskModelState**](RiskModelState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskFeatures

> RiskCatalog RiskFeatures(ctx).Days(days).Execute()

The feature catalogue: what the model reads, and what your surface carries



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
	days := int32(30) // int32 | Days is how far back to measure the organisation's own coverage, 1 to 400. Zero takes thirty. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskFeatures(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskFeatures`: RiskCatalog
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | Days is how far back to measure the organisation&#39;s own coverage, 1 to 400. Zero takes thirty. | 

### Return type

[**RiskCatalog**](RiskCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskLearn

> RiskLearnOut RiskLearn(ctx).RiskLearnIn(riskLearnIn).Execute()

Teach your organisation's own model from its own events



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
	riskLearnIn := *openapiclient.NewRiskLearnIn() // RiskLearnIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskLearn(context.Background()).RiskLearnIn(riskLearnIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskLearn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLearn`: RiskLearnOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskLearn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskLearnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskLearnIn** | [**RiskLearnIn**](RiskLearnIn.md) |  | 

### Return type

[**RiskLearnOut**](RiskLearnOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskPolicy

> RiskPolicyOut RiskPolicy(ctx).Execute()

Your organisation's decision-regime history, and which version is in force



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
	resp, r, err := apiClient.RiskAPI.RiskPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskPolicy`: RiskPolicyOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskPolicyRequest struct via the builder pattern


### Return type

[**RiskPolicyOut**](RiskPolicyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskPublishModel

> RiskPublishOut RiskPublishModel(ctx).Execute()

Publish your organisation's model as a named, immutable value



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
	resp, r, err := apiClient.RiskAPI.RiskPublishModel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskPublishModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskPublishModel`: RiskPublishOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskPublishModel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskPublishModelRequest struct via the builder pattern


### Return type

[**RiskPublishOut**](RiskPublishOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskScore

> RiskScoreOut RiskScore(ctx).RiskScoreIn(riskScoreIn).Execute()

Score one event against your organisation's own model



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
	riskScoreIn := *openapiclient.NewRiskScoreIn() // RiskScoreIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskScore(context.Background()).RiskScoreIn(riskScoreIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskScore`: RiskScoreOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskScoreIn** | [**RiskScoreIn**](RiskScoreIn.md) |  | 

### Return type

[**RiskScoreOut**](RiskScoreOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSearch

> RiskSearchRun RiskSearch(ctx).RiskSearchIn(riskSearchIn).Execute()

Search exhaustively for the model shape that fits your own history



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
	riskSearchIn := *openapiclient.NewRiskSearchIn() // RiskSearchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskSearch(context.Background()).RiskSearchIn(riskSearchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSearch`: RiskSearchRun
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskSearchIn** | [**RiskSearchIn**](RiskSearchIn.md) |  | 

### Return type

[**RiskSearchRun**](RiskSearchRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSearchResult

> RiskSearchReport RiskSearchResult(ctx, id).Execute()

Read back one exhaustive search



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
	id := "srch_2f6a1c" // string | ID is the run, taken from the path. A run another organisation started is simply not there — the same answer an unknown id gives.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskSearchResult(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskSearchResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSearchResult`: RiskSearchReport
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskSearchResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the run, taken from the path. A run another organisation started is simply not there — the same answer an unknown id gives. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskSearchResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskSearchReport**](RiskSearchReport.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSetPolicy

> RiskPolicyOut RiskSetPolicy(ctx).RiskAppetiteIn(riskAppetiteIn).Execute()

State the decision regime: the appetite, the sample, and whether the model is live



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
	riskAppetiteIn := *openapiclient.NewRiskAppetiteIn() // RiskAppetiteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskSetPolicy(context.Background()).RiskAppetiteIn(riskAppetiteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskSetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSetPolicy`: RiskPolicyOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskSetPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskSetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskAppetiteIn** | [**RiskAppetiteIn**](RiskAppetiteIn.md) |  | 

### Return type

[**RiskPolicyOut**](RiskPolicyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskState

> RiskModelState RiskState(ctx).Execute()

Report your organisation's model: what it learned, and what it realised



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
	resp, r, err := apiClient.RiskAPI.RiskState(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskState`: RiskModelState
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskStateRequest struct via the builder pattern


### Return type

[**RiskModelState**](RiskModelState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

