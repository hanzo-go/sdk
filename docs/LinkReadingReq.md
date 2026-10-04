# LinkReadingReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider-side account the sample belongs to. | [optional] 
**CachedInputTokens** | Pointer to **int64** | CachedInputTokens is the window&#39;s cached-prompt-token count. | [optional] 
**Confidence** | Pointer to **string** | Confidence says how real the counters are, as the meter graded itself. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is the window&#39;s spend in cents, as the provider&#39;s meter states it. | [optional] 
**CostLimitCents** | Pointer to **int64** | CostLimitCents is the window&#39;s spend cap in cents, when the meter knows one. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO currency the cost fields are stated in. | [optional] 
**InputTokens** | Pointer to **int64** | InputTokens is the window&#39;s prompt-token count. | [optional] 
**Kind** | Pointer to **string** | Kind is subscription or apikey; anything else is refused. | [optional] 
**Lane** | Pointer to **string** | Lane names the meter&#39;s own lane label for this measurement. | [optional] 
**Machine** | Pointer to **string** | Machine is the machine the collector observed the account on. Required. | [optional] 
**OutputTokens** | Pointer to **int64** | OutputTokens is the window&#39;s completion-token count. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan label the account is on. | [optional] 
**Provider** | Pointer to **string** | Provider is the AI provider whose meter reported this sample. Required. | [optional] 
**Requests** | Pointer to **int64** | Requests is the window&#39;s request count. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when the window resets, RFC 3339, bounded. | [optional] 
**Synthetic** | Pointer to **bool** | Synthetic marks a sample the collector derived rather than observed. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is the window&#39;s total token count. | [optional] 
**UsedPct** | Pointer to **float64** | UsedPct is how much of the window&#39;s allowance is consumed, clamped 0..100. | [optional] 
**Window** | Pointer to **string** | Window is the window class, one of 6h, day, week, month; anything else is refused rather than silently reclassified. | [optional] 
**WindowMinutes** | Pointer to **int32** | WindowMinutes is the window&#39;s length as the meter reported it. | [optional] 
**WindowStart** | Pointer to **string** | WindowStart is when the measured window opened, RFC 3339, bounded to a sane interval around now. | [optional] 

## Methods

### NewLinkReadingReq

`func NewLinkReadingReq() *LinkReadingReq`

NewLinkReadingReq instantiates a new LinkReadingReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkReadingReqWithDefaults

`func NewLinkReadingReqWithDefaults() *LinkReadingReq`

NewLinkReadingReqWithDefaults instantiates a new LinkReadingReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LinkReadingReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LinkReadingReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LinkReadingReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LinkReadingReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *LinkReadingReq) GetCachedInputTokens() int64`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *LinkReadingReq) GetCachedInputTokensOk() (*int64, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *LinkReadingReq) SetCachedInputTokens(v int64)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *LinkReadingReq) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *LinkReadingReq) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *LinkReadingReq) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *LinkReadingReq) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *LinkReadingReq) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *LinkReadingReq) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LinkReadingReq) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LinkReadingReq) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LinkReadingReq) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *LinkReadingReq) GetCostLimitCents() int64`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *LinkReadingReq) GetCostLimitCentsOk() (*int64, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *LinkReadingReq) SetCostLimitCents(v int64)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *LinkReadingReq) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *LinkReadingReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LinkReadingReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LinkReadingReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LinkReadingReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *LinkReadingReq) GetInputTokens() int64`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *LinkReadingReq) GetInputTokensOk() (*int64, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *LinkReadingReq) SetInputTokens(v int64)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *LinkReadingReq) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetKind

`func (o *LinkReadingReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LinkReadingReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LinkReadingReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LinkReadingReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLane

`func (o *LinkReadingReq) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *LinkReadingReq) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *LinkReadingReq) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *LinkReadingReq) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *LinkReadingReq) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *LinkReadingReq) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *LinkReadingReq) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *LinkReadingReq) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *LinkReadingReq) GetOutputTokens() int64`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *LinkReadingReq) GetOutputTokensOk() (*int64, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *LinkReadingReq) SetOutputTokens(v int64)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *LinkReadingReq) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *LinkReadingReq) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LinkReadingReq) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LinkReadingReq) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LinkReadingReq) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *LinkReadingReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LinkReadingReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LinkReadingReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LinkReadingReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *LinkReadingReq) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LinkReadingReq) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LinkReadingReq) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LinkReadingReq) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *LinkReadingReq) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *LinkReadingReq) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *LinkReadingReq) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *LinkReadingReq) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSynthetic

`func (o *LinkReadingReq) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *LinkReadingReq) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *LinkReadingReq) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *LinkReadingReq) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *LinkReadingReq) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *LinkReadingReq) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *LinkReadingReq) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *LinkReadingReq) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *LinkReadingReq) GetUsedPct() float64`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *LinkReadingReq) GetUsedPctOk() (*float64, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *LinkReadingReq) SetUsedPct(v float64)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *LinkReadingReq) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *LinkReadingReq) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *LinkReadingReq) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *LinkReadingReq) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *LinkReadingReq) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *LinkReadingReq) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *LinkReadingReq) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *LinkReadingReq) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *LinkReadingReq) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *LinkReadingReq) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *LinkReadingReq) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *LinkReadingReq) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *LinkReadingReq) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


