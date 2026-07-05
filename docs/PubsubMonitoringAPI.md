# \PubsubMonitoringAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PubsubGetConnz**](PubsubMonitoringAPI.md#PubsubGetConnz) | **Get** /v1/pubsub/connz | Connection details
[**PubsubGetGatewayz**](PubsubMonitoringAPI.md#PubsubGetGatewayz) | **Get** /v1/pubsub/gatewayz | Gateway status
[**PubsubGetJsz**](PubsubMonitoringAPI.md#PubsubGetJsz) | **Get** /v1/pubsub/jsz | JetStream info
[**PubsubGetLeafz**](PubsubMonitoringAPI.md#PubsubGetLeafz) | **Get** /v1/pubsub/leafz | Leaf node info
[**PubsubGetRoutez**](PubsubMonitoringAPI.md#PubsubGetRoutez) | **Get** /v1/pubsub/routez | Cluster routes
[**PubsubGetSubsz**](PubsubMonitoringAPI.md#PubsubGetSubsz) | **Get** /v1/pubsub/subsz | Subscription info
[**PubsubGetVarz**](PubsubMonitoringAPI.md#PubsubGetVarz) | **Get** /v1/pubsub/varz | Server statistics



## PubsubGetConnz

> PubsubGetConnz200Response PubsubGetConnz(ctx).Sort(sort).Limit(limit).Execute()

Connection details



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
	sort := "sort_example" // string | Sort connections by field (optional)
	limit := int32(56) // int32 |  (optional) (default to 1024)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubMonitoringAPI.PubsubGetConnz(context.Background()).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubMonitoringAPI.PubsubGetConnz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetConnz`: PubsubGetConnz200Response
	fmt.Fprintf(os.Stdout, "Response from `PubsubMonitoringAPI.PubsubGetConnz`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetConnzRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort connections by field | 
 **limit** | **int32** |  | [default to 1024]

### Return type

[**PubsubGetConnz200Response**](PubsubGetConnz200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubGetGatewayz

> map[string]interface{} PubsubGetGatewayz(ctx).Execute()

Gateway status



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
	resp, r, err := apiClient.PubsubMonitoringAPI.PubsubGetGatewayz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubMonitoringAPI.PubsubGetGatewayz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetGatewayz`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PubsubMonitoringAPI.PubsubGetGatewayz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetGatewayzRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubGetJsz

> PubsubJetStreamInfo PubsubGetJsz(ctx).Execute()

JetStream info



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
	resp, r, err := apiClient.PubsubMonitoringAPI.PubsubGetJsz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubMonitoringAPI.PubsubGetJsz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetJsz`: PubsubJetStreamInfo
	fmt.Fprintf(os.Stdout, "Response from `PubsubMonitoringAPI.PubsubGetJsz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetJszRequest struct via the builder pattern


### Return type

[**PubsubJetStreamInfo**](PubsubJetStreamInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubGetLeafz

> map[string]interface{} PubsubGetLeafz(ctx).Execute()

Leaf node info



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
	resp, r, err := apiClient.PubsubMonitoringAPI.PubsubGetLeafz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubMonitoringAPI.PubsubGetLeafz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetLeafz`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PubsubMonitoringAPI.PubsubGetLeafz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetLeafzRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubGetRoutez

> map[string]interface{} PubsubGetRoutez(ctx).Execute()

Cluster routes



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
	resp, r, err := apiClient.PubsubMonitoringAPI.PubsubGetRoutez(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubMonitoringAPI.PubsubGetRoutez``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetRoutez`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PubsubMonitoringAPI.PubsubGetRoutez`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetRoutezRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubGetSubsz

> map[string]interface{} PubsubGetSubsz(ctx).Execute()

Subscription info



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
	resp, r, err := apiClient.PubsubMonitoringAPI.PubsubGetSubsz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubMonitoringAPI.PubsubGetSubsz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetSubsz`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PubsubMonitoringAPI.PubsubGetSubsz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetSubszRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubGetVarz

> PubsubServerVarz PubsubGetVarz(ctx).Execute()

Server statistics



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
	resp, r, err := apiClient.PubsubMonitoringAPI.PubsubGetVarz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubMonitoringAPI.PubsubGetVarz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetVarz`: PubsubServerVarz
	fmt.Fprintf(os.Stdout, "Response from `PubsubMonitoringAPI.PubsubGetVarz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetVarzRequest struct via the builder pattern


### Return type

[**PubsubServerVarz**](PubsubServerVarz.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

