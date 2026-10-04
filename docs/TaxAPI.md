# \TaxAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaxFilings**](TaxAPI.md#GetTaxFilings) | **Get** /v1/tax/filings | Lists the caller org&#39;s IRS return exports and where each stands, without their files.
[**GetTaxFilingsById**](TaxAPI.md#GetTaxFilingsById) | **Get** /v1/tax/filings/{id} | Reads one IRS return export WITH its files, rebuilt from the sealed forms and proven identical to what was exported.
[**GetTaxForms**](TaxAPI.md#GetTaxForms) | **Get** /v1/tax/forms | Lists the 1099s the caller org prepared as payer, oldest first, voided and superseded ones included — a return that was furnished stays on the record.
[**GetTaxFormsById**](TaxAPI.md#GetTaxFormsById) | **Get** /v1/tax/forms/{id} | Reads one 1099 the caller org prepared as payer, every box and both parties, TINs masked.
[**GetTaxFormsByIdPdf**](TaxAPI.md#GetTaxFormsByIdPdf) | **Get** /v1/tax/forms/{id}/pdf | Download a 1099 the org prepared, as a PDF
[**GetTaxInbox**](TaxAPI.md#GetTaxInbox) | **Get** /v1/tax/inbox | Lists the Copy B statements furnished to the caller org — every 1099 a payer on Hanzo delivered to it electronically, corrected ones beside the ones they supersede.
[**GetTaxInboxById**](TaxAPI.md#GetTaxInboxById) | **Get** /v1/tax/inbox/{id} | Reads one Copy B furnished to the caller org.
[**GetTaxInboxByIdPdf**](TaxAPI.md#GetTaxInboxByIdPdf) | **Get** /v1/tax/inbox/{id}/pdf | Download a Copy B furnished to the org, as a PDF
[**GetTaxPayments**](TaxAPI.md#GetTaxPayments) | **Get** /v1/tax/payments | Returns the caller org&#39;s tax year as a PAYER: every payment it made to another org in that calendar year, on any rail, as the event plane records it, each classified — reportable or not, in which 1099 box, and the rule that decided it, with its source — and each payee&#39;s totals tested against the year&#39;s threshold.
[**GetTaxProfile**](TaxAPI.md#GetTaxProfile) | **Get** /v1/tax/profile | Returns the caller org&#39;s own tax profile — every line of its Form W-9, W-8BEN or W-8BEN-E, every number masked to its last four characters, where the certification stands, whether a W-8 is valid and when it expires, and whether the org consents to receive its 1099s electronically.
[**GetTaxW9**](TaxAPI.md#GetTaxW9) | **Get** /v1/tax/w9 | Lists the caller org&#39;s W-9 relationships on both sides: the W-9s it asked its payees for, with where each stands, and the requests other orgs made for its own W-9.
[**GetTaxW9ById**](TaxAPI.md#GetTaxW9ById) | **Get** /v1/tax/w9/{id} | Reads one W-9 relationship from the caller&#39;s side.
[**GetTaxW9ByIdTin**](TaxAPI.md#GetTaxW9ByIdTin) | **Get** /v1/tax/w9/{id}/tin | Reads a payee&#39;s TIN in full — and, for a payee on a W-8, its foreign TIN, which the payer&#39;s Form 1042-S carries.
[**PostTaxFilings**](TaxAPI.md#PostTaxFilings) | **Post** /v1/tax/filings | Exports the caller org&#39;s return for a tax year and form, as the files the IRIS Taxpayer Portal takes, and records the export.
[**PostTaxFilingsByIdReceipt**](TaxAPI.md#PostTaxFilingsByIdReceipt) | **Post** /v1/tax/filings/{id}/receipt | Records what IRIS answered for an exported return — the Receipt ID on submission, then the acknowledgment.
[**PostTaxForms**](TaxAPI.md#PostTaxForms) | **Post** /v1/tax/forms | Prepares the caller org&#39;s 1099 drafts for a tax year, as the PAYER: one 1099-NEC and/or 1099-MISC per payee whose reportable payments on Hanzo&#39;s rails reach the year&#39;s threshold, from the same derivation GET /v1/tax/payments answers.
[**PostTaxFormsByIdCorrect**](TaxAPI.md#PostTaxFormsByIdCorrect) | **Post** /v1/tax/forms/{id}/correct | Corrects a furnished (or owed) 1099 with a NEW form that supersedes it.
[**PostTaxFormsByIdFurnish**](TaxAPI.md#PostTaxFormsByIdFurnish) | **Post** /v1/tax/forms/{id}/furnish | Furnishes Copy B of a reviewed 1099 to its payee.
[**PostTaxFormsByIdMailed**](TaxAPI.md#PostTaxFormsByIdMailed) | **Post** /v1/tax/forms/{id}/mailed | Records that the payer mailed a paper Copy B it owed.
[**PostTaxFormsByIdReview**](TaxAPI.md#PostTaxFormsByIdReview) | **Post** /v1/tax/forms/{id}/review | Marks a draft 1099 reviewed — a person has checked it.
[**PostTaxFormsByIdVoid**](TaxAPI.md#PostTaxFormsByIdVoid) | **Post** /v1/tax/forms/{id}/void | Voids a draft or reviewed 1099, which is then never furnished or filed.
[**PostTaxProfileCertify**](TaxAPI.md#PostTaxProfileCertify) | **Post** /v1/tax/profile/certify | Signs the caller org&#39;s certification — Form W-9 Part II, Form W-8BEN Part III or Form W-8BEN-E Part XXX — through /v1/legal&#39;s e-signature.
[**PostTaxW9**](TaxAPI.md#PostTaxW9) | **Post** /v1/tax/w9 | Asks another org for its W-9, as the org that pays it.
[**PostTaxW9ByIdDecline**](TaxAPI.md#PostTaxW9ByIdDecline) | **Post** /v1/tax/w9/{id}/decline | Declines a request for this org&#39;s W-9.
[**PostTaxW9ByIdGrant**](TaxAPI.md#PostTaxW9ByIdGrant) | **Post** /v1/tax/w9/{id}/grant | Grants the payer that asked a live read of this org&#39;s W-9 — every line, the TIN masked, and the full TIN to that payer&#39;s org admins alone, each read audited.
[**PostTaxW9ByIdMatch**](TaxAPI.md#PostTaxW9ByIdMatch) | **Post** /v1/tax/w9/{id}/match | Records the result of IRS TIN Matching for a payee, as the payer.
[**PostTaxW9ByIdRevoke**](TaxAPI.md#PostTaxW9ByIdRevoke) | **Post** /v1/tax/w9/{id}/revoke | Revokes a W-9 this org granted.
[**PutTaxProfile**](TaxAPI.md#PutTaxProfile) | **Put** /v1/tax/profile | Writes the caller org&#39;s tax profile — its Form W-9, or its Form W-8BEN or W-8BEN-E when it is a foreign person — and answers it, every number masked.



## GetTaxFilings

> TaxFilingList GetTaxFilings(ctx).Execute()

Lists the caller org's IRS return exports and where each stands, without their files.



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
	resp, r, err := apiClient.TaxAPI.GetTaxFilings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxFilings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxFilings`: TaxFilingList
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxFilings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxFilingsRequest struct via the builder pattern


### Return type

[**TaxFilingList**](TaxFilingList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxFilingsById

> TaxFiling GetTaxFilingsById(ctx, id).Execute()

Reads one IRS return export WITH its files, rebuilt from the sealed forms and proven identical to what was exported.



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
	id := "id_example" // string | ID is the filing id, \"iris_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.GetTaxFilingsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxFilingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxFilingsById`: TaxFiling
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxFilingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the filing id, \&quot;iris_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxFilingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxFiling**](TaxFiling.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxForms

> TaxFormList GetTaxForms(ctx).Year(year).Execute()

Lists the 1099s the caller org prepared as payer, oldest first, voided and superseded ones included — a return that was furnished stays on the record.



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
	year := int64(789) // int64 | Year narrows the list to one tax year; absent lists every year. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.GetTaxForms(context.Background()).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxForms`: TaxFormList
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int64** | Year narrows the list to one tax year; absent lists every year. | 

### Return type

[**TaxFormList**](TaxFormList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxFormsById

> TaxForm GetTaxFormsById(ctx, id).Execute()

Reads one 1099 the caller org prepared as payer, every box and both parties, TINs masked.



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
	id := "id_example" // string | ID is the form id, \"f1099_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.GetTaxFormsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxFormsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxFormsById`: TaxForm
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxFormsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form id, \&quot;f1099_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxFormsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxForm**](TaxForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxFormsByIdPdf

> GetTaxFormsByIdPdf(ctx, id).Execute()

Download a 1099 the org prepared, as a PDF



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxAPI.GetTaxFormsByIdPdf(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxFormsByIdPdf``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetTaxFormsByIdPdfRequest struct via the builder pattern


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


## GetTaxInbox

> TaxStatementList GetTaxInbox(ctx).Year(year).Execute()

Lists the Copy B statements furnished to the caller org — every 1099 a payer on Hanzo delivered to it electronically, corrected ones beside the ones they supersede.



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
	year := int64(789) // int64 | Year narrows the list to one tax year; absent lists every year. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.GetTaxInbox(context.Background()).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxInbox`: TaxStatementList
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int64** | Year narrows the list to one tax year; absent lists every year. | 

### Return type

[**TaxStatementList**](TaxStatementList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxInboxById

> TaxStatement GetTaxInboxById(ctx, id).Execute()

Reads one Copy B furnished to the caller org.



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
	id := "id_example" // string | ID is the form id, \"f1099_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.GetTaxInboxById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxInboxById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxInboxById`: TaxStatement
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxInboxById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form id, \&quot;f1099_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxInboxByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxStatement**](TaxStatement.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxInboxByIdPdf

> GetTaxInboxByIdPdf(ctx, id).Execute()

Download a Copy B furnished to the org, as a PDF



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxAPI.GetTaxInboxByIdPdf(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxInboxByIdPdf``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetTaxInboxByIdPdfRequest struct via the builder pattern


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


## GetTaxPayments

> TaxLedger GetTaxPayments(ctx).Year(year).Execute()

Returns the caller org's tax year as a PAYER: every payment it made to another org in that calendar year, on any rail, as the event plane records it, each classified — reportable or not, in which 1099 box, and the rule that decided it, with its source — and each payee's totals tested against the year's threshold.



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
	year := int64(789) // int64 | Year is the calendar year the payments were made in, e.g. 2026. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.GetTaxPayments(context.Background()).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxPayments`: TaxLedger
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int64** | Year is the calendar year the payments were made in, e.g. 2026. | 

### Return type

[**TaxLedger**](TaxLedger.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxProfile

> TaxProfile GetTaxProfile(ctx).Execute()

Returns the caller org's own tax profile — every line of its Form W-9, W-8BEN or W-8BEN-E, every number masked to its last four characters, where the certification stands, whether a W-8 is valid and when it expires, and whether the org consents to receive its 1099s electronically.



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
	resp, r, err := apiClient.TaxAPI.GetTaxProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxProfile`: TaxProfile
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxProfileRequest struct via the builder pattern


### Return type

[**TaxProfile**](TaxProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxW9

> TaxW9List GetTaxW9(ctx).Execute()

Lists the caller org's W-9 relationships on both sides: the W-9s it asked its payees for, with where each stands, and the requests other orgs made for its own W-9.



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
	resp, r, err := apiClient.TaxAPI.GetTaxW9(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxW9``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxW9`: TaxW9List
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxW9`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxW9Request struct via the builder pattern


### Return type

[**TaxW9List**](TaxW9List.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxW9ById

> TaxW9 GetTaxW9ById(ctx, id).Execute()

Reads one W-9 relationship from the caller's side.



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
	id := "id_example" // string | ID is the relationship's id, \"w9_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.GetTaxW9ById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxW9ById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxW9ById`: TaxW9
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxW9ById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the relationship&#39;s id, \&quot;w9_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxW9ByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxW9**](TaxW9.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxW9ByIdTin

> TaxTinOut GetTaxW9ByIdTin(ctx, id).Execute()

Reads a payee's TIN in full — and, for a payee on a W-8, its foreign TIN, which the payer's Form 1042-S carries.



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
	id := "id_example" // string | ID is the relationship's id, \"w9_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.GetTaxW9ByIdTin(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.GetTaxW9ByIdTin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxW9ByIdTin`: TaxTinOut
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.GetTaxW9ByIdTin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the relationship&#39;s id, \&quot;w9_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxW9ByIdTinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxTinOut**](TaxTinOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxFilings

> TaxFiling PostTaxFilings(ctx).TaxExportIn(taxExportIn).Execute()

Exports the caller org's return for a tax year and form, as the files the IRIS Taxpayer Portal takes, and records the export.



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
	taxExportIn := *openapiclient.NewTaxExportIn() // TaxExportIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxFilings(context.Background()).TaxExportIn(taxExportIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxFilings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxFilings`: TaxFiling
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxFilings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxFilingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxExportIn** | [**TaxExportIn**](TaxExportIn.md) |  | 

### Return type

[**TaxFiling**](TaxFiling.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxFilingsByIdReceipt

> TaxFiling PostTaxFilingsByIdReceipt(ctx, id).TaxReceiptIn(taxReceiptIn).Execute()

Records what IRIS answered for an exported return — the Receipt ID on submission, then the acknowledgment.



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
	id := "id_example" // string | ID is the filing, from the path.
	taxReceiptIn := *openapiclient.NewTaxReceiptIn() // TaxReceiptIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxFilingsByIdReceipt(context.Background(), id).TaxReceiptIn(taxReceiptIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxFilingsByIdReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxFilingsByIdReceipt`: TaxFiling
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxFilingsByIdReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the filing, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxFilingsByIdReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxReceiptIn** | [**TaxReceiptIn**](TaxReceiptIn.md) |  | 

### Return type

[**TaxFiling**](TaxFiling.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxForms

> TaxPrepared PostTaxForms(ctx).TaxPrepareIn(taxPrepareIn).Execute()

Prepares the caller org's 1099 drafts for a tax year, as the PAYER: one 1099-NEC and/or 1099-MISC per payee whose reportable payments on Hanzo's rails reach the year's threshold, from the same derivation GET /v1/tax/payments answers.



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
	taxPrepareIn := *openapiclient.NewTaxPrepareIn() // TaxPrepareIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxForms(context.Background()).TaxPrepareIn(taxPrepareIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxForms`: TaxPrepared
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxPrepareIn** | [**TaxPrepareIn**](TaxPrepareIn.md) |  | 

### Return type

[**TaxPrepared**](TaxPrepared.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxFormsByIdCorrect

> TaxForm PostTaxFormsByIdCorrect(ctx, id).TaxCorrectIn(taxCorrectIn).Execute()

Corrects a furnished (or owed) 1099 with a NEW form that supersedes it.



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
	id := "id_example" // string | ID is the furnished form to correct, from the path.
	taxCorrectIn := *openapiclient.NewTaxCorrectIn() // TaxCorrectIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxFormsByIdCorrect(context.Background(), id).TaxCorrectIn(taxCorrectIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxFormsByIdCorrect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxFormsByIdCorrect`: TaxForm
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxFormsByIdCorrect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the furnished form to correct, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxFormsByIdCorrectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxCorrectIn** | [**TaxCorrectIn**](TaxCorrectIn.md) |  | 

### Return type

[**TaxForm**](TaxForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxFormsByIdFurnish

> TaxForm PostTaxFormsByIdFurnish(ctx, id).Execute()

Furnishes Copy B of a reviewed 1099 to its payee.



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
	id := "id_example" // string | ID is the form id, \"f1099_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxFormsByIdFurnish(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxFormsByIdFurnish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxFormsByIdFurnish`: TaxForm
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxFormsByIdFurnish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form id, \&quot;f1099_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxFormsByIdFurnishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxForm**](TaxForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxFormsByIdMailed

> TaxForm PostTaxFormsByIdMailed(ctx, id).Execute()

Records that the payer mailed a paper Copy B it owed.



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
	id := "id_example" // string | ID is the form id, \"f1099_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxFormsByIdMailed(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxFormsByIdMailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxFormsByIdMailed`: TaxForm
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxFormsByIdMailed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form id, \&quot;f1099_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxFormsByIdMailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxForm**](TaxForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxFormsByIdReview

> TaxForm PostTaxFormsByIdReview(ctx, id).Execute()

Marks a draft 1099 reviewed — a person has checked it.



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
	id := "id_example" // string | ID is the form id, \"f1099_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxFormsByIdReview(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxFormsByIdReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxFormsByIdReview`: TaxForm
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxFormsByIdReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form id, \&quot;f1099_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxFormsByIdReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxForm**](TaxForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxFormsByIdVoid

> TaxForm PostTaxFormsByIdVoid(ctx, id).Execute()

Voids a draft or reviewed 1099, which is then never furnished or filed.



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
	id := "id_example" // string | ID is the form id, \"f1099_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxFormsByIdVoid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxFormsByIdVoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxFormsByIdVoid`: TaxForm
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxFormsByIdVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form id, \&quot;f1099_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxFormsByIdVoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxForm**](TaxForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxProfileCertify

> TaxProfile PostTaxProfileCertify(ctx).Execute()

Signs the caller org's certification — Form W-9 Part II, Form W-8BEN Part III or Form W-8BEN-E Part XXX — through /v1/legal's e-signature.



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
	resp, r, err := apiClient.TaxAPI.PostTaxProfileCertify(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxProfileCertify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxProfileCertify`: TaxProfile
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxProfileCertify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxProfileCertifyRequest struct via the builder pattern


### Return type

[**TaxProfile**](TaxProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxW9

> TaxW9 PostTaxW9(ctx).TaxW9Request(taxW9Request).Execute()

Asks another org for its W-9, as the org that pays it.



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
	taxW9Request := *openapiclient.NewTaxW9Request() // TaxW9Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxW9(context.Background()).TaxW9Request(taxW9Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxW9``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxW9`: TaxW9
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxW9`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxW9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxW9Request** | [**TaxW9Request**](TaxW9Request.md) |  | 

### Return type

[**TaxW9**](TaxW9.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxW9ByIdDecline

> TaxW9 PostTaxW9ByIdDecline(ctx, id).Execute()

Declines a request for this org's W-9.



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
	id := "id_example" // string | ID is the relationship's id, \"w9_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxW9ByIdDecline(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxW9ByIdDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxW9ByIdDecline`: TaxW9
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxW9ByIdDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the relationship&#39;s id, \&quot;w9_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxW9ByIdDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxW9**](TaxW9.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxW9ByIdGrant

> TaxW9 PostTaxW9ByIdGrant(ctx, id).Execute()

Grants the payer that asked a live read of this org's W-9 — every line, the TIN masked, and the full TIN to that payer's org admins alone, each read audited.



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
	id := "id_example" // string | ID is the relationship's id, \"w9_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxW9ByIdGrant(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxW9ByIdGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxW9ByIdGrant`: TaxW9
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxW9ByIdGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the relationship&#39;s id, \&quot;w9_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxW9ByIdGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxW9**](TaxW9.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxW9ByIdMatch

> TaxW9 PostTaxW9ByIdMatch(ctx, id).TaxMatchIn(taxMatchIn).Execute()

Records the result of IRS TIN Matching for a payee, as the payer.



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
	id := "id_example" // string | ID is the relationship, from the path.
	taxMatchIn := *openapiclient.NewTaxMatchIn() // TaxMatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxW9ByIdMatch(context.Background(), id).TaxMatchIn(taxMatchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxW9ByIdMatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxW9ByIdMatch`: TaxW9
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxW9ByIdMatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the relationship, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxW9ByIdMatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxMatchIn** | [**TaxMatchIn**](TaxMatchIn.md) |  | 

### Return type

[**TaxW9**](TaxW9.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxW9ByIdRevoke

> TaxW9 PostTaxW9ByIdRevoke(ctx, id).Execute()

Revokes a W-9 this org granted.



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
	id := "id_example" // string | ID is the relationship's id, \"w9_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PostTaxW9ByIdRevoke(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PostTaxW9ByIdRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaxW9ByIdRevoke`: TaxW9
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PostTaxW9ByIdRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the relationship&#39;s id, \&quot;w9_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaxW9ByIdRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxW9**](TaxW9.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTaxProfile

> TaxProfile PutTaxProfile(ctx).TaxProfileIn(taxProfileIn).Execute()

Writes the caller org's tax profile — its Form W-9, or its Form W-8BEN or W-8BEN-E when it is a foreign person — and answers it, every number masked.



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
	taxProfileIn := *openapiclient.NewTaxProfileIn() // TaxProfileIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.PutTaxProfile(context.Background()).TaxProfileIn(taxProfileIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.PutTaxProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTaxProfile`: TaxProfile
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.PutTaxProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutTaxProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxProfileIn** | [**TaxProfileIn**](TaxProfileIn.md) |  | 

### Return type

[**TaxProfile**](TaxProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

