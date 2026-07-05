# \SearchChatsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchListChatWorkspaces**](SearchChatsAPI.md#SearchListChatWorkspaces) | **Get** /v1/search/chats | List chat workspaces (experimental)



## SearchListChatWorkspaces

> map[string]interface{} SearchListChatWorkspaces(ctx).Execute()

List chat workspaces (experimental)

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
	resp, r, err := apiClient.SearchChatsAPI.SearchListChatWorkspaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchChatsAPI.SearchListChatWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchListChatWorkspaces`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SearchChatsAPI.SearchListChatWorkspaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchListChatWorkspacesRequest struct via the builder pattern


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

