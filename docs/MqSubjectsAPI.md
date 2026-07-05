# \MqSubjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqGetSubjectInfo**](MqSubjectsAPI.md#MqGetSubjectInfo) | **Get** /v1/mq/subjects/{subject}/info | Get subject info
[**MqListSubjects**](MqSubjectsAPI.md#MqListSubjects) | **Get** /v1/mq/subjects | List active subjects



## MqGetSubjectInfo

> MqSubjectInfo MqGetSubjectInfo(ctx, subject).Execute()

Get subject info



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
	subject := "subject_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqSubjectsAPI.MqGetSubjectInfo(context.Background(), subject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqSubjectsAPI.MqGetSubjectInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetSubjectInfo`: MqSubjectInfo
	fmt.Fprintf(os.Stdout, "Response from `MqSubjectsAPI.MqGetSubjectInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetSubjectInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MqSubjectInfo**](MqSubjectInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListSubjects

> MqListSubjects200Response MqListSubjects(ctx).Limit(limit).Offset(offset).Filter(filter).Execute()

List active subjects



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
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip. (optional) (default to 0)
	filter := "filter_example" // string | Subject filter pattern (supports wildcards). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqSubjectsAPI.MqListSubjects(context.Background()).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqSubjectsAPI.MqListSubjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListSubjects`: MqListSubjects200Response
	fmt.Fprintf(os.Stdout, "Response from `MqSubjectsAPI.MqListSubjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqListSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]
 **filter** | **string** | Subject filter pattern (supports wildcards). | 

### Return type

[**MqListSubjects200Response**](MqListSubjects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

