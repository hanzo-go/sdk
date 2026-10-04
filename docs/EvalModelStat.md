# EvalModelStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int64** | tokens it answered with | [optional] 
**CostCents** | Pointer to **int64** | what this model cost, in cents | [optional] 
**CostPct** | Pointer to **float64** | share of total spend, 0..100 | [optional] 
**ErrorRate** | Pointer to **float64** | share of its calls that failed, 0..1 | [optional] 
**Errors** | Pointer to **int64** | calls to it that did not succeed | [optional] 
**Model** | Pointer to **string** | the model this row is about, or \&quot;other\&quot; for the fold | [optional] 
**ModelCount** | Pointer to **int64** | &gt;0 only on the \&quot;other\&quot; fold | [optional] 
**P50Ms** | Pointer to **float64** | median latency, null when no spans carry it | [optional] 
**P95Ms** | Pointer to **float64** | 95th-percentile latency, null when unknown | [optional] 
**P99Ms** | Pointer to **float64** | 99th-percentile latency, null when unknown | [optional] 
**PromptTokens** | Pointer to **int64** | tokens sent to it | [optional] 
**Provider** | Pointer to **string** | who serves it | [optional] 
**Requests** | Pointer to **int64** | calls to this model in the window | [optional] 
**TotalTokens** | Pointer to **int64** | prompt plus completion | [optional] 

## Methods

### NewEvalModelStat

`func NewEvalModelStat() *EvalModelStat`

NewEvalModelStat instantiates a new EvalModelStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalModelStatWithDefaults

`func NewEvalModelStatWithDefaults() *EvalModelStat`

NewEvalModelStatWithDefaults instantiates a new EvalModelStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *EvalModelStat) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *EvalModelStat) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *EvalModelStat) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *EvalModelStat) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *EvalModelStat) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *EvalModelStat) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *EvalModelStat) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *EvalModelStat) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostPct

`func (o *EvalModelStat) GetCostPct() float64`

GetCostPct returns the CostPct field if non-nil, zero value otherwise.

### GetCostPctOk

`func (o *EvalModelStat) GetCostPctOk() (*float64, bool)`

GetCostPctOk returns a tuple with the CostPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPct

`func (o *EvalModelStat) SetCostPct(v float64)`

SetCostPct sets CostPct field to given value.

### HasCostPct

`func (o *EvalModelStat) HasCostPct() bool`

HasCostPct returns a boolean if a field has been set.

### GetErrorRate

`func (o *EvalModelStat) GetErrorRate() float64`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *EvalModelStat) GetErrorRateOk() (*float64, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *EvalModelStat) SetErrorRate(v float64)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *EvalModelStat) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *EvalModelStat) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *EvalModelStat) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *EvalModelStat) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *EvalModelStat) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetModel

`func (o *EvalModelStat) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalModelStat) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalModelStat) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvalModelStat) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelCount

`func (o *EvalModelStat) GetModelCount() int64`

GetModelCount returns the ModelCount field if non-nil, zero value otherwise.

### GetModelCountOk

`func (o *EvalModelStat) GetModelCountOk() (*int64, bool)`

GetModelCountOk returns a tuple with the ModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCount

`func (o *EvalModelStat) SetModelCount(v int64)`

SetModelCount sets ModelCount field to given value.

### HasModelCount

`func (o *EvalModelStat) HasModelCount() bool`

HasModelCount returns a boolean if a field has been set.

### GetP50Ms

`func (o *EvalModelStat) GetP50Ms() float64`

GetP50Ms returns the P50Ms field if non-nil, zero value otherwise.

### GetP50MsOk

`func (o *EvalModelStat) GetP50MsOk() (*float64, bool)`

GetP50MsOk returns a tuple with the P50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50Ms

`func (o *EvalModelStat) SetP50Ms(v float64)`

SetP50Ms sets P50Ms field to given value.

### HasP50Ms

`func (o *EvalModelStat) HasP50Ms() bool`

HasP50Ms returns a boolean if a field has been set.

### GetP95Ms

`func (o *EvalModelStat) GetP95Ms() float64`

GetP95Ms returns the P95Ms field if non-nil, zero value otherwise.

### GetP95MsOk

`func (o *EvalModelStat) GetP95MsOk() (*float64, bool)`

GetP95MsOk returns a tuple with the P95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95Ms

`func (o *EvalModelStat) SetP95Ms(v float64)`

SetP95Ms sets P95Ms field to given value.

### HasP95Ms

`func (o *EvalModelStat) HasP95Ms() bool`

HasP95Ms returns a boolean if a field has been set.

### GetP99Ms

`func (o *EvalModelStat) GetP99Ms() float64`

GetP99Ms returns the P99Ms field if non-nil, zero value otherwise.

### GetP99MsOk

`func (o *EvalModelStat) GetP99MsOk() (*float64, bool)`

GetP99MsOk returns a tuple with the P99Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99Ms

`func (o *EvalModelStat) SetP99Ms(v float64)`

SetP99Ms sets P99Ms field to given value.

### HasP99Ms

`func (o *EvalModelStat) HasP99Ms() bool`

HasP99Ms returns a boolean if a field has been set.

### GetPromptTokens

`func (o *EvalModelStat) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *EvalModelStat) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *EvalModelStat) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *EvalModelStat) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetProvider

`func (o *EvalModelStat) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EvalModelStat) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EvalModelStat) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EvalModelStat) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *EvalModelStat) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *EvalModelStat) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *EvalModelStat) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *EvalModelStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTotalTokens

`func (o *EvalModelStat) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *EvalModelStat) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *EvalModelStat) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *EvalModelStat) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


