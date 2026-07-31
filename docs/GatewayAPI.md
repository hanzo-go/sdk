# \GatewayAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1GatewayConfig**](GatewayAPI.md#CloudGetV1GatewayConfig) | **Get** /v1/gateway/config | Read returns the EFFECTIVE edge policy the caller is subject to: the platform CORS allowlist and pre-auth per-IP flood cap in force, plus the caller&#39;s own authenticated rate ceiling, edge-cache TTLs and accepted-method allowlist.
[**CloudPutV1GatewayConfig**](GatewayAPI.md#CloudPutV1GatewayConfig) | **Put** /v1/gateway/config | Write updates one policy scope and returns the policy in force after the write.



## CloudGetV1GatewayConfig

> CloudPolicy CloudGetV1GatewayConfig(ctx).Execute()

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
	resp, r, err := apiClient.GatewayAPI.CloudGetV1GatewayConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.CloudGetV1GatewayConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GatewayConfig`: CloudPolicy
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.CloudGetV1GatewayConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GatewayConfigRequest struct via the builder pattern


### Return type

[**CloudPolicy**](CloudPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1GatewayConfig

> CloudPolicy CloudPutV1GatewayConfig(ctx).CloudPolicy(cloudPolicy).Execute()

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
	cloudPolicy := *openapiclient.NewCloudPolicy() // CloudPolicy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.CloudPutV1GatewayConfig(context.Background()).CloudPolicy(cloudPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.CloudPutV1GatewayConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1GatewayConfig`: CloudPolicy
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.CloudPutV1GatewayConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1GatewayConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPolicy** | [**CloudPolicy**](CloudPolicy.md) |  | 

### Return type

[**CloudPolicy**](CloudPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

