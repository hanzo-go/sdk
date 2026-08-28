# HfRepoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Downloads** | Pointer to **int32** |  | [optional] 
**Gated** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Likes** | Pointer to **int32** |  | [optional] 
**PipelineTag** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Siblings** | Pointer to [**[]HfSibling**](HfSibling.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHfRepoInfo

`func NewHfRepoInfo() *HfRepoInfo`

NewHfRepoInfo instantiates a new HfRepoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHfRepoInfoWithDefaults

`func NewHfRepoInfoWithDefaults() *HfRepoInfo`

NewHfRepoInfoWithDefaults instantiates a new HfRepoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloads

`func (o *HfRepoInfo) GetDownloads() int32`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *HfRepoInfo) GetDownloadsOk() (*int32, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *HfRepoInfo) SetDownloads(v int32)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *HfRepoInfo) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetGated

`func (o *HfRepoInfo) GetGated() interface{}`

GetGated returns the Gated field if non-nil, zero value otherwise.

### GetGatedOk

`func (o *HfRepoInfo) GetGatedOk() (*interface{}, bool)`

GetGatedOk returns a tuple with the Gated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGated

`func (o *HfRepoInfo) SetGated(v interface{})`

SetGated sets Gated field to given value.

### HasGated

`func (o *HfRepoInfo) HasGated() bool`

HasGated returns a boolean if a field has been set.

### SetGatedNil

`func (o *HfRepoInfo) SetGatedNil(b bool)`

 SetGatedNil sets the value for Gated to be an explicit nil

### UnsetGated
`func (o *HfRepoInfo) UnsetGated()`

UnsetGated ensures that no value is present for Gated, not even an explicit nil
### GetId

`func (o *HfRepoInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HfRepoInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HfRepoInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HfRepoInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModified

`func (o *HfRepoInfo) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *HfRepoInfo) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *HfRepoInfo) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *HfRepoInfo) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLikes

`func (o *HfRepoInfo) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *HfRepoInfo) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *HfRepoInfo) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *HfRepoInfo) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetPipelineTag

`func (o *HfRepoInfo) GetPipelineTag() string`

GetPipelineTag returns the PipelineTag field if non-nil, zero value otherwise.

### GetPipelineTagOk

`func (o *HfRepoInfo) GetPipelineTagOk() (*string, bool)`

GetPipelineTagOk returns a tuple with the PipelineTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineTag

`func (o *HfRepoInfo) SetPipelineTag(v string)`

SetPipelineTag sets PipelineTag field to given value.

### HasPipelineTag

`func (o *HfRepoInfo) HasPipelineTag() bool`

HasPipelineTag returns a boolean if a field has been set.

### GetPrivate

`func (o *HfRepoInfo) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *HfRepoInfo) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *HfRepoInfo) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *HfRepoInfo) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetSiblings

`func (o *HfRepoInfo) GetSiblings() []HfSibling`

GetSiblings returns the Siblings field if non-nil, zero value otherwise.

### GetSiblingsOk

`func (o *HfRepoInfo) GetSiblingsOk() (*[]HfSibling, bool)`

GetSiblingsOk returns a tuple with the Siblings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiblings

`func (o *HfRepoInfo) SetSiblings(v []HfSibling)`

SetSiblings sets Siblings field to given value.

### HasSiblings

`func (o *HfRepoInfo) HasSiblings() bool`

HasSiblings returns a boolean if a field has been set.

### GetTags

`func (o *HfRepoInfo) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HfRepoInfo) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HfRepoInfo) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HfRepoInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


