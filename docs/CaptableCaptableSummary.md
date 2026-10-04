# CaptableCaptableSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByShareClass** | Pointer to [**[]CaptableCaptableClassHolding**](CaptableCaptableClassHolding.md) | ByShareClass is each share class&#39;s authorized-versus-issued position, in class creation order. | [optional] 
**ByStakeholder** | Pointer to [**[]CaptableCaptableHolding**](CaptableCaptableHolding.md) | ByStakeholder is each stakeholder&#39;s position, largest holding first. | [optional] 
**Company** | Pointer to [**CaptableCaptableSummaryCompany**](CaptableCaptableSummaryCompany.md) | Company names the company the cap table is computed for. | [optional] 
**Convertibles** | Pointer to [**CaptableCaptableConvertibles**](CaptableCaptableConvertibles.md) | Convertibles is the capital on SAFEs and notes that have not converted. | [optional] 
**Rounds** | Pointer to [**CaptableCaptableRoundTotals**](CaptableCaptableRoundTotals.md) | Rounds is the fundraising rollup. | [optional] 
**Totals** | Pointer to [**CaptableCaptableTotals**](CaptableCaptableTotals.md) | Totals is the company-wide share count. | [optional] 

## Methods

### NewCaptableCaptableSummary

`func NewCaptableCaptableSummary() *CaptableCaptableSummary`

NewCaptableCaptableSummary instantiates a new CaptableCaptableSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableSummaryWithDefaults

`func NewCaptableCaptableSummaryWithDefaults() *CaptableCaptableSummary`

NewCaptableCaptableSummaryWithDefaults instantiates a new CaptableCaptableSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByShareClass

`func (o *CaptableCaptableSummary) GetByShareClass() []CaptableCaptableClassHolding`

GetByShareClass returns the ByShareClass field if non-nil, zero value otherwise.

### GetByShareClassOk

`func (o *CaptableCaptableSummary) GetByShareClassOk() (*[]CaptableCaptableClassHolding, bool)`

GetByShareClassOk returns a tuple with the ByShareClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByShareClass

`func (o *CaptableCaptableSummary) SetByShareClass(v []CaptableCaptableClassHolding)`

SetByShareClass sets ByShareClass field to given value.

### HasByShareClass

`func (o *CaptableCaptableSummary) HasByShareClass() bool`

HasByShareClass returns a boolean if a field has been set.

### GetByStakeholder

`func (o *CaptableCaptableSummary) GetByStakeholder() []CaptableCaptableHolding`

GetByStakeholder returns the ByStakeholder field if non-nil, zero value otherwise.

### GetByStakeholderOk

`func (o *CaptableCaptableSummary) GetByStakeholderOk() (*[]CaptableCaptableHolding, bool)`

GetByStakeholderOk returns a tuple with the ByStakeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStakeholder

`func (o *CaptableCaptableSummary) SetByStakeholder(v []CaptableCaptableHolding)`

SetByStakeholder sets ByStakeholder field to given value.

### HasByStakeholder

`func (o *CaptableCaptableSummary) HasByStakeholder() bool`

HasByStakeholder returns a boolean if a field has been set.

### GetCompany

`func (o *CaptableCaptableSummary) GetCompany() CaptableCaptableSummaryCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CaptableCaptableSummary) GetCompanyOk() (*CaptableCaptableSummaryCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CaptableCaptableSummary) SetCompany(v CaptableCaptableSummaryCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CaptableCaptableSummary) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetConvertibles

`func (o *CaptableCaptableSummary) GetConvertibles() CaptableCaptableConvertibles`

GetConvertibles returns the Convertibles field if non-nil, zero value otherwise.

### GetConvertiblesOk

`func (o *CaptableCaptableSummary) GetConvertiblesOk() (*CaptableCaptableConvertibles, bool)`

GetConvertiblesOk returns a tuple with the Convertibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertibles

`func (o *CaptableCaptableSummary) SetConvertibles(v CaptableCaptableConvertibles)`

SetConvertibles sets Convertibles field to given value.

### HasConvertibles

`func (o *CaptableCaptableSummary) HasConvertibles() bool`

HasConvertibles returns a boolean if a field has been set.

### GetRounds

`func (o *CaptableCaptableSummary) GetRounds() CaptableCaptableRoundTotals`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *CaptableCaptableSummary) GetRoundsOk() (*CaptableCaptableRoundTotals, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *CaptableCaptableSummary) SetRounds(v CaptableCaptableRoundTotals)`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *CaptableCaptableSummary) HasRounds() bool`

HasRounds returns a boolean if a field has been set.

### GetTotals

`func (o *CaptableCaptableSummary) GetTotals() CaptableCaptableTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *CaptableCaptableSummary) GetTotalsOk() (*CaptableCaptableTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *CaptableCaptableSummary) SetTotals(v CaptableCaptableTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *CaptableCaptableSummary) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


