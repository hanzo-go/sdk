# BillingAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Enforce** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Over** | Pointer to **bool** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**PeriodSpentCents** | Pointer to **int64** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**RateLimitRpm** | Pointer to **int64** |  | [optional] 
**ResetsAt** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**SoftPct** | Pointer to **int64** |  | [optional] 
**Threshold** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TriggeredAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Warn** | Pointer to **bool** |  | [optional] 

## Methods

### NewBillingAlert

`func NewBillingAlert() *BillingAlert`

NewBillingAlert instantiates a new BillingAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAlertWithDefaults

`func NewBillingAlertWithDefaults() *BillingAlert`

NewBillingAlertWithDefaults instantiates a new BillingAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BillingAlert) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingAlert) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingAlert) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingAlert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingAlert) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingAlert) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingAlert) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingAlert) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEnforce

`func (o *BillingAlert) GetEnforce() bool`

GetEnforce returns the Enforce field if non-nil, zero value otherwise.

### GetEnforceOk

`func (o *BillingAlert) GetEnforceOk() (*bool, bool)`

GetEnforceOk returns a tuple with the Enforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce

`func (o *BillingAlert) SetEnforce(v bool)`

SetEnforce sets Enforce field to given value.

### HasEnforce

`func (o *BillingAlert) HasEnforce() bool`

HasEnforce returns a boolean if a field has been set.

### GetId

`func (o *BillingAlert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingAlert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingAlert) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingAlert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOver

`func (o *BillingAlert) GetOver() bool`

GetOver returns the Over field if non-nil, zero value otherwise.

### GetOverOk

`func (o *BillingAlert) GetOverOk() (*bool, bool)`

GetOverOk returns a tuple with the Over field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver

`func (o *BillingAlert) SetOver(v bool)`

SetOver sets Over field to given value.

### HasOver

`func (o *BillingAlert) HasOver() bool`

HasOver returns a boolean if a field has been set.

### GetPeriod

`func (o *BillingAlert) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BillingAlert) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BillingAlert) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *BillingAlert) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPeriodSpentCents

`func (o *BillingAlert) GetPeriodSpentCents() int64`

GetPeriodSpentCents returns the PeriodSpentCents field if non-nil, zero value otherwise.

### GetPeriodSpentCentsOk

`func (o *BillingAlert) GetPeriodSpentCentsOk() (*int64, bool)`

GetPeriodSpentCentsOk returns a tuple with the PeriodSpentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSpentCents

`func (o *BillingAlert) SetPeriodSpentCents(v int64)`

SetPeriodSpentCents sets PeriodSpentCents field to given value.

### HasPeriodSpentCents

`func (o *BillingAlert) HasPeriodSpentCents() bool`

HasPeriodSpentCents returns a boolean if a field has been set.

### GetProject

`func (o *BillingAlert) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BillingAlert) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BillingAlert) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BillingAlert) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRateLimitRpm

`func (o *BillingAlert) GetRateLimitRpm() int64`

GetRateLimitRpm returns the RateLimitRpm field if non-nil, zero value otherwise.

### GetRateLimitRpmOk

`func (o *BillingAlert) GetRateLimitRpmOk() (*int64, bool)`

GetRateLimitRpmOk returns a tuple with the RateLimitRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitRpm

`func (o *BillingAlert) SetRateLimitRpm(v int64)`

SetRateLimitRpm sets RateLimitRpm field to given value.

### HasRateLimitRpm

`func (o *BillingAlert) HasRateLimitRpm() bool`

HasRateLimitRpm returns a boolean if a field has been set.

### GetResetsAt

`func (o *BillingAlert) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *BillingAlert) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *BillingAlert) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *BillingAlert) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetService

`func (o *BillingAlert) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingAlert) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingAlert) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BillingAlert) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSoftPct

`func (o *BillingAlert) GetSoftPct() int64`

GetSoftPct returns the SoftPct field if non-nil, zero value otherwise.

### GetSoftPctOk

`func (o *BillingAlert) GetSoftPctOk() (*int64, bool)`

GetSoftPctOk returns a tuple with the SoftPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftPct

`func (o *BillingAlert) SetSoftPct(v int64)`

SetSoftPct sets SoftPct field to given value.

### HasSoftPct

`func (o *BillingAlert) HasSoftPct() bool`

HasSoftPct returns a boolean if a field has been set.

### GetThreshold

`func (o *BillingAlert) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *BillingAlert) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *BillingAlert) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *BillingAlert) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTitle

`func (o *BillingAlert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillingAlert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillingAlert) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BillingAlert) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *BillingAlert) GetTriggeredAt() string`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *BillingAlert) GetTriggeredAtOk() (*string, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *BillingAlert) SetTriggeredAt(v string)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *BillingAlert) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingAlert) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingAlert) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingAlert) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillingAlert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *BillingAlert) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BillingAlert) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BillingAlert) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BillingAlert) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWarn

`func (o *BillingAlert) GetWarn() bool`

GetWarn returns the Warn field if non-nil, zero value otherwise.

### GetWarnOk

`func (o *BillingAlert) GetWarnOk() (*bool, bool)`

GetWarnOk returns a tuple with the Warn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarn

`func (o *BillingAlert) SetWarn(v bool)`

SetWarn sets Warn field to given value.

### HasWarn

`func (o *BillingAlert) HasWarn() bool`

HasWarn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


