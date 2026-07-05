# \KmsCertificatesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsGetCertificate**](KmsCertificatesAPI.md#KmsGetCertificate) | **Get** /v1/kms/cert-manager/certificates/{certificateId} | Get a certificate by ID
[**KmsIssueCertificate**](KmsCertificatesAPI.md#KmsIssueCertificate) | **Post** /v1/kms/cert-manager/ca/{caId}/issue-certificate | Issue a certificate from a CA
[**KmsListCertificates**](KmsCertificatesAPI.md#KmsListCertificates) | **Get** /v1/kms/cert-manager/certificates | List certificates
[**KmsRevokeCertificate**](KmsCertificatesAPI.md#KmsRevokeCertificate) | **Delete** /v1/kms/cert-manager/certificates/{certificateId} | Revoke a certificate



## KmsGetCertificate

> KmsGetCertificate200Response KmsGetCertificate(ctx, certificateId).Execute()

Get a certificate by ID

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
	certificateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsCertificatesAPI.KmsGetCertificate(context.Background(), certificateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsCertificatesAPI.KmsGetCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetCertificate`: KmsGetCertificate200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsCertificatesAPI.KmsGetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsGetCertificate200Response**](KmsGetCertificate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsIssueCertificate

> KmsIssueCertificate200Response KmsIssueCertificate(ctx, caId).KmsIssueCertificateRequest(kmsIssueCertificateRequest).Execute()

Issue a certificate from a CA

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
	caId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsIssueCertificateRequest := *openapiclient.NewKmsIssueCertificateRequest("FriendlyName_example", "CommonName_example") // KmsIssueCertificateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsCertificatesAPI.KmsIssueCertificate(context.Background(), caId).KmsIssueCertificateRequest(kmsIssueCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsCertificatesAPI.KmsIssueCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsIssueCertificate`: KmsIssueCertificate200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsCertificatesAPI.KmsIssueCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsIssueCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsIssueCertificateRequest** | [**KmsIssueCertificateRequest**](KmsIssueCertificateRequest.md) |  | 

### Return type

[**KmsIssueCertificate200Response**](KmsIssueCertificate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListCertificates

> KmsListCertificates200Response KmsListCertificates(ctx).CaId(caId).Offset(offset).Limit(limit).Execute()

List certificates

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
	caId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsCertificatesAPI.KmsListCertificates(context.Background()).CaId(caId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsCertificatesAPI.KmsListCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListCertificates`: KmsListCertificates200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsCertificatesAPI.KmsListCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **caId** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**KmsListCertificates200Response**](KmsListCertificates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsRevokeCertificate

> map[string]interface{} KmsRevokeCertificate(ctx, certificateId).KmsRevokeCertificateRequest(kmsRevokeCertificateRequest).Execute()

Revoke a certificate

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
	certificateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsRevokeCertificateRequest := *openapiclient.NewKmsRevokeCertificateRequest() // KmsRevokeCertificateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsCertificatesAPI.KmsRevokeCertificate(context.Background(), certificateId).KmsRevokeCertificateRequest(kmsRevokeCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsCertificatesAPI.KmsRevokeCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsRevokeCertificate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsCertificatesAPI.KmsRevokeCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsRevokeCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsRevokeCertificateRequest** | [**KmsRevokeCertificateRequest**](KmsRevokeCertificateRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

