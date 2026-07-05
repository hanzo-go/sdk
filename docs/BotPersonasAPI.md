# \BotPersonasAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotGetPersona**](BotPersonasAPI.md#BotGetPersona) | **Get** /v1/bot/personas/{slug}/detail | Get persona detail including latest version and owner
[**BotListPersonaVersions**](BotPersonasAPI.md#BotListPersonaVersions) | **Get** /v1/bot/personas/{slug}/versions | List versions of a persona
[**BotListPersonas**](BotPersonasAPI.md#BotListPersonas) | **Get** /v1/bot/personas | List personas (paginated)



## BotGetPersona

> BotGetPersona200Response BotGetPersona(ctx, slug).Execute()

Get persona detail including latest version and owner

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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotPersonasAPI.BotGetPersona(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotPersonasAPI.BotGetPersona``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetPersona`: BotGetPersona200Response
	fmt.Fprintf(os.Stdout, "Response from `BotPersonasAPI.BotGetPersona`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGetPersonaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotGetPersona200Response**](BotGetPersona200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListPersonaVersions

> BotListPersonaVersions200Response BotListPersonaVersions(ctx, slug).Limit(limit).Execute()

List versions of a persona

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
	slug := "slug_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotPersonasAPI.BotListPersonaVersions(context.Background(), slug).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotPersonasAPI.BotListPersonaVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListPersonaVersions`: BotListPersonaVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `BotPersonasAPI.BotListPersonaVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotListPersonaVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 20]

### Return type

[**BotListPersonaVersions200Response**](BotListPersonaVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListPersonas

> BotListPersonas200Response BotListPersonas(ctx).Sort(sort).Limit(limit).Cursor(cursor).Execute()

List personas (paginated)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	sort := "sort_example" // string |  (optional) (default to "updated")
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotPersonasAPI.BotListPersonas(context.Background()).Sort(sort).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotPersonasAPI.BotListPersonas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListPersonas`: BotListPersonas200Response
	fmt.Fprintf(os.Stdout, "Response from `BotPersonasAPI.BotListPersonas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotListPersonasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** |  | [default to &quot;updated&quot;]
 **limit** | **int32** |  | [default to 50]
 **cursor** | **time.Time** |  | 

### Return type

[**BotListPersonas200Response**](BotListPersonas200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

