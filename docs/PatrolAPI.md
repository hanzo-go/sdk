# \PatrolAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPatrolCamera**](PatrolAPI.md#GetPatrolCamera) | **Get** /v1/patrol/camera | Lists the cameras a caller may OPEN.
[**GetPatrolCameraByNameStream**](PatrolAPI.md#GetPatrolCameraByNameStream) | **Get** /v1/patrol/camera/{name}/stream | Mints a short-lived way in to one camera.
[**GetPatrolCameraByNameStreamByTicket**](PatrolAPI.md#GetPatrolCameraByNameStreamByTicket) | **Get** /v1/patrol/camera/{name}/stream/{ticket} | Redeems a minted ticket for what opens the camera.
[**GetPatrolEvent**](PatrolAPI.md#GetPatrolEvent) | **Get** /v1/patrol/event | Lists the feed, newest first.
[**GetPatrolFeed**](PatrolAPI.md#GetPatrolFeed) | **Get** /v1/patrol/feed | Watch the live feed
[**GetPatrolIncident**](PatrolAPI.md#GetPatrolIncident) | **Get** /v1/patrol/incident | Lists responses.
[**GetPatrolIncidentByName**](PatrolAPI.md#GetPatrolIncidentByName) | **Get** /v1/patrol/incident/{name} | Returns one response whole: the incident, every act performed on it, and the SLA clocks those acts settled.
[**GetPatrolKey**](PatrolAPI.md#GetPatrolKey) | **Get** /v1/patrol/key | Lists the key register.
[**GetPatrolKeyByName**](PatrolAPI.md#GetPatrolKeyByName) | **Get** /v1/patrol/key/{name} | Returns one key set from the register.
[**GetPatrolNote**](PatrolAPI.md#GetPatrolNote) | **Get** /v1/patrol/note | Lists internal remarks.
[**GetPatrolReportByName**](PatrolAPI.md#GetPatrolReportByName) | **Get** /v1/patrol/report/{name} | Download a filed report
[**GetPatrolSite**](PatrolAPI.md#GetPatrolSite) | **Get** /v1/patrol/site | Lists the estate.
[**GetPatrolSiteByName**](PatrolAPI.md#GetPatrolSiteByName) | **Get** /v1/patrol/site/{name} | Returns one site whole: the register row plus the panel&#39;s zones and the standing hazards, which ride inside the document.
[**GetPatrolSnapshot**](PatrolAPI.md#GetPatrolSnapshot) | **Get** /v1/patrol/snapshot | Returns the whole picture in one read: sites, units with their position trails, today&#39;s tours and checkpoints, the key register, the recent feed and the open incident.
[**GetPatrolTenantByOrg**](PatrolAPI.md#GetPatrolTenantByOrg) | **Get** /v1/patrol/tenant/{org} | Returns one org&#39;s public face — brand, company, operations, features and terminology.
[**GetPatrolTour**](PatrolAPI.md#GetPatrolTour) | **Get** /v1/patrol/tour | Lists patrol rounds with their checkpoints — the proof-of-presence record.
[**GetPatrolUnit**](PatrolAPI.md#GetPatrolUnit) | **Get** /v1/patrol/unit | Lists the fleet with each unit&#39;s position trail, which is what the map draws.
[**GetPatrolUnitNear**](PatrolAPI.md#GetPatrolUnitNear) | **Get** /v1/patrol/unit/near | Answers which units can reach a site soonest.
[**PostPatrolAlarm**](PatrolAPI.md#PostPatrolAlarm) | **Post** /v1/patrol/alarm | Takes one activation.
[**PostPatrolEvent**](PatrolAPI.md#PostPatrolEvent) | **Post** /v1/patrol/event | Writes one feed row by hand and pushes it to every screen watching.
[**PostPatrolIncident**](PatrolAPI.md#PostPatrolIncident) | **Post** /v1/patrol/incident | Records a response by hand, in state received and carrying the site&#39;s tier and SLA.
[**PostPatrolIncidentByNameStep**](PatrolAPI.md#PostPatrolIncidentByNameStep) | **Post** /v1/patrol/incident/{name}/step | Performs one act on an incident: it stamps the act&#39;s time, moves the state, appends the act to the trail, and runs what the act sets in motion — a unit assigned, a site secured, a report filed, a controller told.
[**PostPatrolKeyByNameMove**](PatrolAPI.md#PostPatrolKeyByNameMove) | **Post** /v1/patrol/key/{name}/move | Records a key set leaving, returning, or being counted, and moves the register row to match.
[**PostPatrolNote**](PatrolAPI.md#PostPatrolNote) | **Post** /v1/patrol/note | Records an internal remark, stamped with who wrote it.
[**PostPatrolPointByNameConfirm**](PatrolAPI.md#PostPatrolPointByNameConfirm) | **Post** /v1/patrol/point/{name}/confirm | Records presence at a checkpoint: who confirmed it, the method, where they stood, and the photo key if one was taken.
[**PostPatrolReport**](PatrolAPI.md#PostPatrolReport) | **Post** /v1/patrol/report | Composes and files a report from one estate read, so the same figures read the same a month later.
[**PostPatrolUnitByNameFix**](PatrolAPI.md#PostPatrolUnitByNameFix) | **Post** /v1/patrol/unit/{name}/fix | Records where a unit is.
[**PostPatrolUnitByNameState**](PatrolAPI.md#PostPatrolUnitByNameState) | **Post** /v1/patrol/unit/{name}/state | Moves a unit between states — on shift, on a break, off, or in trouble.
[**PutPatrolSiteByName**](PatrolAPI.md#PutPatrolSiteByName) | **Put** /v1/patrol/site/{name} | Changes a site&#39;s contract and brief fields.



## GetPatrolCamera

> PatrolPatrolCameraList GetPatrolCamera(ctx).Site(site).Execute()

Lists the cameras a caller may OPEN.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	site := "site_example" // string | Site narrows the list to one site by document name. Empty lists every camera. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolCamera(context.Background()).Site(site).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolCamera``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolCamera`: PatrolPatrolCameraList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolCamera`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolCameraRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **site** | **string** | Site narrows the list to one site by document name. Empty lists every camera. | 

### Return type

[**PatrolPatrolCameraList**](PatrolPatrolCameraList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolCameraByNameStream

> PatrolPatrolStreamOut GetPatrolCameraByNameStream(ctx, name).Execute()

Mints a short-lived way in to one camera.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the camera's document name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolCameraByNameStream(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolCameraByNameStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolCameraByNameStream`: PatrolPatrolStreamOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolCameraByNameStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the camera&#39;s document name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolCameraByNameStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatrolPatrolStreamOut**](PatrolPatrolStreamOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolCameraByNameStreamByTicket

> PatrolPatrolTicketOut GetPatrolCameraByNameStreamByTicket(ctx, name, ticket).Expires(expires).Execute()

Redeems a minted ticket for what opens the camera.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the camera's document name, from the path.
	ticket := "ticket_example" // string | Ticket is the tag the minted address carries, from the path.
	expires := "expires_example" // string | Expires is the moment the tag covers, as the seconds the minted address carries in its `expires` parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolCameraByNameStreamByTicket(context.Background(), name, ticket).Expires(expires).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolCameraByNameStreamByTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolCameraByNameStreamByTicket`: PatrolPatrolTicketOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolCameraByNameStreamByTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the camera&#39;s document name, from the path. | 
**ticket** | **string** | Ticket is the tag the minted address carries, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolCameraByNameStreamByTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expires** | **string** | Expires is the moment the tag covers, as the seconds the minted address carries in its &#x60;expires&#x60; parameter. | 

### Return type

[**PatrolPatrolTicketOut**](PatrolPatrolTicketOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolEvent

> PatrolPatrolEventList GetPatrolEvent(ctx).Level(level).Limit(limit).Execute()

Lists the feed, newest first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	level := "level_example" // string | Level narrows the list to red, amber, green or cyan. Empty lists every level. (optional)
	limit := int64(789) // int64 | Limit caps how many rows are returned. Anything that is not a positive integer uses 80, and higher values are clamped to 80. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolEvent(context.Background()).Level(level).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolEvent`: PatrolPatrolEventList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **string** | Level narrows the list to red, amber, green or cyan. Empty lists every level. | 
 **limit** | **int64** | Limit caps how many rows are returned. Anything that is not a positive integer uses 80, and higher values are clamped to 80. | 

### Return type

[**PatrolPatrolEventList**](PatrolPatrolEventList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolFeed

> GetPatrolFeed(ctx).Execute()

Watch the live feed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PatrolAPI.GetPatrolFeed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolFeedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolIncident

> PatrolPatrolIncidentList GetPatrolIncident(ctx).Open(open).Limit(limit).Execute()

Lists responses.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	open := int64(789) // int64 | Open lists only the incidents that are not closed when it is 1. (optional)
	limit := int64(789) // int64 | Limit caps how many are returned. Anything that is not a positive integer uses 50, and higher values are clamped to 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolIncident(context.Background()).Open(open).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolIncident`: PatrolPatrolIncidentList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolIncident`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **open** | **int64** | Open lists only the incidents that are not closed when it is 1. | 
 **limit** | **int64** | Limit caps how many are returned. Anything that is not a positive integer uses 50, and higher values are clamped to 50. | 

### Return type

[**PatrolPatrolIncidentList**](PatrolPatrolIncidentList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolIncidentByName

> PatrolPatrolIncidentOut GetPatrolIncidentByName(ctx, name).Execute()

Returns one response whole: the incident, every act performed on it, and the SLA clocks those acts settled.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the incident's document name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolIncidentByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolIncidentByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolIncidentByName`: PatrolPatrolIncidentOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolIncidentByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the incident&#39;s document name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolIncidentByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatrolPatrolIncidentOut**](PatrolPatrolIncidentOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolKey

> PatrolPatrolKeyList GetPatrolKey(ctx).Execute()

Lists the key register.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolKey`: PatrolPatrolKeyList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolKeyRequest struct via the builder pattern


### Return type

[**PatrolPatrolKeyList**](PatrolPatrolKeyList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolKeyByName

> PatrolPatrolKeyOut GetPatrolKeyByName(ctx, name).Execute()

Returns one key set from the register.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the key set's reference, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolKeyByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolKeyByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolKeyByName`: PatrolPatrolKeyOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolKeyByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the key set&#39;s reference, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolKeyByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatrolPatrolKeyOut**](PatrolPatrolKeyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolNote

> PatrolPatrolNoteList GetPatrolNote(ctx).Subject(subject).Execute()

Lists internal remarks.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	subject := "subject_example" // string | Subject narrows the list to one subject by exact match. Empty lists every note. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolNote(context.Background()).Subject(subject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolNote`: PatrolPatrolNoteList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subject** | **string** | Subject narrows the list to one subject by exact match. Empty lists every note. | 

### Return type

[**PatrolPatrolNoteList**](PatrolPatrolNoteList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolReportByName

> GetPatrolReportByName(ctx, name).Execute()

Download a filed report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PatrolAPI.GetPatrolReportByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolReportByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolReportByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolSite

> PatrolPatrolSiteList GetPatrolSite(ctx).Limit(limit).Status(status).Execute()

Lists the estate.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	limit := int64(789) // int64 | Limit caps how many sites are returned. Anything that is not a positive integer uses 500, and higher values are clamped to 500. (optional)
	status := "status_example" // string | Status narrows the list to ok, fault, alarm or offline. Empty lists every status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolSite(context.Background()).Limit(limit).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolSite`: PatrolPatrolSiteList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps how many sites are returned. Anything that is not a positive integer uses 500, and higher values are clamped to 500. | 
 **status** | **string** | Status narrows the list to ok, fault, alarm or offline. Empty lists every status. | 

### Return type

[**PatrolPatrolSiteList**](PatrolPatrolSiteList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolSiteByName

> PatrolPatrolSiteOut GetPatrolSiteByName(ctx, name).Execute()

Returns one site whole: the register row plus the panel's zones and the standing hazards, which ride inside the document.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the site's document name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolSiteByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolSiteByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolSiteByName`: PatrolPatrolSiteOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolSiteByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the site&#39;s document name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolSiteByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatrolPatrolSiteOut**](PatrolPatrolSiteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolSnapshot

> PatrolPatrolSnapshotOut GetPatrolSnapshot(ctx).Events(events).Execute()

Returns the whole picture in one read: sites, units with their position trails, today's tours and checkpoints, the key register, the recent feed and the open incident.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	events := int64(789) // int64 | Events caps how many feed rows ride along. Anything that is not a positive integer uses 80, and higher values are clamped to 80. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolSnapshot(context.Background()).Events(events).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolSnapshot`: PatrolPatrolSnapshotOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **events** | **int64** | Events caps how many feed rows ride along. Anything that is not a positive integer uses 80, and higher values are clamped to 80. | 

### Return type

[**PatrolPatrolSnapshotOut**](PatrolPatrolSnapshotOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolTenantByOrg

> PatrolPatrolTenantOut GetPatrolTenantByOrg(ctx, org).Execute()

Returns one org's public face — brand, company, operations, features and terminology.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	org := "org_example" // string | Org is the IAM org the tenant belongs to, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolTenantByOrg(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolTenantByOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolTenantByOrg`: PatrolPatrolTenantOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolTenantByOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the IAM org the tenant belongs to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolTenantByOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatrolPatrolTenantOut**](PatrolPatrolTenantOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolTour

> PatrolPatrolTourList GetPatrolTour(ctx).Date(date).Execute()

Lists patrol rounds with their checkpoints — the proof-of-presence record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	date := "date_example" // string | Date narrows the list to one day, as YYYY-MM-DD. \"today\" means today. Empty lists every round held. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolTour(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolTour``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolTour`: PatrolPatrolTourList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolTour`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolTourRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | Date narrows the list to one day, as YYYY-MM-DD. \&quot;today\&quot; means today. Empty lists every round held. | 

### Return type

[**PatrolPatrolTourList**](PatrolPatrolTourList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolUnit

> PatrolPatrolUnitList GetPatrolUnit(ctx).Execute()

Lists the fleet with each unit's position trail, which is what the map draws.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolUnit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolUnit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolUnit`: PatrolPatrolUnitList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolUnit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolUnitRequest struct via the builder pattern


### Return type

[**PatrolPatrolUnitList**](PatrolPatrolUnitList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatrolUnitNear

> PatrolPatrolNearList GetPatrolUnitNear(ctx).Site(site).Limit(limit).Execute()

Answers which units can reach a site soonest.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	site := "site_example" // string | Site is the site to measure from, by document name. (optional)
	limit := int64(789) // int64 | Limit caps how many units come back. Anything that is not a positive integer uses 3, and higher values are clamped to 10. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.GetPatrolUnitNear(context.Background()).Site(site).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.GetPatrolUnitNear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatrolUnitNear`: PatrolPatrolNearList
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.GetPatrolUnitNear`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatrolUnitNearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **site** | **string** | Site is the site to measure from, by document name. | 
 **limit** | **int64** | Limit caps how many units come back. Anything that is not a positive integer uses 3, and higher values are clamped to 10. | 

### Return type

[**PatrolPatrolNearList**](PatrolPatrolNearList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolAlarm

> PatrolPatrolAlarmOut PostPatrolAlarm(ctx).PatrolPatrolAlarmIn(patrolPatrolAlarmIn).Execute()

Takes one activation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	patrolPatrolAlarmIn := *openapiclient.NewPatrolPatrolAlarmIn() // PatrolPatrolAlarmIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolAlarm(context.Background()).PatrolPatrolAlarmIn(patrolPatrolAlarmIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolAlarm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolAlarm`: PatrolPatrolAlarmOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolAlarm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patrolPatrolAlarmIn** | [**PatrolPatrolAlarmIn**](PatrolPatrolAlarmIn.md) |  | 

### Return type

[**PatrolPatrolAlarmOut**](PatrolPatrolAlarmOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolEvent

> PatrolPatrolEventOut PostPatrolEvent(ctx).PatrolPatrolEventIn(patrolPatrolEventIn).Execute()

Writes one feed row by hand and pushes it to every screen watching.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	patrolPatrolEventIn := *openapiclient.NewPatrolPatrolEventIn() // PatrolPatrolEventIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolEvent(context.Background()).PatrolPatrolEventIn(patrolPatrolEventIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolEvent`: PatrolPatrolEventOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patrolPatrolEventIn** | [**PatrolPatrolEventIn**](PatrolPatrolEventIn.md) |  | 

### Return type

[**PatrolPatrolEventOut**](PatrolPatrolEventOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolIncident

> PatrolPatrolIncidentOut PostPatrolIncident(ctx).PatrolPatrolIncidentIn(patrolPatrolIncidentIn).Execute()

Records a response by hand, in state received and carrying the site's tier and SLA.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	patrolPatrolIncidentIn := *openapiclient.NewPatrolPatrolIncidentIn() // PatrolPatrolIncidentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolIncident(context.Background()).PatrolPatrolIncidentIn(patrolPatrolIncidentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolIncident`: PatrolPatrolIncidentOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolIncident`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patrolPatrolIncidentIn** | [**PatrolPatrolIncidentIn**](PatrolPatrolIncidentIn.md) |  | 

### Return type

[**PatrolPatrolIncidentOut**](PatrolPatrolIncidentOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolIncidentByNameStep

> PatrolPatrolIncidentOut PostPatrolIncidentByNameStep(ctx, name).PatrolPatrolIncidentAct(patrolPatrolIncidentAct).Execute()

Performs one act on an incident: it stamps the act's time, moves the state, appends the act to the trail, and runs what the act sets in motion — a unit assigned, a site secured, a report filed, a controller told.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the incident's document name, from the path.
	patrolPatrolIncidentAct := *openapiclient.NewPatrolPatrolIncidentAct() // PatrolPatrolIncidentAct | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolIncidentByNameStep(context.Background(), name).PatrolPatrolIncidentAct(patrolPatrolIncidentAct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolIncidentByNameStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolIncidentByNameStep`: PatrolPatrolIncidentOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolIncidentByNameStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the incident&#39;s document name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolIncidentByNameStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patrolPatrolIncidentAct** | [**PatrolPatrolIncidentAct**](PatrolPatrolIncidentAct.md) |  | 

### Return type

[**PatrolPatrolIncidentOut**](PatrolPatrolIncidentOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolKeyByNameMove

> PatrolPatrolKeyOut PostPatrolKeyByNameMove(ctx, name).PatrolPatrolMoveIn(patrolPatrolMoveIn).Execute()

Records a key set leaving, returning, or being counted, and moves the register row to match.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the key set's reference, from the path.
	patrolPatrolMoveIn := *openapiclient.NewPatrolPatrolMoveIn() // PatrolPatrolMoveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolKeyByNameMove(context.Background(), name).PatrolPatrolMoveIn(patrolPatrolMoveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolKeyByNameMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolKeyByNameMove`: PatrolPatrolKeyOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolKeyByNameMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the key set&#39;s reference, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolKeyByNameMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patrolPatrolMoveIn** | [**PatrolPatrolMoveIn**](PatrolPatrolMoveIn.md) |  | 

### Return type

[**PatrolPatrolKeyOut**](PatrolPatrolKeyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolNote

> PatrolPatrolNoteOut PostPatrolNote(ctx).PatrolPatrolNoteIn(patrolPatrolNoteIn).Execute()

Records an internal remark, stamped with who wrote it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	patrolPatrolNoteIn := *openapiclient.NewPatrolPatrolNoteIn() // PatrolPatrolNoteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolNote(context.Background()).PatrolPatrolNoteIn(patrolPatrolNoteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolNote`: PatrolPatrolNoteOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patrolPatrolNoteIn** | [**PatrolPatrolNoteIn**](PatrolPatrolNoteIn.md) |  | 

### Return type

[**PatrolPatrolNoteOut**](PatrolPatrolNoteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolPointByNameConfirm

> PatrolPatrolCheckpointOut PostPatrolPointByNameConfirm(ctx, name).PatrolPatrolConfirmIn(patrolPatrolConfirmIn).Execute()

Records presence at a checkpoint: who confirmed it, the method, where they stood, and the photo key if one was taken.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the checkpoint's document name, from the path.
	patrolPatrolConfirmIn := *openapiclient.NewPatrolPatrolConfirmIn() // PatrolPatrolConfirmIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolPointByNameConfirm(context.Background(), name).PatrolPatrolConfirmIn(patrolPatrolConfirmIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolPointByNameConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolPointByNameConfirm`: PatrolPatrolCheckpointOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolPointByNameConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the checkpoint&#39;s document name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolPointByNameConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patrolPatrolConfirmIn** | [**PatrolPatrolConfirmIn**](PatrolPatrolConfirmIn.md) |  | 

### Return type

[**PatrolPatrolCheckpointOut**](PatrolPatrolCheckpointOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolReport

> PatrolPatrolReportOut PostPatrolReport(ctx).PatrolPatrolReportIn(patrolPatrolReportIn).Execute()

Composes and files a report from one estate read, so the same figures read the same a month later.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	patrolPatrolReportIn := *openapiclient.NewPatrolPatrolReportIn() // PatrolPatrolReportIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolReport(context.Background()).PatrolPatrolReportIn(patrolPatrolReportIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolReport`: PatrolPatrolReportOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patrolPatrolReportIn** | [**PatrolPatrolReportIn**](PatrolPatrolReportIn.md) |  | 

### Return type

[**PatrolPatrolReportOut**](PatrolPatrolReportOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolUnitByNameFix

> PatrolPatrolFixOut PostPatrolUnitByNameFix(ctx, name).PatrolPatrolFixIn(patrolPatrolFixIn).Execute()

Records where a unit is.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the unit's call sign, from the path.
	patrolPatrolFixIn := *openapiclient.NewPatrolPatrolFixIn() // PatrolPatrolFixIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolUnitByNameFix(context.Background(), name).PatrolPatrolFixIn(patrolPatrolFixIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolUnitByNameFix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolUnitByNameFix`: PatrolPatrolFixOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolUnitByNameFix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the unit&#39;s call sign, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolUnitByNameFixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patrolPatrolFixIn** | [**PatrolPatrolFixIn**](PatrolPatrolFixIn.md) |  | 

### Return type

[**PatrolPatrolFixOut**](PatrolPatrolFixOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPatrolUnitByNameState

> PatrolPatrolUnitOut PostPatrolUnitByNameState(ctx, name).PatrolPatrolUnitStateIn(patrolPatrolUnitStateIn).Execute()

Moves a unit between states — on shift, on a break, off, or in trouble.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the unit's call sign, from the path.
	patrolPatrolUnitStateIn := *openapiclient.NewPatrolPatrolUnitStateIn() // PatrolPatrolUnitStateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PostPatrolUnitByNameState(context.Background(), name).PatrolPatrolUnitStateIn(patrolPatrolUnitStateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PostPatrolUnitByNameState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPatrolUnitByNameState`: PatrolPatrolUnitOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PostPatrolUnitByNameState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the unit&#39;s call sign, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPatrolUnitByNameStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patrolPatrolUnitStateIn** | [**PatrolPatrolUnitStateIn**](PatrolPatrolUnitStateIn.md) |  | 

### Return type

[**PatrolPatrolUnitOut**](PatrolPatrolUnitOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPatrolSiteByName

> PatrolPatrolSiteOut PutPatrolSiteByName(ctx, name).PatrolPatrolSiteEdit(patrolPatrolSiteEdit).Execute()

Changes a site's contract and brief fields.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the site's document name, from the path.
	patrolPatrolSiteEdit := *openapiclient.NewPatrolPatrolSiteEdit() // PatrolPatrolSiteEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatrolAPI.PutPatrolSiteByName(context.Background(), name).PatrolPatrolSiteEdit(patrolPatrolSiteEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatrolAPI.PutPatrolSiteByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPatrolSiteByName`: PatrolPatrolSiteOut
	fmt.Fprintf(os.Stdout, "Response from `PatrolAPI.PutPatrolSiteByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the site&#39;s document name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPatrolSiteByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patrolPatrolSiteEdit** | [**PatrolPatrolSiteEdit**](PatrolPatrolSiteEdit.md) |  | 

### Return type

[**PatrolPatrolSiteOut**](PatrolPatrolSiteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

