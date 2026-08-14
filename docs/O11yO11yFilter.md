# O11yO11yFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Field is the column to test. | [optional] 
**Op** | Pointer to **string** | Op is how to test it: eq, neq or like. | [optional] 
**Value** | Pointer to **string** | Value is what to test it against. | [optional] 

## Methods

### NewO11yO11yFilter

`func NewO11yO11yFilter() *O11yO11yFilter`

NewO11yO11yFilter instantiates a new O11yO11yFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFilterWithDefaults

`func NewO11yO11yFilterWithDefaults() *O11yO11yFilter`

NewO11yO11yFilterWithDefaults instantiates a new O11yO11yFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *O11yO11yFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *O11yO11yFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *O11yO11yFilter) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *O11yO11yFilter) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOp

`func (o *O11yO11yFilter) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *O11yO11yFilter) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *O11yO11yFilter) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *O11yO11yFilter) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


