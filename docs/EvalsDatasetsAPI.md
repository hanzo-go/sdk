# \EvalsDatasetsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EvalsDatasetItemsPost**](EvalsDatasetsAPI.md#V1EvalsDatasetItemsPost) | **Post** /v1/evals/dataset-items | Add an item (input + optional expected output) to a dataset
[**V1EvalsDatasetsPost**](EvalsDatasetsAPI.md#V1EvalsDatasetsPost) | **Post** /v1/evals/datasets | Create a dataset



## V1EvalsDatasetItemsPost

> EvalsDatasetItem V1EvalsDatasetItemsPost(ctx).EvalsDatasetItemCreate(evalsDatasetItemCreate).Execute()

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
	resp, r, err := apiClient.EvalsDatasetsAPI.V1EvalsDatasetItemsPost(context.Background()).EvalsDatasetItemCreate(evalsDatasetItemCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsDatasetsAPI.V1EvalsDatasetItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EvalsDatasetItemsPost`: EvalsDatasetItem
	fmt.Fprintf(os.Stdout, "Response from `EvalsDatasetsAPI.V1EvalsDatasetItemsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EvalsDatasetItemsPostRequest struct via the builder pattern


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


## V1EvalsDatasetsPost

> EvalsDataset V1EvalsDatasetsPost(ctx).EvalsDatasetCreate(evalsDatasetCreate).Execute()

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
	resp, r, err := apiClient.EvalsDatasetsAPI.V1EvalsDatasetsPost(context.Background()).EvalsDatasetCreate(evalsDatasetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsDatasetsAPI.V1EvalsDatasetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EvalsDatasetsPost`: EvalsDataset
	fmt.Fprintf(os.Stdout, "Response from `EvalsDatasetsAPI.V1EvalsDatasetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EvalsDatasetsPostRequest struct via the builder pattern


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

