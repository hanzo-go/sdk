# LinkReadingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider-side account the sample belongs to. | [optional] 
**CachedInputTokens** | Pointer to **int64** | CachedInputTokens is the window&#39;s cached-prompt-token count. | [optional] 
**Confidence** | Pointer to **string** | Confidence says whether the counters that remain mean anything, as the meter graded itself. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is the window&#39;s spend in cents, as the provider&#39;s meter states it. | [optional] 
**CostLimitCents** | Pointer to **int64** | CostLimitCents is the window&#39;s spend cap in cents, when the meter knows one. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO currency the cost fields are stated in. | [optional] 
**InputTokens** | Pointer to **int64** | InputTokens is the window&#39;s prompt-token count. | [optional] 
**Lane** | Pointer to **string** | Lane names the meter&#39;s own lane label for this measurement. | [optional] 
**Machine** | Pointer to **string** | Machine is the machine the collector observed the account on. | [optional] 
**OutputTokens** | Pointer to **int64** | OutputTokens is the window&#39;s completion-token count. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan label the account is on. | [optional] 
**Requests** | Pointer to **int64** | Requests is the window&#39;s request count. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when the window resets, RFC 3339 UTC. | [optional] 
**Synthetic** | Pointer to **bool** | Synthetic marks a sample the collector derived rather than observed. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is the window&#39;s total token count. | [optional] 
**UsedPct** | Pointer to **float64** | UsedPct is how much of the window&#39;s allowance is consumed, 0..100. | [optional] 
**Window** | Pointer to **string** | Window is the window class: 6h, day, week or month. | [optional] 
**WindowMinutes** | Pointer to **int32** | WindowMinutes is the window&#39;s length as the meter reported it. | [optional] 
**WindowStart** | Pointer to **string** | WindowStart is when the measured window opened, RFC 3339 UTC. | [optional] 

## Methods

### NewLinkReadingView

`func NewLinkReadingView() *LinkReadingView`

NewLinkReadingView instantiates a new LinkReadingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkReadingViewWithDefaults

`func NewLinkReadingViewWithDefaults() *LinkReadingView`

NewLinkReadingViewWithDefaults instantiates a new LinkReadingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LinkReadingView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LinkReadingView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LinkReadingView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LinkReadingView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *LinkReadingView) GetCachedInputTokens() int64`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *LinkReadingView) GetCachedInputTokensOk() (*int64, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *LinkReadingView) SetCachedInputTokens(v int64)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *LinkReadingView) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *LinkReadingView) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *LinkReadingView) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *LinkReadingView) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *LinkReadingView) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *LinkReadingView) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LinkReadingView) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LinkReadingView) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LinkReadingView) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *LinkReadingView) GetCostLimitCents() int64`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *LinkReadingView) GetCostLimitCentsOk() (*int64, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *LinkReadingView) SetCostLimitCents(v int64)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *LinkReadingView) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *LinkReadingView) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LinkReadingView) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LinkReadingView) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LinkReadingView) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *LinkReadingView) GetInputTokens() int64`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *LinkReadingView) GetInputTokensOk() (*int64, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *LinkReadingView) SetInputTokens(v int64)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *LinkReadingView) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetLane

`func (o *LinkReadingView) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *LinkReadingView) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *LinkReadingView) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *LinkReadingView) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *LinkReadingView) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *LinkReadingView) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *LinkReadingView) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *LinkReadingView) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *LinkReadingView) GetOutputTokens() int64`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *LinkReadingView) GetOutputTokensOk() (*int64, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *LinkReadingView) SetOutputTokens(v int64)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *LinkReadingView) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *LinkReadingView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LinkReadingView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LinkReadingView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LinkReadingView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRequests

`func (o *LinkReadingView) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LinkReadingView) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LinkReadingView) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LinkReadingView) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *LinkReadingView) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *LinkReadingView) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *LinkReadingView) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *LinkReadingView) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSynthetic

`func (o *LinkReadingView) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *LinkReadingView) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *LinkReadingView) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *LinkReadingView) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *LinkReadingView) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *LinkReadingView) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *LinkReadingView) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *LinkReadingView) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *LinkReadingView) GetUsedPct() float64`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *LinkReadingView) GetUsedPctOk() (*float64, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *LinkReadingView) SetUsedPct(v float64)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *LinkReadingView) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *LinkReadingView) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *LinkReadingView) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *LinkReadingView) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *LinkReadingView) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *LinkReadingView) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *LinkReadingView) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *LinkReadingView) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *LinkReadingView) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *LinkReadingView) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *LinkReadingView) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *LinkReadingView) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *LinkReadingView) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


