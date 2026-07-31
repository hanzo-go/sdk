# CloudRetentionGrid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cohorts** | Pointer to [**[]CloudRetentionCohort**](CloudRetentionCohort.md) |  | [optional] 
**Interval** | Pointer to **string** | \&quot;month\&quot; | [optional] 
**Periods** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudRetentionGrid

`func NewCloudRetentionGrid() *CloudRetentionGrid`

NewCloudRetentionGrid instantiates a new CloudRetentionGrid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRetentionGridWithDefaults

`func NewCloudRetentionGridWithDefaults() *CloudRetentionGrid`

NewCloudRetentionGridWithDefaults instantiates a new CloudRetentionGrid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCohorts

`func (o *CloudRetentionGrid) GetCohorts() []CloudRetentionCohort`

GetCohorts returns the Cohorts field if non-nil, zero value otherwise.

### GetCohortsOk

`func (o *CloudRetentionGrid) GetCohortsOk() (*[]CloudRetentionCohort, bool)`

GetCohortsOk returns a tuple with the Cohorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohorts

`func (o *CloudRetentionGrid) SetCohorts(v []CloudRetentionCohort)`

SetCohorts sets Cohorts field to given value.

### HasCohorts

`func (o *CloudRetentionGrid) HasCohorts() bool`

HasCohorts returns a boolean if a field has been set.

### GetInterval

`func (o *CloudRetentionGrid) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudRetentionGrid) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudRetentionGrid) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudRetentionGrid) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPeriods

`func (o *CloudRetentionGrid) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *CloudRetentionGrid) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *CloudRetentionGrid) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *CloudRetentionGrid) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


