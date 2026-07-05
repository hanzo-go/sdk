# \AnalyticsMeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsChangeMyPassword**](AnalyticsMeAPI.md#AnalyticsChangeMyPassword) | **Post** /v1/analytics/me/password | Change current user password
[**AnalyticsGetMe**](AnalyticsMeAPI.md#AnalyticsGetMe) | **Get** /v1/analytics/me | Get current authenticated user info
[**AnalyticsGetMyTeams**](AnalyticsMeAPI.md#AnalyticsGetMyTeams) | **Get** /v1/analytics/me/teams | List teams for the current user
[**AnalyticsGetMyWebsites**](AnalyticsMeAPI.md#AnalyticsGetMyWebsites) | **Get** /v1/analytics/me/websites | List websites for the current user



## AnalyticsChangeMyPassword

> AnalyticsUser AnalyticsChangeMyPassword(ctx).AnalyticsChangeMyPasswordRequest(analyticsChangeMyPasswordRequest).Execute()

Change current user password

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
	analyticsChangeMyPasswordRequest := *openapiclient.NewAnalyticsChangeMyPasswordRequest("CurrentPassword_example", "NewPassword_example") // AnalyticsChangeMyPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsMeAPI.AnalyticsChangeMyPassword(context.Background()).AnalyticsChangeMyPasswordRequest(analyticsChangeMyPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsMeAPI.AnalyticsChangeMyPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsChangeMyPassword`: AnalyticsUser
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsMeAPI.AnalyticsChangeMyPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsChangeMyPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsChangeMyPasswordRequest** | [**AnalyticsChangeMyPasswordRequest**](AnalyticsChangeMyPasswordRequest.md) |  | 

### Return type

[**AnalyticsUser**](AnalyticsUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetMe

> AnalyticsGetMe200Response AnalyticsGetMe(ctx).Execute()

Get current authenticated user info

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
	resp, r, err := apiClient.AnalyticsMeAPI.AnalyticsGetMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsMeAPI.AnalyticsGetMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetMe`: AnalyticsGetMe200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsMeAPI.AnalyticsGetMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetMeRequest struct via the builder pattern


### Return type

[**AnalyticsGetMe200Response**](AnalyticsGetMe200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetMyTeams

> []AnalyticsTeam AnalyticsGetMyTeams(ctx).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List teams for the current user

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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsMeAPI.AnalyticsGetMyTeams(context.Background()).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsMeAPI.AnalyticsGetMyTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetMyTeams`: []AnalyticsTeam
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsMeAPI.AnalyticsGetMyTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetMyTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsTeam**](AnalyticsTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetMyWebsites

> []AnalyticsWebsite AnalyticsGetMyWebsites(ctx).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List websites for the current user

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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsMeAPI.AnalyticsGetMyWebsites(context.Background()).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsMeAPI.AnalyticsGetMyWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetMyWebsites`: []AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsMeAPI.AnalyticsGetMyWebsites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetMyWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

