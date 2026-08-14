# O11yO11yUsageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count is how many spans were ingested in the bucket. | [optional] 
**Time** | Pointer to **time.Time** | Time is the bucket start. | [optional] 
**Timestamp** | Pointer to **int32** | Timestamp is the bucket start, as epoch nanoseconds. | [optional] 

## Methods

### NewO11yO11yUsageItem

`func NewO11yO11yUsageItem() *O11yO11yUsageItem`

NewO11yO11yUsageItem instantiates a new O11yO11yUsageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yUsageItemWithDefaults

`func NewO11yO11yUsageItemWithDefaults() *O11yO11yUsageItem`

NewO11yO11yUsageItemWithDefaults instantiates a new O11yO11yUsageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *O11yO11yUsageItem) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *O11yO11yUsageItem) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *O11yO11yUsageItem) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *O11yO11yUsageItem) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTime

`func (o *O11yO11yUsageItem) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *O11yO11yUsageItem) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *O11yO11yUsageItem) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *O11yO11yUsageItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yUsageItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yUsageItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yUsageItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yUsageItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


