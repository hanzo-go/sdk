# \SbomAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1SbomByWildcard1**](SbomAPI.md#CloudGetV1SbomByWildcard1) | **Get** /v1/sbom/{wildcard1} | 
[**CloudGetV1SbomHealth**](SbomAPI.md#CloudGetV1SbomHealth) | **Get** /v1/sbom/health | Health is a pure liveness probe: the service is up; datastore reflects whether the datastore store is connected.
[**CloudPostV1Sbom**](SbomAPI.md#CloudPostV1Sbom) | **Post** /v1/sbom | Ingest persists a CycloneDX SBOM&#39;s components keyed by image digest.



## CloudGetV1SbomByWildcard1

> CloudGetV1SbomByWildcard1(ctx, wildcard1).Execute()



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
	r, err := apiClient.SbomAPI.CloudGetV1SbomByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.CloudGetV1SbomByWildcard1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1SbomByWildcard1Request struct via the builder pattern


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


## CloudGetV1SbomHealth

> CloudSbomHealth CloudGetV1SbomHealth(ctx).Execute()

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
	resp, r, err := apiClient.SbomAPI.CloudGetV1SbomHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.CloudGetV1SbomHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1SbomHealth`: CloudSbomHealth
	fmt.Fprintf(os.Stdout, "Response from `SbomAPI.CloudGetV1SbomHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SbomHealthRequest struct via the builder pattern


### Return type

[**CloudSbomHealth**](CloudSbomHealth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Sbom

> CloudSbomIngested CloudPostV1Sbom(ctx).CloudSbomIngest(cloudSbomIngest).Execute()

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
	cloudSbomIngest := *openapiclient.NewCloudSbomIngest() // CloudSbomIngest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SbomAPI.CloudPostV1Sbom(context.Background()).CloudSbomIngest(cloudSbomIngest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.CloudPostV1Sbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Sbom`: CloudSbomIngested
	fmt.Fprintf(os.Stdout, "Response from `SbomAPI.CloudPostV1Sbom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1SbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSbomIngest** | [**CloudSbomIngest**](CloudSbomIngest.md) |  | 

### Return type

[**CloudSbomIngested**](CloudSbomIngested.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

