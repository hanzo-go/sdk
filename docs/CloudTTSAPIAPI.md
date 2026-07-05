# \CloudTTSAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerGenerateTextToSpeechAudio**](CloudTTSAPIAPI.md#CloudApiControllerGenerateTextToSpeechAudio) | **Post** /v1/cloud/generate-text-to-speech-audio | Api Controller Generate Text To Speech Audio
[**CloudApiControllerGenerateTextToSpeechAudioStream**](CloudTTSAPIAPI.md#CloudApiControllerGenerateTextToSpeechAudioStream) | **Get** /v1/cloud/generate-text-to-speech-audio-stream | Api Controller Generate Text To Speech Audio Stream



## CloudApiControllerGenerateTextToSpeechAudio

> []string CloudApiControllerGenerateTextToSpeechAudio(ctx).Body(body).Execute()

Api Controller Generate Text To Speech Audio



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
	body := map[string]interface{}{ ... } // map[string]interface{} | The text to convert to speech

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudTTSAPIAPI.CloudApiControllerGenerateTextToSpeechAudio(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTTSAPIAPI.CloudApiControllerGenerateTextToSpeechAudio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGenerateTextToSpeechAudio`: []string
	fmt.Fprintf(os.Stdout, "Response from `CloudTTSAPIAPI.CloudApiControllerGenerateTextToSpeechAudio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGenerateTextToSpeechAudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The text to convert to speech | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGenerateTextToSpeechAudioStream

> map[string]interface{} CloudApiControllerGenerateTextToSpeechAudioStream(ctx).StoreId(storeId).MessageId(messageId).Execute()

Api Controller Generate Text To Speech Audio Stream



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
	storeId := "storeId_example" // string | The store ID
	messageId := "messageId_example" // string | The message ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudTTSAPIAPI.CloudApiControllerGenerateTextToSpeechAudioStream(context.Background()).StoreId(storeId).MessageId(messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTTSAPIAPI.CloudApiControllerGenerateTextToSpeechAudioStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGenerateTextToSpeechAudioStream`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudTTSAPIAPI.CloudApiControllerGenerateTextToSpeechAudioStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGenerateTextToSpeechAudioStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | The store ID | 
 **messageId** | **string** | The message ID | 

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

