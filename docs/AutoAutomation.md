# AutoAutomation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Created and Updated are RFC 3339 UTC. | [optional] 
**Draft** | Pointer to **bool** | Draft is a flow with no step yet: it has no instructions, never runs, and becomes an automation when its instructions are saved. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is whether its schedule is armed. | [optional] 
**Id** | Pointer to **string** | ID is the automation&#39;s id, which is also its flow&#39;s. | [optional] 
**Instructions** | Pointer to **string** | Instructions are what the agent is asked to do each run. | [optional] 
**Last** | Pointer to [**AutoLastRun**](AutoLastRun.md) | Last is its most recent run that was not skipped; null before the first. | [optional] 
**Model** | Pointer to **string** | Model is the model the agent thinks with; null for the default. | [optional] 
**Name** | Pointer to **string** | Name is what it is called. | [optional] 
**Next** | Pointer to **string** | Next is when it runs next, RFC 3339 UTC; null when manual or disabled. | [optional] 
**Notify** | Pointer to **bool** | Notify sends a one-line summary to the automation&#39;s person when a run ends. | [optional] 
**Permissions** | Pointer to **string** | Permissions is auto (works and uses connectors without stopping) or ask (changes nothing and ends with the actions it proposes; agent.go agentTask). | [optional] 
**Project** | Pointer to **string** | Project is the Dev project the run works in, by its slug; null for none. | [optional] 
**Schedule** | Pointer to [**AutoSchedule**](AutoSchedule.md) | Schedule is when it runs. | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoAutomation

`func NewAutoAutomation() *AutoAutomation`

NewAutoAutomation instantiates a new AutoAutomation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoAutomationWithDefaults

`func NewAutoAutomationWithDefaults() *AutoAutomation`

NewAutoAutomationWithDefaults instantiates a new AutoAutomation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AutoAutomation) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AutoAutomation) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AutoAutomation) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AutoAutomation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDraft

`func (o *AutoAutomation) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *AutoAutomation) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *AutoAutomation) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *AutoAutomation) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetEnabled

`func (o *AutoAutomation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoAutomation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoAutomation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoAutomation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *AutoAutomation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoAutomation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoAutomation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoAutomation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstructions

`func (o *AutoAutomation) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AutoAutomation) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AutoAutomation) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AutoAutomation) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetLast

`func (o *AutoAutomation) GetLast() AutoLastRun`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *AutoAutomation) GetLastOk() (*AutoLastRun, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *AutoAutomation) SetLast(v AutoLastRun)`

SetLast sets Last field to given value.

### HasLast

`func (o *AutoAutomation) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetModel

`func (o *AutoAutomation) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AutoAutomation) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AutoAutomation) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AutoAutomation) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *AutoAutomation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoAutomation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoAutomation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoAutomation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNext

`func (o *AutoAutomation) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AutoAutomation) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AutoAutomation) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *AutoAutomation) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetNotify

`func (o *AutoAutomation) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *AutoAutomation) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *AutoAutomation) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *AutoAutomation) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetPermissions

`func (o *AutoAutomation) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AutoAutomation) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AutoAutomation) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AutoAutomation) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetProject

`func (o *AutoAutomation) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AutoAutomation) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AutoAutomation) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AutoAutomation) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSchedule

`func (o *AutoAutomation) GetSchedule() AutoSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AutoAutomation) GetScheduleOk() (*AutoSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AutoAutomation) SetSchedule(v AutoSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AutoAutomation) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetUpdated

`func (o *AutoAutomation) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutoAutomation) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutoAutomation) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutoAutomation) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


