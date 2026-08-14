# \EsignAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEsignDocuments**](EsignAPI.md#GetEsignDocuments) | **Get** /v1/esign/documents | Your org&#39;s documents, newest first
[**GetEsignDocumentsById**](EsignAPI.md#GetEsignDocumentsById) | **Get** /v1/esign/documents/{id} | One document with its recipients and field layout
[**GetEsignDocumentsByIdAudit**](EsignAPI.md#GetEsignDocumentsByIdAudit) | **Get** /v1/esign/documents/{id}/audit | The document&#39;s full audit trail, oldest first
[**GetEsignDocumentsByIdDownload**](EsignAPI.md#GetEsignDocumentsByIdDownload) | **Get** /v1/esign/documents/{id}/download | Download the document — the sealed PDF once it is complete
[**GetEsignHealth**](EsignAPI.md#GetEsignHealth) | **Get** /v1/esign/health | Whether the e-signature surface is mounted
[**GetEsignOByOrgSignByToken**](EsignAPI.md#GetEsignOByOrgSignByToken) | **Get** /v1/esign/o/{org}/sign/{token} | Open a document you were asked to sign, using your signing link
[**PostEsignDocuments**](EsignAPI.md#PostEsignDocuments) | **Post** /v1/esign/documents | Upload a PDF and open a draft ready for recipients and fields
[**PostEsignDocumentsByIdFields**](EsignAPI.md#PostEsignDocumentsByIdFields) | **Post** /v1/esign/documents/{id}/fields | Place a field on the page for one recipient to fill
[**PostEsignDocumentsByIdRecipients**](EsignAPI.md#PostEsignDocumentsByIdRecipients) | **Post** /v1/esign/documents/{id}/recipients | Add someone to a draft and mint their signing token
[**PostEsignDocumentsByIdSend**](EsignAPI.md#PostEsignDocumentsByIdSend) | **Post** /v1/esign/documents/{id}/send | Send the document out and get each signer&#39;s link
[**PostEsignOByOrgSignByTokenComplete**](EsignAPI.md#PostEsignOByOrgSignByTokenComplete) | **Post** /v1/esign/o/{org}/sign/{token}/complete | Finish signing — and seal the document if you were the last
[**PostEsignOByOrgSignByTokenFieldsByFieldid**](EsignAPI.md#PostEsignOByOrgSignByTokenFieldsByFieldid) | **Post** /v1/esign/o/{org}/sign/{token}/fields/{fieldId} | Fill in one of your fields
[**PostEsignOByOrgSignByTokenReject**](EsignAPI.md#PostEsignOByOrgSignByTokenReject) | **Post** /v1/esign/o/{org}/sign/{token}/reject | Decline to sign, with an optional reason



## GetEsignDocuments

> GetEsignDocuments(ctx).Execute()

Your org's documents, newest first



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
	r, err := apiClient.EsignAPI.GetEsignDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEsignDocumentsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignDocumentsById

> GetEsignDocumentsById(ctx, id).Execute()

One document with its recipients and field layout



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
	r, err := apiClient.EsignAPI.GetEsignDocumentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignDocumentsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetEsignDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignDocumentsByIdAudit

> GetEsignDocumentsByIdAudit(ctx, id).Execute()

The document's full audit trail, oldest first



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
	r, err := apiClient.EsignAPI.GetEsignDocumentsByIdAudit(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignDocumentsByIdAudit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetEsignDocumentsByIdAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignDocumentsByIdDownload

> GetEsignDocumentsByIdDownload(ctx, id).Execute()

Download the document — the sealed PDF once it is complete



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
	r, err := apiClient.EsignAPI.GetEsignDocumentsByIdDownload(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignDocumentsByIdDownload``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetEsignDocumentsByIdDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignHealth

> GetEsignHealth(ctx).Execute()

Whether the e-signature surface is mounted



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
	r, err := apiClient.EsignAPI.GetEsignHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEsignHealthRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEsignOByOrgSignByToken

> GetEsignOByOrgSignByToken(ctx, org, token).Execute()

Open a document you were asked to sign, using your signing link



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
	org := "org_example" // string | 
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EsignAPI.GetEsignOByOrgSignByToken(context.Background(), org, token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.GetEsignOByOrgSignByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignDocuments

> PostEsignDocuments(ctx).Execute()

Upload a PDF and open a draft ready for recipients and fields



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
	r, err := apiClient.EsignAPI.PostEsignDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostEsignDocumentsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignDocumentsByIdFields

> PostEsignDocumentsByIdFields(ctx, id).Execute()

Place a field on the page for one recipient to fill



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
	r, err := apiClient.EsignAPI.PostEsignDocumentsByIdFields(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignDocumentsByIdFields``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostEsignDocumentsByIdFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignDocumentsByIdRecipients

> PostEsignDocumentsByIdRecipients(ctx, id).Execute()

Add someone to a draft and mint their signing token



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
	r, err := apiClient.EsignAPI.PostEsignDocumentsByIdRecipients(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignDocumentsByIdRecipients``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostEsignDocumentsByIdRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignDocumentsByIdSend

> PostEsignDocumentsByIdSend(ctx, id).Execute()

Send the document out and get each signer's link



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
	r, err := apiClient.EsignAPI.PostEsignDocumentsByIdSend(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignDocumentsByIdSend``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostEsignDocumentsByIdSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignOByOrgSignByTokenComplete

> PostEsignOByOrgSignByTokenComplete(ctx, org, token).Execute()

Finish signing — and seal the document if you were the last



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
	org := "org_example" // string | 
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EsignAPI.PostEsignOByOrgSignByTokenComplete(context.Background(), org, token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignOByOrgSignByTokenComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignOByOrgSignByTokenFieldsByFieldid

> PostEsignOByOrgSignByTokenFieldsByFieldid(ctx, org, token, fieldId).Execute()

Fill in one of your fields



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
	org := "org_example" // string | 
	token := "token_example" // string | 
	fieldId := "fieldId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EsignAPI.PostEsignOByOrgSignByTokenFieldsByFieldid(context.Background(), org, token, fieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignOByOrgSignByTokenFieldsByFieldid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEsignOByOrgSignByTokenReject

> PostEsignOByOrgSignByTokenReject(ctx, org, token).Execute()

Decline to sign, with an optional reason



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
	org := "org_example" // string | 
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EsignAPI.PostEsignOByOrgSignByTokenReject(context.Background(), org, token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsignAPI.PostEsignOByOrgSignByTokenReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

