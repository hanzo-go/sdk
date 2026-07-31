# \UsersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiGetUserTableInfos**](UsersAPI.md#AiGetUserTableInfos) | **Get** /v1/ai/users/table-infos | Table Infos (user)
[**AiGetUsers**](UsersAPI.md#AiGetUsers) | **Get** /v1/ai/users | List users
[**AnalyticsCreateUser**](UsersAPI.md#AnalyticsCreateUser) | **Post** /v1/analytics/users | Create a new user (admin only)
[**AnalyticsDeleteUser**](UsersAPI.md#AnalyticsDeleteUser) | **Delete** /v1/analytics/users/{userId} | Delete user (admin only)
[**AnalyticsGetUser**](UsersAPI.md#AnalyticsGetUser) | **Get** /v1/analytics/users/{userId} | Get user by ID
[**AnalyticsGetUserTeams**](UsersAPI.md#AnalyticsGetUserTeams) | **Get** /v1/analytics/users/{userId}/teams | List teams a user belongs to
[**AnalyticsGetUserUsage**](UsersAPI.md#AnalyticsGetUserUsage) | **Get** /v1/analytics/users/{userId}/usage | Get event usage breakdown for a user (admin only)
[**AnalyticsGetUserWebsites**](UsersAPI.md#AnalyticsGetUserWebsites) | **Get** /v1/analytics/users/{userId}/websites | List websites owned by a user
[**AnalyticsUpdateUser**](UsersAPI.md#AnalyticsUpdateUser) | **Post** /v1/analytics/users/{userId} | Update user
[**BotGetUser**](UsersAPI.md#BotGetUser) | **Get** /v1/bot/users/{handle} | Get user profile by handle
[**BotListUserSkills**](UsersAPI.md#BotListUserSkills) | **Get** /v1/bot/users/{handle}/skills | List skills published by a user
[**BotListUserStars**](UsersAPI.md#BotListUserStars) | **Get** /v1/bot/users/{handle}/stars | List skills starred by a user
[**BotUpdateProfile**](UsersAPI.md#BotUpdateProfile) | **Patch** /v1/bot/users/me | Update current user&#39;s profile
[**CommerceCreateUser**](UsersAPI.md#CommerceCreateUser) | **Post** /v1/commerce/user | Create user
[**CommerceCreateWalletAccount**](UsersAPI.md#CommerceCreateWalletAccount) | **Post** /v1/commerce/user/{userid}/wallet/account | Create wallet account
[**CommerceDeleteUser**](UsersAPI.md#CommerceDeleteUser) | **Delete** /v1/commerce/user/{userid} | Delete user
[**CommerceGetUser**](UsersAPI.md#CommerceGetUser) | **Get** /v1/commerce/user/{userid} | Get user
[**CommerceGetUserOrders**](UsersAPI.md#CommerceGetUserOrders) | **Get** /v1/commerce/user/{userid}/orders | Get user orders
[**CommerceGetUserPaymentMethods**](UsersAPI.md#CommerceGetUserPaymentMethods) | **Get** /v1/commerce/user/{userid}/paymentmethods | Get user payment methods
[**CommerceGetUserReferrals**](UsersAPI.md#CommerceGetUserReferrals) | **Get** /v1/commerce/user/{userid}/referrals | Get user referrals
[**CommerceGetUserReferrers**](UsersAPI.md#CommerceGetUserReferrers) | **Get** /v1/commerce/user/{userid}/referrers | Get user referrers
[**CommerceGetUserTransactions**](UsersAPI.md#CommerceGetUserTransactions) | **Get** /v1/commerce/user/{userid}/transactions | Get user transactions
[**CommerceGetUserWallet**](UsersAPI.md#CommerceGetUserWallet) | **Get** /v1/commerce/user/{userid}/wallet | Get user wallet
[**CommerceGetWalletAccount**](UsersAPI.md#CommerceGetWalletAccount) | **Get** /v1/commerce/user/{userid}/wallet/account/{name} | Get wallet account
[**CommerceListUsers**](UsersAPI.md#CommerceListUsers) | **Get** /v1/commerce/user | List users
[**CommercePatchUser**](UsersAPI.md#CommercePatchUser) | **Patch** /v1/commerce/user/{userid} | Partially update user
[**CommerceResetUserPassword**](UsersAPI.md#CommerceResetUserPassword) | **Get** /v1/commerce/user/{userid}/password/reset | Reset user password (admin)
[**CommerceUpdateUser**](UsersAPI.md#CommerceUpdateUser) | **Put** /v1/commerce/user/{userid} | Update user
[**CommerceWalletPay**](UsersAPI.md#CommerceWalletPay) | **Post** /v1/commerce/user/{userid}/wallet/pay | Send payment from wallet
[**IamApiControllerAddLdap**](UsersAPI.md#IamApiControllerAddLdap) | **Post** /v1/iam/ldaps | Api Controller Add Ldap
[**IamApiControllerAddUser**](UsersAPI.md#IamApiControllerAddUser) | **Post** /v1/iam/users | Api Controller Add User
[**IamApiControllerAddUserKeys**](UsersAPI.md#IamApiControllerAddUserKeys) | **Post** /v1/iam/user-keys | Api Controller Add User Keys
[**IamApiControllerCheckUserPassword**](UsersAPI.md#IamApiControllerCheckUserPassword) | **Post** /v1/iam/auth/check-password | Api Controller Check User Password
[**IamApiControllerDeleteLdap**](UsersAPI.md#IamApiControllerDeleteLdap) | **Delete** /v1/iam/ldaps/{id} | Api Controller Delete Ldap
[**IamApiControllerDeleteUser**](UsersAPI.md#IamApiControllerDeleteUser) | **Delete** /v1/iam/users/{id} | Api Controller Delete User
[**IamApiControllerExitImpersonateUser**](UsersAPI.md#IamApiControllerExitImpersonateUser) | **Post** /v1/iam/impersonation/exit | Api Controller Exit Impersonate User
[**IamApiControllerGetAccount**](UsersAPI.md#IamApiControllerGetAccount) | **Get** /v1/iam/accounts/{id} | Api Controller Get Account
[**IamApiControllerGetEmailAndPhone**](UsersAPI.md#IamApiControllerGetEmailAndPhone) | **Get** /v1/iam/auth/contact | Api Controller Get Email And Phone
[**IamApiControllerGetGlobalUsers**](UsersAPI.md#IamApiControllerGetGlobalUsers) | **Get** /v1/iam/global-users | Api Controller Get Global Users
[**IamApiControllerGetLdap**](UsersAPI.md#IamApiControllerGetLdap) | **Get** /v1/iam/ldaps/{id} | Api Controller Get Ldap
[**IamApiControllerGetLdaps**](UsersAPI.md#IamApiControllerGetLdaps) | **Get** /v1/iam/ldaps | Api Controller Get Ldaps
[**IamApiControllerGetLdapser**](UsersAPI.md#IamApiControllerGetLdapser) | **Get** /v1/iam/ldap-users | Api Controller Get Ldapser
[**IamApiControllerGetSortedUsers**](UsersAPI.md#IamApiControllerGetSortedUsers) | **Get** /v1/iam/sorted-users | Api Controller Get Sorted Users
[**IamApiControllerGetUser**](UsersAPI.md#IamApiControllerGetUser) | **Get** /v1/iam/users/{id} | Api Controller Get User
[**IamApiControllerGetUserCount**](UsersAPI.md#IamApiControllerGetUserCount) | **Get** /v1/iam/user-counts/{id} | Api Controller Get User Count
[**IamApiControllerGetUsers**](UsersAPI.md#IamApiControllerGetUsers) | **Get** /v1/iam/users | Api Controller Get Users
[**IamApiControllerImpersonateUser**](UsersAPI.md#IamApiControllerImpersonateUser) | **Post** /v1/iam/impersonation-user | Api Controller Impersonate User
[**IamApiControllerResetEmailOrPhone**](UsersAPI.md#IamApiControllerResetEmailOrPhone) | **Post** /v1/iam/auth/reset-contact | Api Controller Reset Email Or Phone
[**IamApiControllerSetPassword**](UsersAPI.md#IamApiControllerSetPassword) | **Post** /v1/iam/auth/set-password | Api Controller Set Password
[**IamApiControllerSyncLdapUsers**](UsersAPI.md#IamApiControllerSyncLdapUsers) | **Post** /v1/iam/ldap/sync | Api Controller Sync Ldap Users
[**IamApiControllerUpdateLdap**](UsersAPI.md#IamApiControllerUpdateLdap) | **Put** /v1/iam/ldaps/{id} | Api Controller Update Ldap
[**IamApiControllerUpdateUser**](UsersAPI.md#IamApiControllerUpdateUser) | **Put** /v1/iam/users/{id} | Api Controller Update User
[**IamApiControllerUserInfo**](UsersAPI.md#IamApiControllerUserInfo) | **Get** /oauth/userinfo | Api Controller User Info
[**IamApiControllerUserInfo2**](UsersAPI.md#IamApiControllerUserInfo2) | **Get** /v1/iam/user | Api Controller User Info2
[**IamApiControllerVerifyIdentification**](UsersAPI.md#IamApiControllerVerifyIdentification) | **Post** /v1/iam/auth/identification/verify | Api Controller Verify Identification
[**IamApiControllerWebAuthnSignupBegin**](UsersAPI.md#IamApiControllerWebAuthnSignupBegin) | **Get** /v1/iam/auth/webauthn/signup/begin | Api Controller Web Authn Signup Begin
[**IamApiControllerWebAuthnSignupFinish**](UsersAPI.md#IamApiControllerWebAuthnSignupFinish) | **Post** /v1/iam/auth/webauthn/signup/finish | Api Controller Web Authn Signup Finish



## AiGetUserTableInfos

> AiEnvelope AiGetUserTableInfos(ctx).Execute()

Table Infos (user)

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
	resp, r, err := apiClient.UsersAPI.AiGetUserTableInfos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AiGetUserTableInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetUserTableInfos`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AiGetUserTableInfos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetUserTableInfosRequest struct via the builder pattern


### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiGetUsers

> AiEnvelope AiGetUsers(ctx).Execute()

List users



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
	resp, r, err := apiClient.UsersAPI.AiGetUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AiGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetUsers`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AiGetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetUsersRequest struct via the builder pattern


### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsCreateUser

> AnalyticsUser AnalyticsCreateUser(ctx).AnalyticsCreateUserRequest(analyticsCreateUserRequest).Execute()

Create a new user (admin only)

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
	analyticsCreateUserRequest := *openapiclient.NewAnalyticsCreateUserRequest("Username_example", "Password_example", "Role_example") // AnalyticsCreateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AnalyticsCreateUser(context.Background()).AnalyticsCreateUserRequest(analyticsCreateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AnalyticsCreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCreateUser`: AnalyticsUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AnalyticsCreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsCreateUserRequest** | [**AnalyticsCreateUserRequest**](AnalyticsCreateUserRequest.md) |  | 

### Return type

[**AnalyticsUser**](AnalyticsUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsDeleteUser

> map[string]interface{} AnalyticsDeleteUser(ctx, userId).Execute()

Delete user (admin only)

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AnalyticsDeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AnalyticsDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsDeleteUser`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AnalyticsDeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsDeleteUserRequest struct via the builder pattern


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


## AnalyticsGetUser

> AnalyticsUser AnalyticsGetUser(ctx, userId).Execute()

Get user by ID

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AnalyticsGetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AnalyticsGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetUser`: AnalyticsUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AnalyticsGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsUser**](AnalyticsUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetUserTeams

> []AnalyticsTeam AnalyticsGetUserTeams(ctx, userId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List teams a user belongs to

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AnalyticsGetUserTeams(context.Background(), userId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AnalyticsGetUserTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetUserTeams`: []AnalyticsTeam
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AnalyticsGetUserTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetUserTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsTeam**](AnalyticsTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetUserUsage

> AnalyticsGetUserUsage200Response AnalyticsGetUserUsage(ctx, userId).StartAt(startAt).EndAt(endAt).Execute()

Get event usage breakdown for a user (admin only)

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AnalyticsGetUserUsage(context.Background(), userId).StartAt(startAt).EndAt(endAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AnalyticsGetUserUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetUserUsage`: AnalyticsGetUserUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AnalyticsGetUserUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetUserUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 

### Return type

[**AnalyticsGetUserUsage200Response**](AnalyticsGetUserUsage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetUserWebsites

> []AnalyticsWebsite AnalyticsGetUserWebsites(ctx, userId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List websites owned by a user

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AnalyticsGetUserWebsites(context.Background(), userId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AnalyticsGetUserWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetUserWebsites`: []AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AnalyticsGetUserWebsites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetUserWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsUpdateUser

> AnalyticsUser AnalyticsUpdateUser(ctx, userId).AnalyticsUpdateUserRequest(analyticsUpdateUserRequest).Execute()

Update user

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	analyticsUpdateUserRequest := *openapiclient.NewAnalyticsUpdateUserRequest("Username_example") // AnalyticsUpdateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AnalyticsUpdateUser(context.Background(), userId).AnalyticsUpdateUserRequest(analyticsUpdateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AnalyticsUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsUpdateUser`: AnalyticsUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AnalyticsUpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsUpdateUserRequest** | [**AnalyticsUpdateUserRequest**](AnalyticsUpdateUserRequest.md) |  | 

### Return type

[**AnalyticsUser**](AnalyticsUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGetUser

> BotUser BotGetUser(ctx, handle).Execute()

Get user profile by handle

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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.BotGetUser(context.Background(), handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.BotGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetUser`: BotUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.BotGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotUser**](BotUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListUserSkills

> BotListUserSkills200Response BotListUserSkills(ctx, handle).Execute()

List skills published by a user

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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.BotListUserSkills(context.Background(), handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.BotListUserSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListUserSkills`: BotListUserSkills200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.BotListUserSkills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotListUserSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotListUserSkills200Response**](BotListUserSkills200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListUserStars

> BotListUserStars200Response BotListUserStars(ctx, handle).Execute()

List skills starred by a user

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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.BotListUserStars(context.Background(), handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.BotListUserStars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListUserStars`: BotListUserStars200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.BotListUserStars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotListUserStarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotListUserStars200Response**](BotListUserStars200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotUpdateProfile

> AnalyticsHeartbeat200Response BotUpdateProfile(ctx).BotUpdateProfileRequest(botUpdateProfileRequest).Execute()

Update current user's profile

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
	botUpdateProfileRequest := *openapiclient.NewBotUpdateProfileRequest() // BotUpdateProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.BotUpdateProfile(context.Background()).BotUpdateProfileRequest(botUpdateProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.BotUpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotUpdateProfile`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.BotUpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botUpdateProfileRequest** | [**BotUpdateProfileRequest**](BotUpdateProfileRequest.md) |  | 

### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreateUser

> CommerceUser CommerceCreateUser(ctx).CommerceUser(commerceUser).Execute()

Create user

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
	commerceUser := *openapiclient.NewCommerceUser() // CommerceUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceCreateUser(context.Background()).CommerceUser(commerceUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceCreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateUser`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceCreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceUser** | [**CommerceUser**](CommerceUser.md) |  | 

### Return type

[**CommerceUser**](CommerceUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreateWalletAccount

> map[string]interface{} CommerceCreateWalletAccount(ctx, userid).CommerceCreateWalletAccountRequest(commerceCreateWalletAccountRequest).Execute()

Create wallet account

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
	userid := "userid_example" // string | 
	commerceCreateWalletAccountRequest := *openapiclient.NewCommerceCreateWalletAccountRequest() // CommerceCreateWalletAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceCreateWalletAccount(context.Background(), userid).CommerceCreateWalletAccountRequest(commerceCreateWalletAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceCreateWalletAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateWalletAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceCreateWalletAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateWalletAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceCreateWalletAccountRequest** | [**CommerceCreateWalletAccountRequest**](CommerceCreateWalletAccountRequest.md) |  | 

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


## CommerceDeleteUser

> CommerceDeleteUser(ctx, userid).Execute()

Delete user

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.CommerceDeleteUser(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceDeleteUserRequest struct via the builder pattern


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


## CommerceGetUser

> CommerceUser CommerceGetUser(ctx, userid).Execute()

Get user

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceGetUser(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUser`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceUser**](CommerceUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetUserOrders

> []CommerceOrder CommerceGetUserOrders(ctx, userid).Execute()

Get user orders

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceGetUserOrders(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceGetUserOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserOrders`: []CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceGetUserOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetUserOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetUserPaymentMethods

> []CommercePaymentMethod CommerceGetUserPaymentMethods(ctx, userid).Execute()

Get user payment methods

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceGetUserPaymentMethods(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceGetUserPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserPaymentMethods`: []CommercePaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceGetUserPaymentMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetUserPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommercePaymentMethod**](CommercePaymentMethod.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetUserReferrals

> []CommerceReferral CommerceGetUserReferrals(ctx, userid).Execute()

Get user referrals

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceGetUserReferrals(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceGetUserReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserReferrals`: []CommerceReferral
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceGetUserReferrals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetUserReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommerceReferral**](CommerceReferral.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetUserReferrers

> []CommerceReferrer CommerceGetUserReferrers(ctx, userid).Execute()

Get user referrers

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceGetUserReferrers(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceGetUserReferrers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserReferrers`: []CommerceReferrer
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceGetUserReferrers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetUserReferrersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommerceReferrer**](CommerceReferrer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetUserTransactions

> map[string]CommerceTransactionData CommerceGetUserTransactions(ctx, userid).Execute()

Get user transactions

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceGetUserTransactions(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceGetUserTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserTransactions`: map[string]CommerceTransactionData
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceGetUserTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetUserTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]CommerceTransactionData**](CommerceTransactionData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetUserWallet

> CommerceWallet CommerceGetUserWallet(ctx, userid).Execute()

Get user wallet

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceGetUserWallet(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceGetUserWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserWallet`: CommerceWallet
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceGetUserWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetUserWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceWallet**](CommerceWallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetWalletAccount

> CommerceWalletAccount CommerceGetWalletAccount(ctx, userid, name).Execute()

Get wallet account

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
	userid := "userid_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceGetWalletAccount(context.Background(), userid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceGetWalletAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetWalletAccount`: CommerceWalletAccount
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceGetWalletAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetWalletAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommerceWalletAccount**](CommerceWalletAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListUsers

> CommercePaginatedUsers CommerceListUsers(ctx).Page(page).Display(display).Sort(sort).Q(q).Execute()

List users

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
	page := int32(56) // int32 | Page number (1-indexed) (optional) (default to 1)
	display := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sort := "sort_example" // string | Sort field (prefix with - for descending) (optional) (default to "-UpdatedAt")
	q := "q_example" // string | Search query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceListUsers(context.Background()).Page(page).Display(display).Sort(sort).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListUsers`: CommercePaginatedUsers
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed) | [default to 1]
 **display** | **int32** | Number of items per page | [default to 20]
 **sort** | **string** | Sort field (prefix with - for descending) | [default to &quot;-UpdatedAt&quot;]
 **q** | **string** | Search query | 

### Return type

[**CommercePaginatedUsers**](CommercePaginatedUsers.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommercePatchUser

> CommerceUser CommercePatchUser(ctx, userid).CommerceUser(commerceUser).Execute()

Partially update user

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
	userid := "userid_example" // string | 
	commerceUser := *openapiclient.NewCommerceUser() // CommerceUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommercePatchUser(context.Background(), userid).CommerceUser(commerceUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommercePatchUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchUser`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommercePatchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommercePatchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceUser** | [**CommerceUser**](CommerceUser.md) |  | 

### Return type

[**CommerceUser**](CommerceUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceResetUserPassword

> map[string]interface{} CommerceResetUserPassword(ctx, userid).Execute()

Reset user password (admin)

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
	userid := "userid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceResetUserPassword(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceResetUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceResetUserPassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceResetUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceResetUserPasswordRequest struct via the builder pattern


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


## CommerceUpdateUser

> CommerceUser CommerceUpdateUser(ctx, userid).CommerceUser(commerceUser).Execute()

Update user

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
	userid := "userid_example" // string | 
	commerceUser := *openapiclient.NewCommerceUser() // CommerceUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceUpdateUser(context.Background(), userid).CommerceUser(commerceUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateUser`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceUpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceUser** | [**CommerceUser**](CommerceUser.md) |  | 

### Return type

[**CommerceUser**](CommerceUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceWalletPay

> map[string]interface{} CommerceWalletPay(ctx, userid).CommerceWalletPayRequest(commerceWalletPayRequest).Execute()

Send payment from wallet

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
	userid := "userid_example" // string | 
	commerceWalletPayRequest := *openapiclient.NewCommerceWalletPayRequest() // CommerceWalletPayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CommerceWalletPay(context.Background(), userid).CommerceWalletPayRequest(commerceWalletPayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CommerceWalletPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceWalletPay`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CommerceWalletPay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceWalletPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceWalletPayRequest** | [**CommerceWalletPayRequest**](CommerceWalletPayRequest.md) |  | 

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
	resp, r, err := apiClient.UsersAPI.IamApiControllerAddLdap(context.Background()).IamObjectLdap(iamObjectLdap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerAddLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddLdap`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerAddLdap`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerAddUser(context.Background()).IamObjectUser(iamObjectUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerAddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerAddUser`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerAddUserKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerAddUserKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddUserKeys`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerAddUserKeys`: %v\n", resp)
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

> IamObjectUserinfo IamApiControllerCheckUserPassword(ctx).IamControllersCheckPasswordRequest(iamControllersCheckPasswordRequest).Execute()

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
	iamControllersCheckPasswordRequest := *openapiclient.NewIamControllersCheckPasswordRequest("Owner_example", "Name_example", "Password_example") // IamControllersCheckPasswordRequest | Handler binds the full object.User; only owner, name, password (and ldap, which toggles LDAP-password mode) are consumed. Extra User fields are ignored.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.IamApiControllerCheckUserPassword(context.Background()).IamControllersCheckPasswordRequest(iamControllersCheckPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerCheckUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerCheckUserPassword`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerCheckUserPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerCheckUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamControllersCheckPasswordRequest** | [**IamControllersCheckPasswordRequest**](IamControllersCheckPasswordRequest.md) | Handler binds the full object.User; only owner, name, password (and ldap, which toggles LDAP-password mode) are consumed. Extra User fields are ignored. | 

### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerDeleteLdap(context.Background(), id).IamObjectLdap(iamObjectLdap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerDeleteLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteLdap`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerDeleteLdap`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerDeleteUser(context.Background(), id).IamObjectUser(iamObjectUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerDeleteUser`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerExitImpersonateUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerExitImpersonateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerExitImpersonateUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerExitImpersonateUser`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAccount`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetAccount`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetEmailAndPhone(context.Background()).Username(username).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetEmailAndPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetEmailAndPhone`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetEmailAndPhone`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetGlobalUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetGlobalUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGlobalUsers`: []IamObjectUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetGlobalUsers`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetLdap(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetLdap`: IamObjectLdap
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetLdap`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetLdaps(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetLdaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetLdaps`: []IamObjectLdap
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetLdaps`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetLdapser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetLdapser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetLdapser`: IamControllersLdapResp
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetLdapser`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetSortedUsers(context.Background()).Owner(owner).Sorter(sorter).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetSortedUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSortedUsers`: []IamObjectUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetSortedUsers`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetUser(context.Background(), id).Owner(owner).Email(email).Phone(phone).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUser`: IamObjectUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetUser`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetUserCount(context.Background(), id).Owner(owner).IsOnline(isOnline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetUserCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUserCount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetUserCount`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerGetUsers(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUsers`: []IamObjectUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerGetUsers`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerImpersonateUser(context.Background()).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerImpersonateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerImpersonateUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerImpersonateUser`: %v\n", resp)
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

> IamObjectUserinfo IamApiControllerResetEmailOrPhone(ctx).Type_(type_).Dest(dest).Code(code).Execute()

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
	type_ := "type__example" // string | \\\"email\\\" or \\\"phone\\\"
	dest := "dest_example" // string | 
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.IamApiControllerResetEmailOrPhone(context.Background()).Type_(type_).Dest(dest).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerResetEmailOrPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerResetEmailOrPhone`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerResetEmailOrPhone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerResetEmailOrPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | \\\&quot;email\\\&quot; or \\\&quot;phone\\\&quot; | 
 **dest** | **string** |  | 
 **code** | **string** |  | 

### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerSetPassword(context.Background()).UserOwner(userOwner).UserName(userName).OldPassword(oldPassword).NewPassword(newPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerSetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSetPassword`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerSetPassword`: %v\n", resp)
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

> IamControllersLdapSyncResp IamApiControllerSyncLdapUsers(ctx).Id(id).IamObjectLdapUser(iamObjectLdapUser).Execute()

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
	iamObjectLdapUser := []openapiclient.IamObjectLdapUser{*openapiclient.NewIamObjectLdapUser()} // []IamObjectLdapUser | The LDAP users to sync (JSON array; may be empty). Query id (<owner>/<ldapId>) selects the server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.IamApiControllerSyncLdapUsers(context.Background()).Id(id).IamObjectLdapUser(iamObjectLdapUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerSyncLdapUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSyncLdapUsers`: IamControllersLdapSyncResp
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerSyncLdapUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSyncLdapUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | id | 
 **iamObjectLdapUser** | [**[]IamObjectLdapUser**](IamObjectLdapUser.md) | The LDAP users to sync (JSON array; may be empty). Query id (&lt;owner&gt;/&lt;ldapId&gt;) selects the server. | 

### Return type

[**IamControllersLdapSyncResp**](IamControllersLdapSyncResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerUpdateLdap(context.Background(), id).IamObjectLdap(iamObjectLdap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerUpdateLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateLdap`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerUpdateLdap`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerUpdateUser(context.Background(), id).IamObjectUser(iamObjectUser).UserId(userId).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateUser`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerUpdateUser`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUserInfo`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerUserInfo`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerUserInfo2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerUserInfo2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUserInfo2`: IamControllersLaravelResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerUserInfo2`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerVerifyIdentification(context.Background()).Owner(owner).Name(name).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerVerifyIdentification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerVerifyIdentification`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerVerifyIdentification`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerWebAuthnSignupBegin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerWebAuthnSignupBegin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerWebAuthnSignupBegin`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerWebAuthnSignupBegin`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.IamApiControllerWebAuthnSignupFinish(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.IamApiControllerWebAuthnSignupFinish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerWebAuthnSignupFinish`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.IamApiControllerWebAuthnSignupFinish`: %v\n", resp)
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

