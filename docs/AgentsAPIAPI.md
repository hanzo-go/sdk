# \AgentsAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudAgentsControllerActivity**](AgentsAPIAPI.md#CloudAgentsControllerActivity) | **Get** /v1/agents/activity | 
[**CloudAgentsControllerCreate**](AgentsAPIAPI.md#CloudAgentsControllerCreate) | **Post** /v1/agents | 
[**CloudAgentsControllerDelete**](AgentsAPIAPI.md#CloudAgentsControllerDelete) | **Delete** /v1/agents/{ref} | 
[**CloudAgentsControllerGet**](AgentsAPIAPI.md#CloudAgentsControllerGet) | **Get** /v1/agents/{ref} | 
[**CloudAgentsControllerList**](AgentsAPIAPI.md#CloudAgentsControllerList) | **Get** /v1/agents | 
[**CloudAgentsControllerMetrics**](AgentsAPIAPI.md#CloudAgentsControllerMetrics) | **Get** /v1/agents/metrics | 
[**CloudAgentsControllerRun**](AgentsAPIAPI.md#CloudAgentsControllerRun) | **Post** /v1/agents/{ref}/run | 
[**CloudAgentsControllerRuns**](AgentsAPIAPI.md#CloudAgentsControllerRuns) | **Get** /v1/agents/{ref}/runs | 
[**CloudAgentsControllerUpdate**](AgentsAPIAPI.md#CloudAgentsControllerUpdate) | **Patch** /v1/agents/{ref} | 



## CloudAgentsControllerActivity

> CloudAgentsControllerActivity200Response CloudAgentsControllerActivity(ctx).Execute()





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
	resp, r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerActivity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentsControllerActivity`: CloudAgentsControllerActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPIAPI.CloudAgentsControllerActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerActivityRequest struct via the builder pattern


### Return type

[**CloudAgentsControllerActivity200Response**](CloudAgentsControllerActivity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentsControllerCreate

> CloudAgentsAgent CloudAgentsControllerCreate(ctx).CloudAgentsCreateAgentRequest(cloudAgentsCreateAgentRequest).Execute()





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
	cloudAgentsCreateAgentRequest := *openapiclient.NewCloudAgentsCreateAgentRequest("Name_example", "Model_example") // CloudAgentsCreateAgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerCreate(context.Background()).CloudAgentsCreateAgentRequest(cloudAgentsCreateAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentsControllerCreate`: CloudAgentsAgent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPIAPI.CloudAgentsControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAgentsCreateAgentRequest** | [**CloudAgentsCreateAgentRequest**](CloudAgentsCreateAgentRequest.md) |  | 

### Return type

[**CloudAgentsAgent**](CloudAgentsAgent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentsControllerDelete

> CloudAgentsControllerDelete(ctx, ref).Execute()





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
	ref := "ref_example" // string | The agent's public id (agent_...) or org-unique name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerDelete(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | The agent&#39;s public id (agent_...) or org-unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerDeleteRequest struct via the builder pattern


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


## CloudAgentsControllerGet

> CloudAgentsAgentDetail CloudAgentsControllerGet(ctx, ref).Execute()





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
	ref := "ref_example" // string | The agent's public id (agent_...) or org-unique name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerGet(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentsControllerGet`: CloudAgentsAgentDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPIAPI.CloudAgentsControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | The agent&#39;s public id (agent_...) or org-unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAgentsAgentDetail**](CloudAgentsAgentDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentsControllerList

> CloudAgentsControllerList200Response CloudAgentsControllerList(ctx).Execute()





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
	resp, r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentsControllerList`: CloudAgentsControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPIAPI.CloudAgentsControllerList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerListRequest struct via the builder pattern


### Return type

[**CloudAgentsControllerList200Response**](CloudAgentsControllerList200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentsControllerMetrics

> CloudAgentsMetrics CloudAgentsControllerMetrics(ctx).Range_(range_).Execute()





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
	range_ := "range__example" // string | Window token; one of 24H, 7D, 30D (default 30D). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentsControllerMetrics`: CloudAgentsMetrics
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPIAPI.CloudAgentsControllerMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Window token; one of 24H, 7D, 30D (default 30D). | 

### Return type

[**CloudAgentsMetrics**](CloudAgentsMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentsControllerRun

> CloudAgentsRun CloudAgentsControllerRun(ctx, ref).CloudAgentsRunRequest(cloudAgentsRunRequest).Execute()





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
	ref := "ref_example" // string | The agent's public id (agent_...) or org-unique name.
	cloudAgentsRunRequest := *openapiclient.NewCloudAgentsRunRequest() // CloudAgentsRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerRun(context.Background(), ref).CloudAgentsRunRequest(cloudAgentsRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentsControllerRun`: CloudAgentsRun
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPIAPI.CloudAgentsControllerRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | The agent&#39;s public id (agent_...) or org-unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAgentsRunRequest** | [**CloudAgentsRunRequest**](CloudAgentsRunRequest.md) |  | 

### Return type

[**CloudAgentsRun**](CloudAgentsRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentsControllerRuns

> CloudAgentsControllerRuns200Response CloudAgentsControllerRuns(ctx, ref).Limit(limit).Execute()





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
	ref := "ref_example" // string | The agent's public id (agent_...) or org-unique name.
	limit := int32(56) // int32 | Max runs to return (default 50, max 200). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerRuns(context.Background(), ref).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentsControllerRuns`: CloudAgentsControllerRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPIAPI.CloudAgentsControllerRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | The agent&#39;s public id (agent_...) or org-unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max runs to return (default 50, max 200). | 

### Return type

[**CloudAgentsControllerRuns200Response**](CloudAgentsControllerRuns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentsControllerUpdate

> CloudAgentsAgent CloudAgentsControllerUpdate(ctx, ref).CloudAgentsUpdateAgentRequest(cloudAgentsUpdateAgentRequest).Execute()





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
	ref := "ref_example" // string | The agent's public id (agent_...) or org-unique name.
	cloudAgentsUpdateAgentRequest := *openapiclient.NewCloudAgentsUpdateAgentRequest() // CloudAgentsUpdateAgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPIAPI.CloudAgentsControllerUpdate(context.Background(), ref).CloudAgentsUpdateAgentRequest(cloudAgentsUpdateAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPIAPI.CloudAgentsControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentsControllerUpdate`: CloudAgentsAgent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPIAPI.CloudAgentsControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | The agent&#39;s public id (agent_...) or org-unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentsControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAgentsUpdateAgentRequest** | [**CloudAgentsUpdateAgentRequest**](CloudAgentsUpdateAgentRequest.md) |  | 

### Return type

[**CloudAgentsAgent**](CloudAgentsAgent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

