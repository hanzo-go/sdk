# O11yRuleStateHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to **interface{}** |  | [optional] 
**OverallState** | Pointer to **interface{}** |  | [optional] 
**OverallStateChanged** | Pointer to **bool** |  | [optional] 
**RelatedLogsLink** | Pointer to **string** |  | [optional] 
**RelatedTracesLink** | Pointer to **string** |  | [optional] 
**RuleID** | Pointer to **string** |  | [optional] 
**RuleName** | Pointer to **string** |  | [optional] 
**State** | Pointer to **interface{}** |  | [optional] 
**StateChanged** | Pointer to **bool** |  | [optional] 
**UnixMilli** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewO11yRuleStateHistory

`func NewO11yRuleStateHistory() *O11yRuleStateHistory`

NewO11yRuleStateHistory instantiates a new O11yRuleStateHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yRuleStateHistoryWithDefaults

`func NewO11yRuleStateHistoryWithDefaults() *O11yRuleStateHistory`

NewO11yRuleStateHistoryWithDefaults instantiates a new O11yRuleStateHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *O11yRuleStateHistory) GetFingerprint() int32`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *O11yRuleStateHistory) GetFingerprintOk() (*int32, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *O11yRuleStateHistory) SetFingerprint(v int32)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *O11yRuleStateHistory) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetLabels

`func (o *O11yRuleStateHistory) GetLabels() interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yRuleStateHistory) GetLabelsOk() (*interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yRuleStateHistory) SetLabels(v interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yRuleStateHistory) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *O11yRuleStateHistory) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *O11yRuleStateHistory) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetOverallState

`func (o *O11yRuleStateHistory) GetOverallState() interface{}`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *O11yRuleStateHistory) GetOverallStateOk() (*interface{}, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallState

`func (o *O11yRuleStateHistory) SetOverallState(v interface{})`

SetOverallState sets OverallState field to given value.

### HasOverallState

`func (o *O11yRuleStateHistory) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### SetOverallStateNil

`func (o *O11yRuleStateHistory) SetOverallStateNil(b bool)`

 SetOverallStateNil sets the value for OverallState to be an explicit nil

### UnsetOverallState
`func (o *O11yRuleStateHistory) UnsetOverallState()`

UnsetOverallState ensures that no value is present for OverallState, not even an explicit nil
### GetOverallStateChanged

`func (o *O11yRuleStateHistory) GetOverallStateChanged() bool`

GetOverallStateChanged returns the OverallStateChanged field if non-nil, zero value otherwise.

### GetOverallStateChangedOk

`func (o *O11yRuleStateHistory) GetOverallStateChangedOk() (*bool, bool)`

GetOverallStateChangedOk returns a tuple with the OverallStateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStateChanged

`func (o *O11yRuleStateHistory) SetOverallStateChanged(v bool)`

SetOverallStateChanged sets OverallStateChanged field to given value.

### HasOverallStateChanged

`func (o *O11yRuleStateHistory) HasOverallStateChanged() bool`

HasOverallStateChanged returns a boolean if a field has been set.

### GetRelatedLogsLink

`func (o *O11yRuleStateHistory) GetRelatedLogsLink() string`

GetRelatedLogsLink returns the RelatedLogsLink field if non-nil, zero value otherwise.

### GetRelatedLogsLinkOk

`func (o *O11yRuleStateHistory) GetRelatedLogsLinkOk() (*string, bool)`

GetRelatedLogsLinkOk returns a tuple with the RelatedLogsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedLogsLink

`func (o *O11yRuleStateHistory) SetRelatedLogsLink(v string)`

SetRelatedLogsLink sets RelatedLogsLink field to given value.

### HasRelatedLogsLink

`func (o *O11yRuleStateHistory) HasRelatedLogsLink() bool`

HasRelatedLogsLink returns a boolean if a field has been set.

### GetRelatedTracesLink

`func (o *O11yRuleStateHistory) GetRelatedTracesLink() string`

GetRelatedTracesLink returns the RelatedTracesLink field if non-nil, zero value otherwise.

### GetRelatedTracesLinkOk

`func (o *O11yRuleStateHistory) GetRelatedTracesLinkOk() (*string, bool)`

GetRelatedTracesLinkOk returns a tuple with the RelatedTracesLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTracesLink

`func (o *O11yRuleStateHistory) SetRelatedTracesLink(v string)`

SetRelatedTracesLink sets RelatedTracesLink field to given value.

### HasRelatedTracesLink

`func (o *O11yRuleStateHistory) HasRelatedTracesLink() bool`

HasRelatedTracesLink returns a boolean if a field has been set.

### GetRuleID

`func (o *O11yRuleStateHistory) GetRuleID() string`

GetRuleID returns the RuleID field if non-nil, zero value otherwise.

### GetRuleIDOk

`func (o *O11yRuleStateHistory) GetRuleIDOk() (*string, bool)`

GetRuleIDOk returns a tuple with the RuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleID

`func (o *O11yRuleStateHistory) SetRuleID(v string)`

SetRuleID sets RuleID field to given value.

### HasRuleID

`func (o *O11yRuleStateHistory) HasRuleID() bool`

HasRuleID returns a boolean if a field has been set.

### GetRuleName

`func (o *O11yRuleStateHistory) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *O11yRuleStateHistory) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *O11yRuleStateHistory) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *O11yRuleStateHistory) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetState

`func (o *O11yRuleStateHistory) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *O11yRuleStateHistory) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *O11yRuleStateHistory) SetState(v interface{})`

SetState sets State field to given value.

### HasState

`func (o *O11yRuleStateHistory) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *O11yRuleStateHistory) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *O11yRuleStateHistory) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStateChanged

`func (o *O11yRuleStateHistory) GetStateChanged() bool`

GetStateChanged returns the StateChanged field if non-nil, zero value otherwise.

### GetStateChangedOk

`func (o *O11yRuleStateHistory) GetStateChangedOk() (*bool, bool)`

GetStateChangedOk returns a tuple with the StateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChanged

`func (o *O11yRuleStateHistory) SetStateChanged(v bool)`

SetStateChanged sets StateChanged field to given value.

### HasStateChanged

`func (o *O11yRuleStateHistory) HasStateChanged() bool`

HasStateChanged returns a boolean if a field has been set.

### GetUnixMilli

`func (o *O11yRuleStateHistory) GetUnixMilli() int32`

GetUnixMilli returns the UnixMilli field if non-nil, zero value otherwise.

### GetUnixMilliOk

`func (o *O11yRuleStateHistory) GetUnixMilliOk() (*int32, bool)`

GetUnixMilliOk returns a tuple with the UnixMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixMilli

`func (o *O11yRuleStateHistory) SetUnixMilli(v int32)`

SetUnixMilli sets UnixMilli field to given value.

### HasUnixMilli

`func (o *O11yRuleStateHistory) HasUnixMilli() bool`

HasUnixMilli returns a boolean if a field has been set.

### GetValue

`func (o *O11yRuleStateHistory) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yRuleStateHistory) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yRuleStateHistory) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yRuleStateHistory) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


