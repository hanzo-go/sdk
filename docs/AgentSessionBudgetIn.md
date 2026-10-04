# AgentSessionBudgetIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetMicroUsd** | Pointer to **int64** | BudgetMicroUSD is the new cap, or null to remove the cap for good. | [optional] 
**Id** | Pointer to **string** | ID is the session, from the path. | [optional] 

## Methods

### NewAgentSessionBudgetIn

`func NewAgentSessionBudgetIn() *AgentSessionBudgetIn`

NewAgentSessionBudgetIn instantiates a new AgentSessionBudgetIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSessionBudgetInWithDefaults

`func NewAgentSessionBudgetInWithDefaults() *AgentSessionBudgetIn`

NewAgentSessionBudgetInWithDefaults instantiates a new AgentSessionBudgetIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetMicroUsd

`func (o *AgentSessionBudgetIn) GetBudgetMicroUsd() int64`

GetBudgetMicroUsd returns the BudgetMicroUsd field if non-nil, zero value otherwise.

### GetBudgetMicroUsdOk

`func (o *AgentSessionBudgetIn) GetBudgetMicroUsdOk() (*int64, bool)`

GetBudgetMicroUsdOk returns a tuple with the BudgetMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetMicroUsd

`func (o *AgentSessionBudgetIn) SetBudgetMicroUsd(v int64)`

SetBudgetMicroUsd sets BudgetMicroUsd field to given value.

### HasBudgetMicroUsd

`func (o *AgentSessionBudgetIn) HasBudgetMicroUsd() bool`

HasBudgetMicroUsd returns a boolean if a field has been set.

### GetId

`func (o *AgentSessionBudgetIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentSessionBudgetIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentSessionBudgetIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentSessionBudgetIn) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


