# O11yO11yFieldSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** | DataType is the field&#39;s data type, e.g. string, int64, float64, bool. Required. | 
**Index** | Pointer to **string** | Index is the index expression to put on the column, e.g. minmax, set(N), bloom_filter(P), tokenbf_v1(S,H,SEED). Empty keeps the default. | [optional] 
**IndexGranularity** | Pointer to **int32** | IndexGranularity is the index granularity in rows. | [optional] 
**Name** | **string** | Name is the field to tune. Required. | 
**Selected** | Pointer to **bool** | Selected materializes the field as its own column when true. | [optional] 
**Type** | **string** | Type is where the field lives: attributes or resources. Required. | 

## Methods

### NewO11yO11yFieldSetting

`func NewO11yO11yFieldSetting(dataType string, name string, type_ string, ) *O11yO11yFieldSetting`

NewO11yO11yFieldSetting instantiates a new O11yO11yFieldSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFieldSettingWithDefaults

`func NewO11yO11yFieldSettingWithDefaults() *O11yO11yFieldSetting`

NewO11yO11yFieldSettingWithDefaults instantiates a new O11yO11yFieldSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *O11yO11yFieldSetting) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *O11yO11yFieldSetting) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *O11yO11yFieldSetting) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetIndex

`func (o *O11yO11yFieldSetting) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *O11yO11yFieldSetting) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *O11yO11yFieldSetting) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *O11yO11yFieldSetting) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIndexGranularity

`func (o *O11yO11yFieldSetting) GetIndexGranularity() int32`

GetIndexGranularity returns the IndexGranularity field if non-nil, zero value otherwise.

### GetIndexGranularityOk

`func (o *O11yO11yFieldSetting) GetIndexGranularityOk() (*int32, bool)`

GetIndexGranularityOk returns a tuple with the IndexGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexGranularity

`func (o *O11yO11yFieldSetting) SetIndexGranularity(v int32)`

SetIndexGranularity sets IndexGranularity field to given value.

### HasIndexGranularity

`func (o *O11yO11yFieldSetting) HasIndexGranularity() bool`

HasIndexGranularity returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yFieldSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yFieldSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yFieldSetting) SetName(v string)`

SetName sets Name field to given value.


### GetSelected

`func (o *O11yO11yFieldSetting) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *O11yO11yFieldSetting) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *O11yO11yFieldSetting) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *O11yO11yFieldSetting) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yFieldSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yFieldSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yFieldSetting) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


