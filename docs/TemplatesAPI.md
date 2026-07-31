# \TemplatesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiAddTemplate**](TemplatesAPI.md#AiAddTemplate) | **Post** /v1/ai/templates | Create a template
[**AiDeleteTemplate**](TemplatesAPI.md#AiDeleteTemplate) | **Delete** /v1/ai/templates/{owner}/{name} | Delete a template
[**AiGetTemplate**](TemplatesAPI.md#AiGetTemplate) | **Get** /v1/ai/templates/{owner}/{name} | Retrieve a template
[**AiGetTemplates**](TemplatesAPI.md#AiGetTemplates) | **Get** /v1/ai/templates | List templates
[**AiReplaceTemplate**](TemplatesAPI.md#AiReplaceTemplate) | **Put** /v1/ai/templates/{owner}/{name} | Replace a template
[**AiUpdateTemplate**](TemplatesAPI.md#AiUpdateTemplate) | **Patch** /v1/ai/templates/{owner}/{name} | Update a template
[**CloudDeleteV1TemplatesSlug**](TemplatesAPI.md#CloudDeleteV1TemplatesSlug) | **Delete** /v1/templates/{slug} | Deletes the caller org&#39;s OWN starter kit.
[**CloudGetV1Templates**](TemplatesAPI.md#CloudGetV1Templates) | **Get** /v1/templates | Lists the public starter-kit catalog plus, for a validated caller, that org&#39;s own private kits.
[**CloudGetV1TemplatesSlug**](TemplatesAPI.md#CloudGetV1TemplatesSlug) | **Get** /v1/templates/{slug} | Returns one starter kit: the caller org&#39;s own by that slug, else the public catalog&#39;s.
[**CloudPostV1Templates**](TemplatesAPI.md#CloudPostV1Templates) | **Post** /v1/templates | Creates a starter kit PRIVATE to the caller&#39;s org and answers 201 with the stored kit.
[**CloudPutV1TemplatesSlug**](TemplatesAPI.md#CloudPutV1TemplatesSlug) | **Put** /v1/templates/{slug} | Overwrites the caller org&#39;s OWN starter kit at the path slug, answering the stored kit.



## AiAddTemplate

> AiEnvelope AiAddTemplate(ctx).Body(body).Execute()

Create a template



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.AiAddTemplate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.AiAddTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAddTemplate`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.AiAddTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAddTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiDeleteTemplate

> AiEnvelope AiDeleteTemplate(ctx, owner, name).Execute()

Delete a template



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
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.AiDeleteTemplate(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.AiDeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiDeleteTemplate`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.AiDeleteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiGetTemplate

> AiEnvelope AiGetTemplate(ctx, owner, name).Execute()

Retrieve a template



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
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.AiGetTemplate(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.AiGetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetTemplate`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.AiGetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiGetTemplates

> AiEnvelope AiGetTemplates(ctx).Execute()

List templates



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
	resp, r, err := apiClient.TemplatesAPI.AiGetTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.AiGetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetTemplates`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.AiGetTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetTemplatesRequest struct via the builder pattern


### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiReplaceTemplate

> AiEnvelope AiReplaceTemplate(ctx, owner, name).Body(body).Execute()

Replace a template



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
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.AiReplaceTemplate(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.AiReplaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiReplaceTemplate`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.AiReplaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiReplaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiUpdateTemplate

> AiEnvelope AiUpdateTemplate(ctx, owner, name).Body(body).Execute()

Update a template



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
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.AiUpdateTemplate(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.AiUpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiUpdateTemplate`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.AiUpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1TemplatesSlug

> CloudDeleteV1TemplatesSlug(ctx, slug).Execute()

Deletes the caller org's OWN starter kit.



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
	slug := "acme-portal" // string | Slug is the starter kit to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplatesAPI.CloudDeleteV1TemplatesSlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CloudDeleteV1TemplatesSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the starter kit to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1TemplatesSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Templates

> CloudKitList CloudGetV1Templates(ctx).Execute()

Lists the public starter-kit catalog plus, for a validated caller, that org's own private kits.



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
	resp, r, err := apiClient.TemplatesAPI.CloudGetV1Templates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CloudGetV1Templates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Templates`: CloudKitList
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CloudGetV1Templates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TemplatesRequest struct via the builder pattern


### Return type

[**CloudKitList**](CloudKitList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TemplatesSlug

> CloudStarterKit CloudGetV1TemplatesSlug(ctx, slug).Execute()

Returns one starter kit: the caller org's own by that slug, else the public catalog's.



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
	slug := "folio" // string | Slug is the starter kit to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CloudGetV1TemplatesSlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CloudGetV1TemplatesSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TemplatesSlug`: CloudStarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CloudGetV1TemplatesSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the starter kit to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TemplatesSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudStarterKit**](CloudStarterKit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Templates

> CloudStarterKit CloudPostV1Templates(ctx).CloudPublishKitIn(cloudPublishKitIn).Execute()

Creates a starter kit PRIVATE to the caller's org and answers 201 with the stored kit.



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
	cloudPublishKitIn := *openapiclient.NewCloudPublishKitIn() // CloudPublishKitIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CloudPostV1Templates(context.Background()).CloudPublishKitIn(cloudPublishKitIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CloudPostV1Templates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Templates`: CloudStarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CloudPostV1Templates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPublishKitIn** | [**CloudPublishKitIn**](CloudPublishKitIn.md) |  | 

### Return type

[**CloudStarterKit**](CloudStarterKit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1TemplatesSlug

> CloudStarterKit CloudPutV1TemplatesSlug(ctx, slug).CloudReplaceKitIn(cloudReplaceKitIn).Execute()

Overwrites the caller org's OWN starter kit at the path slug, answering the stored kit.



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
	slug := "acme-portal" // string | Slug is the kit to replace, from the path.
	cloudReplaceKitIn := *openapiclient.NewCloudReplaceKitIn() // CloudReplaceKitIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CloudPutV1TemplatesSlug(context.Background(), slug).CloudReplaceKitIn(cloudReplaceKitIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CloudPutV1TemplatesSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1TemplatesSlug`: CloudStarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CloudPutV1TemplatesSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the kit to replace, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1TemplatesSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudReplaceKitIn** | [**CloudReplaceKitIn**](CloudReplaceKitIn.md) |  | 

### Return type

[**CloudStarterKit**](CloudStarterKit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

