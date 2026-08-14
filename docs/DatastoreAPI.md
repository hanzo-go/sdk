# \DatastoreAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDatastoreByName**](DatastoreAPI.md#DeleteDatastoreByName) | **Delete** /v1/datastore/{name} | Deprovisions one Hanzo Datastore warehouse.
[**GetDatastore**](DatastoreAPI.md#GetDatastore) | **Get** /v1/datastore | Lists the caller org&#39;s Hanzo Datastore warehouses.
[**GetDatastoreByName**](DatastoreAPI.md#GetDatastoreByName) | **Get** /v1/datastore/{name} | Returns one Hanzo Datastore warehouse&#39;s metadata.
[**PostDatastore**](DatastoreAPI.md#PostDatastore) | **Post** /v1/datastore | Provision a Hanzo Datastore instance for your org



## DeleteDatastoreByName

> DeleteDatastoreByName(ctx, name).Execute()

Deprovisions one Hanzo Datastore warehouse.



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
	name := "warehouse" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatastoreAPI.DeleteDatastoreByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoreAPI.DeleteDatastoreByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatastoreByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDatastore

> []ProvisionedSummary GetDatastore(ctx).Execute()

Lists the caller org's Hanzo Datastore warehouses.



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
	resp, r, err := apiClient.DatastoreAPI.GetDatastore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoreAPI.GetDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatastore`: []ProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `DatastoreAPI.GetDatastore`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatastoreRequest struct via the builder pattern


### Return type

[**[]ProvisionedSummary**](ProvisionedSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatastoreByName

> ProvisionedResource GetDatastoreByName(ctx, name).Execute()

Returns one Hanzo Datastore warehouse's metadata.



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
	name := "warehouse" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatastoreAPI.GetDatastoreByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoreAPI.GetDatastoreByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatastoreByName`: ProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `DatastoreAPI.GetDatastoreByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatastoreByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisionedResource**](ProvisionedResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatastore

> ProvisionResult PostDatastore(ctx).ProvisionRequest(provisionRequest).Execute()

Provision a Hanzo Datastore instance for your org



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
	provisionRequest := *openapiclient.NewProvisionRequest() // ProvisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatastoreAPI.PostDatastore(context.Background()).ProvisionRequest(provisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoreAPI.PostDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatastore`: ProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `DatastoreAPI.PostDatastore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionRequest** | [**ProvisionRequest**](ProvisionRequest.md) |  | 

### Return type

[**ProvisionResult**](ProvisionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

