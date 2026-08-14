# O11yO11yLogAggregateBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBy** | Pointer to **map[string]interface{}** | GroupBy carries the group&#39;s key values when the aggregate grouped. | [optional] 
**Timestamp** | Pointer to **int32** | Timestamp is the start of the bucket. | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yO11yLogAggregateBucket

`func NewO11yO11yLogAggregateBucket() *O11yO11yLogAggregateBucket`

NewO11yO11yLogAggregateBucket instantiates a new O11yO11yLogAggregateBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogAggregateBucketWithDefaults

`func NewO11yO11yLogAggregateBucketWithDefaults() *O11yO11yLogAggregateBucket`

NewO11yO11yLogAggregateBucketWithDefaults instantiates a new O11yO11yLogAggregateBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBy

`func (o *O11yO11yLogAggregateBucket) GetGroupBy() map[string]interface{}`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *O11yO11yLogAggregateBucket) GetGroupByOk() (*map[string]interface{}, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *O11yO11yLogAggregateBucket) SetGroupBy(v map[string]interface{})`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *O11yO11yLogAggregateBucket) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yLogAggregateBucket) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yLogAggregateBucket) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yLogAggregateBucket) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yLogAggregateBucket) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yLogAggregateBucket) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yLogAggregateBucket) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yLogAggregateBucket) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yLogAggregateBucket) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *O11yO11yLogAggregateBucket) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *O11yO11yLogAggregateBucket) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


