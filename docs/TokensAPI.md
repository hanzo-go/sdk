# \TokensAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotCreateToken**](TokensAPI.md#BotCreateToken) | **Post** /v1/bot/tokens | Create a new API token
[**BotListTokens**](TokensAPI.md#BotListTokens) | **Get** /v1/bot/tokens | List current user&#39;s API tokens
[**BotRevokeToken**](TokensAPI.md#BotRevokeToken) | **Delete** /v1/bot/tokens/{id} | Revoke an API token
[**IamApiControllerAddToken**](TokensAPI.md#IamApiControllerAddToken) | **Post** /v1/iam/tokens | Api Controller Add Token
[**IamApiControllerDeleteToken**](TokensAPI.md#IamApiControllerDeleteToken) | **Delete** /v1/iam/tokens/{id} | Api Controller Delete Token
[**IamApiControllerGetCaptchaStatus**](TokensAPI.md#IamApiControllerGetCaptchaStatus) | **Get** /v1/iam/captcha/status | Api Controller Get Captcha Status
[**IamApiControllerGetOAuthToken**](TokensAPI.md#IamApiControllerGetOAuthToken) | **Post** /oauth/token | Api Controller Get O Auth Token
[**IamApiControllerGetToken**](TokensAPI.md#IamApiControllerGetToken) | **Get** /v1/iam/tokens/{id} | Api Controller Get Token
[**IamApiControllerGetTokens**](TokensAPI.md#IamApiControllerGetTokens) | **Get** /v1/iam/tokens | Api Controller Get Tokens
[**IamApiControllerRefreshToken**](TokensAPI.md#IamApiControllerRefreshToken) | **Post** /oauth/token/refresh | Api Controller Refresh Token
[**IamApiControllerUpdateToken**](TokensAPI.md#IamApiControllerUpdateToken) | **Put** /v1/iam/tokens/{id} | Api Controller Update Token



## BotCreateToken

> BotCreateToken200Response BotCreateToken(ctx).BotCreateTokenRequest(botCreateTokenRequest).Execute()

Create a new API token

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
	botCreateTokenRequest := *openapiclient.NewBotCreateTokenRequest("Label_example") // BotCreateTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.BotCreateToken(context.Background()).BotCreateTokenRequest(botCreateTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.BotCreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotCreateToken`: BotCreateToken200Response
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.BotCreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botCreateTokenRequest** | [**BotCreateTokenRequest**](BotCreateTokenRequest.md) |  | 

### Return type

[**BotCreateToken200Response**](BotCreateToken200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListTokens

> BotListTokens200Response BotListTokens(ctx).Execute()

List current user's API tokens

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
	resp, r, err := apiClient.TokensAPI.BotListTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.BotListTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListTokens`: BotListTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.BotListTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBotListTokensRequest struct via the builder pattern


### Return type

[**BotListTokens200Response**](BotListTokens200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotRevokeToken

> AnalyticsHeartbeat200Response BotRevokeToken(ctx, id).Execute()

Revoke an API token

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.BotRevokeToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.BotRevokeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotRevokeToken`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.BotRevokeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.TokensAPI.IamApiControllerAddToken(context.Background()).IamObjectToken(iamObjectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.IamApiControllerAddToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddToken`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.IamApiControllerAddToken`: %v\n", resp)
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
	resp, r, err := apiClient.TokensAPI.IamApiControllerDeleteToken(context.Background(), id).IamObjectToken(iamObjectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.IamApiControllerDeleteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteToken`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.IamApiControllerDeleteToken`: %v\n", resp)
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
	resp, r, err := apiClient.TokensAPI.IamApiControllerGetCaptchaStatus(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.IamApiControllerGetCaptchaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetCaptchaStatus`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.IamApiControllerGetCaptchaStatus`: %v\n", resp)
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

> IamObjectTokenWrapper IamApiControllerGetOAuthToken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Code(code).IamControllersTokenRequest(iamControllersTokenRequest).Execute()

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
	iamControllersTokenRequest := *openapiclient.NewIamControllersTokenRequest() // IamControllersTokenRequest | Optional. Params may be supplied as query string (documented above) or in the body; body values fill only fields still empty from the query. JSON (TokenRequest) is tried when the body is non-empty and grant_type is not device_code, otherwise form-urlencoded is parsed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.IamApiControllerGetOAuthToken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Code(code).IamControllersTokenRequest(iamControllersTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.IamApiControllerGetOAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetOAuthToken`: IamObjectTokenWrapper
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.IamApiControllerGetOAuthToken`: %v\n", resp)
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
 **iamControllersTokenRequest** | [**IamControllersTokenRequest**](IamControllersTokenRequest.md) | Optional. Params may be supplied as query string (documented above) or in the body; body values fill only fields still empty from the query. JSON (TokenRequest) is tried when the body is non-empty and grant_type is not device_code, otherwise form-urlencoded is parsed. | 

### Return type

[**IamObjectTokenWrapper**](IamObjectTokenWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
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
	resp, r, err := apiClient.TokensAPI.IamApiControllerGetToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.IamApiControllerGetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetToken`: IamObjectToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.IamApiControllerGetToken`: %v\n", resp)
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
	resp, r, err := apiClient.TokensAPI.IamApiControllerGetTokens(context.Background()).Owner(owner).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.IamApiControllerGetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetTokens`: []IamObjectToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.IamApiControllerGetTokens`: %v\n", resp)
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

> IamObjectTokenWrapper IamApiControllerRefreshToken(ctx).GrantType(grantType).RefreshToken(refreshToken).Scope(scope).ClientId(clientId).ClientSecret(clientSecret).IamControllersTokenRequest(iamControllersTokenRequest).Execute()

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
	iamControllersTokenRequest := *openapiclient.NewIamControllersTokenRequest() // IamControllersTokenRequest | Optional. Read only when client_id is absent from the query/Basic-auth; then JSON (TokenRequest) or form-urlencoded fills client_id, client_secret, grant_type, scope, refresh_token. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.IamApiControllerRefreshToken(context.Background()).GrantType(grantType).RefreshToken(refreshToken).Scope(scope).ClientId(clientId).ClientSecret(clientSecret).IamControllersTokenRequest(iamControllersTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.IamApiControllerRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerRefreshToken`: IamObjectTokenWrapper
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.IamApiControllerRefreshToken`: %v\n", resp)
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
 **iamControllersTokenRequest** | [**IamControllersTokenRequest**](IamControllersTokenRequest.md) | Optional. Read only when client_id is absent from the query/Basic-auth; then JSON (TokenRequest) or form-urlencoded fills client_id, client_secret, grant_type, scope, refresh_token. | 

### Return type

[**IamObjectTokenWrapper**](IamObjectTokenWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
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
	resp, r, err := apiClient.TokensAPI.IamApiControllerUpdateToken(context.Background(), id).IamObjectToken(iamObjectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.IamApiControllerUpdateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateToken`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.IamApiControllerUpdateToken`: %v\n", resp)
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

