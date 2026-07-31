# \CodeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1CodeAsk**](CodeAPI.md#CloudGetV1CodeAsk) | **Get** /v1/code/ask | Answers a question about the caller org&#39;s code with a CITED answer: retrieval packs grounding context, then the synthesizer writes the answer over exactly those spans, which come back alongside it.
[**CloudGetV1CodeFile**](CodeAPI.md#CloudGetV1CodeFile) | **Get** /v1/code/file | Returns the INDEXED content of one file — read_file over the chunks the search tiers hold, for pulling up code an agent just found.
[**CloudGetV1CodeSearch**](CodeAPI.md#CloudGetV1CodeSearch) | **Get** /v1/code/search | Finds code in the caller org&#39;s index across three orthogonal retrieval tiers fused by reciprocal-rank fusion: lexical (FTS5 trigram over code-tokenized text), symbolic (real definition and reference edges), and semantic (embedding cosine over AST-boundary chunks).
[**CloudGetV1CodeTree**](CodeAPI.md#CloudGetV1CodeTree) | **Get** /v1/code/tree | Returns one repository&#39;s file structure with a per-file symbol count — get_repo_structure over the org&#39;s own index, with no git checkout involved.
[**CloudPostV1CodeAsk**](CodeAPI.md#CloudPostV1CodeAsk) | **Post** /v1/code/ask | Is askGet with the question in the request BODY, for a question too long or too awkward to put in a URL.
[**CloudPostV1CodeContext**](CodeAPI.md#CloudPostV1CodeContext) | **Post** /v1/code/context | Packs the most relevant code for a query into a token budget — THE primitive for a coding agent that has to decide what to put in a prompt.
[**CloudPostV1CodeIndex**](CodeAPI.md#CloudPostV1CodeIndex) | **Post** /v1/code/index | (re)indexes a repository for the caller&#39;s org, incrementally: files whose content hash is unchanged are skipped, so re-sending a whole tree is cheap.



## CloudGetV1CodeAsk

> CloudAskAnswer CloudGetV1CodeAsk(ctx).Q(q).Repo(repo).Execute()

Answers a question about the caller org's code with a CITED answer: retrieval packs grounding context, then the synthesizer writes the answer over exactly those spans, which come back alongside it.



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
	q := "where is the per-org SQLite file opened" // string | Q is the question to answer. Required, max 4000 bytes. (optional)
	repo := "cloud" // string | Repo narrows retrieval to one repository. Empty searches every repo the org has indexed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.CloudGetV1CodeAsk(context.Background()).Q(q).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.CloudGetV1CodeAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CodeAsk`: CloudAskAnswer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.CloudGetV1CodeAsk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CodeAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q is the question to answer. Required, max 4000 bytes. | 
 **repo** | **string** | Repo narrows retrieval to one repository. Empty searches every repo the org has indexed. | 

### Return type

[**CloudAskAnswer**](CloudAskAnswer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CodeFile

> CloudFileContent CloudGetV1CodeFile(ctx).Path(path).Repo(repo).Execute()

Returns the INDEXED content of one file — read_file over the chunks the search tiers hold, for pulling up code an agent just found.



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
	path := "apps/code/store.go" // string | Path is the file's repo-relative path. Required. (optional)
	repo := "cloud" // string | Repo is the repository the file belongs to. REQUIRED. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.CloudGetV1CodeFile(context.Background()).Path(path).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.CloudGetV1CodeFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CodeFile`: CloudFileContent
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.CloudGetV1CodeFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CodeFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path is the file&#39;s repo-relative path. Required. | 
 **repo** | **string** | Repo is the repository the file belongs to. REQUIRED. | 

### Return type

[**CloudFileContent**](CloudFileContent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CodeSearch

> CloudSearchResults CloudGetV1CodeSearch(ctx).Q(q).Type_(type_).Repo(repo).Limit(limit).Execute()

Finds code in the caller org's index across three orthogonal retrieval tiers fused by reciprocal-rank fusion: lexical (FTS5 trigram over code-tokenized text), symbolic (real definition and reference edges), and semantic (embedding cosine over AST-boundary chunks).



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
	q := "func openStore" // string | Q is the search query. Required, max 4000 bytes. For type=regex it is a regular expression; for type=symbol it is a symbol name. (optional)
	type_ := "hybrid" // string | Type selects the retrieval tier: \"text\" (FTS5 trigram), \"regex\", \"symbol\" (definitions), \"semantic\" (embeddings) or \"hybrid\". Anything else — including empty — reads as hybrid. (optional)
	repo := "cloud" // string | Repo narrows to one repository. Empty searches every repo the org has indexed. (optional)
	limit := int32(20) // int32 | Limit caps how many spans come back: default 20, maximum 100. A value that is not a positive integer reads as the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.CloudGetV1CodeSearch(context.Background()).Q(q).Type_(type_).Repo(repo).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.CloudGetV1CodeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CodeSearch`: CloudSearchResults
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.CloudGetV1CodeSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CodeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q is the search query. Required, max 4000 bytes. For type&#x3D;regex it is a regular expression; for type&#x3D;symbol it is a symbol name. | 
 **type_** | **string** | Type selects the retrieval tier: \&quot;text\&quot; (FTS5 trigram), \&quot;regex\&quot;, \&quot;symbol\&quot; (definitions), \&quot;semantic\&quot; (embeddings) or \&quot;hybrid\&quot;. Anything else — including empty — reads as hybrid. | 
 **repo** | **string** | Repo narrows to one repository. Empty searches every repo the org has indexed. | 
 **limit** | **int32** | Limit caps how many spans come back: default 20, maximum 100. A value that is not a positive integer reads as the default. | 

### Return type

[**CloudSearchResults**](CloudSearchResults.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CodeTree

> CloudRepoTree CloudGetV1CodeTree(ctx).Repo(repo).Execute()

Returns one repository's file structure with a per-file symbol count — get_repo_structure over the org's own index, with no git checkout involved.



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
	repo := "cloud" // string | Repo is the repository to walk. REQUIRED — a tree is repo-scoped. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.CloudGetV1CodeTree(context.Background()).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.CloudGetV1CodeTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CodeTree`: CloudRepoTree
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.CloudGetV1CodeTree`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CodeTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | Repo is the repository to walk. REQUIRED — a tree is repo-scoped. | 

### Return type

[**CloudRepoTree**](CloudRepoTree.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CodeAsk

> CloudAskAnswer CloudPostV1CodeAsk(ctx).CloudAskPostIn(cloudAskPostIn).Execute()

Is askGet with the question in the request BODY, for a question too long or too awkward to put in a URL.



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
	cloudAskPostIn := *openapiclient.NewCloudAskPostIn() // CloudAskPostIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.CloudPostV1CodeAsk(context.Background()).CloudAskPostIn(cloudAskPostIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.CloudPostV1CodeAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CodeAsk`: CloudAskAnswer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.CloudPostV1CodeAsk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CodeAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAskPostIn** | [**CloudAskPostIn**](CloudAskPostIn.md) |  | 

### Return type

[**CloudAskAnswer**](CloudAskAnswer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CodeContext

> CloudContextBundle CloudPostV1CodeContext(ctx).CloudContextIn(cloudContextIn).Execute()

Packs the most relevant code for a query into a token budget — THE primitive for a coding agent that has to decide what to put in a prompt.



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
	cloudContextIn := *openapiclient.NewCloudContextIn() // CloudContextIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.CloudPostV1CodeContext(context.Background()).CloudContextIn(cloudContextIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.CloudPostV1CodeContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CodeContext`: CloudContextBundle
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.CloudPostV1CodeContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CodeContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudContextIn** | [**CloudContextIn**](CloudContextIn.md) |  | 

### Return type

[**CloudContextBundle**](CloudContextBundle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CodeIndex

> CloudIndexResult CloudPostV1CodeIndex(ctx).CloudIndexIn(cloudIndexIn).Execute()

(re)indexes a repository for the caller's org, incrementally: files whose content hash is unchanged are skipped, so re-sending a whole tree is cheap.



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
	cloudIndexIn := *openapiclient.NewCloudIndexIn() // CloudIndexIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.CloudPostV1CodeIndex(context.Background()).CloudIndexIn(cloudIndexIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.CloudPostV1CodeIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CodeIndex`: CloudIndexResult
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.CloudPostV1CodeIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CodeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudIndexIn** | [**CloudIndexIn**](CloudIndexIn.md) |  | 

### Return type

[**CloudIndexResult**](CloudIndexResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

