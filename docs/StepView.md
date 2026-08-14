# StepView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Automatable** | Pointer to **bool** | Automatable is true when the Business AI can run this step (it names a tool). | [optional] 
**Available** | Pointer to **bool** | Available is true when every dependency is done or skipped. | [optional] 
**BlockedBy** | Pointer to **[]string** | BlockedBy lists the unfinished dependencies keeping the step unavailable. | [optional] 
**Deps** | Pointer to **[]string** | Dependencies are step ids that must be done/skipped before this step is available. The wire key is &#x60;deps&#x60; (the blueprint contract). | [optional] 
**Detail** | Pointer to **string** | Detail is the prose/juncture — what the Guide asks or explains here. | [optional] 
**Draft** | Pointer to **string** |  | [optional] 
**DraftInto** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled is the admin on/off lever; absent reads as enabled. | [optional] 
**Id** | Pointer to **string** | ID is the step&#39;s id, as it appears in the journey (e.g. \&quot;gsuite\&quot;). | [optional] 
**Section** | Pointer to **string** | Section is the phase (section id) this step groups under. | [optional] 
**Signal** | Pointer to **string** | Signal names the machine detector that auto-marks this step done. | [optional] 
**Source** | Pointer to **string** | Source records what marked the state: manual, auto (detected) or agent. | [optional] 
**State** | Pointer to **string** | State is the step&#39;s per-org lifecycle state: todo|in_progress|done|skipped. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Tool** | Pointer to **string** | Tool is the MCP tool the Business AI runs for \&quot;do it for me\&quot;; Args are its default arguments, Draft an optional AI prompt whose output fills the DraftInto arg (default \&quot;brief\&quot;). | [optional] 

## Methods

### NewStepView

`func NewStepView() *StepView`

NewStepView instantiates a new StepView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepViewWithDefaults

`func NewStepViewWithDefaults() *StepView`

NewStepViewWithDefaults instantiates a new StepView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *StepView) GetArgs() map[string]map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *StepView) GetArgsOk() (*map[string]map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *StepView) SetArgs(v map[string]map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *StepView) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetAutomatable

`func (o *StepView) GetAutomatable() bool`

GetAutomatable returns the Automatable field if non-nil, zero value otherwise.

### GetAutomatableOk

`func (o *StepView) GetAutomatableOk() (*bool, bool)`

GetAutomatableOk returns a tuple with the Automatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatable

`func (o *StepView) SetAutomatable(v bool)`

SetAutomatable sets Automatable field to given value.

### HasAutomatable

`func (o *StepView) HasAutomatable() bool`

HasAutomatable returns a boolean if a field has been set.

### GetAvailable

`func (o *StepView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *StepView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *StepView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *StepView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBlockedBy

`func (o *StepView) GetBlockedBy() []string`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *StepView) GetBlockedByOk() (*[]string, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *StepView) SetBlockedBy(v []string)`

SetBlockedBy sets BlockedBy field to given value.

### HasBlockedBy

`func (o *StepView) HasBlockedBy() bool`

HasBlockedBy returns a boolean if a field has been set.

### GetDeps

`func (o *StepView) GetDeps() []string`

GetDeps returns the Deps field if non-nil, zero value otherwise.

### GetDepsOk

`func (o *StepView) GetDepsOk() (*[]string, bool)`

GetDepsOk returns a tuple with the Deps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeps

`func (o *StepView) SetDeps(v []string)`

SetDeps sets Deps field to given value.

### HasDeps

`func (o *StepView) HasDeps() bool`

HasDeps returns a boolean if a field has been set.

### GetDetail

`func (o *StepView) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *StepView) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *StepView) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *StepView) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDraft

`func (o *StepView) GetDraft() string`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *StepView) GetDraftOk() (*string, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *StepView) SetDraft(v string)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *StepView) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetDraftInto

`func (o *StepView) GetDraftInto() string`

GetDraftInto returns the DraftInto field if non-nil, zero value otherwise.

### GetDraftIntoOk

`func (o *StepView) GetDraftIntoOk() (*string, bool)`

GetDraftIntoOk returns a tuple with the DraftInto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftInto

`func (o *StepView) SetDraftInto(v string)`

SetDraftInto sets DraftInto field to given value.

### HasDraftInto

`func (o *StepView) HasDraftInto() bool`

HasDraftInto returns a boolean if a field has been set.

### GetEnabled

`func (o *StepView) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StepView) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StepView) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StepView) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *StepView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StepView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StepView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StepView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *StepView) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *StepView) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *StepView) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *StepView) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSignal

`func (o *StepView) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *StepView) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *StepView) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *StepView) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetSource

`func (o *StepView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StepView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StepView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StepView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *StepView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StepView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StepView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StepView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTitle

`func (o *StepView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StepView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StepView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StepView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *StepView) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *StepView) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *StepView) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *StepView) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


