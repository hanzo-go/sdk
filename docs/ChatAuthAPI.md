# \ChatAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetAuth2faEnable**](ChatAuthAPI.md#ChatGetAuth2faEnable) | **Get** /v1/chat/auth/2fa/enable | Enable 2FA
[**ChatGetAuthGraphToken**](ChatAuthAPI.md#ChatGetAuthGraphToken) | **Get** /v1/chat/auth/graph-token | Get Microsoft Graph token
[**ChatPostAuth2faBackupRegenerate**](ChatAuthAPI.md#ChatPostAuth2faBackupRegenerate) | **Post** /v1/chat/auth/2fa/backup/regenerate | Regenerate 2FA backup codes
[**ChatPostAuth2faConfirm**](ChatAuthAPI.md#ChatPostAuth2faConfirm) | **Post** /v1/chat/auth/2fa/confirm | Confirm 2FA activation
[**ChatPostAuth2faDisable**](ChatAuthAPI.md#ChatPostAuth2faDisable) | **Post** /v1/chat/auth/2fa/disable | Disable 2FA
[**ChatPostAuth2faVerify**](ChatAuthAPI.md#ChatPostAuth2faVerify) | **Post** /v1/chat/auth/2fa/verify | Verify 2FA setup
[**ChatPostAuth2faVerifyTemp**](ChatAuthAPI.md#ChatPostAuth2faVerifyTemp) | **Post** /v1/chat/auth/2fa/verify-temp | Verify 2FA with temporary token
[**ChatPostAuthLogin**](ChatAuthAPI.md#ChatPostAuthLogin) | **Post** /v1/chat/auth/login | Login
[**ChatPostAuthLogout**](ChatAuthAPI.md#ChatPostAuthLogout) | **Post** /v1/chat/auth/logout | Logout
[**ChatPostAuthRefresh**](ChatAuthAPI.md#ChatPostAuthRefresh) | **Post** /v1/chat/auth/refresh | Refresh token
[**ChatPostAuthRegister**](ChatAuthAPI.md#ChatPostAuthRegister) | **Post** /v1/chat/auth/register | Register a new user
[**ChatPostAuthRequestpasswordreset**](ChatAuthAPI.md#ChatPostAuthRequestpasswordreset) | **Post** /v1/chat/auth/requestPasswordReset | Request password reset email
[**ChatPostAuthResetpassword**](ChatAuthAPI.md#ChatPostAuthResetpassword) | **Post** /v1/chat/auth/resetPassword | Reset password with token



## ChatGetAuth2faEnable

> ChatGetAuth2faEnable200Response ChatGetAuth2faEnable(ctx).Execute()

Enable 2FA



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
	resp, r, err := apiClient.ChatAuthAPI.ChatGetAuth2faEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatGetAuth2faEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAuth2faEnable`: ChatGetAuth2faEnable200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatGetAuth2faEnable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAuth2faEnableRequest struct via the builder pattern


### Return type

[**ChatGetAuth2faEnable200Response**](ChatGetAuth2faEnable200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAuthGraphToken

> map[string]interface{} ChatGetAuthGraphToken(ctx).Execute()

Get Microsoft Graph token



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
	resp, r, err := apiClient.ChatAuthAPI.ChatGetAuthGraphToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatGetAuthGraphToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAuthGraphToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatGetAuthGraphToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAuthGraphTokenRequest struct via the builder pattern


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


## ChatPostAuth2faBackupRegenerate

> ChatPostAuth2faBackupRegenerate200Response ChatPostAuth2faBackupRegenerate(ctx).ChatPostAuth2faBackupRegenerateRequest(chatPostAuth2faBackupRegenerateRequest).Execute()

Regenerate 2FA backup codes

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
	chatPostAuth2faBackupRegenerateRequest := *openapiclient.NewChatPostAuth2faBackupRegenerateRequest("Token_example") // ChatPostAuth2faBackupRegenerateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuth2faBackupRegenerate(context.Background()).ChatPostAuth2faBackupRegenerateRequest(chatPostAuth2faBackupRegenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuth2faBackupRegenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuth2faBackupRegenerate`: ChatPostAuth2faBackupRegenerate200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuth2faBackupRegenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuth2faBackupRegenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuth2faBackupRegenerateRequest** | [**ChatPostAuth2faBackupRegenerateRequest**](ChatPostAuth2faBackupRegenerateRequest.md) |  | 

### Return type

[**ChatPostAuth2faBackupRegenerate200Response**](ChatPostAuth2faBackupRegenerate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAuth2faConfirm

> ChatPostAuth2faBackupRegenerate200Response ChatPostAuth2faConfirm(ctx).ChatPostAuth2faBackupRegenerateRequest(chatPostAuth2faBackupRegenerateRequest).Execute()

Confirm 2FA activation

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
	chatPostAuth2faBackupRegenerateRequest := *openapiclient.NewChatPostAuth2faBackupRegenerateRequest("Token_example") // ChatPostAuth2faBackupRegenerateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuth2faConfirm(context.Background()).ChatPostAuth2faBackupRegenerateRequest(chatPostAuth2faBackupRegenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuth2faConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuth2faConfirm`: ChatPostAuth2faBackupRegenerate200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuth2faConfirm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuth2faConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuth2faBackupRegenerateRequest** | [**ChatPostAuth2faBackupRegenerateRequest**](ChatPostAuth2faBackupRegenerateRequest.md) |  | 

### Return type

[**ChatPostAuth2faBackupRegenerate200Response**](ChatPostAuth2faBackupRegenerate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAuth2faDisable

> map[string]interface{} ChatPostAuth2faDisable(ctx).ChatPostAuth2faBackupRegenerateRequest(chatPostAuth2faBackupRegenerateRequest).Execute()

Disable 2FA

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
	chatPostAuth2faBackupRegenerateRequest := *openapiclient.NewChatPostAuth2faBackupRegenerateRequest("Token_example") // ChatPostAuth2faBackupRegenerateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuth2faDisable(context.Background()).ChatPostAuth2faBackupRegenerateRequest(chatPostAuth2faBackupRegenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuth2faDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuth2faDisable`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuth2faDisable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuth2faDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuth2faBackupRegenerateRequest** | [**ChatPostAuth2faBackupRegenerateRequest**](ChatPostAuth2faBackupRegenerateRequest.md) |  | 

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


## ChatPostAuth2faVerify

> map[string]interface{} ChatPostAuth2faVerify(ctx).ChatPostAuth2faBackupRegenerateRequest(chatPostAuth2faBackupRegenerateRequest).Execute()

Verify 2FA setup



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
	chatPostAuth2faBackupRegenerateRequest := *openapiclient.NewChatPostAuth2faBackupRegenerateRequest("Token_example") // ChatPostAuth2faBackupRegenerateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuth2faVerify(context.Background()).ChatPostAuth2faBackupRegenerateRequest(chatPostAuth2faBackupRegenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuth2faVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuth2faVerify`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuth2faVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuth2faVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuth2faBackupRegenerateRequest** | [**ChatPostAuth2faBackupRegenerateRequest**](ChatPostAuth2faBackupRegenerateRequest.md) |  | 

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


## ChatPostAuth2faVerifyTemp

> ChatAuthResponse ChatPostAuth2faVerifyTemp(ctx).ChatPostAuth2faVerifyTempRequest(chatPostAuth2faVerifyTempRequest).Execute()

Verify 2FA with temporary token



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
	chatPostAuth2faVerifyTempRequest := *openapiclient.NewChatPostAuth2faVerifyTempRequest("TempToken_example", "Token_example") // ChatPostAuth2faVerifyTempRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuth2faVerifyTemp(context.Background()).ChatPostAuth2faVerifyTempRequest(chatPostAuth2faVerifyTempRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuth2faVerifyTemp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuth2faVerifyTemp`: ChatAuthResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuth2faVerifyTemp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuth2faVerifyTempRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuth2faVerifyTempRequest** | [**ChatPostAuth2faVerifyTempRequest**](ChatPostAuth2faVerifyTempRequest.md) |  | 

### Return type

[**ChatAuthResponse**](ChatAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAuthLogin

> ChatAuthResponse ChatPostAuthLogin(ctx).ChatPostAuthLoginRequest(chatPostAuthLoginRequest).Execute()

Login



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
	chatPostAuthLoginRequest := *openapiclient.NewChatPostAuthLoginRequest("Email_example", "Password_example") // ChatPostAuthLoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuthLogin(context.Background()).ChatPostAuthLoginRequest(chatPostAuthLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuthLogin`: ChatAuthResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuthLoginRequest** | [**ChatPostAuthLoginRequest**](ChatPostAuthLoginRequest.md) |  | 

### Return type

[**ChatAuthResponse**](ChatAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAuthLogout

> map[string]interface{} ChatPostAuthLogout(ctx).Execute()

Logout



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
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuthLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuthLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuthLogout`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuthLogout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuthLogoutRequest struct via the builder pattern


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


## ChatPostAuthRefresh

> ChatAuthResponse ChatPostAuthRefresh(ctx).ChatPostAuthRefreshRequest(chatPostAuthRefreshRequest).Execute()

Refresh token



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
	chatPostAuthRefreshRequest := *openapiclient.NewChatPostAuthRefreshRequest() // ChatPostAuthRefreshRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuthRefresh(context.Background()).ChatPostAuthRefreshRequest(chatPostAuthRefreshRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuthRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuthRefresh`: ChatAuthResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuthRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuthRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuthRefreshRequest** | [**ChatPostAuthRefreshRequest**](ChatPostAuthRefreshRequest.md) |  | 

### Return type

[**ChatAuthResponse**](ChatAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAuthRegister

> map[string]interface{} ChatPostAuthRegister(ctx).ChatPostAuthRegisterRequest(chatPostAuthRegisterRequest).Execute()

Register a new user

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
	chatPostAuthRegisterRequest := *openapiclient.NewChatPostAuthRegisterRequest("Name_example", "Email_example", "Username_example", "Password_example", "ConfirmPassword_example") // ChatPostAuthRegisterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuthRegister(context.Background()).ChatPostAuthRegisterRequest(chatPostAuthRegisterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuthRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuthRegister`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuthRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuthRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuthRegisterRequest** | [**ChatPostAuthRegisterRequest**](ChatPostAuthRegisterRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAuthRequestpasswordreset

> map[string]interface{} ChatPostAuthRequestpasswordreset(ctx).ChatPostAuthRequestpasswordresetRequest(chatPostAuthRequestpasswordresetRequest).Execute()

Request password reset email

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
	chatPostAuthRequestpasswordresetRequest := *openapiclient.NewChatPostAuthRequestpasswordresetRequest("Email_example") // ChatPostAuthRequestpasswordresetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuthRequestpasswordreset(context.Background()).ChatPostAuthRequestpasswordresetRequest(chatPostAuthRequestpasswordresetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuthRequestpasswordreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuthRequestpasswordreset`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuthRequestpasswordreset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuthRequestpasswordresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuthRequestpasswordresetRequest** | [**ChatPostAuthRequestpasswordresetRequest**](ChatPostAuthRequestpasswordresetRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAuthResetpassword

> map[string]interface{} ChatPostAuthResetpassword(ctx).ChatPostAuthResetpasswordRequest(chatPostAuthResetpasswordRequest).Execute()

Reset password with token

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
	chatPostAuthResetpasswordRequest := *openapiclient.NewChatPostAuthResetpasswordRequest("Token_example", "Password_example", "ConfirmPassword_example", "UserId_example") // ChatPostAuthResetpasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAuthAPI.ChatPostAuthResetpassword(context.Background()).ChatPostAuthResetpasswordRequest(chatPostAuthResetpasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAuthAPI.ChatPostAuthResetpassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAuthResetpassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAuthAPI.ChatPostAuthResetpassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAuthResetpasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAuthResetpasswordRequest** | [**ChatPostAuthResetpasswordRequest**](ChatPostAuthResetpasswordRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

