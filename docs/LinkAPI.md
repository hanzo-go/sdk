# \LinkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLinkById**](LinkAPI.md#DeleteLinkById) | **Delete** /v1/link/{id} | Logs out one account and stops the sessions it was running.
[**GetLink**](LinkAPI.md#GetLink) | **Get** /v1/link | Lists your linked accounts and the devices they sit on.
[**GetLinkById**](LinkAPI.md#GetLinkById) | **Get** /v1/link/{id} | Reads one linked account.
[**GetLinkDevicesByMachine**](LinkAPI.md#GetLinkDevicesByMachine) | **Get** /v1/link/devices/{machine} | Shows one machine: its accounts, usage and live sessions.
[**GetLinkRoute**](LinkAPI.md#GetLinkRoute) | **Get** /v1/link/route | Gets the failover order across your linked accounts.
[**GetLinkUsage**](LinkAPI.md#GetLinkUsage) | **Get** /v1/link/usage | Shows one provider account&#39;s own usage dashboard.
[**GetLinkUsageAccounts**](LinkAPI.md#GetLinkUsageAccounts) | **Get** /v1/link/usage/accounts | Breaks down what the gateway routed through each of your accounts.
[**GetLinkUsageSummary**](LinkAPI.md#GetLinkUsageSummary) | **Get** /v1/link/usage/summary | Shows plan consumption and Hanzo spend side by side.
[**PostLink**](LinkAPI.md#PostLink) | **Post** /v1/link | Registers a signed-in AI provider account on a machine.
[**PostLinkDevicesByMachineRevoke**](LinkAPI.md#PostLinkDevicesByMachineRevoke) | **Post** /v1/link/devices/{machine}/revoke | Logs out every account on one machine and stops its sessions.
[**PostLinkUsage**](LinkAPI.md#PostLinkUsage) | **Post** /v1/link/usage | Reports usage samples from the device collector.



## DeleteLinkById

> RevokeResp DeleteLinkById(ctx, id).Execute()

Logs out one account and stops the sessions it was running.



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
	id := "id_example" // string | ID is the link to act on, from the path. It is scoped to the caller, so another user's or org's id is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAPI.DeleteLinkById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.DeleteLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLinkById`: RevokeResp
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.DeleteLinkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the link to act on, from the path. It is scoped to the caller, so another user&#39;s or org&#39;s id is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RevokeResp**](RevokeResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLink

> LinkList GetLink(ctx).Execute()

Lists your linked accounts and the devices they sit on.



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
	resp, r, err := apiClient.LinkAPI.GetLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.GetLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLink`: LinkList
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.GetLink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkRequest struct via the builder pattern


### Return type

[**LinkList**](LinkList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkById

> LinkView GetLinkById(ctx, id).Execute()

Reads one linked account.



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
	id := "id_example" // string | ID is the link to act on, from the path. It is scoped to the caller, so another user's or org's id is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAPI.GetLinkById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.GetLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkById`: LinkView
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.GetLinkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the link to act on, from the path. It is scoped to the caller, so another user&#39;s or org&#39;s id is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinkView**](LinkView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkDevicesByMachine

> DeviceView GetLinkDevicesByMachine(ctx, machine).Execute()

Shows one machine: its accounts, usage and live sessions.



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
	machine := "machine_example" // string | Machine is the machine to act on, from the path. It is scoped to the caller, so a machine with none of the caller's accounts is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAPI.GetLinkDevicesByMachine(context.Background(), machine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.GetLinkDevicesByMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkDevicesByMachine`: DeviceView
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.GetLinkDevicesByMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine is the machine to act on, from the path. It is scoped to the caller, so a machine with none of the caller&#39;s accounts is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkDevicesByMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceView**](DeviceView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkRoute

> RoutePlan GetLinkRoute(ctx).Execute()

Gets the failover order across your linked accounts.



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
	resp, r, err := apiClient.LinkAPI.GetLinkRoute(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.GetLinkRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkRoute`: RoutePlan
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.GetLinkRoute`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkRouteRequest struct via the builder pattern


### Return type

[**RoutePlan**](RoutePlan.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkUsage

> BoardResp GetLinkUsage(ctx).Provider(provider).Account(account).Window(window).Range_(range_).Execute()

Shows one provider account's own usage dashboard.



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
	provider := "provider_example" // string | Provider is the provider whose meter to read. Required. (optional)
	account := "account_example" // string | Account narrows to one account when a user has several with the provider. (optional)
	window := "window_example" // string | Window selects a window class: 6h, day, week or month. Empty reads all. (optional)
	range_ := "range__example" // string | Range is the period, one of 1h, 24h, 7d or 30d; empty means 24h, and an unknown label is 400, never a quiet fallback. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAPI.GetLinkUsage(context.Background()).Provider(provider).Account(account).Window(window).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.GetLinkUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkUsage`: BoardResp
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.GetLinkUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Provider is the provider whose meter to read. Required. | 
 **account** | **string** | Account narrows to one account when a user has several with the provider. | 
 **window** | **string** | Window selects a window class: 6h, day, week or month. Empty reads all. | 
 **range_** | **string** | Range is the period, one of 1h, 24h, 7d or 30d; empty means 24h, and an unknown label is 400, never a quiet fallback. | 

### Return type

[**BoardResp**](BoardResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkUsageAccounts

> AccountsUsage GetLinkUsageAccounts(ctx).Execute()

Breaks down what the gateway routed through each of your accounts.



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
	resp, r, err := apiClient.LinkAPI.GetLinkUsageAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.GetLinkUsageAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkUsageAccounts`: AccountsUsage
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.GetLinkUsageAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkUsageAccountsRequest struct via the builder pattern


### Return type

[**AccountsUsage**](AccountsUsage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkUsageSummary

> SummaryResp GetLinkUsageSummary(ctx).Range_(range_).Execute()

Shows plan consumption and Hanzo spend side by side.



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
	range_ := "range__example" // string | Range is the period, one of 1h, 24h, 7d or 30d; empty means 24h, and an unknown label is 400, never a silent substitution. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAPI.GetLinkUsageSummary(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.GetLinkUsageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkUsageSummary`: SummaryResp
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.GetLinkUsageSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the period, one of 1h, 24h, 7d or 30d; empty means 24h, and an unknown label is 400, never a silent substitution. | 

### Return type

[**SummaryResp**](SummaryResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLink

> LinkView PostLink(ctx).EnrollReq(enrollReq).Execute()

Registers a signed-in AI provider account on a machine.



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
	enrollReq := *openapiclient.NewEnrollReq() // EnrollReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAPI.PostLink(context.Background()).EnrollReq(enrollReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.PostLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLink`: LinkView
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.PostLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollReq** | [**EnrollReq**](EnrollReq.md) |  | 

### Return type

[**LinkView**](LinkView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLinkDevicesByMachineRevoke

> RevokeResp PostLinkDevicesByMachineRevoke(ctx, machine).Execute()

Logs out every account on one machine and stops its sessions.



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
	machine := "machine_example" // string | Machine is the machine to act on, from the path. It is scoped to the caller, so a machine with none of the caller's accounts is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAPI.PostLinkDevicesByMachineRevoke(context.Background(), machine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.PostLinkDevicesByMachineRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLinkDevicesByMachineRevoke`: RevokeResp
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.PostLinkDevicesByMachineRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine is the machine to act on, from the path. It is scoped to the caller, so a machine with none of the caller&#39;s accounts is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLinkDevicesByMachineRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RevokeResp**](RevokeResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLinkUsage

> IngestResp PostLinkUsage(ctx).IngestReq(ingestReq).Execute()

Reports usage samples from the device collector.



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
	ingestReq := *openapiclient.NewIngestReq() // IngestReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAPI.PostLinkUsage(context.Background()).IngestReq(ingestReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAPI.PostLinkUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLinkUsage`: IngestResp
	fmt.Fprintf(os.Stdout, "Response from `LinkAPI.PostLinkUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLinkUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestReq** | [**IngestReq**](IngestReq.md) |  | 

### Return type

[**IngestResp**](IngestResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

