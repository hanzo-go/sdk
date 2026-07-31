# \STTAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerProcessSpeechToText**](STTAPIAPI.md#CloudApiControllerProcessSpeechToText) | **Post** /v1/cloud/process-speech-to-text | Api Controller Process Speech To Text



## CloudApiControllerProcessSpeechToText

> map[string]interface{} CloudApiControllerProcessSpeechToText(ctx).Audio(audio).StoreId(storeId).Execute()

Api Controller Process Speech To Text



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
	resp, r, err := apiClient.STTAPIAPI.CloudApiControllerProcessSpeechToText(context.Background()).Audio(audio).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `STTAPIAPI.CloudApiControllerProcessSpeechToText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerProcessSpeechToText`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `STTAPIAPI.CloudApiControllerProcessSpeechToText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerProcessSpeechToTextRequest struct via the builder pattern


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

