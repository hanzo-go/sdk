# \ExecFilesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecDownloadFile**](ExecFilesAPI.md#ExecDownloadFile) | **Get** /v1/download/{id} | Download a produced file by id
[**ExecListSessionFiles**](ExecFilesAPI.md#ExecListSessionFiles) | **Get** /v1/files/{session_id} | List the files produced in a session
[**ExecUploadFile**](ExecFilesAPI.md#ExecUploadFile) | **Post** /v1/upload | Upload a file into a session



## ExecDownloadFile

> *os.File ExecDownloadFile(ctx, id).Execute()

Download a produced file by id

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
	resp, r, err := apiClient.ExecFilesAPI.ExecDownloadFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecFilesAPI.ExecDownloadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecDownloadFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExecFilesAPI.ExecDownloadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecDownloadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecListSessionFiles

> ExecListSessionFiles200Response ExecListSessionFiles(ctx, sessionId).Execute()

List the files produced in a session

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecFilesAPI.ExecListSessionFiles(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecFilesAPI.ExecListSessionFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecListSessionFiles`: ExecListSessionFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `ExecFilesAPI.ExecListSessionFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecListSessionFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExecListSessionFiles200Response**](ExecListSessionFiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecUploadFile

> ExecUploadFile200Response ExecUploadFile(ctx).File(file).SessionId(sessionId).Execute()

Upload a file into a session

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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	sessionId := "sessionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecFilesAPI.ExecUploadFile(context.Background()).File(file).SessionId(sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecFilesAPI.ExecUploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecUploadFile`: ExecUploadFile200Response
	fmt.Fprintf(os.Stdout, "Response from `ExecFilesAPI.ExecUploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **sessionId** | **string** |  | 

### Return type

[**ExecUploadFile200Response**](ExecUploadFile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

