# PeriodEarningView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommissionCents** | Pointer to **int64** | CommissionCents is what the caller earned that period, in cents: the sum over each referred org and upline level of margin × that level&#39;s rate. Always ≤ marginCents, by construction. | [optional] 
**MarginCents** | Pointer to **int64** | MarginCents is the margin Hanzo earned in that period on the spend of every org the caller referred, in cents — the base commission is a rate OF. It is the aggregate base, never any one customer&#39;s bill. | [optional] 
**Period** | Pointer to **string** | Period is the accrual bucket: the UTC year-month, \&quot;YYYY-MM\&quot;. Commission is latched at most once per referred org per period, so one row is one month. | [optional] 

## Methods

### NewPeriodEarningView

`func NewPeriodEarningView() *PeriodEarningView`

NewPeriodEarningView instantiates a new PeriodEarningView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodEarningViewWithDefaults

`func NewPeriodEarningViewWithDefaults() *PeriodEarningView`

NewPeriodEarningViewWithDefaults instantiates a new PeriodEarningView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommissionCents

`func (o *PeriodEarningView) GetCommissionCents() int64`

GetCommissionCents returns the CommissionCents field if non-nil, zero value otherwise.

### GetCommissionCentsOk

`func (o *PeriodEarningView) GetCommissionCentsOk() (*int64, bool)`

GetCommissionCentsOk returns a tuple with the CommissionCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionCents

`func (o *PeriodEarningView) SetCommissionCents(v int64)`

SetCommissionCents sets CommissionCents field to given value.

### HasCommissionCents

`func (o *PeriodEarningView) HasCommissionCents() bool`

HasCommissionCents returns a boolean if a field has been set.

### GetMarginCents

`func (o *PeriodEarningView) GetMarginCents() int64`

GetMarginCents returns the MarginCents field if non-nil, zero value otherwise.

### GetMarginCentsOk

`func (o *PeriodEarningView) GetMarginCentsOk() (*int64, bool)`

GetMarginCentsOk returns a tuple with the MarginCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginCents

`func (o *PeriodEarningView) SetMarginCents(v int64)`

SetMarginCents sets MarginCents field to given value.

### HasMarginCents

`func (o *PeriodEarningView) HasMarginCents() bool`

HasMarginCents returns a boolean if a field has been set.

### GetPeriod

`func (o *PeriodEarningView) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *PeriodEarningView) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *PeriodEarningView) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *PeriodEarningView) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


