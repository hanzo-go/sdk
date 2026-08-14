# O11yO11ySpanPercentile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentiles** | Pointer to [**O11yO11yPercentiles**](O11yO11yPercentiles.md) | Percentiles are the peer group&#39;s duration percentiles. | [optional] 
**Position** | Pointer to [**O11yO11yPercentilePosition**](O11yO11yPercentilePosition.md) | Position is where the given duration lands. | [optional] 

## Methods

### NewO11yO11ySpanPercentile

`func NewO11yO11ySpanPercentile() *O11yO11ySpanPercentile`

NewO11yO11ySpanPercentile instantiates a new O11yO11ySpanPercentile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySpanPercentileWithDefaults

`func NewO11yO11ySpanPercentileWithDefaults() *O11yO11ySpanPercentile`

NewO11yO11ySpanPercentileWithDefaults instantiates a new O11yO11ySpanPercentile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentiles

`func (o *O11yO11ySpanPercentile) GetPercentiles() O11yO11yPercentiles`

GetPercentiles returns the Percentiles field if non-nil, zero value otherwise.

### GetPercentilesOk

`func (o *O11yO11ySpanPercentile) GetPercentilesOk() (*O11yO11yPercentiles, bool)`

GetPercentilesOk returns a tuple with the Percentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentiles

`func (o *O11yO11ySpanPercentile) SetPercentiles(v O11yO11yPercentiles)`

SetPercentiles sets Percentiles field to given value.

### HasPercentiles

`func (o *O11yO11ySpanPercentile) HasPercentiles() bool`

HasPercentiles returns a boolean if a field has been set.

### GetPosition

`func (o *O11yO11ySpanPercentile) GetPosition() O11yO11yPercentilePosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *O11yO11ySpanPercentile) GetPositionOk() (*O11yO11yPercentilePosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *O11yO11ySpanPercentile) SetPosition(v O11yO11yPercentilePosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *O11yO11ySpanPercentile) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


