# \SandboxAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSandboxById**](SandboxAPI.md#DeleteSandboxById) | **Delete** /v1/sandbox/{id} | Ends a sandbox and releases the compute behind it.
[**EndSandbox**](SandboxAPI.md#EndSandbox) | **Post** /v1/sandbox/end | End a sandbox and release it
[**GetSandbox**](SandboxAPI.md#GetSandbox) | **Get** /v1/sandbox | Lists the caller org&#39;s sandboxes, newest first.
[**GetSandboxById**](SandboxAPI.md#GetSandboxById) | **Get** /v1/sandbox/{id} | Returns one sandbox: its class, project, image, the runtime it was given, its status and when its lease ends.
[**GetSandboxByIdFs**](SandboxAPI.md#GetSandboxByIdFs) | **Get** /v1/sandbox/{id}/fs | Read a file, or list a directory
[**GetSandboxByIdScreen**](SandboxAPI.md#GetSandboxByIdScreen) | **Get** /v1/sandbox/{id}/screen | The screen, as a page
[**GetSandboxByIdScreenWs**](SandboxAPI.md#GetSandboxByIdScreenWs) | **Get** /v1/sandbox/{id}/screen/ws | The screen, as a socket
[**GetSandboxByIdTerminal**](SandboxAPI.md#GetSandboxByIdTerminal) | **Get** /v1/sandbox/{id}/terminal | The terminal, as a page
[**GetSandboxByIdTerminalWs**](SandboxAPI.md#GetSandboxByIdTerminalWs) | **Get** /v1/sandbox/{id}/terminal/ws | The terminal, as a socket
[**LeaseSandbox**](SandboxAPI.md#LeaseSandbox) | **Post** /v1/sandbox/lease | Lease a sandbox — a real computer — or resume one you hold
[**PostSandbox**](SandboxAPI.md#PostSandbox) | **Post** /v1/sandbox | Leases a sandbox — a real computer — for the caller&#39;s org.
[**PostSandboxByIdExec**](SandboxAPI.md#PostSandboxByIdExec) | **Post** /v1/sandbox/{id}/exec | Runs one command in a sandbox the caller holds and answers with its exit code, stdout and stderr.
[**PostSandboxByIdFs**](SandboxAPI.md#PostSandboxByIdFs) | **Post** /v1/sandbox/{id}/fs | Write a file
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
	openapiclient "github.com/hanzoai/go-sdk"
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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndSandbox

> EndSandbox(ctx).EndIn(endIn).Execute()

End a sandbox and release it



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
	endIn := *openapiclient.NewEndIn() // EndIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SandboxAPI.EndSandbox(context.Background()).EndIn(endIn).Execute()
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
 **endIn** | [**EndIn**](EndIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandbox

> SandboxList GetSandbox(ctx).Project(project).Status(status).Execute()

Lists the caller org's sandboxes, newest first.



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
	project := "project_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.GetSandbox(context.Background()).Project(project).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSandbox`: SandboxList
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

[**SandboxList**](SandboxList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxById

> Sandbox GetSandboxById(ctx, id).Execute()

Returns one sandbox: its class, project, image, the runtime it was given, its status and when its lease ends.



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
	id := "id_example" // string | ID is the sandbox to address, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.GetSandboxById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandboxById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSandboxById`: Sandbox
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

[**Sandbox**](Sandbox.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/hanzoai/go-sdk"
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
	openapiclient "github.com/hanzoai/go-sdk"
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
	openapiclient "github.com/hanzoai/go-sdk"
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
	openapiclient "github.com/hanzoai/go-sdk"
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
	openapiclient "github.com/hanzoai/go-sdk"
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

> Leased LeaseSandbox(ctx).LeaseIn(leaseIn).Execute()

Lease a sandbox — a real computer — or resume one you hold



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
	leaseIn := *openapiclient.NewLeaseIn() // LeaseIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.LeaseSandbox(context.Background()).LeaseIn(leaseIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.LeaseSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeaseSandbox`: Leased
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.LeaseSandbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLeaseSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leaseIn** | [**LeaseIn**](LeaseIn.md) |  | 

### Return type

[**Leased**](Leased.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandbox

> Sandbox PostSandbox(ctx).LeaseIn(leaseIn).Execute()

Leases a sandbox — a real computer — for the caller's org.



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
	leaseIn := *openapiclient.NewLeaseIn() // LeaseIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandbox(context.Background()).LeaseIn(leaseIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandbox`: Sandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.PostSandbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leaseIn** | [**LeaseIn**](LeaseIn.md) |  | 

### Return type

[**Sandbox**](Sandbox.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandboxByIdExec

> ExecResult PostSandboxByIdExec(ctx, id).ExecRequest(execRequest).Execute()

Runs one command in a sandbox the caller holds and answers with its exit code, stdout and stderr.



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
	id := "id_example" // string | ID is the sandbox to run in, from the path.
	execRequest := *openapiclient.NewExecRequest() // ExecRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdExec(context.Background(), id).ExecRequest(execRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdExec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdExec`: ExecResult
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

 **execRequest** | [**ExecRequest**](ExecRequest.md) |  | 

### Return type

[**ExecResult**](ExecResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
	openapiclient "github.com/hanzoai/go-sdk"
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


## PostSandboxByIdScreenTicket

> TicketGrant PostSandboxByIdScreenTicket(ctx, id).Execute()

Mints a short-lived grant to open the screen of a desktop sandbox.



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
	id := "id_example" // string | ID is the sandbox to address, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdScreenTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdScreenTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdScreenTicket`: TicketGrant
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

[**TicketGrant**](TicketGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSandboxByIdTerminalTicket

> TicketGrant PostSandboxByIdTerminalTicket(ctx, id).Execute()

Mints a short-lived grant to open a terminal on a sandbox.



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
	id := "id_example" // string | ID is the sandbox to address, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.PostSandboxByIdTerminalTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.PostSandboxByIdTerminalTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSandboxByIdTerminalTicket`: TicketGrant
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

[**TicketGrant**](TicketGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSandboxFile

> Blob ReadSandboxFile(ctx).PathIn(pathIn).Execute()

Read a file from a sandbox you hold



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
	pathIn := *openapiclient.NewPathIn() // PathIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.ReadSandboxFile(context.Background()).PathIn(pathIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.ReadSandboxFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadSandboxFile`: Blob
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.ReadSandboxFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadSandboxFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pathIn** | [**PathIn**](PathIn.md) |  | 

### Return type

[**Blob**](Blob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunInSandbox

> Ran RunInSandbox(ctx).RunIn(runIn).Execute()

Run a command in a sandbox you hold and read its output



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
	runIn := *openapiclient.NewRunIn() // RunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.RunInSandbox(context.Background()).RunIn(runIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.RunInSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunInSandbox`: Ran
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.RunInSandbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunInSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runIn** | [**RunIn**](RunIn.md) |  | 

### Return type

[**Ran**](Ran.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopRun

> Stopped StopRun(ctx).StopIn(stopIn).Execute()

Stop what a sandbox is running, and keep the sandbox



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
	stopIn := *openapiclient.NewStopIn() // StopIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.StopRun(context.Background()).StopIn(stopIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.StopRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopRun`: Stopped
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.StopRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopIn** | [**StopIn**](StopIn.md) |  | 

### Return type

[**Stopped**](Stopped.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteSandboxFile

> Wrote WriteSandboxFile(ctx).WriteIn(writeIn).Execute()

Write a file into a sandbox you hold



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
	writeIn := *openapiclient.NewWriteIn() // WriteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.WriteSandboxFile(context.Background()).WriteIn(writeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.WriteSandboxFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WriteSandboxFile`: Wrote
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.WriteSandboxFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWriteSandboxFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writeIn** | [**WriteIn**](WriteIn.md) |  | 

### Return type

[**Wrote**](Wrote.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

