# \GitWebhookAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGitWebhook**](GitWebhookAPI.md#PostGitWebhook) | **Post** /v1/git-webhook | Receive a push from the forge and trigger its build



## PostGitWebhook

> Verdict PostGitWebhook(ctx).Push(push).Execute()

Receive a push from the forge and trigger its build



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
	push := *openapiclient.NewPush() // Push |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitWebhookAPI.PostGitWebhook(context.Background()).Push(push).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitWebhookAPI.PostGitWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitWebhook`: Verdict
	fmt.Fprintf(os.Stdout, "Response from `GitWebhookAPI.PostGitWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **push** | [**Push**](Push.md) |  | 

### Return type

[**Verdict**](Verdict.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

