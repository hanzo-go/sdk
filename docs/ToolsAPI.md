# \ToolsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Tools**](ToolsAPI.md#CloudGetV1Tools) | **Get** /v1/tools | ListTools lists every tool the caller&#39;s org and project can reach, from every source, each flagged with whether it is activated.
[**CloudGetV1ToolsActivation**](ToolsAPI.md#CloudGetV1ToolsActivation) | **Get** /v1/tools/activation | GetActivation reports which tools are switched on for the caller&#39;s org and project.
[**CloudGetV1ToolsCatalog**](ToolsAPI.md#CloudGetV1ToolsCatalog) | **Get** /v1/tools/catalog | ListCatalog lists the MCP servers the public registries publish, as we hold them: our canonical copy of registry.modelcontextprotocol.io, plus what we decided about each entry.
[**CloudGetV1ToolsCatalogId**](ToolsAPI.md#CloudGetV1ToolsCatalogId) | **Get** /v1/tools/catalog/{id} | GetListing returns one catalog entry in full: the publisher&#39;s description, its repository and site, every package form with the runtime that launches it, and every hosted endpoint.
[**CloudPatchV1ToolsCatalogId**](ToolsAPI.md#CloudPatchV1ToolsCatalogId) | **Patch** /v1/tools/catalog/{id} | CurateListing sets what WE say about one catalog entry — hidden, featured, official, logo — and answers with the stored listing.
[**CloudPostV1ToolsCall**](ToolsAPI.md#CloudPostV1ToolsCall) | **Post** /v1/tools/call | CallTool runs one of the caller&#39;s activated tools and answers with its output.
[**CloudPostV1ToolsCatalogSync**](ToolsAPI.md#CloudPostV1ToolsCatalogSync) | **Post** /v1/tools/catalog/sync | SyncCatalog pulls the public MCP registry into our canonical copy and reports what changed.
[**CloudPutV1ToolsActivation**](ToolsAPI.md#CloudPutV1ToolsActivation) | **Put** /v1/tools/activation | PutActivation switches tools on and off for the caller&#39;s org and project, and answers with the resulting activated set.



## CloudGetV1Tools

> CloudToolList CloudGetV1Tools(ctx).Source(source).Activated(activated).Execute()

ListTools lists every tool the caller's org and project can reach, from every source, each flagged with whether it is activated.



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
	resp, r, err := apiClient.ToolsAPI.CloudGetV1Tools(context.Background()).Source(source).Activated(activated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CloudGetV1Tools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Tools`: CloudToolList
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CloudGetV1Tools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | Source keeps only tools from one source — connector, function, zap-service, agent, skill or mcp. Empty keeps every source. | 
 **activated** | **string** | Activated keeps only the tools activated for the caller&#39;s org and project, and only when it is exactly the string \&quot;true\&quot;. | 

### Return type

[**CloudToolList**](CloudToolList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ToolsActivation

> CloudActivationSet CloudGetV1ToolsActivation(ctx).Execute()

GetActivation reports which tools are switched on for the caller's org and project.



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
	resp, r, err := apiClient.ToolsAPI.CloudGetV1ToolsActivation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CloudGetV1ToolsActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ToolsActivation`: CloudActivationSet
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CloudGetV1ToolsActivation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ToolsActivationRequest struct via the builder pattern


### Return type

[**CloudActivationSet**](CloudActivationSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ToolsCatalog

> CloudMcpCatalog CloudGetV1ToolsCatalog(ctx).Q(q).Featured(featured).Official(official).Limit(limit).Offset(offset).Execute()

ListCatalog lists the MCP servers the public registries publish, as we hold them: our canonical copy of registry.modelcontextprotocol.io, plus what we decided about each entry.



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
	resp, r, err := apiClient.ToolsAPI.CloudGetV1ToolsCatalog(context.Background()).Q(q).Featured(featured).Official(official).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CloudGetV1ToolsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ToolsCatalog`: CloudMcpCatalog
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CloudGetV1ToolsCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ToolsCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q matches the name, title or description, case-insensitively. | 
 **featured** | **string** | Featured keeps only the listings we put on the front of the shelf, and only when it is exactly the string \&quot;true\&quot;. | 
 **official** | **string** | Official keeps only the vendors&#39; OWN servers — not third-party copies of them — and only when it is exactly the string \&quot;true\&quot;. | 
 **limit** | **int32** | Limit bounds the page: default 50, maximum 200. A value that is not a positive integer reads as the default. | 
 **offset** | **int32** | Offset skips that many listings. | 

### Return type

[**CloudMcpCatalog**](CloudMcpCatalog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ToolsCatalogId

> CloudMCPListing CloudGetV1ToolsCatalogId(ctx, id).Execute()

GetListing returns one catalog entry in full: the publisher's description, its repository and site, every package form with the runtime that launches it, and every hosted endpoint.



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
	resp, r, err := apiClient.ToolsAPI.CloudGetV1ToolsCatalogId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CloudGetV1ToolsCatalogId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ToolsCatalogId`: CloudMCPListing
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CloudGetV1ToolsCatalogId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the listing, from the path. It is the publisher&#39;s reverse-DNS name with its one slash written as an underscore — \&quot;com.stripe_mcp\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ToolsCatalogIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudMCPListing**](CloudMCPListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1ToolsCatalogId

> CloudMCPListing CloudPatchV1ToolsCatalogId(ctx, id).CloudCurateReq(cloudCurateReq).Execute()

CurateListing sets what WE say about one catalog entry — hidden, featured, official, logo — and answers with the stored listing.



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
	cloudCurateReq := *openapiclient.NewCloudCurateReq() // CloudCurateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.CloudPatchV1ToolsCatalogId(context.Background(), id).CloudCurateReq(cloudCurateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CloudPatchV1ToolsCatalogId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1ToolsCatalogId`: CloudMCPListing
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CloudPatchV1ToolsCatalogId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the listing to curate, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1ToolsCatalogIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCurateReq** | [**CloudCurateReq**](CloudCurateReq.md) |  | 

### Return type

[**CloudMCPListing**](CloudMCPListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ToolsCall

> CloudToolResult CloudPostV1ToolsCall(ctx).CloudToolCall(cloudToolCall).Execute()

CallTool runs one of the caller's activated tools and answers with its output.



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
	cloudToolCall := *openapiclient.NewCloudToolCall() // CloudToolCall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.CloudPostV1ToolsCall(context.Background()).CloudToolCall(cloudToolCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CloudPostV1ToolsCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ToolsCall`: CloudToolResult
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CloudPostV1ToolsCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ToolsCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudToolCall** | [**CloudToolCall**](CloudToolCall.md) |  | 

### Return type

[**CloudToolResult**](CloudToolResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ToolsCatalogSync

> CloudMcpCatalogSync CloudPostV1ToolsCatalogSync(ctx).Execute()

SyncCatalog pulls the public MCP registry into our canonical copy and reports what changed.



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
	resp, r, err := apiClient.ToolsAPI.CloudPostV1ToolsCatalogSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CloudPostV1ToolsCatalogSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ToolsCatalogSync`: CloudMcpCatalogSync
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CloudPostV1ToolsCatalogSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ToolsCatalogSyncRequest struct via the builder pattern


### Return type

[**CloudMcpCatalogSync**](CloudMcpCatalogSync.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1ToolsActivation

> CloudActivationSet CloudPutV1ToolsActivation(ctx).CloudActivationReq(cloudActivationReq).Execute()

PutActivation switches tools on and off for the caller's org and project, and answers with the resulting activated set.



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
	cloudActivationReq := *openapiclient.NewCloudActivationReq() // CloudActivationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.CloudPutV1ToolsActivation(context.Background()).CloudActivationReq(cloudActivationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CloudPutV1ToolsActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1ToolsActivation`: CloudActivationSet
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CloudPutV1ToolsActivation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1ToolsActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudActivationReq** | [**CloudActivationReq**](CloudActivationReq.md) |  | 

### Return type

[**CloudActivationSet**](CloudActivationSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

