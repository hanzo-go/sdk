# \CatalogAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1CatalogEntriesByWildcard1**](CatalogAPI.md#CloudDeleteV1CatalogEntriesByWildcard1) | **Delete** /v1/catalog/entries/{wildcard1} | 
[**CloudGetV1Catalog**](CatalogAPI.md#CloudGetV1Catalog) | **Get** /v1/catalog | Browse searches AND browses the cross-org catalog: every project, app and site the fleet has built, whichever org built it.
[**CloudGetV1CatalogEntries**](CatalogAPI.md#CloudGetV1CatalogEntries) | **Get** /v1/catalog/entries | 
[**CloudPostV1CatalogEntries**](CatalogAPI.md#CloudPostV1CatalogEntries) | **Post** /v1/catalog/entries | 
[**CloudPostV1CatalogModels**](CatalogAPI.md#CloudPostV1CatalogModels) | **Post** /v1/catalog/models | 
[**CloudPostV1CatalogModelsRefresh**](CatalogAPI.md#CloudPostV1CatalogModelsRefresh) | **Post** /v1/catalog/models/refresh | 
[**CloudPostV1CatalogSeed**](CatalogAPI.md#CloudPostV1CatalogSeed) | **Post** /v1/catalog/seed | 
[**CloudPutV1CatalogEntriesByWildcard1**](CatalogAPI.md#CloudPutV1CatalogEntriesByWildcard1) | **Put** /v1/catalog/entries/{wildcard1} | 



## CloudDeleteV1CatalogEntriesByWildcard1

> CloudDeleteV1CatalogEntriesByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPI.CloudDeleteV1CatalogEntriesByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CloudDeleteV1CatalogEntriesByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CatalogEntriesByWildcard1Request struct via the builder pattern


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


## CloudGetV1Catalog

> CloudCatalogPage CloudGetV1Catalog(ctx).Q(q).Org(org).Kind(kind).Origin(origin).Archetype(archetype).Language(language).Template(template).Forkable(forkable).Limit(limit).Offset(offset).Execute()

Browse searches AND browses the cross-org catalog: every project, app and site the fleet has built, whichever org built it.



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
	resp, r, err := apiClient.CatalogAPI.CloudGetV1Catalog(context.Background()).Q(q).Org(org).Kind(kind).Origin(origin).Archetype(archetype).Language(language).Template(template).Forkable(forkable).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CloudGetV1Catalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Catalog`: CloudCatalogPage
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.CloudGetV1Catalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CatalogRequest struct via the builder pattern


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

[**CloudCatalogPage**](CloudCatalogPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CatalogEntries

> CloudGetV1CatalogEntries(ctx).Execute()



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
	r, err := apiClient.CatalogAPI.CloudGetV1CatalogEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CloudGetV1CatalogEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CatalogEntriesRequest struct via the builder pattern


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


## CloudPostV1CatalogEntries

> CloudPostV1CatalogEntries(ctx).Execute()



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
	r, err := apiClient.CatalogAPI.CloudPostV1CatalogEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CloudPostV1CatalogEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CatalogEntriesRequest struct via the builder pattern


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


## CloudPostV1CatalogModels

> CloudPostV1CatalogModels(ctx).Execute()



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
	r, err := apiClient.CatalogAPI.CloudPostV1CatalogModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CloudPostV1CatalogModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CatalogModelsRequest struct via the builder pattern


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


## CloudPostV1CatalogModelsRefresh

> CloudPostV1CatalogModelsRefresh(ctx).Execute()



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
	r, err := apiClient.CatalogAPI.CloudPostV1CatalogModelsRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CloudPostV1CatalogModelsRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CatalogModelsRefreshRequest struct via the builder pattern


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


## CloudPostV1CatalogSeed

> CloudPostV1CatalogSeed(ctx).Execute()



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
	r, err := apiClient.CatalogAPI.CloudPostV1CatalogSeed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CloudPostV1CatalogSeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CatalogSeedRequest struct via the builder pattern


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


## CloudPutV1CatalogEntriesByWildcard1

> CloudPutV1CatalogEntriesByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPI.CloudPutV1CatalogEntriesByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CloudPutV1CatalogEntriesByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CatalogEntriesByWildcard1Request struct via the builder pattern


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

