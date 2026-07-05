# \ChatSpeechAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetFilesSpeechConfigGet**](ChatSpeechAPI.md#ChatGetFilesSpeechConfigGet) | **Get** /v1/chat/files/speech/config/get | Get custom speech configuration
[**ChatGetFilesSpeechTtsVoices**](ChatSpeechAPI.md#ChatGetFilesSpeechTtsVoices) | **Get** /v1/chat/files/speech/tts/voices | Get available TTS voices
[**ChatPostFilesSpeechStt**](ChatSpeechAPI.md#ChatPostFilesSpeechStt) | **Post** /v1/chat/files/speech/stt | Speech to text
[**ChatPostFilesSpeechTts**](ChatSpeechAPI.md#ChatPostFilesSpeechTts) | **Post** /v1/chat/files/speech/tts | Stream text to speech
[**ChatPostFilesSpeechTtsManual**](ChatSpeechAPI.md#ChatPostFilesSpeechTtsManual) | **Post** /v1/chat/files/speech/tts/manual | Manual text to speech



## ChatGetFilesSpeechConfigGet

> map[string]interface{} ChatGetFilesSpeechConfigGet(ctx).Execute()

Get custom speech configuration

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
	resp, r, err := apiClient.ChatSpeechAPI.ChatGetFilesSpeechConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatSpeechAPI.ChatGetFilesSpeechConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetFilesSpeechConfigGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatSpeechAPI.ChatGetFilesSpeechConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetFilesSpeechConfigGetRequest struct via the builder pattern


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


## ChatGetFilesSpeechTtsVoices

> map[string]interface{} ChatGetFilesSpeechTtsVoices(ctx).Execute()

Get available TTS voices

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
	resp, r, err := apiClient.ChatSpeechAPI.ChatGetFilesSpeechTtsVoices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatSpeechAPI.ChatGetFilesSpeechTtsVoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetFilesSpeechTtsVoices`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatSpeechAPI.ChatGetFilesSpeechTtsVoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetFilesSpeechTtsVoicesRequest struct via the builder pattern


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


## ChatPostFilesSpeechStt

> map[string]interface{} ChatPostFilesSpeechStt(ctx).Audio(audio).Execute()

Speech to text

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
	audio := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatSpeechAPI.ChatPostFilesSpeechStt(context.Background()).Audio(audio).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatSpeechAPI.ChatPostFilesSpeechStt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostFilesSpeechStt`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatSpeechAPI.ChatPostFilesSpeechStt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostFilesSpeechSttRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audio** | ***os.File** |  | 

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


## ChatPostFilesSpeechTts

> *os.File ChatPostFilesSpeechTts(ctx).ChatPostFilesSpeechTtsRequest(chatPostFilesSpeechTtsRequest).Execute()

Stream text to speech

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
	chatPostFilesSpeechTtsRequest := *openapiclient.NewChatPostFilesSpeechTtsRequest() // ChatPostFilesSpeechTtsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatSpeechAPI.ChatPostFilesSpeechTts(context.Background()).ChatPostFilesSpeechTtsRequest(chatPostFilesSpeechTtsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatSpeechAPI.ChatPostFilesSpeechTts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostFilesSpeechTts`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ChatSpeechAPI.ChatPostFilesSpeechTts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostFilesSpeechTtsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostFilesSpeechTtsRequest** | [**ChatPostFilesSpeechTtsRequest**](ChatPostFilesSpeechTtsRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: audio/mpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostFilesSpeechTtsManual

> map[string]interface{} ChatPostFilesSpeechTtsManual(ctx).Text(text).Voice(voice).Execute()

Manual text to speech

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
	text := "text_example" // string |  (optional)
	voice := "voice_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatSpeechAPI.ChatPostFilesSpeechTtsManual(context.Background()).Text(text).Voice(voice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatSpeechAPI.ChatPostFilesSpeechTtsManual``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostFilesSpeechTtsManual`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatSpeechAPI.ChatPostFilesSpeechTtsManual`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostFilesSpeechTtsManualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** |  | 
 **voice** | **string** |  | 

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

