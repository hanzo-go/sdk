# O11ySpanAggregationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **interface{}** |  | [optional] 
**Field** | Pointer to [**O11yTelemetryFieldKey**](O11yTelemetryFieldKey.md) |  | [optional] 
**Value** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewO11ySpanAggregationResult

`func NewO11ySpanAggregationResult() *O11ySpanAggregationResult`

NewO11ySpanAggregationResult instantiates a new O11ySpanAggregationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySpanAggregationResultWithDefaults

`func NewO11ySpanAggregationResultWithDefaults() *O11ySpanAggregationResult`

NewO11ySpanAggregationResultWithDefaults instantiates a new O11ySpanAggregationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *O11ySpanAggregationResult) GetAggregation() interface{}`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *O11ySpanAggregationResult) GetAggregationOk() (*interface{}, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *O11ySpanAggregationResult) SetAggregation(v interface{})`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *O11ySpanAggregationResult) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### SetAggregationNil

`func (o *O11ySpanAggregationResult) SetAggregationNil(b bool)`

 SetAggregationNil sets the value for Aggregation to be an explicit nil

### UnsetAggregation
`func (o *O11ySpanAggregationResult) UnsetAggregation()`

UnsetAggregation ensures that no value is present for Aggregation, not even an explicit nil
### GetField

`func (o *O11ySpanAggregationResult) GetField() O11yTelemetryFieldKey`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *O11ySpanAggregationResult) GetFieldOk() (*O11yTelemetryFieldKey, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *O11ySpanAggregationResult) SetField(v O11yTelemetryFieldKey)`

SetField sets Field field to given value.

### HasField

`func (o *O11ySpanAggregationResult) HasField() bool`

HasField returns a boolean if a field has been set.

### GetValue

`func (o *O11ySpanAggregationResult) GetValue() map[string]int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11ySpanAggregationResult) GetValueOk() (*map[string]int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11ySpanAggregationResult) SetValue(v map[string]int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11ySpanAggregationResult) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


