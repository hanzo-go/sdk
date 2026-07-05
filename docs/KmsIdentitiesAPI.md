# \KmsIdentitiesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateIdentity**](KmsIdentitiesAPI.md#KmsCreateIdentity) | **Post** /v1/kms/identities | Create a machine identity
[**KmsDeleteIdentity**](KmsIdentitiesAPI.md#KmsDeleteIdentity) | **Delete** /v1/kms/identities/{identityId} | Delete an identity
[**KmsGetIdentity**](KmsIdentitiesAPI.md#KmsGetIdentity) | **Get** /v1/kms/identities/{identityId} | Get an identity by ID
[**KmsUpdateIdentity**](KmsIdentitiesAPI.md#KmsUpdateIdentity) | **Patch** /v1/kms/identities/{identityId} | Update an identity



## KmsCreateIdentity

> KmsCreateIdentity200Response KmsCreateIdentity(ctx).KmsCreateIdentityRequest(kmsCreateIdentityRequest).Execute()

Create a machine identity

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
	kmsCreateIdentityRequest := *openapiclient.NewKmsCreateIdentityRequest("Name_example", "OrganizationId_example", "Role_example") // KmsCreateIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentitiesAPI.KmsCreateIdentity(context.Background()).KmsCreateIdentityRequest(kmsCreateIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentitiesAPI.KmsCreateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateIdentity`: KmsCreateIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentitiesAPI.KmsCreateIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateIdentityRequest** | [**KmsCreateIdentityRequest**](KmsCreateIdentityRequest.md) |  | 

### Return type

[**KmsCreateIdentity200Response**](KmsCreateIdentity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteIdentity

> map[string]interface{} KmsDeleteIdentity(ctx, identityId).Execute()

Delete an identity

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
	identityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentitiesAPI.KmsDeleteIdentity(context.Background(), identityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentitiesAPI.KmsDeleteIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteIdentity`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentitiesAPI.KmsDeleteIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteIdentityRequest struct via the builder pattern


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


## KmsGetIdentity

> KmsCreateIdentity200Response KmsGetIdentity(ctx, identityId).Execute()

Get an identity by ID

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
	identityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentitiesAPI.KmsGetIdentity(context.Background(), identityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentitiesAPI.KmsGetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetIdentity`: KmsCreateIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentitiesAPI.KmsGetIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsCreateIdentity200Response**](KmsCreateIdentity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateIdentity

> KmsCreateIdentity200Response KmsUpdateIdentity(ctx, identityId).KmsUpdateIdentityRequest(kmsUpdateIdentityRequest).Execute()

Update an identity

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
	identityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateIdentityRequest := *openapiclient.NewKmsUpdateIdentityRequest() // KmsUpdateIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentitiesAPI.KmsUpdateIdentity(context.Background(), identityId).KmsUpdateIdentityRequest(kmsUpdateIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentitiesAPI.KmsUpdateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateIdentity`: KmsCreateIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentitiesAPI.KmsUpdateIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateIdentityRequest** | [**KmsUpdateIdentityRequest**](KmsUpdateIdentityRequest.md) |  | 

### Return type

[**KmsCreateIdentity200Response**](KmsCreateIdentity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

