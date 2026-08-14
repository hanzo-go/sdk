# O11yO11yLogFilterKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | Pointer to **string** | DataType is the field&#39;s data type, e.g. string, int64, float64, bool. | [optional] 
**IsColumn** | Pointer to **bool** | IsColumn marks a field materialized as its own column. | [optional] 
**IsJSON** | Pointer to **bool** | IsJSON marks a path into the record&#39;s JSON body. | [optional] 
**Key** | Pointer to **string** | Key is the field&#39;s name. | [optional] 
**Type** | Pointer to **string** | Type is where the field lives: tag or resource. | [optional] 

## Methods

### NewO11yO11yLogFilterKey

`func NewO11yO11yLogFilterKey() *O11yO11yLogFilterKey`

NewO11yO11yLogFilterKey instantiates a new O11yO11yLogFilterKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogFilterKeyWithDefaults

`func NewO11yO11yLogFilterKeyWithDefaults() *O11yO11yLogFilterKey`

NewO11yO11yLogFilterKeyWithDefaults instantiates a new O11yO11yLogFilterKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *O11yO11yLogFilterKey) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *O11yO11yLogFilterKey) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *O11yO11yLogFilterKey) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *O11yO11yLogFilterKey) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetIsColumn

`func (o *O11yO11yLogFilterKey) GetIsColumn() bool`

GetIsColumn returns the IsColumn field if non-nil, zero value otherwise.

### GetIsColumnOk

`func (o *O11yO11yLogFilterKey) GetIsColumnOk() (*bool, bool)`

GetIsColumnOk returns a tuple with the IsColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsColumn

`func (o *O11yO11yLogFilterKey) SetIsColumn(v bool)`

SetIsColumn sets IsColumn field to given value.

### HasIsColumn

`func (o *O11yO11yLogFilterKey) HasIsColumn() bool`

HasIsColumn returns a boolean if a field has been set.

### GetIsJSON

`func (o *O11yO11yLogFilterKey) GetIsJSON() bool`

GetIsJSON returns the IsJSON field if non-nil, zero value otherwise.

### GetIsJSONOk

`func (o *O11yO11yLogFilterKey) GetIsJSONOk() (*bool, bool)`

GetIsJSONOk returns a tuple with the IsJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJSON

`func (o *O11yO11yLogFilterKey) SetIsJSON(v bool)`

SetIsJSON sets IsJSON field to given value.

### HasIsJSON

`func (o *O11yO11yLogFilterKey) HasIsJSON() bool`

HasIsJSON returns a boolean if a field has been set.

### GetKey

`func (o *O11yO11yLogFilterKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yLogFilterKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yLogFilterKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yLogFilterKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yLogFilterKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yLogFilterKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yLogFilterKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yLogFilterKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


