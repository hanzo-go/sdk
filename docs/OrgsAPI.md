# \OrgsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsByOrgEntitlements**](OrgsAPI.md#GetOrgsByOrgEntitlements) | **Get** /v1/orgs/{org}/entitlements | Get lists the products an org has ENABLED — its own intent, which the console&#39;s paid-product sidebar reads to decide what to show.
[**PostOrgs**](OrgsAPI.md#PostOrgs) | **Post** /v1/orgs | Onboard creates the caller&#39;s organization.
[**PostOrgsByOrgEntitlements**](OrgsAPI.md#PostOrgsByOrgEntitlements) | **Post** /v1/orgs/{org}/entitlements | Post turns products on or off for an org and returns the enabled set afterwards.



## GetOrgsByOrgEntitlements

> EntitlementsView GetOrgsByOrgEntitlements(ctx, org).Execute()

Get lists the products an org has ENABLED — its own intent, which the console's paid-product sidebar reads to decide what to show.



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
	org := "org_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetOrgsByOrgEntitlements(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrgsByOrgEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsByOrgEntitlements`: EntitlementsView
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrgsByOrgEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsByOrgEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntitlementsView**](EntitlementsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgs

> OnboardResp PostOrgs(ctx).OnboardReq(onboardReq).Execute()

Onboard creates the caller's organization.



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
	onboardReq := *openapiclient.NewOnboardReq() // OnboardReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.PostOrgs(context.Background()).OnboardReq(onboardReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.PostOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgs`: OnboardResp
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.PostOrgs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onboardReq** | [**OnboardReq**](OnboardReq.md) |  | 

### Return type

[**OnboardResp**](OnboardResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsByOrgEntitlements

> EntitlementsView PostOrgsByOrgEntitlements(ctx, org).MutateReq(mutateReq).Execute()

Post turns products on or off for an org and returns the enabled set afterwards.



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
	org := "org_example" // string | 
	mutateReq := *openapiclient.NewMutateReq() // MutateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.PostOrgsByOrgEntitlements(context.Background(), org).MutateReq(mutateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.PostOrgsByOrgEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsByOrgEntitlements`: EntitlementsView
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.PostOrgsByOrgEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsByOrgEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutateReq** | [**MutateReq**](MutateReq.md) |  | 

### Return type

[**EntitlementsView**](EntitlementsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

