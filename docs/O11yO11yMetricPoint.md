# O11yO11yMetricPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partial** | Pointer to **bool** | Partial marks a point whose bucket the window only partly covers. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp is the point&#39;s time as a Unix timestamp in milliseconds. | [optional] 
**Value** | Pointer to **float64** | Value is the point&#39;s value. | [optional] 
**Values** | Pointer to **[]float64** | Values carries the bucket values of a heatmap point. | [optional] 

## Methods

### NewO11yO11yMetricPoint

`func NewO11yO11yMetricPoint() *O11yO11yMetricPoint`

NewO11yO11yMetricPoint instantiates a new O11yO11yMetricPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricPointWithDefaults

`func NewO11yO11yMetricPointWithDefaults() *O11yO11yMetricPoint`

NewO11yO11yMetricPointWithDefaults instantiates a new O11yO11yMetricPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartial

`func (o *O11yO11yMetricPoint) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *O11yO11yMetricPoint) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *O11yO11yMetricPoint) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *O11yO11yMetricPoint) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yMetricPoint) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yMetricPoint) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yMetricPoint) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yMetricPoint) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yMetricPoint) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yMetricPoint) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yMetricPoint) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yMetricPoint) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *O11yO11yMetricPoint) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *O11yO11yMetricPoint) GetValuesOk() (*[]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *O11yO11yMetricPoint) SetValues(v []float64)`

SetValues sets Values field to given value.

### HasValues

`func (o *O11yO11yMetricPoint) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


