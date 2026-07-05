# \ChatUserAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteUserDelete**](ChatUserAPI.md#ChatDeleteUserDelete) | **Delete** /v1/chat/user/delete | Delete user account
[**ChatGetUser**](ChatUserAPI.md#ChatGetUser) | **Get** /v1/chat/user | Get current user
[**ChatGetUserTerms**](ChatUserAPI.md#ChatGetUserTerms) | **Get** /v1/chat/user/terms | Get terms acceptance status
[**ChatPostUserPlugins**](ChatUserAPI.md#ChatPostUserPlugins) | **Post** /v1/chat/user/plugins | Update user plugins
[**ChatPostUserTermsAccept**](ChatUserAPI.md#ChatPostUserTermsAccept) | **Post** /v1/chat/user/terms/accept | Accept terms of service
[**ChatPostUserVerify**](ChatUserAPI.md#ChatPostUserVerify) | **Post** /v1/chat/user/verify | Verify email with token
[**ChatPostUserVerifyResend**](ChatUserAPI.md#ChatPostUserVerifyResend) | **Post** /v1/chat/user/verify/resend | Resend verification email



## ChatDeleteUserDelete

> map[string]interface{} ChatDeleteUserDelete(ctx).Execute()

Delete user account

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
	resp, r, err := apiClient.ChatUserAPI.ChatDeleteUserDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatUserAPI.ChatDeleteUserDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteUserDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatUserAPI.ChatDeleteUserDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteUserDeleteRequest struct via the builder pattern


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


## ChatGetUser

> ChatUser ChatGetUser(ctx).Execute()

Get current user

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
	resp, r, err := apiClient.ChatUserAPI.ChatGetUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatUserAPI.ChatGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetUser`: ChatUser
	fmt.Fprintf(os.Stdout, "Response from `ChatUserAPI.ChatGetUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetUserRequest struct via the builder pattern


### Return type

[**ChatUser**](ChatUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetUserTerms

> map[string]interface{} ChatGetUserTerms(ctx).Execute()

Get terms acceptance status

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
	resp, r, err := apiClient.ChatUserAPI.ChatGetUserTerms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatUserAPI.ChatGetUserTerms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetUserTerms`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatUserAPI.ChatGetUserTerms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetUserTermsRequest struct via the builder pattern


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


## ChatPostUserPlugins

> map[string]interface{} ChatPostUserPlugins(ctx).ChatPostUserPluginsRequest(chatPostUserPluginsRequest).Execute()

Update user plugins

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
	chatPostUserPluginsRequest := *openapiclient.NewChatPostUserPluginsRequest() // ChatPostUserPluginsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatUserAPI.ChatPostUserPlugins(context.Background()).ChatPostUserPluginsRequest(chatPostUserPluginsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatUserAPI.ChatPostUserPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostUserPlugins`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatUserAPI.ChatPostUserPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostUserPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostUserPluginsRequest** | [**ChatPostUserPluginsRequest**](ChatPostUserPluginsRequest.md) |  | 

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


## ChatPostUserTermsAccept

> map[string]interface{} ChatPostUserTermsAccept(ctx).Execute()

Accept terms of service

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
	resp, r, err := apiClient.ChatUserAPI.ChatPostUserTermsAccept(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatUserAPI.ChatPostUserTermsAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostUserTermsAccept`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatUserAPI.ChatPostUserTermsAccept`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostUserTermsAcceptRequest struct via the builder pattern


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


## ChatPostUserVerify

> map[string]interface{} ChatPostUserVerify(ctx).ChatPostUserVerifyRequest(chatPostUserVerifyRequest).Execute()

Verify email with token

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
	chatPostUserVerifyRequest := *openapiclient.NewChatPostUserVerifyRequest() // ChatPostUserVerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatUserAPI.ChatPostUserVerify(context.Background()).ChatPostUserVerifyRequest(chatPostUserVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatUserAPI.ChatPostUserVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostUserVerify`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatUserAPI.ChatPostUserVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostUserVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostUserVerifyRequest** | [**ChatPostUserVerifyRequest**](ChatPostUserVerifyRequest.md) |  | 

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


## ChatPostUserVerifyResend

> map[string]interface{} ChatPostUserVerifyResend(ctx).ChatPostUserVerifyResendRequest(chatPostUserVerifyResendRequest).Execute()

Resend verification email

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
	chatPostUserVerifyResendRequest := *openapiclient.NewChatPostUserVerifyResendRequest() // ChatPostUserVerifyResendRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatUserAPI.ChatPostUserVerifyResend(context.Background()).ChatPostUserVerifyResendRequest(chatPostUserVerifyResendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatUserAPI.ChatPostUserVerifyResend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostUserVerifyResend`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatUserAPI.ChatPostUserVerifyResend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostUserVerifyResendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostUserVerifyResendRequest** | [**ChatPostUserVerifyResendRequest**](ChatPostUserVerifyResendRequest.md) |  | 

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

