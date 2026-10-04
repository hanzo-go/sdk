# GitTreeEntryJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Mode is the octal git file mode (\&quot;100644\&quot;, \&quot;040000\&quot;, \&quot;120000\&quot;). | [optional] 
**Name** | Pointer to **string** | Name is the entry&#39;s own name, no directory part. | [optional] 
**Path** | Pointer to **string** | Path is the entry&#39;s full repo-relative path. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s byte length; 0 for a directory. | [optional] 
**Type** | Pointer to **string** | Type is \&quot;tree\&quot; for a directory, \&quot;blob\&quot; for a file. | [optional] 

## Methods

### NewGitTreeEntryJSON

`func NewGitTreeEntryJSON() *GitTreeEntryJSON`

NewGitTreeEntryJSON instantiates a new GitTreeEntryJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTreeEntryJSONWithDefaults

`func NewGitTreeEntryJSONWithDefaults() *GitTreeEntryJSON`

NewGitTreeEntryJSONWithDefaults instantiates a new GitTreeEntryJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GitTreeEntryJSON) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GitTreeEntryJSON) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GitTreeEntryJSON) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GitTreeEntryJSON) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GitTreeEntryJSON) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitTreeEntryJSON) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitTreeEntryJSON) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitTreeEntryJSON) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *GitTreeEntryJSON) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitTreeEntryJSON) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitTreeEntryJSON) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitTreeEntryJSON) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *GitTreeEntryJSON) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GitTreeEntryJSON) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GitTreeEntryJSON) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GitTreeEntryJSON) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *GitTreeEntryJSON) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitTreeEntryJSON) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitTreeEntryJSON) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GitTreeEntryJSON) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


