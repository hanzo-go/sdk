# \MeshAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1MeshServices**](MeshAPI.md#CloudGetV1MeshServices) | **Get** /v1/mesh/services | Returns the Zero Trust edge services the caller&#39;s org owns.



## CloudGetV1MeshServices

> CloudMeshServiceList CloudGetV1MeshServices(ctx).Execute()

Returns the Zero Trust edge services the caller's org owns.



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
	resp, r, err := apiClient.MeshAPI.CloudGetV1MeshServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeshAPI.CloudGetV1MeshServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MeshServices`: CloudMeshServiceList
	fmt.Fprintf(os.Stdout, "Response from `MeshAPI.CloudGetV1MeshServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MeshServicesRequest struct via the builder pattern


### Return type

[**CloudMeshServiceList**](CloudMeshServiceList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

