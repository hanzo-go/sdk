# \ConsoleDatasetItemsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreateDatasetItem**](ConsoleDatasetItemsAPI.md#ConsoleCreateDatasetItem) | **Post** /v1/console/dataset-items | Create or upsert a dataset item
[**ConsoleDeleteDatasetItem**](ConsoleDatasetItemsAPI.md#ConsoleDeleteDatasetItem) | **Delete** /v1/console/dataset-items/{id} | Delete a dataset item
[**ConsoleGetDatasetItem**](ConsoleDatasetItemsAPI.md#ConsoleGetDatasetItem) | **Get** /v1/console/dataset-items/{id} | Get a dataset item
[**ConsoleListDatasetItems**](ConsoleDatasetItemsAPI.md#ConsoleListDatasetItems) | **Get** /v1/console/dataset-items | Get dataset items



## ConsoleCreateDatasetItem

> ConsoleDatasetItem ConsoleCreateDatasetItem(ctx).ConsoleCreateDatasetItemRequest(consoleCreateDatasetItemRequest).Execute()

Create or upsert a dataset item

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
	consoleCreateDatasetItemRequest := *openapiclient.NewConsoleCreateDatasetItemRequest("DatasetName_example") // ConsoleCreateDatasetItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetItemsAPI.ConsoleCreateDatasetItem(context.Background()).ConsoleCreateDatasetItemRequest(consoleCreateDatasetItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetItemsAPI.ConsoleCreateDatasetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateDatasetItem`: ConsoleDatasetItem
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetItemsAPI.ConsoleCreateDatasetItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateDatasetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreateDatasetItemRequest** | [**ConsoleCreateDatasetItemRequest**](ConsoleCreateDatasetItemRequest.md) |  | 

### Return type

[**ConsoleDatasetItem**](ConsoleDatasetItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleDeleteDatasetItem

> ConsoleDeleteDatasetItem200Response ConsoleDeleteDatasetItem(ctx, id).Execute()

Delete a dataset item

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetItemsAPI.ConsoleDeleteDatasetItem(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetItemsAPI.ConsoleDeleteDatasetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteDatasetItem`: ConsoleDeleteDatasetItem200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetItemsAPI.ConsoleDeleteDatasetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteDatasetItemRequest struct via the builder pattern


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


## ConsoleGetDatasetItem

> ConsoleDatasetItem ConsoleGetDatasetItem(ctx, id).Execute()

Get a dataset item

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetItemsAPI.ConsoleGetDatasetItem(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetItemsAPI.ConsoleGetDatasetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetDatasetItem`: ConsoleDatasetItem
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetItemsAPI.ConsoleGetDatasetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetDatasetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleDatasetItem**](ConsoleDatasetItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListDatasetItems

> ConsoleListDatasetItems200Response ConsoleListDatasetItems(ctx).DatasetName(datasetName).SourceTraceId(sourceTraceId).SourceObservationId(sourceObservationId).Version(version).Page(page).Limit(limit).Execute()

Get dataset items

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
	datasetName := "datasetName_example" // string |  (optional)
	sourceTraceId := "sourceTraceId_example" // string |  (optional)
	sourceObservationId := "sourceObservationId_example" // string |  (optional)
	version := time.Now() // time.Time | ISO 8601 timestamp. Returns state of dataset at this time. (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetItemsAPI.ConsoleListDatasetItems(context.Background()).DatasetName(datasetName).SourceTraceId(sourceTraceId).SourceObservationId(sourceObservationId).Version(version).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetItemsAPI.ConsoleListDatasetItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListDatasetItems`: ConsoleListDatasetItems200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetItemsAPI.ConsoleListDatasetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListDatasetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasetName** | **string** |  | 
 **sourceTraceId** | **string** |  | 
 **sourceObservationId** | **string** |  | 
 **version** | **time.Time** | ISO 8601 timestamp. Returns state of dataset at this time. | 
 **page** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ConsoleListDatasetItems200Response**](ConsoleListDatasetItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

