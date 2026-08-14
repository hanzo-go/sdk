# \SqlAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSqlByName**](SqlAPI.md#DeleteSqlByName) | **Delete** /v1/sql/{name} | DropSQL deprovisions one Hanzo SQL database.
[**GetSql**](SqlAPI.md#GetSql) | **Get** /v1/sql | ListSQL lists the caller org&#39;s Hanzo SQL databases.
[**GetSqlByName**](SqlAPI.md#GetSqlByName) | **Get** /v1/sql/{name} | GetSQL returns one Hanzo SQL database&#39;s metadata.
[**PostSql**](SqlAPI.md#PostSql) | **Post** /v1/sql | Provision a PostgreSQL database for your org



## DeleteSqlByName

> DeleteSqlByName(ctx, name).Execute()

DropSQL deprovisions one Hanzo SQL database.



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
	name := "orders" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SqlAPI.DeleteSqlByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlAPI.DeleteSqlByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSqlByNameRequest struct via the builder pattern


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


## GetSql

> []ProvisionedSummary GetSql(ctx).Execute()

ListSQL lists the caller org's Hanzo SQL databases.



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
	resp, r, err := apiClient.SqlAPI.GetSql(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlAPI.GetSql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSql`: []ProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `SqlAPI.GetSql`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlRequest struct via the builder pattern


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


## GetSqlByName

> ProvisionedResource GetSqlByName(ctx, name).Execute()

GetSQL returns one Hanzo SQL database's metadata.



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
	name := "orders" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SqlAPI.GetSqlByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlAPI.GetSqlByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSqlByName`: ProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `SqlAPI.GetSqlByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlByNameRequest struct via the builder pattern


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


## PostSql

> ProvisionResult PostSql(ctx).ProvisionRequest(provisionRequest).Execute()

Provision a PostgreSQL database for your org



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
	resp, r, err := apiClient.SqlAPI.PostSql(context.Background()).ProvisionRequest(provisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlAPI.PostSql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSql`: ProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `SqlAPI.PostSql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSqlRequest struct via the builder pattern


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

