# \EnvironmentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEnvironmentByRepo**](EnvironmentAPI.md#DeleteEnvironmentByRepo) | **Delete** /v1/environment/{repo} | Forgets one codebase&#39;s environment, and its secrets with it.
[**DeleteEnvironmentByRepoSecretsByName**](EnvironmentAPI.md#DeleteEnvironmentByRepoSecretsByName) | **Delete** /v1/environment/{repo}/secrets/{name} | Removes one secret from a codebase.
[**GetEnvironment**](EnvironmentAPI.md#GetEnvironment) | **Get** /v1/environment | Lists every codebase in your org that has an environment.
[**GetEnvironmentByRepo**](EnvironmentAPI.md#GetEnvironmentByRepo) | **Get** /v1/environment/{repo} | Reads one codebase&#39;s environment.
[**PutEnvironmentByRepo**](EnvironmentAPI.md#PutEnvironmentByRepo) | **Put** /v1/environment/{repo} | Saves one codebase&#39;s install and start scripts.
[**PutEnvironmentByRepoSecretsByName**](EnvironmentAPI.md#PutEnvironmentByRepoSecretsByName) | **Put** /v1/environment/{repo}/secrets/{name} | Sets one secret on a codebase.



## DeleteEnvironmentByRepo

> DeleteEnvironmentByRepo(ctx, repo).Execute()

Forgets one codebase's environment, and its secrets with it.



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
	repo := "repo_example" // string | Repo is the repository's name in the caller's org — a name, never owner/name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentAPI.DeleteEnvironmentByRepo(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.DeleteEnvironmentByRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s name in the caller&#39;s org — a name, never owner/name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentByRepoRequest struct via the builder pattern


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


## DeleteEnvironmentByRepoSecretsByName

> EnvironmentEnvironment DeleteEnvironmentByRepoSecretsByName(ctx, repo, name).Execute()

Removes one secret from a codebase.



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
	repo := "repo_example" // string | Repo is the repository's name in the caller's org.
	name := "name_example" // string | Name is the secret's environment variable name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.DeleteEnvironmentByRepoSecretsByName(context.Background(), repo, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.DeleteEnvironmentByRepoSecretsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEnvironmentByRepoSecretsByName`: EnvironmentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.DeleteEnvironmentByRepoSecretsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s name in the caller&#39;s org. | 
**name** | **string** | Name is the secret&#39;s environment variable name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentByRepoSecretsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvironmentEnvironment**](EnvironmentEnvironment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironment

> EnvironmentEnvList GetEnvironment(ctx).Execute()

Lists every codebase in your org that has an environment.



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
	resp, r, err := apiClient.EnvironmentAPI.GetEnvironment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironment`: EnvironmentEnvList
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


### Return type

[**EnvironmentEnvList**](EnvironmentEnvList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentByRepo

> EnvironmentEnvironment GetEnvironmentByRepo(ctx, repo).Execute()

Reads one codebase's environment.



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
	repo := "repo_example" // string | Repo is the repository's name in the caller's org — a name, never owner/name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.GetEnvironmentByRepo(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetEnvironmentByRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentByRepo`: EnvironmentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetEnvironmentByRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s name in the caller&#39;s org — a name, never owner/name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentByRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentEnvironment**](EnvironmentEnvironment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEnvironmentByRepo

> EnvironmentEnvironment PutEnvironmentByRepo(ctx, repo).EnvironmentEnvSave(environmentEnvSave).Execute()

Saves one codebase's install and start scripts.



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
	repo := "repo_example" // string | Repo is the repository's name in the caller's org.
	environmentEnvSave := *openapiclient.NewEnvironmentEnvSave() // EnvironmentEnvSave | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.PutEnvironmentByRepo(context.Background(), repo).EnvironmentEnvSave(environmentEnvSave).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.PutEnvironmentByRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutEnvironmentByRepo`: EnvironmentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.PutEnvironmentByRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s name in the caller&#39;s org. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEnvironmentByRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEnvSave** | [**EnvironmentEnvSave**](EnvironmentEnvSave.md) |  | 

### Return type

[**EnvironmentEnvironment**](EnvironmentEnvironment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEnvironmentByRepoSecretsByName

> EnvironmentEnvironment PutEnvironmentByRepoSecretsByName(ctx, repo, name).EnvironmentEnvSecret(environmentEnvSecret).Execute()

Sets one secret on a codebase.



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
	repo := "repo_example" // string | Repo is the repository's name in the caller's org.
	name := "name_example" // string | Name is the environment variable the run exports the value under.
	environmentEnvSecret := *openapiclient.NewEnvironmentEnvSecret() // EnvironmentEnvSecret | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.PutEnvironmentByRepoSecretsByName(context.Background(), repo, name).EnvironmentEnvSecret(environmentEnvSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.PutEnvironmentByRepoSecretsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutEnvironmentByRepoSecretsByName`: EnvironmentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.PutEnvironmentByRepoSecretsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s name in the caller&#39;s org. | 
**name** | **string** | Name is the environment variable the run exports the value under. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEnvironmentByRepoSecretsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentEnvSecret** | [**EnvironmentEnvSecret**](EnvironmentEnvSecret.md) |  | 

### Return type

[**EnvironmentEnvironment**](EnvironmentEnvironment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

