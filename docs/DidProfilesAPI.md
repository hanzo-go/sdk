# \DidProfilesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DidGetProfile**](DidProfilesAPI.md#DidGetProfile) | **Get** /v1/did/profiles/{profile_id} | Get a profile
[**DidGetProfileHistory**](DidProfilesAPI.md#DidGetProfileHistory) | **Get** /v1/did/profiles/{profile_id}/history | Get profile change history
[**DidListProfiles**](DidProfilesAPI.md#DidListProfiles) | **Get** /v1/did/profiles | List profiles
[**DidUpdateProfile**](DidProfilesAPI.md#DidUpdateProfile) | **Put** /v1/did/profiles/{profile_id} | Update a profile



## DidGetProfile

> DidProfile DidGetProfile(ctx, profileId).Execute()

Get a profile

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
	profileId := "profileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidProfilesAPI.DidGetProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidProfilesAPI.DidGetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidGetProfile`: DidProfile
	fmt.Fprintf(os.Stdout, "Response from `DidProfilesAPI.DidGetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DidProfile**](DidProfile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidGetProfileHistory

> DidGetProfileHistory200Response DidGetProfileHistory(ctx, profileId).Limit(limit).Execute()

Get profile change history

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
	profileId := "profileId_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidProfilesAPI.DidGetProfileHistory(context.Background(), profileId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidProfilesAPI.DidGetProfileHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidGetProfileHistory`: DidGetProfileHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `DidProfilesAPI.DidGetProfileHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidGetProfileHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]

### Return type

[**DidGetProfileHistory200Response**](DidGetProfileHistory200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidListProfiles

> DidListProfiles200Response DidListProfiles(ctx).Organization(organization).Type_(type_).Team(team).Status(status).Search(search).Limit(limit).Cursor(cursor).Execute()

List profiles



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
	organization := "organization_example" // string | Organization to list profiles for
	type_ := "type__example" // string | Filter by profile type (optional)
	team := "team_example" // string | Filter by team membership (optional)
	status := "status_example" // string | Filter by status (optional)
	search := "search_example" // string | Search by name or email (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidProfilesAPI.DidListProfiles(context.Background()).Organization(organization).Type_(type_).Team(team).Status(status).Search(search).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidProfilesAPI.DidListProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidListProfiles`: DidListProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `DidProfilesAPI.DidListProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDidListProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | **string** | Organization to list profiles for | 
 **type_** | **string** | Filter by profile type | 
 **team** | **string** | Filter by team membership | 
 **status** | **string** | Filter by status | 
 **search** | **string** | Search by name or email | 
 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**DidListProfiles200Response**](DidListProfiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidUpdateProfile

> DidProfile DidUpdateProfile(ctx, profileId).DidUpdateProfileRequest(didUpdateProfileRequest).Execute()

Update a profile

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
	profileId := "profileId_example" // string | 
	didUpdateProfileRequest := *openapiclient.NewDidUpdateProfileRequest() // DidUpdateProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidProfilesAPI.DidUpdateProfile(context.Background(), profileId).DidUpdateProfileRequest(didUpdateProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidProfilesAPI.DidUpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidUpdateProfile`: DidProfile
	fmt.Fprintf(os.Stdout, "Response from `DidProfilesAPI.DidUpdateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **didUpdateProfileRequest** | [**DidUpdateProfileRequest**](DidUpdateProfileRequest.md) |  | 

### Return type

[**DidProfile**](DidProfile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

