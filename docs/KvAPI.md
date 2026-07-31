# \KvAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1KvName**](KvAPI.md#CloudDeleteV1KvName) | **Delete** /v1/kv/{name} | DropKV deprovisions one Hanzo KV store.
[**CloudGetV1Kv**](KvAPI.md#CloudGetV1Kv) | **Get** /v1/kv | ListKV lists the caller org&#39;s Hanzo KV stores.
[**CloudGetV1KvName**](KvAPI.md#CloudGetV1KvName) | **Get** /v1/kv/{name} | GetKV returns one Hanzo KV store&#39;s metadata.
[**CloudPostV1Kv**](KvAPI.md#CloudPostV1Kv) | **Post** /v1/kv | 



## CloudDeleteV1KvName

> CloudDeleteV1KvName(ctx, name).Execute()

DropKV deprovisions one Hanzo KV store.



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
	name := "sessions" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KvAPI.CloudDeleteV1KvName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.CloudDeleteV1KvName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1KvNameRequest struct via the builder pattern


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


## CloudGetV1Kv

> []CloudProvisionedSummary CloudGetV1Kv(ctx).Execute()

ListKV lists the caller org's Hanzo KV stores.



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
	resp, r, err := apiClient.KvAPI.CloudGetV1Kv(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.CloudGetV1Kv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Kv`: []CloudProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `KvAPI.CloudGetV1Kv`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1KvRequest struct via the builder pattern


### Return type

[**[]CloudProvisionedSummary**](CloudProvisionedSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1KvName

> CloudProvisionedResource CloudGetV1KvName(ctx, name).Execute()

GetKV returns one Hanzo KV store's metadata.



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
	name := "sessions" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvAPI.CloudGetV1KvName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.CloudGetV1KvName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1KvName`: CloudProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `KvAPI.CloudGetV1KvName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1KvNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProvisionedResource**](CloudProvisionedResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Kv

> CloudProvisionResult CloudPostV1Kv(ctx).CloudProvisionRequest(cloudProvisionRequest).Execute()



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
	cloudProvisionRequest := *openapiclient.NewCloudProvisionRequest() // CloudProvisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvAPI.CloudPostV1Kv(context.Background()).CloudProvisionRequest(cloudProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvAPI.CloudPostV1Kv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Kv`: CloudProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `KvAPI.CloudPostV1Kv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1KvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvisionRequest** | [**CloudProvisionRequest**](CloudProvisionRequest.md) |  | 

### Return type

[**CloudProvisionResult**](CloudProvisionResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

