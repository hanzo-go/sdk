# PaasContainerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**AvailableReplicas** | Pointer to **int32** |  | [optional] 
**ReadyReplicas** | Pointer to **int32** |  | [optional] 

## Methods

### NewPaasContainerStatus

`func NewPaasContainerStatus() *PaasContainerStatus`

NewPaasContainerStatus instantiates a new PaasContainerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasContainerStatusWithDefaults

`func NewPaasContainerStatusWithDefaults() *PaasContainerStatus`

NewPaasContainerStatusWithDefaults instantiates a new PaasContainerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaasContainerStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaasContainerStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaasContainerStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaasContainerStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvailableReplicas

`func (o *PaasContainerStatus) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *PaasContainerStatus) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *PaasContainerStatus) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *PaasContainerStatus) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *PaasContainerStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *PaasContainerStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *PaasContainerStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *PaasContainerStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


