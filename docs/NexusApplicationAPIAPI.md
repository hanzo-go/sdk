# \NexusApplicationAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddApplication**](NexusApplicationAPIAPI.md#NexusAddApplication) | **Post** /v1/nexus/add-application | add Application
[**NexusDeleteApplication**](NexusApplicationAPIAPI.md#NexusDeleteApplication) | **Post** /v1/nexus/delete-application | delete Application
[**NexusDeployApplication**](NexusApplicationAPIAPI.md#NexusDeployApplication) | **Post** /v1/nexus/deploy-application | deploy Application
[**NexusGetApplication**](NexusApplicationAPIAPI.md#NexusGetApplication) | **Get** /v1/nexus/get-application | get Application
[**NexusGetApplicationStatus**](NexusApplicationAPIAPI.md#NexusGetApplicationStatus) | **Get** /v1/nexus/get-application-status | get Application Status
[**NexusGetApplications**](NexusApplicationAPIAPI.md#NexusGetApplications) | **Get** /v1/nexus/get-applications | get Applications
[**NexusUndeployApplication**](NexusApplicationAPIAPI.md#NexusUndeployApplication) | **Post** /v1/nexus/undeploy-application | undeploy Application
[**NexusUpdateApplication**](NexusApplicationAPIAPI.md#NexusUpdateApplication) | **Post** /v1/nexus/update-application | update Application



## NexusAddApplication

> NexusResponse NexusAddApplication(ctx).NexusApplication(nexusApplication).Execute()

add Application



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
	nexusApplication := *openapiclient.NewNexusApplication() // NexusApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusApplicationAPIAPI.NexusAddApplication(context.Background()).NexusApplication(nexusApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusApplicationAPIAPI.NexusAddApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddApplication`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusApplicationAPIAPI.NexusAddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusApplication** | [**NexusApplication**](NexusApplication.md) | The details of the application | 

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


## NexusDeleteApplication

> NexusResponse NexusDeleteApplication(ctx).NexusApplication(nexusApplication).Execute()

delete Application



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
	nexusApplication := *openapiclient.NewNexusApplication() // NexusApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusApplicationAPIAPI.NexusDeleteApplication(context.Background()).NexusApplication(nexusApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusApplicationAPIAPI.NexusDeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteApplication`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusApplicationAPIAPI.NexusDeleteApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusApplication** | [**NexusApplication**](NexusApplication.md) | The details of the application | 

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


## NexusDeployApplication

> NexusResponse NexusDeployApplication(ctx).Body(body).Execute()

deploy Application



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
	body := map[string]interface{}{ ... } // map[string]interface{} | The deployment request details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusApplicationAPIAPI.NexusDeployApplication(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusApplicationAPIAPI.NexusDeployApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeployApplication`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusApplicationAPIAPI.NexusDeployApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeployApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The deployment request details | 

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


## NexusGetApplication

> NexusApplication NexusGetApplication(ctx).Id(id).Execute()

get Application



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
	id := "id_example" // string | The id of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusApplicationAPIAPI.NexusGetApplication(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusApplicationAPIAPI.NexusGetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetApplication`: NexusApplication
	fmt.Fprintf(os.Stdout, "Response from `NexusApplicationAPIAPI.NexusGetApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the application | 

### Return type

[**NexusApplication**](NexusApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetApplicationStatus

> map[string]interface{} NexusGetApplicationStatus(ctx).Id(id).Execute()

get Application Status



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
	id := "id_example" // string | The id (owner/name) of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusApplicationAPIAPI.NexusGetApplicationStatus(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusApplicationAPIAPI.NexusGetApplicationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetApplicationStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NexusApplicationAPIAPI.NexusGetApplicationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetApplicationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the application | 

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


## NexusGetApplications

> []NexusApplication NexusGetApplications(ctx).Owner(owner).Execute()

get Applications



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
	owner := "owner_example" // string | The owner of the applications

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusApplicationAPIAPI.NexusGetApplications(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusApplicationAPIAPI.NexusGetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetApplications`: []NexusApplication
	fmt.Fprintf(os.Stdout, "Response from `NexusApplicationAPIAPI.NexusGetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the applications | 

### Return type

[**[]NexusApplication**](NexusApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUndeployApplication

> NexusResponse NexusUndeployApplication(ctx).Body(body).Execute()

undeploy Application



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
	body := map[string]interface{}{ ... } // map[string]interface{} | The deployment request details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusApplicationAPIAPI.NexusUndeployApplication(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusApplicationAPIAPI.NexusUndeployApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUndeployApplication`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusApplicationAPIAPI.NexusUndeployApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUndeployApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The deployment request details | 

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


## NexusUpdateApplication

> NexusResponse NexusUpdateApplication(ctx).Id(id).NexusApplication(nexusApplication).Execute()

update Application



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
	id := "id_example" // string | The id (owner/name) of the application
	nexusApplication := *openapiclient.NewNexusApplication() // NexusApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusApplicationAPIAPI.NexusUpdateApplication(context.Background()).Id(id).NexusApplication(nexusApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusApplicationAPIAPI.NexusUpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateApplication`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusApplicationAPIAPI.NexusUpdateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the application | 
 **nexusApplication** | [**NexusApplication**](NexusApplication.md) | The details of the application | 

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

