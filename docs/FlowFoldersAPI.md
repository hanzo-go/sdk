# \FlowFoldersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowCreateFolder**](FlowFoldersAPI.md#FlowCreateFolder) | **Post** /v1/flow/folders | Create a folder
[**FlowDeleteFolder**](FlowFoldersAPI.md#FlowDeleteFolder) | **Delete** /v1/flow/folders/{id} | Delete a folder
[**FlowGetFolder**](FlowFoldersAPI.md#FlowGetFolder) | **Get** /v1/flow/folders/{id} | Get a folder by id
[**FlowListFolders**](FlowFoldersAPI.md#FlowListFolders) | **Get** /v1/flow/folders | List folders
[**FlowUpdateFolder**](FlowFoldersAPI.md#FlowUpdateFolder) | **Post** /v1/flow/folders/{id} | Update a folder



## FlowCreateFolder

> FlowFolder FlowCreateFolder(ctx).AutoCreateApiKeyRequest(autoCreateApiKeyRequest).Execute()

Create a folder

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
	autoCreateApiKeyRequest := *openapiclient.NewAutoCreateApiKeyRequest("DisplayName_example") // AutoCreateApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFoldersAPI.FlowCreateFolder(context.Background()).AutoCreateApiKeyRequest(autoCreateApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFoldersAPI.FlowCreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateFolder`: FlowFolder
	fmt.Fprintf(os.Stdout, "Response from `FlowFoldersAPI.FlowCreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateApiKeyRequest** | [**AutoCreateApiKeyRequest**](AutoCreateApiKeyRequest.md) |  | 

### Return type

[**FlowFolder**](FlowFolder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowDeleteFolder

> FlowDeleteFolder(ctx, id).Execute()

Delete a folder

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
	r, err := apiClient.FlowFoldersAPI.FlowDeleteFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFoldersAPI.FlowDeleteFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowDeleteFolderRequest struct via the builder pattern


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


## FlowGetFolder

> map[string]interface{} FlowGetFolder(ctx, id).Execute()

Get a folder by id

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
	resp, r, err := apiClient.FlowFoldersAPI.FlowGetFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFoldersAPI.FlowGetFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetFolder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowFoldersAPI.FlowGetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetFolderRequest struct via the builder pattern


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


## FlowListFolders

> []FlowFolder FlowListFolders(ctx).Execute()

List folders

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
	resp, r, err := apiClient.FlowFoldersAPI.FlowListFolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFoldersAPI.FlowListFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListFolders`: []FlowFolder
	fmt.Fprintf(os.Stdout, "Response from `FlowFoldersAPI.FlowListFolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListFoldersRequest struct via the builder pattern


### Return type

[**[]FlowFolder**](FlowFolder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowUpdateFolder

> map[string]interface{} FlowUpdateFolder(ctx, id).AutoUpdateAppConnectionRequest(autoUpdateAppConnectionRequest).Execute()

Update a folder

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
	autoUpdateAppConnectionRequest := *openapiclient.NewAutoUpdateAppConnectionRequest() // AutoUpdateAppConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFoldersAPI.FlowUpdateFolder(context.Background(), id).AutoUpdateAppConnectionRequest(autoUpdateAppConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFoldersAPI.FlowUpdateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpdateFolder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowFoldersAPI.FlowUpdateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpdateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoUpdateAppConnectionRequest** | [**AutoUpdateAppConnectionRequest**](AutoUpdateAppConnectionRequest.md) |  | 

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

