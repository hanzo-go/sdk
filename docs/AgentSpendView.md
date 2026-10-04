# AgentSpendView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByComponent** | Pointer to **map[string]int64** | ByComponent holds only the components with spend. Absent, not zero. | [optional] 
**CapMicroUsd** | Pointer to **int64** | CapMicroUSD is the total this agent may spend within one Period, as an integer number of micro-USD (1,000,000 &#x3D; $1). It is required at creation: a cap of zero would mean no limit and no per-agent spend record at all. | [optional] 
**ConsumedMicroUsd** | Pointer to **int64** | ConsumedMicroUSD is what has been spent in the current period, in micro-USD. It is settled from what the gateway reported, not from the quote. | [optional] 
**MaxTaskMicroUsd** | Pointer to **int64** | MaxTaskMicroUSD is the ceiling for a single run, in micro-USD. A session cannot exceed it even when the period cap still has room, so one runaway task cannot consume a month. | [optional] 
**Period** | Pointer to **string** | Period is the window the cap resets on: day, week or month. | [optional] 
**PeriodStartedAt** | Pointer to **int64** | PeriodStartedAt is when the current period began, as a Unix second. | [optional] 
**Ref** | Pointer to **string** | Ref names the agent this spend belongs to. | [optional] 
**RemainingMicroUsd** | Pointer to **int64** | RemainingMicroUSD is what the cap still allows this period, in micro-USD. | [optional] 

## Methods

### NewAgentSpendView

`func NewAgentSpendView() *AgentSpendView`

NewAgentSpendView instantiates a new AgentSpendView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSpendViewWithDefaults

`func NewAgentSpendViewWithDefaults() *AgentSpendView`

NewAgentSpendViewWithDefaults instantiates a new AgentSpendView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByComponent

`func (o *AgentSpendView) GetByComponent() map[string]int64`

GetByComponent returns the ByComponent field if non-nil, zero value otherwise.

### GetByComponentOk

`func (o *AgentSpendView) GetByComponentOk() (*map[string]int64, bool)`

GetByComponentOk returns a tuple with the ByComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByComponent

`func (o *AgentSpendView) SetByComponent(v map[string]int64)`

SetByComponent sets ByComponent field to given value.

### HasByComponent

`func (o *AgentSpendView) HasByComponent() bool`

HasByComponent returns a boolean if a field has been set.

### GetCapMicroUsd

`func (o *AgentSpendView) GetCapMicroUsd() int64`

GetCapMicroUsd returns the CapMicroUsd field if non-nil, zero value otherwise.

### GetCapMicroUsdOk

`func (o *AgentSpendView) GetCapMicroUsdOk() (*int64, bool)`

GetCapMicroUsdOk returns a tuple with the CapMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapMicroUsd

`func (o *AgentSpendView) SetCapMicroUsd(v int64)`

SetCapMicroUsd sets CapMicroUsd field to given value.

### HasCapMicroUsd

`func (o *AgentSpendView) HasCapMicroUsd() bool`

HasCapMicroUsd returns a boolean if a field has been set.

### GetConsumedMicroUsd

`func (o *AgentSpendView) GetConsumedMicroUsd() int64`

GetConsumedMicroUsd returns the ConsumedMicroUsd field if non-nil, zero value otherwise.

### GetConsumedMicroUsdOk

`func (o *AgentSpendView) GetConsumedMicroUsdOk() (*int64, bool)`

GetConsumedMicroUsdOk returns a tuple with the ConsumedMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedMicroUsd

`func (o *AgentSpendView) SetConsumedMicroUsd(v int64)`

SetConsumedMicroUsd sets ConsumedMicroUsd field to given value.

### HasConsumedMicroUsd

`func (o *AgentSpendView) HasConsumedMicroUsd() bool`

HasConsumedMicroUsd returns a boolean if a field has been set.

### GetMaxTaskMicroUsd

`func (o *AgentSpendView) GetMaxTaskMicroUsd() int64`

GetMaxTaskMicroUsd returns the MaxTaskMicroUsd field if non-nil, zero value otherwise.

### GetMaxTaskMicroUsdOk

`func (o *AgentSpendView) GetMaxTaskMicroUsdOk() (*int64, bool)`

GetMaxTaskMicroUsdOk returns a tuple with the MaxTaskMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskMicroUsd

`func (o *AgentSpendView) SetMaxTaskMicroUsd(v int64)`

SetMaxTaskMicroUsd sets MaxTaskMicroUsd field to given value.

### HasMaxTaskMicroUsd

`func (o *AgentSpendView) HasMaxTaskMicroUsd() bool`

HasMaxTaskMicroUsd returns a boolean if a field has been set.

### GetPeriod

`func (o *AgentSpendView) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AgentSpendView) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AgentSpendView) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AgentSpendView) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPeriodStartedAt

`func (o *AgentSpendView) GetPeriodStartedAt() int64`

GetPeriodStartedAt returns the PeriodStartedAt field if non-nil, zero value otherwise.

### GetPeriodStartedAtOk

`func (o *AgentSpendView) GetPeriodStartedAtOk() (*int64, bool)`

GetPeriodStartedAtOk returns a tuple with the PeriodStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartedAt

`func (o *AgentSpendView) SetPeriodStartedAt(v int64)`

SetPeriodStartedAt sets PeriodStartedAt field to given value.

### HasPeriodStartedAt

`func (o *AgentSpendView) HasPeriodStartedAt() bool`

HasPeriodStartedAt returns a boolean if a field has been set.

### GetRef

`func (o *AgentSpendView) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *AgentSpendView) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *AgentSpendView) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *AgentSpendView) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRemainingMicroUsd

`func (o *AgentSpendView) GetRemainingMicroUsd() int64`

GetRemainingMicroUsd returns the RemainingMicroUsd field if non-nil, zero value otherwise.

### GetRemainingMicroUsdOk

`func (o *AgentSpendView) GetRemainingMicroUsdOk() (*int64, bool)`

GetRemainingMicroUsdOk returns a tuple with the RemainingMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingMicroUsd

`func (o *AgentSpendView) SetRemainingMicroUsd(v int64)`

SetRemainingMicroUsd sets RemainingMicroUsd field to given value.

### HasRemainingMicroUsd

`func (o *AgentSpendView) HasRemainingMicroUsd() bool`

HasRemainingMicroUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


