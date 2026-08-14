# JourneyStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Deps** | Pointer to **[]string** | Dependencies are step ids that must be done/skipped before this step is available. The wire key is &#x60;deps&#x60; (the blueprint contract); the Go field keeps its descriptive name. | [optional] 
**Detail** | Pointer to **string** | the prose/juncture — what the Guide asks/explains here | [optional] 
**Draft** | Pointer to **string** |  | [optional] 
**DraftInto** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled is the admin on/off lever. A NIL pointer reads as ENABLED (absence &#x3D;&#x3D; on): a legacy/org curriculum that omits the field keeps every step, and only an explicit &#x60;enabled: false&#x60; (an admin disable) drops a step from the journey. See on() in blueprint.go and the Blueprint.Curriculum() projection. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Section** | Pointer to **string** | the phase (section id) this step groups under | [optional] 
**Signal** | Pointer to **string** | Signal, when set, names a machine detector (detect.go). When the detector reports the org&#39;s real state present, the step auto-marks done. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Tool** | Pointer to **string** | Tool, when set, is the MCP tool the Business AI runs for \&quot;do it for me\&quot;. Args are its default arguments; Draft is an optional AI prompt whose output fills the DraftInto arg (default \&quot;brief\&quot;). | [optional] 

## Methods

### NewJourneyStep

`func NewJourneyStep() *JourneyStep`

NewJourneyStep instantiates a new JourneyStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyStepWithDefaults

`func NewJourneyStepWithDefaults() *JourneyStep`

NewJourneyStepWithDefaults instantiates a new JourneyStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *JourneyStep) GetArgs() map[string]map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *JourneyStep) GetArgsOk() (*map[string]map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *JourneyStep) SetArgs(v map[string]map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *JourneyStep) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetDeps

`func (o *JourneyStep) GetDeps() []string`

GetDeps returns the Deps field if non-nil, zero value otherwise.

### GetDepsOk

`func (o *JourneyStep) GetDepsOk() (*[]string, bool)`

GetDepsOk returns a tuple with the Deps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeps

`func (o *JourneyStep) SetDeps(v []string)`

SetDeps sets Deps field to given value.

### HasDeps

`func (o *JourneyStep) HasDeps() bool`

HasDeps returns a boolean if a field has been set.

### GetDetail

`func (o *JourneyStep) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *JourneyStep) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *JourneyStep) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *JourneyStep) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDraft

`func (o *JourneyStep) GetDraft() string`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *JourneyStep) GetDraftOk() (*string, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *JourneyStep) SetDraft(v string)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *JourneyStep) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetDraftInto

`func (o *JourneyStep) GetDraftInto() string`

GetDraftInto returns the DraftInto field if non-nil, zero value otherwise.

### GetDraftIntoOk

`func (o *JourneyStep) GetDraftIntoOk() (*string, bool)`

GetDraftIntoOk returns a tuple with the DraftInto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftInto

`func (o *JourneyStep) SetDraftInto(v string)`

SetDraftInto sets DraftInto field to given value.

### HasDraftInto

`func (o *JourneyStep) HasDraftInto() bool`

HasDraftInto returns a boolean if a field has been set.

### GetEnabled

`func (o *JourneyStep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JourneyStep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JourneyStep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JourneyStep) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *JourneyStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JourneyStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JourneyStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JourneyStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *JourneyStep) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *JourneyStep) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *JourneyStep) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *JourneyStep) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSignal

`func (o *JourneyStep) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *JourneyStep) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *JourneyStep) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *JourneyStep) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetTitle

`func (o *JourneyStep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *JourneyStep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *JourneyStep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *JourneyStep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *JourneyStep) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *JourneyStep) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *JourneyStep) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *JourneyStep) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


