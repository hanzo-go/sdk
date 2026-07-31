# \DestinationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1DestinationsPlatform**](DestinationsAPI.md#CloudDeleteV1DestinationsPlatform) | **Delete** /v1/destinations/{platform} | Forgets a destination for the caller&#39;s org: every credential held in KMS, then the stored config.
[**CloudGetV1Destinations**](DestinationsAPI.md#CloudGetV1Destinations) | **Get** /v1/destinations | Reports every destination this deployment can forward to, each with the caller org&#39;s connection state: whether it is connected, whether it is enabled, whether a credential resolves right now, and the config fields the console renders for it.
[**CloudGetV1DestinationsPlatform**](DestinationsAPI.md#CloudGetV1DestinationsPlatform) | **Get** /v1/destinations/{platform} | Reports one destination&#39;s card for the caller&#39;s org — its config fields, its connection state, and whether a credential resolves right now.
[**CloudPostV1DestinationsByPlatform**](DestinationsAPI.md#CloudPostV1DestinationsByPlatform) | **Post** /v1/destinations/{platform} | 
[**CloudPostV1DestinationsPlatformTest**](DestinationsAPI.md#CloudPostV1DestinationsPlatformTest) | **Post** /v1/destinations/{platform}/test | Sends ONE synthetic pageview through the connected destination end to end and reports what the platform said.



## CloudDeleteV1DestinationsPlatform

> CloudDestinationDisconnected CloudDeleteV1DestinationsPlatform(ctx, platform).Execute()

Forgets a destination for the caller's org: every credential held in KMS, then the stored config.



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
	platform := "ga4" // string | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationsAPI.CloudDeleteV1DestinationsPlatform(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CloudDeleteV1DestinationsPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1DestinationsPlatform`: CloudDestinationDisconnected
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CloudDeleteV1DestinationsPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1DestinationsPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDestinationDisconnected**](CloudDestinationDisconnected.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Destinations

> CloudDestinationList CloudGetV1Destinations(ctx).Execute()

Reports every destination this deployment can forward to, each with the caller org's connection state: whether it is connected, whether it is enabled, whether a credential resolves right now, and the config fields the console renders for it.



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
	resp, r, err := apiClient.DestinationsAPI.CloudGetV1Destinations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CloudGetV1Destinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Destinations`: CloudDestinationList
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CloudGetV1Destinations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DestinationsRequest struct via the builder pattern


### Return type

[**CloudDestinationList**](CloudDestinationList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DestinationsPlatform

> CloudDestinationStatus CloudGetV1DestinationsPlatform(ctx, platform).Execute()

Reports one destination's card for the caller's org — its config fields, its connection state, and whether a credential resolves right now.



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
	platform := "ga4" // string | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationsAPI.CloudGetV1DestinationsPlatform(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CloudGetV1DestinationsPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DestinationsPlatform`: CloudDestinationStatus
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CloudGetV1DestinationsPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DestinationsPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDestinationStatus**](CloudDestinationStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1DestinationsByPlatform

> CloudDestinationStatus CloudPostV1DestinationsByPlatform(ctx, platform).RequestBody(requestBody).Execute()



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
	platform := "platform_example" // string | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationsAPI.CloudPostV1DestinationsByPlatform(context.Background(), platform).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CloudPostV1DestinationsByPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1DestinationsByPlatform`: CloudDestinationStatus
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CloudPostV1DestinationsByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1DestinationsByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**CloudDestinationStatus**](CloudDestinationStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1DestinationsPlatformTest

> CloudDestinationTest CloudPostV1DestinationsPlatformTest(ctx, platform).Execute()

Sends ONE synthetic pageview through the connected destination end to end and reports what the platform said.



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
	platform := "ga4" // string | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationsAPI.CloudPostV1DestinationsPlatformTest(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CloudPostV1DestinationsPlatformTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1DestinationsPlatformTest`: CloudDestinationTest
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CloudPostV1DestinationsPlatformTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1DestinationsPlatformTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDestinationTest**](CloudDestinationTest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

