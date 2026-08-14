# \CompatAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIamApplication**](CompatAPI.md#DeleteIamApplication) | **Delete** /v1/iam/application | Removes an application.
[**GetIamApplication**](CompatAPI.md#GetIamApplication) | **Get** /v1/iam/application | Returns one application: its sign-in methods, its allowed redirect URIs and the client credentials your integration authenticates with.
[**PostIamAddApplication**](CompatAPI.md#PostIamAddApplication) | **Post** /v1/iam/add-application | Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.
[**PostIamAddOrganization**](CompatAPI.md#PostIamAddOrganization) | **Post** /v1/iam/add-organization | Creates an organization — the account everything else in your directory hangs from.
[**PostIamAddProject**](CompatAPI.md#PostIamAddProject) | **Post** /v1/iam/add-project | Creates a project inside your organization — the scope people pick between when their work is separated by product or client rather than by team.
[**PostIamAddProvider**](CompatAPI.md#PostIamAddProvider) | **Post** /v1/iam/add-provider | Adds an identity provider your people can sign in with, or a service your applications send through — a social or enterprise login, an email or SMS sender, a storage or payment connector.
[**PostIamAddRole**](CompatAPI.md#PostIamAddRole) | **Post** /v1/iam/add-role | Creates a role — a named group of people that permissions are granted to.
[**PostIamAddUser**](CompatAPI.md#PostIamAddUser) | **Post** /v1/iam/add-user | Adds a person to your organization and, if you send a password, sets the one they will sign in with.
[**PostIamAddWorkspace**](CompatAPI.md#PostIamAddWorkspace) | **Post** /v1/iam/add-workspace | Creates a workspace inside your organization — the scope a team works in, alongside projects rather than instead of them.
[**PostIamApplication**](CompatAPI.md#PostIamApplication) | **Post** /v1/iam/application | Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.
[**PostIamDeleteApplication**](CompatAPI.md#PostIamDeleteApplication) | **Post** /v1/iam/delete-application | Deletes an application.
[**PostIamDeleteOrganization**](CompatAPI.md#PostIamDeleteOrganization) | **Post** /v1/iam/delete-organization | Deletes an organization and everything named inside it — its users, applications, roles, projects and workspaces.
[**PostIamDeleteProject**](CompatAPI.md#PostIamDeleteProject) | **Post** /v1/iam/delete-project | Deletes a project.
[**PostIamDeleteProvider**](CompatAPI.md#PostIamDeleteProvider) | **Post** /v1/iam/delete-provider | Removes a provider.
[**PostIamDeleteRole**](CompatAPI.md#PostIamDeleteRole) | **Post** /v1/iam/delete-role | Deletes a role.
[**PostIamDeleteUser**](CompatAPI.md#PostIamDeleteUser) | **Post** /v1/iam/delete-user | Removes a person from your organization.
[**PostIamDeleteWorkspace**](CompatAPI.md#PostIamDeleteWorkspace) | **Post** /v1/iam/delete-workspace | Deletes a workspace.
[**PostIamUpdateApplication**](CompatAPI.md#PostIamUpdateApplication) | **Post** /v1/iam/update-application | Updates one of your applications — its display, its sign-in methods and the redirect URIs it is allowed to return to.
[**PostIamUpdateOrganization**](CompatAPI.md#PostIamUpdateOrganization) | **Post** /v1/iam/update-organization | Updates your organization — its display, its default settings and the sign-in rules everyone in it inherits.
[**PostIamUpdateProvider**](CompatAPI.md#PostIamUpdateProvider) | **Post** /v1/iam/update-provider | Updates a provider&#39;s settings or rotates the credentials it holds.
[**PostIamUpdateRole**](CompatAPI.md#PostIamUpdateRole) | **Post** /v1/iam/update-role | Updates a role&#39;s members or the roles it includes.
[**PostIamUpdateUser**](CompatAPI.md#PostIamUpdateUser) | **Post** /v1/iam/update-user | Updates one of your users&#39; profile, roles or credentials.
[**PutIamApplication**](CompatAPI.md#PutIamApplication) | **Put** /v1/iam/application | Changes an application&#39;s display, its sign-in methods and the redirect URIs it may return to — the call that makes login work from a new host.



## DeleteIamApplication

> IamDeleteResult DeleteIamApplication(ctx).Owner(owner).Name(name).Execute()

Removes an application.



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
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.DeleteIamApplication(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.DeleteIamApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIamApplication`: IamDeleteResult
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.DeleteIamApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamDeleteResult**](IamDeleteResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamApplication

> IamApplication GetIamApplication(ctx).Owner(owner).Name(name).Execute()

Returns one application: its sign-in methods, its allowed redirect URIs and the client credentials your integration authenticates with.



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
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.GetIamApplication(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.GetIamApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIamApplication`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.GetIamApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIamApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddApplication

> IamResponse PostIamAddApplication(ctx).IamApplication(iamApplication).Execute()

Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.



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
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamAddApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamAddApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddApplication`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamAddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddOrganization

> IamResponse PostIamAddOrganization(ctx).IamCreateOrganizationInput(iamCreateOrganizationInput).Execute()

Creates an organization — the account everything else in your directory hangs from.



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
	iamCreateOrganizationInput := *openapiclient.NewIamCreateOrganizationInput() // IamCreateOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamAddOrganization(context.Background()).IamCreateOrganizationInput(iamCreateOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamAddOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddOrganization`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamAddOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamCreateOrganizationInput** | [**IamCreateOrganizationInput**](IamCreateOrganizationInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddProject

> IamResponse PostIamAddProject(ctx).IamInput(iamInput).Execute()

Creates a project inside your organization — the scope people pick between when their work is separated by product or client rather than by team.



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
	iamInput := *openapiclient.NewIamInput() // IamInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamAddProject(context.Background()).IamInput(iamInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamAddProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddProject`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamAddProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamInput** | [**IamInput**](IamInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddProvider

> IamResponse PostIamAddProvider(ctx).IamProvider(iamProvider).Execute()

Adds an identity provider your people can sign in with, or a service your applications send through — a social or enterprise login, an email or SMS sender, a storage or payment connector.



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
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamAddProvider(context.Background()).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamAddProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddProvider`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamAddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProvider** | [**IamProvider**](IamProvider.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddRole

> IamResponse PostIamAddRole(ctx).IamRolesInput(iamRolesInput).Execute()

Creates a role — a named group of people that permissions are granted to.



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
	iamRolesInput := *openapiclient.NewIamRolesInput() // IamRolesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamAddRole(context.Background()).IamRolesInput(iamRolesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddRole`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamAddRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesInput** | [**IamRolesInput**](IamRolesInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddUser

> IamResponse PostIamAddUser(ctx).IamUserBody(iamUserBody).Execute()

Adds a person to your organization and, if you send a password, sets the one they will sign in with.



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
	iamUserBody := *openapiclient.NewIamUserBody() // IamUserBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamAddUser(context.Background()).IamUserBody(iamUserBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamAddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddUser`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamAddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUserBody** | [**IamUserBody**](IamUserBody.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamAddWorkspace

> IamResponse PostIamAddWorkspace(ctx).IamWorkspacesInput(iamWorkspacesInput).Execute()

Creates a workspace inside your organization — the scope a team works in, alongside projects rather than instead of them.



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
	iamWorkspacesInput := *openapiclient.NewIamWorkspacesInput() // IamWorkspacesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamAddWorkspace(context.Background()).IamWorkspacesInput(iamWorkspacesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamAddWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamAddWorkspace`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamAddWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamAddWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWorkspacesInput** | [**IamWorkspacesInput**](IamWorkspacesInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamApplication

> IamApplication PostIamApplication(ctx).IamApplication(iamApplication).Execute()

Registers an application in your organization — one product or site your people sign in to, with its own client credentials, sign-in methods and allowed redirect URIs.



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
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamApplication`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteApplication

> IamResponse PostIamDeleteApplication(ctx).IamApplication(iamApplication).Execute()

Deletes an application.



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
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamDeleteApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamDeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteApplication`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamDeleteApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteOrganization

> IamResponse PostIamDeleteOrganization(ctx).IamDeleteOrganizationInput(iamDeleteOrganizationInput).Execute()

Deletes an organization and everything named inside it — its users, applications, roles, projects and workspaces.



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
	iamDeleteOrganizationInput := *openapiclient.NewIamDeleteOrganizationInput() // IamDeleteOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamDeleteOrganization(context.Background()).IamDeleteOrganizationInput(iamDeleteOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamDeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteOrganization`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamDeleteOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamDeleteOrganizationInput** | [**IamDeleteOrganizationInput**](IamDeleteOrganizationInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteProject

> IamResponse PostIamDeleteProject(ctx).IamProjectsRef(iamProjectsRef).Execute()

Deletes a project.



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
	iamProjectsRef := *openapiclient.NewIamProjectsRef() // IamProjectsRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamDeleteProject(context.Background()).IamProjectsRef(iamProjectsRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteProject`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamDeleteProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProjectsRef** | [**IamProjectsRef**](IamProjectsRef.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteProvider

> IamResponse PostIamDeleteProvider(ctx).IamProvider(iamProvider).Execute()

Removes a provider.



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
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamDeleteProvider(context.Background()).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamDeleteProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteProvider`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamDeleteProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProvider** | [**IamProvider**](IamProvider.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteRole

> IamResponse PostIamDeleteRole(ctx).IamRolesRef(iamRolesRef).Execute()

Deletes a role.



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
	iamRolesRef := *openapiclient.NewIamRolesRef() // IamRolesRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamDeleteRole(context.Background()).IamRolesRef(iamRolesRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamDeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteRole`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamDeleteRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesRef** | [**IamRolesRef**](IamRolesRef.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteUser

> IamResponse PostIamDeleteUser(ctx).IamUserBody(iamUserBody).Execute()

Removes a person from your organization.



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
	iamUserBody := *openapiclient.NewIamUserBody() // IamUserBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamDeleteUser(context.Background()).IamUserBody(iamUserBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteUser`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamDeleteUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUserBody** | [**IamUserBody**](IamUserBody.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamDeleteWorkspace

> IamResponse PostIamDeleteWorkspace(ctx).IamWorkspacesRef(iamWorkspacesRef).Execute()

Deletes a workspace.



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
	iamWorkspacesRef := *openapiclient.NewIamWorkspacesRef() // IamWorkspacesRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamDeleteWorkspace(context.Background()).IamWorkspacesRef(iamWorkspacesRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamDeleteWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamDeleteWorkspace`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamDeleteWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamDeleteWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamWorkspacesRef** | [**IamWorkspacesRef**](IamWorkspacesRef.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateApplication

> IamResponse PostIamUpdateApplication(ctx).IamApplication(iamApplication).Execute()

Updates one of your applications — its display, its sign-in methods and the redirect URIs it is allowed to return to.



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
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamUpdateApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamUpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateApplication`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamUpdateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateOrganization

> IamResponse PostIamUpdateOrganization(ctx).IamUpdateOrganizationInput(iamUpdateOrganizationInput).Execute()

Updates your organization — its display, its default settings and the sign-in rules everyone in it inherits.



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
	iamUpdateOrganizationInput := *openapiclient.NewIamUpdateOrganizationInput() // IamUpdateOrganizationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamUpdateOrganization(context.Background()).IamUpdateOrganizationInput(iamUpdateOrganizationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamUpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateOrganization`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamUpdateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUpdateOrganizationInput** | [**IamUpdateOrganizationInput**](IamUpdateOrganizationInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateProvider

> IamResponse PostIamUpdateProvider(ctx).IamProvider(iamProvider).Execute()

Updates a provider's settings or rotates the credentials it holds.



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
	iamProvider := *openapiclient.NewIamProvider() // IamProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamUpdateProvider(context.Background()).IamProvider(iamProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamUpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateProvider`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamUpdateProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProvider** | [**IamProvider**](IamProvider.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateRole

> IamResponse PostIamUpdateRole(ctx).IamRolesInput(iamRolesInput).Execute()

Updates a role's members or the roles it includes.



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
	iamRolesInput := *openapiclient.NewIamRolesInput() // IamRolesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamUpdateRole(context.Background()).IamRolesInput(iamRolesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamUpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateRole`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamUpdateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamRolesInput** | [**IamRolesInput**](IamRolesInput.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIamUpdateUser

> IamResponse PostIamUpdateUser(ctx).IamUserBody(iamUserBody).Execute()

Updates one of your users' profile, roles or credentials.



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
	iamUserBody := *openapiclient.NewIamUserBody() // IamUserBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PostIamUpdateUser(context.Background()).IamUserBody(iamUserBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PostIamUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIamUpdateUser`: IamResponse
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PostIamUpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIamUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamUserBody** | [**IamUserBody**](IamUserBody.md) |  | 

### Return type

[**IamResponse**](IamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIamApplication

> IamApplication PutIamApplication(ctx).IamApplication(iamApplication).Execute()

Changes an application's display, its sign-in methods and the redirect URIs it may return to — the call that makes login work from a new host.



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
	iamApplication := *openapiclient.NewIamApplication() // IamApplication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompatAPI.PutIamApplication(context.Background()).IamApplication(iamApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompatAPI.PutIamApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIamApplication`: IamApplication
	fmt.Fprintf(os.Stdout, "Response from `CompatAPI.PutIamApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIamApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamApplication** | [**IamApplication**](IamApplication.md) |  | 

### Return type

[**IamApplication**](IamApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

