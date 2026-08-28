# \ExecAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecFilesBySid**](ExecAPI.md#GetExecFilesBySid) | **Get** /v1/exec/files/{sid} | Lists the files in an execution session.
[**PostExec**](ExecAPI.md#PostExec) | **Post** /v1/exec | Run a code snippet in a sandboxed interpreter
[**PostExecProgrammatic**](ExecAPI.md#PostExecProgrammatic) | **Post** /v1/exec/programmatic | Answers 501 — this deployment does not serve programmatic tool calling.
[**PostExecUpload**](ExecAPI.md#PostExecUpload) | **Post** /v1/exec/upload | Upload a file into an execution session



## GetExecFilesBySid

> []Listing GetExecFilesBySid(ctx, sid).Execute()

Lists the files in an execution session.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sid := "sid_example" // string | SID is the session identifier — the sandbox this listing is of. The URL is the addressing authority: a path segment binds after the body and after the query, so the address decides which session is read whatever else is sent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecAPI.GetExecFilesBySid(context.Background(), sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecAPI.GetExecFilesBySid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecFilesBySid`: []Listing
	fmt.Fprintf(os.Stdout, "Response from `ExecAPI.GetExecFilesBySid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | SID is the session identifier — the sandbox this listing is of. The URL is the addressing authority: a path segment binds after the body and after the query, so the address decides which session is read whatever else is sent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecFilesBySidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Listing**](Listing.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExec

> CodeResult PostExec(ctx).CodeRun(codeRun).Execute()

Run a code snippet in a sandboxed interpreter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	codeRun := *openapiclient.NewCodeRun("Code_example", "Lang_example") // CodeRun | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecAPI.PostExec(context.Background()).CodeRun(codeRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecAPI.PostExec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExec`: CodeResult
	fmt.Fprintf(os.Stdout, "Response from `ExecAPI.PostExec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeRun** | [**CodeRun**](CodeRun.md) |  | 

### Return type

[**CodeResult**](CodeResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExecProgrammatic

> PostExecProgrammatic(ctx).Execute()

Answers 501 — this deployment does not serve programmatic tool calling.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecAPI.PostExecProgrammatic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecAPI.PostExecProgrammatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostExecProgrammaticRequest struct via the builder pattern


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


## PostExecUpload

> PostExecUpload(ctx).Execute()

Upload a file into an execution session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecAPI.PostExecUpload(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecAPI.PostExecUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostExecUploadRequest struct via the builder pattern


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

