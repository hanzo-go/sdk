# \FlowOauthAppsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowDeleteOAuthApp**](FlowOauthAppsAPI.md#FlowDeleteOAuthApp) | **Delete** /v1/flow/oauth-apps/{id} | Delete an OAuth app (EE)
[**FlowListOAuthApps**](FlowOauthAppsAPI.md#FlowListOAuthApps) | **Get** /v1/flow/oauth-apps | List OAuth app configurations (EE)
[**FlowUpsertOAuthApp**](FlowOauthAppsAPI.md#FlowUpsertOAuthApp) | **Post** /v1/flow/oauth-apps | Upsert an OAuth app (EE)



## FlowDeleteOAuthApp

> FlowDeleteOAuthApp(ctx, id).Execute()

Delete an OAuth app (EE)

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowOauthAppsAPI.FlowDeleteOAuthApp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowOauthAppsAPI.FlowDeleteOAuthApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowDeleteOAuthAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## FlowListOAuthApps

> map[string]interface{} FlowListOAuthApps(ctx).Execute()

List OAuth app configurations (EE)

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
	resp, r, err := apiClient.FlowOauthAppsAPI.FlowListOAuthApps(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowOauthAppsAPI.FlowListOAuthApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListOAuthApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowOauthAppsAPI.FlowListOAuthApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListOAuthAppsRequest struct via the builder pattern


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


## FlowUpsertOAuthApp

> map[string]interface{} FlowUpsertOAuthApp(ctx).FlowUpsertOAuthAppRequest(flowUpsertOAuthAppRequest).Execute()

Upsert an OAuth app (EE)

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
	flowUpsertOAuthAppRequest := *openapiclient.NewFlowUpsertOAuthAppRequest("PieceName_example", "ClientId_example") // FlowUpsertOAuthAppRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowOauthAppsAPI.FlowUpsertOAuthApp(context.Background()).FlowUpsertOAuthAppRequest(flowUpsertOAuthAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowOauthAppsAPI.FlowUpsertOAuthApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpsertOAuthApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowOauthAppsAPI.FlowUpsertOAuthApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpsertOAuthAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowUpsertOAuthAppRequest** | [**FlowUpsertOAuthAppRequest**](FlowUpsertOAuthAppRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

