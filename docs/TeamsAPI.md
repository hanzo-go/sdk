# \TeamsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsAddTeamUser**](TeamsAPI.md#AnalyticsAddTeamUser) | **Post** /v1/analytics/teams/{teamId}/users | Add a user to a team
[**AnalyticsCreateTeam**](TeamsAPI.md#AnalyticsCreateTeam) | **Post** /v1/analytics/teams | Create a new team
[**AnalyticsDeleteTeam**](TeamsAPI.md#AnalyticsDeleteTeam) | **Delete** /v1/analytics/teams/{teamId} | Delete team (owner only)
[**AnalyticsGetTeam**](TeamsAPI.md#AnalyticsGetTeam) | **Get** /v1/analytics/teams/{teamId} | Get team by ID (includes members)
[**AnalyticsGetTeamUser**](TeamsAPI.md#AnalyticsGetTeamUser) | **Get** /v1/analytics/teams/{teamId}/users/{userId} | Get a team member
[**AnalyticsGetTeamUsers**](TeamsAPI.md#AnalyticsGetTeamUsers) | **Get** /v1/analytics/teams/{teamId}/users | List team members
[**AnalyticsGetTeamWebsites**](TeamsAPI.md#AnalyticsGetTeamWebsites) | **Get** /v1/analytics/teams/{teamId}/websites | List websites belonging to a team
[**AnalyticsJoinTeam**](TeamsAPI.md#AnalyticsJoinTeam) | **Post** /v1/analytics/teams/join | Join a team using an access code
[**AnalyticsRemoveTeamUser**](TeamsAPI.md#AnalyticsRemoveTeamUser) | **Delete** /v1/analytics/teams/{teamId}/users/{userId} | Remove user from team
[**AnalyticsUpdateTeam**](TeamsAPI.md#AnalyticsUpdateTeam) | **Post** /v1/analytics/teams/{teamId} | Update team (owner only)
[**AnalyticsUpdateTeamUser**](TeamsAPI.md#AnalyticsUpdateTeamUser) | **Post** /v1/analytics/teams/{teamId}/users/{userId} | Update team member role (owner only)
[**GatewayCreateTeam**](TeamsAPI.md#GatewayCreateTeam) | **Post** /v1/gateway/team/new | Create team
[**GatewayGetTeamInfo**](TeamsAPI.md#GatewayGetTeamInfo) | **Get** /v1/gateway/team/info | Get team info
[**GatewayListTeams**](TeamsAPI.md#GatewayListTeams) | **Get** /v1/gateway/team/list | List teams



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
	resp, r, err := apiClient.TeamsAPI.AnalyticsAddTeamUser(context.Background(), teamId).AnalyticsAddTeamUserRequest(analyticsAddTeamUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsAddTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsAddTeamUser`: AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsAddTeamUser`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsCreateTeam(context.Background()).AnalyticsCreateTeamRequest(analyticsCreateTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsCreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCreateTeam`: AnalyticsTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsCreateTeam`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsDeleteTeam(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsDeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsDeleteTeam`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsDeleteTeam`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsGetTeam(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsGetTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTeam`: AnalyticsTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsGetTeam`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsGetTeamUser(context.Background(), teamId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsGetTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTeamUser`: AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsGetTeamUser`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsGetTeamUsers(context.Background(), teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsGetTeamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTeamUsers`: []AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsGetTeamUsers`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsGetTeamWebsites(context.Background(), teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsGetTeamWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetTeamWebsites`: []AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsGetTeamWebsites`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsJoinTeam(context.Background()).AnalyticsJoinTeamRequest(analyticsJoinTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsJoinTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsJoinTeam`: AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsJoinTeam`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsRemoveTeamUser(context.Background(), teamId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsRemoveTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRemoveTeamUser`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsRemoveTeamUser`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsUpdateTeam(context.Background(), teamId).AnalyticsUpdateTeamRequest(analyticsUpdateTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsUpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsUpdateTeam`: AnalyticsTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsUpdateTeam`: %v\n", resp)
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
	resp, r, err := apiClient.TeamsAPI.AnalyticsUpdateTeamUser(context.Background(), teamId, userId).AnalyticsUpdateTeamUserRequest(analyticsUpdateTeamUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AnalyticsUpdateTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsUpdateTeamUser`: AnalyticsTeamUser
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AnalyticsUpdateTeamUser`: %v\n", resp)
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


## GatewayCreateTeam

> GatewayTeam GatewayCreateTeam(ctx).GatewayCreateTeamRequest(gatewayCreateTeamRequest).Execute()

Create team

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
	gatewayCreateTeamRequest := *openapiclient.NewGatewayCreateTeamRequest() // GatewayCreateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GatewayCreateTeam(context.Background()).GatewayCreateTeamRequest(gatewayCreateTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GatewayCreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayCreateTeam`: GatewayTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GatewayCreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayCreateTeamRequest** | [**GatewayCreateTeamRequest**](GatewayCreateTeamRequest.md) |  | 

### Return type

[**GatewayTeam**](GatewayTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetTeamInfo

> GatewayTeam GatewayGetTeamInfo(ctx).TeamId(teamId).Execute()

Get team info

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
	teamId := "teamId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GatewayGetTeamInfo(context.Background()).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GatewayGetTeamInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGetTeamInfo`: GatewayTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GatewayGetTeamInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetTeamInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string** |  | 

### Return type

[**GatewayTeam**](GatewayTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayListTeams

> []GatewayTeam GatewayListTeams(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GatewayListTeams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GatewayListTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayListTeams`: []GatewayTeam
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GatewayListTeams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListTeamsRequest struct via the builder pattern


### Return type

[**[]GatewayTeam**](GatewayTeam.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

