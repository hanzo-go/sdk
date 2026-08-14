# RiskModelValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address names this value by its own content: the model&#39;s shape, the geometry seed, its position in the window, its threshold, its masses as IEEE-754 bits and the fold watermark behind them. Nothing else — no clock, no counter and deliberately NOT the organisation, so an identical model has one name and a name is never an authority. Holding another organisation&#39;s address resolves nothing. | [optional] 
**At** | Pointer to **string** | At is when it was published, RFC 3339, on the server clock. You do not supply it: a record whose date the audited party chose is not a record. | [optional] 
**Learned** | Pointer to **int32** | Learned is how many events are behind the masses. | [optional] 
**Sequence** | Pointer to **int32** | Sequence is this value&#39;s place in YOUR organisation&#39;s own history, from 1 and contiguous until retention disposes of the oldest. | [optional] 
**Shape** | Pointer to **string** | Shape NAMES the model space the masses are only meaningful against, as &#x60;&lt;family&gt;:&lt;digest&gt;&#x60; — the KIND of model, and that family&#39;s own digest over the feature inventory in order and the detector&#39;s geometry parameters. Compare it with the &#x60;shape&#x60; on your model state (GET /v1/risk/state): equal means adopting this value restores masses into the space already running, and different means adopting it REPLANTS the model into the space this value describes. That is what makes a searched shape installable.  A DIFFERENT FAMILY IS NOT ADOPTABLE AT ALL, and that is the one difference the family term makes here: a different geometry in the same family is a replant, and a different family is a refusal naming both — its masses do not describe your model in any space. | [optional] 
**Warmed** | Pointer to **string** | Warmed is how far your own event surface had been folded in when this value was published, RFC 3339. It is part of the address because two models with identical masses reached by different routes disagree about what is left to fold, and one of them will re-teach history the other will not. | [optional] 

## Methods

### NewRiskModelValue

`func NewRiskModelValue() *RiskModelValue`

NewRiskModelValue instantiates a new RiskModelValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskModelValueWithDefaults

`func NewRiskModelValueWithDefaults() *RiskModelValue`

NewRiskModelValueWithDefaults instantiates a new RiskModelValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RiskModelValue) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RiskModelValue) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RiskModelValue) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RiskModelValue) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAt

`func (o *RiskModelValue) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RiskModelValue) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RiskModelValue) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *RiskModelValue) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetLearned

`func (o *RiskModelValue) GetLearned() int32`

GetLearned returns the Learned field if non-nil, zero value otherwise.

### GetLearnedOk

`func (o *RiskModelValue) GetLearnedOk() (*int32, bool)`

GetLearnedOk returns a tuple with the Learned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearned

`func (o *RiskModelValue) SetLearned(v int32)`

SetLearned sets Learned field to given value.

### HasLearned

`func (o *RiskModelValue) HasLearned() bool`

HasLearned returns a boolean if a field has been set.

### GetSequence

`func (o *RiskModelValue) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *RiskModelValue) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *RiskModelValue) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *RiskModelValue) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetShape

`func (o *RiskModelValue) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *RiskModelValue) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *RiskModelValue) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *RiskModelValue) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetWarmed

`func (o *RiskModelValue) GetWarmed() string`

GetWarmed returns the Warmed field if non-nil, zero value otherwise.

### GetWarmedOk

`func (o *RiskModelValue) GetWarmedOk() (*string, bool)`

GetWarmedOk returns a tuple with the Warmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmed

`func (o *RiskModelValue) SetWarmed(v string)`

SetWarmed sets Warmed field to given value.

### HasWarmed

`func (o *RiskModelValue) HasWarmed() bool`

HasWarmed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


