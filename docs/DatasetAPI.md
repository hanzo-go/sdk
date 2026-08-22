# \DatasetAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RiskCreateDataset**](DatasetAPI.md#RiskCreateDataset) | **Post** /v1/dataset | Declare the next version of a dataset
[**RiskDataset**](DatasetAPI.md#RiskDataset) | **Get** /v1/dataset/{name} | Describe every version of one dataset
[**RiskDatasetLineage**](DatasetAPI.md#RiskDatasetLineage) | **Get** /v1/dataset/{name}/lineage | Show where a version&#39;s rows came from, and whether that can still be demonstrated
[**RiskDatasets**](DatasetAPI.md#RiskDatasets) | **Get** /v1/dataset | List this org&#39;s datasets
[**RiskDeleteDataset**](DatasetAPI.md#RiskDeleteDataset) | **Delete** /v1/dataset/{name} | Dispose of one dataset and every version of it
[**RiskExportDataset**](DatasetAPI.md#RiskExportDataset) | **Get** /v1/dataset/{name}/export | Read a version&#39;s rows back, one page at a time
[**RiskMaterializeDataset**](DatasetAPI.md#RiskMaterializeDataset) | **Post** /v1/dataset/{name}/materialize | Materialise the declared version into immutable rows



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
	resp, r, err := apiClient.DatasetAPI.RiskCreateDataset(context.Background()).RiskDatasetSpec(riskDatasetSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.RiskCreateDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskCreateDataset`: RiskDataset
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.RiskCreateDataset`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
	resp, r, err := apiClient.DatasetAPI.RiskDataset(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.RiskDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDataset`: RiskDatasetVersions
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.RiskDataset`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
	resp, r, err := apiClient.DatasetAPI.RiskDatasetLineage(context.Background(), name).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.RiskDatasetLineage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDatasetLineage`: RiskLineage
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.RiskDatasetLineage`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
	resp, r, err := apiClient.DatasetAPI.RiskDatasets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.RiskDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDatasets`: RiskDatasetList
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.RiskDatasets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskDatasetsRequest struct via the builder pattern


### Return type

[**RiskDatasetList**](RiskDatasetList.md)

### Authorization

[bearer](../README.md#bearer)

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
	resp, r, err := apiClient.DatasetAPI.RiskDeleteDataset(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.RiskDeleteDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskDeleteDataset`: RiskDatasetDisposal
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.RiskDeleteDataset`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.DatasetAPI.RiskExportDataset(context.Background(), name).Version(version).Split(split).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.RiskExportDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskExportDataset`: RiskDatasetRows
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.RiskExportDataset`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.DatasetAPI.RiskMaterializeDataset(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.RiskMaterializeDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskMaterializeDataset`: RiskDataset
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.RiskMaterializeDataset`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

