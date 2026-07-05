# \RegistryRepositoriesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryDeleteRepository**](RegistryRepositoriesAPI.md#RegistryDeleteRepository) | **Delete** /v1/registry/projects/{name}/repositories/{repo} | Delete repository
[**RegistryGetRepository**](RegistryRepositoriesAPI.md#RegistryGetRepository) | **Get** /v1/registry/projects/{name}/repositories/{repo} | Get repository
[**RegistryListRepositories**](RegistryRepositoriesAPI.md#RegistryListRepositories) | **Get** /v1/registry/projects/{name}/repositories | List repositories



## RegistryDeleteRepository

> map[string]interface{} RegistryDeleteRepository(ctx, name, repo).Execute()

Delete repository

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
	name := "name_example" // string | 
	repo := "repo_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryRepositoriesAPI.RegistryDeleteRepository(context.Background(), name, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryRepositoriesAPI.RegistryDeleteRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryDeleteRepository`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryRepositoriesAPI.RegistryDeleteRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryDeleteRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryGetRepository

> RegistryRepository RegistryGetRepository(ctx, name, repo).Execute()

Get repository

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
	name := "name_example" // string | 
	repo := "repo_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryRepositoriesAPI.RegistryGetRepository(context.Background(), name, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryRepositoriesAPI.RegistryGetRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryGetRepository`: RegistryRepository
	fmt.Fprintf(os.Stdout, "Response from `RegistryRepositoriesAPI.RegistryGetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RegistryRepository**](RegistryRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryListRepositories

> []RegistryRepository RegistryListRepositories(ctx, name).Q(q).Page(page).PageSize(pageSize).Sort(sort).Execute()

List repositories

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
	name := "name_example" // string | 
	q := "q_example" // string | Search query (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	sort := "sort_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryRepositoriesAPI.RegistryListRepositories(context.Background(), name).Q(q).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryRepositoriesAPI.RegistryListRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryListRepositories`: []RegistryRepository
	fmt.Fprintf(os.Stdout, "Response from `RegistryRepositoriesAPI.RegistryListRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryListRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Search query | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sort** | **string** |  | 

### Return type

[**[]RegistryRepository**](RegistryRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

