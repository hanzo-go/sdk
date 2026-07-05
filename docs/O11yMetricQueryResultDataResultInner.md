# O11yMetricQueryResultDataResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **map[string]string** |  | [optional] 
**Value** | Pointer to **[]interface{}** | [unix_timestamp, value] for vector/scalar | [optional] 
**Values** | Pointer to **[][]interface{}** | [[unix_timestamp, value], ...] for matrix | [optional] 

## Methods

### NewO11yMetricQueryResultDataResultInner

`func NewO11yMetricQueryResultDataResultInner() *O11yMetricQueryResultDataResultInner`

NewO11yMetricQueryResultDataResultInner instantiates a new O11yMetricQueryResultDataResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMetricQueryResultDataResultInnerWithDefaults

`func NewO11yMetricQueryResultDataResultInnerWithDefaults() *O11yMetricQueryResultDataResultInner`

NewO11yMetricQueryResultDataResultInnerWithDefaults instantiates a new O11yMetricQueryResultDataResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *O11yMetricQueryResultDataResultInner) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *O11yMetricQueryResultDataResultInner) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *O11yMetricQueryResultDataResultInner) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *O11yMetricQueryResultDataResultInner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetValue

`func (o *O11yMetricQueryResultDataResultInner) GetValue() []interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yMetricQueryResultDataResultInner) GetValueOk() (*[]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yMetricQueryResultDataResultInner) SetValue(v []interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yMetricQueryResultDataResultInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *O11yMetricQueryResultDataResultInner) GetValues() [][]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *O11yMetricQueryResultDataResultInner) GetValuesOk() (*[][]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *O11yMetricQueryResultDataResultInner) SetValues(v [][]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *O11yMetricQueryResultDataResultInner) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


