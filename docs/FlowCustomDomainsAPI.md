# \FlowCustomDomainsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowCreateCustomDomain**](FlowCustomDomainsAPI.md#FlowCreateCustomDomain) | **Post** /v1/flow/custom-domains | Add a custom domain (EE)
[**FlowDeleteCustomDomain**](FlowCustomDomainsAPI.md#FlowDeleteCustomDomain) | **Delete** /v1/flow/custom-domains/{id} | Remove a custom domain (EE)
[**FlowListCustomDomains**](FlowCustomDomainsAPI.md#FlowListCustomDomains) | **Get** /v1/flow/custom-domains | List custom domains (EE)



## FlowCreateCustomDomain

> map[string]interface{} FlowCreateCustomDomain(ctx).FlowCreateCustomDomainRequest(flowCreateCustomDomainRequest).Execute()

Add a custom domain (EE)

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
	flowCreateCustomDomainRequest := *openapiclient.NewFlowCreateCustomDomainRequest("Domain_example") // FlowCreateCustomDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowCustomDomainsAPI.FlowCreateCustomDomain(context.Background()).FlowCreateCustomDomainRequest(flowCreateCustomDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowCustomDomainsAPI.FlowCreateCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateCustomDomain`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowCustomDomainsAPI.FlowCreateCustomDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowCreateCustomDomainRequest** | [**FlowCreateCustomDomainRequest**](FlowCreateCustomDomainRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowDeleteCustomDomain

> FlowDeleteCustomDomain(ctx, id).Execute()

Remove a custom domain (EE)

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowCustomDomainsAPI.FlowDeleteCustomDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowCustomDomainsAPI.FlowDeleteCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowDeleteCustomDomainRequest struct via the builder pattern


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


## FlowListCustomDomains

> map[string]interface{} FlowListCustomDomains(ctx).Execute()

List custom domains (EE)

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
	resp, r, err := apiClient.FlowCustomDomainsAPI.FlowListCustomDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowCustomDomainsAPI.FlowListCustomDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListCustomDomains`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowCustomDomainsAPI.FlowListCustomDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListCustomDomainsRequest struct via the builder pattern


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

