# \SyncAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSyncById**](SyncAPI.md#DeleteSyncById) | **Delete** /v1/sync/{id} | Delete removes one sync and tears down the outbound mirror it derived, answering 204.
[**GetSync**](SyncAPI.md#GetSync) | **Get** /v1/sync | List returns every sync link the caller&#39;s org has, each with its two endpoints, its direction and trigger policy, and the time it last reconciled.
[**GetSyncById**](SyncAPI.md#GetSyncById) | **Get** /v1/sync/{id} | Get returns one sync by id.
[**PatchSyncById**](SyncAPI.md#PatchSyncById) | **Patch** /v1/sync/{id} | Patch updates one sync&#39;s mutable policy — direction, trigger and actor — in place.
[**PostSync**](SyncAPI.md#PostSync) | **Post** /v1/sync | Create declares a sync between two endpoints and returns it.
[**PostSyncByIdRun**](SyncAPI.md#PostSyncByIdRun) | **Post** /v1/sync/{id}/run | Run reconciles one sync now — the manual re-sync, and the initial import for a link created without run&#x3D;true.



## DeleteSyncById

> DeleteSyncById(ctx, id).Execute()

Delete removes one sync and tears down the outbound mirror it derived, answering 204.



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
	id := "sync_1" // string | ID is the sync to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncAPI.DeleteSyncById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.DeleteSyncById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sync to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSync

> SyncList GetSync(ctx).Execute()

List returns every sync link the caller's org has, each with its two endpoints, its direction and trigger policy, and the time it last reconciled.



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
	resp, r, err := apiClient.SyncAPI.GetSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.GetSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSync`: SyncList
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.GetSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncRequest struct via the builder pattern


### Return type

[**SyncList**](SyncList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncById

> SyncView GetSyncById(ctx, id).Execute()

Get returns one sync by id.



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
	id := "sync_1" // string | ID is the sync to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.GetSyncById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.GetSyncById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncById`: SyncView
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.GetSyncById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sync to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncView**](SyncView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSyncById

> SyncView PatchSyncById(ctx, id).PatchSyncIn(patchSyncIn).Execute()

Patch updates one sync's mutable policy — direction, trigger and actor — in place.



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
	id := "sync_1" // string | ID is the sync to update, from the path.
	patchSyncIn := *openapiclient.NewPatchSyncIn() // PatchSyncIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.PatchSyncById(context.Background(), id).PatchSyncIn(patchSyncIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.PatchSyncById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSyncById`: SyncView
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.PatchSyncById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sync to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSyncByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchSyncIn** | [**PatchSyncIn**](PatchSyncIn.md) |  | 

### Return type

[**SyncView**](SyncView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSync

> SyncView PostSync(ctx).SyncReq(syncReq).Execute()

Create declares a sync between two endpoints and returns it.



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
	syncReq := *openapiclient.NewSyncReq() // SyncReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.PostSync(context.Background()).SyncReq(syncReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.PostSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSync`: SyncView
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.PostSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncReq** | [**SyncReq**](SyncReq.md) |  | 

### Return type

[**SyncView**](SyncView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSyncByIdRun

> SyncQueued PostSyncByIdRun(ctx, id).Execute()

Run reconciles one sync now — the manual re-sync, and the initial import for a link created without run=true.



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
	id := "sync_1" // string | ID is the sync to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.PostSyncByIdRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.PostSyncByIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSyncByIdRun`: SyncQueued
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.PostSyncByIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sync to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncByIdRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncQueued**](SyncQueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

