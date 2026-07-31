# \OrgsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1OrgsOrgEntitlements**](OrgsAPI.md#CloudGetV1OrgsOrgEntitlements) | **Get** /v1/orgs/{org}/entitlements | Get lists the products an org has ENABLED — its own intent, which the console&#39;s paid-product sidebar reads to decide what to show.
[**CloudPostV1OrgsOrgEntitlements**](OrgsAPI.md#CloudPostV1OrgsOrgEntitlements) | **Post** /v1/orgs/{org}/entitlements | Post turns products on or off for an org and returns the enabled set afterwards.



## CloudGetV1OrgsOrgEntitlements

> CloudEntitlementsView CloudGetV1OrgsOrgEntitlements(ctx, org).Execute()

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
	resp, r, err := apiClient.OrgsAPI.CloudGetV1OrgsOrgEntitlements(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.CloudGetV1OrgsOrgEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1OrgsOrgEntitlements`: CloudEntitlementsView
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.CloudGetV1OrgsOrgEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1OrgsOrgEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudEntitlementsView**](CloudEntitlementsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1OrgsOrgEntitlements

> CloudEntitlementsView CloudPostV1OrgsOrgEntitlements(ctx, org).CloudMutateReq(cloudMutateReq).Execute()

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
	cloudMutateReq := *openapiclient.NewCloudMutateReq() // CloudMutateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.CloudPostV1OrgsOrgEntitlements(context.Background(), org).CloudMutateReq(cloudMutateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.CloudPostV1OrgsOrgEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1OrgsOrgEntitlements`: CloudEntitlementsView
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.CloudPostV1OrgsOrgEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1OrgsOrgEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudMutateReq** | [**CloudMutateReq**](CloudMutateReq.md) |  | 

### Return type

[**CloudEntitlementsView**](CloudEntitlementsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

