# \BaseAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBaseBases**](BaseAPI.md#GetBaseBases) | **Get** /v1/base/bases | Lists every Base the caller can reach, one per org their token carries.
[**GetBaseBasesByOrg**](BaseAPI.md#GetBaseBasesByOrg) | **Get** /v1/base/bases/{org} | Describes ONE org&#39;s Base — whether its store exists, and what it occupies.
[**GetBaseGraphql**](BaseAPI.md#GetBaseGraphql) | **Get** /v1/base/graphql | Read your org Base&#39;s GraphQL schema
[**GetBaseHealth**](BaseAPI.md#GetBaseHealth) | **Get** /v1/base/health | Reports that the base subsystem is serving.
[**PostBaseGraphql**](BaseAPI.md#PostBaseGraphql) | **Post** /v1/base/graphql | Query or write your org Base with GraphQL



## GetBaseBases

> []BaseBaseView GetBaseBases(ctx).Execute()

Lists every Base the caller can reach, one per org their token carries.



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
	resp, r, err := apiClient.BaseAPI.GetBaseBases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseAPI.GetBaseBases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseBases`: []BaseBaseView
	fmt.Fprintf(os.Stdout, "Response from `BaseAPI.GetBaseBases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseBasesRequest struct via the builder pattern


### Return type

[**[]BaseBaseView**](BaseBaseView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBaseBasesByOrg

> BaseBaseView GetBaseBasesByOrg(ctx, org).Execute()

Describes ONE org's Base — whether its store exists, and what it occupies.



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
	org := "org_example" // string | Org is the org whose Base to describe, from the path. An org the caller's token does not carry is not found — the same answer a nonexistent one gets, so the listing cannot be used to discover which orgs exist.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaseAPI.GetBaseBasesByOrg(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseAPI.GetBaseBasesByOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseBasesByOrg`: BaseBaseView
	fmt.Fprintf(os.Stdout, "Response from `BaseAPI.GetBaseBasesByOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the org whose Base to describe, from the path. An org the caller&#39;s token does not carry is not found — the same answer a nonexistent one gets, so the listing cannot be used to discover which orgs exist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseBasesByOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseBaseView**](BaseBaseView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBaseGraphql

> GetBaseGraphql(ctx).Execute()

Read your org Base's GraphQL schema



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
	r, err := apiClient.BaseAPI.GetBaseGraphql(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseAPI.GetBaseGraphql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseGraphqlRequest struct via the builder pattern


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


## GetBaseHealth

> BaseBaseHealth GetBaseHealth(ctx).Execute()

Reports that the base subsystem is serving.



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
	resp, r, err := apiClient.BaseAPI.GetBaseHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseAPI.GetBaseHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseHealth`: BaseBaseHealth
	fmt.Fprintf(os.Stdout, "Response from `BaseAPI.GetBaseHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseHealthRequest struct via the builder pattern


### Return type

[**BaseBaseHealth**](BaseBaseHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBaseGraphql

> PostBaseGraphql(ctx).Execute()

Query or write your org Base with GraphQL



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
	r, err := apiClient.BaseAPI.PostBaseGraphql(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseAPI.PostBaseGraphql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBaseGraphqlRequest struct via the builder pattern


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

