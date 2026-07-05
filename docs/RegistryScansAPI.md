# \RegistryScansAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryGetScanReport**](RegistryScansAPI.md#RegistryGetScanReport) | **Get** /v1/registry/projects/{name}/repositories/{repo}/artifacts/{digest}/scan/report | Get scan report
[**RegistryTriggerScan**](RegistryScansAPI.md#RegistryTriggerScan) | **Post** /v1/registry/projects/{name}/repositories/{repo}/artifacts/{digest}/scan | Trigger vulnerability scan



## RegistryGetScanReport

> RegistryScanReport RegistryGetScanReport(ctx, name, repo, digest).Execute()

Get scan report

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
	repo := "repo_example" // string | 
	digest := "digest_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryScansAPI.RegistryGetScanReport(context.Background(), name, repo, digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryScansAPI.RegistryGetScanReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryGetScanReport`: RegistryScanReport
	fmt.Fprintf(os.Stdout, "Response from `RegistryScansAPI.RegistryGetScanReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 
**digest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryGetScanReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RegistryScanReport**](RegistryScanReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryTriggerScan

> RegistryTriggerScan(ctx, name, repo, digest).Execute()

Trigger vulnerability scan

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
	repo := "repo_example" // string | 
	digest := "digest_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegistryScansAPI.RegistryTriggerScan(context.Background(), name, repo, digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryScansAPI.RegistryTriggerScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 
**digest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryTriggerScanRequest struct via the builder pattern


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

