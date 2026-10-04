# \IndexAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndexIndexesByUid**](IndexAPI.md#DeleteIndexIndexesByUid) | **Delete** /v1/index/indexes/{uid} | Deletes an index and everything in it.
[**DeleteIndexIndexesByUidDocumentsById**](IndexAPI.md#DeleteIndexIndexesByUidDocumentsById) | **Delete** /v1/index/indexes/{uid}/documents/{id} | Deletes one document by its primary key.
[**GetIndexHealth**](IndexAPI.md#GetIndexHealth) | **Get** /v1/index/health | Reports whether the search plane can serve.
[**GetIndexIndexes**](IndexAPI.md#GetIndexIndexes) | **Get** /v1/index/indexes | Lists the indexes your org holds.
[**GetIndexIndexesByUid**](IndexAPI.md#GetIndexIndexesByUid) | **Get** /v1/index/indexes/{uid} | Reads one index&#39;s definition.
[**GetIndexIndexesByUidDocuments**](IndexAPI.md#GetIndexIndexesByUidDocuments) | **Get** /v1/index/indexes/{uid}/documents | Pages through the documents in an index.
[**GetIndexIndexesByUidDocumentsById**](IndexAPI.md#GetIndexIndexesByUidDocumentsById) | **Get** /v1/index/indexes/{uid}/documents/{id} | Reads one document by its primary key.
[**GetIndexIndexesByUidSettings**](IndexAPI.md#GetIndexIndexesByUidSettings) | **Get** /v1/index/indexes/{uid}/settings | Reads an index&#39;s filterable attributes.
[**GetIndexStats**](IndexAPI.md#GetIndexStats) | **Get** /v1/index/stats | Counts the documents in each of your indexes.
[**GetIndexTasksByUid**](IndexAPI.md#GetIndexTasksByUid) | **Get** /v1/index/tasks/{uid} | Checks a write task, which has already finished.
[**GetIndexVersion**](IndexAPI.md#GetIndexVersion) | **Get** /v1/index/version | Identifies the search implementation answering.
[**PatchIndexIndexesByUidSettings**](IndexAPI.md#PatchIndexIndexesByUidSettings) | **Patch** /v1/index/indexes/{uid}/settings | Sets which attributes an index can be filtered on.
[**PostIndexIndexes**](IndexAPI.md#PostIndexIndexes) | **Post** /v1/index/indexes | Creates an index.
[**PostIndexIndexesByUidDocuments**](IndexAPI.md#PostIndexIndexesByUidDocuments) | **Post** /v1/index/indexes/{uid}/documents | Add or replace documents in an index
[**PostIndexIndexesByUidDocumentsDeleteBatch**](IndexAPI.md#PostIndexIndexesByUidDocumentsDeleteBatch) | **Post** /v1/index/indexes/{uid}/documents/delete-batch | Delete many documents by primary key in one call
[**PostIndexIndexesByUidSearch**](IndexAPI.md#PostIndexIndexesByUidSearch) | **Post** /v1/index/indexes/{uid}/search | Searches an index, forgiving typos.
[**PutIndexIndexesByUidDocuments**](IndexAPI.md#PutIndexIndexesByUidDocuments) | **Put** /v1/index/indexes/{uid}/documents | Add or update documents in an index



## DeleteIndexIndexesByUid

> IndexIndexEnqueued DeleteIndexIndexesByUid(ctx, uid).Execute()

Deletes an index and everything in it.



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.DeleteIndexIndexesByUid(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.DeleteIndexIndexesByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIndexIndexesByUid`: IndexIndexEnqueued
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.DeleteIndexIndexesByUid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndexIndexesByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexIndexEnqueued**](IndexIndexEnqueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndexIndexesByUidDocumentsById

> IndexIndexEnqueued DeleteIndexIndexesByUidDocumentsById(ctx, uid, id).Execute()

Deletes one document by its primary key.



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
	uid := "uid_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.DeleteIndexIndexesByUidDocumentsById(context.Background(), uid, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.DeleteIndexIndexesByUidDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIndexIndexesByUidDocumentsById`: IndexIndexEnqueued
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.DeleteIndexIndexesByUidDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndexIndexesByUidDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IndexIndexEnqueued**](IndexIndexEnqueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexHealth

> IndexIndexHealth GetIndexHealth(ctx).Execute()

Reports whether the search plane can serve.



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
	resp, r, err := apiClient.IndexAPI.GetIndexHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexHealth`: IndexIndexHealth
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexHealthRequest struct via the builder pattern


### Return type

[**IndexIndexHealth**](IndexIndexHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexes

> IndexIndexList GetIndexIndexes(ctx).Execute()

Lists the indexes your org holds.



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
	resp, r, err := apiClient.IndexAPI.GetIndexIndexes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexIndexes`: IndexIndexList
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexIndexes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexIndexesRequest struct via the builder pattern


### Return type

[**IndexIndexList**](IndexIndexList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexesByUid

> IndexIndexView GetIndexIndexesByUid(ctx, uid).Execute()

Reads one index's definition.



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.GetIndexIndexesByUid(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexesByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexIndexesByUid`: IndexIndexView
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexIndexesByUid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexIndexesByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexIndexView**](IndexIndexView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexesByUidDocuments

> IndexIndexDocuments GetIndexIndexesByUidDocuments(ctx, uid).Limit(limit).Offset(offset).Execute()

Pages through the documents in an index.



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
	uid := "uid_example" // string | 
	limit := "limit_example" // string |  (optional)
	offset := "offset_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.GetIndexIndexesByUidDocuments(context.Background(), uid).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexesByUidDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexIndexesByUidDocuments`: IndexIndexDocuments
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexIndexesByUidDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexIndexesByUidDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**IndexIndexDocuments**](IndexIndexDocuments.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexesByUidDocumentsById

> interface{} GetIndexIndexesByUidDocumentsById(ctx, uid, id).Execute()

Reads one document by its primary key.



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
	uid := "uid_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.GetIndexIndexesByUidDocumentsById(context.Background(), uid, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexesByUidDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexIndexesByUidDocumentsById`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexIndexesByUidDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexIndexesByUidDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexesByUidSettings

> IndexIndexSettings GetIndexIndexesByUidSettings(ctx, uid).Execute()

Reads an index's filterable attributes.



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.GetIndexIndexesByUidSettings(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexesByUidSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexIndexesByUidSettings`: IndexIndexSettings
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexIndexesByUidSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexIndexesByUidSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexIndexSettings**](IndexIndexSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexStats

> IndexIndexStats GetIndexStats(ctx).Execute()

Counts the documents in each of your indexes.



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
	resp, r, err := apiClient.IndexAPI.GetIndexStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexStats`: IndexIndexStats
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexStatsRequest struct via the builder pattern


### Return type

[**IndexIndexStats**](IndexIndexStats.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexTasksByUid

> IndexIndexTask GetIndexTasksByUid(ctx, uid).Execute()

Checks a write task, which has already finished.



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
	uid := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.GetIndexTasksByUid(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexTasksByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexTasksByUid`: IndexIndexTask
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexTasksByUid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexTasksByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexIndexTask**](IndexIndexTask.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexVersion

> IndexIndexVersion GetIndexVersion(ctx).Execute()

Identifies the search implementation answering.



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
	resp, r, err := apiClient.IndexAPI.GetIndexVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexVersion`: IndexIndexVersion
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.GetIndexVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexVersionRequest struct via the builder pattern


### Return type

[**IndexIndexVersion**](IndexIndexVersion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIndexIndexesByUidSettings

> IndexIndexEnqueued PatchIndexIndexesByUidSettings(ctx, uid).IndexIndexFilter(indexIndexFilter).Execute()

Sets which attributes an index can be filtered on.



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
	uid := "uid_example" // string | 
	indexIndexFilter := *openapiclient.NewIndexIndexFilter() // IndexIndexFilter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.PatchIndexIndexesByUidSettings(context.Background(), uid).IndexIndexFilter(indexIndexFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PatchIndexIndexesByUidSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchIndexIndexesByUidSettings`: IndexIndexEnqueued
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.PatchIndexIndexesByUidSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIndexIndexesByUidSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexIndexFilter** | [**IndexIndexFilter**](IndexIndexFilter.md) |  | 

### Return type

[**IndexIndexEnqueued**](IndexIndexEnqueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIndexIndexes

> IndexIndexEnqueued PostIndexIndexes(ctx).IndexIndexNew(indexIndexNew).Execute()

Creates an index.



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
	indexIndexNew := *openapiclient.NewIndexIndexNew() // IndexIndexNew | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.PostIndexIndexes(context.Background()).IndexIndexNew(indexIndexNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PostIndexIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIndexIndexes`: IndexIndexEnqueued
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.PostIndexIndexes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIndexIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexIndexNew** | [**IndexIndexNew**](IndexIndexNew.md) |  | 

### Return type

[**IndexIndexEnqueued**](IndexIndexEnqueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIndexIndexesByUidDocuments

> IndexEnqueued PostIndexIndexesByUidDocuments(ctx, uid).RequestBody(requestBody).Execute()

Add or replace documents in an index



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
	uid := "uid_example" // string | 
	requestBody := []interface{}{interface{}(123)} // []interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.PostIndexIndexesByUidDocuments(context.Background(), uid).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PostIndexIndexesByUidDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIndexIndexesByUidDocuments`: IndexEnqueued
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.PostIndexIndexesByUidDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIndexIndexesByUidDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** |  | 

### Return type

[**IndexEnqueued**](IndexEnqueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIndexIndexesByUidDocumentsDeleteBatch

> IndexEnqueued PostIndexIndexesByUidDocumentsDeleteBatch(ctx, uid).RequestBody(requestBody).Execute()

Delete many documents by primary key in one call



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
	uid := "uid_example" // string | 
	requestBody := []interface{}{interface{}(123)} // []interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.PostIndexIndexesByUidDocumentsDeleteBatch(context.Background(), uid).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PostIndexIndexesByUidDocumentsDeleteBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIndexIndexesByUidDocumentsDeleteBatch`: IndexEnqueued
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.PostIndexIndexesByUidDocumentsDeleteBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIndexIndexesByUidDocumentsDeleteBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** |  | 

### Return type

[**IndexEnqueued**](IndexEnqueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIndexIndexesByUidSearch

> IndexIndexHits PostIndexIndexesByUidSearch(ctx, uid).IndexIndexQuery(indexIndexQuery).Execute()

Searches an index, forgiving typos.



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
	uid := "uid_example" // string | 
	indexIndexQuery := *openapiclient.NewIndexIndexQuery() // IndexIndexQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.PostIndexIndexesByUidSearch(context.Background(), uid).IndexIndexQuery(indexIndexQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PostIndexIndexesByUidSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIndexIndexesByUidSearch`: IndexIndexHits
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.PostIndexIndexesByUidSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIndexIndexesByUidSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexIndexQuery** | [**IndexIndexQuery**](IndexIndexQuery.md) |  | 

### Return type

[**IndexIndexHits**](IndexIndexHits.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIndexIndexesByUidDocuments

> IndexEnqueued PutIndexIndexesByUidDocuments(ctx, uid).RequestBody(requestBody).Execute()

Add or update documents in an index



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
	uid := "uid_example" // string | 
	requestBody := []interface{}{interface{}(123)} // []interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexAPI.PutIndexIndexesByUidDocuments(context.Background(), uid).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PutIndexIndexesByUidDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIndexIndexesByUidDocuments`: IndexEnqueued
	fmt.Fprintf(os.Stdout, "Response from `IndexAPI.PutIndexIndexesByUidDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIndexIndexesByUidDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** |  | 

### Return type

[**IndexEnqueued**](IndexEnqueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

