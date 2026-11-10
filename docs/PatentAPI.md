# \PatentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPatentApprovals**](PatentAPI.md#GetPatentApprovals) | **Get** /v1/patent/approvals | Returns the decision trail — each person&#39;s current word on each disclosure.
[**GetPatentDisclosures**](PatentAPI.md#GetPatentDisclosures) | **Get** /v1/patent/disclosures | Returns the canonical disclosures, each carrying its current approvals.
[**GetPatentHealth**](PatentAPI.md#GetPatentHealth) | **Get** /v1/patent/health | Reports what this deployment actually guarantees about the records it keeps, in honest states.
[**GetPatentIdeas**](PatentAPI.md#GetPatentIdeas) | **Get** /v1/patent/ideas | Returns the caller org&#39;s canonical conception entries, oldest appended first, which is the order the log says they were thought.
[**GetPatentImplementations**](PatentAPI.md#GetPatentImplementations) | **Get** /v1/patent/implementations | Returns the canonical implementation records with their citations.
[**GetPatentProof**](PatentAPI.md#GetPatentProof) | **Get** /v1/patent/proof | Folds the caller&#39;s record set and answers the ordering question over it.
[**GetPatentTimeline**](PatentAPI.md#GetPatentTimeline) | **Get** /v1/patent/timeline | Returns every retained version of one entry, in append order.
[**PostPatentApprovals**](PatentAPI.md#PostPatentApprovals) | **Post** /v1/patent/approvals | Appends one person&#39;s decision about one disclosure.
[**PostPatentDisclosures**](PatentAPI.md#PostPatentDisclosures) | **Post** /v1/patent/disclosures | Appends the assembled document and the record ids it rests on.
[**PostPatentGrants**](PatentAPI.md#PostPatentGrants) | **Post** /v1/patent/grants | Makes the decision a recording never implied: whether an idea, implementation or disclosure stays private to its author&#39;s project, opens to the org, or goes public.
[**PostPatentIdeas**](PatentAPI.md#PostPatentIdeas) | **Post** /v1/patent/ideas | Appends conception entries.
[**PostPatentImplementations**](PatentAPI.md#PostPatentImplementations) | **Post** /v1/patent/implementations | Appends the machine&#39;s half of the record, and the idea citation is the load-bearing field: it is what lets a reader see that the conception came first.



## GetPatentApprovals

> PatentApprovalsOut GetPatentApprovals(ctx).Disclosure(disclosure).Project(project).Execute()

Returns the decision trail — each person's current word on each disclosure.



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
	disclosure := "disclosure_example" // string | Disclosure narrows to the decisions on one document. (optional)
	project := "project_example" // string | Project narrows to one project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.GetPatentApprovals(context.Background()).Disclosure(disclosure).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.GetPatentApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatentApprovals`: PatentApprovalsOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.GetPatentApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatentApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disclosure** | **string** | Disclosure narrows to the decisions on one document. | 
 **project** | **string** | Project narrows to one project. | 

### Return type

[**PatentApprovalsOut**](PatentApprovalsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatentDisclosures

> PatentDiscsOut GetPatentDisclosures(ctx).Id(id).Status(status).Project(project).Execute()

Returns the canonical disclosures, each carrying its current approvals.



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
	id := "id_example" // string | ID reads one disclosure. (optional)
	status := "status_example" // string | Status narrows to one status. (optional)
	project := "project_example" // string | Project narrows to one project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.GetPatentDisclosures(context.Background()).Id(id).Status(status).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.GetPatentDisclosures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatentDisclosures`: PatentDiscsOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.GetPatentDisclosures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatentDisclosuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID reads one disclosure. | 
 **status** | **string** | Status narrows to one status. | 
 **project** | **string** | Project narrows to one project. | 

### Return type

[**PatentDiscsOut**](PatentDiscsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatentHealth

> PatentPatentHealth GetPatentHealth(ctx).Execute()

Reports what this deployment actually guarantees about the records it keeps, in honest states.



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
	resp, r, err := apiClient.PatentAPI.GetPatentHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.GetPatentHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatentHealth`: PatentPatentHealth
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.GetPatentHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatentHealthRequest struct via the builder pattern


### Return type

[**PatentPatentHealth**](PatentPatentHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatentIdeas

> PatentIdeasOut GetPatentIdeas(ctx).Id(id).Kind(kind).Project(project).Execute()

Returns the caller org's canonical conception entries, oldest appended first, which is the order the log says they were thought.



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
	id := "id_example" // string | ID reads one entry, which is how a single idea is addressed — there is no second route for it. (optional)
	kind := "kind_example" // string | Kind narrows to one kind of conception entry. (optional)
	project := "project_example" // string | Project narrows to one project; empty reads the org's whole set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.GetPatentIdeas(context.Background()).Id(id).Kind(kind).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.GetPatentIdeas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatentIdeas`: PatentIdeasOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.GetPatentIdeas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatentIdeasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID reads one entry, which is how a single idea is addressed — there is no second route for it. | 
 **kind** | **string** | Kind narrows to one kind of conception entry. | 
 **project** | **string** | Project narrows to one project; empty reads the org&#39;s whole set. | 

### Return type

[**PatentIdeasOut**](PatentIdeasOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatentImplementations

> PatentImplsOut GetPatentImplementations(ctx).Id(id).Idea(idea).Project(project).Execute()

Returns the canonical implementation records with their citations.



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
	id := "id_example" // string | ID reads one implementation. (optional)
	idea := "idea_example" // string | Idea narrows to implementations citing one conception entry — which is how a correction finds every artifact built on the idea it invalidates. (optional)
	project := "project_example" // string | Project narrows to one project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.GetPatentImplementations(context.Background()).Id(id).Idea(idea).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.GetPatentImplementations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatentImplementations`: PatentImplsOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.GetPatentImplementations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatentImplementationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID reads one implementation. | 
 **idea** | **string** | Idea narrows to implementations citing one conception entry — which is how a correction finds every artifact built on the idea it invalidates. | 
 **project** | **string** | Project narrows to one project. | 

### Return type

[**PatentImplsOut**](PatentImplsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatentProof

> PatentPatentProof GetPatentProof(ctx).Disclosure(disclosure).Idea(idea).Project(project).Execute()

Folds the caller's record set and answers the ordering question over it.



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
	disclosure := "disclosure_example" // string | Disclosure folds the record set one document rests on, and reports on that set alone — which is what an attorney asks for. (optional)
	idea := "idea_example" // string | Idea narrows to the one entry and the artifacts citing it. (optional)
	project := "project_example" // string | Project narrows inside the org; empty reads the caller's scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.GetPatentProof(context.Background()).Disclosure(disclosure).Idea(idea).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.GetPatentProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatentProof`: PatentPatentProof
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.GetPatentProof`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatentProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disclosure** | **string** | Disclosure folds the record set one document rests on, and reports on that set alone — which is what an attorney asks for. | 
 **idea** | **string** | Idea narrows to the one entry and the artifacts citing it. | 
 **project** | **string** | Project narrows inside the org; empty reads the caller&#39;s scope. | 

### Return type

[**PatentPatentProof**](PatentPatentProof.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatentTimeline

> PatentPatentTimeline GetPatentTimeline(ctx).Id(id).Kind(kind).Project(project).Execute()

Returns every retained version of one entry, in append order.



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
	id := "id_example" // string | ID is the entry's stable id. Required. (optional)
	kind := "kind_example" // string | Kind is idea, implementation or disclosure — which log to read. (optional)
	project := "project_example" // string | Project narrows inside the org; empty is the caller's own scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.GetPatentTimeline(context.Background()).Id(id).Kind(kind).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.GetPatentTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatentTimeline`: PatentPatentTimeline
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.GetPatentTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatentTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID is the entry&#39;s stable id. Required. | 
 **kind** | **string** | Kind is idea, implementation or disclosure — which log to read. | 
 **project** | **string** | Project narrows inside the org; empty is the caller&#39;s own scope. | 

### Return type

[**PatentPatentTimeline**](PatentPatentTimeline.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatentApprovals

> PatentPatentWriteOut PostPatentApprovals(ctx).PatentApprovalBatch(patentApprovalBatch).Execute()

Appends one person's decision about one disclosure.



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
	patentApprovalBatch := *openapiclient.NewPatentApprovalBatch() // PatentApprovalBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.PostPatentApprovals(context.Background()).PatentApprovalBatch(patentApprovalBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.PostPatentApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatentApprovals`: PatentPatentWriteOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.PostPatentApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatentApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patentApprovalBatch** | [**PatentApprovalBatch**](PatentApprovalBatch.md) |  | 

### Return type

[**PatentPatentWriteOut**](PatentPatentWriteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatentDisclosures

> PatentPatentWriteOut PostPatentDisclosures(ctx).PatentDisclosureBatch(patentDisclosureBatch).Execute()

Appends the assembled document and the record ids it rests on.



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
	patentDisclosureBatch := *openapiclient.NewPatentDisclosureBatch() // PatentDisclosureBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.PostPatentDisclosures(context.Background()).PatentDisclosureBatch(patentDisclosureBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.PostPatentDisclosures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatentDisclosures`: PatentPatentWriteOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.PostPatentDisclosures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatentDisclosuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patentDisclosureBatch** | [**PatentDisclosureBatch**](PatentDisclosureBatch.md) |  | 

### Return type

[**PatentPatentWriteOut**](PatentPatentWriteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatentGrants

> PatentPatentGrantOut PostPatentGrants(ctx).PatentGrantIn(patentGrantIn).Execute()

Makes the decision a recording never implied: whether an idea, implementation or disclosure stays private to its author's project, opens to the org, or goes public.



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
	patentGrantIn := *openapiclient.NewPatentGrantIn() // PatentGrantIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.PostPatentGrants(context.Background()).PatentGrantIn(patentGrantIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.PostPatentGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatentGrants`: PatentPatentGrantOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.PostPatentGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatentGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patentGrantIn** | [**PatentGrantIn**](PatentGrantIn.md) |  | 

### Return type

[**PatentPatentGrantOut**](PatentPatentGrantOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatentIdeas

> PatentPatentWriteOut PostPatentIdeas(ctx).PatentIdeaBatch(patentIdeaBatch).Execute()

Appends conception entries.



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
	patentIdeaBatch := *openapiclient.NewPatentIdeaBatch() // PatentIdeaBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.PostPatentIdeas(context.Background()).PatentIdeaBatch(patentIdeaBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.PostPatentIdeas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatentIdeas`: PatentPatentWriteOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.PostPatentIdeas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatentIdeasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patentIdeaBatch** | [**PatentIdeaBatch**](PatentIdeaBatch.md) |  | 

### Return type

[**PatentPatentWriteOut**](PatentPatentWriteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatentImplementations

> PatentPatentWriteOut PostPatentImplementations(ctx).PatentImplBatch(patentImplBatch).Execute()

Appends the machine's half of the record, and the idea citation is the load-bearing field: it is what lets a reader see that the conception came first.



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
	patentImplBatch := *openapiclient.NewPatentImplBatch() // PatentImplBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatentAPI.PostPatentImplementations(context.Background()).PatentImplBatch(patentImplBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatentAPI.PostPatentImplementations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatentImplementations`: PatentPatentWriteOut
	fmt.Fprintf(os.Stdout, "Response from `PatentAPI.PostPatentImplementations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatentImplementationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patentImplBatch** | [**PatentImplBatch**](PatentImplBatch.md) |  | 

### Return type

[**PatentPatentWriteOut**](PatentPatentWriteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

