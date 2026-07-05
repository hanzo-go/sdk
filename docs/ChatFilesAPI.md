# \ChatFilesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteFiles**](ChatFilesAPI.md#ChatDeleteFiles) | **Delete** /v1/chat/files | Delete files
[**ChatGetFiles**](ChatFilesAPI.md#ChatGetFiles) | **Get** /v1/chat/files | List user files
[**ChatGetFilesAgentByagentId**](ChatFilesAPI.md#ChatGetFilesAgentByagentId) | **Get** /v1/chat/files/agent/{agent_id} | Get files for an agent
[**ChatGetFilesCodeDownloadBysessionIdByfileid**](ChatFilesAPI.md#ChatGetFilesCodeDownloadBysessionIdByfileid) | **Get** /v1/chat/files/code/download/{session_id}/{fileId} | Download code execution output
[**ChatGetFilesConfig**](ChatFilesAPI.md#ChatGetFilesConfig) | **Get** /v1/chat/files/config | Get file upload configuration
[**ChatGetFilesDownloadByuseridByfileId**](ChatFilesAPI.md#ChatGetFilesDownloadByuseridByfileId) | **Get** /v1/chat/files/download/{userId}/{file_id} | Download a file
[**ChatPostFiles**](ChatFilesAPI.md#ChatPostFiles) | **Post** /v1/chat/files | Upload a file



## ChatDeleteFiles

> map[string]interface{} ChatDeleteFiles(ctx).ChatDeleteFilesRequest(chatDeleteFilesRequest).Execute()

Delete files

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
	chatDeleteFilesRequest := *openapiclient.NewChatDeleteFilesRequest([]openapiclient.ChatDeleteFilesRequestFilesInner{*openapiclient.NewChatDeleteFilesRequestFilesInner()}) // ChatDeleteFilesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatFilesAPI.ChatDeleteFiles(context.Background()).ChatDeleteFilesRequest(chatDeleteFilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatFilesAPI.ChatDeleteFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteFiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatFilesAPI.ChatDeleteFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatDeleteFilesRequest** | [**ChatDeleteFilesRequest**](ChatDeleteFilesRequest.md) |  | 

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


## ChatGetFiles

> []ChatFile ChatGetFiles(ctx).Execute()

List user files

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
	resp, r, err := apiClient.ChatFilesAPI.ChatGetFiles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatFilesAPI.ChatGetFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetFiles`: []ChatFile
	fmt.Fprintf(os.Stdout, "Response from `ChatFilesAPI.ChatGetFiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetFilesRequest struct via the builder pattern


### Return type

[**[]ChatFile**](ChatFile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetFilesAgentByagentId

> []ChatFile ChatGetFilesAgentByagentId(ctx, agentId).Execute()

Get files for an agent

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
	agentId := "agentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatFilesAPI.ChatGetFilesAgentByagentId(context.Background(), agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatFilesAPI.ChatGetFilesAgentByagentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetFilesAgentByagentId`: []ChatFile
	fmt.Fprintf(os.Stdout, "Response from `ChatFilesAPI.ChatGetFilesAgentByagentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetFilesAgentByagentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChatFile**](ChatFile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetFilesCodeDownloadBysessionIdByfileid

> *os.File ChatGetFilesCodeDownloadBysessionIdByfileid(ctx, sessionId, fileId).Execute()

Download code execution output

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
	sessionId := "sessionId_example" // string | 
	fileId := "fileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatFilesAPI.ChatGetFilesCodeDownloadBysessionIdByfileid(context.Background(), sessionId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatFilesAPI.ChatGetFilesCodeDownloadBysessionIdByfileid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetFilesCodeDownloadBysessionIdByfileid`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ChatFilesAPI.ChatGetFilesCodeDownloadBysessionIdByfileid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetFilesCodeDownloadBysessionIdByfileidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetFilesConfig

> map[string]interface{} ChatGetFilesConfig(ctx).Execute()

Get file upload configuration

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
	resp, r, err := apiClient.ChatFilesAPI.ChatGetFilesConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatFilesAPI.ChatGetFilesConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetFilesConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatFilesAPI.ChatGetFilesConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetFilesConfigRequest struct via the builder pattern


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


## ChatGetFilesDownloadByuseridByfileId

> *os.File ChatGetFilesDownloadByuseridByfileId(ctx, userId, fileId).Execute()

Download a file

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
	userId := "userId_example" // string | 
	fileId := "fileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatFilesAPI.ChatGetFilesDownloadByuseridByfileId(context.Background(), userId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatFilesAPI.ChatGetFilesDownloadByuseridByfileId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetFilesDownloadByuseridByfileId`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ChatFilesAPI.ChatGetFilesDownloadByuseridByfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetFilesDownloadByuseridByfileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostFiles

> ChatFile ChatPostFiles(ctx).File(file).FileId(fileId).Endpoint(endpoint).AgentId(agentId).ToolResource(toolResource).MessageFile(messageFile).Execute()

Upload a file

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
	file := os.NewFile(1234, "some_file") // *os.File | 
	fileId := "fileId_example" // string |  (optional)
	endpoint := "endpoint_example" // string |  (optional)
	agentId := "agentId_example" // string |  (optional)
	toolResource := "toolResource_example" // string |  (optional)
	messageFile := "messageFile_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatFilesAPI.ChatPostFiles(context.Background()).File(file).FileId(fileId).Endpoint(endpoint).AgentId(agentId).ToolResource(toolResource).MessageFile(messageFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatFilesAPI.ChatPostFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostFiles`: ChatFile
	fmt.Fprintf(os.Stdout, "Response from `ChatFilesAPI.ChatPostFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **fileId** | **string** |  | 
 **endpoint** | **string** |  | 
 **agentId** | **string** |  | 
 **toolResource** | **string** |  | 
 **messageFile** | **string** |  | 

### Return type

[**ChatFile**](ChatFile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

