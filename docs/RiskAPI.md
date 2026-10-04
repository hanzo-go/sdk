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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> RiskRiskModelState RiskAdoptModel(ctx).RiskRiskAdoptIn(riskRiskAdoptIn).Execute()

Put one of your organisation's own published model values in force



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
	riskRiskAdoptIn := *openapiclient.NewRiskRiskAdoptIn() // RiskRiskAdoptIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskAdoptModel(context.Background()).RiskRiskAdoptIn(riskRiskAdoptIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskAdoptModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAdoptModel`: RiskRiskModelState
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskAdoptModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskAdoptModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskRiskAdoptIn** | [**RiskRiskAdoptIn**](RiskRiskAdoptIn.md) |  | 

### Return type

[**RiskRiskModelState**](RiskRiskModelState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskFeatures

> RiskRiskCatalog RiskFeatures(ctx).Days(days).Execute()

The feature catalogue: what the model reads, and what your surface carries



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
	days := int64(30) // int64 | Days is how far back to measure the organisation's own coverage, 1 to 400. Zero takes thirty. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskFeatures(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskFeatures`: RiskRiskCatalog
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int64** | Days is how far back to measure the organisation&#39;s own coverage, 1 to 400. Zero takes thirty. | 

### Return type

[**RiskRiskCatalog**](RiskRiskCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskLearn

> RiskRiskLearnOut RiskLearn(ctx).RiskRiskLearnIn(riskRiskLearnIn).Execute()

Teach your organisation's own model from its own events



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
	riskRiskLearnIn := *openapiclient.NewRiskRiskLearnIn() // RiskRiskLearnIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskLearn(context.Background()).RiskRiskLearnIn(riskRiskLearnIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskLearn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLearn`: RiskRiskLearnOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskLearn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskLearnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskRiskLearnIn** | [**RiskRiskLearnIn**](RiskRiskLearnIn.md) |  | 

### Return type

[**RiskRiskLearnOut**](RiskRiskLearnOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskPolicy

> RiskRiskPolicyOut RiskPolicy(ctx).Execute()

Your organisation's decision-regime history, and which version is in force



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
	resp, r, err := apiClient.RiskAPI.RiskPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskPolicy`: RiskRiskPolicyOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskPolicyRequest struct via the builder pattern


### Return type

[**RiskRiskPolicyOut**](RiskRiskPolicyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskPublishModel

> RiskRiskPublishOut RiskPublishModel(ctx).Execute()

Publish your organisation's model as a named, immutable value



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
	resp, r, err := apiClient.RiskAPI.RiskPublishModel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskPublishModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskPublishModel`: RiskRiskPublishOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskPublishModel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskPublishModelRequest struct via the builder pattern


### Return type

[**RiskRiskPublishOut**](RiskRiskPublishOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskScore

> RiskRiskScoreOut RiskScore(ctx).RiskRiskScoreIn(riskRiskScoreIn).Execute()

Score one event against your organisation's own model



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
	riskRiskScoreIn := *openapiclient.NewRiskRiskScoreIn() // RiskRiskScoreIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskScore(context.Background()).RiskRiskScoreIn(riskRiskScoreIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskScore`: RiskRiskScoreOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskRiskScoreIn** | [**RiskRiskScoreIn**](RiskRiskScoreIn.md) |  | 

### Return type

[**RiskRiskScoreOut**](RiskRiskScoreOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSearch

> RiskRiskSearchRun RiskSearch(ctx).RiskRiskSearchIn(riskRiskSearchIn).Execute()

Search exhaustively for the model shape that fits your own history



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
	riskRiskSearchIn := *openapiclient.NewRiskRiskSearchIn() // RiskRiskSearchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskSearch(context.Background()).RiskRiskSearchIn(riskRiskSearchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSearch`: RiskRiskSearchRun
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskRiskSearchIn** | [**RiskRiskSearchIn**](RiskRiskSearchIn.md) |  | 

### Return type

[**RiskRiskSearchRun**](RiskRiskSearchRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSearchResult

> RiskRiskSearchReport RiskSearchResult(ctx, id).Execute()

Read back one exhaustive search



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
	id := "srch_2f6a1c" // string | ID is the run, taken from the path. A run another organisation started is simply not there — the same answer an unknown id gives.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskSearchResult(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskSearchResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSearchResult`: RiskRiskSearchReport
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

[**RiskRiskSearchReport**](RiskRiskSearchReport.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSetPolicy

> RiskRiskPolicyOut RiskSetPolicy(ctx).RiskRiskAppetiteIn(riskRiskAppetiteIn).Execute()

State the decision regime: the appetite, the sample, and whether the model is live



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
	riskRiskAppetiteIn := *openapiclient.NewRiskRiskAppetiteIn() // RiskRiskAppetiteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskSetPolicy(context.Background()).RiskRiskAppetiteIn(riskRiskAppetiteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskSetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSetPolicy`: RiskRiskPolicyOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskSetPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskSetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskRiskAppetiteIn** | [**RiskRiskAppetiteIn**](RiskRiskAppetiteIn.md) |  | 

### Return type

[**RiskRiskPolicyOut**](RiskRiskPolicyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskState

> RiskRiskModelState RiskState(ctx).Execute()

Report your organisation's model: what it learned, and what it realised



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
	resp, r, err := apiClient.RiskAPI.RiskState(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskState`: RiskRiskModelState
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskStateRequest struct via the builder pattern


### Return type

[**RiskRiskModelState**](RiskRiskModelState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

