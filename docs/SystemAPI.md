# \SystemAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiGetSystemInfo**](SystemAPI.md#AiGetSystemInfo) | **Get** /v1/ai/system | System
[**IamApiControllerAddAdapter**](SystemAPI.md#IamApiControllerAddAdapter) | **Post** /v1/iam/adapters | Api Controller Add Adapter
[**IamApiControllerAddCert**](SystemAPI.md#IamApiControllerAddCert) | **Post** /v1/iam/certs | Api Controller Add Cert
[**IamApiControllerAddForm**](SystemAPI.md#IamApiControllerAddForm) | **Post** /v1/iam/forms | Api Controller Add Form
[**IamApiControllerAddRecord**](SystemAPI.md#IamApiControllerAddRecord) | **Post** /v1/iam/records | Api Controller Add Record
[**IamApiControllerAddSyncer**](SystemAPI.md#IamApiControllerAddSyncer) | **Post** /v1/iam/syncers | Api Controller Add Syncer
[**IamApiControllerAddTicket**](SystemAPI.md#IamApiControllerAddTicket) | **Post** /v1/iam/tickets | Api Controller Add Ticket
[**IamApiControllerAddTicketMessage**](SystemAPI.md#IamApiControllerAddTicketMessage) | **Post** /v1/iam/ticket-messages | Api Controller Add Ticket Message
[**IamApiControllerDeleteAdapter**](SystemAPI.md#IamApiControllerDeleteAdapter) | **Delete** /v1/iam/adapters/{id} | Api Controller Delete Adapter
[**IamApiControllerDeleteCert**](SystemAPI.md#IamApiControllerDeleteCert) | **Delete** /v1/iam/certs/{id} | Api Controller Delete Cert
[**IamApiControllerDeleteForm**](SystemAPI.md#IamApiControllerDeleteForm) | **Delete** /v1/iam/forms/{id} | Api Controller Delete Form
[**IamApiControllerDeleteSyncer**](SystemAPI.md#IamApiControllerDeleteSyncer) | **Delete** /v1/iam/syncers/{id} | Api Controller Delete Syncer
[**IamApiControllerDeleteTicket**](SystemAPI.md#IamApiControllerDeleteTicket) | **Delete** /v1/iam/tickets/{id} | Api Controller Delete Ticket
[**IamApiControllerGetAdapter**](SystemAPI.md#IamApiControllerGetAdapter) | **Get** /v1/iam/adapters/{id} | Api Controller Get Adapter
[**IamApiControllerGetAdapters**](SystemAPI.md#IamApiControllerGetAdapters) | **Get** /v1/iam/adapters | Api Controller Get Adapters
[**IamApiControllerGetCert**](SystemAPI.md#IamApiControllerGetCert) | **Get** /v1/iam/certs/{id} | Api Controller Get Cert
[**IamApiControllerGetCerts**](SystemAPI.md#IamApiControllerGetCerts) | **Get** /v1/iam/certs | Api Controller Get Certs
[**IamApiControllerGetDashboard**](SystemAPI.md#IamApiControllerGetDashboard) | **Get** /v1/iam/dashboard | Api Controller Get Dashboard
[**IamApiControllerGetForm**](SystemAPI.md#IamApiControllerGetForm) | **Get** /v1/iam/forms/{id} | Api Controller Get Form
[**IamApiControllerGetForms**](SystemAPI.md#IamApiControllerGetForms) | **Get** /v1/iam/forms | Api Controller Get Forms
[**IamApiControllerGetGlobalCerts**](SystemAPI.md#IamApiControllerGetGlobalCerts) | **Get** /v1/iam/global-certs | Api Controller Get Global Certs
[**IamApiControllerGetGlobalForms**](SystemAPI.md#IamApiControllerGetGlobalForms) | **Get** /v1/iam/global-forms | Api Controller Get Global Forms
[**IamApiControllerGetMetrics**](SystemAPI.md#IamApiControllerGetMetrics) | **Get** /v1/iam/metrics | Api Controller Get Metrics
[**IamApiControllerGetPrometheusInfo**](SystemAPI.md#IamApiControllerGetPrometheusInfo) | **Get** /v1/iam/metrics/prometheus | Api Controller Get Prometheus Info
[**IamApiControllerGetRecords**](SystemAPI.md#IamApiControllerGetRecords) | **Get** /v1/iam/records | Api Controller Get Records
[**IamApiControllerGetRecordsByFilter**](SystemAPI.md#IamApiControllerGetRecordsByFilter) | **Get** /v1/iam/records-filters/{id} | Api Controller Get Records By Filter
[**IamApiControllerGetSyncer**](SystemAPI.md#IamApiControllerGetSyncer) | **Get** /v1/iam/syncers/{id} | Api Controller Get Syncer
[**IamApiControllerGetSyncers**](SystemAPI.md#IamApiControllerGetSyncers) | **Get** /v1/iam/syncers | Api Controller Get Syncers
[**IamApiControllerGetSystemInfo**](SystemAPI.md#IamApiControllerGetSystemInfo) | **Get** /v1/iam/system | Api Controller Get System Info
[**IamApiControllerGetTicket**](SystemAPI.md#IamApiControllerGetTicket) | **Get** /v1/iam/tickets/{id} | Api Controller Get Ticket
[**IamApiControllerGetTickets**](SystemAPI.md#IamApiControllerGetTickets) | **Get** /v1/iam/tickets | Api Controller Get Tickets
[**IamApiControllerGetVersionInfo**](SystemAPI.md#IamApiControllerGetVersionInfo) | **Get** /v1/iam/version-infos/{id} | Api Controller Get Version Info
[**IamApiControllerGetWebhookEventType**](SystemAPI.md#IamApiControllerGetWebhookEventType) | **Get** /v1/iam/webhook-events/{id} | Api Controller Get Webhook Event Type
[**IamApiControllerGetWechatQRCode**](SystemAPI.md#IamApiControllerGetWechatQRCode) | **Get** /v1/iam/qrcodes/{id} | Api Controller Get Wechat QR Code
[**IamApiControllerHandleOfficialAccountEvent**](SystemAPI.md#IamApiControllerHandleOfficialAccountEvent) | **Post** /v1/iam/webhook | Api Controller Handle Official Account Event
[**IamApiControllerHealth**](SystemAPI.md#IamApiControllerHealth) | **Get** /v1/iam/health | Api Controller Health
[**IamApiControllerRefreshEngines**](SystemAPI.md#IamApiControllerRefreshEngines) | **Post** /v1/iam/refresh-engines | Api Controller Refresh Engines
[**IamApiControllerRunSyncer**](SystemAPI.md#IamApiControllerRunSyncer) | **Get** /v1/iam/syncers/run | Api Controller Run Syncer
[**IamApiControllerSendEmail**](SystemAPI.md#IamApiControllerSendEmail) | **Post** /v1/iam/messaging/email | Api Controller Send Email
[**IamApiControllerSendNotification**](SystemAPI.md#IamApiControllerSendNotification) | **Post** /v1/iam/messaging/notification | Api Controller Send Notification
[**IamApiControllerSendSms**](SystemAPI.md#IamApiControllerSendSms) | **Post** /v1/iam/messaging/sms | Api Controller Send Sms
[**IamApiControllerUpdateAdapter**](SystemAPI.md#IamApiControllerUpdateAdapter) | **Put** /v1/iam/adapters/{id} | Api Controller Update Adapter
[**IamApiControllerUpdateCert**](SystemAPI.md#IamApiControllerUpdateCert) | **Put** /v1/iam/certs/{id} | Api Controller Update Cert
[**IamApiControllerUpdateForm**](SystemAPI.md#IamApiControllerUpdateForm) | **Put** /v1/iam/forms/{id} | Api Controller Update Form
[**IamApiControllerUpdateSyncer**](SystemAPI.md#IamApiControllerUpdateSyncer) | **Put** /v1/iam/syncers/{id} | Api Controller Update Syncer
[**IamApiControllerUpdateTicket**](SystemAPI.md#IamApiControllerUpdateTicket) | **Put** /v1/iam/tickets/{id} | Api Controller Update Ticket
[**TasksTasksCluster**](SystemAPI.md#TasksTasksCluster) | **Get** /v1/tasks/cluster | Cluster status (open probe)
[**TasksTasksClusterHealth**](SystemAPI.md#TasksTasksClusterHealth) | **Get** /v1/tasks/cluster/health | Cluster health (open probe)
[**TasksTasksHealth**](SystemAPI.md#TasksTasksHealth) | **Get** /v1/tasks/health | Liveness probe
[**TasksTasksSettings**](SystemAPI.md#TasksTasksSettings) | **Get** /v1/tasks/settings | Capability flags (open bootstrap)



## AiGetSystemInfo

> AiEnvelope AiGetSystemInfo(ctx).Execute()

System

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
	resp, r, err := apiClient.SystemAPI.AiGetSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.AiGetSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetSystemInfo`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.AiGetSystemInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetSystemInfoRequest struct via the builder pattern


### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddAdapter

> IamControllersResponse IamApiControllerAddAdapter(ctx).IamObjectAdapter(iamObjectAdapter).Execute()

Api Controller Add Adapter



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
	iamObjectAdapter := *openapiclient.NewIamObjectAdapter() // IamObjectAdapter | The details of the adapter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerAddAdapter(context.Background()).IamObjectAdapter(iamObjectAdapter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerAddAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddAdapter`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerAddAdapter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectAdapter** | [**IamObjectAdapter**](IamObjectAdapter.md) | The details of the adapter | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddCert

> IamControllersResponse IamApiControllerAddCert(ctx).IamObjectCert(iamObjectCert).Execute()

Api Controller Add Cert



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
	iamObjectCert := *openapiclient.NewIamObjectCert() // IamObjectCert | The details of the cert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerAddCert(context.Background()).IamObjectCert(iamObjectCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerAddCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddCert`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerAddCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectCert** | [**IamObjectCert**](IamObjectCert.md) | The details of the cert | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddForm

> IamControllersResponse IamApiControllerAddForm(ctx).IamObjectForm(iamObjectForm).Execute()

Api Controller Add Form



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
	iamObjectForm := *openapiclient.NewIamObjectForm() // IamObjectForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerAddForm(context.Background()).IamObjectForm(iamObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerAddForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddForm`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerAddForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectForm** | [**IamObjectForm**](IamObjectForm.md) | The details of the form | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddRecord

> IamControllersResponse IamApiControllerAddRecord(ctx).Body(body).Execute()

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
	body := map[string]interface{}{ ... } // map[string]interface{} | The details of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerAddRecord(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerAddRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddRecord`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerAddRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The details of the record | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddSyncer

> IamControllersResponse IamApiControllerAddSyncer(ctx).IamObjectSyncer(iamObjectSyncer).Execute()

Api Controller Add Syncer



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
	iamObjectSyncer := *openapiclient.NewIamObjectSyncer() // IamObjectSyncer | The details of the syncer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerAddSyncer(context.Background()).IamObjectSyncer(iamObjectSyncer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerAddSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddSyncer`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerAddSyncer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddSyncerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectSyncer** | [**IamObjectSyncer**](IamObjectSyncer.md) | The details of the syncer | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddTicket

> IamControllersResponse IamApiControllerAddTicket(ctx).IamObjectTicket(iamObjectTicket).Execute()

Api Controller Add Ticket



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
	iamObjectTicket := *openapiclient.NewIamObjectTicket() // IamObjectTicket | The details of the ticket

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerAddTicket(context.Background()).IamObjectTicket(iamObjectTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerAddTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddTicket`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerAddTicket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectTicket** | [**IamObjectTicket**](IamObjectTicket.md) | The details of the ticket | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddTicketMessage

> IamControllersResponse IamApiControllerAddTicketMessage(ctx).Id(id).IamObjectTicketMessage(iamObjectTicketMessage).Execute()

Api Controller Add Ticket Message



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
	id := "id_example" // string | The id ( owner/name ) of the ticket
	iamObjectTicketMessage := *openapiclient.NewIamObjectTicketMessage() // IamObjectTicketMessage | The message to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerAddTicketMessage(context.Background()).Id(id).IamObjectTicketMessage(iamObjectTicketMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerAddTicketMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddTicketMessage`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerAddTicketMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddTicketMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the ticket | 
 **iamObjectTicketMessage** | [**IamObjectTicketMessage**](IamObjectTicketMessage.md) | The message to add | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteAdapter

> IamControllersResponse IamApiControllerDeleteAdapter(ctx, id).IamObjectAdapter(iamObjectAdapter).Execute()

Api Controller Delete Adapter



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectAdapter := *openapiclient.NewIamObjectAdapter() // IamObjectAdapter | The details of the adapter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerDeleteAdapter(context.Background(), id).IamObjectAdapter(iamObjectAdapter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerDeleteAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteAdapter`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerDeleteAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectAdapter** | [**IamObjectAdapter**](IamObjectAdapter.md) | The details of the adapter | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteCert

> IamControllersResponse IamApiControllerDeleteCert(ctx, id).IamObjectCert(iamObjectCert).Execute()

Api Controller Delete Cert



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectCert := *openapiclient.NewIamObjectCert() // IamObjectCert | The details of the cert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerDeleteCert(context.Background(), id).IamObjectCert(iamObjectCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerDeleteCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteCert`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerDeleteCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectCert** | [**IamObjectCert**](IamObjectCert.md) | The details of the cert | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteForm

> IamControllersResponse IamApiControllerDeleteForm(ctx, id).IamObjectForm(iamObjectForm).Execute()

Api Controller Delete Form



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectForm := *openapiclient.NewIamObjectForm() // IamObjectForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerDeleteForm(context.Background(), id).IamObjectForm(iamObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerDeleteForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteForm`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerDeleteForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectForm** | [**IamObjectForm**](IamObjectForm.md) | The details of the form | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteSyncer

> IamControllersResponse IamApiControllerDeleteSyncer(ctx, id).IamObjectSyncer(iamObjectSyncer).Execute()

Api Controller Delete Syncer



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectSyncer := *openapiclient.NewIamObjectSyncer() // IamObjectSyncer | The details of the syncer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerDeleteSyncer(context.Background(), id).IamObjectSyncer(iamObjectSyncer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerDeleteSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteSyncer`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerDeleteSyncer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteSyncerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectSyncer** | [**IamObjectSyncer**](IamObjectSyncer.md) | The details of the syncer | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteTicket

> IamControllersResponse IamApiControllerDeleteTicket(ctx, id).IamObjectTicket(iamObjectTicket).Execute()

Api Controller Delete Ticket



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectTicket := *openapiclient.NewIamObjectTicket() // IamObjectTicket | The details of the ticket

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerDeleteTicket(context.Background(), id).IamObjectTicket(iamObjectTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerDeleteTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteTicket`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerDeleteTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectTicket** | [**IamObjectTicket**](IamObjectTicket.md) | The details of the ticket | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetAdapter

> IamObjectAdapter IamApiControllerGetAdapter(ctx, id).Execute()

Api Controller Get Adapter



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
	id := "id_example" // string | The id ( owner/name ) of the adapter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetAdapter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAdapter`: IamObjectAdapter
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the adapter | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectAdapter**](IamObjectAdapter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetAdapters

> []IamObjectAdapter IamApiControllerGetAdapters(ctx).Owner(owner).Execute()

Api Controller Get Adapters



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
	owner := "owner_example" // string | The owner of adapters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetAdapters(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetAdapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAdapters`: []IamObjectAdapter
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetAdapters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetAdaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of adapters | 

### Return type

[**[]IamObjectAdapter**](IamObjectAdapter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetCert

> IamObjectCert IamApiControllerGetCert(ctx, id).Execute()

Api Controller Get Cert



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
	id := "id_example" // string | The id ( owner/name ) of the cert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetCert(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetCert`: IamObjectCert
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the cert | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectCert**](IamObjectCert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetCerts

> []IamObjectCert IamApiControllerGetCerts(ctx).Owner(owner).Execute()

Api Controller Get Certs



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
	owner := "owner_example" // string | The owner of certs

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetCerts(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetCerts`: []IamObjectCert
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetCerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of certs | 

### Return type

[**[]IamObjectCert**](IamObjectCert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetDashboard

> IamControllersResponse IamApiControllerGetDashboard(ctx).Execute()

Api Controller Get Dashboard



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
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetDashboard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetDashboard`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetDashboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetDashboardRequest struct via the builder pattern


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetForm

> IamObjectForm IamApiControllerGetForm(ctx, id).Execute()

Api Controller Get Form



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
	id := "id_example" // string | The id (owner/name) of form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetForm`: IamObjectForm
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id (owner/name) of form | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectForm**](IamObjectForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetForms

> []IamObjectForm IamApiControllerGetForms(ctx).Owner(owner).Execute()

Api Controller Get Forms



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
	owner := "owner_example" // string | The owner of form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetForms(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetForms`: []IamObjectForm
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of form | 

### Return type

[**[]IamObjectForm**](IamObjectForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetGlobalCerts

> []IamObjectCert IamApiControllerGetGlobalCerts(ctx).Execute()

Api Controller Get Global Certs



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
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetGlobalCerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetGlobalCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGlobalCerts`: []IamObjectCert
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetGlobalCerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetGlobalCertsRequest struct via the builder pattern


### Return type

[**[]IamObjectCert**](IamObjectCert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetGlobalForms

> []IamObjectForm IamApiControllerGetGlobalForms(ctx).Execute()

Api Controller Get Global Forms



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
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetGlobalForms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetGlobalForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGlobalForms`: []IamObjectForm
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetGlobalForms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetGlobalFormsRequest struct via the builder pattern


### Return type

[**[]IamObjectForm**](IamObjectForm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetMetrics

> map[string]interface{} IamApiControllerGetMetrics(ctx).Execute()

Api Controller Get Metrics



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
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetMetricsRequest struct via the builder pattern


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


## IamApiControllerGetPrometheusInfo

> IamObjectPrometheusInfo IamApiControllerGetPrometheusInfo(ctx).Execute()

Api Controller Get Prometheus Info



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
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetPrometheusInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetPrometheusInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPrometheusInfo`: IamObjectPrometheusInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetPrometheusInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPrometheusInfoRequest struct via the builder pattern


### Return type

[**IamObjectPrometheusInfo**](IamObjectPrometheusInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetRecords

> map[string]interface{} IamApiControllerGetRecords(ctx).PageSize(pageSize).P(p).Execute()

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
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetRecords(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetRecords`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

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


## IamApiControllerGetRecordsByFilter

> map[string]interface{} IamApiControllerGetRecordsByFilter(ctx, id).Body(body).Execute()

Api Controller Get Records By Filter



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
	id := "id_example" // string | Resource identifier (owner/name)
	body := "body_example" // string | filter Record message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetRecordsByFilter(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetRecordsByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetRecordsByFilter`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetRecordsByFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetRecordsByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | filter Record message | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSyncer

> IamObjectSyncer IamApiControllerGetSyncer(ctx, id).Execute()

Api Controller Get Syncer



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
	id := "id_example" // string | The id ( owner/name ) of the syncer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetSyncer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSyncer`: IamObjectSyncer
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetSyncer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the syncer | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSyncerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectSyncer**](IamObjectSyncer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSyncers

> []IamObjectSyncer IamApiControllerGetSyncers(ctx).Owner(owner).Execute()

Api Controller Get Syncers



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
	owner := "owner_example" // string | The owner of syncers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetSyncers(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetSyncers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSyncers`: []IamObjectSyncer
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetSyncers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSyncersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of syncers | 

### Return type

[**[]IamObjectSyncer**](IamObjectSyncer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSystemInfo

> IamUtilSystemInfo IamApiControllerGetSystemInfo(ctx).Execute()

Api Controller Get System Info



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
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSystemInfo`: IamUtilSystemInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetSystemInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSystemInfoRequest struct via the builder pattern


### Return type

[**IamUtilSystemInfo**](IamUtilSystemInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetTicket

> IamObjectTicket IamApiControllerGetTicket(ctx, id).Execute()

Api Controller Get Ticket



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
	id := "id_example" // string | The id ( owner/name ) of the ticket

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetTicket`: IamObjectTicket
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the ticket | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectTicket**](IamObjectTicket.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetTickets

> []IamObjectTicket IamApiControllerGetTickets(ctx).Owner(owner).Execute()

Api Controller Get Tickets



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
	owner := "owner_example" // string | The owner of tickets

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetTickets(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetTickets`: []IamObjectTicket
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetTickets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of tickets | 

### Return type

[**[]IamObjectTicket**](IamObjectTicket.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetVersionInfo

> IamUtilVersionInfo IamApiControllerGetVersionInfo(ctx, id).Execute()

Api Controller Get Version Info



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
	id := "id_example" // string | Resource identifier (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetVersionInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetVersionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetVersionInfo`: IamUtilVersionInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetVersionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetVersionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamUtilVersionInfo**](IamUtilVersionInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetWebhookEventType

> IamControllersResponse IamApiControllerGetWebhookEventType(ctx, id).Ticket(ticket).Execute()

Api Controller Get Webhook Event Type

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
	ticket := "ticket_example" // string | The eventId of QRCode
	id := "id_example" // string | Resource identifier (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetWebhookEventType(context.Background(), id).Ticket(ticket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetWebhookEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWebhookEventType`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetWebhookEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetWebhookEventTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ticket** | **string** | The eventId of QRCode | 


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetWechatQRCode

> IamControllersResponse IamApiControllerGetWechatQRCode(ctx, id).Execute()

Api Controller Get Wechat QR Code

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
	id := "id_example" // string | The id ( owner/name ) of provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerGetWechatQRCode(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerGetWechatQRCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWechatQRCode`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerGetWechatQRCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetWechatQRCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerHandleOfficialAccountEvent

> IamControllersResponse IamApiControllerHandleOfficialAccountEvent(ctx).IamControllersWechatEvent(iamControllersWechatEvent).Execute()

Api Controller Handle Official Account Event

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
	iamControllersWechatEvent := *openapiclient.NewIamControllersWechatEvent("Event_example", "Ticket_example") // IamControllersWechatEvent | WeChat official-account event callback (XML). GET pre-verify (echostr) is a separate path; POST delivers the event XML. Query signature/timestamp/nonce are also required for signature verification.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerHandleOfficialAccountEvent(context.Background()).IamControllersWechatEvent(iamControllersWechatEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerHandleOfficialAccountEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerHandleOfficialAccountEvent`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerHandleOfficialAccountEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerHandleOfficialAccountEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamControllersWechatEvent** | [**IamControllersWechatEvent**](IamControllersWechatEvent.md) | WeChat official-account event callback (XML). GET pre-verify (echostr) is a separate path; POST delivers the event XML. Query signature/timestamp/nonce are also required for signature verification. | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerHealth

> IamControllersResponse IamApiControllerHealth(ctx).Execute()

Api Controller Health



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
	resp, r, err := apiClient.SystemAPI.IamApiControllerHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerHealth`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerHealthRequest struct via the builder pattern


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerRefreshEngines

> IamControllersResponse IamApiControllerRefreshEngines(ctx).M(m).T(t).Execute()

Api Controller Refresh Engines



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
	m := "m_example" // string | Hash for request validation
	t := "t_example" // string | Timestamp for request validation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerRefreshEngines(context.Background()).M(m).T(t).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerRefreshEngines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerRefreshEngines`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerRefreshEngines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerRefreshEnginesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **m** | **string** | Hash for request validation | 
 **t** | **string** | Timestamp for request validation | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerRunSyncer

> IamControllersResponse IamApiControllerRunSyncer(ctx).IamObjectSyncer(iamObjectSyncer).Execute()

Api Controller Run Syncer



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
	iamObjectSyncer := *openapiclient.NewIamObjectSyncer() // IamObjectSyncer | The details of the syncer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerRunSyncer(context.Background()).IamObjectSyncer(iamObjectSyncer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerRunSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerRunSyncer`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerRunSyncer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerRunSyncerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectSyncer** | [**IamObjectSyncer**](IamObjectSyncer.md) | The details of the syncer | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSendEmail

> IamControllersResponse IamApiControllerSendEmail(ctx).ClientId(clientId).ClientSecret(clientSecret).IamControllersEmailForm(iamControllersEmailForm).Execute()

Api Controller Send Email



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
	clientId := "clientId_example" // string | The clientId of the application
	clientSecret := "clientSecret_example" // string | The clientSecret of the application
	iamControllersEmailForm := *openapiclient.NewIamControllersEmailForm() // IamControllersEmailForm | Details of the email request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerSendEmail(context.Background()).ClientId(clientId).ClientSecret(clientSecret).IamControllersEmailForm(iamControllersEmailForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerSendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSendEmail`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerSendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | The clientId of the application | 
 **clientSecret** | **string** | The clientSecret of the application | 
 **iamControllersEmailForm** | [**IamControllersEmailForm**](IamControllersEmailForm.md) | Details of the email request | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSendNotification

> IamControllersResponse IamApiControllerSendNotification(ctx).IamControllersNotificationForm(iamControllersNotificationForm).Execute()

Api Controller Send Notification



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
	iamControllersNotificationForm := *openapiclient.NewIamControllersNotificationForm() // IamControllersNotificationForm | Details of the notification request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerSendNotification(context.Background()).IamControllersNotificationForm(iamControllersNotificationForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerSendNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSendNotification`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerSendNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSendNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamControllersNotificationForm** | [**IamControllersNotificationForm**](IamControllersNotificationForm.md) | Details of the notification request | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSendSms

> IamControllersResponse IamApiControllerSendSms(ctx).ClientId(clientId).ClientSecret(clientSecret).IamControllersSmsForm(iamControllersSmsForm).Execute()

Api Controller Send Sms



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
	clientId := "clientId_example" // string | The clientId of the application
	clientSecret := "clientSecret_example" // string | The clientSecret of the application
	iamControllersSmsForm := *openapiclient.NewIamControllersSmsForm() // IamControllersSmsForm | Details of the sms request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerSendSms(context.Background()).ClientId(clientId).ClientSecret(clientSecret).IamControllersSmsForm(iamControllersSmsForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerSendSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSendSms`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerSendSms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSendSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | The clientId of the application | 
 **clientSecret** | **string** | The clientSecret of the application | 
 **iamControllersSmsForm** | [**IamControllersSmsForm**](IamControllersSmsForm.md) | Details of the sms request | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateAdapter

> IamControllersResponse IamApiControllerUpdateAdapter(ctx, id).IamObjectAdapter(iamObjectAdapter).Execute()

Api Controller Update Adapter



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
	id := "id_example" // string | The id ( owner/name ) of the adapter
	iamObjectAdapter := *openapiclient.NewIamObjectAdapter() // IamObjectAdapter | The details of the adapter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerUpdateAdapter(context.Background(), id).IamObjectAdapter(iamObjectAdapter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerUpdateAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateAdapter`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerUpdateAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the adapter | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectAdapter** | [**IamObjectAdapter**](IamObjectAdapter.md) | The details of the adapter | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateCert

> IamControllersResponse IamApiControllerUpdateCert(ctx, id).IamObjectCert(iamObjectCert).Execute()

Api Controller Update Cert



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
	id := "id_example" // string | The id ( owner/name ) of the cert
	iamObjectCert := *openapiclient.NewIamObjectCert() // IamObjectCert | The details of the cert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerUpdateCert(context.Background(), id).IamObjectCert(iamObjectCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerUpdateCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateCert`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerUpdateCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the cert | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectCert** | [**IamObjectCert**](IamObjectCert.md) | The details of the cert | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateForm

> IamControllersResponse IamApiControllerUpdateForm(ctx, id).IamObjectForm(iamObjectForm).Execute()

Api Controller Update Form



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
	id := "id_example" // string | The id (owner/name) of the form
	iamObjectForm := *openapiclient.NewIamObjectForm() // IamObjectForm | The details of the form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerUpdateForm(context.Background(), id).IamObjectForm(iamObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerUpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateForm`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerUpdateForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id (owner/name) of the form | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectForm** | [**IamObjectForm**](IamObjectForm.md) | The details of the form | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateSyncer

> IamControllersResponse IamApiControllerUpdateSyncer(ctx, id).IamObjectSyncer(iamObjectSyncer).Execute()

Api Controller Update Syncer



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
	id := "id_example" // string | The id ( owner/name ) of the syncer
	iamObjectSyncer := *openapiclient.NewIamObjectSyncer() // IamObjectSyncer | The details of the syncer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerUpdateSyncer(context.Background(), id).IamObjectSyncer(iamObjectSyncer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerUpdateSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateSyncer`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerUpdateSyncer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the syncer | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateSyncerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectSyncer** | [**IamObjectSyncer**](IamObjectSyncer.md) | The details of the syncer | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateTicket

> IamControllersResponse IamApiControllerUpdateTicket(ctx, id).IamObjectTicket(iamObjectTicket).Execute()

Api Controller Update Ticket



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
	id := "id_example" // string | The id ( owner/name ) of the ticket
	iamObjectTicket := *openapiclient.NewIamObjectTicket() // IamObjectTicket | The details of the ticket

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.IamApiControllerUpdateTicket(context.Background(), id).IamObjectTicket(iamObjectTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.IamApiControllerUpdateTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateTicket`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.IamApiControllerUpdateTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the ticket | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectTicket** | [**IamObjectTicket**](IamObjectTicket.md) | The details of the ticket | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksCluster

> map[string]interface{} TasksTasksCluster(ctx).Execute()

Cluster status (open probe)

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
	resp, r, err := apiClient.SystemAPI.TasksTasksCluster(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.TasksTasksCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksCluster`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.TasksTasksCluster`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksClusterRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksClusterHealth

> map[string]interface{} TasksTasksClusterHealth(ctx).Execute()

Cluster health (open probe)

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
	resp, r, err := apiClient.SystemAPI.TasksTasksClusterHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.TasksTasksClusterHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksClusterHealth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.TasksTasksClusterHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksClusterHealthRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksHealth

> EvalsGetV1EvalsHealth200Response TasksTasksHealth(ctx).Execute()

Liveness probe

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
	resp, r, err := apiClient.SystemAPI.TasksTasksHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.TasksTasksHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksHealth`: EvalsGetV1EvalsHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.TasksTasksHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksHealthRequest struct via the builder pattern


### Return type

[**EvalsGetV1EvalsHealth200Response**](EvalsGetV1EvalsHealth200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksSettings

> map[string]interface{} TasksTasksSettings(ctx).Execute()

Capability flags (open bootstrap)

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
	resp, r, err := apiClient.SystemAPI.TasksTasksSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.TasksTasksSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.TasksTasksSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksSettingsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

