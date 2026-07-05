# \DoLoadBalancersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoCreateLoadBalancer**](DoLoadBalancersAPI.md#DoCreateLoadBalancer) | **Post** /v1/load-balancers | Create a load balancer
[**DoDeleteLoadBalancer**](DoLoadBalancersAPI.md#DoDeleteLoadBalancer) | **Delete** /v1/load-balancers/{id} | Delete one load balancer (owned)
[**DoGetLoadBalancer**](DoLoadBalancersAPI.md#DoGetLoadBalancer) | **Get** /v1/load-balancers/{id} | Get one load balancer (owned)
[**DoListLoadBalancers**](DoLoadBalancersAPI.md#DoListLoadBalancers) | **Get** /v1/load-balancers | List the caller&#39;s load balancers



## DoCreateLoadBalancer

> DoLoadBalancer DoCreateLoadBalancer(ctx).DoLoadBalancerCreate(doLoadBalancerCreate).Execute()

Create a load balancer

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
	doLoadBalancerCreate := *openapiclient.NewDoLoadBalancerCreate("Name_example", "Region_example") // DoLoadBalancerCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DoLoadBalancersAPI.DoCreateLoadBalancer(context.Background()).DoLoadBalancerCreate(doLoadBalancerCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoLoadBalancersAPI.DoCreateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoCreateLoadBalancer`: DoLoadBalancer
	fmt.Fprintf(os.Stdout, "Response from `DoLoadBalancersAPI.DoCreateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDoCreateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **doLoadBalancerCreate** | [**DoLoadBalancerCreate**](DoLoadBalancerCreate.md) |  | 

### Return type

[**DoLoadBalancer**](DoLoadBalancer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoDeleteLoadBalancer

> DoDeleteLoadBalancer(ctx, id).Execute()

Delete one load balancer (owned)

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
	id := "id_example" // string | DO load balancer id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DoLoadBalancersAPI.DoDeleteLoadBalancer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoLoadBalancersAPI.DoDeleteLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DO load balancer id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoDeleteLoadBalancerRequest struct via the builder pattern


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


## DoGetLoadBalancer

> DoLoadBalancer DoGetLoadBalancer(ctx, id).Execute()

Get one load balancer (owned)

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
	id := "id_example" // string | DO load balancer id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DoLoadBalancersAPI.DoGetLoadBalancer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoLoadBalancersAPI.DoGetLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoGetLoadBalancer`: DoLoadBalancer
	fmt.Fprintf(os.Stdout, "Response from `DoLoadBalancersAPI.DoGetLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DO load balancer id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoGetLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DoLoadBalancer**](DoLoadBalancer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoListLoadBalancers

> DoListLoadBalancers200Response DoListLoadBalancers(ctx).Execute()

List the caller's load balancers

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
	resp, r, err := apiClient.DoLoadBalancersAPI.DoListLoadBalancers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoLoadBalancersAPI.DoListLoadBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoListLoadBalancers`: DoListLoadBalancers200Response
	fmt.Fprintf(os.Stdout, "Response from `DoLoadBalancersAPI.DoListLoadBalancers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDoListLoadBalancersRequest struct via the builder pattern


### Return type

[**DoListLoadBalancers200Response**](DoListLoadBalancers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

