# ChatFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Filepath** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChatFile

`func NewChatFile() *ChatFile`

NewChatFile instantiates a new ChatFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatFileWithDefaults

`func NewChatFileWithDefaults() *ChatFile`

NewChatFileWithDefaults instantiates a new ChatFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *ChatFile) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ChatFile) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ChatFile) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ChatFile) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetUser

`func (o *ChatFile) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatFile) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatFile) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ChatFile) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetFilename

`func (o *ChatFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ChatFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ChatFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ChatFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFilepath

`func (o *ChatFile) GetFilepath() string`

GetFilepath returns the Filepath field if non-nil, zero value otherwise.

### GetFilepathOk

`func (o *ChatFile) GetFilepathOk() (*string, bool)`

GetFilepathOk returns a tuple with the Filepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilepath

`func (o *ChatFile) SetFilepath(v string)`

SetFilepath sets Filepath field to given value.

### HasFilepath

`func (o *ChatFile) HasFilepath() bool`

HasFilepath returns a boolean if a field has been set.

### GetType

`func (o *ChatFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChatFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHeight

`func (o *ChatFile) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ChatFile) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ChatFile) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ChatFile) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *ChatFile) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ChatFile) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ChatFile) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ChatFile) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetSize

`func (o *ChatFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ChatFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ChatFile) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ChatFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *ChatFile) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChatFile) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChatFile) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ChatFile) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChatFile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatFile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatFile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChatFile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


