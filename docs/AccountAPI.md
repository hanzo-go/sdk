# \AccountAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiGetAccount**](AccountAPI.md#AiGetAccount) | **Get** /v1/ai/account | Account
[**CommerceAccountExists**](AccountAPI.md#CommerceAccountExists) | **Get** /v1/commerce/account/exists/{emailorusername} | Check if account exists
[**CommerceAccountLogin**](AccountAPI.md#CommerceAccountLogin) | **Post** /v1/commerce/account/login | Login to account
[**CommerceAccountWithdraw**](AccountAPI.md#CommerceAccountWithdraw) | **Post** /v1/commerce/account/withdraw | Withdraw funds
[**CommerceConfirmPasswordReset**](AccountAPI.md#CommerceConfirmPasswordReset) | **Post** /v1/commerce/account/confirm/{tokenid} | Confirm password reset
[**CommerceCreateAccount**](AccountAPI.md#CommerceCreateAccount) | **Post** /v1/commerce/account/create | Create new account
[**CommerceCreatePaymentMethod**](AccountAPI.md#CommerceCreatePaymentMethod) | **Post** /v1/commerce/account/paymentmethod/{paymentmethodtype} | Create payment method
[**CommerceEnableAccount**](AccountAPI.md#CommerceEnableAccount) | **Post** /v1/commerce/account/enable/{tokenid} | Enable account with token
[**CommerceGetAccount**](AccountAPI.md#CommerceGetAccount) | **Get** /v1/commerce/account | Get current account
[**CommerceGetAccountOrder**](AccountAPI.md#CommerceGetAccountOrder) | **Get** /v1/commerce/account/order/{orderid} | Get account order
[**CommercePatchAccount**](AccountAPI.md#CommercePatchAccount) | **Patch** /v1/commerce/account | Partially update account
[**CommercePatchAccountOrder**](AccountAPI.md#CommercePatchAccountOrder) | **Patch** /v1/commerce/account/order/{orderid} | Update account order
[**CommerceRequestPasswordReset**](AccountAPI.md#CommerceRequestPasswordReset) | **Post** /v1/commerce/account/reset | Request password reset
[**CommerceUpdateAccount**](AccountAPI.md#CommerceUpdateAccount) | **Put** /v1/commerce/account | Update account



## AiGetAccount

> AiEnvelope AiGetAccount(ctx).Execute()

Account

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
	resp, r, err := apiClient.AccountAPI.AiGetAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AiGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetAccount`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AiGetAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetAccountRequest struct via the builder pattern


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


## CommerceAccountExists

> CommerceAccountExists200Response CommerceAccountExists(ctx, emailorusername).Execute()

Check if account exists

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
	emailorusername := "emailorusername_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceAccountExists(context.Background(), emailorusername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceAccountExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceAccountExists`: CommerceAccountExists200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceAccountExists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailorusername** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceAccountExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceAccountExists200Response**](CommerceAccountExists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceAccountLogin

> CommerceUser CommerceAccountLogin(ctx).CommerceAccountLoginRequest(commerceAccountLoginRequest).Execute()

Login to account

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
	commerceAccountLoginRequest := *openapiclient.NewCommerceAccountLoginRequest("Email_example", "Password_example") // CommerceAccountLoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceAccountLogin(context.Background()).CommerceAccountLoginRequest(commerceAccountLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceAccountLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceAccountLogin`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceAccountLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceAccountLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceAccountLoginRequest** | [**CommerceAccountLoginRequest**](CommerceAccountLoginRequest.md) |  | 

### Return type

[**CommerceUser**](CommerceUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceAccountWithdraw

> map[string]interface{} CommerceAccountWithdraw(ctx).CommerceAccountWithdrawRequest(commerceAccountWithdrawRequest).Execute()

Withdraw funds

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
	commerceAccountWithdrawRequest := *openapiclient.NewCommerceAccountWithdrawRequest() // CommerceAccountWithdrawRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceAccountWithdraw(context.Background()).CommerceAccountWithdrawRequest(commerceAccountWithdrawRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceAccountWithdraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceAccountWithdraw`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceAccountWithdraw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceAccountWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceAccountWithdrawRequest** | [**CommerceAccountWithdrawRequest**](CommerceAccountWithdrawRequest.md) |  | 

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


## CommerceConfirmPasswordReset

> map[string]interface{} CommerceConfirmPasswordReset(ctx, tokenid).CommerceConfirmPasswordResetRequest(commerceConfirmPasswordResetRequest).Execute()

Confirm password reset

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
	tokenid := "tokenid_example" // string | 
	commerceConfirmPasswordResetRequest := *openapiclient.NewCommerceConfirmPasswordResetRequest("Password_example") // CommerceConfirmPasswordResetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceConfirmPasswordReset(context.Background(), tokenid).CommerceConfirmPasswordResetRequest(commerceConfirmPasswordResetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceConfirmPasswordReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceConfirmPasswordReset`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceConfirmPasswordReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceConfirmPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceConfirmPasswordResetRequest** | [**CommerceConfirmPasswordResetRequest**](CommerceConfirmPasswordResetRequest.md) |  | 

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


## CommerceCreateAccount

> CommerceUser CommerceCreateAccount(ctx).CommerceCreateAccountRequest(commerceCreateAccountRequest).Execute()

Create new account

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
	commerceCreateAccountRequest := *openapiclient.NewCommerceCreateAccountRequest("Email_example", "Password_example", "FirstName_example", "LastName_example") // CommerceCreateAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceCreateAccount(context.Background()).CommerceCreateAccountRequest(commerceCreateAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceCreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateAccount`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceCreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceCreateAccountRequest** | [**CommerceCreateAccountRequest**](CommerceCreateAccountRequest.md) |  | 

### Return type

[**CommerceUser**](CommerceUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreatePaymentMethod

> CommercePaymentMethod CommerceCreatePaymentMethod(ctx, paymentmethodtype).CommercePaymentMethod(commercePaymentMethod).Execute()

Create payment method

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
	paymentmethodtype := "paymentmethodtype_example" // string | 
	commercePaymentMethod := *openapiclient.NewCommercePaymentMethod() // CommercePaymentMethod | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceCreatePaymentMethod(context.Background(), paymentmethodtype).CommercePaymentMethod(commercePaymentMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceCreatePaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreatePaymentMethod`: CommercePaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceCreatePaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentmethodtype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreatePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commercePaymentMethod** | [**CommercePaymentMethod**](CommercePaymentMethod.md) |  | 

### Return type

[**CommercePaymentMethod**](CommercePaymentMethod.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceEnableAccount

> map[string]interface{} CommerceEnableAccount(ctx, tokenid).CommerceCommerceEnableAccountRequest(commerceCommerceEnableAccountRequest).Execute()

Enable account with token

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
	tokenid := "tokenid_example" // string | 
	commerceCommerceEnableAccountRequest := *openapiclient.NewCommerceCommerceEnableAccountRequest() // CommerceCommerceEnableAccountRequest | Optional. Read (json.DecodeBytes) only when the org has two-stage enablement (SignUpOptions.TwoStageEnabled); otherwise ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceEnableAccount(context.Background(), tokenid).CommerceCommerceEnableAccountRequest(commerceCommerceEnableAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceEnableAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceEnableAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceEnableAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceEnableAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceCommerceEnableAccountRequest** | [**CommerceCommerceEnableAccountRequest**](CommerceCommerceEnableAccountRequest.md) | Optional. Read (json.DecodeBytes) only when the org has two-stage enablement (SignUpOptions.TwoStageEnabled); otherwise ignored. | 

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


## CommerceGetAccount

> CommerceUser CommerceGetAccount(ctx).Execute()

Get current account

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
	resp, r, err := apiClient.AccountAPI.CommerceGetAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetAccount`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceGetAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetAccountRequest struct via the builder pattern


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


## CommerceGetAccountOrder

> CommerceOrder CommerceGetAccountOrder(ctx, orderid).Execute()

Get account order

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
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceGetAccountOrder(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceGetAccountOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetAccountOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceGetAccountOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetAccountOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommercePatchAccount

> CommerceUser CommercePatchAccount(ctx).CommerceUser(commerceUser).Execute()

Partially update account

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
	resp, r, err := apiClient.AccountAPI.CommercePatchAccount(context.Background()).CommerceUser(commerceUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommercePatchAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchAccount`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommercePatchAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommercePatchAccountRequest struct via the builder pattern


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


## CommercePatchAccountOrder

> CommerceOrder CommercePatchAccountOrder(ctx, orderid).CommerceOrder(commerceOrder).Execute()

Update account order

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
	orderid := "orderid_example" // string | 
	commerceOrder := *openapiclient.NewCommerceOrder() // CommerceOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommercePatchAccountOrder(context.Background(), orderid).CommerceOrder(commerceOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommercePatchAccountOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchAccountOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommercePatchAccountOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommercePatchAccountOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceOrder** | [**CommerceOrder**](CommerceOrder.md) |  | 

### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceRequestPasswordReset

> map[string]interface{} CommerceRequestPasswordReset(ctx).CommerceRequestPasswordResetRequest(commerceRequestPasswordResetRequest).Execute()

Request password reset

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
	commerceRequestPasswordResetRequest := *openapiclient.NewCommerceRequestPasswordResetRequest("Email_example") // CommerceRequestPasswordResetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CommerceRequestPasswordReset(context.Background()).CommerceRequestPasswordResetRequest(commerceRequestPasswordResetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceRequestPasswordReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceRequestPasswordReset`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceRequestPasswordReset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceRequestPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceRequestPasswordResetRequest** | [**CommerceRequestPasswordResetRequest**](CommerceRequestPasswordResetRequest.md) |  | 

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


## CommerceUpdateAccount

> CommerceUser CommerceUpdateAccount(ctx).CommerceUser(commerceUser).Execute()

Update account

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
	resp, r, err := apiClient.AccountAPI.CommerceUpdateAccount(context.Background()).CommerceUser(commerceUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CommerceUpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateAccount`: CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CommerceUpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceUpdateAccountRequest struct via the builder pattern


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

