# \LegalAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLegalDocuments**](LegalAPI.md#GetLegalDocuments) | **Get** /v1/legal/documents | Returns the org&#39;s generated documents, newest first, WITHOUT their rendered content — fetch one document to read its body.
[**GetLegalDocumentsById**](LegalAPI.md#GetLegalDocumentsById) | **Get** /v1/legal/documents/{id} | Returns one of the org&#39;s documents WITH its rendered body.
[**GetLegalFilings**](LegalAPI.md#GetLegalFilings) | **Get** /v1/legal/filings | Returns the org&#39;s filing records, newest first — which documents were filed where, through which provider, and what the filing&#39;s honest status is.
[**GetLegalHealth**](LegalAPI.md#GetLegalHealth) | **Get** /v1/legal/health | Reports that the legal subsystem is serving and how many built-in templates its catalog carries.
[**GetLegalTemplates**](LegalAPI.md#GetLegalTemplates) | **Get** /v1/legal/templates | Returns the org&#39;s effective template catalog: every built-in template, with any the org has overridden replaced by its own latest version.
[**GetLegalTemplatesById**](LegalAPI.md#GetLegalTemplatesById) | **Get** /v1/legal/templates/{id} | Returns one template resolved for the caller&#39;s org — the org&#39;s own override if it has saved one, else the built-in — with its full text/template body and its declared merge fields.
[**PostLegalDocuments**](LegalAPI.md#PostLegalDocuments) | **Post** /v1/legal/documents | Renders a document from a template and the caller&#39;s own merge data, seals it in the org&#39;s store, and returns it with its rendered body.
[**PostLegalDocumentsByIdSign**](LegalAPI.md#PostLegalDocumentsByIdSign) | **Post** /v1/legal/documents/{id}/sign | Opens an e-signature request over one document and moves it to out_for_signature, returning the provider&#39;s reference for the request.
[**PostLegalDocumentsByIdSignComplete**](LegalAPI.md#PostLegalDocumentsByIdSignComplete) | **Post** /v1/legal/documents/{id}/sign/complete | Record that a generated document&#39;s signature request completed
[**PostLegalFilings**](LegalAPI.md#PostLegalFilings) | **Post** /v1/legal/filings | Records a filing of one or more of the org&#39;s documents with a state or agency, and returns the tracking record.
[**PutLegalTemplatesById**](LegalAPI.md#PutLegalTemplatesById) | **Put** /v1/legal/templates/{id} | Saves the org&#39;s own version of a template — a custom NDA, a house MSA — and returns it with its new version number.



## GetLegalDocuments

> DocumentPage GetLegalDocuments(ctx).Limit(limit).Execute()

Returns the org's generated documents, newest first, WITHOUT their rendered content — fetch one document to read its body.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	limit := int32(56) // int32 | Limit bounds the page. Absent or unparseable means the store's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.GetLegalDocuments(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.GetLegalDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegalDocuments`: DocumentPage
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.GetLegalDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds the page. Absent or unparseable means the store&#39;s own default. | 

### Return type

[**DocumentPage**](DocumentPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegalDocumentsById

> DocumentReply GetLegalDocumentsById(ctx, id).Execute()

Returns one of the org's documents WITH its rendered body.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the document's server-minted handle, \"doc_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.GetLegalDocumentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.GetLegalDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegalDocumentsById`: DocumentReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.GetLegalDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document&#39;s server-minted handle, \&quot;doc_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumentReply**](DocumentReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegalFilings

> FilingPage GetLegalFilings(ctx).Limit(limit).Execute()

Returns the org's filing records, newest first — which documents were filed where, through which provider, and what the filing's honest status is.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	limit := int32(56) // int32 | Limit bounds the page. Absent or unparseable means the store's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.GetLegalFilings(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.GetLegalFilings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegalFilings`: FilingPage
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.GetLegalFilings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalFilingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds the page. Absent or unparseable means the store&#39;s own default. | 

### Return type

[**FilingPage**](FilingPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegalHealth

> LegalHealth GetLegalHealth(ctx).Execute()

Reports that the legal subsystem is serving and how many built-in templates its catalog carries.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.GetLegalHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.GetLegalHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegalHealth`: LegalHealth
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.GetLegalHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalHealthRequest struct via the builder pattern


### Return type

[**LegalHealth**](LegalHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegalTemplates

> TemplateCatalog GetLegalTemplates(ctx).Execute()

Returns the org's effective template catalog: every built-in template, with any the org has overridden replaced by its own latest version.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.GetLegalTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.GetLegalTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegalTemplates`: TemplateCatalog
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.GetLegalTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalTemplatesRequest struct via the builder pattern


### Return type

[**TemplateCatalog**](TemplateCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegalTemplatesById

> TemplateReply GetLegalTemplatesById(ctx, id).Execute()

Returns one template resolved for the caller's org — the org's own override if it has saved one, else the built-in — with its full text/template body and its declared merge fields.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the template's stable id, e.g. \"nda\" or \"safe\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.GetLegalTemplatesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.GetLegalTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLegalTemplatesById`: TemplateReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.GetLegalTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the template&#39;s stable id, e.g. \&quot;nda\&quot; or \&quot;safe\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateReply**](TemplateReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLegalDocuments

> DocumentReply PostLegalDocuments(ctx).GenerateRequest(generateRequest).Execute()

Renders a document from a template and the caller's own merge data, seals it in the org's store, and returns it with its rendered body.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	generateRequest := *openapiclient.NewGenerateRequest() // GenerateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.PostLegalDocuments(context.Background()).GenerateRequest(generateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.PostLegalDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLegalDocuments`: DocumentReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.PostLegalDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLegalDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateRequest** | [**GenerateRequest**](GenerateRequest.md) |  | 

### Return type

[**DocumentReply**](DocumentReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLegalDocumentsByIdSign

> SignReply PostLegalDocumentsByIdSign(ctx, id).SignRequest(signRequest).Execute()

Opens an e-signature request over one document and moves it to out_for_signature, returning the provider's reference for the request.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "doc_1f…" // string | ID is the document to send for signature, from the path.
	signRequest := *openapiclient.NewSignRequest() // SignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.PostLegalDocumentsByIdSign(context.Background(), id).SignRequest(signRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.PostLegalDocumentsByIdSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLegalDocumentsByIdSign`: SignReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.PostLegalDocumentsByIdSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document to send for signature, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLegalDocumentsByIdSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signRequest** | [**SignRequest**](SignRequest.md) |  | 

### Return type

[**SignReply**](SignReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLegalDocumentsByIdSignComplete

> PostLegalDocumentsByIdSignComplete(ctx, id).Execute()

Record that a generated document's signature request completed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LegalAPI.PostLegalDocumentsByIdSignComplete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.PostLegalDocumentsByIdSignComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLegalDocumentsByIdSignCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLegalFilings

> FilingReply PostLegalFilings(ctx).FilingRequest(filingRequest).Execute()

Records a filing of one or more of the org's documents with a state or agency, and returns the tracking record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	filingRequest := *openapiclient.NewFilingRequest() // FilingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.PostLegalFilings(context.Background()).FilingRequest(filingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.PostLegalFilings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLegalFilings`: FilingReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.PostLegalFilings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLegalFilingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filingRequest** | [**FilingRequest**](FilingRequest.md) |  | 

### Return type

[**FilingReply**](FilingReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLegalTemplatesById

> TemplateReply PutLegalTemplatesById(ctx, id).TemplateOverride(templateOverride).Execute()

Saves the org's own version of a template — a custom NDA, a house MSA — and returns it with its new version number.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "nda" // string | ID is the template to override, from the path. Overriding a built-in id inherits that built-in's category, title and counsel-review posture.
	templateOverride := *openapiclient.NewTemplateOverride() // TemplateOverride | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.PutLegalTemplatesById(context.Background(), id).TemplateOverride(templateOverride).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.PutLegalTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLegalTemplatesById`: TemplateReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.PutLegalTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the template to override, from the path. Overriding a built-in id inherits that built-in&#39;s category, title and counsel-review posture. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLegalTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateOverride** | [**TemplateOverride**](TemplateOverride.md) |  | 

### Return type

[**TemplateReply**](TemplateReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

