# GuideJourneyStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]interface{}** | Args are the tool&#39;s default arguments, merged under whatever the caller passes at run time, so a step ships with the arguments that make it work. | [optional] 
**Deps** | Pointer to **[]string** | Dependencies are step ids that must be done/skipped before this step is available. The wire key is &#x60;deps&#x60; (the blueprint contract); the Go field keeps its descriptive name. | [optional] 
**Detail** | Pointer to **string** | Detail is the juncture — what the Guide explains, or asks for, at this step. | [optional] 
**Draft** | Pointer to **string** | Draft, when set, is the prompt the embedded AI answers first; its output is folded into one of Args before the tool runs, so the model writes the content and the tool only delivers it. | [optional] 
**DraftInto** | Pointer to **string** | DraftInto names the argument the drafted text lands in. Empty means \&quot;brief\&quot;. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is the admin on/off lever. A NIL pointer reads as ENABLED (absence &#x3D;&#x3D; on): a legacy/org curriculum that omits the field keeps every step, and only an explicit &#x60;enabled: false&#x60; (an admin disable) drops a step from the journey. See on() in blueprint.go and the Blueprint.Curriculum() projection. | [optional] 
**Id** | Pointer to **string** | ID is the stable slug the whole plane addresses this step by — the value in &#x60;deps&#x60;, in &#x60;next&#x60;, in the progress rows, and in the URL of every step route. Renaming it orphans an org&#39;s recorded progress for this step. | [optional] 
**Section** | Pointer to **string** | Section is the id of the phase this step groups under. A disabled section takes its steps out of the journey with it. | [optional] 
**Signal** | Pointer to **string** | Signal, when set, names a machine detector (detect.go). When the detector reports the org&#39;s real state present, the step auto-marks done. | [optional] 
**Title** | Pointer to **string** | Title is the one-line quest as a person reads it in the checklist. | [optional] 
**Tool** | Pointer to **string** | Tool, when set, names the MCP tool the Business AI runs for \&quot;do it for me\&quot;. A step with no tool can only be completed by a person; it is the field the &#x60;automatable&#x60; flag on every projection of this step is derived from. | [optional] 

## Methods

### NewGuideJourneyStep

`func NewGuideJourneyStep() *GuideJourneyStep`

NewGuideJourneyStep instantiates a new GuideJourneyStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuideJourneyStepWithDefaults

`func NewGuideJourneyStepWithDefaults() *GuideJourneyStep`

NewGuideJourneyStepWithDefaults instantiates a new GuideJourneyStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *GuideJourneyStep) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *GuideJourneyStep) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *GuideJourneyStep) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *GuideJourneyStep) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetDeps

`func (o *GuideJourneyStep) GetDeps() []string`

GetDeps returns the Deps field if non-nil, zero value otherwise.

### GetDepsOk

`func (o *GuideJourneyStep) GetDepsOk() (*[]string, bool)`

GetDepsOk returns a tuple with the Deps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeps

`func (o *GuideJourneyStep) SetDeps(v []string)`

SetDeps sets Deps field to given value.

### HasDeps

`func (o *GuideJourneyStep) HasDeps() bool`

HasDeps returns a boolean if a field has been set.

### GetDetail

`func (o *GuideJourneyStep) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GuideJourneyStep) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GuideJourneyStep) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GuideJourneyStep) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDraft

`func (o *GuideJourneyStep) GetDraft() string`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *GuideJourneyStep) GetDraftOk() (*string, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *GuideJourneyStep) SetDraft(v string)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *GuideJourneyStep) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetDraftInto

`func (o *GuideJourneyStep) GetDraftInto() string`

GetDraftInto returns the DraftInto field if non-nil, zero value otherwise.

### GetDraftIntoOk

`func (o *GuideJourneyStep) GetDraftIntoOk() (*string, bool)`

GetDraftIntoOk returns a tuple with the DraftInto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftInto

`func (o *GuideJourneyStep) SetDraftInto(v string)`

SetDraftInto sets DraftInto field to given value.

### HasDraftInto

`func (o *GuideJourneyStep) HasDraftInto() bool`

HasDraftInto returns a boolean if a field has been set.

### GetEnabled

`func (o *GuideJourneyStep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GuideJourneyStep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GuideJourneyStep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GuideJourneyStep) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *GuideJourneyStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuideJourneyStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuideJourneyStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GuideJourneyStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *GuideJourneyStep) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *GuideJourneyStep) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *GuideJourneyStep) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *GuideJourneyStep) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSignal

`func (o *GuideJourneyStep) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *GuideJourneyStep) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *GuideJourneyStep) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *GuideJourneyStep) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetTitle

`func (o *GuideJourneyStep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GuideJourneyStep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GuideJourneyStep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GuideJourneyStep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *GuideJourneyStep) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *GuideJourneyStep) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *GuideJourneyStep) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *GuideJourneyStep) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


