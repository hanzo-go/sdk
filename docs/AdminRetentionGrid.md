# AdminRetentionGrid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **string** |  | [optional] 
**Periods** | Pointer to **int32** |  | [optional] 
**Cohorts** | Pointer to [**[]AdminRetentionCohort**](AdminRetentionCohort.md) |  | [optional] 

## Methods

### NewAdminRetentionGrid

`func NewAdminRetentionGrid() *AdminRetentionGrid`

NewAdminRetentionGrid instantiates a new AdminRetentionGrid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminRetentionGridWithDefaults

`func NewAdminRetentionGridWithDefaults() *AdminRetentionGrid`

NewAdminRetentionGridWithDefaults instantiates a new AdminRetentionGrid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *AdminRetentionGrid) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AdminRetentionGrid) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AdminRetentionGrid) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AdminRetentionGrid) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPeriods

`func (o *AdminRetentionGrid) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *AdminRetentionGrid) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *AdminRetentionGrid) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *AdminRetentionGrid) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.

### GetCohorts

`func (o *AdminRetentionGrid) GetCohorts() []AdminRetentionCohort`

GetCohorts returns the Cohorts field if non-nil, zero value otherwise.

### GetCohortsOk

`func (o *AdminRetentionGrid) GetCohortsOk() (*[]AdminRetentionCohort, bool)`

GetCohortsOk returns a tuple with the Cohorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohorts

`func (o *AdminRetentionGrid) SetCohorts(v []AdminRetentionCohort)`

SetCohorts sets Cohorts field to given value.

### HasCohorts

`func (o *AdminRetentionGrid) HasCohorts() bool`

HasCohorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


