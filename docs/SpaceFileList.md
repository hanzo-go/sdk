# SpaceFileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drive** | Pointer to **string** | Drive is the drive that was listed. | [optional] 
**Files** | Pointer to [**[]SpaceFileItem**](SpaceFileItem.md) | Files are the entries at this level, names RELATIVE to Folder. | [optional] 
**Folder** | Pointer to **string** | Folder is the sub-folder the listing was scoped to, cleaned. Empty for the drive&#39;s own root. | [optional] 
**Space** | Pointer to **string** | Space is the space that was listed. | [optional] 
**Total** | Pointer to **int64** | Total is how many entries came back. The listing is BOUNDED, so a drive with more files than the cap answers the cap and this says so — it is not a count of what the drive holds. | [optional] 

## Methods

### NewSpaceFileList

`func NewSpaceFileList() *SpaceFileList`

NewSpaceFileList instantiates a new SpaceFileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceFileListWithDefaults

`func NewSpaceFileListWithDefaults() *SpaceFileList`

NewSpaceFileListWithDefaults instantiates a new SpaceFileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrive

`func (o *SpaceFileList) GetDrive() string`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *SpaceFileList) GetDriveOk() (*string, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *SpaceFileList) SetDrive(v string)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *SpaceFileList) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### GetFiles

`func (o *SpaceFileList) GetFiles() []SpaceFileItem`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SpaceFileList) GetFilesOk() (*[]SpaceFileItem, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SpaceFileList) SetFiles(v []SpaceFileItem)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SpaceFileList) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFolder

`func (o *SpaceFileList) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *SpaceFileList) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *SpaceFileList) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *SpaceFileList) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetSpace

`func (o *SpaceFileList) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *SpaceFileList) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *SpaceFileList) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *SpaceFileList) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTotal

`func (o *SpaceFileList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SpaceFileList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SpaceFileList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SpaceFileList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


