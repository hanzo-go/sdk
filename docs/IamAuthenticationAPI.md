# \IamAuthenticationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerCallback**](IamAuthenticationAPI.md#IamApiControllerCallback) | **Post** /oauth/callback | Api Controller Callback
[**IamApiControllerDeviceAuth**](IamAuthenticationAPI.md#IamApiControllerDeviceAuth) | **Post** /v1/iam/auth/device | Api Controller Device Auth
[**IamApiControllerFaceIDSigninBegin**](IamAuthenticationAPI.md#IamApiControllerFaceIDSigninBegin) | **Get** /v1/iam/auth/faceid/begin | Api Controller Face ID Signin Begin
[**IamApiControllerGetApplicationLogin**](IamAuthenticationAPI.md#IamApiControllerGetApplicationLogin) | **Get** /v1/iam/auth/app-login | Api Controller Get Application Login
[**IamApiControllerGetCaptcha**](IamAuthenticationAPI.md#IamApiControllerGetCaptcha) | **Get** /v1/iam/captcha | Api Controller Get Captcha
[**IamApiControllerIntrospectToken**](IamAuthenticationAPI.md#IamApiControllerIntrospectToken) | **Post** /oauth/introspect | Api Controller Introspect Token
[**IamApiControllerLogin**](IamAuthenticationAPI.md#IamApiControllerLogin) | **Post** /v1/iam/auth/login | Api Controller Login
[**IamApiControllerLogout**](IamAuthenticationAPI.md#IamApiControllerLogout) | **Post** /v1/iam/auth/logout | Api Controller Logout
[**IamApiControllerSignup**](IamAuthenticationAPI.md#IamApiControllerSignup) | **Post** /v1/iam/auth/signup | Api Controller Signup
[**IamApiControllerSsoLogoutGet**](IamAuthenticationAPI.md#IamApiControllerSsoLogoutGet) | **Get** /v1/iam/sso-logout | Api Controller Sso Logout
[**IamApiControllerSsoLogoutPost**](IamAuthenticationAPI.md#IamApiControllerSsoLogoutPost) | **Post** /v1/iam/sso-logout | Api Controller Sso Logout
[**IamApiControllerUnlink**](IamAuthenticationAPI.md#IamApiControllerUnlink) | **Post** /v1/iam/unlink | Api Controller Unlink
[**IamApiControllerWebAuthnSigninBegin**](IamAuthenticationAPI.md#IamApiControllerWebAuthnSigninBegin) | **Get** /v1/iam/auth/webauthn/signin/begin | Api Controller Web Authn Signin Begin
[**IamApiControllerWebAuthnSigninFinish**](IamAuthenticationAPI.md#IamApiControllerWebAuthnSigninFinish) | **Post** /v1/iam/auth/webauthn/signin/finish | Api Controller Web Authn Signin Finish
[**IamRootControllerGetJwks**](IamAuthenticationAPI.md#IamRootControllerGetJwks) | **Get** /.well-known/jwks | Root Controller Get Jwks
[**IamRootControllerGetJwksByApplication**](IamAuthenticationAPI.md#IamRootControllerGetJwksByApplication) | **Get** /.well-known/{application}/jwks | Root Controller Get Jwks By Application
[**IamRootControllerGetOidcDiscovery**](IamAuthenticationAPI.md#IamRootControllerGetOidcDiscovery) | **Get** /.well-known/openid-configuration | Root Controller Get Oidc Discovery
[**IamRootControllerGetOidcDiscoveryByApplication**](IamAuthenticationAPI.md#IamRootControllerGetOidcDiscoveryByApplication) | **Get** /.well-known/{application}/openid-configuration | Root Controller Get Oidc Discovery By Application
[**IamRootControllerGetWebFinger**](IamAuthenticationAPI.md#IamRootControllerGetWebFinger) | **Get** /.well-known/webfinger | Root Controller Get Web Finger
[**IamRootControllerGetWebFingerByApplication**](IamAuthenticationAPI.md#IamRootControllerGetWebFingerByApplication) | **Get** /.well-known/{application}/webfinger | Root Controller Get Web Finger By Application



## IamApiControllerCallback

> IamObjectUserinfo IamApiControllerCallback(ctx).Execute()

Api Controller Callback



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
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerCallback`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerCallback`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerCallbackRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeviceAuth

> IamObjectDeviceAuthResponse IamApiControllerDeviceAuth(ctx).Execute()

Api Controller Device Auth



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
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerDeviceAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerDeviceAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeviceAuth`: IamObjectDeviceAuthResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerDeviceAuth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeviceAuthRequest struct via the builder pattern


### Return type

[**IamObjectDeviceAuthResponse**](IamObjectDeviceAuthResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerFaceIDSigninBegin

> IamControllersResponse IamApiControllerFaceIDSigninBegin(ctx).Owner(owner).Name(name).Execute()

Api Controller Face ID Signin Begin



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
	owner := "owner_example" // string | owner
	name := "name_example" // string | name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerFaceIDSigninBegin(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerFaceIDSigninBegin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerFaceIDSigninBegin`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerFaceIDSigninBegin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerFaceIDSigninBeginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | owner | 
 **name** | **string** | name | 

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


## IamApiControllerGetApplicationLogin

> IamControllersResponse IamApiControllerGetApplicationLogin(ctx).ClientId(clientId).ResponseType(responseType).RedirectUri(redirectUri).Scope(scope).State(state).Execute()

Api Controller Get Application Login



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
	clientId := "clientId_example" // string | client id
	responseType := "responseType_example" // string | response type
	redirectUri := "redirectUri_example" // string | redirect uri
	scope := "scope_example" // string | scope
	state := "state_example" // string | state

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerGetApplicationLogin(context.Background()).ClientId(clientId).ResponseType(responseType).RedirectUri(redirectUri).Scope(scope).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerGetApplicationLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetApplicationLogin`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerGetApplicationLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetApplicationLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | client id | 
 **responseType** | **string** | response type | 
 **redirectUri** | **string** | redirect uri | 
 **scope** | **string** | scope | 
 **state** | **string** | state | 

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


## IamApiControllerGetCaptcha

> IamObjectUserinfo IamApiControllerGetCaptcha(ctx).Execute()

Api Controller Get Captcha

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
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerGetCaptcha(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerGetCaptcha``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetCaptcha`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerGetCaptcha`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetCaptchaRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerIntrospectToken

> IamObjectIntrospectionResponse IamApiControllerIntrospectToken(ctx).Token(token).TokenTypeHint(tokenTypeHint).Execute()

Api Controller Introspect Token



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
	token := "token_example" // string | access_token's value or refresh_token's value
	tokenTypeHint := "tokenTypeHint_example" // string | the token type access_token or refresh_token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerIntrospectToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerIntrospectToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerIntrospectToken`: IamObjectIntrospectionResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerIntrospectToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerIntrospectTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | access_token&#39;s value or refresh_token&#39;s value | 
 **tokenTypeHint** | **string** | the token type access_token or refresh_token | 

### Return type

[**IamObjectIntrospectionResponse**](IamObjectIntrospectionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerLogin

> IamControllersResponse IamApiControllerLogin(ctx).ClientId(clientId).ResponseType(responseType).RedirectUri(redirectUri).Body(body).Scope(scope).State(state).Nonce(nonce).CodeChallengeMethod(codeChallengeMethod).CodeChallenge(codeChallenge).Execute()

Api Controller Login



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
	clientId := "clientId_example" // string | clientId
	responseType := "responseType_example" // string | responseType
	redirectUri := "redirectUri_example" // string | redirectUri
	body := map[string]interface{}{ ... } // map[string]interface{} | Login information
	scope := "scope_example" // string | scope (optional)
	state := "state_example" // string | state (optional)
	nonce := "nonce_example" // string | nonce (optional)
	codeChallengeMethod := "codeChallengeMethod_example" // string | code_challenge_method (optional)
	codeChallenge := "codeChallenge_example" // string | code_challenge (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerLogin(context.Background()).ClientId(clientId).ResponseType(responseType).RedirectUri(redirectUri).Body(body).Scope(scope).State(state).Nonce(nonce).CodeChallengeMethod(codeChallengeMethod).CodeChallenge(codeChallenge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerLogin`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | clientId | 
 **responseType** | **string** | responseType | 
 **redirectUri** | **string** | redirectUri | 
 **body** | **map[string]interface{}** | Login information | 
 **scope** | **string** | scope | 
 **state** | **string** | state | 
 **nonce** | **string** | nonce | 
 **codeChallengeMethod** | **string** | code_challenge_method | 
 **codeChallenge** | **string** | code_challenge | 

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


## IamApiControllerLogout

> IamControllersResponse IamApiControllerLogout(ctx).IdTokenHint(idTokenHint).PostLogoutRedirectUri(postLogoutRedirectUri).State(state).Execute()

Api Controller Logout



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
	idTokenHint := "idTokenHint_example" // string | id_token_hint (optional)
	postLogoutRedirectUri := "postLogoutRedirectUri_example" // string | post_logout_redirect_uri (optional)
	state := "state_example" // string | state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerLogout(context.Background()).IdTokenHint(idTokenHint).PostLogoutRedirectUri(postLogoutRedirectUri).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerLogout`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerLogout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idTokenHint** | **string** | id_token_hint | 
 **postLogoutRedirectUri** | **string** | post_logout_redirect_uri | 
 **state** | **string** | state | 

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


## IamApiControllerSignup

> IamControllersResponse IamApiControllerSignup(ctx).Username(username).Password(password).Execute()

Api Controller Signup



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
	username := "username_example" // string | The username to sign up
	password := "password_example" // string | The password

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerSignup(context.Background()).Username(username).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerSignup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSignup`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerSignup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | The username to sign up | 
 **password** | **string** | The password | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSsoLogoutGet

> IamControllersResponse IamApiControllerSsoLogoutGet(ctx).LogoutAll(logoutAll).Execute()

Api Controller Sso Logout



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
	logoutAll := "logoutAll_example" // string | Whether to logout from all sessions. Accepted values: 'true', '1', or empty (default: true). Any other value means false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerSsoLogoutGet(context.Background()).LogoutAll(logoutAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerSsoLogoutGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSsoLogoutGet`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerSsoLogoutGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSsoLogoutGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logoutAll** | **string** | Whether to logout from all sessions. Accepted values: &#39;true&#39;, &#39;1&#39;, or empty (default: true). Any other value means false. | 

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


## IamApiControllerSsoLogoutPost

> IamControllersResponse IamApiControllerSsoLogoutPost(ctx).LogoutAll(logoutAll).Execute()

Api Controller Sso Logout



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
	logoutAll := "logoutAll_example" // string | Whether to logout from all sessions. Accepted values: 'true', '1', or empty (default: true). Any other value means false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerSsoLogoutPost(context.Background()).LogoutAll(logoutAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerSsoLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSsoLogoutPost`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerSsoLogoutPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSsoLogoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logoutAll** | **string** | Whether to logout from all sessions. Accepted values: &#39;true&#39;, &#39;1&#39;, or empty (default: true). Any other value means false. | 

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


## IamApiControllerUnlink

> IamObjectUserinfo IamApiControllerUnlink(ctx).Execute()

Api Controller Unlink

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
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerUnlink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerUnlink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUnlink`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerUnlink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUnlinkRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerWebAuthnSigninBegin

> map[string]interface{} IamApiControllerWebAuthnSigninBegin(ctx).Owner(owner).Name(name).Execute()

Api Controller Web Authn Signin Begin



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
	owner := "owner_example" // string | owner
	name := "name_example" // string | name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerWebAuthnSigninBegin(context.Background()).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerWebAuthnSigninBegin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerWebAuthnSigninBegin`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerWebAuthnSigninBegin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerWebAuthnSigninBeginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | owner | 
 **name** | **string** | name | 

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


## IamApiControllerWebAuthnSigninFinish

> IamControllersResponse IamApiControllerWebAuthnSigninFinish(ctx).Body(body).Execute()

Api Controller Web Authn Signin Finish



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
	body := map[string]interface{}{ ... } // map[string]interface{} | authenticator assertion Response

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamApiControllerWebAuthnSigninFinish(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamApiControllerWebAuthnSigninFinish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerWebAuthnSigninFinish`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamApiControllerWebAuthnSigninFinish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerWebAuthnSigninFinishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | authenticator assertion Response | 

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


## IamRootControllerGetJwks

> map[string]interface{} IamRootControllerGetJwks(ctx).Execute()

Root Controller Get Jwks

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
	resp, r, err := apiClient.IamAuthenticationAPI.IamRootControllerGetJwks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamRootControllerGetJwks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamRootControllerGetJwks`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamRootControllerGetJwks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamRootControllerGetJwksRequest struct via the builder pattern


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


## IamRootControllerGetJwksByApplication

> map[string]interface{} IamRootControllerGetJwksByApplication(ctx, application).Execute()

Root Controller Get Jwks By Application

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
	application := "application_example" // string | application name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamRootControllerGetJwksByApplication(context.Background(), application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamRootControllerGetJwksByApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamRootControllerGetJwksByApplication`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamRootControllerGetJwksByApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application name | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamRootControllerGetJwksByApplicationRequest struct via the builder pattern


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


## IamRootControllerGetOidcDiscovery

> IamObjectOidcDiscovery IamRootControllerGetOidcDiscovery(ctx).Execute()

Root Controller Get Oidc Discovery



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
	resp, r, err := apiClient.IamAuthenticationAPI.IamRootControllerGetOidcDiscovery(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamRootControllerGetOidcDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamRootControllerGetOidcDiscovery`: IamObjectOidcDiscovery
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamRootControllerGetOidcDiscovery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamRootControllerGetOidcDiscoveryRequest struct via the builder pattern


### Return type

[**IamObjectOidcDiscovery**](IamObjectOidcDiscovery.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamRootControllerGetOidcDiscoveryByApplication

> IamObjectOidcDiscovery IamRootControllerGetOidcDiscoveryByApplication(ctx, application).Execute()

Root Controller Get Oidc Discovery By Application



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
	application := "application_example" // string | application name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamRootControllerGetOidcDiscoveryByApplication(context.Background(), application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamRootControllerGetOidcDiscoveryByApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamRootControllerGetOidcDiscoveryByApplication`: IamObjectOidcDiscovery
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamRootControllerGetOidcDiscoveryByApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application name | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamRootControllerGetOidcDiscoveryByApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectOidcDiscovery**](IamObjectOidcDiscovery.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamRootControllerGetWebFinger

> IamObjectWebFinger IamRootControllerGetWebFinger(ctx).Resource(resource).Execute()

Root Controller Get Web Finger

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
	resource := "resource_example" // string | resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamRootControllerGetWebFinger(context.Background()).Resource(resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamRootControllerGetWebFinger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamRootControllerGetWebFinger`: IamObjectWebFinger
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamRootControllerGetWebFinger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamRootControllerGetWebFingerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource** | **string** | resource | 

### Return type

[**IamObjectWebFinger**](IamObjectWebFinger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamRootControllerGetWebFingerByApplication

> IamObjectWebFinger IamRootControllerGetWebFingerByApplication(ctx, application).Resource(resource).Execute()

Root Controller Get Web Finger By Application

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
	application := "application_example" // string | application name
	resource := "resource_example" // string | resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAuthenticationAPI.IamRootControllerGetWebFingerByApplication(context.Background(), application).Resource(resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAuthenticationAPI.IamRootControllerGetWebFingerByApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamRootControllerGetWebFingerByApplication`: IamObjectWebFinger
	fmt.Fprintf(os.Stdout, "Response from `IamAuthenticationAPI.IamRootControllerGetWebFingerByApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application name | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamRootControllerGetWebFingerByApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resource** | **string** | resource | 

### Return type

[**IamObjectWebFinger**](IamObjectWebFinger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

