# TrainResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to **int64** | Budget bounds what the job&#39;s compute costs, in US cents; 0 is none beyond the payer&#39;s balance. Only a job Hanzo runs costs anything, so only it takes one. | [optional] 
**Devices** | Pointer to **[]string** | Devices are the accelerator kinds the job may run on — cuda, rocm, metal, vulkan, cpu; none named is any its base runs. | [optional] 
**Machines** | Pointer to **[]string** | Machines are the org&#39;s linked machines by name; the first leads. When it names none, Hanzo&#39;s executor runs the job and its compute is charged by the device-second. | [optional] 
**MaxSeconds** | Pointer to **int64** | MaxSeconds bounds the device-seconds the job&#39;s tasks use together; 0 is none. | [optional] 
**Steps** | Pointer to **int64** | Steps bounds optimizer steps; 0 runs the plan. | [optional] 

## Methods

### NewTrainResources

`func NewTrainResources() *TrainResources`

NewTrainResources instantiates a new TrainResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainResourcesWithDefaults

`func NewTrainResourcesWithDefaults() *TrainResources`

NewTrainResourcesWithDefaults instantiates a new TrainResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *TrainResources) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *TrainResources) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *TrainResources) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *TrainResources) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetDevices

`func (o *TrainResources) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *TrainResources) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *TrainResources) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *TrainResources) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetMachines

`func (o *TrainResources) GetMachines() []string`

GetMachines returns the Machines field if non-nil, zero value otherwise.

### GetMachinesOk

`func (o *TrainResources) GetMachinesOk() (*[]string, bool)`

GetMachinesOk returns a tuple with the Machines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachines

`func (o *TrainResources) SetMachines(v []string)`

SetMachines sets Machines field to given value.

### HasMachines

`func (o *TrainResources) HasMachines() bool`

HasMachines returns a boolean if a field has been set.

### GetMaxSeconds

`func (o *TrainResources) GetMaxSeconds() int64`

GetMaxSeconds returns the MaxSeconds field if non-nil, zero value otherwise.

### GetMaxSecondsOk

`func (o *TrainResources) GetMaxSecondsOk() (*int64, bool)`

GetMaxSecondsOk returns a tuple with the MaxSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSeconds

`func (o *TrainResources) SetMaxSeconds(v int64)`

SetMaxSeconds sets MaxSeconds field to given value.

### HasMaxSeconds

`func (o *TrainResources) HasMaxSeconds() bool`

HasMaxSeconds returns a boolean if a field has been set.

### GetSteps

`func (o *TrainResources) GetSteps() int64`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TrainResources) GetStepsOk() (*int64, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TrainResources) SetSteps(v int64)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TrainResources) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


