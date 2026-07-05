# \AdminIdentityAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminListApplications**](AdminIdentityAPI.md#AdminAdminListApplications) | **Get** /v1/admin/applications | IAM applications (verbatim passthrough)
[**AdminAdminListOrgs**](AdminIdentityAPI.md#AdminAdminListOrgs) | **Get** /v1/admin/orgs | Tenant directory
[**AdminAdminListRoles**](AdminIdentityAPI.md#AdminAdminListRoles) | **Get** /v1/admin/roles | IAM roles (verbatim passthrough)
[**AdminAdminListUsers**](AdminIdentityAPI.md#AdminAdminListUsers) | **Get** /v1/admin/users | Cross-org user directory
[**AdminAdminMe**](AdminIdentityAPI.md#AdminAdminMe) | **Get** /v1/admin/me | Validated operator identity



## AdminAdminListApplications

> AdminRawList AdminAdminListApplications(ctx).Owner(owner).P(p).PageSize(pageSize).Execute()

IAM applications (verbatim passthrough)

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
	owner := "owner_example" // string |  (optional)
	p := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminIdentityAPI.AdminAdminListApplications(context.Background()).Owner(owner).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminIdentityAPI.AdminAdminListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminListApplications`: AdminRawList
	fmt.Fprintf(os.Stdout, "Response from `AdminIdentityAPI.AdminAdminListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **p** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**AdminRawList**](AdminRawList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminListOrgs

> AdminOrgList AdminAdminListOrgs(ctx).Execute()

Tenant directory

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
	resp, r, err := apiClient.AdminIdentityAPI.AdminAdminListOrgs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminIdentityAPI.AdminAdminListOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminListOrgs`: AdminOrgList
	fmt.Fprintf(os.Stdout, "Response from `AdminIdentityAPI.AdminAdminListOrgs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminListOrgsRequest struct via the builder pattern


### Return type

[**AdminOrgList**](AdminOrgList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminListRoles

> AdminRawList AdminAdminListRoles(ctx).Owner(owner).P(p).PageSize(pageSize).Execute()

IAM roles (verbatim passthrough)

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
	owner := "owner_example" // string |  (optional)
	p := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminIdentityAPI.AdminAdminListRoles(context.Background()).Owner(owner).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminIdentityAPI.AdminAdminListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminListRoles`: AdminRawList
	fmt.Fprintf(os.Stdout, "Response from `AdminIdentityAPI.AdminAdminListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **p** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**AdminRawList**](AdminRawList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminListUsers

> AdminUserList AdminAdminListUsers(ctx).Org(org).P(p).PageSize(pageSize).Q(q).Execute()

Cross-org user directory

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
	org := "org_example" // string | Filter to one org (owner) (optional)
	p := int32(56) // int32 | 1-based page (optional)
	pageSize := int32(56) // int32 |  (optional)
	q := "q_example" // string | Free-text name filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminIdentityAPI.AdminAdminListUsers(context.Background()).Org(org).P(p).PageSize(pageSize).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminIdentityAPI.AdminAdminListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminListUsers`: AdminUserList
	fmt.Fprintf(os.Stdout, "Response from `AdminIdentityAPI.AdminAdminListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Filter to one org (owner) | 
 **p** | **int32** | 1-based page | 
 **pageSize** | **int32** |  | 
 **q** | **string** | Free-text name filter | 

### Return type

[**AdminUserList**](AdminUserList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminMe

> AdminAdminMe200Response AdminAdminMe(ctx).Execute()

Validated operator identity

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
	resp, r, err := apiClient.AdminIdentityAPI.AdminAdminMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminIdentityAPI.AdminAdminMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminMe`: AdminAdminMe200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminIdentityAPI.AdminAdminMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminMeRequest struct via the builder pattern


### Return type

[**AdminAdminMe200Response**](AdminAdminMe200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

