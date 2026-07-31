# \AdsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1AdsCampaignsId**](AdsAPI.md#CloudDeleteV1AdsCampaignsId) | **Delete** /v1/ads/campaigns/{id} | Removes one of the caller org&#39;s campaigns and answers 204 with no body.
[**CloudGetV1AdsCampaigns**](AdsAPI.md#CloudGetV1AdsCampaigns) | **Get** /v1/ads/campaigns | Returns the caller org&#39;s ad campaigns, most recently updated first, optionally narrowed to one lifecycle status.
[**CloudGetV1AdsCampaignsId**](AdsAPI.md#CloudGetV1AdsCampaignsId) | **Get** /v1/ads/campaigns/{id} | Returns one of the caller org&#39;s campaigns.
[**CloudGetV1AdsSummary**](AdsAPI.md#CloudGetV1AdsSummary) | **Get** /v1/ads/summary | Rolls the caller org&#39;s ad campaigns up into four numbers: how many campaigns exist, how many are active, and the summed budget and spend across all of them.
[**CloudPostV1AdsCampaigns**](AdsAPI.md#CloudPostV1AdsCampaigns) | **Post** /v1/ads/campaigns | Registers a new ad campaign for the caller&#39;s org and answers 201 with the stored row.
[**CloudPostV1AdsCampaignsByIdLaunch**](AdsAPI.md#CloudPostV1AdsCampaignsByIdLaunch) | **Post** /v1/ads/campaigns/{id}/launch | 
[**CloudPutV1AdsCampaignsId**](AdsAPI.md#CloudPutV1AdsCampaignsId) | **Put** /v1/ads/campaigns/{id} | Replaces the user-owned fields of one of the caller org&#39;s campaigns and answers the stored row.



## CloudDeleteV1AdsCampaignsId

> CloudDeleteV1AdsCampaignsId(ctx, id).Execute()

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
	r, err := apiClient.AdsAPI.CloudDeleteV1AdsCampaignsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CloudDeleteV1AdsCampaignsId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1AdsCampaignsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdsCampaigns

> CloudCampaignList CloudGetV1AdsCampaigns(ctx).Status(status).Limit(limit).Execute()

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
	resp, r, err := apiClient.AdsAPI.CloudGetV1AdsCampaigns(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CloudGetV1AdsCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdsCampaigns`: CloudCampaignList
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CloudGetV1AdsCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdsCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status filters to one lifecycle state (draft, active, paused, completed). Empty returns every campaign the org has. | 
 **limit** | **int32** | Limit caps how many campaigns come back: default 200, maximum 1000. A value that is not a positive integer reads as the default. | 

### Return type

[**CloudCampaignList**](CloudCampaignList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdsCampaignsId

> CloudAdCampaign CloudGetV1AdsCampaignsId(ctx, id).Execute()

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
	resp, r, err := apiClient.AdsAPI.CloudGetV1AdsCampaignsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CloudGetV1AdsCampaignsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdsCampaignsId`: CloudAdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CloudGetV1AdsCampaignsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdsCampaignsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAdCampaign**](CloudAdCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdsSummary

> CloudAdSummary CloudGetV1AdsSummary(ctx).Execute()

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
	resp, r, err := apiClient.AdsAPI.CloudGetV1AdsSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CloudGetV1AdsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdsSummary`: CloudAdSummary
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CloudGetV1AdsSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdsSummaryRequest struct via the builder pattern


### Return type

[**CloudAdSummary**](CloudAdSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdsCampaigns

> CloudAdCampaign CloudPostV1AdsCampaigns(ctx).CloudCampaignInput(cloudCampaignInput).Execute()

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
	cloudCampaignInput := *openapiclient.NewCloudCampaignInput() // CloudCampaignInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.CloudPostV1AdsCampaigns(context.Background()).CloudCampaignInput(cloudCampaignInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CloudPostV1AdsCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdsCampaigns`: CloudAdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CloudPostV1AdsCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdsCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCampaignInput** | [**CloudCampaignInput**](CloudCampaignInput.md) |  | 

### Return type

[**CloudAdCampaign**](CloudAdCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdsCampaignsByIdLaunch

> CloudPostV1AdsCampaignsByIdLaunch(ctx, id).Execute()



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
	r, err := apiClient.AdsAPI.CloudPostV1AdsCampaignsByIdLaunch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CloudPostV1AdsCampaignsByIdLaunch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AdsCampaignsByIdLaunchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1AdsCampaignsId

> CloudAdCampaign CloudPutV1AdsCampaignsId(ctx, id).CloudUpdateCampaignIn(cloudUpdateCampaignIn).Execute()

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
	cloudUpdateCampaignIn := *openapiclient.NewCloudUpdateCampaignIn() // CloudUpdateCampaignIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.CloudPutV1AdsCampaignsId(context.Background(), id).CloudUpdateCampaignIn(cloudUpdateCampaignIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CloudPutV1AdsCampaignsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1AdsCampaignsId`: CloudAdCampaign
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CloudPutV1AdsCampaignsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1AdsCampaignsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudUpdateCampaignIn** | [**CloudUpdateCampaignIn**](CloudUpdateCampaignIn.md) |  | 

### Return type

[**CloudAdCampaign**](CloudAdCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

