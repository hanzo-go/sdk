# CloudCaptableSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByShareClass** | Pointer to [**[]CloudCaptableClassHolding**](CloudCaptableClassHolding.md) | ByShareClass is each share class&#39;s authorized-versus-issued position, in class creation order. | [optional] 
**ByStakeholder** | Pointer to [**[]CloudCaptableHolding**](CloudCaptableHolding.md) | ByStakeholder is each stakeholder&#39;s position, largest holding first. | [optional] 
**Company** | Pointer to [**CloudCaptableSummaryCompany**](CloudCaptableSummaryCompany.md) | Company names the company the cap table is computed for. | [optional] 
**Convertibles** | Pointer to [**CloudCaptableConvertibles**](CloudCaptableConvertibles.md) | Convertibles is the capital on SAFEs and notes that have not converted. | [optional] 
**Rounds** | Pointer to [**CloudCaptableRoundTotals**](CloudCaptableRoundTotals.md) | Rounds is the fundraising rollup. | [optional] 
**Totals** | Pointer to [**CloudCaptableTotals**](CloudCaptableTotals.md) | Totals is the company-wide share count. | [optional] 

## Methods

### NewCloudCaptableSummary

`func NewCloudCaptableSummary() *CloudCaptableSummary`

NewCloudCaptableSummary instantiates a new CloudCaptableSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableSummaryWithDefaults

`func NewCloudCaptableSummaryWithDefaults() *CloudCaptableSummary`

NewCloudCaptableSummaryWithDefaults instantiates a new CloudCaptableSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByShareClass

`func (o *CloudCaptableSummary) GetByShareClass() []CloudCaptableClassHolding`

GetByShareClass returns the ByShareClass field if non-nil, zero value otherwise.

### GetByShareClassOk

`func (o *CloudCaptableSummary) GetByShareClassOk() (*[]CloudCaptableClassHolding, bool)`

GetByShareClassOk returns a tuple with the ByShareClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByShareClass

`func (o *CloudCaptableSummary) SetByShareClass(v []CloudCaptableClassHolding)`

SetByShareClass sets ByShareClass field to given value.

### HasByShareClass

`func (o *CloudCaptableSummary) HasByShareClass() bool`

HasByShareClass returns a boolean if a field has been set.

### GetByStakeholder

`func (o *CloudCaptableSummary) GetByStakeholder() []CloudCaptableHolding`

GetByStakeholder returns the ByStakeholder field if non-nil, zero value otherwise.

### GetByStakeholderOk

`func (o *CloudCaptableSummary) GetByStakeholderOk() (*[]CloudCaptableHolding, bool)`

GetByStakeholderOk returns a tuple with the ByStakeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStakeholder

`func (o *CloudCaptableSummary) SetByStakeholder(v []CloudCaptableHolding)`

SetByStakeholder sets ByStakeholder field to given value.

### HasByStakeholder

`func (o *CloudCaptableSummary) HasByStakeholder() bool`

HasByStakeholder returns a boolean if a field has been set.

### GetCompany

`func (o *CloudCaptableSummary) GetCompany() CloudCaptableSummaryCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CloudCaptableSummary) GetCompanyOk() (*CloudCaptableSummaryCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CloudCaptableSummary) SetCompany(v CloudCaptableSummaryCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CloudCaptableSummary) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetConvertibles

`func (o *CloudCaptableSummary) GetConvertibles() CloudCaptableConvertibles`

GetConvertibles returns the Convertibles field if non-nil, zero value otherwise.

### GetConvertiblesOk

`func (o *CloudCaptableSummary) GetConvertiblesOk() (*CloudCaptableConvertibles, bool)`

GetConvertiblesOk returns a tuple with the Convertibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertibles

`func (o *CloudCaptableSummary) SetConvertibles(v CloudCaptableConvertibles)`

SetConvertibles sets Convertibles field to given value.

### HasConvertibles

`func (o *CloudCaptableSummary) HasConvertibles() bool`

HasConvertibles returns a boolean if a field has been set.

### GetRounds

`func (o *CloudCaptableSummary) GetRounds() CloudCaptableRoundTotals`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *CloudCaptableSummary) GetRoundsOk() (*CloudCaptableRoundTotals, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *CloudCaptableSummary) SetRounds(v CloudCaptableRoundTotals)`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *CloudCaptableSummary) HasRounds() bool`

HasRounds returns a boolean if a field has been set.

### GetTotals

`func (o *CloudCaptableSummary) GetTotals() CloudCaptableTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *CloudCaptableSummary) GetTotalsOk() (*CloudCaptableTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *CloudCaptableSummary) SetTotals(v CloudCaptableTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *CloudCaptableSummary) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


