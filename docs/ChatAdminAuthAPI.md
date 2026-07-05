# \ChatAdminAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetAdminOauthOpenid**](ChatAdminAuthAPI.md#ChatGetAdminOauthOpenid) | **Get** /v1/chat/admin/oauth/openid | Initiate admin OpenID login
[**ChatGetAdminOauthOpenidCallback**](ChatAdminAuthAPI.md#ChatGetAdminOauthOpenidCallback) | **Get** /v1/chat/admin/oauth/openid/callback | Admin OpenID callback
[**ChatGetAdminOauthOpenidCheck**](ChatAdminAuthAPI.md#ChatGetAdminOauthOpenidCheck) | **Get** /v1/chat/admin/oauth/openid/check | Check if OpenID is configured for admin
[**ChatGetAdminVerify**](ChatAdminAuthAPI.md#ChatGetAdminVerify) | **Get** /v1/chat/admin/verify | Verify admin session
[**ChatPostAdminLoginLocal**](ChatAdminAuthAPI.md#ChatPostAdminLoginLocal) | **Post** /v1/chat/admin/login/local | Admin local login
[**ChatPostAdminOauthExchange**](ChatAdminAuthAPI.md#ChatPostAdminOauthExchange) | **Post** /v1/chat/admin/oauth/exchange | Exchange OAuth code for admin tokens



## ChatGetAdminOauthOpenid

> ChatGetAdminOauthOpenid(ctx).Execute()

Initiate admin OpenID login

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
	r, err := apiClient.ChatAdminAuthAPI.ChatGetAdminOauthOpenid(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAdminAuthAPI.ChatGetAdminOauthOpenid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAdminOauthOpenidRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAdminOauthOpenidCallback

> ChatGetAdminOauthOpenidCallback(ctx).Execute()

Admin OpenID callback

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
	r, err := apiClient.ChatAdminAuthAPI.ChatGetAdminOauthOpenidCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAdminAuthAPI.ChatGetAdminOauthOpenidCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAdminOauthOpenidCallbackRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAdminOauthOpenidCheck

> map[string]interface{} ChatGetAdminOauthOpenidCheck(ctx).Execute()

Check if OpenID is configured for admin

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
	resp, r, err := apiClient.ChatAdminAuthAPI.ChatGetAdminOauthOpenidCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAdminAuthAPI.ChatGetAdminOauthOpenidCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAdminOauthOpenidCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAdminAuthAPI.ChatGetAdminOauthOpenidCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAdminOauthOpenidCheckRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAdminVerify

> ChatGetAdminVerify200Response ChatGetAdminVerify(ctx).Execute()

Verify admin session

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
	resp, r, err := apiClient.ChatAdminAuthAPI.ChatGetAdminVerify(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAdminAuthAPI.ChatGetAdminVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAdminVerify`: ChatGetAdminVerify200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAdminAuthAPI.ChatGetAdminVerify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAdminVerifyRequest struct via the builder pattern


### Return type

[**ChatGetAdminVerify200Response**](ChatGetAdminVerify200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAdminLoginLocal

> map[string]interface{} ChatPostAdminLoginLocal(ctx).ChatPostAdminLoginLocalRequest(chatPostAdminLoginLocalRequest).Execute()

Admin local login

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
	chatPostAdminLoginLocalRequest := *openapiclient.NewChatPostAdminLoginLocalRequest("Email_example", "Password_example") // ChatPostAdminLoginLocalRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAdminAuthAPI.ChatPostAdminLoginLocal(context.Background()).ChatPostAdminLoginLocalRequest(chatPostAdminLoginLocalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAdminAuthAPI.ChatPostAdminLoginLocal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAdminLoginLocal`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAdminAuthAPI.ChatPostAdminLoginLocal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAdminLoginLocalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAdminLoginLocalRequest** | [**ChatPostAdminLoginLocalRequest**](ChatPostAdminLoginLocalRequest.md) |  | 

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


## ChatPostAdminOauthExchange

> ChatPostAdminOauthExchange200Response ChatPostAdminOauthExchange(ctx).ChatPostAdminOauthExchangeRequest(chatPostAdminOauthExchangeRequest).Execute()

Exchange OAuth code for admin tokens

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
	chatPostAdminOauthExchangeRequest := *openapiclient.NewChatPostAdminOauthExchangeRequest("Code_example") // ChatPostAdminOauthExchangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAdminAuthAPI.ChatPostAdminOauthExchange(context.Background()).ChatPostAdminOauthExchangeRequest(chatPostAdminOauthExchangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAdminAuthAPI.ChatPostAdminOauthExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAdminOauthExchange`: ChatPostAdminOauthExchange200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAdminAuthAPI.ChatPostAdminOauthExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAdminOauthExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAdminOauthExchangeRequest** | [**ChatPostAdminOauthExchangeRequest**](ChatPostAdminOauthExchangeRequest.md) |  | 

### Return type

[**ChatPostAdminOauthExchange200Response**](ChatPostAdminOauthExchange200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

