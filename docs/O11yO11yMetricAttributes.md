# O11yO11yMetricAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]O11yO11yMetricAttribute**](O11yO11yMetricAttribute.md) | Attributes are the keys, each with its values. | [optional] 
**TotalKeys** | Pointer to **int64** | TotalKeys is how many keys the metric has. | [optional] 

## Methods

### NewO11yO11yMetricAttributes

`func NewO11yO11yMetricAttributes() *O11yO11yMetricAttributes`

NewO11yO11yMetricAttributes instantiates a new O11yO11yMetricAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricAttributesWithDefaults

`func NewO11yO11yMetricAttributesWithDefaults() *O11yO11yMetricAttributes`

NewO11yO11yMetricAttributesWithDefaults instantiates a new O11yO11yMetricAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *O11yO11yMetricAttributes) GetAttributes() []O11yO11yMetricAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *O11yO11yMetricAttributes) GetAttributesOk() (*[]O11yO11yMetricAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *O11yO11yMetricAttributes) SetAttributes(v []O11yO11yMetricAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *O11yO11yMetricAttributes) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTotalKeys

`func (o *O11yO11yMetricAttributes) GetTotalKeys() int64`

GetTotalKeys returns the TotalKeys field if non-nil, zero value otherwise.

### GetTotalKeysOk

`func (o *O11yO11yMetricAttributes) GetTotalKeysOk() (*int64, bool)`

GetTotalKeysOk returns a tuple with the TotalKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKeys

`func (o *O11yO11yMetricAttributes) SetTotalKeys(v int64)`

SetTotalKeys sets TotalKeys field to given value.

### HasTotalKeys

`func (o *O11yO11yMetricAttributes) HasTotalKeys() bool`

HasTotalKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


