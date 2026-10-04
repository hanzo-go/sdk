# \GraphAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphAnswer**](GraphAPI.md#GraphAnswer) | **Post** /v1/graph/answer | Answers a question from the whole graph and cites the assertions it rests on.
[**GraphAssert**](GraphAPI.md#GraphAssert) | **Post** /v1/graph | Records a batch of assertions and counts what became of each.
[**GraphCommunities**](GraphAPI.md#GraphCommunities) | **Post** /v1/graph/communities | Partitions the edge graph into sets of entities more densely connected to each other than to the rest.
[**GraphDerive**](GraphAPI.md#GraphDerive) | **Post** /v1/graph/derive | Concludes what the organization&#39;s rules derive from the graph, with a proof for every conclusion: the rule, and each support as the assertion in force it rests on or as the derived atom it is.
[**GraphDiff**](GraphAPI.md#GraphDiff) | **Post** /v1/graph/diff | Reports what came into force, was superseded and was retracted between two points.
[**GraphErase**](GraphAPI.md#GraphErase) | **Post** /v1/graph/erase | Removes every assertion that names an entity and returns a receipt.
[**GraphExtract**](GraphAPI.md#GraphExtract) | **Post** /v1/graph/extract | Reads the relations a source states and returns them, recording nothing.
[**GraphIngest**](GraphAPI.md#GraphIngest) | **Post** /v1/graph/ingest | Reads a source and records what it states, through the same admission as assert.
[**GraphNeighbors**](GraphAPI.md#GraphNeighbors) | **Post** /v1/graph/neighbors | Walks the in-force edges from a set of seeds and lists every entity reached, bounded.
[**GraphPath**](GraphAPI.md#GraphPath) | **Post** /v1/graph/path | Finds the shortest chain of in-force edges from one entity to another.
[**GraphRead**](GraphAPI.md#GraphRead) | **Get** /v1/graph | Lists the assertions recorded, every version, oldest first.
[**GraphResolve**](GraphAPI.md#GraphResolve) | **Post** /v1/graph/resolve | Answers what holds for one entity and relation at a point, winner first, with every weaker account that disagreed.
[**GraphSearch**](GraphAPI.md#GraphSearch) | **Get** /v1/graph/search | Finds assertions by their words, where read finds them by their keys, best match first.
[**GraphVocabulary**](GraphAPI.md#GraphVocabulary) | **Get** /v1/graph/vocabulary | Lists the relations in use, the schema declared for them and the rule that settles a conflict.
[**PostGraphGraphql**](GraphAPI.md#PostGraphGraphql) | **Post** /v1/graph/graphql | Ask the graph in one request, traversing.



## GraphAnswer

> GraphGraphAnswerOut GraphAnswer(ctx).GraphGraphAnswerIn(graphGraphAnswerIn).Execute()

Answers a question from the whole graph and cites the assertions it rests on.



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
	graphGraphAnswerIn := *openapiclient.NewGraphGraphAnswerIn("Question_example") // GraphGraphAnswerIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphAnswer(context.Background()).GraphGraphAnswerIn(graphGraphAnswerIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphAnswer`: GraphGraphAnswerOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphAnswerIn** | [**GraphGraphAnswerIn**](GraphGraphAnswerIn.md) |  | 

### Return type

[**GraphGraphAnswerOut**](GraphGraphAnswerOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphAssert

> GraphGraphAssertOut GraphAssert(ctx).GraphGraphAssertIn(graphGraphAssertIn).Execute()

Records a batch of assertions and counts what became of each.



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
	graphGraphAssertIn := *openapiclient.NewGraphGraphAssertIn([]openapiclient.GraphGraphFact{*openapiclient.NewGraphGraphFact("At_example", "Entity_example", "Relation_example", "Source_example")}) // GraphGraphAssertIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphAssert(context.Background()).GraphGraphAssertIn(graphGraphAssertIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphAssert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphAssert`: GraphGraphAssertOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphAssert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphAssertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphAssertIn** | [**GraphGraphAssertIn**](GraphGraphAssertIn.md) |  | 

### Return type

[**GraphGraphAssertOut**](GraphGraphAssertOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphCommunities

> GraphGraphCommunitiesOut GraphCommunities(ctx).GraphGraphCommunitiesIn(graphGraphCommunitiesIn).Execute()

Partitions the edge graph into sets of entities more densely connected to each other than to the rest.



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
	graphGraphCommunitiesIn := *openapiclient.NewGraphGraphCommunitiesIn() // GraphGraphCommunitiesIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphCommunities(context.Background()).GraphGraphCommunitiesIn(graphGraphCommunitiesIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphCommunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphCommunities`: GraphGraphCommunitiesOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphCommunities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphCommunitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphCommunitiesIn** | [**GraphGraphCommunitiesIn**](GraphGraphCommunitiesIn.md) |  | 

### Return type

[**GraphGraphCommunitiesOut**](GraphGraphCommunitiesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphDerive

> GraphGraphDeriveOut GraphDerive(ctx).GraphGraphDeriveIn(graphGraphDeriveIn).Execute()

Concludes what the organization's rules derive from the graph, with a proof for every conclusion: the rule, and each support as the assertion in force it rests on or as the derived atom it is.



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
	graphGraphDeriveIn := *openapiclient.NewGraphGraphDeriveIn() // GraphGraphDeriveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphDerive(context.Background()).GraphGraphDeriveIn(graphGraphDeriveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphDerive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphDerive`: GraphGraphDeriveOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphDerive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphDeriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphDeriveIn** | [**GraphGraphDeriveIn**](GraphGraphDeriveIn.md) |  | 

### Return type

[**GraphGraphDeriveOut**](GraphGraphDeriveOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphDiff

> GraphGraphDiffOut GraphDiff(ctx).GraphGraphDiffIn(graphGraphDiffIn).Execute()

Reports what came into force, was superseded and was retracted between two points.



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
	graphGraphDiffIn := *openapiclient.NewGraphGraphDiffIn() // GraphGraphDiffIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphDiff(context.Background()).GraphGraphDiffIn(graphGraphDiffIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphDiff`: GraphGraphDiffOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphDiff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphDiffIn** | [**GraphGraphDiffIn**](GraphGraphDiffIn.md) |  | 

### Return type

[**GraphGraphDiffOut**](GraphGraphDiffOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphErase

> GraphGraphEraseOut GraphErase(ctx).GraphGraphEraseIn(graphGraphEraseIn).Execute()

Removes every assertion that names an entity and returns a receipt.



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
	graphGraphEraseIn := *openapiclient.NewGraphGraphEraseIn("Entity_example", "Reason_example") // GraphGraphEraseIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphErase(context.Background()).GraphGraphEraseIn(graphGraphEraseIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphErase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphErase`: GraphGraphEraseOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphErase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphEraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphEraseIn** | [**GraphGraphEraseIn**](GraphGraphEraseIn.md) |  | 

### Return type

[**GraphGraphEraseOut**](GraphGraphEraseOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphExtract

> GraphGraphExtractOut GraphExtract(ctx).GraphGraphSourceIn(graphGraphSourceIn).Execute()

Reads the relations a source states and returns them, recording nothing.



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
	graphGraphSourceIn := *openapiclient.NewGraphGraphSourceIn("Source_example", "Text_example") // GraphGraphSourceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphExtract(context.Background()).GraphGraphSourceIn(graphGraphSourceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphExtract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphExtract`: GraphGraphExtractOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphExtract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphExtractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphSourceIn** | [**GraphGraphSourceIn**](GraphGraphSourceIn.md) |  | 

### Return type

[**GraphGraphExtractOut**](GraphGraphExtractOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphIngest

> GraphGraphAssertOut GraphIngest(ctx).GraphGraphSourceIn(graphGraphSourceIn).Execute()

Reads a source and records what it states, through the same admission as assert.



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
	graphGraphSourceIn := *openapiclient.NewGraphGraphSourceIn("Source_example", "Text_example") // GraphGraphSourceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphIngest(context.Background()).GraphGraphSourceIn(graphGraphSourceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphIngest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphIngest`: GraphGraphAssertOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphIngest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphIngestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphSourceIn** | [**GraphGraphSourceIn**](GraphGraphSourceIn.md) |  | 

### Return type

[**GraphGraphAssertOut**](GraphGraphAssertOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphNeighbors

> GraphGraphNeighborsOut GraphNeighbors(ctx).GraphGraphNeighborsIn(graphGraphNeighborsIn).Execute()

Walks the in-force edges from a set of seeds and lists every entity reached, bounded.



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
	graphGraphNeighborsIn := *openapiclient.NewGraphGraphNeighborsIn([]string{"Seeds_example"}) // GraphGraphNeighborsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphNeighbors(context.Background()).GraphGraphNeighborsIn(graphGraphNeighborsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphNeighbors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphNeighbors`: GraphGraphNeighborsOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphNeighbors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphNeighborsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphNeighborsIn** | [**GraphGraphNeighborsIn**](GraphGraphNeighborsIn.md) |  | 

### Return type

[**GraphGraphNeighborsOut**](GraphGraphNeighborsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphPath

> GraphGraphPathOut GraphPath(ctx).GraphGraphPathIn(graphGraphPathIn).Execute()

Finds the shortest chain of in-force edges from one entity to another.



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
	graphGraphPathIn := *openapiclient.NewGraphGraphPathIn("From_example", "To_example") // GraphGraphPathIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphPath(context.Background()).GraphGraphPathIn(graphGraphPathIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphPath`: GraphGraphPathOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphPathIn** | [**GraphGraphPathIn**](GraphGraphPathIn.md) |  | 

### Return type

[**GraphGraphPathOut**](GraphGraphPathOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphRead

> GraphGraphReadOut GraphRead(ctx).Entity(entity).Relation(relation).Value(value).AsOf(asOf).AsKnown(asKnown).Limit(limit).Execute()

Lists the assertions recorded, every version, oldest first.



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
	entity := "acme/svc/api" // string | Entity narrows to what was asserted ABOUT one entity. Absent matches every entity. (optional)
	relation := "relation_example" // string | Relation narrows to one relation. Absent matches every relation. (optional)
	value := "value_example" // string | Value narrows to assertions pointing AT one value, which is how the edges into an entity are read. A property's scalar is matched byte for byte and an edge's value by its key, folded as every key is. (optional)
	asOf := "asOf_example" // string | AsOf bounds the read to statements begun by an instant of the world, RFC 3339. Absent reads every instant. (optional)
	asKnown := "2026-09-01T00:00:00Z" // string | AsKnown bounds the read to what this plane had heard by an instant, RFC 3339. Absent reads everything it holds. (optional)
	limit := int64(789) // int64 | Limit caps how many assertions come back. Absent, zero, or anything above the walk ceiling is the ceiling. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphRead(context.Background()).Entity(entity).Relation(relation).Value(value).AsOf(asOf).AsKnown(asKnown).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphRead`: GraphGraphReadOut
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
 **value** | **string** | Value narrows to assertions pointing AT one value, which is how the edges into an entity are read. A property&#39;s scalar is matched byte for byte and an edge&#39;s value by its key, folded as every key is. | 
 **asOf** | **string** | AsOf bounds the read to statements begun by an instant of the world, RFC 3339. Absent reads every instant. | 
 **asKnown** | **string** | AsKnown bounds the read to what this plane had heard by an instant, RFC 3339. Absent reads everything it holds. | 
 **limit** | **int64** | Limit caps how many assertions come back. Absent, zero, or anything above the walk ceiling is the ceiling. | 

### Return type

[**GraphGraphReadOut**](GraphGraphReadOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphResolve

> GraphGraphResolveOut GraphResolve(ctx).GraphGraphResolveIn(graphGraphResolveIn).Execute()

Answers what holds for one entity and relation at a point, winner first, with every weaker account that disagreed.



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
	graphGraphResolveIn := *openapiclient.NewGraphGraphResolveIn("Entity_example", "Relation_example") // GraphGraphResolveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphResolve(context.Background()).GraphGraphResolveIn(graphGraphResolveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphResolve`: GraphGraphResolveOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphResolve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphGraphResolveIn** | [**GraphGraphResolveIn**](GraphGraphResolveIn.md) |  | 

### Return type

[**GraphGraphResolveOut**](GraphGraphResolveOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphSearch

> GraphGraphReadOut GraphSearch(ctx).Q(q).Relation(relation).AsOf(asOf).AsKnown(asKnown).Limit(limit).Execute()

Finds assertions by their words, where read finds them by their keys, best match first.



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
	q := "billing owner" // string | Q is what to look for: words, matched as prefixes, all of them required. Punctuation is text here rather than syntax, so an entity key searches as itself.
	relation := "relation_example" // string | Relation narrows to one relation. Absent matches every relation. (optional)
	asOf := "asOf_example" // string | AsOf bounds the search to statements begun by an instant of the world, RFC 3339. Absent searches every instant. (optional)
	asKnown := "asKnown_example" // string | AsKnown bounds the search to what this plane had heard by an instant, RFC 3339. Absent searches everything it holds. (optional)
	limit := int64(789) // int64 | Limit caps how many assertions come back. Absent, zero, or anything above the walk ceiling is the ceiling. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.GraphSearch(context.Background()).Q(q).Relation(relation).AsOf(asOf).AsKnown(asKnown).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.GraphSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphSearch`: GraphGraphReadOut
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
 **asOf** | **string** | AsOf bounds the search to statements begun by an instant of the world, RFC 3339. Absent searches every instant. | 
 **asKnown** | **string** | AsKnown bounds the search to what this plane had heard by an instant, RFC 3339. Absent searches everything it holds. | 
 **limit** | **int64** | Limit caps how many assertions come back. Absent, zero, or anything above the walk ceiling is the ceiling. | 

### Return type

[**GraphGraphReadOut**](GraphGraphReadOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphVocabulary

> GraphGraphVocabularyOut GraphVocabulary(ctx).Execute()

Lists the relations in use, the schema declared for them and the rule that settles a conflict.



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
	// response from `GraphVocabulary`: GraphGraphVocabularyOut
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.GraphVocabulary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGraphVocabularyRequest struct via the builder pattern


### Return type

[**GraphGraphVocabularyOut**](GraphGraphVocabularyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

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

