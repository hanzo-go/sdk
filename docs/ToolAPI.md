# \ToolAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToolMcpServersById**](ToolAPI.md#DeleteToolMcpServersById) | **Delete** /v1/tool/mcp/servers/{id} | Deregisters one of the caller org&#39;s external MCP servers, so its tools leave the registry.
[**DeleteToolPluginsAuthoredById**](ToolAPI.md#DeleteToolPluginsAuthoredById) | **Delete** /v1/tool/plugins/authored/{id} | Removes one of the caller org&#39;s built plugins, so the runtime can no longer load it.
[**DeleteToolSkillsById**](ToolAPI.md#DeleteToolSkillsById) | **Delete** /v1/tool/skills/{id} | Removes one of the caller org&#39;s authored skills.
[**GetTool**](ToolAPI.md#GetTool) | **Get** /v1/tool | Lists every tool the caller&#39;s org and project can reach, from every source, each flagged with whether it is activated.
[**GetToolActivation**](ToolAPI.md#GetToolActivation) | **Get** /v1/tool/activation | Reports which tools are switched on for the caller&#39;s org and project.
[**GetToolCatalog**](ToolAPI.md#GetToolCatalog) | **Get** /v1/tool/catalog | Lists the MCP servers the public registries publish, as we hold them: our canonical copy of registry.modelcontextprotocol.io, plus what we decided about each entry.
[**GetToolCatalogById**](ToolAPI.md#GetToolCatalogById) | **Get** /v1/tool/catalog/{id} | Returns one catalog entry in full: the publisher&#39;s description, its repository and site, every package form with the runtime that launches it, and every hosted endpoint.
[**GetToolKit**](ToolAPI.md#GetToolKit) | **Get** /v1/tool/kit | Answers the caller&#39;s kit: what their coding runs carry of the org&#39;s tool plane.
[**GetToolMcpServers**](ToolAPI.md#GetToolMcpServers) | **Get** /v1/tool/mcp/servers | Lists the external MCP servers the caller&#39;s org has registered.
[**GetToolPlugins**](ToolAPI.md#GetToolPlugins) | **Get** /v1/tool/plugins | Reports what this deployment actually mounted: every subsystem the composition root declared and whether it is switched on.
[**GetToolPluginsAuthored**](ToolAPI.md#GetToolPluginsAuthored) | **Get** /v1/tool/plugins/authored | Lists the plugins the caller&#39;s org BUILT, newest first, each with the TypeScript as authored.
[**GetToolSkills**](ToolAPI.md#GetToolSkills) | **Get** /v1/tool/skills | Lists the skills the caller&#39;s org can reach — the brand&#39;s embedded catalogue plus the org&#39;s own authored ones — with each one&#39;s activation flag.
[**GetToolSkillsAuthored**](ToolAPI.md#GetToolSkillsAuthored) | **Get** /v1/tool/skills/authored | Lists the caller org&#39;s OWN skills with their SKILL.md bodies.
[**PostToolCall**](ToolAPI.md#PostToolCall) | **Post** /v1/tool/call | Runs one of the caller&#39;s activated tools and answers with its output.
[**PostToolMcpServers**](ToolAPI.md#PostToolMcpServers) | **Post** /v1/tool/mcp/servers | Gives the caller&#39;s org one more external MCP server, so its tools join the org&#39;s tool plane and the fleet&#39;s MCP server.
[**PostToolPluginsBuild**](ToolAPI.md#PostToolPluginsBuild) | **Post** /v1/tool/plugins/build | Builds and stores one plugin for the caller&#39;s org.
[**PostToolSkills**](ToolAPI.md#PostToolSkills) | **Post** /v1/tool/skills | Adds or revises one of the caller org&#39;s own skills, and answers 201 with the stored record.
[**PutToolActivation**](ToolAPI.md#PutToolActivation) | **Put** /v1/tool/activation | Switches tools on and off for the caller&#39;s org and project, and answers with the resulting activated set.
[**PutToolKit**](ToolAPI.md#PutToolKit) | **Put** /v1/tool/kit | Mutes and unmutes names of the org&#39;s kit for the caller&#39;s own runs, and answers the caller&#39;s kit after.



## DeleteToolMcpServersById

> DeleteToolMcpServersById(ctx, id).Execute()

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
	r, err := apiClient.ToolAPI.DeleteToolMcpServersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.DeleteToolMcpServersById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteToolMcpServersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteToolPluginsAuthoredById

> ToolPluginDeleted DeleteToolPluginsAuthoredById(ctx, id).Execute()

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
	resp, r, err := apiClient.ToolAPI.DeleteToolPluginsAuthoredById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.DeleteToolPluginsAuthoredById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToolPluginsAuthoredById`: ToolPluginDeleted
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.DeleteToolPluginsAuthoredById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plugin to remove, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteToolPluginsAuthoredByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToolPluginDeleted**](ToolPluginDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToolSkillsById

> ToolSkillDeleted DeleteToolSkillsById(ctx, id).Execute()

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
	resp, r, err := apiClient.ToolAPI.DeleteToolSkillsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.DeleteToolSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToolSkillsById`: ToolSkillDeleted
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.DeleteToolSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the skill to remove, from the path. It is the skill&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteToolSkillsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToolSkillDeleted**](ToolSkillDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTool

> ToolToolList GetTool(ctx).Source(source).Activated(activated).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetTool(context.Background()).Source(source).Activated(activated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTool`: ToolToolList
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetTool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | Source keeps only tools from one source — connector, function, zap-service, agent, skill or mcp. Empty keeps every source. | 
 **activated** | **string** | Activated keeps only the tools activated for the caller&#39;s org and project, and only when it is exactly the string \&quot;true\&quot;. | 

### Return type

[**ToolToolList**](ToolToolList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolActivation

> ToolActivationSet GetToolActivation(ctx).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetToolActivation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolActivation`: ToolActivationSet
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolActivation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolActivationRequest struct via the builder pattern


### Return type

[**ToolActivationSet**](ToolActivationSet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolCatalog

> ToolMcpCatalog GetToolCatalog(ctx).Q(q).Featured(featured).Official(official).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetToolCatalog(context.Background()).Q(q).Featured(featured).Official(official).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolCatalog`: ToolMcpCatalog
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetToolCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q matches the name, title or description, case-insensitively. | 
 **featured** | **string** | Featured keeps only the listings we put on the front of the shelf, and only when it is exactly the string \&quot;true\&quot;. | 
 **official** | **string** | Official keeps only the vendors&#39; OWN servers — not third-party copies of them — and only when it is exactly the string \&quot;true\&quot;. | 
 **limit** | **int64** | Limit bounds the page: default 50, maximum 200. A value that is not a positive integer reads as the default. | 
 **offset** | **int64** | Offset skips that many listings. | 

### Return type

[**ToolMcpCatalog**](ToolMcpCatalog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolCatalogById

> ToolMCPListing GetToolCatalogById(ctx, id).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetToolCatalogById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolCatalogById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolCatalogById`: ToolMCPListing
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolCatalogById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the listing, from the path. It is the publisher&#39;s reverse-DNS name with its one slash written as an underscore — \&quot;com.stripe_mcp\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolCatalogByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToolMCPListing**](ToolMCPListing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolKit

> ToolKit GetToolKit(ctx).Execute()

Answers the caller's kit: what their coding runs carry of the org's tool plane.



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
	resp, r, err := apiClient.ToolAPI.GetToolKit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolKit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolKit`: ToolKit
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolKit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolKitRequest struct via the builder pattern


### Return type

[**ToolKit**](ToolKit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolMcpServers

> ToolMcpServerList GetToolMcpServers(ctx).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetToolMcpServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolMcpServers`: ToolMcpServerList
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolMcpServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolMcpServersRequest struct via the builder pattern


### Return type

[**ToolMcpServerList**](ToolMcpServerList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolPlugins

> ToolPluginMountList GetToolPlugins(ctx).All(all).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetToolPlugins(context.Background()).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolPlugins`: ToolPluginMountList
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetToolPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **string** | All includes the configured-but-disabled subsystems too, but only when it is exactly the string \&quot;true\&quot;. Otherwise only the running ones are reported. | 

### Return type

[**ToolPluginMountList**](ToolPluginMountList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolPluginsAuthored

> ToolAuthoredPluginList GetToolPluginsAuthored(ctx).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetToolPluginsAuthored(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolPluginsAuthored``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolPluginsAuthored`: ToolAuthoredPluginList
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolPluginsAuthored`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolPluginsAuthoredRequest struct via the builder pattern


### Return type

[**ToolAuthoredPluginList**](ToolAuthoredPluginList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolSkills

> ToolSourceToolList GetToolSkills(ctx).Activated(activated).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetToolSkills(context.Background()).Activated(activated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolSkills`: ToolSourceToolList
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetToolSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activated** | **string** | Activated keeps only the tools activated for the caller&#39;s org and project, and only when it is exactly the string \&quot;true\&quot;. | 

### Return type

[**ToolSourceToolList**](ToolSourceToolList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolSkillsAuthored

> ToolAuthoredSkillList GetToolSkillsAuthored(ctx).Execute()

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
	resp, r, err := apiClient.ToolAPI.GetToolSkillsAuthored(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.GetToolSkillsAuthored``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolSkillsAuthored`: ToolAuthoredSkillList
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.GetToolSkillsAuthored`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolSkillsAuthoredRequest struct via the builder pattern


### Return type

[**ToolAuthoredSkillList**](ToolAuthoredSkillList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolCall

> ToolToolResult PostToolCall(ctx).ToolToolCall(toolToolCall).Execute()

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
	toolToolCall := *openapiclient.NewToolToolCall() // ToolToolCall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolAPI.PostToolCall(context.Background()).ToolToolCall(toolToolCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.PostToolCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolCall`: ToolToolResult
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.PostToolCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostToolCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolToolCall** | [**ToolToolCall**](ToolToolCall.md) |  | 

### Return type

[**ToolToolResult**](ToolToolResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolMcpServers

> ToolMCPServer PostToolMcpServers(ctx).ToolCreateServerReq(toolCreateServerReq).Execute()

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
	toolCreateServerReq := *openapiclient.NewToolCreateServerReq() // ToolCreateServerReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolAPI.PostToolMcpServers(context.Background()).ToolCreateServerReq(toolCreateServerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.PostToolMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolMcpServers`: ToolMCPServer
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.PostToolMcpServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostToolMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolCreateServerReq** | [**ToolCreateServerReq**](ToolCreateServerReq.md) |  | 

### Return type

[**ToolMCPServer**](ToolMCPServer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolPluginsBuild

> ToolBuildOut PostToolPluginsBuild(ctx).ToolBuildRequest(toolBuildRequest).Execute()

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
	toolBuildRequest := *openapiclient.NewToolBuildRequest() // ToolBuildRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolAPI.PostToolPluginsBuild(context.Background()).ToolBuildRequest(toolBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.PostToolPluginsBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolPluginsBuild`: ToolBuildOut
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.PostToolPluginsBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostToolPluginsBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolBuildRequest** | [**ToolBuildRequest**](ToolBuildRequest.md) |  | 

### Return type

[**ToolBuildOut**](ToolBuildOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToolSkills

> ToolSkillWritten PostToolSkills(ctx).ToolSkillIn(toolSkillIn).Execute()

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
	toolSkillIn := *openapiclient.NewToolSkillIn() // ToolSkillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolAPI.PostToolSkills(context.Background()).ToolSkillIn(toolSkillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.PostToolSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToolSkills`: ToolSkillWritten
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.PostToolSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostToolSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolSkillIn** | [**ToolSkillIn**](ToolSkillIn.md) |  | 

### Return type

[**ToolSkillWritten**](ToolSkillWritten.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutToolActivation

> ToolActivationSet PutToolActivation(ctx).ToolActivationReq(toolActivationReq).Execute()

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
	toolActivationReq := *openapiclient.NewToolActivationReq() // ToolActivationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolAPI.PutToolActivation(context.Background()).ToolActivationReq(toolActivationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.PutToolActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutToolActivation`: ToolActivationSet
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.PutToolActivation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutToolActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolActivationReq** | [**ToolActivationReq**](ToolActivationReq.md) |  | 

### Return type

[**ToolActivationSet**](ToolActivationSet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutToolKit

> ToolKit PutToolKit(ctx).ToolMuteReq(toolMuteReq).Execute()

Mutes and unmutes names of the org's kit for the caller's own runs, and answers the caller's kit after.



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
	toolMuteReq := *openapiclient.NewToolMuteReq() // ToolMuteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolAPI.PutToolKit(context.Background()).ToolMuteReq(toolMuteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolAPI.PutToolKit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutToolKit`: ToolKit
	fmt.Fprintf(os.Stdout, "Response from `ToolAPI.PutToolKit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutToolKitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolMuteReq** | [**ToolMuteReq**](ToolMuteReq.md) |  | 

### Return type

[**ToolKit**](ToolKit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

