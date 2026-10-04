# AgentCodingTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]AgentCodingEntry**](AgentCodingEntry.md) | Entries are the directory&#39;s immediate children, directories first. | [optional] 
**Ref** | Pointer to **string** | Ref is the branch this was read at: the run&#39;s own once the forge holds it, its base until then. | [optional] 

## Methods

### NewAgentCodingTree

`func NewAgentCodingTree() *AgentCodingTree`

NewAgentCodingTree instantiates a new AgentCodingTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingTreeWithDefaults

`func NewAgentCodingTreeWithDefaults() *AgentCodingTree`

NewAgentCodingTreeWithDefaults instantiates a new AgentCodingTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *AgentCodingTree) GetEntries() []AgentCodingEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AgentCodingTree) GetEntriesOk() (*[]AgentCodingEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AgentCodingTree) SetEntries(v []AgentCodingEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AgentCodingTree) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetRef

`func (o *AgentCodingTree) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *AgentCodingTree) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *AgentCodingTree) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *AgentCodingTree) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


