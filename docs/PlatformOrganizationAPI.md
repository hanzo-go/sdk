# \PlatformOrganizationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformOrganizationAll**](PlatformOrganizationAPI.md#PlatformOrganizationAll) | **Get** /v1/platform/organization/all | List organizations
[**PlatformOrganizationCreate**](PlatformOrganizationAPI.md#PlatformOrganizationCreate) | **Post** /v1/platform/organization/create | Create an organization



## PlatformOrganizationAll

> PlatformTRPCResult PlatformOrganizationAll(ctx).Execute()

List organizations

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
	resp, r, err := apiClient.PlatformOrganizationAPI.PlatformOrganizationAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformOrganizationAPI.PlatformOrganizationAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformOrganizationAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformOrganizationAPI.PlatformOrganizationAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformOrganizationAllRequest struct via the builder pattern


### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformOrganizationCreate

> PlatformTRPCResult PlatformOrganizationCreate(ctx).PlatformOrganizationCreateRequest(platformOrganizationCreateRequest).Execute()

Create an organization

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
	platformOrganizationCreateRequest := *openapiclient.NewPlatformOrganizationCreateRequest() // PlatformOrganizationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformOrganizationAPI.PlatformOrganizationCreate(context.Background()).PlatformOrganizationCreateRequest(platformOrganizationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformOrganizationAPI.PlatformOrganizationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformOrganizationCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformOrganizationAPI.PlatformOrganizationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformOrganizationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformOrganizationCreateRequest** | [**PlatformOrganizationCreateRequest**](PlatformOrganizationCreateRequest.md) |  | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

