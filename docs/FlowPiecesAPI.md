# \FlowPiecesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowGetPiece**](FlowPiecesAPI.md#FlowGetPiece) | **Get** /v1/flow/pieces/{name} | Get piece metadata by name
[**FlowGetPieceOptions**](FlowPiecesAPI.md#FlowGetPieceOptions) | **Post** /v1/flow/pieces/options | Get dynamic property options for a piece action/trigger
[**FlowGetPieceRegistry**](FlowPiecesAPI.md#FlowGetPieceRegistry) | **Get** /v1/flow/pieces/registry | Get piece registry metadata
[**FlowGetScopedPiece**](FlowPiecesAPI.md#FlowGetScopedPiece) | **Get** /v1/flow/pieces/{scope}/{name} | Get scoped piece metadata
[**FlowListPieceCategories**](FlowPiecesAPI.md#FlowListPieceCategories) | **Get** /v1/flow/pieces/categories | List piece categories
[**FlowListPieceVersions**](FlowPiecesAPI.md#FlowListPieceVersions) | **Get** /v1/flow/pieces/versions | List available versions of a piece
[**FlowListPieces**](FlowPiecesAPI.md#FlowListPieces) | **Get** /v1/flow/pieces | List available pieces
[**FlowSyncPieces**](FlowPiecesAPI.md#FlowSyncPieces) | **Post** /v1/flow/pieces/sync | Sync piece metadata from registry



## FlowGetPiece

> FlowPiece FlowGetPiece(ctx, name).Version(version).Execute()

Get piece metadata by name

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
	name := "name_example" // string | 
	version := "version_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowPiecesAPI.FlowGetPiece(context.Background(), name).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPiecesAPI.FlowGetPiece``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetPiece`: FlowPiece
	fmt.Fprintf(os.Stdout, "Response from `FlowPiecesAPI.FlowGetPiece`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetPieceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** |  | 

### Return type

[**FlowPiece**](FlowPiece.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetPieceOptions

> map[string]interface{} FlowGetPieceOptions(ctx).AutoGetPieceOptionsRequest(autoGetPieceOptionsRequest).Execute()

Get dynamic property options for a piece action/trigger

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
	autoGetPieceOptionsRequest := *openapiclient.NewAutoGetPieceOptionsRequest("FlowId_example", "FlowVersionId_example", "ActionOrTriggerName_example", "PropertyName_example", "PieceName_example", "PieceVersion_example") // AutoGetPieceOptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowPiecesAPI.FlowGetPieceOptions(context.Background()).AutoGetPieceOptionsRequest(autoGetPieceOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPiecesAPI.FlowGetPieceOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetPieceOptions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowPiecesAPI.FlowGetPieceOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetPieceOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoGetPieceOptionsRequest** | [**AutoGetPieceOptionsRequest**](AutoGetPieceOptionsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetPieceRegistry

> map[string]interface{} FlowGetPieceRegistry(ctx).Execute()

Get piece registry metadata

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
	resp, r, err := apiClient.FlowPiecesAPI.FlowGetPieceRegistry(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPiecesAPI.FlowGetPieceRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetPieceRegistry`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowPiecesAPI.FlowGetPieceRegistry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetPieceRegistryRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetScopedPiece

> FlowPiece FlowGetScopedPiece(ctx, scope, name).Version(version).Execute()

Get scoped piece metadata

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
	scope := "scope_example" // string | 
	name := "name_example" // string | 
	version := "version_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowPiecesAPI.FlowGetScopedPiece(context.Background(), scope, name).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPiecesAPI.FlowGetScopedPiece``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetScopedPiece`: FlowPiece
	fmt.Fprintf(os.Stdout, "Response from `FlowPiecesAPI.FlowGetScopedPiece`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetScopedPieceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** |  | 

### Return type

[**FlowPiece**](FlowPiece.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowListPieceCategories

> []string FlowListPieceCategories(ctx).Execute()

List piece categories

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
	resp, r, err := apiClient.FlowPiecesAPI.FlowListPieceCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPiecesAPI.FlowListPieceCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListPieceCategories`: []string
	fmt.Fprintf(os.Stdout, "Response from `FlowPiecesAPI.FlowListPieceCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListPieceCategoriesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowListPieceVersions

> map[string]interface{} FlowListPieceVersions(ctx).Name(name).Execute()

List available versions of a piece

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowPiecesAPI.FlowListPieceVersions(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPiecesAPI.FlowListPieceVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListPieceVersions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowPiecesAPI.FlowListPieceVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListPieceVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowListPieces

> []FlowPiece FlowListPieces(ctx).SearchQuery(searchQuery).Categories(categories).IncludeHidden(includeHidden).IncludeTags(includeTags).SortBy(sortBy).OrderBy(orderBy).Execute()

List available pieces

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
	searchQuery := "searchQuery_example" // string |  (optional)
	categories := []string{"Inner_example"} // []string |  (optional)
	includeHidden := true // bool |  (optional) (default to false)
	includeTags := true // bool |  (optional) (default to false)
	sortBy := "sortBy_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowPiecesAPI.FlowListPieces(context.Background()).SearchQuery(searchQuery).Categories(categories).IncludeHidden(includeHidden).IncludeTags(includeTags).SortBy(sortBy).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPiecesAPI.FlowListPieces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListPieces`: []FlowPiece
	fmt.Fprintf(os.Stdout, "Response from `FlowPiecesAPI.FlowListPieces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListPiecesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchQuery** | **string** |  | 
 **categories** | **[]string** |  | 
 **includeHidden** | **bool** |  | [default to false]
 **includeTags** | **bool** |  | [default to false]
 **sortBy** | **string** |  | 
 **orderBy** | **string** |  | 

### Return type

[**[]FlowPiece**](FlowPiece.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowSyncPieces

> map[string]interface{} FlowSyncPieces(ctx).Execute()

Sync piece metadata from registry

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
	resp, r, err := apiClient.FlowPiecesAPI.FlowSyncPieces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPiecesAPI.FlowSyncPieces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowSyncPieces`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowPiecesAPI.FlowSyncPieces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowSyncPiecesRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

