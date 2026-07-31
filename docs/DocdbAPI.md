# \DocdbAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1DocdbName**](DocdbAPI.md#CloudDeleteV1DocdbName) | **Delete** /v1/docdb/{name} | DropDocDB deprovisions one Hanzo DocDB database.
[**CloudGetV1Docdb**](DocdbAPI.md#CloudGetV1Docdb) | **Get** /v1/docdb | ListDocDB lists the caller org&#39;s Hanzo DocDB document databases.
[**CloudGetV1DocdbName**](DocdbAPI.md#CloudGetV1DocdbName) | **Get** /v1/docdb/{name} | GetDocDB returns one Hanzo DocDB database&#39;s metadata.
[**CloudPostV1Docdb**](DocdbAPI.md#CloudPostV1Docdb) | **Post** /v1/docdb | 



## CloudDeleteV1DocdbName

> CloudDeleteV1DocdbName(ctx, name).Execute()

DropDocDB deprovisions one Hanzo DocDB database.



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
	name := "sessions" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DocdbAPI.CloudDeleteV1DocdbName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocdbAPI.CloudDeleteV1DocdbName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1DocdbNameRequest struct via the builder pattern


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


## CloudGetV1Docdb

> []CloudProvisionedSummary CloudGetV1Docdb(ctx).Execute()

ListDocDB lists the caller org's Hanzo DocDB document databases.



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
	resp, r, err := apiClient.DocdbAPI.CloudGetV1Docdb(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocdbAPI.CloudGetV1Docdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Docdb`: []CloudProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `DocdbAPI.CloudGetV1Docdb`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DocdbRequest struct via the builder pattern


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


## CloudGetV1DocdbName

> CloudProvisionedResource CloudGetV1DocdbName(ctx, name).Execute()

GetDocDB returns one Hanzo DocDB database's metadata.



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
	name := "sessions" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocdbAPI.CloudGetV1DocdbName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocdbAPI.CloudGetV1DocdbName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DocdbName`: CloudProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `DocdbAPI.CloudGetV1DocdbName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DocdbNameRequest struct via the builder pattern


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


## CloudPostV1Docdb

> CloudProvisionResult CloudPostV1Docdb(ctx).CloudProvisionRequest(cloudProvisionRequest).Execute()



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
	resp, r, err := apiClient.DocdbAPI.CloudPostV1Docdb(context.Background()).CloudProvisionRequest(cloudProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocdbAPI.CloudPostV1Docdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Docdb`: CloudProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `DocdbAPI.CloudPostV1Docdb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1DocdbRequest struct via the builder pattern


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

