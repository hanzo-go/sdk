# \AutoFlowsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoCountFlows**](AutoFlowsAPI.md#AutoCountFlows) | **Get** /v1/auto/flows/count | Count flows
[**AutoCreateFlow**](AutoFlowsAPI.md#AutoCreateFlow) | **Post** /v1/auto/flows | Create a flow
[**AutoDeleteFlow**](AutoFlowsAPI.md#AutoDeleteFlow) | **Delete** /v1/auto/flows/{id} | Delete a flow
[**AutoGetFlow**](AutoFlowsAPI.md#AutoGetFlow) | **Get** /v1/auto/flows/{id} | Get a flow by id
[**AutoGetFlowTemplate**](AutoFlowsAPI.md#AutoGetFlowTemplate) | **Get** /v1/auto/flows/{id}/template | Export flow as template
[**AutoListFlows**](AutoFlowsAPI.md#AutoListFlows) | **Get** /v1/auto/flows | List flows
[**AutoUpdateFlow**](AutoFlowsAPI.md#AutoUpdateFlow) | **Post** /v1/auto/flows/{id} | Apply an operation to a flow



## AutoCountFlows

> map[string]interface{} AutoCountFlows(ctx).FolderId(folderId).Execute()

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
	resp, r, err := apiClient.AutoFlowsAPI.AutoCountFlows(context.Background()).FolderId(folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowsAPI.AutoCountFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCountFlows`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowsAPI.AutoCountFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoCountFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string** |  | 

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


## AutoCreateFlow

> AutoFlow AutoCreateFlow(ctx).AutoCreateFlowRequest(autoCreateFlowRequest).Execute()

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
	resp, r, err := apiClient.AutoFlowsAPI.AutoCreateFlow(context.Background()).AutoCreateFlowRequest(autoCreateFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowsAPI.AutoCreateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreateFlow`: AutoFlow
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowsAPI.AutoCreateFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateFlowRequest** | [**AutoCreateFlowRequest**](AutoCreateFlowRequest.md) |  | 

### Return type

[**AutoFlow**](AutoFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoDeleteFlow

> AutoDeleteFlow(ctx, id).Execute()

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
	r, err := apiClient.AutoFlowsAPI.AutoDeleteFlow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowsAPI.AutoDeleteFlow``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAutoDeleteFlowRequest struct via the builder pattern


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


## AutoGetFlow

> AutoFlow AutoGetFlow(ctx, id).VersionId(versionId).Execute()

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
	resp, r, err := apiClient.AutoFlowsAPI.AutoGetFlow(context.Background(), id).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowsAPI.AutoGetFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoGetFlow`: AutoFlow
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowsAPI.AutoGetFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoGetFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionId** | **string** |  | 

### Return type

[**AutoFlow**](AutoFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoGetFlowTemplate

> map[string]interface{} AutoGetFlowTemplate(ctx, id).Execute()

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
	resp, r, err := apiClient.AutoFlowsAPI.AutoGetFlowTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowsAPI.AutoGetFlowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoGetFlowTemplate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowsAPI.AutoGetFlowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoGetFlowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AutoListFlows

> AutoListFlows200Response AutoListFlows(ctx).FolderId(folderId).Status(status).Name(name).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.AutoFlowsAPI.AutoListFlows(context.Background()).FolderId(folderId).Status(status).Name(name).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowsAPI.AutoListFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListFlows`: AutoListFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowsAPI.AutoListFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string** |  | 
 **status** | **string** |  | 
 **name** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 10]

### Return type

[**AutoListFlows200Response**](AutoListFlows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoUpdateFlow

> AutoFlow AutoUpdateFlow(ctx, id).AutoUpdateFlowRequest(autoUpdateFlowRequest).Execute()

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
	resp, r, err := apiClient.AutoFlowsAPI.AutoUpdateFlow(context.Background(), id).AutoUpdateFlowRequest(autoUpdateFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowsAPI.AutoUpdateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoUpdateFlow`: AutoFlow
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowsAPI.AutoUpdateFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoUpdateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoUpdateFlowRequest** | [**AutoUpdateFlowRequest**](AutoUpdateFlowRequest.md) |  | 

### Return type

[**AutoFlow**](AutoFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

