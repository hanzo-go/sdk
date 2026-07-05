# \PlatformDeployWorkerAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformDeployWorkerCancelDeployment**](PlatformDeployWorkerAPI.md#PlatformDeployWorkerCancelDeployment) | **Post** /v1/platform/cancel-deployment | Cancel an in-progress worker deployment
[**PlatformDeployWorkerDeploy**](PlatformDeployWorkerAPI.md#PlatformDeployWorkerDeploy) | **Post** /v1/platform/deploy | Submit deployment job to worker queue



## PlatformDeployWorkerCancelDeployment

> PlatformDeployWorkerCancelDeployment200Response PlatformDeployWorkerCancelDeployment(ctx).PlatformCancelDeploymentJob(platformCancelDeploymentJob).Execute()

Cancel an in-progress worker deployment

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
	platformCancelDeploymentJob := openapiclient.platform_CancelDeploymentJob{PlatformCancelDeploymentJobApplication: openapiclient.NewPlatformCancelDeploymentJobApplication("ApplicationId_example", "ApplicationType_example")} // PlatformCancelDeploymentJob | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDeployWorkerAPI.PlatformDeployWorkerCancelDeployment(context.Background()).PlatformCancelDeploymentJob(platformCancelDeploymentJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeployWorkerAPI.PlatformDeployWorkerCancelDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDeployWorkerCancelDeployment`: PlatformDeployWorkerCancelDeployment200Response
	fmt.Fprintf(os.Stdout, "Response from `PlatformDeployWorkerAPI.PlatformDeployWorkerCancelDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeployWorkerCancelDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformCancelDeploymentJob** | [**PlatformCancelDeploymentJob**](PlatformCancelDeploymentJob.md) |  | 

### Return type

[**PlatformDeployWorkerCancelDeployment200Response**](PlatformDeployWorkerCancelDeployment200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformDeployWorkerDeploy

> PlatformDeployWorkerDeploy200Response PlatformDeployWorkerDeploy(ctx).PlatformDeployJob(platformDeployJob).Execute()

Submit deployment job to worker queue

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
	platformDeployJob := openapiclient.platform_DeployJob{PlatformDeployJobApplication: openapiclient.NewPlatformDeployJobApplication("ApplicationId_example", "Type_example", "ApplicationType_example", "ServerId_example")} // PlatformDeployJob | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDeployWorkerAPI.PlatformDeployWorkerDeploy(context.Background()).PlatformDeployJob(platformDeployJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeployWorkerAPI.PlatformDeployWorkerDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDeployWorkerDeploy`: PlatformDeployWorkerDeploy200Response
	fmt.Fprintf(os.Stdout, "Response from `PlatformDeployWorkerAPI.PlatformDeployWorkerDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeployWorkerDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDeployJob** | [**PlatformDeployJob**](PlatformDeployJob.md) |  | 

### Return type

[**PlatformDeployWorkerDeploy200Response**](PlatformDeployWorkerDeploy200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

