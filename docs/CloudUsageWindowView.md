# CloudUsageWindowView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the linked provider account the window belongs to. | [optional] 
**CachedInputTokens** | Pointer to **int32** | CachedInputTokens is the prompt tokens served from the provider&#39;s cache; omitted when unknown. | [optional] 
**Confidence** | Pointer to **string** | Confidence says how much the counters beside it mean — a meter that reported only a percentage leaves them at zero, and this is how a reader tells that from a true zero. | [optional] 
**CostCents** | Pointer to **int32** | CostCents is what the window cost on the PROVIDER&#39;s own plan, in US cents. It is not a Hanzo charge. | [optional] 
**CostLimitCents** | Pointer to **int32** | CostLimitCents is the plan&#39;s spend ceiling for the window, in US cents. | [optional] 
**Currency** | Pointer to **string** | Currency is the provider&#39;s currency when it is not US cents. | [optional] 
**InputTokens** | Pointer to **int32** | InputTokens is prompt tokens consumed in the window; omitted when unknown. | [optional] 
**Lane** | Pointer to **string** | Lane is the meter lane this instance belongs to, e.g. a provider&#39;s own rolling-window meter. | [optional] 
**Machine** | Pointer to **string** | Machine is the host whose meter reported the window. | [optional] 
**OutputTokens** | Pointer to **int32** | OutputTokens is completion tokens produced in the window; omitted when unknown. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan the account is on, e.g. a Claude Max plan. | [optional] 
**Requests** | Pointer to **int32** | Requests is how many requests were made in the window; omitted when the meter did not report it. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when this window rolls over, RFC3339 UTC; omitted when unknown. | [optional] 
**Synthetic** | Pointer to **bool** | Synthetic marks an instance the meter inferred rather than read. | [optional] 
**TotalTokens** | Pointer to **int32** | TotalTokens is the window&#39;s total tokens; omitted when unknown. | [optional] 
**UsedPct** | Pointer to **float32** | UsedPct is how much of the window&#39;s allowance is consumed, 0–100. | [optional] 
**Window** | Pointer to **string** | Window is the window class: 6h, day, week or month. | [optional] 
**WindowMinutes** | Pointer to **int32** | WindowMinutes is the window&#39;s real length in minutes when the meter reported one; omitted when it did not. | [optional] 
**WindowStart** | Pointer to **string** | WindowStart is when this window opened, RFC3339 UTC; omitted when unknown. | [optional] 

## Methods

### NewCloudUsageWindowView

`func NewCloudUsageWindowView() *CloudUsageWindowView`

NewCloudUsageWindowView instantiates a new CloudUsageWindowView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUsageWindowViewWithDefaults

`func NewCloudUsageWindowViewWithDefaults() *CloudUsageWindowView`

NewCloudUsageWindowViewWithDefaults instantiates a new CloudUsageWindowView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudUsageWindowView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudUsageWindowView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudUsageWindowView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudUsageWindowView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *CloudUsageWindowView) GetCachedInputTokens() int32`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *CloudUsageWindowView) GetCachedInputTokensOk() (*int32, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *CloudUsageWindowView) SetCachedInputTokens(v int32)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *CloudUsageWindowView) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *CloudUsageWindowView) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CloudUsageWindowView) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CloudUsageWindowView) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CloudUsageWindowView) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *CloudUsageWindowView) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *CloudUsageWindowView) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *CloudUsageWindowView) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *CloudUsageWindowView) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *CloudUsageWindowView) GetCostLimitCents() int32`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *CloudUsageWindowView) GetCostLimitCentsOk() (*int32, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *CloudUsageWindowView) SetCostLimitCents(v int32)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *CloudUsageWindowView) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudUsageWindowView) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudUsageWindowView) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudUsageWindowView) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudUsageWindowView) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *CloudUsageWindowView) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *CloudUsageWindowView) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *CloudUsageWindowView) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *CloudUsageWindowView) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetLane

`func (o *CloudUsageWindowView) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *CloudUsageWindowView) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *CloudUsageWindowView) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *CloudUsageWindowView) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *CloudUsageWindowView) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *CloudUsageWindowView) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *CloudUsageWindowView) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *CloudUsageWindowView) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *CloudUsageWindowView) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *CloudUsageWindowView) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *CloudUsageWindowView) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *CloudUsageWindowView) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *CloudUsageWindowView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CloudUsageWindowView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CloudUsageWindowView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CloudUsageWindowView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRequests

`func (o *CloudUsageWindowView) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudUsageWindowView) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudUsageWindowView) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudUsageWindowView) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *CloudUsageWindowView) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *CloudUsageWindowView) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *CloudUsageWindowView) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *CloudUsageWindowView) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSynthetic

`func (o *CloudUsageWindowView) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *CloudUsageWindowView) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *CloudUsageWindowView) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *CloudUsageWindowView) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *CloudUsageWindowView) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *CloudUsageWindowView) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *CloudUsageWindowView) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *CloudUsageWindowView) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *CloudUsageWindowView) GetUsedPct() float32`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *CloudUsageWindowView) GetUsedPctOk() (*float32, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *CloudUsageWindowView) SetUsedPct(v float32)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *CloudUsageWindowView) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *CloudUsageWindowView) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *CloudUsageWindowView) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *CloudUsageWindowView) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *CloudUsageWindowView) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *CloudUsageWindowView) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *CloudUsageWindowView) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *CloudUsageWindowView) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *CloudUsageWindowView) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *CloudUsageWindowView) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *CloudUsageWindowView) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *CloudUsageWindowView) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *CloudUsageWindowView) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


