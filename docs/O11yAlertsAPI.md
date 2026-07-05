# \O11yAlertsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**O11yCreateAlertRule**](O11yAlertsAPI.md#O11yCreateAlertRule) | **Post** /v1/o11y/rules | Create alert rule
[**O11yCreateNotificationChannel**](O11yAlertsAPI.md#O11yCreateNotificationChannel) | **Post** /v1/o11y/channels | Create notification channel
[**O11yDeleteAlertRule**](O11yAlertsAPI.md#O11yDeleteAlertRule) | **Delete** /v1/o11y/rules/{id} | Delete alert rule
[**O11yDeleteNotificationChannel**](O11yAlertsAPI.md#O11yDeleteNotificationChannel) | **Delete** /v1/o11y/channels/{id} | Delete notification channel
[**O11yGetAlertRule**](O11yAlertsAPI.md#O11yGetAlertRule) | **Get** /v1/o11y/rules/{id} | Get alert rule
[**O11yListAlertRules**](O11yAlertsAPI.md#O11yListAlertRules) | **Get** /v1/o11y/rules | List alert rules
[**O11yListNotificationChannels**](O11yAlertsAPI.md#O11yListNotificationChannels) | **Get** /v1/o11y/channels | List notification channels
[**O11yUpdateAlertRule**](O11yAlertsAPI.md#O11yUpdateAlertRule) | **Put** /v1/o11y/rules/{id} | Update alert rule
[**O11yUpdateNotificationChannel**](O11yAlertsAPI.md#O11yUpdateNotificationChannel) | **Put** /v1/o11y/channels/{id} | Update notification channel



## O11yCreateAlertRule

> O11yAlertRule O11yCreateAlertRule(ctx).O11yAlertRuleCreate(o11yAlertRuleCreate).Execute()

Create alert rule

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
	o11yAlertRuleCreate := *openapiclient.NewO11yAlertRuleCreate("Name_example", "Datasource_example", "Query_example", *openapiclient.NewO11yAlertRuleCreateCondition(*openapiclient.NewO11yAlertRuleCreateConditionEvaluator("Type_example", []float32{float32(123)}))) // O11yAlertRuleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAlertsAPI.O11yCreateAlertRule(context.Background()).O11yAlertRuleCreate(o11yAlertRuleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yCreateAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yCreateAlertRule`: O11yAlertRule
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yCreateAlertRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yCreateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yAlertRuleCreate** | [**O11yAlertRuleCreate**](O11yAlertRuleCreate.md) |  | 

### Return type

[**O11yAlertRule**](O11yAlertRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yCreateNotificationChannel

> O11yNotificationChannel O11yCreateNotificationChannel(ctx).O11yNotificationChannelCreate(o11yNotificationChannelCreate).Execute()

Create notification channel

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
	o11yNotificationChannelCreate := *openapiclient.NewO11yNotificationChannelCreate("Name_example", "Type_example", *openapiclient.NewO11yNotificationChannelCreateConfig()) // O11yNotificationChannelCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAlertsAPI.O11yCreateNotificationChannel(context.Background()).O11yNotificationChannelCreate(o11yNotificationChannelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yCreateNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yCreateNotificationChannel`: O11yNotificationChannel
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yCreateNotificationChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yCreateNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yNotificationChannelCreate** | [**O11yNotificationChannelCreate**](O11yNotificationChannelCreate.md) |  | 

### Return type

[**O11yNotificationChannel**](O11yNotificationChannel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yDeleteAlertRule

> map[string]interface{} O11yDeleteAlertRule(ctx, id).Execute()

Delete alert rule

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAlertsAPI.O11yDeleteAlertRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yDeleteAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yDeleteAlertRule`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yDeleteAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yDeleteAlertRuleRequest struct via the builder pattern


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


## O11yDeleteNotificationChannel

> map[string]interface{} O11yDeleteNotificationChannel(ctx, id).Execute()

Delete notification channel

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAlertsAPI.O11yDeleteNotificationChannel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yDeleteNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yDeleteNotificationChannel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yDeleteNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yDeleteNotificationChannelRequest struct via the builder pattern


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


## O11yGetAlertRule

> O11yAlertRule O11yGetAlertRule(ctx, id).Execute()

Get alert rule

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAlertsAPI.O11yGetAlertRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yGetAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yGetAlertRule`: O11yAlertRule
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yGetAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yGetAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yAlertRule**](O11yAlertRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yListAlertRules

> []O11yAlertRule O11yListAlertRules(ctx).Severity(severity).State(state).Page(page).PageSize(pageSize).Execute()

List alert rules

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
	severity := "severity_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAlertsAPI.O11yListAlertRules(context.Background()).Severity(severity).State(state).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yListAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yListAlertRules`: []O11yAlertRule
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yListAlertRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yListAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **severity** | **string** |  | 
 **state** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**[]O11yAlertRule**](O11yAlertRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yListNotificationChannels

> []O11yNotificationChannel O11yListNotificationChannels(ctx).Execute()

List notification channels

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
	resp, r, err := apiClient.O11yAlertsAPI.O11yListNotificationChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yListNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yListNotificationChannels`: []O11yNotificationChannel
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yListNotificationChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yListNotificationChannelsRequest struct via the builder pattern


### Return type

[**[]O11yNotificationChannel**](O11yNotificationChannel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yUpdateAlertRule

> O11yAlertRule O11yUpdateAlertRule(ctx, id).O11yAlertRuleCreate(o11yAlertRuleCreate).Execute()

Update alert rule

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	o11yAlertRuleCreate := *openapiclient.NewO11yAlertRuleCreate("Name_example", "Datasource_example", "Query_example", *openapiclient.NewO11yAlertRuleCreateCondition(*openapiclient.NewO11yAlertRuleCreateConditionEvaluator("Type_example", []float32{float32(123)}))) // O11yAlertRuleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAlertsAPI.O11yUpdateAlertRule(context.Background(), id).O11yAlertRuleCreate(o11yAlertRuleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yUpdateAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yUpdateAlertRule`: O11yAlertRule
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yUpdateAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yUpdateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yAlertRuleCreate** | [**O11yAlertRuleCreate**](O11yAlertRuleCreate.md) |  | 

### Return type

[**O11yAlertRule**](O11yAlertRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yUpdateNotificationChannel

> O11yNotificationChannel O11yUpdateNotificationChannel(ctx, id).O11yNotificationChannelCreate(o11yNotificationChannelCreate).Execute()

Update notification channel

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	o11yNotificationChannelCreate := *openapiclient.NewO11yNotificationChannelCreate("Name_example", "Type_example", *openapiclient.NewO11yNotificationChannelCreateConfig()) // O11yNotificationChannelCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAlertsAPI.O11yUpdateNotificationChannel(context.Background(), id).O11yNotificationChannelCreate(o11yNotificationChannelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAlertsAPI.O11yUpdateNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yUpdateNotificationChannel`: O11yNotificationChannel
	fmt.Fprintf(os.Stdout, "Response from `O11yAlertsAPI.O11yUpdateNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yUpdateNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yNotificationChannelCreate** | [**O11yNotificationChannelCreate**](O11yNotificationChannelCreate.md) |  | 

### Return type

[**O11yNotificationChannel**](O11yNotificationChannel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

