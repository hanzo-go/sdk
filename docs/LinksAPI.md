# \LinksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLinksById**](LinksAPI.md#DeleteLinksById) | **Delete** /v1/links/{id} | Logs out one account and stops the sessions it was running.
[**GetLinks**](LinksAPI.md#GetLinks) | **Get** /v1/links | Lists your linked accounts and the devices they sit on.
[**GetLinksById**](LinksAPI.md#GetLinksById) | **Get** /v1/links/{id} | Reads one linked account.
[**GetLinksDevicesByMachine**](LinksAPI.md#GetLinksDevicesByMachine) | **Get** /v1/links/devices/{machine} | Shows one machine: its accounts, usage and live sessions.
[**GetLinksRoute**](LinksAPI.md#GetLinksRoute) | **Get** /v1/links/route | Gets the failover order across your linked accounts.
[**GetLinksUsage**](LinksAPI.md#GetLinksUsage) | **Get** /v1/links/usage | Shows one provider account&#39;s own usage dashboard.
[**GetLinksUsageAccounts**](LinksAPI.md#GetLinksUsageAccounts) | **Get** /v1/links/usage/accounts | Breaks down what the gateway routed through each of your accounts.
[**GetLinksUsageSummary**](LinksAPI.md#GetLinksUsageSummary) | **Get** /v1/links/usage/summary | Shows plan consumption and Hanzo spend side by side.
[**PostLinks**](LinksAPI.md#PostLinks) | **Post** /v1/links | Registers a signed-in AI provider account on a machine.
[**PostLinksDevicesByMachineRevoke**](LinksAPI.md#PostLinksDevicesByMachineRevoke) | **Post** /v1/links/devices/{machine}/revoke | Logs out every account on one machine and stops its sessions.
[**PostLinksUsage**](LinksAPI.md#PostLinksUsage) | **Post** /v1/links/usage | Reports usage samples from the device collector.



## DeleteLinksById

> RevokeResp DeleteLinksById(ctx, id).Execute()

Logs out one account and stops the sessions it was running.



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
	id := "id_example" // string | ID is the link to act on, from the path. It is scoped to the caller, so another user's or org's id is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.DeleteLinksById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.DeleteLinksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLinksById`: RevokeResp
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.DeleteLinksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the link to act on, from the path. It is scoped to the caller, so another user&#39;s or org&#39;s id is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RevokeResp**](RevokeResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinks

> LinkList GetLinks(ctx).Execute()

Lists your linked accounts and the devices they sit on.



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
	resp, r, err := apiClient.LinksAPI.GetLinks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.GetLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinks`: LinkList
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.GetLinks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinksRequest struct via the builder pattern


### Return type

[**LinkList**](LinkList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinksById

> LinkView GetLinksById(ctx, id).Execute()

Reads one linked account.



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
	id := "id_example" // string | ID is the link to act on, from the path. It is scoped to the caller, so another user's or org's id is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.GetLinksById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.GetLinksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinksById`: LinkView
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.GetLinksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the link to act on, from the path. It is scoped to the caller, so another user&#39;s or org&#39;s id is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinkView**](LinkView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinksDevicesByMachine

> DeviceView GetLinksDevicesByMachine(ctx, machine).Execute()

Shows one machine: its accounts, usage and live sessions.



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
	machine := "machine_example" // string | Machine is the machine to act on, from the path. It is scoped to the caller, so a machine with none of the caller's accounts is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.GetLinksDevicesByMachine(context.Background(), machine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.GetLinksDevicesByMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinksDevicesByMachine`: DeviceView
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.GetLinksDevicesByMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine is the machine to act on, from the path. It is scoped to the caller, so a machine with none of the caller&#39;s accounts is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinksDevicesByMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceView**](DeviceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinksRoute

> RoutePlan GetLinksRoute(ctx).Execute()

Gets the failover order across your linked accounts.



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
	resp, r, err := apiClient.LinksAPI.GetLinksRoute(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.GetLinksRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinksRoute`: RoutePlan
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.GetLinksRoute`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinksRouteRequest struct via the builder pattern


### Return type

[**RoutePlan**](RoutePlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinksUsage

> BoardResp GetLinksUsage(ctx).Provider(provider).Account(account).Window(window).Range_(range_).Execute()

Shows one provider account's own usage dashboard.



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
	provider := "provider_example" // string | Provider is the provider whose meter to read. Required. (optional)
	account := "account_example" // string | Account narrows to one account when a user has several with the provider. (optional)
	window := "window_example" // string | Window selects a window class: 6h, day, week or month. Empty reads all. (optional)
	range_ := "range__example" // string | Range is the period, one of 1h, 24h, 7d or 30d; empty means 24h, and an unknown label is 400, never a quiet fallback. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.GetLinksUsage(context.Background()).Provider(provider).Account(account).Window(window).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.GetLinksUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinksUsage`: BoardResp
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.GetLinksUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinksUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Provider is the provider whose meter to read. Required. | 
 **account** | **string** | Account narrows to one account when a user has several with the provider. | 
 **window** | **string** | Window selects a window class: 6h, day, week or month. Empty reads all. | 
 **range_** | **string** | Range is the period, one of 1h, 24h, 7d or 30d; empty means 24h, and an unknown label is 400, never a quiet fallback. | 

### Return type

[**BoardResp**](BoardResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinksUsageAccounts

> AccountsUsage GetLinksUsageAccounts(ctx).Execute()

Breaks down what the gateway routed through each of your accounts.



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
	resp, r, err := apiClient.LinksAPI.GetLinksUsageAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.GetLinksUsageAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinksUsageAccounts`: AccountsUsage
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.GetLinksUsageAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinksUsageAccountsRequest struct via the builder pattern


### Return type

[**AccountsUsage**](AccountsUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinksUsageSummary

> SummaryResp GetLinksUsageSummary(ctx).Range_(range_).Execute()

Shows plan consumption and Hanzo spend side by side.



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
	range_ := "range__example" // string | Range is the period, one of 1h, 24h, 7d or 30d; empty means 24h, and an unknown label is 400, never a silent substitution. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.GetLinksUsageSummary(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.GetLinksUsageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinksUsageSummary`: SummaryResp
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.GetLinksUsageSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinksUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the period, one of 1h, 24h, 7d or 30d; empty means 24h, and an unknown label is 400, never a silent substitution. | 

### Return type

[**SummaryResp**](SummaryResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLinks

> LinkView PostLinks(ctx).EnrollReq(enrollReq).Execute()

Registers a signed-in AI provider account on a machine.



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
	enrollReq := *openapiclient.NewEnrollReq() // EnrollReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.PostLinks(context.Background()).EnrollReq(enrollReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.PostLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLinks`: LinkView
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.PostLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollReq** | [**EnrollReq**](EnrollReq.md) |  | 

### Return type

[**LinkView**](LinkView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLinksDevicesByMachineRevoke

> RevokeResp PostLinksDevicesByMachineRevoke(ctx, machine).Execute()

Logs out every account on one machine and stops its sessions.



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
	machine := "machine_example" // string | Machine is the machine to act on, from the path. It is scoped to the caller, so a machine with none of the caller's accounts is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.PostLinksDevicesByMachineRevoke(context.Background(), machine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.PostLinksDevicesByMachineRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLinksDevicesByMachineRevoke`: RevokeResp
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.PostLinksDevicesByMachineRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine is the machine to act on, from the path. It is scoped to the caller, so a machine with none of the caller&#39;s accounts is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLinksDevicesByMachineRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RevokeResp**](RevokeResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLinksUsage

> IngestResp PostLinksUsage(ctx).IngestReq(ingestReq).Execute()

Reports usage samples from the device collector.



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
	ingestReq := *openapiclient.NewIngestReq() // IngestReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.PostLinksUsage(context.Background()).IngestReq(ingestReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.PostLinksUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLinksUsage`: IngestResp
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.PostLinksUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLinksUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestReq** | [**IngestReq**](IngestReq.md) |  | 

### Return type

[**IngestResp**](IngestResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

