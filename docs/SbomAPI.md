# \SbomAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSbomByWildcard1**](SbomAPI.md#GetSbomByWildcard1) | **Get** /v1/sbom/{wildcard1} | Resolve everything inside a container image
[**GetSbomHealth**](SbomAPI.md#GetSbomHealth) | **Get** /v1/sbom/health | Health is a pure liveness probe: the service is up; datastore reflects whether the datastore store is connected.
[**PostSbom**](SbomAPI.md#PostSbom) | **Post** /v1/sbom | Ingest persists a CycloneDX SBOM&#39;s components keyed by image digest.



## GetSbomByWildcard1

> GetSbomByWildcard1(ctx, wildcard1).Execute()

Resolve everything inside a container image



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.GetSbomByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomByWildcard1Request struct via the builder pattern


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


## GetSbomHealth

> SbomHealth GetSbomHealth(ctx).Execute()

Health is a pure liveness probe: the service is up; datastore reflects whether the datastore store is connected.



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
	resp, r, err := apiClient.SbomAPI.GetSbomHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomHealth`: SbomHealth
	fmt.Fprintf(os.Stdout, "Response from `SbomAPI.GetSbomHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomHealthRequest struct via the builder pattern


### Return type

[**SbomHealth**](SbomHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSbom

> SbomIngested PostSbom(ctx).SbomIngest(sbomIngest).Execute()

Ingest persists a CycloneDX SBOM's components keyed by image digest.



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
	sbomIngest := *openapiclient.NewSbomIngest() // SbomIngest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SbomAPI.PostSbom(context.Background()).SbomIngest(sbomIngest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.PostSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSbom`: SbomIngested
	fmt.Fprintf(os.Stdout, "Response from `SbomAPI.PostSbom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sbomIngest** | [**SbomIngest**](SbomIngest.md) |  | 

### Return type

[**SbomIngested**](SbomIngested.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

