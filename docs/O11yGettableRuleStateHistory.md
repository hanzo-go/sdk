# O11yGettableRuleStateHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to [**[]O11yLabel**](O11yLabel.md) |  | [optional] 
**OverallState** | Pointer to **interface{}** |  | [optional] 
**OverallStateChanged** | Pointer to **bool** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**RuleName** | Pointer to **string** |  | [optional] 
**State** | Pointer to **interface{}** |  | [optional] 
**StateChanged** | Pointer to **bool** |  | [optional] 
**UnixMilli** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewO11yGettableRuleStateHistory

`func NewO11yGettableRuleStateHistory() *O11yGettableRuleStateHistory`

NewO11yGettableRuleStateHistory instantiates a new O11yGettableRuleStateHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableRuleStateHistoryWithDefaults

`func NewO11yGettableRuleStateHistoryWithDefaults() *O11yGettableRuleStateHistory`

NewO11yGettableRuleStateHistoryWithDefaults instantiates a new O11yGettableRuleStateHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *O11yGettableRuleStateHistory) GetFingerprint() int32`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *O11yGettableRuleStateHistory) GetFingerprintOk() (*int32, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *O11yGettableRuleStateHistory) SetFingerprint(v int32)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *O11yGettableRuleStateHistory) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetLabels

`func (o *O11yGettableRuleStateHistory) GetLabels() []O11yLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yGettableRuleStateHistory) GetLabelsOk() (*[]O11yLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yGettableRuleStateHistory) SetLabels(v []O11yLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yGettableRuleStateHistory) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOverallState

`func (o *O11yGettableRuleStateHistory) GetOverallState() interface{}`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *O11yGettableRuleStateHistory) GetOverallStateOk() (*interface{}, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallState

`func (o *O11yGettableRuleStateHistory) SetOverallState(v interface{})`

SetOverallState sets OverallState field to given value.

### HasOverallState

`func (o *O11yGettableRuleStateHistory) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### SetOverallStateNil

`func (o *O11yGettableRuleStateHistory) SetOverallStateNil(b bool)`

 SetOverallStateNil sets the value for OverallState to be an explicit nil

### UnsetOverallState
`func (o *O11yGettableRuleStateHistory) UnsetOverallState()`

UnsetOverallState ensures that no value is present for OverallState, not even an explicit nil
### GetOverallStateChanged

`func (o *O11yGettableRuleStateHistory) GetOverallStateChanged() bool`

GetOverallStateChanged returns the OverallStateChanged field if non-nil, zero value otherwise.

### GetOverallStateChangedOk

`func (o *O11yGettableRuleStateHistory) GetOverallStateChangedOk() (*bool, bool)`

GetOverallStateChangedOk returns a tuple with the OverallStateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStateChanged

`func (o *O11yGettableRuleStateHistory) SetOverallStateChanged(v bool)`

SetOverallStateChanged sets OverallStateChanged field to given value.

### HasOverallStateChanged

`func (o *O11yGettableRuleStateHistory) HasOverallStateChanged() bool`

HasOverallStateChanged returns a boolean if a field has been set.

### GetRuleId

`func (o *O11yGettableRuleStateHistory) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *O11yGettableRuleStateHistory) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *O11yGettableRuleStateHistory) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *O11yGettableRuleStateHistory) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *O11yGettableRuleStateHistory) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *O11yGettableRuleStateHistory) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *O11yGettableRuleStateHistory) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *O11yGettableRuleStateHistory) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetState

`func (o *O11yGettableRuleStateHistory) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *O11yGettableRuleStateHistory) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *O11yGettableRuleStateHistory) SetState(v interface{})`

SetState sets State field to given value.

### HasState

`func (o *O11yGettableRuleStateHistory) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *O11yGettableRuleStateHistory) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *O11yGettableRuleStateHistory) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStateChanged

`func (o *O11yGettableRuleStateHistory) GetStateChanged() bool`

GetStateChanged returns the StateChanged field if non-nil, zero value otherwise.

### GetStateChangedOk

`func (o *O11yGettableRuleStateHistory) GetStateChangedOk() (*bool, bool)`

GetStateChangedOk returns a tuple with the StateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChanged

`func (o *O11yGettableRuleStateHistory) SetStateChanged(v bool)`

SetStateChanged sets StateChanged field to given value.

### HasStateChanged

`func (o *O11yGettableRuleStateHistory) HasStateChanged() bool`

HasStateChanged returns a boolean if a field has been set.

### GetUnixMilli

`func (o *O11yGettableRuleStateHistory) GetUnixMilli() int32`

GetUnixMilli returns the UnixMilli field if non-nil, zero value otherwise.

### GetUnixMilliOk

`func (o *O11yGettableRuleStateHistory) GetUnixMilliOk() (*int32, bool)`

GetUnixMilliOk returns a tuple with the UnixMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixMilli

`func (o *O11yGettableRuleStateHistory) SetUnixMilli(v int32)`

SetUnixMilli sets UnixMilli field to given value.

### HasUnixMilli

`func (o *O11yGettableRuleStateHistory) HasUnixMilli() bool`

HasUnixMilli returns a boolean if a field has been set.

### GetValue

`func (o *O11yGettableRuleStateHistory) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yGettableRuleStateHistory) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yGettableRuleStateHistory) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yGettableRuleStateHistory) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


