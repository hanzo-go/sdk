# \AdAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdCampaignsById**](AdAPI.md#DeleteAdCampaignsById) | **Delete** /v1/ad/campaigns/{id} | Removes one of the caller org&#39;s campaigns and answers 204 with no body.
[**GetAdCampaigns**](AdAPI.md#GetAdCampaigns) | **Get** /v1/ad/campaigns | Returns the caller org&#39;s ad campaigns, most recently updated first, optionally narrowed to one lifecycle status.
[**GetAdCampaignsById**](AdAPI.md#GetAdCampaignsById) | **Get** /v1/ad/campaigns/{id} | Returns one of the caller org&#39;s campaigns.
[**GetAdSummary**](AdAPI.md#GetAdSummary) | **Get** /v1/ad/summary | Rolls the caller org&#39;s ad campaigns up into four numbers: how many campaigns exist, how many are active, and the summed budget and spend across all of them.
[**PostAdCampaigns**](AdAPI.md#PostAdCampaigns) | **Post** /v1/ad/campaigns | Registers a new ad campaign for the caller&#39;s org and answers 201 with the stored row.
[**PostAdCampaignsByIdLaunch**](AdAPI.md#PostAdCampaignsByIdLaunch) | **Post** /v1/ad/campaigns/{id}/launch | Run one of your stored campaigns on its ad network
[**PutAdCampaignsById**](AdAPI.md#PutAdCampaignsById) | **Put** /v1/ad/campaigns/{id} | Replaces the user-owned fields of one of the caller org&#39;s campaigns and answers the stored row.



## DeleteAdCampaignsById

> DeleteAdCampaignsById(ctx, id).Execute()

Removes one of the caller org's campaigns and answers 204 with no body.



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
	id := "camp_2f9c1d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdAPI.DeleteAdCampaignsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.DeleteAdCampaignsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAdCampaignsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAdCampaigns

> CampaignList GetAdCampaigns(ctx).Status(status).Limit(limit).Execute()

Returns the caller org's ad campaigns, most recently updated first, optionally narrowed to one lifecycle status.



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
	status := "active" // string | Status filters to one lifecycle state (draft, active, paused, completed). Empty returns every campaign the org has. (optional)
	limit := int64(50) // int64 | Limit caps how many campaigns come back: default 200, maximum 1000. A value that is not a positive integer reads as the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAPI.GetAdCampaigns(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.GetAdCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdCampaigns`: CampaignList
	fmt.Fprintf(os.Stdout, "Response from `AdAPI.GetAdCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status filters to one lifecycle state (draft, active, paused, completed). Empty returns every campaign the org has. | 
 **limit** | **int64** | Limit caps how many campaigns come back: default 200, maximum 1000. A value that is not a positive integer reads as the default. | 

### Return type

[**CampaignList**](CampaignList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdCampaignsById

> AdCampaign GetAdCampaignsById(ctx, id).Execute()

Returns one of the caller org's campaigns.



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
	id := "camp_2f9c1d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAPI.GetAdCampaignsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.GetAdCampaignsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdCampaignsById`: AdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdAPI.GetAdCampaignsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdCampaignsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdCampaign**](AdCampaign.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdSummary

> AdSummary GetAdSummary(ctx).Execute()

Rolls the caller org's ad campaigns up into four numbers: how many campaigns exist, how many are active, and the summed budget and spend across all of them.



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
	resp, r, err := apiClient.AdAPI.GetAdSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.GetAdSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdSummary`: AdSummary
	fmt.Fprintf(os.Stdout, "Response from `AdAPI.GetAdSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdSummaryRequest struct via the builder pattern


### Return type

[**AdSummary**](AdSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdCampaigns

> AdCampaign PostAdCampaigns(ctx).CampaignInput(campaignInput).Execute()

Registers a new ad campaign for the caller's org and answers 201 with the stored row.



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
	campaignInput := *openapiclient.NewCampaignInput() // CampaignInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAPI.PostAdCampaigns(context.Background()).CampaignInput(campaignInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.PostAdCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdCampaigns`: AdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdAPI.PostAdCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignInput** | [**CampaignInput**](CampaignInput.md) |  | 

### Return type

[**AdCampaign**](AdCampaign.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdCampaignsByIdLaunch

> PostAdCampaignsByIdLaunch(ctx, id).Execute()

Run one of your stored campaigns on its ad network



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdAPI.PostAdCampaignsByIdLaunch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.PostAdCampaignsByIdLaunch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAdCampaignsByIdLaunchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PutAdCampaignsById

> AdCampaign PutAdCampaignsById(ctx, id).UpdateCampaignIn(updateCampaignIn).Execute()

Replaces the user-owned fields of one of the caller org's campaigns and answers the stored row.



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
	id := "camp_2f9c1d" // string | 
	updateCampaignIn := *openapiclient.NewUpdateCampaignIn() // UpdateCampaignIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAPI.PutAdCampaignsById(context.Background(), id).UpdateCampaignIn(updateCampaignIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.PutAdCampaignsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAdCampaignsById`: AdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdAPI.PutAdCampaignsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAdCampaignsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCampaignIn** | [**UpdateCampaignIn**](UpdateCampaignIn.md) |  | 

### Return type

[**AdCampaign**](AdCampaign.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

