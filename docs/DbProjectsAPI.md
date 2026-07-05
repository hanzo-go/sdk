# \DbProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbCreateProject**](DbProjectsAPI.md#DbCreateProject) | **Post** /v1/db/projects | Create project
[**DbDeleteProject**](DbProjectsAPI.md#DbDeleteProject) | **Delete** /v1/db/projects/{id} | Delete project
[**DbGetConnectionUri**](DbProjectsAPI.md#DbGetConnectionUri) | **Get** /v1/db/projects/{id}/connection_uri | Get connection URI
[**DbGetProject**](DbProjectsAPI.md#DbGetProject) | **Get** /v1/db/projects/{id} | Get project
[**DbListProjects**](DbProjectsAPI.md#DbListProjects) | **Get** /v1/db/projects | List projects
[**DbUpdateProject**](DbProjectsAPI.md#DbUpdateProject) | **Put** /v1/db/projects/{id} | Update project



## DbCreateProject

> DbCreateProject201Response DbCreateProject(ctx).DbCreateProjectRequest(dbCreateProjectRequest).Execute()

Create project

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
	dbCreateProjectRequest := *openapiclient.NewDbCreateProjectRequest(*openapiclient.NewDbProjectCreate("Name_example")) // DbCreateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbProjectsAPI.DbCreateProject(context.Background()).DbCreateProjectRequest(dbCreateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbProjectsAPI.DbCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbCreateProject`: DbCreateProject201Response
	fmt.Fprintf(os.Stdout, "Response from `DbProjectsAPI.DbCreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbCreateProjectRequest** | [**DbCreateProjectRequest**](DbCreateProjectRequest.md) |  | 

### Return type

[**DbCreateProject201Response**](DbCreateProject201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbDeleteProject

> DbGetProject200Response DbDeleteProject(ctx, id).Execute()

Delete project

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
	resp, r, err := apiClient.DbProjectsAPI.DbDeleteProject(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbProjectsAPI.DbDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbDeleteProject`: DbGetProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DbProjectsAPI.DbDeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbGetProject200Response**](DbGetProject200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbGetConnectionUri

> DbConnectionUri DbGetConnectionUri(ctx, id).RoleName(roleName).BranchId(branchId).EndpointId(endpointId).DatabaseName(databaseName).Pooled(pooled).Execute()

Get connection URI

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
	roleName := "roleName_example" // string | 
	branchId := "branchId_example" // string |  (optional)
	endpointId := "endpointId_example" // string |  (optional)
	databaseName := "databaseName_example" // string |  (optional) (default to "neondb")
	pooled := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbProjectsAPI.DbGetConnectionUri(context.Background(), id).RoleName(roleName).BranchId(branchId).EndpointId(endpointId).DatabaseName(databaseName).Pooled(pooled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbProjectsAPI.DbGetConnectionUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbGetConnectionUri`: DbConnectionUri
	fmt.Fprintf(os.Stdout, "Response from `DbProjectsAPI.DbGetConnectionUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbGetConnectionUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleName** | **string** |  | 
 **branchId** | **string** |  | 
 **endpointId** | **string** |  | 
 **databaseName** | **string** |  | [default to &quot;neondb&quot;]
 **pooled** | **bool** |  | [default to true]

### Return type

[**DbConnectionUri**](DbConnectionUri.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbGetProject

> DbGetProject200Response DbGetProject(ctx, id).Execute()

Get project

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
	resp, r, err := apiClient.DbProjectsAPI.DbGetProject(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbProjectsAPI.DbGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbGetProject`: DbGetProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DbProjectsAPI.DbGetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbGetProject200Response**](DbGetProject200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbListProjects

> DbListProjects200Response DbListProjects(ctx).Cursor(cursor).Limit(limit).Execute()

List projects

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
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbProjectsAPI.DbListProjects(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbProjectsAPI.DbListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbListProjects`: DbListProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `DbProjectsAPI.DbListProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 10]

### Return type

[**DbListProjects200Response**](DbListProjects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbUpdateProject

> DbGetProject200Response DbUpdateProject(ctx, id).DbUpdateProjectRequest(dbUpdateProjectRequest).Execute()

Update project

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
	dbUpdateProjectRequest := *openapiclient.NewDbUpdateProjectRequest(*openapiclient.NewDbUpdateProjectRequestProject()) // DbUpdateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbProjectsAPI.DbUpdateProject(context.Background(), id).DbUpdateProjectRequest(dbUpdateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbProjectsAPI.DbUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbUpdateProject`: DbGetProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DbProjectsAPI.DbUpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbUpdateProjectRequest** | [**DbUpdateProjectRequest**](DbUpdateProjectRequest.md) |  | 

### Return type

[**DbGetProject200Response**](DbGetProject200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

