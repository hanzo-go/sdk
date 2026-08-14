# \DestinationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDestinationsByPlatform**](DestinationsAPI.md#DeleteDestinationsByPlatform) | **Delete** /v1/destinations/{platform} | Forgets a destination for the caller&#39;s org: every credential held in KMS, then the stored config.
[**GetDestinations**](DestinationsAPI.md#GetDestinations) | **Get** /v1/destinations | Reports every destination this deployment can forward to, each with the caller org&#39;s connection state: whether it is connected, whether it is enabled, whether a credential resolves right now, and the config fields the console renders for it.
[**GetDestinationsByPlatform**](DestinationsAPI.md#GetDestinationsByPlatform) | **Get** /v1/destinations/{platform} | Reports one destination&#39;s card for the caller&#39;s org — its config fields, its connection state, and whether a credential resolves right now.
[**PostDestinationsByPlatform**](DestinationsAPI.md#PostDestinationsByPlatform) | **Post** /v1/destinations/{platform} | Connect one conversion destination for your org, or update the one you have
[**PostDestinationsByPlatformTest**](DestinationsAPI.md#PostDestinationsByPlatformTest) | **Post** /v1/destinations/{platform}/test | Sends ONE synthetic pageview through the connected destination end to end and reports what the platform said.



## DeleteDestinationsByPlatform

> DestinationDisconnected DeleteDestinationsByPlatform(ctx, platform).Execute()

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
	resp, r, err := apiClient.DestinationsAPI.DeleteDestinationsByPlatform(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.DeleteDestinationsByPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDestinationsByPlatform`: DestinationDisconnected
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.DeleteDestinationsByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationsByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DestinationDisconnected**](DestinationDisconnected.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinations

> DestinationList GetDestinations(ctx).Execute()

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
	resp, r, err := apiClient.DestinationsAPI.GetDestinations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.GetDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestinations`: DestinationList
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.GetDestinations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationsRequest struct via the builder pattern


### Return type

[**DestinationList**](DestinationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinationsByPlatform

> DestinationStatus GetDestinationsByPlatform(ctx, platform).Execute()

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
	resp, r, err := apiClient.DestinationsAPI.GetDestinationsByPlatform(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.GetDestinationsByPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestinationsByPlatform`: DestinationStatus
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.GetDestinationsByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationsByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DestinationStatus**](DestinationStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDestinationsByPlatform

> DestinationStatus PostDestinationsByPlatform(ctx, platform).RequestBody(requestBody).Execute()

Connect one conversion destination for your org, or update the one you have



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
	resp, r, err := apiClient.DestinationsAPI.PostDestinationsByPlatform(context.Background(), platform).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.PostDestinationsByPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDestinationsByPlatform`: DestinationStatus
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.PostDestinationsByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDestinationsByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**DestinationStatus**](DestinationStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDestinationsByPlatformTest

> DestinationTest PostDestinationsByPlatformTest(ctx, platform).Execute()

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
	resp, r, err := apiClient.DestinationsAPI.PostDestinationsByPlatformTest(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.PostDestinationsByPlatformTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDestinationsByPlatformTest`: DestinationTest
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.PostDestinationsByPlatformTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | posthog | umami. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDestinationsByPlatformTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DestinationTest**](DestinationTest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

