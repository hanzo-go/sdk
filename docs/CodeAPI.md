# \CodeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCodeAsk**](CodeAPI.md#GetCodeAsk) | **Get** /v1/code/ask | Answers a question about the caller org&#39;s code with a CITED answer: retrieval packs grounding context, then the synthesizer writes the answer over exactly those spans, which come back alongside it.
[**GetCodeFile**](CodeAPI.md#GetCodeFile) | **Get** /v1/code/file | Returns the INDEXED content of one file — read_file over the chunks the search tiers hold, for pulling up code an agent just found.
[**GetCodeSearch**](CodeAPI.md#GetCodeSearch) | **Get** /v1/code/search | Finds code in the caller org&#39;s index across three orthogonal retrieval tiers fused by reciprocal-rank fusion: lexical (FTS5 trigram over code-tokenized text), symbolic (real definition and reference edges), and semantic (embedding cosine over AST-boundary chunks).
[**GetCodeTree**](CodeAPI.md#GetCodeTree) | **Get** /v1/code/tree | Returns one repository&#39;s file structure with a per-file symbol count — get_repo_structure over the org&#39;s own index, with no git checkout involved.
[**PostCodeAsk**](CodeAPI.md#PostCodeAsk) | **Post** /v1/code/ask | Is askGet with the question in the request BODY, for a question too long or too awkward to put in a URL.
[**PostCodeContext**](CodeAPI.md#PostCodeContext) | **Post** /v1/code/context | Packs the most relevant code for a query into a token budget — THE primitive for a coding agent that has to decide what to put in a prompt.
[**PostCodeIndex**](CodeAPI.md#PostCodeIndex) | **Post** /v1/code/index | (re)indexes a repository for the caller&#39;s org, incrementally: files whose content hash is unchanged are skipped, so re-sending a whole tree is cheap.
[**PostCodeLspComplete**](CodeAPI.md#PostCodeLspComplete) | **Post** /v1/code/lsp/complete | Offers the candidates a language server has at a position, typed and resolved through the repository&#39;s dependencies rather than guessed from text.
[**PostCodeLspDiagnostics**](CodeAPI.md#PostCodeLspDiagnostics) | **Post** /v1/code/lsp/diagnostics | Reports every problem the language server finds in one file — compile errors, type errors and lints, each with its span and its severity (1 error, 2 warning, 3 information, 4 hint).
[**PostCodeLspHover**](CodeAPI.md#PostCodeLspHover) | **Post** /v1/code/lsp/hover | Renders the type and documentation of the symbol at a position, as the language server itself renders it.
[**PostCodeLspLocate**](CodeAPI.md#PostCodeLspLocate) | **Post** /v1/code/lsp/locate | Finds where a symbol lives: its definition, its references, its type or its implementations, chosen by relation (definition, reference, type, implementation — empty means definition).
[**PostCodeLspSymbols**](CodeAPI.md#PostCodeLspSymbols) | **Post** /v1/code/lsp/symbols | Outlines one file: every declaration in it, with its kind and its span.



## GetCodeAsk

> AskAnswer GetCodeAsk(ctx).Q(q).Repo(repo).Execute()

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
	resp, r, err := apiClient.CodeAPI.GetCodeAsk(context.Background()).Q(q).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.GetCodeAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCodeAsk`: AskAnswer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.GetCodeAsk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q is the question to answer. Required, max 4000 bytes. | 
 **repo** | **string** | Repo narrows retrieval to one repository. Empty searches every repo the org has indexed. | 

### Return type

[**AskAnswer**](AskAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCodeFile

> FileContent GetCodeFile(ctx).Path(path).Repo(repo).Execute()

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
	resp, r, err := apiClient.CodeAPI.GetCodeFile(context.Background()).Path(path).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.GetCodeFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCodeFile`: FileContent
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.GetCodeFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path is the file&#39;s repo-relative path. Required. | 
 **repo** | **string** | Repo is the repository the file belongs to. REQUIRED. | 

### Return type

[**FileContent**](FileContent.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCodeSearch

> SearchResults GetCodeSearch(ctx).Q(q).Type_(type_).Repo(repo).Limit(limit).Execute()

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
	resp, r, err := apiClient.CodeAPI.GetCodeSearch(context.Background()).Q(q).Type_(type_).Repo(repo).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.GetCodeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCodeSearch`: SearchResults
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.GetCodeSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q is the search query. Required, max 4000 bytes. For type&#x3D;regex it is a regular expression; for type&#x3D;symbol it is a symbol name. | 
 **type_** | **string** | Type selects the retrieval tier: \&quot;text\&quot; (FTS5 trigram), \&quot;regex\&quot;, \&quot;symbol\&quot; (definitions), \&quot;semantic\&quot; (embeddings) or \&quot;hybrid\&quot;. Anything else — including empty — reads as hybrid. | 
 **repo** | **string** | Repo narrows to one repository. Empty searches every repo the org has indexed. | 
 **limit** | **int32** | Limit caps how many spans come back: default 20, maximum 100. A value that is not a positive integer reads as the default. | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCodeTree

> RepoTree GetCodeTree(ctx).Repo(repo).Execute()

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
	resp, r, err := apiClient.CodeAPI.GetCodeTree(context.Background()).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.GetCodeTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCodeTree`: RepoTree
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.GetCodeTree`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | Repo is the repository to walk. REQUIRED — a tree is repo-scoped. | 

### Return type

[**RepoTree**](RepoTree.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCodeAsk

> AskAnswer PostCodeAsk(ctx).AskPostIn(askPostIn).Execute()

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
	askPostIn := *openapiclient.NewAskPostIn() // AskPostIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.PostCodeAsk(context.Background()).AskPostIn(askPostIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.PostCodeAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCodeAsk`: AskAnswer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.PostCodeAsk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodeAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **askPostIn** | [**AskPostIn**](AskPostIn.md) |  | 

### Return type

[**AskAnswer**](AskAnswer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCodeContext

> ContextBundle PostCodeContext(ctx).ContextIn(contextIn).Execute()

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
	contextIn := *openapiclient.NewContextIn() // ContextIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.PostCodeContext(context.Background()).ContextIn(contextIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.PostCodeContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCodeContext`: ContextBundle
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.PostCodeContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodeContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextIn** | [**ContextIn**](ContextIn.md) |  | 

### Return type

[**ContextBundle**](ContextBundle.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCodeIndex

> IndexResult PostCodeIndex(ctx).IndexIn(indexIn).Execute()

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
	indexIn := *openapiclient.NewIndexIn() // IndexIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.PostCodeIndex(context.Background()).IndexIn(indexIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.PostCodeIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCodeIndex`: IndexResult
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.PostCodeIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexIn** | [**IndexIn**](IndexIn.md) |  | 

### Return type

[**IndexResult**](IndexResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCodeLspComplete

> Answer PostCodeLspComplete(ctx).Query(query).Execute()

Offers the candidates a language server has at a position, typed and resolved through the repository's dependencies rather than guessed from text.



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.PostCodeLspComplete(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.PostCodeLspComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCodeLspComplete`: Answer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.PostCodeLspComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodeLspCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 

### Return type

[**Answer**](Answer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCodeLspDiagnostics

> Answer PostCodeLspDiagnostics(ctx).Query(query).Execute()

Reports every problem the language server finds in one file — compile errors, type errors and lints, each with its span and its severity (1 error, 2 warning, 3 information, 4 hint).



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.PostCodeLspDiagnostics(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.PostCodeLspDiagnostics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCodeLspDiagnostics`: Answer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.PostCodeLspDiagnostics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodeLspDiagnosticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 

### Return type

[**Answer**](Answer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCodeLspHover

> Answer PostCodeLspHover(ctx).Query(query).Execute()

Renders the type and documentation of the symbol at a position, as the language server itself renders it.



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.PostCodeLspHover(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.PostCodeLspHover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCodeLspHover`: Answer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.PostCodeLspHover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodeLspHoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 

### Return type

[**Answer**](Answer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCodeLspLocate

> Answer PostCodeLspLocate(ctx).Query(query).Execute()

Finds where a symbol lives: its definition, its references, its type or its implementations, chosen by relation (definition, reference, type, implementation — empty means definition).



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.PostCodeLspLocate(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.PostCodeLspLocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCodeLspLocate`: Answer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.PostCodeLspLocate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodeLspLocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 

### Return type

[**Answer**](Answer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCodeLspSymbols

> Answer PostCodeLspSymbols(ctx).Query(query).Execute()

Outlines one file: every declaration in it, with its kind and its span.



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeAPI.PostCodeLspSymbols(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeAPI.PostCodeLspSymbols``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCodeLspSymbols`: Answer
	fmt.Fprintf(os.Stdout, "Response from `CodeAPI.PostCodeLspSymbols`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCodeLspSymbolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 

### Return type

[**Answer**](Answer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

