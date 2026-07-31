# \CompanyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Company**](CompanyAPI.md#CloudGetV1Company) | **Get** /v1/company | Get returns the caller org&#39;s formation and the stages reachable from it, or 404 when the org has not begun one.
[**CloudGetV1CompanyRegister**](CompanyAPI.md#CloudGetV1CompanyRegister) | **Get** /v1/company/register | ListRegister returns the platform&#39;s whole formation register, newest activity first — every org&#39;s formation, not the caller&#39;s.
[**CloudGetV1CompanyRegisterSummary**](CompanyAPI.md#CloudGetV1CompanyRegisterSummary) | **Get** /v1/company/register/summary | SummarizeRegister counts the platform&#39;s formations by stage — the register&#39;s shape in one read, so a queue that is growing is visible as a number rather than inferred by paging the list.
[**CloudGetV1CompanyReview**](CompanyAPI.md#CloudGetV1CompanyReview) | **Get** /v1/company/review | ReviewQueue reports the founders whose KYC is not yet settled, oldest formation first, so the queue drains in the order founders have been waiting.
[**CloudPostV1Company**](CompanyAPI.md#CloudPostV1Company) | **Post** /v1/company | Begin starts the org&#39;s one formation and returns it with the stages reachable from it.
[**CloudPostV1CompanyAdvance**](CompanyAPI.md#CloudPostV1CompanyAdvance) | **Post** /v1/company/advance | Advance runs the ONE guarded transition of the formation machine.
[**CloudPostV1CompanyDocuments**](CompanyAPI.md#CloudPostV1CompanyDocuments) | **Post** /v1/company/documents | GenerateDocuments renders the formation documents for the chosen structure and jurisdiction, ingests each into the org&#39;s data room, and submits the state filing through the filing seam.
[**CloudPostV1CompanyEsign**](CompanyAPI.md#CloudPostV1CompanyEsign) | **Post** /v1/company/esign | RequestEsign sends the generated formation documents for signature by every founder and records the provider&#39;s reference on the formation.
[**CloudPostV1CompanyEsignComplete**](CompanyAPI.md#CloudPostV1CompanyEsignComplete) | **Post** /v1/company/esign/complete | CompleteEsign records whether the formation documents have been signed.
[**CloudPostV1CompanyFounders**](CompanyAPI.md#CloudPostV1CompanyFounders) | **Post** /v1/company/founders | SetFounders replaces the formation&#39;s founders.
[**CloudPostV1CompanyFundraiseDeck**](CompanyAPI.md#CloudPostV1CompanyFundraiseDeck) | **Post** /v1/company/fundraise/deck | 
[**CloudPostV1CompanyFundraiseRound**](CompanyAPI.md#CloudPostV1CompanyFundraiseRound) | **Post** /v1/company/fundraise/round | RecordRound records a fundraising round on the org&#39;s canonical cap table.
[**CloudPostV1CompanyFundraiseSafe**](CompanyAPI.md#CloudPostV1CompanyFundraiseSafe) | **Post** /v1/company/fundraise/safe | RequestSafe raises an e-signature request over documents already in the org&#39;s data room — a SAFE, a convertible note, or any other fundraising paper.
[**CloudPostV1CompanyGenesis**](CompanyAPI.md#CloudPostV1CompanyGenesis) | **Post** /v1/company/genesis | RecordGenesis seeds the canonical cap table with the founding allocation (stakeholders, a common share class, issued shares) and anchors the deterministic equity-genesis root on-chain.
[**CloudPostV1CompanyImportCaptable**](CompanyAPI.md#CloudPostV1CompanyImportCaptable) | **Post** /v1/company/import/captable | ImportCapTable reads an existing company&#39;s cap table from a Google Sheet and adds its stakeholders to the canonical cap table.
[**CloudPostV1CompanyImportDocuments**](CompanyAPI.md#CloudPostV1CompanyImportDocuments) | **Post** /v1/company/import/documents | ImportDocuments ingests an existing company&#39;s corporate documents from a Google Drive folder into the org&#39;s data room.
[**CloudPostV1CompanyKyc**](CompanyAPI.md#CloudPostV1CompanyKyc) | **Post** /v1/company/kyc | StartKYC opens an identity-verification session for every founder with the wired provider and records each session&#39;s reference on the formation.
[**CloudPostV1CompanyKycDecision**](CompanyAPI.md#CloudPostV1CompanyKycDecision) | **Post** /v1/company/kyc/decision | DecideKYC records a privileged reviewer&#39;s MANUAL decision on a founder&#39;s KYC — the human-in-the-loop path, and the ONLY route to a pass when no real provider is wired.
[**CloudPostV1CompanyKycRefresh**](CompanyAPI.md#CloudPostV1CompanyKycRefresh) | **Post** /v1/company/kyc/refresh | RefreshKYC reconciles each pending founder&#39;s KYC with the WIRED provider — the PULL path to a provider-reported terminal status.
[**CloudPostV1CompanyPayment**](CompanyAPI.md#CloudPostV1CompanyPayment) | **Post** /v1/company/payment | 
[**CloudPostV1CompanySkip**](CompanyAPI.md#CloudPostV1CompanySkip) | **Post** /v1/company/skip | Skip marks the org as already incorporated and moves it onto the import path, so an existing company brings its documents and cap table in instead of forming a new entity.
[**CloudPutV1CompanyStructure**](CompanyAPI.md#CloudPutV1CompanyStructure) | **Put** /v1/company/structure | SetStructure records the entity kind, the state of formation and the proposed name.



## CloudGetV1Company

> CloudFormationView CloudGetV1Company(ctx).Execute()

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
	resp, r, err := apiClient.CompanyAPI.CloudGetV1Company(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudGetV1Company``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Company`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudGetV1Company`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CompanyRequest struct via the builder pattern


### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CompanyRegister

> CloudRegisterPage CloudGetV1CompanyRegister(ctx).Stage(stage).Structure(structure).Limit(limit).Offset(offset).Execute()

ListRegister returns the platform's whole formation register, newest activity first — every org's formation, not the caller's.



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
	resp, r, err := apiClient.CompanyAPI.CloudGetV1CompanyRegister(context.Background()).Stage(stage).Structure(structure).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudGetV1CompanyRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CompanyRegister`: CloudRegisterPage
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudGetV1CompanyRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CompanyRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stage** | **string** | Stage keeps only formations at that stage. Empty means any. | 
 **structure** | **string** | Structure keeps only formations of that entity kind. Empty means any. | 
 **limit** | **int32** | Limit bounds the page; 0 or less means the default of 200. | 
 **offset** | **int32** | Offset skips that many rows. | 

### Return type

[**CloudRegisterPage**](CloudRegisterPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CompanyRegisterSummary

> CloudRegisterCounts CloudGetV1CompanyRegisterSummary(ctx).Execute()

SummarizeRegister counts the platform's formations by stage — the register's shape in one read, so a queue that is growing is visible as a number rather than inferred by paging the list.



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
	resp, r, err := apiClient.CompanyAPI.CloudGetV1CompanyRegisterSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudGetV1CompanyRegisterSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CompanyRegisterSummary`: CloudRegisterCounts
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudGetV1CompanyRegisterSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CompanyRegisterSummaryRequest struct via the builder pattern


### Return type

[**CloudRegisterCounts**](CloudRegisterCounts.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CompanyReview

> CloudReviewQueue CloudGetV1CompanyReview(ctx).Limit(limit).Execute()

ReviewQueue reports the founders whose KYC is not yet settled, oldest formation first, so the queue drains in the order founders have been waiting.



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
	resp, r, err := apiClient.CompanyAPI.CloudGetV1CompanyReview(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudGetV1CompanyReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CompanyReview`: CloudReviewQueue
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudGetV1CompanyReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CompanyReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds how many formations are scanned; 0 or less means the default of 200. | 

### Return type

[**CloudReviewQueue**](CloudReviewQueue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Company

> CloudFormationView CloudPostV1Company(ctx).CloudBeginIn(cloudBeginIn).Execute()

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
	cloudBeginIn := *openapiclient.NewCloudBeginIn() // CloudBeginIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1Company(context.Background()).CloudBeginIn(cloudBeginIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1Company``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Company`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1Company`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudBeginIn** | [**CloudBeginIn**](CloudBeginIn.md) |  | 

### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyAdvance

> CloudFormationView CloudPostV1CompanyAdvance(ctx).CloudAdvanceIn(cloudAdvanceIn).Execute()

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
	cloudAdvanceIn := *openapiclient.NewCloudAdvanceIn() // CloudAdvanceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyAdvance(context.Background()).CloudAdvanceIn(cloudAdvanceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyAdvance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyAdvance`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyAdvance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyAdvanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAdvanceIn** | [**CloudAdvanceIn**](CloudAdvanceIn.md) |  | 

### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyDocuments

> CloudFormationView CloudPostV1CompanyDocuments(ctx).Execute()

GenerateDocuments renders the formation documents for the chosen structure and jurisdiction, ingests each into the org's data room, and submits the state filing through the filing seam.



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
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyDocuments`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyDocuments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyDocumentsRequest struct via the builder pattern


### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyEsign

> CloudEsignOut CloudPostV1CompanyEsign(ctx).Execute()

RequestEsign sends the generated formation documents for signature by every founder and records the provider's reference on the formation.



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
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyEsign(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyEsign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyEsign`: CloudEsignOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyEsign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyEsignRequest struct via the builder pattern


### Return type

[**CloudEsignOut**](CloudEsignOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyEsignComplete

> CloudFormationView CloudPostV1CompanyEsignComplete(ctx).CloudEsignCompleteIn(cloudEsignCompleteIn).Execute()

CompleteEsign records whether the formation documents have been signed.



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
	cloudEsignCompleteIn := *openapiclient.NewCloudEsignCompleteIn() // CloudEsignCompleteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyEsignComplete(context.Background()).CloudEsignCompleteIn(cloudEsignCompleteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyEsignComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyEsignComplete`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyEsignComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyEsignCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudEsignCompleteIn** | [**CloudEsignCompleteIn**](CloudEsignCompleteIn.md) |  | 

### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyFounders

> CloudFormationView CloudPostV1CompanyFounders(ctx).CloudFoundersIn(cloudFoundersIn).Execute()

SetFounders replaces the formation's founders.



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
	cloudFoundersIn := *openapiclient.NewCloudFoundersIn() // CloudFoundersIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyFounders(context.Background()).CloudFoundersIn(cloudFoundersIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyFounders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyFounders`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyFounders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyFoundersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudFoundersIn** | [**CloudFoundersIn**](CloudFoundersIn.md) |  | 

### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyFundraiseDeck

> CloudDeckOut CloudPostV1CompanyFundraiseDeck(ctx).Body(body).Execute()



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
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyFundraiseDeck(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyFundraiseDeck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyFundraiseDeck`: CloudDeckOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyFundraiseDeck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyFundraiseDeckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**CloudDeckOut**](CloudDeckOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyFundraiseRound

> CloudRoundOut CloudPostV1CompanyFundraiseRound(ctx).CloudRoundInput(cloudRoundInput).Execute()

RecordRound records a fundraising round on the org's canonical cap table.



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
	cloudRoundInput := *openapiclient.NewCloudRoundInput() // CloudRoundInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyFundraiseRound(context.Background()).CloudRoundInput(cloudRoundInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyFundraiseRound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyFundraiseRound`: CloudRoundOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyFundraiseRound`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyFundraiseRoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRoundInput** | [**CloudRoundInput**](CloudRoundInput.md) |  | 

### Return type

[**CloudRoundOut**](CloudRoundOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyFundraiseSafe

> CloudSafeOut CloudPostV1CompanyFundraiseSafe(ctx).CloudSafeIn(cloudSafeIn).Execute()

RequestSafe raises an e-signature request over documents already in the org's data room — a SAFE, a convertible note, or any other fundraising paper.



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
	cloudSafeIn := *openapiclient.NewCloudSafeIn() // CloudSafeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyFundraiseSafe(context.Background()).CloudSafeIn(cloudSafeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyFundraiseSafe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyFundraiseSafe`: CloudSafeOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyFundraiseSafe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyFundraiseSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSafeIn** | [**CloudSafeIn**](CloudSafeIn.md) |  | 

### Return type

[**CloudSafeOut**](CloudSafeOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyGenesis

> CloudFormationView CloudPostV1CompanyGenesis(ctx).Execute()

RecordGenesis seeds the canonical cap table with the founding allocation (stakeholders, a common share class, issued shares) and anchors the deterministic equity-genesis root on-chain.



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
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyGenesis(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyGenesis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyGenesis`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyGenesis`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyGenesisRequest struct via the builder pattern


### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyImportCaptable

> CloudImportCapTableOut CloudPostV1CompanyImportCaptable(ctx).CloudImportCapTableIn(cloudImportCapTableIn).Execute()

ImportCapTable reads an existing company's cap table from a Google Sheet and adds its stakeholders to the canonical cap table.



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
	cloudImportCapTableIn := *openapiclient.NewCloudImportCapTableIn() // CloudImportCapTableIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyImportCaptable(context.Background()).CloudImportCapTableIn(cloudImportCapTableIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyImportCaptable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyImportCaptable`: CloudImportCapTableOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyImportCaptable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyImportCaptableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudImportCapTableIn** | [**CloudImportCapTableIn**](CloudImportCapTableIn.md) |  | 

### Return type

[**CloudImportCapTableOut**](CloudImportCapTableOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyImportDocuments

> CloudImportDocumentsOut CloudPostV1CompanyImportDocuments(ctx).CloudImportDocumentsIn(cloudImportDocumentsIn).Execute()

ImportDocuments ingests an existing company's corporate documents from a Google Drive folder into the org's data room.



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
	cloudImportDocumentsIn := *openapiclient.NewCloudImportDocumentsIn() // CloudImportDocumentsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyImportDocuments(context.Background()).CloudImportDocumentsIn(cloudImportDocumentsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyImportDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyImportDocuments`: CloudImportDocumentsOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyImportDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyImportDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudImportDocumentsIn** | [**CloudImportDocumentsIn**](CloudImportDocumentsIn.md) |  | 

### Return type

[**CloudImportDocumentsOut**](CloudImportDocumentsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyKyc

> CloudKycStartOut CloudPostV1CompanyKyc(ctx).Execute()

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
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyKyc(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyKyc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyKyc`: CloudKycStartOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyKyc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyKycRequest struct via the builder pattern


### Return type

[**CloudKycStartOut**](CloudKycStartOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyKycDecision

> CloudFormationView CloudPostV1CompanyKycDecision(ctx).CloudDecisionIn(cloudDecisionIn).Execute()

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
	cloudDecisionIn := *openapiclient.NewCloudDecisionIn() // CloudDecisionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyKycDecision(context.Background()).CloudDecisionIn(cloudDecisionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyKycDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyKycDecision`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyKycDecision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyKycDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudDecisionIn** | [**CloudDecisionIn**](CloudDecisionIn.md) |  | 

### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyKycRefresh

> CloudKycRefreshOut CloudPostV1CompanyKycRefresh(ctx).Execute()

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
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyKycRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyKycRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyKycRefresh`: CloudKycRefreshOut
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyKycRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyKycRefreshRequest struct via the builder pattern


### Return type

[**CloudKycRefreshOut**](CloudKycRefreshOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanyPayment

> CloudFormationView CloudPostV1CompanyPayment(ctx).Execute()



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
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanyPayment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanyPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanyPayment`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanyPayment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanyPaymentRequest struct via the builder pattern


### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CompanySkip

> CloudFormationView CloudPostV1CompanySkip(ctx).Execute()

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
	resp, r, err := apiClient.CompanyAPI.CloudPostV1CompanySkip(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPostV1CompanySkip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CompanySkip`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPostV1CompanySkip`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CompanySkipRequest struct via the builder pattern


### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1CompanyStructure

> CloudFormationView CloudPutV1CompanyStructure(ctx).CloudStructureIn(cloudStructureIn).Execute()

SetStructure records the entity kind, the state of formation and the proposed name.



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
	cloudStructureIn := *openapiclient.NewCloudStructureIn() // CloudStructureIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CloudPutV1CompanyStructure(context.Background()).CloudStructureIn(cloudStructureIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CloudPutV1CompanyStructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1CompanyStructure`: CloudFormationView
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CloudPutV1CompanyStructure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CompanyStructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudStructureIn** | [**CloudStructureIn**](CloudStructureIn.md) |  | 

### Return type

[**CloudFormationView**](CloudFormationView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

