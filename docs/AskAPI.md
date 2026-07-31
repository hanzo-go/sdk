# \AskAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudPostV1Ask**](AskAPI.md#CloudPostV1Ask) | **Post** /v1/ask | 



## CloudPostV1Ask

> CloudPostV1Ask(ctx).CloudAskRequest(cloudAskRequest).Execute()



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
	cloudAskRequest := *openapiclient.NewCloudAskRequest() // CloudAskRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AskAPI.CloudPostV1Ask(context.Background()).CloudAskRequest(cloudAskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AskAPI.CloudPostV1Ask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAskRequest** | [**CloudAskRequest**](CloudAskRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

