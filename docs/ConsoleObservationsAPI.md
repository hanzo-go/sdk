# \ConsoleObservationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleGetObservation**](ConsoleObservationsAPI.md#ConsoleGetObservation) | **Get** /v1/console/observations/{observationId} | Get an observation by ID
[**ConsoleListObservations**](ConsoleObservationsAPI.md#ConsoleListObservations) | **Get** /v1/console/observations | Get a list of observations



## ConsoleGetObservation

> ConsoleObservation ConsoleGetObservation(ctx, observationId).Execute()

Get an observation by ID

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
	observationId := "observationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleObservationsAPI.ConsoleGetObservation(context.Background(), observationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleObservationsAPI.ConsoleGetObservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetObservation`: ConsoleObservation
	fmt.Fprintf(os.Stdout, "Response from `ConsoleObservationsAPI.ConsoleGetObservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetObservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleObservation**](ConsoleObservation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListObservations

> ConsoleListObservations200Response ConsoleListObservations(ctx).Page(page).Limit(limit).Name(name).UserId(userId).Type_(type_).TraceId(traceId).Level(level).ParentObservationId(parentObservationId).Environment(environment).FromStartTime(fromStartTime).ToStartTime(toStartTime).Version(version).Execute()

Get a list of observations

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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 50)
	name := "name_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	traceId := "traceId_example" // string |  (optional)
	level := "level_example" // string |  (optional)
	parentObservationId := "parentObservationId_example" // string |  (optional)
	environment := []string{"Inner_example"} // []string |  (optional)
	fromStartTime := time.Now() // time.Time |  (optional)
	toStartTime := time.Now() // time.Time |  (optional)
	version := "version_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleObservationsAPI.ConsoleListObservations(context.Background()).Page(page).Limit(limit).Name(name).UserId(userId).Type_(type_).TraceId(traceId).Level(level).ParentObservationId(parentObservationId).Environment(environment).FromStartTime(fromStartTime).ToStartTime(toStartTime).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleObservationsAPI.ConsoleListObservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListObservations`: ConsoleListObservations200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleObservationsAPI.ConsoleListObservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListObservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 50]
 **name** | **string** |  | 
 **userId** | **string** |  | 
 **type_** | **string** |  | 
 **traceId** | **string** |  | 
 **level** | **string** |  | 
 **parentObservationId** | **string** |  | 
 **environment** | **[]string** |  | 
 **fromStartTime** | **time.Time** |  | 
 **toStartTime** | **time.Time** |  | 
 **version** | **string** |  | 

### Return type

[**ConsoleListObservations200Response**](ConsoleListObservations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

