# \RolesPermissionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddEnforcer**](RolesPermissionsAPI.md#IamApiControllerAddEnforcer) | **Post** /v1/iam/enforcers | Api Controller Add Enforcer
[**IamApiControllerAddModel**](RolesPermissionsAPI.md#IamApiControllerAddModel) | **Post** /v1/iam/models | Api Controller Add Model
[**IamApiControllerAddPermission**](RolesPermissionsAPI.md#IamApiControllerAddPermission) | **Post** /v1/iam/permissions | Api Controller Add Permission
[**IamApiControllerAddPolicy**](RolesPermissionsAPI.md#IamApiControllerAddPolicy) | **Post** /v1/iam/policies | Api Controller Add Policy
[**IamApiControllerAddRole**](RolesPermissionsAPI.md#IamApiControllerAddRole) | **Post** /v1/iam/roles | Api Controller Add Role
[**IamApiControllerBatchEnforce**](RolesPermissionsAPI.md#IamApiControllerBatchEnforce) | **Post** /v1/iam/enforce/batch | Api Controller Batch Enforce
[**IamApiControllerDeleteEnforcer**](RolesPermissionsAPI.md#IamApiControllerDeleteEnforcer) | **Delete** /v1/iam/enforcers/{id} | Api Controller Delete Enforcer
[**IamApiControllerDeleteModel**](RolesPermissionsAPI.md#IamApiControllerDeleteModel) | **Delete** /v1/iam/models/{id} | Api Controller Delete Model
[**IamApiControllerDeletePermission**](RolesPermissionsAPI.md#IamApiControllerDeletePermission) | **Delete** /v1/iam/permissions/{id} | Api Controller Delete Permission
[**IamApiControllerDeleteRole**](RolesPermissionsAPI.md#IamApiControllerDeleteRole) | **Delete** /v1/iam/roles/{id} | Api Controller Delete Role
[**IamApiControllerEnforce**](RolesPermissionsAPI.md#IamApiControllerEnforce) | **Post** /v1/iam/enforce | Api Controller Enforce
[**IamApiControllerGetAllActions**](RolesPermissionsAPI.md#IamApiControllerGetAllActions) | **Get** /v1/iam/all-actions | Api Controller Get All Actions
[**IamApiControllerGetAllObjects**](RolesPermissionsAPI.md#IamApiControllerGetAllObjects) | **Get** /v1/iam/all-objects | Api Controller Get All Objects
[**IamApiControllerGetAllRoles**](RolesPermissionsAPI.md#IamApiControllerGetAllRoles) | **Get** /v1/iam/all-roles | Api Controller Get All Roles
[**IamApiControllerGetEnforcer**](RolesPermissionsAPI.md#IamApiControllerGetEnforcer) | **Get** /v1/iam/enforcers/{id} | Api Controller Get Enforcer
[**IamApiControllerGetEnforcers**](RolesPermissionsAPI.md#IamApiControllerGetEnforcers) | **Get** /v1/iam/enforcers | Api Controller Get Enforcers
[**IamApiControllerGetFilteredPolicies**](RolesPermissionsAPI.md#IamApiControllerGetFilteredPolicies) | **Get** /v1/iam/filtered-policies | Api Controller Get Filtered Policies
[**IamApiControllerGetModel**](RolesPermissionsAPI.md#IamApiControllerGetModel) | **Get** /v1/iam/models/{id} | Api Controller Get Model
[**IamApiControllerGetModels**](RolesPermissionsAPI.md#IamApiControllerGetModels) | **Get** /v1/iam/models | Api Controller Get Models
[**IamApiControllerGetPermission**](RolesPermissionsAPI.md#IamApiControllerGetPermission) | **Get** /v1/iam/permissions/{id} | Api Controller Get Permission
[**IamApiControllerGetPermissions**](RolesPermissionsAPI.md#IamApiControllerGetPermissions) | **Get** /v1/iam/permissions | Api Controller Get Permissions
[**IamApiControllerGetPermissionsByRole**](RolesPermissionsAPI.md#IamApiControllerGetPermissionsByRole) | **Get** /v1/iam/permissions-by-roles/{id} | Api Controller Get Permissions By Role
[**IamApiControllerGetPermissionsBySubmitter**](RolesPermissionsAPI.md#IamApiControllerGetPermissionsBySubmitter) | **Get** /v1/iam/permissions-by-submitters/{id} | Api Controller Get Permissions By Submitter
[**IamApiControllerGetPolicies**](RolesPermissionsAPI.md#IamApiControllerGetPolicies) | **Get** /v1/iam/policies | Api Controller Get Policies
[**IamApiControllerGetRole**](RolesPermissionsAPI.md#IamApiControllerGetRole) | **Get** /v1/iam/roles/{id} | Api Controller Get Role
[**IamApiControllerGetRoles**](RolesPermissionsAPI.md#IamApiControllerGetRoles) | **Get** /v1/iam/roles | Api Controller Get Roles
[**IamApiControllerRemovePolicy**](RolesPermissionsAPI.md#IamApiControllerRemovePolicy) | **Post** /v1/iam/remove-policy | Api Controller Remove Policy
[**IamApiControllerRunCasbinCommand**](RolesPermissionsAPI.md#IamApiControllerRunCasbinCommand) | **Get** /v1/iam/run-casbin-command | Api Controller Run Casbin Command
[**IamApiControllerUpdateEnforcer**](RolesPermissionsAPI.md#IamApiControllerUpdateEnforcer) | **Put** /v1/iam/enforcers/{id} | Api Controller Update Enforcer
[**IamApiControllerUpdateModel**](RolesPermissionsAPI.md#IamApiControllerUpdateModel) | **Put** /v1/iam/models/{id} | Api Controller Update Model
[**IamApiControllerUpdatePermission**](RolesPermissionsAPI.md#IamApiControllerUpdatePermission) | **Put** /v1/iam/permissions/{id} | Api Controller Update Permission
[**IamApiControllerUpdatePolicy**](RolesPermissionsAPI.md#IamApiControllerUpdatePolicy) | **Put** /v1/iam/policies/{id} | Api Controller Update Policy
[**IamApiControllerUpdateRole**](RolesPermissionsAPI.md#IamApiControllerUpdateRole) | **Put** /v1/iam/roles/{id} | Api Controller Update Role



## IamApiControllerAddEnforcer

> IamObjectEnforcer IamApiControllerAddEnforcer(ctx).Body(body).Execute()

Api Controller Add Enforcer



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
	body := map[string]interface{}{ ... } // map[string]interface{} | The enforcer object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerAddEnforcer(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerAddEnforcer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddEnforcer`: IamObjectEnforcer
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerAddEnforcer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddEnforcerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The enforcer object | 

### Return type

[**IamObjectEnforcer**](IamObjectEnforcer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddModel

> IamControllersResponse IamApiControllerAddModel(ctx).IamObjectModel(iamObjectModel).Execute()

Api Controller Add Model



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
	iamObjectModel := *openapiclient.NewIamObjectModel() // IamObjectModel | The details of the model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerAddModel(context.Background()).IamObjectModel(iamObjectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerAddModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddModel`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerAddModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectModel** | [**IamObjectModel**](IamObjectModel.md) | The details of the model | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddPermission

> IamControllersResponse IamApiControllerAddPermission(ctx).IamObjectPermission(iamObjectPermission).Execute()

Api Controller Add Permission



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
	iamObjectPermission := *openapiclient.NewIamObjectPermission() // IamObjectPermission | The details of the permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerAddPermission(context.Background()).IamObjectPermission(iamObjectPermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerAddPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddPermission`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerAddPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectPermission** | [**IamObjectPermission**](IamObjectPermission.md) | The details of the permission | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddPolicy

> map[string]interface{} IamApiControllerAddPolicy(ctx).Id(id).Body(body).Execute()

Api Controller Add Policy



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
	id := "id_example" // string | The id ( owner/name )  of enforcer
	body := map[string]interface{}{ ... } // map[string]interface{} | The policy to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerAddPolicy(context.Background()).Id(id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerAddPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerAddPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name )  of enforcer | 
 **body** | **map[string]interface{}** | The policy to add | 

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


## IamApiControllerAddRole

> IamControllersResponse IamApiControllerAddRole(ctx).IamObjectRole(iamObjectRole).Execute()

Api Controller Add Role



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
	iamObjectRole := *openapiclient.NewIamObjectRole() // IamObjectRole | The details of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerAddRole(context.Background()).IamObjectRole(iamObjectRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddRole`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerAddRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectRole** | [**IamObjectRole**](IamObjectRole.md) | The details of the role | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerBatchEnforce

> IamControllersResponse IamApiControllerBatchEnforce(ctx).RequestBody(requestBody).PermissionId(permissionId).ModelId(modelId).Owner(owner).Execute()

Api Controller Batch Enforce



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
	requestBody := []string{"Property_example"} // []string | array of casbin requests
	permissionId := "permissionId_example" // string | permission id (optional)
	modelId := "modelId_example" // string | model id (optional)
	owner := "owner_example" // string | owner (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerBatchEnforce(context.Background()).RequestBody(requestBody).PermissionId(permissionId).ModelId(modelId).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerBatchEnforce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerBatchEnforce`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerBatchEnforce`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerBatchEnforceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | array of casbin requests | 
 **permissionId** | **string** | permission id | 
 **modelId** | **string** | model id | 
 **owner** | **string** | owner | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteEnforcer

> IamObjectEnforcer IamApiControllerDeleteEnforcer(ctx, id).IamObjectEnforcer(iamObjectEnforcer).Execute()

Api Controller Delete Enforcer



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectEnforcer := *openapiclient.NewIamObjectEnforcer() // IamObjectEnforcer | The enforcer object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerDeleteEnforcer(context.Background(), id).IamObjectEnforcer(iamObjectEnforcer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerDeleteEnforcer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteEnforcer`: IamObjectEnforcer
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerDeleteEnforcer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteEnforcerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectEnforcer** | [**IamObjectEnforcer**](IamObjectEnforcer.md) | The enforcer object | 

### Return type

[**IamObjectEnforcer**](IamObjectEnforcer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteModel

> IamControllersResponse IamApiControllerDeleteModel(ctx, id).IamObjectModel(iamObjectModel).Execute()

Api Controller Delete Model



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectModel := *openapiclient.NewIamObjectModel() // IamObjectModel | The details of the model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerDeleteModel(context.Background(), id).IamObjectModel(iamObjectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerDeleteModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteModel`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerDeleteModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectModel** | [**IamObjectModel**](IamObjectModel.md) | The details of the model | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeletePermission

> IamControllersResponse IamApiControllerDeletePermission(ctx, id).IamObjectPermission(iamObjectPermission).Execute()

Api Controller Delete Permission



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectPermission := *openapiclient.NewIamObjectPermission() // IamObjectPermission | The details of the permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerDeletePermission(context.Background(), id).IamObjectPermission(iamObjectPermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerDeletePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeletePermission`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerDeletePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeletePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectPermission** | [**IamObjectPermission**](IamObjectPermission.md) | The details of the permission | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteRole

> IamControllersResponse IamApiControllerDeleteRole(ctx, id).IamObjectRole(iamObjectRole).Execute()

Api Controller Delete Role



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectRole := *openapiclient.NewIamObjectRole() // IamObjectRole | The details of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerDeleteRole(context.Background(), id).IamObjectRole(iamObjectRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerDeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteRole`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerDeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectRole** | [**IamObjectRole**](IamObjectRole.md) | The details of the role | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerEnforce

> IamControllersResponse IamApiControllerEnforce(ctx).RequestBody(requestBody).PermissionId(permissionId).ModelId(modelId).ResourceId(resourceId).Owner(owner).Execute()

Api Controller Enforce



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
	requestBody := []string{"Property_example"} // []string | Casbin request
	permissionId := "permissionId_example" // string | permission id (optional)
	modelId := "modelId_example" // string | model id (optional)
	resourceId := "resourceId_example" // string | resource id (optional)
	owner := "owner_example" // string | owner (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerEnforce(context.Background()).RequestBody(requestBody).PermissionId(permissionId).ModelId(modelId).ResourceId(resourceId).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerEnforce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerEnforce`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerEnforce`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerEnforceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Casbin request | 
 **permissionId** | **string** | permission id | 
 **modelId** | **string** | model id | 
 **resourceId** | **string** | resource id | 
 **owner** | **string** | owner | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetAllActions

> IamControllersResponse IamApiControllerGetAllActions(ctx).UserId(userId).Execute()

Api Controller Get All Actions



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
	userId := "userId_example" // string | user id like built-in/admin (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetAllActions(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetAllActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAllActions`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetAllActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetAllActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | user id like built-in/admin | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetAllObjects

> IamControllersResponse IamApiControllerGetAllObjects(ctx).UserId(userId).Execute()

Api Controller Get All Objects



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
	userId := "userId_example" // string | user id like built-in/admin (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetAllObjects(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetAllObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAllObjects`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetAllObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetAllObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | user id like built-in/admin | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetAllRoles

> IamControllersResponse IamApiControllerGetAllRoles(ctx).UserId(userId).Execute()

Api Controller Get All Roles



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
	userId := "userId_example" // string | user id like built-in/admin (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetAllRoles(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetAllRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAllRoles`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetAllRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetAllRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | user id like built-in/admin | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetEnforcer

> IamObjectEnforcer IamApiControllerGetEnforcer(ctx, id).Execute()

Api Controller Get Enforcer



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
	id := "id_example" // string | The id ( owner/name )  of enforcer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetEnforcer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetEnforcer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetEnforcer`: IamObjectEnforcer
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetEnforcer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name )  of enforcer | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetEnforcerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectEnforcer**](IamObjectEnforcer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetEnforcers

> []IamObjectEnforcer IamApiControllerGetEnforcers(ctx).Owner(owner).Execute()

Api Controller Get Enforcers



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
	owner := "owner_example" // string | The owner of enforcers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetEnforcers(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetEnforcers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetEnforcers`: []IamObjectEnforcer
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetEnforcers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetEnforcersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of enforcers | 

### Return type

[**[]IamObjectEnforcer**](IamObjectEnforcer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetFilteredPolicies

> []map[string]interface{} IamApiControllerGetFilteredPolicies(ctx).Id(id).IamObjectFilter(iamObjectFilter).Execute()

Api Controller Get Filtered Policies



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
	id := "id_example" // string | The id ( owner/name )  of enforcer
	iamObjectFilter := []openapiclient.IamObjectFilter{*openapiclient.NewIamObjectFilter()} // []IamObjectFilter | Array of filter objects for multiple filters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetFilteredPolicies(context.Background()).Id(id).IamObjectFilter(iamObjectFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetFilteredPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetFilteredPolicies`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetFilteredPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetFilteredPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name )  of enforcer | 
 **iamObjectFilter** | [**[]IamObjectFilter**](IamObjectFilter.md) | Array of filter objects for multiple filters | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetModel

> IamObjectModel IamApiControllerGetModel(ctx, id).Execute()

Api Controller Get Model



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
	id := "id_example" // string | The id ( owner/name ) of the model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetModel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetModel`: IamObjectModel
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the model | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectModel**](IamObjectModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetModels

> []IamObjectModel IamApiControllerGetModels(ctx).Owner(owner).Execute()

Api Controller Get Models



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
	owner := "owner_example" // string | The owner of models

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetModels(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetModels`: []IamObjectModel
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of models | 

### Return type

[**[]IamObjectModel**](IamObjectModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPermission

> IamObjectPermission IamApiControllerGetPermission(ctx, id).Execute()

Api Controller Get Permission



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
	id := "id_example" // string | The id ( owner/name ) of the permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetPermission(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPermission`: IamObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectPermission**](IamObjectPermission.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPermissions

> []IamObjectPermission IamApiControllerGetPermissions(ctx).Owner(owner).Execute()

Api Controller Get Permissions



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
	owner := "owner_example" // string | The owner of permissions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetPermissions(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPermissions`: []IamObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of permissions | 

### Return type

[**[]IamObjectPermission**](IamObjectPermission.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPermissionsByRole

> []IamObjectPermission IamApiControllerGetPermissionsByRole(ctx, id).Execute()

Api Controller Get Permissions By Role



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
	id := "id_example" // string | The id ( owner/name ) of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetPermissionsByRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetPermissionsByRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPermissionsByRole`: []IamObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetPermissionsByRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPermissionsByRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IamObjectPermission**](IamObjectPermission.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPermissionsBySubmitter

> []IamObjectPermission IamApiControllerGetPermissionsBySubmitter(ctx, id).Execute()

Api Controller Get Permissions By Submitter



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
	id := "id_example" // string | Resource identifier (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetPermissionsBySubmitter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetPermissionsBySubmitter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPermissionsBySubmitter`: []IamObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetPermissionsBySubmitter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPermissionsBySubmitterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IamObjectPermission**](IamObjectPermission.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPolicies

> []map[string]interface{} IamApiControllerGetPolicies(ctx).Id(id).AdapterId(adapterId).Execute()

Api Controller Get Policies



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
	id := "id_example" // string | The id ( owner/name )  of enforcer
	adapterId := "adapterId_example" // string | The adapter id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetPolicies(context.Background()).Id(id).AdapterId(adapterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPolicies`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name )  of enforcer | 
 **adapterId** | **string** | The adapter id | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetRole

> IamObjectRole IamApiControllerGetRole(ctx, id).Execute()

Api Controller Get Role



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
	id := "id_example" // string | The id ( owner/name ) of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetRole`: IamObjectRole
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectRole**](IamObjectRole.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetRoles

> []IamObjectRole IamApiControllerGetRoles(ctx).Owner(owner).Execute()

Api Controller Get Roles



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
	owner := "owner_example" // string | The owner of roles

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerGetRoles(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerGetRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetRoles`: []IamObjectRole
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerGetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of roles | 

### Return type

[**[]IamObjectRole**](IamObjectRole.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerRemovePolicy

> map[string]interface{} IamApiControllerRemovePolicy(ctx).Id(id).Body(body).Execute()

Api Controller Remove Policy



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
	id := "id_example" // string | The id ( owner/name )  of enforcer
	body := map[string]interface{}{ ... } // map[string]interface{} | The policy to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerRemovePolicy(context.Background()).Id(id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerRemovePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerRemovePolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerRemovePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerRemovePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name )  of enforcer | 
 **body** | **map[string]interface{}** | The policy to remove | 

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


## IamApiControllerRunCasbinCommand

> IamControllersResponse IamApiControllerRunCasbinCommand(ctx).Execute()

Api Controller Run Casbin Command



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
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerRunCasbinCommand(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerRunCasbinCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerRunCasbinCommand`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerRunCasbinCommand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerRunCasbinCommandRequest struct via the builder pattern


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateEnforcer

> IamObjectEnforcer IamApiControllerUpdateEnforcer(ctx, id).Body(body).Execute()

Api Controller Update Enforcer



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
	id := "id_example" // string | The id ( owner/name )  of enforcer
	body := map[string]interface{}{ ... } // map[string]interface{} | The enforcer object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerUpdateEnforcer(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerUpdateEnforcer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateEnforcer`: IamObjectEnforcer
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerUpdateEnforcer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name )  of enforcer | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateEnforcerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | The enforcer object | 

### Return type

[**IamObjectEnforcer**](IamObjectEnforcer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateModel

> IamControllersResponse IamApiControllerUpdateModel(ctx, id).IamObjectModel(iamObjectModel).Execute()

Api Controller Update Model



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
	id := "id_example" // string | The id ( owner/name ) of the model
	iamObjectModel := *openapiclient.NewIamObjectModel() // IamObjectModel | The details of the model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerUpdateModel(context.Background(), id).IamObjectModel(iamObjectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerUpdateModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateModel`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerUpdateModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the model | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectModel** | [**IamObjectModel**](IamObjectModel.md) | The details of the model | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdatePermission

> IamControllersResponse IamApiControllerUpdatePermission(ctx, id).IamObjectPermission(iamObjectPermission).Execute()

Api Controller Update Permission



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
	id := "id_example" // string | The id ( owner/name ) of the permission
	iamObjectPermission := *openapiclient.NewIamObjectPermission() // IamObjectPermission | The details of the permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerUpdatePermission(context.Background(), id).IamObjectPermission(iamObjectPermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerUpdatePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdatePermission`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerUpdatePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectPermission** | [**IamObjectPermission**](IamObjectPermission.md) | The details of the permission | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdatePolicy

> map[string]interface{} IamApiControllerUpdatePolicy(ctx, id).RequestBody(requestBody).Execute()

Api Controller Update Policy



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
	id := "id_example" // string | The id ( owner/name )  of enforcer
	requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Array containing old and new policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerUpdatePolicy(context.Background(), id).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerUpdatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdatePolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerUpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name )  of enforcer | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** | Array containing old and new policy | 

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


## IamApiControllerUpdateRole

> IamControllersResponse IamApiControllerUpdateRole(ctx, id).IamObjectRole(iamObjectRole).Execute()

Api Controller Update Role



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
	id := "id_example" // string | The id ( owner/name ) of the role
	iamObjectRole := *openapiclient.NewIamObjectRole() // IamObjectRole | The details of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesPermissionsAPI.IamApiControllerUpdateRole(context.Background(), id).IamObjectRole(iamObjectRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesPermissionsAPI.IamApiControllerUpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateRole`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesPermissionsAPI.IamApiControllerUpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectRole** | [**IamObjectRole**](IamObjectRole.md) | The details of the role | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

