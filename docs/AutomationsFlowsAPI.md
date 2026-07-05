# \AutomationsFlowsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationsApplyOperation**](AutomationsFlowsAPI.md#AutomationsApplyOperation) | **Post** /v1/automations/flows/{id}/operations | Apply a flow operation to the latest version
[**AutomationsCreateFlow**](AutomationsFlowsAPI.md#AutomationsCreateFlow) | **Post** /v1/automations/flows | Create a flow (with an initial draft version)
[**AutomationsCreateVersion**](AutomationsFlowsAPI.md#AutomationsCreateVersion) | **Post** /v1/automations/flows/{id}/versions | Create a draft version
[**AutomationsDeleteFlow**](AutomationsFlowsAPI.md#AutomationsDeleteFlow) | **Delete** /v1/automations/flows/{id} | Delete a flow (with its versions and runs)
[**AutomationsDisableFlow**](AutomationsFlowsAPI.md#AutomationsDisableFlow) | **Post** /v1/automations/flows/{id}/disable | Disable a flow (removes any POLLING schedule)
[**AutomationsEnableFlow**](AutomationsFlowsAPI.md#AutomationsEnableFlow) | **Post** /v1/automations/flows/{id}/enable | Enable a flow (POLLING triggers create a schedule)
[**AutomationsGetFlow**](AutomationsFlowsAPI.md#AutomationsGetFlow) | **Get** /v1/automations/flows/{id} | Get a flow and its latest version
[**AutomationsListFlows**](AutomationsFlowsAPI.md#AutomationsListFlows) | **Get** /v1/automations/flows | List flows
[**AutomationsListVersions**](AutomationsFlowsAPI.md#AutomationsListVersions) | **Get** /v1/automations/flows/{id}/versions | List a flow&#39;s versions
[**AutomationsUpdateFlow**](AutomationsFlowsAPI.md#AutomationsUpdateFlow) | **Patch** /v1/automations/flows/{id} | Update flow metadata



## AutomationsApplyOperation

> AutomationsFlowVersion AutomationsApplyOperation(ctx, id).AutomationsFlowOperation(automationsFlowOperation).Execute()

Apply a flow operation to the latest version



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
	automationsFlowOperation := *openapiclient.NewAutomationsFlowOperation("Type_example", interface{}(123)) // AutomationsFlowOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsApplyOperation(context.Background(), id).AutomationsFlowOperation(automationsFlowOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsApplyOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsApplyOperation`: AutomationsFlowVersion
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsApplyOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsApplyOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **automationsFlowOperation** | [**AutomationsFlowOperation**](AutomationsFlowOperation.md) |  | 

### Return type

[**AutomationsFlowVersion**](AutomationsFlowVersion.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsCreateFlow

> AutomationsPopulatedFlow AutomationsCreateFlow(ctx).AutomationsCreateFlowRequest(automationsCreateFlowRequest).Execute()

Create a flow (with an initial draft version)

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
	automationsCreateFlowRequest := *openapiclient.NewAutomationsCreateFlowRequest() // AutomationsCreateFlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsCreateFlow(context.Background()).AutomationsCreateFlowRequest(automationsCreateFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsCreateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsCreateFlow`: AutomationsPopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsCreateFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsCreateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationsCreateFlowRequest** | [**AutomationsCreateFlowRequest**](AutomationsCreateFlowRequest.md) |  | 

### Return type

[**AutomationsPopulatedFlow**](AutomationsPopulatedFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsCreateVersion

> AutomationsFlowVersion AutomationsCreateVersion(ctx, id).AutomationsCreateVersionRequest(automationsCreateVersionRequest).Execute()

Create a draft version

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
	automationsCreateVersionRequest := *openapiclient.NewAutomationsCreateVersionRequest() // AutomationsCreateVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsCreateVersion(context.Background(), id).AutomationsCreateVersionRequest(automationsCreateVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsCreateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsCreateVersion`: AutomationsFlowVersion
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsCreateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **automationsCreateVersionRequest** | [**AutomationsCreateVersionRequest**](AutomationsCreateVersionRequest.md) |  | 

### Return type

[**AutomationsFlowVersion**](AutomationsFlowVersion.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsDeleteFlow

> AutomationsDeleteFlow(ctx, id).Execute()

Delete a flow (with its versions and runs)

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
	r, err := apiClient.AutomationsFlowsAPI.AutomationsDeleteFlow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsDeleteFlow``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAutomationsDeleteFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsDisableFlow

> AutomationsFlow AutomationsDisableFlow(ctx, id).Execute()

Disable a flow (removes any POLLING schedule)

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
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsDisableFlow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsDisableFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsDisableFlow`: AutomationsFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsDisableFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsDisableFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationsFlow**](AutomationsFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsEnableFlow

> AutomationsFlow AutomationsEnableFlow(ctx, id).Execute()

Enable a flow (POLLING triggers create a schedule)

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
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsEnableFlow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsEnableFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsEnableFlow`: AutomationsFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsEnableFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsEnableFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationsFlow**](AutomationsFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsGetFlow

> AutomationsPopulatedFlow AutomationsGetFlow(ctx, id).Execute()

Get a flow and its latest version

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
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsGetFlow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsGetFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsGetFlow`: AutomationsPopulatedFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsGetFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsGetFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationsPopulatedFlow**](AutomationsPopulatedFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsListFlows

> AutomationsListFlows200Response AutomationsListFlows(ctx).Limit(limit).Execute()

List flows

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
	limit := int32(56) // int32 |  (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsListFlows(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsListFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsListFlows`: AutomationsListFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsListFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsListFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 200]

### Return type

[**AutomationsListFlows200Response**](AutomationsListFlows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsListVersions

> AutomationsListVersions200Response AutomationsListVersions(ctx, id).Limit(limit).Execute()

List a flow's versions

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
	limit := int32(56) // int32 |  (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsListVersions(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsListVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsListVersions`: AutomationsListVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 200]

### Return type

[**AutomationsListVersions200Response**](AutomationsListVersions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsUpdateFlow

> AutomationsFlow AutomationsUpdateFlow(ctx, id).AutomationsPatchFlowRequest(automationsPatchFlowRequest).Execute()

Update flow metadata

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
	automationsPatchFlowRequest := *openapiclient.NewAutomationsPatchFlowRequest() // AutomationsPatchFlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsFlowsAPI.AutomationsUpdateFlow(context.Background(), id).AutomationsPatchFlowRequest(automationsPatchFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsFlowsAPI.AutomationsUpdateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsUpdateFlow`: AutomationsFlow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsFlowsAPI.AutomationsUpdateFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsUpdateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **automationsPatchFlowRequest** | [**AutomationsPatchFlowRequest**](AutomationsPatchFlowRequest.md) |  | 

### Return type

[**AutomationsFlow**](AutomationsFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

