# \VpcsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1VpcsId**](VpcsAPI.md#CloudDeleteV1VpcsId) | **Delete** /v1/vpcs/{id} | DeleteVpc removes one of the caller org&#39;s VPCs and answers 204.
[**CloudGetV1Vpcs**](VpcsAPI.md#CloudGetV1Vpcs) | **Get** /v1/vpcs | ListVpcs returns every VPC the caller&#39;s org owns, under the friendly names the org created them with.
[**CloudGetV1VpcsId**](VpcsAPI.md#CloudGetV1VpcsId) | **Get** /v1/vpcs/{id} | GetVpc returns one of the caller org&#39;s VPCs by id.
[**CloudPostV1Vpcs**](VpcsAPI.md#CloudPostV1Vpcs) | **Post** /v1/vpcs | CreateVpc creates a VPC in the caller&#39;s org namespace and answers 201 with it.



## CloudDeleteV1VpcsId

> CloudDeleteV1VpcsId(ctx, id).Execute()

DeleteVpc removes one of the caller org's VPCs and answers 204.



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
	r, err := apiClient.VpcsAPI.CloudDeleteV1VpcsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcsAPI.CloudDeleteV1VpcsId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1VpcsIdRequest struct via the builder pattern


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


## CloudGetV1Vpcs

> CloudVpcList CloudGetV1Vpcs(ctx).Execute()

ListVpcs returns every VPC the caller's org owns, under the friendly names the org created them with.



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
	resp, r, err := apiClient.VpcsAPI.CloudGetV1Vpcs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcsAPI.CloudGetV1Vpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Vpcs`: CloudVpcList
	fmt.Fprintf(os.Stdout, "Response from `VpcsAPI.CloudGetV1Vpcs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1VpcsRequest struct via the builder pattern


### Return type

[**CloudVpcList**](CloudVpcList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1VpcsId

> CloudVpcView CloudGetV1VpcsId(ctx, id).Execute()

GetVpc returns one of the caller org's VPCs by id.



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
	resp, r, err := apiClient.VpcsAPI.CloudGetV1VpcsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcsAPI.CloudGetV1VpcsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1VpcsId`: CloudVpcView
	fmt.Fprintf(os.Stdout, "Response from `VpcsAPI.CloudGetV1VpcsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DigitalOcean resource id (a UUID), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1VpcsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudVpcView**](CloudVpcView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Vpcs

> CloudVpcView CloudPostV1Vpcs(ctx).CloudCreateVPCReq(cloudCreateVPCReq).Execute()

CreateVpc creates a VPC in the caller's org namespace and answers 201 with it.



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
	cloudCreateVPCReq := *openapiclient.NewCloudCreateVPCReq() // CloudCreateVPCReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcsAPI.CloudPostV1Vpcs(context.Background()).CloudCreateVPCReq(cloudCreateVPCReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcsAPI.CloudPostV1Vpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Vpcs`: CloudVpcView
	fmt.Fprintf(os.Stdout, "Response from `VpcsAPI.CloudPostV1Vpcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1VpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateVPCReq** | [**CloudCreateVPCReq**](CloudCreateVPCReq.md) |  | 

### Return type

[**CloudVpcView**](CloudVpcView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

