# \StorageProviderAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerGetStorageProviders**](StorageProviderAPIAPI.md#CloudApiControllerGetStorageProviders) | **Get** /v1/cloud/get-storage-providers | Api Controller Get Storage Providers



## CloudApiControllerGetStorageProviders

> []CloudObjectProvider CloudApiControllerGetStorageProviders(ctx).Execute()

Api Controller Get Storage Providers



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
	resp, r, err := apiClient.StorageProviderAPIAPI.CloudApiControllerGetStorageProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageProviderAPIAPI.CloudApiControllerGetStorageProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetStorageProviders`: []CloudObjectProvider
	fmt.Fprintf(os.Stdout, "Response from `StorageProviderAPIAPI.CloudApiControllerGetStorageProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetStorageProvidersRequest struct via the builder pattern


### Return type

[**[]CloudObjectProvider**](CloudObjectProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

