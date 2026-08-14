# \IndexAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndexIndexesByUid**](IndexAPI.md#DeleteIndexIndexesByUid) | **Delete** /v1/index/indexes/{uid} | Delete an index and everything in it
[**DeleteIndexIndexesByUidDocumentsById**](IndexAPI.md#DeleteIndexIndexesByUidDocumentsById) | **Delete** /v1/index/indexes/{uid}/documents/{id} | Delete one document by its primary key
[**GetIndexHealth**](IndexAPI.md#GetIndexHealth) | **Get** /v1/index/health | Report whether the search plane can serve
[**GetIndexIndexes**](IndexAPI.md#GetIndexIndexes) | **Get** /v1/index/indexes | List the indexes your org holds
[**GetIndexIndexesByUid**](IndexAPI.md#GetIndexIndexesByUid) | **Get** /v1/index/indexes/{uid} | Read one index&#39;s definition
[**GetIndexIndexesByUidDocuments**](IndexAPI.md#GetIndexIndexesByUidDocuments) | **Get** /v1/index/indexes/{uid}/documents | Page through the documents in an index
[**GetIndexIndexesByUidDocumentsById**](IndexAPI.md#GetIndexIndexesByUidDocumentsById) | **Get** /v1/index/indexes/{uid}/documents/{id} | Read one document by its primary key
[**GetIndexIndexesByUidSettings**](IndexAPI.md#GetIndexIndexesByUidSettings) | **Get** /v1/index/indexes/{uid}/settings | Read an index&#39;s filterable attributes
[**GetIndexStats**](IndexAPI.md#GetIndexStats) | **Get** /v1/index/stats | Count the documents in each of your indexes
[**GetIndexTasksByUid**](IndexAPI.md#GetIndexTasksByUid) | **Get** /v1/index/tasks/{uid} | Check a write task, which has already finished
[**GetIndexVersion**](IndexAPI.md#GetIndexVersion) | **Get** /v1/index/version | Identify the search implementation answering
[**PatchIndexIndexesByUidSettings**](IndexAPI.md#PatchIndexIndexesByUidSettings) | **Patch** /v1/index/indexes/{uid}/settings | Set which attributes an index can be filtered on
[**PostIndexIndexes**](IndexAPI.md#PostIndexIndexes) | **Post** /v1/index/indexes | Create an index
[**PostIndexIndexesByUidDocuments**](IndexAPI.md#PostIndexIndexesByUidDocuments) | **Post** /v1/index/indexes/{uid}/documents | Add or replace documents in an index
[**PostIndexIndexesByUidDocumentsDeleteBatch**](IndexAPI.md#PostIndexIndexesByUidDocumentsDeleteBatch) | **Post** /v1/index/indexes/{uid}/documents/delete-batch | Delete many documents by primary key in one call
[**PostIndexIndexesByUidSearch**](IndexAPI.md#PostIndexIndexesByUidSearch) | **Post** /v1/index/indexes/{uid}/search | Search an index, forgiving typos
[**PutIndexIndexesByUidDocuments**](IndexAPI.md#PutIndexIndexesByUidDocuments) | **Put** /v1/index/indexes/{uid}/documents | Add or update documents in an index



## DeleteIndexIndexesByUid

> DeleteIndexIndexesByUid(ctx, uid).Execute()

Delete an index and everything in it



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.DeleteIndexIndexesByUid(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.DeleteIndexIndexesByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndexIndexesByUidDocumentsById

> DeleteIndexIndexesByUidDocumentsById(ctx, uid, id).Execute()

Delete one document by its primary key



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
	uid := "uid_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.DeleteIndexIndexesByUidDocumentsById(context.Background(), uid, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.DeleteIndexIndexesByUidDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexHealth

> GetIndexHealth(ctx).Execute()

Report whether the search plane can serve



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
	r, err := apiClient.IndexAPI.GetIndexHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexHealthRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexes

> GetIndexIndexes(ctx).Execute()

List the indexes your org holds



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
	r, err := apiClient.IndexAPI.GetIndexIndexes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexIndexesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexesByUid

> GetIndexIndexesByUid(ctx, uid).Execute()

Read one index's definition



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.GetIndexIndexesByUid(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexesByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexesByUidDocuments

> GetIndexIndexesByUidDocuments(ctx, uid).Execute()

Page through the documents in an index



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.GetIndexIndexesByUidDocuments(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexesByUidDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexesByUidDocumentsById

> GetIndexIndexesByUidDocumentsById(ctx, uid, id).Execute()

Read one document by its primary key



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
	uid := "uid_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.GetIndexIndexesByUidDocumentsById(context.Background(), uid, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexesByUidDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexIndexesByUidSettings

> GetIndexIndexesByUidSettings(ctx, uid).Execute()

Read an index's filterable attributes



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.GetIndexIndexesByUidSettings(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexIndexesByUidSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexStats

> GetIndexStats(ctx).Execute()

Count the documents in each of your indexes



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
	r, err := apiClient.IndexAPI.GetIndexStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexStatsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexTasksByUid

> GetIndexTasksByUid(ctx, uid).Execute()

Check a write task, which has already finished



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.GetIndexTasksByUid(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexTasksByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexTasksByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexVersion

> GetIndexVersion(ctx).Execute()

Identify the search implementation answering



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
	r, err := apiClient.IndexAPI.GetIndexVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.GetIndexVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexVersionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIndexIndexesByUidSettings

> PatchIndexIndexesByUidSettings(ctx, uid).Execute()

Set which attributes an index can be filtered on



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.PatchIndexIndexesByUidSettings(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PatchIndexIndexesByUidSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIndexIndexes

> PostIndexIndexes(ctx).Execute()

Create an index



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
	r, err := apiClient.IndexAPI.PostIndexIndexes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PostIndexIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIndexIndexesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIndexIndexesByUidDocuments

> PostIndexIndexesByUidDocuments(ctx, uid).Execute()

Add or replace documents in an index



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.PostIndexIndexesByUidDocuments(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PostIndexIndexesByUidDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIndexIndexesByUidDocumentsDeleteBatch

> PostIndexIndexesByUidDocumentsDeleteBatch(ctx, uid).Execute()

Delete many documents by primary key in one call



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.PostIndexIndexesByUidDocumentsDeleteBatch(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PostIndexIndexesByUidDocumentsDeleteBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIndexIndexesByUidSearch

> PostIndexIndexesByUidSearch(ctx, uid).Execute()

Search an index, forgiving typos



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.PostIndexIndexesByUidSearch(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PostIndexIndexesByUidSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIndexIndexesByUidDocuments

> PutIndexIndexesByUidDocuments(ctx, uid).Execute()

Add or update documents in an index



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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IndexAPI.PutIndexIndexesByUidDocuments(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexAPI.PutIndexIndexesByUidDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

