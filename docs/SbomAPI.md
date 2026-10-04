# \SbomAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSbomHealth**](SbomAPI.md#GetSbomHealth) | **Get** /v1/sbom/health | Health is a pure liveness probe: the service is up; datastore reflects whether the datastore store is connected.



## GetSbomHealth

> SbomSbomHealth GetSbomHealth(ctx).Execute()

Health is a pure liveness probe: the service is up; datastore reflects whether the datastore store is connected.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SbomAPI.GetSbomHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbomHealth`: SbomSbomHealth
	fmt.Fprintf(os.Stdout, "Response from `SbomAPI.GetSbomHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomHealthRequest struct via the builder pattern


### Return type

[**SbomSbomHealth**](SbomSbomHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

