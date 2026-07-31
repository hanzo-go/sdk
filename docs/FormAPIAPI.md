# \FormAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddForm**](FormAPIAPI.md#CloudApiControllerAddForm) | **Post** /v1/cloud/add-form | Api Controller Add Form
[**CloudApiControllerDeleteForm**](FormAPIAPI.md#CloudApiControllerDeleteForm) | **Post** /v1/cloud/delete-form | Api Controller Delete Form
[**CloudApiControllerGetForm**](FormAPIAPI.md#CloudApiControllerGetForm) | **Get** /v1/cloud/get-form | Api Controller Get Form
[**CloudApiControllerGetFormData**](FormAPIAPI.md#CloudApiControllerGetFormData) | **Get** /v1/cloud/get-form-data | Api Controller Get Form Data
[**CloudApiControllerGetForms**](FormAPIAPI.md#CloudApiControllerGetForms) | **Get** /v1/cloud/get-forms | Api Controller Get Forms
[**CloudApiControllerGetGlobalForms**](FormAPIAPI.md#CloudApiControllerGetGlobalForms) | **Get** /v1/cloud/get-global-forms | Api Controller Get Global Forms
[**CloudApiControllerUpdateForm**](FormAPIAPI.md#CloudApiControllerUpdateForm) | **Post** /v1/cloud/update-form | Api Controller Update Form



## CloudApiControllerAddForm

> CloudControllersResponse CloudApiControllerAddForm(ctx).CloudObjectForm(cloudObjectForm).Execute()

Api Controller Add Form



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
	cloudObjectForm := *openapiclient.NewCloudObjectForm() // CloudObjectForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPIAPI.CloudApiControllerAddForm(context.Background()).CloudObjectForm(cloudObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPIAPI.CloudApiControllerAddForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddForm`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `FormAPIAPI.CloudApiControllerAddForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectForm** | [**CloudObjectForm**](CloudObjectForm.md) | The details of the form | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteForm

> CloudControllersResponse CloudApiControllerDeleteForm(ctx).CloudObjectForm(cloudObjectForm).Execute()

Api Controller Delete Form



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
	cloudObjectForm := *openapiclient.NewCloudObjectForm() // CloudObjectForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPIAPI.CloudApiControllerDeleteForm(context.Background()).CloudObjectForm(cloudObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPIAPI.CloudApiControllerDeleteForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteForm`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `FormAPIAPI.CloudApiControllerDeleteForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectForm** | [**CloudObjectForm**](CloudObjectForm.md) | The details of the form | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetForm

> CloudObjectForm CloudApiControllerGetForm(ctx).Id(id).Execute()

Api Controller Get Form



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
	id := "id_example" // string | The id (owner/name) of form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPIAPI.CloudApiControllerGetForm(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPIAPI.CloudApiControllerGetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetForm`: CloudObjectForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPIAPI.CloudApiControllerGetForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of form | 

### Return type

[**CloudObjectForm**](CloudObjectForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetFormData

> []CloudObjectForm CloudApiControllerGetFormData(ctx).Owner(owner).Execute()

Api Controller Get Form Data



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
	owner := "owner_example" // string | The owner of form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPIAPI.CloudApiControllerGetFormData(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPIAPI.CloudApiControllerGetFormData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetFormData`: []CloudObjectForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPIAPI.CloudApiControllerGetFormData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetFormDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of form | 

### Return type

[**[]CloudObjectForm**](CloudObjectForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetForms

> []CloudObjectForm CloudApiControllerGetForms(ctx).Owner(owner).Execute()

Api Controller Get Forms



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
	owner := "owner_example" // string | The owner of form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPIAPI.CloudApiControllerGetForms(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPIAPI.CloudApiControllerGetForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetForms`: []CloudObjectForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPIAPI.CloudApiControllerGetForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of form | 

### Return type

[**[]CloudObjectForm**](CloudObjectForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetGlobalForms

> []CloudObjectForm CloudApiControllerGetGlobalForms(ctx).Execute()

Api Controller Get Global Forms



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
	resp, r, err := apiClient.FormAPIAPI.CloudApiControllerGetGlobalForms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPIAPI.CloudApiControllerGetGlobalForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalForms`: []CloudObjectForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPIAPI.CloudApiControllerGetGlobalForms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalFormsRequest struct via the builder pattern


### Return type

[**[]CloudObjectForm**](CloudObjectForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateForm

> CloudControllersResponse CloudApiControllerUpdateForm(ctx).Id(id).CloudObjectForm(cloudObjectForm).Execute()

Api Controller Update Form



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
	id := "id_example" // string | The id (owner/name) of the form
	cloudObjectForm := *openapiclient.NewCloudObjectForm() // CloudObjectForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPIAPI.CloudApiControllerUpdateForm(context.Background()).Id(id).CloudObjectForm(cloudObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPIAPI.CloudApiControllerUpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateForm`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `FormAPIAPI.CloudApiControllerUpdateForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the form | 
 **cloudObjectForm** | [**CloudObjectForm**](CloudObjectForm.md) | The details of the form | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

