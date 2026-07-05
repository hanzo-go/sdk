# \CloudFormAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddForm**](CloudFormAPIAPI.md#CloudApiControllerAddForm) | **Post** /v1/cloud/add-form | Api Controller Add Form
[**CloudApiControllerDeleteForm**](CloudFormAPIAPI.md#CloudApiControllerDeleteForm) | **Post** /v1/cloud/delete-form | Api Controller Delete Form
[**CloudApiControllerGetForm**](CloudFormAPIAPI.md#CloudApiControllerGetForm) | **Get** /v1/cloud/get-form | Api Controller Get Form
[**CloudApiControllerGetFormData**](CloudFormAPIAPI.md#CloudApiControllerGetFormData) | **Get** /v1/cloud/get-form-data | Api Controller Get Form Data
[**CloudApiControllerGetForms**](CloudFormAPIAPI.md#CloudApiControllerGetForms) | **Get** /v1/cloud/get-forms | Api Controller Get Forms
[**CloudApiControllerGetGlobalForms**](CloudFormAPIAPI.md#CloudApiControllerGetGlobalForms) | **Get** /v1/cloud/get-global-forms | Api Controller Get Global Forms
[**CloudApiControllerUpdateForm**](CloudFormAPIAPI.md#CloudApiControllerUpdateForm) | **Post** /v1/cloud/update-form | Api Controller Update Form



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
	resp, r, err := apiClient.CloudFormAPIAPI.CloudApiControllerAddForm(context.Background()).CloudObjectForm(cloudObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudFormAPIAPI.CloudApiControllerAddForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddForm`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudFormAPIAPI.CloudApiControllerAddForm`: %v\n", resp)
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
	resp, r, err := apiClient.CloudFormAPIAPI.CloudApiControllerDeleteForm(context.Background()).CloudObjectForm(cloudObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudFormAPIAPI.CloudApiControllerDeleteForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteForm`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudFormAPIAPI.CloudApiControllerDeleteForm`: %v\n", resp)
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
	resp, r, err := apiClient.CloudFormAPIAPI.CloudApiControllerGetForm(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudFormAPIAPI.CloudApiControllerGetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetForm`: CloudObjectForm
	fmt.Fprintf(os.Stdout, "Response from `CloudFormAPIAPI.CloudApiControllerGetForm`: %v\n", resp)
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
	resp, r, err := apiClient.CloudFormAPIAPI.CloudApiControllerGetFormData(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudFormAPIAPI.CloudApiControllerGetFormData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetFormData`: []CloudObjectForm
	fmt.Fprintf(os.Stdout, "Response from `CloudFormAPIAPI.CloudApiControllerGetFormData`: %v\n", resp)
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
	resp, r, err := apiClient.CloudFormAPIAPI.CloudApiControllerGetForms(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudFormAPIAPI.CloudApiControllerGetForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetForms`: []CloudObjectForm
	fmt.Fprintf(os.Stdout, "Response from `CloudFormAPIAPI.CloudApiControllerGetForms`: %v\n", resp)
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
	resp, r, err := apiClient.CloudFormAPIAPI.CloudApiControllerGetGlobalForms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudFormAPIAPI.CloudApiControllerGetGlobalForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalForms`: []CloudObjectForm
	fmt.Fprintf(os.Stdout, "Response from `CloudFormAPIAPI.CloudApiControllerGetGlobalForms`: %v\n", resp)
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
	resp, r, err := apiClient.CloudFormAPIAPI.CloudApiControllerUpdateForm(context.Background()).Id(id).CloudObjectForm(cloudObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudFormAPIAPI.CloudApiControllerUpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateForm`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudFormAPIAPI.CloudApiControllerUpdateForm`: %v\n", resp)
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

