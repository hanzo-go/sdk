# AutoAutomationIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled arms its schedule. Absent is true. | [optional] 
**Instructions** | Pointer to **string** | Instructions are what the agent is asked to do each run. Required. | [optional] 
**Model** | Pointer to **string** | Model is the model id to think with. Optional; absent is the default. | [optional] 
**Name** | Pointer to **string** | Name is what it is called. Required. | [optional] 
**Notify** | Pointer to **bool** | Notify sends a one-line summary when a run ends. | [optional] 
**Permissions** | Pointer to **string** | Permissions is auto (works without stopping) or ask (proposes, changes nothing). Absent is ask. | [optional] 
**Project** | Pointer to **string** | Project is a Dev project&#39;s slug to work in. Optional. | [optional] 
**Schedule** | Pointer to [**AutoSchedule**](AutoSchedule.md) | Schedule is when it runs. Absent runs it only on demand. | [optional] 

## Methods

### NewAutoAutomationIn

`func NewAutoAutomationIn() *AutoAutomationIn`

NewAutoAutomationIn instantiates a new AutoAutomationIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoAutomationInWithDefaults

`func NewAutoAutomationInWithDefaults() *AutoAutomationIn`

NewAutoAutomationInWithDefaults instantiates a new AutoAutomationIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoAutomationIn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoAutomationIn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoAutomationIn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoAutomationIn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInstructions

`func (o *AutoAutomationIn) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AutoAutomationIn) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AutoAutomationIn) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AutoAutomationIn) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *AutoAutomationIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AutoAutomationIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AutoAutomationIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AutoAutomationIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *AutoAutomationIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoAutomationIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoAutomationIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoAutomationIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotify

`func (o *AutoAutomationIn) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *AutoAutomationIn) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *AutoAutomationIn) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *AutoAutomationIn) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetPermissions

`func (o *AutoAutomationIn) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AutoAutomationIn) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AutoAutomationIn) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AutoAutomationIn) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetProject

`func (o *AutoAutomationIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AutoAutomationIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AutoAutomationIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AutoAutomationIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSchedule

`func (o *AutoAutomationIn) GetSchedule() AutoSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AutoAutomationIn) GetScheduleOk() (*AutoSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AutoAutomationIn) SetSchedule(v AutoSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AutoAutomationIn) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


