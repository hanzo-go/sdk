# Suggestion

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

### NewSuggestion

`func NewSuggestion() *Suggestion`

NewSuggestion instantiates a new Suggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestionWithDefaults

`func NewSuggestionWithDefaults() *Suggestion`

NewSuggestionWithDefaults instantiates a new Suggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatable

`func (o *Suggestion) GetAutomatable() bool`

GetAutomatable returns the Automatable field if non-nil, zero value otherwise.

### GetAutomatableOk

`func (o *Suggestion) GetAutomatableOk() (*bool, bool)`

GetAutomatableOk returns a tuple with the Automatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatable

`func (o *Suggestion) SetAutomatable(v bool)`

SetAutomatable sets Automatable field to given value.

### HasAutomatable

`func (o *Suggestion) HasAutomatable() bool`

HasAutomatable returns a boolean if a field has been set.

### GetDetail

`func (o *Suggestion) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Suggestion) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Suggestion) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Suggestion) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetRationale

`func (o *Suggestion) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *Suggestion) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *Suggestion) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *Suggestion) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### GetStepId

`func (o *Suggestion) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *Suggestion) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *Suggestion) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *Suggestion) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetTitle

`func (o *Suggestion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Suggestion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Suggestion) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Suggestion) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnlocks

`func (o *Suggestion) GetUnlocks() int64`

GetUnlocks returns the Unlocks field if non-nil, zero value otherwise.

### GetUnlocksOk

`func (o *Suggestion) GetUnlocksOk() (*int64, bool)`

GetUnlocksOk returns a tuple with the Unlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlocks

`func (o *Suggestion) SetUnlocks(v int64)`

SetUnlocks sets Unlocks field to given value.

### HasUnlocks

`func (o *Suggestion) HasUnlocks() bool`

HasUnlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


