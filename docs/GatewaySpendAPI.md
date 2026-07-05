# \GatewaySpendAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayCalculateSpend**](GatewaySpendAPI.md#GatewayCalculateSpend) | **Post** /v1/gateway/spend/calculate | Calculate spend for request
[**GatewayGetSpendLogs**](GatewaySpendAPI.md#GatewayGetSpendLogs) | **Get** /v1/gateway/spend/logs | Get spend logs



## GatewayCalculateSpend

> GatewayCalculateSpend200Response GatewayCalculateSpend(ctx).GatewayCalculateSpendRequest(gatewayCalculateSpendRequest).Execute()

Calculate spend for request

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
	gatewayCalculateSpendRequest := *openapiclient.NewGatewayCalculateSpendRequest() // GatewayCalculateSpendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaySpendAPI.GatewayCalculateSpend(context.Background()).GatewayCalculateSpendRequest(gatewayCalculateSpendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaySpendAPI.GatewayCalculateSpend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayCalculateSpend`: GatewayCalculateSpend200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewaySpendAPI.GatewayCalculateSpend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCalculateSpendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayCalculateSpendRequest** | [**GatewayCalculateSpendRequest**](GatewayCalculateSpendRequest.md) |  | 

### Return type

[**GatewayCalculateSpend200Response**](GatewayCalculateSpend200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetSpendLogs

> []GatewayGetSpendLogs200ResponseInner GatewayGetSpendLogs(ctx).UserId(userId).TeamId(teamId).StartDate(startDate).EndDate(endDate).Execute()

Get spend logs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	userId := "userId_example" // string |  (optional)
	teamId := "teamId_example" // string |  (optional)
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaySpendAPI.GatewayGetSpendLogs(context.Background()).UserId(userId).TeamId(teamId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaySpendAPI.GatewayGetSpendLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGetSpendLogs`: []GatewayGetSpendLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GatewaySpendAPI.GatewayGetSpendLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetSpendLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **teamId** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

[**[]GatewayGetSpendLogs200ResponseInner**](GatewayGetSpendLogs200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

