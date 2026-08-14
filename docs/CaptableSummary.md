# CaptableSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByShareClass** | Pointer to [**[]CaptableClassHolding**](CaptableClassHolding.md) | ByShareClass is each share class&#39;s authorized-versus-issued position, in class creation order. | [optional] 
**ByStakeholder** | Pointer to [**[]CaptableHolding**](CaptableHolding.md) | ByStakeholder is each stakeholder&#39;s position, largest holding first. | [optional] 
**Company** | Pointer to [**CaptableSummaryCompany**](CaptableSummaryCompany.md) | Company names the company the cap table is computed for. | [optional] 
**Convertibles** | Pointer to [**CaptableConvertibles**](CaptableConvertibles.md) | Convertibles is the capital on SAFEs and notes that have not converted. | [optional] 
**Rounds** | Pointer to [**CaptableRoundTotals**](CaptableRoundTotals.md) | Rounds is the fundraising rollup. | [optional] 
**Totals** | Pointer to [**CaptableTotals**](CaptableTotals.md) | Totals is the company-wide share count. | [optional] 

## Methods

### NewCaptableSummary

`func NewCaptableSummary() *CaptableSummary`

NewCaptableSummary instantiates a new CaptableSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableSummaryWithDefaults

`func NewCaptableSummaryWithDefaults() *CaptableSummary`

NewCaptableSummaryWithDefaults instantiates a new CaptableSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByShareClass

`func (o *CaptableSummary) GetByShareClass() []CaptableClassHolding`

GetByShareClass returns the ByShareClass field if non-nil, zero value otherwise.

### GetByShareClassOk

`func (o *CaptableSummary) GetByShareClassOk() (*[]CaptableClassHolding, bool)`

GetByShareClassOk returns a tuple with the ByShareClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByShareClass

`func (o *CaptableSummary) SetByShareClass(v []CaptableClassHolding)`

SetByShareClass sets ByShareClass field to given value.

### HasByShareClass

`func (o *CaptableSummary) HasByShareClass() bool`

HasByShareClass returns a boolean if a field has been set.

### GetByStakeholder

`func (o *CaptableSummary) GetByStakeholder() []CaptableHolding`

GetByStakeholder returns the ByStakeholder field if non-nil, zero value otherwise.

### GetByStakeholderOk

`func (o *CaptableSummary) GetByStakeholderOk() (*[]CaptableHolding, bool)`

GetByStakeholderOk returns a tuple with the ByStakeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStakeholder

`func (o *CaptableSummary) SetByStakeholder(v []CaptableHolding)`

SetByStakeholder sets ByStakeholder field to given value.

### HasByStakeholder

`func (o *CaptableSummary) HasByStakeholder() bool`

HasByStakeholder returns a boolean if a field has been set.

### GetCompany

`func (o *CaptableSummary) GetCompany() CaptableSummaryCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CaptableSummary) GetCompanyOk() (*CaptableSummaryCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CaptableSummary) SetCompany(v CaptableSummaryCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CaptableSummary) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetConvertibles

`func (o *CaptableSummary) GetConvertibles() CaptableConvertibles`

GetConvertibles returns the Convertibles field if non-nil, zero value otherwise.

### GetConvertiblesOk

`func (o *CaptableSummary) GetConvertiblesOk() (*CaptableConvertibles, bool)`

GetConvertiblesOk returns a tuple with the Convertibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertibles

`func (o *CaptableSummary) SetConvertibles(v CaptableConvertibles)`

SetConvertibles sets Convertibles field to given value.

### HasConvertibles

`func (o *CaptableSummary) HasConvertibles() bool`

HasConvertibles returns a boolean if a field has been set.

### GetRounds

`func (o *CaptableSummary) GetRounds() CaptableRoundTotals`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *CaptableSummary) GetRoundsOk() (*CaptableRoundTotals, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *CaptableSummary) SetRounds(v CaptableRoundTotals)`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *CaptableSummary) HasRounds() bool`

HasRounds returns a boolean if a field has been set.

### GetTotals

`func (o *CaptableSummary) GetTotals() CaptableTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *CaptableSummary) GetTotalsOk() (*CaptableTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *CaptableSummary) SetTotals(v CaptableTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *CaptableSummary) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


