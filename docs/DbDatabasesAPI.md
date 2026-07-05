# \DbDatabasesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbCreateDatabase**](DbDatabasesAPI.md#DbCreateDatabase) | **Post** /v1/db/projects/{id}/databases | Create database
[**DbDeleteDatabase**](DbDatabasesAPI.md#DbDeleteDatabase) | **Delete** /v1/db/projects/{id}/databases/{name} | Delete database
[**DbGetDatabase**](DbDatabasesAPI.md#DbGetDatabase) | **Get** /v1/db/projects/{id}/databases/{name} | Get database
[**DbListDatabases**](DbDatabasesAPI.md#DbListDatabases) | **Get** /v1/db/projects/{id}/databases | List databases
[**DbUpdateDatabase**](DbDatabasesAPI.md#DbUpdateDatabase) | **Put** /v1/db/projects/{id}/databases/{name} | Update database



## DbCreateDatabase

> DbCreateDatabase201Response DbCreateDatabase(ctx, id).DbCreateDatabaseRequest(dbCreateDatabaseRequest).Execute()

Create database

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
	dbCreateDatabaseRequest := *openapiclient.NewDbCreateDatabaseRequest(*openapiclient.NewDbDatabaseCreate("Name_example", "OwnerName_example")) // DbCreateDatabaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbDatabasesAPI.DbCreateDatabase(context.Background(), id).DbCreateDatabaseRequest(dbCreateDatabaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbDatabasesAPI.DbCreateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbCreateDatabase`: DbCreateDatabase201Response
	fmt.Fprintf(os.Stdout, "Response from `DbDatabasesAPI.DbCreateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbCreateDatabaseRequest** | [**DbCreateDatabaseRequest**](DbCreateDatabaseRequest.md) |  | 

### Return type

[**DbCreateDatabase201Response**](DbCreateDatabase201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbDeleteDatabase

> DbCreateDatabase201Response DbDeleteDatabase(ctx, id, name).BranchId(branchId).Execute()

Delete database

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
	name := "name_example" // string | 
	branchId := "branchId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbDatabasesAPI.DbDeleteDatabase(context.Background(), id, name).BranchId(branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbDatabasesAPI.DbDeleteDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbDeleteDatabase`: DbCreateDatabase201Response
	fmt.Fprintf(os.Stdout, "Response from `DbDatabasesAPI.DbDeleteDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbDeleteDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branchId** | **string** |  | 

### Return type

[**DbCreateDatabase201Response**](DbCreateDatabase201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbGetDatabase

> DbCreateDatabase201Response DbGetDatabase(ctx, id, name).BranchId(branchId).Execute()

Get database

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
	name := "name_example" // string | 
	branchId := "branchId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbDatabasesAPI.DbGetDatabase(context.Background(), id, name).BranchId(branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbDatabasesAPI.DbGetDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbGetDatabase`: DbCreateDatabase201Response
	fmt.Fprintf(os.Stdout, "Response from `DbDatabasesAPI.DbGetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branchId** | **string** |  | 

### Return type

[**DbCreateDatabase201Response**](DbCreateDatabase201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbListDatabases

> DbListDatabases200Response DbListDatabases(ctx, id).BranchId(branchId).Execute()

List databases

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
	branchId := "branchId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbDatabasesAPI.DbListDatabases(context.Background(), id).BranchId(branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbDatabasesAPI.DbListDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbListDatabases`: DbListDatabases200Response
	fmt.Fprintf(os.Stdout, "Response from `DbDatabasesAPI.DbListDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchId** | **string** |  | 

### Return type

[**DbListDatabases200Response**](DbListDatabases200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbUpdateDatabase

> DbCreateDatabase201Response DbUpdateDatabase(ctx, id, name).DbUpdateDatabaseRequest(dbUpdateDatabaseRequest).Execute()

Update database

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
	name := "name_example" // string | 
	dbUpdateDatabaseRequest := *openapiclient.NewDbUpdateDatabaseRequest(*openapiclient.NewDbUpdateDatabaseRequestDatabase()) // DbUpdateDatabaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbDatabasesAPI.DbUpdateDatabase(context.Background(), id, name).DbUpdateDatabaseRequest(dbUpdateDatabaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbDatabasesAPI.DbUpdateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbUpdateDatabase`: DbCreateDatabase201Response
	fmt.Fprintf(os.Stdout, "Response from `DbDatabasesAPI.DbUpdateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbUpdateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dbUpdateDatabaseRequest** | [**DbUpdateDatabaseRequest**](DbUpdateDatabaseRequest.md) |  | 

### Return type

[**DbCreateDatabase201Response**](DbCreateDatabase201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

