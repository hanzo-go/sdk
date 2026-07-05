# \CommerceFormsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCreateForm**](CommerceFormsAPI.md#CommerceCreateForm) | **Post** /v1/commerce/form | Create form
[**CommerceGetForm**](CommerceFormsAPI.md#CommerceGetForm) | **Get** /v1/commerce/form/{formid} | Get form
[**CommerceListForms**](CommerceFormsAPI.md#CommerceListForms) | **Get** /v1/commerce/form | List forms
[**CommerceSubmitForm**](CommerceFormsAPI.md#CommerceSubmitForm) | **Post** /v1/commerce/form/{formid}/submit | Submit form
[**CommerceSubscribeForm**](CommerceFormsAPI.md#CommerceSubscribeForm) | **Post** /v1/commerce/form/{formid}/subscribe | Subscribe via form



## CommerceCreateForm

> CommerceForm CommerceCreateForm(ctx).CommerceForm(commerceForm).Execute()

Create form

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
	commerceForm := *openapiclient.NewCommerceForm() // CommerceForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceFormsAPI.CommerceCreateForm(context.Background()).CommerceForm(commerceForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceFormsAPI.CommerceCreateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateForm`: CommerceForm
	fmt.Fprintf(os.Stdout, "Response from `CommerceFormsAPI.CommerceCreateForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceForm** | [**CommerceForm**](CommerceForm.md) |  | 

### Return type

[**CommerceForm**](CommerceForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetForm

> CommerceForm CommerceGetForm(ctx, formid).Execute()

Get form

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
	formid := "formid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceFormsAPI.CommerceGetForm(context.Background(), formid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceFormsAPI.CommerceGetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetForm`: CommerceForm
	fmt.Fprintf(os.Stdout, "Response from `CommerceFormsAPI.CommerceGetForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceForm**](CommerceForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListForms

> CommercePaginatedForms CommerceListForms(ctx).Execute()

List forms

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
	resp, r, err := apiClient.CommerceFormsAPI.CommerceListForms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceFormsAPI.CommerceListForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListForms`: CommercePaginatedForms
	fmt.Fprintf(os.Stdout, "Response from `CommerceFormsAPI.CommerceListForms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListFormsRequest struct via the builder pattern


### Return type

[**CommercePaginatedForms**](CommercePaginatedForms.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSubmitForm

> map[string]interface{} CommerceSubmitForm(ctx, formid).RequestBody(requestBody).Execute()

Submit form

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
	formid := "formid_example" // string | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceFormsAPI.CommerceSubmitForm(context.Background(), formid).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceFormsAPI.CommerceSubmitForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSubmitForm`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceFormsAPI.CommerceSubmitForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSubmitFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSubscribeForm

> map[string]interface{} CommerceSubscribeForm(ctx, formid).CommerceSubscribeFormRequest(commerceSubscribeFormRequest).Execute()

Subscribe via form

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
	formid := "formid_example" // string | 
	commerceSubscribeFormRequest := *openapiclient.NewCommerceSubscribeFormRequest() // CommerceSubscribeFormRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceFormsAPI.CommerceSubscribeForm(context.Background(), formid).CommerceSubscribeFormRequest(commerceSubscribeFormRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceFormsAPI.CommerceSubscribeForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSubscribeForm`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceFormsAPI.CommerceSubscribeForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSubscribeFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceSubscribeFormRequest** | [**CommerceSubscribeFormRequest**](CommerceSubscribeFormRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

