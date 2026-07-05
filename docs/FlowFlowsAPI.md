# \FlowFlowsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowCountFlows**](FlowFlowsAPI.md#FlowCountFlows) | **Get** /v1/flow/flows/count | Count flows
[**FlowCreateFlow**](FlowFlowsAPI.md#FlowCreateFlow) | **Post** /v1/flow/flows | Create a flow
[**FlowDeleteFlow**](FlowFlowsAPI.md#FlowDeleteFlow) | **Delete** /v1/flow/flows/{id} | Delete a flow
[**FlowGetFlow**](FlowFlowsAPI.md#FlowGetFlow) | **Get** /v1/flow/flows/{id} | Get a flow by id
[**FlowGetFlowTemplate**](FlowFlowsAPI.md#FlowGetFlowTemplate) | **Get** /v1/flow/flows/{id}/template | Export flow as template
[**FlowGetHumanInputForm**](FlowFlowsAPI.md#FlowGetHumanInputForm) | **Get** /v1/flow/human-input/form/{flowId} | Get human input form definition for a flow
[**FlowGetStepFile**](FlowFlowsAPI.md#FlowGetStepFile) | **Get** /v1/flow/step-files | Get a file produced by a flow step
[**FlowListFlows**](FlowFlowsAPI.md#FlowListFlows) | **Get** /v1/flow/flows | List flows
[**FlowUpdateFlow**](FlowFlowsAPI.md#FlowUpdateFlow) | **Post** /v1/flow/flows/{id} | Apply an operation to a flow



## FlowCountFlows

> FlowCountFlows200Response FlowCountFlows(ctx).FolderId(folderId).Execute()

Count flows

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
	folderId := "folderId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowsAPI.FlowCountFlows(context.Background()).FolderId(folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowCountFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCountFlows`: FlowCountFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowsAPI.FlowCountFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCountFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string** |  | 

### Return type

[**FlowCountFlows200Response**](FlowCountFlows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowCreateFlow

> FlowFlow FlowCreateFlow(ctx).AutoCreateFlowRequest(autoCreateFlowRequest).Execute()

Create a flow

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
	autoCreateFlowRequest := *openapiclient.NewAutoCreateFlowRequest("DisplayName_example") // AutoCreateFlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowsAPI.FlowCreateFlow(context.Background()).AutoCreateFlowRequest(autoCreateFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowCreateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateFlow`: FlowFlow
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowsAPI.FlowCreateFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateFlowRequest** | [**AutoCreateFlowRequest**](AutoCreateFlowRequest.md) |  | 

### Return type

[**FlowFlow**](FlowFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowDeleteFlow

> FlowDeleteFlow(ctx, id).Execute()

Delete a flow

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
	r, err := apiClient.FlowFlowsAPI.FlowDeleteFlow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowDeleteFlow``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowDeleteFlowRequest struct via the builder pattern


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


## FlowGetFlow

> FlowFlow FlowGetFlow(ctx, id).VersionId(versionId).Execute()

Get a flow by id

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
	versionId := "versionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowsAPI.FlowGetFlow(context.Background(), id).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowGetFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetFlow`: FlowFlow
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowsAPI.FlowGetFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionId** | **string** |  | 

### Return type

[**FlowFlow**](FlowFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetFlowTemplate

> FlowTemplate FlowGetFlowTemplate(ctx, id).Execute()

Export flow as template

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
	resp, r, err := apiClient.FlowFlowsAPI.FlowGetFlowTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowGetFlowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetFlowTemplate`: FlowTemplate
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowsAPI.FlowGetFlowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetFlowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowTemplate**](FlowTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetHumanInputForm

> map[string]interface{} FlowGetHumanInputForm(ctx, flowId).Execute()

Get human input form definition for a flow

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
	flowId := "flowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowsAPI.FlowGetHumanInputForm(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowGetHumanInputForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetHumanInputForm`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowsAPI.FlowGetHumanInputForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetHumanInputFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetStepFile

> *os.File FlowGetStepFile(ctx).Id(id).Execute()

Get a file produced by a flow step

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
	resp, r, err := apiClient.FlowFlowsAPI.FlowGetStepFile(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowGetStepFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetStepFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowsAPI.FlowGetStepFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetStepFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowListFlows

> FlowListFlows200Response FlowListFlows(ctx).FolderId(folderId).Status(status).Name(name).Cursor(cursor).Limit(limit).Execute()

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
	folderId := "folderId_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowsAPI.FlowListFlows(context.Background()).FolderId(folderId).Status(status).Name(name).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowListFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListFlows`: FlowListFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowsAPI.FlowListFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string** |  | 
 **status** | **string** |  | 
 **name** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 10]

### Return type

[**FlowListFlows200Response**](FlowListFlows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowUpdateFlow

> FlowFlow FlowUpdateFlow(ctx, id).AutoUpdateFlowRequest(autoUpdateFlowRequest).Execute()

Apply an operation to a flow

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
	autoUpdateFlowRequest := *openapiclient.NewAutoUpdateFlowRequest("Type_example", map[string]interface{}(123)) // AutoUpdateFlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowsAPI.FlowUpdateFlow(context.Background(), id).AutoUpdateFlowRequest(autoUpdateFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowsAPI.FlowUpdateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpdateFlow`: FlowFlow
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowsAPI.FlowUpdateFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpdateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoUpdateFlowRequest** | [**AutoUpdateFlowRequest**](AutoUpdateFlowRequest.md) |  | 

### Return type

[**FlowFlow**](FlowFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

