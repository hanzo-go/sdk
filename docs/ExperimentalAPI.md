# \ExperimentalAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchGetExperimentalFeatures**](ExperimentalAPI.md#SearchGetExperimentalFeatures) | **Get** /v1/search/experimental-features | Get runtime-togglable experimental features
[**SearchUpdateExperimentalFeatures**](ExperimentalAPI.md#SearchUpdateExperimentalFeatures) | **Patch** /v1/search/experimental-features | Toggle experimental features



## SearchGetExperimentalFeatures

> SearchRuntimeTogglableFeatures SearchGetExperimentalFeatures(ctx).Execute()

Get runtime-togglable experimental features

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
	resp, r, err := apiClient.ExperimentalAPI.SearchGetExperimentalFeatures(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.SearchGetExperimentalFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetExperimentalFeatures`: SearchRuntimeTogglableFeatures
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.SearchGetExperimentalFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetExperimentalFeaturesRequest struct via the builder pattern


### Return type

[**SearchRuntimeTogglableFeatures**](SearchRuntimeTogglableFeatures.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateExperimentalFeatures

> SearchRuntimeTogglableFeatures SearchUpdateExperimentalFeatures(ctx).SearchRuntimeTogglableFeatures(searchRuntimeTogglableFeatures).Execute()

Toggle experimental features

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
	searchRuntimeTogglableFeatures := *openapiclient.NewSearchRuntimeTogglableFeatures() // SearchRuntimeTogglableFeatures | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.SearchUpdateExperimentalFeatures(context.Background()).SearchRuntimeTogglableFeatures(searchRuntimeTogglableFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.SearchUpdateExperimentalFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateExperimentalFeatures`: SearchRuntimeTogglableFeatures
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.SearchUpdateExperimentalFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateExperimentalFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRuntimeTogglableFeatures** | [**SearchRuntimeTogglableFeatures**](SearchRuntimeTogglableFeatures.md) |  | 

### Return type

[**SearchRuntimeTogglableFeatures**](SearchRuntimeTogglableFeatures.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

