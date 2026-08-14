# \SecurityAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityFindings**](SecurityAPI.md#GetSecurityFindings) | **Get** /v1/security/findings | Is the org&#39;s findings — rule, severity, path, line, masked preview and fingerprint — newest first, across scans or within one.
[**GetSecurityFindingsById**](SecurityAPI.md#GetSecurityFindingsById) | **Get** /v1/security/findings/{id} | Returns a single finding: which rule fired, where (path and line), the masked preview and the SHA-256 fingerprint of the secret — the raw secret is not stored and cannot be read back.
[**GetSecurityHealth**](SecurityAPI.md#GetSecurityHealth) | **Get** /v1/security/health | Reports that the scanning subsystem is serving and how many secret-detection rules the engine holds.
[**GetSecurityRules**](SecurityAPI.md#GetSecurityRules) | **Get** /v1/security/rules | Is the secret-detection catalog the engine scans with.
[**GetSecurityScans**](SecurityAPI.md#GetSecurityScans) | **Get** /v1/security/scans | Is the org&#39;s scan history, newest first, each as the same summary the submission answered — files read, findings fired, tally by severity.
[**GetSecurityScansById**](SecurityAPI.md#GetSecurityScansById) | **Get** /v1/security/scans/{id} | Returns one scan together with every finding on it, so the detail view is one round-trip rather than a list call per scan.
[**PostSecurityScans**](SecurityAPI.md#PostSecurityScans) | **Post** /v1/security/scans | Runs the detection engine over a batch of files and answers 201 with the scan summary: how many files were read, how many findings fired, and the tally by severity.



## GetSecurityFindings

> FindingList GetSecurityFindings(ctx).ScanId(scanId).MinSeverity(minSeverity).Limit(limit).Execute()

Is the org's findings — rule, severity, path, line, masked preview and fingerprint — newest first, across scans or within one.



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
	scanId := "scanId_example" // string | ScanID narrows to a single scan. (optional)
	minSeverity := "minSeverity_example" // string | MinSeverity drops everything below that rank: critical, high, medium or low. A value outside that set is refused rather than quietly ignored, so a filter typo cannot read as \"no findings\". (optional)
	limit := int32(56) // int32 | Limit caps the page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetSecurityFindings(context.Background()).ScanId(scanId).MinSeverity(minSeverity).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityFindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityFindings`: FindingList
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetSecurityFindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityFindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scanId** | **string** | ScanID narrows to a single scan. | 
 **minSeverity** | **string** | MinSeverity drops everything below that rank: critical, high, medium or low. A value outside that set is refused rather than quietly ignored, so a filter typo cannot read as \&quot;no findings\&quot;. | 
 **limit** | **int32** | Limit caps the page. | 

### Return type

[**FindingList**](FindingList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityFindingsById

> FindingView GetSecurityFindingsById(ctx, id).Execute()

Returns a single finding: which rule fired, where (path and line), the masked preview and the SHA-256 fingerprint of the secret — the raw secret is not stored and cannot be read back.



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
	id := "id_example" // string | ID is the finding the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetSecurityFindingsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityFindingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityFindingsById`: FindingView
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetSecurityFindingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the finding the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityFindingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindingView**](FindingView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityHealth

> Ruleset GetSecurityHealth(ctx).Execute()

Reports that the scanning subsystem is serving and how many secret-detection rules the engine holds.



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
	resp, r, err := apiClient.SecurityAPI.GetSecurityHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityHealth`: Ruleset
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetSecurityHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityHealthRequest struct via the builder pattern


### Return type

[**Ruleset**](Ruleset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRules

> RuleList GetSecurityRules(ctx).Execute()

Is the secret-detection catalog the engine scans with.



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
	resp, r, err := apiClient.SecurityAPI.GetSecurityRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityRules`: RuleList
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetSecurityRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRulesRequest struct via the builder pattern


### Return type

[**RuleList**](RuleList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityScans

> ScanList GetSecurityScans(ctx).Limit(limit).Execute()

Is the org's scan history, newest first, each as the same summary the submission answered — files read, findings fired, tally by severity.



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
	limit := int32(56) // int32 | Limit caps the page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetSecurityScans(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityScans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityScans`: ScanList
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetSecurityScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the page. | 

### Return type

[**ScanList**](ScanList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityScansById

> ScanDetail GetSecurityScansById(ctx, id).Execute()

Returns one scan together with every finding on it, so the detail view is one round-trip rather than a list call per scan.



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
	id := "id_example" // string | ID is the scan the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetSecurityScansById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityScansById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityScansById`: ScanDetail
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetSecurityScansById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the scan the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityScansByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScanDetail**](ScanDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecurityScans

> ScanView PostSecurityScans(ctx).SubmitReq(submitReq).Execute()

Runs the detection engine over a batch of files and answers 201 with the scan summary: how many files were read, how many findings fired, and the tally by severity.



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
	submitReq := *openapiclient.NewSubmitReq([]openapiclient.Scan{*openapiclient.NewScan()}) // SubmitReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.PostSecurityScans(context.Background()).SubmitReq(submitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.PostSecurityScans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSecurityScans`: ScanView
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.PostSecurityScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSecurityScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitReq** | [**SubmitReq**](SubmitReq.md) |  | 

### Return type

[**ScanView**](ScanView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

