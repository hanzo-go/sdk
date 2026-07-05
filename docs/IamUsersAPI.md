# \IamUsersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddLdap**](IamUsersAPI.md#IamApiControllerAddLdap) | **Post** /v1/iam/ldaps | Api Controller Add Ldap
[**IamApiControllerAddUser**](IamUsersAPI.md#IamApiControllerAddUser) | **Post** /v1/iam/users | Api Controller Add User
[**IamApiControllerAddUserKeys**](IamUsersAPI.md#IamApiControllerAddUserKeys) | **Post** /v1/iam/user-keys | Api Controller Add User Keys
[**IamApiControllerCheckUserPassword**](IamUsersAPI.md#IamApiControllerCheckUserPassword) | **Post** /v1/iam/auth/check-password | Api Controller Check User Password
[**IamApiControllerDeleteLdap**](IamUsersAPI.md#IamApiControllerDeleteLdap) | **Delete** /v1/iam/ldaps/{id} | Api Controller Delete Ldap
[**IamApiControllerDeleteUser**](IamUsersAPI.md#IamApiControllerDeleteUser) | **Delete** /v1/iam/users/{id} | Api Controller Delete User
[**IamApiControllerExitImpersonateUser**](IamUsersAPI.md#IamApiControllerExitImpersonateUser) | **Post** /v1/iam/impersonation/exit | Api Controller Exit Impersonate User
[**IamApiControllerGetAccount**](IamUsersAPI.md#IamApiControllerGetAccount) | **Get** /v1/iam/accounts/{id} | Api Controller Get Account
[**IamApiControllerGetEmailAndPhone**](IamUsersAPI.md#IamApiControllerGetEmailAndPhone) | **Get** /v1/iam/auth/contact | Api Controller Get Email And Phone
[**IamApiControllerGetGlobalUsers**](IamUsersAPI.md#IamApiControllerGetGlobalUsers) | **Get** /v1/iam/global-users | Api Controller Get Global Users
[**IamApiControllerGetLdap**](IamUsersAPI.md#IamApiControllerGetLdap) | **Get** /v1/iam/ldaps/{id} | Api Controller Get Ldap
[**IamApiControllerGetLdaps**](IamUsersAPI.md#IamApiControllerGetLdaps) | **Get** /v1/iam/ldaps | Api Controller Get Ldaps
[**IamApiControllerGetLdapser**](IamUsersAPI.md#IamApiControllerGetLdapser) | **Get** /v1/iam/ldap-users | Api Controller Get Ldapser
[**IamApiControllerGetSortedUsers**](IamUsersAPI.md#IamApiControllerGetSortedUsers) | **Get** /v1/iam/sorted-users | Api Controller Get Sorted Users
[**IamApiControllerGetUser**](IamUsersAPI.md#IamApiControllerGetUser) | **Get** /v1/iam/users/{id} | Api Controller Get User
[**IamApiControllerGetUserCount**](IamUsersAPI.md#IamApiControllerGetUserCount) | **Get** /v1/iam/user-counts/{id} | Api Controller Get User Count
[**IamApiControllerGetUsers**](IamUsersAPI.md#IamApiControllerGetUsers) | **Get** /v1/iam/users | Api Controller Get Users
[**IamApiControllerImpersonateUser**](IamUsersAPI.md#IamApiControllerImpersonateUser) | **Post** /v1/iam/impersonation-user | Api Controller Impersonate User
[**IamApiControllerResetEmailOrPhone**](IamUsersAPI.md#IamApiControllerResetEmailOrPhone) | **Post** /v1/iam/auth/reset-contact | Api Controller Reset Email Or Phone
[**IamApiControllerSetPassword**](IamUsersAPI.md#IamApiControllerSetPassword) | **Post** /v1/iam/auth/set-password | Api Controller Set Password
[**IamApiControllerSyncLdapUsers**](IamUsersAPI.md#IamApiControllerSyncLdapUsers) | **Post** /v1/iam/ldap/sync | Api Controller Sync Ldap Users
[**IamApiControllerUpdateLdap**](IamUsersAPI.md#IamApiControllerUpdateLdap) | **Put** /v1/iam/ldaps/{id} | Api Controller Update Ldap
[**IamApiControllerUpdateUser**](IamUsersAPI.md#IamApiControllerUpdateUser) | **Put** /v1/iam/users/{id} | Api Controller Update User
[**IamApiControllerUserInfo**](IamUsersAPI.md#IamApiControllerUserInfo) | **Get** /oauth/userinfo | Api Controller User Info
[**IamApiControllerUserInfo2**](IamUsersAPI.md#IamApiControllerUserInfo2) | **Get** /v1/iam/user | Api Controller User Info2
[**IamApiControllerVerifyIdentification**](IamUsersAPI.md#IamApiControllerVerifyIdentification) | **Post** /v1/iam/auth/identification/verify | Api Controller Verify Identification
[**IamApiControllerWebAuthnSignupBegin**](IamUsersAPI.md#IamApiControllerWebAuthnSignupBegin) | **Get** /v1/iam/auth/webauthn/signup/begin | Api Controller Web Authn Signup Begin
[**IamApiControllerWebAuthnSignupFinish**](IamUsersAPI.md#IamApiControllerWebAuthnSignupFinish) | **Post** /v1/iam/auth/webauthn/signup/finish | Api Controller Web Authn Signup Finish



## IamApiControllerAddLdap

> IamControllersResponse IamApiControllerAddLdap(ctx).IamObjectLdap(iamObjectLdap).Execute()

Api Controller Add Ldap



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
	iamObjectLdap := *openapiclient.NewIamObjectLdap() // IamObjectLdap | The details of the ldap

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerAddLdap(context.Background()).IamObjectLdap(iamObjectLdap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerAddLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddLdap`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerAddLdap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectLdap** | [**IamObjectLdap**](IamObjectLdap.md) | The details of the ldap | 

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


## IamApiControllerAddUser

> IamControllersResponse IamApiControllerAddUser(ctx).IamObjectUser(iamObjectUser).Execute()

Api Controller Add User



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
	iamObjectUser := *openapiclient.NewIamObjectUser() // IamObjectUser | The details of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerAddUser(context.Background()).IamObjectUser(iamObjectUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerAddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerAddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectUser** | [**IamObjectUser**](IamObjectUser.md) | The details of the user | 

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


## IamApiControllerAddUserKeys

> IamObjectUserinfo IamApiControllerAddUserKeys(ctx).Execute()

Api Controller Add User Keys

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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerAddUserKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerAddUserKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddUserKeys`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerAddUserKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddUserKeysRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerCheckUserPassword

> IamObjectUserinfo IamApiControllerCheckUserPassword(ctx).Execute()

Api Controller Check User Password

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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerCheckUserPassword(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerCheckUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerCheckUserPassword`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerCheckUserPassword`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerCheckUserPasswordRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteLdap

> IamControllersResponse IamApiControllerDeleteLdap(ctx, id).IamObjectLdap(iamObjectLdap).Execute()

Api Controller Delete Ldap



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
	iamObjectLdap := *openapiclient.NewIamObjectLdap() // IamObjectLdap | The details of the ldap

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerDeleteLdap(context.Background(), id).IamObjectLdap(iamObjectLdap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerDeleteLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteLdap`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerDeleteLdap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectLdap** | [**IamObjectLdap**](IamObjectLdap.md) | The details of the ldap | 

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


## IamApiControllerDeleteUser

> IamControllersResponse IamApiControllerDeleteUser(ctx, id).IamObjectUser(iamObjectUser).Execute()

Api Controller Delete User



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
	iamObjectUser := *openapiclient.NewIamObjectUser() // IamObjectUser | The details of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerDeleteUser(context.Background(), id).IamObjectUser(iamObjectUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerDeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectUser** | [**IamObjectUser**](IamObjectUser.md) | The details of the user | 

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


## IamApiControllerExitImpersonateUser

> IamControllersResponse IamApiControllerExitImpersonateUser(ctx).Execute()

Api Controller Exit Impersonate User



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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerExitImpersonateUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerExitImpersonateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerExitImpersonateUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerExitImpersonateUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerExitImpersonateUserRequest struct via the builder pattern


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


## IamApiControllerGetAccount

> IamControllersResponse IamApiControllerGetAccount(ctx, id).Execute()

Api Controller Get Account



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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAccount`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## IamApiControllerGetEmailAndPhone

> IamControllersResponse IamApiControllerGetEmailAndPhone(ctx).Username(username).Organization(organization).Execute()

Api Controller Get Email And Phone



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
	username := "username_example" // string | The username of the user
	organization := "organization_example" // string | The organization of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetEmailAndPhone(context.Background()).Username(username).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetEmailAndPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetEmailAndPhone`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetEmailAndPhone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetEmailAndPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | The username of the user | 
 **organization** | **string** | The organization of the user | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetGlobalUsers

> []IamObjectUser IamApiControllerGetGlobalUsers(ctx).Execute()

Api Controller Get Global Users



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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetGlobalUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetGlobalUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGlobalUsers`: []IamObjectUser
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetGlobalUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetGlobalUsersRequest struct via the builder pattern


### Return type

[**[]IamObjectUser**](IamObjectUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetLdap

> IamObjectLdap IamApiControllerGetLdap(ctx, id).Execute()

Api Controller Get Ldap



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
	id := "id_example" // string | id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetLdap(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetLdap`: IamObjectLdap
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetLdap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectLdap**](IamObjectLdap.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetLdaps

> []IamObjectLdap IamApiControllerGetLdaps(ctx).Owner(owner).Execute()

Api Controller Get Ldaps



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
	owner := "owner_example" // string | owner (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetLdaps(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetLdaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetLdaps`: []IamObjectLdap
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetLdaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetLdapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | owner | 

### Return type

[**[]IamObjectLdap**](IamObjectLdap.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetLdapser

> IamControllersLdapResp IamApiControllerGetLdapser(ctx).Execute()

Api Controller Get Ldapser



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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetLdapser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetLdapser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetLdapser`: IamControllersLdapResp
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetLdapser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetLdapserRequest struct via the builder pattern


### Return type

[**IamControllersLdapResp**](IamControllersLdapResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSortedUsers

> []IamObjectUser IamApiControllerGetSortedUsers(ctx).Owner(owner).Sorter(sorter).Limit(limit).Execute()

Api Controller Get Sorted Users

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
	owner := "owner_example" // string | The owner of users
	sorter := "sorter_example" // string | The DB column name to sort by, e.g., created_time
	limit := "limit_example" // string | The count of users to return, e.g., 25

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetSortedUsers(context.Background()).Owner(owner).Sorter(sorter).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetSortedUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSortedUsers`: []IamObjectUser
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetSortedUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSortedUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of users | 
 **sorter** | **string** | The DB column name to sort by, e.g., created_time | 
 **limit** | **string** | The count of users to return, e.g., 25 | 

### Return type

[**[]IamObjectUser**](IamObjectUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetUser

> IamObjectUser IamApiControllerGetUser(ctx, id).Owner(owner).Email(email).Phone(phone).UserId(userId).Execute()

Api Controller Get User



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
	id := "id_example" // string | The id ( owner/name ) of the user
	owner := "owner_example" // string | The owner of the user (optional)
	email := "email_example" // string | The email of the user (optional)
	phone := "phone_example" // string | The phone of the user (optional)
	userId := "userId_example" // string | The userId of the user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetUser(context.Background(), id).Owner(owner).Email(email).Phone(phone).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUser`: IamObjectUser
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **owner** | **string** | The owner of the user | 
 **email** | **string** | The email of the user | 
 **phone** | **string** | The phone of the user | 
 **userId** | **string** | The userId of the user | 

### Return type

[**IamObjectUser**](IamObjectUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetUserCount

> map[string]interface{} IamApiControllerGetUserCount(ctx, id).Owner(owner).IsOnline(isOnline).Execute()

Api Controller Get User Count

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
	owner := "owner_example" // string | The owner of users
	isOnline := "isOnline_example" // string | The filter for query, 1 for online, 0 for offline, empty string for all users
	id := "id_example" // string | Resource identifier (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetUserCount(context.Background(), id).Owner(owner).IsOnline(isOnline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetUserCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUserCount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetUserCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetUserCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of users | 
 **isOnline** | **string** | The filter for query, 1 for online, 0 for offline, empty string for all users | 


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


## IamApiControllerGetUsers

> []IamObjectUser IamApiControllerGetUsers(ctx).Owner(owner).Execute()

Api Controller Get Users

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
	owner := "owner_example" // string | The owner of users

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerGetUsers(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUsers`: []IamObjectUser
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerGetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of users | 

### Return type

[**[]IamObjectUser**](IamObjectUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerImpersonateUser

> IamControllersResponse IamApiControllerImpersonateUser(ctx).Username(username).Execute()

Api Controller Impersonate User



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
	username := "username_example" // string | The username to impersonate (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerImpersonateUser(context.Background()).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerImpersonateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerImpersonateUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerImpersonateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerImpersonateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | The username to impersonate (owner/name) | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerResetEmailOrPhone

> IamObjectUserinfo IamApiControllerResetEmailOrPhone(ctx).Execute()

Api Controller Reset Email Or Phone

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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerResetEmailOrPhone(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerResetEmailOrPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerResetEmailOrPhone`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerResetEmailOrPhone`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerResetEmailOrPhoneRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSetPassword

> IamControllersResponse IamApiControllerSetPassword(ctx).UserOwner(userOwner).UserName(userName).OldPassword(oldPassword).NewPassword(newPassword).Execute()

Api Controller Set Password



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
	userOwner := "userOwner_example" // string | The owner of the user
	userName := "userName_example" // string | The name of the user
	oldPassword := "oldPassword_example" // string | The old password of the user
	newPassword := "newPassword_example" // string | The new password of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerSetPassword(context.Background()).UserOwner(userOwner).UserName(userName).OldPassword(oldPassword).NewPassword(newPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerSetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSetPassword`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerSetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userOwner** | **string** | The owner of the user | 
 **userName** | **string** | The name of the user | 
 **oldPassword** | **string** | The old password of the user | 
 **newPassword** | **string** | The new password of the user | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSyncLdapUsers

> IamControllersLdapSyncResp IamApiControllerSyncLdapUsers(ctx).Id(id).Execute()

Api Controller Sync Ldap Users



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
	id := "id_example" // string | id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerSyncLdapUsers(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerSyncLdapUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSyncLdapUsers`: IamControllersLdapSyncResp
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerSyncLdapUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSyncLdapUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | id | 

### Return type

[**IamControllersLdapSyncResp**](IamControllersLdapSyncResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateLdap

> IamControllersResponse IamApiControllerUpdateLdap(ctx, id).IamObjectLdap(iamObjectLdap).Execute()

Api Controller Update Ldap



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
	iamObjectLdap := *openapiclient.NewIamObjectLdap() // IamObjectLdap | The details of the ldap

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerUpdateLdap(context.Background(), id).IamObjectLdap(iamObjectLdap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerUpdateLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateLdap`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerUpdateLdap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectLdap** | [**IamObjectLdap**](IamObjectLdap.md) | The details of the ldap | 

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


## IamApiControllerUpdateUser

> IamControllersResponse IamApiControllerUpdateUser(ctx, id).IamObjectUser(iamObjectUser).UserId(userId).Owner(owner).Execute()

Api Controller Update User



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
	id := "id_example" // string | The id ( owner/name ) of the user
	iamObjectUser := *openapiclient.NewIamObjectUser() // IamObjectUser | The details of the user
	userId := "userId_example" // string | The userId (UUID) of the user (optional)
	owner := "owner_example" // string | The owner of the user (required when using userId) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerUpdateUser(context.Background(), id).IamObjectUser(iamObjectUser).UserId(userId).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerUpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectUser** | [**IamObjectUser**](IamObjectUser.md) | The details of the user | 
 **userId** | **string** | The userId (UUID) of the user | 
 **owner** | **string** | The owner of the user (required when using userId) | 

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


## IamApiControllerUserInfo

> IamObjectUserinfo IamApiControllerUserInfo(ctx).Execute()

Api Controller User Info



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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUserInfo`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUserInfoRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUserInfo2

> IamControllersLaravelResponse IamApiControllerUserInfo2(ctx).Execute()

Api Controller User Info2



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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerUserInfo2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerUserInfo2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUserInfo2`: IamControllersLaravelResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerUserInfo2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUserInfo2Request struct via the builder pattern


### Return type

[**IamControllersLaravelResponse**](IamControllersLaravelResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerVerifyIdentification

> IamControllersResponse IamApiControllerVerifyIdentification(ctx).Owner(owner).Name(name).Provider(provider).Execute()

Api Controller Verify Identification



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
	owner := "owner_example" // string | The owner of the user (optional, defaults to logged-in user) (optional)
	name := "name_example" // string | The name of the user (optional, defaults to logged-in user) (optional)
	provider := "provider_example" // string | The name of the ID Verification provider (optional, auto-selected if not provided) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerVerifyIdentification(context.Background()).Owner(owner).Name(name).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerVerifyIdentification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerVerifyIdentification`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerVerifyIdentification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerVerifyIdentificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the user (optional, defaults to logged-in user) | 
 **name** | **string** | The name of the user (optional, defaults to logged-in user) | 
 **provider** | **string** | The name of the ID Verification provider (optional, auto-selected if not provided) | 

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


## IamApiControllerWebAuthnSignupBegin

> map[string]interface{} IamApiControllerWebAuthnSignupBegin(ctx).Execute()

Api Controller Web Authn Signup Begin



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
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerWebAuthnSignupBegin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerWebAuthnSignupBegin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerWebAuthnSignupBegin`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerWebAuthnSignupBegin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerWebAuthnSignupBeginRequest struct via the builder pattern


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


## IamApiControllerWebAuthnSignupFinish

> IamControllersResponse IamApiControllerWebAuthnSignupFinish(ctx).Body(body).Execute()

Api Controller Web Authn Signup Finish



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
	body := map[string]interface{}{ ... } // map[string]interface{} | authenticator attestation Response

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamUsersAPI.IamApiControllerWebAuthnSignupFinish(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamUsersAPI.IamApiControllerWebAuthnSignupFinish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerWebAuthnSignupFinish`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamUsersAPI.IamApiControllerWebAuthnSignupFinish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerWebAuthnSignupFinishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | authenticator attestation Response | 

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

