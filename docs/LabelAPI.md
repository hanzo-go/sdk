# \LabelAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RiskDisposeLabels**](LabelAPI.md#RiskDisposeLabels) | **Post** /v1/label/dispose | Dispose of this tenant&#39;s expired assertions, whole records only
[**RiskHoldLabels**](LabelAPI.md#RiskHoldLabels) | **Post** /v1/label/hold | Place or release a litigation hold on named records
[**RiskLabel**](LabelAPI.md#RiskLabel) | **Post** /v1/label | Assert ground truth about events
[**RiskLabelCoverage**](LabelAPI.md#RiskLabelCoverage) | **Get** /v1/label/coverage | How much of the window has matured, and how much of that is judged
[**RiskLabelVocabulary**](LabelAPI.md#RiskLabelVocabulary) | **Get** /v1/label/vocabulary | The closed vocabularies and the precedence rule that resolves a conflict
[**RiskLabels**](LabelAPI.md#RiskLabels) | **Get** /v1/label | Read the assertions this tenant has recorded
[**RiskResolveLabels**](LabelAPI.md#RiskResolveLabels) | **Post** /v1/label/resolve | Resolve the label in force for named events, as of each event&#39;s own horizon



## RiskDisposeLabels

> RiskDisposeOut RiskDisposeLabels(ctx).RiskDisposeIn(riskDisposeIn).Execute()

Dispose of this tenant's expired assertions, whole records only



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
	riskDisposeIn := *openapiclient.NewRiskDisposeIn() // RiskDisposeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.RiskDisposeLabels(context.Background()).RiskDisposeIn(riskDisposeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.RiskDisposeLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDisposeLabels`: RiskDisposeOut
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.RiskDisposeLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskDisposeLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskDisposeIn** | [**RiskDisposeIn**](RiskDisposeIn.md) |  | 

### Return type

[**RiskDisposeOut**](RiskDisposeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskHoldLabels

> RiskHoldOut RiskHoldLabels(ctx).RiskHoldIn(riskHoldIn).Execute()

Place or release a litigation hold on named records



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
	riskHoldIn := *openapiclient.NewRiskHoldIn() // RiskHoldIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.RiskHoldLabels(context.Background()).RiskHoldIn(riskHoldIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.RiskHoldLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskHoldLabels`: RiskHoldOut
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.RiskHoldLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskHoldLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskHoldIn** | [**RiskHoldIn**](RiskHoldIn.md) |  | 

### Return type

[**RiskHoldOut**](RiskHoldOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskLabel

> RiskLabelOut RiskLabel(ctx).RiskLabelIn(riskLabelIn).Execute()

Assert ground truth about events



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
	riskLabelIn := *openapiclient.NewRiskLabelIn() // RiskLabelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.RiskLabel(context.Background()).RiskLabelIn(riskLabelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.RiskLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLabel`: RiskLabelOut
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.RiskLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskLabelIn** | [**RiskLabelIn**](RiskLabelIn.md) |  | 

### Return type

[**RiskLabelOut**](RiskLabelOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskLabelCoverage

> RiskLabelCoverage RiskLabelCoverage(ctx).From(from).To(to).Horizon(horizon).Execute()

How much of the window has matured, and how much of that is judged



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
	from := "from_example" // string | From and To bound the EVENT window, half-open, RFC 3339.  Unstated, the window is the 90 days ENDING where maturity begins — `to` is the horizon ago, not now. A default window running to now under a default horizon could not contain one matured event, so every count below it would be zero however much ground truth the tenant held. (optional)
	to := "to_example" // string |  (optional)
	horizon := int32(56) // int32 | Horizon is the maturity horizon in days the coverage is measured under. Unstated takes 120. It also moves the default window, which ends where maturity begins. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.RiskLabelCoverage(context.Background()).From(from).To(to).Horizon(horizon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.RiskLabelCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLabelCoverage`: RiskLabelCoverage
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.RiskLabelCoverage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskLabelCoverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | From and To bound the EVENT window, half-open, RFC 3339.  Unstated, the window is the 90 days ENDING where maturity begins — &#x60;to&#x60; is the horizon ago, not now. A default window running to now under a default horizon could not contain one matured event, so every count below it would be zero however much ground truth the tenant held. | 
 **to** | **string** |  | 

### Return type

[**RiskLabelCoverage**](RiskLabelCoverage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskLabelVocabulary

> RiskLabelVocabulary RiskLabelVocabulary(ctx).Execute()

The closed vocabularies and the precedence rule that resolves a conflict



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
	resp, r, err := apiClient.LabelAPI.RiskLabelVocabulary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.RiskLabelVocabulary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLabelVocabulary`: RiskLabelVocabulary
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.RiskLabelVocabulary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskLabelVocabularyRequest struct via the builder pattern


### Return type

[**RiskLabelVocabulary**](RiskLabelVocabulary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskLabels

> RiskLabelsOut RiskLabels(ctx).Kind(kind).Subject(subject).Source(source).From(from).To(to).Limit(limit).Execute()

Read the assertions this tenant has recorded



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
	kind := "kind_example" // string | Kind and Subject narrow to one entity. (optional)
	subject := "subject_example" // string |  (optional)
	source := "source_example" // string | Source narrows to one asserter — the read that answers \"what has commerce told us\", separately from \"what has an analyst told us\". (optional)
	from := "from_example" // string | From and To bound the EVENT time, half-open, RFC 3339. (optional)
	to := "to_example" // string |  (optional)
	limit := int32(56) // int32 | Limit caps the page. Out of range takes the plane's own bound. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.RiskLabels(context.Background()).Kind(kind).Subject(subject).Source(source).From(from).To(to).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.RiskLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLabels`: RiskLabelsOut
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.RiskLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Kind and Subject narrow to one entity. | 
 **subject** | **string** |  | 
 **source** | **string** | Source narrows to one asserter — the read that answers \&quot;what has commerce told us\&quot;, separately from \&quot;what has an analyst told us\&quot;. | 
 **from** | **string** | From and To bound the EVENT time, half-open, RFC 3339. | 
 **to** | **string** |  | 
 **limit** | **int32** | Limit caps the page. Out of range takes the plane&#39;s own bound. | 

### Return type

[**RiskLabelsOut**](RiskLabelsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskResolveLabels

> RiskResolveOut RiskResolveLabels(ctx).RiskResolveIn(riskResolveIn).Execute()

Resolve the label in force for named events, as of each event's own horizon



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
	riskResolveIn := *openapiclient.NewRiskResolveIn() // RiskResolveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelAPI.RiskResolveLabels(context.Background()).RiskResolveIn(riskResolveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelAPI.RiskResolveLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskResolveLabels`: RiskResolveOut
	fmt.Fprintf(os.Stdout, "Response from `LabelAPI.RiskResolveLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskResolveLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskResolveIn** | [**RiskResolveIn**](RiskResolveIn.md) |  | 

### Return type

[**RiskResolveOut**](RiskResolveOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

