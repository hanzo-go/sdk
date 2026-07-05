# \ChatPresetsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetPresets**](ChatPresetsAPI.md#ChatGetPresets) | **Get** /v1/chat/presets | List user presets
[**ChatPostPresets**](ChatPresetsAPI.md#ChatPostPresets) | **Post** /v1/chat/presets | Create or update a preset
[**ChatPostPresetsDelete**](ChatPresetsAPI.md#ChatPostPresetsDelete) | **Post** /v1/chat/presets/delete | Delete a preset



## ChatGetPresets

> []ChatPreset ChatGetPresets(ctx).Execute()

List user presets

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
	resp, r, err := apiClient.ChatPresetsAPI.ChatGetPresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPresetsAPI.ChatGetPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPresets`: []ChatPreset
	fmt.Fprintf(os.Stdout, "Response from `ChatPresetsAPI.ChatGetPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPresetsRequest struct via the builder pattern


### Return type

[**[]ChatPreset**](ChatPreset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostPresets

> ChatPreset ChatPostPresets(ctx).ChatPreset(chatPreset).Execute()

Create or update a preset

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
	chatPreset := *openapiclient.NewChatPreset() // ChatPreset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPresetsAPI.ChatPostPresets(context.Background()).ChatPreset(chatPreset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPresetsAPI.ChatPostPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostPresets`: ChatPreset
	fmt.Fprintf(os.Stdout, "Response from `ChatPresetsAPI.ChatPostPresets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostPresetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPreset** | [**ChatPreset**](ChatPreset.md) |  | 

### Return type

[**ChatPreset**](ChatPreset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostPresetsDelete

> map[string]interface{} ChatPostPresetsDelete(ctx).ChatPostPresetsDeleteRequest(chatPostPresetsDeleteRequest).Execute()

Delete a preset

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
	chatPostPresetsDeleteRequest := *openapiclient.NewChatPostPresetsDeleteRequest() // ChatPostPresetsDeleteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPresetsAPI.ChatPostPresetsDelete(context.Background()).ChatPostPresetsDeleteRequest(chatPostPresetsDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPresetsAPI.ChatPostPresetsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostPresetsDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPresetsAPI.ChatPostPresetsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostPresetsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostPresetsDeleteRequest** | [**ChatPostPresetsDeleteRequest**](ChatPostPresetsDeleteRequest.md) |  | 

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

