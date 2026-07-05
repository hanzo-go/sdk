# \SecurityScansAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecuritySecurityGetScan**](SecurityScansAPI.md#SecuritySecurityGetScan) | **Get** /v1/security/scans/{id} | Get a scan with its findings
[**SecuritySecurityListScans**](SecurityScansAPI.md#SecuritySecurityListScans) | **Get** /v1/security/scans | List scans
[**SecuritySecuritySubmitScan**](SecurityScansAPI.md#SecuritySecuritySubmitScan) | **Post** /v1/security/scans | Submit a scan



## SecuritySecurityGetScan

> SecuritySecurityGetScan200Response SecuritySecurityGetScan(ctx, id).Execute()

Get a scan with its findings

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
	resp, r, err := apiClient.SecurityScansAPI.SecuritySecurityGetScan(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityScansAPI.SecuritySecurityGetScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySecurityGetScan`: SecuritySecurityGetScan200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityScansAPI.SecuritySecurityGetScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecurityGetScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecuritySecurityGetScan200Response**](SecuritySecurityGetScan200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySecurityListScans

> SecuritySecurityListScans200Response SecuritySecurityListScans(ctx).Limit(limit).Execute()

List scans

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
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityScansAPI.SecuritySecurityListScans(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityScansAPI.SecuritySecurityListScans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySecurityListScans`: SecuritySecurityListScans200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityScansAPI.SecuritySecurityListScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecurityListScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 

### Return type

[**SecuritySecurityListScans200Response**](SecuritySecurityListScans200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySecuritySubmitScan

> SecurityScan SecuritySecuritySubmitScan(ctx).SecurityScanRequest(securityScanRequest).Execute()

Submit a scan

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
	securityScanRequest := *openapiclient.NewSecurityScanRequest([]openapiclient.SecurityFileInput{*openapiclient.NewSecurityFileInput("Path_example", "Content_example")}) // SecurityScanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityScansAPI.SecuritySecuritySubmitScan(context.Background()).SecurityScanRequest(securityScanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityScansAPI.SecuritySecuritySubmitScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySecuritySubmitScan`: SecurityScan
	fmt.Fprintf(os.Stdout, "Response from `SecurityScansAPI.SecuritySecuritySubmitScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecuritySubmitScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityScanRequest** | [**SecurityScanRequest**](SecurityScanRequest.md) |  | 

### Return type

[**SecurityScan**](SecurityScan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

