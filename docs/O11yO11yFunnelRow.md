# O11yO11yFunnelRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]map[string]interface{}** | Data are the row&#39;s columns, keyed by column name. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is the row&#39;s time. | [optional] 

## Methods

### NewO11yO11yFunnelRow

`func NewO11yO11yFunnelRow() *O11yO11yFunnelRow`

NewO11yO11yFunnelRow instantiates a new O11yO11yFunnelRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFunnelRowWithDefaults

`func NewO11yO11yFunnelRowWithDefaults() *O11yO11yFunnelRow`

NewO11yO11yFunnelRowWithDefaults instantiates a new O11yO11yFunnelRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yFunnelRow) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yFunnelRow) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yFunnelRow) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yFunnelRow) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yFunnelRow) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yFunnelRow) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yFunnelRow) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yFunnelRow) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


