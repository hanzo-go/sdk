# \BlueprintAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Blueprint**](BlueprintAPI.md#CloudGetV1Blueprint) | **Get** /v1/blueprint | Returns every deployable blueprint with its service count and estimated monthly compute cost.
[**CloudGetV1BlueprintHealth**](BlueprintAPI.md#CloudGetV1BlueprintHealth) | **Get** /v1/blueprint/health | Reports blueprint liveness and echoes the compute rate card in force.
[**CloudGetV1BlueprintSbom**](BlueprintAPI.md#CloudGetV1BlueprintSbom) | **Get** /v1/blueprint/sbom | 



## CloudGetV1Blueprint

> CloudBlueprintIndex CloudGetV1Blueprint(ctx).Execute()

Returns every deployable blueprint with its service count and estimated monthly compute cost.



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
	resp, r, err := apiClient.BlueprintAPI.CloudGetV1Blueprint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintAPI.CloudGetV1Blueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Blueprint`: CloudBlueprintIndex
	fmt.Fprintf(os.Stdout, "Response from `BlueprintAPI.CloudGetV1Blueprint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BlueprintRequest struct via the builder pattern


### Return type

[**CloudBlueprintIndex**](CloudBlueprintIndex.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BlueprintHealth

> CloudBlueprintHealth CloudGetV1BlueprintHealth(ctx).Execute()

Reports blueprint liveness and echoes the compute rate card in force.



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
	resp, r, err := apiClient.BlueprintAPI.CloudGetV1BlueprintHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintAPI.CloudGetV1BlueprintHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BlueprintHealth`: CloudBlueprintHealth
	fmt.Fprintf(os.Stdout, "Response from `BlueprintAPI.CloudGetV1BlueprintHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BlueprintHealthRequest struct via the builder pattern


### Return type

[**CloudBlueprintHealth**](CloudBlueprintHealth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1BlueprintSbom

> CloudGetV1BlueprintSbom(ctx).Execute()



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
	r, err := apiClient.BlueprintAPI.CloudGetV1BlueprintSbom(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintAPI.CloudGetV1BlueprintSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BlueprintSbomRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

