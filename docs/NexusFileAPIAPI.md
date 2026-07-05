# \NexusFileAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusActivateFile**](NexusFileAPIAPI.md#NexusActivateFile) | **Post** /v1/nexus/activate-file | activate File
[**NexusAddFile**](NexusFileAPIAPI.md#NexusAddFile) | **Post** /v1/nexus/add-file | add File
[**NexusDeleteFile**](NexusFileAPIAPI.md#NexusDeleteFile) | **Post** /v1/nexus/delete-file | delete File
[**NexusGetActiveFile**](NexusFileAPIAPI.md#NexusGetActiveFile) | **Get** /v1/nexus/get-active-file | get Active File
[**NexusUpdateFile**](NexusFileAPIAPI.md#NexusUpdateFile) | **Post** /v1/nexus/update-file | update File
[**NexusUploadFile**](NexusFileAPIAPI.md#NexusUploadFile) | **Post** /v1/nexus/upload-file | upload File



## NexusActivateFile

> NexusResponse NexusActivateFile(ctx).Key(key).Filename(filename).Execute()

activate File



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
	key := "key_example" // string | The key of the file
	filename := "filename_example" // string | The name of the file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFileAPIAPI.NexusActivateFile(context.Background()).Key(key).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFileAPIAPI.NexusActivateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusActivateFile`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusFileAPIAPI.NexusActivateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusActivateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | The key of the file | 
 **filename** | **string** | The name of the file | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusAddFile

> NexusResponse NexusAddFile(ctx).Store(store).Key(key).IsLeaf(isLeaf).Filename(filename).Execute()

add File



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
	store := "store_example" // string | The store of the file
	key := "key_example" // string | The key of the file
	isLeaf := "isLeaf_example" // string | Whether the file is a leaf node
	filename := "filename_example" // string | The name of the file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFileAPIAPI.NexusAddFile(context.Background()).Store(store).Key(key).IsLeaf(isLeaf).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFileAPIAPI.NexusAddFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddFile`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusFileAPIAPI.NexusAddFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **store** | **string** | The store of the file | 
 **key** | **string** | The key of the file | 
 **isLeaf** | **string** | Whether the file is a leaf node | 
 **filename** | **string** | The name of the file | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteFile

> NexusResponse NexusDeleteFile(ctx).Store(store).Key(key).IsLeaf(isLeaf).Execute()

delete File



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
	store := "store_example" // string | The store of the file
	key := "key_example" // string | The key of the file
	isLeaf := "isLeaf_example" // string | Whether the file is a leaf node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFileAPIAPI.NexusDeleteFile(context.Background()).Store(store).Key(key).IsLeaf(isLeaf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFileAPIAPI.NexusDeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteFile`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusFileAPIAPI.NexusDeleteFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **store** | **string** | The store of the file | 
 **key** | **string** | The key of the file | 
 **isLeaf** | **string** | Whether the file is a leaf node | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetActiveFile

> string NexusGetActiveFile(ctx).Prefix(prefix).Execute()

get Active File



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
	prefix := "prefix_example" // string | The prefix of the file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFileAPIAPI.NexusGetActiveFile(context.Background()).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFileAPIAPI.NexusGetActiveFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetActiveFile`: string
	fmt.Fprintf(os.Stdout, "Response from `NexusFileAPIAPI.NexusGetActiveFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetActiveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | The prefix of the file | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateFile

> NexusResponse NexusUpdateFile(ctx).StoreId(storeId).Key(key).NexusFile(nexusFile).Execute()

update File



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
	storeId := "storeId_example" // string | The store id of the file
	key := "key_example" // string | The key of the file
	nexusFile := *openapiclient.NewNexusFile() // NexusFile | The details of the file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFileAPIAPI.NexusUpdateFile(context.Background()).StoreId(storeId).Key(key).NexusFile(nexusFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFileAPIAPI.NexusUpdateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateFile`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusFileAPIAPI.NexusUpdateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | The store id of the file | 
 **key** | **string** | The key of the file | 
 **nexusFile** | [**NexusFile**](NexusFile.md) | The details of the file | 

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


## NexusUploadFile

> NexusResponse NexusUploadFile(ctx).File(file).Type_(type_).Name(name).Execute()

upload File



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
	file := "file_example" // string | The base64 encoded file data
	type_ := "type__example" // string | The file type/extension
	name := "name_example" // string | The file name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusFileAPIAPI.NexusUploadFile(context.Background()).File(file).Type_(type_).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusFileAPIAPI.NexusUploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUploadFile`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusFileAPIAPI.NexusUploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string** | The base64 encoded file data | 
 **type_** | **string** | The file type/extension | 
 **name** | **string** | The file name | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

