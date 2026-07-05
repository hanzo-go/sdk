# \KmsSecretFoldersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateSecretFolder**](KmsSecretFoldersAPI.md#KmsCreateSecretFolder) | **Post** /v1/kms/folders | Create a secret folder
[**KmsDeleteSecretFolder**](KmsSecretFoldersAPI.md#KmsDeleteSecretFolder) | **Delete** /v1/kms/folders/{folderId} | Delete a secret folder
[**KmsListSecretFolders**](KmsSecretFoldersAPI.md#KmsListSecretFolders) | **Get** /v1/kms/folders | List secret folders
[**KmsUpdateSecretFolder**](KmsSecretFoldersAPI.md#KmsUpdateSecretFolder) | **Patch** /v1/kms/folders/{folderId} | Update a secret folder



## KmsCreateSecretFolder

> KmsCreateSecretFolder200Response KmsCreateSecretFolder(ctx).KmsCreateSecretFolderRequest(kmsCreateSecretFolderRequest).Execute()

Create a secret folder

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
	kmsCreateSecretFolderRequest := *openapiclient.NewKmsCreateSecretFolderRequest("WorkspaceId_example", "Environment_example", "Name_example") // KmsCreateSecretFolderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretFoldersAPI.KmsCreateSecretFolder(context.Background()).KmsCreateSecretFolderRequest(kmsCreateSecretFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretFoldersAPI.KmsCreateSecretFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateSecretFolder`: KmsCreateSecretFolder200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretFoldersAPI.KmsCreateSecretFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateSecretFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateSecretFolderRequest** | [**KmsCreateSecretFolderRequest**](KmsCreateSecretFolderRequest.md) |  | 

### Return type

[**KmsCreateSecretFolder200Response**](KmsCreateSecretFolder200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteSecretFolder

> map[string]interface{} KmsDeleteSecretFolder(ctx, folderId).KmsDeleteSecretFolderRequest(kmsDeleteSecretFolderRequest).Execute()

Delete a secret folder

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
	folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsDeleteSecretFolderRequest := *openapiclient.NewKmsDeleteSecretFolderRequest("WorkspaceId_example", "Environment_example") // KmsDeleteSecretFolderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretFoldersAPI.KmsDeleteSecretFolder(context.Background(), folderId).KmsDeleteSecretFolderRequest(kmsDeleteSecretFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretFoldersAPI.KmsDeleteSecretFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteSecretFolder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretFoldersAPI.KmsDeleteSecretFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteSecretFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsDeleteSecretFolderRequest** | [**KmsDeleteSecretFolderRequest**](KmsDeleteSecretFolderRequest.md) |  | 

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


## KmsListSecretFolders

> KmsListSecretFolders200Response KmsListSecretFolders(ctx).WorkspaceId(workspaceId).Environment(environment).Directory(directory).Execute()

List secret folders

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	environment := "environment_example" // string | 
	directory := "directory_example" // string |  (optional) (default to "/")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretFoldersAPI.KmsListSecretFolders(context.Background()).WorkspaceId(workspaceId).Environment(environment).Directory(directory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretFoldersAPI.KmsListSecretFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListSecretFolders`: KmsListSecretFolders200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretFoldersAPI.KmsListSecretFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListSecretFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 
 **environment** | **string** |  | 
 **directory** | **string** |  | [default to &quot;/&quot;]

### Return type

[**KmsListSecretFolders200Response**](KmsListSecretFolders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateSecretFolder

> KmsCreateSecretFolder200Response KmsUpdateSecretFolder(ctx, folderId).KmsUpdateSecretFolderRequest(kmsUpdateSecretFolderRequest).Execute()

Update a secret folder

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
	folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateSecretFolderRequest := *openapiclient.NewKmsUpdateSecretFolderRequest() // KmsUpdateSecretFolderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretFoldersAPI.KmsUpdateSecretFolder(context.Background(), folderId).KmsUpdateSecretFolderRequest(kmsUpdateSecretFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretFoldersAPI.KmsUpdateSecretFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateSecretFolder`: KmsCreateSecretFolder200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretFoldersAPI.KmsUpdateSecretFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateSecretFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateSecretFolderRequest** | [**KmsUpdateSecretFolderRequest**](KmsUpdateSecretFolderRequest.md) |  | 

### Return type

[**KmsCreateSecretFolder200Response**](KmsCreateSecretFolder200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

