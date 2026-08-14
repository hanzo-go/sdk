# O11yO11yFilterKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | Pointer to **string** | DataType is the attribute&#39;s value type — string, int64, float64 or bool. | [optional] 
**IsColumn** | Pointer to **bool** | IsColumn marks an attribute stored as its own column. | [optional] 
**IsJSON** | Pointer to **bool** | IsJSON marks an attribute read out of a JSON body. | [optional] 
**Key** | Pointer to **string** | Key is the attribute name. | [optional] 
**Type** | Pointer to **string** | Type says where the attribute lives — tag or resource. | [optional] 

## Methods

### NewO11yO11yFilterKey

`func NewO11yO11yFilterKey() *O11yO11yFilterKey`

NewO11yO11yFilterKey instantiates a new O11yO11yFilterKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFilterKeyWithDefaults

`func NewO11yO11yFilterKeyWithDefaults() *O11yO11yFilterKey`

NewO11yO11yFilterKeyWithDefaults instantiates a new O11yO11yFilterKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *O11yO11yFilterKey) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *O11yO11yFilterKey) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *O11yO11yFilterKey) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *O11yO11yFilterKey) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetIsColumn

`func (o *O11yO11yFilterKey) GetIsColumn() bool`

GetIsColumn returns the IsColumn field if non-nil, zero value otherwise.

### GetIsColumnOk

`func (o *O11yO11yFilterKey) GetIsColumnOk() (*bool, bool)`

GetIsColumnOk returns a tuple with the IsColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsColumn

`func (o *O11yO11yFilterKey) SetIsColumn(v bool)`

SetIsColumn sets IsColumn field to given value.

### HasIsColumn

`func (o *O11yO11yFilterKey) HasIsColumn() bool`

HasIsColumn returns a boolean if a field has been set.

### GetIsJSON

`func (o *O11yO11yFilterKey) GetIsJSON() bool`

GetIsJSON returns the IsJSON field if non-nil, zero value otherwise.

### GetIsJSONOk

`func (o *O11yO11yFilterKey) GetIsJSONOk() (*bool, bool)`

GetIsJSONOk returns a tuple with the IsJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJSON

`func (o *O11yO11yFilterKey) SetIsJSON(v bool)`

SetIsJSON sets IsJSON field to given value.

### HasIsJSON

`func (o *O11yO11yFilterKey) HasIsJSON() bool`

HasIsJSON returns a boolean if a field has been set.

### GetKey

`func (o *O11yO11yFilterKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yFilterKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yFilterKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yFilterKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yFilterKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yFilterKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yFilterKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yFilterKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


