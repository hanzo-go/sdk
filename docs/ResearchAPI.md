# \ResearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1ResearchArtifacts**](ResearchAPI.md#CloudGetV1ResearchArtifacts) | **Get** /v1/research/artifacts | ListResearchArtifacts returns the caller org&#39;s research-diary feed newest-first — the snapshots and reports tied to its runs, as metadata and content addresses; the bytes themselves are fetched by hash.
[**CloudGetV1ResearchArtifactsBySha256**](ResearchAPI.md#CloudGetV1ResearchArtifactsBySha256) | **Get** /v1/research/artifacts/{sha256} | 
[**CloudGetV1ResearchExperiments**](ResearchAPI.md#CloudGetV1ResearchExperiments) | **Get** /v1/research/experiments | ListExperiments returns the caller org&#39;s CANONICAL experiments — the deterministic deduped view over the versioned history.
[**CloudGetV1ResearchProjects**](ResearchAPI.md#CloudGetV1ResearchProjects) | **Get** /v1/research/projects | ListResearchProjects returns every research project in the caller&#39;s org with its real totals — canonical and retained side by side — which is the ops board&#39;s \&quot;every project + real totals\&quot; view.
[**CloudGetV1ResearchTotals**](ResearchAPI.md#CloudGetV1ResearchTotals) | **Get** /v1/research/totals | GetResearchTotals returns the caller org&#39;s headline aggregate plus a per-kind breakdown — the observatory&#39;s poll target.
[**CloudPostV1ResearchArtifacts**](ResearchAPI.md#CloudPostV1ResearchArtifacts) | **Post** /v1/research/artifacts | RecordResearchArtifact records one research-diary artifact — a board snapshot or a generated report — CONTENT-ADDRESSED inside the trust boundary.
[**CloudPostV1ResearchExperiments**](ResearchAPI.md#CloudPostV1ResearchExperiments) | **Post** /v1/research/experiments | IngestExperiments appends one batch of experiment and attempt versions to the caller org&#39;s evidence store, idempotently by content, then rolls it up to the analytics plane best-effort.
[**CloudPostV1ResearchGrants**](ResearchAPI.md#CloudPostV1ResearchGrants) | **Post** /v1/research/grants | GrantResearchVisibility records the SEPARATE authorization an upload never implies: a record&#39;s visibility (private, org or public) and, for a run, its training and commons-publication consent.



## CloudGetV1ResearchArtifacts

> CloudArtifactsOut CloudGetV1ResearchArtifacts(ctx).Project(project).Run(run).Since(since).Execute()

ListResearchArtifacts returns the caller org's research-diary feed newest-first — the snapshots and reports tied to its runs, as metadata and content addresses; the bytes themselves are fetched by hash.



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
	project := "project_example" // string | Project narrows to one project. Empty takes the caller's project scope. (optional)
	run := "run_example" // string | Run narrows to one run's artifacts by its stable id. (optional)
	since := int32(56) // int32 | Since bounds the feed to artifacts recorded at or after this unix second. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.CloudGetV1ResearchArtifacts(context.Background()).Project(project).Run(run).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.CloudGetV1ResearchArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ResearchArtifacts`: CloudArtifactsOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.CloudGetV1ResearchArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ResearchArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows to one project. Empty takes the caller&#39;s project scope. | 
 **run** | **string** | Run narrows to one run&#39;s artifacts by its stable id. | 
 **since** | **int32** | Since bounds the feed to artifacts recorded at or after this unix second. | 

### Return type

[**CloudArtifactsOut**](CloudArtifactsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ResearchArtifactsBySha256

> CloudGetV1ResearchArtifactsBySha256(ctx, sha256).Execute()



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
	sha256 := "sha256_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResearchAPI.CloudGetV1ResearchArtifactsBySha256(context.Background(), sha256).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.CloudGetV1ResearchArtifactsBySha256``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sha256** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ResearchArtifactsBySha256Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetV1ResearchExperiments

> CloudExperimentsOut CloudGetV1ResearchExperiments(ctx).Project(project).Kind(kind).Execute()

ListExperiments returns the caller org's CANONICAL experiments — the deterministic deduped view over the versioned history.



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
	project := "project_example" // string | Project narrows to one project. Empty reads the org's whole set across projects. (optional)
	kind := "kind_example" // string | Kind narrows to one discriminator: benchmark, kernel-perf, training, ablation or policy-eval. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.CloudGetV1ResearchExperiments(context.Background()).Project(project).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.CloudGetV1ResearchExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ResearchExperiments`: CloudExperimentsOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.CloudGetV1ResearchExperiments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ResearchExperimentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows to one project. Empty reads the org&#39;s whole set across projects. | 
 **kind** | **string** | Kind narrows to one discriminator: benchmark, kernel-perf, training, ablation or policy-eval. | 

### Return type

[**CloudExperimentsOut**](CloudExperimentsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ResearchProjects

> CloudProjectsOut CloudGetV1ResearchProjects(ctx).Execute()

ListResearchProjects returns every research project in the caller's org with its real totals — canonical and retained side by side — which is the ops board's \"every project + real totals\" view.



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
	resp, r, err := apiClient.ResearchAPI.CloudGetV1ResearchProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.CloudGetV1ResearchProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ResearchProjects`: CloudProjectsOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.CloudGetV1ResearchProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ResearchProjectsRequest struct via the builder pattern


### Return type

[**CloudProjectsOut**](CloudProjectsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ResearchTotals

> CloudResearchTotals CloudGetV1ResearchTotals(ctx).Project(project).Execute()

GetResearchTotals returns the caller org's headline aggregate plus a per-kind breakdown — the observatory's poll target.



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
	project := "project_example" // string | Project narrows the aggregate to one project. Empty aggregates the whole org. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.CloudGetV1ResearchTotals(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.CloudGetV1ResearchTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ResearchTotals`: CloudResearchTotals
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.CloudGetV1ResearchTotals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ResearchTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows the aggregate to one project. Empty aggregates the whole org. | 

### Return type

[**CloudResearchTotals**](CloudResearchTotals.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ResearchArtifacts

> CloudArtifactOut CloudPostV1ResearchArtifacts(ctx).CloudArtifact(cloudArtifact).Execute()

RecordResearchArtifact records one research-diary artifact — a board snapshot or a generated report — CONTENT-ADDRESSED inside the trust boundary.



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
	cloudArtifact := *openapiclient.NewCloudArtifact() // CloudArtifact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.CloudPostV1ResearchArtifacts(context.Background()).CloudArtifact(cloudArtifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.CloudPostV1ResearchArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ResearchArtifacts`: CloudArtifactOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.CloudPostV1ResearchArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ResearchArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudArtifact** | [**CloudArtifact**](CloudArtifact.md) |  | 

### Return type

[**CloudArtifactOut**](CloudArtifactOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ResearchExperiments

> CloudIngestOut CloudPostV1ResearchExperiments(ctx).CloudIngestRequest(cloudIngestRequest).Execute()

IngestExperiments appends one batch of experiment and attempt versions to the caller org's evidence store, idempotently by content, then rolls it up to the analytics plane best-effort.



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
	cloudIngestRequest := *openapiclient.NewCloudIngestRequest() // CloudIngestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.CloudPostV1ResearchExperiments(context.Background()).CloudIngestRequest(cloudIngestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.CloudPostV1ResearchExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ResearchExperiments`: CloudIngestOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.CloudPostV1ResearchExperiments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ResearchExperimentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudIngestRequest** | [**CloudIngestRequest**](CloudIngestRequest.md) |  | 

### Return type

[**CloudIngestOut**](CloudIngestOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ResearchGrants

> CloudGrantOut CloudPostV1ResearchGrants(ctx).CloudGrantRequest(cloudGrantRequest).Execute()

GrantResearchVisibility records the SEPARATE authorization an upload never implies: a record's visibility (private, org or public) and, for a run, its training and commons-publication consent.



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
	cloudGrantRequest := *openapiclient.NewCloudGrantRequest() // CloudGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.CloudPostV1ResearchGrants(context.Background()).CloudGrantRequest(cloudGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.CloudPostV1ResearchGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ResearchGrants`: CloudGrantOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.CloudPostV1ResearchGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ResearchGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudGrantRequest** | [**CloudGrantRequest**](CloudGrantRequest.md) |  | 

### Return type

[**CloudGrantOut**](CloudGrantOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

