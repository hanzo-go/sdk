# AgentCodingEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the entry&#39;s own name, no directory part. | [optional] 
**Path** | Pointer to **string** | Path is the entry&#39;s full repo-relative path. | [optional] 
**Size** | Pointer to **int64** | Size is a file&#39;s byte length; 0 for a directory. | [optional] 
**Type** | Pointer to **string** | Type is \&quot;tree\&quot; for a directory and \&quot;blob\&quot; for anything else a reader opens. | [optional] 

## Methods

### NewAgentCodingEntry

`func NewAgentCodingEntry() *AgentCodingEntry`

NewAgentCodingEntry instantiates a new AgentCodingEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingEntryWithDefaults

`func NewAgentCodingEntryWithDefaults() *AgentCodingEntry`

NewAgentCodingEntryWithDefaults instantiates a new AgentCodingEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentCodingEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentCodingEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentCodingEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentCodingEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *AgentCodingEntry) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AgentCodingEntry) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AgentCodingEntry) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AgentCodingEntry) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *AgentCodingEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AgentCodingEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AgentCodingEntry) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *AgentCodingEntry) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *AgentCodingEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentCodingEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentCodingEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentCodingEntry) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


