# LinkIngestReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**CachedInputTokens** | Pointer to **int64** |  | [optional] 
**Confidence** | Pointer to **string** |  | [optional] 
**CostCents** | Pointer to **int64** |  | [optional] 
**CostLimitCents** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**InputTokens** | Pointer to **int64** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Lane** | Pointer to **string** |  | [optional] 
**Machine** | Pointer to **string** |  | [optional] 
**OutputTokens** | Pointer to **int64** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Requests** | Pointer to **int64** |  | [optional] 
**ResetsAt** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**[]LinkReadingReq**](LinkReadingReq.md) | Samples is the batch form, up to 256 samples; leave it empty to send one sample inline on the same fields. | [optional] 
**Synthetic** | Pointer to **bool** |  | [optional] 
**TotalTokens** | Pointer to **int64** |  | [optional] 
**UsedPct** | Pointer to **float64** |  | [optional] 
**Window** | Pointer to **string** |  | [optional] 
**WindowMinutes** | Pointer to **int32** |  | [optional] 
**WindowStart** | Pointer to **string** |  | [optional] 

## Methods

### NewLinkIngestReq

`func NewLinkIngestReq() *LinkIngestReq`

NewLinkIngestReq instantiates a new LinkIngestReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkIngestReqWithDefaults

`func NewLinkIngestReqWithDefaults() *LinkIngestReq`

NewLinkIngestReqWithDefaults instantiates a new LinkIngestReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LinkIngestReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LinkIngestReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LinkIngestReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LinkIngestReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *LinkIngestReq) GetCachedInputTokens() int64`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *LinkIngestReq) GetCachedInputTokensOk() (*int64, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *LinkIngestReq) SetCachedInputTokens(v int64)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *LinkIngestReq) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *LinkIngestReq) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *LinkIngestReq) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *LinkIngestReq) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *LinkIngestReq) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *LinkIngestReq) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LinkIngestReq) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LinkIngestReq) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LinkIngestReq) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *LinkIngestReq) GetCostLimitCents() int64`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *LinkIngestReq) GetCostLimitCentsOk() (*int64, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *LinkIngestReq) SetCostLimitCents(v int64)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *LinkIngestReq) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *LinkIngestReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LinkIngestReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LinkIngestReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LinkIngestReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *LinkIngestReq) GetInputTokens() int64`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *LinkIngestReq) GetInputTokensOk() (*int64, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *LinkIngestReq) SetInputTokens(v int64)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *LinkIngestReq) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetKind

`func (o *LinkIngestReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LinkIngestReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LinkIngestReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LinkIngestReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLane

`func (o *LinkIngestReq) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *LinkIngestReq) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *LinkIngestReq) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *LinkIngestReq) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *LinkIngestReq) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *LinkIngestReq) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *LinkIngestReq) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *LinkIngestReq) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *LinkIngestReq) GetOutputTokens() int64`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *LinkIngestReq) GetOutputTokensOk() (*int64, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *LinkIngestReq) SetOutputTokens(v int64)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *LinkIngestReq) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *LinkIngestReq) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LinkIngestReq) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LinkIngestReq) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LinkIngestReq) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *LinkIngestReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LinkIngestReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LinkIngestReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LinkIngestReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *LinkIngestReq) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LinkIngestReq) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LinkIngestReq) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LinkIngestReq) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *LinkIngestReq) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *LinkIngestReq) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *LinkIngestReq) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *LinkIngestReq) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSamples

`func (o *LinkIngestReq) GetSamples() []LinkReadingReq`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *LinkIngestReq) GetSamplesOk() (*[]LinkReadingReq, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *LinkIngestReq) SetSamples(v []LinkReadingReq)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *LinkIngestReq) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSynthetic

`func (o *LinkIngestReq) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *LinkIngestReq) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *LinkIngestReq) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *LinkIngestReq) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *LinkIngestReq) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *LinkIngestReq) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *LinkIngestReq) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *LinkIngestReq) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *LinkIngestReq) GetUsedPct() float64`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *LinkIngestReq) GetUsedPctOk() (*float64, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *LinkIngestReq) SetUsedPct(v float64)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *LinkIngestReq) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *LinkIngestReq) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *LinkIngestReq) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *LinkIngestReq) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *LinkIngestReq) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *LinkIngestReq) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *LinkIngestReq) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *LinkIngestReq) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *LinkIngestReq) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *LinkIngestReq) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *LinkIngestReq) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *LinkIngestReq) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *LinkIngestReq) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


