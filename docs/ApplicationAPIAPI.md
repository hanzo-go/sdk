# \ApplicationAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddApplication**](ApplicationAPIAPI.md#CloudApiControllerAddApplication) | **Post** /v1/cloud/add-application | Api Controller Add Application
[**CloudApiControllerDeleteApplication**](ApplicationAPIAPI.md#CloudApiControllerDeleteApplication) | **Post** /v1/cloud/delete-application | Api Controller Delete Application
[**CloudApiControllerDeployApplication**](ApplicationAPIAPI.md#CloudApiControllerDeployApplication) | **Post** /v1/cloud/deploy-application | Api Controller Deploy Application
[**CloudApiControllerGetApplication**](ApplicationAPIAPI.md#CloudApiControllerGetApplication) | **Get** /v1/cloud/get-application | Api Controller Get Application
[**CloudApiControllerGetApplicationStatus**](ApplicationAPIAPI.md#CloudApiControllerGetApplicationStatus) | **Get** /v1/cloud/get-application-status | Api Controller Get Application Status
[**CloudApiControllerGetApplications**](ApplicationAPIAPI.md#CloudApiControllerGetApplications) | **Get** /v1/cloud/get-applications | Api Controller Get Applications
[**CloudApiControllerUndeployApplication**](ApplicationAPIAPI.md#CloudApiControllerUndeployApplication) | **Post** /v1/cloud/undeploy-application | Api Controller Undeploy Application
[**CloudApiControllerUpdateApplication**](ApplicationAPIAPI.md#CloudApiControllerUpdateApplication) | **Post** /v1/cloud/update-application | Api Controller Update Application



## CloudApiControllerAddApplication

> CloudControllersResponse CloudApiControllerAddApplication(ctx).CloudObjectApplication(cloudObjectApplication).Execute()

Api Controller Add Application



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
	cloudObjectApplication := *openapiclient.NewCloudObjectApplication() // CloudObjectApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPIAPI.CloudApiControllerAddApplication(context.Background()).CloudObjectApplication(cloudObjectApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPIAPI.CloudApiControllerAddApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddApplication`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPIAPI.CloudApiControllerAddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectApplication** | [**CloudObjectApplication**](CloudObjectApplication.md) | The details of the application | 

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


## CloudApiControllerDeleteApplication

> CloudControllersResponse CloudApiControllerDeleteApplication(ctx).CloudObjectApplication(cloudObjectApplication).Execute()

Api Controller Delete Application



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
	cloudObjectApplication := *openapiclient.NewCloudObjectApplication() // CloudObjectApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPIAPI.CloudApiControllerDeleteApplication(context.Background()).CloudObjectApplication(cloudObjectApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPIAPI.CloudApiControllerDeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteApplication`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPIAPI.CloudApiControllerDeleteApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectApplication** | [**CloudObjectApplication**](CloudObjectApplication.md) | The details of the application | 

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


## CloudApiControllerDeployApplication

> CloudControllersResponse CloudApiControllerDeployApplication(ctx).Body(body).Execute()

Api Controller Deploy Application



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
	resp, r, err := apiClient.ApplicationAPIAPI.CloudApiControllerDeployApplication(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPIAPI.CloudApiControllerDeployApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeployApplication`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPIAPI.CloudApiControllerDeployApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeployApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The deployment request details | 

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


## CloudApiControllerGetApplication

> CloudObjectApplication CloudApiControllerGetApplication(ctx).Id(id).Execute()

Api Controller Get Application



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
	id := "id_example" // string | The id of application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPIAPI.CloudApiControllerGetApplication(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPIAPI.CloudApiControllerGetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetApplication`: CloudObjectApplication
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPIAPI.CloudApiControllerGetApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of application | 

### Return type

[**CloudObjectApplication**](CloudObjectApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetApplicationStatus

> map[string]interface{} CloudApiControllerGetApplicationStatus(ctx).Id(id).Execute()

Api Controller Get Application Status



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
	resp, r, err := apiClient.ApplicationAPIAPI.CloudApiControllerGetApplicationStatus(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPIAPI.CloudApiControllerGetApplicationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetApplicationStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPIAPI.CloudApiControllerGetApplicationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetApplicationStatusRequest struct via the builder pattern


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


## CloudApiControllerGetApplications

> []CloudObjectApplication CloudApiControllerGetApplications(ctx).Owner(owner).Execute()

Api Controller Get Applications



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
	owner := "owner_example" // string | The owner of applications

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPIAPI.CloudApiControllerGetApplications(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPIAPI.CloudApiControllerGetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetApplications`: []CloudObjectApplication
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPIAPI.CloudApiControllerGetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of applications | 

### Return type

[**[]CloudObjectApplication**](CloudObjectApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUndeployApplication

> CloudControllersResponse CloudApiControllerUndeployApplication(ctx).Body(body).Execute()

Api Controller Undeploy Application



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
	resp, r, err := apiClient.ApplicationAPIAPI.CloudApiControllerUndeployApplication(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPIAPI.CloudApiControllerUndeployApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUndeployApplication`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPIAPI.CloudApiControllerUndeployApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUndeployApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The deployment request details | 

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


## CloudApiControllerUpdateApplication

> CloudControllersResponse CloudApiControllerUpdateApplication(ctx).Id(id).CloudObjectApplication(cloudObjectApplication).Execute()

Api Controller Update Application



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
	cloudObjectApplication := *openapiclient.NewCloudObjectApplication() // CloudObjectApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPIAPI.CloudApiControllerUpdateApplication(context.Background()).Id(id).CloudObjectApplication(cloudObjectApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPIAPI.CloudApiControllerUpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateApplication`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPIAPI.CloudApiControllerUpdateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the application | 
 **cloudObjectApplication** | [**CloudObjectApplication**](CloudObjectApplication.md) | The details of the application | 

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

