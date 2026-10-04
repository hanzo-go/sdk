# \RegistryAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegistryImages**](RegistryAPI.md#GetRegistryImages) | **Get** /v1/registry/images | Lists the org&#39;s container repositories, read live from the OCI catalog and filtered server-side to the org&#39;s namespace — the page can only ever hold the caller&#39;s own images.
[**GetRegistryPackages**](RegistryAPI.md#GetRegistryPackages) | **Get** /v1/registry/packages | Lists the org&#39;s npm packages — &#x60;&lt;org&gt;&#x60; and &#x60;@&lt;org&gt;/…&#x60; — from the npm registry&#39;s search index, optionally narrowed by a query within that scope.
[**GetRegistryProjects**](RegistryAPI.md#GetRegistryProjects) | **Get** /v1/registry/projects | Lists the namespaces the caller can see with what each holds: the org&#39;s slug, its repository count on the OCI catalog, and its package count on the npm registry.
[**GetRegistryStatus**](RegistryAPI.md#GetRegistryStatus) | **Get** /v1/registry/status | Reports whether the OCI and npm registries are reachable and, when the OCI half is auth-gated, which token realm its challenge advertises — an honest lens for \&quot;is the registry plane up\&quot;, never a fabricated ok.
[**GetRegistryTags**](RegistryAPI.md#GetRegistryTags) | **Get** /v1/registry/tags | Lists one org-owned repository&#39;s tags, read live from the OCI registry.
[**PostRegistryToken**](RegistryAPI.md#PostRegistryToken) | **Post** /v1/registry/token | Mints a short-lived, pull-only registry token for exactly one of the org&#39;s images, through the same IAM realm the docker CLI authenticates against.



## GetRegistryImages

> RegistryRegistryImageList GetRegistryImages(ctx).Execute()

Lists the org's container repositories, read live from the OCI catalog and filtered server-side to the org's namespace — the page can only ever hold the caller's own images.



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
	resp, r, err := apiClient.RegistryAPI.GetRegistryImages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.GetRegistryImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistryImages`: RegistryRegistryImageList
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.GetRegistryImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryImagesRequest struct via the builder pattern


### Return type

[**RegistryRegistryImageList**](RegistryRegistryImageList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryPackages

> RegistryRegistryPackageList GetRegistryPackages(ctx).Query(query).Execute()

Lists the org's npm packages — `<org>` and `@<org>/…` — from the npm registry's search index, optionally narrowed by a query within that scope.



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
	query := "query_example" // string | Query narrows the listing within the org's scope when present; the org boundary itself is never widened by it. It rides the query string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.GetRegistryPackages(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.GetRegistryPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistryPackages`: RegistryRegistryPackageList
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.GetRegistryPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query narrows the listing within the org&#39;s scope when present; the org boundary itself is never widened by it. It rides the query string. | 

### Return type

[**RegistryRegistryPackageList**](RegistryRegistryPackageList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryProjects

> RegistryRegistryProjectList GetRegistryProjects(ctx).Execute()

Lists the namespaces the caller can see with what each holds: the org's slug, its repository count on the OCI catalog, and its package count on the npm registry.



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
	resp, r, err := apiClient.RegistryAPI.GetRegistryProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.GetRegistryProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistryProjects`: RegistryRegistryProjectList
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.GetRegistryProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryProjectsRequest struct via the builder pattern


### Return type

[**RegistryRegistryProjectList**](RegistryRegistryProjectList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryStatus

> RegistryRegistryStatus GetRegistryStatus(ctx).Execute()

Reports whether the OCI and npm registries are reachable and, when the OCI half is auth-gated, which token realm its challenge advertises — an honest lens for \"is the registry plane up\", never a fabricated ok.



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
	resp, r, err := apiClient.RegistryAPI.GetRegistryStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.GetRegistryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistryStatus`: RegistryRegistryStatus
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.GetRegistryStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryStatusRequest struct via the builder pattern


### Return type

[**RegistryRegistryStatus**](RegistryRegistryStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryTags

> RegistryRegistryTagList GetRegistryTags(ctx).Image(image).Execute()

Lists one org-owned repository's tags, read live from the OCI registry.



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
	image := "image_example" // string | Image is the repository name inside the org's namespace, as returned by the images op. It rides the query string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.GetRegistryTags(context.Background()).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.GetRegistryTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistryTags`: RegistryRegistryTagList
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.GetRegistryTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | **string** | Image is the repository name inside the org&#39;s namespace, as returned by the images op. It rides the query string. | 

### Return type

[**RegistryRegistryTagList**](RegistryRegistryTagList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRegistryToken

> RegistryRegistryToken PostRegistryToken(ctx).RegistryRegistryMint(registryRegistryMint).Execute()

Mints a short-lived, pull-only registry token for exactly one of the org's images, through the same IAM realm the docker CLI authenticates against.



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
	registryRegistryMint := *openapiclient.NewRegistryRegistryMint() // RegistryRegistryMint | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.PostRegistryToken(context.Background()).RegistryRegistryMint(registryRegistryMint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.PostRegistryToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRegistryToken`: RegistryRegistryToken
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.PostRegistryToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRegistryTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryRegistryMint** | [**RegistryRegistryMint**](RegistryRegistryMint.md) |  | 

### Return type

[**RegistryRegistryToken**](RegistryRegistryToken.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

