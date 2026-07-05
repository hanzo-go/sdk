# \ChatAssistantsV2API

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteAssistantsV2Byid**](ChatAssistantsV2API.md#ChatDeleteAssistantsV2Byid) | **Delete** /v1/chat/assistants/v2/{id} | Delete an assistant (v2)
[**ChatGetAssistantsV2**](ChatAssistantsV2API.md#ChatGetAssistantsV2) | **Get** /v1/chat/assistants/v2 | List assistants (v2)
[**ChatGetAssistantsV2Byid**](ChatAssistantsV2API.md#ChatGetAssistantsV2Byid) | **Get** /v1/chat/assistants/v2/{id} | Retrieve an assistant (v2)
[**ChatGetAssistantsV2Documents**](ChatAssistantsV2API.md#ChatGetAssistantsV2Documents) | **Get** /v1/chat/assistants/v2/documents | Get assistant documents (v2)
[**ChatGetAssistantsV2Tools**](ChatAssistantsV2API.md#ChatGetAssistantsV2Tools) | **Get** /v1/chat/assistants/v2/tools | List available assistant tools (v2)
[**ChatPatchAssistantsV2Byid**](ChatAssistantsV2API.md#ChatPatchAssistantsV2Byid) | **Patch** /v1/chat/assistants/v2/{id} | Modify an assistant (v2)
[**ChatPostAssistantsV2**](ChatAssistantsV2API.md#ChatPostAssistantsV2) | **Post** /v1/chat/assistants/v2 | Create an assistant (v2)
[**ChatPostAssistantsV2Chat**](ChatAssistantsV2API.md#ChatPostAssistantsV2Chat) | **Post** /v1/chat/assistants/v2/chat | Chat with an assistant (v2)
[**ChatPostAssistantsV2ChatAbort**](ChatAssistantsV2API.md#ChatPostAssistantsV2ChatAbort) | **Post** /v1/chat/assistants/v2/chat/abort | Abort assistant chat (v2)



## ChatDeleteAssistantsV2Byid

> map[string]interface{} ChatDeleteAssistantsV2Byid(ctx, id).Execute()

Delete an assistant (v2)

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAssistantsV2API.ChatDeleteAssistantsV2Byid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatDeleteAssistantsV2Byid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteAssistantsV2Byid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatDeleteAssistantsV2Byid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteAssistantsV2ByidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ChatGetAssistantsV2

> map[string]interface{} ChatGetAssistantsV2(ctx).Execute()

List assistants (v2)

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
	resp, r, err := apiClient.ChatAssistantsV2API.ChatGetAssistantsV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatGetAssistantsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAssistantsV2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatGetAssistantsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAssistantsV2Request struct via the builder pattern


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


## ChatGetAssistantsV2Byid

> map[string]interface{} ChatGetAssistantsV2Byid(ctx, id).Execute()

Retrieve an assistant (v2)

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAssistantsV2API.ChatGetAssistantsV2Byid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatGetAssistantsV2Byid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAssistantsV2Byid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatGetAssistantsV2Byid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAssistantsV2ByidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ChatGetAssistantsV2Documents

> map[string]interface{} ChatGetAssistantsV2Documents(ctx).Execute()

Get assistant documents (v2)

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
	resp, r, err := apiClient.ChatAssistantsV2API.ChatGetAssistantsV2Documents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatGetAssistantsV2Documents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAssistantsV2Documents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatGetAssistantsV2Documents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAssistantsV2DocumentsRequest struct via the builder pattern


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


## ChatGetAssistantsV2Tools

> map[string]interface{} ChatGetAssistantsV2Tools(ctx).Execute()

List available assistant tools (v2)

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
	resp, r, err := apiClient.ChatAssistantsV2API.ChatGetAssistantsV2Tools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatGetAssistantsV2Tools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAssistantsV2Tools`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatGetAssistantsV2Tools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAssistantsV2ToolsRequest struct via the builder pattern


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


## ChatPatchAssistantsV2Byid

> map[string]interface{} ChatPatchAssistantsV2Byid(ctx, id).Body(body).Execute()

Modify an assistant (v2)

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
	id := "id_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAssistantsV2API.ChatPatchAssistantsV2Byid(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatPatchAssistantsV2Byid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchAssistantsV2Byid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatPatchAssistantsV2Byid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchAssistantsV2ByidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAssistantsV2

> map[string]interface{} ChatPostAssistantsV2(ctx).Body(body).Execute()

Create an assistant (v2)

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAssistantsV2API.ChatPostAssistantsV2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatPostAssistantsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatPostAssistantsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAssistantsV2Chat

> string ChatPostAssistantsV2Chat(ctx).Body(body).Execute()

Chat with an assistant (v2)

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAssistantsV2API.ChatPostAssistantsV2Chat(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatPostAssistantsV2Chat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV2Chat`: string
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatPostAssistantsV2Chat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV2ChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAssistantsV2ChatAbort

> map[string]interface{} ChatPostAssistantsV2ChatAbort(ctx).Execute()

Abort assistant chat (v2)

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
	resp, r, err := apiClient.ChatAssistantsV2API.ChatPostAssistantsV2ChatAbort(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV2API.ChatPostAssistantsV2ChatAbort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV2ChatAbort`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV2API.ChatPostAssistantsV2ChatAbort`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV2ChatAbortRequest struct via the builder pattern


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

