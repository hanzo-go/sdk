# AgentCodingFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Additions** | Pointer to **int64** | Additions counts the lines the change adds, over the whole file whether or not all of its patch is carried. | [optional] 
**Deletions** | Pointer to **int64** | Deletions counts the lines the change removes, the same way. | [optional] 
**From** | Pointer to **string** | From is where a renamed file was. Omitted for every other status. | [optional] 
**Patch** | Pointer to **string** | Patch is the file&#39;s unified-diff hunks, from its first \&quot;@@\&quot;. Empty for a binary file and for a rename that moved no line. | [optional] 
**Path** | Pointer to **string** | Path is where the file is on the branch — for a deleted file, where it was. | [optional] 
**Status** | Pointer to **string** | Status is added, modified, deleted or renamed. | [optional] 
**Truncated** | Pointer to **bool** | Truncated marks a patch cut short: at 64 KiB for one file, or past 1 MiB of patch across the change, when later files carry none. The counts are whole. | [optional] 

## Methods

### NewAgentCodingFile

`func NewAgentCodingFile() *AgentCodingFile`

NewAgentCodingFile instantiates a new AgentCodingFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingFileWithDefaults

`func NewAgentCodingFileWithDefaults() *AgentCodingFile`

NewAgentCodingFileWithDefaults instantiates a new AgentCodingFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditions

`func (o *AgentCodingFile) GetAdditions() int64`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *AgentCodingFile) GetAdditionsOk() (*int64, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *AgentCodingFile) SetAdditions(v int64)`

SetAdditions sets Additions field to given value.

### HasAdditions

`func (o *AgentCodingFile) HasAdditions() bool`

HasAdditions returns a boolean if a field has been set.

### GetDeletions

`func (o *AgentCodingFile) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *AgentCodingFile) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *AgentCodingFile) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *AgentCodingFile) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetFrom

`func (o *AgentCodingFile) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AgentCodingFile) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AgentCodingFile) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *AgentCodingFile) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetPatch

`func (o *AgentCodingFile) GetPatch() string`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *AgentCodingFile) GetPatchOk() (*string, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *AgentCodingFile) SetPatch(v string)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *AgentCodingFile) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetPath

`func (o *AgentCodingFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AgentCodingFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AgentCodingFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AgentCodingFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStatus

`func (o *AgentCodingFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentCodingFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentCodingFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentCodingFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTruncated

`func (o *AgentCodingFile) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *AgentCodingFile) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *AgentCodingFile) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *AgentCodingFile) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


