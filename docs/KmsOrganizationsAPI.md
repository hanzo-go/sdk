# \KmsOrganizationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsGetOrganization**](KmsOrganizationsAPI.md#KmsGetOrganization) | **Get** /v1/kms/organization/{organizationId} | Get an organization by ID
[**KmsListOrganizations**](KmsOrganizationsAPI.md#KmsListOrganizations) | **Get** /v1/kms/organization | List organizations the user belongs to
[**KmsUpdateOrganization**](KmsOrganizationsAPI.md#KmsUpdateOrganization) | **Patch** /v1/kms/organization/{organizationId} | Update an organization



## KmsGetOrganization

> KmsGetOrganization200Response KmsGetOrganization(ctx, organizationId).Execute()

Get an organization by ID

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsOrganizationsAPI.KmsGetOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsOrganizationsAPI.KmsGetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetOrganization`: KmsGetOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsOrganizationsAPI.KmsGetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsGetOrganization200Response**](KmsGetOrganization200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListOrganizations

> KmsListOrganizations200Response KmsListOrganizations(ctx).Execute()

List organizations the user belongs to

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
	resp, r, err := apiClient.KmsOrganizationsAPI.KmsListOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsOrganizationsAPI.KmsListOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListOrganizations`: KmsListOrganizations200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsOrganizationsAPI.KmsListOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsListOrganizationsRequest struct via the builder pattern


### Return type

[**KmsListOrganizations200Response**](KmsListOrganizations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateOrganization

> KmsGetOrganization200Response KmsUpdateOrganization(ctx, organizationId).KmsUpdateOrganizationRequest(kmsUpdateOrganizationRequest).Execute()

Update an organization

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateOrganizationRequest := *openapiclient.NewKmsUpdateOrganizationRequest() // KmsUpdateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsOrganizationsAPI.KmsUpdateOrganization(context.Background(), organizationId).KmsUpdateOrganizationRequest(kmsUpdateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsOrganizationsAPI.KmsUpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateOrganization`: KmsGetOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsOrganizationsAPI.KmsUpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateOrganizationRequest** | [**KmsUpdateOrganizationRequest**](KmsUpdateOrganizationRequest.md) |  | 

### Return type

[**KmsGetOrganization200Response**](KmsGetOrganization200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

