# \DbBranchesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbCreateBranch**](DbBranchesAPI.md#DbCreateBranch) | **Post** /v1/db/projects/{id}/branches | Create branch
[**DbDeleteBranch**](DbBranchesAPI.md#DbDeleteBranch) | **Delete** /v1/db/projects/{id}/branches/{branch_id} | Delete branch
[**DbGetBranch**](DbBranchesAPI.md#DbGetBranch) | **Get** /v1/db/projects/{id}/branches/{branch_id} | Get branch
[**DbListBranches**](DbBranchesAPI.md#DbListBranches) | **Get** /v1/db/projects/{id}/branches | List branches
[**DbRestoreBranch**](DbBranchesAPI.md#DbRestoreBranch) | **Post** /v1/db/projects/{id}/branches/{branch_id}/restore | Point-in-time restore



## DbCreateBranch

> DbCreateBranch201Response DbCreateBranch(ctx, id).DbCreateBranchRequest(dbCreateBranchRequest).Execute()

Create branch



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
	dbCreateBranchRequest := *openapiclient.NewDbCreateBranchRequest() // DbCreateBranchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbBranchesAPI.DbCreateBranch(context.Background(), id).DbCreateBranchRequest(dbCreateBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbBranchesAPI.DbCreateBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbCreateBranch`: DbCreateBranch201Response
	fmt.Fprintf(os.Stdout, "Response from `DbBranchesAPI.DbCreateBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbCreateBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbCreateBranchRequest** | [**DbCreateBranchRequest**](DbCreateBranchRequest.md) |  | 

### Return type

[**DbCreateBranch201Response**](DbCreateBranch201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbDeleteBranch

> DbGetBranch200Response DbDeleteBranch(ctx, id, branchId).Execute()

Delete branch

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
	resp, r, err := apiClient.DbBranchesAPI.DbDeleteBranch(context.Background(), id, branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbBranchesAPI.DbDeleteBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbDeleteBranch`: DbGetBranch200Response
	fmt.Fprintf(os.Stdout, "Response from `DbBranchesAPI.DbDeleteBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**branchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbDeleteBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbGetBranch200Response**](DbGetBranch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbGetBranch

> DbGetBranch200Response DbGetBranch(ctx, id, branchId).Execute()

Get branch

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
	resp, r, err := apiClient.DbBranchesAPI.DbGetBranch(context.Background(), id, branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbBranchesAPI.DbGetBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbGetBranch`: DbGetBranch200Response
	fmt.Fprintf(os.Stdout, "Response from `DbBranchesAPI.DbGetBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**branchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbGetBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbGetBranch200Response**](DbGetBranch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbListBranches

> DbListBranches200Response DbListBranches(ctx, id).Execute()

List branches

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
	resp, r, err := apiClient.DbBranchesAPI.DbListBranches(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbBranchesAPI.DbListBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbListBranches`: DbListBranches200Response
	fmt.Fprintf(os.Stdout, "Response from `DbBranchesAPI.DbListBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbListBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbListBranches200Response**](DbListBranches200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbRestoreBranch

> DbRestoreBranch200Response DbRestoreBranch(ctx, id, branchId).DbRestoreBranchRequest(dbRestoreBranchRequest).Execute()

Point-in-time restore



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
	dbRestoreBranchRequest := *openapiclient.NewDbRestoreBranchRequest() // DbRestoreBranchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbBranchesAPI.DbRestoreBranch(context.Background(), id, branchId).DbRestoreBranchRequest(dbRestoreBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbBranchesAPI.DbRestoreBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbRestoreBranch`: DbRestoreBranch200Response
	fmt.Fprintf(os.Stdout, "Response from `DbBranchesAPI.DbRestoreBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**branchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbRestoreBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dbRestoreBranchRequest** | [**DbRestoreBranchRequest**](DbRestoreBranchRequest.md) |  | 

### Return type

[**DbRestoreBranch200Response**](DbRestoreBranch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

