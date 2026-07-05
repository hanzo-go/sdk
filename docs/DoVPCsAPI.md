# \DoVPCsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoCreateVpc**](DoVPCsAPI.md#DoCreateVpc) | **Post** /v1/vpcs | Create a VPC
[**DoDeleteVpc**](DoVPCsAPI.md#DoDeleteVpc) | **Delete** /v1/vpcs/{id} | Delete one VPC (owned)
[**DoGetVpc**](DoVPCsAPI.md#DoGetVpc) | **Get** /v1/vpcs/{id} | Get one VPC (owned)
[**DoListVpcs**](DoVPCsAPI.md#DoListVpcs) | **Get** /v1/vpcs | List the caller&#39;s VPCs



## DoCreateVpc

> DoVpc DoCreateVpc(ctx).DoVpcCreate(doVpcCreate).Execute()

Create a VPC

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
	doVpcCreate := *openapiclient.NewDoVpcCreate("Name_example", "Region_example") // DoVpcCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DoVPCsAPI.DoCreateVpc(context.Background()).DoVpcCreate(doVpcCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoVPCsAPI.DoCreateVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoCreateVpc`: DoVpc
	fmt.Fprintf(os.Stdout, "Response from `DoVPCsAPI.DoCreateVpc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDoCreateVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **doVpcCreate** | [**DoVpcCreate**](DoVpcCreate.md) |  | 

### Return type

[**DoVpc**](DoVpc.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoDeleteVpc

> DoDeleteVpc(ctx, id).Execute()

Delete one VPC (owned)

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
	id := "id_example" // string | DO VPC id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DoVPCsAPI.DoDeleteVpc(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoVPCsAPI.DoDeleteVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DO VPC id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoDeleteVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoGetVpc

> DoVpc DoGetVpc(ctx, id).Execute()

Get one VPC (owned)

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
	id := "id_example" // string | DO VPC id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DoVPCsAPI.DoGetVpc(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoVPCsAPI.DoGetVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoGetVpc`: DoVpc
	fmt.Fprintf(os.Stdout, "Response from `DoVPCsAPI.DoGetVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DO VPC id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoGetVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DoVpc**](DoVpc.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoListVpcs

> DoListVpcs200Response DoListVpcs(ctx).Execute()

List the caller's VPCs

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
	resp, r, err := apiClient.DoVPCsAPI.DoListVpcs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoVPCsAPI.DoListVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoListVpcs`: DoListVpcs200Response
	fmt.Fprintf(os.Stdout, "Response from `DoVPCsAPI.DoListVpcs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDoListVpcsRequest struct via the builder pattern


### Return type

[**DoListVpcs200Response**](DoListVpcs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

