# GuideSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automatable** | Pointer to **bool** | Automatable is true when the step names a tool, so the Business AI can do it rather than only describe it. | [optional] 
**Detail** | Pointer to **string** | Detail is the step&#39;s own prose — what it asks for. | [optional] 
**Rationale** | Pointer to **string** | Rationale is why this step is being suggested NOW, written for the person reading it. It explains the ranking, not the step. | [optional] 
**StepId** | Pointer to **string** | StepID is the checklist step being recommended — the id every step route takes, so a caller can act on the suggestion directly. | [optional] 
**Title** | Pointer to **string** | Title is the step&#39;s own one-line quest. | [optional] 
**Unlocks** | Pointer to **int64** | Unlocks is how many downstream steps completing this one immediately makes available (its leverage) — the primary ranking key. | [optional] 

## Methods

### NewGuideSuggestion

`func NewGuideSuggestion() *GuideSuggestion`

NewGuideSuggestion instantiates a new GuideSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuideSuggestionWithDefaults

`func NewGuideSuggestionWithDefaults() *GuideSuggestion`

NewGuideSuggestionWithDefaults instantiates a new GuideSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatable

`func (o *GuideSuggestion) GetAutomatable() bool`

GetAutomatable returns the Automatable field if non-nil, zero value otherwise.

### GetAutomatableOk

`func (o *GuideSuggestion) GetAutomatableOk() (*bool, bool)`

GetAutomatableOk returns a tuple with the Automatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatable

`func (o *GuideSuggestion) SetAutomatable(v bool)`

SetAutomatable sets Automatable field to given value.

### HasAutomatable

`func (o *GuideSuggestion) HasAutomatable() bool`

HasAutomatable returns a boolean if a field has been set.

### GetDetail

`func (o *GuideSuggestion) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GuideSuggestion) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GuideSuggestion) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GuideSuggestion) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetRationale

`func (o *GuideSuggestion) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *GuideSuggestion) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *GuideSuggestion) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *GuideSuggestion) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### GetStepId

`func (o *GuideSuggestion) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *GuideSuggestion) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *GuideSuggestion) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *GuideSuggestion) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetTitle

`func (o *GuideSuggestion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GuideSuggestion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GuideSuggestion) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GuideSuggestion) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnlocks

`func (o *GuideSuggestion) GetUnlocks() int64`

GetUnlocks returns the Unlocks field if non-nil, zero value otherwise.

### GetUnlocksOk

`func (o *GuideSuggestion) GetUnlocksOk() (*int64, bool)`

GetUnlocksOk returns a tuple with the Unlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlocks

`func (o *GuideSuggestion) SetUnlocks(v int64)`

SetUnlocks sets Unlocks field to given value.

### HasUnlocks

`func (o *GuideSuggestion) HasUnlocks() bool`

HasUnlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


