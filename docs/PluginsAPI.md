# \PluginsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePluginsAuthoredById**](PluginsAPI.md#DeletePluginsAuthoredById) | **Delete** /v1/plugins/authored/{id} | Removes one of the caller org&#39;s built plugins, so the runtime can no longer load it.
[**GetPlugins**](PluginsAPI.md#GetPlugins) | **Get** /v1/plugins | Reports what this deployment actually mounted: every subsystem the composition root declared and whether it is switched on.
[**GetPluginsAuthored**](PluginsAPI.md#GetPluginsAuthored) | **Get** /v1/plugins/authored | Lists the plugins the caller&#39;s org BUILT, newest first, each with the TypeScript as authored.
[**PostPluginsBuild**](PluginsAPI.md#PostPluginsBuild) | **Post** /v1/plugins/build | Build a plugin for your org from TypeScript, or from an API spec a model writes it from



## DeletePluginsAuthoredById

> PluginDeleted DeletePluginsAuthoredById(ctx, id).Execute()

Removes one of the caller org's built plugins, so the runtime can no longer load it.



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
	id := "id_example" // string | ID is the plugin to remove, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.DeletePluginsAuthoredById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.DeletePluginsAuthoredById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePluginsAuthoredById`: PluginDeleted
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.DeletePluginsAuthoredById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plugin to remove, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginsAuthoredByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginDeleted**](PluginDeleted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlugins

> PluginMountList GetPlugins(ctx).All(all).Execute()

Reports what this deployment actually mounted: every subsystem the composition root declared and whether it is switched on.



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
	all := "all_example" // string | All includes the configured-but-disabled subsystems too, but only when it is exactly the string \"true\". Otherwise only the running ones are reported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.GetPlugins(context.Background()).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlugins`: PluginMountList
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **string** | All includes the configured-but-disabled subsystems too, but only when it is exactly the string \&quot;true\&quot;. Otherwise only the running ones are reported. | 

### Return type

[**PluginMountList**](PluginMountList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginsAuthored

> AuthoredPluginList GetPluginsAuthored(ctx).Execute()

Lists the plugins the caller's org BUILT, newest first, each with the TypeScript as authored.



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
	resp, r, err := apiClient.PluginsAPI.GetPluginsAuthored(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPluginsAuthored``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPluginsAuthored`: AuthoredPluginList
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPluginsAuthored`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginsAuthoredRequest struct via the builder pattern


### Return type

[**AuthoredPluginList**](AuthoredPluginList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPluginsBuild

> BuildOut PostPluginsBuild(ctx).BuildRequest(buildRequest).Execute()

Build a plugin for your org from TypeScript, or from an API spec a model writes it from



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
	buildRequest := *openapiclient.NewBuildRequest() // BuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PostPluginsBuild(context.Background()).BuildRequest(buildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PostPluginsBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPluginsBuild`: BuildOut
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PostPluginsBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPluginsBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildRequest** | [**BuildRequest**](BuildRequest.md) |  | 

### Return type

[**BuildOut**](BuildOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

