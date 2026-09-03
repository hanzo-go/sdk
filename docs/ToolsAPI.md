# \ToolsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToolsMcpServersById**](ToolsAPI.md#DeleteToolsMcpServersById) | **Delete** /v1/tools/mcp/servers/{id} | Deregisters one of the caller org&#39;s external MCP servers, so its tools leave the registry.
[**DeleteToolsPluginsAuthoredById**](ToolsAPI.md#DeleteToolsPluginsAuthoredById) | **Delete** /v1/tools/plugins/authored/{id} | Removes one of the caller org&#39;s built plugins, so the runtime can no longer load it.
[**DeleteToolsSkillsById**](ToolsAPI.md#DeleteToolsSkillsById) | **Delete** /v1/tools/skills/{id} | Removes one of the caller org&#39;s authored skills.
[**GetTools**](ToolsAPI.md#GetTools) | **Get** /v1/tools | Lists every tool the caller&#39;s org and project can reach, from every source, each flagged with whether it is activated.
[**GetToolsActivation**](ToolsAPI.md#GetToolsActivation) | **Get** /v1/tools/activation | Reports which tools are switched on for the caller&#39;s org and project.
[**GetToolsCatalog**](ToolsAPI.md#GetToolsCatalog) | **Get** /v1/tools/catalog | Lists the MCP servers the public registries publish, as we hold them: our canonical copy of registry.modelcontextprotocol.io, plus what we decided about each entry.
[**GetToolsCatalogById**](ToolsAPI.md#GetToolsCatalogById) | **Get** /v1/tools/catalog/{id} | Returns one catalog entry in full: the publisher&#39;s description, its repository and site, every package form with the runtime that launches it, and every hosted endpoint.
[**GetToolsMcpServers**](ToolsAPI.md#GetToolsMcpServers) | **Get** /v1/tools/mcp/servers | Lists the external MCP servers the caller&#39;s org has registered.
[**GetToolsPlugins**](ToolsAPI.md#GetToolsPlugins) | **Get** /v1/tools/plugins | Reports what this deployment actually mounted: every subsystem the composition root declared and whether it is switched on.
[**GetToolsPluginsAuthored**](ToolsAPI.md#GetToolsPluginsAuthored) | **Get** /v1/tools/plugins/authored | Lists the plugins the caller&#39;s org BUILT, newest first, each with the TypeScript as authored.
[**GetToolsSkills**](ToolsAPI.md#GetToolsSkills) | **Get** /v1/tools/skills | Lists the skills the caller&#39;s org can reach — the brand&#39;s embedded catalogue plus the org&#39;s own authored ones — with each one&#39;s activation flag.
[**GetToolsSkillsAuthored**](ToolsAPI.md#GetToolsSkillsAuthored) | **Get** /v1/tools/skills/authored | Lists the caller org&#39;s OWN skills with their SKILL.md bodies.
[**PatchToolsCatalogById**](ToolsAPI.md#PatchToolsCatalogById) | **Patch** /v1/tools/catalog/{id} | Sets what WE say about one catalog entry — hidden, featured, official, logo — and answers with the stored listing.
[**PostToolsCall**](ToolsAPI.md#PostToolsCall) | **Post** /v1/tools/call | Runs one of the caller&#39;s activated tools and answers with its output.
[**PostToolsCatalogSync**](ToolsAPI.md#PostToolsCatalogSync) | **Post** /v1/tools/catalog/sync | Pulls the public MCP registry into our canonical copy and reports what changed.
[**PostToolsMcpServers**](ToolsAPI.md#PostToolsMcpServers) | **Post** /v1/tools/mcp/servers | Gives the caller&#39;s org one more external MCP server, so its tools join the org&#39;s tool plane and the fleet&#39;s MCP server.
[**PostToolsPluginsBuild**](ToolsAPI.md#PostToolsPluginsBuild) | **Post** /v1/tools/plugins/build | Builds and stores one plugin for the caller&#39;s org.
[**PostToolsSkills**](ToolsAPI.md#PostToolsSkills) | **Post** /v1/tools/skills | Adds or revises one of the caller org&#39;s own skills, and answers 201 with the stored record.
[**PutToolsActivation**](ToolsAPI.md#PutToolsActivation) | **Put** /v1/tools/activation | Switches tools on and off for the caller&#39;s org and project, and answers with the resulting activated set.



## DeleteToolsMcpServersById

> DeleteToolsMcpServersById(ctx, id).Execute()

Deregisters one of the caller org's external MCP servers, so its tools leave the registry.



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
	id := "id_example" // string | ID is the server to deregister, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToolsAPI.DeleteToolsMcpServersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.DeleteToolsMcpServersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the server to deregister, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteToolsMcpServersByIdRequest struct via the builder pattern


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


## DeleteToolsPluginsAuthoredById

> PluginDeleted DeleteToolsPluginsAuthoredById(ctx, id).Execute()

Removes one of the caller org's built plugins, so the runtime can no longer load it.



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
	id := "id_example" // string | ID is the plugin to remove, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.DeleteToolsPluginsAuthoredById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.DeleteToolsPluginsAuthoredById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToolsPluginsAuthoredById`: PluginDeleted
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.DeleteToolsPluginsAuthoredById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plugin to remove, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteToolsPluginsAuthoredByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginDeleted**](PluginDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToolsSkillsById

> SkillDeleted DeleteToolsSkillsById(ctx, id).Execute()

Removes one of the caller org's authored skills.



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
	id := "id_example" // string | ID is the skill to remove, from the path. It is the skill's name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.DeleteToolsSkillsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.DeleteToolsSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToolsSkillsById`: SkillDeleted
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.DeleteToolsSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the skill to remove, from the path. It is the skill&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteToolsSkillsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkillDeleted**](SkillDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	q := "q_example" // string | Q matches the name, title or description, case-insensitively. (optional)
	featured := "featured_example" // string | Featured keeps only the listings we put on the front of the shelf, and only when it is exactly the string \"true\". (optional)
	official := "official_example" // string | Official keeps only the vendors' OWN servers — not third-party copies of them — and only when it is exactly the string \"true\". (optional)
	limit := int64(789) // int64 | Limit bounds the page: default 50, maximum 200. A value that is not a positive integer reads as the default. (optional)
	offset := int64(789) // int64 | Offset skips that many listings. (optional)

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
 **limit** | **int64** | Limit bounds the page: default 50, maximum 200. A value that is not a positive integer reads as the default. | 
 **offset** | **int64** | Offset skips that many listings. | 

### Return type

[**McpCatalog**](McpCatalog.md)

### Authorization

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsMcpServers

> McpServerList GetToolsMcpServers(ctx).Execute()

Lists the external MCP servers the caller's org has registered.



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
	resp, r, err := apiClient.ToolsAPI.GetToolsMcpServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsMcpServers`: McpServerList
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsMcpServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsMcpServersRequest struct via the builder pattern


### Return type

[**McpServerList**](McpServerList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsPlugins

> PluginMountList GetToolsPlugins(ctx).All(all).Execute()

Reports what this deployment actually mounted: every subsystem the composition root declared and whether it is switched on.



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
	all := "all_example" // string | All includes the configured-but-disabled subsystems too, but only when it is exactly the string \"true\". Otherwise only the running ones are reported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.GetToolsPlugins(context.Background()).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsPlugins`: PluginMountList
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **string** | All includes the configured-but-disabled subsystems too, but only when it is exactly the string \&quot;true\&quot;. Otherwise only the running ones are reported. | 

### Return type

[**PluginMountList**](PluginMountList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsPluginsAuthored

> AuthoredPluginList GetToolsPluginsAuthored(ctx).Execute()

Lists the plugins the caller's org BUILT, newest first, each with the TypeScript as authored.



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
	resp, r, err := apiClient.ToolsAPI.GetToolsPluginsAuthored(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsPluginsAuthored``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsPluginsAuthored`: AuthoredPluginList
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsPluginsAuthored`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsPluginsAuthoredRequest struct via the builder pattern


### Return type

[**AuthoredPluginList**](AuthoredPluginList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsSkills

> SourceToolList GetToolsSkills(ctx).Activated(activated).Execute()

Lists the skills the caller's org can reach — the brand's embedded catalogue plus the org's own authored ones — with each one's activation flag.



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
	activated := "activated_example" // string | Activated keeps only the tools activated for the caller's org and project, and only when it is exactly the string \"true\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.GetToolsSkills(context.Background()).Activated(activated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsSkills`: SourceToolList
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activated** | **string** | Activated keeps only the tools activated for the caller&#39;s org and project, and only when it is exactly the string \&quot;true\&quot;. | 

### Return type

[**SourceToolList**](SourceToolList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsSkillsAuthored

> AuthoredSkillList GetToolsSkillsAuthored(ctx).Execute()

Lists the caller org's OWN skills with their SKILL.md bodies.



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
	resp, r, err := apiClient.ToolsAPI.GetToolsSkillsAuthored(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsSkillsAuthored``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsSkillsAuthored`: AuthoredSkillList
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsSkillsAuthored`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsSkillsAuthoredRequest struct via the builder pattern


### Return type

[**AuthoredSkillList**](AuthoredSkillList.md)

### Authorization

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolsMcpServers

> MCPServer PostToolsMcpServers(ctx).CreateServerReq(createServerReq).Execute()

Gives the caller's org one more external MCP server, so its tools join the org's tool plane and the fleet's MCP server.



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
	createServerReq := *openapiclient.NewCreateServerReq() // CreateServerReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.PostToolsMcpServers(context.Background()).CreateServerReq(createServerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.PostToolsMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolsMcpServers`: MCPServer
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.PostToolsMcpServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostToolsMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerReq** | [**CreateServerReq**](CreateServerReq.md) |  | 

### Return type

[**MCPServer**](MCPServer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolsPluginsBuild

> BuildOut PostToolsPluginsBuild(ctx).BuildRequest(buildRequest).Execute()

Builds and stores one plugin for the caller's org.



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
	buildRequest := *openapiclient.NewBuildRequest() // BuildRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.PostToolsPluginsBuild(context.Background()).BuildRequest(buildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.PostToolsPluginsBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolsPluginsBuild`: BuildOut
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.PostToolsPluginsBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostToolsPluginsBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildRequest** | [**BuildRequest**](BuildRequest.md) |  | 

### Return type

[**BuildOut**](BuildOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolsSkills

> SkillWritten PostToolsSkills(ctx).SkillIn(skillIn).Execute()

Adds or revises one of the caller org's own skills, and answers 201 with the stored record.



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
	skillIn := *openapiclient.NewSkillIn() // SkillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.PostToolsSkills(context.Background()).SkillIn(skillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.PostToolsSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolsSkills`: SkillWritten
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.PostToolsSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostToolsSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skillIn** | [**SkillIn**](SkillIn.md) |  | 

### Return type

[**SkillWritten**](SkillWritten.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

