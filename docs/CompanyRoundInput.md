# CompanyRoundInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the round&#39;s name on the cap table, e.g. \&quot;Seed\&quot;. Required. | [optional] 
**PreMoneyValuation** | Pointer to **float64** | PreMoneyValuation is the valuation the round prices off, before the new money. | [optional] 
**PricePerShare** | Pointer to **float64** | PricePerShare is the per-share price of a priced round. | [optional] 
**RoundType** | Pointer to **string** | RoundType is PRICED, SAFE or CONVERTIBLE_NOTE. Defaults to PRICED. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the cap table&#39;s share class the round issues into. | [optional] 
**TargetAmount** | Pointer to **float64** | TargetAmount is the amount the round is raising, recorded verbatim on the canonical cap table&#39;s rounds.create contract. | [optional] 

## Methods

### NewCompanyRoundInput

`func NewCompanyRoundInput() *CompanyRoundInput`

NewCompanyRoundInput instantiates a new CompanyRoundInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyRoundInputWithDefaults

`func NewCompanyRoundInputWithDefaults() *CompanyRoundInput`

NewCompanyRoundInputWithDefaults instantiates a new CompanyRoundInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompanyRoundInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyRoundInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyRoundInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyRoundInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreMoneyValuation

`func (o *CompanyRoundInput) GetPreMoneyValuation() float64`

GetPreMoneyValuation returns the PreMoneyValuation field if non-nil, zero value otherwise.

### GetPreMoneyValuationOk

`func (o *CompanyRoundInput) GetPreMoneyValuationOk() (*float64, bool)`

GetPreMoneyValuationOk returns a tuple with the PreMoneyValuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreMoneyValuation

`func (o *CompanyRoundInput) SetPreMoneyValuation(v float64)`

SetPreMoneyValuation sets PreMoneyValuation field to given value.

### HasPreMoneyValuation

`func (o *CompanyRoundInput) HasPreMoneyValuation() bool`

HasPreMoneyValuation returns a boolean if a field has been set.

### GetPricePerShare

`func (o *CompanyRoundInput) GetPricePerShare() float64`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CompanyRoundInput) GetPricePerShareOk() (*float64, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CompanyRoundInput) SetPricePerShare(v float64)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CompanyRoundInput) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetRoundType

`func (o *CompanyRoundInput) GetRoundType() string`

GetRoundType returns the RoundType field if non-nil, zero value otherwise.

### GetRoundTypeOk

`func (o *CompanyRoundInput) GetRoundTypeOk() (*string, bool)`

GetRoundTypeOk returns a tuple with the RoundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundType

`func (o *CompanyRoundInput) SetRoundType(v string)`

SetRoundType sets RoundType field to given value.

### HasRoundType

`func (o *CompanyRoundInput) HasRoundType() bool`

HasRoundType returns a boolean if a field has been set.

### GetShareClassId

`func (o *CompanyRoundInput) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CompanyRoundInput) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CompanyRoundInput) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CompanyRoundInput) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetTargetAmount

`func (o *CompanyRoundInput) GetTargetAmount() float64`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *CompanyRoundInput) GetTargetAmountOk() (*float64, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *CompanyRoundInput) SetTargetAmount(v float64)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *CompanyRoundInput) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


