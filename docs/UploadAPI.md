# \UploadAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotGetUploadUrl**](UploadAPI.md#BotGetUploadUrl) | **Post** /v1/bot/upload/url | Generate a presigned upload URL



## BotGetUploadUrl

> BotGetUploadUrl200Response BotGetUploadUrl(ctx).BotGetUploadUrlRequest(botGetUploadUrlRequest).Execute()

Generate a presigned upload URL

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
	botGetUploadUrlRequest := *openapiclient.NewBotGetUploadUrlRequest("Filename_example") // BotGetUploadUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadAPI.BotGetUploadUrl(context.Background()).BotGetUploadUrlRequest(botGetUploadUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadAPI.BotGetUploadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetUploadUrl`: BotGetUploadUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `UploadAPI.BotGetUploadUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotGetUploadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botGetUploadUrlRequest** | [**BotGetUploadUrlRequest**](BotGetUploadUrlRequest.md) |  | 

### Return type

[**BotGetUploadUrl200Response**](BotGetUploadUrl200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

