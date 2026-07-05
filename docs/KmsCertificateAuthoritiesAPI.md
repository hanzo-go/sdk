# \KmsCertificateAuthoritiesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateCertificateAuthority**](KmsCertificateAuthoritiesAPI.md#KmsCreateCertificateAuthority) | **Post** /v1/kms/cert-manager/ca | Create a certificate authority
[**KmsDeleteCertificateAuthority**](KmsCertificateAuthoritiesAPI.md#KmsDeleteCertificateAuthority) | **Delete** /v1/kms/cert-manager/ca/{caId} | Delete a certificate authority
[**KmsGetCertificateAuthority**](KmsCertificateAuthoritiesAPI.md#KmsGetCertificateAuthority) | **Get** /v1/kms/cert-manager/ca/{caId} | Get a certificate authority by ID
[**KmsUpdateCertificateAuthority**](KmsCertificateAuthoritiesAPI.md#KmsUpdateCertificateAuthority) | **Patch** /v1/kms/cert-manager/ca/{caId} | Update a certificate authority



## KmsCreateCertificateAuthority

> KmsCreateCertificateAuthority200Response KmsCreateCertificateAuthority(ctx).KmsCreateCertificateAuthorityRequest(kmsCreateCertificateAuthorityRequest).Execute()

Create a certificate authority

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
	kmsCreateCertificateAuthorityRequest := *openapiclient.NewKmsCreateCertificateAuthorityRequest("ProjectId_example", "Type_example", "FriendlyName_example", "CommonName_example", "KeyAlgorithm_example") // KmsCreateCertificateAuthorityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsCertificateAuthoritiesAPI.KmsCreateCertificateAuthority(context.Background()).KmsCreateCertificateAuthorityRequest(kmsCreateCertificateAuthorityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsCertificateAuthoritiesAPI.KmsCreateCertificateAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateCertificateAuthority`: KmsCreateCertificateAuthority200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsCertificateAuthoritiesAPI.KmsCreateCertificateAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateCertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateCertificateAuthorityRequest** | [**KmsCreateCertificateAuthorityRequest**](KmsCreateCertificateAuthorityRequest.md) |  | 

### Return type

[**KmsCreateCertificateAuthority200Response**](KmsCreateCertificateAuthority200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteCertificateAuthority

> map[string]interface{} KmsDeleteCertificateAuthority(ctx, caId).Execute()

Delete a certificate authority

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsCertificateAuthoritiesAPI.KmsDeleteCertificateAuthority(context.Background(), caId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsCertificateAuthoritiesAPI.KmsDeleteCertificateAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteCertificateAuthority`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsCertificateAuthoritiesAPI.KmsDeleteCertificateAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteCertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetCertificateAuthority

> KmsCreateCertificateAuthority200Response KmsGetCertificateAuthority(ctx, caId).Execute()

Get a certificate authority by ID

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsCertificateAuthoritiesAPI.KmsGetCertificateAuthority(context.Background(), caId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsCertificateAuthoritiesAPI.KmsGetCertificateAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetCertificateAuthority`: KmsCreateCertificateAuthority200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsCertificateAuthoritiesAPI.KmsGetCertificateAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetCertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsCreateCertificateAuthority200Response**](KmsCreateCertificateAuthority200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateCertificateAuthority

> KmsCreateCertificateAuthority200Response KmsUpdateCertificateAuthority(ctx, caId).DnsUpdateZoneRequest(dnsUpdateZoneRequest).Execute()

Update a certificate authority

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
	dnsUpdateZoneRequest := *openapiclient.NewDnsUpdateZoneRequest() // DnsUpdateZoneRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsCertificateAuthoritiesAPI.KmsUpdateCertificateAuthority(context.Background(), caId).DnsUpdateZoneRequest(dnsUpdateZoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsCertificateAuthoritiesAPI.KmsUpdateCertificateAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateCertificateAuthority`: KmsCreateCertificateAuthority200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsCertificateAuthoritiesAPI.KmsUpdateCertificateAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateCertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsUpdateZoneRequest** | [**DnsUpdateZoneRequest**](DnsUpdateZoneRequest.md) |  | 

### Return type

[**KmsCreateCertificateAuthority200Response**](KmsCreateCertificateAuthority200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

