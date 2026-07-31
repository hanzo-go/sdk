# CloudMachineList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Machines** | Pointer to [**[]CloudMachineView**](CloudMachineView.md) | Machines is every machine the org has: Visor-provisioned and BYO together. | [optional] 

## Methods

### NewCloudMachineList

`func NewCloudMachineList() *CloudMachineList`

NewCloudMachineList instantiates a new CloudMachineList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMachineListWithDefaults

`func NewCloudMachineListWithDefaults() *CloudMachineList`

NewCloudMachineListWithDefaults instantiates a new CloudMachineList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachines

`func (o *CloudMachineList) GetMachines() []CloudMachineView`

GetMachines returns the Machines field if non-nil, zero value otherwise.

### GetMachinesOk

`func (o *CloudMachineList) GetMachinesOk() (*[]CloudMachineView, bool)`

GetMachinesOk returns a tuple with the Machines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachines

`func (o *CloudMachineList) SetMachines(v []CloudMachineView)`

SetMachines sets Machines field to given value.

### HasMachines

`func (o *CloudMachineList) HasMachines() bool`

HasMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


