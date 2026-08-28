# \GraphAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphAssert**](GraphAPI.md#GraphAssert) | **Post** /v1/graph | Assert what is true of an entity
[**GraphNeighbors**](GraphAPI.md#GraphNeighbors) | **Post** /v1/graph/neighbors | Walk the edges from a seed set, bounded
[**GraphRead**](GraphAPI.md#GraphRead) | **Get** /v1/graph | Read the assertions this organization has recorded
[**GraphResolve**](GraphAPI.md#GraphResolve) | **Post** /v1/graph/resolve | What is in force about an entity as of an instant, and what disagreed
[**GraphSearch**](GraphAPI.md#GraphSearch) | **Get** /v1/graph/search | Find assertions by their text rather than by an entity key
[**GraphVocabulary**](GraphAPI.md#GraphVocabulary) | **Get** /v1/graph/vocabulary | The relations in use, and the rule that resolves a conflict
[**PostGraphGraphql**](GraphAPI.md#PostGraphGraphql) | **Post** /v1/graph/graphql | Ask the graph in one request, traversing.



## GraphAssert

> GraphAssertOut GraphAssert(ctx).GraphAssertIn(graphAssertIn).Execute()

Assert what is true of an entity

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
	graphAssertIn := *openapiclient.NewGraphAssertIn() // GraphAssertIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphAssert(context.Background()).GraphAssertIn(graphAssertIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphAssert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphAssert`: GraphAssertOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphAssert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphAssertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphAssertIn** | [**GraphAssertIn**](GraphAssertIn.md) |  | 

### Return type

[**GraphAssertOut**](GraphAssertOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphNeighbors

> GraphNeighborsOut GraphNeighbors(ctx).GraphNeighborsIn(graphNeighborsIn).Execute()

Walk the edges from a seed set, bounded

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
	graphNeighborsIn := *openapiclient.NewGraphNeighborsIn() // GraphNeighborsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphNeighbors(context.Background()).GraphNeighborsIn(graphNeighborsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphNeighbors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphNeighbors`: GraphNeighborsOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphNeighbors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphNeighborsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphNeighborsIn** | [**GraphNeighborsIn**](GraphNeighborsIn.md) |  | 

### Return type

[**GraphNeighborsOut**](GraphNeighborsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphRead

> GraphReadOut GraphRead(ctx).Entity(entity).Relation(relation).Value(value).AsOf(asOf).Limit(limit).Execute()

Read the assertions this organization has recorded

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
	entity := "entity_example" // string | Entity narrows to what was asserted ABOUT one entity. Absent matches every entity. (optional)
	relation := "relation_example" // string | Relation narrows to one relation. Absent matches every relation. (optional)
	value := "value_example" // string | Value narrows to assertions pointing AT one value, which is how the edges into an entity are read. (optional)
	asOf := "asOf_example" // string | AsOf bounds the read to what was knowable at an instant, RFC 3339. Absent reads everything this plane holds. (optional)
	limit := int32(56) // int32 | Limit caps how many assertions come back. Absent, zero, or anything above the walk ceiling is the ceiling. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphRead(context.Background()).Entity(entity).Relation(relation).Value(value).AsOf(asOf).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphRead`: GraphReadOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string** | Entity narrows to what was asserted ABOUT one entity. Absent matches every entity. | 
 **relation** | **string** | Relation narrows to one relation. Absent matches every relation. | 
 **value** | **string** | Value narrows to assertions pointing AT one value, which is how the edges into an entity are read. | 
 **asOf** | **string** | AsOf bounds the read to what was knowable at an instant, RFC 3339. Absent reads everything this plane holds. | 
 **limit** | **int32** | Limit caps how many assertions come back. Absent, zero, or anything above the walk ceiling is the ceiling. | 

### Return type

[**GraphReadOut**](GraphReadOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphResolve

> GraphResolveOut GraphResolve(ctx).GraphResolveIn(graphResolveIn).Execute()

What is in force about an entity as of an instant, and what disagreed

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
	graphResolveIn := *openapiclient.NewGraphResolveIn() // GraphResolveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphResolve(context.Background()).GraphResolveIn(graphResolveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphResolve`: GraphResolveOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphResolve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphResolveIn** | [**GraphResolveIn**](GraphResolveIn.md) |  | 

### Return type

[**GraphResolveOut**](GraphResolveOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphSearch

> GraphReadOut GraphSearch(ctx).Q(q).Relation(relation).AsOf(asOf).Limit(limit).Execute()

Find assertions by their text rather than by an entity key



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
	q := "q_example" // string | Q is what to look for: words, matched as prefixes, all of them required. Punctuation is text here rather than syntax, so an entity key searches as itself. (optional)
	relation := "relation_example" // string | Relation narrows to one relation. Absent matches every relation. (optional)
	asOf := "asOf_example" // string | AsOf bounds the search to what was knowable at an instant, RFC 3339. Absent searches everything this plane holds. (optional)
	limit := int32(56) // int32 | Limit caps how many assertions come back. Absent, zero, or anything above the walk ceiling is the ceiling. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphSearch(context.Background()).Q(q).Relation(relation).AsOf(asOf).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphSearch`: GraphReadOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q is what to look for: words, matched as prefixes, all of them required. Punctuation is text here rather than syntax, so an entity key searches as itself. | 
 **relation** | **string** | Relation narrows to one relation. Absent matches every relation. | 
 **asOf** | **string** | AsOf bounds the search to what was knowable at an instant, RFC 3339. Absent searches everything this plane holds. | 
 **limit** | **int32** | Limit caps how many assertions come back. Absent, zero, or anything above the walk ceiling is the ceiling. | 

### Return type

[**GraphReadOut**](GraphReadOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphVocabulary

> GraphVocabularyOut GraphVocabulary(ctx).Execute()

The relations in use, and the rule that resolves a conflict

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
	resp, r, err := apiClient.GraphAPI.GraphVocabulary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphVocabulary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphVocabulary`: GraphVocabularyOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphVocabulary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGraphVocabularyRequest struct via the builder pattern


### Return type

[**GraphVocabularyOut**](GraphVocabularyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGraphGraphql

> GraphQLOut PostGraphGraphql(ctx).GraphQLIn(graphQLIn).Execute()

Ask the graph in one request, traversing.



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
	graphQLIn := *openapiclient.NewGraphQLIn() // GraphQLIn |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.PostGraphGraphql(context.Background()).GraphQLIn(graphQLIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.PostGraphGraphql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGraphGraphql`: GraphQLOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.PostGraphGraphql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGraphGraphqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLIn** | [**GraphQLIn**](GraphQLIn.md) |  | 

### Return type

[**GraphQLOut**](GraphQLOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

