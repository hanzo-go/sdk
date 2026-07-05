# \ChatActionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteAgentsActionsByagentIdByactionId**](ChatActionsAPI.md#ChatDeleteAgentsActionsByagentIdByactionId) | **Delete** /v1/chat/agents/actions/{agent_id}/{action_id} | Delete an agent action
[**ChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodel**](ChatActionsAPI.md#ChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodel) | **Delete** /v1/chat/assistants/v1/actions/{assistant_id}/{action_id}/{model} | Delete an assistant action (v1)
[**ChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodel**](ChatActionsAPI.md#ChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodel) | **Delete** /v1/chat/assistants/v2/actions/{assistant_id}/{action_id}/{model} | Delete an assistant action (v2)
[**ChatGetActionsByactionIdOauthCallback**](ChatActionsAPI.md#ChatGetActionsByactionIdOauthCallback) | **Get** /v1/chat/actions/{action_id}/oauth/callback | Action OAuth callback
[**ChatGetAgentsActions**](ChatActionsAPI.md#ChatGetAgentsActions) | **Get** /v1/chat/agents/actions | List agent actions
[**ChatPostActionsByactionIdOauthBind**](ChatActionsAPI.md#ChatPostActionsByactionIdOauthBind) | **Post** /v1/chat/actions/{action_id}/oauth/bind | Set CSRF cookie for action OAuth flow
[**ChatPostAgentsActionsByagentId**](ChatActionsAPI.md#ChatPostAgentsActionsByagentId) | **Post** /v1/chat/agents/actions/{agent_id} | Add or update actions for an agent
[**ChatPostAssistantsV1ActionsByassistantId**](ChatActionsAPI.md#ChatPostAssistantsV1ActionsByassistantId) | **Post** /v1/chat/assistants/v1/actions/{assistant_id} | Add or update actions for an assistant (v1)
[**ChatPostAssistantsV2ActionsByassistantId**](ChatActionsAPI.md#ChatPostAssistantsV2ActionsByassistantId) | **Post** /v1/chat/assistants/v2/actions/{assistant_id} | Add or update actions for an assistant (v2)



## ChatDeleteAgentsActionsByagentIdByactionId

> map[string]interface{} ChatDeleteAgentsActionsByagentIdByactionId(ctx, agentId, actionId).Execute()

Delete an agent action

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
	agentId := "agentId_example" // string | 
	actionId := "actionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatActionsAPI.ChatDeleteAgentsActionsByagentIdByactionId(context.Background(), agentId, actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatDeleteAgentsActionsByagentIdByactionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteAgentsActionsByagentIdByactionId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatActionsAPI.ChatDeleteAgentsActionsByagentIdByactionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteAgentsActionsByagentIdByactionIdRequest struct via the builder pattern


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


## ChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodel

> map[string]interface{} ChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodel(ctx, assistantId, actionId, model).Execute()

Delete an assistant action (v1)

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
	assistantId := "assistantId_example" // string | 
	actionId := "actionId_example" // string | 
	model := "model_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatActionsAPI.ChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodel(context.Background(), assistantId, actionId, model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatActionsAPI.ChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 
**actionId** | **string** |  | 
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteAssistantsV1ActionsByassistantIdByactionIdBymodelRequest struct via the builder pattern


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


## ChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodel

> map[string]interface{} ChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodel(ctx, assistantId, actionId, model).Execute()

Delete an assistant action (v2)

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
	assistantId := "assistantId_example" // string | 
	actionId := "actionId_example" // string | 
	model := "model_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatActionsAPI.ChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodel(context.Background(), assistantId, actionId, model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatActionsAPI.ChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 
**actionId** | **string** |  | 
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteAssistantsV2ActionsByassistantIdByactionIdBymodelRequest struct via the builder pattern


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


## ChatGetActionsByactionIdOauthCallback

> ChatGetActionsByactionIdOauthCallback(ctx, actionId).Code(code).State(state).Execute()

Action OAuth callback

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
	actionId := "actionId_example" // string | 
	code := "code_example" // string |  (optional)
	state := "state_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChatActionsAPI.ChatGetActionsByactionIdOauthCallback(context.Background(), actionId).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatGetActionsByactionIdOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetActionsByactionIdOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** |  | 
 **state** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsActions

> []ChatAction ChatGetAgentsActions(ctx).Execute()

List agent actions

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
	resp, r, err := apiClient.ChatActionsAPI.ChatGetAgentsActions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatGetAgentsActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsActions`: []ChatAction
	fmt.Fprintf(os.Stdout, "Response from `ChatActionsAPI.ChatGetAgentsActions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsActionsRequest struct via the builder pattern


### Return type

[**[]ChatAction**](ChatAction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostActionsByactionIdOauthBind

> map[string]interface{} ChatPostActionsByactionIdOauthBind(ctx, actionId).Execute()

Set CSRF cookie for action OAuth flow

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
	actionId := "actionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatActionsAPI.ChatPostActionsByactionIdOauthBind(context.Background(), actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatPostActionsByactionIdOauthBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostActionsByactionIdOauthBind`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatActionsAPI.ChatPostActionsByactionIdOauthBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostActionsByactionIdOauthBindRequest struct via the builder pattern


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


## ChatPostAgentsActionsByagentId

> map[string]interface{} ChatPostAgentsActionsByagentId(ctx, agentId).ChatPostAgentsActionsByagentIdRequest(chatPostAgentsActionsByagentIdRequest).Execute()

Add or update actions for an agent

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
	agentId := "agentId_example" // string | 
	chatPostAgentsActionsByagentIdRequest := *openapiclient.NewChatPostAgentsActionsByagentIdRequest([]openapiclient.ChatFunctionTool{*openapiclient.NewChatFunctionTool()}, *openapiclient.NewChatActionMetadata()) // ChatPostAgentsActionsByagentIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatActionsAPI.ChatPostAgentsActionsByagentId(context.Background(), agentId).ChatPostAgentsActionsByagentIdRequest(chatPostAgentsActionsByagentIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatPostAgentsActionsByagentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsActionsByagentId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatActionsAPI.ChatPostAgentsActionsByagentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsActionsByagentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatPostAgentsActionsByagentIdRequest** | [**ChatPostAgentsActionsByagentIdRequest**](ChatPostAgentsActionsByagentIdRequest.md) |  | 

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


## ChatPostAssistantsV1ActionsByassistantId

> map[string]interface{} ChatPostAssistantsV1ActionsByassistantId(ctx, assistantId).ChatPostAgentsActionsByagentIdRequest(chatPostAgentsActionsByagentIdRequest).Execute()

Add or update actions for an assistant (v1)

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
	assistantId := "assistantId_example" // string | 
	chatPostAgentsActionsByagentIdRequest := *openapiclient.NewChatPostAgentsActionsByagentIdRequest([]openapiclient.ChatFunctionTool{*openapiclient.NewChatFunctionTool()}, *openapiclient.NewChatActionMetadata()) // ChatPostAgentsActionsByagentIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatActionsAPI.ChatPostAssistantsV1ActionsByassistantId(context.Background(), assistantId).ChatPostAgentsActionsByagentIdRequest(chatPostAgentsActionsByagentIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatPostAssistantsV1ActionsByassistantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV1ActionsByassistantId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatActionsAPI.ChatPostAssistantsV1ActionsByassistantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV1ActionsByassistantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatPostAgentsActionsByagentIdRequest** | [**ChatPostAgentsActionsByagentIdRequest**](ChatPostAgentsActionsByagentIdRequest.md) |  | 

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


## ChatPostAssistantsV2ActionsByassistantId

> map[string]interface{} ChatPostAssistantsV2ActionsByassistantId(ctx, assistantId).Body(body).Execute()

Add or update actions for an assistant (v2)

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
	assistantId := "assistantId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatActionsAPI.ChatPostAssistantsV2ActionsByassistantId(context.Background(), assistantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatActionsAPI.ChatPostAssistantsV2ActionsByassistantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAssistantsV2ActionsByassistantId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatActionsAPI.ChatPostAssistantsV2ActionsByassistantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAssistantsV2ActionsByassistantIdRequest struct via the builder pattern


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

