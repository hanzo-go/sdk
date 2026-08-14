# \GatewayAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayTraffic**](GatewayAPI.md#GatewayTraffic) | **Get** /v1/gateway/traffic | Report who is calling this org&#39;s API right now
[**GetGatewayConfig**](GatewayAPI.md#GetGatewayConfig) | **Get** /v1/gateway/config | Read returns the EFFECTIVE edge policy the caller is subject to: the platform CORS allowlist and pre-auth per-IP flood cap in force, plus the caller&#39;s own authenticated rate ceiling, edge-cache TTLs and accepted-method allowlist.
[**PutGatewayConfig**](GatewayAPI.md#PutGatewayConfig) | **Put** /v1/gateway/config | Write updates one policy scope and returns the policy in force after the write.



## GatewayTraffic

> TrafficView GatewayTraffic(ctx).Execute()

Report who is calling this org's API right now



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
	resp, r, err := apiClient.GatewayAPI.GatewayTraffic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GatewayTraffic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayTraffic`: TrafficView
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GatewayTraffic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayTrafficRequest struct via the builder pattern


### Return type

[**TrafficView**](TrafficView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayConfig

> Policy GetGatewayConfig(ctx).Execute()

Read returns the EFFECTIVE edge policy the caller is subject to: the platform CORS allowlist and pre-auth per-IP flood cap in force, plus the caller's own authenticated rate ceiling, edge-cache TTLs and accepted-method allowlist.



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
	resp, r, err := apiClient.GatewayAPI.GetGatewayConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetGatewayConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGatewayConfig`: Policy
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetGatewayConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayConfigRequest struct via the builder pattern


### Return type

[**Policy**](Policy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutGatewayConfig

> Policy PutGatewayConfig(ctx).Policy(policy).Execute()

Write updates one policy scope and returns the policy in force after the write.



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
	policy := *openapiclient.NewPolicy() // Policy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.PutGatewayConfig(context.Background()).Policy(policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.PutGatewayConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutGatewayConfig`: Policy
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.PutGatewayConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutGatewayConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**Policy**](Policy.md) |  | 

### Return type

[**Policy**](Policy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

