# \DestinationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDestinationByPlatform**](DestinationAPI.md#DeleteDestinationByPlatform) | **Delete** /v1/destination/{platform} | Forgets a destination for the caller&#39;s org: every credential held in KMS, then the stored config.
[**GetDestination**](DestinationAPI.md#GetDestination) | **Get** /v1/destination | Reports every destination this deployment can forward to, each with the caller org&#39;s connection state: whether it is connected, whether it is enabled, whether a credential resolves right now, and the config fields the console renders for it.
[**GetDestinationByPlatform**](DestinationAPI.md#GetDestinationByPlatform) | **Get** /v1/destination/{platform} | Reports one destination&#39;s card for the caller&#39;s org — its config fields, its connection state, and whether a credential resolves right now.
[**PostDestinationByPlatform**](DestinationAPI.md#PostDestinationByPlatform) | **Post** /v1/destination/{platform} | Connect one conversion destination for your org, or update the one you have
[**PostDestinationByPlatformTest**](DestinationAPI.md#PostDestinationByPlatformTest) | **Post** /v1/destination/{platform}/test | Sends ONE synthetic pageview through the connected destination end to end and reports what the platform said.



## DeleteDestinationByPlatform

> DestinationDisconnected DeleteDestinationByPlatform(ctx, platform).Execute()

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
	platform := "ga4" // string | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | insights | analytics.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationAPI.DeleteDestinationByPlatform(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationAPI.DeleteDestinationByPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDestinationByPlatform`: DestinationDisconnected
	fmt.Fprintf(os.Stdout, "Response from `DestinationAPI.DeleteDestinationByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | insights | analytics. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DestinationDisconnected**](DestinationDisconnected.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestination

> DestinationList GetDestination(ctx).Execute()

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
	resp, r, err := apiClient.DestinationAPI.GetDestination(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationAPI.GetDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestination`: DestinationList
	fmt.Fprintf(os.Stdout, "Response from `DestinationAPI.GetDestination`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationRequest struct via the builder pattern


### Return type

[**DestinationList**](DestinationList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinationByPlatform

> DestinationStatus GetDestinationByPlatform(ctx, platform).Execute()

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
	platform := "ga4" // string | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | insights | analytics.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationAPI.GetDestinationByPlatform(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationAPI.GetDestinationByPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestinationByPlatform`: DestinationStatus
	fmt.Fprintf(os.Stdout, "Response from `DestinationAPI.GetDestinationByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | insights | analytics. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DestinationStatus**](DestinationStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDestinationByPlatform

> DestinationStatus PostDestinationByPlatform(ctx, platform).RequestBody(requestBody).Execute()

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
	resp, r, err := apiClient.DestinationAPI.PostDestinationByPlatform(context.Background(), platform).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationAPI.PostDestinationByPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDestinationByPlatform`: DestinationStatus
	fmt.Fprintf(os.Stdout, "Response from `DestinationAPI.PostDestinationByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDestinationByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**DestinationStatus**](DestinationStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDestinationByPlatformTest

> DestinationTest PostDestinationByPlatformTest(ctx, platform).Execute()

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
	platform := "ga4" // string | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | insights | analytics.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationAPI.PostDestinationByPlatformTest(context.Background(), platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationAPI.PostDestinationByPlatformTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDestinationByPlatformTest`: DestinationTest
	fmt.Fprintf(os.Stdout, "Response from `DestinationAPI.PostDestinationByPlatformTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform is the destination to act on, from the path: ga4 | meta | tiktok | linkedin | x | reddit | insights | analytics. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDestinationByPlatformTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DestinationTest**](DestinationTest.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

