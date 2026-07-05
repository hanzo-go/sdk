# \ChatSettingsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetUserSettingsFavorites**](ChatSettingsAPI.md#ChatGetUserSettingsFavorites) | **Get** /v1/chat/user/settings/favorites | Get user favorites
[**ChatPostUserSettingsFavorites**](ChatSettingsAPI.md#ChatPostUserSettingsFavorites) | **Post** /v1/chat/user/settings/favorites | Update user favorites



## ChatGetUserSettingsFavorites

> map[string]interface{} ChatGetUserSettingsFavorites(ctx).Execute()

Get user favorites

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
	resp, r, err := apiClient.ChatSettingsAPI.ChatGetUserSettingsFavorites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatSettingsAPI.ChatGetUserSettingsFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetUserSettingsFavorites`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatSettingsAPI.ChatGetUserSettingsFavorites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetUserSettingsFavoritesRequest struct via the builder pattern


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


## ChatPostUserSettingsFavorites

> map[string]interface{} ChatPostUserSettingsFavorites(ctx).Body(body).Execute()

Update user favorites

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatSettingsAPI.ChatPostUserSettingsFavorites(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatSettingsAPI.ChatPostUserSettingsFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostUserSettingsFavorites`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatSettingsAPI.ChatPostUserSettingsFavorites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostUserSettingsFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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

