# \NexusTemplateAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddTemplate**](NexusTemplateAPIAPI.md#NexusAddTemplate) | **Post** /v1/nexus/add-template | add Template
[**NexusDeleteTemplate**](NexusTemplateAPIAPI.md#NexusDeleteTemplate) | **Post** /v1/nexus/delete-template | delete Template
[**NexusGetTemplate**](NexusTemplateAPIAPI.md#NexusGetTemplate) | **Get** /v1/nexus/get-template | get Template
[**NexusGetTemplates**](NexusTemplateAPIAPI.md#NexusGetTemplates) | **Get** /v1/nexus/get-templates | get Templates
[**NexusUpdateTemplate**](NexusTemplateAPIAPI.md#NexusUpdateTemplate) | **Post** /v1/nexus/update-template | update Template



## NexusAddTemplate

> NexusResponse NexusAddTemplate(ctx).NexusTemplate(nexusTemplate).Execute()

add Template



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
	nexusTemplate := *openapiclient.NewNexusTemplate() // NexusTemplate | The details of the template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTemplateAPIAPI.NexusAddTemplate(context.Background()).NexusTemplate(nexusTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTemplateAPIAPI.NexusAddTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddTemplate`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusTemplateAPIAPI.NexusAddTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusTemplate** | [**NexusTemplate**](NexusTemplate.md) | The details of the template | 

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


## NexusDeleteTemplate

> NexusResponse NexusDeleteTemplate(ctx).NexusTemplate(nexusTemplate).Execute()

delete Template



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
	nexusTemplate := *openapiclient.NewNexusTemplate() // NexusTemplate | The details of the template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTemplateAPIAPI.NexusDeleteTemplate(context.Background()).NexusTemplate(nexusTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTemplateAPIAPI.NexusDeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteTemplate`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusTemplateAPIAPI.NexusDeleteTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusTemplate** | [**NexusTemplate**](NexusTemplate.md) | The details of the template | 

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


## NexusGetTemplate

> NexusTemplate NexusGetTemplate(ctx).Id(id).Execute()

get Template



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
	id := "id_example" // string | The id of the template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTemplateAPIAPI.NexusGetTemplate(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTemplateAPIAPI.NexusGetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetTemplate`: NexusTemplate
	fmt.Fprintf(os.Stdout, "Response from `NexusTemplateAPIAPI.NexusGetTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the template | 

### Return type

[**NexusTemplate**](NexusTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetTemplates

> []NexusTemplate NexusGetTemplates(ctx).Owner(owner).Execute()

get Templates



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
	owner := "owner_example" // string | The owner of the templates

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTemplateAPIAPI.NexusGetTemplates(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTemplateAPIAPI.NexusGetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetTemplates`: []NexusTemplate
	fmt.Fprintf(os.Stdout, "Response from `NexusTemplateAPIAPI.NexusGetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the templates | 

### Return type

[**[]NexusTemplate**](NexusTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateTemplate

> NexusResponse NexusUpdateTemplate(ctx).Id(id).NexusTemplate(nexusTemplate).Execute()

update Template



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
	nexusTemplate := *openapiclient.NewNexusTemplate() // NexusTemplate | The details of the template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTemplateAPIAPI.NexusUpdateTemplate(context.Background()).Id(id).NexusTemplate(nexusTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTemplateAPIAPI.NexusUpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateTemplate`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusTemplateAPIAPI.NexusUpdateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the template | 
 **nexusTemplate** | [**NexusTemplate**](NexusTemplate.md) | The details of the template | 

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

