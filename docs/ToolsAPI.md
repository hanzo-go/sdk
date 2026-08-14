# \ToolsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTools**](ToolsAPI.md#GetTools) | **Get** /v1/tools | Lists every tool the caller&#39;s org and project can reach, from every source, each flagged with whether it is activated.
[**GetToolsActivation**](ToolsAPI.md#GetToolsActivation) | **Get** /v1/tools/activation | Reports which tools are switched on for the caller&#39;s org and project.
[**GetToolsCatalog**](ToolsAPI.md#GetToolsCatalog) | **Get** /v1/tools/catalog | Lists the MCP servers the public registries publish, as we hold them: our canonical copy of registry.modelcontextprotocol.io, plus what we decided about each entry.
[**GetToolsCatalogById**](ToolsAPI.md#GetToolsCatalogById) | **Get** /v1/tools/catalog/{id} | Returns one catalog entry in full: the publisher&#39;s description, its repository and site, every package form with the runtime that launches it, and every hosted endpoint.
[**PatchToolsCatalogById**](ToolsAPI.md#PatchToolsCatalogById) | **Patch** /v1/tools/catalog/{id} | Sets what WE say about one catalog entry — hidden, featured, official, logo — and answers with the stored listing.
[**PostToolsCall**](ToolsAPI.md#PostToolsCall) | **Post** /v1/tools/call | Runs one of the caller&#39;s activated tools and answers with its output.
[**PostToolsCatalogSync**](ToolsAPI.md#PostToolsCatalogSync) | **Post** /v1/tools/catalog/sync | Pulls the public MCP registry into our canonical copy and reports what changed.
[**PutToolsActivation**](ToolsAPI.md#PutToolsActivation) | **Put** /v1/tools/activation | Switches tools on and off for the caller&#39;s org and project, and answers with the resulting activated set.



## GetTools

> ToolList GetTools(ctx).Source(source).Activated(activated).Execute()

Lists every tool the caller's org and project can reach, from every source, each flagged with whether it is activated.



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
	source := "source_example" // string | Source keeps only tools from one source — connector, function, zap-service, agent, skill or mcp. Empty keeps every source. (optional)
	activated := "activated_example" // string | Activated keeps only the tools activated for the caller's org and project, and only when it is exactly the string \"true\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.GetTools(context.Background()).Source(source).Activated(activated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTools`: ToolList
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetTools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | Source keeps only tools from one source — connector, function, zap-service, agent, skill or mcp. Empty keeps every source. | 
 **activated** | **string** | Activated keeps only the tools activated for the caller&#39;s org and project, and only when it is exactly the string \&quot;true\&quot;. | 

### Return type

[**ToolList**](ToolList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsActivation

> ActivationSet GetToolsActivation(ctx).Execute()

Reports which tools are switched on for the caller's org and project.



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
	resp, r, err := apiClient.ToolsAPI.GetToolsActivation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsActivation`: ActivationSet
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsActivation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsActivationRequest struct via the builder pattern


### Return type

[**ActivationSet**](ActivationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsCatalog

> McpCatalog GetToolsCatalog(ctx).Q(q).Featured(featured).Official(official).Limit(limit).Offset(offset).Execute()

Lists the MCP servers the public registries publish, as we hold them: our canonical copy of registry.modelcontextprotocol.io, plus what we decided about each entry.



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
	q := "q_example" // string | Q matches the name, title or description, case-insensitively. (optional)
	featured := "featured_example" // string | Featured keeps only the listings we put on the front of the shelf, and only when it is exactly the string \"true\". (optional)
	official := "official_example" // string | Official keeps only the vendors' OWN servers — not third-party copies of them — and only when it is exactly the string \"true\". (optional)
	limit := int32(56) // int32 | Limit bounds the page: default 50, maximum 200. A value that is not a positive integer reads as the default. (optional)
	offset := int32(56) // int32 | Offset skips that many listings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.GetToolsCatalog(context.Background()).Q(q).Featured(featured).Official(official).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsCatalog`: McpCatalog
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q matches the name, title or description, case-insensitively. | 
 **featured** | **string** | Featured keeps only the listings we put on the front of the shelf, and only when it is exactly the string \&quot;true\&quot;. | 
 **official** | **string** | Official keeps only the vendors&#39; OWN servers — not third-party copies of them — and only when it is exactly the string \&quot;true\&quot;. | 
 **limit** | **int32** | Limit bounds the page: default 50, maximum 200. A value that is not a positive integer reads as the default. | 
 **offset** | **int32** | Offset skips that many listings. | 

### Return type

[**McpCatalog**](McpCatalog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsCatalogById

> MCPListing GetToolsCatalogById(ctx, id).Execute()

Returns one catalog entry in full: the publisher's description, its repository and site, every package form with the runtime that launches it, and every hosted endpoint.



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
	id := "id_example" // string | ID is the listing, from the path. It is the publisher's reverse-DNS name with its one slash written as an underscore — \"com.stripe_mcp\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.GetToolsCatalogById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsCatalogById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsCatalogById`: MCPListing
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsCatalogById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the listing, from the path. It is the publisher&#39;s reverse-DNS name with its one slash written as an underscore — \&quot;com.stripe_mcp\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsCatalogByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MCPListing**](MCPListing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchToolsCatalogById

> MCPListing PatchToolsCatalogById(ctx, id).CurateReq(curateReq).Execute()

Sets what WE say about one catalog entry — hidden, featured, official, logo — and answers with the stored listing.



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
	id := "id_example" // string | ID is the listing to curate, from the path.
	curateReq := *openapiclient.NewCurateReq() // CurateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.PatchToolsCatalogById(context.Background(), id).CurateReq(curateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.PatchToolsCatalogById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchToolsCatalogById`: MCPListing
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.PatchToolsCatalogById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the listing to curate, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchToolsCatalogByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **curateReq** | [**CurateReq**](CurateReq.md) |  | 

### Return type

[**MCPListing**](MCPListing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolsCall

> ToolResult PostToolsCall(ctx).ToolCall(toolCall).Execute()

Runs one of the caller's activated tools and answers with its output.



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
	toolCall := *openapiclient.NewToolCall() // ToolCall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.PostToolsCall(context.Background()).ToolCall(toolCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.PostToolsCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolsCall`: ToolResult
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.PostToolsCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostToolsCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolCall** | [**ToolCall**](ToolCall.md) |  | 

### Return type

[**ToolResult**](ToolResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolsCatalogSync

> McpCatalogSync PostToolsCatalogSync(ctx).Execute()

Pulls the public MCP registry into our canonical copy and reports what changed.



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
	resp, r, err := apiClient.ToolsAPI.PostToolsCatalogSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.PostToolsCatalogSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolsCatalogSync`: McpCatalogSync
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.PostToolsCatalogSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostToolsCatalogSyncRequest struct via the builder pattern


### Return type

[**McpCatalogSync**](McpCatalogSync.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutToolsActivation

> ActivationSet PutToolsActivation(ctx).ActivationReq(activationReq).Execute()

Switches tools on and off for the caller's org and project, and answers with the resulting activated set.



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
	activationReq := *openapiclient.NewActivationReq() // ActivationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.PutToolsActivation(context.Background()).ActivationReq(activationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.PutToolsActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutToolsActivation`: ActivationSet
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.PutToolsActivation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutToolsActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activationReq** | [**ActivationReq**](ActivationReq.md) |  | 

### Return type

[**ActivationSet**](ActivationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

