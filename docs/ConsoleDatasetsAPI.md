# \ConsoleDatasetsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreateDataset**](ConsoleDatasetsAPI.md#ConsoleCreateDataset) | **Post** /v1/console/datasets | Create a dataset
[**ConsoleDeleteDatasetRun**](ConsoleDatasetsAPI.md#ConsoleDeleteDatasetRun) | **Delete** /v1/console/datasets/{datasetName}/runs/{runName} | Delete a dataset run and all its run items
[**ConsoleGetDataset**](ConsoleDatasetsAPI.md#ConsoleGetDataset) | **Get** /v1/console/datasets/{datasetName} | Get a dataset by name
[**ConsoleGetDatasetRun**](ConsoleDatasetsAPI.md#ConsoleGetDatasetRun) | **Get** /v1/console/datasets/{datasetName}/runs/{runName} | Get a dataset run and its items
[**ConsoleListDatasetRuns**](ConsoleDatasetsAPI.md#ConsoleListDatasetRuns) | **Get** /v1/console/datasets/{datasetName}/runs | Get dataset runs
[**ConsoleListDatasets**](ConsoleDatasetsAPI.md#ConsoleListDatasets) | **Get** /v1/console/datasets | Get all datasets



## ConsoleCreateDataset

> ConsoleDataset ConsoleCreateDataset(ctx).ConsoleCreateDatasetRequest(consoleCreateDatasetRequest).Execute()

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
	consoleCreateDatasetRequest := *openapiclient.NewConsoleCreateDatasetRequest("Name_example") // ConsoleCreateDatasetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetsAPI.ConsoleCreateDataset(context.Background()).ConsoleCreateDatasetRequest(consoleCreateDatasetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetsAPI.ConsoleCreateDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateDataset`: ConsoleDataset
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetsAPI.ConsoleCreateDataset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreateDatasetRequest** | [**ConsoleCreateDatasetRequest**](ConsoleCreateDatasetRequest.md) |  | 

### Return type

[**ConsoleDataset**](ConsoleDataset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleDeleteDatasetRun

> ConsoleDeleteDatasetItem200Response ConsoleDeleteDatasetRun(ctx, datasetName, runName).Execute()

Delete a dataset run and all its run items

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
	datasetName := "datasetName_example" // string | 
	runName := "runName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetsAPI.ConsoleDeleteDatasetRun(context.Background(), datasetName, runName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetsAPI.ConsoleDeleteDatasetRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteDatasetRun`: ConsoleDeleteDatasetItem200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetsAPI.ConsoleDeleteDatasetRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetName** | **string** |  | 
**runName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteDatasetRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsoleDeleteDatasetItem200Response**](ConsoleDeleteDatasetItem200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetDataset

> ConsoleDataset ConsoleGetDataset(ctx, datasetName).Execute()

Get a dataset by name

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
	datasetName := "datasetName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetsAPI.ConsoleGetDataset(context.Background(), datasetName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetsAPI.ConsoleGetDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetDataset`: ConsoleDataset
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetsAPI.ConsoleGetDataset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleDataset**](ConsoleDataset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetDatasetRun

> ConsoleGetDatasetRun200Response ConsoleGetDatasetRun(ctx, datasetName, runName).Execute()

Get a dataset run and its items

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
	datasetName := "datasetName_example" // string | 
	runName := "runName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetsAPI.ConsoleGetDatasetRun(context.Background(), datasetName, runName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetsAPI.ConsoleGetDatasetRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetDatasetRun`: ConsoleGetDatasetRun200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetsAPI.ConsoleGetDatasetRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetName** | **string** |  | 
**runName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetDatasetRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsoleGetDatasetRun200Response**](ConsoleGetDatasetRun200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListDatasetRuns

> ConsoleListDatasetRuns200Response ConsoleListDatasetRuns(ctx, datasetName).Page(page).Limit(limit).Execute()

Get dataset runs

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
	datasetName := "datasetName_example" // string | 
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetsAPI.ConsoleListDatasetRuns(context.Background(), datasetName).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetsAPI.ConsoleListDatasetRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListDatasetRuns`: ConsoleListDatasetRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetsAPI.ConsoleListDatasetRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListDatasetRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ConsoleListDatasetRuns200Response**](ConsoleListDatasetRuns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListDatasets

> ConsoleListDatasets200Response ConsoleListDatasets(ctx).Page(page).Limit(limit).Execute()

Get all datasets

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
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetsAPI.ConsoleListDatasets(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetsAPI.ConsoleListDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListDatasets`: ConsoleListDatasets200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetsAPI.ConsoleListDatasets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListDatasetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ConsoleListDatasets200Response**](ConsoleListDatasets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

