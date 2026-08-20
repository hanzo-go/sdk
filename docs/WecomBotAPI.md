# \WecomBotAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWecomBotCallbackByBotid**](WecomBotAPI.md#GetWecomBotCallbackByBotid) | **Get** /v1/wecom-bot/callback/{botId} | Verify WeChat work bot callback URL
[**PostWecomBotCallbackByBotid**](WecomBotAPI.md#PostWecomBotCallbackByBotid) | **Post** /v1/wecom-bot/callback/{botId} | Process WeChat work bot messages



## GetWecomBotCallbackByBotid

> GetWecomBotCallbackByBotid(ctx, botId).Execute()

Verify WeChat work bot callback URL



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
	botId := "botId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WecomBotAPI.GetWecomBotCallbackByBotid(context.Background(), botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WecomBotAPI.GetWecomBotCallbackByBotid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWecomBotCallbackByBotidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWecomBotCallbackByBotid

> PostWecomBotCallbackByBotid(ctx, botId).Execute()

Process WeChat work bot messages



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
	botId := "botId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WecomBotAPI.PostWecomBotCallbackByBotid(context.Background(), botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WecomBotAPI.PostWecomBotCallbackByBotid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWecomBotCallbackByBotidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

