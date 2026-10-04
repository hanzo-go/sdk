# TeamTeamFileIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to **string** | File is the blob id — the uuid the file was uploaded under with POST /v1/team/files/{space}. A blob of another space cannot be named: the download is keyed by this message&#39;s space, so a foreign id resolves to nothing. | [optional] 
**Name** | Pointer to **string** | Name is the file&#39;s name as a person should see it. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s length in bytes. | [optional] 
**Type** | Pointer to **string** | Type is the file&#39;s media type, e.g. \&quot;image/png\&quot;. | [optional] 

## Methods

### NewTeamTeamFileIn

`func NewTeamTeamFileIn() *TeamTeamFileIn`

NewTeamTeamFileIn instantiates a new TeamTeamFileIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamFileInWithDefaults

`func NewTeamTeamFileInWithDefaults() *TeamTeamFileIn`

NewTeamTeamFileInWithDefaults instantiates a new TeamTeamFileIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *TeamTeamFileIn) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TeamTeamFileIn) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TeamTeamFileIn) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *TeamTeamFileIn) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamFileIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamFileIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamFileIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamFileIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *TeamTeamFileIn) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TeamTeamFileIn) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TeamTeamFileIn) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TeamTeamFileIn) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *TeamTeamFileIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamTeamFileIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamTeamFileIn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TeamTeamFileIn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


