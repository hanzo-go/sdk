# \CompaniesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmCreateCompany**](CompaniesAPI.md#CrmCreateCompany) | **Post** /v1/crm/companies | Create a company
[**CrmDeleteCompany**](CompaniesAPI.md#CrmDeleteCompany) | **Delete** /v1/crm/companies/{id} | Delete a company (clears dangling contact/opportunity refs)
[**CrmGetCompany**](CompaniesAPI.md#CrmGetCompany) | **Get** /v1/crm/companies/{id} | Company detail
[**CrmListCompanies**](CompaniesAPI.md#CrmListCompanies) | **Get** /v1/crm/companies | List companies
[**CrmUpdateCompany**](CompaniesAPI.md#CrmUpdateCompany) | **Put** /v1/crm/companies/{id} | Update a company



## CrmCreateCompany

> CrmCompany CrmCreateCompany(ctx).CrmCompanyInput(crmCompanyInput).Execute()

Create a company

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
	crmCompanyInput := *openapiclient.NewCrmCompanyInput("Name_example") // CrmCompanyInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CrmCreateCompany(context.Background()).CrmCompanyInput(crmCompanyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CrmCreateCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmCreateCompany`: CrmCompany
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CrmCreateCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCrmCreateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crmCompanyInput** | [**CrmCompanyInput**](CrmCompanyInput.md) |  | 

### Return type

[**CrmCompany**](CrmCompany.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmDeleteCompany

> CrmDeleteCompany(ctx, id).Execute()

Delete a company (clears dangling contact/opportunity refs)

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
	r, err := apiClient.CompaniesAPI.CrmDeleteCompany(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CrmDeleteCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrmDeleteCompanyRequest struct via the builder pattern


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


## CrmGetCompany

> CrmCompany CrmGetCompany(ctx, id).Execute()

Company detail

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
	resp, r, err := apiClient.CompaniesAPI.CrmGetCompany(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CrmGetCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmGetCompany`: CrmCompany
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CrmGetCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrmGetCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CrmCompany**](CrmCompany.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmListCompanies

> CrmListCompanies200Response CrmListCompanies(ctx).Limit(limit).Execute()

List companies

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
	limit := int32(56) // int32 |  (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CrmListCompanies(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CrmListCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmListCompanies`: CrmListCompanies200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CrmListCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCrmListCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 200]

### Return type

[**CrmListCompanies200Response**](CrmListCompanies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmUpdateCompany

> CrmCompany CrmUpdateCompany(ctx, id).CrmCompanyInput(crmCompanyInput).Execute()

Update a company

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
	crmCompanyInput := *openapiclient.NewCrmCompanyInput("Name_example") // CrmCompanyInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CrmUpdateCompany(context.Background(), id).CrmCompanyInput(crmCompanyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CrmUpdateCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmUpdateCompany`: CrmCompany
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CrmUpdateCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrmUpdateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crmCompanyInput** | [**CrmCompanyInput**](CrmCompanyInput.md) |  | 

### Return type

[**CrmCompany**](CrmCompany.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

