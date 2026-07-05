# \PaasDOKSAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasAddDOKSNodePool**](PaasDOKSAPI.md#PaasAddDOKSNodePool) | **Post** /v1/paas/cluster/doks/{orgId}/node-pools | Add node pool
[**PaasDeleteDOKSNodePool**](PaasDOKSAPI.md#PaasDeleteDOKSNodePool) | **Delete** /v1/paas/cluster/doks/{orgId}/node-pools/{poolId} | Delete node pool
[**PaasDestroyDOKS**](PaasDOKSAPI.md#PaasDestroyDOKS) | **Delete** /v1/paas/cluster/doks/{orgId} | Destroy DOKS cluster
[**PaasGetDOKSFleet**](PaasDOKSAPI.md#PaasGetDOKSFleet) | **Get** /v1/paas/cluster/doks/fleet | Fleet overview (all org clusters)
[**PaasGetDOKSKubeconfig**](PaasDOKSAPI.md#PaasGetDOKSKubeconfig) | **Get** /v1/paas/cluster/doks/{orgId}/kubeconfig | Download kubeconfig
[**PaasGetDOKSOptions**](PaasDOKSAPI.md#PaasGetDOKSOptions) | **Get** /v1/paas/cluster/doks/options | Available regions and node sizes
[**PaasGetDOKSPricing**](PaasDOKSAPI.md#PaasGetDOKSPricing) | **Get** /v1/paas/cluster/doks/pricing/{sizeSlug} | Get droplet pricing
[**PaasGetDOKSStatus**](PaasDOKSAPI.md#PaasGetDOKSStatus) | **Get** /v1/paas/cluster/doks/{orgId}/status | Get cluster status (polls DO API)
[**PaasListDOKSNodePools**](PaasDOKSAPI.md#PaasListDOKSNodePools) | **Get** /v1/paas/cluster/doks/{orgId}/node-pools | List node pools
[**PaasProvisionDOKS**](PaasDOKSAPI.md#PaasProvisionDOKS) | **Post** /v1/paas/cluster/doks/provision | Provision new DOKS cluster for org
[**PaasUpdateDOKSNodePool**](PaasDOKSAPI.md#PaasUpdateDOKSNodePool) | **Put** /v1/paas/cluster/doks/{orgId}/node-pools/{poolId} | Update node pool
[**PaasUpgradeDOKSHA**](PaasDOKSAPI.md#PaasUpgradeDOKSHA) | **Post** /v1/paas/cluster/doks/{orgId}/upgrade-ha | Upgrade to HA control plane



## PaasAddDOKSNodePool

> map[string]interface{} PaasAddDOKSNodePool(ctx, orgId).PaasNodePool(paasNodePool).Execute()

Add node pool

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
	orgId := "orgId_example" // string | 
	paasNodePool := *openapiclient.NewPaasNodePool() // PaasNodePool | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasAddDOKSNodePool(context.Background(), orgId).PaasNodePool(paasNodePool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasAddDOKSNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasAddDOKSNodePool`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasAddDOKSNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasAddDOKSNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paasNodePool** | [**PaasNodePool**](PaasNodePool.md) |  | 

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


## PaasDeleteDOKSNodePool

> map[string]interface{} PaasDeleteDOKSNodePool(ctx, orgId, poolId).Execute()

Delete node pool

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
	orgId := "orgId_example" // string | 
	poolId := "poolId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasDeleteDOKSNodePool(context.Background(), orgId, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasDeleteDOKSNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasDeleteDOKSNodePool`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasDeleteDOKSNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasDeleteDOKSNodePoolRequest struct via the builder pattern


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


## PaasDestroyDOKS

> map[string]interface{} PaasDestroyDOKS(ctx, orgId).Confirm(confirm).Execute()

Destroy DOKS cluster

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
	orgId := "orgId_example" // string | 
	confirm := true // bool | Must be true to confirm destruction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasDestroyDOKS(context.Background(), orgId).Confirm(confirm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasDestroyDOKS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasDestroyDOKS`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasDestroyDOKS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasDestroyDOKSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirm** | **bool** | Must be true to confirm destruction | 

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


## PaasGetDOKSFleet

> PaasFleetOverview PaasGetDOKSFleet(ctx).Execute()

Fleet overview (all org clusters)

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
	resp, r, err := apiClient.PaasDOKSAPI.PaasGetDOKSFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasGetDOKSFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetDOKSFleet`: PaasFleetOverview
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasGetDOKSFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetDOKSFleetRequest struct via the builder pattern


### Return type

[**PaasFleetOverview**](PaasFleetOverview.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasGetDOKSKubeconfig

> string PaasGetDOKSKubeconfig(ctx, orgId).Execute()

Download kubeconfig

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasGetDOKSKubeconfig(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasGetDOKSKubeconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetDOKSKubeconfig`: string
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasGetDOKSKubeconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetDOKSKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasGetDOKSOptions

> PaasGetDOKSOptions200Response PaasGetDOKSOptions(ctx).Execute()

Available regions and node sizes

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
	resp, r, err := apiClient.PaasDOKSAPI.PaasGetDOKSOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasGetDOKSOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetDOKSOptions`: PaasGetDOKSOptions200Response
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasGetDOKSOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetDOKSOptionsRequest struct via the builder pattern


### Return type

[**PaasGetDOKSOptions200Response**](PaasGetDOKSOptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasGetDOKSPricing

> PaasGetDOKSPricing200Response PaasGetDOKSPricing(ctx, sizeSlug).Execute()

Get droplet pricing

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
	sizeSlug := "sizeSlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasGetDOKSPricing(context.Background(), sizeSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasGetDOKSPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetDOKSPricing`: PaasGetDOKSPricing200Response
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasGetDOKSPricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sizeSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetDOKSPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaasGetDOKSPricing200Response**](PaasGetDOKSPricing200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasGetDOKSStatus

> PaasDOKSCluster PaasGetDOKSStatus(ctx, orgId).Execute()

Get cluster status (polls DO API)

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasGetDOKSStatus(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasGetDOKSStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetDOKSStatus`: PaasDOKSCluster
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasGetDOKSStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetDOKSStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaasDOKSCluster**](PaasDOKSCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasListDOKSNodePools

> []PaasNodePool PaasListDOKSNodePools(ctx, orgId).Execute()

List node pools

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasListDOKSNodePools(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasListDOKSNodePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasListDOKSNodePools`: []PaasNodePool
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasListDOKSNodePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasListDOKSNodePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PaasNodePool**](PaasNodePool.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasProvisionDOKS

> PaasDOKSCluster PaasProvisionDOKS(ctx).PaasProvisionDOKSRequest(paasProvisionDOKSRequest).Execute()

Provision new DOKS cluster for org

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
	paasProvisionDOKSRequest := *openapiclient.NewPaasProvisionDOKSRequest("OrgId_example", "Name_example", "Region_example", "Version_example", *openapiclient.NewPaasProvisionDOKSRequestNodePool("Name_example", "Size_example", int32(123))) // PaasProvisionDOKSRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasProvisionDOKS(context.Background()).PaasProvisionDOKSRequest(paasProvisionDOKSRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasProvisionDOKS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasProvisionDOKS`: PaasDOKSCluster
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasProvisionDOKS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaasProvisionDOKSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paasProvisionDOKSRequest** | [**PaasProvisionDOKSRequest**](PaasProvisionDOKSRequest.md) |  | 

### Return type

[**PaasDOKSCluster**](PaasDOKSCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasUpdateDOKSNodePool

> map[string]interface{} PaasUpdateDOKSNodePool(ctx, orgId, poolId).PaasUpdateDOKSNodePoolRequest(paasUpdateDOKSNodePoolRequest).Execute()

Update node pool

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
	orgId := "orgId_example" // string | 
	poolId := "poolId_example" // string | 
	paasUpdateDOKSNodePoolRequest := *openapiclient.NewPaasUpdateDOKSNodePoolRequest() // PaasUpdateDOKSNodePoolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasUpdateDOKSNodePool(context.Background(), orgId, poolId).PaasUpdateDOKSNodePoolRequest(paasUpdateDOKSNodePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasUpdateDOKSNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasUpdateDOKSNodePool`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasUpdateDOKSNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasUpdateDOKSNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **paasUpdateDOKSNodePoolRequest** | [**PaasUpdateDOKSNodePoolRequest**](PaasUpdateDOKSNodePoolRequest.md) |  | 

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


## PaasUpgradeDOKSHA

> map[string]interface{} PaasUpgradeDOKSHA(ctx, orgId).Execute()

Upgrade to HA control plane

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasDOKSAPI.PaasUpgradeDOKSHA(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasDOKSAPI.PaasUpgradeDOKSHA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasUpgradeDOKSHA`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasDOKSAPI.PaasUpgradeDOKSHA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasUpgradeDOKSHARequest struct via the builder pattern


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

