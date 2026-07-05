# \FrameworkDocumentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FrameworkCancelDocument**](FrameworkDocumentsAPI.md#FrameworkCancelDocument) | **Post** /v1/framework/{doctype}/{name}/cancel | Cancel a submitted document (docstatus 1→2)
[**FrameworkCreateDocument**](FrameworkDocumentsAPI.md#FrameworkCreateDocument) | **Post** /v1/framework/{doctype} | Create a document
[**FrameworkDeleteDocument**](FrameworkDocumentsAPI.md#FrameworkDeleteDocument) | **Delete** /v1/framework/{doctype}/{name} | Delete a document
[**FrameworkGetDocument**](FrameworkDocumentsAPI.md#FrameworkGetDocument) | **Get** /v1/framework/{doctype}/{name} | Get a document
[**FrameworkListDocuments**](FrameworkDocumentsAPI.md#FrameworkListDocuments) | **Get** /v1/framework/{doctype} | List documents of a DocType
[**FrameworkSubmitDocument**](FrameworkDocumentsAPI.md#FrameworkSubmitDocument) | **Post** /v1/framework/{doctype}/{name}/submit | Submit a document (docstatus 0→1)
[**FrameworkUpdateDocument**](FrameworkDocumentsAPI.md#FrameworkUpdateDocument) | **Put** /v1/framework/{doctype}/{name} | Update a draft document



## FrameworkCancelDocument

> FrameworkDocument FrameworkCancelDocument(ctx, doctype, name).Execute()

Cancel a submitted document (docstatus 1→2)

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocumentsAPI.FrameworkCancelDocument(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocumentsAPI.FrameworkCancelDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkCancelDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocumentsAPI.FrameworkCancelDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkCancelDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkCreateDocument

> FrameworkDocument FrameworkCreateDocument(ctx, doctype).FrameworkDocument(frameworkDocument).Execute()

Create a document

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
	doctype := "doctype_example" // string | 
	frameworkDocument := *openapiclient.NewFrameworkDocument() // FrameworkDocument | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocumentsAPI.FrameworkCreateDocument(context.Background(), doctype).FrameworkDocument(frameworkDocument).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocumentsAPI.FrameworkCreateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkCreateDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocumentsAPI.FrameworkCreateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frameworkDocument** | [**FrameworkDocument**](FrameworkDocument.md) |  | 

### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkDeleteDocument

> FrameworkDeleteDocument(ctx, doctype, name).Execute()

Delete a document

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkDocumentsAPI.FrameworkDeleteDocument(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocumentsAPI.FrameworkDeleteDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkDeleteDocumentRequest struct via the builder pattern


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


## FrameworkGetDocument

> FrameworkDocument FrameworkGetDocument(ctx, doctype, name).Execute()

Get a document

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocumentsAPI.FrameworkGetDocument(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocumentsAPI.FrameworkGetDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkGetDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocumentsAPI.FrameworkGetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkListDocuments

> FrameworkListDocuments200Response FrameworkListDocuments(ctx, doctype).Filters(filters).Fields(fields).OrderBy(orderBy).Limit(limit).Execute()

List documents of a DocType

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
	doctype := "doctype_example" // string | 
	filters := "filters_example" // string | 'JSON object of field→value filters (declared fields, name, or docstatus)' (optional)
	fields := "fields_example" // string | Comma list or JSON array of field names to project (optional)
	orderBy := "orderBy_example" // string | field [asc|desc] (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocumentsAPI.FrameworkListDocuments(context.Background(), doctype).Filters(filters).Fields(fields).OrderBy(orderBy).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocumentsAPI.FrameworkListDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkListDocuments`: FrameworkListDocuments200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocumentsAPI.FrameworkListDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | &#39;JSON object of field→value filters (declared fields, name, or docstatus)&#39; | 
 **fields** | **string** | Comma list or JSON array of field names to project | 
 **orderBy** | **string** | field [asc|desc] | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**FrameworkListDocuments200Response**](FrameworkListDocuments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkSubmitDocument

> FrameworkDocument FrameworkSubmitDocument(ctx, doctype, name).Execute()

Submit a document (docstatus 0→1)

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocumentsAPI.FrameworkSubmitDocument(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocumentsAPI.FrameworkSubmitDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkSubmitDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocumentsAPI.FrameworkSubmitDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkSubmitDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkUpdateDocument

> FrameworkDocument FrameworkUpdateDocument(ctx, doctype, name).FrameworkDocument(frameworkDocument).Execute()

Update a draft document

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 
	frameworkDocument := *openapiclient.NewFrameworkDocument() // FrameworkDocument | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocumentsAPI.FrameworkUpdateDocument(context.Background(), doctype, name).FrameworkDocument(frameworkDocument).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocumentsAPI.FrameworkUpdateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkUpdateDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocumentsAPI.FrameworkUpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkUpdateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **frameworkDocument** | [**FrameworkDocument**](FrameworkDocument.md) |  | 

### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

