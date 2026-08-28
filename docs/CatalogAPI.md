# \CatalogAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalog**](CatalogAPI.md#GetCatalog) | **Get** /v1/catalog | Browse searches AND browses the cross-org catalog: every project, app and site the fleet has built, whichever org built it.



## GetCatalog

> CatalogPage GetCatalog(ctx).Q(q).Org(org).Kind(kind).Origin(origin).Archetype(archetype).Language(language).Template(template).Forkable(forkable).Limit(limit).Offset(offset).Execute()

Browse searches AND browses the cross-org catalog: every project, app and site the fleet has built, whichever org built it.



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
	q := "q_example" // string | Q is the free-text query the lexical index scores relevance on. Empty is a browse rather than a search — the same request either way. (optional)
	org := "org_example" // string | Org narrows to one builder org: hanzo | lux | zoo. Case-insensitive. (optional)
	kind := "kind_example" // string | Kind narrows to repo | site. Case-insensitive. (optional)
	origin := "template" // string | Origin narrows to what a row IS to you: template | community | third-party | product. This is the axis the two hanzo.app lanes are cut on. (optional)
	archetype := "archetype_example" // string | Archetype narrows to one project archetype. Case-insensitive. (optional)
	language := "typescript" // string | Language narrows to one implementation language. Case-insensitive. (optional)
	template := "template_example" // string | Template narrows a lane to ONE lineage: the id of the parent everything returned was forked from. (optional)
	forkable := "true" // string | Forkable is tri-state: \"true\" selects the forkable rows, \"false\" selects the rest, and anything else — including absent — applies no filter at all. (optional)
	limit := "20" // string | Limit caps the page at 200, default 50. A value that is not a non-negative integer falls back to the default. (optional)
	offset := "offset_example" // string | Offset is where the page starts, default 0, with the same tolerance. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.GetCatalog(context.Background()).Q(q).Org(org).Kind(kind).Origin(origin).Archetype(archetype).Language(language).Template(template).Forkable(forkable).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalog`: CatalogPage
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q is the free-text query the lexical index scores relevance on. Empty is a browse rather than a search — the same request either way. | 
 **org** | **string** | Org narrows to one builder org: hanzo | lux | zoo. Case-insensitive. | 
 **kind** | **string** | Kind narrows to repo | site. Case-insensitive. | 
 **origin** | **string** | Origin narrows to what a row IS to you: template | community | third-party | product. This is the axis the two hanzo.app lanes are cut on. | 
 **archetype** | **string** | Archetype narrows to one project archetype. Case-insensitive. | 
 **language** | **string** | Language narrows to one implementation language. Case-insensitive. | 
 **template** | **string** | Template narrows a lane to ONE lineage: the id of the parent everything returned was forked from. | 
 **forkable** | **string** | Forkable is tri-state: \&quot;true\&quot; selects the forkable rows, \&quot;false\&quot; selects the rest, and anything else — including absent — applies no filter at all. | 
 **limit** | **string** | Limit caps the page at 200, default 50. A value that is not a non-negative integer falls back to the default. | 
 **offset** | **string** | Offset is where the page starts, default 0, with the same tolerance. | 

### Return type

[**CatalogPage**](CatalogPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

