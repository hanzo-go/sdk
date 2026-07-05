# \AdminOverviewAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminOverview**](AdminOverviewAPI.md#AdminAdminOverview) | **Get** /v1/admin/overview | Fleet overview tiles
[**AdminAdminProducts**](AdminOverviewAPI.md#AdminAdminProducts) | **Get** /v1/admin/products | Product / workload registry
[**AdminAdminSync**](AdminOverviewAPI.md#AdminAdminSync) | **Post** /v1/admin/sync | Trigger a fleet re-read
[**AdminAdminUsage**](AdminOverviewAPI.md#AdminAdminUsage) | **Get** /v1/admin/usage | Fleet usage roll-up



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
	resp, r, err := apiClient.AdminOverviewAPI.AdminAdminOverview(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminOverviewAPI.AdminAdminOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminOverview`: AdminAdminOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminOverviewAPI.AdminAdminOverview`: %v\n", resp)
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
	resp, r, err := apiClient.AdminOverviewAPI.AdminAdminProducts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminOverviewAPI.AdminAdminProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminProducts`: AdminAdminProducts200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminOverviewAPI.AdminAdminProducts`: %v\n", resp)
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
	resp, r, err := apiClient.AdminOverviewAPI.AdminAdminSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminOverviewAPI.AdminAdminSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminSync`: AdminAdminSync200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminOverviewAPI.AdminAdminSync`: %v\n", resp)
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
	resp, r, err := apiClient.AdminOverviewAPI.AdminAdminUsage(context.Background()).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminOverviewAPI.AdminAdminUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminUsage`: AdminAdminUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminOverviewAPI.AdminAdminUsage`: %v\n", resp)
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

