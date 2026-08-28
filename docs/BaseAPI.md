# \BaseAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBaseBases**](BaseAPI.md#GetBaseBases) | **Get** /v1/base/bases | Lists every Base the caller can reach, one per org their token carries.
[**GetBaseBasesByOrg**](BaseAPI.md#GetBaseBasesByOrg) | **Get** /v1/base/bases/{org} | Describes ONE org&#39;s Base — whether its store exists, and what it occupies.
[**GetBaseHealth**](BaseAPI.md#GetBaseHealth) | **Get** /v1/base/health | Reports that the base subsystem is serving.



## GetBaseBases

> []BaseView GetBaseBases(ctx).Execute()

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
	// response from `GetBaseBases`: []BaseView
	fmt.Fprintf(os.Stdout, "Response from `BaseAPI.GetBaseBases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseBasesRequest struct via the builder pattern


### Return type

[**[]BaseView**](BaseView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBaseBasesByOrg

> BaseView GetBaseBasesByOrg(ctx, org).Execute()

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
	// response from `GetBaseBasesByOrg`: BaseView
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

[**BaseView**](BaseView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBaseHealth

> BaseHealth GetBaseHealth(ctx).Execute()

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
	// response from `GetBaseHealth`: BaseHealth
	fmt.Fprintf(os.Stdout, "Response from `BaseAPI.GetBaseHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseHealthRequest struct via the builder pattern


### Return type

[**BaseHealth**](BaseHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

