# \ResearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResearchArtifacts**](ResearchAPI.md#GetResearchArtifacts) | **Get** /v1/research/artifacts | Returns the caller org&#39;s research-diary feed newest-first — the snapshots and reports tied to its runs, as metadata and content addresses; the bytes themselves are fetched by hash.
[**GetResearchArtifactsBySha256**](ResearchAPI.md#GetResearchArtifactsBySha256) | **Get** /v1/research/artifacts/{sha256} | Fetch one recorded artifact&#39;s bytes by its content hash.
[**GetResearchExperiments**](ResearchAPI.md#GetResearchExperiments) | **Get** /v1/research/experiments | Returns the caller org&#39;s CANONICAL experiments — the deterministic deduped view over the versioned history.
[**GetResearchProjects**](ResearchAPI.md#GetResearchProjects) | **Get** /v1/research/projects | Returns every research project in the caller&#39;s org with its real totals — canonical and retained side by side — which is the ops board&#39;s \&quot;every project + real totals\&quot; view.
[**GetResearchTotals**](ResearchAPI.md#GetResearchTotals) | **Get** /v1/research/totals | Returns the caller org&#39;s headline aggregate plus a per-kind breakdown — the observatory&#39;s poll target.
[**PostResearchArtifacts**](ResearchAPI.md#PostResearchArtifacts) | **Post** /v1/research/artifacts | Records one research-diary artifact — a board snapshot or a generated report — CONTENT-ADDRESSED inside the trust boundary.
[**PostResearchExperiments**](ResearchAPI.md#PostResearchExperiments) | **Post** /v1/research/experiments | Appends one batch of experiment and attempt versions to the caller org&#39;s evidence store, idempotently by content, then rolls it up to the analytics plane best-effort.
[**PostResearchGrants**](ResearchAPI.md#PostResearchGrants) | **Post** /v1/research/grants | Records the SEPARATE authorization an upload never implies: a record&#39;s visibility (private, org or public) and, for a run, its training and commons-publication consent.



## GetResearchArtifacts

> ArtifactsOut GetResearchArtifacts(ctx).Project(project).Run(run).Since(since).Execute()

Returns the caller org's research-diary feed newest-first — the snapshots and reports tied to its runs, as metadata and content addresses; the bytes themselves are fetched by hash.



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
	resp, r, err := apiClient.ResearchAPI.GetResearchArtifacts(context.Background()).Project(project).Run(run).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.GetResearchArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResearchArtifacts`: ArtifactsOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.GetResearchArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResearchArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows to one project. Empty takes the caller&#39;s project scope. | 
 **run** | **string** | Run narrows to one run&#39;s artifacts by its stable id. | 
 **since** | **int32** | Since bounds the feed to artifacts recorded at or after this unix second. | 

### Return type

[**ArtifactsOut**](ArtifactsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResearchArtifactsBySha256

> GetResearchArtifactsBySha256(ctx, sha256).Execute()

Fetch one recorded artifact's bytes by its content hash.



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
	r, err := apiClient.ResearchAPI.GetResearchArtifactsBySha256(context.Background(), sha256).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.GetResearchArtifactsBySha256``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetResearchArtifactsBySha256Request struct via the builder pattern


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


## GetResearchExperiments

> ExperimentsOut GetResearchExperiments(ctx).Project(project).Kind(kind).Execute()

Returns the caller org's CANONICAL experiments — the deterministic deduped view over the versioned history.



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
	resp, r, err := apiClient.ResearchAPI.GetResearchExperiments(context.Background()).Project(project).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.GetResearchExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResearchExperiments`: ExperimentsOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.GetResearchExperiments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResearchExperimentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows to one project. Empty reads the org&#39;s whole set across projects. | 
 **kind** | **string** | Kind narrows to one discriminator: benchmark, kernel-perf, training, ablation or policy-eval. | 

### Return type

[**ExperimentsOut**](ExperimentsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResearchProjects

> ProjectsOut GetResearchProjects(ctx).Execute()

Returns every research project in the caller's org with its real totals — canonical and retained side by side — which is the ops board's \"every project + real totals\" view.



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
	resp, r, err := apiClient.ResearchAPI.GetResearchProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.GetResearchProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResearchProjects`: ProjectsOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.GetResearchProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetResearchProjectsRequest struct via the builder pattern


### Return type

[**ProjectsOut**](ProjectsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResearchTotals

> ResearchTotals GetResearchTotals(ctx).Project(project).Execute()

Returns the caller org's headline aggregate plus a per-kind breakdown — the observatory's poll target.



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
	resp, r, err := apiClient.ResearchAPI.GetResearchTotals(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.GetResearchTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResearchTotals`: ResearchTotals
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.GetResearchTotals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResearchTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows the aggregate to one project. Empty aggregates the whole org. | 

### Return type

[**ResearchTotals**](ResearchTotals.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostResearchArtifacts

> ArtifactOut PostResearchArtifacts(ctx).ResearchArtifact(researchArtifact).Execute()

Records one research-diary artifact — a board snapshot or a generated report — CONTENT-ADDRESSED inside the trust boundary.



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
	researchArtifact := *openapiclient.NewResearchArtifact() // ResearchArtifact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.PostResearchArtifacts(context.Background()).ResearchArtifact(researchArtifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.PostResearchArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResearchArtifacts`: ArtifactOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.PostResearchArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostResearchArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **researchArtifact** | [**ResearchArtifact**](ResearchArtifact.md) |  | 

### Return type

[**ArtifactOut**](ArtifactOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostResearchExperiments

> IngestOut PostResearchExperiments(ctx).IngestRequest(ingestRequest).Execute()

Appends one batch of experiment and attempt versions to the caller org's evidence store, idempotently by content, then rolls it up to the analytics plane best-effort.



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
	ingestRequest := *openapiclient.NewIngestRequest() // IngestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.PostResearchExperiments(context.Background()).IngestRequest(ingestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.PostResearchExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResearchExperiments`: IngestOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.PostResearchExperiments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostResearchExperimentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestRequest** | [**IngestRequest**](IngestRequest.md) |  | 

### Return type

[**IngestOut**](IngestOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostResearchGrants

> GrantOut PostResearchGrants(ctx).GrantRequest(grantRequest).Execute()

Records the SEPARATE authorization an upload never implies: a record's visibility (private, org or public) and, for a run, its training and commons-publication consent.



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
	grantRequest := *openapiclient.NewGrantRequest() // GrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.PostResearchGrants(context.Background()).GrantRequest(grantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.PostResearchGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResearchGrants`: GrantOut
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.PostResearchGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostResearchGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantRequest** | [**GrantRequest**](GrantRequest.md) |  | 

### Return type

[**GrantOut**](GrantOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

