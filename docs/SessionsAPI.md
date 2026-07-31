# \SessionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiAddSession**](SessionsAPI.md#AiAddSession) | **Post** /v1/ai/sessions | Create a session
[**AiDeleteSession**](SessionsAPI.md#AiDeleteSession) | **Delete** /v1/ai/sessions/{owner}/{name} | Delete a session
[**AiGetSession**](SessionsAPI.md#AiGetSession) | **Get** /v1/ai/sessions/{owner}/{name} | Retrieve a session
[**AiGetSessions**](SessionsAPI.md#AiGetSessions) | **Get** /v1/ai/sessions | List sessions
[**AiIsSessionDuplicated**](SessionsAPI.md#AiIsSessionDuplicated) | **Get** /v1/ai/sessions/duplicated | Duplicated (session)
[**AiReplaceSession**](SessionsAPI.md#AiReplaceSession) | **Put** /v1/ai/sessions/{owner}/{name} | Replace a session
[**AiUpdateSession**](SessionsAPI.md#AiUpdateSession) | **Patch** /v1/ai/sessions/{owner}/{name} | Update a session
[**AnalyticsGetSession**](SessionsAPI.md#AnalyticsGetSession) | **Get** /v1/analytics/websites/{websiteId}/sessions/{sessionId} | Get a single session by ID
[**AnalyticsGetSessionActivity**](SessionsAPI.md#AnalyticsGetSessionActivity) | **Get** /v1/analytics/websites/{websiteId}/sessions/{sessionId}/activity | Get activity log for a session
[**AnalyticsGetSessionDataProperties**](SessionsAPI.md#AnalyticsGetSessionDataProperties) | **Get** /v1/analytics/websites/{websiteId}/session-data/properties | Get distinct session data property names
[**AnalyticsGetSessionDataValues**](SessionsAPI.md#AnalyticsGetSessionDataValues) | **Get** /v1/analytics/websites/{websiteId}/session-data/values | Get session data values for a property
[**AnalyticsGetSessionProperties**](SessionsAPI.md#AnalyticsGetSessionProperties) | **Get** /v1/analytics/websites/{websiteId}/sessions/{sessionId}/properties | Get custom session data properties
[**AnalyticsGetSessionStats**](SessionsAPI.md#AnalyticsGetSessionStats) | **Get** /v1/analytics/websites/{websiteId}/sessions/stats | Get aggregate session statistics
[**AnalyticsGetSessions**](SessionsAPI.md#AnalyticsGetSessions) | **Get** /v1/analytics/websites/{websiteId}/sessions | Get paginated list of sessions
[**AnalyticsGetSessionsWeekly**](SessionsAPI.md#AnalyticsGetSessionsWeekly) | **Get** /v1/analytics/websites/{websiteId}/sessions/weekly | Get weekly session breakdown
[**IamApiControllerAddSession**](SessionsAPI.md#IamApiControllerAddSession) | **Post** /v1/iam/sessions | Api Controller Add Session
[**IamApiControllerDeleteSession**](SessionsAPI.md#IamApiControllerDeleteSession) | **Delete** /v1/iam/sessions/{id} | Api Controller Delete Session
[**IamApiControllerGetSessions**](SessionsAPI.md#IamApiControllerGetSessions) | **Get** /v1/iam/sessions | Api Controller Get Sessions
[**IamApiControllerGetSingleSession**](SessionsAPI.md#IamApiControllerGetSingleSession) | **Get** /v1/iam/sessions/{id} | Api Controller Get Single Session
[**IamApiControllerIsSessionDuplicated**](SessionsAPI.md#IamApiControllerIsSessionDuplicated) | **Get** /v1/iam/is-session-duplicated | Api Controller Is Session Duplicated
[**IamApiControllerUpdateSession**](SessionsAPI.md#IamApiControllerUpdateSession) | **Put** /v1/iam/sessions/{id} | Api Controller Update Session



## AiAddSession

> AiEnvelope AiAddSession(ctx).Body(body).Execute()

Create a session



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AiAddSession(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AiAddSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAddSession`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AiAddSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAddSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiDeleteSession

> AiEnvelope AiDeleteSession(ctx, owner, name).Execute()

Delete a session



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
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AiDeleteSession(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AiDeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiDeleteSession`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AiDeleteSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiDeleteSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiGetSession

> AiEnvelope AiGetSession(ctx, owner, name).Execute()

Retrieve a session



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
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AiGetSession(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AiGetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetSession`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AiGetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiGetSessions

> AiEnvelope AiGetSessions(ctx).Execute()

List sessions



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
	resp, r, err := apiClient.SessionsAPI.AiGetSessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AiGetSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetSessions`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AiGetSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetSessionsRequest struct via the builder pattern


### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiIsSessionDuplicated

> AiEnvelope AiIsSessionDuplicated(ctx).Execute()

Duplicated (session)

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
	resp, r, err := apiClient.SessionsAPI.AiIsSessionDuplicated(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AiIsSessionDuplicated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiIsSessionDuplicated`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AiIsSessionDuplicated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiIsSessionDuplicatedRequest struct via the builder pattern


### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiReplaceSession

> AiEnvelope AiReplaceSession(ctx, owner, name).Body(body).Execute()

Replace a session



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
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AiReplaceSession(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AiReplaceSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiReplaceSession`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AiReplaceSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiReplaceSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiUpdateSession

> AiEnvelope AiUpdateSession(ctx, owner, name).Body(body).Execute()

Update a session



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
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AiUpdateSession(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AiUpdateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiUpdateSession`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AiUpdateSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiUpdateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetSession

> AnalyticsSession AnalyticsGetSession(ctx, websiteId, sessionId).Execute()

Get a single session by ID

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AnalyticsGetSession(context.Background(), websiteId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AnalyticsGetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSession`: AnalyticsSession
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AnalyticsGetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AnalyticsSession**](AnalyticsSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetSessionActivity

> []AnalyticsGetSessionActivity200ResponseInner AnalyticsGetSessionActivity(ctx, websiteId, sessionId).StartAt(startAt).EndAt(endAt).Execute()

Get activity log for a session

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AnalyticsGetSessionActivity(context.Background(), websiteId, sessionId).StartAt(startAt).EndAt(endAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AnalyticsGetSessionActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSessionActivity`: []AnalyticsGetSessionActivity200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AnalyticsGetSessionActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSessionActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 

### Return type

[**[]AnalyticsGetSessionActivity200ResponseInner**](AnalyticsGetSessionActivity200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetSessionDataProperties

> []map[string]interface{} AnalyticsGetSessionDataProperties(ctx, websiteId).StartAt(startAt).EndAt(endAt).PropertyName(propertyName).Execute()

Get distinct session data property names

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds
	propertyName := "propertyName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AnalyticsGetSessionDataProperties(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).PropertyName(propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AnalyticsGetSessionDataProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSessionDataProperties`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AnalyticsGetSessionDataProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSessionDataPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **propertyName** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetSessionDataValues

> []map[string]interface{} AnalyticsGetSessionDataValues(ctx, websiteId).StartAt(startAt).EndAt(endAt).PropertyName(propertyName).Execute()

Get session data values for a property

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds
	propertyName := "propertyName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AnalyticsGetSessionDataValues(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).PropertyName(propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AnalyticsGetSessionDataValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSessionDataValues`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AnalyticsGetSessionDataValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSessionDataValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **propertyName** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetSessionProperties

> []map[string]interface{} AnalyticsGetSessionProperties(ctx, websiteId, sessionId).Execute()

Get custom session data properties

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AnalyticsGetSessionProperties(context.Background(), websiteId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AnalyticsGetSessionProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSessionProperties`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AnalyticsGetSessionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSessionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetSessionStats

> map[string]AnalyticsGetSessionStats200ResponseValue AnalyticsGetSessionStats(ctx, websiteId).StartAt(startAt).EndAt(endAt).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()

Get aggregate session statistics

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds
	url := "url_example" // string |  (optional)
	referrer := "referrer_example" // string |  (optional)
	title := "title_example" // string |  (optional)
	os := "os_example" // string |  (optional)
	browser := "browser_example" // string |  (optional)
	device := "device_example" // string |  (optional)
	country := "country_example" // string |  (optional)
	region := "region_example" // string |  (optional)
	city := "city_example" // string |  (optional)
	tag := "tag_example" // string |  (optional)
	host := "host_example" // string |  (optional)
	language := "language_example" // string |  (optional)
	event := "event_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AnalyticsGetSessionStats(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AnalyticsGetSessionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSessionStats`: map[string]AnalyticsGetSessionStats200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AnalyticsGetSessionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSessionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **url** | **string** |  | 
 **referrer** | **string** |  | 
 **title** | **string** |  | 
 **os** | **string** |  | 
 **browser** | **string** |  | 
 **device** | **string** |  | 
 **country** | **string** |  | 
 **region** | **string** |  | 
 **city** | **string** |  | 
 **tag** | **string** |  | 
 **host** | **string** |  | 
 **language** | **string** |  | 
 **event** | **string** |  | 

### Return type

[**map[string]AnalyticsGetSessionStats200ResponseValue**](AnalyticsGetSessionStats200ResponseValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetSessions

> []AnalyticsSession AnalyticsGetSessions(ctx, websiteId).StartAt(startAt).EndAt(endAt).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

Get paginated list of sessions

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AnalyticsGetSessions(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AnalyticsGetSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSessions`: []AnalyticsSession
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AnalyticsGetSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsSession**](AnalyticsSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetSessionsWeekly

> []map[string]interface{} AnalyticsGetSessionsWeekly(ctx, websiteId).StartAt(startAt).EndAt(endAt).Timezone(timezone).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

Get weekly session breakdown

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds
	timezone := "America/Los_Angeles" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.AnalyticsGetSessionsWeekly(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Timezone(timezone).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.AnalyticsGetSessionsWeekly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSessionsWeekly`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.AnalyticsGetSessionsWeekly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSessionsWeeklyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **timezone** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddSession

> IamControllersResponse IamApiControllerAddSession(ctx).IamObjectSession(iamObjectSession).Execute()

Api Controller Add Session



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
	iamObjectSession := *openapiclient.NewIamObjectSession() // IamObjectSession | The session object to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.IamApiControllerAddSession(context.Background()).IamObjectSession(iamObjectSession).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.IamApiControllerAddSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddSession`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.IamApiControllerAddSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectSession** | [**IamObjectSession**](IamObjectSession.md) | The session object to add | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteSession

> IamControllersResponse IamApiControllerDeleteSession(ctx, id).IamObjectSession(iamObjectSession).Execute()

Api Controller Delete Session



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectSession := *openapiclient.NewIamObjectSession() // IamObjectSession | The session object to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.IamApiControllerDeleteSession(context.Background(), id).IamObjectSession(iamObjectSession).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.IamApiControllerDeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteSession`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.IamApiControllerDeleteSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectSession** | [**IamObjectSession**](IamObjectSession.md) | The session object to delete | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSessions

> []string IamApiControllerGetSessions(ctx).Owner(owner).Execute()

Api Controller Get Sessions



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
	owner := "owner_example" // string | The organization name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.IamApiControllerGetSessions(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.IamApiControllerGetSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSessions`: []string
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.IamApiControllerGetSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The organization name | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSingleSession

> []string IamApiControllerGetSingleSession(ctx, id).SessionPkId(sessionPkId).Execute()

Api Controller Get Single Session



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
	sessionPkId := "sessionPkId_example" // string | The session ID in format: organization/user/application (e.g., built-in/admin/app-built-in)
	id := "id_example" // string | Resource identifier (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.IamApiControllerGetSingleSession(context.Background(), id).SessionPkId(sessionPkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.IamApiControllerGetSingleSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSingleSession`: []string
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.IamApiControllerGetSingleSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSingleSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionPkId** | **string** | The session ID in format: organization/user/application (e.g., built-in/admin/app-built-in) | 


### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerIsSessionDuplicated

> []string IamApiControllerIsSessionDuplicated(ctx).SessionPkId(sessionPkId).SessionId(sessionId).Execute()

Api Controller Is Session Duplicated



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
	sessionPkId := "sessionPkId_example" // string | The session ID in format: organization/user/application (e.g., built-in/admin/app-built-in)
	sessionId := "sessionId_example" // string | The specific session ID to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.IamApiControllerIsSessionDuplicated(context.Background()).SessionPkId(sessionPkId).SessionId(sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.IamApiControllerIsSessionDuplicated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerIsSessionDuplicated`: []string
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.IamApiControllerIsSessionDuplicated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerIsSessionDuplicatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionPkId** | **string** | The session ID in format: organization/user/application (e.g., built-in/admin/app-built-in) | 
 **sessionId** | **string** | The specific session ID to check | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateSession

> IamControllersResponse IamApiControllerUpdateSession(ctx, id).IamObjectSession(iamObjectSession).Execute()

Api Controller Update Session



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectSession := *openapiclient.NewIamObjectSession() // IamObjectSession | The session object to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.IamApiControllerUpdateSession(context.Background(), id).IamObjectSession(iamObjectSession).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.IamApiControllerUpdateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateSession`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.IamApiControllerUpdateSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectSession** | [**IamObjectSession**](IamObjectSession.md) | The session object to update | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

