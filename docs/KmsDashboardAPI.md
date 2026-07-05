# \KmsDashboardAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsGetDashboardSecretsOverview**](KmsDashboardAPI.md#KmsGetDashboardSecretsOverview) | **Get** /v1/kms/dashboard/{projectId}/secrets-overview | Get secrets overview for the dashboard



## KmsGetDashboardSecretsOverview

> KmsDashboardSecrets KmsGetDashboardSecretsOverview(ctx, projectId).Environments(environments).SecretPath(secretPath).Execute()

Get secrets overview for the dashboard

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	environments := "environments_example" // string |  (optional)
	secretPath := "secretPath_example" // string |  (optional) (default to "/")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsDashboardAPI.KmsGetDashboardSecretsOverview(context.Background(), projectId).Environments(environments).SecretPath(secretPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsDashboardAPI.KmsGetDashboardSecretsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetDashboardSecretsOverview`: KmsDashboardSecrets
	fmt.Fprintf(os.Stdout, "Response from `KmsDashboardAPI.KmsGetDashboardSecretsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetDashboardSecretsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environments** | **string** |  | 
 **secretPath** | **string** |  | [default to &quot;/&quot;]

### Return type

[**KmsDashboardSecrets**](KmsDashboardSecrets.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

