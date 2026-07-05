# \BotSearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotSearchPersonas**](BotSearchAPI.md#BotSearchPersonas) | **Get** /v1/bot/search/personas | Lexical search for personas
[**BotSearchSkills**](BotSearchAPI.md#BotSearchSkills) | **Get** /v1/bot/search/skills | Hybrid vector + lexical search for skills



## BotSearchPersonas

> BotSearchPersonas200Response BotSearchPersonas(ctx).Q(q).Limit(limit).Execute()

Lexical search for personas

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
	q := "q_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotSearchAPI.BotSearchPersonas(context.Background()).Q(q).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotSearchAPI.BotSearchPersonas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotSearchPersonas`: BotSearchPersonas200Response
	fmt.Fprintf(os.Stdout, "Response from `BotSearchAPI.BotSearchPersonas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotSearchPersonasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**BotSearchPersonas200Response**](BotSearchPersonas200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotSearchSkills

> BotSearchPersonas200Response BotSearchSkills(ctx).Q(q).Limit(limit).Execute()

Hybrid vector + lexical search for skills

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
	q := "q_example" // string | Search query
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotSearchAPI.BotSearchSkills(context.Background()).Q(q).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotSearchAPI.BotSearchSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotSearchSkills`: BotSearchPersonas200Response
	fmt.Fprintf(os.Stdout, "Response from `BotSearchAPI.BotSearchSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotSearchSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search query | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**BotSearchPersonas200Response**](BotSearchPersonas200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

