# CloudSampleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the linked account the window was metered from. | [optional] 
**CachedInputTokens** | Pointer to **int32** | CachedInputTokens is the prompt tokens the provider served from cache. | [optional] 
**Confidence** | Pointer to **string** | Confidence says how much the counters below mean. | [optional] 
**CostCents** | Pointer to **int32** | CostCents is what the window cost on the PROVIDER&#39;s own plan, in US cents. | [optional] 
**CostLimitCents** | Pointer to **int32** | CostLimitCents is the plan&#39;s spend ceiling for the window, in US cents. | [optional] 
**Currency** | Pointer to **string** | Currency is the provider&#39;s currency when it is not US cents. | [optional] 
**InputTokens** | Pointer to **int32** | InputTokens is prompt tokens consumed in the window. | [optional] 
**Kind** | Pointer to **string** | Kind is subscription or apikey. Empty is accepted; anything else is refused. | [optional] 
**Lane** | Pointer to **string** | Lane is the meter lane within the account. | [optional] 
**Machine** | Pointer to **string** | Machine is the host whose meter read the window. Required. | [optional] 
**OutputTokens** | Pointer to **int32** | OutputTokens is completion tokens produced in the window. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan the account is on, e.g. a Claude Max plan. | [optional] 
**Provider** | Pointer to **string** | Provider is the upstream the account belongs to, e.g. anthropic. Required. | [optional] 
**Requests** | Pointer to **int32** | Requests is how many requests the window covers. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when the measured window rolls over, RFC3339. Empty is allowed; anything else that is not RFC3339 is refused. | [optional] 
**Synthetic** | Pointer to **bool** | Synthetic marks a window the meter inferred rather than read. | [optional] 
**TotalTokens** | Pointer to **int32** | TotalTokens is the window&#39;s total tokens. | [optional] 
**UsedPct** | Pointer to **float32** | UsedPct is how much of the window&#39;s allowance is consumed, 0–100. | [optional] 
**Window** | Pointer to **string** | Window is the window class: 6h, day, week or month. Required, and a class this surface does not know is refused rather than rewritten. | [optional] 
**WindowMinutes** | Pointer to **int32** | WindowMinutes is the window&#39;s real length in minutes, as the meter reports it. | [optional] 
**WindowStart** | Pointer to **string** | WindowStart is when the measured window opened, RFC3339. Empty is allowed; anything else that is not RFC3339 is refused. | [optional] 

## Methods

### NewCloudSampleReq

`func NewCloudSampleReq() *CloudSampleReq`

NewCloudSampleReq instantiates a new CloudSampleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSampleReqWithDefaults

`func NewCloudSampleReqWithDefaults() *CloudSampleReq`

NewCloudSampleReqWithDefaults instantiates a new CloudSampleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudSampleReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudSampleReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudSampleReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudSampleReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *CloudSampleReq) GetCachedInputTokens() int32`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *CloudSampleReq) GetCachedInputTokensOk() (*int32, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *CloudSampleReq) SetCachedInputTokens(v int32)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *CloudSampleReq) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *CloudSampleReq) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CloudSampleReq) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CloudSampleReq) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CloudSampleReq) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *CloudSampleReq) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *CloudSampleReq) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *CloudSampleReq) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *CloudSampleReq) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *CloudSampleReq) GetCostLimitCents() int32`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *CloudSampleReq) GetCostLimitCentsOk() (*int32, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *CloudSampleReq) SetCostLimitCents(v int32)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *CloudSampleReq) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudSampleReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudSampleReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudSampleReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudSampleReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *CloudSampleReq) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *CloudSampleReq) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *CloudSampleReq) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *CloudSampleReq) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetKind

`func (o *CloudSampleReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudSampleReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudSampleReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudSampleReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLane

`func (o *CloudSampleReq) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *CloudSampleReq) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *CloudSampleReq) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *CloudSampleReq) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *CloudSampleReq) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *CloudSampleReq) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *CloudSampleReq) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *CloudSampleReq) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *CloudSampleReq) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *CloudSampleReq) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *CloudSampleReq) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *CloudSampleReq) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *CloudSampleReq) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CloudSampleReq) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CloudSampleReq) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CloudSampleReq) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *CloudSampleReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudSampleReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudSampleReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudSampleReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *CloudSampleReq) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudSampleReq) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudSampleReq) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudSampleReq) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *CloudSampleReq) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *CloudSampleReq) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *CloudSampleReq) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *CloudSampleReq) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSynthetic

`func (o *CloudSampleReq) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *CloudSampleReq) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *CloudSampleReq) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *CloudSampleReq) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *CloudSampleReq) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *CloudSampleReq) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *CloudSampleReq) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *CloudSampleReq) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *CloudSampleReq) GetUsedPct() float32`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *CloudSampleReq) GetUsedPctOk() (*float32, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *CloudSampleReq) SetUsedPct(v float32)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *CloudSampleReq) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *CloudSampleReq) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *CloudSampleReq) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *CloudSampleReq) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *CloudSampleReq) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *CloudSampleReq) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *CloudSampleReq) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *CloudSampleReq) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *CloudSampleReq) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *CloudSampleReq) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *CloudSampleReq) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *CloudSampleReq) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *CloudSampleReq) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


