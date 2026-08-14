# RetentionGrid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cohorts** | Pointer to [**[]RetentionCohort**](RetentionCohort.md) |  | [optional] 
**Interval** | Pointer to **string** | \&quot;month\&quot; | [optional] 
**Periods** | Pointer to **int32** |  | [optional] 

## Methods

### NewRetentionGrid

`func NewRetentionGrid() *RetentionGrid`

NewRetentionGrid instantiates a new RetentionGrid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionGridWithDefaults

`func NewRetentionGridWithDefaults() *RetentionGrid`

NewRetentionGridWithDefaults instantiates a new RetentionGrid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCohorts

`func (o *RetentionGrid) GetCohorts() []RetentionCohort`

GetCohorts returns the Cohorts field if non-nil, zero value otherwise.

### GetCohortsOk

`func (o *RetentionGrid) GetCohortsOk() (*[]RetentionCohort, bool)`

GetCohortsOk returns a tuple with the Cohorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohorts

`func (o *RetentionGrid) SetCohorts(v []RetentionCohort)`

SetCohorts sets Cohorts field to given value.

### HasCohorts

`func (o *RetentionGrid) HasCohorts() bool`

HasCohorts returns a boolean if a field has been set.

### GetInterval

`func (o *RetentionGrid) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RetentionGrid) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RetentionGrid) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *RetentionGrid) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPeriods

`func (o *RetentionGrid) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *RetentionGrid) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *RetentionGrid) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *RetentionGrid) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


