# ObjectItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etag** | Pointer to **string** |  | [optional] 
**IsDir** | Pointer to **bool** | true for a folder (common prefix) | [optional] 
**Key** | Pointer to **string** | key RELATIVE to the requested prefix | [optional] 
**LastModified** | Pointer to **int32** | unix seconds (0 for a folder) | [optional] 
**Size** | Pointer to **int32** | bytes (0 for a folder) | [optional] 

## Methods

### NewObjectItem

`func NewObjectItem() *ObjectItem`

NewObjectItem instantiates a new ObjectItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectItemWithDefaults

`func NewObjectItemWithDefaults() *ObjectItem`

NewObjectItemWithDefaults instantiates a new ObjectItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *ObjectItem) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *ObjectItem) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *ObjectItem) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *ObjectItem) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetIsDir

`func (o *ObjectItem) GetIsDir() bool`

GetIsDir returns the IsDir field if non-nil, zero value otherwise.

### GetIsDirOk

`func (o *ObjectItem) GetIsDirOk() (*bool, bool)`

GetIsDirOk returns a tuple with the IsDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDir

`func (o *ObjectItem) SetIsDir(v bool)`

SetIsDir sets IsDir field to given value.

### HasIsDir

`func (o *ObjectItem) HasIsDir() bool`

HasIsDir returns a boolean if a field has been set.

### GetKey

`func (o *ObjectItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ObjectItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ObjectItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ObjectItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastModified

`func (o *ObjectItem) GetLastModified() int32`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ObjectItem) GetLastModifiedOk() (*int32, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ObjectItem) SetLastModified(v int32)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ObjectItem) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetSize

`func (o *ObjectItem) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ObjectItem) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ObjectItem) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ObjectItem) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


