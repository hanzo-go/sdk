# BillingSpendAlertCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to **int64** | Spend cap ceiling in USD cents. | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "usd"]
**Project** | Pointer to **string** | Scope project axis; \&quot;\&quot; or \&quot;default\&quot; &#x3D; org-wide. | [optional] 
**Service** | Pointer to **string** | Scope service axis; \&quot;\&quot; &#x3D; all services. | [optional] 
**Enforce** | Pointer to **bool** | true &#x3D; HARD cap; false &#x3D; soft warn. | [optional] [default to false]
**SoftPct** | Pointer to **int32** | Soft-warn percent; 0 defaults to 80. | [optional] 
**RateLimitRpm** | Pointer to **int32** | Per-minute request ceiling; 0 &#x3D; none. | [optional] 

## Methods

### NewBillingSpendAlertCreate

`func NewBillingSpendAlertCreate() *BillingSpendAlertCreate`

NewBillingSpendAlertCreate instantiates a new BillingSpendAlertCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSpendAlertCreateWithDefaults

`func NewBillingSpendAlertCreateWithDefaults() *BillingSpendAlertCreate`

NewBillingSpendAlertCreateWithDefaults instantiates a new BillingSpendAlertCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BillingSpendAlertCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillingSpendAlertCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillingSpendAlertCreate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BillingSpendAlertCreate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetThreshold

`func (o *BillingSpendAlertCreate) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *BillingSpendAlertCreate) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *BillingSpendAlertCreate) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *BillingSpendAlertCreate) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingSpendAlertCreate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingSpendAlertCreate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingSpendAlertCreate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingSpendAlertCreate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetProject

`func (o *BillingSpendAlertCreate) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BillingSpendAlertCreate) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BillingSpendAlertCreate) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BillingSpendAlertCreate) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetService

`func (o *BillingSpendAlertCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingSpendAlertCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingSpendAlertCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BillingSpendAlertCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetEnforce

`func (o *BillingSpendAlertCreate) GetEnforce() bool`

GetEnforce returns the Enforce field if non-nil, zero value otherwise.

### GetEnforceOk

`func (o *BillingSpendAlertCreate) GetEnforceOk() (*bool, bool)`

GetEnforceOk returns a tuple with the Enforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce

`func (o *BillingSpendAlertCreate) SetEnforce(v bool)`

SetEnforce sets Enforce field to given value.

### HasEnforce

`func (o *BillingSpendAlertCreate) HasEnforce() bool`

HasEnforce returns a boolean if a field has been set.

### GetSoftPct

`func (o *BillingSpendAlertCreate) GetSoftPct() int32`

GetSoftPct returns the SoftPct field if non-nil, zero value otherwise.

### GetSoftPctOk

`func (o *BillingSpendAlertCreate) GetSoftPctOk() (*int32, bool)`

GetSoftPctOk returns a tuple with the SoftPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftPct

`func (o *BillingSpendAlertCreate) SetSoftPct(v int32)`

SetSoftPct sets SoftPct field to given value.

### HasSoftPct

`func (o *BillingSpendAlertCreate) HasSoftPct() bool`

HasSoftPct returns a boolean if a field has been set.

### GetRateLimitRpm

`func (o *BillingSpendAlertCreate) GetRateLimitRpm() int32`

GetRateLimitRpm returns the RateLimitRpm field if non-nil, zero value otherwise.

### GetRateLimitRpmOk

`func (o *BillingSpendAlertCreate) GetRateLimitRpmOk() (*int32, bool)`

GetRateLimitRpmOk returns a tuple with the RateLimitRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitRpm

`func (o *BillingSpendAlertCreate) SetRateLimitRpm(v int32)`

SetRateLimitRpm sets RateLimitRpm field to given value.

### HasRateLimitRpm

`func (o *BillingSpendAlertCreate) HasRateLimitRpm() bool`

HasRateLimitRpm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


