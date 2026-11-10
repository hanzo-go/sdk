# UsageKai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the warehouse was not connected or a query blipped. The totals below are then honest zeros, NOT measured ones. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is what they cost the org, in US cents. | [optional] 
**Decisions** | Pointer to **int64** | Decisions is how many decisions Kai answered in the window. | [optional] 
**InputTokens** | Pointer to **int64** | InputTokens is the input tokens those decisions were billed on. | [optional] 

## Methods

### NewUsageKai

`func NewUsageKai() *UsageKai`

NewUsageKai instantiates a new UsageKai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageKaiWithDefaults

`func NewUsageKaiWithDefaults() *UsageKai`

NewUsageKaiWithDefaults instantiates a new UsageKai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *UsageKai) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *UsageKai) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *UsageKai) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *UsageKai) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCostCents

`func (o *UsageKai) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *UsageKai) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *UsageKai) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *UsageKai) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetDecisions

`func (o *UsageKai) GetDecisions() int64`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *UsageKai) GetDecisionsOk() (*int64, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *UsageKai) SetDecisions(v int64)`

SetDecisions sets Decisions field to given value.

### HasDecisions

`func (o *UsageKai) HasDecisions() bool`

HasDecisions returns a boolean if a field has been set.

### GetInputTokens

`func (o *UsageKai) GetInputTokens() int64`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *UsageKai) GetInputTokensOk() (*int64, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *UsageKai) SetInputTokens(v int64)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *UsageKai) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


