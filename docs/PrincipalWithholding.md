# PrincipalWithholding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount is the amount to withhold, U.S. dollars; empty when the rate is. | [optional] 
**Chapter** | Pointer to **string** | Chapter is \&quot;3\&quot; (IRC chapter 3, a payment to a foreign person) or \&quot;backup\&quot; (IRC §3406), and empty when nothing is withheld or it cannot yet be decided. | [optional] 
**Rate** | Pointer to **string** | Rate is the rate in percent: \&quot;30\&quot;, \&quot;24\&quot;, \&quot;10\&quot;, \&quot;0\&quot;. Empty when it cannot be decided on the facts held — see facts_required. | [optional] 
**Reason** | Pointer to **string** | Reason says why, in words. | [optional] 
**Rule** | Pointer to [**PrincipalRule**](PrincipalRule.md) | Rule is the rule that decided it. | [optional] 

## Methods

### NewPrincipalWithholding

`func NewPrincipalWithholding() *PrincipalWithholding`

NewPrincipalWithholding instantiates a new PrincipalWithholding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalWithholdingWithDefaults

`func NewPrincipalWithholdingWithDefaults() *PrincipalWithholding`

NewPrincipalWithholdingWithDefaults instantiates a new PrincipalWithholding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PrincipalWithholding) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PrincipalWithholding) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PrincipalWithholding) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PrincipalWithholding) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChapter

`func (o *PrincipalWithholding) GetChapter() string`

GetChapter returns the Chapter field if non-nil, zero value otherwise.

### GetChapterOk

`func (o *PrincipalWithholding) GetChapterOk() (*string, bool)`

GetChapterOk returns a tuple with the Chapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapter

`func (o *PrincipalWithholding) SetChapter(v string)`

SetChapter sets Chapter field to given value.

### HasChapter

`func (o *PrincipalWithholding) HasChapter() bool`

HasChapter returns a boolean if a field has been set.

### GetRate

`func (o *PrincipalWithholding) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *PrincipalWithholding) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *PrincipalWithholding) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *PrincipalWithholding) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetReason

`func (o *PrincipalWithholding) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PrincipalWithholding) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PrincipalWithholding) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PrincipalWithholding) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRule

`func (o *PrincipalWithholding) GetRule() PrincipalRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *PrincipalWithholding) GetRuleOk() (*PrincipalRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *PrincipalWithholding) SetRule(v PrincipalRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *PrincipalWithholding) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


