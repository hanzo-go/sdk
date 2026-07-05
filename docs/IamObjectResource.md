# IamObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FileFormat** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int64** |  | [optional] 
**FileType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectResource

`func NewIamObjectResource() *IamObjectResource`

NewIamObjectResource instantiates a new IamObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectResourceWithDefaults

`func NewIamObjectResourceWithDefaults() *IamObjectResource`

NewIamObjectResourceWithDefaults instantiates a new IamObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *IamObjectResource) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamObjectResource) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamObjectResource) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamObjectResource) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectResource) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectResource) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectResource) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectResource) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *IamObjectResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamObjectResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamObjectResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamObjectResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileFormat

`func (o *IamObjectResource) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *IamObjectResource) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *IamObjectResource) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *IamObjectResource) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetFileName

`func (o *IamObjectResource) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *IamObjectResource) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *IamObjectResource) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *IamObjectResource) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *IamObjectResource) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *IamObjectResource) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *IamObjectResource) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *IamObjectResource) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileType

`func (o *IamObjectResource) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *IamObjectResource) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *IamObjectResource) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *IamObjectResource) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetName

`func (o *IamObjectResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectResource) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectResource) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectResource) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectResource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *IamObjectResource) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IamObjectResource) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IamObjectResource) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IamObjectResource) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProvider

`func (o *IamObjectResource) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IamObjectResource) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IamObjectResource) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IamObjectResource) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetTag

`func (o *IamObjectResource) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamObjectResource) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamObjectResource) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamObjectResource) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUrl

`func (o *IamObjectResource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IamObjectResource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IamObjectResource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IamObjectResource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUser

`func (o *IamObjectResource) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamObjectResource) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamObjectResource) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamObjectResource) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


