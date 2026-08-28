# \EntitlementAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntitlement**](EntitlementAPI.md#GetEntitlement) | **Get** /v1/entitlement | Projection reports which console apps the CALLER&#39;s org may open, and the plan slug that decides it.
[**GetEntitlementOrgsByOrg**](EntitlementAPI.md#GetEntitlementOrgsByOrg) | **Get** /v1/entitlement/orgs/{org} | Get lists the products an org has ENABLED — its own intent, which the console&#39;s paid-product sidebar reads to decide what to show.
[**PostEntitlementOrgsByOrg**](EntitlementAPI.md#PostEntitlementOrgsByOrg) | **Post** /v1/entitlement/orgs/{org} | Post turns products on or off for an org and returns the enabled set afterwards.



## GetEntitlement

> ProjectionView GetEntitlement(ctx).Execute()

Projection reports which console apps the CALLER's org may open, and the plan slug that decides it.



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
	// response from `GetEntitlement`: ProjectionView
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


### Return type

[**ProjectionView**](ProjectionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementOrgsByOrg

> EntitlementsView GetEntitlementOrgsByOrg(ctx, org).Execute()

Get lists the products an org has ENABLED — its own intent, which the console's paid-product sidebar reads to decide what to show.



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
	// response from `GetEntitlementOrgsByOrg`: EntitlementsView
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

[**EntitlementsView**](EntitlementsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEntitlementOrgsByOrg

> EntitlementsView PostEntitlementOrgsByOrg(ctx, org).MutateReq(mutateReq).Execute()

Post turns products on or off for an org and returns the enabled set afterwards.



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
	mutateReq := *openapiclient.NewMutateReq() // MutateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.PostEntitlementOrgsByOrg(context.Background(), org).MutateReq(mutateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.PostEntitlementOrgsByOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEntitlementOrgsByOrg`: EntitlementsView
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

 **mutateReq** | [**MutateReq**](MutateReq.md) |  | 

### Return type

[**EntitlementsView**](EntitlementsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

