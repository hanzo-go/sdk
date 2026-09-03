# Alert

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

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Alert) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Alert) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Alert) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Alert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Alert) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Alert) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Alert) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Alert) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEnforce

`func (o *Alert) GetEnforce() bool`

GetEnforce returns the Enforce field if non-nil, zero value otherwise.

### GetEnforceOk

`func (o *Alert) GetEnforceOk() (*bool, bool)`

GetEnforceOk returns a tuple with the Enforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce

`func (o *Alert) SetEnforce(v bool)`

SetEnforce sets Enforce field to given value.

### HasEnforce

`func (o *Alert) HasEnforce() bool`

HasEnforce returns a boolean if a field has been set.

### GetId

`func (o *Alert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Alert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOver

`func (o *Alert) GetOver() bool`

GetOver returns the Over field if non-nil, zero value otherwise.

### GetOverOk

`func (o *Alert) GetOverOk() (*bool, bool)`

GetOverOk returns a tuple with the Over field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver

`func (o *Alert) SetOver(v bool)`

SetOver sets Over field to given value.

### HasOver

`func (o *Alert) HasOver() bool`

HasOver returns a boolean if a field has been set.

### GetPeriod

`func (o *Alert) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Alert) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Alert) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Alert) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPeriodSpentCents

`func (o *Alert) GetPeriodSpentCents() int64`

GetPeriodSpentCents returns the PeriodSpentCents field if non-nil, zero value otherwise.

### GetPeriodSpentCentsOk

`func (o *Alert) GetPeriodSpentCentsOk() (*int64, bool)`

GetPeriodSpentCentsOk returns a tuple with the PeriodSpentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSpentCents

`func (o *Alert) SetPeriodSpentCents(v int64)`

SetPeriodSpentCents sets PeriodSpentCents field to given value.

### HasPeriodSpentCents

`func (o *Alert) HasPeriodSpentCents() bool`

HasPeriodSpentCents returns a boolean if a field has been set.

### GetProject

`func (o *Alert) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Alert) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Alert) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Alert) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRateLimitRpm

`func (o *Alert) GetRateLimitRpm() int64`

GetRateLimitRpm returns the RateLimitRpm field if non-nil, zero value otherwise.

### GetRateLimitRpmOk

`func (o *Alert) GetRateLimitRpmOk() (*int64, bool)`

GetRateLimitRpmOk returns a tuple with the RateLimitRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitRpm

`func (o *Alert) SetRateLimitRpm(v int64)`

SetRateLimitRpm sets RateLimitRpm field to given value.

### HasRateLimitRpm

`func (o *Alert) HasRateLimitRpm() bool`

HasRateLimitRpm returns a boolean if a field has been set.

### GetResetsAt

`func (o *Alert) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *Alert) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *Alert) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *Alert) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetService

`func (o *Alert) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Alert) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Alert) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *Alert) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSoftPct

`func (o *Alert) GetSoftPct() int64`

GetSoftPct returns the SoftPct field if non-nil, zero value otherwise.

### GetSoftPctOk

`func (o *Alert) GetSoftPctOk() (*int64, bool)`

GetSoftPctOk returns a tuple with the SoftPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftPct

`func (o *Alert) SetSoftPct(v int64)`

SetSoftPct sets SoftPct field to given value.

### HasSoftPct

`func (o *Alert) HasSoftPct() bool`

HasSoftPct returns a boolean if a field has been set.

### GetThreshold

`func (o *Alert) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Alert) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Alert) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Alert) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTitle

`func (o *Alert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Alert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Alert) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Alert) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *Alert) GetTriggeredAt() string`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *Alert) GetTriggeredAtOk() (*string, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *Alert) SetTriggeredAt(v string)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *Alert) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Alert) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Alert) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Alert) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Alert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *Alert) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Alert) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Alert) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Alert) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWarn

`func (o *Alert) GetWarn() bool`

GetWarn returns the Warn field if non-nil, zero value otherwise.

### GetWarnOk

`func (o *Alert) GetWarnOk() (*bool, bool)`

GetWarnOk returns a tuple with the Warn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarn

`func (o *Alert) SetWarn(v bool)`

SetWarn sets Warn field to given value.

### HasWarn

`func (o *Alert) HasWarn() bool`

HasWarn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


