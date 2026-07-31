# CloudObjectFileInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** |  | 
**Name** | **string** |  | 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**TokenCount** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorText** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudObjectFileInput

`func NewCloudObjectFileInput(owner string, name string, ) *CloudObjectFileInput`

NewCloudObjectFileInput instantiates a new CloudObjectFileInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectFileInputWithDefaults

`func NewCloudObjectFileInputWithDefaults() *CloudObjectFileInput`

NewCloudObjectFileInputWithDefaults instantiates a new CloudObjectFileInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *CloudObjectFileInput) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectFileInput) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectFileInput) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetName

`func (o *CloudObjectFileInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectFileInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectFileInput) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedTime

`func (o *CloudObjectFileInput) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectFileInput) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectFileInput) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectFileInput) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetFilename

`func (o *CloudObjectFileInput) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CloudObjectFileInput) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CloudObjectFileInput) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CloudObjectFileInput) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetSize

`func (o *CloudObjectFileInput) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudObjectFileInput) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudObjectFileInput) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CloudObjectFileInput) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStore

`func (o *CloudObjectFileInput) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CloudObjectFileInput) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CloudObjectFileInput) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *CloudObjectFileInput) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetStorageProvider

`func (o *CloudObjectFileInput) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *CloudObjectFileInput) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *CloudObjectFileInput) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *CloudObjectFileInput) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetUrl

`func (o *CloudObjectFileInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudObjectFileInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudObjectFileInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudObjectFileInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTokenCount

`func (o *CloudObjectFileInput) GetTokenCount() int32`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *CloudObjectFileInput) GetTokenCountOk() (*int32, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *CloudObjectFileInput) SetTokenCount(v int32)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *CloudObjectFileInput) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetStatus

`func (o *CloudObjectFileInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudObjectFileInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudObjectFileInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudObjectFileInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorText

`func (o *CloudObjectFileInput) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *CloudObjectFileInput) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *CloudObjectFileInput) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *CloudObjectFileInput) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


