# \NexusStorageProviderAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusGetStorageProviders**](NexusStorageProviderAPIAPI.md#NexusGetStorageProviders) | **Get** /v1/nexus/get-storage-providers | get Storage Providers



## NexusGetStorageProviders

> []NexusProvider NexusGetStorageProviders(ctx).Execute()

get Storage Providers



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
	resp, r, err := apiClient.NexusStorageProviderAPIAPI.NexusGetStorageProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusStorageProviderAPIAPI.NexusGetStorageProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetStorageProviders`: []NexusProvider
	fmt.Fprintf(os.Stdout, "Response from `NexusStorageProviderAPIAPI.NexusGetStorageProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetStorageProvidersRequest struct via the builder pattern


### Return type

[**[]NexusProvider**](NexusProvider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

