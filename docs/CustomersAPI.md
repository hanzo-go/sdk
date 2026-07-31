# \CustomersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminCustomerDetail**](CustomersAPI.md#AdminAdminCustomerDetail) | **Get** /v1/admin/customers/{org} | One customer&#39;s detail
[**AdminAdminGrantCredit**](CustomersAPI.md#AdminAdminGrantCredit) | **Post** /v1/admin/customers/{org}/credit | Grant credit (commerce deposit)
[**AdminAdminListCustomers**](CustomersAPI.md#AdminAdminListCustomers) | **Get** /v1/admin/customers | Fleet customer list
[**AdminAdminReactivateCustomer**](CustomersAPI.md#AdminAdminReactivateCustomer) | **Post** /v1/admin/customers/{org}/reactivate | Reactivate a customer (IAM isForbidden&#x3D;false)
[**AdminAdminSuspendCustomer**](CustomersAPI.md#AdminAdminSuspendCustomer) | **Post** /v1/admin/customers/{org}/suspend | Suspend a customer (IAM isForbidden&#x3D;true)



## AdminAdminCustomerDetail

> AdminAdminCustomerDetail200Response AdminAdminCustomerDetail(ctx, org).Execute()

One customer's detail

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
	org := "org_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.AdminAdminCustomerDetail(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.AdminAdminCustomerDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminCustomerDetail`: AdminAdminCustomerDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.AdminAdminCustomerDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminCustomerDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminAdminCustomerDetail200Response**](AdminAdminCustomerDetail200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminGrantCredit

> AdminAdminGrantCredit200Response AdminAdminGrantCredit(ctx, org).AdminCreditRequest(adminCreditRequest).Execute()

Grant credit (commerce deposit)

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
	org := "org_example" // string | 
	adminCreditRequest := *openapiclient.NewAdminCreditRequest(int64(123)) // AdminCreditRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.AdminAdminGrantCredit(context.Background(), org).AdminCreditRequest(adminCreditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.AdminAdminGrantCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminGrantCredit`: AdminAdminGrantCredit200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.AdminAdminGrantCredit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminGrantCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminCreditRequest** | [**AdminCreditRequest**](AdminCreditRequest.md) |  | 

### Return type

[**AdminAdminGrantCredit200Response**](AdminAdminGrantCredit200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminListCustomers

> AdminAdminListCustomers200Response AdminAdminListCustomers(ctx).Execute()

Fleet customer list

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
	resp, r, err := apiClient.CustomersAPI.AdminAdminListCustomers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.AdminAdminListCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminListCustomers`: AdminAdminListCustomers200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.AdminAdminListCustomers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminListCustomersRequest struct via the builder pattern


### Return type

[**AdminAdminListCustomers200Response**](AdminAdminListCustomers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminReactivateCustomer

> AdminAdminReactivateCustomer200Response AdminAdminReactivateCustomer(ctx, org).Execute()

Reactivate a customer (IAM isForbidden=false)

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
	org := "org_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.AdminAdminReactivateCustomer(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.AdminAdminReactivateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminReactivateCustomer`: AdminAdminReactivateCustomer200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.AdminAdminReactivateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminReactivateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminAdminReactivateCustomer200Response**](AdminAdminReactivateCustomer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminSuspendCustomer

> AdminAdminReactivateCustomer200Response AdminAdminSuspendCustomer(ctx, org).Execute()

Suspend a customer (IAM isForbidden=true)

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
	org := "org_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.AdminAdminSuspendCustomer(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.AdminAdminSuspendCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminSuspendCustomer`: AdminAdminReactivateCustomer200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.AdminAdminSuspendCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminSuspendCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminAdminReactivateCustomer200Response**](AdminAdminReactivateCustomer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

