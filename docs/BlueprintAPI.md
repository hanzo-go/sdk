# \BlueprintAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlueprint**](BlueprintAPI.md#GetBlueprint) | **Get** /v1/blueprint | Returns every deployable blueprint with its service count and estimated monthly compute cost.
[**GetBlueprintHealth**](BlueprintAPI.md#GetBlueprintHealth) | **Get** /v1/blueprint/health | Reports blueprint liveness and echoes the compute rate card in force.
[**GetBlueprintSbom**](BlueprintAPI.md#GetBlueprintSbom) | **Get** /v1/blueprint/sbom | A blueprint&#39;s bill of images and what running it costs



## GetBlueprint

> BlueprintIndex GetBlueprint(ctx).Execute()

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
	resp, r, err := apiClient.BlueprintAPI.GetBlueprint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintAPI.GetBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlueprint`: BlueprintIndex
	fmt.Fprintf(os.Stdout, "Response from `BlueprintAPI.GetBlueprint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlueprintRequest struct via the builder pattern


### Return type

[**BlueprintIndex**](BlueprintIndex.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlueprintHealth

> BlueprintHealth GetBlueprintHealth(ctx).Execute()

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
	resp, r, err := apiClient.BlueprintAPI.GetBlueprintHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintAPI.GetBlueprintHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlueprintHealth`: BlueprintHealth
	fmt.Fprintf(os.Stdout, "Response from `BlueprintAPI.GetBlueprintHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlueprintHealthRequest struct via the builder pattern


### Return type

[**BlueprintHealth**](BlueprintHealth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlueprintSbom

> GetBlueprintSbom(ctx).Execute()

A blueprint's bill of images and what running it costs



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
	r, err := apiClient.BlueprintAPI.GetBlueprintSbom(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintAPI.GetBlueprintSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlueprintSbomRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

