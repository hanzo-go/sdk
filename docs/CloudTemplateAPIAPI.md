# \CloudTemplateAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddTemplate**](CloudTemplateAPIAPI.md#CloudApiControllerAddTemplate) | **Post** /v1/cloud/add-template | Api Controller Add Template
[**CloudApiControllerDeleteTemplate**](CloudTemplateAPIAPI.md#CloudApiControllerDeleteTemplate) | **Post** /v1/cloud/delete-template | Api Controller Delete Template
[**CloudApiControllerGetTemplate**](CloudTemplateAPIAPI.md#CloudApiControllerGetTemplate) | **Get** /v1/cloud/get-template | Api Controller Get Template
[**CloudApiControllerGetTemplates**](CloudTemplateAPIAPI.md#CloudApiControllerGetTemplates) | **Get** /v1/cloud/get-templates | Api Controller Get Templates
[**CloudApiControllerUpdateTemplate**](CloudTemplateAPIAPI.md#CloudApiControllerUpdateTemplate) | **Post** /v1/cloud/update-template | Api Controller Update Template



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
	resp, r, err := apiClient.CloudTemplateAPIAPI.CloudApiControllerAddTemplate(context.Background()).CloudObjectTemplate(cloudObjectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTemplateAPIAPI.CloudApiControllerAddTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddTemplate`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudTemplateAPIAPI.CloudApiControllerAddTemplate`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTemplateAPIAPI.CloudApiControllerDeleteTemplate(context.Background()).CloudObjectTemplate(cloudObjectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTemplateAPIAPI.CloudApiControllerDeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteTemplate`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudTemplateAPIAPI.CloudApiControllerDeleteTemplate`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTemplateAPIAPI.CloudApiControllerGetTemplate(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTemplateAPIAPI.CloudApiControllerGetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetTemplate`: CloudObjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `CloudTemplateAPIAPI.CloudApiControllerGetTemplate`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTemplateAPIAPI.CloudApiControllerGetTemplates(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTemplateAPIAPI.CloudApiControllerGetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetTemplates`: []CloudObjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `CloudTemplateAPIAPI.CloudApiControllerGetTemplates`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTemplateAPIAPI.CloudApiControllerUpdateTemplate(context.Background()).Id(id).CloudObjectTemplate(cloudObjectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTemplateAPIAPI.CloudApiControllerUpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateTemplate`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudTemplateAPIAPI.CloudApiControllerUpdateTemplate`: %v\n", resp)
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

