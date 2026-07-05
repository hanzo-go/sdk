# \ConsoleScoresAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreateScore**](ConsoleScoresAPI.md#ConsoleCreateScore) | **Post** /v1/console/scores | Create a score
[**ConsoleDeleteScore**](ConsoleScoresAPI.md#ConsoleDeleteScore) | **Delete** /v1/console/scores/{scoreId} | Delete a score
[**ConsoleGetScore**](ConsoleScoresAPI.md#ConsoleGetScore) | **Get** /v1/console/scores/{scoreId} | Get a score by ID
[**ConsoleListScores**](ConsoleScoresAPI.md#ConsoleListScores) | **Get** /v1/console/scores | Get all scores



## ConsoleCreateScore

> ConsoleCreateComment200Response ConsoleCreateScore(ctx).ConsoleCreateScoreRequest(consoleCreateScoreRequest).Execute()

Create a score

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
	consoleCreateScoreRequest := *openapiclient.NewConsoleCreateScoreRequest("Name_example", interface{}(123)) // ConsoleCreateScoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleScoresAPI.ConsoleCreateScore(context.Background()).ConsoleCreateScoreRequest(consoleCreateScoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleScoresAPI.ConsoleCreateScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateScore`: ConsoleCreateComment200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleScoresAPI.ConsoleCreateScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreateScoreRequest** | [**ConsoleCreateScoreRequest**](ConsoleCreateScoreRequest.md) |  | 

### Return type

[**ConsoleCreateComment200Response**](ConsoleCreateComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleDeleteScore

> map[string]interface{} ConsoleDeleteScore(ctx, scoreId).Execute()

Delete a score

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
	scoreId := "scoreId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleScoresAPI.ConsoleDeleteScore(context.Background(), scoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleScoresAPI.ConsoleDeleteScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteScore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConsoleScoresAPI.ConsoleDeleteScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteScoreRequest struct via the builder pattern


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


## ConsoleGetScore

> ConsoleScore ConsoleGetScore(ctx, scoreId).Execute()

Get a score by ID

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
	scoreId := "scoreId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleScoresAPI.ConsoleGetScore(context.Background(), scoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleScoresAPI.ConsoleGetScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetScore`: ConsoleScore
	fmt.Fprintf(os.Stdout, "Response from `ConsoleScoresAPI.ConsoleGetScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleScore**](ConsoleScore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListScores

> ConsoleListScores200Response ConsoleListScores(ctx).Page(page).Limit(limit).UserId(userId).Name(name).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Source(source).Operator(operator).Value(value).ScoreIds(scoreIds).ConfigId(configId).DataType(dataType).Environment(environment).Execute()

Get all scores

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
	userId := "userId_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	fromTimestamp := time.Now() // time.Time |  (optional)
	toTimestamp := time.Now() // time.Time |  (optional)
	source := "source_example" // string |  (optional)
	operator := "operator_example" // string |  (optional)
	value := float32(8.14) // float32 |  (optional)
	scoreIds := "scoreIds_example" // string |  (optional)
	configId := "configId_example" // string |  (optional)
	dataType := "dataType_example" // string |  (optional)
	environment := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleScoresAPI.ConsoleListScores(context.Background()).Page(page).Limit(limit).UserId(userId).Name(name).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Source(source).Operator(operator).Value(value).ScoreIds(scoreIds).ConfigId(configId).DataType(dataType).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleScoresAPI.ConsoleListScores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListScores`: ConsoleListScores200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleScoresAPI.ConsoleListScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 50]
 **userId** | **string** |  | 
 **name** | **string** |  | 
 **fromTimestamp** | **time.Time** |  | 
 **toTimestamp** | **time.Time** |  | 
 **source** | **string** |  | 
 **operator** | **string** |  | 
 **value** | **float32** |  | 
 **scoreIds** | **string** |  | 
 **configId** | **string** |  | 
 **dataType** | **string** |  | 
 **environment** | **[]string** |  | 

### Return type

[**ConsoleListScores200Response**](ConsoleListScores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

