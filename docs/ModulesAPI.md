# \ModulesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FrameworkGetModule**](ModulesAPI.md#FrameworkGetModule) | **Get** /v1/framework/modules/{module} | Inspect a module and its install state in this org
[**FrameworkInstallModule**](ModulesAPI.md#FrameworkInstallModule) | **Post** /v1/framework/modules/{module}/install | Install a module&#39;s DocType fixtures into this org (idempotent)
[**FrameworkListModules**](ModulesAPI.md#FrameworkListModules) | **Get** /v1/framework/modules | List registered app-lane modules



## FrameworkGetModule

> FrameworkGetModule200Response FrameworkGetModule(ctx, module).Execute()

Inspect a module and its install state in this org

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
	module := "module_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModulesAPI.FrameworkGetModule(context.Background(), module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModulesAPI.FrameworkGetModule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkGetModule`: FrameworkGetModule200Response
	fmt.Fprintf(os.Stdout, "Response from `ModulesAPI.FrameworkGetModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**module** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkGetModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FrameworkGetModule200Response**](FrameworkGetModule200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkInstallModule

> FrameworkInstallModule200Response FrameworkInstallModule(ctx, module).Execute()

Install a module's DocType fixtures into this org (idempotent)

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
	module := "module_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModulesAPI.FrameworkInstallModule(context.Background(), module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModulesAPI.FrameworkInstallModule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkInstallModule`: FrameworkInstallModule200Response
	fmt.Fprintf(os.Stdout, "Response from `ModulesAPI.FrameworkInstallModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**module** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkInstallModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FrameworkInstallModule200Response**](FrameworkInstallModule200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkListModules

> FrameworkListModules200Response FrameworkListModules(ctx).Execute()

List registered app-lane modules

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
	resp, r, err := apiClient.ModulesAPI.FrameworkListModules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModulesAPI.FrameworkListModules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkListModules`: FrameworkListModules200Response
	fmt.Fprintf(os.Stdout, "Response from `ModulesAPI.FrameworkListModules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkListModulesRequest struct via the builder pattern


### Return type

[**FrameworkListModules200Response**](FrameworkListModules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

