# AgentSessionBudgetView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetMicroUsd** | Pointer to **int64** | BudgetMicroUSD is this session&#39;s own ceiling in micro-USD, beyond the agent&#39;s. A replacement must strictly exceed what the session has already consumed, and removing it is one-way. | [optional] 
**BudgetRemoved** | Pointer to **bool** | BudgetRemoved says the session&#39;s own cap was taken off. It cannot be put back. | [optional] 
**ConsumedMicroUsd** | Pointer to **int64** | ConsumedMicroUSD is what has been spent in the current period, in micro-USD. It is settled from what the gateway reported, not from the quote. | [optional] 
**Id** | Pointer to **string** | ID names the session. | [optional] 
**Status** | Pointer to **string** | Status is the session&#39;s state: running, paused at its cap, or done. | [optional] 

## Methods

### NewAgentSessionBudgetView

`func NewAgentSessionBudgetView() *AgentSessionBudgetView`

NewAgentSessionBudgetView instantiates a new AgentSessionBudgetView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSessionBudgetViewWithDefaults

`func NewAgentSessionBudgetViewWithDefaults() *AgentSessionBudgetView`

NewAgentSessionBudgetViewWithDefaults instantiates a new AgentSessionBudgetView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetMicroUsd

`func (o *AgentSessionBudgetView) GetBudgetMicroUsd() int64`

GetBudgetMicroUsd returns the BudgetMicroUsd field if non-nil, zero value otherwise.

### GetBudgetMicroUsdOk

`func (o *AgentSessionBudgetView) GetBudgetMicroUsdOk() (*int64, bool)`

GetBudgetMicroUsdOk returns a tuple with the BudgetMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetMicroUsd

`func (o *AgentSessionBudgetView) SetBudgetMicroUsd(v int64)`

SetBudgetMicroUsd sets BudgetMicroUsd field to given value.

### HasBudgetMicroUsd

`func (o *AgentSessionBudgetView) HasBudgetMicroUsd() bool`

HasBudgetMicroUsd returns a boolean if a field has been set.

### GetBudgetRemoved

`func (o *AgentSessionBudgetView) GetBudgetRemoved() bool`

GetBudgetRemoved returns the BudgetRemoved field if non-nil, zero value otherwise.

### GetBudgetRemovedOk

`func (o *AgentSessionBudgetView) GetBudgetRemovedOk() (*bool, bool)`

GetBudgetRemovedOk returns a tuple with the BudgetRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetRemoved

`func (o *AgentSessionBudgetView) SetBudgetRemoved(v bool)`

SetBudgetRemoved sets BudgetRemoved field to given value.

### HasBudgetRemoved

`func (o *AgentSessionBudgetView) HasBudgetRemoved() bool`

HasBudgetRemoved returns a boolean if a field has been set.

### GetConsumedMicroUsd

`func (o *AgentSessionBudgetView) GetConsumedMicroUsd() int64`

GetConsumedMicroUsd returns the ConsumedMicroUsd field if non-nil, zero value otherwise.

### GetConsumedMicroUsdOk

`func (o *AgentSessionBudgetView) GetConsumedMicroUsdOk() (*int64, bool)`

GetConsumedMicroUsdOk returns a tuple with the ConsumedMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedMicroUsd

`func (o *AgentSessionBudgetView) SetConsumedMicroUsd(v int64)`

SetConsumedMicroUsd sets ConsumedMicroUsd field to given value.

### HasConsumedMicroUsd

`func (o *AgentSessionBudgetView) HasConsumedMicroUsd() bool`

HasConsumedMicroUsd returns a boolean if a field has been set.

### GetId

`func (o *AgentSessionBudgetView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentSessionBudgetView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentSessionBudgetView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentSessionBudgetView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AgentSessionBudgetView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentSessionBudgetView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentSessionBudgetView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentSessionBudgetView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


