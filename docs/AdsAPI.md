# \AdsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdsCampaignsById**](AdsAPI.md#DeleteAdsCampaignsById) | **Delete** /v1/ads/campaigns/{id} | Removes one of the caller org&#39;s campaigns and answers 204 with no body.
[**GetAdsCampaigns**](AdsAPI.md#GetAdsCampaigns) | **Get** /v1/ads/campaigns | Returns the caller org&#39;s ad campaigns, most recently updated first, optionally narrowed to one lifecycle status.
[**GetAdsCampaignsById**](AdsAPI.md#GetAdsCampaignsById) | **Get** /v1/ads/campaigns/{id} | Returns one of the caller org&#39;s campaigns.
[**GetAdsSummary**](AdsAPI.md#GetAdsSummary) | **Get** /v1/ads/summary | Rolls the caller org&#39;s ad campaigns up into four numbers: how many campaigns exist, how many are active, and the summed budget and spend across all of them.
[**PostAdsCampaigns**](AdsAPI.md#PostAdsCampaigns) | **Post** /v1/ads/campaigns | Registers a new ad campaign for the caller&#39;s org and answers 201 with the stored row.
[**PostAdsCampaignsByIdLaunch**](AdsAPI.md#PostAdsCampaignsByIdLaunch) | **Post** /v1/ads/campaigns/{id}/launch | Run one of your stored campaigns on its ad network
[**PutAdsCampaignsById**](AdsAPI.md#PutAdsCampaignsById) | **Put** /v1/ads/campaigns/{id} | Replaces the user-owned fields of one of the caller org&#39;s campaigns and answers the stored row.



## DeleteAdsCampaignsById

> DeleteAdsCampaignsById(ctx, id).Execute()

Removes one of the caller org's campaigns and answers 204 with no body.



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
	id := "camp_2f9c1d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdsAPI.DeleteAdsCampaignsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.DeleteAdsCampaignsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdsCampaignsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdsCampaigns

> CampaignList GetAdsCampaigns(ctx).Status(status).Limit(limit).Execute()

Returns the caller org's ad campaigns, most recently updated first, optionally narrowed to one lifecycle status.



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
	status := "active" // string | Status filters to one lifecycle state (draft, active, paused, completed). Empty returns every campaign the org has. (optional)
	limit := int32(50) // int32 | Limit caps how many campaigns come back: default 200, maximum 1000. A value that is not a positive integer reads as the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetAdsCampaigns(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAdsCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdsCampaigns`: CampaignList
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAdsCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdsCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status filters to one lifecycle state (draft, active, paused, completed). Empty returns every campaign the org has. | 
 **limit** | **int32** | Limit caps how many campaigns come back: default 200, maximum 1000. A value that is not a positive integer reads as the default. | 

### Return type

[**CampaignList**](CampaignList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdsCampaignsById

> AdCampaign GetAdsCampaignsById(ctx, id).Execute()

Returns one of the caller org's campaigns.



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
	id := "camp_2f9c1d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetAdsCampaignsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAdsCampaignsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdsCampaignsById`: AdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAdsCampaignsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdsCampaignsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdCampaign**](AdCampaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdsSummary

> AdSummary GetAdsSummary(ctx).Execute()

Rolls the caller org's ad campaigns up into four numbers: how many campaigns exist, how many are active, and the summed budget and spend across all of them.



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
	resp, r, err := apiClient.AdsAPI.GetAdsSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAdsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdsSummary`: AdSummary
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAdsSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdsSummaryRequest struct via the builder pattern


### Return type

[**AdSummary**](AdSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdsCampaigns

> AdCampaign PostAdsCampaigns(ctx).CampaignInput(campaignInput).Execute()

Registers a new ad campaign for the caller's org and answers 201 with the stored row.



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
	campaignInput := *openapiclient.NewCampaignInput() // CampaignInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.PostAdsCampaigns(context.Background()).CampaignInput(campaignInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.PostAdsCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdsCampaigns`: AdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.PostAdsCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdsCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignInput** | [**CampaignInput**](CampaignInput.md) |  | 

### Return type

[**AdCampaign**](AdCampaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdsCampaignsByIdLaunch

> PostAdsCampaignsByIdLaunch(ctx, id).Execute()

Run one of your stored campaigns on its ad network



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdsAPI.PostAdsCampaignsByIdLaunch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.PostAdsCampaignsByIdLaunch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdsCampaignsByIdLaunchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAdsCampaignsById

> AdCampaign PutAdsCampaignsById(ctx, id).UpdateCampaignIn(updateCampaignIn).Execute()

Replaces the user-owned fields of one of the caller org's campaigns and answers the stored row.



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
	id := "camp_2f9c1d" // string | 
	updateCampaignIn := *openapiclient.NewUpdateCampaignIn() // UpdateCampaignIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.PutAdsCampaignsById(context.Background(), id).UpdateCampaignIn(updateCampaignIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.PutAdsCampaignsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAdsCampaignsById`: AdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.PutAdsCampaignsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAdsCampaignsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCampaignIn** | [**UpdateCampaignIn**](UpdateCampaignIn.md) |  | 

### Return type

[**AdCampaign**](AdCampaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

