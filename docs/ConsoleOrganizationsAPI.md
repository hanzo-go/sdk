# \ConsoleOrganizationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleDeleteOrganizationMembership**](ConsoleOrganizationsAPI.md#ConsoleDeleteOrganizationMembership) | **Delete** /v1/console/organizations/memberships | Delete an organization membership
[**ConsoleDeleteProjectMembership**](ConsoleOrganizationsAPI.md#ConsoleDeleteProjectMembership) | **Delete** /v1/console/projects/{projectId}/memberships | Delete a project membership
[**ConsoleListOrganizationApiKeys**](ConsoleOrganizationsAPI.md#ConsoleListOrganizationApiKeys) | **Get** /v1/console/organizations/apiKeys | Get all API keys for the organization
[**ConsoleListOrganizationMemberships**](ConsoleOrganizationsAPI.md#ConsoleListOrganizationMemberships) | **Get** /v1/console/organizations/memberships | Get all memberships for the organization
[**ConsoleListOrganizationProjects**](ConsoleOrganizationsAPI.md#ConsoleListOrganizationProjects) | **Get** /v1/console/organizations/projects | Get all projects for the organization
[**ConsoleListProjectMemberships**](ConsoleOrganizationsAPI.md#ConsoleListProjectMemberships) | **Get** /v1/console/projects/{projectId}/memberships | Get all memberships for a project
[**ConsoleUpdateOrganizationMembership**](ConsoleOrganizationsAPI.md#ConsoleUpdateOrganizationMembership) | **Put** /v1/console/organizations/memberships | Create or update an organization membership
[**ConsoleUpdateProjectMembership**](ConsoleOrganizationsAPI.md#ConsoleUpdateProjectMembership) | **Put** /v1/console/projects/{projectId}/memberships | Create or update a project membership



## ConsoleDeleteOrganizationMembership

> ConsoleDeleteOrganizationMembership200Response ConsoleDeleteOrganizationMembership(ctx).ConsoleDeleteOrganizationMembershipRequest(consoleDeleteOrganizationMembershipRequest).Execute()

Delete an organization membership

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
	consoleDeleteOrganizationMembershipRequest := *openapiclient.NewConsoleDeleteOrganizationMembershipRequest("UserId_example") // ConsoleDeleteOrganizationMembershipRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleOrganizationsAPI.ConsoleDeleteOrganizationMembership(context.Background()).ConsoleDeleteOrganizationMembershipRequest(consoleDeleteOrganizationMembershipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleOrganizationsAPI.ConsoleDeleteOrganizationMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteOrganizationMembership`: ConsoleDeleteOrganizationMembership200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleOrganizationsAPI.ConsoleDeleteOrganizationMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteOrganizationMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleDeleteOrganizationMembershipRequest** | [**ConsoleDeleteOrganizationMembershipRequest**](ConsoleDeleteOrganizationMembershipRequest.md) |  | 

### Return type

[**ConsoleDeleteOrganizationMembership200Response**](ConsoleDeleteOrganizationMembership200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleDeleteProjectMembership

> ConsoleDeleteOrganizationMembership200Response ConsoleDeleteProjectMembership(ctx, projectId).ConsoleDeleteOrganizationMembershipRequest(consoleDeleteOrganizationMembershipRequest).Execute()

Delete a project membership

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
	consoleDeleteOrganizationMembershipRequest := *openapiclient.NewConsoleDeleteOrganizationMembershipRequest("UserId_example") // ConsoleDeleteOrganizationMembershipRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleOrganizationsAPI.ConsoleDeleteProjectMembership(context.Background(), projectId).ConsoleDeleteOrganizationMembershipRequest(consoleDeleteOrganizationMembershipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleOrganizationsAPI.ConsoleDeleteProjectMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteProjectMembership`: ConsoleDeleteOrganizationMembership200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleOrganizationsAPI.ConsoleDeleteProjectMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteProjectMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consoleDeleteOrganizationMembershipRequest** | [**ConsoleDeleteOrganizationMembershipRequest**](ConsoleDeleteOrganizationMembershipRequest.md) |  | 

### Return type

[**ConsoleDeleteOrganizationMembership200Response**](ConsoleDeleteOrganizationMembership200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListOrganizationApiKeys

> ConsoleListOrganizationApiKeys200Response ConsoleListOrganizationApiKeys(ctx).Execute()

Get all API keys for the organization

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
	resp, r, err := apiClient.ConsoleOrganizationsAPI.ConsoleListOrganizationApiKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleOrganizationsAPI.ConsoleListOrganizationApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListOrganizationApiKeys`: ConsoleListOrganizationApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleOrganizationsAPI.ConsoleListOrganizationApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListOrganizationApiKeysRequest struct via the builder pattern


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


## ConsoleListOrganizationMemberships

> ConsoleListOrganizationMemberships200Response ConsoleListOrganizationMemberships(ctx).Execute()

Get all memberships for the organization

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
	resp, r, err := apiClient.ConsoleOrganizationsAPI.ConsoleListOrganizationMemberships(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleOrganizationsAPI.ConsoleListOrganizationMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListOrganizationMemberships`: ConsoleListOrganizationMemberships200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleOrganizationsAPI.ConsoleListOrganizationMemberships`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListOrganizationMembershipsRequest struct via the builder pattern


### Return type

[**ConsoleListOrganizationMemberships200Response**](ConsoleListOrganizationMemberships200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListOrganizationProjects

> ConsoleListOrganizationProjects200Response ConsoleListOrganizationProjects(ctx).Execute()

Get all projects for the organization

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
	resp, r, err := apiClient.ConsoleOrganizationsAPI.ConsoleListOrganizationProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleOrganizationsAPI.ConsoleListOrganizationProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListOrganizationProjects`: ConsoleListOrganizationProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleOrganizationsAPI.ConsoleListOrganizationProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListOrganizationProjectsRequest struct via the builder pattern


### Return type

[**ConsoleListOrganizationProjects200Response**](ConsoleListOrganizationProjects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListProjectMemberships

> ConsoleListOrganizationMemberships200Response ConsoleListProjectMemberships(ctx, projectId).Execute()

Get all memberships for a project

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
	resp, r, err := apiClient.ConsoleOrganizationsAPI.ConsoleListProjectMemberships(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleOrganizationsAPI.ConsoleListProjectMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListProjectMemberships`: ConsoleListOrganizationMemberships200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleOrganizationsAPI.ConsoleListProjectMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListProjectMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleListOrganizationMemberships200Response**](ConsoleListOrganizationMemberships200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleUpdateOrganizationMembership

> ConsoleMembership ConsoleUpdateOrganizationMembership(ctx).ConsoleUpdateOrganizationMembershipRequest(consoleUpdateOrganizationMembershipRequest).Execute()

Create or update an organization membership

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
	consoleUpdateOrganizationMembershipRequest := *openapiclient.NewConsoleUpdateOrganizationMembershipRequest("UserId_example", "Role_example") // ConsoleUpdateOrganizationMembershipRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleOrganizationsAPI.ConsoleUpdateOrganizationMembership(context.Background()).ConsoleUpdateOrganizationMembershipRequest(consoleUpdateOrganizationMembershipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleOrganizationsAPI.ConsoleUpdateOrganizationMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleUpdateOrganizationMembership`: ConsoleMembership
	fmt.Fprintf(os.Stdout, "Response from `ConsoleOrganizationsAPI.ConsoleUpdateOrganizationMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleUpdateOrganizationMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleUpdateOrganizationMembershipRequest** | [**ConsoleUpdateOrganizationMembershipRequest**](ConsoleUpdateOrganizationMembershipRequest.md) |  | 

### Return type

[**ConsoleMembership**](ConsoleMembership.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleUpdateProjectMembership

> ConsoleMembership ConsoleUpdateProjectMembership(ctx, projectId).ConsoleUpdateOrganizationMembershipRequest(consoleUpdateOrganizationMembershipRequest).Execute()

Create or update a project membership

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
	consoleUpdateOrganizationMembershipRequest := *openapiclient.NewConsoleUpdateOrganizationMembershipRequest("UserId_example", "Role_example") // ConsoleUpdateOrganizationMembershipRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleOrganizationsAPI.ConsoleUpdateProjectMembership(context.Background(), projectId).ConsoleUpdateOrganizationMembershipRequest(consoleUpdateOrganizationMembershipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleOrganizationsAPI.ConsoleUpdateProjectMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleUpdateProjectMembership`: ConsoleMembership
	fmt.Fprintf(os.Stdout, "Response from `ConsoleOrganizationsAPI.ConsoleUpdateProjectMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleUpdateProjectMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consoleUpdateOrganizationMembershipRequest** | [**ConsoleUpdateOrganizationMembershipRequest**](ConsoleUpdateOrganizationMembershipRequest.md) |  | 

### Return type

[**ConsoleMembership**](ConsoleMembership.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

