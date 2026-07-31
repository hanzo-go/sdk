# \RegistryAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1RegistryImages**](RegistryAPI.md#CloudGetV1RegistryImages) | **Get** /v1/registry/images | Images lists the org&#39;s container repositories, read live from the OCI catalog and filtered server-side to the org&#39;s namespace — the page can only ever hold the caller&#39;s own images.
[**CloudGetV1RegistryPackages**](RegistryAPI.md#CloudGetV1RegistryPackages) | **Get** /v1/registry/packages | Packages lists the org&#39;s npm packages — &#x60;&lt;org&gt;&#x60; and &#x60;@&lt;org&gt;/…&#x60; — from the npm registry&#39;s search index, optionally narrowed by a query within that scope.
[**CloudGetV1RegistryProjects**](RegistryAPI.md#CloudGetV1RegistryProjects) | **Get** /v1/registry/projects | Projects lists the namespaces the caller can see with what each holds: the org&#39;s slug, its repository count on the OCI catalog, and its package count on the npm registry.
[**CloudGetV1RegistryStatus**](RegistryAPI.md#CloudGetV1RegistryStatus) | **Get** /v1/registry/status | Status reports whether the OCI and npm registries are reachable and, when the OCI half is auth-gated, which token realm its challenge advertises — an honest lens for \&quot;is the registry plane up\&quot;, never a fabricated ok.
[**CloudGetV1RegistryTags**](RegistryAPI.md#CloudGetV1RegistryTags) | **Get** /v1/registry/tags | Tags lists one org-owned repository&#39;s tags, read live from the OCI registry.
[**CloudPostV1RegistryToken**](RegistryAPI.md#CloudPostV1RegistryToken) | **Post** /v1/registry/token | Token mints a short-lived, pull-only registry token for exactly one of the org&#39;s images, through the same IAM realm the docker CLI authenticates against.



## CloudGetV1RegistryImages

> CloudRegistryImageList CloudGetV1RegistryImages(ctx).Execute()

Images lists the org's container repositories, read live from the OCI catalog and filtered server-side to the org's namespace — the page can only ever hold the caller's own images.



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
	resp, r, err := apiClient.RegistryAPI.CloudGetV1RegistryImages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.CloudGetV1RegistryImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1RegistryImages`: CloudRegistryImageList
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.CloudGetV1RegistryImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1RegistryImagesRequest struct via the builder pattern


### Return type

[**CloudRegistryImageList**](CloudRegistryImageList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1RegistryPackages

> CloudRegistryPackageList CloudGetV1RegistryPackages(ctx).Query(query).Execute()

Packages lists the org's npm packages — `<org>` and `@<org>/…` — from the npm registry's search index, optionally narrowed by a query within that scope.



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
	query := "query_example" // string | Query narrows the listing within the org's scope when present; the org boundary itself is never widened by it. It rides the query string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.CloudGetV1RegistryPackages(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.CloudGetV1RegistryPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1RegistryPackages`: CloudRegistryPackageList
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.CloudGetV1RegistryPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1RegistryPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query narrows the listing within the org&#39;s scope when present; the org boundary itself is never widened by it. It rides the query string. | 

### Return type

[**CloudRegistryPackageList**](CloudRegistryPackageList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1RegistryProjects

> CloudRegistryProjectList CloudGetV1RegistryProjects(ctx).Execute()

Projects lists the namespaces the caller can see with what each holds: the org's slug, its repository count on the OCI catalog, and its package count on the npm registry.



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
	resp, r, err := apiClient.RegistryAPI.CloudGetV1RegistryProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.CloudGetV1RegistryProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1RegistryProjects`: CloudRegistryProjectList
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.CloudGetV1RegistryProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1RegistryProjectsRequest struct via the builder pattern


### Return type

[**CloudRegistryProjectList**](CloudRegistryProjectList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1RegistryStatus

> CloudRegistryStatus CloudGetV1RegistryStatus(ctx).Execute()

Status reports whether the OCI and npm registries are reachable and, when the OCI half is auth-gated, which token realm its challenge advertises — an honest lens for \"is the registry plane up\", never a fabricated ok.



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
	resp, r, err := apiClient.RegistryAPI.CloudGetV1RegistryStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.CloudGetV1RegistryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1RegistryStatus`: CloudRegistryStatus
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.CloudGetV1RegistryStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1RegistryStatusRequest struct via the builder pattern


### Return type

[**CloudRegistryStatus**](CloudRegistryStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1RegistryTags

> CloudRegistryTagList CloudGetV1RegistryTags(ctx).Image(image).Execute()

Tags lists one org-owned repository's tags, read live from the OCI registry.



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
	image := "image_example" // string | Image is the repository name inside the org's namespace, as returned by the images op. It rides the query string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.CloudGetV1RegistryTags(context.Background()).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.CloudGetV1RegistryTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1RegistryTags`: CloudRegistryTagList
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.CloudGetV1RegistryTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1RegistryTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | **string** | Image is the repository name inside the org&#39;s namespace, as returned by the images op. It rides the query string. | 

### Return type

[**CloudRegistryTagList**](CloudRegistryTagList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1RegistryToken

> CloudRegistryToken CloudPostV1RegistryToken(ctx).CloudRegistryMint(cloudRegistryMint).Execute()

Token mints a short-lived, pull-only registry token for exactly one of the org's images, through the same IAM realm the docker CLI authenticates against.



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
	cloudRegistryMint := *openapiclient.NewCloudRegistryMint() // CloudRegistryMint | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryAPI.CloudPostV1RegistryToken(context.Background()).CloudRegistryMint(cloudRegistryMint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryAPI.CloudPostV1RegistryToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1RegistryToken`: CloudRegistryToken
	fmt.Fprintf(os.Stdout, "Response from `RegistryAPI.CloudPostV1RegistryToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1RegistryTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRegistryMint** | [**CloudRegistryMint**](CloudRegistryMint.md) |  | 

### Return type

[**CloudRegistryToken**](CloudRegistryToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

