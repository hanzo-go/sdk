# TrainClaimIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bases** | Pointer to **[]string** | Bases are the base models its trainer runs. | [optional] 
**Devices** | Pointer to **[]string** | Devices are its accelerators, one entry per device: cuda, rocm, metal, vulkan, cpu. A task runs on those of one kind its job allows. | [optional] 
**Machine** | Pointer to **string** | Machine is the machine claiming, by name: one of the org&#39;s linked machines, or for a platform claim the Hanzo host. | [optional] 
**Platform** | Pointer to **bool** | Platform claims for Hanzo&#39;s executor: a task of any org&#39;s job that names no machines. A SuperAdmin&#39;s claim alone, and on the platform&#39;s audit trail. | [optional] 
**Supports** | Pointer to [**TrainSupports**](TrainSupports.md) | Supports is what its build runs. | [optional] 

## Methods

### NewTrainClaimIn

`func NewTrainClaimIn() *TrainClaimIn`

NewTrainClaimIn instantiates a new TrainClaimIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainClaimInWithDefaults

`func NewTrainClaimInWithDefaults() *TrainClaimIn`

NewTrainClaimInWithDefaults instantiates a new TrainClaimIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBases

`func (o *TrainClaimIn) GetBases() []string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *TrainClaimIn) GetBasesOk() (*[]string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *TrainClaimIn) SetBases(v []string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *TrainClaimIn) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetDevices

`func (o *TrainClaimIn) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *TrainClaimIn) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *TrainClaimIn) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *TrainClaimIn) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetMachine

`func (o *TrainClaimIn) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *TrainClaimIn) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *TrainClaimIn) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *TrainClaimIn) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetPlatform

`func (o *TrainClaimIn) GetPlatform() bool`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *TrainClaimIn) GetPlatformOk() (*bool, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *TrainClaimIn) SetPlatform(v bool)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *TrainClaimIn) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSupports

`func (o *TrainClaimIn) GetSupports() TrainSupports`

GetSupports returns the Supports field if non-nil, zero value otherwise.

### GetSupportsOk

`func (o *TrainClaimIn) GetSupportsOk() (*TrainSupports, bool)`

GetSupportsOk returns a tuple with the Supports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupports

`func (o *TrainClaimIn) SetSupports(v TrainSupports)`

SetSupports sets Supports field to given value.

### HasSupports

`func (o *TrainClaimIn) HasSupports() bool`

HasSupports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


