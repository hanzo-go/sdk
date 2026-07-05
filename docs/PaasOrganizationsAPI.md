# \PaasOrganizationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasCreateOrganization**](PaasOrganizationsAPI.md#PaasCreateOrganization) | **Post** /v1/paas/org | Create organization
[**PaasDeleteOrganization**](PaasOrganizationsAPI.md#PaasDeleteOrganization) | **Delete** /v1/paas/org/{orgId} | Delete organization
[**PaasGetOrganization**](PaasOrganizationsAPI.md#PaasGetOrganization) | **Get** /v1/paas/org/{orgId} | Get organization
[**PaasListOrganizations**](PaasOrganizationsAPI.md#PaasListOrganizations) | **Get** /v1/paas/org | List organizations
[**PaasUpdateOrganization**](PaasOrganizationsAPI.md#PaasUpdateOrganization) | **Put** /v1/paas/org/{orgId} | Update organization



## PaasCreateOrganization

> PaasOrganization PaasCreateOrganization(ctx).PaasCreateOrganizationRequest(paasCreateOrganizationRequest).Execute()

Create organization

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
	paasCreateOrganizationRequest := *openapiclient.NewPaasCreateOrganizationRequest("Name_example") // PaasCreateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasOrganizationsAPI.PaasCreateOrganization(context.Background()).PaasCreateOrganizationRequest(paasCreateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasOrganizationsAPI.PaasCreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasCreateOrganization`: PaasOrganization
	fmt.Fprintf(os.Stdout, "Response from `PaasOrganizationsAPI.PaasCreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaasCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paasCreateOrganizationRequest** | [**PaasCreateOrganizationRequest**](PaasCreateOrganizationRequest.md) |  | 

### Return type

[**PaasOrganization**](PaasOrganization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasDeleteOrganization

> map[string]interface{} PaasDeleteOrganization(ctx, orgId).Execute()

Delete organization

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
	resp, r, err := apiClient.PaasOrganizationsAPI.PaasDeleteOrganization(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasOrganizationsAPI.PaasDeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasDeleteOrganization`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasOrganizationsAPI.PaasDeleteOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasDeleteOrganizationRequest struct via the builder pattern


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


## PaasGetOrganization

> PaasOrganization PaasGetOrganization(ctx, orgId).Execute()

Get organization

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
	resp, r, err := apiClient.PaasOrganizationsAPI.PaasGetOrganization(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasOrganizationsAPI.PaasGetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetOrganization`: PaasOrganization
	fmt.Fprintf(os.Stdout, "Response from `PaasOrganizationsAPI.PaasGetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaasOrganization**](PaasOrganization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasListOrganizations

> []PaasOrganization PaasListOrganizations(ctx).Execute()

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
	resp, r, err := apiClient.PaasOrganizationsAPI.PaasListOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasOrganizationsAPI.PaasListOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasListOrganizations`: []PaasOrganization
	fmt.Fprintf(os.Stdout, "Response from `PaasOrganizationsAPI.PaasListOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaasListOrganizationsRequest struct via the builder pattern


### Return type

[**[]PaasOrganization**](PaasOrganization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasUpdateOrganization

> map[string]interface{} PaasUpdateOrganization(ctx, orgId).PaasUpdateOrganizationRequest(paasUpdateOrganizationRequest).Execute()

Update organization

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
	paasUpdateOrganizationRequest := *openapiclient.NewPaasUpdateOrganizationRequest() // PaasUpdateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasOrganizationsAPI.PaasUpdateOrganization(context.Background(), orgId).PaasUpdateOrganizationRequest(paasUpdateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasOrganizationsAPI.PaasUpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasUpdateOrganization`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasOrganizationsAPI.PaasUpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paasUpdateOrganizationRequest** | [**PaasUpdateOrganizationRequest**](PaasUpdateOrganizationRequest.md) |  | 

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

