# PlatformUpdateApplicationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MemoryReservation** | Pointer to **int32** |  | [optional] 
**MemoryLimit** | Pointer to **int32** |  | [optional] 
**CpuReservation** | Pointer to **int32** |  | [optional] 
**CpuLimit** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformUpdateApplicationInput

`func NewPlatformUpdateApplicationInput(applicationId string, ) *PlatformUpdateApplicationInput`

NewPlatformUpdateApplicationInput instantiates a new PlatformUpdateApplicationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformUpdateApplicationInputWithDefaults

`func NewPlatformUpdateApplicationInputWithDefaults() *PlatformUpdateApplicationInput`

NewPlatformUpdateApplicationInputWithDefaults instantiates a new PlatformUpdateApplicationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PlatformUpdateApplicationInput) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PlatformUpdateApplicationInput) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PlatformUpdateApplicationInput) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetName

`func (o *PlatformUpdateApplicationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformUpdateApplicationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformUpdateApplicationInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformUpdateApplicationInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformUpdateApplicationInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformUpdateApplicationInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformUpdateApplicationInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformUpdateApplicationInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemoryReservation

`func (o *PlatformUpdateApplicationInput) GetMemoryReservation() int32`

GetMemoryReservation returns the MemoryReservation field if non-nil, zero value otherwise.

### GetMemoryReservationOk

`func (o *PlatformUpdateApplicationInput) GetMemoryReservationOk() (*int32, bool)`

GetMemoryReservationOk returns a tuple with the MemoryReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryReservation

`func (o *PlatformUpdateApplicationInput) SetMemoryReservation(v int32)`

SetMemoryReservation sets MemoryReservation field to given value.

### HasMemoryReservation

`func (o *PlatformUpdateApplicationInput) HasMemoryReservation() bool`

HasMemoryReservation returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *PlatformUpdateApplicationInput) GetMemoryLimit() int32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *PlatformUpdateApplicationInput) GetMemoryLimitOk() (*int32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *PlatformUpdateApplicationInput) SetMemoryLimit(v int32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *PlatformUpdateApplicationInput) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetCpuReservation

`func (o *PlatformUpdateApplicationInput) GetCpuReservation() int32`

GetCpuReservation returns the CpuReservation field if non-nil, zero value otherwise.

### GetCpuReservationOk

`func (o *PlatformUpdateApplicationInput) GetCpuReservationOk() (*int32, bool)`

GetCpuReservationOk returns a tuple with the CpuReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuReservation

`func (o *PlatformUpdateApplicationInput) SetCpuReservation(v int32)`

SetCpuReservation sets CpuReservation field to given value.

### HasCpuReservation

`func (o *PlatformUpdateApplicationInput) HasCpuReservation() bool`

HasCpuReservation returns a boolean if a field has been set.

### GetCpuLimit

`func (o *PlatformUpdateApplicationInput) GetCpuLimit() int32`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *PlatformUpdateApplicationInput) GetCpuLimitOk() (*int32, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *PlatformUpdateApplicationInput) SetCpuLimit(v int32)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *PlatformUpdateApplicationInput) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetReplicas

`func (o *PlatformUpdateApplicationInput) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PlatformUpdateApplicationInput) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PlatformUpdateApplicationInput) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PlatformUpdateApplicationInput) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetCommand

`func (o *PlatformUpdateApplicationInput) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PlatformUpdateApplicationInput) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PlatformUpdateApplicationInput) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *PlatformUpdateApplicationInput) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


