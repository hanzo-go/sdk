# \OperativeConfigurationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperativeGetEnvConfig**](OperativeConfigurationAPI.md#OperativeGetEnvConfig) | **Get** /v1/operative/env-config.json | Get environment configuration



## OperativeGetEnvConfig

> OperativeEnvConfig OperativeGetEnvConfig(ctx).Execute()

Get environment configuration



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
	resp, r, err := apiClient.OperativeConfigurationAPI.OperativeGetEnvConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeConfigurationAPI.OperativeGetEnvConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeGetEnvConfig`: OperativeEnvConfig
	fmt.Fprintf(os.Stdout, "Response from `OperativeConfigurationAPI.OperativeGetEnvConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeGetEnvConfigRequest struct via the builder pattern


### Return type

[**OperativeEnvConfig**](OperativeEnvConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

