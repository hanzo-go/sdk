# AutoWhenOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | Next is when it would first run, RFC 3339 UTC; null when it runs only on demand. | [optional] 
**Says** | Pointer to **string** | Says is the schedule said back in plain words. | [optional] 
**Schedule** | Pointer to [**AutoSchedule**](AutoSchedule.md) | Schedule is the schedule, ready to send as an automation&#39;s &#x60;schedule&#x60;. | [optional] 

## Methods

### NewAutoWhenOut

`func NewAutoWhenOut() *AutoWhenOut`

NewAutoWhenOut instantiates a new AutoWhenOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoWhenOutWithDefaults

`func NewAutoWhenOutWithDefaults() *AutoWhenOut`

NewAutoWhenOutWithDefaults instantiates a new AutoWhenOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *AutoWhenOut) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AutoWhenOut) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AutoWhenOut) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *AutoWhenOut) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetSays

`func (o *AutoWhenOut) GetSays() string`

GetSays returns the Says field if non-nil, zero value otherwise.

### GetSaysOk

`func (o *AutoWhenOut) GetSaysOk() (*string, bool)`

GetSaysOk returns a tuple with the Says field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSays

`func (o *AutoWhenOut) SetSays(v string)`

SetSays sets Says field to given value.

### HasSays

`func (o *AutoWhenOut) HasSays() bool`

HasSays returns a boolean if a field has been set.

### GetSchedule

`func (o *AutoWhenOut) GetSchedule() AutoSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AutoWhenOut) GetScheduleOk() (*AutoSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AutoWhenOut) SetSchedule(v AutoSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AutoWhenOut) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


