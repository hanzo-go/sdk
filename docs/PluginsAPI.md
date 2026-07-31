# \PluginsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1PluginsAuthoredId**](PluginsAPI.md#CloudDeleteV1PluginsAuthoredId) | **Delete** /v1/plugins/authored/{id} | DeleteAuthoredPlugin removes one of the caller org&#39;s built plugins, so the runtime can no longer load it.
[**CloudGetV1Plugins**](PluginsAPI.md#CloudGetV1Plugins) | **Get** /v1/plugins | ListPlugins reports what this deployment actually mounted: every subsystem the composition root declared and whether it is switched on.
[**CloudGetV1PluginsAuthored**](PluginsAPI.md#CloudGetV1PluginsAuthored) | **Get** /v1/plugins/authored | ListAuthoredPlugins lists the plugins the caller&#39;s org BUILT, newest first, each with the TypeScript as authored.
[**CloudPostV1PluginsBuild**](PluginsAPI.md#CloudPostV1PluginsBuild) | **Post** /v1/plugins/build | 



## CloudDeleteV1PluginsAuthoredId

> CloudPluginDeleted CloudDeleteV1PluginsAuthoredId(ctx, id).Execute()

DeleteAuthoredPlugin removes one of the caller org's built plugins, so the runtime can no longer load it.



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
	resp, r, err := apiClient.PluginsAPI.CloudDeleteV1PluginsAuthoredId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.CloudDeleteV1PluginsAuthoredId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1PluginsAuthoredId`: CloudPluginDeleted
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.CloudDeleteV1PluginsAuthoredId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plugin to remove, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1PluginsAuthoredIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudPluginDeleted**](CloudPluginDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Plugins

> CloudPluginMountList CloudGetV1Plugins(ctx).All(all).Execute()

ListPlugins reports what this deployment actually mounted: every subsystem the composition root declared and whether it is switched on.



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
	resp, r, err := apiClient.PluginsAPI.CloudGetV1Plugins(context.Background()).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.CloudGetV1Plugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Plugins`: CloudPluginMountList
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.CloudGetV1Plugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **string** | All includes the configured-but-disabled subsystems too, but only when it is exactly the string \&quot;true\&quot;. Otherwise only the running ones are reported. | 

### Return type

[**CloudPluginMountList**](CloudPluginMountList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PluginsAuthored

> CloudAuthoredPluginList CloudGetV1PluginsAuthored(ctx).Execute()

ListAuthoredPlugins lists the plugins the caller's org BUILT, newest first, each with the TypeScript as authored.



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
	resp, r, err := apiClient.PluginsAPI.CloudGetV1PluginsAuthored(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.CloudGetV1PluginsAuthored``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PluginsAuthored`: CloudAuthoredPluginList
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.CloudGetV1PluginsAuthored`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PluginsAuthoredRequest struct via the builder pattern


### Return type

[**CloudAuthoredPluginList**](CloudAuthoredPluginList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PluginsBuild

> CloudBuildOut CloudPostV1PluginsBuild(ctx).CloudBuildRequest(cloudBuildRequest).Execute()



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
	cloudBuildRequest := *openapiclient.NewCloudBuildRequest() // CloudBuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.CloudPostV1PluginsBuild(context.Background()).CloudBuildRequest(cloudBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.CloudPostV1PluginsBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PluginsBuild`: CloudBuildOut
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.CloudPostV1PluginsBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PluginsBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudBuildRequest** | [**CloudBuildRequest**](CloudBuildRequest.md) |  | 

### Return type

[**CloudBuildOut**](CloudBuildOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

