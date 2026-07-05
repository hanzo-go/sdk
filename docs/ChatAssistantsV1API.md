# \ChatAssistantsV1API

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteAssistantsV1Byid**](ChatAssistantsV1API.md#ChatDeleteAssistantsV1Byid) | **Delete** /v1/chat/assistants/v1/{id} | Delete an assistant (v1)
[**ChatGetAssistantsV1**](ChatAssistantsV1API.md#ChatGetAssistantsV1) | **Get** /v1/chat/assistants/v1 | List assistants (v1)
[**ChatGetAssistantsV1Byid**](ChatAssistantsV1API.md#ChatGetAssistantsV1Byid) | **Get** /v1/chat/assistants/v1/{id} | Retrieve an assistant (v1)
[**ChatGetAssistantsV1Documents**](ChatAssistantsV1API.md#ChatGetAssistantsV1Documents) | **Get** /v1/chat/assistants/v1/documents | Get assistant documents
[**ChatGetAssistantsV1Tools**](ChatAssistantsV1API.md#ChatGetAssistantsV1Tools) | **Get** /v1/chat/assistants/v1/tools | List available assistant tools
[**ChatPatchAssistantsV1Byid**](ChatAssistantsV1API.md#ChatPatchAssistantsV1Byid) | **Patch** /v1/chat/assistants/v1/{id} | Modify an assistant (v1)
[**ChatPostAssistantsV1**](ChatAssistantsV1API.md#ChatPostAssistantsV1) | **Post** /v1/chat/assistants/v1 | Create an assistant (v1)
[**ChatPostAssistantsV1Chat**](ChatAssistantsV1API.md#ChatPostAssistantsV1Chat) | **Post** /v1/chat/assistants/v1/chat | Chat with an assistant (v1)
[**ChatPostAssistantsV1ChatAbort**](ChatAssistantsV1API.md#ChatPostAssistantsV1ChatAbort) | **Post** /v1/chat/assistants/v1/chat/abort | Abort assistant chat (v1)



## ChatDeleteAssistantsV1Byid

> map[string]interface{} ChatDeleteAssistantsV1Byid(ctx, id).Execute()

Delete an assistant (v1)

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatDeleteAssistantsV1Byid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatDeleteAssistantsV1Byid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteAssistantsV1Byid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatDeleteAssistantsV1Byid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteAssistantsV1ByidRequest struct via the builder pattern


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


## ChatGetAssistantsV1

> map[string]interface{} ChatGetAssistantsV1(ctx).Execute()

List assistants (v1)

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatGetAssistantsV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatGetAssistantsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAssistantsV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatGetAssistantsV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAssistantsV1Request struct via the builder pattern


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


## ChatGetAssistantsV1Byid

> map[string]interface{} ChatGetAssistantsV1Byid(ctx, id).Execute()

Retrieve an assistant (v1)

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatGetAssistantsV1Byid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatGetAssistantsV1Byid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAssistantsV1Byid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatGetAssistantsV1Byid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAssistantsV1ByidRequest struct via the builder pattern


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


## ChatGetAssistantsV1Documents

> map[string]interface{} ChatGetAssistantsV1Documents(ctx).Execute()

Get assistant documents

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatGetAssistantsV1Documents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatGetAssistantsV1Documents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAssistantsV1Documents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatGetAssistantsV1Documents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAssistantsV1DocumentsRequest struct via the builder pattern


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


## ChatGetAssistantsV1Tools

> map[string]interface{} ChatGetAssistantsV1Tools(ctx).Execute()

List available assistant tools

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatGetAssistantsV1Tools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatGetAssistantsV1Tools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAssistantsV1Tools`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatGetAssistantsV1Tools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAssistantsV1ToolsRequest struct via the builder pattern


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


## ChatPatchAssistantsV1Byid

> map[string]interface{} ChatPatchAssistantsV1Byid(ctx, id).Body(body).Execute()

Modify an assistant (v1)

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatPatchAssistantsV1Byid(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatPatchAssistantsV1Byid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchAssistantsV1Byid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatPatchAssistantsV1Byid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchAssistantsV1ByidRequest struct via the builder pattern


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


## ChatPostAssistantsV1

> map[string]interface{} ChatPostAssistantsV1(ctx).Body(body).Execute()

Create an assistant (v1)

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatPostAssistantsV1(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatPostAssistantsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatPostAssistantsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV1Request struct via the builder pattern


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


## ChatPostAssistantsV1Chat

> string ChatPostAssistantsV1Chat(ctx).Body(body).Execute()

Chat with an assistant (v1)

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatPostAssistantsV1Chat(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatPostAssistantsV1Chat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV1Chat`: string
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatPostAssistantsV1Chat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV1ChatRequest struct via the builder pattern


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


## ChatPostAssistantsV1ChatAbort

> map[string]interface{} ChatPostAssistantsV1ChatAbort(ctx).Execute()

Abort assistant chat (v1)

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
	resp, r, err := apiClient.ChatAssistantsV1API.ChatPostAssistantsV1ChatAbort(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAssistantsV1API.ChatPostAssistantsV1ChatAbort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV1ChatAbort`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAssistantsV1API.ChatPostAssistantsV1ChatAbort`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV1ChatAbortRequest struct via the builder pattern


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

