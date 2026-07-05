# \AnalyticsTeamsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsAddTeamUser**](AnalyticsTeamsAPI.md#AnalyticsAddTeamUser) | **Post** /v1/analytics/teams/{teamId}/users | Add a user to a team
[**AnalyticsCreateTeam**](AnalyticsTeamsAPI.md#AnalyticsCreateTeam) | **Post** /v1/analytics/teams | Create a new team
[**AnalyticsDeleteTeam**](AnalyticsTeamsAPI.md#AnalyticsDeleteTeam) | **Delete** /v1/analytics/teams/{teamId} | Delete team (owner only)
[**AnalyticsGetTeam**](AnalyticsTeamsAPI.md#AnalyticsGetTeam) | **Get** /v1/analytics/teams/{teamId} | Get team by ID (includes members)
[**AnalyticsGetTeamUser**](AnalyticsTeamsAPI.md#AnalyticsGetTeamUser) | **Get** /v1/analytics/teams/{teamId}/users/{userId} | Get a team member
[**AnalyticsGetTeamUsers**](AnalyticsTeamsAPI.md#AnalyticsGetTeamUsers) | **Get** /v1/analytics/teams/{teamId}/users | List team members
[**AnalyticsGetTeamWebsites**](AnalyticsTeamsAPI.md#AnalyticsGetTeamWebsites) | **Get** /v1/analytics/teams/{teamId}/websites | List websites belonging to a team
[**AnalyticsJoinTeam**](AnalyticsTeamsAPI.md#AnalyticsJoinTeam) | **Post** /v1/analytics/teams/join | Join a team using an access code
[**AnalyticsRemoveTeamUser**](AnalyticsTeamsAPI.md#AnalyticsRemoveTeamUser) | **Delete** /v1/analytics/teams/{teamId}/users/{userId} | Remove user from team
[**AnalyticsUpdateTeam**](AnalyticsTeamsAPI.md#AnalyticsUpdateTeam) | **Post** /v1/analytics/teams/{teamId} | Update team (owner only)
[**AnalyticsUpdateTeamUser**](AnalyticsTeamsAPI.md#AnalyticsUpdateTeamUser) | **Post** /v1/analytics/teams/{teamId}/users/{userId} | Update team member role (owner only)



## AnalyticsAddTeamUser

> AnalyticsTeamUser AnalyticsAddTeamUser(ctx, teamId).AnalyticsAddTeamUserRequest(analyticsAddTeamUserRequest).Execute()

Add a user to a team

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	analyticsAddTeamUserRequest := *openapiclient.NewAnalyticsAddTeamUserRequest("UserId_example", "Role_example") // AnalyticsAddTeamUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsAddTeamUser(context.Background(), teamId).AnalyticsAddTeamUserRequest(analyticsAddTeamUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsAddTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsAddTeamUser`: AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsAddTeamUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsAddTeamUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsAddTeamUserRequest** | [**AnalyticsAddTeamUserRequest**](AnalyticsAddTeamUserRequest.md) |  | 

### Return type

[**AnalyticsTeamUser**](AnalyticsTeamUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsCreateTeam

> AnalyticsTeam AnalyticsCreateTeam(ctx).AnalyticsCreateTeamRequest(analyticsCreateTeamRequest).Execute()

Create a new team

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
	analyticsCreateTeamRequest := *openapiclient.NewAnalyticsCreateTeamRequest("Name_example") // AnalyticsCreateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsCreateTeam(context.Background()).AnalyticsCreateTeamRequest(analyticsCreateTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsCreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCreateTeam`: AnalyticsTeam
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsCreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsCreateTeamRequest** | [**AnalyticsCreateTeamRequest**](AnalyticsCreateTeamRequest.md) |  | 

### Return type

[**AnalyticsTeam**](AnalyticsTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsDeleteTeam

> map[string]interface{} AnalyticsDeleteTeam(ctx, teamId).Execute()

Delete team (owner only)

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsDeleteTeam(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsDeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsDeleteTeam`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsDeleteTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsDeleteTeamRequest struct via the builder pattern


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


## AnalyticsGetTeam

> AnalyticsTeam AnalyticsGetTeam(ctx, teamId).Execute()

Get team by ID (includes members)

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsGetTeam(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsGetTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTeam`: AnalyticsTeam
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsGetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsTeam**](AnalyticsTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetTeamUser

> AnalyticsTeamUser AnalyticsGetTeamUser(ctx, teamId, userId).Execute()

Get a team member

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsGetTeamUser(context.Background(), teamId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsGetTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTeamUser`: AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsGetTeamUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetTeamUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AnalyticsTeamUser**](AnalyticsTeamUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetTeamUsers

> []AnalyticsTeamUser AnalyticsGetTeamUsers(ctx, teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List team members

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsGetTeamUsers(context.Background(), teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsGetTeamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTeamUsers`: []AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsGetTeamUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetTeamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsTeamUser**](AnalyticsTeamUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetTeamWebsites

> []AnalyticsWebsite AnalyticsGetTeamWebsites(ctx, teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List websites belonging to a team

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsGetTeamWebsites(context.Background(), teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsGetTeamWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTeamWebsites`: []AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsGetTeamWebsites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetTeamWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsJoinTeam

> AnalyticsTeamUser AnalyticsJoinTeam(ctx).AnalyticsJoinTeamRequest(analyticsJoinTeamRequest).Execute()

Join a team using an access code

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
	analyticsJoinTeamRequest := *openapiclient.NewAnalyticsJoinTeamRequest("AccessCode_example") // AnalyticsJoinTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsJoinTeam(context.Background()).AnalyticsJoinTeamRequest(analyticsJoinTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsJoinTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsJoinTeam`: AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsJoinTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsJoinTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsJoinTeamRequest** | [**AnalyticsJoinTeamRequest**](AnalyticsJoinTeamRequest.md) |  | 

### Return type

[**AnalyticsTeamUser**](AnalyticsTeamUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRemoveTeamUser

> map[string]interface{} AnalyticsRemoveTeamUser(ctx, teamId, userId).Execute()

Remove user from team

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsRemoveTeamUser(context.Background(), teamId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsRemoveTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRemoveTeamUser`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsRemoveTeamUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRemoveTeamUserRequest struct via the builder pattern


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


## AnalyticsUpdateTeam

> AnalyticsTeam AnalyticsUpdateTeam(ctx, teamId).AnalyticsUpdateTeamRequest(analyticsUpdateTeamRequest).Execute()

Update team (owner only)

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	analyticsUpdateTeamRequest := *openapiclient.NewAnalyticsUpdateTeamRequest() // AnalyticsUpdateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsUpdateTeam(context.Background(), teamId).AnalyticsUpdateTeamRequest(analyticsUpdateTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsUpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsUpdateTeam`: AnalyticsTeam
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsUpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsUpdateTeamRequest** | [**AnalyticsUpdateTeamRequest**](AnalyticsUpdateTeamRequest.md) |  | 

### Return type

[**AnalyticsTeam**](AnalyticsTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsUpdateTeamUser

> AnalyticsTeamUser AnalyticsUpdateTeamUser(ctx, teamId, userId).AnalyticsUpdateTeamUserRequest(analyticsUpdateTeamUserRequest).Execute()

Update team member role (owner only)

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
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	analyticsUpdateTeamUserRequest := *openapiclient.NewAnalyticsUpdateTeamUserRequest("Role_example") // AnalyticsUpdateTeamUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsTeamsAPI.AnalyticsUpdateTeamUser(context.Background(), teamId, userId).AnalyticsUpdateTeamUserRequest(analyticsUpdateTeamUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsTeamsAPI.AnalyticsUpdateTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsUpdateTeamUser`: AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsTeamsAPI.AnalyticsUpdateTeamUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsUpdateTeamUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **analyticsUpdateTeamUserRequest** | [**AnalyticsUpdateTeamUserRequest**](AnalyticsUpdateTeamUserRequest.md) |  | 

### Return type

[**AnalyticsTeamUser**](AnalyticsTeamUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

