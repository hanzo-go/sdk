# \GatewayTeamsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayCreateTeam**](GatewayTeamsAPI.md#GatewayCreateTeam) | **Post** /v1/gateway/team/new | Create team
[**GatewayGetTeamInfo**](GatewayTeamsAPI.md#GatewayGetTeamInfo) | **Get** /v1/gateway/team/info | Get team info
[**GatewayListTeams**](GatewayTeamsAPI.md#GatewayListTeams) | **Get** /v1/gateway/team/list | List teams



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
	resp, r, err := apiClient.GatewayTeamsAPI.GatewayCreateTeam(context.Background()).GatewayCreateTeamRequest(gatewayCreateTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTeamsAPI.GatewayCreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayCreateTeam`: GatewayTeam
	fmt.Fprintf(os.Stdout, "Response from `GatewayTeamsAPI.GatewayCreateTeam`: %v\n", resp)
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
	resp, r, err := apiClient.GatewayTeamsAPI.GatewayGetTeamInfo(context.Background()).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTeamsAPI.GatewayGetTeamInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGetTeamInfo`: GatewayTeam
	fmt.Fprintf(os.Stdout, "Response from `GatewayTeamsAPI.GatewayGetTeamInfo`: %v\n", resp)
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
	resp, r, err := apiClient.GatewayTeamsAPI.GatewayListTeams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTeamsAPI.GatewayListTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayListTeams`: []GatewayTeam
	fmt.Fprintf(os.Stdout, "Response from `GatewayTeamsAPI.GatewayListTeams`: %v\n", resp)
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

