# \DatastoreAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1DatastoreName**](DatastoreAPI.md#CloudDeleteV1DatastoreName) | **Delete** /v1/datastore/{name} | DropDatastore deprovisions one Hanzo Datastore warehouse.
[**CloudGetV1Datastore**](DatastoreAPI.md#CloudGetV1Datastore) | **Get** /v1/datastore | ListDatastore lists the caller org&#39;s Hanzo Datastore warehouses.
[**CloudGetV1DatastoreName**](DatastoreAPI.md#CloudGetV1DatastoreName) | **Get** /v1/datastore/{name} | GetDatastore returns one Hanzo Datastore warehouse&#39;s metadata.
[**CloudPostV1Datastore**](DatastoreAPI.md#CloudPostV1Datastore) | **Post** /v1/datastore | 



## CloudDeleteV1DatastoreName

> CloudDeleteV1DatastoreName(ctx, name).Execute()

DropDatastore deprovisions one Hanzo Datastore warehouse.



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
	r, err := apiClient.DatastoreAPI.CloudDeleteV1DatastoreName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoreAPI.CloudDeleteV1DatastoreName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1DatastoreNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Datastore

> []CloudProvisionedSummary CloudGetV1Datastore(ctx).Execute()

ListDatastore lists the caller org's Hanzo Datastore warehouses.



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
	resp, r, err := apiClient.DatastoreAPI.CloudGetV1Datastore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoreAPI.CloudGetV1Datastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Datastore`: []CloudProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `DatastoreAPI.CloudGetV1Datastore`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DatastoreRequest struct via the builder pattern


### Return type

[**[]CloudProvisionedSummary**](CloudProvisionedSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DatastoreName

> CloudProvisionedResource CloudGetV1DatastoreName(ctx, name).Execute()

GetDatastore returns one Hanzo Datastore warehouse's metadata.



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
	resp, r, err := apiClient.DatastoreAPI.CloudGetV1DatastoreName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoreAPI.CloudGetV1DatastoreName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DatastoreName`: CloudProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `DatastoreAPI.CloudGetV1DatastoreName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DatastoreNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProvisionedResource**](CloudProvisionedResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Datastore

> CloudProvisionResult CloudPostV1Datastore(ctx).CloudProvisionRequest(cloudProvisionRequest).Execute()



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
	cloudProvisionRequest := *openapiclient.NewCloudProvisionRequest() // CloudProvisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatastoreAPI.CloudPostV1Datastore(context.Background()).CloudProvisionRequest(cloudProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoreAPI.CloudPostV1Datastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Datastore`: CloudProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `DatastoreAPI.CloudPostV1Datastore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1DatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvisionRequest** | [**CloudProvisionRequest**](CloudProvisionRequest.md) |  | 

### Return type

[**CloudProvisionResult**](CloudProvisionResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

