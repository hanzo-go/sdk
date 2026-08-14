# O11yO11yMetricAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the attribute&#39;s name. | [optional] 
**ValueCount** | Pointer to **int32** | ValueCount is how many distinct values the attribute has. | [optional] 
**Values** | Pointer to **[]string** | Values are the attribute&#39;s distinct values. | [optional] 

## Methods

### NewO11yO11yMetricAttribute

`func NewO11yO11yMetricAttribute() *O11yO11yMetricAttribute`

NewO11yO11yMetricAttribute instantiates a new O11yO11yMetricAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricAttributeWithDefaults

`func NewO11yO11yMetricAttributeWithDefaults() *O11yO11yMetricAttribute`

NewO11yO11yMetricAttributeWithDefaults instantiates a new O11yO11yMetricAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *O11yO11yMetricAttribute) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yMetricAttribute) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yMetricAttribute) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yMetricAttribute) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValueCount

`func (o *O11yO11yMetricAttribute) GetValueCount() int32`

GetValueCount returns the ValueCount field if non-nil, zero value otherwise.

### GetValueCountOk

`func (o *O11yO11yMetricAttribute) GetValueCountOk() (*int32, bool)`

GetValueCountOk returns a tuple with the ValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCount

`func (o *O11yO11yMetricAttribute) SetValueCount(v int32)`

SetValueCount sets ValueCount field to given value.

### HasValueCount

`func (o *O11yO11yMetricAttribute) HasValueCount() bool`

HasValueCount returns a boolean if a field has been set.

### GetValues

`func (o *O11yO11yMetricAttribute) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *O11yO11yMetricAttribute) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *O11yO11yMetricAttribute) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *O11yO11yMetricAttribute) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


