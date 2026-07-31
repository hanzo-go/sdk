# \SyncAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1SyncId**](SyncAPI.md#CloudDeleteV1SyncId) | **Delete** /v1/sync/{id} | Delete removes one sync and tears down the outbound mirror it derived, answering 204.
[**CloudGetV1Sync**](SyncAPI.md#CloudGetV1Sync) | **Get** /v1/sync | List returns every sync link the caller&#39;s org has, each with its two endpoints, its direction and trigger policy, and the time it last reconciled.
[**CloudGetV1SyncId**](SyncAPI.md#CloudGetV1SyncId) | **Get** /v1/sync/{id} | Get returns one sync by id.
[**CloudPatchV1SyncId**](SyncAPI.md#CloudPatchV1SyncId) | **Patch** /v1/sync/{id} | Patch updates one sync&#39;s mutable policy — direction, trigger and actor — in place.
[**CloudPostV1Sync**](SyncAPI.md#CloudPostV1Sync) | **Post** /v1/sync | Create declares a sync between two endpoints and returns it.
[**CloudPostV1SyncIdRun**](SyncAPI.md#CloudPostV1SyncIdRun) | **Post** /v1/sync/{id}/run | Run reconciles one sync now — the manual re-sync, and the initial import for a link created without run&#x3D;true.



## CloudDeleteV1SyncId

> CloudDeleteV1SyncId(ctx, id).Execute()

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
	r, err := apiClient.SyncAPI.CloudDeleteV1SyncId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.CloudDeleteV1SyncId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1SyncIdRequest struct via the builder pattern


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


## CloudGetV1Sync

> CloudSyncList CloudGetV1Sync(ctx).Execute()

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
	resp, r, err := apiClient.SyncAPI.CloudGetV1Sync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.CloudGetV1Sync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Sync`: CloudSyncList
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.CloudGetV1Sync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SyncRequest struct via the builder pattern


### Return type

[**CloudSyncList**](CloudSyncList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1SyncId

> CloudSyncView CloudGetV1SyncId(ctx, id).Execute()

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
	resp, r, err := apiClient.SyncAPI.CloudGetV1SyncId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.CloudGetV1SyncId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1SyncId`: CloudSyncView
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.CloudGetV1SyncId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sync to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SyncIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSyncView**](CloudSyncView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1SyncId

> CloudSyncView CloudPatchV1SyncId(ctx, id).CloudPatchSyncIn(cloudPatchSyncIn).Execute()

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
	cloudPatchSyncIn := *openapiclient.NewCloudPatchSyncIn() // CloudPatchSyncIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.CloudPatchV1SyncId(context.Background(), id).CloudPatchSyncIn(cloudPatchSyncIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.CloudPatchV1SyncId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1SyncId`: CloudSyncView
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.CloudPatchV1SyncId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sync to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1SyncIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPatchSyncIn** | [**CloudPatchSyncIn**](CloudPatchSyncIn.md) |  | 

### Return type

[**CloudSyncView**](CloudSyncView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Sync

> CloudSyncView CloudPostV1Sync(ctx).CloudSyncReq(cloudSyncReq).Execute()

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
	cloudSyncReq := *openapiclient.NewCloudSyncReq() // CloudSyncReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.CloudPostV1Sync(context.Background()).CloudSyncReq(cloudSyncReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.CloudPostV1Sync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Sync`: CloudSyncView
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.CloudPostV1Sync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1SyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSyncReq** | [**CloudSyncReq**](CloudSyncReq.md) |  | 

### Return type

[**CloudSyncView**](CloudSyncView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1SyncIdRun

> CloudSyncQueued CloudPostV1SyncIdRun(ctx, id).Execute()

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
	resp, r, err := apiClient.SyncAPI.CloudPostV1SyncIdRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.CloudPostV1SyncIdRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1SyncIdRun`: CloudSyncQueued
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.CloudPostV1SyncIdRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sync to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1SyncIdRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSyncQueued**](CloudSyncQueued.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

