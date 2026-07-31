# \SearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotSearchPersonas**](SearchAPI.md#BotSearchPersonas) | **Get** /v1/bot/search/personas | Lexical search for personas
[**BotSearchSkills**](SearchAPI.md#BotSearchSkills) | **Get** /v1/bot/search/skills | Hybrid vector + lexical search for skills
[**CloudDeleteV1SearchName**](SearchAPI.md#CloudDeleteV1SearchName) | **Delete** /v1/search/{name} | DropSearch deletes one search index from the shared backend and removes its metadata row.
[**CloudGetV1Search**](SearchAPI.md#CloudGetV1Search) | **Get** /v1/search | ListSearch lists the caller org&#39;s search indexes.
[**CloudGetV1SearchName**](SearchAPI.md#CloudGetV1SearchName) | **Get** /v1/search/{name} | GetSearch returns one search index&#39;s metadata.
[**CloudPostV1Search**](SearchAPI.md#CloudPostV1Search) | **Post** /v1/search | 
[**CommerceSearchNotes**](SearchAPI.md#CommerceSearchNotes) | **Post** /v1/commerce/search/note | Search notes
[**CommerceSearchOrders**](SearchAPI.md#CommerceSearchOrders) | **Get** /v1/commerce/search/order | Search orders
[**CommerceSearchUsers**](SearchAPI.md#CommerceSearchUsers) | **Get** /v1/commerce/search/user | Search users
[**SearchSearchGet**](SearchAPI.md#SearchSearchGet) | **Get** /v1/search/indexes/{indexUid}/search | Search documents (GET)
[**SearchSearchPost**](SearchAPI.md#SearchSearchPost) | **Post** /v1/search/indexes/{indexUid}/search | Search documents (POST)
[**VectorRecommendPoints**](SearchAPI.md#VectorRecommendPoints) | **Post** /v1/vector/collections/{collection_name}/points/recommend | Recommend points
[**VectorSearchBatch**](SearchAPI.md#VectorSearchBatch) | **Post** /v1/vector/collections/{collection_name}/points/search/batch | Batch search
[**VectorSearchPoints**](SearchAPI.md#VectorSearchPoints) | **Post** /v1/vector/collections/{collection_name}/points/search | Search points



## BotSearchPersonas

> BotSearchPersonas200Response BotSearchPersonas(ctx).Q(q).Limit(limit).Execute()

Lexical search for personas

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
	q := "q_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.BotSearchPersonas(context.Background()).Q(q).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.BotSearchPersonas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotSearchPersonas`: BotSearchPersonas200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.BotSearchPersonas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotSearchPersonasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**BotSearchPersonas200Response**](BotSearchPersonas200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotSearchSkills

> BotSearchPersonas200Response BotSearchSkills(ctx).Q(q).Limit(limit).Execute()

Hybrid vector + lexical search for skills

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
	q := "q_example" // string | Search query
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.BotSearchSkills(context.Background()).Q(q).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.BotSearchSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotSearchSkills`: BotSearchPersonas200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.BotSearchSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotSearchSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search query | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**BotSearchPersonas200Response**](BotSearchPersonas200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1SearchName

> CloudDeleteV1SearchName(ctx, name).Execute()

DropSearch deletes one search index from the shared backend and removes its metadata row.



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
	name := "products" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchAPI.CloudDeleteV1SearchName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CloudDeleteV1SearchName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1SearchNameRequest struct via the builder pattern


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


## CloudGetV1Search

> []CloudProvisionedSummary CloudGetV1Search(ctx).Execute()

ListSearch lists the caller org's search indexes.



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
	resp, r, err := apiClient.SearchAPI.CloudGetV1Search(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CloudGetV1Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Search`: []CloudProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CloudGetV1Search`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SearchRequest struct via the builder pattern


### Return type

[**[]CloudProvisionedSummary**](CloudProvisionedSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1SearchName

> CloudProvisionedResource CloudGetV1SearchName(ctx, name).Execute()

GetSearch returns one search index's metadata.



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
	name := "products" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CloudGetV1SearchName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CloudGetV1SearchName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1SearchName`: CloudProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CloudGetV1SearchName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SearchNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProvisionedResource**](CloudProvisionedResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Search

> CloudProvisionResult CloudPostV1Search(ctx).CloudProvisionRequest(cloudProvisionRequest).Execute()



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
	cloudProvisionRequest := *openapiclient.NewCloudProvisionRequest() // CloudProvisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CloudPostV1Search(context.Background()).CloudProvisionRequest(cloudProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CloudPostV1Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Search`: CloudProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CloudPostV1Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1SearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvisionRequest** | [**CloudProvisionRequest**](CloudProvisionRequest.md) |  | 

### Return type

[**CloudProvisionResult**](CloudProvisionResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSearchNotes

> []CommerceNote CommerceSearchNotes(ctx).CommerceSearchNotesRequest(commerceSearchNotesRequest).Execute()

Search notes

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
	commerceSearchNotesRequest := *openapiclient.NewCommerceSearchNotesRequest() // CommerceSearchNotesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CommerceSearchNotes(context.Background()).CommerceSearchNotesRequest(commerceSearchNotesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CommerceSearchNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSearchNotes`: []CommerceNote
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CommerceSearchNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSearchNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceSearchNotesRequest** | [**CommerceSearchNotesRequest**](CommerceSearchNotesRequest.md) |  | 

### Return type

[**[]CommerceNote**](CommerceNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSearchOrders

> []CommerceOrder CommerceSearchOrders(ctx).Q(q).Execute()

Search orders

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
	q := "q_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CommerceSearchOrders(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CommerceSearchOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSearchOrders`: []CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CommerceSearchOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSearchOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 

### Return type

[**[]CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSearchUsers

> []CommerceUser CommerceSearchUsers(ctx).Q(q).Execute()

Search users

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
	q := "q_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CommerceSearchUsers(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CommerceSearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSearchUsers`: []CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CommerceSearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 

### Return type

[**[]CommerceUser**](CommerceUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSearchGet

> SearchSearchResult SearchSearchGet(ctx, indexUid).Q(q).Offset(offset).Limit(limit).AttributesToRetrieve(attributesToRetrieve).AttributesToHighlight(attributesToHighlight).AttributesToCrop(attributesToCrop).CropLength(cropLength).Filter(filter).Sort(sort).Facets(facets).ShowMatchesPosition(showMatchesPosition).ShowRankingScore(showRankingScore).MatchingStrategy(matchingStrategy).Execute()

Search documents (GET)

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
	indexUid := "indexUid_example" // string | Unique index identifier
	q := "q_example" // string | Search query (optional)
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 20)
	attributesToRetrieve := "attributesToRetrieve_example" // string | Comma-separated list of attributes to return (optional)
	attributesToHighlight := "attributesToHighlight_example" // string |  (optional)
	attributesToCrop := "attributesToCrop_example" // string |  (optional)
	cropLength := int32(56) // int32 |  (optional) (default to 10)
	filter := "filter_example" // string | Filter expression (optional)
	sort := "sort_example" // string | Comma-separated sort rules (optional)
	facets := "facets_example" // string | Comma-separated facet attributes (optional)
	showMatchesPosition := true // bool |  (optional)
	showRankingScore := true // bool |  (optional)
	matchingStrategy := "matchingStrategy_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchSearchGet(context.Background(), indexUid).Q(q).Offset(offset).Limit(limit).AttributesToRetrieve(attributesToRetrieve).AttributesToHighlight(attributesToHighlight).AttributesToCrop(attributesToCrop).CropLength(cropLength).Filter(filter).Sort(sort).Facets(facets).ShowMatchesPosition(showMatchesPosition).ShowRankingScore(showRankingScore).MatchingStrategy(matchingStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSearchGet`: SearchSearchResult
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Search query | 
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **attributesToRetrieve** | **string** | Comma-separated list of attributes to return | 
 **attributesToHighlight** | **string** |  | 
 **attributesToCrop** | **string** |  | 
 **cropLength** | **int32** |  | [default to 10]
 **filter** | **string** | Filter expression | 
 **sort** | **string** | Comma-separated sort rules | 
 **facets** | **string** | Comma-separated facet attributes | 
 **showMatchesPosition** | **bool** |  | 
 **showRankingScore** | **bool** |  | 
 **matchingStrategy** | **string** |  | 

### Return type

[**SearchSearchResult**](SearchSearchResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSearchPost

> SearchSearchResult SearchSearchPost(ctx, indexUid).SearchSearchQuery(searchSearchQuery).Execute()

Search documents (POST)

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
	indexUid := "indexUid_example" // string | Unique index identifier
	searchSearchQuery := *openapiclient.NewSearchSearchQuery() // SearchSearchQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchSearchPost(context.Background(), indexUid).SearchSearchQuery(searchSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSearchPost`: SearchSearchResult
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchSearchQuery** | [**SearchSearchQuery**](SearchSearchQuery.md) |  | 

### Return type

[**SearchSearchResult**](SearchSearchResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorRecommendPoints

> VectorRecommendPoints200Response VectorRecommendPoints(ctx, collectionName).VectorRecommendPointsRequest(vectorRecommendPointsRequest).Execute()

Recommend points

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
	collectionName := "collectionName_example" // string | 
	vectorRecommendPointsRequest := *openapiclient.NewVectorRecommendPointsRequest([]openapiclient.VectorPointId{openapiclient.vector_PointId{Int32: new(int32)}}) // VectorRecommendPointsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.VectorRecommendPoints(context.Background(), collectionName).VectorRecommendPointsRequest(vectorRecommendPointsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.VectorRecommendPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorRecommendPoints`: VectorRecommendPoints200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.VectorRecommendPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorRecommendPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorRecommendPointsRequest** | [**VectorRecommendPointsRequest**](VectorRecommendPointsRequest.md) |  | 

### Return type

[**VectorRecommendPoints200Response**](VectorRecommendPoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorSearchBatch

> VectorSearchBatch200Response VectorSearchBatch(ctx, collectionName).VectorSearchBatchRequest(vectorSearchBatchRequest).Execute()

Batch search

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
	collectionName := "collectionName_example" // string | 
	vectorSearchBatchRequest := *openapiclient.NewVectorSearchBatchRequest([]openapiclient.VectorSearchRequest{*openapiclient.NewVectorSearchRequest(openapiclient.vector_VectorQuery{VectorNamedVector: openapiclient.NewVectorNamedVector()})}) // VectorSearchBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.VectorSearchBatch(context.Background(), collectionName).VectorSearchBatchRequest(vectorSearchBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.VectorSearchBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorSearchBatch`: VectorSearchBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.VectorSearchBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorSearchBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorSearchBatchRequest** | [**VectorSearchBatchRequest**](VectorSearchBatchRequest.md) |  | 

### Return type

[**VectorSearchBatch200Response**](VectorSearchBatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorSearchPoints

> VectorRecommendPoints200Response VectorSearchPoints(ctx, collectionName).VectorSearchRequest(vectorSearchRequest).Execute()

Search points

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
	collectionName := "collectionName_example" // string | 
	vectorSearchRequest := *openapiclient.NewVectorSearchRequest(openapiclient.vector_VectorQuery{VectorNamedVector: openapiclient.NewVectorNamedVector()}) // VectorSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.VectorSearchPoints(context.Background(), collectionName).VectorSearchRequest(vectorSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.VectorSearchPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorSearchPoints`: VectorRecommendPoints200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.VectorSearchPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorSearchPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorSearchRequest** | [**VectorSearchRequest**](VectorSearchRequest.md) |  | 

### Return type

[**VectorRecommendPoints200Response**](VectorRecommendPoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

