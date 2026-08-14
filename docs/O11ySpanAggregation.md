# O11ySpanAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **interface{}** |  | [optional] 
**Field** | Pointer to [**O11yTelemetryFieldKey**](O11yTelemetryFieldKey.md) |  | [optional] 

## Methods

### NewO11ySpanAggregation

`func NewO11ySpanAggregation() *O11ySpanAggregation`

NewO11ySpanAggregation instantiates a new O11ySpanAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySpanAggregationWithDefaults

`func NewO11ySpanAggregationWithDefaults() *O11ySpanAggregation`

NewO11ySpanAggregationWithDefaults instantiates a new O11ySpanAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *O11ySpanAggregation) GetAggregation() interface{}`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *O11ySpanAggregation) GetAggregationOk() (*interface{}, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *O11ySpanAggregation) SetAggregation(v interface{})`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *O11ySpanAggregation) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### SetAggregationNil

`func (o *O11ySpanAggregation) SetAggregationNil(b bool)`

 SetAggregationNil sets the value for Aggregation to be an explicit nil

### UnsetAggregation
`func (o *O11ySpanAggregation) UnsetAggregation()`

UnsetAggregation ensures that no value is present for Aggregation, not even an explicit nil
### GetField

`func (o *O11ySpanAggregation) GetField() O11yTelemetryFieldKey`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *O11ySpanAggregation) GetFieldOk() (*O11yTelemetryFieldKey, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *O11ySpanAggregation) SetField(v O11yTelemetryFieldKey)`

SetField sets Field field to given value.

### HasField

`func (o *O11ySpanAggregation) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


