# \BillingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscription**](BillingAPI.md#CancelSubscription) | **Post** /v1/billing/subscriptions/{id}/cancel | End a subscription
[**CollectInvoice**](BillingAPI.md#CollectInvoice) | **Post** /v1/billing/invoices/{id}/collect | Collect an issued invoice from credits, balance, then card
[**DeleteBillingAlertsById**](BillingAPI.md#DeleteBillingAlertsById) | **Delete** /v1/billing/alerts/{id} | Removes one of the caller&#39;s spend caps and answers 204.
[**DeleteBillingMethodsById**](BillingAPI.md#DeleteBillingMethodsById) | **Delete** /v1/billing/methods/{id} | Removes one card or account the caller has saved.
[**DeleteBillingPortalMethodsById**](BillingAPI.md#DeleteBillingPortalMethodsById) | **Delete** /v1/billing/portal/methods/{id} | DetachPortalMethod is DetachMethod at the address a hosted checkout addresses it by.
[**GetBillingAccounts**](BillingAPI.md#GetBillingAccounts) | **Get** /v1/billing/accounts | Answers the caller&#39;s billing accounts: the org itself, its currency, when it was opened, and the caller&#39;s own standing in it.
[**GetBillingAccountsByIdMembers**](BillingAPI.md#GetBillingAccountsByIdMembers) | **Get** /v1/billing/accounts/{id}/members | Answers one billing account&#39;s roster.
[**GetBillingAlerts**](BillingAPI.md#GetBillingAlerts) | **Get** /v1/billing/alerts | Lists this org&#39;s spend caps: the ceiling, its scope, whether it enforces, and how much of it has been spent this period.
[**GetBillingAlertsAuthorize**](BillingAPI.md#GetBillingAlertsAuthorize) | **Get** /v1/billing/alerts/authorize | Answers whether one proposed spend fits inside this org&#39;s caps.
[**GetBillingBalance**](BillingAPI.md#GetBillingBalance) | **Get** /v1/billing/balance | Prepaid credit the caller&#39;s org can still spend
[**GetBillingCreditBalance**](BillingAPI.md#GetBillingCreditBalance) | **Get** /v1/billing/credit-balance | Answers what the caller can spend right now, one entry per currency.
[**GetBillingCreditBalanceBreakdown**](BillingAPI.md#GetBillingCreditBalanceBreakdown) | **Get** /v1/billing/credit-balance/breakdown | Answers that same spendable credit split by grant tag, with the earliest expiry under each and the total across all of them.
[**GetBillingCredits**](BillingAPI.md#GetBillingCredits) | **Get** /v1/billing/credits | Lists the caller&#39;s credit grants — every one of them, spent and lapsed and voided included.
[**GetBillingCryptoDepositById**](BillingAPI.md#GetBillingCryptoDepositById) | **Get** /v1/billing/crypto/deposit/{id} | Reads one of the caller&#39;s own deposit intents back — pending, confirming, or succeeded.
[**GetBillingCryptoOptions**](BillingAPI.md#GetBillingCryptoOptions) | **Get** /v1/billing/crypto/options | Answers which chains and tokens the crypto rail accepts — what an asset picker renders.
[**GetBillingInvoices**](BillingAPI.md#GetBillingInvoices) | **Get** /v1/billing/invoices | Lists the caller&#39;s invoices, newest first, with the count beside them.
[**GetBillingInvoicesByIdPdf**](BillingAPI.md#GetBillingInvoicesByIdPdf) | **Get** /v1/billing/invoices/{id}/pdf | Download one invoice as a PDF
[**GetBillingLedger**](BillingAPI.md#GetBillingLedger) | **Get** /v1/billing/ledger | Answers the org&#39;s own postings inside &#x60;range&#x3D;&#x60;, each as a signed entry: a DEPOSIT CREDITS the wallet (positive, account &#x60;credits:&lt;org&gt;&#x60;) and every other posting DEBITS it (negative, account &#x60;usage:&lt;org&gt;&#x60;), described by its notes or its tags.
[**GetBillingMethods**](BillingAPI.md#GetBillingMethods) | **Get** /v1/billing/methods | Cards and accounts on file for the caller
[**GetBillingPayouts**](BillingAPI.md#GetBillingPayouts) | **Get** /v1/billing/payouts | Answers the org&#39;s outbound payouts, newest first — amount, destination, status, and the failure reason where one applies.
[**GetBillingPlans**](BillingAPI.md#GetBillingPlans) | **Get** /v1/billing/plans | The plan catalog, priced with whatever offer is in force
[**GetBillingPortalMethods**](BillingAPI.md#GetBillingPortalMethods) | **Get** /v1/billing/portal/methods | Cards and accounts on file for the caller
[**GetBillingSettings**](BillingAPI.md#GetBillingSettings) | **Get** /v1/billing/settings | Answers the PUBLIC half of this org&#39;s processor configuration — the ids a browser needs to tokenize a card, and the environment it must tokenize against.
[**GetBillingSubscriptions**](BillingAPI.md#GetBillingSubscriptions) | **Get** /v1/billing/subscriptions | Lists the plans the caller holds, with the count beside them.
[**GetBillingTier**](BillingAPI.md#GetBillingTier) | **Get** /v1/billing/tier | Answers which tier the caller is on, what it allows, and what is left to spend.
[**GetBillingTransactions**](BillingAPI.md#GetBillingTransactions) | **Get** /v1/billing/transactions | Answers one page of the caller&#39;s own ledger, newest first: what moved, how much, when, and what it was tagged with.
[**GetBillingUsage**](BillingAPI.md#GetBillingUsage) | **Get** /v1/billing/usage | Every billed call the caller&#39;s org made, attributed to a product
[**GetBillingUsageAccounts**](BillingAPI.md#GetBillingUsageAccounts) | **Get** /v1/billing/usage/accounts | Answers per-account totals for the linked provider accounts the gateway ROUTED this caller&#39;s traffic through — requests, prompt and completion tokens, recorded cost — plus their honest sum.
[**GetBillingUsageRollup**](BillingAPI.md#GetBillingUsageRollup) | **Get** /v1/billing/usage/rollup | Answers the caller&#39;s month: what their plan includes, what has been consumed against it, and the wallet beside it.
[**GetBillingWire**](BillingAPI.md#GetBillingWire) | **Get** /v1/billing/wire | Answers where to send a wire top-up: the receiving bank details, with the caller&#39;s own payment reference.
[**GetInvoice**](BillingAPI.md#GetInvoice) | **Get** /v1/billing/invoices/{id} | Read one invoice
[**IssueInvoice**](BillingAPI.md#IssueInvoice) | **Post** /v1/billing/invoices/{id}/issue | Issue a draft invoice, making it collectible
[**PatchBillingAlertsById**](BillingAPI.md#PatchBillingAlertsById) | **Patch** /v1/billing/alerts/{id} | Changes one spend cap: raise or lower the ceiling, flip enforcement, retune the rate limit.
[**PostBillingAlerts**](BillingAPI.md#PostBillingAlerts) | **Post** /v1/billing/alerts | Opens a spend cap on the caller&#39;s own org.
[**PostBillingCryptoDeposit**](BillingAPI.md#PostBillingCryptoDeposit) | **Post** /v1/billing/crypto/deposit | Issues a deposit address the caller can send crypto to, on the asset they ask for.
[**PostBillingMethods**](BillingAPI.md#PostBillingMethods) | **Post** /v1/billing/methods | Save a card or account for the caller
[**PostBillingMode**](BillingAPI.md#PostBillingMode) | **Post** /v1/billing/mode | Moves this org between sandbox money and real money.
[**PostBillingPortalMethods**](BillingAPI.md#PostBillingPortalMethods) | **Post** /v1/billing/portal/methods | Save a card or account for the caller
[**PostBillingRechargeRunAll**](BillingAPI.md#PostBillingRechargeRunAll) | **Post** /v1/billing/recharge/run-all | Sweeps every org&#39;s auto-recharge and answers what it did.
[**PostBillingSubscribeCard**](BillingAPI.md#PostBillingSubscribeCard) | **Post** /v1/billing/subscribe/card | Buy a plan with a card
[**PostBillingTopup**](BillingAPI.md#PostBillingTopup) | **Post** /v1/billing/topup | Charges a card the caller already saved and credits the balance.
[**PostBillingTopupToken**](BillingAPI.md#PostBillingTopupToken) | **Post** /v1/billing/topup/token | Charges a single-use card token and credits the caller&#39;s balance.
[**RaiseInvoice**](BillingAPI.md#RaiseInvoice) | **Post** /v1/billing/invoices | Raise a draft invoice against a customer
[**ReactivateSubscription**](BillingAPI.md#ReactivateSubscription) | **Post** /v1/billing/subscriptions/{id}/reactivate | Put a canceled subscription back on its plan
[**VoidInvoice**](BillingAPI.md#VoidInvoice) | **Post** /v1/billing/invoices/{id}/void | Void a draft or issued invoice



## CancelSubscription

> Subscription CancelSubscription(ctx, id).SubscriptionRef(subscriptionRef).Execute()

End a subscription



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
	id := "id_example" // string | 
	subscriptionRef := *openapiclient.NewSubscriptionRef() // SubscriptionRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CancelSubscription(context.Background(), id).SubscriptionRef(subscriptionRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CancelSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CancelSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionRef** | [**SubscriptionRef**](SubscriptionRef.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectInvoice

> Collected CollectInvoice(ctx, id).Execute()

Collect an issued invoice from credits, balance, then card



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
	id := "id_example" // string | ID is the invoice id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CollectInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CollectInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectInvoice`: Collected
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CollectInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the invoice id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collected**](Collected.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBillingAlertsById

> DeleteBillingAlertsById(ctx, id).Execute()

Removes one of the caller's spend caps and answers 204.



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
	id := "id_example" // string | ID is the cap to remove, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingAPI.DeleteBillingAlertsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DeleteBillingAlertsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cap to remove, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillingAlertsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBillingMethodsById

> Detachment DeleteBillingMethodsById(ctx, id).Execute()

Removes one card or account the caller has saved.



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
	id := "id_example" // string | ID is the saved method to detach, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.DeleteBillingMethodsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DeleteBillingMethodsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBillingMethodsById`: Detachment
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.DeleteBillingMethodsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the saved method to detach, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillingMethodsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Detachment**](Detachment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBillingPortalMethodsById

> Detachment DeleteBillingPortalMethodsById(ctx, id).Execute()

DetachPortalMethod is DetachMethod at the address a hosted checkout addresses it by.



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
	id := "id_example" // string | ID is the saved method to detach, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.DeleteBillingPortalMethodsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DeleteBillingPortalMethodsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBillingPortalMethodsById`: Detachment
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.DeleteBillingPortalMethodsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the saved method to detach, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillingPortalMethodsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Detachment**](Detachment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingAccounts

> []BillingAccount GetBillingAccounts(ctx).Execute()

Answers the caller's billing accounts: the org itself, its currency, when it was opened, and the caller's own standing in it.



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
	resp, r, err := apiClient.BillingAPI.GetBillingAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingAccounts`: []BillingAccount
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingAccountsRequest struct via the builder pattern


### Return type

[**[]BillingAccount**](BillingAccount.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingAccountsByIdMembers

> []Holder GetBillingAccountsByIdMembers(ctx, id).Execute()

Answers one billing account's roster.



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
	id := "id_example" // string | ID is the billing account id, which for this store is the org's own id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetBillingAccountsByIdMembers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingAccountsByIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingAccountsByIdMembers`: []Holder
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingAccountsByIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the billing account id, which for this store is the org&#39;s own id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingAccountsByIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Holder**](Holder.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingAlerts

> []Alert GetBillingAlerts(ctx).Execute()

Lists this org's spend caps: the ceiling, its scope, whether it enforces, and how much of it has been spent this period.



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
	resp, r, err := apiClient.BillingAPI.GetBillingAlerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingAlerts`: []Alert
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingAlertsRequest struct via the builder pattern


### Return type

[**[]Alert**](Alert.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingAlertsAuthorize

> CapVerdict GetBillingAlertsAuthorize(ctx).Project(project).Service(service).Amount(amount).Pv(pv).Execute()

Answers whether one proposed spend fits inside this org's caps.



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
	project := "project_example" // string | Project narrows the verdict to one project's caps. Empty is the org-wide row. (optional)
	service := "service_example" // string | Service narrows it to one service's caps. Empty is every service. (optional)
	amount := "amount_example" // string | Amount is the proposed spend in cents. (optional)
	pv := "pv_example" // string | PV is \"1\" when the caller ESTABLISHED the project rather than merely carrying a claim of one. An unproven project may not deny traffic. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetBillingAlertsAuthorize(context.Background()).Project(project).Service(service).Amount(amount).Pv(pv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingAlertsAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingAlertsAuthorize`: CapVerdict
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingAlertsAuthorize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingAlertsAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows the verdict to one project&#39;s caps. Empty is the org-wide row. | 
 **service** | **string** | Service narrows it to one service&#39;s caps. Empty is every service. | 
 **amount** | **string** | Amount is the proposed spend in cents. | 
 **pv** | **string** | PV is \&quot;1\&quot; when the caller ESTABLISHED the project rather than merely carrying a claim of one. An unproven project may not deny traffic. | 

### Return type

[**CapVerdict**](CapVerdict.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingBalance

> GetBillingBalance(ctx).Execute()

Prepaid credit the caller's org can still spend



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
	r, err := apiClient.BillingAPI.GetBillingBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingBalanceRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingCreditBalance

> CreditBalance GetBillingCreditBalance(ctx).Execute()

Answers what the caller can spend right now, one entry per currency.



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
	resp, r, err := apiClient.BillingAPI.GetBillingCreditBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingCreditBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingCreditBalance`: CreditBalance
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingCreditBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingCreditBalanceRequest struct via the builder pattern


### Return type

[**CreditBalance**](CreditBalance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingCreditBalanceBreakdown

> interface{} GetBillingCreditBalanceBreakdown(ctx).Execute()

Answers that same spendable credit split by grant tag, with the earliest expiry under each and the total across all of them.



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
	resp, r, err := apiClient.BillingAPI.GetBillingCreditBalanceBreakdown(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingCreditBalanceBreakdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingCreditBalanceBreakdown`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingCreditBalanceBreakdown`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingCreditBalanceBreakdownRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingCredits

> CreditGrants GetBillingCredits(ctx).Execute()

Lists the caller's credit grants — every one of them, spent and lapsed and voided included.



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
	resp, r, err := apiClient.BillingAPI.GetBillingCredits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingCredits`: CreditGrants
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingCredits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingCreditsRequest struct via the builder pattern


### Return type

[**CreditGrants**](CreditGrants.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingCryptoDepositById

> CryptoDeposit GetBillingCryptoDepositById(ctx, id).Execute()

Reads one of the caller's own deposit intents back — pending, confirming, or succeeded.



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
	id := "id_example" // string | ID is the deposit intent id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetBillingCryptoDepositById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingCryptoDepositById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingCryptoDepositById`: CryptoDeposit
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingCryptoDepositById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the deposit intent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingCryptoDepositByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CryptoDeposit**](CryptoDeposit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingCryptoOptions

> CryptoOptions GetBillingCryptoOptions(ctx).Execute()

Answers which chains and tokens the crypto rail accepts — what an asset picker renders.



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
	resp, r, err := apiClient.BillingAPI.GetBillingCryptoOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingCryptoOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingCryptoOptions`: CryptoOptions
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingCryptoOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingCryptoOptionsRequest struct via the builder pattern


### Return type

[**CryptoOptions**](CryptoOptions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingInvoices

> Invoices GetBillingInvoices(ctx).Execute()

Lists the caller's invoices, newest first, with the count beside them.



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
	resp, r, err := apiClient.BillingAPI.GetBillingInvoices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingInvoices`: Invoices
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingInvoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingInvoicesRequest struct via the builder pattern


### Return type

[**Invoices**](Invoices.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingInvoicesByIdPdf

> GetBillingInvoicesByIdPdf(ctx, id).Execute()

Download one invoice as a PDF



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingAPI.GetBillingInvoicesByIdPdf(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingInvoicesByIdPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingInvoicesByIdPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingLedger

> []FinanceLedgerEntry GetBillingLedger(ctx).Range_(range_).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetBillingLedger(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingLedger`: []FinanceLedgerEntry
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingLedger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingLedgerRequest struct via the builder pattern


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


## GetBillingMethods

> GetBillingMethods(ctx).Execute()

Cards and accounts on file for the caller



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
	r, err := apiClient.BillingAPI.GetBillingMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingMethodsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingPayouts

> []Payout GetBillingPayouts(ctx).Execute()

Answers the org's outbound payouts, newest first — amount, destination, status, and the failure reason where one applies.



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
	resp, r, err := apiClient.BillingAPI.GetBillingPayouts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingPayouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingPayouts`: []Payout
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingPayouts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingPayoutsRequest struct via the builder pattern


### Return type

[**[]Payout**](Payout.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingPlans

> GetBillingPlans(ctx).Execute()

The plan catalog, priced with whatever offer is in force



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
	r, err := apiClient.BillingAPI.GetBillingPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingPlansRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingPortalMethods

> GetBillingPortalMethods(ctx).Execute()

Cards and accounts on file for the caller



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
	r, err := apiClient.BillingAPI.GetBillingPortalMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingPortalMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingPortalMethodsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingSettings

> PaymentConfig GetBillingSettings(ctx).Execute()

Answers the PUBLIC half of this org's processor configuration — the ids a browser needs to tokenize a card, and the environment it must tokenize against.



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
	resp, r, err := apiClient.BillingAPI.GetBillingSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingSettings`: PaymentConfig
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingSettingsRequest struct via the builder pattern


### Return type

[**PaymentConfig**](PaymentConfig.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingSubscriptions

> Subscriptions GetBillingSubscriptions(ctx).Execute()

Lists the plans the caller holds, with the count beside them.



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
	resp, r, err := apiClient.BillingAPI.GetBillingSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingSubscriptions`: Subscriptions
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingSubscriptionsRequest struct via the builder pattern


### Return type

[**Subscriptions**](Subscriptions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingTier

> Tier GetBillingTier(ctx).Execute()

Answers which tier the caller is on, what it allows, and what is left to spend.



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
	resp, r, err := apiClient.BillingAPI.GetBillingTier(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingTier`: Tier
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingTier`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingTierRequest struct via the builder pattern


### Return type

[**Tier**](Tier.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingTransactions

> Transactions GetBillingTransactions(ctx).Currency(currency).Limit(limit).Offset(offset).Execute()

Answers one page of the caller's own ledger, newest first: what moved, how much, when, and what it was tagged with.



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
	currency := "currency_example" // string | Currency filters to one currency. Empty reads every currency. (optional)
	limit := "limit_example" // string | Limit is the page size; absent or non-positive takes the default 100. (optional)
	offset := "offset_example" // string | Offset is how far into the history the page starts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetBillingTransactions(context.Background()).Currency(currency).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingTransactions`: Transactions
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Currency filters to one currency. Empty reads every currency. | 
 **limit** | **string** | Limit is the page size; absent or non-positive takes the default 100. | 
 **offset** | **string** | Offset is how far into the history the page starts. | 

### Return type

[**Transactions**](Transactions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingUsage

> GetBillingUsage(ctx).Execute()

Every billed call the caller's org made, attributed to a product



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
	r, err := apiClient.BillingAPI.GetBillingUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingUsageRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingUsageAccounts

> Accounts GetBillingUsageAccounts(ctx).Execute()

Answers per-account totals for the linked provider accounts the gateway ROUTED this caller's traffic through — requests, prompt and completion tokens, recorded cost — plus their honest sum.



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
	resp, r, err := apiClient.BillingAPI.GetBillingUsageAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingUsageAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingUsageAccounts`: Accounts
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingUsageAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingUsageAccountsRequest struct via the builder pattern


### Return type

[**Accounts**](Accounts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingUsageRollup

> Rollup GetBillingUsageRollup(ctx).Execute()

Answers the caller's month: what their plan includes, what has been consumed against it, and the wallet beside it.



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
	resp, r, err := apiClient.BillingAPI.GetBillingUsageRollup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingUsageRollup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingUsageRollup`: Rollup
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingUsageRollup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingUsageRollupRequest struct via the builder pattern


### Return type

[**Rollup**](Rollup.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingWire

> WireInstructions GetBillingWire(ctx).Execute()

Answers where to send a wire top-up: the receiving bank details, with the caller's own payment reference.



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
	resp, r, err := apiClient.BillingAPI.GetBillingWire(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingWire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingWire`: WireInstructions
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingWire`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingWireRequest struct via the builder pattern


### Return type

[**WireInstructions**](WireInstructions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> Invoice GetInvoice(ctx, id).Execute()

Read one invoice



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
	id := "id_example" // string | ID is the invoice id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the invoice id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueInvoice

> Invoice IssueInvoice(ctx, id).Execute()

Issue a draft invoice, making it collectible



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
	id := "id_example" // string | ID is the invoice id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.IssueInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.IssueInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.IssueInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the invoice id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchBillingAlertsById

> Alert PatchBillingAlertsById(ctx, id).AlertPatch(alertPatch).Execute()

Changes one spend cap: raise or lower the ceiling, flip enforcement, retune the rate limit.



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
	id := "id_example" // string | 
	alertPatch := *openapiclient.NewAlertPatch() // AlertPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PatchBillingAlertsById(context.Background(), id).AlertPatch(alertPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PatchBillingAlertsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchBillingAlertsById`: Alert
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PatchBillingAlertsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBillingAlertsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertPatch** | [**AlertPatch**](AlertPatch.md) |  | 

### Return type

[**Alert**](Alert.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingAlerts

> Alert PostBillingAlerts(ctx).AlertSpec(alertSpec).Execute()

Opens a spend cap on the caller's own org.



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
	alertSpec := *openapiclient.NewAlertSpec() // AlertSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PostBillingAlerts(context.Background()).AlertSpec(alertSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBillingAlerts`: Alert
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PostBillingAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertSpec** | [**AlertSpec**](AlertSpec.md) |  | 

### Return type

[**Alert**](Alert.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingCryptoDeposit

> CryptoDeposit PostBillingCryptoDeposit(ctx).CryptoAsset(cryptoAsset).Execute()

Issues a deposit address the caller can send crypto to, on the asset they ask for.



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
	cryptoAsset := *openapiclient.NewCryptoAsset() // CryptoAsset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PostBillingCryptoDeposit(context.Background()).CryptoAsset(cryptoAsset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingCryptoDeposit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBillingCryptoDeposit`: CryptoDeposit
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PostBillingCryptoDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingCryptoDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cryptoAsset** | [**CryptoAsset**](CryptoAsset.md) |  | 

### Return type

[**CryptoDeposit**](CryptoDeposit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingMethods

> PostBillingMethods(ctx).Execute()

Save a card or account for the caller



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
	r, err := apiClient.BillingAPI.PostBillingMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingMethodsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingMode

> Mode PostBillingMode(ctx).ModeIn(modeIn).Execute()

Moves this org between sandbox money and real money.



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
	modeIn := *openapiclient.NewModeIn() // ModeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PostBillingMode(context.Background()).ModeIn(modeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBillingMode`: Mode
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PostBillingMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modeIn** | [**ModeIn**](ModeIn.md) |  | 

### Return type

[**Mode**](Mode.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingPortalMethods

> PostBillingPortalMethods(ctx).Execute()

Save a card or account for the caller



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
	r, err := apiClient.BillingAPI.PostBillingPortalMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingPortalMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingPortalMethodsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingRechargeRunAll

> Recharge PostBillingRechargeRunAll(ctx).Execute()

Sweeps every org's auto-recharge and answers what it did.



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
	resp, r, err := apiClient.BillingAPI.PostBillingRechargeRunAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingRechargeRunAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBillingRechargeRunAll`: Recharge
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PostBillingRechargeRunAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingRechargeRunAllRequest struct via the builder pattern


### Return type

[**Recharge**](Recharge.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingSubscribeCard

> PostBillingSubscribeCard(ctx).Execute()

Buy a plan with a card



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
	r, err := apiClient.BillingAPI.PostBillingSubscribeCard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingSubscribeCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingSubscribeCardRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingTopup

> Charged PostBillingTopup(ctx).TopupIn(topupIn).XIdempotencyKey(xIdempotencyKey).Execute()

Charges a card the caller already saved and credits the balance.



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
	topupIn := *openapiclient.NewTopupIn() // TopupIn | 
	xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PostBillingTopup(context.Background()).TopupIn(topupIn).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingTopup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBillingTopup`: Charged
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PostBillingTopup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingTopupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topupIn** | [**TopupIn**](TopupIn.md) |  | 
 **xIdempotencyKey** | **string** |  | 

### Return type

[**Charged**](Charged.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingTopupToken

> Charged PostBillingTopupToken(ctx).TopupIn(topupIn).XIdempotencyKey(xIdempotencyKey).Execute()

Charges a single-use card token and credits the caller's balance.



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
	topupIn := *openapiclient.NewTopupIn() // TopupIn | 
	xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PostBillingTopupToken(context.Background()).TopupIn(topupIn).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PostBillingTopupToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBillingTopupToken`: Charged
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PostBillingTopupToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBillingTopupTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topupIn** | [**TopupIn**](TopupIn.md) |  | 
 **xIdempotencyKey** | **string** |  | 

### Return type

[**Charged**](Charged.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RaiseInvoice

> Invoice RaiseInvoice(ctx).RaiseIn(raiseIn).Execute()

Raise a draft invoice against a customer



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
	raiseIn := *openapiclient.NewRaiseIn() // RaiseIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RaiseInvoice(context.Background()).RaiseIn(raiseIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RaiseInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RaiseInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RaiseInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRaiseInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **raiseIn** | [**RaiseIn**](RaiseIn.md) |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactivateSubscription

> Subscription ReactivateSubscription(ctx, id).SubscriptionRef(subscriptionRef).Execute()

Put a canceled subscription back on its plan



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
	id := "id_example" // string | 
	subscriptionRef := *openapiclient.NewSubscriptionRef() // SubscriptionRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ReactivateSubscription(context.Background(), id).SubscriptionRef(subscriptionRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ReactivateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactivateSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ReactivateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionRef** | [**SubscriptionRef**](SubscriptionRef.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidInvoice

> Invoice VoidInvoice(ctx, id).Execute()

Void a draft or issued invoice



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
	id := "id_example" // string | ID is the invoice id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.VoidInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.VoidInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.VoidInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the invoice id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

