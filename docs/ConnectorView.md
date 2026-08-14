# ConnectorView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account names the connected external account. Absent until the org connects. | [optional] 
**Configured** | Pointer to **bool** | Configured is true when this deployment holds OAuth credentials for the provider. | [optional] 
**DocCount** | Pointer to **int32** | DocCount is the live count of this provider&#39;s documents in the org&#39;s store. | [optional] 
**Error** | Pointer to **string** | Error is the last sync failure, if any. Absent until the org connects. | [optional] 
**Kind** | Pointer to **string** | Kind is \&quot;native\&quot; for a first-party Go connector, \&quot;piece\&quot; for a long-tail one. | [optional] 
**LastSync** | Pointer to **string** | LastSync is when the last pull finished. Absent until the org connects. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector&#39;s id. | [optional] 
**Status** | Pointer to **string** | Status is connected, disconnected, syncing or error. | [optional] 

## Methods

### NewConnectorView

`func NewConnectorView() *ConnectorView`

NewConnectorView instantiates a new ConnectorView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorViewWithDefaults

`func NewConnectorViewWithDefaults() *ConnectorView`

NewConnectorViewWithDefaults instantiates a new ConnectorView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ConnectorView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ConnectorView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ConnectorView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ConnectorView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetConfigured

`func (o *ConnectorView) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *ConnectorView) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *ConnectorView) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *ConnectorView) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDocCount

`func (o *ConnectorView) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *ConnectorView) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *ConnectorView) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *ConnectorView) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetError

`func (o *ConnectorView) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConnectorView) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConnectorView) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConnectorView) HasError() bool`

HasError returns a boolean if a field has been set.

### GetKind

`func (o *ConnectorView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectorView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectorView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectorView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastSync

`func (o *ConnectorView) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ConnectorView) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ConnectorView) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ConnectorView) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetProvider

`func (o *ConnectorView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConnectorView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConnectorView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ConnectorView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectorView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectorView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


