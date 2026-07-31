# \FinanceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1FinanceAccounts**](FinanceAPI.md#CloudGetV1FinanceAccounts) | **Get** /v1/finance/accounts | ListFinanceAccounts returns the ledger accounts the caller may see, with their balances.
[**CloudGetV1FinanceBalance**](FinanceAPI.md#CloudGetV1FinanceBalance) | **Get** /v1/finance/balance | 
[**CloudGetV1FinanceCredits**](FinanceAPI.md#CloudGetV1FinanceCredits) | **Get** /v1/finance/credits | 
[**CloudGetV1FinanceInvoices**](FinanceAPI.md#CloudGetV1FinanceInvoices) | **Get** /v1/finance/invoices | 
[**CloudGetV1FinanceLedger**](FinanceAPI.md#CloudGetV1FinanceLedger) | **Get** /v1/finance/ledger | 
[**CloudGetV1FinancePaymentMethods**](FinanceAPI.md#CloudGetV1FinancePaymentMethods) | **Get** /v1/finance/payment-methods | 
[**CloudGetV1FinanceTreasury**](FinanceAPI.md#CloudGetV1FinanceTreasury) | **Get** /v1/finance/treasury | GetTreasury returns the reserve fund&#39;s health and the current revenue-share policy for any validated caller.
[**CloudGetV1FinanceUsage**](FinanceAPI.md#CloudGetV1FinanceUsage) | **Get** /v1/finance/usage | 



## CloudGetV1FinanceAccounts

> CloudAccountsOut CloudGetV1FinanceAccounts(ctx).Scope(scope).Org(org).Execute()

ListFinanceAccounts returns the ledger accounts the caller may see, with their balances.



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
	scope := "scope_example" // string | Scope is \"house\" to read the reserve/revenue/payout house accounts. SuperAdmin only. (optional)
	org := "org_example" // string | Org names another tenant to read. SuperAdmin only; ignored when scope=house. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinanceAPI.CloudGetV1FinanceAccounts(context.Background()).Scope(scope).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.CloudGetV1FinanceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FinanceAccounts`: CloudAccountsOut
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.CloudGetV1FinanceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FinanceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope is \&quot;house\&quot; to read the reserve/revenue/payout house accounts. SuperAdmin only. | 
 **org** | **string** | Org names another tenant to read. SuperAdmin only; ignored when scope&#x3D;house. | 

### Return type

[**CloudAccountsOut**](CloudAccountsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FinanceBalance

> CloudGetV1FinanceBalance(ctx).Execute()



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
	r, err := apiClient.FinanceAPI.CloudGetV1FinanceBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.CloudGetV1FinanceBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FinanceBalanceRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FinanceCredits

> CloudGetV1FinanceCredits(ctx).Execute()



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
	r, err := apiClient.FinanceAPI.CloudGetV1FinanceCredits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.CloudGetV1FinanceCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FinanceCreditsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FinanceInvoices

> CloudGetV1FinanceInvoices(ctx).Execute()



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
	r, err := apiClient.FinanceAPI.CloudGetV1FinanceInvoices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.CloudGetV1FinanceInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FinanceInvoicesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FinanceLedger

> CloudGetV1FinanceLedger(ctx).Execute()



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
	r, err := apiClient.FinanceAPI.CloudGetV1FinanceLedger(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.CloudGetV1FinanceLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FinanceLedgerRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FinancePaymentMethods

> CloudGetV1FinancePaymentMethods(ctx).Execute()



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
	r, err := apiClient.FinanceAPI.CloudGetV1FinancePaymentMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.CloudGetV1FinancePaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FinancePaymentMethodsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FinanceTreasury

> CloudTreasuryReport CloudGetV1FinanceTreasury(ctx).Execute()

GetTreasury returns the reserve fund's health and the current revenue-share policy for any validated caller.



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
	resp, r, err := apiClient.FinanceAPI.CloudGetV1FinanceTreasury(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.CloudGetV1FinanceTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FinanceTreasury`: CloudTreasuryReport
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.CloudGetV1FinanceTreasury`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FinanceTreasuryRequest struct via the builder pattern


### Return type

[**CloudTreasuryReport**](CloudTreasuryReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FinanceUsage

> CloudGetV1FinanceUsage(ctx).Execute()



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
	r, err := apiClient.FinanceAPI.CloudGetV1FinanceUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.CloudGetV1FinanceUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FinanceUsageRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

