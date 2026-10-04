# GuideCurriculum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Steps** | Pointer to [**[]GuideJourneyStep**](GuideJourneyStep.md) | Steps are the enabled steps in authoring order. Order is the tiebreak the next-step logic walks, so it is part of the contract rather than cosmetic. | [optional] 
**Title** | Pointer to **string** | Title is the playbook&#39;s name as it heads the checklist. | [optional] 
**Version** | Pointer to **string** | Version identifies the authored playbook this journey was projected from, so two orgs on different playbooks can be told apart. It is the blueprint&#39;s own &#x60;version&#x60; string, not the store&#39;s numeric revision. | [optional] 

## Methods

### NewGuideCurriculum

`func NewGuideCurriculum() *GuideCurriculum`

NewGuideCurriculum instantiates a new GuideCurriculum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuideCurriculumWithDefaults

`func NewGuideCurriculumWithDefaults() *GuideCurriculum`

NewGuideCurriculumWithDefaults instantiates a new GuideCurriculum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteps

`func (o *GuideCurriculum) GetSteps() []GuideJourneyStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *GuideCurriculum) GetStepsOk() (*[]GuideJourneyStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *GuideCurriculum) SetSteps(v []GuideJourneyStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *GuideCurriculum) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTitle

`func (o *GuideCurriculum) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GuideCurriculum) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GuideCurriculum) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GuideCurriculum) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *GuideCurriculum) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GuideCurriculum) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GuideCurriculum) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GuideCurriculum) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


