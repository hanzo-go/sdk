# \ResearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResearchGetResearchArtifactBlob**](ResearchAPI.md#ResearchGetResearchArtifactBlob) | **Get** /v1/research/artifacts/{sha256} | Retrieve an artifact&#39;s bytes by content hash (org-scoped)
[**ResearchGetResearchTotals**](ResearchAPI.md#ResearchGetResearchTotals) | **Get** /v1/research/totals | Headline aggregate (canonical + retained) plus per-kind breakdown
[**ResearchIngestResearch**](ResearchAPI.md#ResearchIngestResearch) | **Post** /v1/research/experiments | Ingest a batch of experiments and attempts (idempotent by content)
[**ResearchListExperiments**](ResearchAPI.md#ResearchListExperiments) | **Get** /v1/research/experiments | List canonical experiments (latest answered version per stable id)
[**ResearchListResearchArtifacts**](ResearchAPI.md#ResearchListResearchArtifacts) | **Get** /v1/research/artifacts | The chronological diary feed (newest-first)
[**ResearchListResearchProjects**](ResearchAPI.md#ResearchListResearchProjects) | **Get** /v1/research/projects | Every project in the org with its real totals (canonical + retained)
[**ResearchRecordResearchArtifact**](ResearchAPI.md#ResearchRecordResearchArtifact) | **Post** /v1/research/artifacts | Record a diary artifact (idempotent by sha256 content hash)
[**ResearchSetResearchGrant**](ResearchAPI.md#ResearchSetResearchGrant) | **Post** /v1/research/grants | Set visibility/consent for a run (by id) or an artifact (by sha256)



## ResearchGetResearchArtifactBlob

> *os.File ResearchGetResearchArtifactBlob(ctx, sha256).Project(project).Execute()

Retrieve an artifact's bytes by content hash (org-scoped)

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
	sha256 := "sha256_example" // string | The artifact content hash.
	project := "project_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.ResearchGetResearchArtifactBlob(context.Background(), sha256).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.ResearchGetResearchArtifactBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchGetResearchArtifactBlob`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.ResearchGetResearchArtifactBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sha256** | **string** | The artifact content hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResearchGetResearchArtifactBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchGetResearchTotals

> ResearchTotals ResearchGetResearchTotals(ctx).Project(project).Execute()

Headline aggregate (canonical + retained) plus per-kind breakdown

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
	project := "project_example" // string | Scope the aggregate to one project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.ResearchGetResearchTotals(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.ResearchGetResearchTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchGetResearchTotals`: ResearchTotals
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.ResearchGetResearchTotals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResearchGetResearchTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Scope the aggregate to one project. | 

### Return type

[**ResearchTotals**](ResearchTotals.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchIngestResearch

> ResearchIngestResult ResearchIngestResearch(ctx).ResearchIngestRequest(researchIngestRequest).Execute()

Ingest a batch of experiments and attempts (idempotent by content)



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
	researchIngestRequest := *openapiclient.NewResearchIngestRequest() // ResearchIngestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.ResearchIngestResearch(context.Background()).ResearchIngestRequest(researchIngestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.ResearchIngestResearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchIngestResearch`: ResearchIngestResult
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.ResearchIngestResearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResearchIngestResearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **researchIngestRequest** | [**ResearchIngestRequest**](ResearchIngestRequest.md) |  | 

### Return type

[**ResearchIngestResult**](ResearchIngestResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchListExperiments

> ResearchListExperiments200Response ResearchListExperiments(ctx).Project(project).Kind(kind).Execute()

List canonical experiments (latest answered version per stable id)

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
	project := "project_example" // string | Narrow to one project; omit for the org's whole cross-project view. (optional)
	kind := "kind_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.ResearchListExperiments(context.Background()).Project(project).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.ResearchListExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchListExperiments`: ResearchListExperiments200Response
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.ResearchListExperiments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResearchListExperimentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Narrow to one project; omit for the org&#39;s whole cross-project view. | 
 **kind** | **string** |  | 

### Return type

[**ResearchListExperiments200Response**](ResearchListExperiments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchListResearchArtifacts

> ResearchListResearchArtifacts200Response ResearchListResearchArtifacts(ctx).Run(run).Project(project).Since(since).Execute()

The chronological diary feed (newest-first)

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
	run := "run_example" // string | Filter to one run/snapshot id. (optional)
	project := "project_example" // string | Defaults to the caller's project scope. (optional)
	since := int64(789) // int64 | Unix-seconds lower bound on ts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.ResearchListResearchArtifacts(context.Background()).Run(run).Project(project).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.ResearchListResearchArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchListResearchArtifacts`: ResearchListResearchArtifacts200Response
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.ResearchListResearchArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResearchListResearchArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **run** | **string** | Filter to one run/snapshot id. | 
 **project** | **string** | Defaults to the caller&#39;s project scope. | 
 **since** | **int64** | Unix-seconds lower bound on ts. | 

### Return type

[**ResearchListResearchArtifacts200Response**](ResearchListResearchArtifacts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchListResearchProjects

> ResearchListResearchProjects200Response ResearchListResearchProjects(ctx).Execute()

Every project in the org with its real totals (canonical + retained)

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
	resp, r, err := apiClient.ResearchAPI.ResearchListResearchProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.ResearchListResearchProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchListResearchProjects`: ResearchListResearchProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.ResearchListResearchProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResearchListResearchProjectsRequest struct via the builder pattern


### Return type

[**ResearchListResearchProjects200Response**](ResearchListResearchProjects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchRecordResearchArtifact

> ResearchRecordResearchArtifact200Response ResearchRecordResearchArtifact(ctx).ResearchArtifact(researchArtifact).Execute()

Record a diary artifact (idempotent by sha256 content hash)



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
	researchArtifact := *openapiclient.NewResearchArtifact(string(123), "Kind_example") // ResearchArtifact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.ResearchRecordResearchArtifact(context.Background()).ResearchArtifact(researchArtifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.ResearchRecordResearchArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchRecordResearchArtifact`: ResearchRecordResearchArtifact200Response
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.ResearchRecordResearchArtifact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResearchRecordResearchArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **researchArtifact** | [**ResearchArtifact**](ResearchArtifact.md) |  | 

### Return type

[**ResearchRecordResearchArtifact200Response**](ResearchRecordResearchArtifact200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchSetResearchGrant

> ResearchSetResearchGrant200Response ResearchSetResearchGrant(ctx).ResearchGrantRequest(researchGrantRequest).Execute()

Set visibility/consent for a run (by id) or an artifact (by sha256)



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
	researchGrantRequest := *openapiclient.NewResearchGrantRequest() // ResearchGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResearchAPI.ResearchSetResearchGrant(context.Background()).ResearchGrantRequest(researchGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResearchAPI.ResearchSetResearchGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchSetResearchGrant`: ResearchSetResearchGrant200Response
	fmt.Fprintf(os.Stdout, "Response from `ResearchAPI.ResearchSetResearchGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResearchSetResearchGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **researchGrantRequest** | [**ResearchGrantRequest**](ResearchGrantRequest.md) |  | 

### Return type

[**ResearchSetResearchGrant200Response**](ResearchSetResearchGrant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

