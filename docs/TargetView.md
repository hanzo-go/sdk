# TargetView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**Metrics**](Metrics.md) |  | [optional] 
**MetricsAt** | Pointer to **string** |  | [optional] 
**Running** | Pointer to **int32** |  | [optional] 
**Sessions** | Pointer to **int32** |  | [optional] 
**Spec** | Pointer to [**Spec**](Spec.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewTargetView

`func NewTargetView() *TargetView`

NewTargetView instantiates a new TargetView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetViewWithDefaults

`func NewTargetViewWithDefaults() *TargetView`

NewTargetViewWithDefaults instantiates a new TargetView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *TargetView) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *TargetView) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *TargetView) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *TargetView) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TargetView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TargetView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TargetView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TargetView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *TargetView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *TargetView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TargetView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TargetView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TargetView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TargetView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TargetView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *TargetView) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TargetView) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TargetView) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TargetView) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *TargetView) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TargetView) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TargetView) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *TargetView) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsAt

`func (o *TargetView) GetMetricsAt() string`

GetMetricsAt returns the MetricsAt field if non-nil, zero value otherwise.

### GetMetricsAtOk

`func (o *TargetView) GetMetricsAtOk() (*string, bool)`

GetMetricsAtOk returns a tuple with the MetricsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsAt

`func (o *TargetView) SetMetricsAt(v string)`

SetMetricsAt sets MetricsAt field to given value.

### HasMetricsAt

`func (o *TargetView) HasMetricsAt() bool`

HasMetricsAt returns a boolean if a field has been set.

### GetRunning

`func (o *TargetView) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *TargetView) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *TargetView) SetRunning(v int32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *TargetView) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSessions

`func (o *TargetView) GetSessions() int32`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *TargetView) GetSessionsOk() (*int32, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *TargetView) SetSessions(v int32)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *TargetView) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSpec

`func (o *TargetView) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TargetView) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TargetView) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *TargetView) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *TargetView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TargetView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TargetView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TargetView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TargetView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TargetView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TargetView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TargetView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


