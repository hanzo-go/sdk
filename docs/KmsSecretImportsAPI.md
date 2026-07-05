# \KmsSecretImportsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateSecretImport**](KmsSecretImportsAPI.md#KmsCreateSecretImport) | **Post** /v1/kms/secret-imports | Create a secret import
[**KmsDeleteSecretImport**](KmsSecretImportsAPI.md#KmsDeleteSecretImport) | **Delete** /v1/kms/secret-imports/{importId} | Delete a secret import
[**KmsListSecretImports**](KmsSecretImportsAPI.md#KmsListSecretImports) | **Get** /v1/kms/secret-imports | List secret imports
[**KmsUpdateSecretImport**](KmsSecretImportsAPI.md#KmsUpdateSecretImport) | **Patch** /v1/kms/secret-imports/{importId} | Update a secret import



## KmsCreateSecretImport

> KmsCreateSecretImport200Response KmsCreateSecretImport(ctx).KmsCreateSecretImportRequest(kmsCreateSecretImportRequest).Execute()

Create a secret import

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
	kmsCreateSecretImportRequest := *openapiclient.NewKmsCreateSecretImportRequest("WorkspaceId_example", "Environment_example", "Directory_example", *openapiclient.NewKmsCreateSecretImportRequestImport("Environment_example", "Path_example")) // KmsCreateSecretImportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretImportsAPI.KmsCreateSecretImport(context.Background()).KmsCreateSecretImportRequest(kmsCreateSecretImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretImportsAPI.KmsCreateSecretImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateSecretImport`: KmsCreateSecretImport200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretImportsAPI.KmsCreateSecretImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateSecretImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateSecretImportRequest** | [**KmsCreateSecretImportRequest**](KmsCreateSecretImportRequest.md) |  | 

### Return type

[**KmsCreateSecretImport200Response**](KmsCreateSecretImport200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteSecretImport

> map[string]interface{} KmsDeleteSecretImport(ctx, importId).Execute()

Delete a secret import

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
	importId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretImportsAPI.KmsDeleteSecretImport(context.Background(), importId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretImportsAPI.KmsDeleteSecretImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteSecretImport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretImportsAPI.KmsDeleteSecretImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteSecretImportRequest struct via the builder pattern


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


## KmsListSecretImports

> KmsListSecretImports200Response KmsListSecretImports(ctx).WorkspaceId(workspaceId).Environment(environment).Directory(directory).Execute()

List secret imports

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
	resp, r, err := apiClient.KmsSecretImportsAPI.KmsListSecretImports(context.Background()).WorkspaceId(workspaceId).Environment(environment).Directory(directory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretImportsAPI.KmsListSecretImports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListSecretImports`: KmsListSecretImports200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretImportsAPI.KmsListSecretImports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListSecretImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 
 **environment** | **string** |  | 
 **directory** | **string** |  | [default to &quot;/&quot;]

### Return type

[**KmsListSecretImports200Response**](KmsListSecretImports200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateSecretImport

> KmsCreateSecretImport200Response KmsUpdateSecretImport(ctx, importId).KmsUpdateSecretImportRequest(kmsUpdateSecretImportRequest).Execute()

Update a secret import

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
	importId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateSecretImportRequest := *openapiclient.NewKmsUpdateSecretImportRequest() // KmsUpdateSecretImportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretImportsAPI.KmsUpdateSecretImport(context.Background(), importId).KmsUpdateSecretImportRequest(kmsUpdateSecretImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretImportsAPI.KmsUpdateSecretImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateSecretImport`: KmsCreateSecretImport200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretImportsAPI.KmsUpdateSecretImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateSecretImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateSecretImportRequest** | [**KmsUpdateSecretImportRequest**](KmsUpdateSecretImportRequest.md) |  | 

### Return type

[**KmsCreateSecretImport200Response**](KmsCreateSecretImport200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

