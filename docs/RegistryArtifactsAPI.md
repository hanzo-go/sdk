# \RegistryArtifactsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryCreateTag**](RegistryArtifactsAPI.md#RegistryCreateTag) | **Post** /v1/registry/projects/{name}/repositories/{repo}/artifacts/{digest}/tags/{tag} | Create tag
[**RegistryDeleteArtifact**](RegistryArtifactsAPI.md#RegistryDeleteArtifact) | **Delete** /v1/registry/projects/{name}/repositories/{repo}/artifacts/{digest} | Delete artifact
[**RegistryDeleteTag**](RegistryArtifactsAPI.md#RegistryDeleteTag) | **Delete** /v1/registry/projects/{name}/repositories/{repo}/artifacts/{digest}/tags/{tag} | Delete tag
[**RegistryGetArtifact**](RegistryArtifactsAPI.md#RegistryGetArtifact) | **Get** /v1/registry/projects/{name}/repositories/{repo}/artifacts/{digest} | Get artifact
[**RegistryListArtifacts**](RegistryArtifactsAPI.md#RegistryListArtifacts) | **Get** /v1/registry/projects/{name}/repositories/{repo}/artifacts | List artifacts
[**RegistryListTags**](RegistryArtifactsAPI.md#RegistryListTags) | **Get** /v1/registry/projects/{name}/repositories/{repo}/artifacts/{digest}/tags | List tags



## RegistryCreateTag

> map[string]interface{} RegistryCreateTag(ctx, name, repo, digest, tag).Execute()

Create tag

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
	digest := "digest_example" // string | 
	tag := "tag_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryArtifactsAPI.RegistryCreateTag(context.Background(), name, repo, digest, tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryArtifactsAPI.RegistryCreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryCreateTag`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryArtifactsAPI.RegistryCreateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 
**digest** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryCreateTagRequest struct via the builder pattern


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


## RegistryDeleteArtifact

> map[string]interface{} RegistryDeleteArtifact(ctx, name, repo, digest).Execute()

Delete artifact

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
	digest := "digest_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryArtifactsAPI.RegistryDeleteArtifact(context.Background(), name, repo, digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryArtifactsAPI.RegistryDeleteArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryDeleteArtifact`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryArtifactsAPI.RegistryDeleteArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 
**digest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryDeleteArtifactRequest struct via the builder pattern


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


## RegistryDeleteTag

> map[string]interface{} RegistryDeleteTag(ctx, name, repo, digest, tag).Execute()

Delete tag

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
	digest := "digest_example" // string | 
	tag := "tag_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryArtifactsAPI.RegistryDeleteTag(context.Background(), name, repo, digest, tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryArtifactsAPI.RegistryDeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryDeleteTag`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryArtifactsAPI.RegistryDeleteTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 
**digest** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryDeleteTagRequest struct via the builder pattern


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


## RegistryGetArtifact

> RegistryArtifact RegistryGetArtifact(ctx, name, repo, digest).WithScanOverview(withScanOverview).Execute()

Get artifact

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
	digest := "digest_example" // string | Artifact digest (sha256:...)
	withScanOverview := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryArtifactsAPI.RegistryGetArtifact(context.Background(), name, repo, digest).WithScanOverview(withScanOverview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryArtifactsAPI.RegistryGetArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryGetArtifact`: RegistryArtifact
	fmt.Fprintf(os.Stdout, "Response from `RegistryArtifactsAPI.RegistryGetArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 
**digest** | **string** | Artifact digest (sha256:...) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryGetArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **withScanOverview** | **bool** |  | [default to false]

### Return type

[**RegistryArtifact**](RegistryArtifact.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryListArtifacts

> []RegistryArtifact RegistryListArtifacts(ctx, name, repo).Type_(type_).WithTag(withTag).WithScanOverview(withScanOverview).Page(page).PageSize(pageSize).Execute()

List artifacts

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
	type_ := "type__example" // string | Filter by artifact type (optional)
	withTag := true // bool |  (optional) (default to true)
	withScanOverview := true // bool |  (optional) (default to false)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryArtifactsAPI.RegistryListArtifacts(context.Background(), name, repo).Type_(type_).WithTag(withTag).WithScanOverview(withScanOverview).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryArtifactsAPI.RegistryListArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryListArtifacts`: []RegistryArtifact
	fmt.Fprintf(os.Stdout, "Response from `RegistryArtifactsAPI.RegistryListArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryListArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | Filter by artifact type | 
 **withTag** | **bool** |  | [default to true]
 **withScanOverview** | **bool** |  | [default to false]
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]

### Return type

[**[]RegistryArtifact**](RegistryArtifact.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryListTags

> []RegistryTag RegistryListTags(ctx, name, repo, digest).Page(page).PageSize(pageSize).Execute()

List tags

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
	digest := "digest_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryArtifactsAPI.RegistryListTags(context.Background(), name, repo, digest).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryArtifactsAPI.RegistryListTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryListTags`: []RegistryTag
	fmt.Fprintf(os.Stdout, "Response from `RegistryArtifactsAPI.RegistryListTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**repo** | **string** |  | 
**digest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryListTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]

### Return type

[**[]RegistryTag**](RegistryTag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

