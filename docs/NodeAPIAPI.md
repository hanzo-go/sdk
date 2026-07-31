# \NodeAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddNode**](NodeAPIAPI.md#CloudApiControllerAddNode) | **Post** /v1/cloud/add-node | Api Controller Add Node
[**CloudApiControllerDeleteNode**](NodeAPIAPI.md#CloudApiControllerDeleteNode) | **Post** /v1/cloud/delete-node | Api Controller Delete Node
[**CloudApiControllerGetNode**](NodeAPIAPI.md#CloudApiControllerGetNode) | **Get** /v1/cloud/get-node | Api Controller Get Node
[**CloudApiControllerGetNodes**](NodeAPIAPI.md#CloudApiControllerGetNodes) | **Get** /v1/cloud/get-nodes | Api Controller Get Nodes
[**CloudApiControllerUpdateNode**](NodeAPIAPI.md#CloudApiControllerUpdateNode) | **Post** /v1/cloud/update-node | Api Controller Update Node



## CloudApiControllerAddNode

> CloudControllersResponse CloudApiControllerAddNode(ctx).CloudObjectNode(cloudObjectNode).Execute()

Api Controller Add Node



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
	cloudObjectNode := *openapiclient.NewCloudObjectNode() // CloudObjectNode | The details of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPIAPI.CloudApiControllerAddNode(context.Background()).CloudObjectNode(cloudObjectNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPIAPI.CloudApiControllerAddNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddNode`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeAPIAPI.CloudApiControllerAddNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectNode** | [**CloudObjectNode**](CloudObjectNode.md) | The details of the node | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteNode

> CloudControllersResponse CloudApiControllerDeleteNode(ctx).CloudObjectNode(cloudObjectNode).Execute()

Api Controller Delete Node



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
	cloudObjectNode := *openapiclient.NewCloudObjectNode() // CloudObjectNode | The details of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPIAPI.CloudApiControllerDeleteNode(context.Background()).CloudObjectNode(cloudObjectNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPIAPI.CloudApiControllerDeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteNode`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeAPIAPI.CloudApiControllerDeleteNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectNode** | [**CloudObjectNode**](CloudObjectNode.md) | The details of the node | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetNode

> CloudObjectNode CloudApiControllerGetNode(ctx).Id(id).Execute()

Api Controller Get Node



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
	id := "id_example" // string | The id ( owner/name ) of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPIAPI.CloudApiControllerGetNode(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPIAPI.CloudApiControllerGetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetNode`: CloudObjectNode
	fmt.Fprintf(os.Stdout, "Response from `NodeAPIAPI.CloudApiControllerGetNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the node | 

### Return type

[**CloudObjectNode**](CloudObjectNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetNodes

> CloudObjectNode CloudApiControllerGetNodes(ctx).PageSize(pageSize).P(p).Execute()

Api Controller Get Nodes



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
	pageSize := "pageSize_example" // string | The size of each page
	p := "p_example" // string | The number of the page

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPIAPI.CloudApiControllerGetNodes(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPIAPI.CloudApiControllerGetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetNodes`: CloudObjectNode
	fmt.Fprintf(os.Stdout, "Response from `NodeAPIAPI.CloudApiControllerGetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

### Return type

[**CloudObjectNode**](CloudObjectNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateNode

> CloudControllersResponse CloudApiControllerUpdateNode(ctx).Id(id).CloudObjectNode(cloudObjectNode).Execute()

Api Controller Update Node



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
	id := "id_example" // string | The id ( owner/name ) of the node
	cloudObjectNode := *openapiclient.NewCloudObjectNode() // CloudObjectNode | The details of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPIAPI.CloudApiControllerUpdateNode(context.Background()).Id(id).CloudObjectNode(cloudObjectNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPIAPI.CloudApiControllerUpdateNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateNode`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeAPIAPI.CloudApiControllerUpdateNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the node | 
 **cloudObjectNode** | [**CloudObjectNode**](CloudObjectNode.md) | The details of the node | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

