# \BotUsersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotGetUser**](BotUsersAPI.md#BotGetUser) | **Get** /v1/bot/users/{handle} | Get user profile by handle
[**BotListUserSkills**](BotUsersAPI.md#BotListUserSkills) | **Get** /v1/bot/users/{handle}/skills | List skills published by a user
[**BotListUserStars**](BotUsersAPI.md#BotListUserStars) | **Get** /v1/bot/users/{handle}/stars | List skills starred by a user
[**BotUpdateProfile**](BotUsersAPI.md#BotUpdateProfile) | **Patch** /v1/bot/users/me | Update current user&#39;s profile



## BotGetUser

> BotUser BotGetUser(ctx, handle).Execute()

Get user profile by handle

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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotUsersAPI.BotGetUser(context.Background(), handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotUsersAPI.BotGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetUser`: BotUser
	fmt.Fprintf(os.Stdout, "Response from `BotUsersAPI.BotGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotUser**](BotUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListUserSkills

> BotListUserSkills200Response BotListUserSkills(ctx, handle).Execute()

List skills published by a user

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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotUsersAPI.BotListUserSkills(context.Background(), handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotUsersAPI.BotListUserSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListUserSkills`: BotListUserSkills200Response
	fmt.Fprintf(os.Stdout, "Response from `BotUsersAPI.BotListUserSkills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotListUserSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotListUserSkills200Response**](BotListUserSkills200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListUserStars

> BotListUserStars200Response BotListUserStars(ctx, handle).Execute()

List skills starred by a user

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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotUsersAPI.BotListUserStars(context.Background(), handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotUsersAPI.BotListUserStars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListUserStars`: BotListUserStars200Response
	fmt.Fprintf(os.Stdout, "Response from `BotUsersAPI.BotListUserStars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotListUserStarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotListUserStars200Response**](BotListUserStars200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotUpdateProfile

> AnalyticsHeartbeat200Response BotUpdateProfile(ctx).BotUpdateProfileRequest(botUpdateProfileRequest).Execute()

Update current user's profile

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
	botUpdateProfileRequest := *openapiclient.NewBotUpdateProfileRequest() // BotUpdateProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotUsersAPI.BotUpdateProfile(context.Background()).BotUpdateProfileRequest(botUpdateProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotUsersAPI.BotUpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotUpdateProfile`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `BotUsersAPI.BotUpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botUpdateProfileRequest** | [**BotUpdateProfileRequest**](BotUpdateProfileRequest.md) |  | 

### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

