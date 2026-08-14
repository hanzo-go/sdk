# IngestReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**CachedInputTokens** | Pointer to **int32** |  | [optional] 
**Confidence** | Pointer to **string** |  | [optional] 
**CostCents** | Pointer to **int32** |  | [optional] 
**CostLimitCents** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**InputTokens** | Pointer to **int32** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Lane** | Pointer to **string** |  | [optional] 
**Machine** | Pointer to **string** |  | [optional] 
**OutputTokens** | Pointer to **int32** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Requests** | Pointer to **int32** |  | [optional] 
**ResetsAt** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**[]ReadingReq**](ReadingReq.md) | Samples is the batch form, up to 256 samples; leave it empty to send one sample inline on the same fields. | [optional] 
**Synthetic** | Pointer to **bool** |  | [optional] 
**TotalTokens** | Pointer to **int32** |  | [optional] 
**UsedPct** | Pointer to **float32** |  | [optional] 
**Window** | Pointer to **string** |  | [optional] 
**WindowMinutes** | Pointer to **int32** |  | [optional] 
**WindowStart** | Pointer to **string** |  | [optional] 

## Methods

### NewIngestReq

`func NewIngestReq() *IngestReq`

NewIngestReq instantiates a new IngestReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestReqWithDefaults

`func NewIngestReqWithDefaults() *IngestReq`

NewIngestReqWithDefaults instantiates a new IngestReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *IngestReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IngestReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IngestReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IngestReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCachedInputTokens

`func (o *IngestReq) GetCachedInputTokens() int32`

GetCachedInputTokens returns the CachedInputTokens field if non-nil, zero value otherwise.

### GetCachedInputTokensOk

`func (o *IngestReq) GetCachedInputTokensOk() (*int32, bool)`

GetCachedInputTokensOk returns a tuple with the CachedInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedInputTokens

`func (o *IngestReq) SetCachedInputTokens(v int32)`

SetCachedInputTokens sets CachedInputTokens field to given value.

### HasCachedInputTokens

`func (o *IngestReq) HasCachedInputTokens() bool`

HasCachedInputTokens returns a boolean if a field has been set.

### GetConfidence

`func (o *IngestReq) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *IngestReq) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *IngestReq) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *IngestReq) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCostCents

`func (o *IngestReq) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *IngestReq) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *IngestReq) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *IngestReq) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostLimitCents

`func (o *IngestReq) GetCostLimitCents() int32`

GetCostLimitCents returns the CostLimitCents field if non-nil, zero value otherwise.

### GetCostLimitCentsOk

`func (o *IngestReq) GetCostLimitCentsOk() (*int32, bool)`

GetCostLimitCentsOk returns a tuple with the CostLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLimitCents

`func (o *IngestReq) SetCostLimitCents(v int32)`

SetCostLimitCents sets CostLimitCents field to given value.

### HasCostLimitCents

`func (o *IngestReq) HasCostLimitCents() bool`

HasCostLimitCents returns a boolean if a field has been set.

### GetCurrency

`func (o *IngestReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IngestReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IngestReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IngestReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInputTokens

`func (o *IngestReq) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *IngestReq) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *IngestReq) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *IngestReq) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetKind

`func (o *IngestReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IngestReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IngestReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IngestReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLane

`func (o *IngestReq) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *IngestReq) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *IngestReq) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *IngestReq) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetMachine

`func (o *IngestReq) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *IngestReq) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *IngestReq) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *IngestReq) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOutputTokens

`func (o *IngestReq) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *IngestReq) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *IngestReq) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *IngestReq) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetPlan

`func (o *IngestReq) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *IngestReq) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *IngestReq) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *IngestReq) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *IngestReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IngestReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IngestReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IngestReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *IngestReq) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *IngestReq) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *IngestReq) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *IngestReq) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetResetsAt

`func (o *IngestReq) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *IngestReq) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *IngestReq) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *IngestReq) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetSamples

`func (o *IngestReq) GetSamples() []ReadingReq`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *IngestReq) GetSamplesOk() (*[]ReadingReq, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *IngestReq) SetSamples(v []ReadingReq)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *IngestReq) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSynthetic

`func (o *IngestReq) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *IngestReq) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *IngestReq) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *IngestReq) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTotalTokens

`func (o *IngestReq) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *IngestReq) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *IngestReq) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *IngestReq) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedPct

`func (o *IngestReq) GetUsedPct() float32`

GetUsedPct returns the UsedPct field if non-nil, zero value otherwise.

### GetUsedPctOk

`func (o *IngestReq) GetUsedPctOk() (*float32, bool)`

GetUsedPctOk returns a tuple with the UsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPct

`func (o *IngestReq) SetUsedPct(v float32)`

SetUsedPct sets UsedPct field to given value.

### HasUsedPct

`func (o *IngestReq) HasUsedPct() bool`

HasUsedPct returns a boolean if a field has been set.

### GetWindow

`func (o *IngestReq) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *IngestReq) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *IngestReq) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *IngestReq) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *IngestReq) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *IngestReq) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *IngestReq) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *IngestReq) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetWindowStart

`func (o *IngestReq) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *IngestReq) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *IngestReq) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *IngestReq) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


