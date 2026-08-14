# \VpcsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVpcsById**](VpcsAPI.md#DeleteVpcsById) | **Delete** /v1/vpcs/{id} | Removes one of the caller org&#39;s VPCs and answers 204.
[**GetVpcs**](VpcsAPI.md#GetVpcs) | **Get** /v1/vpcs | Returns every VPC the caller&#39;s org owns, under the friendly names the org created them with.
[**GetVpcsById**](VpcsAPI.md#GetVpcsById) | **Get** /v1/vpcs/{id} | Returns one of the caller org&#39;s VPCs by id.
[**PostVpcs**](VpcsAPI.md#PostVpcs) | **Post** /v1/vpcs | Creates a VPC in the caller&#39;s org namespace and answers 201 with it.



## DeleteVpcsById

> DeleteVpcsById(ctx, id).Execute()

Removes one of the caller org's VPCs and answers 204.



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
	r, err := apiClient.VpcsAPI.DeleteVpcsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcsAPI.DeleteVpcsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVpcsByIdRequest struct via the builder pattern


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


## GetVpcs

> VpcList GetVpcs(ctx).Execute()

Returns every VPC the caller's org owns, under the friendly names the org created them with.



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
	resp, r, err := apiClient.VpcsAPI.GetVpcs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcsAPI.GetVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpcs`: VpcList
	fmt.Fprintf(os.Stdout, "Response from `VpcsAPI.GetVpcs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcsRequest struct via the builder pattern


### Return type

[**VpcList**](VpcList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpcsById

> VpcView GetVpcsById(ctx, id).Execute()

Returns one of the caller org's VPCs by id.



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
	resp, r, err := apiClient.VpcsAPI.GetVpcsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcsAPI.GetVpcsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpcsById`: VpcView
	fmt.Fprintf(os.Stdout, "Response from `VpcsAPI.GetVpcsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DigitalOcean resource id (a UUID), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VpcView**](VpcView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVpcs

> VpcView PostVpcs(ctx).CreateVPCReq(createVPCReq).Execute()

Creates a VPC in the caller's org namespace and answers 201 with it.



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
	createVPCReq := *openapiclient.NewCreateVPCReq() // CreateVPCReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcsAPI.PostVpcs(context.Background()).CreateVPCReq(createVPCReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcsAPI.PostVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVpcs`: VpcView
	fmt.Fprintf(os.Stdout, "Response from `VpcsAPI.PostVpcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVPCReq** | [**CreateVPCReq**](CreateVPCReq.md) |  | 

### Return type

[**VpcView**](VpcView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

