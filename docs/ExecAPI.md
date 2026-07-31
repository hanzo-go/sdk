# \ExecAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecExecCode**](ExecAPI.md#ExecExecCode) | **Post** /v1/exec | Execute code in a sandboxed session
[**ExecExecProgrammatic**](ExecAPI.md#ExecExecProgrammatic) | **Post** /v1/exec/programmatic | Programmatic code execution (sibling of /v1/exec, same executor contract)



## ExecExecCode

> ExecExecResult ExecExecCode(ctx).ExecExecRequest(execExecRequest).Execute()

Execute code in a sandboxed session

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
	execExecRequest := *openapiclient.NewExecExecRequest("Lang_example", "Code_example") // ExecExecRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecAPI.ExecExecCode(context.Background()).ExecExecRequest(execExecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecAPI.ExecExecCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecExecCode`: ExecExecResult
	fmt.Fprintf(os.Stdout, "Response from `ExecAPI.ExecExecCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecExecCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **execExecRequest** | [**ExecExecRequest**](ExecExecRequest.md) |  | 

### Return type

[**ExecExecResult**](ExecExecResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecExecProgrammatic

> ExecExecResult ExecExecProgrammatic(ctx).ExecExecRequest(execExecRequest).Execute()

Programmatic code execution (sibling of /v1/exec, same executor contract)

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
	execExecRequest := *openapiclient.NewExecExecRequest("Lang_example", "Code_example") // ExecExecRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecAPI.ExecExecProgrammatic(context.Background()).ExecExecRequest(execExecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecAPI.ExecExecProgrammatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecExecProgrammatic`: ExecExecResult
	fmt.Fprintf(os.Stdout, "Response from `ExecAPI.ExecExecProgrammatic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecExecProgrammaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **execExecRequest** | [**ExecExecRequest**](ExecExecRequest.md) |  | 

### Return type

[**ExecExecResult**](ExecExecResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

