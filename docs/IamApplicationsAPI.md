# \IamApplicationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddApplication**](IamApplicationsAPI.md#IamApiControllerAddApplication) | **Post** /v1/iam/applications | Api Controller Add Application
[**IamApiControllerDeleteApplication**](IamApplicationsAPI.md#IamApiControllerDeleteApplication) | **Delete** /v1/iam/applications/{id} | Api Controller Delete Application
[**IamApiControllerGetApplication**](IamApplicationsAPI.md#IamApiControllerGetApplication) | **Get** /v1/iam/applications/{id} | Api Controller Get Application
[**IamApiControllerGetApplications**](IamApplicationsAPI.md#IamApiControllerGetApplications) | **Get** /v1/iam/applications | Api Controller Get Applications
[**IamApiControllerGetOrganizationApplications**](IamApplicationsAPI.md#IamApiControllerGetOrganizationApplications) | **Get** /v1/iam/organizations/applications | Api Controller Get Organization Applications
[**IamApiControllerGetUserApplication**](IamApplicationsAPI.md#IamApiControllerGetUserApplication) | **Get** /v1/iam/users/application | Api Controller Get User Application
[**IamApiControllerUpdateApplication**](IamApplicationsAPI.md#IamApiControllerUpdateApplication) | **Put** /v1/iam/applications/{id} | Api Controller Update Application



## IamApiControllerAddApplication

> IamControllersResponse IamApiControllerAddApplication(ctx).IamObjectApplication(iamObjectApplication).Execute()

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
	iamObjectApplication := *openapiclient.NewIamObjectApplication() // IamObjectApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamApplicationsAPI.IamApiControllerAddApplication(context.Background()).IamObjectApplication(iamObjectApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamApplicationsAPI.IamApiControllerAddApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddApplication`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamApplicationsAPI.IamApiControllerAddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectApplication** | [**IamObjectApplication**](IamObjectApplication.md) | The details of the application | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteApplication

> IamControllersResponse IamApiControllerDeleteApplication(ctx, id).IamObjectApplication(iamObjectApplication).Execute()

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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectApplication := *openapiclient.NewIamObjectApplication() // IamObjectApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamApplicationsAPI.IamApiControllerDeleteApplication(context.Background(), id).IamObjectApplication(iamObjectApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamApplicationsAPI.IamApiControllerDeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteApplication`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamApplicationsAPI.IamApiControllerDeleteApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectApplication** | [**IamObjectApplication**](IamObjectApplication.md) | The details of the application | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetApplication

> IamObjectApplication IamApiControllerGetApplication(ctx, id).Execute()

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
	id := "id_example" // string | The id ( owner/name ) of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamApplicationsAPI.IamApiControllerGetApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamApplicationsAPI.IamApiControllerGetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetApplication`: IamObjectApplication
	fmt.Fprintf(os.Stdout, "Response from `IamApplicationsAPI.IamApiControllerGetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectApplication**](IamObjectApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetApplications

> []IamObjectApplication IamApiControllerGetApplications(ctx).Owner(owner).Execute()

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
	owner := "owner_example" // string | The owner of applications.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamApplicationsAPI.IamApiControllerGetApplications(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamApplicationsAPI.IamApiControllerGetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetApplications`: []IamObjectApplication
	fmt.Fprintf(os.Stdout, "Response from `IamApplicationsAPI.IamApiControllerGetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of applications. | 

### Return type

[**[]IamObjectApplication**](IamObjectApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetOrganizationApplications

> []IamObjectApplication IamApiControllerGetOrganizationApplications(ctx).Organization(organization).Execute()

Api Controller Get Organization Applications



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
	organization := "organization_example" // string | The organization name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamApplicationsAPI.IamApiControllerGetOrganizationApplications(context.Background()).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamApplicationsAPI.IamApiControllerGetOrganizationApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetOrganizationApplications`: []IamObjectApplication
	fmt.Fprintf(os.Stdout, "Response from `IamApplicationsAPI.IamApiControllerGetOrganizationApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetOrganizationApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | **string** | The organization name | 

### Return type

[**[]IamObjectApplication**](IamObjectApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetUserApplication

> IamObjectApplication IamApiControllerGetUserApplication(ctx).Id(id).Execute()

Api Controller Get User Application



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
	id := "id_example" // string | The id ( owner/name ) of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamApplicationsAPI.IamApiControllerGetUserApplication(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamApplicationsAPI.IamApiControllerGetUserApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUserApplication`: IamObjectApplication
	fmt.Fprintf(os.Stdout, "Response from `IamApplicationsAPI.IamApiControllerGetUserApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetUserApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the user | 

### Return type

[**IamObjectApplication**](IamObjectApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateApplication

> IamControllersResponse IamApiControllerUpdateApplication(ctx, id).IamObjectApplication(iamObjectApplication).Execute()

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
	id := "id_example" // string | The id ( owner/name ) of the application
	iamObjectApplication := *openapiclient.NewIamObjectApplication() // IamObjectApplication | The details of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamApplicationsAPI.IamApiControllerUpdateApplication(context.Background(), id).IamObjectApplication(iamObjectApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamApplicationsAPI.IamApiControllerUpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateApplication`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamApplicationsAPI.IamApiControllerUpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectApplication** | [**IamObjectApplication**](IamObjectApplication.md) | The details of the application | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

