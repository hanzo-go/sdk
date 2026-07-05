# \ChatPromptsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeletePromptsBypromptid**](ChatPromptsAPI.md#ChatDeletePromptsBypromptid) | **Delete** /v1/chat/prompts/{promptId} | Delete a prompt
[**ChatDeletePromptsGroupsBygroupid**](ChatPromptsAPI.md#ChatDeletePromptsGroupsBygroupid) | **Delete** /v1/chat/prompts/groups/{groupId} | Delete a prompt group
[**ChatGetPrompts**](ChatPromptsAPI.md#ChatGetPrompts) | **Get** /v1/chat/prompts | List user prompts
[**ChatGetPromptsAll**](ChatPromptsAPI.md#ChatGetPromptsAll) | **Get** /v1/chat/prompts/all | Get all prompt groups (ACL-aware)
[**ChatGetPromptsBypromptid**](ChatPromptsAPI.md#ChatGetPromptsBypromptid) | **Get** /v1/chat/prompts/{promptId} | Get a prompt
[**ChatGetPromptsGroups**](ChatPromptsAPI.md#ChatGetPromptsGroups) | **Get** /v1/chat/prompts/groups | List prompt groups (paginated)
[**ChatGetPromptsGroupsBygroupid**](ChatPromptsAPI.md#ChatGetPromptsGroupsBygroupid) | **Get** /v1/chat/prompts/groups/{groupId} | Get a prompt group by ID
[**ChatPatchPromptsBypromptidTagsProduction**](ChatPromptsAPI.md#ChatPatchPromptsBypromptidTagsProduction) | **Patch** /v1/chat/prompts/{promptId}/tags/production | Make a prompt the production version
[**ChatPatchPromptsGroupsBygroupid**](ChatPromptsAPI.md#ChatPatchPromptsGroupsBygroupid) | **Patch** /v1/chat/prompts/groups/{groupId} | Update a prompt group
[**ChatPostPrompts**](ChatPromptsAPI.md#ChatPostPrompts) | **Post** /v1/chat/prompts | Create a new prompt group with initial prompt
[**ChatPostPromptsGroupsBygroupidPrompts**](ChatPromptsAPI.md#ChatPostPromptsGroupsBygroupidPrompts) | **Post** /v1/chat/prompts/groups/{groupId}/prompts | Add a prompt to an existing group



## ChatDeletePromptsBypromptid

> map[string]interface{} ChatDeletePromptsBypromptid(ctx, promptId).Execute()

Delete a prompt

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
	promptId := "promptId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatDeletePromptsBypromptid(context.Background(), promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatDeletePromptsBypromptid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeletePromptsBypromptid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatDeletePromptsBypromptid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeletePromptsBypromptidRequest struct via the builder pattern


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


## ChatDeletePromptsGroupsBygroupid

> map[string]interface{} ChatDeletePromptsGroupsBygroupid(ctx, groupId).Execute()

Delete a prompt group

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatDeletePromptsGroupsBygroupid(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatDeletePromptsGroupsBygroupid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeletePromptsGroupsBygroupid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatDeletePromptsGroupsBygroupid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeletePromptsGroupsBygroupidRequest struct via the builder pattern


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


## ChatGetPrompts

> map[string]interface{} ChatGetPrompts(ctx).GroupId(groupId).Execute()

List user prompts

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
	groupId := "groupId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatGetPrompts(context.Background()).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatGetPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPrompts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatGetPrompts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPromptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** |  | 

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


## ChatGetPromptsAll

> map[string]interface{} ChatGetPromptsAll(ctx).Name(name).Category(category).Execute()

Get all prompt groups (ACL-aware)

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
	name := "name_example" // string |  (optional)
	category := "category_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatGetPromptsAll(context.Background()).Name(name).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatGetPromptsAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPromptsAll`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatGetPromptsAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPromptsAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **category** | **string** |  | 

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


## ChatGetPromptsBypromptid

> map[string]interface{} ChatGetPromptsBypromptid(ctx, promptId).Execute()

Get a prompt

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
	promptId := "promptId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatGetPromptsBypromptid(context.Background(), promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatGetPromptsBypromptid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPromptsBypromptid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatGetPromptsBypromptid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPromptsBypromptidRequest struct via the builder pattern


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


## ChatGetPromptsGroups

> map[string]interface{} ChatGetPromptsGroups(ctx).PageSize(pageSize).Limit(limit).Cursor(cursor).Name(name).Category(category).Execute()

List prompt groups (paginated)

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
	pageSize := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	category := "category_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatGetPromptsGroups(context.Background()).PageSize(pageSize).Limit(limit).Cursor(cursor).Name(name).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatGetPromptsGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPromptsGroups`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatGetPromptsGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPromptsGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 
 **name** | **string** |  | 
 **category** | **string** |  | 

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


## ChatGetPromptsGroupsBygroupid

> map[string]interface{} ChatGetPromptsGroupsBygroupid(ctx, groupId).Execute()

Get a prompt group by ID

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatGetPromptsGroupsBygroupid(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatGetPromptsGroupsBygroupid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetPromptsGroupsBygroupid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatGetPromptsGroupsBygroupid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetPromptsGroupsBygroupidRequest struct via the builder pattern


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


## ChatPatchPromptsBypromptidTagsProduction

> map[string]interface{} ChatPatchPromptsBypromptidTagsProduction(ctx, promptId).Execute()

Make a prompt the production version

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
	promptId := "promptId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatPatchPromptsBypromptidTagsProduction(context.Background(), promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatPatchPromptsBypromptidTagsProduction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchPromptsBypromptidTagsProduction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatPatchPromptsBypromptidTagsProduction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchPromptsBypromptidTagsProductionRequest struct via the builder pattern


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


## ChatPatchPromptsGroupsBygroupid

> map[string]interface{} ChatPatchPromptsGroupsBygroupid(ctx, groupId).Body(body).Execute()

Update a prompt group

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
	groupId := "groupId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatPatchPromptsGroupsBygroupid(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatPatchPromptsGroupsBygroupid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchPromptsGroupsBygroupid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatPatchPromptsGroupsBygroupid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchPromptsGroupsBygroupidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostPrompts

> map[string]interface{} ChatPostPrompts(ctx).ChatPostPromptsRequest(chatPostPromptsRequest).Execute()

Create a new prompt group with initial prompt

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
	chatPostPromptsRequest := *openapiclient.NewChatPostPromptsRequest(map[string]interface{}(123), *openapiclient.NewChatPostPromptsRequestGroup("Name_example")) // ChatPostPromptsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatPostPrompts(context.Background()).ChatPostPromptsRequest(chatPostPromptsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatPostPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostPrompts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatPostPrompts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostPromptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostPromptsRequest** | [**ChatPostPromptsRequest**](ChatPostPromptsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostPromptsGroupsBygroupidPrompts

> map[string]interface{} ChatPostPromptsGroupsBygroupidPrompts(ctx, groupId).ChatPostPromptsGroupsBygroupidPromptsRequest(chatPostPromptsGroupsBygroupidPromptsRequest).Execute()

Add a prompt to an existing group

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
	groupId := "groupId_example" // string | 
	chatPostPromptsGroupsBygroupidPromptsRequest := *openapiclient.NewChatPostPromptsGroupsBygroupidPromptsRequest(map[string]interface{}(123)) // ChatPostPromptsGroupsBygroupidPromptsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatPromptsAPI.ChatPostPromptsGroupsBygroupidPrompts(context.Background(), groupId).ChatPostPromptsGroupsBygroupidPromptsRequest(chatPostPromptsGroupsBygroupidPromptsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatPromptsAPI.ChatPostPromptsGroupsBygroupidPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostPromptsGroupsBygroupidPrompts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatPromptsAPI.ChatPostPromptsGroupsBygroupidPrompts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostPromptsGroupsBygroupidPromptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatPostPromptsGroupsBygroupidPromptsRequest** | [**ChatPostPromptsGroupsBygroupidPromptsRequest**](ChatPostPromptsGroupsBygroupidPromptsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

