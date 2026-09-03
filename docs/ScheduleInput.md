# ScheduleInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the campaign id from the path. | [optional] 
**ScheduledAt** | Pointer to **int64** | ScheduledAt is the unix send time. 0 clears the schedule. | [optional] 

## Methods

### NewScheduleInput

`func NewScheduleInput() *ScheduleInput`

NewScheduleInput instantiates a new ScheduleInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleInputWithDefaults

`func NewScheduleInputWithDefaults() *ScheduleInput`

NewScheduleInputWithDefaults instantiates a new ScheduleInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduledAt

`func (o *ScheduleInput) GetScheduledAt() int64`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *ScheduleInput) GetScheduledAtOk() (*int64, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *ScheduleInput) SetScheduledAt(v int64)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *ScheduleInput) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


