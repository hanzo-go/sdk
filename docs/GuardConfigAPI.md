# \GuardConfigAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GuardGetConfig**](GuardConfigAPI.md#GuardGetConfig) | **Get** /v1/guard/config | Get current configuration
[**GuardUpdateConfig**](GuardConfigAPI.md#GuardUpdateConfig) | **Put** /v1/guard/config | Update configuration



## GuardGetConfig

> GuardSanitizeConfig GuardGetConfig(ctx).Execute()

Get current configuration

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
	resp, r, err := apiClient.GuardConfigAPI.GuardGetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardConfigAPI.GuardGetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GuardGetConfig`: GuardSanitizeConfig
	fmt.Fprintf(os.Stdout, "Response from `GuardConfigAPI.GuardGetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGuardGetConfigRequest struct via the builder pattern


### Return type

[**GuardSanitizeConfig**](GuardSanitizeConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GuardUpdateConfig

> GuardUpdateConfig(ctx).GuardSanitizeConfig(guardSanitizeConfig).Execute()

Update configuration



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
	guardSanitizeConfig := *openapiclient.NewGuardSanitizeConfig() // GuardSanitizeConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GuardConfigAPI.GuardUpdateConfig(context.Background()).GuardSanitizeConfig(guardSanitizeConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardConfigAPI.GuardUpdateConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guardSanitizeConfig** | [**GuardSanitizeConfig**](GuardSanitizeConfig.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

