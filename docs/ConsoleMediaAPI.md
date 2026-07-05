# \ConsoleMediaAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleGetMedia**](ConsoleMediaAPI.md#ConsoleGetMedia) | **Get** /v1/console/media/{mediaId} | Get a media record
[**ConsoleGetMediaUploadUrl**](ConsoleMediaAPI.md#ConsoleGetMediaUploadUrl) | **Post** /v1/console/media | Get a presigned upload URL for a media record
[**ConsolePatchMedia**](ConsoleMediaAPI.md#ConsolePatchMedia) | **Patch** /v1/console/media/{mediaId} | Patch a media record (update upload status)



## ConsoleGetMedia

> ConsoleMediaRecord ConsoleGetMedia(ctx, mediaId).Execute()

Get a media record

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
	mediaId := "mediaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleMediaAPI.ConsoleGetMedia(context.Background(), mediaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleMediaAPI.ConsoleGetMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetMedia`: ConsoleMediaRecord
	fmt.Fprintf(os.Stdout, "Response from `ConsoleMediaAPI.ConsoleGetMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleMediaRecord**](ConsoleMediaRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetMediaUploadUrl

> ConsoleGetMediaUploadUrl200Response ConsoleGetMediaUploadUrl(ctx).ConsoleGetMediaUploadUrlRequest(consoleGetMediaUploadUrlRequest).Execute()

Get a presigned upload URL for a media record

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
	consoleGetMediaUploadUrlRequest := *openapiclient.NewConsoleGetMediaUploadUrlRequest("TraceId_example", "ContentType_example", int32(123), "Sha256Hash_example", "Field_example") // ConsoleGetMediaUploadUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleMediaAPI.ConsoleGetMediaUploadUrl(context.Background()).ConsoleGetMediaUploadUrlRequest(consoleGetMediaUploadUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleMediaAPI.ConsoleGetMediaUploadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetMediaUploadUrl`: ConsoleGetMediaUploadUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleMediaAPI.ConsoleGetMediaUploadUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetMediaUploadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleGetMediaUploadUrlRequest** | [**ConsoleGetMediaUploadUrlRequest**](ConsoleGetMediaUploadUrlRequest.md) |  | 

### Return type

[**ConsoleGetMediaUploadUrl200Response**](ConsoleGetMediaUploadUrl200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsolePatchMedia

> ConsolePatchMedia(ctx, mediaId).ConsolePatchMediaRequest(consolePatchMediaRequest).Execute()

Patch a media record (update upload status)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	mediaId := "mediaId_example" // string | 
	consolePatchMediaRequest := *openapiclient.NewConsolePatchMediaRequest(time.Now(), int32(123)) // ConsolePatchMediaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsoleMediaAPI.ConsolePatchMedia(context.Background(), mediaId).ConsolePatchMediaRequest(consolePatchMediaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleMediaAPI.ConsolePatchMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsolePatchMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consolePatchMediaRequest** | [**ConsolePatchMediaRequest**](ConsolePatchMediaRequest.md) |  | 

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

