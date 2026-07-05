# \AnalyticsAdminAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsAdminListUsers**](AnalyticsAdminAPI.md#AnalyticsAdminListUsers) | **Get** /v1/analytics/admin/users | List all users (admin only)
[**AnalyticsAdminListWebsites**](AnalyticsAdminAPI.md#AnalyticsAdminListWebsites) | **Get** /v1/analytics/admin/websites | List all websites for a user (admin only)



## AnalyticsAdminListUsers

> []AnalyticsAdminListUsers200ResponseInner AnalyticsAdminListUsers(ctx).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List all users (admin only)

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
	resp, r, err := apiClient.AnalyticsAdminAPI.AnalyticsAdminListUsers(context.Background()).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAdminAPI.AnalyticsAdminListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsAdminListUsers`: []AnalyticsAdminListUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAdminAPI.AnalyticsAdminListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsAdminListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsAdminListUsers200ResponseInner**](AnalyticsAdminListUsers200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsAdminListWebsites

> []AnalyticsWebsite AnalyticsAdminListWebsites(ctx).UserId(userId).IncludeOwnedTeams(includeOwnedTeams).IncludeAllTeams(includeAllTeams).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List all websites for a user (admin only)

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeOwnedTeams := "includeOwnedTeams_example" // string |  (optional)
	includeAllTeams := "includeAllTeams_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAdminAPI.AnalyticsAdminListWebsites(context.Background()).UserId(userId).IncludeOwnedTeams(includeOwnedTeams).IncludeAllTeams(includeAllTeams).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAdminAPI.AnalyticsAdminListWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsAdminListWebsites`: []AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAdminAPI.AnalyticsAdminListWebsites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsAdminListWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **includeOwnedTeams** | **string** |  | 
 **includeAllTeams** | **string** |  | 
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

