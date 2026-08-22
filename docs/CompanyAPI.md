# \CompanyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompany**](CompanyAPI.md#GetCompany) | **Get** /v1/company | Get returns the caller org&#39;s formation and the stages reachable from it, or 404 when the org has not begun one.
[**GetCompanyRegister**](CompanyAPI.md#GetCompanyRegister) | **Get** /v1/company/register | Returns the platform&#39;s whole formation register, newest activity first — every org&#39;s formation, not the caller&#39;s.
[**GetCompanyRegisterSummary**](CompanyAPI.md#GetCompanyRegisterSummary) | **Get** /v1/company/register/summary | Counts the platform&#39;s formations by stage — the register&#39;s shape in one read, so a queue that is growing is visible as a number rather than inferred by paging the list.
[**GetCompanyReview**](CompanyAPI.md#GetCompanyReview) | **Get** /v1/company/review | Reports the founders whose KYC is not yet settled, oldest formation first, so the queue drains in the order founders have been waiting.
[**PostCompany**](CompanyAPI.md#PostCompany) | **Post** /v1/company | Begin starts the org&#39;s one formation and returns it with the stages reachable from it.
[**PostCompanyAdvance**](CompanyAPI.md#PostCompanyAdvance) | **Post** /v1/company/advance | Advance runs the ONE guarded transition of the formation machine.
[**PostCompanyDocuments**](CompanyAPI.md#PostCompanyDocuments) | **Post** /v1/company/documents | Renders the formation documents for the chosen structure and jurisdiction, ingests each into the org&#39;s data room, and submits the state filing through the filing seam.
[**PostCompanyEin**](CompanyAPI.md#PostCompanyEin) | **Post** /v1/company/ein | Opens the EIN application and answers what it owes.
[**PostCompanyEsign**](CompanyAPI.md#PostCompanyEsign) | **Post** /v1/company/esign | Sends the generated formation documents for signature by every founder and records the provider&#39;s reference on the formation.
[**PostCompanyEsignComplete**](CompanyAPI.md#PostCompanyEsignComplete) | **Post** /v1/company/esign/complete | Records whether the formation documents have been signed.
[**PostCompanyFounders**](CompanyAPI.md#PostCompanyFounders) | **Post** /v1/company/founders | Replaces the formation&#39;s founders.
[**PostCompanyFundraiseDeck**](CompanyAPI.md#PostCompanyFundraiseDeck) | **Post** /v1/company/fundraise/deck | Share a pitch deck in the org&#39;s data room
[**PostCompanyFundraiseRound**](CompanyAPI.md#PostCompanyFundraiseRound) | **Post** /v1/company/fundraise/round | Records a fundraising round on the org&#39;s canonical cap table.
[**PostCompanyFundraiseSafe**](CompanyAPI.md#PostCompanyFundraiseSafe) | **Post** /v1/company/fundraise/safe | Raises an e-signature request over documents already in the org&#39;s data room — a SAFE, a convertible note, or any other fundraising paper.
[**PostCompanyGenesis**](CompanyAPI.md#PostCompanyGenesis) | **Post** /v1/company/genesis | Seeds the canonical cap table with the founding allocation (stakeholders, a common share class, issued shares) and anchors the deterministic equity-genesis root on-chain.
[**PostCompanyImportCaptable**](CompanyAPI.md#PostCompanyImportCaptable) | **Post** /v1/company/import/captable | Reads an existing company&#39;s cap table from a Google Sheet and adds its stakeholders to the canonical cap table.
[**PostCompanyImportDocuments**](CompanyAPI.md#PostCompanyImportDocuments) | **Post** /v1/company/import/documents | Ingests an existing company&#39;s corporate documents from a Google Drive folder into the org&#39;s data room.
[**PostCompanyKyc**](CompanyAPI.md#PostCompanyKyc) | **Post** /v1/company/kyc | StartKYC opens an identity-verification session for every founder with the wired provider and records each session&#39;s reference on the formation.
[**PostCompanyKycDecision**](CompanyAPI.md#PostCompanyKycDecision) | **Post** /v1/company/kyc/decision | DecideKYC records a privileged reviewer&#39;s MANUAL decision on a founder&#39;s KYC — the human-in-the-loop path, and the ONLY route to a pass when no real provider is wired.
[**PostCompanyKycRefresh**](CompanyAPI.md#PostCompanyKycRefresh) | **Post** /v1/company/kyc/refresh | RefreshKYC reconciles each pending founder&#39;s KYC with the WIRED provider — the PULL path to a provider-reported terminal status.
[**PostCompanyPayment**](CompanyAPI.md#PostCompanyPayment) | **Post** /v1/company/payment | Charges the caller&#39;s own org the one-time Hanzo Company formation fee.
[**PostCompanySkip**](CompanyAPI.md#PostCompanySkip) | **Post** /v1/company/skip | Skip marks the org as already incorporated and moves it onto the import path, so an existing company brings its documents and cap table in instead of forming a new entity.
[**PostCompanyTariff**](CompanyAPI.md#PostCompanyTariff) | **Post** /v1/company/tariff | Itemises what a formation costs before anyone commits to it.
[**PutCompanyStructure**](CompanyAPI.md#PutCompanyStructure) | **Put** /v1/company/structure | Records the entity kind, the state of formation and the proposed name.



## GetCompany

> FormationView GetCompany(ctx).Execute()

Get returns the caller org's formation and the stages reachable from it, or 404 when the org has not begun one.



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
	resp, r, err := apiClient.CompanyAPI.GetCompany(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.GetCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompany`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.GetCompany`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyRegister

> RegisterPage GetCompanyRegister(ctx).Stage(stage).Structure(structure).Limit(limit).Offset(offset).Execute()

Returns the platform's whole formation register, newest activity first — every org's formation, not the caller's.



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
	stage := "founders" // string | Stage keeps only formations at that stage. Empty means any. (optional)
	structure := "structure_example" // string | Structure keeps only formations of that entity kind. Empty means any. (optional)
	limit := int32(50) // int32 | Limit bounds the page; 0 or less means the default of 200. (optional)
	offset := int32(56) // int32 | Offset skips that many rows. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.GetCompanyRegister(context.Background()).Stage(stage).Structure(structure).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.GetCompanyRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyRegister`: RegisterPage
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.GetCompanyRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stage** | **string** | Stage keeps only formations at that stage. Empty means any. | 
 **structure** | **string** | Structure keeps only formations of that entity kind. Empty means any. | 
 **limit** | **int32** | Limit bounds the page; 0 or less means the default of 200. | 
 **offset** | **int32** | Offset skips that many rows. | 

### Return type

[**RegisterPage**](RegisterPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyRegisterSummary

> RegisterCounts GetCompanyRegisterSummary(ctx).Execute()

Counts the platform's formations by stage — the register's shape in one read, so a queue that is growing is visible as a number rather than inferred by paging the list.



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
	resp, r, err := apiClient.CompanyAPI.GetCompanyRegisterSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.GetCompanyRegisterSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyRegisterSummary`: RegisterCounts
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.GetCompanyRegisterSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRegisterSummaryRequest struct via the builder pattern


### Return type

[**RegisterCounts**](RegisterCounts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyReview

> ReviewQueue GetCompanyReview(ctx).Limit(limit).Execute()

Reports the founders whose KYC is not yet settled, oldest formation first, so the queue drains in the order founders have been waiting.



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
	limit := int32(50) // int32 | Limit bounds how many formations are scanned; 0 or less means the default of 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.GetCompanyReview(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.GetCompanyReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyReview`: ReviewQueue
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.GetCompanyReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds how many formations are scanned; 0 or less means the default of 200. | 

### Return type

[**ReviewQueue**](ReviewQueue.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompany

> FormationView PostCompany(ctx).BeginIn(beginIn).Execute()

Begin starts the org's one formation and returns it with the stages reachable from it.



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
	beginIn := *openapiclient.NewBeginIn() // BeginIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompany(context.Background()).BeginIn(beginIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompany`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beginIn** | [**BeginIn**](BeginIn.md) |  | 

### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyAdvance

> FormationView PostCompanyAdvance(ctx).AdvanceIn(advanceIn).Execute()

Advance runs the ONE guarded transition of the formation machine.



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
	advanceIn := *openapiclient.NewAdvanceIn() // AdvanceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyAdvance(context.Background()).AdvanceIn(advanceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyAdvance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyAdvance`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyAdvance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyAdvanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **advanceIn** | [**AdvanceIn**](AdvanceIn.md) |  | 

### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyDocuments

> FormationView PostCompanyDocuments(ctx).Execute()

Renders the formation documents for the chosen structure and jurisdiction, ingests each into the org's data room, and submits the state filing through the filing seam.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyDocuments`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyDocuments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyDocumentsRequest struct via the builder pattern


### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyEin

> EIN PostCompanyEin(ctx).EinIn(einIn).Execute()

Opens the EIN application and answers what it owes.



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
	einIn := *openapiclient.NewEinIn() // EinIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyEin(context.Background()).EinIn(einIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyEin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyEin`: EIN
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyEin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyEinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **einIn** | [**EinIn**](EinIn.md) |  | 

### Return type

[**EIN**](EIN.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyEsign

> EsignOut PostCompanyEsign(ctx).Execute()

Sends the generated formation documents for signature by every founder and records the provider's reference on the formation.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyEsign(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyEsign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyEsign`: EsignOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyEsign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyEsignRequest struct via the builder pattern


### Return type

[**EsignOut**](EsignOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyEsignComplete

> FormationView PostCompanyEsignComplete(ctx).EsignCompleteIn(esignCompleteIn).Execute()

Records whether the formation documents have been signed.



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
	esignCompleteIn := *openapiclient.NewEsignCompleteIn() // EsignCompleteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyEsignComplete(context.Background()).EsignCompleteIn(esignCompleteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyEsignComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyEsignComplete`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyEsignComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyEsignCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esignCompleteIn** | [**EsignCompleteIn**](EsignCompleteIn.md) |  | 

### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyFounders

> FormationView PostCompanyFounders(ctx).FoundersIn(foundersIn).Execute()

Replaces the formation's founders.



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
	foundersIn := *openapiclient.NewFoundersIn() // FoundersIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyFounders(context.Background()).FoundersIn(foundersIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyFounders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyFounders`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyFounders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyFoundersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foundersIn** | [**FoundersIn**](FoundersIn.md) |  | 

### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyFundraiseDeck

> DeckOut PostCompanyFundraiseDeck(ctx).Body(body).Execute()

Share a pitch deck in the org's data room



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyFundraiseDeck(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyFundraiseDeck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyFundraiseDeck`: DeckOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyFundraiseDeck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyFundraiseDeckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**DeckOut**](DeckOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyFundraiseRound

> RoundOut PostCompanyFundraiseRound(ctx).RoundInput(roundInput).Execute()

Records a fundraising round on the org's canonical cap table.



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
	roundInput := *openapiclient.NewRoundInput() // RoundInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyFundraiseRound(context.Background()).RoundInput(roundInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyFundraiseRound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyFundraiseRound`: RoundOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyFundraiseRound`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyFundraiseRoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roundInput** | [**RoundInput**](RoundInput.md) |  | 

### Return type

[**RoundOut**](RoundOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyFundraiseSafe

> SafeOut PostCompanyFundraiseSafe(ctx).SafeIn(safeIn).Execute()

Raises an e-signature request over documents already in the org's data room — a SAFE, a convertible note, or any other fundraising paper.



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
	safeIn := *openapiclient.NewSafeIn() // SafeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyFundraiseSafe(context.Background()).SafeIn(safeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyFundraiseSafe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyFundraiseSafe`: SafeOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyFundraiseSafe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyFundraiseSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **safeIn** | [**SafeIn**](SafeIn.md) |  | 

### Return type

[**SafeOut**](SafeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyGenesis

> FormationView PostCompanyGenesis(ctx).Execute()

Seeds the canonical cap table with the founding allocation (stakeholders, a common share class, issued shares) and anchors the deterministic equity-genesis root on-chain.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyGenesis(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyGenesis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyGenesis`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyGenesis`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyGenesisRequest struct via the builder pattern


### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyImportCaptable

> ImportCapTableOut PostCompanyImportCaptable(ctx).ImportCapTableIn(importCapTableIn).Execute()

Reads an existing company's cap table from a Google Sheet and adds its stakeholders to the canonical cap table.



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
	importCapTableIn := *openapiclient.NewImportCapTableIn() // ImportCapTableIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyImportCaptable(context.Background()).ImportCapTableIn(importCapTableIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyImportCaptable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyImportCaptable`: ImportCapTableOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyImportCaptable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyImportCaptableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importCapTableIn** | [**ImportCapTableIn**](ImportCapTableIn.md) |  | 

### Return type

[**ImportCapTableOut**](ImportCapTableOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyImportDocuments

> ImportDocumentsOut PostCompanyImportDocuments(ctx).ImportDocumentsIn(importDocumentsIn).Execute()

Ingests an existing company's corporate documents from a Google Drive folder into the org's data room.



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
	importDocumentsIn := *openapiclient.NewImportDocumentsIn() // ImportDocumentsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyImportDocuments(context.Background()).ImportDocumentsIn(importDocumentsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyImportDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyImportDocuments`: ImportDocumentsOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyImportDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyImportDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importDocumentsIn** | [**ImportDocumentsIn**](ImportDocumentsIn.md) |  | 

### Return type

[**ImportDocumentsOut**](ImportDocumentsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyKyc

> KycStartOut PostCompanyKyc(ctx).Execute()

StartKYC opens an identity-verification session for every founder with the wired provider and records each session's reference on the formation.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyKyc(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyKyc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyKyc`: KycStartOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyKyc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyKycRequest struct via the builder pattern


### Return type

[**KycStartOut**](KycStartOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyKycDecision

> FormationView PostCompanyKycDecision(ctx).DecisionIn(decisionIn).Execute()

DecideKYC records a privileged reviewer's MANUAL decision on a founder's KYC — the human-in-the-loop path, and the ONLY route to a pass when no real provider is wired.



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
	decisionIn := *openapiclient.NewDecisionIn() // DecisionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyKycDecision(context.Background()).DecisionIn(decisionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyKycDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyKycDecision`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyKycDecision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyKycDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionIn** | [**DecisionIn**](DecisionIn.md) |  | 

### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyKycRefresh

> KycRefreshOut PostCompanyKycRefresh(ctx).Execute()

RefreshKYC reconciles each pending founder's KYC with the WIRED provider — the PULL path to a provider-reported terminal status.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyKycRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyKycRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyKycRefresh`: KycRefreshOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyKycRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyKycRefreshRequest struct via the builder pattern


### Return type

[**KycRefreshOut**](KycRefreshOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyPayment

> FormationView PostCompanyPayment(ctx).Execute()

Charges the caller's own org the one-time Hanzo Company formation fee.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyPayment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyPayment`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyPayment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyPaymentRequest struct via the builder pattern


### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanySkip

> FormationView PostCompanySkip(ctx).Execute()

Skip marks the org as already incorporated and moves it onto the import path, so an existing company brings its documents and cap table in instead of forming a new entity.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanySkip(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanySkip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanySkip`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanySkip`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanySkipRequest struct via the builder pattern


### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyTariff

> Tariff PostCompanyTariff(ctx).TariffIn(tariffIn).Execute()

Itemises what a formation costs before anyone commits to it.



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
	tariffIn := *openapiclient.NewTariffIn() // TariffIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyTariff(context.Background()).TariffIn(tariffIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyTariff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyTariff`: Tariff
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyTariff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyTariffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tariffIn** | [**TariffIn**](TariffIn.md) |  | 

### Return type

[**Tariff**](Tariff.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyStructure

> FormationView PutCompanyStructure(ctx).StructureIn(structureIn).Execute()

Records the entity kind, the state of formation and the proposed name.



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
	structureIn := *openapiclient.NewStructureIn() // StructureIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PutCompanyStructure(context.Background()).StructureIn(structureIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PutCompanyStructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyStructure`: FormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PutCompanyStructure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyStructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **structureIn** | [**StructureIn**](StructureIn.md) |  | 

### Return type

[**FormationView**](FormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

