# \AccessTokensAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCreateAccessToken**](AccessTokensAPI.md#CommerceCreateAccessToken) | **Post** /v1/commerce/access/{mode}/{id} | Create or delete access token
[**CommerceDeleteAccessToken**](AccessTokensAPI.md#CommerceDeleteAccessToken) | **Delete** /v1/commerce/access/{mode}/{id} | Delete access token
[**CommerceGetAccessToken**](AccessTokensAPI.md#CommerceGetAccessToken) | **Get** /v1/commerce/access/{mode}/{id} | Get access token



## CommerceCreateAccessToken

> map[string]interface{} CommerceCreateAccessToken(ctx, mode, id).Execute()

Create or delete access token

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
	mode := "mode_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.CommerceCreateAccessToken(context.Background(), mode, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.CommerceCreateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateAccessToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.CommerceCreateAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mode** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateAccessTokenRequest struct via the builder pattern


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


## CommerceDeleteAccessToken

> CommerceDeleteAccessToken(ctx, mode, id).Execute()

Delete access token

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
	mode := "mode_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessTokensAPI.CommerceDeleteAccessToken(context.Background(), mode, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.CommerceDeleteAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mode** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceDeleteAccessTokenRequest struct via the builder pattern


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


## CommerceGetAccessToken

> CommerceAccessToken CommerceGetAccessToken(ctx, mode, id).Email(email).Password(password).Execute()

Get access token

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
	mode := "mode_example" // string | 
	id := "id_example" // string | 
	email := "email_example" // string |  (optional)
	password := "password_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.CommerceGetAccessToken(context.Background(), mode, id).Email(email).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.CommerceGetAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetAccessToken`: CommerceAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.CommerceGetAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mode** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **email** | **string** |  | 
 **password** | **string** |  | 

### Return type

[**CommerceAccessToken**](CommerceAccessToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

