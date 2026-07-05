# AnalyticsDateRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **time.Time** |  | 
**EndDate** | **time.Time** |  | 
**Num** | Pointer to **float32** |  | [optional] 
**Offset** | Pointer to **float32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewAnalyticsDateRange

`func NewAnalyticsDateRange(startDate time.Time, endDate time.Time, ) *AnalyticsDateRange`

NewAnalyticsDateRange instantiates a new AnalyticsDateRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsDateRangeWithDefaults

`func NewAnalyticsDateRangeWithDefaults() *AnalyticsDateRange`

NewAnalyticsDateRangeWithDefaults instantiates a new AnalyticsDateRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *AnalyticsDateRange) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AnalyticsDateRange) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AnalyticsDateRange) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *AnalyticsDateRange) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AnalyticsDateRange) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AnalyticsDateRange) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetNum

`func (o *AnalyticsDateRange) GetNum() float32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *AnalyticsDateRange) GetNumOk() (*float32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *AnalyticsDateRange) SetNum(v float32)`

SetNum sets Num field to given value.

### HasNum

`func (o *AnalyticsDateRange) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetOffset

`func (o *AnalyticsDateRange) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AnalyticsDateRange) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AnalyticsDateRange) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AnalyticsDateRange) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetUnit

`func (o *AnalyticsDateRange) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AnalyticsDateRange) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AnalyticsDateRange) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AnalyticsDateRange) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *AnalyticsDateRange) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AnalyticsDateRange) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AnalyticsDateRange) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AnalyticsDateRange) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


