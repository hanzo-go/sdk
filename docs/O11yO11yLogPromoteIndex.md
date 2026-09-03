# O11yO11yLogPromoteIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldDataType** | Pointer to **string** | FieldDataType is the path&#39;s data type, e.g. string, number, bool. | [optional] 
**Granularity** | Pointer to **int64** | Granularity is the index granularity in rows. | [optional] 
**Type** | Pointer to **string** | Type is the index type, e.g. minmax, set(N), bloom_filter(P). | [optional] 

## Methods

### NewO11yO11yLogPromoteIndex

`func NewO11yO11yLogPromoteIndex() *O11yO11yLogPromoteIndex`

NewO11yO11yLogPromoteIndex instantiates a new O11yO11yLogPromoteIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogPromoteIndexWithDefaults

`func NewO11yO11yLogPromoteIndexWithDefaults() *O11yO11yLogPromoteIndex`

NewO11yO11yLogPromoteIndexWithDefaults instantiates a new O11yO11yLogPromoteIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldDataType

`func (o *O11yO11yLogPromoteIndex) GetFieldDataType() string`

GetFieldDataType returns the FieldDataType field if non-nil, zero value otherwise.

### GetFieldDataTypeOk

`func (o *O11yO11yLogPromoteIndex) GetFieldDataTypeOk() (*string, bool)`

GetFieldDataTypeOk returns a tuple with the FieldDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDataType

`func (o *O11yO11yLogPromoteIndex) SetFieldDataType(v string)`

SetFieldDataType sets FieldDataType field to given value.

### HasFieldDataType

`func (o *O11yO11yLogPromoteIndex) HasFieldDataType() bool`

HasFieldDataType returns a boolean if a field has been set.

### GetGranularity

`func (o *O11yO11yLogPromoteIndex) GetGranularity() int64`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *O11yO11yLogPromoteIndex) GetGranularityOk() (*int64, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *O11yO11yLogPromoteIndex) SetGranularity(v int64)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *O11yO11yLogPromoteIndex) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yLogPromoteIndex) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yLogPromoteIndex) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yLogPromoteIndex) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yLogPromoteIndex) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


