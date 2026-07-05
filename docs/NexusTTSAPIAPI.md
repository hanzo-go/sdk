# \NexusTTSAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusGenerateTextToSpeechAudio**](NexusTTSAPIAPI.md#NexusGenerateTextToSpeechAudio) | **Post** /v1/nexus/generate-text-to-speech-audio | generate Text To Speech Audio
[**NexusGenerateTextToSpeechAudioStream**](NexusTTSAPIAPI.md#NexusGenerateTextToSpeechAudioStream) | **Get** /v1/nexus/generate-text-to-speech-audio-stream | generate Text To Speech Audio Stream



## NexusGenerateTextToSpeechAudio

> *os.File NexusGenerateTextToSpeechAudio(ctx).Body(body).Execute()

generate Text To Speech Audio



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
	resp, r, err := apiClient.NexusTTSAPIAPI.NexusGenerateTextToSpeechAudio(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTTSAPIAPI.NexusGenerateTextToSpeechAudio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGenerateTextToSpeechAudio`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `NexusTTSAPIAPI.NexusGenerateTextToSpeechAudio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGenerateTextToSpeechAudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The text to convert to speech | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGenerateTextToSpeechAudioStream

> string NexusGenerateTextToSpeechAudioStream(ctx).StoreId(storeId).MessageId(messageId).Execute()

generate Text To Speech Audio Stream



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
	resp, r, err := apiClient.NexusTTSAPIAPI.NexusGenerateTextToSpeechAudioStream(context.Background()).StoreId(storeId).MessageId(messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTTSAPIAPI.NexusGenerateTextToSpeechAudioStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGenerateTextToSpeechAudioStream`: string
	fmt.Fprintf(os.Stdout, "Response from `NexusTTSAPIAPI.NexusGenerateTextToSpeechAudioStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGenerateTextToSpeechAudioStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | The store ID | 
 **messageId** | **string** | The message ID | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

