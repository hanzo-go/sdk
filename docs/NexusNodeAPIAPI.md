# \NexusNodeAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddNode**](NexusNodeAPIAPI.md#NexusAddNode) | **Post** /v1/nexus/add-node | add Node
[**NexusDeleteNode**](NexusNodeAPIAPI.md#NexusDeleteNode) | **Post** /v1/nexus/delete-node | delete Node
[**NexusGetNode**](NexusNodeAPIAPI.md#NexusGetNode) | **Get** /v1/nexus/get-node | get Node
[**NexusGetNodes**](NexusNodeAPIAPI.md#NexusGetNodes) | **Get** /v1/nexus/get-nodes | get Nodes
[**NexusUpdateNode**](NexusNodeAPIAPI.md#NexusUpdateNode) | **Post** /v1/nexus/update-node | update Node



## NexusAddNode

> NexusResponse NexusAddNode(ctx).NexusNode(nexusNode).Execute()

add Node



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
	nexusNode := *openapiclient.NewNexusNode() // NexusNode | The details of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusNodeAPIAPI.NexusAddNode(context.Background()).NexusNode(nexusNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusNodeAPIAPI.NexusAddNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddNode`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusNodeAPIAPI.NexusAddNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusNode** | [**NexusNode**](NexusNode.md) | The details of the node | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteNode

> NexusResponse NexusDeleteNode(ctx).NexusNode(nexusNode).Execute()

delete Node



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
	nexusNode := *openapiclient.NewNexusNode() // NexusNode | The details of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusNodeAPIAPI.NexusDeleteNode(context.Background()).NexusNode(nexusNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusNodeAPIAPI.NexusDeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteNode`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusNodeAPIAPI.NexusDeleteNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusNode** | [**NexusNode**](NexusNode.md) | The details of the node | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetNode

> NexusNode NexusGetNode(ctx).Id(id).Execute()

get Node



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
	id := "id_example" // string | The id (owner/name) of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusNodeAPIAPI.NexusGetNode(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusNodeAPIAPI.NexusGetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetNode`: NexusNode
	fmt.Fprintf(os.Stdout, "Response from `NexusNodeAPIAPI.NexusGetNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the node | 

### Return type

[**NexusNode**](NexusNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetNodes

> NexusNode NexusGetNodes(ctx).PageSize(pageSize).P(p).Execute()

get Nodes



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
	p := "p_example" // string | The page number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusNodeAPIAPI.NexusGetNodes(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusNodeAPIAPI.NexusGetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetNodes`: NexusNode
	fmt.Fprintf(os.Stdout, "Response from `NexusNodeAPIAPI.NexusGetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The page number | 

### Return type

[**NexusNode**](NexusNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateNode

> NexusResponse NexusUpdateNode(ctx).Id(id).NexusNode(nexusNode).Execute()

update Node



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
	id := "id_example" // string | The id (owner/name) of the node
	nexusNode := *openapiclient.NewNexusNode() // NexusNode | The details of the node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusNodeAPIAPI.NexusUpdateNode(context.Background()).Id(id).NexusNode(nexusNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusNodeAPIAPI.NexusUpdateNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateNode`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusNodeAPIAPI.NexusUpdateNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the node | 
 **nexusNode** | [**NexusNode**](NexusNode.md) | The details of the node | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

