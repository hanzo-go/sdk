# HfModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Downloads** | Pointer to **int32** |  | [optional] 
**Gated** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LibraryName** | Pointer to **string** |  | [optional] 
**Likes** | Pointer to **int32** |  | [optional] 
**PipelineTag** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHfModel

`func NewHfModel() *HfModel`

NewHfModel instantiates a new HfModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHfModelWithDefaults

`func NewHfModelWithDefaults() *HfModel`

NewHfModelWithDefaults instantiates a new HfModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloads

`func (o *HfModel) GetDownloads() int32`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *HfModel) GetDownloadsOk() (*int32, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *HfModel) SetDownloads(v int32)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *HfModel) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetGated

`func (o *HfModel) GetGated() interface{}`

GetGated returns the Gated field if non-nil, zero value otherwise.

### GetGatedOk

`func (o *HfModel) GetGatedOk() (*interface{}, bool)`

GetGatedOk returns a tuple with the Gated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGated

`func (o *HfModel) SetGated(v interface{})`

SetGated sets Gated field to given value.

### HasGated

`func (o *HfModel) HasGated() bool`

HasGated returns a boolean if a field has been set.

### SetGatedNil

`func (o *HfModel) SetGatedNil(b bool)`

 SetGatedNil sets the value for Gated to be an explicit nil

### UnsetGated
`func (o *HfModel) UnsetGated()`

UnsetGated ensures that no value is present for Gated, not even an explicit nil
### GetId

`func (o *HfModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HfModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HfModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HfModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModified

`func (o *HfModel) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *HfModel) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *HfModel) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *HfModel) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLibraryName

`func (o *HfModel) GetLibraryName() string`

GetLibraryName returns the LibraryName field if non-nil, zero value otherwise.

### GetLibraryNameOk

`func (o *HfModel) GetLibraryNameOk() (*string, bool)`

GetLibraryNameOk returns a tuple with the LibraryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryName

`func (o *HfModel) SetLibraryName(v string)`

SetLibraryName sets LibraryName field to given value.

### HasLibraryName

`func (o *HfModel) HasLibraryName() bool`

HasLibraryName returns a boolean if a field has been set.

### GetLikes

`func (o *HfModel) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *HfModel) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *HfModel) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *HfModel) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetPipelineTag

`func (o *HfModel) GetPipelineTag() string`

GetPipelineTag returns the PipelineTag field if non-nil, zero value otherwise.

### GetPipelineTagOk

`func (o *HfModel) GetPipelineTagOk() (*string, bool)`

GetPipelineTagOk returns a tuple with the PipelineTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineTag

`func (o *HfModel) SetPipelineTag(v string)`

SetPipelineTag sets PipelineTag field to given value.

### HasPipelineTag

`func (o *HfModel) HasPipelineTag() bool`

HasPipelineTag returns a boolean if a field has been set.

### GetPrivate

`func (o *HfModel) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *HfModel) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *HfModel) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *HfModel) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetTags

`func (o *HfModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HfModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HfModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HfModel) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


