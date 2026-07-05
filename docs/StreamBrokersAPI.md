# \StreamBrokersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamGetBrokerConfig**](StreamBrokersAPI.md#StreamGetBrokerConfig) | **Get** /v1/stream/brokers/config | Get broker configuration
[**StreamListBrokers**](StreamBrokersAPI.md#StreamListBrokers) | **Get** /v1/stream/brokers | List broker instances



## StreamGetBrokerConfig

> StreamGetBrokerConfig200Response StreamGetBrokerConfig(ctx).Execute()

Get broker configuration

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
	resp, r, err := apiClient.StreamBrokersAPI.StreamGetBrokerConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamBrokersAPI.StreamGetBrokerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamGetBrokerConfig`: StreamGetBrokerConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `StreamBrokersAPI.StreamGetBrokerConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStreamGetBrokerConfigRequest struct via the builder pattern


### Return type

[**StreamGetBrokerConfig200Response**](StreamGetBrokerConfig200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamListBrokers

> StreamListBrokers200Response StreamListBrokers(ctx).Execute()

List broker instances



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
	resp, r, err := apiClient.StreamBrokersAPI.StreamListBrokers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamBrokersAPI.StreamListBrokers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamListBrokers`: StreamListBrokers200Response
	fmt.Fprintf(os.Stdout, "Response from `StreamBrokersAPI.StreamListBrokers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStreamListBrokersRequest struct via the builder pattern


### Return type

[**StreamListBrokers200Response**](StreamListBrokers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

