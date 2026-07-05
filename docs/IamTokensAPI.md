# \IamTokensAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddToken**](IamTokensAPI.md#IamApiControllerAddToken) | **Post** /v1/iam/tokens | Api Controller Add Token
[**IamApiControllerDeleteToken**](IamTokensAPI.md#IamApiControllerDeleteToken) | **Delete** /v1/iam/tokens/{id} | Api Controller Delete Token
[**IamApiControllerGetCaptchaStatus**](IamTokensAPI.md#IamApiControllerGetCaptchaStatus) | **Get** /v1/iam/captcha/status | Api Controller Get Captcha Status
[**IamApiControllerGetOAuthToken**](IamTokensAPI.md#IamApiControllerGetOAuthToken) | **Post** /oauth/token | Api Controller Get O Auth Token
[**IamApiControllerGetToken**](IamTokensAPI.md#IamApiControllerGetToken) | **Get** /v1/iam/tokens/{id} | Api Controller Get Token
[**IamApiControllerGetTokens**](IamTokensAPI.md#IamApiControllerGetTokens) | **Get** /v1/iam/tokens | Api Controller Get Tokens
[**IamApiControllerRefreshToken**](IamTokensAPI.md#IamApiControllerRefreshToken) | **Post** /oauth/token/refresh | Api Controller Refresh Token
[**IamApiControllerUpdateToken**](IamTokensAPI.md#IamApiControllerUpdateToken) | **Put** /v1/iam/tokens/{id} | Api Controller Update Token



## IamApiControllerAddToken

> IamControllersResponse IamApiControllerAddToken(ctx).IamObjectToken(iamObjectToken).Execute()

Api Controller Add Token



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
	iamObjectToken := *openapiclient.NewIamObjectToken() // IamObjectToken | Details of the token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamTokensAPI.IamApiControllerAddToken(context.Background()).IamObjectToken(iamObjectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamTokensAPI.IamApiControllerAddToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddToken`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamTokensAPI.IamApiControllerAddToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectToken** | [**IamObjectToken**](IamObjectToken.md) | Details of the token | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteToken

> IamControllersResponse IamApiControllerDeleteToken(ctx, id).IamObjectToken(iamObjectToken).Execute()

Api Controller Delete Token



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectToken := *openapiclient.NewIamObjectToken() // IamObjectToken | Details of the token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamTokensAPI.IamApiControllerDeleteToken(context.Background(), id).IamObjectToken(iamObjectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamTokensAPI.IamApiControllerDeleteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteToken`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamTokensAPI.IamApiControllerDeleteToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectToken** | [**IamObjectToken**](IamObjectToken.md) | Details of the token | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetCaptchaStatus

> IamControllersResponse IamApiControllerGetCaptchaStatus(ctx).Id(id).Execute()

Api Controller Get Captcha Status



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
	id := "id_example" // string | The id ( owner/name ) of user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamTokensAPI.IamApiControllerGetCaptchaStatus(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamTokensAPI.IamApiControllerGetCaptchaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetCaptchaStatus`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamTokensAPI.IamApiControllerGetCaptchaStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetCaptchaStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of user | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetOAuthToken

> IamObjectTokenWrapper IamApiControllerGetOAuthToken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Code(code).Execute()

Api Controller Get O Auth Token



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
	grantType := "grantType_example" // string | OAuth grant type
	clientId := "clientId_example" // string | OAuth client id
	clientSecret := "clientSecret_example" // string | OAuth client secret
	code := "code_example" // string | OAuth code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamTokensAPI.IamApiControllerGetOAuthToken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamTokensAPI.IamApiControllerGetOAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetOAuthToken`: IamObjectTokenWrapper
	fmt.Fprintf(os.Stdout, "Response from `IamTokensAPI.IamApiControllerGetOAuthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetOAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | OAuth grant type | 
 **clientId** | **string** | OAuth client id | 
 **clientSecret** | **string** | OAuth client secret | 
 **code** | **string** | OAuth code | 

### Return type

[**IamObjectTokenWrapper**](IamObjectTokenWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetToken

> IamObjectToken IamApiControllerGetToken(ctx, id).Execute()

Api Controller Get Token



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
	id := "id_example" // string | The token ID in format: organization/token-name (e.g., built-in/token-123456)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamTokensAPI.IamApiControllerGetToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamTokensAPI.IamApiControllerGetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetToken`: IamObjectToken
	fmt.Fprintf(os.Stdout, "Response from `IamTokensAPI.IamApiControllerGetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The token ID in format: organization/token-name (e.g., built-in/token-123456) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectToken**](IamObjectToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetTokens

> []IamObjectToken IamApiControllerGetTokens(ctx).Owner(owner).PageSize(pageSize).P(p).Execute()

Api Controller Get Tokens



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
	owner := "owner_example" // string | The organization name (e.g., built-in)
	pageSize := "pageSize_example" // string | The size of each page
	p := "p_example" // string | The number of the page

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamTokensAPI.IamApiControllerGetTokens(context.Background()).Owner(owner).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamTokensAPI.IamApiControllerGetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetTokens`: []IamObjectToken
	fmt.Fprintf(os.Stdout, "Response from `IamTokensAPI.IamApiControllerGetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The organization name (e.g., built-in) | 
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

### Return type

[**[]IamObjectToken**](IamObjectToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerRefreshToken

> IamObjectTokenWrapper IamApiControllerRefreshToken(ctx).GrantType(grantType).RefreshToken(refreshToken).Scope(scope).ClientId(clientId).ClientSecret(clientSecret).Execute()

Api Controller Refresh Token



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
	grantType := "grantType_example" // string | OAuth grant type
	refreshToken := "refreshToken_example" // string | OAuth refresh token
	scope := "scope_example" // string | OAuth scope
	clientId := "clientId_example" // string | OAuth client id
	clientSecret := "clientSecret_example" // string | OAuth client secret (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamTokensAPI.IamApiControllerRefreshToken(context.Background()).GrantType(grantType).RefreshToken(refreshToken).Scope(scope).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamTokensAPI.IamApiControllerRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerRefreshToken`: IamObjectTokenWrapper
	fmt.Fprintf(os.Stdout, "Response from `IamTokensAPI.IamApiControllerRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | OAuth grant type | 
 **refreshToken** | **string** | OAuth refresh token | 
 **scope** | **string** | OAuth scope | 
 **clientId** | **string** | OAuth client id | 
 **clientSecret** | **string** | OAuth client secret | 

### Return type

[**IamObjectTokenWrapper**](IamObjectTokenWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateToken

> IamControllersResponse IamApiControllerUpdateToken(ctx, id).IamObjectToken(iamObjectToken).Execute()

Api Controller Update Token



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
	id := "id_example" // string | The token ID in format: organization/token-name (e.g., built-in/token-123456)
	iamObjectToken := *openapiclient.NewIamObjectToken() // IamObjectToken | Details of the token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamTokensAPI.IamApiControllerUpdateToken(context.Background(), id).IamObjectToken(iamObjectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamTokensAPI.IamApiControllerUpdateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateToken`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamTokensAPI.IamApiControllerUpdateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The token ID in format: organization/token-name (e.g., built-in/token-123456) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectToken** | [**IamObjectToken**](IamObjectToken.md) | Details of the token | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

