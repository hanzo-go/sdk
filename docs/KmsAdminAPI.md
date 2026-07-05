# \KmsAdminAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsGetServerConfig**](KmsAdminAPI.md#KmsGetServerConfig) | **Get** /v1/kms/admin/config | Get server configuration
[**KmsUpdateServerConfig**](KmsAdminAPI.md#KmsUpdateServerConfig) | **Patch** /v1/kms/admin/config | Update server configuration



## KmsGetServerConfig

> KmsGetServerConfig200Response KmsGetServerConfig(ctx).Execute()

Get server configuration

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
	resp, r, err := apiClient.KmsAdminAPI.KmsGetServerConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAdminAPI.KmsGetServerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetServerConfig`: KmsGetServerConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsAdminAPI.KmsGetServerConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetServerConfigRequest struct via the builder pattern


### Return type

[**KmsGetServerConfig200Response**](KmsGetServerConfig200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateServerConfig

> map[string]interface{} KmsUpdateServerConfig(ctx).KmsUpdateServerConfigRequest(kmsUpdateServerConfigRequest).Execute()

Update server configuration

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
	kmsUpdateServerConfigRequest := *openapiclient.NewKmsUpdateServerConfigRequest() // KmsUpdateServerConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAdminAPI.KmsUpdateServerConfig(context.Background()).KmsUpdateServerConfigRequest(kmsUpdateServerConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAdminAPI.KmsUpdateServerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateServerConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsAdminAPI.KmsUpdateServerConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateServerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsUpdateServerConfigRequest** | [**KmsUpdateServerConfigRequest**](KmsUpdateServerConfigRequest.md) |  | 

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

