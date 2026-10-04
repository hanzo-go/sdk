# GuideOverviewView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to **bool** | Custom is true when the org replaced the shared playbook with one of its own — the difference between \&quot;everyone&#39;s checklist\&quot; and \&quot;the one you authored\&quot;. | [optional] 
**Funnel** | Pointer to [**GuideFunnel**](GuideFunnel.md) | Funnel is the org&#39;s analytics lens, present only where the read asked for it — absent means it was not requested, never that the org has no traffic. | [optional] 
**Progress** | Pointer to [**GuideProgressView**](GuideProgressView.md) | Progress is how far through the journey the org is. | [optional] 
**Steps** | Pointer to [**[]GuideStepView**](GuideStepView.md) | Steps are every enabled step with the org&#39;s own state folded in, in authoring order. | [optional] 
**Title** | Pointer to **string** | Title is the playbook&#39;s name as it heads the checklist. | [optional] 
**Version** | Pointer to **string** | Version identifies the playbook this journey came from, so a caller can tell that the checklist itself changed under them. | [optional] 

## Methods

### NewGuideOverviewView

`func NewGuideOverviewView() *GuideOverviewView`

NewGuideOverviewView instantiates a new GuideOverviewView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuideOverviewViewWithDefaults

`func NewGuideOverviewViewWithDefaults() *GuideOverviewView`

NewGuideOverviewViewWithDefaults instantiates a new GuideOverviewView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *GuideOverviewView) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *GuideOverviewView) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *GuideOverviewView) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *GuideOverviewView) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetFunnel

`func (o *GuideOverviewView) GetFunnel() GuideFunnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *GuideOverviewView) GetFunnelOk() (*GuideFunnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *GuideOverviewView) SetFunnel(v GuideFunnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *GuideOverviewView) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetProgress

`func (o *GuideOverviewView) GetProgress() GuideProgressView`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *GuideOverviewView) GetProgressOk() (*GuideProgressView, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *GuideOverviewView) SetProgress(v GuideProgressView)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *GuideOverviewView) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSteps

`func (o *GuideOverviewView) GetSteps() []GuideStepView`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *GuideOverviewView) GetStepsOk() (*[]GuideStepView, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *GuideOverviewView) SetSteps(v []GuideStepView)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *GuideOverviewView) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTitle

`func (o *GuideOverviewView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GuideOverviewView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GuideOverviewView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GuideOverviewView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *GuideOverviewView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GuideOverviewView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GuideOverviewView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GuideOverviewView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


