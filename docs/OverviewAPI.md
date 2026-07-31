# \OverviewAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminOverview**](OverviewAPI.md#AdminAdminOverview) | **Get** /v1/admin/overview | Fleet overview tiles
[**AdminAdminProducts**](OverviewAPI.md#AdminAdminProducts) | **Get** /v1/admin/products | Product / workload registry
[**AdminAdminSync**](OverviewAPI.md#AdminAdminSync) | **Post** /v1/admin/sync | Trigger a fleet re-read
[**AdminAdminUsage**](OverviewAPI.md#AdminAdminUsage) | **Get** /v1/admin/usage | Fleet usage roll-up



## AdminAdminOverview

> AdminAdminOverview200Response AdminAdminOverview(ctx).Execute()

Fleet overview tiles

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
	resp, r, err := apiClient.OverviewAPI.AdminAdminOverview(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverviewAPI.AdminAdminOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminOverview`: AdminAdminOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `OverviewAPI.AdminAdminOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminOverviewRequest struct via the builder pattern


### Return type

[**AdminAdminOverview200Response**](AdminAdminOverview200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminProducts

> AdminAdminProducts200Response AdminAdminProducts(ctx).Execute()

Product / workload registry

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
	resp, r, err := apiClient.OverviewAPI.AdminAdminProducts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverviewAPI.AdminAdminProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminProducts`: AdminAdminProducts200Response
	fmt.Fprintf(os.Stdout, "Response from `OverviewAPI.AdminAdminProducts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminProductsRequest struct via the builder pattern


### Return type

[**AdminAdminProducts200Response**](AdminAdminProducts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminSync

> AdminAdminSync200Response AdminAdminSync(ctx).Execute()

Trigger a fleet re-read

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
	resp, r, err := apiClient.OverviewAPI.AdminAdminSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverviewAPI.AdminAdminSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminSync`: AdminAdminSync200Response
	fmt.Fprintf(os.Stdout, "Response from `OverviewAPI.AdminAdminSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminSyncRequest struct via the builder pattern


### Return type

[**AdminAdminSync200Response**](AdminAdminSync200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminUsage

> AdminAdminUsage200Response AdminAdminUsage(ctx).Org(org).Execute()

Fleet usage roll-up

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
	org := "org_example" // string | Scope to one org (default fleet) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OverviewAPI.AdminAdminUsage(context.Background()).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverviewAPI.AdminAdminUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminUsage`: AdminAdminUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `OverviewAPI.AdminAdminUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Scope to one org (default fleet) | 

### Return type

[**AdminAdminUsage200Response**](AdminAdminUsage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

