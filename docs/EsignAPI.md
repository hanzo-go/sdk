# \EsignAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEsignDocuments**](EsignAPI.md#GetEsignDocuments) | **Get** /v1/esign/documents | Returns your org&#39;s documents, newest first.
[**GetEsignDocumentsById**](EsignAPI.md#GetEsignDocumentsById) | **Get** /v1/esign/documents/{id} | Returns one document with its recipients and field layout.
[**GetEsignDocumentsByIdAudit**](EsignAPI.md#GetEsignDocumentsByIdAudit) | **Get** /v1/esign/documents/{id}/audit | Returns the document&#39;s full audit trail, oldest first.
[**GetEsignDocumentsByIdDownload**](EsignAPI.md#GetEsignDocumentsByIdDownload) | **Get** /v1/esign/documents/{id}/download | Returns the document — the sealed PDF once it is complete.
[**GetEsignHealth**](EsignAPI.md#GetEsignHealth) | **Get** /v1/esign/health | Reports whether the e-signature surface is mounted.
[**GetEsignOByOrgSignByToken**](EsignAPI.md#GetEsignOByOrgSignByToken) | **Get** /v1/esign/o/{org}/sign/{token} | Opens a document you were asked to sign, using your signing link.
[**PostEsignDocuments**](EsignAPI.md#PostEsignDocuments) | **Post** /v1/esign/documents | Uploads a PDF and opens a draft ready for recipients and fields.
[**PostEsignDocumentsByIdFields**](EsignAPI.md#PostEsignDocumentsByIdFields) | **Post** /v1/esign/documents/{id}/fields | Places a field on the page for one recipient to fill.
[**PostEsignDocumentsByIdRecipients**](EsignAPI.md#PostEsignDocumentsByIdRecipients) | **Post** /v1/esign/documents/{id}/recipients | Adds someone to a draft and mints their signing token.
[**PostEsignDocumentsByIdSend**](EsignAPI.md#PostEsignDocumentsByIdSend) | **Post** /v1/esign/documents/{id}/send | Sends the document out and answers each signer&#39;s link.
[**PostEsignOByOrgSignByTokenComplete**](EsignAPI.md#PostEsignOByOrgSignByTokenComplete) | **Post** /v1/esign/o/{org}/sign/{token}/complete | Finishes your signing — and seals the document if you were the last.
[**PostEsignOByOrgSignByTokenFieldsByFieldid**](EsignAPI.md#PostEsignOByOrgSignByTokenFieldsByFieldid) | **Post** /v1/esign/o/{org}/sign/{token}/fields/{fieldId} | Fills in one of your fields.
[**PostEsignOByOrgSignByTokenReject**](EsignAPI.md#PostEsignOByOrgSignByTokenReject) | **Post** /v1/esign/o/{org}/sign/{token}/reject | Declines to sign, with an optional reason.



## GetEsignDocuments

> EsignEsignDocuments GetEsignDocuments(ctx).Execute()

Returns your org's documents, newest first.



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
	resp, r, err := apiClient.EsignAPI.GetEsignDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEsignDocuments`: EsignEsignDocuments
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.GetEsignDocuments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEsignDocumentsRequest struct via the builder pattern


### Return type

[**EsignEsignDocuments**](EsignEsignDocuments.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignDocumentsById

> EsignEsignDocument GetEsignDocumentsById(ctx, id).Execute()

Returns one document with its recipients and field layout.



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
	id := "id_example" // string | ID is the document to act on. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id belonging to another tenant is simply not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.GetEsignDocumentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEsignDocumentsById`: EsignEsignDocument
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.GetEsignDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document to act on. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id belonging to another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEsignDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsignEsignDocument**](EsignEsignDocument.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignDocumentsByIdAudit

> EsignEsignTrail GetEsignDocumentsByIdAudit(ctx, id).Execute()

Returns the document's full audit trail, oldest first.



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
	id := "id_example" // string | ID is the document to act on. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id belonging to another tenant is simply not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.GetEsignDocumentsByIdAudit(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignDocumentsByIdAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEsignDocumentsByIdAudit`: EsignEsignTrail
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.GetEsignDocumentsByIdAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document to act on. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id belonging to another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEsignDocumentsByIdAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsignEsignTrail**](EsignEsignTrail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignDocumentsByIdDownload

> EsignEsignPDF GetEsignDocumentsByIdDownload(ctx, id).Execute()

Returns the document — the sealed PDF once it is complete.



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
	id := "id_example" // string | ID is the document to act on. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id belonging to another tenant is simply not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.GetEsignDocumentsByIdDownload(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignDocumentsByIdDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEsignDocumentsByIdDownload`: EsignEsignPDF
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.GetEsignDocumentsByIdDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document to act on. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id belonging to another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEsignDocumentsByIdDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsignEsignPDF**](EsignEsignPDF.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignHealth

> EsignEsignHealth GetEsignHealth(ctx).Execute()

Reports whether the e-signature surface is mounted.



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
	resp, r, err := apiClient.EsignAPI.GetEsignHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEsignHealth`: EsignEsignHealth
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.GetEsignHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEsignHealthRequest struct via the builder pattern


### Return type

[**EsignEsignHealth**](EsignEsignHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignOByOrgSignByToken

> EsignEsignSession GetEsignOByOrgSignByToken(ctx, org, token).Execute()

Opens a document you were asked to sign, using your signing link.



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
	org := "org_example" // string | 
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.GetEsignOByOrgSignByToken(context.Background(), org, token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignOByOrgSignByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEsignOByOrgSignByToken`: EsignEsignSession
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.GetEsignOByOrgSignByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEsignOByOrgSignByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EsignEsignSession**](EsignEsignSession.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignDocuments

> EsignEsignDocument PostEsignDocuments(ctx).EsignEsignUploadIn(esignEsignUploadIn).Execute()

Uploads a PDF and opens a draft ready for recipients and fields.



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
	esignEsignUploadIn := *openapiclient.NewEsignEsignUploadIn() // EsignEsignUploadIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.PostEsignDocuments(context.Background()).EsignEsignUploadIn(esignEsignUploadIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEsignDocuments`: EsignEsignDocument
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.PostEsignDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEsignDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esignEsignUploadIn** | [**EsignEsignUploadIn**](EsignEsignUploadIn.md) |  | 

### Return type

[**EsignEsignDocument**](EsignEsignDocument.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignDocumentsByIdFields

> EsignEsignPlacement PostEsignDocumentsByIdFields(ctx, id).EsignEsignFieldIn(esignEsignFieldIn).Execute()

Places a field on the page for one recipient to fill.



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
	esignEsignFieldIn := *openapiclient.NewEsignEsignFieldIn() // EsignEsignFieldIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.PostEsignDocumentsByIdFields(context.Background(), id).EsignEsignFieldIn(esignEsignFieldIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignDocumentsByIdFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEsignDocumentsByIdFields`: EsignEsignPlacement
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.PostEsignDocumentsByIdFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEsignDocumentsByIdFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esignEsignFieldIn** | [**EsignEsignFieldIn**](EsignEsignFieldIn.md) |  | 

### Return type

[**EsignEsignPlacement**](EsignEsignPlacement.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignDocumentsByIdRecipients

> EsignEsignInvite PostEsignDocumentsByIdRecipients(ctx, id).EsignEsignRecipientIn(esignEsignRecipientIn).Execute()

Adds someone to a draft and mints their signing token.



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
	esignEsignRecipientIn := *openapiclient.NewEsignEsignRecipientIn() // EsignEsignRecipientIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.PostEsignDocumentsByIdRecipients(context.Background(), id).EsignEsignRecipientIn(esignEsignRecipientIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignDocumentsByIdRecipients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEsignDocumentsByIdRecipients`: EsignEsignInvite
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.PostEsignDocumentsByIdRecipients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEsignDocumentsByIdRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esignEsignRecipientIn** | [**EsignEsignRecipientIn**](EsignEsignRecipientIn.md) |  | 

### Return type

[**EsignEsignInvite**](EsignEsignInvite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignDocumentsByIdSend

> EsignEsignLinks PostEsignDocumentsByIdSend(ctx, id).Execute()

Sends the document out and answers each signer's link.



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
	id := "id_example" // string | ID is the document to send. The URL is the addressing authority.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.PostEsignDocumentsByIdSend(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignDocumentsByIdSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEsignDocumentsByIdSend`: EsignEsignLinks
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.PostEsignDocumentsByIdSend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document to send. The URL is the addressing authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEsignDocumentsByIdSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsignEsignLinks**](EsignEsignLinks.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignOByOrgSignByTokenComplete

> EsignEsignCompletion PostEsignOByOrgSignByTokenComplete(ctx, org, token).Execute()

Finishes your signing — and seals the document if you were the last.



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
	org := "org_example" // string | 
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.PostEsignOByOrgSignByTokenComplete(context.Background(), org, token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignOByOrgSignByTokenComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEsignOByOrgSignByTokenComplete`: EsignEsignCompletion
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.PostEsignOByOrgSignByTokenComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEsignOByOrgSignByTokenCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EsignEsignCompletion**](EsignEsignCompletion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignOByOrgSignByTokenFieldsByFieldid

> EsignEsignInsertion PostEsignOByOrgSignByTokenFieldsByFieldid(ctx, org, token, fieldId).EsignEsignValueIn(esignEsignValueIn).Execute()

Fills in one of your fields.



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
	org := "org_example" // string | 
	token := "token_example" // string | 
	fieldId := "fieldId_example" // string | 
	esignEsignValueIn := *openapiclient.NewEsignEsignValueIn() // EsignEsignValueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.PostEsignOByOrgSignByTokenFieldsByFieldid(context.Background(), org, token, fieldId).EsignEsignValueIn(esignEsignValueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignOByOrgSignByTokenFieldsByFieldid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEsignOByOrgSignByTokenFieldsByFieldid`: EsignEsignInsertion
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.PostEsignOByOrgSignByTokenFieldsByFieldid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**token** | **string** |  | 
**fieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEsignOByOrgSignByTokenFieldsByFieldidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **esignEsignValueIn** | [**EsignEsignValueIn**](EsignEsignValueIn.md) |  | 

### Return type

[**EsignEsignInsertion**](EsignEsignInsertion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignOByOrgSignByTokenReject

> EsignEsignRejection PostEsignOByOrgSignByTokenReject(ctx, org, token).EsignEsignRejectIn(esignEsignRejectIn).Execute()

Declines to sign, with an optional reason.



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
	org := "org_example" // string | 
	token := "token_example" // string | 
	esignEsignRejectIn := *openapiclient.NewEsignEsignRejectIn() // EsignEsignRejectIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsignAPI.PostEsignOByOrgSignByTokenReject(context.Background(), org, token).EsignEsignRejectIn(esignEsignRejectIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignOByOrgSignByTokenReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEsignOByOrgSignByTokenReject`: EsignEsignRejection
	fmt.Fprintf(os.Stdout, "Response from `EsignAPI.PostEsignOByOrgSignByTokenReject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEsignOByOrgSignByTokenRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **esignEsignRejectIn** | [**EsignEsignRejectIn**](EsignEsignRejectIn.md) |  | 

### Return type

[**EsignEsignRejection**](EsignEsignRejection.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

