# BillingAlertSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Enforce** | Pointer to **bool** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**RateLimitRpm** | Pointer to **int64** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**SoftPct** | Pointer to **int64** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingAlertSpec

`func NewBillingAlertSpec() *BillingAlertSpec`

NewBillingAlertSpec instantiates a new BillingAlertSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAlertSpecWithDefaults

`func NewBillingAlertSpecWithDefaults() *BillingAlertSpec`

NewBillingAlertSpecWithDefaults instantiates a new BillingAlertSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *BillingAlertSpec) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingAlertSpec) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingAlertSpec) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingAlertSpec) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEnforce

`func (o *BillingAlertSpec) GetEnforce() bool`

GetEnforce returns the Enforce field if non-nil, zero value otherwise.

### GetEnforceOk

`func (o *BillingAlertSpec) GetEnforceOk() (*bool, bool)`

GetEnforceOk returns a tuple with the Enforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce

`func (o *BillingAlertSpec) SetEnforce(v bool)`

SetEnforce sets Enforce field to given value.

### HasEnforce

`func (o *BillingAlertSpec) HasEnforce() bool`

HasEnforce returns a boolean if a field has been set.

### GetProject

`func (o *BillingAlertSpec) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BillingAlertSpec) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BillingAlertSpec) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BillingAlertSpec) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRateLimitRpm

`func (o *BillingAlertSpec) GetRateLimitRpm() int64`

GetRateLimitRpm returns the RateLimitRpm field if non-nil, zero value otherwise.

### GetRateLimitRpmOk

`func (o *BillingAlertSpec) GetRateLimitRpmOk() (*int64, bool)`

GetRateLimitRpmOk returns a tuple with the RateLimitRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitRpm

`func (o *BillingAlertSpec) SetRateLimitRpm(v int64)`

SetRateLimitRpm sets RateLimitRpm field to given value.

### HasRateLimitRpm

`func (o *BillingAlertSpec) HasRateLimitRpm() bool`

HasRateLimitRpm returns a boolean if a field has been set.

### GetService

`func (o *BillingAlertSpec) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingAlertSpec) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingAlertSpec) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BillingAlertSpec) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSoftPct

`func (o *BillingAlertSpec) GetSoftPct() int64`

GetSoftPct returns the SoftPct field if non-nil, zero value otherwise.

### GetSoftPctOk

`func (o *BillingAlertSpec) GetSoftPctOk() (*int64, bool)`

GetSoftPctOk returns a tuple with the SoftPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftPct

`func (o *BillingAlertSpec) SetSoftPct(v int64)`

SetSoftPct sets SoftPct field to given value.

### HasSoftPct

`func (o *BillingAlertSpec) HasSoftPct() bool`

HasSoftPct returns a boolean if a field has been set.

### GetSubject

`func (o *BillingAlertSpec) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *BillingAlertSpec) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *BillingAlertSpec) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *BillingAlertSpec) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetThreshold

`func (o *BillingAlertSpec) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *BillingAlertSpec) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *BillingAlertSpec) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *BillingAlertSpec) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTitle

`func (o *BillingAlertSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillingAlertSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillingAlertSpec) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BillingAlertSpec) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


