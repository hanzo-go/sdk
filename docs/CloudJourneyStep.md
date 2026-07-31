# CloudJourneyStep

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

### NewCloudJourneyStep

`func NewCloudJourneyStep() *CloudJourneyStep`

NewCloudJourneyStep instantiates a new CloudJourneyStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudJourneyStepWithDefaults

`func NewCloudJourneyStepWithDefaults() *CloudJourneyStep`

NewCloudJourneyStepWithDefaults instantiates a new CloudJourneyStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *CloudJourneyStep) GetArgs() map[string]map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CloudJourneyStep) GetArgsOk() (*map[string]map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CloudJourneyStep) SetArgs(v map[string]map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *CloudJourneyStep) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetDeps

`func (o *CloudJourneyStep) GetDeps() []string`

GetDeps returns the Deps field if non-nil, zero value otherwise.

### GetDepsOk

`func (o *CloudJourneyStep) GetDepsOk() (*[]string, bool)`

GetDepsOk returns a tuple with the Deps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeps

`func (o *CloudJourneyStep) SetDeps(v []string)`

SetDeps sets Deps field to given value.

### HasDeps

`func (o *CloudJourneyStep) HasDeps() bool`

HasDeps returns a boolean if a field has been set.

### GetDetail

`func (o *CloudJourneyStep) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CloudJourneyStep) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CloudJourneyStep) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CloudJourneyStep) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDraft

`func (o *CloudJourneyStep) GetDraft() string`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *CloudJourneyStep) GetDraftOk() (*string, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *CloudJourneyStep) SetDraft(v string)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *CloudJourneyStep) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetDraftInto

`func (o *CloudJourneyStep) GetDraftInto() string`

GetDraftInto returns the DraftInto field if non-nil, zero value otherwise.

### GetDraftIntoOk

`func (o *CloudJourneyStep) GetDraftIntoOk() (*string, bool)`

GetDraftIntoOk returns a tuple with the DraftInto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftInto

`func (o *CloudJourneyStep) SetDraftInto(v string)`

SetDraftInto sets DraftInto field to given value.

### HasDraftInto

`func (o *CloudJourneyStep) HasDraftInto() bool`

HasDraftInto returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudJourneyStep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudJourneyStep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudJourneyStep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudJourneyStep) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *CloudJourneyStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudJourneyStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudJourneyStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudJourneyStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *CloudJourneyStep) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *CloudJourneyStep) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *CloudJourneyStep) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *CloudJourneyStep) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSignal

`func (o *CloudJourneyStep) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *CloudJourneyStep) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *CloudJourneyStep) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *CloudJourneyStep) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetTitle

`func (o *CloudJourneyStep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudJourneyStep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudJourneyStep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudJourneyStep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTool

`func (o *CloudJourneyStep) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *CloudJourneyStep) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *CloudJourneyStep) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *CloudJourneyStep) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


