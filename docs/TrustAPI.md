# \TrustAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrustByKindById**](TrustAPI.md#DeleteTrustByKindById) | **Delete** /v1/trust/{kind}/{id} | Removes one record from a section of your organization&#39;s trust centre.
[**GetTrust**](TrustAPI.md#GetTrust) | **Get** /v1/trust | Reads YOUR organization&#39;s whole trust centre, including the addresses of your own gated documents.
[**GetTrustControls**](TrustAPI.md#GetTrustControls) | **Get** /v1/trust/controls | Lists every control your organization publishes, with the counts.
[**GetTrustControlsById**](TrustAPI.md#GetTrustControlsById) | **Get** /v1/trust/controls/{id} | Reads one control by id.
[**GetTrustCoverage**](TrustAPI.md#GetTrustCoverage) | **Get** /v1/trust/coverage | Reads coverage: per framework, how many clauses have an automated control behind them, how many are partial, and how many have none — each carrying the unit it is counted in, because \&quot;12 of 20\&quot; is not a fact until you know what the 20 are.
[**GetTrustCoverageByFramework**](TrustAPI.md#GetTrustCoverageByFramework) | **Get** /v1/trust/coverage/{framework} | Reads one framework clause by clause: every clause the standard publishes, what covers it, and which controls stand behind it — so a coverage number can be checked line by line rather than taken on trust.
[**GetTrustDocuments**](TrustAPI.md#GetTrustDocuments) | **Get** /v1/trust/documents | Lists your organization&#39;s documents.
[**GetTrustEvidence**](TrustAPI.md#GetTrustEvidence) | **Get** /v1/trust/evidence | Reads the audit rows that stand behind one control, over a window.
[**GetTrustFaq**](TrustAPI.md#GetTrustFaq) | **Get** /v1/trust/faq | Lists your knowledge base — the questions a reviewer asks, answered once.
[**GetTrustFrameworks**](TrustAPI.md#GetTrustFrameworks) | **Get** /v1/trust/frameworks | Lists the frameworks coverage is computed against, and how many clauses each publishes.
[**GetTrustPolicies**](TrustAPI.md#GetTrustPolicies) | **Get** /v1/trust/policies | Lists your organization&#39;s published policies.
[**GetTrustProfile**](TrustAPI.md#GetTrustProfile) | **Get** /v1/trust/profile | Reads your organization&#39;s trust-centre profile — the name, tagline and summary a visitor sees, whether the centre is published, and where to send somebody who wants a gated document.
[**GetTrustPublishedByOrg**](TrustAPI.md#GetTrustPublishedByOrg) | **Get** /v1/trust/published/{org} | Reads a published trust centre — the whole thing in one answer: the organization&#39;s profile, its control inventory, coverage computed against each framework&#39;s whole published clause list, its documents, subprocessors, policies, knowledge base, updates and risk profile.
[**GetTrustRisk**](TrustAPI.md#GetTrustRisk) | **Get** /v1/trust/risk | Reads your risk profile — the label and value pairs describing what your organization handles and how.
[**GetTrustSubprocessors**](TrustAPI.md#GetTrustSubprocessors) | **Get** /v1/trust/subprocessors | Lists the third parties your organization sends data to, each naming what it is for.
[**GetTrustUpdates**](TrustAPI.md#GetTrustUpdates) | **Get** /v1/trust/updates | Lists your trust-centre updates, newest as you ordered them.
[**PutTrustByKindById**](TrustAPI.md#PutTrustByKindById) | **Put** /v1/trust/{kind}/{id} | Writes one record into a section of YOUR organization&#39;s trust centre — profile, control, document, subprocessor, policy, faq, update or risk.



## DeleteTrustByKindById

> Dropped DeleteTrustByKindById(ctx, kind, id).Execute()

Removes one record from a section of your organization's trust centre.



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
	kind := "faq" // string | Kind is the section — profile, control, document, subprocessor, policy, faq, update or risk. Anything else is not found. The URL is the authority: a value here is bound from the path, which zip binds last.
	id := "where-is-data-held" // string | ID is the record's id within that section. The single-valued sections (profile, risk) hold one record whatever id is named.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrustAPI.DeleteTrustByKindById(context.Background(), kind, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.DeleteTrustByKindById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrustByKindById`: Dropped
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.DeleteTrustByKindById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** | Kind is the section — profile, control, document, subprocessor, policy, faq, update or risk. Anything else is not found. The URL is the authority: a value here is bound from the path, which zip binds last. | 
**id** | **string** | ID is the record&#39;s id within that section. The single-valued sections (profile, risk) hold one record whatever id is named. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrustByKindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Dropped**](Dropped.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrust

> Centre GetTrust(ctx).Execute()

Reads YOUR organization's whole trust centre, including the addresses of your own gated documents.



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
	resp, r, err := apiClient.TrustAPI.GetTrust(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrust``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrust`: Centre
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrust`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustRequest struct via the builder pattern


### Return type

[**Centre**](Centre.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustControls

> ControlList GetTrustControls(ctx).Execute()

Lists every control your organization publishes, with the counts.



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
	resp, r, err := apiClient.TrustAPI.GetTrustControls(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustControls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustControls`: ControlList
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustControls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustControlsRequest struct via the builder pattern


### Return type

[**ControlList**](ControlList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustControlsById

> interface{} GetTrustControlsById(ctx, id).Execute()

Reads one control by id.



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
	id := "iam.pkce.s256" // string | ID is the control's id, dotted lowercase — \"iam.pkce.s256\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrustAPI.GetTrustControlsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustControlsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustControlsById`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustControlsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the control&#39;s id, dotted lowercase — \&quot;iam.pkce.s256\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustControlsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetTrustCoverage

> TrustCoverage GetTrustCoverage(ctx).Execute()

Reads coverage: per framework, how many clauses have an automated control behind them, how many are partial, and how many have none — each carrying the unit it is counted in, because \"12 of 20\" is not a fact until you know what the 20 are.



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
	resp, r, err := apiClient.TrustAPI.GetTrustCoverage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustCoverage`: TrustCoverage
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustCoverage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustCoverageRequest struct via the builder pattern


### Return type

[**TrustCoverage**](TrustCoverage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustCoverageByFramework

> ClauseCoverage GetTrustCoverageByFramework(ctx, framework).Execute()

Reads one framework clause by clause: every clause the standard publishes, what covers it, and which controls stand behind it — so a coverage number can be checked line by line rather than taken on trust.



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
	framework := "soc2" // string | Framework is the framework id — \"soc2\", \"iso27001\", \"nist80053\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrustAPI.GetTrustCoverageByFramework(context.Background(), framework).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustCoverageByFramework``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustCoverageByFramework`: ClauseCoverage
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustCoverageByFramework`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**framework** | **string** | Framework is the framework id — \&quot;soc2\&quot;, \&quot;iso27001\&quot;, \&quot;nist80053\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustCoverageByFrameworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClauseCoverage**](ClauseCoverage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustDocuments

> TrustDocuments GetTrustDocuments(ctx).Execute()

Lists your organization's documents.



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
	resp, r, err := apiClient.TrustAPI.GetTrustDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustDocuments`: TrustDocuments
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustDocuments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustDocumentsRequest struct via the builder pattern


### Return type

[**TrustDocuments**](TrustDocuments.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustEvidence

> interface{} GetTrustEvidence(ctx).Control(control).From(from).To(to).Limit(limit).Execute()

Reads the audit rows that stand behind one control, over a window.



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
	control := "iam.refresh.rotation" // string | Control is the control id whose trail to read. Required. (optional)
	from := "2026-01-01" // string | From is the inclusive lower bound, an RFC 3339 date or instant (\"2026-01-01\" or \"2026-01-01T00:00:00Z\"). Empty leaves it unbounded. A malformed bound is refused rather than silently widening the window. (optional)
	to := "to_example" // string | To is the upper bound, same form and same tolerance. (optional)
	limit := "50" // string | Limit caps the rows returned, 1..1000, default 100. It is a string because an unparseable value is refused rather than read as zero. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrustAPI.GetTrustEvidence(context.Background()).Control(control).From(from).To(to).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustEvidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustEvidence`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustEvidence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **control** | **string** | Control is the control id whose trail to read. Required. | 
 **from** | **string** | From is the inclusive lower bound, an RFC 3339 date or instant (\&quot;2026-01-01\&quot; or \&quot;2026-01-01T00:00:00Z\&quot;). Empty leaves it unbounded. A malformed bound is refused rather than silently widening the window. | 
 **to** | **string** | To is the upper bound, same form and same tolerance. | 
 **limit** | **string** | Limit caps the rows returned, 1..1000, default 100. It is a string because an unparseable value is refused rather than read as zero. | 

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


## GetTrustFaq

> FaqList GetTrustFaq(ctx).Execute()

Lists your knowledge base — the questions a reviewer asks, answered once.



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
	resp, r, err := apiClient.TrustAPI.GetTrustFaq(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustFaq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustFaq`: FaqList
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustFaq`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustFaqRequest struct via the builder pattern


### Return type

[**FaqList**](FaqList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustFrameworks

> FrameworkList GetTrustFrameworks(ctx).Execute()

Lists the frameworks coverage is computed against, and how many clauses each publishes.



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
	resp, r, err := apiClient.TrustAPI.GetTrustFrameworks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustFrameworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustFrameworks`: FrameworkList
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustFrameworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustFrameworksRequest struct via the builder pattern


### Return type

[**FrameworkList**](FrameworkList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustPolicies

> PolicyList GetTrustPolicies(ctx).Execute()

Lists your organization's published policies.



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
	resp, r, err := apiClient.TrustAPI.GetTrustPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustPolicies`: PolicyList
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustPoliciesRequest struct via the builder pattern


### Return type

[**PolicyList**](PolicyList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustProfile

> interface{} GetTrustProfile(ctx).Execute()

Reads your organization's trust-centre profile — the name, tagline and summary a visitor sees, whether the centre is published, and where to send somebody who wants a gated document.



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
	resp, r, err := apiClient.TrustAPI.GetTrustProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustProfile`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustProfileRequest struct via the builder pattern


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


## GetTrustPublishedByOrg

> Centre GetTrustPublishedByOrg(ctx, org).Execute()

Reads a published trust centre — the whole thing in one answer: the organization's profile, its control inventory, coverage computed against each framework's whole published clause list, its documents, subprocessors, policies, knowledge base, updates and risk profile.



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
	org := "hanzo" // string | Org is the organization's slug — the name in its address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrustAPI.GetTrustPublishedByOrg(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustPublishedByOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustPublishedByOrg`: Centre
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustPublishedByOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the organization&#39;s slug — the name in its address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustPublishedByOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Centre**](Centre.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustRisk

> interface{} GetTrustRisk(ctx).Execute()

Reads your risk profile — the label and value pairs describing what your organization handles and how.



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
	resp, r, err := apiClient.TrustAPI.GetTrustRisk(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustRisk`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustRisk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustRiskRequest struct via the builder pattern


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


## GetTrustSubprocessors

> SubprocessorList GetTrustSubprocessors(ctx).Execute()

Lists the third parties your organization sends data to, each naming what it is for.



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
	resp, r, err := apiClient.TrustAPI.GetTrustSubprocessors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustSubprocessors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustSubprocessors`: SubprocessorList
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustSubprocessors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustSubprocessorsRequest struct via the builder pattern


### Return type

[**SubprocessorList**](SubprocessorList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustUpdates

> UpdateList GetTrustUpdates(ctx).Execute()

Lists your trust-centre updates, newest as you ordered them.



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
	resp, r, err := apiClient.TrustAPI.GetTrustUpdates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.GetTrustUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustUpdates`: UpdateList
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.GetTrustUpdates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustUpdatesRequest struct via the builder pattern


### Return type

[**UpdateList**](UpdateList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTrustByKindById

> Written PutTrustByKindById(ctx, kind, id).SectionWrite(sectionWrite).Execute()

Writes one record into a section of YOUR organization's trust centre — profile, control, document, subprocessor, policy, faq, update or risk.



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
	kind := "subprocessor" // string | Kind is the section being written. The URL is the authority.
	id := "acme-cloud" // string | ID is the record's id. Omit it on a create and one is minted; the single-valued sections (profile, risk) hold one record whatever is named.
	sectionWrite := *openapiclient.NewSectionWrite() // SectionWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrustAPI.PutTrustByKindById(context.Background(), kind, id).SectionWrite(sectionWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustAPI.PutTrustByKindById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTrustByKindById`: Written
	fmt.Fprintf(os.Stdout, "Response from `TrustAPI.PutTrustByKindById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** | Kind is the section being written. The URL is the authority. | 
**id** | **string** | ID is the record&#39;s id. Omit it on a create and one is minted; the single-valued sections (profile, risk) hold one record whatever is named. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTrustByKindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sectionWrite** | [**SectionWrite**](SectionWrite.md) |  | 

### Return type

[**Written**](Written.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

