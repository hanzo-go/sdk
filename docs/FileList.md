# FileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drive** | Pointer to **string** | Drive is the drive that was listed. | [optional] 
**Files** | Pointer to [**[]FileItem**](FileItem.md) | Files are the entries at this level, names RELATIVE to Folder. | [optional] 
**Folder** | Pointer to **string** | Folder is the sub-folder the listing was scoped to, cleaned. Empty for the drive&#39;s own root. | [optional] 
**Space** | Pointer to **string** | Space is the space that was listed. | [optional] 
**Total** | Pointer to **int64** | Total is how many entries came back. The listing is BOUNDED, so a drive with more files than the cap answers the cap and this says so — it is not a count of what the drive holds. | [optional] 

## Methods

### NewFileList

`func NewFileList() *FileList`

NewFileList instantiates a new FileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileListWithDefaults

`func NewFileListWithDefaults() *FileList`

NewFileListWithDefaults instantiates a new FileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrive

`func (o *FileList) GetDrive() string`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *FileList) GetDriveOk() (*string, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *FileList) SetDrive(v string)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *FileList) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### GetFiles

`func (o *FileList) GetFiles() []FileItem`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileList) GetFilesOk() (*[]FileItem, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileList) SetFiles(v []FileItem)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FileList) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFolder

`func (o *FileList) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *FileList) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *FileList) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *FileList) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetSpace

`func (o *FileList) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *FileList) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *FileList) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *FileList) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTotal

`func (o *FileList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FileList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FileList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FileList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


