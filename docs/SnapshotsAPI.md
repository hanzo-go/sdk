# \SnapshotsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCreateSnapshot**](SnapshotsAPI.md#SearchCreateSnapshot) | **Post** /v1/search/snapshots | Create a database snapshot
[**VectorCreateSnapshot**](SnapshotsAPI.md#VectorCreateSnapshot) | **Post** /v1/vector/collections/{collection_name}/snapshots | Create snapshot
[**VectorDeleteSnapshot**](SnapshotsAPI.md#VectorDeleteSnapshot) | **Delete** /v1/vector/collections/{collection_name}/snapshots/{snapshot_name} | Delete snapshot
[**VectorDownloadSnapshot**](SnapshotsAPI.md#VectorDownloadSnapshot) | **Get** /v1/vector/collections/{collection_name}/snapshots/{snapshot_name} | Download snapshot
[**VectorListSnapshots**](SnapshotsAPI.md#VectorListSnapshots) | **Get** /v1/vector/collections/{collection_name}/snapshots | List snapshots



## SearchCreateSnapshot

> SearchSummarizedTaskView SearchCreateSnapshot(ctx).Execute()

Create a database snapshot

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
	resp, r, err := apiClient.SnapshotsAPI.SearchCreateSnapshot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.SearchCreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCreateSnapshot`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.SearchCreateSnapshot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCreateSnapshotRequest struct via the builder pattern


### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorCreateSnapshot

> VectorCreateSnapshot200Response VectorCreateSnapshot(ctx, collectionName).Execute()

Create snapshot

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.VectorCreateSnapshot(context.Background(), collectionName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.VectorCreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorCreateSnapshot`: VectorCreateSnapshot200Response
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.VectorCreateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VectorCreateSnapshot200Response**](VectorCreateSnapshot200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorDeleteSnapshot

> VectorCreateCollection200Response VectorDeleteSnapshot(ctx, collectionName, snapshotName).Execute()

Delete snapshot

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
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.VectorDeleteSnapshot(context.Background(), collectionName, snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.VectorDeleteSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorDeleteSnapshot`: VectorCreateCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.VectorDeleteSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VectorCreateCollection200Response**](VectorCreateCollection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorDownloadSnapshot

> *os.File VectorDownloadSnapshot(ctx, collectionName, snapshotName).Execute()

Download snapshot

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
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.VectorDownloadSnapshot(context.Background(), collectionName, snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.VectorDownloadSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorDownloadSnapshot`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.VectorDownloadSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorDownloadSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorListSnapshots

> VectorListSnapshots200Response VectorListSnapshots(ctx, collectionName).Execute()

List snapshots

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.VectorListSnapshots(context.Background(), collectionName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.VectorListSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorListSnapshots`: VectorListSnapshots200Response
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.VectorListSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorListSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VectorListSnapshots200Response**](VectorListSnapshots200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

