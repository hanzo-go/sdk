# \BooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBooksAccounts**](BooksAPI.md#GetBooksAccounts) | **Get** /v1/books/accounts | Returns the org&#39;s chart of accounts — the seeded fixed chart every posting key in the ledger refers to.
[**GetBooksBankTransactions**](BooksAPI.md#GetBooksBankTransactions) | **Get** /v1/books/bank/transactions | Returns the org&#39;s normalized bank transactions, newest first — every row the import and connector paths have ingested, with its amount in exact cents, its direction, and whether it has been matched to a voucher yet.
[**GetBooksBankUnreconciled**](BooksAPI.md#GetBooksBankUnreconciled) | **Get** /v1/books/bank/unreconciled | Returns the org&#39;s unmatched bank inflows and their open clarifying questions — the queue a human answers so an unexplained deposit is never guessed into revenue.
[**GetBooksExport**](BooksAPI.md#GetBooksExport) | **Get** /v1/books/export | Returns the complete financial package for the caller&#39;s org over (from, to]: the trial balance, the P&amp;L, the balance sheet, and the GL detail behind them — the four statements a tax preparer or an investor asks for, assembled from the one ledger in a single read so they cannot disagree with each other.
[**GetBooksGl**](BooksAPI.md#GetBooksGl) | **Get** /v1/books/gl | ListGL returns the org&#39;s most recent GL Entry rows, newest first.
[**GetBooksInbox**](BooksAPI.md#GetBooksInbox) | **Get** /v1/books/inbox | Returns the org&#39;s open document queue — everything uploaded but not yet booked, newest first, each with its extracted summary and the confidence the scanner resolved its category at.
[**GetBooksMetrics**](BooksAPI.md#GetBooksMetrics) | **Get** /v1/books/metrics | Metrics returns the org&#39;s deterministic SaaS-metrics snapshot over an optional (from, to] window — MRR, ARR, revenue, COGS, burn, gross margin, net income, cash, deferred revenue, monthly burn and runway — as raw int64-cent figures AND the same figures already formatted.
[**GetBooksPnl**](BooksAPI.md#GetBooksPnl) | **Get** /v1/books/pnl | Returns the org&#39;s accrual-basis Profit &amp; Loss over an optional (from, to] window of RFC3339 posting times: recognized revenue, matched cost, and the net.
[**GetBooksPosition**](BooksAPI.md#GetBooksPosition) | **Get** /v1/books/position | Returns the org&#39;s Balance Sheet as of &#x60;to&#x60; (empty &#x3D; all time), with the Assets &#x3D;&#x3D; Liabilities + Equity equation proof.
[**GetBooksQuestions**](BooksAPI.md#GetBooksQuestions) | **Get** /v1/books/questions | Returns the clarifying questions the caller&#39;s own recent GL raises — the unusual postings a founder should look at (outliers, reversals, round-offs, uncosted revenue, an overdrawn wallet), sharpest first.
[**GetBooksRules**](BooksAPI.md#GetBooksRules) | **Get** /v1/books/rules | Returns the org&#39;s auto-categorization rules, highest priority first.
[**GetBooksTransactions**](BooksAPI.md#GetBooksTransactions) | **Get** /v1/books/transactions | Returns the org&#39;s booked ledger as a single-line register, newest first: one row per voucher, with its date, description, vendor, category, source and amount in exact cents.
[**GetBooksTrial**](BooksAPI.md#GetBooksTrial) | **Get** /v1/books/trial | Returns the org&#39;s trial balance over an optional [from, to] window of RFC3339 posting times, including the opening/closing columns and the TotalDebit &#x3D;&#x3D; TotalCredit proof that the books balance.
[**GetBooksVendors**](BooksAPI.md#GetBooksVendors) | **Get** /v1/books/vendors | Returns the org&#39;s vendor book: each canonical vendor, the alias spellings a receipt may print it under, and the expense account new bills from it default to.
[**PostBooksAsk**](BooksAPI.md#PostBooksAsk) | **Post** /v1/books/ask | Answers a plain-language question about the caller&#39;s own books — \&quot;what is my MRR?\&quot;, \&quot;how long is my runway?\&quot; — with figures taken from their ledger, never a guessed number.
[**PostBooksBankExchange**](BooksAPI.md#PostBooksBankExchange) | **Post** /v1/books/bank/exchange | Finish connecting a bank account (not yet available)
[**PostBooksBankImport**](BooksAPI.md#PostBooksBankImport) | **Post** /v1/books/bank/import | Import a bank statement file into your books
[**PostBooksBankSync**](BooksAPI.md#PostBooksBankSync) | **Post** /v1/books/bank/sync | Pulls every connected bank (Plaid/Teller) for the caller&#39;s org, maps each fetched transaction to a posting and books it idempotently, then advances that connector&#39;s cursor so the next sync resumes where this one stopped.
[**PostBooksBankToken**](BooksAPI.md#PostBooksBankToken) | **Post** /v1/books/bank/token | Begin connecting a bank account (not yet available)
[**PostBooksInbox**](BooksAPI.md#PostBooksInbox) | **Post** /v1/books/inbox | Queue a document for later scanning
[**PostBooksRules**](BooksAPI.md#PostBooksRules) | **Post** /v1/books/rules | Creates or updates one auto-categorization rule, keyed by its pattern — writing a pattern that already exists REPLACES that row&#39;s category and priority.
[**PostBooksScan**](BooksAPI.md#PostBooksScan) | **Post** /v1/books/scan | Scan a receipt or invoice into a proposed voucher
[**PostBooksScanBook**](BooksAPI.md#PostBooksScanBook) | **Post** /v1/books/scan/book | Posts a reviewed scanned bill to the ledger.
[**PostBooksSync**](BooksAPI.md#PostBooksSync) | **Post** /v1/books/sync | Sync ingests the caller&#39;s OWN org from commerce into BOTH ledgers (live and sandbox) and reports how many new vouchers posted to each.
[**PostBooksVendors**](BooksAPI.md#PostBooksVendors) | **Post** /v1/books/vendors | Creates or updates one vendor in the org&#39;s vendor book, keyed by its canonical name — writing a canonical name that already exists REPLACES that row&#39;s aliases and default category.



## GetBooksAccounts

> []Account GetBooksAccounts(ctx).Sandbox(sandbox).Execute()

Returns the org's chart of accounts — the seeded fixed chart every posting key in the ledger refers to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksAccounts(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksAccounts`: []Account
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**[]Account**](Account.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksBankTransactions

> []BankTxnRow GetBooksBankTransactions(ctx).Sandbox(sandbox).Limit(limit).Execute()

Returns the org's normalized bank transactions, newest first — every row the import and connector paths have ingested, with its amount in exact cents, its direction, and whether it has been matched to a voucher yet.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	limit := int64(100) // int64 | Limit caps how many rows come back; 500 when absent or not positive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksBankTransactions(context.Background()).Sandbox(sandbox).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksBankTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksBankTransactions`: []BankTxnRow
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksBankTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksBankTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **limit** | **int64** | Limit caps how many rows come back; 500 when absent or not positive. | 

### Return type

[**[]BankTxnRow**](BankTxnRow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksBankUnreconciled

> UnreconciledOut GetBooksBankUnreconciled(ctx).Sandbox(sandbox).Execute()

Returns the org's unmatched bank inflows and their open clarifying questions — the queue a human answers so an unexplained deposit is never guessed into revenue.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksBankUnreconciled(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksBankUnreconciled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksBankUnreconciled`: UnreconciledOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksBankUnreconciled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksBankUnreconciledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**UnreconciledOut**](UnreconciledOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksExport

> FinancialPackage GetBooksExport(ctx).Sandbox(sandbox).From(from).To(to).Format(format).Limit(limit).Execute()

Returns the complete financial package for the caller's org over (from, to]: the trial balance, the P&L, the balance sheet, and the GL detail behind them — the four statements a tax preparer or an investor asks for, assembled from the one ledger in a single read so they cannot disagree with each other.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "2026-01-01T00:00:00Z" // string | From is the RFC3339 start of the window, exclusive. Empty means all time. (optional)
	to := "2026-12-31T23:59:59Z" // string | To is the RFC3339 end of the window, inclusive. Empty means up to now. (optional)
	format := "json" // string | Format is the export encoding. Only \"json\" is supported; empty means json. (optional)
	limit := int64(789) // int64 | Limit caps the GL detail rows included as the audit trail; 5000 when absent or not positive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksExport(context.Background()).Sandbox(sandbox).From(from).To(to).Format(format).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksExport`: FinancialPackage
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the window, exclusive. Empty means all time. | 
 **to** | **string** | To is the RFC3339 end of the window, inclusive. Empty means up to now. | 
 **format** | **string** | Format is the export encoding. Only \&quot;json\&quot; is supported; empty means json. | 
 **limit** | **int64** | Limit caps the GL detail rows included as the audit trail; 5000 when absent or not positive. | 

### Return type

[**FinancialPackage**](FinancialPackage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksGl

> []GLRow GetBooksGl(ctx).Sandbox(sandbox).Limit(limit).Execute()

ListGL returns the org's most recent GL Entry rows, newest first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	limit := int64(100) // int64 | Limit caps how many rows come back; 500 when absent or not positive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksGl(context.Background()).Sandbox(sandbox).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksGl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksGl`: []GLRow
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksGl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksGlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **limit** | **int64** | Limit caps how many rows come back; 500 when absent or not positive. | 

### Return type

[**[]GLRow**](GLRow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksInbox

> InboxOut GetBooksInbox(ctx).Sandbox(sandbox).Execute()

Returns the org's open document queue — everything uploaded but not yet booked, newest first, each with its extracted summary and the confidence the scanner resolved its category at.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksInbox(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksInbox`: InboxOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**InboxOut**](InboxOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksMetrics

> MetricsResponse GetBooksMetrics(ctx).Sandbox(sandbox).From(from).To(to).Execute()

Metrics returns the org's deterministic SaaS-metrics snapshot over an optional (from, to] window — MRR, ARR, revenue, COGS, burn, gross margin, net income, cash, deferred revenue, monthly burn and runway — as raw int64-cent figures AND the same figures already formatted.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "2026-01-01T00:00:00Z" // string | From is the RFC3339 start of the window, exclusive. Empty means all time. (optional)
	to := "2026-06-30T23:59:59Z" // string | To is the RFC3339 end of the window, inclusive. Empty means up to now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksMetrics(context.Background()).Sandbox(sandbox).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksMetrics`: MetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the window, exclusive. Empty means all time. | 
 **to** | **string** | To is the RFC3339 end of the window, inclusive. Empty means up to now. | 

### Return type

[**MetricsResponse**](MetricsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksPnl

> PnL GetBooksPnl(ctx).Sandbox(sandbox).From(from).To(to).Execute()

Returns the org's accrual-basis Profit & Loss over an optional (from, to] window of RFC3339 posting times: recognized revenue, matched cost, and the net.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "2026-01-01T00:00:00Z" // string | From is the RFC3339 start of the window, exclusive. Empty means all time. (optional)
	to := "2026-03-31T23:59:59Z" // string | To is the RFC3339 end of the window, inclusive. Empty means up to now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksPnl(context.Background()).Sandbox(sandbox).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksPnl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksPnl`: PnL
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksPnl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksPnlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the window, exclusive. Empty means all time. | 
 **to** | **string** | To is the RFC3339 end of the window, inclusive. Empty means up to now. | 

### Return type

[**PnL**](PnL.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksPosition

> BalanceSheet GetBooksPosition(ctx).Sandbox(sandbox).To(to).Execute()

Returns the org's Balance Sheet as of `to` (empty = all time), with the Assets == Liabilities + Equity equation proof.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	to := "2026-03-31T23:59:59Z" // string | To is the RFC3339 instant the statement is struck as of. Empty means all time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksPosition(context.Background()).Sandbox(sandbox).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksPosition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksPosition`: BalanceSheet
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksPosition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **to** | **string** | To is the RFC3339 instant the statement is struck as of. Empty means all time. | 

### Return type

[**BalanceSheet**](BalanceSheet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksQuestions

> QuestionsResponse GetBooksQuestions(ctx).Sandbox(sandbox).Execute()

Returns the clarifying questions the caller's own recent GL raises — the unusual postings a founder should look at (outliers, reversals, round-offs, uncosted revenue, an overdrawn wallet), sharpest first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksQuestions(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksQuestions`: QuestionsResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksQuestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**QuestionsResponse**](QuestionsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksRules

> RulesOut GetBooksRules(ctx).Sandbox(sandbox).Execute()

Returns the org's auto-categorization rules, highest priority first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksRules(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksRules`: RulesOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**RulesOut**](RulesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksTransactions

> TransactionsOut GetBooksTransactions(ctx).Sandbox(sandbox).From(from).To(to).Category(category).Vendor(vendor).Limit(limit).Execute()

Returns the org's booked ledger as a single-line register, newest first: one row per voucher, with its date, description, vendor, category, source and amount in exact cents.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "from_example" // string | From is the RFC3339 start of the posting-time window, inclusive. (optional)
	to := "to_example" // string | To is the RFC3339 end of the posting-time window, inclusive. (optional)
	category := "software" // string | Category filters to one COA account, named by number (\"5300\") or by category slug (\"software\"). (optional)
	vendor := "vendor_example" // string | Vendor filters to rows whose vendor or description contains this text, case-insensitively. (optional)
	limit := int64(50) // int64 | Limit caps how many rows come back; 200 when absent or not positive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksTransactions(context.Background()).Sandbox(sandbox).From(from).To(to).Category(category).Vendor(vendor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksTransactions`: TransactionsOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the posting-time window, inclusive. | 
 **to** | **string** | To is the RFC3339 end of the posting-time window, inclusive. | 
 **category** | **string** | Category filters to one COA account, named by number (\&quot;5300\&quot;) or by category slug (\&quot;software\&quot;). | 
 **vendor** | **string** | Vendor filters to rows whose vendor or description contains this text, case-insensitively. | 
 **limit** | **int64** | Limit caps how many rows come back; 200 when absent or not positive. | 

### Return type

[**TransactionsOut**](TransactionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksTrial

> TrialBalance GetBooksTrial(ctx).Sandbox(sandbox).From(from).To(to).Execute()

Returns the org's trial balance over an optional [from, to] window of RFC3339 posting times, including the opening/closing columns and the TotalDebit == TotalCredit proof that the books balance.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "2026-01-01T00:00:00Z" // string | From is the RFC3339 start of the window, exclusive. Empty means all time. (optional)
	to := "2026-03-31T23:59:59Z" // string | To is the RFC3339 end of the window, inclusive. Empty means up to now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksTrial(context.Background()).Sandbox(sandbox).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksTrial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksTrial`: TrialBalance
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksTrial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the window, exclusive. Empty means all time. | 
 **to** | **string** | To is the RFC3339 end of the window, inclusive. Empty means up to now. | 

### Return type

[**TrialBalance**](TrialBalance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooksVendors

> VendorsOut GetBooksVendors(ctx).Sandbox(sandbox).Execute()

Returns the org's vendor book: each canonical vendor, the alias spellings a receipt may print it under, and the expense account new bills from it default to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooksVendors(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooksVendors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooksVendors`: VendorsOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooksVendors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**VendorsOut**](VendorsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksAsk

> AskResponse PostBooksAsk(ctx).AskRequest(askRequest).Execute()

Answers a plain-language question about the caller's own books — \"what is my MRR?\", \"how long is my runway?\" — with figures taken from their ledger, never a guessed number.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	askRequest := *openapiclient.NewAskRequest() // AskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksAsk(context.Background()).AskRequest(askRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksAsk`: AskResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksAsk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **askRequest** | [**AskRequest**](AskRequest.md) |  | 

### Return type

[**AskResponse**](AskResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksBankExchange

> PostBooksBankExchange(ctx).Execute()

Finish connecting a bank account (not yet available)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BooksAPI.PostBooksBankExchange(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksBankExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksBankExchangeRequest struct via the builder pattern


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


## PostBooksBankImport

> BankTally PostBooksBankImport(ctx).Body(body).Execute()

Import a bank statement file into your books



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksBankImport(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksBankImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksBankImport`: BankTally
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksBankImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksBankImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**BankTally**](BankTally.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksBankSync

> BankTally PostBooksBankSync(ctx).Execute()

Pulls every connected bank (Plaid/Teller) for the caller's org, maps each fetched transaction to a posting and books it idempotently, then advances that connector's cursor so the next sync resumes where this one stopped.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksBankSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksBankSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksBankSync`: BankTally
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksBankSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksBankSyncRequest struct via the builder pattern


### Return type

[**BankTally**](BankTally.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksBankToken

> PostBooksBankToken(ctx).Execute()

Begin connecting a bank account (not yet available)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BooksAPI.PostBooksBankToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksBankToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksBankTokenRequest struct via the builder pattern


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


## PostBooksInbox

> InboxItem PostBooksInbox(ctx).Body(body).Execute()

Queue a document for later scanning



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksInbox(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksInbox`: InboxItem
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**InboxItem**](InboxItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksRules

> Rule PostBooksRules(ctx).Rule(rule).Execute()

Creates or updates one auto-categorization rule, keyed by its pattern — writing a pattern that already exists REPLACES that row's category and priority.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	rule := *openapiclient.NewRule() // Rule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksRules(context.Background()).Rule(rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksRules`: Rule
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rule** | [**Rule**](Rule.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksScan

> ScanDraft PostBooksScan(ctx).Body(body).Execute()

Scan a receipt or invoice into a proposed voucher



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksScan(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksScan`: ScanDraft
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**ScanDraft**](ScanDraft.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksScanBook

> BookResponse PostBooksScanBook(ctx).BookRequest(bookRequest).Execute()

Posts a reviewed scanned bill to the ledger.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	bookRequest := *openapiclient.NewBookRequest() // BookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksScanBook(context.Background()).BookRequest(bookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksScanBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksScanBook`: BookResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksScanBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksScanBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookRequest** | [**BookRequest**](BookRequest.md) |  | 

### Return type

[**BookResponse**](BookResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksSync

> SyncTally PostBooksSync(ctx).Execute()

Sync ingests the caller's OWN org from commerce into BOTH ledgers (live and sandbox) and reports how many new vouchers posted to each.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksSync`: SyncTally
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksSyncRequest struct via the builder pattern


### Return type

[**SyncTally**](SyncTally.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBooksVendors

> VendorRow PostBooksVendors(ctx).VendorRow(vendorRow).Execute()

Creates or updates one vendor in the org's vendor book, keyed by its canonical name — writing a canonical name that already exists REPLACES that row's aliases and default category.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	vendorRow := *openapiclient.NewVendorRow() // VendorRow | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.PostBooksVendors(context.Background()).VendorRow(vendorRow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.PostBooksVendors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooksVendors`: VendorRow
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.PostBooksVendors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBooksVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendorRow** | [**VendorRow**](VendorRow.md) |  | 

### Return type

[**VendorRow**](VendorRow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

