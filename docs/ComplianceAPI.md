# \ComplianceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComplianceAccreditation**](ComplianceAPI.md#GetComplianceAccreditation) | **Get** /v1/compliance/accreditation | Returns the org&#39;s tracked accreditation-state records, newest first — evidence entries the org keeps, never a platform certification.
[**GetComplianceAccreditationById**](ComplianceAPI.md#GetComplianceAccreditationById) | **Get** /v1/compliance/accreditation/{id} | Returns one tracked accreditation record.
[**GetComplianceAudit**](ComplianceAPI.md#GetComplianceAudit) | **Get** /v1/compliance/audit | AuditRead is the compliance read of the SHARED tamper-evident audit plane — the SOC 2 posture surface (privileged actions: who started/decided what, when).
[**GetComplianceHealth**](ComplianceAPI.md#GetComplianceHealth) | **Get** /v1/compliance/health | Reports subsystem liveness and the wired verification provider.
[**GetComplianceRecords**](ComplianceAPI.md#GetComplianceRecords) | **Get** /v1/compliance/records | ListRecords is the unified compliance-record view for the org: its verifications and accreditation records together, each provider-reported or tracked, never platform-asserted.
[**GetComplianceStatus**](ComplianceAPI.md#GetComplianceStatus) | **Get** /v1/compliance/status | Status is the org&#39;s honest posture read: the wired provider and the per-status tally of its verifications.
[**GetComplianceSubjects**](ComplianceAPI.md#GetComplianceSubjects) | **Get** /v1/compliance/subjects | Returns the org&#39;s subjects as PII-MINIMIZED summaries — no name or email, only whether an email is on file.
[**GetComplianceSubjectsById**](ComplianceAPI.md#GetComplianceSubjectsById) | **Get** /v1/compliance/subjects/{id} | Returns one subject WITH its contact PII — the only surface that returns it, and only to the owning org.
[**GetComplianceVerifications**](ComplianceAPI.md#GetComplianceVerifications) | **Get** /v1/compliance/verifications | Returns the org&#39;s KYC/KYB verifications, newest first — opaque subject references and provider-reported statuses only, no subject PII.
[**GetComplianceVerificationsById**](ComplianceAPI.md#GetComplianceVerificationsById) | **Get** /v1/compliance/verifications/{id} | Returns one verification — its opaque subject reference and provider-reported status, no subject PII.
[**PostComplianceAccreditation**](ComplianceAPI.md#PostComplianceAccreditation) | **Post** /v1/compliance/accreditation | Records an ASSERTED accreditation state for a subject — the subject&#39;s own assertion, with no verifier.
[**PostComplianceAccreditationByIdDecision**](ComplianceAPI.md#PostComplianceAccreditationByIdDecision) | **Post** /v1/compliance/accreditation/{id}/decision | Records an org reviewer&#39;s decision on an accreditation record — a reviewer confirmation, a provider verification the reviewer has evidence of (a CPA/attorney letter, a verifier report), a rejection, or an expiry.
[**PostComplianceSubjects**](ComplianceAPI.md#PostComplianceSubjects) | **Post** /v1/compliance/subjects | Records a party the org is verifying as part of its own onboarding/compliance — a team member, vendor, customer, or counterparty.
[**PostComplianceVerifications**](ComplianceAPI.md#PostComplianceVerifications) | **Post** /v1/compliance/verifications | Begins a KYC/KYB verification of a subject through the wired provider — an existing subject by id, or one created inline from the request.
[**PostComplianceVerificationsByIdDecision**](ComplianceAPI.md#PostComplianceVerificationsByIdDecision) | **Post** /v1/compliance/verifications/{id}/decision | Records a privileged reviewer&#39;s MANUAL decision on a verification — the human-in-the-loop path, and the ONLY route to a passing status when no real provider is wired.
[**PostComplianceVerificationsByIdRefresh**](ComplianceAPI.md#PostComplianceVerificationsByIdRefresh) | **Post** /v1/compliance/verifications/{id}/refresh | Polls the wired provider for its current decision and records it, ATTRIBUTED to the provider — the internal PULL reconcile.
[**PostComplianceVerificationsWebhook**](ComplianceAPI.md#PostComplianceVerificationsWebhook) | **Post** /v1/compliance/verifications/webhook | Provider push that settles a verification, authenticated by HMAC signature



## GetComplianceAccreditation

> ComplianceAccList GetComplianceAccreditation(ctx).Limit(limit).Execute()

Returns the org's tracked accreditation-state records, newest first — evidence entries the org keeps, never a platform certification.



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
	limit := int64(789) // int64 | Limit caps the rows returned; non-positive means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetComplianceAccreditation(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceAccreditation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceAccreditation`: ComplianceAccList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceAccreditation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceAccreditationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows returned; non-positive means the server default. | 

### Return type

[**ComplianceAccList**](ComplianceAccList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceAccreditationById

> ComplianceAccView GetComplianceAccreditationById(ctx, id).Execute()

Returns one tracked accreditation record.



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
	id := "acc_1" // string | ID is the accreditation record to read, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetComplianceAccreditationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceAccreditationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceAccreditationById`: ComplianceAccView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceAccreditationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the accreditation record to read, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceAccreditationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComplianceAccView**](ComplianceAccView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceAudit

> ComplianceAuditList GetComplianceAudit(ctx).Result(result).Execute()

AuditRead is the compliance read of the SHARED tamper-evident audit plane — the SOC 2 posture surface (privileged actions: who started/decided what, when).



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
	result := "result_example" // string | Result filters rows by outcome result: success, deny, or error; empty means all. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetComplianceAudit(context.Background()).Result(result).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceAudit`: ComplianceAuditList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **result** | **string** | Result filters rows by outcome result: success, deny, or error; empty means all. | 

### Return type

[**ComplianceAuditList**](ComplianceAuditList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceHealth

> ComplianceHealthView GetComplianceHealth(ctx).Execute()

Reports subsystem liveness and the wired verification provider.



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
	resp, r, err := apiClient.ComplianceAPI.GetComplianceHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceHealth`: ComplianceHealthView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceHealthRequest struct via the builder pattern


### Return type

[**ComplianceHealthView**](ComplianceHealthView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceRecords

> ComplianceRecordList GetComplianceRecords(ctx).Limit(limit).Execute()

ListRecords is the unified compliance-record view for the org: its verifications and accreditation records together, each provider-reported or tracked, never platform-asserted.



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
	limit := int64(789) // int64 | Limit caps the rows returned; non-positive means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetComplianceRecords(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceRecords`: ComplianceRecordList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows returned; non-positive means the server default. | 

### Return type

[**ComplianceRecordList**](ComplianceRecordList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceStatus

> ComplianceStatusView GetComplianceStatus(ctx).Execute()

Status is the org's honest posture read: the wired provider and the per-status tally of its verifications.



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
	resp, r, err := apiClient.ComplianceAPI.GetComplianceStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceStatus`: ComplianceStatusView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceStatusRequest struct via the builder pattern


### Return type

[**ComplianceStatusView**](ComplianceStatusView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceSubjects

> ComplianceSubjectList GetComplianceSubjects(ctx).Limit(limit).Execute()

Returns the org's subjects as PII-MINIMIZED summaries — no name or email, only whether an email is on file.



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
	limit := int64(789) // int64 | Limit caps the rows returned; non-positive means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetComplianceSubjects(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceSubjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceSubjects`: ComplianceSubjectList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceSubjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows returned; non-positive means the server default. | 

### Return type

[**ComplianceSubjectList**](ComplianceSubjectList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceSubjectsById

> ComplianceSubject GetComplianceSubjectsById(ctx, id).Execute()

Returns one subject WITH its contact PII — the only surface that returns it, and only to the owning org.



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
	id := "sub_1" // string | ID is the subject to read, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetComplianceSubjectsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceSubjectsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceSubjectsById`: ComplianceSubject
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceSubjectsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the subject to read, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceSubjectsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComplianceSubject**](ComplianceSubject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceVerifications

> ComplianceCheckList GetComplianceVerifications(ctx).Limit(limit).Execute()

Returns the org's KYC/KYB verifications, newest first — opaque subject references and provider-reported statuses only, no subject PII.



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
	limit := int64(789) // int64 | Limit caps the rows returned; non-positive means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetComplianceVerifications(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceVerifications`: ComplianceCheckList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceVerifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows returned; non-positive means the server default. | 

### Return type

[**ComplianceCheckList**](ComplianceCheckList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceVerificationsById

> ComplianceCheckView GetComplianceVerificationsById(ctx, id).Execute()

Returns one verification — its opaque subject reference and provider-reported status, no subject PII.



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
	id := "chk_1" // string | ID is the verification to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetComplianceVerificationsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetComplianceVerificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceVerificationsById`: ComplianceCheckView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetComplianceVerificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the verification to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceVerificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComplianceCheckView**](ComplianceCheckView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComplianceAccreditation

> ComplianceAccView PostComplianceAccreditation(ctx).ComplianceAccreditationReq(complianceAccreditationReq).Execute()

Records an ASSERTED accreditation state for a subject — the subject's own assertion, with no verifier.



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
	complianceAccreditationReq := *openapiclient.NewComplianceAccreditationReq() // ComplianceAccreditationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.PostComplianceAccreditation(context.Background()).ComplianceAccreditationReq(complianceAccreditationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.PostComplianceAccreditation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostComplianceAccreditation`: ComplianceAccView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.PostComplianceAccreditation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostComplianceAccreditationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **complianceAccreditationReq** | [**ComplianceAccreditationReq**](ComplianceAccreditationReq.md) |  | 

### Return type

[**ComplianceAccView**](ComplianceAccView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComplianceAccreditationByIdDecision

> ComplianceAccView PostComplianceAccreditationByIdDecision(ctx, id).ComplianceAccreditationDecision(complianceAccreditationDecision).Execute()

Records an org reviewer's decision on an accreditation record — a reviewer confirmation, a provider verification the reviewer has evidence of (a CPA/attorney letter, a verifier report), a rejection, or an expiry.



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
	id := "acc_1" // string | ID is the accreditation record to decide, from the path.
	complianceAccreditationDecision := *openapiclient.NewComplianceAccreditationDecision() // ComplianceAccreditationDecision | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.PostComplianceAccreditationByIdDecision(context.Background(), id).ComplianceAccreditationDecision(complianceAccreditationDecision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.PostComplianceAccreditationByIdDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostComplianceAccreditationByIdDecision`: ComplianceAccView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.PostComplianceAccreditationByIdDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the accreditation record to decide, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostComplianceAccreditationByIdDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **complianceAccreditationDecision** | [**ComplianceAccreditationDecision**](ComplianceAccreditationDecision.md) |  | 

### Return type

[**ComplianceAccView**](ComplianceAccView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComplianceSubjects

> ComplianceSubject PostComplianceSubjects(ctx).ComplianceSubjectReq(complianceSubjectReq).Execute()

Records a party the org is verifying as part of its own onboarding/compliance — a team member, vendor, customer, or counterparty.



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
	complianceSubjectReq := *openapiclient.NewComplianceSubjectReq() // ComplianceSubjectReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.PostComplianceSubjects(context.Background()).ComplianceSubjectReq(complianceSubjectReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.PostComplianceSubjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostComplianceSubjects`: ComplianceSubject
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.PostComplianceSubjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostComplianceSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **complianceSubjectReq** | [**ComplianceSubjectReq**](ComplianceSubjectReq.md) |  | 

### Return type

[**ComplianceSubject**](ComplianceSubject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComplianceVerifications

> ComplianceCheckView PostComplianceVerifications(ctx).ComplianceVerificationReq(complianceVerificationReq).Execute()

Begins a KYC/KYB verification of a subject through the wired provider — an existing subject by id, or one created inline from the request.



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
	complianceVerificationReq := *openapiclient.NewComplianceVerificationReq() // ComplianceVerificationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.PostComplianceVerifications(context.Background()).ComplianceVerificationReq(complianceVerificationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.PostComplianceVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostComplianceVerifications`: ComplianceCheckView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.PostComplianceVerifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostComplianceVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **complianceVerificationReq** | [**ComplianceVerificationReq**](ComplianceVerificationReq.md) |  | 

### Return type

[**ComplianceCheckView**](ComplianceCheckView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComplianceVerificationsByIdDecision

> ComplianceCheckView PostComplianceVerificationsByIdDecision(ctx, id).ComplianceVerificationDecision(complianceVerificationDecision).Execute()

Records a privileged reviewer's MANUAL decision on a verification — the human-in-the-loop path, and the ONLY route to a passing status when no real provider is wired.



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
	id := "chk_1" // string | ID is the verification to decide, from the path.
	complianceVerificationDecision := *openapiclient.NewComplianceVerificationDecision() // ComplianceVerificationDecision | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.PostComplianceVerificationsByIdDecision(context.Background(), id).ComplianceVerificationDecision(complianceVerificationDecision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.PostComplianceVerificationsByIdDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostComplianceVerificationsByIdDecision`: ComplianceCheckView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.PostComplianceVerificationsByIdDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the verification to decide, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostComplianceVerificationsByIdDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **complianceVerificationDecision** | [**ComplianceVerificationDecision**](ComplianceVerificationDecision.md) |  | 

### Return type

[**ComplianceCheckView**](ComplianceCheckView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComplianceVerificationsByIdRefresh

> ComplianceCheckView PostComplianceVerificationsByIdRefresh(ctx, id).Execute()

Polls the wired provider for its current decision and records it, ATTRIBUTED to the provider — the internal PULL reconcile.



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
	id := "chk_1" // string | ID is the verification to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.PostComplianceVerificationsByIdRefresh(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.PostComplianceVerificationsByIdRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostComplianceVerificationsByIdRefresh`: ComplianceCheckView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.PostComplianceVerificationsByIdRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the verification to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostComplianceVerificationsByIdRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComplianceCheckView**](ComplianceCheckView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComplianceVerificationsWebhook

> PostComplianceVerificationsWebhook(ctx).Execute()

Provider push that settles a verification, authenticated by HMAC signature



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
	r, err := apiClient.ComplianceAPI.PostComplianceVerificationsWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.PostComplianceVerificationsWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostComplianceVerificationsWebhookRequest struct via the builder pattern


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

