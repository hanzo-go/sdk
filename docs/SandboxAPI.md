# \SandboxAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSandboxById**](SandboxAPI.md#DeleteSandboxById) | **Delete** /v1/sandbox/{id} | Ends a sandbox and releases the compute behind it.
[**EndSandbox**](SandboxAPI.md#EndSandbox) | **Post** /v1/sandbox/end | End a sandbox and release it
[**GetSandbox**](SandboxAPI.md#GetSandbox) | **Get** /v1/sandbox | Lists the sandboxes the caller holds, newest first.
[**GetSandboxById**](SandboxAPI.md#GetSandboxById) | **Get** /v1/sandbox/{id} | Returns one sandbox: its class, project, image, the runtime it was given, its status and when its lease ends.
[**GetSandboxByIdFs**](SandboxAPI.md#GetSandboxByIdFs) | **Get** /v1/sandbox/{id}/fs | Read a file, or list a directory
[**GetSandboxByIdPorts**](SandboxAPI.md#GetSandboxByIdPorts) | **Get** /v1/sandbox/{id}/ports | Lists the TCP ports something listens on in a sandbox the caller holds, each with the preview host that serves it — what the sandbox&#39;s Browser can open.
[**GetSandboxByIdScreen**](SandboxAPI.md#GetSandboxByIdScreen) | **Get** /v1/sandbox/{id}/screen | The screen, as a page
[**GetSandboxByIdScreenWs**](SandboxAPI.md#GetSandboxByIdScreenWs) | **Get** /v1/sandbox/{id}/screen/ws | The screen, as a socket
[**GetSandboxByIdTerminal**](SandboxAPI.md#GetSandboxByIdTerminal) | **Get** /v1/sandbox/{id}/terminal | The terminal, as a page
[**GetSandboxByIdTerminalWs**](SandboxAPI.md#GetSandboxByIdTerminalWs) | **Get** /v1/sandbox/{id}/terminal/ws | The terminal, as a socket
[**LeaseSandbox**](SandboxAPI.md#LeaseSandbox) | **Post** /v1/sandbox/lease | Lease a sandbox — a real computer — or resume one you hold
[**PostSandbox**](SandboxAPI.md#PostSandbox) | **Post** /v1/sandbox | Leases a sandbox — a real computer — for the caller&#39;s org.
[**PostSandboxByIdExec**](SandboxAPI.md#PostSandboxByIdExec) | **Post** /v1/sandbox/{id}/exec | Runs one command in a sandbox the caller holds and answers with its exit code, stdout and stderr.
[**PostSandboxByIdFs**](SandboxAPI.md#PostSandboxByIdFs) | **Post** /v1/sandbox/{id}/fs | Write a file
[**PostSandboxByIdPause**](SandboxAPI.md#PostSandboxByIdPause) | **Post** /v1/sandbox/{id}/pause | Stops the pod, keeps the row and the volume, and settles the tail: POST /v1/sandbox/:id/pause.
[**PostSandboxByIdPreview**](SandboxAPI.md#PostSandboxByIdPreview) | **Post** /v1/sandbox/{id}/preview | Opens a port of a sandbox the caller holds in a browser.
[**PostSandboxByIdResume**](SandboxAPI.md#PostSandboxByIdResume) | **Post** /v1/sandbox/{id}/resume | Gives a parked sandbox a pod again: POST /v1/sandbox/:id/resume.
[**PostSandboxByIdScreenTicket**](SandboxAPI.md#PostSandboxByIdScreenTicket) | **Post** /v1/sandbox/{id}/screen/ticket | Mints a short-lived grant to open the screen of a desktop sandbox.
[**PostSandboxByIdTerminalTicket**](SandboxAPI.md#PostSandboxByIdTerminalTicket) | **Post** /v1/sandbox/{id}/terminal/ticket | Mints a short-lived grant to open a terminal on a sandbox.
[**ReadSandboxFile**](SandboxAPI.md#ReadSandboxFile) | **Post** /v1/sandbox/read | Read a file from a sandbox you hold
[**RunInSandbox**](SandboxAPI.md#RunInSandbox) | **Post** /v1/sandbox/run | Run a command in a sandbox you hold and read its output
[**StopRun**](SandboxAPI.md#StopRun) | **Post** /v1/sandbox/stop | Stop what a sandbox is running, and keep the sandbox
[**WriteSandboxFile**](SandboxAPI.md#WriteSandboxFile) | **Post** /v1/sandbox/write | Write a file into a sandbox you hold



## DeleteSandboxById

> DeleteSandboxById(ctx, id).Purge(purge).Execute()

Ends a sandbox and releases the compute behind it.



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
	id := "id_example" // string | ID is the sandbox to end, from the path.
	purge := "purge_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.DeleteSandboxById(context.Background(), id).Purge(purge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.DeleteSandboxById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox to end, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSandboxByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndSandbox

> EndSandbox(ctx).SandboxEndIn(sandboxEndIn).Execute()

End a sandbox and release it



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
	sandboxEndIn := *openapiclient.NewSandboxEndIn() // SandboxEndIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.EndSandbox(context.Background()).SandboxEndIn(sandboxEndIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.EndSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxEndIn** | [**SandboxEndIn**](SandboxEndIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandbox

> SandboxSandboxList GetSandbox(ctx).Project(project).Status(status).Execute()

Lists the sandboxes the caller holds, newest first.



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
	project := "project_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.GetSandbox(context.Background()).Project(project).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSandbox`: SandboxSandboxList
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.GetSandbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**SandboxSandboxList**](SandboxSandboxList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxById

> SandboxSandbox GetSandboxById(ctx, id).Execute()

Returns one sandbox: its class, project, image, the runtime it was given, its status and when its lease ends.



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
	id := "id_example" // string | ID is the sandbox to address, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.GetSandboxById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandboxById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSandboxById`: SandboxSandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.GetSandboxById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox to address, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SandboxSandbox**](SandboxSandbox.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxByIdFs

> GetSandboxByIdFs(ctx, id).Execute()

Read a file, or list a directory



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.GetSandboxByIdFs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandboxByIdFs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxByIdFsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSandboxByIdPorts

> SandboxPorts GetSandboxByIdPorts(ctx, id).Execute()

Lists the TCP ports something listens on in a sandbox the caller holds, each with the preview host that serves it — what the sandbox's Browser can open.



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
	id := "id_example" // string | ID is the sandbox, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.GetSandboxByIdPorts(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandboxByIdPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSandboxByIdPorts`: SandboxPorts
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.GetSandboxByIdPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxByIdPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SandboxPorts**](SandboxPorts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxByIdScreen

> GetSandboxByIdScreen(ctx, id).Execute()

The screen, as a page



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.GetSandboxByIdScreen(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandboxByIdScreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxByIdScreenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSandboxByIdScreenWs

> GetSandboxByIdScreenWs(ctx, id).Execute()

The screen, as a socket



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.GetSandboxByIdScreenWs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandboxByIdScreenWs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxByIdScreenWsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSandboxByIdTerminal

> GetSandboxByIdTerminal(ctx, id).Execute()

The terminal, as a page



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.GetSandboxByIdTerminal(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandboxByIdTerminal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxByIdTerminalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSandboxByIdTerminalWs

> GetSandboxByIdTerminalWs(ctx, id).Execute()

The terminal, as a socket



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.GetSandboxByIdTerminalWs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandboxByIdTerminalWs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxByIdTerminalWsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## LeaseSandbox

> SandboxLeased LeaseSandbox(ctx).SandboxLeaseIn(sandboxLeaseIn).Execute()

Lease a sandbox — a real computer — or resume one you hold



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
	sandboxLeaseIn := *openapiclient.NewSandboxLeaseIn() // SandboxLeaseIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.LeaseSandbox(context.Background()).SandboxLeaseIn(sandboxLeaseIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.LeaseSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeaseSandbox`: SandboxLeased
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.LeaseSandbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLeaseSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxLeaseIn** | [**SandboxLeaseIn**](SandboxLeaseIn.md) |  | 

### Return type

[**SandboxLeased**](SandboxLeased.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandbox

> SandboxSandbox PostSandbox(ctx).SandboxSandboxIn(sandboxSandboxIn).Execute()

Leases a sandbox — a real computer — for the caller's org.



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
	sandboxSandboxIn := *openapiclient.NewSandboxSandboxIn() // SandboxSandboxIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandbox(context.Background()).SandboxSandboxIn(sandboxSandboxIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandbox`: SandboxSandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.PostSandbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxSandboxIn** | [**SandboxSandboxIn**](SandboxSandboxIn.md) |  | 

### Return type

[**SandboxSandbox**](SandboxSandbox.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandboxByIdExec

> SandboxExecResult PostSandboxByIdExec(ctx, id).SandboxExecRequest(sandboxExecRequest).Execute()

Runs one command in a sandbox the caller holds and answers with its exit code, stdout and stderr.



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
	id := "id_example" // string | ID is the sandbox to run in, from the path.
	sandboxExecRequest := *openapiclient.NewSandboxExecRequest() // SandboxExecRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdExec(context.Background(), id).SandboxExecRequest(sandboxExecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdExec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdExec`: SandboxExecResult
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.PostSandboxByIdExec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox to run in, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxByIdExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sandboxExecRequest** | [**SandboxExecRequest**](SandboxExecRequest.md) |  | 

### Return type

[**SandboxExecResult**](SandboxExecResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandboxByIdFs

> PostSandboxByIdFs(ctx, id).Execute()

Write a file



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.PostSandboxByIdFs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdFs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxByIdFsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostSandboxByIdPause

> SandboxSandbox PostSandboxByIdPause(ctx, id).Execute()

Stops the pod, keeps the row and the volume, and settles the tail: POST /v1/sandbox/:id/pause.



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
	id := "id_example" // string | ID is the sandbox to address, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdPause(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdPause`: SandboxSandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.PostSandboxByIdPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox to address, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxByIdPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SandboxSandbox**](SandboxSandbox.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandboxByIdPreview

> SandboxPreviewGrant PostSandboxByIdPreview(ctx, id).SandboxPreviewIn(sandboxPreviewIn).Execute()

Opens a port of a sandbox the caller holds in a browser.



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
	id := "id_example" // string | ID is the sandbox, from the path.
	sandboxPreviewIn := *openapiclient.NewSandboxPreviewIn() // SandboxPreviewIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdPreview(context.Background(), id).SandboxPreviewIn(sandboxPreviewIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdPreview`: SandboxPreviewGrant
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.PostSandboxByIdPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxByIdPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sandboxPreviewIn** | [**SandboxPreviewIn**](SandboxPreviewIn.md) |  | 

### Return type

[**SandboxPreviewGrant**](SandboxPreviewGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandboxByIdResume

> SandboxSandbox PostSandboxByIdResume(ctx, id).Execute()

Gives a parked sandbox a pod again: POST /v1/sandbox/:id/resume.



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
	id := "id_example" // string | ID is the sandbox to address, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdResume(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdResume`: SandboxSandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.PostSandboxByIdResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox to address, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxByIdResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SandboxSandbox**](SandboxSandbox.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandboxByIdScreenTicket

> SandboxTicketGrant PostSandboxByIdScreenTicket(ctx, id).Execute()

Mints a short-lived grant to open the screen of a desktop sandbox.



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
	id := "id_example" // string | ID is the sandbox to address, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdScreenTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdScreenTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdScreenTicket`: SandboxTicketGrant
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.PostSandboxByIdScreenTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox to address, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxByIdScreenTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SandboxTicketGrant**](SandboxTicketGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandboxByIdTerminalTicket

> SandboxTicketGrant PostSandboxByIdTerminalTicket(ctx, id).Execute()

Mints a short-lived grant to open a terminal on a sandbox.



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
	id := "id_example" // string | ID is the sandbox to address, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdTerminalTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdTerminalTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdTerminalTicket`: SandboxTicketGrant
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.PostSandboxByIdTerminalTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sandbox to address, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxByIdTerminalTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SandboxTicketGrant**](SandboxTicketGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSandboxFile

> SandboxBlob ReadSandboxFile(ctx).SandboxPathIn(sandboxPathIn).Execute()

Read a file from a sandbox you hold



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
	sandboxPathIn := *openapiclient.NewSandboxPathIn() // SandboxPathIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.ReadSandboxFile(context.Background()).SandboxPathIn(sandboxPathIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.ReadSandboxFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadSandboxFile`: SandboxBlob
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.ReadSandboxFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadSandboxFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxPathIn** | [**SandboxPathIn**](SandboxPathIn.md) |  | 

### Return type

[**SandboxBlob**](SandboxBlob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunInSandbox

> SandboxRan RunInSandbox(ctx).SandboxRunIn(sandboxRunIn).Execute()

Run a command in a sandbox you hold and read its output



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
	sandboxRunIn := *openapiclient.NewSandboxRunIn() // SandboxRunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.RunInSandbox(context.Background()).SandboxRunIn(sandboxRunIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.RunInSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunInSandbox`: SandboxRan
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.RunInSandbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunInSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxRunIn** | [**SandboxRunIn**](SandboxRunIn.md) |  | 

### Return type

[**SandboxRan**](SandboxRan.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopRun

> SandboxStopped StopRun(ctx).SandboxStopIn(sandboxStopIn).Execute()

Stop what a sandbox is running, and keep the sandbox



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
	sandboxStopIn := *openapiclient.NewSandboxStopIn() // SandboxStopIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.StopRun(context.Background()).SandboxStopIn(sandboxStopIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.StopRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopRun`: SandboxStopped
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.StopRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxStopIn** | [**SandboxStopIn**](SandboxStopIn.md) |  | 

### Return type

[**SandboxStopped**](SandboxStopped.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteSandboxFile

> SandboxWrote WriteSandboxFile(ctx).SandboxWriteIn(sandboxWriteIn).Execute()

Write a file into a sandbox you hold



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
	sandboxWriteIn := *openapiclient.NewSandboxWriteIn() // SandboxWriteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.WriteSandboxFile(context.Background()).SandboxWriteIn(sandboxWriteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.WriteSandboxFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WriteSandboxFile`: SandboxWrote
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.WriteSandboxFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWriteSandboxFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxWriteIn** | [**SandboxWriteIn**](SandboxWriteIn.md) |  | 

### Return type

[**SandboxWrote**](SandboxWrote.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

