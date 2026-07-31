# \DatasetsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvalsPostV1EvalsDatasetItems**](DatasetsAPI.md#EvalsPostV1EvalsDatasetItems) | **Post** /v1/evals/dataset-items | Add an item (input + optional expected output) to a dataset
[**EvalsPostV1EvalsDatasets**](DatasetsAPI.md#EvalsPostV1EvalsDatasets) | **Post** /v1/evals/datasets | Create a dataset



## EvalsPostV1EvalsDatasetItems

> EvalsDatasetItem EvalsPostV1EvalsDatasetItems(ctx).EvalsDatasetItemCreate(evalsDatasetItemCreate).Execute()

Add an item (input + optional expected output) to a dataset

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
	evalsDatasetItemCreate := *openapiclient.NewEvalsDatasetItemCreate("DatasetName_example", interface{}(123)) // EvalsDatasetItemCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetsAPI.EvalsPostV1EvalsDatasetItems(context.Background()).EvalsDatasetItemCreate(evalsDatasetItemCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetsAPI.EvalsPostV1EvalsDatasetItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvalsPostV1EvalsDatasetItems`: EvalsDatasetItem
	fmt.Fprintf(os.Stdout, "Response from `DatasetsAPI.EvalsPostV1EvalsDatasetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvalsPostV1EvalsDatasetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalsDatasetItemCreate** | [**EvalsDatasetItemCreate**](EvalsDatasetItemCreate.md) |  | 

### Return type

[**EvalsDatasetItem**](EvalsDatasetItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvalsPostV1EvalsDatasets

> EvalsDataset EvalsPostV1EvalsDatasets(ctx).EvalsDatasetCreate(evalsDatasetCreate).Execute()

Create a dataset

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
	evalsDatasetCreate := *openapiclient.NewEvalsDatasetCreate("Name_example") // EvalsDatasetCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetsAPI.EvalsPostV1EvalsDatasets(context.Background()).EvalsDatasetCreate(evalsDatasetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetsAPI.EvalsPostV1EvalsDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvalsPostV1EvalsDatasets`: EvalsDataset
	fmt.Fprintf(os.Stdout, "Response from `DatasetsAPI.EvalsPostV1EvalsDatasets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvalsPostV1EvalsDatasetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalsDatasetCreate** | [**EvalsDatasetCreate**](EvalsDatasetCreate.md) |  | 

### Return type

[**EvalsDataset**](EvalsDataset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

