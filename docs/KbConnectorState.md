# KbConnectorState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**Configured** | Pointer to **bool** | Deployment has OAuth credentials for it | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DocCount** | Pointer to **int32** | Ingested kb-source documents from this provider | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewKbConnectorState

`func NewKbConnectorState() *KbConnectorState`

NewKbConnectorState instantiates a new KbConnectorState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKbConnectorStateWithDefaults

`func NewKbConnectorStateWithDefaults() *KbConnectorState`

NewKbConnectorStateWithDefaults instantiates a new KbConnectorState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *KbConnectorState) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KbConnectorState) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KbConnectorState) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KbConnectorState) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetConfigured

`func (o *KbConnectorState) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *KbConnectorState) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *KbConnectorState) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *KbConnectorState) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetStatus

`func (o *KbConnectorState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KbConnectorState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KbConnectorState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KbConnectorState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDocCount

`func (o *KbConnectorState) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *KbConnectorState) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *KbConnectorState) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *KbConnectorState) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetKind

`func (o *KbConnectorState) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KbConnectorState) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KbConnectorState) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KbConnectorState) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAccount

`func (o *KbConnectorState) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *KbConnectorState) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *KbConnectorState) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *KbConnectorState) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLastSync

`func (o *KbConnectorState) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *KbConnectorState) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *KbConnectorState) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *KbConnectorState) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetError

`func (o *KbConnectorState) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *KbConnectorState) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *KbConnectorState) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *KbConnectorState) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


