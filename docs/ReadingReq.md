# ReadingReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider-side account the sample belongs to. | [optional] 
**CachedInputTokens** | Pointer to **int32** | CachedInputTokens is the window&#39;s cached-prompt-token count. | [optional] 
**Confidence** | Pointer to **string** | Confidence says how real the counters are, as the meter graded itself. | [optional] 
**CostCents** | Pointer to **int32** | CostCents is the window&#39;s spend in cents, as the provider&#39;s meter states it. | [optional] 
**CostLimitCents** | Pointer to **int32** | CostLimitCents is the window&#39;s spend cap in cents, when the meter knows one. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO currency the cost fields are stated in. | [optional] 
**InputTokens** | Pointer to **int32** | InputTokens is the window&#39;s prompt-token count. | [optional] 
**Kind** | Pointer to **string** | Kind is subscription or apikey; anything else is refused. | [optional] 
**Lane** | Pointer to **string** | Lane names the meter&#39;s own lane label for this measurement. | [optional] 
**Machine** | Pointer to **string** | Machine is the machine the collector observed the account on. Required. | [optional] 
**OutputTokens** | Pointer to **int32** | OutputTokens is the window&#39;s completion-token count. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan label the account is on. | [optional] 
**Provider** | Pointer to **string** | Provider is the AI provider whose meter reported this sample. Required. | [optional] 
**Requests** | Pointer to **int32** | Requests is the window&#39;s request count. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when the window resets, RFC 3339, bounded. | [optional] 
**Synthetic** | Pointer to **bool** | Synthetic marks a sample the collector derived rather than observed. | [optional] 
**TotalTokens** | Pointer to **int32** | TotalTokens is the window&#39;s total token count. | [optional] 
**UsedPct** | Pointer to **float32** | UsedPct is how much of the window&#39;s allowance is consumed, clamped 0..100. | [optional] 
**Window** | Pointer to **string** | Window is the window class, one of 6h, day, week, month; anything else is refused rather than silently reclassified. | [optional] 
**WindowMinutes** | Pointer to **int32** | WindowMinutes is the window&#39;s length as the meter reported it. | [optional] 
**WindowStart** | Pointer to **string** | WindowStart is when the measured window opened, RFC 3339, bounded to a sane interval around now. | [optional] 

## Methods

### NewReadingReq

`func NewReadingReq() *ReadingReq`

NewReadingReq instantiates a new ReadingReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadingReqWithDefaults

`func NewReadingReqWithDefaults() *ReadingReq`

NewReadingReqWithDefaults instantiates a new ReadingReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ReadingReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ReadingReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ReadingReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ReadingReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *ReadingReq) GetCachedInputTokens() int32`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *ReadingReq) GetCachedInputTokensOk() (*int32, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *ReadingReq) SetCachedInputTokens(v int32)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *ReadingReq) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *ReadingReq) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ReadingReq) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ReadingReq) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ReadingReq) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *ReadingReq) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ReadingReq) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ReadingReq) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ReadingReq) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *ReadingReq) GetCostLimitCents() int32`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *ReadingReq) GetCostLimitCentsOk() (*int32, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *ReadingReq) SetCostLimitCents(v int32)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *ReadingReq) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *ReadingReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReadingReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReadingReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ReadingReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *ReadingReq) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *ReadingReq) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *ReadingReq) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *ReadingReq) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetKind

`func (o *ReadingReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReadingReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReadingReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ReadingReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLane

`func (o *ReadingReq) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *ReadingReq) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *ReadingReq) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *ReadingReq) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *ReadingReq) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *ReadingReq) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *ReadingReq) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *ReadingReq) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *ReadingReq) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *ReadingReq) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *ReadingReq) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *ReadingReq) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *ReadingReq) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ReadingReq) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ReadingReq) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ReadingReq) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *ReadingReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ReadingReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ReadingReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ReadingReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *ReadingReq) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ReadingReq) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ReadingReq) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ReadingReq) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *ReadingReq) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *ReadingReq) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *ReadingReq) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *ReadingReq) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSynthetic

`func (o *ReadingReq) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *ReadingReq) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *ReadingReq) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *ReadingReq) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *ReadingReq) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *ReadingReq) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *ReadingReq) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *ReadingReq) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *ReadingReq) GetUsedPct() float32`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *ReadingReq) GetUsedPctOk() (*float32, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *ReadingReq) SetUsedPct(v float32)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *ReadingReq) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *ReadingReq) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ReadingReq) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ReadingReq) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *ReadingReq) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *ReadingReq) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *ReadingReq) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *ReadingReq) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *ReadingReq) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *ReadingReq) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *ReadingReq) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *ReadingReq) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *ReadingReq) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


