# \BalancersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBalancersById**](BalancersAPI.md#DeleteBalancersById) | **Delete** /v1/balancers/{id} | Removes one of the caller org&#39;s load balancers and answers 204.
[**GetBalancers**](BalancersAPI.md#GetBalancers) | **Get** /v1/balancers | Returns every load balancer the caller&#39;s org owns, under the friendly names the org created them with.
[**GetBalancersById**](BalancersAPI.md#GetBalancersById) | **Get** /v1/balancers/{id} | Returns one of the caller org&#39;s load balancers by id.
[**PostBalancers**](BalancersAPI.md#PostBalancers) | **Post** /v1/balancers | Creates a load balancer in the caller&#39;s org namespace and answers 201 with it.



## DeleteBalancersById

> DeleteBalancersById(ctx, id).Execute()

Removes one of the caller org's load balancers and answers 204.



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
	id := "id_example" // string | ID is the DigitalOcean resource id (a UUID), from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BalancersAPI.DeleteBalancersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.DeleteBalancersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DigitalOcean resource id (a UUID), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBalancersByIdRequest struct via the builder pattern


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


## GetBalancers

> LbList GetBalancers(ctx).Execute()

Returns every load balancer the caller's org owns, under the friendly names the org created them with.



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
	resp, r, err := apiClient.BalancersAPI.GetBalancers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.GetBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalancers`: LbList
	fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.GetBalancers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancersRequest struct via the builder pattern


### Return type

[**LbList**](LbList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalancersById

> LbView GetBalancersById(ctx, id).Execute()

Returns one of the caller org's load balancers by id.



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
	id := "id_example" // string | ID is the DigitalOcean resource id (a UUID), from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BalancersAPI.GetBalancersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.GetBalancersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalancersById`: LbView
	fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.GetBalancersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DigitalOcean resource id (a UUID), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LbView**](LbView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBalancers

> LbView PostBalancers(ctx).CreateLBReq(createLBReq).Execute()

Creates a load balancer in the caller's org namespace and answers 201 with it.



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
	createLBReq := *openapiclient.NewCreateLBReq() // CreateLBReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BalancersAPI.PostBalancers(context.Background()).CreateLBReq(createLBReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.PostBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBalancers`: LbView
	fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.PostBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLBReq** | [**CreateLBReq**](CreateLBReq.md) |  | 

### Return type

[**LbView**](LbView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

