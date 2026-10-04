# AffiliateAffiliateEarnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int64** | AccruedCents is lifetime commission accrued, in cents. | [optional] 
**ByPeriod** | Pointer to [**[]AffiliatePeriodEarningView**](AffiliatePeriodEarningView.md) | ByPeriod is the per-period ledger: the margin earned against and the commission taken from it. | [optional] 
**ByReferredOrg** | Pointer to [**[]AffiliateOrgEarningView**](AffiliateOrgEarningView.md) | ByReferredOrg is each referral&#39;s aggregate contribution — the affiliate&#39;s OWN share, never the referred org&#39;s spend. | [optional] 
**IsAffiliate** | Pointer to **bool** | IsAffiliate says whether the caller org has an affiliate record. On false it is the ONLY field present — there is no ledger to report, and the zeros you might expect are absent rather than reported as earnings of nothing. | [optional] 
**MarginBps** | Pointer to **int64** | MarginBps is the platform gross-margin fraction commission is a rate OF. | [optional] 
**PaidCents** | Pointer to **int64** | PaidCents is lifetime commission already paid out, in cents. | [optional] 
**PendingCents** | Pointer to **int64** | PendingCents is accrued minus paid — what the platform still owes. | [optional] 

## Methods

### NewAffiliateAffiliateEarnings

`func NewAffiliateAffiliateEarnings() *AffiliateAffiliateEarnings`

NewAffiliateAffiliateEarnings instantiates a new AffiliateAffiliateEarnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateAffiliateEarningsWithDefaults

`func NewAffiliateAffiliateEarningsWithDefaults() *AffiliateAffiliateEarnings`

NewAffiliateAffiliateEarningsWithDefaults instantiates a new AffiliateAffiliateEarnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *AffiliateAffiliateEarnings) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AffiliateAffiliateEarnings) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AffiliateAffiliateEarnings) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AffiliateAffiliateEarnings) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetByPeriod

`func (o *AffiliateAffiliateEarnings) GetByPeriod() []AffiliatePeriodEarningView`

GetByPeriod returns the ByPeriod field if non-nil, zero value otherwise.

### GetByPeriodOk

`func (o *AffiliateAffiliateEarnings) GetByPeriodOk() (*[]AffiliatePeriodEarningView, bool)`

GetByPeriodOk returns a tuple with the ByPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByPeriod

`func (o *AffiliateAffiliateEarnings) SetByPeriod(v []AffiliatePeriodEarningView)`

SetByPeriod sets ByPeriod field to given value.

### HasByPeriod

`func (o *AffiliateAffiliateEarnings) HasByPeriod() bool`

HasByPeriod returns a boolean if a field has been set.

### GetByReferredOrg

`func (o *AffiliateAffiliateEarnings) GetByReferredOrg() []AffiliateOrgEarningView`

GetByReferredOrg returns the ByReferredOrg field if non-nil, zero value otherwise.

### GetByReferredOrgOk

`func (o *AffiliateAffiliateEarnings) GetByReferredOrgOk() (*[]AffiliateOrgEarningView, bool)`

GetByReferredOrgOk returns a tuple with the ByReferredOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByReferredOrg

`func (o *AffiliateAffiliateEarnings) SetByReferredOrg(v []AffiliateOrgEarningView)`

SetByReferredOrg sets ByReferredOrg field to given value.

### HasByReferredOrg

`func (o *AffiliateAffiliateEarnings) HasByReferredOrg() bool`

HasByReferredOrg returns a boolean if a field has been set.

### GetIsAffiliate

`func (o *AffiliateAffiliateEarnings) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *AffiliateAffiliateEarnings) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *AffiliateAffiliateEarnings) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *AffiliateAffiliateEarnings) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetMarginBps

`func (o *AffiliateAffiliateEarnings) GetMarginBps() int64`

GetMarginBps returns the MarginBps field if non-nil, zero value otherwise.

### GetMarginBpsOk

`func (o *AffiliateAffiliateEarnings) GetMarginBpsOk() (*int64, bool)`

GetMarginBpsOk returns a tuple with the MarginBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBps

`func (o *AffiliateAffiliateEarnings) SetMarginBps(v int64)`

SetMarginBps sets MarginBps field to given value.

### HasMarginBps

`func (o *AffiliateAffiliateEarnings) HasMarginBps() bool`

HasMarginBps returns a boolean if a field has been set.

### GetPaidCents

`func (o *AffiliateAffiliateEarnings) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AffiliateAffiliateEarnings) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AffiliateAffiliateEarnings) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AffiliateAffiliateEarnings) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AffiliateAffiliateEarnings) GetPendingCents() int64`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AffiliateAffiliateEarnings) GetPendingCentsOk() (*int64, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AffiliateAffiliateEarnings) SetPendingCents(v int64)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AffiliateAffiliateEarnings) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


