# \TemplateAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddTemplate**](TemplateAPIAPI.md#CloudApiControllerAddTemplate) | **Post** /v1/cloud/add-template | Api Controller Add Template
[**CloudApiControllerDeleteTemplate**](TemplateAPIAPI.md#CloudApiControllerDeleteTemplate) | **Post** /v1/cloud/delete-template | Api Controller Delete Template
[**CloudApiControllerGetTemplate**](TemplateAPIAPI.md#CloudApiControllerGetTemplate) | **Get** /v1/cloud/get-template | Api Controller Get Template
[**CloudApiControllerGetTemplates**](TemplateAPIAPI.md#CloudApiControllerGetTemplates) | **Get** /v1/cloud/get-templates | Api Controller Get Templates
[**CloudApiControllerUpdateTemplate**](TemplateAPIAPI.md#CloudApiControllerUpdateTemplate) | **Post** /v1/cloud/update-template | Api Controller Update Template



## CloudApiControllerAddTemplate

> CloudControllersResponse CloudApiControllerAddTemplate(ctx).CloudObjectTemplate(cloudObjectTemplate).Execute()

Api Controller Add Template



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
	cloudObjectTemplate := *openapiclient.NewCloudObjectTemplate() // CloudObjectTemplate | The details of the template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPIAPI.CloudApiControllerAddTemplate(context.Background()).CloudObjectTemplate(cloudObjectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPIAPI.CloudApiControllerAddTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddTemplate`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPIAPI.CloudApiControllerAddTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectTemplate** | [**CloudObjectTemplate**](CloudObjectTemplate.md) | The details of the template | 

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


## CloudApiControllerDeleteTemplate

> CloudControllersResponse CloudApiControllerDeleteTemplate(ctx).CloudObjectTemplate(cloudObjectTemplate).Execute()

Api Controller Delete Template



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
	cloudObjectTemplate := *openapiclient.NewCloudObjectTemplate() // CloudObjectTemplate | The details of the template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPIAPI.CloudApiControllerDeleteTemplate(context.Background()).CloudObjectTemplate(cloudObjectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPIAPI.CloudApiControllerDeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteTemplate`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPIAPI.CloudApiControllerDeleteTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectTemplate** | [**CloudObjectTemplate**](CloudObjectTemplate.md) | The details of the template | 

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


## CloudApiControllerGetTemplate

> CloudObjectTemplate CloudApiControllerGetTemplate(ctx).Id(id).Execute()

Api Controller Get Template



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
	id := "id_example" // string | The id of template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPIAPI.CloudApiControllerGetTemplate(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPIAPI.CloudApiControllerGetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetTemplate`: CloudObjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPIAPI.CloudApiControllerGetTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of template | 

### Return type

[**CloudObjectTemplate**](CloudObjectTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetTemplates

> []CloudObjectTemplate CloudApiControllerGetTemplates(ctx).Owner(owner).Execute()

Api Controller Get Templates



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
	owner := "owner_example" // string | The owner of templates

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPIAPI.CloudApiControllerGetTemplates(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPIAPI.CloudApiControllerGetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetTemplates`: []CloudObjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPIAPI.CloudApiControllerGetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of templates | 

### Return type

[**[]CloudObjectTemplate**](CloudObjectTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateTemplate

> CloudControllersResponse CloudApiControllerUpdateTemplate(ctx).Id(id).CloudObjectTemplate(cloudObjectTemplate).Execute()

Api Controller Update Template



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
	id := "id_example" // string | The id (owner/name) of the template
	cloudObjectTemplate := *openapiclient.NewCloudObjectTemplate() // CloudObjectTemplate | The details of the template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPIAPI.CloudApiControllerUpdateTemplate(context.Background()).Id(id).CloudObjectTemplate(cloudObjectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPIAPI.CloudApiControllerUpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateTemplate`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPIAPI.CloudApiControllerUpdateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the template | 
 **cloudObjectTemplate** | [**CloudObjectTemplate**](CloudObjectTemplate.md) | The details of the template | 

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

