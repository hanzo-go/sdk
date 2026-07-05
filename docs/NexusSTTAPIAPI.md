# \NexusSTTAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusProcessSpeechToText**](NexusSTTAPIAPI.md#NexusProcessSpeechToText) | **Post** /v1/nexus/process-speech-to-text | process Speech To Text



## NexusProcessSpeechToText

> map[string]interface{} NexusProcessSpeechToText(ctx).Audio(audio).StoreId(storeId).Execute()

process Speech To Text



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
	audio := os.NewFile(1234, "some_file") // *os.File | The audio file to convert to text
	storeId := "storeId_example" // string | The store ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusSTTAPIAPI.NexusProcessSpeechToText(context.Background()).Audio(audio).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusSTTAPIAPI.NexusProcessSpeechToText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusProcessSpeechToText`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusSTTAPIAPI.NexusProcessSpeechToText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusProcessSpeechToTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audio** | ***os.File** | The audio file to convert to text | 
 **storeId** | **string** | The store ID | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

