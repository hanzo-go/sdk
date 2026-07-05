# \CloudRecordAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddRecord**](CloudRecordAPIAPI.md#CloudApiControllerAddRecord) | **Post** /v1/cloud/add-record | Api Controller Add Record
[**CloudApiControllerAddRecords**](CloudRecordAPIAPI.md#CloudApiControllerAddRecords) | **Post** /v1/cloud/add-records | Api Controller Add Records
[**CloudApiControllerCommitRecord**](CloudRecordAPIAPI.md#CloudApiControllerCommitRecord) | **Post** /v1/cloud/commit-record | Api Controller Commit Record
[**CloudApiControllerCommitRecordSecond**](CloudRecordAPIAPI.md#CloudApiControllerCommitRecordSecond) | **Post** /v1/cloud/commit-record-second | Api Controller Commit Record Second
[**CloudApiControllerDeleteRecord**](CloudRecordAPIAPI.md#CloudApiControllerDeleteRecord) | **Post** /v1/cloud/delete-record | Api Controller Delete Record
[**CloudApiControllerGetRecord**](CloudRecordAPIAPI.md#CloudApiControllerGetRecord) | **Get** /v1/cloud/get-record | Api Controller Get Record
[**CloudApiControllerGetRecords**](CloudRecordAPIAPI.md#CloudApiControllerGetRecords) | **Get** /v1/cloud/get-records | Api Controller Get Records
[**CloudApiControllerQueryRecord**](CloudRecordAPIAPI.md#CloudApiControllerQueryRecord) | **Get** /v1/cloud/query-record | Api Controller Query Record
[**CloudApiControllerQueryRecordSecond**](CloudRecordAPIAPI.md#CloudApiControllerQueryRecordSecond) | **Get** /v1/cloud/query-record-second | Api Controller Query Record Second
[**CloudApiControllerUpdateRecord**](CloudRecordAPIAPI.md#CloudApiControllerUpdateRecord) | **Post** /v1/cloud/update-record | Api Controller Update Record



## CloudApiControllerAddRecord

> CloudControllersResponse CloudApiControllerAddRecord(ctx).CloudObjectRecord(cloudObjectRecord).Execute()

Api Controller Add Record



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
	cloudObjectRecord := *openapiclient.NewCloudObjectRecord() // CloudObjectRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerAddRecord(context.Background()).CloudObjectRecord(cloudObjectRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerAddRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddRecord`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerAddRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectRecord** | [**CloudObjectRecord**](CloudObjectRecord.md) | The details of the record | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerAddRecords

> CloudControllersResponse CloudApiControllerAddRecords(ctx).CloudObjectRecord(cloudObjectRecord).Execute()

Api Controller Add Records



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
	cloudObjectRecord := []openapiclient.CloudObjectRecord{*openapiclient.NewCloudObjectRecord()} // []CloudObjectRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerAddRecords(context.Background()).CloudObjectRecord(cloudObjectRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerAddRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddRecords`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerAddRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectRecord** | [**[]CloudObjectRecord**](CloudObjectRecord.md) | The details of the record | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerCommitRecord

> CloudControllersResponse CloudApiControllerCommitRecord(ctx).CloudObjectRecord(cloudObjectRecord).Execute()

Api Controller Commit Record



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
	cloudObjectRecord := *openapiclient.NewCloudObjectRecord() // CloudObjectRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerCommitRecord(context.Background()).CloudObjectRecord(cloudObjectRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerCommitRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerCommitRecord`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerCommitRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerCommitRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectRecord** | [**CloudObjectRecord**](CloudObjectRecord.md) | The details of the record | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerCommitRecordSecond

> CloudControllersResponse CloudApiControllerCommitRecordSecond(ctx).CloudObjectRecord(cloudObjectRecord).Execute()

Api Controller Commit Record Second



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
	cloudObjectRecord := *openapiclient.NewCloudObjectRecord() // CloudObjectRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerCommitRecordSecond(context.Background()).CloudObjectRecord(cloudObjectRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerCommitRecordSecond``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerCommitRecordSecond`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerCommitRecordSecond`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerCommitRecordSecondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectRecord** | [**CloudObjectRecord**](CloudObjectRecord.md) | The details of the record | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteRecord

> CloudControllersResponse CloudApiControllerDeleteRecord(ctx).CloudObjectRecord(cloudObjectRecord).Execute()

Api Controller Delete Record



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
	cloudObjectRecord := *openapiclient.NewCloudObjectRecord() // CloudObjectRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerDeleteRecord(context.Background()).CloudObjectRecord(cloudObjectRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerDeleteRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteRecord`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerDeleteRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectRecord** | [**CloudObjectRecord**](CloudObjectRecord.md) | The details of the record | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetRecord

> CloudObjectRecord CloudApiControllerGetRecord(ctx).Id(id).Execute()

Api Controller Get Record



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
	id := "id_example" // string | The id ( owner/name ) of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerGetRecord(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerGetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetRecord`: CloudObjectRecord
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerGetRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the record | 

### Return type

[**CloudObjectRecord**](CloudObjectRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetRecords

> CloudObjectRecord CloudApiControllerGetRecords(ctx).PageSize(pageSize).P(p).Execute()

Api Controller Get Records



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
	p := "p_example" // string | The number of the page

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerGetRecords(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerGetRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetRecords`: CloudObjectRecord
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerGetRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

### Return type

[**CloudObjectRecord**](CloudObjectRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerQueryRecord

> CloudObjectRecord CloudApiControllerQueryRecord(ctx).Id(id).Execute()

Api Controller Query Record



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
	id := "id_example" // string | The id ( owner/name ) of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerQueryRecord(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerQueryRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerQueryRecord`: CloudObjectRecord
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerQueryRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerQueryRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the record | 

### Return type

[**CloudObjectRecord**](CloudObjectRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerQueryRecordSecond

> CloudObjectRecord CloudApiControllerQueryRecordSecond(ctx).Id(id).Execute()

Api Controller Query Record Second



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
	id := "id_example" // string | The id ( owner/name ) of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerQueryRecordSecond(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerQueryRecordSecond``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerQueryRecordSecond`: CloudObjectRecord
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerQueryRecordSecond`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerQueryRecordSecondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the record | 

### Return type

[**CloudObjectRecord**](CloudObjectRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateRecord

> CloudControllersResponse CloudApiControllerUpdateRecord(ctx).Id(id).CloudObjectRecord(cloudObjectRecord).Execute()

Api Controller Update Record



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
	id := "id_example" // string | The id ( owner/name ) of the record
	cloudObjectRecord := *openapiclient.NewCloudObjectRecord() // CloudObjectRecord | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRecordAPIAPI.CloudApiControllerUpdateRecord(context.Background()).Id(id).CloudObjectRecord(cloudObjectRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRecordAPIAPI.CloudApiControllerUpdateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateRecord`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRecordAPIAPI.CloudApiControllerUpdateRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the record | 
 **cloudObjectRecord** | [**CloudObjectRecord**](CloudObjectRecord.md) | The details of the record | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

