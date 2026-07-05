# \CommerceUsersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCreateUser**](CommerceUsersAPI.md#CommerceCreateUser) | **Post** /v1/commerce/user | Create user
[**CommerceCreateWalletAccount**](CommerceUsersAPI.md#CommerceCreateWalletAccount) | **Post** /v1/commerce/user/{userid}/wallet/account | Create wallet account
[**CommerceDeleteUser**](CommerceUsersAPI.md#CommerceDeleteUser) | **Delete** /v1/commerce/user/{userid} | Delete user
[**CommerceGetUser**](CommerceUsersAPI.md#CommerceGetUser) | **Get** /v1/commerce/user/{userid} | Get user
[**CommerceGetUserOrders**](CommerceUsersAPI.md#CommerceGetUserOrders) | **Get** /v1/commerce/user/{userid}/orders | Get user orders
[**CommerceGetUserPaymentMethods**](CommerceUsersAPI.md#CommerceGetUserPaymentMethods) | **Get** /v1/commerce/user/{userid}/paymentmethods | Get user payment methods
[**CommerceGetUserReferrals**](CommerceUsersAPI.md#CommerceGetUserReferrals) | **Get** /v1/commerce/user/{userid}/referrals | Get user referrals
[**CommerceGetUserReferrers**](CommerceUsersAPI.md#CommerceGetUserReferrers) | **Get** /v1/commerce/user/{userid}/referrers | Get user referrers
[**CommerceGetUserTransactions**](CommerceUsersAPI.md#CommerceGetUserTransactions) | **Get** /v1/commerce/user/{userid}/transactions | Get user transactions
[**CommerceGetUserWallet**](CommerceUsersAPI.md#CommerceGetUserWallet) | **Get** /v1/commerce/user/{userid}/wallet | Get user wallet
[**CommerceGetWalletAccount**](CommerceUsersAPI.md#CommerceGetWalletAccount) | **Get** /v1/commerce/user/{userid}/wallet/account/{name} | Get wallet account
[**CommerceListUsers**](CommerceUsersAPI.md#CommerceListUsers) | **Get** /v1/commerce/user | List users
[**CommercePatchUser**](CommerceUsersAPI.md#CommercePatchUser) | **Patch** /v1/commerce/user/{userid} | Partially update user
[**CommerceResetUserPassword**](CommerceUsersAPI.md#CommerceResetUserPassword) | **Get** /v1/commerce/user/{userid}/password/reset | Reset user password (admin)
[**CommerceUpdateUser**](CommerceUsersAPI.md#CommerceUpdateUser) | **Put** /v1/commerce/user/{userid} | Update user
[**CommerceWalletPay**](CommerceUsersAPI.md#CommerceWalletPay) | **Post** /v1/commerce/user/{userid}/wallet/pay | Send payment from wallet



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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceCreateUser(context.Background()).CommerceUser(commerceUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceCreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateUser`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceCreateUser`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceCreateWalletAccount(context.Background(), userid).CommerceCreateWalletAccountRequest(commerceCreateWalletAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceCreateWalletAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateWalletAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceCreateWalletAccount`: %v\n", resp)
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
	r, err := apiClient.CommerceUsersAPI.CommerceDeleteUser(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceDeleteUser``: %v\n", err)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceGetUser(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUser`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceGetUser`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceGetUserOrders(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceGetUserOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserOrders`: []CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceGetUserOrders`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceGetUserPaymentMethods(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceGetUserPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserPaymentMethods`: []CommercePaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceGetUserPaymentMethods`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceGetUserReferrals(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceGetUserReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserReferrals`: []CommerceReferral
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceGetUserReferrals`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceGetUserReferrers(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceGetUserReferrers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserReferrers`: []CommerceReferrer
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceGetUserReferrers`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceGetUserTransactions(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceGetUserTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserTransactions`: map[string]CommerceTransactionData
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceGetUserTransactions`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceGetUserWallet(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceGetUserWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetUserWallet`: CommerceWallet
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceGetUserWallet`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceGetWalletAccount(context.Background(), userid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceGetWalletAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetWalletAccount`: CommerceWalletAccount
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceGetWalletAccount`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceListUsers(context.Background()).Page(page).Display(display).Sort(sort).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListUsers`: CommercePaginatedUsers
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceListUsers`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommercePatchUser(context.Background(), userid).CommerceUser(commerceUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommercePatchUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchUser`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommercePatchUser`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceResetUserPassword(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceResetUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceResetUserPassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceResetUserPassword`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceUpdateUser(context.Background(), userid).CommerceUser(commerceUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateUser`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceUpdateUser`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceUsersAPI.CommerceWalletPay(context.Background(), userid).CommerceWalletPayRequest(commerceWalletPayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceUsersAPI.CommerceWalletPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceWalletPay`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceUsersAPI.CommerceWalletPay`: %v\n", resp)
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

