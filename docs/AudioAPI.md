# \AudioAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAudioFoley**](AudioAPI.md#PostAudioFoley) | **Post** /v1/audio/foley | Serves the generative audio verbs — /v1/audio/voice (TTS), /music, /foley — that the Zen family serves natively.
[**PostAudioMusic**](AudioAPI.md#PostAudioMusic) | **Post** /v1/audio/music | Serves the generative audio verbs — /v1/audio/voice (TTS), /music, /foley — that the Zen family serves natively.
[**PostAudioSpeech**](AudioAPI.md#PostAudioSpeech) | **Post** /v1/audio/speech | The OpenAI-compatible TTS endpoint (POST /v1/audio/speech).
[**PostAudioTranscriptions**](AudioAPI.md#PostAudioTranscriptions) | **Post** /v1/audio/transcriptions | The OpenAI-compatible STT endpoint (POST /v1/audio/transcriptions, multipart: file + model [+ language + response_format]).
[**PostAudioVoice**](AudioAPI.md#PostAudioVoice) | **Post** /v1/audio/voice | Serves the generative audio verbs — /v1/audio/voice (TTS), /music, /foley — that the Zen family serves natively.



## PostAudioFoley

> PostAudioFoley(ctx).Execute()

Serves the generative audio verbs — /v1/audio/voice (TTS), /music, /foley — that the Zen family serves natively.



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
	r, err := apiClient.AudioAPI.PostAudioFoley(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.PostAudioFoley``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAudioFoleyRequest struct via the builder pattern


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


## PostAudioMusic

> PostAudioMusic(ctx).Execute()

Serves the generative audio verbs — /v1/audio/voice (TTS), /music, /foley — that the Zen family serves natively.



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
	r, err := apiClient.AudioAPI.PostAudioMusic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.PostAudioMusic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAudioMusicRequest struct via the builder pattern


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


## PostAudioSpeech

> PostAudioSpeech(ctx).Execute()

The OpenAI-compatible TTS endpoint (POST /v1/audio/speech).



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
	r, err := apiClient.AudioAPI.PostAudioSpeech(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.PostAudioSpeech``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAudioSpeechRequest struct via the builder pattern


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


## PostAudioTranscriptions

> PostAudioTranscriptions(ctx).Execute()

The OpenAI-compatible STT endpoint (POST /v1/audio/transcriptions, multipart: file + model [+ language + response_format]).



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
	r, err := apiClient.AudioAPI.PostAudioTranscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.PostAudioTranscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAudioTranscriptionsRequest struct via the builder pattern


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


## PostAudioVoice

> PostAudioVoice(ctx).Execute()

Serves the generative audio verbs — /v1/audio/voice (TTS), /music, /foley — that the Zen family serves natively.



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
	r, err := apiClient.AudioAPI.PostAudioVoice(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.PostAudioVoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAudioVoiceRequest struct via the builder pattern


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

