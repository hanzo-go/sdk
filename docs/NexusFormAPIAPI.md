# \NexusFormAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddForm**](NexusFormAPIAPI.md#NexusAddForm) | **Post** /v1/nexus/add-form | add Form
[**NexusDeleteForm**](NexusFormAPIAPI.md#NexusDeleteForm) | **Post** /v1/nexus/delete-form | delete Form
[**NexusGetForm**](NexusFormAPIAPI.md#NexusGetForm) | **Get** /v1/nexus/get-form | get Form
[**NexusGetFormData**](NexusFormAPIAPI.md#NexusGetFormData) | **Get** /v1/nexus/get-form-data | get Form Data
[**NexusGetForms**](NexusFormAPIAPI.md#NexusGetForms) | **Get** /v1/nexus/get-forms | get Forms
[**NexusGetGlobalForms**](NexusFormAPIAPI.md#NexusGetGlobalForms) | **Get** /v1/nexus/get-global-forms | get Global Forms
[**NexusUpdateForm**](NexusFormAPIAPI.md#NexusUpdateForm) | **Post** /v1/nexus/update-form | update Form



## NexusAddForm

> NexusResponse NexusAddForm(ctx).NexusForm(nexusForm).Execute()

add Form



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
	nexusForm := *openapiclient.NewNexusForm() // NexusForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFormAPIAPI.NexusAddForm(context.Background()).NexusForm(nexusForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFormAPIAPI.NexusAddForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddForm`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusFormAPIAPI.NexusAddForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusForm** | [**NexusForm**](NexusForm.md) | The details of the form | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteForm

> NexusResponse NexusDeleteForm(ctx).NexusForm(nexusForm).Execute()

delete Form



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
	nexusForm := *openapiclient.NewNexusForm() // NexusForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFormAPIAPI.NexusDeleteForm(context.Background()).NexusForm(nexusForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFormAPIAPI.NexusDeleteForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteForm`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusFormAPIAPI.NexusDeleteForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusForm** | [**NexusForm**](NexusForm.md) | The details of the form | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetForm

> NexusForm NexusGetForm(ctx).Id(id).Execute()

get Form



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFormAPIAPI.NexusGetForm(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFormAPIAPI.NexusGetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetForm`: NexusForm
	fmt.Fprintf(os.Stdout, "Response from `NexusFormAPIAPI.NexusGetForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the form | 

### Return type

[**NexusForm**](NexusForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetFormData

> []NexusForm NexusGetFormData(ctx).Owner(owner).Execute()

get Form Data



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
	owner := "owner_example" // string | The owner of the forms

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFormAPIAPI.NexusGetFormData(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFormAPIAPI.NexusGetFormData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetFormData`: []NexusForm
	fmt.Fprintf(os.Stdout, "Response from `NexusFormAPIAPI.NexusGetFormData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetFormDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the forms | 

### Return type

[**[]NexusForm**](NexusForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetForms

> []NexusForm NexusGetForms(ctx).Owner(owner).Execute()

get Forms



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
	owner := "owner_example" // string | The owner of the forms

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFormAPIAPI.NexusGetForms(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFormAPIAPI.NexusGetForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetForms`: []NexusForm
	fmt.Fprintf(os.Stdout, "Response from `NexusFormAPIAPI.NexusGetForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the forms | 

### Return type

[**[]NexusForm**](NexusForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalForms

> []NexusForm NexusGetGlobalForms(ctx).Execute()

get Global Forms



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
	resp, r, err := apiClient.NexusFormAPIAPI.NexusGetGlobalForms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFormAPIAPI.NexusGetGlobalForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalForms`: []NexusForm
	fmt.Fprintf(os.Stdout, "Response from `NexusFormAPIAPI.NexusGetGlobalForms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalFormsRequest struct via the builder pattern


### Return type

[**[]NexusForm**](NexusForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateForm

> NexusResponse NexusUpdateForm(ctx).Id(id).NexusForm(nexusForm).Execute()

update Form



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
	nexusForm := *openapiclient.NewNexusForm() // NexusForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFormAPIAPI.NexusUpdateForm(context.Background()).Id(id).NexusForm(nexusForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFormAPIAPI.NexusUpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateForm`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusFormAPIAPI.NexusUpdateForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the form | 
 **nexusForm** | [**NexusForm**](NexusForm.md) | The details of the form | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

