# \MeetAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMeetHealth**](MeetAPI.md#GetMeetHealth) | **Get** /v1/meet/health | Health reports whether the office can mint join tokens.
[**GetMeetSession**](MeetAPI.md#GetMeetSession) | **Get** /v1/meet/session | What this caller may open a room in
[**MeetCall**](MeetAPI.md#MeetCall) | **Get** /v1/meet/call | Where a room&#39;s call happens
[**MeetRecordRead**](MeetAPI.md#MeetRecordRead) | **Get** /v1/meet/record | What is being recorded in a room, and where the file goes
[**MeetRecordStart**](MeetAPI.md#MeetRecordStart) | **Post** /v1/meet/record | Start recording a room, or return the recording already running
[**MeetRecordStop**](MeetAPI.md#MeetRecordStop) | **Delete** /v1/meet/record | Stop a room&#39;s recording
[**PostMeetGettoken**](MeetAPI.md#PostMeetGettoken) | **Post** /v1/meet/getToken | Mint a join token for one video room



## GetMeetHealth

> MeetHealth GetMeetHealth(ctx).Execute()

Health reports whether the office can mint join tokens.



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
	resp, r, err := apiClient.MeetAPI.GetMeetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetAPI.GetMeetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeetHealth`: MeetHealth
	fmt.Fprintf(os.Stdout, "Response from `MeetAPI.GetMeetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeetHealthRequest struct via the builder pattern


### Return type

[**MeetHealth**](MeetHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeetSession

> GetMeetSession(ctx).Execute()

What this caller may open a room in



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
	r, err := apiClient.MeetAPI.GetMeetSession(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetAPI.GetMeetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeetSessionRequest struct via the builder pattern


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


## MeetCall

> Venue MeetCall(ctx).Workspace(workspace).Room(room).Execute()

Where a room's call happens



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
	workspace := "workspace_example" // string | Workspace is the workspace uuid holding the room, as GET /v1/team/rooms reports it. It is the segment the caller's membership is checked against.
	room := "room_example" // string | Room is the room's own id within that workspace, as GET /v1/team/rooms reports it. It is opaque here: meet keeps no rooms and cannot say whether one exists, only whether this caller may be seated in the workspace holding it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetAPI.MeetCall(context.Background()).Workspace(workspace).Room(room).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetAPI.MeetCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeetCall`: Venue
	fmt.Fprintf(os.Stdout, "Response from `MeetAPI.MeetCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeetCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspace** | **string** | Workspace is the workspace uuid holding the room, as GET /v1/team/rooms reports it. It is the segment the caller&#39;s membership is checked against. | 
 **room** | **string** | Room is the room&#39;s own id within that workspace, as GET /v1/team/rooms reports it. It is opaque here: meet keeps no rooms and cannot say whether one exists, only whether this caller may be seated in the workspace holding it. | 

### Return type

[**Venue**](Venue.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeetRecordRead

> Recording MeetRecordRead(ctx).Room(room).Execute()

What is being recorded in a room, and where the file goes



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
	room := "room_example" // string | Room is the LiveKit room, named the way the office client names one (`<workspace>_<name>_<id>`). Its leading segment is what binds the room to a tenant, and it is the segment the caller's membership is checked against.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetAPI.MeetRecordRead(context.Background()).Room(room).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetAPI.MeetRecordRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeetRecordRead`: Recording
	fmt.Fprintf(os.Stdout, "Response from `MeetAPI.MeetRecordRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeetRecordReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **room** | **string** | Room is the LiveKit room, named the way the office client names one (&#x60;&lt;workspace&gt;_&lt;name&gt;_&lt;id&gt;&#x60;). Its leading segment is what binds the room to a tenant, and it is the segment the caller&#39;s membership is checked against. | 

### Return type

[**Recording**](Recording.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeetRecordStart

> Recording MeetRecordStart(ctx).RecordIn(recordIn).Execute()

Start recording a room, or return the recording already running



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
	recordIn := *openapiclient.NewRecordIn("Room_example") // RecordIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetAPI.MeetRecordStart(context.Background()).RecordIn(recordIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetAPI.MeetRecordStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeetRecordStart`: Recording
	fmt.Fprintf(os.Stdout, "Response from `MeetAPI.MeetRecordStart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeetRecordStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordIn** | [**RecordIn**](RecordIn.md) |  | 

### Return type

[**Recording**](Recording.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeetRecordStop

> Recording MeetRecordStop(ctx).Room(room).Execute()

Stop a room's recording



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
	room := "room_example" // string | Room is the LiveKit room, named the way the office client names one (`<workspace>_<name>_<id>`). Its leading segment is what binds the room to a tenant, and it is the segment the caller's membership is checked against.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetAPI.MeetRecordStop(context.Background()).Room(room).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetAPI.MeetRecordStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeetRecordStop`: Recording
	fmt.Fprintf(os.Stdout, "Response from `MeetAPI.MeetRecordStop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeetRecordStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **room** | **string** | Room is the LiveKit room, named the way the office client names one (&#x60;&lt;workspace&gt;_&lt;name&gt;_&lt;id&gt;&#x60;). Its leading segment is what binds the room to a tenant, and it is the segment the caller&#39;s membership is checked against. | 

### Return type

[**Recording**](Recording.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMeetGettoken

> PostMeetGettoken(ctx).Execute()

Mint a join token for one video room



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
	r, err := apiClient.MeetAPI.PostMeetGettoken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetAPI.PostMeetGettoken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostMeetGettokenRequest struct via the builder pattern


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

