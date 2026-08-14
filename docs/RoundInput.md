# RoundInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the round&#39;s name on the cap table, e.g. \&quot;Seed\&quot;. Required. | [optional] 
**PreMoneyValuation** | Pointer to **float32** | PreMoneyValuation is the valuation the round prices off, before the new money. | [optional] 
**PricePerShare** | Pointer to **float32** | PricePerShare is the per-share price of a priced round. | [optional] 
**RoundType** | Pointer to **string** | RoundType is PRICED, SAFE or CONVERTIBLE_NOTE. Defaults to PRICED. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the cap table&#39;s share class the round issues into. | [optional] 
**TargetAmount** | Pointer to **float32** | TargetAmount is the amount the round is raising, recorded verbatim on the canonical cap table&#39;s rounds.create contract. | [optional] 

## Methods

### NewRoundInput

`func NewRoundInput() *RoundInput`

NewRoundInput instantiates a new RoundInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoundInputWithDefaults

`func NewRoundInputWithDefaults() *RoundInput`

NewRoundInputWithDefaults instantiates a new RoundInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoundInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoundInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoundInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoundInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreMoneyValuation

`func (o *RoundInput) GetPreMoneyValuation() float32`

GetPreMoneyValuation returns the PreMoneyValuation field if non-nil, zero value otherwise.

### GetPreMoneyValuationOk

`func (o *RoundInput) GetPreMoneyValuationOk() (*float32, bool)`

GetPreMoneyValuationOk returns a tuple with the PreMoneyValuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreMoneyValuation

`func (o *RoundInput) SetPreMoneyValuation(v float32)`

SetPreMoneyValuation sets PreMoneyValuation field to given value.

### HasPreMoneyValuation

`func (o *RoundInput) HasPreMoneyValuation() bool`

HasPreMoneyValuation returns a boolean if a field has been set.

### GetPricePerShare

`func (o *RoundInput) GetPricePerShare() float32`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *RoundInput) GetPricePerShareOk() (*float32, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *RoundInput) SetPricePerShare(v float32)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *RoundInput) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetRoundType

`func (o *RoundInput) GetRoundType() string`

GetRoundType returns the RoundType field if non-nil, zero value otherwise.

### GetRoundTypeOk

`func (o *RoundInput) GetRoundTypeOk() (*string, bool)`

GetRoundTypeOk returns a tuple with the RoundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundType

`func (o *RoundInput) SetRoundType(v string)`

SetRoundType sets RoundType field to given value.

### HasRoundType

`func (o *RoundInput) HasRoundType() bool`

HasRoundType returns a boolean if a field has been set.

### GetShareClassId

`func (o *RoundInput) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *RoundInput) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *RoundInput) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *RoundInput) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetTargetAmount

`func (o *RoundInput) GetTargetAmount() float32`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *RoundInput) GetTargetAmountOk() (*float32, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *RoundInput) SetTargetAmount(v float32)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *RoundInput) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


