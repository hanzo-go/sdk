# O11yO11yLogFilterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**O11yO11yLogFilterKey**](O11yO11yLogFilterKey.md) | Key is the field the predicate tests. | [optional] 
**Op** | Pointer to **string** | Op is the comparison, e.g. &#x3D;, !&#x3D;, in, contains. | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yO11yLogFilterItem

`func NewO11yO11yLogFilterItem() *O11yO11yLogFilterItem`

NewO11yO11yLogFilterItem instantiates a new O11yO11yLogFilterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogFilterItemWithDefaults

`func NewO11yO11yLogFilterItemWithDefaults() *O11yO11yLogFilterItem`

NewO11yO11yLogFilterItemWithDefaults instantiates a new O11yO11yLogFilterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *O11yO11yLogFilterItem) GetKey() O11yO11yLogFilterKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yLogFilterItem) GetKeyOk() (*O11yO11yLogFilterKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yLogFilterItem) SetKey(v O11yO11yLogFilterKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yLogFilterItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOp

`func (o *O11yO11yLogFilterItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *O11yO11yLogFilterItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *O11yO11yLogFilterItem) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *O11yO11yLogFilterItem) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yLogFilterItem) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yLogFilterItem) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yLogFilterItem) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yLogFilterItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *O11yO11yLogFilterItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *O11yO11yLogFilterItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


