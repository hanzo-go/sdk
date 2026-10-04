# TeamTeamFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to **string** | File is the blob id, readable at GET /v1/team/files/{space}/{name}?file&#x3D;{file}. | [optional] 
**Id** | Pointer to **string** | ID is the attachment document&#39;s own id. | [optional] 
**Name** | Pointer to **string** | Name is the file&#39;s name as the person uploaded it. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s length in bytes, as the uploader stated it. | [optional] 
**Type** | Pointer to **string** | Type is the media type the uploader declared. The download route serves what the BYTES are, not this; it is a hint for drawing an icon. | [optional] 

## Methods

### NewTeamTeamFile

`func NewTeamTeamFile() *TeamTeamFile`

NewTeamTeamFile instantiates a new TeamTeamFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamFileWithDefaults

`func NewTeamTeamFileWithDefaults() *TeamTeamFile`

NewTeamTeamFileWithDefaults instantiates a new TeamTeamFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *TeamTeamFile) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TeamTeamFile) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TeamTeamFile) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *TeamTeamFile) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *TeamTeamFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TeamTeamFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TeamTeamFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TeamTeamFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *TeamTeamFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamTeamFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamTeamFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TeamTeamFile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


