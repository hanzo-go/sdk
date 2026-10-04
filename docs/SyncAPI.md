# \SyncAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSyncById**](SyncAPI.md#DeleteSyncById) | **Delete** /v1/sync/{id} | Removes one sync and tears down the outbound mirror it derived, answering 204.
[**GetSync**](SyncAPI.md#GetSync) | **Get** /v1/sync | Returns every sync link the caller&#39;s org has, each with its two endpoints, its direction and trigger policy, the time it last reconciled, and for a repo link its native copy on the forge — the address to clone it from, its default branch and whether it is in step.
[**GetSyncById**](SyncAPI.md#GetSyncById) | **Get** /v1/sync/{id} | Returns one sync by id.
[**PatchSyncById**](SyncAPI.md#PatchSyncById) | **Patch** /v1/sync/{id} | Updates one sync&#39;s mutable policy — direction, trigger and actor — in place.
[**PostSync**](SyncAPI.md#PostSync) | **Post** /v1/sync | Declares a sync between two endpoints and returns it.
[**PostSyncByIdRun**](SyncAPI.md#PostSyncByIdRun) | **Post** /v1/sync/{id}/run | Reconciles one sync now — the manual re-sync, and the initial import for a link created without run&#x3D;true.



## DeleteSyncById

> DeleteSyncById(ctx, id).Execute()

Removes one sync and tears down the outbound mirror it derived, answering 204.



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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSync

> SyncSyncList GetSync(ctx).Execute()

Returns every sync link the caller's org has, each with its two endpoints, its direction and trigger policy, the time it last reconciled, and for a repo link its native copy on the forge — the address to clone it from, its default branch and whether it is in step.



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
	resp, r, err := apiClient.SyncAPI.GetSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.GetSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSync`: SyncSyncList
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.GetSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncRequest struct via the builder pattern


### Return type

[**SyncSyncList**](SyncSyncList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncById

> SyncSyncView GetSyncById(ctx, id).Execute()

Returns one sync by id.



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
	id := "sync_1" // string | ID is the sync to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.GetSyncById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.GetSyncById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncById`: SyncSyncView
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

[**SyncSyncView**](SyncSyncView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSyncById

> SyncSyncView PatchSyncById(ctx, id).SyncPatchSyncIn(syncPatchSyncIn).Execute()

Updates one sync's mutable policy — direction, trigger and actor — in place.



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
	id := "sync_1" // string | ID is the sync to update, from the path.
	syncPatchSyncIn := *openapiclient.NewSyncPatchSyncIn() // SyncPatchSyncIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.PatchSyncById(context.Background(), id).SyncPatchSyncIn(syncPatchSyncIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.PatchSyncById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSyncById`: SyncSyncView
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

 **syncPatchSyncIn** | [**SyncPatchSyncIn**](SyncPatchSyncIn.md) |  | 

### Return type

[**SyncSyncView**](SyncSyncView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSync

> SyncSyncView PostSync(ctx).SyncSyncReq(syncSyncReq).Execute()

Declares a sync between two endpoints and returns it.



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
	syncSyncReq := *openapiclient.NewSyncSyncReq() // SyncSyncReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.PostSync(context.Background()).SyncSyncReq(syncSyncReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.PostSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSync`: SyncSyncView
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.PostSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncSyncReq** | [**SyncSyncReq**](SyncSyncReq.md) |  | 

### Return type

[**SyncSyncView**](SyncSyncView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSyncByIdRun

> SyncSyncQueued PostSyncByIdRun(ctx, id).Execute()

Reconciles one sync now — the manual re-sync, and the initial import for a link created without run=true.



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
	id := "sync_1" // string | ID is the sync to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.PostSyncByIdRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.PostSyncByIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSyncByIdRun`: SyncSyncQueued
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

[**SyncSyncQueued**](SyncSyncQueued.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

