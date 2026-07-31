# CloudTargetView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**CloudMetrics**](CloudMetrics.md) |  | [optional] 
**MetricsAt** | Pointer to **string** |  | [optional] 
**Running** | Pointer to **int32** |  | [optional] 
**Sessions** | Pointer to **int32** |  | [optional] 
**Spec** | Pointer to [**CloudSpec**](CloudSpec.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudTargetView

`func NewCloudTargetView() *CloudTargetView`

NewCloudTargetView instantiates a new CloudTargetView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTargetViewWithDefaults

`func NewCloudTargetViewWithDefaults() *CloudTargetView`

NewCloudTargetViewWithDefaults instantiates a new CloudTargetView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *CloudTargetView) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CloudTargetView) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CloudTargetView) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CloudTargetView) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudTargetView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudTargetView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudTargetView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudTargetView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *CloudTargetView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudTargetView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudTargetView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudTargetView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *CloudTargetView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudTargetView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudTargetView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudTargetView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudTargetView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudTargetView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudTargetView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudTargetView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *CloudTargetView) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudTargetView) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudTargetView) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudTargetView) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *CloudTargetView) GetMetrics() CloudMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CloudTargetView) GetMetricsOk() (*CloudMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CloudTargetView) SetMetrics(v CloudMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CloudTargetView) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsAt

`func (o *CloudTargetView) GetMetricsAt() string`

GetMetricsAt returns the MetricsAt field if non-nil, zero value otherwise.

### GetMetricsAtOk

`func (o *CloudTargetView) GetMetricsAtOk() (*string, bool)`

GetMetricsAtOk returns a tuple with the MetricsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsAt

`func (o *CloudTargetView) SetMetricsAt(v string)`

SetMetricsAt sets MetricsAt field to given value.

### HasMetricsAt

`func (o *CloudTargetView) HasMetricsAt() bool`

HasMetricsAt returns a boolean if a field has been set.

### GetRunning

`func (o *CloudTargetView) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *CloudTargetView) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *CloudTargetView) SetRunning(v int32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *CloudTargetView) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSessions

`func (o *CloudTargetView) GetSessions() int32`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *CloudTargetView) GetSessionsOk() (*int32, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *CloudTargetView) SetSessions(v int32)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *CloudTargetView) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSpec

`func (o *CloudTargetView) GetSpec() CloudSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudTargetView) GetSpecOk() (*CloudSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudTargetView) SetSpec(v CloudSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudTargetView) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CloudTargetView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudTargetView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudTargetView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudTargetView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudTargetView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudTargetView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudTargetView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudTargetView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


