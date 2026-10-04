# AutoSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the local time, HH:MM on a 24-hour clock, for daily, weekdays and weekly. Hourly takes its minute, and runs on the hour without one. | [optional] 
**Cron** | Pointer to **string** | Cron is a five-field cron expression (minute hour day-of-month month day-of-week), for kind cron only, read in TZ. | [optional] 
**Day** | Pointer to **string** | Day is the weekday a weekly schedule runs on: mon, tue, wed, thu, fri, sat or sun. | [optional] 
**Kind** | Pointer to **string** | Kind is manual, hourly, daily, weekdays, weekly or cron. Manual runs only when asked. | [optional] 
**Tz** | Pointer to **string** | TZ is the IANA time zone every time here is read in. Empty is UTC. | [optional] 

## Methods

### NewAutoSchedule

`func NewAutoSchedule() *AutoSchedule`

NewAutoSchedule instantiates a new AutoSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScheduleWithDefaults

`func NewAutoScheduleWithDefaults() *AutoSchedule`

NewAutoScheduleWithDefaults instantiates a new AutoSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *AutoSchedule) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AutoSchedule) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AutoSchedule) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *AutoSchedule) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetCron

`func (o *AutoSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *AutoSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *AutoSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *AutoSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDay

`func (o *AutoSchedule) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *AutoSchedule) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *AutoSchedule) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *AutoSchedule) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetKind

`func (o *AutoSchedule) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AutoSchedule) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AutoSchedule) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AutoSchedule) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTz

`func (o *AutoSchedule) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *AutoSchedule) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *AutoSchedule) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *AutoSchedule) HasTz() bool`

HasTz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


