# \NexusRecordAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddRecord**](NexusRecordAPIAPI.md#NexusAddRecord) | **Post** /v1/nexus/add-record | add Record
[**NexusAddRecords**](NexusRecordAPIAPI.md#NexusAddRecords) | **Post** /v1/nexus/add-records | add Records
[**NexusCommitRecord**](NexusRecordAPIAPI.md#NexusCommitRecord) | **Post** /v1/nexus/commit-record | commit Record
[**NexusCommitRecordSecond**](NexusRecordAPIAPI.md#NexusCommitRecordSecond) | **Post** /v1/nexus/commit-record-second | commit Record Second
[**NexusDeleteRecord**](NexusRecordAPIAPI.md#NexusDeleteRecord) | **Post** /v1/nexus/delete-record | delete Record
[**NexusGetRecord**](NexusRecordAPIAPI.md#NexusGetRecord) | **Get** /v1/nexus/get-record | get Record
[**NexusGetRecords**](NexusRecordAPIAPI.md#NexusGetRecords) | **Get** /v1/nexus/get-records | get Records
[**NexusQueryRecord**](NexusRecordAPIAPI.md#NexusQueryRecord) | **Get** /v1/nexus/query-record | query Record
[**NexusQueryRecordSecond**](NexusRecordAPIAPI.md#NexusQueryRecordSecond) | **Get** /v1/nexus/query-record-second | query Record Second
[**NexusUpdateRecord**](NexusRecordAPIAPI.md#NexusUpdateRecord) | **Post** /v1/nexus/update-record | update Record



## NexusAddRecord

> NexusResponse NexusAddRecord(ctx).NexusRecord(nexusRecord).Execute()

add Record



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
	nexusRecord := *openapiclient.NewNexusRecord() // NexusRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusAddRecord(context.Background()).NexusRecord(nexusRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusAddRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddRecord`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusAddRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusRecord** | [**NexusRecord**](NexusRecord.md) | The details of the record | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusAddRecords

> NexusResponse NexusAddRecords(ctx).NexusRecord(nexusRecord).Execute()

add Records



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
	nexusRecord := []openapiclient.NexusRecord{*openapiclient.NewNexusRecord()} // []NexusRecord | The details of the records

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusAddRecords(context.Background()).NexusRecord(nexusRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusAddRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddRecords`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusAddRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusRecord** | [**[]NexusRecord**](NexusRecord.md) | The details of the records | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusCommitRecord

> NexusResponse NexusCommitRecord(ctx).NexusRecord(nexusRecord).Execute()

commit Record



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
	nexusRecord := *openapiclient.NewNexusRecord() // NexusRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusCommitRecord(context.Background()).NexusRecord(nexusRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusCommitRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusCommitRecord`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusCommitRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusCommitRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusRecord** | [**NexusRecord**](NexusRecord.md) | The details of the record | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusCommitRecordSecond

> NexusResponse NexusCommitRecordSecond(ctx).NexusRecord(nexusRecord).Execute()

commit Record Second



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
	nexusRecord := *openapiclient.NewNexusRecord() // NexusRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusCommitRecordSecond(context.Background()).NexusRecord(nexusRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusCommitRecordSecond``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusCommitRecordSecond`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusCommitRecordSecond`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusCommitRecordSecondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusRecord** | [**NexusRecord**](NexusRecord.md) | The details of the record | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteRecord

> NexusResponse NexusDeleteRecord(ctx).NexusRecord(nexusRecord).Execute()

delete Record



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
	nexusRecord := *openapiclient.NewNexusRecord() // NexusRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusDeleteRecord(context.Background()).NexusRecord(nexusRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusDeleteRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteRecord`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusDeleteRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusRecord** | [**NexusRecord**](NexusRecord.md) | The details of the record | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetRecord

> NexusRecord NexusGetRecord(ctx).Id(id).Execute()

get Record



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
	id := "id_example" // string | The id (owner/name) of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusGetRecord(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusGetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetRecord`: NexusRecord
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusGetRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the record | 

### Return type

[**NexusRecord**](NexusRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetRecords

> NexusRecord NexusGetRecords(ctx).PageSize(pageSize).P(p).Execute()

get Records



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
	pageSize := "pageSize_example" // string | The size of each page
	p := "p_example" // string | The page number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusGetRecords(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusGetRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetRecords`: NexusRecord
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusGetRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The page number | 

### Return type

[**NexusRecord**](NexusRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusQueryRecord

> NexusRecord NexusQueryRecord(ctx).Id(id).Execute()

query Record



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
	id := "id_example" // string | The id (owner/name) of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusQueryRecord(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusQueryRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusQueryRecord`: NexusRecord
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusQueryRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusQueryRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the record | 

### Return type

[**NexusRecord**](NexusRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusQueryRecordSecond

> NexusRecord NexusQueryRecordSecond(ctx).Id(id).Execute()

query Record Second



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
	id := "id_example" // string | The id (owner/name) of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusQueryRecordSecond(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusQueryRecordSecond``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusQueryRecordSecond`: NexusRecord
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusQueryRecordSecond`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusQueryRecordSecondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the record | 

### Return type

[**NexusRecord**](NexusRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateRecord

> NexusResponse NexusUpdateRecord(ctx).Id(id).NexusRecord(nexusRecord).Execute()

update Record



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
	id := "id_example" // string | The id (owner/name) of the record
	nexusRecord := *openapiclient.NewNexusRecord() // NexusRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusRecordAPIAPI.NexusUpdateRecord(context.Background()).Id(id).NexusRecord(nexusRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusRecordAPIAPI.NexusUpdateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateRecord`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusRecordAPIAPI.NexusUpdateRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the record | 
 **nexusRecord** | [**NexusRecord**](NexusRecord.md) | The details of the record | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

