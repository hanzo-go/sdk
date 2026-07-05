# \ConsolePromptsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreatePrompt**](ConsolePromptsAPI.md#ConsoleCreatePrompt) | **Post** /v1/console/prompts | Create a new prompt version
[**ConsoleDeletePrompt**](ConsolePromptsAPI.md#ConsoleDeletePrompt) | **Delete** /v1/console/prompts/{promptName} | Delete prompt versions
[**ConsoleGetPrompt**](ConsolePromptsAPI.md#ConsoleGetPrompt) | **Get** /v1/console/prompts/{promptName} | Get a prompt by name
[**ConsoleGetPromptVersion**](ConsolePromptsAPI.md#ConsoleGetPromptVersion) | **Get** /v1/console/prompts/{promptName}/versions/{promptVersion} | Get a specific prompt version
[**ConsoleListPrompts**](ConsolePromptsAPI.md#ConsoleListPrompts) | **Get** /v1/console/prompts | Get a list of prompt names with versions and labels



## ConsoleCreatePrompt

> ConsolePrompt ConsoleCreatePrompt(ctx).ConsoleCreatePromptRequest(consoleCreatePromptRequest).Execute()

Create a new prompt version

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
	consoleCreatePromptRequest := *openapiclient.NewConsoleCreatePromptRequest("Name_example", interface{}(123), "Type_example") // ConsoleCreatePromptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsolePromptsAPI.ConsoleCreatePrompt(context.Background()).ConsoleCreatePromptRequest(consoleCreatePromptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsolePromptsAPI.ConsoleCreatePrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreatePrompt`: ConsolePrompt
	fmt.Fprintf(os.Stdout, "Response from `ConsolePromptsAPI.ConsoleCreatePrompt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreatePromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreatePromptRequest** | [**ConsoleCreatePromptRequest**](ConsoleCreatePromptRequest.md) |  | 

### Return type

[**ConsolePrompt**](ConsolePrompt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleDeletePrompt

> map[string]interface{} ConsoleDeletePrompt(ctx, promptName).Label(label).Version(version).Execute()

Delete prompt versions



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
	promptName := "promptName_example" // string | 
	label := "label_example" // string |  (optional)
	version := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsolePromptsAPI.ConsoleDeletePrompt(context.Background(), promptName).Label(label).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsolePromptsAPI.ConsoleDeletePrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeletePrompt`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConsolePromptsAPI.ConsoleDeletePrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeletePromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **string** |  | 
 **version** | **int32** |  | 

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


## ConsoleGetPrompt

> ConsolePrompt ConsoleGetPrompt(ctx, promptName).Version(version).Label(label).Execute()

Get a prompt by name

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
	promptName := "promptName_example" // string | 
	version := int32(56) // int32 |  (optional)
	label := "label_example" // string | Defaults to \"production\" if no label or version is set (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsolePromptsAPI.ConsoleGetPrompt(context.Background(), promptName).Version(version).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsolePromptsAPI.ConsoleGetPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetPrompt`: ConsolePrompt
	fmt.Fprintf(os.Stdout, "Response from `ConsolePromptsAPI.ConsoleGetPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 
 **label** | **string** | Defaults to \&quot;production\&quot; if no label or version is set | 

### Return type

[**ConsolePrompt**](ConsolePrompt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetPromptVersion

> ConsolePrompt ConsoleGetPromptVersion(ctx, promptName, promptVersion).Execute()

Get a specific prompt version

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
	promptName := "promptName_example" // string | 
	promptVersion := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsolePromptsAPI.ConsoleGetPromptVersion(context.Background(), promptName, promptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsolePromptsAPI.ConsoleGetPromptVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetPromptVersion`: ConsolePrompt
	fmt.Fprintf(os.Stdout, "Response from `ConsolePromptsAPI.ConsoleGetPromptVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptName** | **string** |  | 
**promptVersion** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetPromptVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsolePrompt**](ConsolePrompt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListPrompts

> ConsoleListPrompts200Response ConsoleListPrompts(ctx).Name(name).Label(label).Tag(tag).Page(page).Limit(limit).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).Execute()

Get a list of prompt names with versions and labels

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
	name := "name_example" // string |  (optional)
	label := "label_example" // string |  (optional)
	tag := "tag_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	fromUpdatedAt := time.Now() // time.Time |  (optional)
	toUpdatedAt := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsolePromptsAPI.ConsoleListPrompts(context.Background()).Name(name).Label(label).Tag(tag).Page(page).Limit(limit).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsolePromptsAPI.ConsoleListPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListPrompts`: ConsoleListPrompts200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsolePromptsAPI.ConsoleListPrompts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListPromptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **label** | **string** |  | 
 **tag** | **string** |  | 
 **page** | **int32** |  | 
 **limit** | **int32** |  | 
 **fromUpdatedAt** | **time.Time** |  | 
 **toUpdatedAt** | **time.Time** |  | 

### Return type

[**ConsoleListPrompts200Response**](ConsoleListPrompts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

