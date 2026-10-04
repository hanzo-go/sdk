# ComputeMachineList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Machines** | Pointer to [**[]ComputeMachineView**](ComputeMachineView.md) | Machines is every machine the org has: Visor-provisioned and BYO together. | [optional] 

## Methods

### NewComputeMachineList

`func NewComputeMachineList() *ComputeMachineList`

NewComputeMachineList instantiates a new ComputeMachineList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeMachineListWithDefaults

`func NewComputeMachineListWithDefaults() *ComputeMachineList`

NewComputeMachineListWithDefaults instantiates a new ComputeMachineList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachines

`func (o *ComputeMachineList) GetMachines() []ComputeMachineView`

GetMachines returns the Machines field if non-nil, zero value otherwise.

### GetMachinesOk

`func (o *ComputeMachineList) GetMachinesOk() (*[]ComputeMachineView, bool)`

GetMachinesOk returns a tuple with the Machines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachines

`func (o *ComputeMachineList) SetMachines(v []ComputeMachineView)`

SetMachines sets Machines field to given value.

### HasMachines

`func (o *ComputeMachineList) HasMachines() bool`

HasMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


