# \ShareAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsGetSharedWebsite**](ShareAPI.md#AnalyticsGetSharedWebsite) | **Get** /v1/analytics/share/{shareId} | Get a shared website by share ID (no auth required)
[**CloudGetV1Share**](ShareAPI.md#CloudGetV1Share) | **Get** /v1/share | ListShares returns the tunnel shares the caller&#39;s org currently has open, across every environment that org has enabled.
[**CloudPostV1ShareEnable**](ShareAPI.md#CloudPostV1ShareEnable) | **Post** /v1/share/enable | Enable provisions the caller org&#39;s tunnel account and returns the credential the &#x60;hanzo share&#x60; CLI needs to run a tunnel.



## AnalyticsGetSharedWebsite

> AnalyticsGetSharedWebsite200Response AnalyticsGetSharedWebsite(ctx, shareId).Execute()

Get a shared website by share ID (no auth required)

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
	shareId := "shareId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.AnalyticsGetSharedWebsite(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.AnalyticsGetSharedWebsite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSharedWebsite`: AnalyticsGetSharedWebsite200Response
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.AnalyticsGetSharedWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSharedWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsGetSharedWebsite200Response**](AnalyticsGetSharedWebsite200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Share

> CloudSharesOut CloudGetV1Share(ctx).Execute()

ListShares returns the tunnel shares the caller's org currently has open, across every environment that org has enabled.



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
	resp, r, err := apiClient.ShareAPI.CloudGetV1Share(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.CloudGetV1Share``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Share`: CloudSharesOut
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.CloudGetV1Share`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ShareRequest struct via the builder pattern


### Return type

[**CloudSharesOut**](CloudSharesOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ShareEnable

> CloudEnableResp CloudPostV1ShareEnable(ctx).Execute()

Enable provisions the caller org's tunnel account and returns the credential the `hanzo share` CLI needs to run a tunnel.



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
	resp, r, err := apiClient.ShareAPI.CloudPostV1ShareEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.CloudPostV1ShareEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ShareEnable`: CloudEnableResp
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.CloudPostV1ShareEnable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ShareEnableRequest struct via the builder pattern


### Return type

[**CloudEnableResp**](CloudEnableResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

