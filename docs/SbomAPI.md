# \SbomAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSbomByWildcard1**](SbomAPI.md#GetSbomByWildcard1) | **Get** /v1/sbom/{wildcard1} | Returns everything inside one container image, addressed by its digest or by its image ref.
[**GetSbomHealth**](SbomAPI.md#GetSbomHealth) | **Get** /v1/sbom/health | Health is a pure liveness probe: the service is up; datastore reflects whether the datastore store is connected.



## GetSbomByWildcard1

> SbomSbomView GetSbomByWildcard1(ctx, wildcard1).Execute()

Returns everything inside one container image, addressed by its digest or by its image ref.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	wildcard1 := "wildcard1_example" // string | Ref is the image to resolve, either its content-addressed digest (`sha256:…`) or its full reference (`oci.hanzo.ai/hanzo/cloud:v1`). Both are matched, so either spelling of one image answers the same components. It is the greedy tail of the address, so slashes, a tag and a digest all travel in it whole, and a percent-encoded ref is decoded before it is looked up. Empty is a 400, never a scan of the store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SbomAPI.GetSbomByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomByWildcard1`: SbomSbomView
	fmt.Fprintf(os.Stdout, "Response from `SbomAPI.GetSbomByWildcard1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** | Ref is the image to resolve, either its content-addressed digest (&#x60;sha256:…&#x60;) or its full reference (&#x60;oci.hanzo.ai/hanzo/cloud:v1&#x60;). Both are matched, so either spelling of one image answers the same components. It is the greedy tail of the address, so slashes, a tag and a digest all travel in it whole, and a percent-encoded ref is decoded before it is looked up. Empty is a 400, never a scan of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SbomSbomView**](SbomSbomView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomHealth

> SbomSbomHealth GetSbomHealth(ctx).Execute()

Health is a pure liveness probe: the service is up; datastore reflects whether the datastore store is connected.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SbomAPI.GetSbomHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomHealth`: SbomSbomHealth
	fmt.Fprintf(os.Stdout, "Response from `SbomAPI.GetSbomHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomHealthRequest struct via the builder pattern


### Return type

[**SbomSbomHealth**](SbomSbomHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

