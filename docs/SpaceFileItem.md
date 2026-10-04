# SpaceFileItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etag** | Pointer to **string** | ETag is the store&#39;s entity tag for the bytes currently at this name, with the quotes the store wraps it in stripped. It is an opaque VERSION and not a checksum to verify against: a single-part upload&#39;s tag happens to be the MD5 of the content and a multipart upload&#39;s is not, and nothing here says which this was. Compare two reads of one file to learn whether it changed; absent for a folder, and for a file the store reports none for. | [optional] 
**IsFolder** | Pointer to **bool** | Folder is true for a folder entry, which is emergent from \&quot;/\&quot; in the names beneath it rather than a thing that was created. | [optional] 
**ModifiedAt** | Pointer to **int64** | ModifiedAt is when the file was last written, in unix seconds, and 0 for a folder. | [optional] 
**Name** | Pointer to **string** | Name is the entry&#39;s name RELATIVE to the folder that was listed. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s size in bytes, and 0 for a folder. | [optional] 

## Methods

### NewSpaceFileItem

`func NewSpaceFileItem() *SpaceFileItem`

NewSpaceFileItem instantiates a new SpaceFileItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceFileItemWithDefaults

`func NewSpaceFileItemWithDefaults() *SpaceFileItem`

NewSpaceFileItemWithDefaults instantiates a new SpaceFileItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *SpaceFileItem) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *SpaceFileItem) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *SpaceFileItem) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *SpaceFileItem) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetIsFolder

`func (o *SpaceFileItem) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *SpaceFileItem) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *SpaceFileItem) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *SpaceFileItem) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SpaceFileItem) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SpaceFileItem) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SpaceFileItem) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SpaceFileItem) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *SpaceFileItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceFileItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceFileItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpaceFileItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *SpaceFileItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SpaceFileItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SpaceFileItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SpaceFileItem) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


