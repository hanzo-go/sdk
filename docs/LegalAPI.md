# \LegalAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1LegalDocuments**](LegalAPI.md#CloudGetV1LegalDocuments) | **Get** /v1/legal/documents | ListLegalDocuments returns the org&#39;s generated documents, newest first, WITHOUT their rendered content — fetch one document to read its body.
[**CloudGetV1LegalDocumentsId**](LegalAPI.md#CloudGetV1LegalDocumentsId) | **Get** /v1/legal/documents/{id} | GetLegalDocument returns one of the org&#39;s documents WITH its rendered body.
[**CloudGetV1LegalFilings**](LegalAPI.md#CloudGetV1LegalFilings) | **Get** /v1/legal/filings | ListLegalFilings returns the org&#39;s filing records, newest first — which documents were filed where, through which provider, and what the filing&#39;s honest status is.
[**CloudGetV1LegalHealth**](LegalAPI.md#CloudGetV1LegalHealth) | **Get** /v1/legal/health | LegalHealth reports that the legal subsystem is serving and how many built-in templates its catalog carries.
[**CloudGetV1LegalTemplates**](LegalAPI.md#CloudGetV1LegalTemplates) | **Get** /v1/legal/templates | ListLegalTemplates returns the org&#39;s effective template catalog: every built-in template, with any the org has overridden replaced by its own latest version.
[**CloudGetV1LegalTemplatesId**](LegalAPI.md#CloudGetV1LegalTemplatesId) | **Get** /v1/legal/templates/{id} | GetLegalTemplate returns one template resolved for the caller&#39;s org — the org&#39;s own override if it has saved one, else the built-in — with its full text/template body and its declared merge fields.
[**CloudPostV1LegalDocuments**](LegalAPI.md#CloudPostV1LegalDocuments) | **Post** /v1/legal/documents | GenerateLegalDocument renders a document from a template and the caller&#39;s own merge data, seals it in the org&#39;s store, and returns it with its rendered body.
[**CloudPostV1LegalDocumentsByIdSignComplete**](LegalAPI.md#CloudPostV1LegalDocumentsByIdSignComplete) | **Post** /v1/legal/documents/{id}/sign/complete | 
[**CloudPostV1LegalDocumentsIdSign**](LegalAPI.md#CloudPostV1LegalDocumentsIdSign) | **Post** /v1/legal/documents/{id}/sign | RequestLegalSignature opens an e-signature request over one document and moves it to out_for_signature, returning the provider&#39;s reference for the request.
[**CloudPostV1LegalFilings**](LegalAPI.md#CloudPostV1LegalFilings) | **Post** /v1/legal/filings | CreateLegalFiling records a filing of one or more of the org&#39;s documents with a state or agency, and returns the tracking record.
[**CloudPutV1LegalTemplatesId**](LegalAPI.md#CloudPutV1LegalTemplatesId) | **Put** /v1/legal/templates/{id} | SaveLegalTemplateOverride saves the org&#39;s own version of a template — a custom NDA, a house MSA — and returns it with its new version number.



## CloudGetV1LegalDocuments

> CloudDocumentPage CloudGetV1LegalDocuments(ctx).Limit(limit).Execute()

ListLegalDocuments returns the org's generated documents, newest first, WITHOUT their rendered content — fetch one document to read its body.



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
	limit := int32(56) // int32 | Limit bounds the page. Absent or unparseable means the store's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.CloudGetV1LegalDocuments(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudGetV1LegalDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1LegalDocuments`: CloudDocumentPage
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudGetV1LegalDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1LegalDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds the page. Absent or unparseable means the store&#39;s own default. | 

### Return type

[**CloudDocumentPage**](CloudDocumentPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1LegalDocumentsId

> CloudDocumentReply CloudGetV1LegalDocumentsId(ctx, id).Execute()

GetLegalDocument returns one of the org's documents WITH its rendered body.



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
	id := "id_example" // string | ID is the document's server-minted handle, \"doc_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.CloudGetV1LegalDocumentsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudGetV1LegalDocumentsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1LegalDocumentsId`: CloudDocumentReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudGetV1LegalDocumentsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document&#39;s server-minted handle, \&quot;doc_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1LegalDocumentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDocumentReply**](CloudDocumentReply.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1LegalFilings

> CloudFilingPage CloudGetV1LegalFilings(ctx).Limit(limit).Execute()

ListLegalFilings returns the org's filing records, newest first — which documents were filed where, through which provider, and what the filing's honest status is.



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
	limit := int32(56) // int32 | Limit bounds the page. Absent or unparseable means the store's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.CloudGetV1LegalFilings(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudGetV1LegalFilings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1LegalFilings`: CloudFilingPage
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudGetV1LegalFilings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1LegalFilingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds the page. Absent or unparseable means the store&#39;s own default. | 

### Return type

[**CloudFilingPage**](CloudFilingPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1LegalHealth

> CloudLegalHealth CloudGetV1LegalHealth(ctx).Execute()

LegalHealth reports that the legal subsystem is serving and how many built-in templates its catalog carries.



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
	resp, r, err := apiClient.LegalAPI.CloudGetV1LegalHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudGetV1LegalHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1LegalHealth`: CloudLegalHealth
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudGetV1LegalHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1LegalHealthRequest struct via the builder pattern


### Return type

[**CloudLegalHealth**](CloudLegalHealth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1LegalTemplates

> CloudTemplateCatalog CloudGetV1LegalTemplates(ctx).Execute()

ListLegalTemplates returns the org's effective template catalog: every built-in template, with any the org has overridden replaced by its own latest version.



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
	resp, r, err := apiClient.LegalAPI.CloudGetV1LegalTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudGetV1LegalTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1LegalTemplates`: CloudTemplateCatalog
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudGetV1LegalTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1LegalTemplatesRequest struct via the builder pattern


### Return type

[**CloudTemplateCatalog**](CloudTemplateCatalog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1LegalTemplatesId

> CloudTemplateReply CloudGetV1LegalTemplatesId(ctx, id).Execute()

GetLegalTemplate returns one template resolved for the caller's org — the org's own override if it has saved one, else the built-in — with its full text/template body and its declared merge fields.



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
	id := "id_example" // string | ID is the template's stable id, e.g. \"nda\" or \"safe\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.CloudGetV1LegalTemplatesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudGetV1LegalTemplatesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1LegalTemplatesId`: CloudTemplateReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudGetV1LegalTemplatesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the template&#39;s stable id, e.g. \&quot;nda\&quot; or \&quot;safe\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1LegalTemplatesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudTemplateReply**](CloudTemplateReply.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1LegalDocuments

> CloudDocumentReply CloudPostV1LegalDocuments(ctx).CloudGenerateRequest(cloudGenerateRequest).Execute()

GenerateLegalDocument renders a document from a template and the caller's own merge data, seals it in the org's store, and returns it with its rendered body.



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
	cloudGenerateRequest := *openapiclient.NewCloudGenerateRequest() // CloudGenerateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.CloudPostV1LegalDocuments(context.Background()).CloudGenerateRequest(cloudGenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudPostV1LegalDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1LegalDocuments`: CloudDocumentReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudPostV1LegalDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1LegalDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudGenerateRequest** | [**CloudGenerateRequest**](CloudGenerateRequest.md) |  | 

### Return type

[**CloudDocumentReply**](CloudDocumentReply.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1LegalDocumentsByIdSignComplete

> CloudPostV1LegalDocumentsByIdSignComplete(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LegalAPI.CloudPostV1LegalDocumentsByIdSignComplete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudPostV1LegalDocumentsByIdSignComplete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1LegalDocumentsByIdSignCompleteRequest struct via the builder pattern


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


## CloudPostV1LegalDocumentsIdSign

> CloudSignReply CloudPostV1LegalDocumentsIdSign(ctx, id).CloudSignRequest(cloudSignRequest).Execute()

RequestLegalSignature opens an e-signature request over one document and moves it to out_for_signature, returning the provider's reference for the request.



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
	id := "doc_1f…" // string | ID is the document to send for signature, from the path.
	cloudSignRequest := *openapiclient.NewCloudSignRequest() // CloudSignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.CloudPostV1LegalDocumentsIdSign(context.Background(), id).CloudSignRequest(cloudSignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudPostV1LegalDocumentsIdSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1LegalDocumentsIdSign`: CloudSignReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudPostV1LegalDocumentsIdSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document to send for signature, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1LegalDocumentsIdSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudSignRequest** | [**CloudSignRequest**](CloudSignRequest.md) |  | 

### Return type

[**CloudSignReply**](CloudSignReply.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1LegalFilings

> CloudFilingReply CloudPostV1LegalFilings(ctx).CloudFilingRequest(cloudFilingRequest).Execute()

CreateLegalFiling records a filing of one or more of the org's documents with a state or agency, and returns the tracking record.



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
	cloudFilingRequest := *openapiclient.NewCloudFilingRequest() // CloudFilingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.CloudPostV1LegalFilings(context.Background()).CloudFilingRequest(cloudFilingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudPostV1LegalFilings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1LegalFilings`: CloudFilingReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudPostV1LegalFilings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1LegalFilingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudFilingRequest** | [**CloudFilingRequest**](CloudFilingRequest.md) |  | 

### Return type

[**CloudFilingReply**](CloudFilingReply.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1LegalTemplatesId

> CloudTemplateReply CloudPutV1LegalTemplatesId(ctx, id).CloudTemplateOverride(cloudTemplateOverride).Execute()

SaveLegalTemplateOverride saves the org's own version of a template — a custom NDA, a house MSA — and returns it with its new version number.



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
	id := "nda" // string | ID is the template to override, from the path. Overriding a built-in id inherits that built-in's category, title and counsel-review posture.
	cloudTemplateOverride := *openapiclient.NewCloudTemplateOverride() // CloudTemplateOverride | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegalAPI.CloudPutV1LegalTemplatesId(context.Background(), id).CloudTemplateOverride(cloudTemplateOverride).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegalAPI.CloudPutV1LegalTemplatesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1LegalTemplatesId`: CloudTemplateReply
	fmt.Fprintf(os.Stdout, "Response from `LegalAPI.CloudPutV1LegalTemplatesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the template to override, from the path. Overriding a built-in id inherits that built-in&#39;s category, title and counsel-review posture. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1LegalTemplatesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudTemplateOverride** | [**CloudTemplateOverride**](CloudTemplateOverride.md) |  | 

### Return type

[**CloudTemplateReply**](CloudTemplateReply.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

