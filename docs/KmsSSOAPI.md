# \KmsSSOAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateSsoConfig**](KmsSSOAPI.md#KmsCreateSsoConfig) | **Post** /v1/kms/sso/config | Create SSO configuration
[**KmsGetSsoConfig**](KmsSSOAPI.md#KmsGetSsoConfig) | **Get** /v1/kms/sso/config | Get SSO configuration for an organization



## KmsCreateSsoConfig

> map[string]interface{} KmsCreateSsoConfig(ctx).KmsCreateSsoConfigRequest(kmsCreateSsoConfigRequest).Execute()

Create SSO configuration

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
	kmsCreateSsoConfigRequest := *openapiclient.NewKmsCreateSsoConfigRequest("OrgId_example", "Type_example") // KmsCreateSsoConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSSOAPI.KmsCreateSsoConfig(context.Background()).KmsCreateSsoConfigRequest(kmsCreateSsoConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSSOAPI.KmsCreateSsoConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateSsoConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsSSOAPI.KmsCreateSsoConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateSsoConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateSsoConfigRequest** | [**KmsCreateSsoConfigRequest**](KmsCreateSsoConfigRequest.md) |  | 

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


## KmsGetSsoConfig

> KmsGetSsoConfig200Response KmsGetSsoConfig(ctx).OrgId(orgId).Execute()

Get SSO configuration for an organization

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSSOAPI.KmsGetSsoConfig(context.Background()).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSSOAPI.KmsGetSsoConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetSsoConfig`: KmsGetSsoConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSSOAPI.KmsGetSsoConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetSsoConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** |  | 

### Return type

[**KmsGetSsoConfig200Response**](KmsGetSsoConfig200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

