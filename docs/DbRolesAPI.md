# \DbRolesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbCreateRole**](DbRolesAPI.md#DbCreateRole) | **Post** /v1/db/projects/{id}/roles | Create role
[**DbDeleteRole**](DbRolesAPI.md#DbDeleteRole) | **Delete** /v1/db/projects/{id}/roles/{name} | Delete role
[**DbGetRole**](DbRolesAPI.md#DbGetRole) | **Get** /v1/db/projects/{id}/roles/{name} | Get role
[**DbListRoles**](DbRolesAPI.md#DbListRoles) | **Get** /v1/db/projects/{id}/roles | List roles
[**DbResetRolePassword**](DbRolesAPI.md#DbResetRolePassword) | **Post** /v1/db/projects/{id}/roles/{name}/reset_password | Reset role password



## DbCreateRole

> DbCreateRole201Response DbCreateRole(ctx, id).DbCreateRoleRequest(dbCreateRoleRequest).Execute()

Create role

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
	dbCreateRoleRequest := *openapiclient.NewDbCreateRoleRequest(*openapiclient.NewDbRoleCreate("Name_example")) // DbCreateRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbRolesAPI.DbCreateRole(context.Background(), id).DbCreateRoleRequest(dbCreateRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbRolesAPI.DbCreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbCreateRole`: DbCreateRole201Response
	fmt.Fprintf(os.Stdout, "Response from `DbRolesAPI.DbCreateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbCreateRoleRequest** | [**DbCreateRoleRequest**](DbCreateRoleRequest.md) |  | 

### Return type

[**DbCreateRole201Response**](DbCreateRole201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbDeleteRole

> DbCreateRole201Response DbDeleteRole(ctx, id, name).BranchId(branchId).Execute()

Delete role

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
	resp, r, err := apiClient.DbRolesAPI.DbDeleteRole(context.Background(), id, name).BranchId(branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbRolesAPI.DbDeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbDeleteRole`: DbCreateRole201Response
	fmt.Fprintf(os.Stdout, "Response from `DbRolesAPI.DbDeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branchId** | **string** |  | 

### Return type

[**DbCreateRole201Response**](DbCreateRole201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbGetRole

> DbCreateRole201Response DbGetRole(ctx, id, name).BranchId(branchId).Execute()

Get role

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
	resp, r, err := apiClient.DbRolesAPI.DbGetRole(context.Background(), id, name).BranchId(branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbRolesAPI.DbGetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbGetRole`: DbCreateRole201Response
	fmt.Fprintf(os.Stdout, "Response from `DbRolesAPI.DbGetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branchId** | **string** |  | 

### Return type

[**DbCreateRole201Response**](DbCreateRole201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbListRoles

> DbListRoles200Response DbListRoles(ctx, id).BranchId(branchId).Execute()

List roles

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
	resp, r, err := apiClient.DbRolesAPI.DbListRoles(context.Background(), id).BranchId(branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbRolesAPI.DbListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbListRoles`: DbListRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `DbRolesAPI.DbListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchId** | **string** |  | 

### Return type

[**DbListRoles200Response**](DbListRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbResetRolePassword

> DbCreateRole201Response DbResetRolePassword(ctx, id, name).BranchId(branchId).Execute()

Reset role password

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
	resp, r, err := apiClient.DbRolesAPI.DbResetRolePassword(context.Background(), id, name).BranchId(branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbRolesAPI.DbResetRolePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbResetRolePassword`: DbCreateRole201Response
	fmt.Fprintf(os.Stdout, "Response from `DbRolesAPI.DbResetRolePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbResetRolePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branchId** | **string** |  | 

### Return type

[**DbCreateRole201Response**](DbCreateRole201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

