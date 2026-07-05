# \EngineServeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EngineCreateServingEndpoint**](EngineServeAPI.md#EngineCreateServingEndpoint) | **Post** /v1/engine/serve/endpoints | Create serving endpoint
[**EngineDeleteServingEndpoint**](EngineServeAPI.md#EngineDeleteServingEndpoint) | **Delete** /v1/engine/serve/endpoints/{name} | Delete serving endpoint
[**EngineGetServingEndpoint**](EngineServeAPI.md#EngineGetServingEndpoint) | **Get** /v1/engine/serve/endpoints/{name} | Get serving endpoint
[**EngineGetServingMetrics**](EngineServeAPI.md#EngineGetServingMetrics) | **Get** /v1/engine/serve/endpoints/{name}/metrics | Get serving endpoint metrics
[**EngineListServingEndpoints**](EngineServeAPI.md#EngineListServingEndpoints) | **Get** /v1/engine/serve/endpoints | List serving endpoints
[**EngineUpdateServingEndpoint**](EngineServeAPI.md#EngineUpdateServingEndpoint) | **Put** /v1/engine/serve/endpoints/{name} | Update serving endpoint



## EngineCreateServingEndpoint

> EngineServingEndpoint EngineCreateServingEndpoint(ctx).EngineServingEndpointCreate(engineServingEndpointCreate).Execute()

Create serving endpoint

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
	engineServingEndpointCreate := *openapiclient.NewEngineServingEndpointCreate("Name_example", "Model_example") // EngineServingEndpointCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineServeAPI.EngineCreateServingEndpoint(context.Background()).EngineServingEndpointCreate(engineServingEndpointCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineServeAPI.EngineCreateServingEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineCreateServingEndpoint`: EngineServingEndpoint
	fmt.Fprintf(os.Stdout, "Response from `EngineServeAPI.EngineCreateServingEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineCreateServingEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineServingEndpointCreate** | [**EngineServingEndpointCreate**](EngineServingEndpointCreate.md) |  | 

### Return type

[**EngineServingEndpoint**](EngineServingEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineDeleteServingEndpoint

> map[string]interface{} EngineDeleteServingEndpoint(ctx, name).Execute()

Delete serving endpoint

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
	resp, r, err := apiClient.EngineServeAPI.EngineDeleteServingEndpoint(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineServeAPI.EngineDeleteServingEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineDeleteServingEndpoint`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineServeAPI.EngineDeleteServingEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineDeleteServingEndpointRequest struct via the builder pattern


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


## EngineGetServingEndpoint

> EngineServingEndpoint EngineGetServingEndpoint(ctx, name).Execute()

Get serving endpoint

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
	resp, r, err := apiClient.EngineServeAPI.EngineGetServingEndpoint(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineServeAPI.EngineGetServingEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetServingEndpoint`: EngineServingEndpoint
	fmt.Fprintf(os.Stdout, "Response from `EngineServeAPI.EngineGetServingEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetServingEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngineServingEndpoint**](EngineServingEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineGetServingMetrics

> []EngineServingMetrics EngineGetServingMetrics(ctx, name).From(from).To(to).Granularity(granularity).Execute()

Get serving endpoint metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	granularity := "granularity_example" // string |  (optional) (default to "hour")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineServeAPI.EngineGetServingMetrics(context.Background(), name).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineServeAPI.EngineGetServingMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetServingMetrics`: []EngineServingMetrics
	fmt.Fprintf(os.Stdout, "Response from `EngineServeAPI.EngineGetServingMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetServingMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **granularity** | **string** |  | [default to &quot;hour&quot;]

### Return type

[**[]EngineServingMetrics**](EngineServingMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineListServingEndpoints

> EngineListServingEndpoints200Response EngineListServingEndpoints(ctx).Status(status).Page(page).PageSize(pageSize).Execute()

List serving endpoints

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
	status := "status_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineServeAPI.EngineListServingEndpoints(context.Background()).Status(status).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineServeAPI.EngineListServingEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineListServingEndpoints`: EngineListServingEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `EngineServeAPI.EngineListServingEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineListServingEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**EngineListServingEndpoints200Response**](EngineListServingEndpoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineUpdateServingEndpoint

> EngineServingEndpoint EngineUpdateServingEndpoint(ctx, name).EngineUpdateServingEndpointRequest(engineUpdateServingEndpointRequest).Execute()

Update serving endpoint

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
	engineUpdateServingEndpointRequest := *openapiclient.NewEngineUpdateServingEndpointRequest() // EngineUpdateServingEndpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineServeAPI.EngineUpdateServingEndpoint(context.Background(), name).EngineUpdateServingEndpointRequest(engineUpdateServingEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineServeAPI.EngineUpdateServingEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineUpdateServingEndpoint`: EngineServingEndpoint
	fmt.Fprintf(os.Stdout, "Response from `EngineServeAPI.EngineUpdateServingEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineUpdateServingEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engineUpdateServingEndpointRequest** | [**EngineUpdateServingEndpointRequest**](EngineUpdateServingEndpointRequest.md) |  | 

### Return type

[**EngineServingEndpoint**](EngineServingEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

