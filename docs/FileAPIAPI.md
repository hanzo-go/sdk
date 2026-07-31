# \FileAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerActivateFile**](FileAPIAPI.md#CloudApiControllerActivateFile) | **Post** /v1/cloud/activate-file | Api Controller Activate File
[**CloudApiControllerAddFile**](FileAPIAPI.md#CloudApiControllerAddFile) | **Post** /v1/cloud/add-file | Api Controller Add File
[**CloudApiControllerDeleteFile**](FileAPIAPI.md#CloudApiControllerDeleteFile) | **Post** /v1/cloud/delete-file | Api Controller Delete File
[**CloudApiControllerGetActiveFile**](FileAPIAPI.md#CloudApiControllerGetActiveFile) | **Get** /v1/cloud/get-active-file | Api Controller Get Active File
[**CloudApiControllerUpdateFile**](FileAPIAPI.md#CloudApiControllerUpdateFile) | **Post** /v1/cloud/update-file | Api Controller Update File
[**CloudApiControllerUploadFile**](FileAPIAPI.md#CloudApiControllerUploadFile) | **Post** /v1/cloud/upload-file | Api Controller Upload File



## CloudApiControllerActivateFile

> CloudControllersResponse CloudApiControllerActivateFile(ctx).Key(key).Filename(filename).Execute()

Api Controller Activate File



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
	resp, r, err := apiClient.FileAPIAPI.CloudApiControllerActivateFile(context.Background()).Key(key).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPIAPI.CloudApiControllerActivateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerActivateFile`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPIAPI.CloudApiControllerActivateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerActivateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | The key of the file | 
 **filename** | **string** | The name of the file | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerAddFile

> CloudControllersResponse CloudApiControllerAddFile(ctx).Store(store).Key(key).IsLeaf(isLeaf).Filename(filename).CloudObjectFileInput(cloudObjectFileInput).Execute()

Api Controller Add File



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
	isLeaf := "isLeaf_example" // string | if is leaf
	filename := "filename_example" // string | The name of the file
	cloudObjectFileInput := *openapiclient.NewCloudObjectFileInput("Owner_example", "Name_example") // CloudObjectFileInput | The file object to register. owner and name are required and identify the file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPIAPI.CloudApiControllerAddFile(context.Background()).Store(store).Key(key).IsLeaf(isLeaf).Filename(filename).CloudObjectFileInput(cloudObjectFileInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPIAPI.CloudApiControllerAddFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddFile`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPIAPI.CloudApiControllerAddFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **store** | **string** | The store of the file | 
 **key** | **string** | The key of the file | 
 **isLeaf** | **string** | if is leaf | 
 **filename** | **string** | The name of the file | 
 **cloudObjectFileInput** | [**CloudObjectFileInput**](CloudObjectFileInput.md) | The file object to register. owner and name are required and identify the file. | 

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


## CloudApiControllerDeleteFile

> CloudControllersResponse CloudApiControllerDeleteFile(ctx).Store(store).Key(key).IsLeaf(isLeaf).CloudObjectFileInput(cloudObjectFileInput).Execute()

Api Controller Delete File



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
	isLeaf := "isLeaf_example" // string | if is leaf
	cloudObjectFileInput := *openapiclient.NewCloudObjectFileInput("Owner_example", "Name_example") // CloudObjectFileInput | The file to delete. owner and name are required and identify the file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPIAPI.CloudApiControllerDeleteFile(context.Background()).Store(store).Key(key).IsLeaf(isLeaf).CloudObjectFileInput(cloudObjectFileInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPIAPI.CloudApiControllerDeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteFile`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPIAPI.CloudApiControllerDeleteFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **store** | **string** | The store of the file | 
 **key** | **string** | The key of the file | 
 **isLeaf** | **string** | if is leaf | 
 **cloudObjectFileInput** | [**CloudObjectFileInput**](CloudObjectFileInput.md) | The file to delete. owner and name are required and identify the file. | 

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


## CloudApiControllerGetActiveFile

> map[string]interface{} CloudApiControllerGetActiveFile(ctx).Prefix(prefix).Execute()

Api Controller Get Active File



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
	resp, r, err := apiClient.FileAPIAPI.CloudApiControllerGetActiveFile(context.Background()).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPIAPI.CloudApiControllerGetActiveFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetActiveFile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FileAPIAPI.CloudApiControllerGetActiveFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetActiveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | The prefix of the file | 

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


## CloudApiControllerUpdateFile

> CloudControllersResponse CloudApiControllerUpdateFile(ctx).StoreId(storeId).Key(key).CloudObjectFile(cloudObjectFile).Execute()

Api Controller Update File



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
	cloudObjectFile := *openapiclient.NewCloudObjectFile() // CloudObjectFile | The details of the File

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPIAPI.CloudApiControllerUpdateFile(context.Background()).StoreId(storeId).Key(key).CloudObjectFile(cloudObjectFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPIAPI.CloudApiControllerUpdateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateFile`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPIAPI.CloudApiControllerUpdateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | The store id of the file | 
 **key** | **string** | The key of the file | 
 **cloudObjectFile** | [**CloudObjectFile**](CloudObjectFile.md) | The details of the File | 

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


## CloudApiControllerUploadFile

> CloudControllersResponse CloudApiControllerUploadFile(ctx).File(file).Type_(type_).Name(name).Execute()

Api Controller Upload File



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
	resp, r, err := apiClient.FileAPIAPI.CloudApiControllerUploadFile(context.Background()).File(file).Type_(type_).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPIAPI.CloudApiControllerUploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUploadFile`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPIAPI.CloudApiControllerUploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string** | The base64 encoded file data | 
 **type_** | **string** | The file type/extension | 
 **name** | **string** | The file name | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

