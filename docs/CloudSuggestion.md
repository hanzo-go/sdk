# CloudSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automatable** | Pointer to **bool** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Rationale** | Pointer to **string** |  | [optional] 
**StepId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Unlocks** | Pointer to **int32** | Unlocks is how many downstream steps completing this one immediately makes available (its leverage) — the primary ranking key. | [optional] 

## Methods

### NewCloudSuggestion

`func NewCloudSuggestion() *CloudSuggestion`

NewCloudSuggestion instantiates a new CloudSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSuggestionWithDefaults

`func NewCloudSuggestionWithDefaults() *CloudSuggestion`

NewCloudSuggestionWithDefaults instantiates a new CloudSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatable

`func (o *CloudSuggestion) GetAutomatable() bool`

GetAutomatable returns the Automatable field if non-nil, zero value otherwise.

### GetAutomatableOk

`func (o *CloudSuggestion) GetAutomatableOk() (*bool, bool)`

GetAutomatableOk returns a tuple with the Automatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatable

`func (o *CloudSuggestion) SetAutomatable(v bool)`

SetAutomatable sets Automatable field to given value.

### HasAutomatable

`func (o *CloudSuggestion) HasAutomatable() bool`

HasAutomatable returns a boolean if a field has been set.

### GetDetail

`func (o *CloudSuggestion) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CloudSuggestion) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CloudSuggestion) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CloudSuggestion) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetRationale

`func (o *CloudSuggestion) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *CloudSuggestion) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *CloudSuggestion) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *CloudSuggestion) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### GetStepId

`func (o *CloudSuggestion) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *CloudSuggestion) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *CloudSuggestion) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *CloudSuggestion) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetTitle

`func (o *CloudSuggestion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudSuggestion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudSuggestion) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudSuggestion) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnlocks

`func (o *CloudSuggestion) GetUnlocks() int32`

GetUnlocks returns the Unlocks field if non-nil, zero value otherwise.

### GetUnlocksOk

`func (o *CloudSuggestion) GetUnlocksOk() (*int32, bool)`

GetUnlocksOk returns a tuple with the Unlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlocks

`func (o *CloudSuggestion) SetUnlocks(v int32)`

SetUnlocks sets Unlocks field to given value.

### HasUnlocks

`func (o *CloudSuggestion) HasUnlocks() bool`

HasUnlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


