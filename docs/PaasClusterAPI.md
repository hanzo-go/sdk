# \PaasClusterAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasAddClusterDomain**](PaasClusterAPI.md#PaasAddClusterDomain) | **Post** /v1/paas/cluster/domains | Add cluster domain
[**PaasGetClusterInfo**](PaasClusterAPI.md#PaasGetClusterInfo) | **Get** /v1/paas/cluster/info | Get cluster info
[**PaasGetSetupStatus**](PaasClusterAPI.md#PaasGetSetupStatus) | **Get** /v1/paas/cluster/setup-status | Check cluster setup status
[**PaasListClusterDomains**](PaasClusterAPI.md#PaasListClusterDomains) | **Get** /v1/paas/cluster/domains | List cluster domains
[**PaasRemoveClusterDomain**](PaasClusterAPI.md#PaasRemoveClusterDomain) | **Delete** /v1/paas/cluster/domains/{domain} | Remove cluster domain



## PaasAddClusterDomain

> map[string]interface{} PaasAddClusterDomain(ctx).FlowCreateCustomDomainRequest(flowCreateCustomDomainRequest).Execute()

Add cluster domain

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
	resp, r, err := apiClient.PaasClusterAPI.PaasAddClusterDomain(context.Background()).FlowCreateCustomDomainRequest(flowCreateCustomDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasClusterAPI.PaasAddClusterDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasAddClusterDomain`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasClusterAPI.PaasAddClusterDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaasAddClusterDomainRequest struct via the builder pattern


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


## PaasGetClusterInfo

> PaasCluster PaasGetClusterInfo(ctx).Execute()

Get cluster info

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
	resp, r, err := apiClient.PaasClusterAPI.PaasGetClusterInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasClusterAPI.PaasGetClusterInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetClusterInfo`: PaasCluster
	fmt.Fprintf(os.Stdout, "Response from `PaasClusterAPI.PaasGetClusterInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetClusterInfoRequest struct via the builder pattern


### Return type

[**PaasCluster**](PaasCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasGetSetupStatus

> PaasGetSetupStatus200Response PaasGetSetupStatus(ctx).Execute()

Check cluster setup status

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
	resp, r, err := apiClient.PaasClusterAPI.PaasGetSetupStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasClusterAPI.PaasGetSetupStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetSetupStatus`: PaasGetSetupStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `PaasClusterAPI.PaasGetSetupStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetSetupStatusRequest struct via the builder pattern


### Return type

[**PaasGetSetupStatus200Response**](PaasGetSetupStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasListClusterDomains

> []string PaasListClusterDomains(ctx).Execute()

List cluster domains

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
	resp, r, err := apiClient.PaasClusterAPI.PaasListClusterDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasClusterAPI.PaasListClusterDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasListClusterDomains`: []string
	fmt.Fprintf(os.Stdout, "Response from `PaasClusterAPI.PaasListClusterDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaasListClusterDomainsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasRemoveClusterDomain

> map[string]interface{} PaasRemoveClusterDomain(ctx, domain).Execute()

Remove cluster domain

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
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasClusterAPI.PaasRemoveClusterDomain(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasClusterAPI.PaasRemoveClusterDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasRemoveClusterDomain`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasClusterAPI.PaasRemoveClusterDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasRemoveClusterDomainRequest struct via the builder pattern


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

