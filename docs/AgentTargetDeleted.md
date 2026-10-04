# AgentTargetDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | Deleted is true when the target was removed. | [optional] 
**Id** | Pointer to **string** | ID is the target that was removed. | [optional] 

## Methods

### NewAgentTargetDeleted

`func NewAgentTargetDeleted() *AgentTargetDeleted`

NewAgentTargetDeleted instantiates a new AgentTargetDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTargetDeletedWithDefaults

`func NewAgentTargetDeletedWithDefaults() *AgentTargetDeleted`

NewAgentTargetDeletedWithDefaults instantiates a new AgentTargetDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *AgentTargetDeleted) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *AgentTargetDeleted) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *AgentTargetDeleted) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *AgentTargetDeleted) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *AgentTargetDeleted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentTargetDeleted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentTargetDeleted) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentTargetDeleted) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


