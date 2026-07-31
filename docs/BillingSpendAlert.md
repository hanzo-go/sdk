# BillingSpendAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** | The org the cap is scoped to — server-pinned to the validated X-Org-Id, never a client field. | 
**Title** | Pointer to **string** |  | [optional] 
**Threshold** | **int64** | The scope&#39;s spend cap / alert ceiling in USD cents (Money). 0 &#x3D; no cap (a rate-limit-only row). | 
**Currency** | **string** |  | [default to "usd"]
**Project** | **string** | Scope project axis; \&quot;\&quot; &#x3D; wildcard (the org-wide default row). | 
**Service** | **string** | Scope service axis (server-derived route/provider); \&quot;\&quot; &#x3D; wildcard (all services). | 
**Enforce** | **bool** | false (default) &#x3D; SOFT (warn only, never blocks); true &#x3D; HARD cap (402 on breach). | 
**SoftPct** | **int32** | Soft-warn threshold as a percent of &#x60;threshold&#x60; (unset/0 defaults to 80). | 
**RateLimitRpm** | **int32** | Requests-per-minute ceiling for the scope (429 when exceeded). 0 &#x3D; none. | 
**Period** | Pointer to **string** | The UTC calendar month the cap is measured over, e.g. \&quot;2026-07\&quot; (Schedule window; resets on the 1st). | [optional] 
**ResetsAt** | Pointer to **time.Time** | Midnight UTC on the first of the next month — when period spend resets to 0. | [optional] 
**PeriodSpentCents** | Pointer to **int64** | DERIVED spend for the scope in &#x60;period&#x60;, in USD cents, summed from the deduped usage ledger. Absent if the aggregation errored. | [optional] 
**Over** | Pointer to **bool** | DERIVED — true iff a further billable request would breach the cap (the SAME boundary &#x60;enforce&#x60; uses, so it never drifts from the verdict). | [optional] 
**Warn** | Pointer to **bool** | DERIVED — true iff period spend has reached the soft-warn threshold (&#x60;softPct&#x60; of &#x60;threshold&#x60;). | [optional] 
**TriggeredAt** | Pointer to **string** | ISO timestamp the alert last fired; empty if never. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBillingSpendAlert

`func NewBillingSpendAlert(id string, userId string, threshold int64, currency string, project string, service string, enforce bool, softPct int32, rateLimitRpm int32, ) *BillingSpendAlert`

NewBillingSpendAlert instantiates a new BillingSpendAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSpendAlertWithDefaults

`func NewBillingSpendAlertWithDefaults() *BillingSpendAlert`

NewBillingSpendAlertWithDefaults instantiates a new BillingSpendAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingSpendAlert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSpendAlert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSpendAlert) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *BillingSpendAlert) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BillingSpendAlert) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BillingSpendAlert) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTitle

`func (o *BillingSpendAlert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillingSpendAlert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillingSpendAlert) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BillingSpendAlert) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetThreshold

`func (o *BillingSpendAlert) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *BillingSpendAlert) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *BillingSpendAlert) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.


### GetCurrency

`func (o *BillingSpendAlert) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingSpendAlert) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingSpendAlert) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetProject

`func (o *BillingSpendAlert) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BillingSpendAlert) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BillingSpendAlert) SetProject(v string)`

SetProject sets Project field to given value.


### GetService

`func (o *BillingSpendAlert) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingSpendAlert) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingSpendAlert) SetService(v string)`

SetService sets Service field to given value.


### GetEnforce

`func (o *BillingSpendAlert) GetEnforce() bool`

GetEnforce returns the Enforce field if non-nil, zero value otherwise.

### GetEnforceOk

`func (o *BillingSpendAlert) GetEnforceOk() (*bool, bool)`

GetEnforceOk returns a tuple with the Enforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce

`func (o *BillingSpendAlert) SetEnforce(v bool)`

SetEnforce sets Enforce field to given value.


### GetSoftPct

`func (o *BillingSpendAlert) GetSoftPct() int32`

GetSoftPct returns the SoftPct field if non-nil, zero value otherwise.

### GetSoftPctOk

`func (o *BillingSpendAlert) GetSoftPctOk() (*int32, bool)`

GetSoftPctOk returns a tuple with the SoftPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftPct

`func (o *BillingSpendAlert) SetSoftPct(v int32)`

SetSoftPct sets SoftPct field to given value.


### GetRateLimitRpm

`func (o *BillingSpendAlert) GetRateLimitRpm() int32`

GetRateLimitRpm returns the RateLimitRpm field if non-nil, zero value otherwise.

### GetRateLimitRpmOk

`func (o *BillingSpendAlert) GetRateLimitRpmOk() (*int32, bool)`

GetRateLimitRpmOk returns a tuple with the RateLimitRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitRpm

`func (o *BillingSpendAlert) SetRateLimitRpm(v int32)`

SetRateLimitRpm sets RateLimitRpm field to given value.


### GetPeriod

`func (o *BillingSpendAlert) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BillingSpendAlert) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BillingSpendAlert) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *BillingSpendAlert) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetResetsAt

`func (o *BillingSpendAlert) GetResetsAt() time.Time`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *BillingSpendAlert) GetResetsAtOk() (*time.Time, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *BillingSpendAlert) SetResetsAt(v time.Time)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *BillingSpendAlert) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetPeriodSpentCents

`func (o *BillingSpendAlert) GetPeriodSpentCents() int64`

GetPeriodSpentCents returns the PeriodSpentCents field if non-nil, zero value otherwise.

### GetPeriodSpentCentsOk

`func (o *BillingSpendAlert) GetPeriodSpentCentsOk() (*int64, bool)`

GetPeriodSpentCentsOk returns a tuple with the PeriodSpentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSpentCents

`func (o *BillingSpendAlert) SetPeriodSpentCents(v int64)`

SetPeriodSpentCents sets PeriodSpentCents field to given value.

### HasPeriodSpentCents

`func (o *BillingSpendAlert) HasPeriodSpentCents() bool`

HasPeriodSpentCents returns a boolean if a field has been set.

### GetOver

`func (o *BillingSpendAlert) GetOver() bool`

GetOver returns the Over field if non-nil, zero value otherwise.

### GetOverOk

`func (o *BillingSpendAlert) GetOverOk() (*bool, bool)`

GetOverOk returns a tuple with the Over field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver

`func (o *BillingSpendAlert) SetOver(v bool)`

SetOver sets Over field to given value.

### HasOver

`func (o *BillingSpendAlert) HasOver() bool`

HasOver returns a boolean if a field has been set.

### GetWarn

`func (o *BillingSpendAlert) GetWarn() bool`

GetWarn returns the Warn field if non-nil, zero value otherwise.

### GetWarnOk

`func (o *BillingSpendAlert) GetWarnOk() (*bool, bool)`

GetWarnOk returns a tuple with the Warn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarn

`func (o *BillingSpendAlert) SetWarn(v bool)`

SetWarn sets Warn field to given value.

### HasWarn

`func (o *BillingSpendAlert) HasWarn() bool`

HasWarn returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *BillingSpendAlert) GetTriggeredAt() string`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *BillingSpendAlert) GetTriggeredAtOk() (*string, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *BillingSpendAlert) SetTriggeredAt(v string)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *BillingSpendAlert) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingSpendAlert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingSpendAlert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingSpendAlert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingSpendAlert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingSpendAlert) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingSpendAlert) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingSpendAlert) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillingSpendAlert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


