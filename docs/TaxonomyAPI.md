# \TaxonomyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaxonomyCategoriesById**](TaxonomyAPI.md#DeleteTaxonomyCategoriesById) | **Delete** /v1/taxonomy/categories/{id} | Removes one empty category.
[**DeleteTaxonomyTaxaById**](TaxonomyAPI.md#DeleteTaxonomyTaxaById) | **Delete** /v1/taxonomy/taxa/{id} | Removes one product from the catalogue.
[**GetTaxonomy**](TaxonomyAPI.md#GetTaxonomy) | **Get** /v1/taxonomy | Read returns the product catalogue as this caller sees it: the PLATFORM catalogue — Hanzo&#39;s own products, the part that is true for everyone — plus the caller&#39;s own org&#39;s rows, every category in display order and each carrying the products filed under it in theirs.
[**PutTaxonomyCategoriesById**](TaxonomyAPI.md#PutTaxonomyCategoriesById) | **Put** /v1/taxonomy/categories/{id} | Creates or replaces one category and returns it as stored.
[**PutTaxonomyTaxaById**](TaxonomyAPI.md#PutTaxonomyTaxaById) | **Put** /v1/taxonomy/taxa/{id} | Creates or replaces one product and returns it as stored.



## DeleteTaxonomyCategoriesById

> Deleted DeleteTaxonomyCategoriesById(ctx, id).Execute()

Removes one empty category.



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
	id := "id_example" // string | ID is the slug to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.DeleteTaxonomyCategoriesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.DeleteTaxonomyCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTaxonomyCategoriesById`: Deleted
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

[**Deleted**](Deleted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaxonomyTaxaById

> Deleted DeleteTaxonomyTaxaById(ctx, id).Execute()

Removes one product from the catalogue.



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
	id := "id_example" // string | ID is the slug to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.DeleteTaxonomyTaxaById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.DeleteTaxonomyTaxaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTaxonomyTaxaById`: Deleted
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

[**Deleted**](Deleted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxonomy

> Taxonomy GetTaxonomy(ctx).Brand(brand).Execute()

Read returns the product catalogue as this caller sees it: the PLATFORM catalogue — Hanzo's own products, the part that is true for everyone — plus the caller's own org's rows, every category in display order and each carrying the products filed under it in theirs.



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
	brand := "brand_example" // string | Brand returns only what that brand's console shows — the categories it admits, and within them the taxa scoped to it. Empty returns everything. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.GetTaxonomy(context.Background()).Brand(brand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.GetTaxonomy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxonomy`: Taxonomy
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

[**Taxonomy**](Taxonomy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTaxonomyCategoriesById

> Category PutTaxonomyCategoriesById(ctx, id).CategoryIn(categoryIn).Execute()

Creates or replaces one category and returns it as stored.



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
	id := "id_example" // string | ID is the category slug to write, from the path.
	categoryIn := *openapiclient.NewCategoryIn() // CategoryIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.PutTaxonomyCategoriesById(context.Background(), id).CategoryIn(categoryIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.PutTaxonomyCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTaxonomyCategoriesById`: Category
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

 **categoryIn** | [**CategoryIn**](CategoryIn.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTaxonomyTaxaById

> Taxon PutTaxonomyTaxaById(ctx, id).TaxonIn(taxonIn).Execute()

Creates or replaces one product and returns it as stored.



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
	id := "id_example" // string | ID is the taxon slug to write, from the path.
	taxonIn := *openapiclient.NewTaxonIn() // TaxonIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.PutTaxonomyTaxaById(context.Background(), id).TaxonIn(taxonIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.PutTaxonomyTaxaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTaxonomyTaxaById`: Taxon
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

 **taxonIn** | [**TaxonIn**](TaxonIn.md) |  | 

### Return type

[**Taxon**](Taxon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

