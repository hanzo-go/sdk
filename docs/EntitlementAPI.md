# \EntitlementAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntitlement**](EntitlementAPI.md#GetEntitlement) | **Get** /v1/entitlement | Reports which console apps the CALLER&#39;s org may open, and the plan slug that decides it.
[**GetEntitlementOrgsByOrg**](EntitlementAPI.md#GetEntitlementOrgsByOrg) | **Get** /v1/entitlement/orgs/{org} | Lists the products an org has ENABLED — its own intent, which the console&#39;s paid-product sidebar reads to decide what to show.
[**PostEntitlementOrgsByOrg**](EntitlementAPI.md#PostEntitlementOrgsByOrg) | **Post** /v1/entitlement/orgs/{org} | Turns products on or off for an org and returns the enabled set afterwards.



## GetEntitlement

> EntitlementProjectionView GetEntitlement(ctx).Execute()

Reports which console apps the CALLER's org may open, and the plan slug that decides it.



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
	resp, r, err := apiClient.EntitlementAPI.GetEntitlement(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.GetEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlement`: EntitlementProjectionView
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


### Return type

[**EntitlementProjectionView**](EntitlementProjectionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementOrgsByOrg

> EntitlementEntitlementsView GetEntitlementOrgsByOrg(ctx, org).Execute()

Lists the products an org has ENABLED — its own intent, which the console's paid-product sidebar reads to decide what to show.



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
	org := "org_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.GetEntitlementOrgsByOrg(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.GetEntitlementOrgsByOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementOrgsByOrg`: EntitlementEntitlementsView
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.GetEntitlementOrgsByOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementOrgsByOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntitlementEntitlementsView**](EntitlementEntitlementsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEntitlementOrgsByOrg

> EntitlementEntitlementsView PostEntitlementOrgsByOrg(ctx, org).EntitlementMutateReq(entitlementMutateReq).Execute()

Turns products on or off for an org and returns the enabled set afterwards.



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
	org := "org_example" // string | 
	entitlementMutateReq := *openapiclient.NewEntitlementMutateReq() // EntitlementMutateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.PostEntitlementOrgsByOrg(context.Background(), org).EntitlementMutateReq(entitlementMutateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.PostEntitlementOrgsByOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEntitlementOrgsByOrg`: EntitlementEntitlementsView
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.PostEntitlementOrgsByOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEntitlementOrgsByOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlementMutateReq** | [**EntitlementMutateReq**](EntitlementMutateReq.md) |  | 

### Return type

[**EntitlementEntitlementsView**](EntitlementEntitlementsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

