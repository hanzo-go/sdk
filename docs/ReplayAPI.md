# \ReplayAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostReplay**](ReplayAPI.md#PostReplay) | **Post** /v1/replay | Record a session-replay snapshot batch



## PostReplay

> CaptureResult PostReplay(ctx).ReplayBody(replayBody).Execute()

Record a session-replay snapshot batch



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
	replayBody := *openapiclient.NewReplayBody() // ReplayBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplayAPI.PostReplay(context.Background()).ReplayBody(replayBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplayAPI.PostReplay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReplay`: CaptureResult
	fmt.Fprintf(os.Stdout, "Response from `ReplayAPI.PostReplay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replayBody** | [**ReplayBody**](ReplayBody.md) |  | 

### Return type

[**CaptureResult**](CaptureResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

