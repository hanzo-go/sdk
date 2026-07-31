# \FindingsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecuritySecurityGetFinding**](FindingsAPI.md#SecuritySecurityGetFinding) | **Get** /v1/security/findings/{id} | Get a finding
[**SecuritySecurityListFindings**](FindingsAPI.md#SecuritySecurityListFindings) | **Get** /v1/security/findings | List findings



## SecuritySecurityGetFinding

> SecurityFinding SecuritySecurityGetFinding(ctx, id).Execute()

Get a finding

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
	resp, r, err := apiClient.FindingsAPI.SecuritySecurityGetFinding(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.SecuritySecurityGetFinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySecurityGetFinding`: SecurityFinding
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.SecuritySecurityGetFinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecurityGetFindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityFinding**](SecurityFinding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySecurityListFindings

> SecuritySecurityListFindings200Response SecuritySecurityListFindings(ctx).ScanId(scanId).MinSeverity(minSeverity).Limit(limit).Execute()

List findings

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
	scanId := "scanId_example" // string |  (optional)
	minSeverity := "minSeverity_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.SecuritySecurityListFindings(context.Background()).ScanId(scanId).MinSeverity(minSeverity).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.SecuritySecurityListFindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySecurityListFindings`: SecuritySecurityListFindings200Response
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.SecuritySecurityListFindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecurityListFindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scanId** | **string** |  | 
 **minSeverity** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**SecuritySecurityListFindings200Response**](SecuritySecurityListFindings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

