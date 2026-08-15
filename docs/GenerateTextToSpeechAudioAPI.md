# \GenerateTextToSpeechAudioAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGenerateTextToSpeechAudio**](GenerateTextToSpeechAudioAPI.md#PostGenerateTextToSpeechAudio) | **Post** /v1/generate-text-to-speech-audio | Convert text to speech



## PostGenerateTextToSpeechAudio

> PostGenerateTextToSpeechAudio(ctx).Execute()

Convert text to speech



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
	r, err := apiClient.GenerateTextToSpeechAudioAPI.PostGenerateTextToSpeechAudio(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenerateTextToSpeechAudioAPI.PostGenerateTextToSpeechAudio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostGenerateTextToSpeechAudioRequest struct via the builder pattern


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

