# \TaxonomyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaxonomyCategoriesById**](TaxonomyAPI.md#DeleteTaxonomyCategoriesById) | **Delete** /v1/taxonomy/categories/{id} | Removes one empty category.
[**DeleteTaxonomyTaxaById**](TaxonomyAPI.md#DeleteTaxonomyTaxaById) | **Delete** /v1/taxonomy/taxa/{id} | Removes one product from the catalogue.
[**GetTaxonomy**](TaxonomyAPI.md#GetTaxonomy) | **Get** /v1/taxonomy | Returns the product catalogue as this caller sees it: the PLATFORM catalogue — Hanzo&#39;s own products, the part that is true for everyone — plus the caller&#39;s own org&#39;s rows, every category in display order and each carrying the products filed under it in theirs.
[**PutTaxonomyCategoriesById**](TaxonomyAPI.md#PutTaxonomyCategoriesById) | **Put** /v1/taxonomy/categories/{id} | Creates or replaces one category and returns it as stored.
[**PutTaxonomyTaxaById**](TaxonomyAPI.md#PutTaxonomyTaxaById) | **Put** /v1/taxonomy/taxa/{id} | Creates or replaces one product and returns it as stored.



## DeleteTaxonomyCategoriesById

> TaxonomyDeleted DeleteTaxonomyCategoriesById(ctx, id).Execute()

Removes one empty category.



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
	id := "id_example" // string | ID is the slug to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.DeleteTaxonomyCategoriesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.DeleteTaxonomyCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTaxonomyCategoriesById`: TaxonomyDeleted
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.DeleteTaxonomyCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the slug to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaxonomyCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxonomyDeleted**](TaxonomyDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaxonomyTaxaById

> TaxonomyDeleted DeleteTaxonomyTaxaById(ctx, id).Execute()

Removes one product from the catalogue.



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
	id := "id_example" // string | ID is the slug to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.DeleteTaxonomyTaxaById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.DeleteTaxonomyTaxaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTaxonomyTaxaById`: TaxonomyDeleted
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.DeleteTaxonomyTaxaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the slug to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaxonomyTaxaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxonomyDeleted**](TaxonomyDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxonomy

> TaxonomyTaxonomy GetTaxonomy(ctx).Brand(brand).Execute()

Returns the product catalogue as this caller sees it: the PLATFORM catalogue — Hanzo's own products, the part that is true for everyone — plus the caller's own org's rows, every category in display order and each carrying the products filed under it in theirs.



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
	brand := "brand_example" // string | Brand returns only what that brand's console shows — the categories it admits, and within them the taxa scoped to it. Empty returns everything. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.GetTaxonomy(context.Background()).Brand(brand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.GetTaxonomy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxonomy`: TaxonomyTaxonomy
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.GetTaxonomy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxonomyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brand** | **string** | Brand returns only what that brand&#39;s console shows — the categories it admits, and within them the taxa scoped to it. Empty returns everything. | 

### Return type

[**TaxonomyTaxonomy**](TaxonomyTaxonomy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTaxonomyCategoriesById

> TaxonomyCategory PutTaxonomyCategoriesById(ctx, id).TaxonomyCategoryIn(taxonomyCategoryIn).Execute()

Creates or replaces one category and returns it as stored.



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
	id := "id_example" // string | ID is the category slug to write, from the path.
	taxonomyCategoryIn := *openapiclient.NewTaxonomyCategoryIn() // TaxonomyCategoryIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.PutTaxonomyCategoriesById(context.Background(), id).TaxonomyCategoryIn(taxonomyCategoryIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.PutTaxonomyCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTaxonomyCategoriesById`: TaxonomyCategory
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.PutTaxonomyCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the category slug to write, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTaxonomyCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxonomyCategoryIn** | [**TaxonomyCategoryIn**](TaxonomyCategoryIn.md) |  | 

### Return type

[**TaxonomyCategory**](TaxonomyCategory.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTaxonomyTaxaById

> TaxonomyTaxon PutTaxonomyTaxaById(ctx, id).TaxonomyTaxonIn(taxonomyTaxonIn).Execute()

Creates or replaces one product and returns it as stored.



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
	id := "id_example" // string | ID is the taxon slug to write, from the path.
	taxonomyTaxonIn := *openapiclient.NewTaxonomyTaxonIn() // TaxonomyTaxonIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.PutTaxonomyTaxaById(context.Background(), id).TaxonomyTaxonIn(taxonomyTaxonIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.PutTaxonomyTaxaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTaxonomyTaxaById`: TaxonomyTaxon
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.PutTaxonomyTaxaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the taxon slug to write, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTaxonomyTaxaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxonomyTaxonIn** | [**TaxonomyTaxonIn**](TaxonomyTaxonIn.md) |  | 

### Return type

[**TaxonomyTaxon**](TaxonomyTaxon.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

