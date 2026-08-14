# O11yO11yQueueRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | Data holds the row&#39;s cells keyed by column name; each cell&#39;s JSON type is the column&#39;s own, so the bytes pass through verbatim. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp anchors the row in time. | [optional] 

## Methods

### NewO11yO11yQueueRow

`func NewO11yO11yQueueRow() *O11yO11yQueueRow`

NewO11yO11yQueueRow instantiates a new O11yO11yQueueRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueueRowWithDefaults

`func NewO11yO11yQueueRowWithDefaults() *O11yO11yQueueRow`

NewO11yO11yQueueRowWithDefaults instantiates a new O11yO11yQueueRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yQueueRow) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yQueueRow) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yQueueRow) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yQueueRow) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yQueueRow) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yQueueRow) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yQueueRow) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yQueueRow) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


