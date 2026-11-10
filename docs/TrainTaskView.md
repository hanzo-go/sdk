# TrainTaskView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** | Device is the accelerator class it runs on — gpu or cpu — and Devices how many. | [optional] 
**Devices** | Pointer to **int64** |  | [optional] 
**Held** | Pointer to **int64** |  | [optional] 
**Host** | Pointer to **string** | Host is the machine that claimed it. | [optional] 
**Id** | Pointer to **string** | ID names the task. | [optional] 
**Machine** | Pointer to **string** | Machine is the machine it was placed on, or empty for any. | [optional] 
**Role** | Pointer to **string** | Role is lead or join. | [optional] 
**Seconds** | Pointer to **int64** | Seconds are the device-seconds its reports metered, and Held those held for it. | [optional] 
**Seen** | Pointer to **int64** | Seen is its last report, unix seconds. | [optional] 
**Status** | Pointer to **string** | Status is open, claimed, done or lost. | [optional] 

## Methods

### NewTrainTaskView

`func NewTrainTaskView() *TrainTaskView`

NewTrainTaskView instantiates a new TrainTaskView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainTaskViewWithDefaults

`func NewTrainTaskViewWithDefaults() *TrainTaskView`

NewTrainTaskViewWithDefaults instantiates a new TrainTaskView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *TrainTaskView) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *TrainTaskView) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *TrainTaskView) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *TrainTaskView) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDevices

`func (o *TrainTaskView) GetDevices() int64`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *TrainTaskView) GetDevicesOk() (*int64, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *TrainTaskView) SetDevices(v int64)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *TrainTaskView) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetHeld

`func (o *TrainTaskView) GetHeld() int64`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *TrainTaskView) GetHeldOk() (*int64, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *TrainTaskView) SetHeld(v int64)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *TrainTaskView) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetHost

`func (o *TrainTaskView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TrainTaskView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TrainTaskView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TrainTaskView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *TrainTaskView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainTaskView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainTaskView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainTaskView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMachine

`func (o *TrainTaskView) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *TrainTaskView) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *TrainTaskView) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *TrainTaskView) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetRole

`func (o *TrainTaskView) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TrainTaskView) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TrainTaskView) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *TrainTaskView) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSeconds

`func (o *TrainTaskView) GetSeconds() int64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *TrainTaskView) GetSecondsOk() (*int64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *TrainTaskView) SetSeconds(v int64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *TrainTaskView) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetSeen

`func (o *TrainTaskView) GetSeen() int64`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *TrainTaskView) GetSeenOk() (*int64, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *TrainTaskView) SetSeen(v int64)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *TrainTaskView) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetStatus

`func (o *TrainTaskView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrainTaskView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrainTaskView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrainTaskView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


