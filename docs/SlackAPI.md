# \SlackAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsSlackCommands**](SlackAPI.md#IntegrationsSlackCommands) | **Post** /v1/integrations/slack/commands | Slack slash command webhook
[**IntegrationsSlackEvents**](SlackAPI.md#IntegrationsSlackEvents) | **Post** /v1/integrations/slack/events | Slack Events API webhook
[**IntegrationsSlackLinkCallback**](SlackAPI.md#IntegrationsSlackLinkCallback) | **Get** /v1/integrations/slack/link/callback | hanzo.id OIDC callback — bind Slack↔Hanzo (leg 3)
[**IntegrationsSlackLinkSlack**](SlackAPI.md#IntegrationsSlackLinkSlack) | **Get** /v1/integrations/slack/link/slack | Slack sign-in callback (leg 2)
[**IntegrationsSlackLinkStart**](SlackAPI.md#IntegrationsSlackLinkStart) | **Get** /v1/integrations/slack/link | Begin the per-user account link (leg 1 — Slack sign-in)



## IntegrationsSlackCommands

> IntegrationsSlackCommands(ctx).Command(command).Text(text).TeamId(teamId).UserId(userId).ChannelId(channelId).ResponseUrl(responseUrl).Execute()

Slack slash command webhook



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
	command := "command_example" // string |  (optional)
	text := "text_example" // string |  (optional)
	teamId := "teamId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	channelId := "channelId_example" // string |  (optional)
	responseUrl := "responseUrl_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackAPI.IntegrationsSlackCommands(context.Background()).Command(command).Text(text).TeamId(teamId).UserId(userId).ChannelId(channelId).ResponseUrl(responseUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.IntegrationsSlackCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsSlackCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | **string** |  | 
 **text** | **string** |  | 
 **teamId** | **string** |  | 
 **userId** | **string** |  | 
 **channelId** | **string** |  | 
 **responseUrl** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsSlackEvents

> IntegrationsSlackEvents(ctx).IntegrationsSlackEventEnvelope(integrationsSlackEventEnvelope).Execute()

Slack Events API webhook



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
	integrationsSlackEventEnvelope := *openapiclient.NewIntegrationsSlackEventEnvelope() // IntegrationsSlackEventEnvelope | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackAPI.IntegrationsSlackEvents(context.Background()).IntegrationsSlackEventEnvelope(integrationsSlackEventEnvelope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.IntegrationsSlackEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsSlackEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationsSlackEventEnvelope** | [**IntegrationsSlackEventEnvelope**](IntegrationsSlackEventEnvelope.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsSlackLinkCallback

> string IntegrationsSlackLinkCallback(ctx).Code(code).State(state).Error_(error_).Execute()

hanzo.id OIDC callback — bind Slack↔Hanzo (leg 3)



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
	code := "code_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	error_ := "error__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackAPI.IntegrationsSlackLinkCallback(context.Background()).Code(code).State(state).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.IntegrationsSlackLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsSlackLinkCallback`: string
	fmt.Fprintf(os.Stdout, "Response from `SlackAPI.IntegrationsSlackLinkCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsSlackLinkCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **state** | **string** |  | 
 **error_** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsSlackLinkSlack

> IntegrationsSlackLinkSlack(ctx).Code(code).State(state).Error_(error_).Execute()

Slack sign-in callback (leg 2)



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
	code := "code_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	error_ := "error__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackAPI.IntegrationsSlackLinkSlack(context.Background()).Code(code).State(state).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.IntegrationsSlackLinkSlack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsSlackLinkSlackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **state** | **string** |  | 
 **error_** | **string** |  | 

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


## IntegrationsSlackLinkStart

> IntegrationsSlackLinkStart(ctx).State(state).Execute()

Begin the per-user account link (leg 1 — Slack sign-in)



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
	state := "state_example" // string | Signed, single-use link state

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackAPI.IntegrationsSlackLinkStart(context.Background()).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.IntegrationsSlackLinkStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsSlackLinkStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | Signed, single-use link state | 

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

