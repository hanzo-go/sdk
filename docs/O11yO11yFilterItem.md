# O11yO11yFilterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**O11yO11yAttributeKey**](O11yO11yAttributeKey.md) | Key is the attribute tested. | [optional] 
**Op** | Pointer to **string** | Operator is the comparison, e.g. &#x3D;, !&#x3D;, in, contains. | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value is what it is tested against; its type follows the attribute&#39;s. | [optional] 

## Methods

### NewO11yO11yFilterItem

`func NewO11yO11yFilterItem() *O11yO11yFilterItem`

NewO11yO11yFilterItem instantiates a new O11yO11yFilterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFilterItemWithDefaults

`func NewO11yO11yFilterItemWithDefaults() *O11yO11yFilterItem`

NewO11yO11yFilterItemWithDefaults instantiates a new O11yO11yFilterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *O11yO11yFilterItem) GetKey() O11yO11yAttributeKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yFilterItem) GetKeyOk() (*O11yO11yAttributeKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yFilterItem) SetKey(v O11yO11yAttributeKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yFilterItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOp

`func (o *O11yO11yFilterItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *O11yO11yFilterItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *O11yO11yFilterItem) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *O11yO11yFilterItem) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yFilterItem) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yFilterItem) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yFilterItem) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yFilterItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


