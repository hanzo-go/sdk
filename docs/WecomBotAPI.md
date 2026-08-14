# \WecomBotAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWecomBotCallbackByBotid**](WecomBotAPI.md#GetWecomBotCallbackByBotid) | **Get** /v1/wecom-bot/callback/{botId} | Verify WeChat work bot callback URL
[**PostWecomBotCallbackByBotid**](WecomBotAPI.md#PostWecomBotCallbackByBotid) | **Post** /v1/wecom-bot/callback/{botId} | Process WeChat work bot messages



## GetWecomBotCallbackByBotid

> GetWecomBotCallbackByBotid(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WecomBotAPI.GetWecomBotCallbackByBotid(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WecomBotAPI.GetWecomBotCallbackByBotid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWecomBotCallbackByBotidRequest struct via the builder pattern


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


## PostWecomBotCallbackByBotid

> PostWecomBotCallbackByBotid(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WecomBotAPI.PostWecomBotCallbackByBotid(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WecomBotAPI.PostWecomBotCallbackByBotid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostWecomBotCallbackByBotidRequest struct via the builder pattern


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

