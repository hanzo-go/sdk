# \FinanceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinanceAccounts**](FinanceAPI.md#GetFinanceAccounts) | **Get** /v1/finance/accounts | Returns the ledger accounts the caller may see, with their balances.
[**GetFinanceBalance**](FinanceAPI.md#GetFinanceBalance) | **Get** /v1/finance/balance | Answers the org&#39;s spendable prepaid balance typed for the finance surfaces: &#x60;availableCents&#x60;, &#x60;pendingCents&#x60;, &#x60;dueCents&#x60; and the &#x60;asOf&#x60; instant it was read.
[**GetFinanceCredits**](FinanceAPI.md#GetFinanceCredits) | **Get** /v1/finance/credits | Answers the money PUT IN to the org&#39;s wallet — each staff grant, promo and settled top-up as a positive row with its id, label, cents and grant time.
[**GetFinanceInvoices**](FinanceAPI.md#GetFinanceInvoices) | **Get** /v1/finance/invoices | Answers an empty typed array, always.
[**GetFinanceLedger**](FinanceAPI.md#GetFinanceLedger) | **Get** /v1/finance/ledger | Answers the org&#39;s own postings inside &#x60;range&#x3D;&#x60;, each as a signed entry: a DEPOSIT CREDITS the wallet (positive, account &#x60;credits:&lt;org&gt;&#x60;) and every other posting DEBITS it (negative, account &#x60;usage:&lt;org&gt;&#x60;), described by its notes or its tags.
[**GetFinancePaymentMethods**](FinanceAPI.md#GetFinancePaymentMethods) | **Get** /v1/finance/payment-methods | Answers the masked card descriptors for the caller&#39;s resolved WALLET — id, brand, last four, expiry, default flag — reshaped into the finance contract.
[**GetFinanceTreasury**](FinanceAPI.md#GetFinanceTreasury) | **Get** /v1/finance/treasury | Returns the reserve fund&#39;s health and the current revenue-share policy for any validated caller.
[**GetFinanceUsage**](FinanceAPI.md#GetFinanceUsage) | **Get** /v1/finance/usage | Answers metered spend inside &#x60;range&#x3D;&#x60;: the window total, a time series to plot, and one line per usage TAG.



## GetFinanceAccounts

> AccountsOut GetFinanceAccounts(ctx).Scope(scope).Org(org).Execute()

Returns the ledger accounts the caller may see, with their balances.



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
	resp, r, err := apiClient.FinanceAPI.GetFinanceAccounts(context.Background()).Scope(scope).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetFinanceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccounts`: AccountsOut
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetFinanceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope is \&quot;house\&quot; to read the reserve/revenue/payout house accounts. SuperAdmin only. | 
 **org** | **string** | Org names another tenant to read. SuperAdmin only; ignored when scope&#x3D;house. | 

### Return type

[**AccountsOut**](AccountsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceBalance

> FinanceBalanceView GetFinanceBalance(ctx).Execute()

Answers the org's spendable prepaid balance typed for the finance surfaces: `availableCents`, `pendingCents`, `dueCents` and the `asOf` instant it was read.



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
	resp, r, err := apiClient.FinanceAPI.GetFinanceBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetFinanceBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBalance`: FinanceBalanceView
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetFinanceBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBalanceRequest struct via the builder pattern


### Return type

[**FinanceBalanceView**](FinanceBalanceView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceCredits

> []FinanceCredit GetFinanceCredits(ctx).Execute()

Answers the money PUT IN to the org's wallet — each staff grant, promo and settled top-up as a positive row with its id, label, cents and grant time.



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
	resp, r, err := apiClient.FinanceAPI.GetFinanceCredits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetFinanceCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceCredits`: []FinanceCredit
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetFinanceCredits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceCreditsRequest struct via the builder pattern


### Return type

[**[]FinanceCredit**](FinanceCredit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceInvoices

> []FinanceInvoice GetFinanceInvoices(ctx).Execute()

Answers an empty typed array, always.



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
	resp, r, err := apiClient.FinanceAPI.GetFinanceInvoices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetFinanceInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoices`: []FinanceInvoice
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetFinanceInvoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesRequest struct via the builder pattern


### Return type

[**[]FinanceInvoice**](FinanceInvoice.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceLedger

> []FinanceLedgerEntry GetFinanceLedger(ctx).Range_(range_).Execute()

Answers the org's own postings inside `range=`, each as a signed entry: a DEPOSIT CREDITS the wallet (positive, account `credits:<org>`) and every other posting DEBITS it (negative, account `usage:<org>`), described by its notes or its tags.



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
	range_ := "30d" // string | Range is the window: 24h, 7d, 30d or 90d. Anything else — including absent — is 30d, so a typo silently widens the window to a month rather than failing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinanceAPI.GetFinanceLedger(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetFinanceLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceLedger`: []FinanceLedgerEntry
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetFinanceLedger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the window: 24h, 7d, 30d or 90d. Anything else — including absent — is 30d, so a typo silently widens the window to a month rather than failing. | 

### Return type

[**[]FinanceLedgerEntry**](FinanceLedgerEntry.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancePaymentMethods

> []FinancePaymentMethod GetFinancePaymentMethods(ctx).Execute()

Answers the masked card descriptors for the caller's resolved WALLET — id, brand, last four, expiry, default flag — reshaped into the finance contract.



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
	resp, r, err := apiClient.FinanceAPI.GetFinancePaymentMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetFinancePaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinancePaymentMethods`: []FinancePaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetFinancePaymentMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancePaymentMethodsRequest struct via the builder pattern


### Return type

[**[]FinancePaymentMethod**](FinancePaymentMethod.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTreasury

> TreasuryReport GetFinanceTreasury(ctx).Execute()

Returns the reserve fund's health and the current revenue-share policy for any validated caller.



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
	resp, r, err := apiClient.FinanceAPI.GetFinanceTreasury(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetFinanceTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTreasury`: TreasuryReport
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetFinanceTreasury`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTreasuryRequest struct via the builder pattern


### Return type

[**TreasuryReport**](TreasuryReport.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceUsage

> FinanceUsageView GetFinanceUsage(ctx).Range_(range_).Execute()

Answers metered spend inside `range=`: the window total, a time series to plot, and one line per usage TAG.



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
	range_ := "24h" // string | Range is the window: 24h, 7d, 30d or 90d. Anything else — including absent — is 30d, so a typo silently widens the window to a month rather than failing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinanceAPI.GetFinanceUsage(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.GetFinanceUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceUsage`: FinanceUsageView
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.GetFinanceUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the window: 24h, 7d, 30d or 90d. Anything else — including absent — is 30d, so a typo silently widens the window to a month rather than failing. | 

### Return type

[**FinanceUsageView**](FinanceUsageView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

