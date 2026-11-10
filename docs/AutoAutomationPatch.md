# AutoAutomationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled arms or disarms its schedule. | [optional] 
**Event** | Pointer to [**AutoOutcomeTrigger**](AutoOutcomeTrigger.md) | Event replaces the trigger with a product outcome; null clears it. | [optional] 
**Id** | Pointer to **string** | ID is the automation, from the path. | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notify** | Pointer to **bool** |  | [optional] 
**Permissions** | Pointer to **string** |  | [optional] 
**PostTo** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**AutoSchedule**](AutoSchedule.md) |  | [optional] 
**Target** | Pointer to **string** | Target replaces the registered execution target; null clears it. | [optional] 

## Methods

### NewAutoAutomationPatch

`func NewAutoAutomationPatch() *AutoAutomationPatch`

NewAutoAutomationPatch instantiates a new AutoAutomationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoAutomationPatchWithDefaults

`func NewAutoAutomationPatchWithDefaults() *AutoAutomationPatch`

NewAutoAutomationPatchWithDefaults instantiates a new AutoAutomationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoAutomationPatch) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoAutomationPatch) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoAutomationPatch) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoAutomationPatch) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvent

`func (o *AutoAutomationPatch) GetEvent() AutoOutcomeTrigger`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AutoAutomationPatch) GetEventOk() (*AutoOutcomeTrigger, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AutoAutomationPatch) SetEvent(v AutoOutcomeTrigger)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AutoAutomationPatch) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *AutoAutomationPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoAutomationPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoAutomationPatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoAutomationPatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstructions

`func (o *AutoAutomationPatch) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AutoAutomationPatch) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AutoAutomationPatch) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AutoAutomationPatch) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *AutoAutomationPatch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AutoAutomationPatch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AutoAutomationPatch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AutoAutomationPatch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *AutoAutomationPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoAutomationPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoAutomationPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoAutomationPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotify

`func (o *AutoAutomationPatch) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *AutoAutomationPatch) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *AutoAutomationPatch) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *AutoAutomationPatch) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetPermissions

`func (o *AutoAutomationPatch) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AutoAutomationPatch) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AutoAutomationPatch) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AutoAutomationPatch) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPostTo

`func (o *AutoAutomationPatch) GetPostTo() []string`

GetPostTo returns the PostTo field if non-nil, zero value otherwise.

### GetPostToOk

`func (o *AutoAutomationPatch) GetPostToOk() (*[]string, bool)`

GetPostToOk returns a tuple with the PostTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTo

`func (o *AutoAutomationPatch) SetPostTo(v []string)`

SetPostTo sets PostTo field to given value.

### HasPostTo

`func (o *AutoAutomationPatch) HasPostTo() bool`

HasPostTo returns a boolean if a field has been set.

### GetProject

`func (o *AutoAutomationPatch) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AutoAutomationPatch) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AutoAutomationPatch) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AutoAutomationPatch) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSchedule

`func (o *AutoAutomationPatch) GetSchedule() AutoSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AutoAutomationPatch) GetScheduleOk() (*AutoSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AutoAutomationPatch) SetSchedule(v AutoSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AutoAutomationPatch) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTarget

`func (o *AutoAutomationPatch) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AutoAutomationPatch) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AutoAutomationPatch) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AutoAutomationPatch) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


