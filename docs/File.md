# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**ErrorText** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**TokenCount** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewFile

`func NewFile() *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *File) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *File) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *File) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *File) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetErrorText

`func (o *File) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *File) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *File) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *File) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetFilename

`func (o *File) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *File) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *File) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *File) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetName

`func (o *File) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *File) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *File) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *File) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *File) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *File) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *File) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *File) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSize

`func (o *File) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *File) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *File) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *File) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *File) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *File) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *File) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *File) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageProvider

`func (o *File) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *File) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *File) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *File) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetStore

`func (o *File) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *File) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *File) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *File) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTokenCount

`func (o *File) GetTokenCount() int32`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *File) GetTokenCountOk() (*int32, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *File) SetTokenCount(v int32)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *File) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetUrl

`func (o *File) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *File) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *File) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *File) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


