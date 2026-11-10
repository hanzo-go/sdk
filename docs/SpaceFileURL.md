# SpaceFileURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int64** | Expiry is how many seconds the URL stays valid. A presigned URL has no server-side revocation, so this IS its revocation window. | [optional] 
**File** | Pointer to **string** | File is the file name the URL was signed for, relative to the drive and path-cleaned — so it is what the store will actually read or write, which is not always the string the caller sent. The signature covers this one space, this one drive and this one file: a URL minted here reaches nothing else. | [optional] 
**Method** | Pointer to **string** | Method is the verb the URL is signed for — \&quot;GET\&quot; to read, \&quot;PUT\&quot; to write. | [optional] 
**Url** | Pointer to **string** | URL is the presigned URL, signed against the PUBLIC host so a browser can follow it. | [optional] 

## Methods

### NewSpaceFileURL

`func NewSpaceFileURL() *SpaceFileURL`

NewSpaceFileURL instantiates a new SpaceFileURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceFileURLWithDefaults

`func NewSpaceFileURLWithDefaults() *SpaceFileURL`

NewSpaceFileURLWithDefaults instantiates a new SpaceFileURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *SpaceFileURL) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *SpaceFileURL) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *SpaceFileURL) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *SpaceFileURL) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetFile

`func (o *SpaceFileURL) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SpaceFileURL) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SpaceFileURL) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *SpaceFileURL) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetMethod

`func (o *SpaceFileURL) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SpaceFileURL) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SpaceFileURL) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SpaceFileURL) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *SpaceFileURL) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SpaceFileURL) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SpaceFileURL) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SpaceFileURL) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


