# \AutoPiecesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoGetPiece**](AutoPiecesAPI.md#AutoGetPiece) | **Get** /v1/auto/pieces/{name} | Get piece metadata by name
[**AutoGetPieceOptions**](AutoPiecesAPI.md#AutoGetPieceOptions) | **Post** /v1/auto/pieces/options | Get dynamic property options for a piece action/trigger
[**AutoListPieceCategories**](AutoPiecesAPI.md#AutoListPieceCategories) | **Get** /v1/auto/pieces/categories | List piece categories
[**AutoListPieceVersions**](AutoPiecesAPI.md#AutoListPieceVersions) | **Get** /v1/auto/pieces/versions | List available versions of a piece
[**AutoListPieces**](AutoPiecesAPI.md#AutoListPieces) | **Get** /v1/auto/pieces | List available pieces



## AutoGetPiece

> AutoPiece AutoGetPiece(ctx, name).Version(version).Execute()

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
	resp, r, err := apiClient.AutoPiecesAPI.AutoGetPiece(context.Background(), name).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPiecesAPI.AutoGetPiece``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoGetPiece`: AutoPiece
	fmt.Fprintf(os.Stdout, "Response from `AutoPiecesAPI.AutoGetPiece`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoGetPieceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** |  | 

### Return type

[**AutoPiece**](AutoPiece.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoGetPieceOptions

> map[string]interface{} AutoGetPieceOptions(ctx).AutoGetPieceOptionsRequest(autoGetPieceOptionsRequest).Execute()

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
	resp, r, err := apiClient.AutoPiecesAPI.AutoGetPieceOptions(context.Background()).AutoGetPieceOptionsRequest(autoGetPieceOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPiecesAPI.AutoGetPieceOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoGetPieceOptions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoPiecesAPI.AutoGetPieceOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoGetPieceOptionsRequest struct via the builder pattern


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


## AutoListPieceCategories

> map[string]interface{} AutoListPieceCategories(ctx).Execute()

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
	resp, r, err := apiClient.AutoPiecesAPI.AutoListPieceCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPiecesAPI.AutoListPieceCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListPieceCategories`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoPiecesAPI.AutoListPieceCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutoListPieceCategoriesRequest struct via the builder pattern


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


## AutoListPieceVersions

> map[string]interface{} AutoListPieceVersions(ctx).Name(name).Execute()

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
	resp, r, err := apiClient.AutoPiecesAPI.AutoListPieceVersions(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPiecesAPI.AutoListPieceVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListPieceVersions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoPiecesAPI.AutoListPieceVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListPieceVersionsRequest struct via the builder pattern


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


## AutoListPieces

> []AutoPiece AutoListPieces(ctx).SearchQuery(searchQuery).Categories(categories).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPiecesAPI.AutoListPieces(context.Background()).SearchQuery(searchQuery).Categories(categories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPiecesAPI.AutoListPieces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListPieces`: []AutoPiece
	fmt.Fprintf(os.Stdout, "Response from `AutoPiecesAPI.AutoListPieces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListPiecesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchQuery** | **string** |  | 
 **categories** | **[]string** |  | 

### Return type

[**[]AutoPiece**](AutoPiece.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

