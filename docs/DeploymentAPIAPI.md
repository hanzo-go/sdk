# \DeploymentAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerGetK8sStatus**](DeploymentAPIAPI.md#CloudApiControllerGetK8sStatus) | **Get** /v1/cloud/get-k8s-status | Api Controller Get K8s Status



## CloudApiControllerGetK8sStatus

> map[string]interface{} CloudApiControllerGetK8sStatus(ctx).Execute()

Api Controller Get K8s Status



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
	resp, r, err := apiClient.DeploymentAPIAPI.CloudApiControllerGetK8sStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPIAPI.CloudApiControllerGetK8sStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetK8sStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPIAPI.CloudApiControllerGetK8sStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetK8sStatusRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

