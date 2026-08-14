# \EmbedAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmbed**](EmbedAPI.md#GetEmbed) | **Get** /v1/embed | Reports whether one of this brand&#39;s shared embedded apps (cms, erp, help) may be framed by the caller and is actually running, so a console module can choose between the embed and the provision panel.



## GetEmbed

> EmbedStatusResp GetEmbed(ctx).App(app).Execute()

Reports whether one of this brand's shared embedded apps (cms, erp, help) may be framed by the caller and is actually running, so a console module can choose between the embed and the provision panel.



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
	app := "cms" // string | App is the embedded app to report on: cms (Content Studio), erp or help. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbedAPI.GetEmbed(context.Background()).App(app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.GetEmbed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmbed`: EmbedStatusResp
	fmt.Fprintf(os.Stdout, "Response from `EmbedAPI.GetEmbed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmbedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **string** | App is the embedded app to report on: cms (Content Studio), erp or help. | 

### Return type

[**EmbedStatusResp**](EmbedStatusResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

