# TreeEntryJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Mode is the octal git file mode (\&quot;100644\&quot;, \&quot;040000\&quot;, \&quot;120000\&quot;). | [optional] 
**Name** | Pointer to **string** | Name is the entry&#39;s own name, no directory part. | [optional] 
**Path** | Pointer to **string** | Path is the entry&#39;s full repo-relative path. | [optional] 
**Size** | Pointer to **int32** | Size is the file&#39;s byte length; 0 for a directory. | [optional] 
**Type** | Pointer to **string** | Type is \&quot;tree\&quot; for a directory, \&quot;blob\&quot; for a file. | [optional] 

## Methods

### NewTreeEntryJSON

`func NewTreeEntryJSON() *TreeEntryJSON`

NewTreeEntryJSON instantiates a new TreeEntryJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreeEntryJSONWithDefaults

`func NewTreeEntryJSONWithDefaults() *TreeEntryJSON`

NewTreeEntryJSONWithDefaults instantiates a new TreeEntryJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *TreeEntryJSON) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TreeEntryJSON) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TreeEntryJSON) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TreeEntryJSON) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *TreeEntryJSON) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TreeEntryJSON) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TreeEntryJSON) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TreeEntryJSON) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *TreeEntryJSON) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TreeEntryJSON) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TreeEntryJSON) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TreeEntryJSON) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *TreeEntryJSON) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TreeEntryJSON) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TreeEntryJSON) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TreeEntryJSON) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *TreeEntryJSON) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TreeEntryJSON) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TreeEntryJSON) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TreeEntryJSON) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


