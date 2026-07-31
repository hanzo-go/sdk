# \ComplianceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1ComplianceAccreditation**](ComplianceAPI.md#CloudGetV1ComplianceAccreditation) | **Get** /v1/compliance/accreditation | ListAccreditation returns the org&#39;s tracked accreditation-state records, newest first — evidence entries the org keeps, never a platform certification.
[**CloudGetV1ComplianceAccreditationId**](ComplianceAPI.md#CloudGetV1ComplianceAccreditationId) | **Get** /v1/compliance/accreditation/{id} | GetAccreditation returns one tracked accreditation record.
[**CloudGetV1ComplianceAudit**](ComplianceAPI.md#CloudGetV1ComplianceAudit) | **Get** /v1/compliance/audit | AuditRead is the compliance-scoped read of the SHARED tamper-evident audit plane — the SOC 2 posture surface (privileged actions: who started/decided what, when).
[**CloudGetV1ComplianceHealth**](ComplianceAPI.md#CloudGetV1ComplianceHealth) | **Get** /v1/compliance/health | Health reports subsystem liveness and the wired verification provider.
[**CloudGetV1ComplianceRecords**](ComplianceAPI.md#CloudGetV1ComplianceRecords) | **Get** /v1/compliance/records | ListRecords is the unified compliance-record view for the org: its verifications and accreditation records together, each provider-reported or tracked, never platform-asserted.
[**CloudGetV1ComplianceStatus**](ComplianceAPI.md#CloudGetV1ComplianceStatus) | **Get** /v1/compliance/status | Status is the org&#39;s honest posture read: the wired provider and the per-status tally of its verifications.
[**CloudGetV1ComplianceSubjects**](ComplianceAPI.md#CloudGetV1ComplianceSubjects) | **Get** /v1/compliance/subjects | ListSubjects returns the org&#39;s subjects as PII-MINIMIZED summaries — no name or email, only whether an email is on file.
[**CloudGetV1ComplianceSubjectsId**](ComplianceAPI.md#CloudGetV1ComplianceSubjectsId) | **Get** /v1/compliance/subjects/{id} | GetSubject returns one subject WITH its contact PII — the only surface that returns it, and only to the owning org.
[**CloudGetV1ComplianceVerifications**](ComplianceAPI.md#CloudGetV1ComplianceVerifications) | **Get** /v1/compliance/verifications | ListVerifications returns the org&#39;s KYC/KYB verifications, newest first — opaque subject references and provider-reported statuses only, no subject PII.
[**CloudGetV1ComplianceVerificationsId**](ComplianceAPI.md#CloudGetV1ComplianceVerificationsId) | **Get** /v1/compliance/verifications/{id} | GetVerification returns one verification — its opaque subject reference and provider-reported status, no subject PII.
[**CloudPostV1ComplianceAccreditation**](ComplianceAPI.md#CloudPostV1ComplianceAccreditation) | **Post** /v1/compliance/accreditation | CreateAccreditation records an ASSERTED accreditation state for a subject — the subject&#39;s own assertion, with no verifier.
[**CloudPostV1ComplianceAccreditationIdDecision**](ComplianceAPI.md#CloudPostV1ComplianceAccreditationIdDecision) | **Post** /v1/compliance/accreditation/{id}/decision | DecideAccreditation records an org reviewer&#39;s decision on an accreditation record — a reviewer confirmation, a provider verification the reviewer has evidence of (a CPA/attorney letter, a verifier report), a rejection, or an expiry.
[**CloudPostV1ComplianceSubjects**](ComplianceAPI.md#CloudPostV1ComplianceSubjects) | **Post** /v1/compliance/subjects | CreateSubject records a party the org is verifying as part of its own onboarding/compliance — a team member, vendor, customer, or counterparty.
[**CloudPostV1ComplianceVerifications**](ComplianceAPI.md#CloudPostV1ComplianceVerifications) | **Post** /v1/compliance/verifications | StartVerification begins a KYC/KYB verification of a subject through the wired provider — an existing subject by id, or one created inline from the request.
[**CloudPostV1ComplianceVerificationsIdDecision**](ComplianceAPI.md#CloudPostV1ComplianceVerificationsIdDecision) | **Post** /v1/compliance/verifications/{id}/decision | DecideVerification records a privileged reviewer&#39;s MANUAL decision on a verification — the human-in-the-loop path, and the ONLY route to a passing status when no real provider is wired.
[**CloudPostV1ComplianceVerificationsIdRefresh**](ComplianceAPI.md#CloudPostV1ComplianceVerificationsIdRefresh) | **Post** /v1/compliance/verifications/{id}/refresh | RefreshVerification polls the wired provider for its current decision and records it, ATTRIBUTED to the provider — the internal PULL reconcile.
[**CloudPostV1ComplianceVerificationsWebhook**](ComplianceAPI.md#CloudPostV1ComplianceVerificationsWebhook) | **Post** /v1/compliance/verifications/webhook | 



## CloudGetV1ComplianceAccreditation

> CloudAccList CloudGetV1ComplianceAccreditation(ctx).Limit(limit).Execute()

ListAccreditation returns the org's tracked accreditation-state records, newest first — evidence entries the org keeps, never a platform certification.



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
	limit := int32(56) // int32 | Limit caps the rows returned; non-positive means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceAccreditation(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceAccreditation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceAccreditation`: CloudAccList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceAccreditation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceAccreditationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; non-positive means the server default. | 

### Return type

[**CloudAccList**](CloudAccList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceAccreditationId

> CloudAccView CloudGetV1ComplianceAccreditationId(ctx, id).Execute()

GetAccreditation returns one tracked accreditation record.



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
	id := "acc_1" // string | ID is the accreditation record to read, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceAccreditationId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceAccreditationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceAccreditationId`: CloudAccView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceAccreditationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the accreditation record to read, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceAccreditationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAccView**](CloudAccView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceAudit

> CloudAuditList CloudGetV1ComplianceAudit(ctx).Result(result).Execute()

AuditRead is the compliance-scoped read of the SHARED tamper-evident audit plane — the SOC 2 posture surface (privileged actions: who started/decided what, when).



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
	result := "result_example" // string | Result filters rows by outcome result: success, deny, or error; empty means all. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceAudit(context.Background()).Result(result).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceAudit`: CloudAuditList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **result** | **string** | Result filters rows by outcome result: success, deny, or error; empty means all. | 

### Return type

[**CloudAuditList**](CloudAuditList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceHealth

> CloudHealthView CloudGetV1ComplianceHealth(ctx).Execute()

Health reports subsystem liveness and the wired verification provider.



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
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceHealth`: CloudHealthView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceHealthRequest struct via the builder pattern


### Return type

[**CloudHealthView**](CloudHealthView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceRecords

> CloudRecordList CloudGetV1ComplianceRecords(ctx).Limit(limit).Execute()

ListRecords is the unified compliance-record view for the org: its verifications and accreditation records together, each provider-reported or tracked, never platform-asserted.



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
	limit := int32(56) // int32 | Limit caps the rows returned; non-positive means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceRecords(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceRecords`: CloudRecordList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; non-positive means the server default. | 

### Return type

[**CloudRecordList**](CloudRecordList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceStatus

> CloudStatusView CloudGetV1ComplianceStatus(ctx).Execute()

Status is the org's honest posture read: the wired provider and the per-status tally of its verifications.



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
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceStatus`: CloudStatusView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceStatusRequest struct via the builder pattern


### Return type

[**CloudStatusView**](CloudStatusView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceSubjects

> CloudSubjectList CloudGetV1ComplianceSubjects(ctx).Limit(limit).Execute()

ListSubjects returns the org's subjects as PII-MINIMIZED summaries — no name or email, only whether an email is on file.



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
	limit := int32(56) // int32 | Limit caps the rows returned; non-positive means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceSubjects(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceSubjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceSubjects`: CloudSubjectList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceSubjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; non-positive means the server default. | 

### Return type

[**CloudSubjectList**](CloudSubjectList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceSubjectsId

> CloudSubject CloudGetV1ComplianceSubjectsId(ctx, id).Execute()

GetSubject returns one subject WITH its contact PII — the only surface that returns it, and only to the owning org.



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
	id := "sub_1" // string | ID is the subject to read, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceSubjectsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceSubjectsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceSubjectsId`: CloudSubject
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceSubjectsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the subject to read, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceSubjectsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSubject**](CloudSubject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceVerifications

> CloudCheckList CloudGetV1ComplianceVerifications(ctx).Limit(limit).Execute()

ListVerifications returns the org's KYC/KYB verifications, newest first — opaque subject references and provider-reported statuses only, no subject PII.



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
	limit := int32(56) // int32 | Limit caps the rows returned; non-positive means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceVerifications(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceVerifications`: CloudCheckList
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceVerifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; non-positive means the server default. | 

### Return type

[**CloudCheckList**](CloudCheckList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComplianceVerificationsId

> CloudCheckView CloudGetV1ComplianceVerificationsId(ctx, id).Execute()

GetVerification returns one verification — its opaque subject reference and provider-reported status, no subject PII.



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
	id := "chk_1" // string | ID is the verification to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudGetV1ComplianceVerificationsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudGetV1ComplianceVerificationsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ComplianceVerificationsId`: CloudCheckView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudGetV1ComplianceVerificationsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the verification to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComplianceVerificationsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCheckView**](CloudCheckView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComplianceAccreditation

> CloudAccView CloudPostV1ComplianceAccreditation(ctx).CloudAccreditationReq(cloudAccreditationReq).Execute()

CreateAccreditation records an ASSERTED accreditation state for a subject — the subject's own assertion, with no verifier.



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
	cloudAccreditationReq := *openapiclient.NewCloudAccreditationReq() // CloudAccreditationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudPostV1ComplianceAccreditation(context.Background()).CloudAccreditationReq(cloudAccreditationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudPostV1ComplianceAccreditation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ComplianceAccreditation`: CloudAccView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudPostV1ComplianceAccreditation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ComplianceAccreditationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAccreditationReq** | [**CloudAccreditationReq**](CloudAccreditationReq.md) |  | 

### Return type

[**CloudAccView**](CloudAccView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComplianceAccreditationIdDecision

> CloudAccView CloudPostV1ComplianceAccreditationIdDecision(ctx, id).CloudAccreditationDecision(cloudAccreditationDecision).Execute()

DecideAccreditation records an org reviewer's decision on an accreditation record — a reviewer confirmation, a provider verification the reviewer has evidence of (a CPA/attorney letter, a verifier report), a rejection, or an expiry.



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
	id := "acc_1" // string | ID is the accreditation record to decide, from the path.
	cloudAccreditationDecision := *openapiclient.NewCloudAccreditationDecision() // CloudAccreditationDecision | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudPostV1ComplianceAccreditationIdDecision(context.Background(), id).CloudAccreditationDecision(cloudAccreditationDecision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudPostV1ComplianceAccreditationIdDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ComplianceAccreditationIdDecision`: CloudAccView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudPostV1ComplianceAccreditationIdDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the accreditation record to decide, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ComplianceAccreditationIdDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAccreditationDecision** | [**CloudAccreditationDecision**](CloudAccreditationDecision.md) |  | 

### Return type

[**CloudAccView**](CloudAccView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComplianceSubjects

> CloudSubject CloudPostV1ComplianceSubjects(ctx).CloudSubjectReq(cloudSubjectReq).Execute()

CreateSubject records a party the org is verifying as part of its own onboarding/compliance — a team member, vendor, customer, or counterparty.



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
	cloudSubjectReq := *openapiclient.NewCloudSubjectReq() // CloudSubjectReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudPostV1ComplianceSubjects(context.Background()).CloudSubjectReq(cloudSubjectReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudPostV1ComplianceSubjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ComplianceSubjects`: CloudSubject
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudPostV1ComplianceSubjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ComplianceSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSubjectReq** | [**CloudSubjectReq**](CloudSubjectReq.md) |  | 

### Return type

[**CloudSubject**](CloudSubject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComplianceVerifications

> CloudCheckView CloudPostV1ComplianceVerifications(ctx).CloudVerificationReq(cloudVerificationReq).Execute()

StartVerification begins a KYC/KYB verification of a subject through the wired provider — an existing subject by id, or one created inline from the request.



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
	cloudVerificationReq := *openapiclient.NewCloudVerificationReq() // CloudVerificationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudPostV1ComplianceVerifications(context.Background()).CloudVerificationReq(cloudVerificationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudPostV1ComplianceVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ComplianceVerifications`: CloudCheckView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudPostV1ComplianceVerifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ComplianceVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudVerificationReq** | [**CloudVerificationReq**](CloudVerificationReq.md) |  | 

### Return type

[**CloudCheckView**](CloudCheckView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComplianceVerificationsIdDecision

> CloudCheckView CloudPostV1ComplianceVerificationsIdDecision(ctx, id).CloudVerificationDecision(cloudVerificationDecision).Execute()

DecideVerification records a privileged reviewer's MANUAL decision on a verification — the human-in-the-loop path, and the ONLY route to a passing status when no real provider is wired.



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
	id := "chk_1" // string | ID is the verification to decide, from the path.
	cloudVerificationDecision := *openapiclient.NewCloudVerificationDecision() // CloudVerificationDecision | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudPostV1ComplianceVerificationsIdDecision(context.Background(), id).CloudVerificationDecision(cloudVerificationDecision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudPostV1ComplianceVerificationsIdDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ComplianceVerificationsIdDecision`: CloudCheckView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudPostV1ComplianceVerificationsIdDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the verification to decide, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ComplianceVerificationsIdDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudVerificationDecision** | [**CloudVerificationDecision**](CloudVerificationDecision.md) |  | 

### Return type

[**CloudCheckView**](CloudCheckView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComplianceVerificationsIdRefresh

> CloudCheckView CloudPostV1ComplianceVerificationsIdRefresh(ctx, id).Execute()

RefreshVerification polls the wired provider for its current decision and records it, ATTRIBUTED to the provider — the internal PULL reconcile.



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
	id := "chk_1" // string | ID is the verification to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CloudPostV1ComplianceVerificationsIdRefresh(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudPostV1ComplianceVerificationsIdRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ComplianceVerificationsIdRefresh`: CloudCheckView
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CloudPostV1ComplianceVerificationsIdRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the verification to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ComplianceVerificationsIdRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCheckView**](CloudCheckView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComplianceVerificationsWebhook

> CloudPostV1ComplianceVerificationsWebhook(ctx).Execute()



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
	r, err := apiClient.ComplianceAPI.CloudPostV1ComplianceVerificationsWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CloudPostV1ComplianceVerificationsWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ComplianceVerificationsWebhookRequest struct via the builder pattern


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

