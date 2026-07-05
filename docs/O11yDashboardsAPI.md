# \O11yDashboardsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**O11yCreateDashboard**](O11yDashboardsAPI.md#O11yCreateDashboard) | **Post** /v1/o11y/dashboards | Create dashboard
[**O11yDeleteDashboard**](O11yDashboardsAPI.md#O11yDeleteDashboard) | **Delete** /v1/o11y/dashboards/{uid} | Delete dashboard
[**O11yGetDashboard**](O11yDashboardsAPI.md#O11yGetDashboard) | **Get** /v1/o11y/dashboards/{uid} | Get dashboard
[**O11yListDashboards**](O11yDashboardsAPI.md#O11yListDashboards) | **Get** /v1/o11y/dashboards | List dashboards
[**O11yUpdateDashboard**](O11yDashboardsAPI.md#O11yUpdateDashboard) | **Put** /v1/o11y/dashboards/{uid} | Update dashboard



## O11yCreateDashboard

> O11yDashboard O11yCreateDashboard(ctx).O11yDashboardCreate(o11yDashboardCreate).Execute()

Create dashboard

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
	o11yDashboardCreate := *openapiclient.NewO11yDashboardCreate("Title_example") // O11yDashboardCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yDashboardsAPI.O11yCreateDashboard(context.Background()).O11yDashboardCreate(o11yDashboardCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yDashboardsAPI.O11yCreateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yCreateDashboard`: O11yDashboard
	fmt.Fprintf(os.Stdout, "Response from `O11yDashboardsAPI.O11yCreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yDashboardCreate** | [**O11yDashboardCreate**](O11yDashboardCreate.md) |  | 

### Return type

[**O11yDashboard**](O11yDashboard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yDeleteDashboard

> map[string]interface{} O11yDeleteDashboard(ctx, uid).Execute()

Delete dashboard

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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yDashboardsAPI.O11yDeleteDashboard(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yDashboardsAPI.O11yDeleteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yDeleteDashboard`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `O11yDashboardsAPI.O11yDeleteDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yDeleteDashboardRequest struct via the builder pattern


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


## O11yGetDashboard

> O11yDashboard O11yGetDashboard(ctx, uid).Execute()

Get dashboard

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
	uid := "uid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yDashboardsAPI.O11yGetDashboard(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yDashboardsAPI.O11yGetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yGetDashboard`: O11yDashboard
	fmt.Fprintf(os.Stdout, "Response from `O11yDashboardsAPI.O11yGetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yDashboard**](O11yDashboard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yListDashboards

> []O11yListDashboards200ResponseInner O11yListDashboards(ctx).Tag(tag).Page(page).PageSize(pageSize).Execute()

List dashboards

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
	tag := "tag_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yDashboardsAPI.O11yListDashboards(context.Background()).Tag(tag).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yDashboardsAPI.O11yListDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yListDashboards`: []O11yListDashboards200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `O11yDashboardsAPI.O11yListDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yListDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**[]O11yListDashboards200ResponseInner**](O11yListDashboards200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yUpdateDashboard

> O11yDashboard O11yUpdateDashboard(ctx, uid).O11yDashboardCreate(o11yDashboardCreate).Execute()

Update dashboard

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
	uid := "uid_example" // string | 
	o11yDashboardCreate := *openapiclient.NewO11yDashboardCreate("Title_example") // O11yDashboardCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yDashboardsAPI.O11yUpdateDashboard(context.Background(), uid).O11yDashboardCreate(o11yDashboardCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yDashboardsAPI.O11yUpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yUpdateDashboard`: O11yDashboard
	fmt.Fprintf(os.Stdout, "Response from `O11yDashboardsAPI.O11yUpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yDashboardCreate** | [**O11yDashboardCreate**](O11yDashboardCreate.md) |  | 

### Return type

[**O11yDashboard**](O11yDashboard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

