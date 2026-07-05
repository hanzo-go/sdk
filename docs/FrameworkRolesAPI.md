# \FrameworkRolesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FrameworkAssignRole**](FrameworkRolesAPI.md#FrameworkAssignRole) | **Post** /v1/framework/roles | Assign a role to a user
[**FrameworkListRoles**](FrameworkRolesAPI.md#FrameworkListRoles) | **Get** /v1/framework/roles | List per-org role assignments
[**FrameworkRevokeRole**](FrameworkRolesAPI.md#FrameworkRevokeRole) | **Delete** /v1/framework/roles/{user}/{role} | Revoke a role from a user



## FrameworkAssignRole

> FrameworkRole FrameworkAssignRole(ctx).FrameworkRole(frameworkRole).Execute()

Assign a role to a user

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
	frameworkRole := *openapiclient.NewFrameworkRole("User_example", "Role_example") // FrameworkRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkRolesAPI.FrameworkAssignRole(context.Background()).FrameworkRole(frameworkRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkRolesAPI.FrameworkAssignRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkAssignRole`: FrameworkRole
	fmt.Fprintf(os.Stdout, "Response from `FrameworkRolesAPI.FrameworkAssignRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkAssignRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameworkRole** | [**FrameworkRole**](FrameworkRole.md) |  | 

### Return type

[**FrameworkRole**](FrameworkRole.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkListRoles

> FrameworkListRoles200Response FrameworkListRoles(ctx).Execute()

List per-org role assignments

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
	resp, r, err := apiClient.FrameworkRolesAPI.FrameworkListRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkRolesAPI.FrameworkListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkListRoles`: FrameworkListRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameworkRolesAPI.FrameworkListRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkListRolesRequest struct via the builder pattern


### Return type

[**FrameworkListRoles200Response**](FrameworkListRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkRevokeRole

> FrameworkRevokeRole(ctx, user, role).Execute()

Revoke a role from a user

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
	user := "user_example" // string | 
	role := "role_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkRolesAPI.FrameworkRevokeRole(context.Background(), user, role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkRolesAPI.FrameworkRevokeRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** |  | 
**role** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkRevokeRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

