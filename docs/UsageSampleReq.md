# UsageSampleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the linked account the window was metered from. | [optional] 
**CachedInputTokens** | Pointer to **int64** | CachedInputTokens is the prompt tokens the provider served from cache. | [optional] 
**Confidence** | Pointer to **string** | Confidence says how much the counters below mean. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is what the window cost on the PROVIDER&#39;s own plan, in US cents. | [optional] 
**CostLimitCents** | Pointer to **int64** | CostLimitCents is the plan&#39;s spend ceiling for the window, in US cents. | [optional] 
**Currency** | Pointer to **string** | Currency is the provider&#39;s currency when it is not US cents. | [optional] 
**InputTokens** | Pointer to **int64** | InputTokens is prompt tokens consumed in the window. | [optional] 
**Kind** | Pointer to **string** | Kind is subscription or apikey. Empty is accepted; anything else is refused. | [optional] 
**Lane** | Pointer to **string** | Lane is the meter lane within the account. | [optional] 
**Machine** | Pointer to **string** | Machine is the host whose meter read the window. Required. | [optional] 
**OutputTokens** | Pointer to **int64** | OutputTokens is completion tokens produced in the window. | [optional] 
**Plan** | Pointer to **string** | Plan is the subscription plan the account is on, as the provider names it. | [optional] 
**Provider** | Pointer to **string** | Provider is the upstream the account belongs to, e.g. anthropic. Required. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many requests the window covers. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when the measured window rolls over, RFC3339. Empty is allowed; anything else that is not RFC3339 is refused. | [optional] 
**Synthetic** | Pointer to **bool** | Synthetic marks a window the meter inferred rather than read. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is the window&#39;s total tokens. | [optional] 
**UsedPct** | Pointer to **float64** | UsedPct is how much of the window&#39;s allowance is consumed, 0–100. | [optional] 
**Window** | Pointer to **string** | Window is the window class: 6h, day, week or month. Required, and a class this surface does not know is refused rather than rewritten. | [optional] 
**WindowMinutes** | Pointer to **int32** | WindowMinutes is the window&#39;s real length in minutes, as the meter reports it. | [optional] 
**WindowStart** | Pointer to **string** | WindowStart is when the measured window opened, RFC3339. Empty is allowed; anything else that is not RFC3339 is refused. | [optional] 

## Methods

### NewUsageSampleReq

`func NewUsageSampleReq() *UsageSampleReq`

NewUsageSampleReq instantiates a new UsageSampleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSampleReqWithDefaults

`func NewUsageSampleReqWithDefaults() *UsageSampleReq`

NewUsageSampleReqWithDefaults instantiates a new UsageSampleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UsageSampleReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UsageSampleReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UsageSampleReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UsageSampleReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *UsageSampleReq) GetCachedInputTokens() int64`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *UsageSampleReq) GetCachedInputTokensOk() (*int64, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *UsageSampleReq) SetCachedInputTokens(v int64)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *UsageSampleReq) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *UsageSampleReq) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *UsageSampleReq) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *UsageSampleReq) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *UsageSampleReq) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *UsageSampleReq) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *UsageSampleReq) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *UsageSampleReq) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *UsageSampleReq) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *UsageSampleReq) GetCostLimitCents() int64`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *UsageSampleReq) GetCostLimitCentsOk() (*int64, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *UsageSampleReq) SetCostLimitCents(v int64)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *UsageSampleReq) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *UsageSampleReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UsageSampleReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UsageSampleReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UsageSampleReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *UsageSampleReq) GetInputTokens() int64`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *UsageSampleReq) GetInputTokensOk() (*int64, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *UsageSampleReq) SetInputTokens(v int64)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *UsageSampleReq) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetKind

`func (o *UsageSampleReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UsageSampleReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UsageSampleReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UsageSampleReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLane

`func (o *UsageSampleReq) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *UsageSampleReq) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *UsageSampleReq) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *UsageSampleReq) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *UsageSampleReq) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *UsageSampleReq) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *UsageSampleReq) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *UsageSampleReq) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *UsageSampleReq) GetOutputTokens() int64`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *UsageSampleReq) GetOutputTokensOk() (*int64, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *UsageSampleReq) SetOutputTokens(v int64)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *UsageSampleReq) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *UsageSampleReq) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UsageSampleReq) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UsageSampleReq) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UsageSampleReq) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *UsageSampleReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UsageSampleReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UsageSampleReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UsageSampleReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *UsageSampleReq) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageSampleReq) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageSampleReq) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *UsageSampleReq) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *UsageSampleReq) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *UsageSampleReq) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *UsageSampleReq) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *UsageSampleReq) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSynthetic

`func (o *UsageSampleReq) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *UsageSampleReq) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *UsageSampleReq) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *UsageSampleReq) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *UsageSampleReq) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *UsageSampleReq) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *UsageSampleReq) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *UsageSampleReq) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *UsageSampleReq) GetUsedPct() float64`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *UsageSampleReq) GetUsedPctOk() (*float64, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *UsageSampleReq) SetUsedPct(v float64)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *UsageSampleReq) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *UsageSampleReq) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *UsageSampleReq) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *UsageSampleReq) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *UsageSampleReq) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *UsageSampleReq) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *UsageSampleReq) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *UsageSampleReq) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *UsageSampleReq) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *UsageSampleReq) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *UsageSampleReq) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *UsageSampleReq) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *UsageSampleReq) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


