# \IamSystemAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddAdapter**](IamSystemAPI.md#IamApiControllerAddAdapter) | **Post** /v1/iam/adapters | Api Controller Add Adapter
[**IamApiControllerAddCert**](IamSystemAPI.md#IamApiControllerAddCert) | **Post** /v1/iam/certs | Api Controller Add Cert
[**IamApiControllerAddForm**](IamSystemAPI.md#IamApiControllerAddForm) | **Post** /v1/iam/forms | Api Controller Add Form
[**IamApiControllerAddRecord**](IamSystemAPI.md#IamApiControllerAddRecord) | **Post** /v1/iam/records | Api Controller Add Record
[**IamApiControllerAddSyncer**](IamSystemAPI.md#IamApiControllerAddSyncer) | **Post** /v1/iam/syncers | Api Controller Add Syncer
[**IamApiControllerAddTicket**](IamSystemAPI.md#IamApiControllerAddTicket) | **Post** /v1/iam/tickets | Api Controller Add Ticket
[**IamApiControllerAddTicketMessage**](IamSystemAPI.md#IamApiControllerAddTicketMessage) | **Post** /v1/iam/ticket-messages | Api Controller Add Ticket Message
[**IamApiControllerDeleteAdapter**](IamSystemAPI.md#IamApiControllerDeleteAdapter) | **Delete** /v1/iam/adapters/{id} | Api Controller Delete Adapter
[**IamApiControllerDeleteCert**](IamSystemAPI.md#IamApiControllerDeleteCert) | **Delete** /v1/iam/certs/{id} | Api Controller Delete Cert
[**IamApiControllerDeleteForm**](IamSystemAPI.md#IamApiControllerDeleteForm) | **Delete** /v1/iam/forms/{id} | Api Controller Delete Form
[**IamApiControllerDeleteSyncer**](IamSystemAPI.md#IamApiControllerDeleteSyncer) | **Delete** /v1/iam/syncers/{id} | Api Controller Delete Syncer
[**IamApiControllerDeleteTicket**](IamSystemAPI.md#IamApiControllerDeleteTicket) | **Delete** /v1/iam/tickets/{id} | Api Controller Delete Ticket
[**IamApiControllerGetAdapter**](IamSystemAPI.md#IamApiControllerGetAdapter) | **Get** /v1/iam/adapters/{id} | Api Controller Get Adapter
[**IamApiControllerGetAdapters**](IamSystemAPI.md#IamApiControllerGetAdapters) | **Get** /v1/iam/adapters | Api Controller Get Adapters
[**IamApiControllerGetCert**](IamSystemAPI.md#IamApiControllerGetCert) | **Get** /v1/iam/certs/{id} | Api Controller Get Cert
[**IamApiControllerGetCerts**](IamSystemAPI.md#IamApiControllerGetCerts) | **Get** /v1/iam/certs | Api Controller Get Certs
[**IamApiControllerGetDashboard**](IamSystemAPI.md#IamApiControllerGetDashboard) | **Get** /v1/iam/dashboard | Api Controller Get Dashboard
[**IamApiControllerGetForm**](IamSystemAPI.md#IamApiControllerGetForm) | **Get** /v1/iam/forms/{id} | Api Controller Get Form
[**IamApiControllerGetForms**](IamSystemAPI.md#IamApiControllerGetForms) | **Get** /v1/iam/forms | Api Controller Get Forms
[**IamApiControllerGetGlobalCerts**](IamSystemAPI.md#IamApiControllerGetGlobalCerts) | **Get** /v1/iam/global-certs | Api Controller Get Global Certs
[**IamApiControllerGetGlobalForms**](IamSystemAPI.md#IamApiControllerGetGlobalForms) | **Get** /v1/iam/global-forms | Api Controller Get Global Forms
[**IamApiControllerGetMetrics**](IamSystemAPI.md#IamApiControllerGetMetrics) | **Get** /v1/iam/metrics | Api Controller Get Metrics
[**IamApiControllerGetPrometheusInfo**](IamSystemAPI.md#IamApiControllerGetPrometheusInfo) | **Get** /v1/iam/metrics/prometheus | Api Controller Get Prometheus Info
[**IamApiControllerGetRecords**](IamSystemAPI.md#IamApiControllerGetRecords) | **Get** /v1/iam/records | Api Controller Get Records
[**IamApiControllerGetRecordsByFilter**](IamSystemAPI.md#IamApiControllerGetRecordsByFilter) | **Get** /v1/iam/records-filters/{id} | Api Controller Get Records By Filter
[**IamApiControllerGetSyncer**](IamSystemAPI.md#IamApiControllerGetSyncer) | **Get** /v1/iam/syncers/{id} | Api Controller Get Syncer
[**IamApiControllerGetSyncers**](IamSystemAPI.md#IamApiControllerGetSyncers) | **Get** /v1/iam/syncers | Api Controller Get Syncers
[**IamApiControllerGetSystemInfo**](IamSystemAPI.md#IamApiControllerGetSystemInfo) | **Get** /v1/iam/system | Api Controller Get System Info
[**IamApiControllerGetTicket**](IamSystemAPI.md#IamApiControllerGetTicket) | **Get** /v1/iam/tickets/{id} | Api Controller Get Ticket
[**IamApiControllerGetTickets**](IamSystemAPI.md#IamApiControllerGetTickets) | **Get** /v1/iam/tickets | Api Controller Get Tickets
[**IamApiControllerGetVersionInfo**](IamSystemAPI.md#IamApiControllerGetVersionInfo) | **Get** /v1/iam/version-infos/{id} | Api Controller Get Version Info
[**IamApiControllerGetWebhookEventType**](IamSystemAPI.md#IamApiControllerGetWebhookEventType) | **Get** /v1/iam/webhook-events/{id} | Api Controller Get Webhook Event Type
[**IamApiControllerGetWechatQRCode**](IamSystemAPI.md#IamApiControllerGetWechatQRCode) | **Get** /v1/iam/qrcodes/{id} | Api Controller Get Wechat QR Code
[**IamApiControllerHandleOfficialAccountEvent**](IamSystemAPI.md#IamApiControllerHandleOfficialAccountEvent) | **Post** /v1/iam/webhook | Api Controller Handle Official Account Event
[**IamApiControllerHealth**](IamSystemAPI.md#IamApiControllerHealth) | **Get** /v1/iam/health | Api Controller Health
[**IamApiControllerRefreshEngines**](IamSystemAPI.md#IamApiControllerRefreshEngines) | **Post** /v1/iam/refresh-engines | Api Controller Refresh Engines
[**IamApiControllerRunSyncer**](IamSystemAPI.md#IamApiControllerRunSyncer) | **Get** /v1/iam/syncers/run | Api Controller Run Syncer
[**IamApiControllerSendEmail**](IamSystemAPI.md#IamApiControllerSendEmail) | **Post** /v1/iam/messaging/email | Api Controller Send Email
[**IamApiControllerSendNotification**](IamSystemAPI.md#IamApiControllerSendNotification) | **Post** /v1/iam/messaging/notification | Api Controller Send Notification
[**IamApiControllerSendSms**](IamSystemAPI.md#IamApiControllerSendSms) | **Post** /v1/iam/messaging/sms | Api Controller Send Sms
[**IamApiControllerUpdateAdapter**](IamSystemAPI.md#IamApiControllerUpdateAdapter) | **Put** /v1/iam/adapters/{id} | Api Controller Update Adapter
[**IamApiControllerUpdateCert**](IamSystemAPI.md#IamApiControllerUpdateCert) | **Put** /v1/iam/certs/{id} | Api Controller Update Cert
[**IamApiControllerUpdateForm**](IamSystemAPI.md#IamApiControllerUpdateForm) | **Put** /v1/iam/forms/{id} | Api Controller Update Form
[**IamApiControllerUpdateSyncer**](IamSystemAPI.md#IamApiControllerUpdateSyncer) | **Put** /v1/iam/syncers/{id} | Api Controller Update Syncer
[**IamApiControllerUpdateTicket**](IamSystemAPI.md#IamApiControllerUpdateTicket) | **Put** /v1/iam/tickets/{id} | Api Controller Update Ticket



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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerAddAdapter(context.Background()).IamObjectAdapter(iamObjectAdapter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerAddAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddAdapter`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerAddAdapter`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerAddCert(context.Background()).IamObjectCert(iamObjectCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerAddCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddCert`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerAddCert`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerAddForm(context.Background()).IamObjectForm(iamObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerAddForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddForm`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerAddForm`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerAddRecord(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerAddRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddRecord`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerAddRecord`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerAddSyncer(context.Background()).IamObjectSyncer(iamObjectSyncer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerAddSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddSyncer`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerAddSyncer`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerAddTicket(context.Background()).IamObjectTicket(iamObjectTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerAddTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddTicket`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerAddTicket`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerAddTicketMessage(context.Background()).Id(id).IamObjectTicketMessage(iamObjectTicketMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerAddTicketMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddTicketMessage`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerAddTicketMessage`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerDeleteAdapter(context.Background(), id).IamObjectAdapter(iamObjectAdapter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerDeleteAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteAdapter`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerDeleteAdapter`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerDeleteCert(context.Background(), id).IamObjectCert(iamObjectCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerDeleteCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteCert`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerDeleteCert`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerDeleteForm(context.Background(), id).IamObjectForm(iamObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerDeleteForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteForm`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerDeleteForm`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerDeleteSyncer(context.Background(), id).IamObjectSyncer(iamObjectSyncer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerDeleteSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteSyncer`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerDeleteSyncer`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerDeleteTicket(context.Background(), id).IamObjectTicket(iamObjectTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerDeleteTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteTicket`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerDeleteTicket`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetAdapter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAdapter`: IamObjectAdapter
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetAdapter`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetAdapters(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetAdapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetAdapters`: []IamObjectAdapter
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetAdapters`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetCert(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetCert`: IamObjectCert
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetCert`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetCerts(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetCerts`: []IamObjectCert
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetCerts`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetDashboard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetDashboard`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetDashboard`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetForm`: IamObjectForm
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetForm`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetForms(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetForms`: []IamObjectForm
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetForms`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetGlobalCerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetGlobalCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGlobalCerts`: []IamObjectCert
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetGlobalCerts`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetGlobalForms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetGlobalForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGlobalForms`: []IamObjectForm
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetGlobalForms`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetMetrics`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetPrometheusInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetPrometheusInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPrometheusInfo`: IamObjectPrometheusInfo
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetPrometheusInfo`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetRecords(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetRecords`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetRecords`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetRecordsByFilter(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetRecordsByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetRecordsByFilter`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetRecordsByFilter`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetSyncer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSyncer`: IamObjectSyncer
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetSyncer`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetSyncers(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetSyncers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSyncers`: []IamObjectSyncer
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetSyncers`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSystemInfo`: IamUtilSystemInfo
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetSystemInfo`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetTicket`: IamObjectTicket
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetTicket`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetTickets(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetTickets`: []IamObjectTicket
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetTickets`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetVersionInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetVersionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetVersionInfo`: IamUtilVersionInfo
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetVersionInfo`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetWebhookEventType(context.Background(), id).Ticket(ticket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetWebhookEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWebhookEventType`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetWebhookEventType`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerGetWechatQRCode(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerGetWechatQRCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWechatQRCode`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerGetWechatQRCode`: %v\n", resp)
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

> IamControllersResponse IamApiControllerHandleOfficialAccountEvent(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerHandleOfficialAccountEvent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerHandleOfficialAccountEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerHandleOfficialAccountEvent`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerHandleOfficialAccountEvent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerHandleOfficialAccountEventRequest struct via the builder pattern


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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerHealth`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerHealth`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerRefreshEngines(context.Background()).M(m).T(t).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerRefreshEngines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerRefreshEngines`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerRefreshEngines`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerRunSyncer(context.Background()).IamObjectSyncer(iamObjectSyncer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerRunSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerRunSyncer`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerRunSyncer`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerSendEmail(context.Background()).ClientId(clientId).ClientSecret(clientSecret).IamControllersEmailForm(iamControllersEmailForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerSendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSendEmail`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerSendEmail`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerSendNotification(context.Background()).IamControllersNotificationForm(iamControllersNotificationForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerSendNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSendNotification`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerSendNotification`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerSendSms(context.Background()).ClientId(clientId).ClientSecret(clientSecret).IamControllersSmsForm(iamControllersSmsForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerSendSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSendSms`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerSendSms`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerUpdateAdapter(context.Background(), id).IamObjectAdapter(iamObjectAdapter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerUpdateAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateAdapter`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerUpdateAdapter`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerUpdateCert(context.Background(), id).IamObjectCert(iamObjectCert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerUpdateCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateCert`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerUpdateCert`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerUpdateForm(context.Background(), id).IamObjectForm(iamObjectForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerUpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateForm`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerUpdateForm`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerUpdateSyncer(context.Background(), id).IamObjectSyncer(iamObjectSyncer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerUpdateSyncer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateSyncer`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerUpdateSyncer`: %v\n", resp)
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
	resp, r, err := apiClient.IamSystemAPI.IamApiControllerUpdateTicket(context.Background(), id).IamObjectTicket(iamObjectTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSystemAPI.IamApiControllerUpdateTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateTicket`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSystemAPI.IamApiControllerUpdateTicket`: %v\n", resp)
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

