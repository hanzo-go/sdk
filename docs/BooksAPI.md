# \BooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1BooksAccounts**](BooksAPI.md#CloudGetV1BooksAccounts) | **Get** /v1/books/accounts | ListAccounts returns the org&#39;s chart of accounts — the seeded fixed chart every posting key in the ledger refers to.
[**CloudGetV1BooksBalanceSheet**](BooksAPI.md#CloudGetV1BooksBalanceSheet) | **Get** /v1/books/balance-sheet | BalanceSheet returns the org&#39;s Balance Sheet as of &#x60;to&#x60; (empty &#x3D; all time), with the Assets &#x3D;&#x3D; Liabilities + Equity equation proof.
[**CloudGetV1BooksBankTransactions**](BooksAPI.md#CloudGetV1BooksBankTransactions) | **Get** /v1/books/bank/transactions | ListBankTransactions returns the org&#39;s normalized bank transactions, newest first — every row the import and connector paths have ingested, with its amount in exact cents, its direction, and whether it has been matched to a voucher yet.
[**CloudGetV1BooksBankUnreconciled**](BooksAPI.md#CloudGetV1BooksBankUnreconciled) | **Get** /v1/books/bank/unreconciled | ListUnreconciled returns the org&#39;s unmatched bank inflows and their open clarifying questions — the queue a human answers so an unexplained deposit is never guessed into revenue.
[**CloudGetV1BooksExport**](BooksAPI.md#CloudGetV1BooksExport) | **Get** /v1/books/export | ExportPackage returns the complete financial package for the caller&#39;s org over (from, to]: the trial balance, the P&amp;L, the balance sheet, and the GL detail behind them — the four statements a tax preparer or an investor asks for, assembled from the one ledger in a single read so they cannot disagree with each other.
[**CloudGetV1BooksGl**](BooksAPI.md#CloudGetV1BooksGl) | **Get** /v1/books/gl | ListGL returns the org&#39;s most recent GL Entry rows, newest first.
[**CloudGetV1BooksInbox**](BooksAPI.md#CloudGetV1BooksInbox) | **Get** /v1/books/inbox | ListInbox returns the org&#39;s open document queue — everything uploaded but not yet booked, newest first, each with its extracted summary and the confidence the scanner resolved its category at.
[**CloudGetV1BooksMetrics**](BooksAPI.md#CloudGetV1BooksMetrics) | **Get** /v1/books/metrics | Metrics returns the org&#39;s deterministic SaaS-metrics snapshot over an optional (from, to] window — MRR, ARR, revenue, COGS, burn, gross margin, net income, cash, deferred revenue, monthly burn and runway — as raw int64-cent figures AND the same figures already formatted.
[**CloudGetV1BooksPnl**](BooksAPI.md#CloudGetV1BooksPnl) | **Get** /v1/books/pnl | ProfitAndLoss returns the org&#39;s accrual-basis Profit &amp; Loss over an optional (from, to] window of RFC3339 posting times: recognized revenue, matched cost, and the net.
[**CloudGetV1BooksQuestions**](BooksAPI.md#CloudGetV1BooksQuestions) | **Get** /v1/books/questions | ListQuestions returns the clarifying questions the caller&#39;s own recent GL raises — the unusual postings a founder should look at (outliers, reversals, round-offs, uncosted revenue, an overdrawn wallet), sharpest first.
[**CloudGetV1BooksRules**](BooksAPI.md#CloudGetV1BooksRules) | **Get** /v1/books/rules | ListRules returns the org&#39;s auto-categorization rules, highest priority first.
[**CloudGetV1BooksTransactions**](BooksAPI.md#CloudGetV1BooksTransactions) | **Get** /v1/books/transactions | ListTransactions returns the org&#39;s booked ledger as a single-line register, newest first: one row per voucher, with its date, description, vendor, category, source and amount in exact cents.
[**CloudGetV1BooksTrialBalance**](BooksAPI.md#CloudGetV1BooksTrialBalance) | **Get** /v1/books/trial-balance | TrialBalance returns the org&#39;s trial balance over an optional [from, to] window of RFC3339 posting times, including the opening/closing columns and the TotalDebit &#x3D;&#x3D; TotalCredit proof that the books balance.
[**CloudGetV1BooksVendors**](BooksAPI.md#CloudGetV1BooksVendors) | **Get** /v1/books/vendors | ListVendors returns the org&#39;s vendor book: each canonical vendor, the alias spellings a receipt may print it under, and the expense account new bills from it default to.
[**CloudPostV1BooksAsk**](BooksAPI.md#CloudPostV1BooksAsk) | **Post** /v1/books/ask | AskBooks answers a plain-language question about the caller&#39;s own books — \&quot;what is my MRR?\&quot;, \&quot;how long is my runway?\&quot; — with figures taken from their ledger, never a guessed number.
[**CloudPostV1BooksBankExchange**](BooksAPI.md#CloudPostV1BooksBankExchange) | **Post** /v1/books/bank/exchange | 
[**CloudPostV1BooksBankImport**](BooksAPI.md#CloudPostV1BooksBankImport) | **Post** /v1/books/bank/import | 
[**CloudPostV1BooksBankLinkToken**](BooksAPI.md#CloudPostV1BooksBankLinkToken) | **Post** /v1/books/bank/link-token | 
[**CloudPostV1BooksBankSync**](BooksAPI.md#CloudPostV1BooksBankSync) | **Post** /v1/books/bank/sync | SyncBank pulls every connected bank (Plaid/Teller) for the caller&#39;s org, maps each fetched transaction to a posting and books it idempotently, then advances that connector&#39;s cursor so the next sync resumes where this one stopped.
[**CloudPostV1BooksInbox**](BooksAPI.md#CloudPostV1BooksInbox) | **Post** /v1/books/inbox | 
[**CloudPostV1BooksRules**](BooksAPI.md#CloudPostV1BooksRules) | **Post** /v1/books/rules | UpsertRule creates or updates one auto-categorization rule, keyed by its pattern — writing a pattern that already exists REPLACES that row&#39;s category and priority.
[**CloudPostV1BooksScan**](BooksAPI.md#CloudPostV1BooksScan) | **Post** /v1/books/scan | 
[**CloudPostV1BooksScanBook**](BooksAPI.md#CloudPostV1BooksScanBook) | **Post** /v1/books/scan/book | BookScan posts a reviewed scanned bill to the ledger.
[**CloudPostV1BooksSync**](BooksAPI.md#CloudPostV1BooksSync) | **Post** /v1/books/sync | Sync ingests the caller&#39;s OWN org from commerce into BOTH ledgers (live and sandbox) and reports how many new vouchers posted to each.
[**CloudPostV1BooksVendors**](BooksAPI.md#CloudPostV1BooksVendors) | **Post** /v1/books/vendors | UpsertVendor creates or updates one vendor in the org&#39;s vendor book, keyed by its canonical name — writing a canonical name that already exists REPLACES that row&#39;s aliases and default category.



## CloudGetV1BooksAccounts

> []CloudAccount CloudGetV1BooksAccounts(ctx).Sandbox(sandbox).Execute()

ListAccounts returns the org's chart of accounts — the seeded fixed chart every posting key in the ledger refers to.



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
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksAccounts(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksAccounts`: []CloudAccount
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**[]CloudAccount**](CloudAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksBalanceSheet

> CloudBalanceSheet CloudGetV1BooksBalanceSheet(ctx).Sandbox(sandbox).To(to).Execute()

BalanceSheet returns the org's Balance Sheet as of `to` (empty = all time), with the Assets == Liabilities + Equity equation proof.



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
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	to := "2026-03-31T23:59:59Z" // string | To is the RFC3339 instant the statement is struck as of. Empty means all time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksBalanceSheet(context.Background()).Sandbox(sandbox).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksBalanceSheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksBalanceSheet`: CloudBalanceSheet
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksBalanceSheet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksBalanceSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **to** | **string** | To is the RFC3339 instant the statement is struck as of. Empty means all time. | 

### Return type

[**CloudBalanceSheet**](CloudBalanceSheet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksBankTransactions

> []CloudBankTxnRow CloudGetV1BooksBankTransactions(ctx).Sandbox(sandbox).Limit(limit).Execute()

ListBankTransactions returns the org's normalized bank transactions, newest first — every row the import and connector paths have ingested, with its amount in exact cents, its direction, and whether it has been matched to a voucher yet.



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
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	limit := int32(100) // int32 | Limit caps how many rows come back; 500 when absent or not positive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksBankTransactions(context.Background()).Sandbox(sandbox).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksBankTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksBankTransactions`: []CloudBankTxnRow
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksBankTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksBankTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **limit** | **int32** | Limit caps how many rows come back; 500 when absent or not positive. | 

### Return type

[**[]CloudBankTxnRow**](CloudBankTxnRow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksBankUnreconciled

> CloudUnreconciledOut CloudGetV1BooksBankUnreconciled(ctx).Sandbox(sandbox).Execute()

ListUnreconciled returns the org's unmatched bank inflows and their open clarifying questions — the queue a human answers so an unexplained deposit is never guessed into revenue.



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
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksBankUnreconciled(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksBankUnreconciled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksBankUnreconciled`: CloudUnreconciledOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksBankUnreconciled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksBankUnreconciledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**CloudUnreconciledOut**](CloudUnreconciledOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksExport

> CloudFinancialPackage CloudGetV1BooksExport(ctx).Sandbox(sandbox).From(from).To(to).Format(format).Limit(limit).Execute()

ExportPackage returns the complete financial package for the caller's org over (from, to]: the trial balance, the P&L, the balance sheet, and the GL detail behind them — the four statements a tax preparer or an investor asks for, assembled from the one ledger in a single read so they cannot disagree with each other.



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
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "2026-01-01T00:00:00Z" // string | From is the RFC3339 start of the window, exclusive. Empty means all time. (optional)
	to := "2026-12-31T23:59:59Z" // string | To is the RFC3339 end of the window, inclusive. Empty means up to now. (optional)
	format := "json" // string | Format is the export encoding. Only \"json\" is supported; empty means json. (optional)
	limit := int32(56) // int32 | Limit caps the GL detail rows included as the audit trail; 5000 when absent or not positive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksExport(context.Background()).Sandbox(sandbox).From(from).To(to).Format(format).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksExport`: CloudFinancialPackage
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the window, exclusive. Empty means all time. | 
 **to** | **string** | To is the RFC3339 end of the window, inclusive. Empty means up to now. | 
 **format** | **string** | Format is the export encoding. Only \&quot;json\&quot; is supported; empty means json. | 
 **limit** | **int32** | Limit caps the GL detail rows included as the audit trail; 5000 when absent or not positive. | 

### Return type

[**CloudFinancialPackage**](CloudFinancialPackage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksGl

> []CloudGLRow CloudGetV1BooksGl(ctx).Sandbox(sandbox).Limit(limit).Execute()

ListGL returns the org's most recent GL Entry rows, newest first.



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
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	limit := int32(100) // int32 | Limit caps how many rows come back; 500 when absent or not positive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksGl(context.Background()).Sandbox(sandbox).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksGl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksGl`: []CloudGLRow
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksGl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksGlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **limit** | **int32** | Limit caps how many rows come back; 500 when absent or not positive. | 

### Return type

[**[]CloudGLRow**](CloudGLRow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksInbox

> CloudInboxOut CloudGetV1BooksInbox(ctx).Sandbox(sandbox).Execute()

ListInbox returns the org's open document queue — everything uploaded but not yet booked, newest first, each with its extracted summary and the confidence the scanner resolved its category at.



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
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksInbox(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksInbox`: CloudInboxOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**CloudInboxOut**](CloudInboxOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksMetrics

> CloudMetricsResponse CloudGetV1BooksMetrics(ctx).Sandbox(sandbox).From(from).To(to).Execute()

Metrics returns the org's deterministic SaaS-metrics snapshot over an optional (from, to] window — MRR, ARR, revenue, COGS, burn, gross margin, net income, cash, deferred revenue, monthly burn and runway — as raw int64-cent figures AND the same figures already formatted.



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
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "2026-01-01T00:00:00Z" // string | From is the RFC3339 start of the window, exclusive. Empty means all time. (optional)
	to := "2026-06-30T23:59:59Z" // string | To is the RFC3339 end of the window, inclusive. Empty means up to now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksMetrics(context.Background()).Sandbox(sandbox).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksMetrics`: CloudMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the window, exclusive. Empty means all time. | 
 **to** | **string** | To is the RFC3339 end of the window, inclusive. Empty means up to now. | 

### Return type

[**CloudMetricsResponse**](CloudMetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksPnl

> CloudPnL CloudGetV1BooksPnl(ctx).Sandbox(sandbox).From(from).To(to).Execute()

ProfitAndLoss returns the org's accrual-basis Profit & Loss over an optional (from, to] window of RFC3339 posting times: recognized revenue, matched cost, and the net.



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
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "2026-01-01T00:00:00Z" // string | From is the RFC3339 start of the window, exclusive. Empty means all time. (optional)
	to := "2026-03-31T23:59:59Z" // string | To is the RFC3339 end of the window, inclusive. Empty means up to now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksPnl(context.Background()).Sandbox(sandbox).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksPnl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksPnl`: CloudPnL
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksPnl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksPnlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the window, exclusive. Empty means all time. | 
 **to** | **string** | To is the RFC3339 end of the window, inclusive. Empty means up to now. | 

### Return type

[**CloudPnL**](CloudPnL.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksQuestions

> CloudQuestionsResponse CloudGetV1BooksQuestions(ctx).Sandbox(sandbox).Execute()

ListQuestions returns the clarifying questions the caller's own recent GL raises — the unusual postings a founder should look at (outliers, reversals, round-offs, uncosted revenue, an overdrawn wallet), sharpest first.



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
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksQuestions(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksQuestions`: CloudQuestionsResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksQuestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**CloudQuestionsResponse**](CloudQuestionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksRules

> CloudRulesOut CloudGetV1BooksRules(ctx).Sandbox(sandbox).Execute()

ListRules returns the org's auto-categorization rules, highest priority first.



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
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksRules(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksRules`: CloudRulesOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**CloudRulesOut**](CloudRulesOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksTransactions

> CloudTransactionsOut CloudGetV1BooksTransactions(ctx).Sandbox(sandbox).From(from).To(to).Category(category).Vendor(vendor).Limit(limit).Execute()

ListTransactions returns the org's booked ledger as a single-line register, newest first: one row per voucher, with its date, description, vendor, category, source and amount in exact cents.



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
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "from_example" // string | From is the RFC3339 start of the posting-time window, inclusive. (optional)
	to := "to_example" // string | To is the RFC3339 end of the posting-time window, inclusive. (optional)
	category := "software" // string | Category filters to one COA account, named by number (\"5300\") or by category slug (\"software\"). (optional)
	vendor := "vendor_example" // string | Vendor filters to rows whose vendor or description contains this text, case-insensitively. (optional)
	limit := int32(50) // int32 | Limit caps how many rows come back; 200 when absent or not positive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksTransactions(context.Background()).Sandbox(sandbox).From(from).To(to).Category(category).Vendor(vendor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksTransactions`: CloudTransactionsOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the posting-time window, inclusive. | 
 **to** | **string** | To is the RFC3339 end of the posting-time window, inclusive. | 
 **category** | **string** | Category filters to one COA account, named by number (\&quot;5300\&quot;) or by category slug (\&quot;software\&quot;). | 
 **vendor** | **string** | Vendor filters to rows whose vendor or description contains this text, case-insensitively. | 
 **limit** | **int32** | Limit caps how many rows come back; 200 when absent or not positive. | 

### Return type

[**CloudTransactionsOut**](CloudTransactionsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksTrialBalance

> CloudTrialBalance CloudGetV1BooksTrialBalance(ctx).Sandbox(sandbox).From(from).To(to).Execute()

TrialBalance returns the org's trial balance over an optional [from, to] window of RFC3339 posting times, including the opening/closing columns and the TotalDebit == TotalCredit proof that the books balance.



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
	sandbox := "sandbox_example" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\". (optional)
	from := "2026-01-01T00:00:00Z" // string | From is the RFC3339 start of the window, exclusive. Empty means all time. (optional)
	to := "2026-03-31T23:59:59Z" // string | To is the RFC3339 end of the window, inclusive. Empty means up to now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksTrialBalance(context.Background()).Sandbox(sandbox).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksTrialBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksTrialBalance`: CloudTrialBalance
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksTrialBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksTrialBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;. | 
 **from** | **string** | From is the RFC3339 start of the window, exclusive. Empty means all time. | 
 **to** | **string** | To is the RFC3339 end of the window, inclusive. Empty means up to now. | 

### Return type

[**CloudTrialBalance**](CloudTrialBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BooksVendors

> CloudVendorsOut CloudGetV1BooksVendors(ctx).Sandbox(sandbox).Execute()

ListVendors returns the org's vendor book: each canonical vendor, the alias spellings a receipt may print it under, and the expense account new bills from it default to.



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
	sandbox := "false" // string | Sandbox reads the org's SANDBOX ledger when it is exactly \"true\"; anything else reads the live one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudGetV1BooksVendors(context.Background()).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudGetV1BooksVendors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BooksVendors`: CloudVendorsOut
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudGetV1BooksVendors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BooksVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandbox** | **string** | Sandbox reads the org&#39;s SANDBOX ledger when it is exactly \&quot;true\&quot;; anything else reads the live one. | 

### Return type

[**CloudVendorsOut**](CloudVendorsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksAsk

> CloudAskResponse CloudPostV1BooksAsk(ctx).CloudAskRequest(cloudAskRequest).Execute()

AskBooks answers a plain-language question about the caller's own books — \"what is my MRR?\", \"how long is my runway?\" — with figures taken from their ledger, never a guessed number.



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
	cloudAskRequest := *openapiclient.NewCloudAskRequest() // CloudAskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksAsk(context.Background()).CloudAskRequest(cloudAskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksAsk`: CloudAskResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksAsk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAskRequest** | [**CloudAskRequest**](CloudAskRequest.md) |  | 

### Return type

[**CloudAskResponse**](CloudAskResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksBankExchange

> CloudPostV1BooksBankExchange(ctx).Execute()



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
	r, err := apiClient.BooksAPI.CloudPostV1BooksBankExchange(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksBankExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksBankExchangeRequest struct via the builder pattern


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


## CloudPostV1BooksBankImport

> CloudBankTally CloudPostV1BooksBankImport(ctx).Body(body).Execute()



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
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksBankImport(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksBankImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksBankImport`: CloudBankTally
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksBankImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksBankImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**CloudBankTally**](CloudBankTally.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksBankLinkToken

> CloudPostV1BooksBankLinkToken(ctx).Execute()



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
	r, err := apiClient.BooksAPI.CloudPostV1BooksBankLinkToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksBankLinkToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksBankLinkTokenRequest struct via the builder pattern


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


## CloudPostV1BooksBankSync

> CloudBankTally CloudPostV1BooksBankSync(ctx).Execute()

SyncBank pulls every connected bank (Plaid/Teller) for the caller's org, maps each fetched transaction to a posting and books it idempotently, then advances that connector's cursor so the next sync resumes where this one stopped.



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
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksBankSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksBankSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksBankSync`: CloudBankTally
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksBankSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksBankSyncRequest struct via the builder pattern


### Return type

[**CloudBankTally**](CloudBankTally.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksInbox

> CloudInboxItem CloudPostV1BooksInbox(ctx).Body(body).Execute()



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
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksInbox(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksInbox`: CloudInboxItem
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**CloudInboxItem**](CloudInboxItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksRules

> CloudRule CloudPostV1BooksRules(ctx).CloudRule(cloudRule).Execute()

UpsertRule creates or updates one auto-categorization rule, keyed by its pattern — writing a pattern that already exists REPLACES that row's category and priority.



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
	cloudRule := *openapiclient.NewCloudRule() // CloudRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksRules(context.Background()).CloudRule(cloudRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksRules`: CloudRule
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRule** | [**CloudRule**](CloudRule.md) |  | 

### Return type

[**CloudRule**](CloudRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksScan

> CloudScanDraft CloudPostV1BooksScan(ctx).Body(body).Execute()



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
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksScan(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksScan`: CloudScanDraft
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**CloudScanDraft**](CloudScanDraft.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksScanBook

> CloudBookResponse CloudPostV1BooksScanBook(ctx).CloudBookRequest(cloudBookRequest).Execute()

BookScan posts a reviewed scanned bill to the ledger.



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
	cloudBookRequest := *openapiclient.NewCloudBookRequest() // CloudBookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksScanBook(context.Background()).CloudBookRequest(cloudBookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksScanBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksScanBook`: CloudBookResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksScanBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksScanBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudBookRequest** | [**CloudBookRequest**](CloudBookRequest.md) |  | 

### Return type

[**CloudBookResponse**](CloudBookResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksSync

> CloudSyncTally CloudPostV1BooksSync(ctx).Execute()

Sync ingests the caller's OWN org from commerce into BOTH ledgers (live and sandbox) and reports how many new vouchers posted to each.



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
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksSync`: CloudSyncTally
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksSyncRequest struct via the builder pattern


### Return type

[**CloudSyncTally**](CloudSyncTally.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BooksVendors

> CloudVendorRow CloudPostV1BooksVendors(ctx).CloudVendorRow(cloudVendorRow).Execute()

UpsertVendor creates or updates one vendor in the org's vendor book, keyed by its canonical name — writing a canonical name that already exists REPLACES that row's aliases and default category.



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
	cloudVendorRow := *openapiclient.NewCloudVendorRow() // CloudVendorRow | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.CloudPostV1BooksVendors(context.Background()).CloudVendorRow(cloudVendorRow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.CloudPostV1BooksVendors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BooksVendors`: CloudVendorRow
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.CloudPostV1BooksVendors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BooksVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudVendorRow** | [**CloudVendorRow**](CloudVendorRow.md) |  | 

### Return type

[**CloudVendorRow**](CloudVendorRow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

