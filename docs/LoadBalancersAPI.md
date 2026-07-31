# \LoadBalancersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1LoadBalancersId**](LoadBalancersAPI.md#CloudDeleteV1LoadBalancersId) | **Delete** /v1/load-balancers/{id} | DeleteLoadBalancer removes one of the caller org&#39;s load balancers and answers 204.
[**CloudGetV1LoadBalancers**](LoadBalancersAPI.md#CloudGetV1LoadBalancers) | **Get** /v1/load-balancers | ListLoadBalancers returns every load balancer the caller&#39;s org owns, under the friendly names the org created them with.
[**CloudGetV1LoadBalancersId**](LoadBalancersAPI.md#CloudGetV1LoadBalancersId) | **Get** /v1/load-balancers/{id} | GetLoadBalancer returns one of the caller org&#39;s load balancers by id.
[**CloudPostV1LoadBalancers**](LoadBalancersAPI.md#CloudPostV1LoadBalancers) | **Post** /v1/load-balancers | CreateLoadBalancer creates a load balancer in the caller&#39;s org namespace and answers 201 with it.



## CloudDeleteV1LoadBalancersId

> CloudDeleteV1LoadBalancersId(ctx, id).Execute()

DeleteLoadBalancer removes one of the caller org's load balancers and answers 204.



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
	r, err := apiClient.LoadBalancersAPI.CloudDeleteV1LoadBalancersId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersAPI.CloudDeleteV1LoadBalancersId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1LoadBalancersIdRequest struct via the builder pattern


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


## CloudGetV1LoadBalancers

> CloudLbList CloudGetV1LoadBalancers(ctx).Execute()

ListLoadBalancers returns every load balancer the caller's org owns, under the friendly names the org created them with.



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
	resp, r, err := apiClient.LoadBalancersAPI.CloudGetV1LoadBalancers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersAPI.CloudGetV1LoadBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1LoadBalancers`: CloudLbList
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancersAPI.CloudGetV1LoadBalancers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1LoadBalancersRequest struct via the builder pattern


### Return type

[**CloudLbList**](CloudLbList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1LoadBalancersId

> CloudLbView CloudGetV1LoadBalancersId(ctx, id).Execute()

GetLoadBalancer returns one of the caller org's load balancers by id.



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
	resp, r, err := apiClient.LoadBalancersAPI.CloudGetV1LoadBalancersId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersAPI.CloudGetV1LoadBalancersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1LoadBalancersId`: CloudLbView
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancersAPI.CloudGetV1LoadBalancersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DigitalOcean resource id (a UUID), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1LoadBalancersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLbView**](CloudLbView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1LoadBalancers

> CloudLbView CloudPostV1LoadBalancers(ctx).CloudCreateLBReq(cloudCreateLBReq).Execute()

CreateLoadBalancer creates a load balancer in the caller's org namespace and answers 201 with it.



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
	cloudCreateLBReq := *openapiclient.NewCloudCreateLBReq() // CloudCreateLBReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancersAPI.CloudPostV1LoadBalancers(context.Background()).CloudCreateLBReq(cloudCreateLBReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersAPI.CloudPostV1LoadBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1LoadBalancers`: CloudLbView
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancersAPI.CloudPostV1LoadBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1LoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateLBReq** | [**CloudCreateLBReq**](CloudCreateLBReq.md) |  | 

### Return type

[**CloudLbView**](CloudLbView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

