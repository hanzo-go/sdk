# LeaderboardPulseSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **[]int64** | Days is one count a day, aligned with PulseView.Days. | [optional] 
**Hours** | Pointer to **[]int64** | Hours is one count an hour, aligned with PulseView.Hours. Absent for a series recorded by the day alone. | [optional] 

## Methods

### NewLeaderboardPulseSeries

`func NewLeaderboardPulseSeries() *LeaderboardPulseSeries`

NewLeaderboardPulseSeries instantiates a new LeaderboardPulseSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardPulseSeriesWithDefaults

`func NewLeaderboardPulseSeriesWithDefaults() *LeaderboardPulseSeries`

NewLeaderboardPulseSeriesWithDefaults instantiates a new LeaderboardPulseSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *LeaderboardPulseSeries) GetDays() []int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *LeaderboardPulseSeries) GetDaysOk() (*[]int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *LeaderboardPulseSeries) SetDays(v []int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *LeaderboardPulseSeries) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetHours

`func (o *LeaderboardPulseSeries) GetHours() []int64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *LeaderboardPulseSeries) GetHoursOk() (*[]int64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *LeaderboardPulseSeries) SetHours(v []int64)`

SetHours sets Hours field to given value.

### HasHours

`func (o *LeaderboardPulseSeries) HasHours() bool`

HasHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


