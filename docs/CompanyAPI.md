# \CompanyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompany**](CompanyAPI.md#GetCompany) | **Get** /v1/company | Returns the caller org&#39;s formation and the stages reachable from it, or 404 when the org has not begun one.
[**GetCompanyRegister**](CompanyAPI.md#GetCompanyRegister) | **Get** /v1/company/register | Returns the platform&#39;s whole formation register, newest activity first — every org&#39;s formation, not the caller&#39;s.
[**GetCompanyRegisterSummary**](CompanyAPI.md#GetCompanyRegisterSummary) | **Get** /v1/company/register/summary | Counts the platform&#39;s formations by stage — the register&#39;s shape in one read, so a queue that is growing is visible as a number rather than inferred by paging the list.
[**GetCompanyReview**](CompanyAPI.md#GetCompanyReview) | **Get** /v1/company/review | Reports the founders whose KYC is not yet settled, oldest formation first, so the queue drains in the order founders have been waiting.
[**PostCompany**](CompanyAPI.md#PostCompany) | **Post** /v1/company | Starts the org&#39;s one formation and returns it with the stages reachable from it.
[**PostCompanyAdvance**](CompanyAPI.md#PostCompanyAdvance) | **Post** /v1/company/advance | Runs the ONE guarded transition of the formation machine.
[**PostCompanyDocuments**](CompanyAPI.md#PostCompanyDocuments) | **Post** /v1/company/documents | Renders the formation documents for the chosen structure and jurisdiction, ingests each into the org&#39;s data room, and submits the state filing through the filing client.
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
[**PostCompanyKyc**](CompanyAPI.md#PostCompanyKyc) | **Post** /v1/company/kyc | Opens an identity-verification session for every founder with the wired provider and records each session&#39;s reference on the formation.
[**PostCompanyKycRefresh**](CompanyAPI.md#PostCompanyKycRefresh) | **Post** /v1/company/kyc/refresh | RefreshKYC reconciles each pending founder&#39;s KYC with the WIRED provider — the PULL path to a provider-reported terminal status.
[**PostCompanyPayment**](CompanyAPI.md#PostCompanyPayment) | **Post** /v1/company/payment | Charges the caller&#39;s own org the one-time Hanzo Company formation fee.
[**PostCompanySkip**](CompanyAPI.md#PostCompanySkip) | **Post** /v1/company/skip | Marks the org as already incorporated and moves it onto the import path, so an existing company brings its documents and cap table in instead of forming a new entity.
[**PostCompanyTariff**](CompanyAPI.md#PostCompanyTariff) | **Post** /v1/company/tariff | Itemises what a formation costs before anyone commits to it.
[**PutCompanyStructure**](CompanyAPI.md#PutCompanyStructure) | **Put** /v1/company/structure | Records the entity kind, the state of formation and the proposed name.



## GetCompany

> CompanyFormationView GetCompany(ctx).Execute()

Returns the caller org's formation and the stages reachable from it, or 404 when the org has not begun one.



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
	resp, r, err := apiClient.CompanyAPI.GetCompany(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.GetCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompany`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.GetCompany`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyRegister

> CompanyRegisterPage GetCompanyRegister(ctx).Stage(stage).Structure(structure).Limit(limit).Offset(offset).Execute()

Returns the platform's whole formation register, newest activity first — every org's formation, not the caller's.



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
	stage := "founders" // string | Stage keeps only formations at that stage. Empty means any. (optional)
	structure := "structure_example" // string | Structure keeps only formations of that entity kind. Empty means any. (optional)
	limit := int64(50) // int64 | Limit bounds the page; 0 or less means the default of 200. (optional)
	offset := int64(789) // int64 | Offset skips that many rows. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.GetCompanyRegister(context.Background()).Stage(stage).Structure(structure).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.GetCompanyRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyRegister`: CompanyRegisterPage
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
 **limit** | **int64** | Limit bounds the page; 0 or less means the default of 200. | 
 **offset** | **int64** | Offset skips that many rows. | 

### Return type

[**CompanyRegisterPage**](CompanyRegisterPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyRegisterSummary

> CompanyRegisterCounts GetCompanyRegisterSummary(ctx).Execute()

Counts the platform's formations by stage — the register's shape in one read, so a queue that is growing is visible as a number rather than inferred by paging the list.



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
	resp, r, err := apiClient.CompanyAPI.GetCompanyRegisterSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.GetCompanyRegisterSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyRegisterSummary`: CompanyRegisterCounts
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.GetCompanyRegisterSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRegisterSummaryRequest struct via the builder pattern


### Return type

[**CompanyRegisterCounts**](CompanyRegisterCounts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyReview

> CompanyReviewQueue GetCompanyReview(ctx).Limit(limit).Execute()

Reports the founders whose KYC is not yet settled, oldest formation first, so the queue drains in the order founders have been waiting.



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
	limit := int64(50) // int64 | Limit bounds how many formations are scanned; 0 or less means the default of 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.GetCompanyReview(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.GetCompanyReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyReview`: CompanyReviewQueue
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.GetCompanyReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit bounds how many formations are scanned; 0 or less means the default of 200. | 

### Return type

[**CompanyReviewQueue**](CompanyReviewQueue.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompany

> CompanyFormationView PostCompany(ctx).CompanyBeginIn(companyBeginIn).Execute()

Starts the org's one formation and returns it with the stages reachable from it.



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
	companyBeginIn := *openapiclient.NewCompanyBeginIn() // CompanyBeginIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompany(context.Background()).CompanyBeginIn(companyBeginIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompany`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyBeginIn** | [**CompanyBeginIn**](CompanyBeginIn.md) |  | 

### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyAdvance

> CompanyFormationView PostCompanyAdvance(ctx).CompanyAdvanceIn(companyAdvanceIn).Execute()

Runs the ONE guarded transition of the formation machine.



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
	companyAdvanceIn := *openapiclient.NewCompanyAdvanceIn() // CompanyAdvanceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyAdvance(context.Background()).CompanyAdvanceIn(companyAdvanceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyAdvance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyAdvance`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyAdvance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyAdvanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyAdvanceIn** | [**CompanyAdvanceIn**](CompanyAdvanceIn.md) |  | 

### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyDocuments

> CompanyFormationView PostCompanyDocuments(ctx).Execute()

Renders the formation documents for the chosen structure and jurisdiction, ingests each into the org's data room, and submits the state filing through the filing client.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyDocuments`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyDocuments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyDocumentsRequest struct via the builder pattern


### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyEin

> CompanyEIN PostCompanyEin(ctx).CompanyEinIn(companyEinIn).Execute()

Opens the EIN application and answers what it owes.



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
	companyEinIn := *openapiclient.NewCompanyEinIn() // CompanyEinIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyEin(context.Background()).CompanyEinIn(companyEinIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyEin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyEin`: CompanyEIN
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyEin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyEinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyEinIn** | [**CompanyEinIn**](CompanyEinIn.md) |  | 

### Return type

[**CompanyEIN**](CompanyEIN.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyEsign

> CompanyEsignOut PostCompanyEsign(ctx).Execute()

Sends the generated formation documents for signature by every founder and records the provider's reference on the formation.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyEsign(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyEsign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyEsign`: CompanyEsignOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyEsign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyEsignRequest struct via the builder pattern


### Return type

[**CompanyEsignOut**](CompanyEsignOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyEsignComplete

> CompanyFormationView PostCompanyEsignComplete(ctx).CompanyEsignCompleteIn(companyEsignCompleteIn).Execute()

Records whether the formation documents have been signed.



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
	companyEsignCompleteIn := *openapiclient.NewCompanyEsignCompleteIn() // CompanyEsignCompleteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyEsignComplete(context.Background()).CompanyEsignCompleteIn(companyEsignCompleteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyEsignComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyEsignComplete`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyEsignComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyEsignCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyEsignCompleteIn** | [**CompanyEsignCompleteIn**](CompanyEsignCompleteIn.md) |  | 

### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyFounders

> CompanyFormationView PostCompanyFounders(ctx).CompanyFoundersIn(companyFoundersIn).Execute()

Replaces the formation's founders.



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
	companyFoundersIn := *openapiclient.NewCompanyFoundersIn() // CompanyFoundersIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyFounders(context.Background()).CompanyFoundersIn(companyFoundersIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyFounders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyFounders`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyFounders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyFoundersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyFoundersIn** | [**CompanyFoundersIn**](CompanyFoundersIn.md) |  | 

### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> CompanyRoundOut PostCompanyFundraiseRound(ctx).CompanyRoundInput(companyRoundInput).Execute()

Records a fundraising round on the org's canonical cap table.



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
	companyRoundInput := *openapiclient.NewCompanyRoundInput() // CompanyRoundInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyFundraiseRound(context.Background()).CompanyRoundInput(companyRoundInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyFundraiseRound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyFundraiseRound`: CompanyRoundOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyFundraiseRound`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyFundraiseRoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyRoundInput** | [**CompanyRoundInput**](CompanyRoundInput.md) |  | 

### Return type

[**CompanyRoundOut**](CompanyRoundOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyFundraiseSafe

> CompanySafeOut PostCompanyFundraiseSafe(ctx).CompanySafeIn(companySafeIn).Execute()

Raises an e-signature request over documents already in the org's data room — a SAFE, a convertible note, or any other fundraising paper.



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
	companySafeIn := *openapiclient.NewCompanySafeIn() // CompanySafeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyFundraiseSafe(context.Background()).CompanySafeIn(companySafeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyFundraiseSafe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyFundraiseSafe`: CompanySafeOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyFundraiseSafe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyFundraiseSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companySafeIn** | [**CompanySafeIn**](CompanySafeIn.md) |  | 

### Return type

[**CompanySafeOut**](CompanySafeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyGenesis

> CompanyFormationView PostCompanyGenesis(ctx).Execute()

Seeds the canonical cap table with the founding allocation (stakeholders, a common share class, issued shares) and anchors the deterministic equity-genesis root on-chain.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyGenesis(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyGenesis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyGenesis`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyGenesis`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyGenesisRequest struct via the builder pattern


### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyImportCaptable

> CompanyImportCapTableOut PostCompanyImportCaptable(ctx).CompanyImportCapTableIn(companyImportCapTableIn).Execute()

Reads an existing company's cap table from a Google Sheet and adds its stakeholders to the canonical cap table.



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
	companyImportCapTableIn := *openapiclient.NewCompanyImportCapTableIn() // CompanyImportCapTableIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyImportCaptable(context.Background()).CompanyImportCapTableIn(companyImportCapTableIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyImportCaptable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyImportCaptable`: CompanyImportCapTableOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyImportCaptable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyImportCaptableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyImportCapTableIn** | [**CompanyImportCapTableIn**](CompanyImportCapTableIn.md) |  | 

### Return type

[**CompanyImportCapTableOut**](CompanyImportCapTableOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyImportDocuments

> CompanyImportDocumentsOut PostCompanyImportDocuments(ctx).CompanyImportDocumentsIn(companyImportDocumentsIn).Execute()

Ingests an existing company's corporate documents from a Google Drive folder into the org's data room.



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
	companyImportDocumentsIn := *openapiclient.NewCompanyImportDocumentsIn() // CompanyImportDocumentsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyImportDocuments(context.Background()).CompanyImportDocumentsIn(companyImportDocumentsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyImportDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyImportDocuments`: CompanyImportDocumentsOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyImportDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyImportDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyImportDocumentsIn** | [**CompanyImportDocumentsIn**](CompanyImportDocumentsIn.md) |  | 

### Return type

[**CompanyImportDocumentsOut**](CompanyImportDocumentsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyKyc

> CompanyKycStartOut PostCompanyKyc(ctx).Execute()

Opens an identity-verification session for every founder with the wired provider and records each session's reference on the formation.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyKyc(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyKyc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyKyc`: CompanyKycStartOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyKyc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyKycRequest struct via the builder pattern


### Return type

[**CompanyKycStartOut**](CompanyKycStartOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyKycRefresh

> CompanyKycRefreshOut PostCompanyKycRefresh(ctx).Execute()

RefreshKYC reconciles each pending founder's KYC with the WIRED provider — the PULL path to a provider-reported terminal status.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyKycRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyKycRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyKycRefresh`: CompanyKycRefreshOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyKycRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyKycRefreshRequest struct via the builder pattern


### Return type

[**CompanyKycRefreshOut**](CompanyKycRefreshOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyPayment

> CompanyFormationView PostCompanyPayment(ctx).Execute()

Charges the caller's own org the one-time Hanzo Company formation fee.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanyPayment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyPayment`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyPayment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyPaymentRequest struct via the builder pattern


### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanySkip

> CompanyFormationView PostCompanySkip(ctx).Execute()

Marks the org as already incorporated and moves it onto the import path, so an existing company brings its documents and cap table in instead of forming a new entity.



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
	resp, r, err := apiClient.CompanyAPI.PostCompanySkip(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanySkip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanySkip`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanySkip`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanySkipRequest struct via the builder pattern


### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyTariff

> CompanyTariff PostCompanyTariff(ctx).CompanyTariffIn(companyTariffIn).Execute()

Itemises what a formation costs before anyone commits to it.



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
	companyTariffIn := *openapiclient.NewCompanyTariffIn() // CompanyTariffIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PostCompanyTariff(context.Background()).CompanyTariffIn(companyTariffIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PostCompanyTariff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyTariff`: CompanyTariff
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PostCompanyTariff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyTariffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyTariffIn** | [**CompanyTariffIn**](CompanyTariffIn.md) |  | 

### Return type

[**CompanyTariff**](CompanyTariff.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyStructure

> CompanyFormationView PutCompanyStructure(ctx).CompanyStructureIn(companyStructureIn).Execute()

Records the entity kind, the state of formation and the proposed name.



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
	companyStructureIn := *openapiclient.NewCompanyStructureIn() // CompanyStructureIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.PutCompanyStructure(context.Background()).CompanyStructureIn(companyStructureIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.PutCompanyStructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyStructure`: CompanyFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.PutCompanyStructure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyStructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyStructureIn** | [**CompanyStructureIn**](CompanyStructureIn.md) |  | 

### Return type

[**CompanyFormationView**](CompanyFormationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

