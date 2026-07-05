# \DidDirectoriesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DidCreateTeam**](DidDirectoriesAPI.md#DidCreateTeam) | **Post** /v1/did/directories/{organization}/teams | Create a team
[**DidGetDirectory**](DidDirectoriesAPI.md#DidGetDirectory) | **Get** /v1/did/directories/{organization} | Get organization directory
[**DidGetTeamMembers**](DidDirectoriesAPI.md#DidGetTeamMembers) | **Get** /v1/did/directories/{organization}/teams/{team} | Get team members
[**DidListTeams**](DidDirectoriesAPI.md#DidListTeams) | **Get** /v1/did/directories/{organization}/teams | List teams



## DidCreateTeam

> DidTeam DidCreateTeam(ctx, organization).DidCreateTeamRequest(didCreateTeamRequest).Execute()

Create a team

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
	organization := "organization_example" // string | 
	didCreateTeamRequest := *openapiclient.NewDidCreateTeamRequest("Name_example") // DidCreateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidDirectoriesAPI.DidCreateTeam(context.Background(), organization).DidCreateTeamRequest(didCreateTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidDirectoriesAPI.DidCreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidCreateTeam`: DidTeam
	fmt.Fprintf(os.Stdout, "Response from `DidDirectoriesAPI.DidCreateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **didCreateTeamRequest** | [**DidCreateTeamRequest**](DidCreateTeamRequest.md) |  | 

### Return type

[**DidTeam**](DidTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidGetDirectory

> DidDirectory DidGetDirectory(ctx, organization).Execute()

Get organization directory

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
	organization := "organization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidDirectoriesAPI.DidGetDirectory(context.Background(), organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidDirectoriesAPI.DidGetDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidGetDirectory`: DidDirectory
	fmt.Fprintf(os.Stdout, "Response from `DidDirectoriesAPI.DidGetDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidGetDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DidDirectory**](DidDirectory.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidGetTeamMembers

> DidGetTeamMembers200Response DidGetTeamMembers(ctx, organization, team).Execute()

Get team members

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
	organization := "organization_example" // string | 
	team := "team_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidDirectoriesAPI.DidGetTeamMembers(context.Background(), organization, team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidDirectoriesAPI.DidGetTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidGetTeamMembers`: DidGetTeamMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DidDirectoriesAPI.DidGetTeamMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**team** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidGetTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DidGetTeamMembers200Response**](DidGetTeamMembers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidListTeams

> DidListTeams200Response DidListTeams(ctx, organization).Execute()

List teams

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
	organization := "organization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidDirectoriesAPI.DidListTeams(context.Background(), organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidDirectoriesAPI.DidListTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidListTeams`: DidListTeams200Response
	fmt.Fprintf(os.Stdout, "Response from `DidDirectoriesAPI.DidListTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidListTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DidListTeams200Response**](DidListTeams200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

