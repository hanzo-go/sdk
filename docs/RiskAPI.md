# \RiskAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRiskHealth**](RiskAPI.md#GetRiskHealth) | **Get** /v1/risk/health | Whether the risk model plane can actually work right now
[**RiskAdoptModel**](RiskAPI.md#RiskAdoptModel) | **Put** /v1/risk/state/model | Put one of your organisation&#39;s own published model values in force
[**RiskClearReference**](RiskAPI.md#RiskClearReference) | **Delete** /v1/risk/reference/{set} | Removes one of your organisation&#39;s overrides.
[**RiskCreateDataset**](RiskAPI.md#RiskCreateDataset) | **Post** /v1/risk/datasets | Declare the next version of a dataset
[**RiskDataset**](RiskAPI.md#RiskDataset) | **Get** /v1/risk/datasets/{name} | Describe every version of one dataset
[**RiskDatasetLineage**](RiskAPI.md#RiskDatasetLineage) | **Get** /v1/risk/datasets/{name}/lineage | Show where a version&#39;s rows came from, and whether that can still be demonstrated
[**RiskDatasets**](RiskAPI.md#RiskDatasets) | **Get** /v1/risk/datasets | List this org&#39;s datasets
[**RiskDeleteDataset**](RiskAPI.md#RiskDeleteDataset) | **Delete** /v1/risk/datasets/{name} | Dispose of one dataset and every version of it
[**RiskDisposeLabels**](RiskAPI.md#RiskDisposeLabels) | **Post** /v1/risk/labels/dispose | Dispose of this tenant&#39;s expired assertions, whole records only
[**RiskExportDataset**](RiskAPI.md#RiskExportDataset) | **Get** /v1/risk/datasets/{name}/export | Read a version&#39;s rows back, one page at a time
[**RiskFeatures**](RiskAPI.md#RiskFeatures) | **Get** /v1/risk/features | The feature catalogue: what the model reads, and what your surface carries
[**RiskHoldLabels**](RiskAPI.md#RiskHoldLabels) | **Post** /v1/risk/labels/hold | Place or release a litigation hold on named records
[**RiskLabel**](RiskAPI.md#RiskLabel) | **Post** /v1/risk/labels | Assert ground truth about events
[**RiskLabelCoverage**](RiskAPI.md#RiskLabelCoverage) | **Get** /v1/risk/labels/coverage | How much of the window has matured, and how much of that is judged
[**RiskLabelVocabulary**](RiskAPI.md#RiskLabelVocabulary) | **Get** /v1/risk/labels/vocabulary | The closed vocabularies and the precedence rule that resolves a conflict
[**RiskLabels**](RiskAPI.md#RiskLabels) | **Get** /v1/risk/labels | Read the assertions this tenant has recorded
[**RiskLearn**](RiskAPI.md#RiskLearn) | **Post** /v1/risk/learn | Teach your organisation&#39;s own model from its own events
[**RiskMaterializeDataset**](RiskAPI.md#RiskMaterializeDataset) | **Post** /v1/risk/datasets/{name}/materialize | Materialise the declared version into immutable rows
[**RiskPolicy**](RiskAPI.md#RiskPolicy) | **Get** /v1/risk/policy | Your organisation&#39;s decision-regime history, and which version is in force
[**RiskPublishModel**](RiskAPI.md#RiskPublishModel) | **Post** /v1/risk/state/model | Publish your organisation&#39;s model as a named, immutable value
[**RiskReference**](RiskAPI.md#RiskReference) | **Get** /v1/risk/reference/{set} | Reference describes one set and lists your org&#39;s overrides in it.
[**RiskReferenceSets**](RiskAPI.md#RiskReferenceSets) | **Get** /v1/risk/reference | Lists every set this plane publishes, with its version and how fresh it is.
[**RiskRefreshReference**](RiskAPI.md#RiskRefreshReference) | **Post** /v1/risk/reference/refresh | Takes a new version of one set.
[**RiskResolveLabels**](RiskAPI.md#RiskResolveLabels) | **Post** /v1/risk/labels/resolve | Resolve the label in force for named events, as of each event&#39;s own horizon
[**RiskResolveReference**](RiskAPI.md#RiskResolveReference) | **Post** /v1/risk/reference/resolve | Looks keys up against the reference plane.
[**RiskScore**](RiskAPI.md#RiskScore) | **Post** /v1/risk/score | Score one event against your organisation&#39;s own model
[**RiskSearch**](RiskAPI.md#RiskSearch) | **Post** /v1/risk/search | Search exhaustively for the model shape that fits your own history
[**RiskSearchResult**](RiskAPI.md#RiskSearchResult) | **Get** /v1/risk/search/{id} | Read back one exhaustive search
[**RiskSetPolicy**](RiskAPI.md#RiskSetPolicy) | **Put** /v1/risk/policy | State the decision regime: the appetite, the sample, and whether the model is live
[**RiskSetReference**](RiskAPI.md#RiskSetReference) | **Put** /v1/risk/reference/{set} | Writes your organisation&#39;s own allow and deny entries over a set.
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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskClearReference

> ClearReferenceOut RiskClearReference(ctx, set).Key(key).Execute()

Removes one of your organisation's overrides.



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
	set := "domain" // string | 
	key := "partner.example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskClearReference(context.Background(), set).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskClearReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskClearReference`: ClearReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskClearReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskClearReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 

### Return type

[**ClearReferenceOut**](ClearReferenceOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskCreateDataset

> RiskDataset RiskCreateDataset(ctx).RiskDatasetSpec(riskDatasetSpec).Execute()

Declare the next version of a dataset



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
	riskDatasetSpec := *openapiclient.NewRiskDatasetSpec() // RiskDatasetSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskCreateDataset(context.Background()).RiskDatasetSpec(riskDatasetSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskCreateDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskCreateDataset`: RiskDataset
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskCreateDataset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskCreateDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskDatasetSpec** | [**RiskDatasetSpec**](RiskDatasetSpec.md) |  | 

### Return type

[**RiskDataset**](RiskDataset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskDataset

> RiskDatasetVersions RiskDataset(ctx, name).Execute()

Describe every version of one dataset



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
	name := "signups" // string | Name is the dataset, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskDataset(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDataset`: RiskDatasetVersions
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskDataset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskDatasetVersions**](RiskDatasetVersions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskDatasetLineage

> RiskLineage RiskDatasetLineage(ctx, name).Version(version).Execute()

Show where a version's rows came from, and whether that can still be demonstrated



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
	name := "signups" // string | Name is the dataset, from the path.
	version := int32(1) // int32 | Version is the version to trace. Zero takes the newest published one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskDatasetLineage(context.Background(), name).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskDatasetLineage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDatasetLineage`: RiskLineage
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskDatasetLineage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskDatasetLineageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | Version is the version to trace. Zero takes the newest published one. | 

### Return type

[**RiskLineage**](RiskLineage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskDatasets

> RiskDatasetList RiskDatasets(ctx).Execute()

List this org's datasets



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
	resp, r, err := apiClient.RiskAPI.RiskDatasets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDatasets`: RiskDatasetList
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskDatasets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskDatasetsRequest struct via the builder pattern


### Return type

[**RiskDatasetList**](RiskDatasetList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskDeleteDataset

> RiskDatasetDisposal RiskDeleteDataset(ctx, name).Execute()

Dispose of one dataset and every version of it



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
	name := "signups" // string | Name is the dataset, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskDeleteDataset(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskDeleteDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDeleteDataset`: RiskDatasetDisposal
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskDeleteDataset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskDeleteDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskDatasetDisposal**](RiskDatasetDisposal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	riskDisposeIn := *openapiclient.NewRiskDisposeIn() // RiskDisposeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskDisposeLabels(context.Background()).RiskDisposeIn(riskDisposeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskDisposeLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDisposeLabels`: RiskDisposeOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskDisposeLabels`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskExportDataset

> RiskDatasetRows RiskExportDataset(ctx, name).Version(version).Split(split).Offset(offset).Limit(limit).Execute()

Read a version's rows back, one page at a time



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
	name := "signups" // string | Name is the dataset, from the path.
	version := int32(1) // int32 | Version is the version to read. Zero takes the newest published one. (optional)
	split := "train" // string | Split narrows to train, val or test. Empty reads every split. (optional)
	offset := int32(56) // int32 | Offset is where the page starts, in the version's own row order (by id, which is derived from the row and therefore stable forever). (optional)
	limit := int32(500) // int32 | Limit is how many rows to return. Zero and anything above the plane's bound take the bound. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskExportDataset(context.Background(), name).Version(version).Split(split).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskExportDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskExportDataset`: RiskDatasetRows
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskExportDataset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskExportDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | Version is the version to read. Zero takes the newest published one. | 
 **split** | **string** | Split narrows to train, val or test. Empty reads every split. | 
 **offset** | **int32** | Offset is where the page starts, in the version&#39;s own row order (by id, which is derived from the row and therefore stable forever). | 
 **limit** | **int32** | Limit is how many rows to return. Zero and anything above the plane&#39;s bound take the bound. | 

### Return type

[**RiskDatasetRows**](RiskDatasetRows.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	riskHoldIn := *openapiclient.NewRiskHoldIn() // RiskHoldIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskHoldLabels(context.Background()).RiskHoldIn(riskHoldIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskHoldLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskHoldLabels`: RiskHoldOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskHoldLabels`: %v\n", resp)
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	riskLabelIn := *openapiclient.NewRiskLabelIn() // RiskLabelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskLabel(context.Background()).RiskLabelIn(riskLabelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLabel`: RiskLabelOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskLabel`: %v\n", resp)
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	from := "from_example" // string | From and To bound the EVENT window, half-open, RFC 3339.  Unstated, the window is the 90 days ENDING where maturity begins — `to` is the horizon ago, not now. A default window running to now under a default horizon could not contain one matured event, so every count below it would be zero however much ground truth the tenant held. (optional)
	to := "to_example" // string |  (optional)
	horizon := int32(56) // int32 | Horizon is the maturity horizon in days the coverage is measured under. Unstated takes 120. It also moves the default window, which ends where maturity begins. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskLabelCoverage(context.Background()).From(from).To(to).Horizon(horizon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskLabelCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLabelCoverage`: RiskLabelCoverage
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskLabelCoverage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskLabelCoverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | From and To bound the EVENT window, half-open, RFC 3339.  Unstated, the window is the 90 days ENDING where maturity begins — &#x60;to&#x60; is the horizon ago, not now. A default window running to now under a default horizon could not contain one matured event, so every count below it would be zero however much ground truth the tenant held. | 
 **to** | **string** |  | 
 **horizon** | **int32** | Horizon is the maturity horizon in days the coverage is measured under. Unstated takes 120. It also moves the default window, which ends where maturity begins. | 

### Return type

[**RiskLabelCoverage**](RiskLabelCoverage.md)

### Authorization

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskLabelVocabulary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskLabelVocabulary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLabelVocabulary`: RiskLabelVocabulary
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskLabelVocabulary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskLabelVocabularyRequest struct via the builder pattern


### Return type

[**RiskLabelVocabulary**](RiskLabelVocabulary.md)

### Authorization

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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
	resp, r, err := apiClient.RiskAPI.RiskLabels(context.Background()).Kind(kind).Subject(subject).Source(source).From(from).To(to).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskLabels`: RiskLabelsOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskLabels`: %v\n", resp)
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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskMaterializeDataset

> RiskDataset RiskMaterializeDataset(ctx, name).Execute()

Materialise the declared version into immutable rows



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
	name := "signups" // string | Name is the dataset, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskMaterializeDataset(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskMaterializeDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskMaterializeDataset`: RiskDataset
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskMaterializeDataset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the dataset, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskMaterializeDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskDataset**](RiskDataset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskReference

> ReferenceOut RiskReference(ctx, set).After(after).Limit(limit).Execute()

Reference describes one set and lists your org's overrides in it.



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
	set := "domain" // string | 
	after := "after_example" // string | After pages the override listing: the last key of the previous page. (optional)
	limit := int32(50) // int32 | Limit caps the override listing: default 200, maximum 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskReference(context.Background(), set).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskReference`: ReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | After pages the override listing: the last key of the previous page. | 
 **limit** | **int32** | Limit caps the override listing: default 200, maximum 1000. | 

### Return type

[**ReferenceOut**](ReferenceOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskReferenceSets

> ReferenceSetsOut RiskReferenceSets(ctx).Execute()

Lists every set this plane publishes, with its version and how fresh it is.



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
	resp, r, err := apiClient.RiskAPI.RiskReferenceSets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskReferenceSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskReferenceSets`: ReferenceSetsOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskReferenceSets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskReferenceSetsRequest struct via the builder pattern


### Return type

[**ReferenceSetsOut**](ReferenceSetsOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskRefreshReference

> RefreshReferenceOut RiskRefreshReference(ctx).RefreshReferenceIn(refreshReferenceIn).Execute()

Takes a new version of one set.



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
	refreshReferenceIn := *openapiclient.NewRefreshReferenceIn() // RefreshReferenceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskRefreshReference(context.Background()).RefreshReferenceIn(refreshReferenceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskRefreshReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskRefreshReference`: RefreshReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskRefreshReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskRefreshReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshReferenceIn** | [**RefreshReferenceIn**](RefreshReferenceIn.md) |  | 

### Return type

[**RefreshReferenceOut**](RefreshReferenceOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	riskResolveIn := *openapiclient.NewRiskResolveIn() // RiskResolveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskResolveLabels(context.Background()).RiskResolveIn(riskResolveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskResolveLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskResolveLabels`: RiskResolveOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskResolveLabels`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskResolveReference

> ResolveReferenceOut RiskResolveReference(ctx).ResolveReferenceIn(resolveReferenceIn).Execute()

Looks keys up against the reference plane.



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
	resolveReferenceIn := *openapiclient.NewResolveReferenceIn() // ResolveReferenceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskResolveReference(context.Background()).ResolveReferenceIn(resolveReferenceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskResolveReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskResolveReference`: ResolveReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskResolveReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskResolveReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resolveReferenceIn** | [**ResolveReferenceIn**](ResolveReferenceIn.md) |  | 

### Return type

[**ResolveReferenceOut**](ResolveReferenceOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSetReference

> SetReferenceOut RiskSetReference(ctx, set).SetReferenceIn(setReferenceIn).Execute()

Writes your organisation's own allow and deny entries over a set.



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
	set := "domain" // string | 
	setReferenceIn := *openapiclient.NewSetReferenceIn() // SetReferenceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAPI.RiskSetReference(context.Background(), set).SetReferenceIn(setReferenceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAPI.RiskSetReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSetReference`: SetReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `RiskAPI.RiskSetReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskSetReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setReferenceIn** | [**SetReferenceIn**](SetReferenceIn.md) |  | 

### Return type

[**SetReferenceOut**](SetReferenceOut.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

