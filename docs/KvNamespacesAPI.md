# \KvNamespacesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvCreateNamespace**](KvNamespacesAPI.md#KvCreateNamespace) | **Post** /v1/kv/namespaces | Create namespace
[**KvDeleteNamespace**](KvNamespacesAPI.md#KvDeleteNamespace) | **Delete** /v1/kv/namespaces/{name} | Delete namespace and all keys
[**KvGetNamespace**](KvNamespacesAPI.md#KvGetNamespace) | **Get** /v1/kv/namespaces/{name} | Get namespace
[**KvListNamespaces**](KvNamespacesAPI.md#KvListNamespaces) | **Get** /v1/kv/namespaces | List namespaces



## KvCreateNamespace

> KvNamespace KvCreateNamespace(ctx).KvCreateNamespaceRequest(kvCreateNamespaceRequest).Execute()

Create namespace

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
	kvCreateNamespaceRequest := *openapiclient.NewKvCreateNamespaceRequest("Name_example") // KvCreateNamespaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvNamespacesAPI.KvCreateNamespace(context.Background()).KvCreateNamespaceRequest(kvCreateNamespaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvNamespacesAPI.KvCreateNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvCreateNamespace`: KvNamespace
	fmt.Fprintf(os.Stdout, "Response from `KvNamespacesAPI.KvCreateNamespace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvCreateNamespaceRequest** | [**KvCreateNamespaceRequest**](KvCreateNamespaceRequest.md) |  | 

### Return type

[**KvNamespace**](KvNamespace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvDeleteNamespace

> map[string]interface{} KvDeleteNamespace(ctx, name).Execute()

Delete namespace and all keys

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvNamespacesAPI.KvDeleteNamespace(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvNamespacesAPI.KvDeleteNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvDeleteNamespace`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KvNamespacesAPI.KvDeleteNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvDeleteNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvGetNamespace

> KvNamespace KvGetNamespace(ctx, name).Execute()

Get namespace

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvNamespacesAPI.KvGetNamespace(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvNamespacesAPI.KvGetNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvGetNamespace`: KvNamespace
	fmt.Fprintf(os.Stdout, "Response from `KvNamespacesAPI.KvGetNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvGetNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KvNamespace**](KvNamespace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvListNamespaces

> []KvNamespace KvListNamespaces(ctx).Execute()

List namespaces

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
	resp, r, err := apiClient.KvNamespacesAPI.KvListNamespaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvNamespacesAPI.KvListNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvListNamespaces`: []KvNamespace
	fmt.Fprintf(os.Stdout, "Response from `KvNamespacesAPI.KvListNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKvListNamespacesRequest struct via the builder pattern


### Return type

[**[]KvNamespace**](KvNamespace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

