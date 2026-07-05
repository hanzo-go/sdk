# \ChatPermissionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetPermissionsByresourcetypeByresourceid**](ChatPermissionsAPI.md#ChatGetPermissionsByresourcetypeByresourceid) | **Get** /v1/chat/permissions/{resourceType}/{resourceId} | Get all permissions for a resource
[**ChatGetPermissionsByresourcetypeByresourceidEffective**](ChatPermissionsAPI.md#ChatGetPermissionsByresourcetypeByresourceidEffective) | **Get** /v1/chat/permissions/{resourceType}/{resourceId}/effective | Get effective permissions for a specific resource
[**ChatGetPermissionsByresourcetypeEffectiveAll**](ChatPermissionsAPI.md#ChatGetPermissionsByresourcetypeEffectiveAll) | **Get** /v1/chat/permissions/{resourceType}/effective/all | Get effective permissions for all accessible resources
[**ChatGetPermissionsByresourcetypeRoles**](ChatPermissionsAPI.md#ChatGetPermissionsByresourcetypeRoles) | **Get** /v1/chat/permissions/{resourceType}/roles | Get available roles for a resource type
[**ChatGetPermissionsSearchPrincipals**](ChatPermissionsAPI.md#ChatGetPermissionsSearchPrincipals) | **Get** /v1/chat/permissions/search-principals | Search for users and groups to grant permissions
[**ChatPutPermissionsByresourcetypeByresourceid**](ChatPermissionsAPI.md#ChatPutPermissionsByresourcetypeByresourceid) | **Put** /v1/chat/permissions/{resourceType}/{resourceId} | Bulk update permissions for a resource



## ChatGetPermissionsByresourcetypeByresourceid

> map[string]interface{} ChatGetPermissionsByresourcetypeByresourceid(ctx, resourceType, resourceId).Execute()

Get all permissions for a resource

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
	resourceType := "resourceType_example" // string | 
	resourceId := "resourceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPermissionsAPI.ChatGetPermissionsByresourcetypeByresourceid(context.Background(), resourceType, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPermissionsAPI.ChatGetPermissionsByresourcetypeByresourceid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPermissionsByresourcetypeByresourceid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPermissionsAPI.ChatGetPermissionsByresourcetypeByresourceid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** |  | 
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPermissionsByresourcetypeByresourceidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetPermissionsByresourcetypeByresourceidEffective

> map[string]interface{} ChatGetPermissionsByresourcetypeByresourceidEffective(ctx, resourceType, resourceId).Execute()

Get effective permissions for a specific resource

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
	resourceType := "resourceType_example" // string | 
	resourceId := "resourceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPermissionsAPI.ChatGetPermissionsByresourcetypeByresourceidEffective(context.Background(), resourceType, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPermissionsAPI.ChatGetPermissionsByresourcetypeByresourceidEffective``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPermissionsByresourcetypeByresourceidEffective`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPermissionsAPI.ChatGetPermissionsByresourcetypeByresourceidEffective`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** |  | 
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPermissionsByresourcetypeByresourceidEffectiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetPermissionsByresourcetypeEffectiveAll

> map[string]interface{} ChatGetPermissionsByresourcetypeEffectiveAll(ctx, resourceType).Execute()

Get effective permissions for all accessible resources

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
	resourceType := "resourceType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPermissionsAPI.ChatGetPermissionsByresourcetypeEffectiveAll(context.Background(), resourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPermissionsAPI.ChatGetPermissionsByresourcetypeEffectiveAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPermissionsByresourcetypeEffectiveAll`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPermissionsAPI.ChatGetPermissionsByresourcetypeEffectiveAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPermissionsByresourcetypeEffectiveAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetPermissionsByresourcetypeRoles

> map[string]interface{} ChatGetPermissionsByresourcetypeRoles(ctx, resourceType).Execute()

Get available roles for a resource type

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
	resourceType := "resourceType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPermissionsAPI.ChatGetPermissionsByresourcetypeRoles(context.Background(), resourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPermissionsAPI.ChatGetPermissionsByresourcetypeRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPermissionsByresourcetypeRoles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPermissionsAPI.ChatGetPermissionsByresourcetypeRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPermissionsByresourcetypeRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetPermissionsSearchPrincipals

> map[string]interface{} ChatGetPermissionsSearchPrincipals(ctx).Query(query).Execute()

Search for users and groups to grant permissions

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
	query := "query_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPermissionsAPI.ChatGetPermissionsSearchPrincipals(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPermissionsAPI.ChatGetPermissionsSearchPrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPermissionsSearchPrincipals`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPermissionsAPI.ChatGetPermissionsSearchPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPermissionsSearchPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPutPermissionsByresourcetypeByresourceid

> map[string]interface{} ChatPutPermissionsByresourcetypeByresourceid(ctx, resourceType, resourceId).Body(body).Execute()

Bulk update permissions for a resource

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
	resourceType := "resourceType_example" // string | 
	resourceId := "resourceId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPermissionsAPI.ChatPutPermissionsByresourcetypeByresourceid(context.Background(), resourceType, resourceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPermissionsAPI.ChatPutPermissionsByresourcetypeByresourceid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutPermissionsByresourcetypeByresourceid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPermissionsAPI.ChatPutPermissionsByresourcetypeByresourceid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** |  | 
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutPermissionsByresourcetypeByresourceidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

