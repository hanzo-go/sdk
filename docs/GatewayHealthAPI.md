# \GatewayHealthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayLivelinessCheck**](GatewayHealthAPI.md#GatewayLivelinessCheck) | **Get** /v1/gateway/health/liveliness | Liveliness check
[**GatewayReadinessCheck**](GatewayHealthAPI.md#GatewayReadinessCheck) | **Get** /v1/gateway/health/readiness | Readiness check



## GatewayLivelinessCheck

> string GatewayLivelinessCheck(ctx).Execute()

Liveliness check

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
	resp, r, err := apiClient.GatewayHealthAPI.GatewayLivelinessCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayHealthAPI.GatewayLivelinessCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayLivelinessCheck`: string
	fmt.Fprintf(os.Stdout, "Response from `GatewayHealthAPI.GatewayLivelinessCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayLivelinessCheckRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayReadinessCheck

> GatewayReadinessCheck200Response GatewayReadinessCheck(ctx).Execute()

Readiness check

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
	resp, r, err := apiClient.GatewayHealthAPI.GatewayReadinessCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayHealthAPI.GatewayReadinessCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayReadinessCheck`: GatewayReadinessCheck200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayHealthAPI.GatewayReadinessCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayReadinessCheckRequest struct via the builder pattern


### Return type

[**GatewayReadinessCheck200Response**](GatewayReadinessCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

