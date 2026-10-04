# BillingAlertPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enforce** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**RateLimitRpm** | Pointer to **int64** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**SoftPct** | Pointer to **int64** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingAlertPatch

`func NewBillingAlertPatch() *BillingAlertPatch`

NewBillingAlertPatch instantiates a new BillingAlertPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAlertPatchWithDefaults

`func NewBillingAlertPatchWithDefaults() *BillingAlertPatch`

NewBillingAlertPatchWithDefaults instantiates a new BillingAlertPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforce

`func (o *BillingAlertPatch) GetEnforce() bool`

GetEnforce returns the Enforce field if non-nil, zero value otherwise.

### GetEnforceOk

`func (o *BillingAlertPatch) GetEnforceOk() (*bool, bool)`

GetEnforceOk returns a tuple with the Enforce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce

`func (o *BillingAlertPatch) SetEnforce(v bool)`

SetEnforce sets Enforce field to given value.

### HasEnforce

`func (o *BillingAlertPatch) HasEnforce() bool`

HasEnforce returns a boolean if a field has been set.

### GetId

`func (o *BillingAlertPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingAlertPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingAlertPatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingAlertPatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *BillingAlertPatch) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BillingAlertPatch) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BillingAlertPatch) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BillingAlertPatch) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRateLimitRpm

`func (o *BillingAlertPatch) GetRateLimitRpm() int64`

GetRateLimitRpm returns the RateLimitRpm field if non-nil, zero value otherwise.

### GetRateLimitRpmOk

`func (o *BillingAlertPatch) GetRateLimitRpmOk() (*int64, bool)`

GetRateLimitRpmOk returns a tuple with the RateLimitRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitRpm

`func (o *BillingAlertPatch) SetRateLimitRpm(v int64)`

SetRateLimitRpm sets RateLimitRpm field to given value.

### HasRateLimitRpm

`func (o *BillingAlertPatch) HasRateLimitRpm() bool`

HasRateLimitRpm returns a boolean if a field has been set.

### GetService

`func (o *BillingAlertPatch) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingAlertPatch) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingAlertPatch) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BillingAlertPatch) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSoftPct

`func (o *BillingAlertPatch) GetSoftPct() int64`

GetSoftPct returns the SoftPct field if non-nil, zero value otherwise.

### GetSoftPctOk

`func (o *BillingAlertPatch) GetSoftPctOk() (*int64, bool)`

GetSoftPctOk returns a tuple with the SoftPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftPct

`func (o *BillingAlertPatch) SetSoftPct(v int64)`

SetSoftPct sets SoftPct field to given value.

### HasSoftPct

`func (o *BillingAlertPatch) HasSoftPct() bool`

HasSoftPct returns a boolean if a field has been set.

### GetSubject

`func (o *BillingAlertPatch) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *BillingAlertPatch) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *BillingAlertPatch) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *BillingAlertPatch) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetThreshold

`func (o *BillingAlertPatch) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *BillingAlertPatch) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *BillingAlertPatch) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *BillingAlertPatch) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTitle

`func (o *BillingAlertPatch) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillingAlertPatch) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillingAlertPatch) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BillingAlertPatch) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


