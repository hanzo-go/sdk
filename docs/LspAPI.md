# \LspAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostLspComplete**](LspAPI.md#PostLspComplete) | **Post** /v1/lsp/complete | Offers the candidates a language server has at a position, typed and resolved through the repository&#39;s dependencies rather than guessed from text.
[**PostLspDiagnostics**](LspAPI.md#PostLspDiagnostics) | **Post** /v1/lsp/diagnostics | Reports every problem the language server finds in one file — compile errors, type errors and lints, each with its span and its severity (1 error, 2 warning, 3 information, 4 hint).
[**PostLspHover**](LspAPI.md#PostLspHover) | **Post** /v1/lsp/hover | Renders the type and documentation of the symbol at a position, as the language server itself renders it.
[**PostLspLocate**](LspAPI.md#PostLspLocate) | **Post** /v1/lsp/locate | Finds where a symbol lives: its definition, its references, its type or its implementations, chosen by relation (definition, reference, type, implementation — empty means definition).
[**PostLspSymbols**](LspAPI.md#PostLspSymbols) | **Post** /v1/lsp/symbols | Outlines one file: every declaration in it, with its kind and its span.



## PostLspComplete

> Answer PostLspComplete(ctx).Query(query).Execute()

Offers the candidates a language server has at a position, typed and resolved through the repository's dependencies rather than guessed from text.



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LspAPI.PostLspComplete(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LspAPI.PostLspComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLspComplete`: Answer
	fmt.Fprintf(os.Stdout, "Response from `LspAPI.PostLspComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLspCompleteRequest struct via the builder pattern


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


## PostLspDiagnostics

> Answer PostLspDiagnostics(ctx).Query(query).Execute()

Reports every problem the language server finds in one file — compile errors, type errors and lints, each with its span and its severity (1 error, 2 warning, 3 information, 4 hint).



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LspAPI.PostLspDiagnostics(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LspAPI.PostLspDiagnostics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLspDiagnostics`: Answer
	fmt.Fprintf(os.Stdout, "Response from `LspAPI.PostLspDiagnostics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLspDiagnosticsRequest struct via the builder pattern


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


## PostLspHover

> Answer PostLspHover(ctx).Query(query).Execute()

Renders the type and documentation of the symbol at a position, as the language server itself renders it.



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LspAPI.PostLspHover(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LspAPI.PostLspHover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLspHover`: Answer
	fmt.Fprintf(os.Stdout, "Response from `LspAPI.PostLspHover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLspHoverRequest struct via the builder pattern


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


## PostLspLocate

> Answer PostLspLocate(ctx).Query(query).Execute()

Finds where a symbol lives: its definition, its references, its type or its implementations, chosen by relation (definition, reference, type, implementation — empty means definition).



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LspAPI.PostLspLocate(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LspAPI.PostLspLocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLspLocate`: Answer
	fmt.Fprintf(os.Stdout, "Response from `LspAPI.PostLspLocate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLspLocateRequest struct via the builder pattern


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


## PostLspSymbols

> Answer PostLspSymbols(ctx).Query(query).Execute()

Outlines one file: every declaration in it, with its kind and its span.



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
	query := *openapiclient.NewQuery() // Query | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LspAPI.PostLspSymbols(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LspAPI.PostLspSymbols``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLspSymbols`: Answer
	fmt.Fprintf(os.Stdout, "Response from `LspAPI.PostLspSymbols`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLspSymbolsRequest struct via the builder pattern


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

