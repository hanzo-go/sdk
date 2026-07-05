# \ConsoleProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreateProject**](ConsoleProjectsAPI.md#ConsoleCreateProject) | **Post** /v1/console/projects | Create a new project (requires organization-scoped API key)
[**ConsoleCreateProjectApiKey**](ConsoleProjectsAPI.md#ConsoleCreateProjectApiKey) | **Post** /v1/console/projects/{projectId}/apiKeys | Create an API key for a project
[**ConsoleDeleteProject**](ConsoleProjectsAPI.md#ConsoleDeleteProject) | **Delete** /v1/console/projects/{projectId} | Delete a project (async)
[**ConsoleDeleteProjectApiKey**](ConsoleProjectsAPI.md#ConsoleDeleteProjectApiKey) | **Delete** /v1/console/projects/{projectId}/apiKeys/{apiKeyId} | Delete a project API key
[**ConsoleGetProject**](ConsoleProjectsAPI.md#ConsoleGetProject) | **Get** /v1/console/projects | Get project associated with API key
[**ConsoleGetProjectById**](ConsoleProjectsAPI.md#ConsoleGetProjectById) | **Get** /v1/console/projects/{projectId} | Get a project by ID
[**ConsoleListProjectApiKeys**](ConsoleProjectsAPI.md#ConsoleListProjectApiKeys) | **Get** /v1/console/projects/{projectId}/apiKeys | Get all API keys for a project
[**ConsoleUpdateProject**](ConsoleProjectsAPI.md#ConsoleUpdateProject) | **Put** /v1/console/projects/{projectId} | Update a project



## ConsoleCreateProject

> ConsoleProject ConsoleCreateProject(ctx).ConsoleCreateProjectRequest(consoleCreateProjectRequest).Execute()

Create a new project (requires organization-scoped API key)

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
	consoleCreateProjectRequest := *openapiclient.NewConsoleCreateProjectRequest("Name_example") // ConsoleCreateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleProjectsAPI.ConsoleCreateProject(context.Background()).ConsoleCreateProjectRequest(consoleCreateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleProjectsAPI.ConsoleCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateProject`: ConsoleProject
	fmt.Fprintf(os.Stdout, "Response from `ConsoleProjectsAPI.ConsoleCreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreateProjectRequest** | [**ConsoleCreateProjectRequest**](ConsoleCreateProjectRequest.md) |  | 

### Return type

[**ConsoleProject**](ConsoleProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleCreateProjectApiKey

> ConsoleCreateProjectApiKey200Response ConsoleCreateProjectApiKey(ctx, projectId).ConsoleCreateProjectApiKeyRequest(consoleCreateProjectApiKeyRequest).Execute()

Create an API key for a project

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
	projectId := "projectId_example" // string | 
	consoleCreateProjectApiKeyRequest := *openapiclient.NewConsoleCreateProjectApiKeyRequest() // ConsoleCreateProjectApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleProjectsAPI.ConsoleCreateProjectApiKey(context.Background(), projectId).ConsoleCreateProjectApiKeyRequest(consoleCreateProjectApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleProjectsAPI.ConsoleCreateProjectApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateProjectApiKey`: ConsoleCreateProjectApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleProjectsAPI.ConsoleCreateProjectApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateProjectApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consoleCreateProjectApiKeyRequest** | [**ConsoleCreateProjectApiKeyRequest**](ConsoleCreateProjectApiKeyRequest.md) |  | 

### Return type

[**ConsoleCreateProjectApiKey200Response**](ConsoleCreateProjectApiKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleDeleteProject

> ConsoleDeleteProject202Response ConsoleDeleteProject(ctx, projectId).Execute()

Delete a project (async)

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
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleProjectsAPI.ConsoleDeleteProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleProjectsAPI.ConsoleDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteProject`: ConsoleDeleteProject202Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleProjectsAPI.ConsoleDeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleDeleteProject202Response**](ConsoleDeleteProject202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleDeleteProjectApiKey

> ConsoleDeleteProjectApiKey200Response ConsoleDeleteProjectApiKey(ctx, projectId, apiKeyId).Execute()

Delete a project API key

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
	projectId := "projectId_example" // string | 
	apiKeyId := "apiKeyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleProjectsAPI.ConsoleDeleteProjectApiKey(context.Background(), projectId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleProjectsAPI.ConsoleDeleteProjectApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteProjectApiKey`: ConsoleDeleteProjectApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleProjectsAPI.ConsoleDeleteProjectApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteProjectApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsoleDeleteProjectApiKey200Response**](ConsoleDeleteProjectApiKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetProject

> ConsoleGetProject200Response ConsoleGetProject(ctx).Execute()

Get project associated with API key

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
	resp, r, err := apiClient.ConsoleProjectsAPI.ConsoleGetProject(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleProjectsAPI.ConsoleGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetProject`: ConsoleGetProject200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleProjectsAPI.ConsoleGetProject`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetProjectRequest struct via the builder pattern


### Return type

[**ConsoleGetProject200Response**](ConsoleGetProject200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetProjectById

> ConsoleProject ConsoleGetProjectById(ctx, projectId).Execute()

Get a project by ID

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
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleProjectsAPI.ConsoleGetProjectById(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleProjectsAPI.ConsoleGetProjectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetProjectById`: ConsoleProject
	fmt.Fprintf(os.Stdout, "Response from `ConsoleProjectsAPI.ConsoleGetProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetProjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleProject**](ConsoleProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListProjectApiKeys

> ConsoleListOrganizationApiKeys200Response ConsoleListProjectApiKeys(ctx, projectId).Execute()

Get all API keys for a project

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
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleProjectsAPI.ConsoleListProjectApiKeys(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleProjectsAPI.ConsoleListProjectApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListProjectApiKeys`: ConsoleListOrganizationApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleProjectsAPI.ConsoleListProjectApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListProjectApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleListOrganizationApiKeys200Response**](ConsoleListOrganizationApiKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleUpdateProject

> ConsoleProject ConsoleUpdateProject(ctx, projectId).ConsoleUpdateProjectRequest(consoleUpdateProjectRequest).Execute()

Update a project

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
	projectId := "projectId_example" // string | 
	consoleUpdateProjectRequest := *openapiclient.NewConsoleUpdateProjectRequest("Name_example") // ConsoleUpdateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleProjectsAPI.ConsoleUpdateProject(context.Background(), projectId).ConsoleUpdateProjectRequest(consoleUpdateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleProjectsAPI.ConsoleUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleUpdateProject`: ConsoleProject
	fmt.Fprintf(os.Stdout, "Response from `ConsoleProjectsAPI.ConsoleUpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consoleUpdateProjectRequest** | [**ConsoleUpdateProjectRequest**](ConsoleUpdateProjectRequest.md) |  | 

### Return type

[**ConsoleProject**](ConsoleProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

