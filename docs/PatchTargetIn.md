# PatchTargetIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | ID is the target to update, from the path. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**Metrics**](Metrics.md) | present &#x3D;&gt; a heartbeat; the server stamps its time | [optional] 
**Spec** | Pointer to [**Spec**](Spec.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchTargetIn

`func NewPatchTargetIn() *PatchTargetIn`

NewPatchTargetIn instantiates a new PatchTargetIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchTargetInWithDefaults

`func NewPatchTargetInWithDefaults() *PatchTargetIn`

NewPatchTargetInWithDefaults instantiates a new PatchTargetIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *PatchTargetIn) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *PatchTargetIn) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *PatchTargetIn) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *PatchTargetIn) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHost

`func (o *PatchTargetIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PatchTargetIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PatchTargetIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PatchTargetIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *PatchTargetIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchTargetIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchTargetIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchTargetIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *PatchTargetIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatchTargetIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatchTargetIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatchTargetIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *PatchTargetIn) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchTargetIn) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchTargetIn) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchTargetIn) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *PatchTargetIn) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *PatchTargetIn) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *PatchTargetIn) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *PatchTargetIn) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetSpec

`func (o *PatchTargetIn) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PatchTargetIn) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PatchTargetIn) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PatchTargetIn) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *PatchTargetIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchTargetIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchTargetIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchTargetIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


