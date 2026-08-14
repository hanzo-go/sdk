# O11yO11yAttributeKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | Pointer to **string** | DataType is the attribute&#39;s value type — string, int64, float64 or bool. | [optional] 
**IsColumn** | Pointer to **bool** | IsColumn marks an attribute materialized as its own column. | [optional] 
**IsJSON** | Pointer to **bool** | IsJSON marks an attribute extracted from a JSON body. | [optional] 
**Key** | Pointer to **string** | Key is the attribute&#39;s name. | [optional] 
**Type** | Pointer to **string** | Type says where the attribute lives: tag or resource. | [optional] 

## Methods

### NewO11yO11yAttributeKey

`func NewO11yO11yAttributeKey() *O11yO11yAttributeKey`

NewO11yO11yAttributeKey instantiates a new O11yO11yAttributeKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAttributeKeyWithDefaults

`func NewO11yO11yAttributeKeyWithDefaults() *O11yO11yAttributeKey`

NewO11yO11yAttributeKeyWithDefaults instantiates a new O11yO11yAttributeKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *O11yO11yAttributeKey) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *O11yO11yAttributeKey) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *O11yO11yAttributeKey) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *O11yO11yAttributeKey) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetIsColumn

`func (o *O11yO11yAttributeKey) GetIsColumn() bool`

GetIsColumn returns the IsColumn field if non-nil, zero value otherwise.

### GetIsColumnOk

`func (o *O11yO11yAttributeKey) GetIsColumnOk() (*bool, bool)`

GetIsColumnOk returns a tuple with the IsColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsColumn

`func (o *O11yO11yAttributeKey) SetIsColumn(v bool)`

SetIsColumn sets IsColumn field to given value.

### HasIsColumn

`func (o *O11yO11yAttributeKey) HasIsColumn() bool`

HasIsColumn returns a boolean if a field has been set.

### GetIsJSON

`func (o *O11yO11yAttributeKey) GetIsJSON() bool`

GetIsJSON returns the IsJSON field if non-nil, zero value otherwise.

### GetIsJSONOk

`func (o *O11yO11yAttributeKey) GetIsJSONOk() (*bool, bool)`

GetIsJSONOk returns a tuple with the IsJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJSON

`func (o *O11yO11yAttributeKey) SetIsJSON(v bool)`

SetIsJSON sets IsJSON field to given value.

### HasIsJSON

`func (o *O11yO11yAttributeKey) HasIsJSON() bool`

HasIsJSON returns a boolean if a field has been set.

### GetKey

`func (o *O11yO11yAttributeKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yAttributeKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yAttributeKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yAttributeKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yAttributeKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yAttributeKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yAttributeKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yAttributeKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


