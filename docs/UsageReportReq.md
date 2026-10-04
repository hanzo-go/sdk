# UsageReportReq

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
**Machine** | Pointer to **string** | Machine is the host whose meter read the window. Required on every sample. | [optional] 
**OutputTokens** | Pointer to **int64** | OutputTokens is completion tokens produced in the window. | [optional] 
**Plan** | Pointer to **string** | Plan is the subscription plan the account is on, as the provider names it. | [optional] 
**Provider** | Pointer to **string** | Provider is the upstream the account belongs to, e.g. anthropic. Required on every sample. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many requests the window covers. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when the measured window rolls over, RFC3339. | [optional] 
**Samples** | Pointer to [**[]UsageSampleReq**](UsageSampleReq.md) | Samples is the batch form: every lane a poller measured, in one call. When it is non-empty the top-level sample fields are ignored. | [optional] 
**Synthetic** | Pointer to **bool** | Synthetic marks a window the meter inferred rather than read. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is the window&#39;s total tokens. | [optional] 
**UsedPct** | Pointer to **float64** | UsedPct is how much of the window&#39;s allowance is consumed, 0–100. | [optional] 
**Window** | Pointer to **string** | Window is the window class: 6h, day, week or month. Required, and a class this surface does not know is refused rather than rewritten. | [optional] 
**WindowMinutes** | Pointer to **int32** | WindowMinutes is the window&#39;s real length in minutes, as the meter reports it. | [optional] 
**WindowStart** | Pointer to **string** | WindowStart is when the measured window opened, RFC3339. This is how a backfill states WHICH window it measured; the server always owns the observation clock, so there is no timestamp field. | [optional] 

## Methods

### NewUsageReportReq

`func NewUsageReportReq() *UsageReportReq`

NewUsageReportReq instantiates a new UsageReportReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageReportReqWithDefaults

`func NewUsageReportReqWithDefaults() *UsageReportReq`

NewUsageReportReqWithDefaults instantiates a new UsageReportReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UsageReportReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UsageReportReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UsageReportReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UsageReportReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *UsageReportReq) GetCachedInputTokens() int64`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *UsageReportReq) GetCachedInputTokensOk() (*int64, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *UsageReportReq) SetCachedInputTokens(v int64)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *UsageReportReq) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *UsageReportReq) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *UsageReportReq) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *UsageReportReq) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *UsageReportReq) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *UsageReportReq) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *UsageReportReq) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *UsageReportReq) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *UsageReportReq) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *UsageReportReq) GetCostLimitCents() int64`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *UsageReportReq) GetCostLimitCentsOk() (*int64, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *UsageReportReq) SetCostLimitCents(v int64)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *UsageReportReq) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *UsageReportReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UsageReportReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UsageReportReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UsageReportReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *UsageReportReq) GetInputTokens() int64`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *UsageReportReq) GetInputTokensOk() (*int64, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *UsageReportReq) SetInputTokens(v int64)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *UsageReportReq) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetKind

`func (o *UsageReportReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UsageReportReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UsageReportReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UsageReportReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLane

`func (o *UsageReportReq) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *UsageReportReq) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *UsageReportReq) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *UsageReportReq) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *UsageReportReq) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *UsageReportReq) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *UsageReportReq) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *UsageReportReq) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *UsageReportReq) GetOutputTokens() int64`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *UsageReportReq) GetOutputTokensOk() (*int64, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *UsageReportReq) SetOutputTokens(v int64)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *UsageReportReq) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *UsageReportReq) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UsageReportReq) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UsageReportReq) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UsageReportReq) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *UsageReportReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UsageReportReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UsageReportReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UsageReportReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *UsageReportReq) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageReportReq) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageReportReq) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *UsageReportReq) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *UsageReportReq) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *UsageReportReq) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *UsageReportReq) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *UsageReportReq) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSamples

`func (o *UsageReportReq) GetSamples() []UsageSampleReq`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *UsageReportReq) GetSamplesOk() (*[]UsageSampleReq, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *UsageReportReq) SetSamples(v []UsageSampleReq)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *UsageReportReq) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSynthetic

`func (o *UsageReportReq) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *UsageReportReq) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *UsageReportReq) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *UsageReportReq) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *UsageReportReq) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *UsageReportReq) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *UsageReportReq) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *UsageReportReq) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *UsageReportReq) GetUsedPct() float64`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *UsageReportReq) GetUsedPctOk() (*float64, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *UsageReportReq) SetUsedPct(v float64)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *UsageReportReq) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *UsageReportReq) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *UsageReportReq) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *UsageReportReq) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *UsageReportReq) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *UsageReportReq) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *UsageReportReq) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *UsageReportReq) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *UsageReportReq) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *UsageReportReq) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *UsageReportReq) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *UsageReportReq) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *UsageReportReq) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


