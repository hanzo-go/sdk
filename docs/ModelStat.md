# ModelStat

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

### NewModelStat

`func NewModelStat() *ModelStat`

NewModelStat instantiates a new ModelStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelStatWithDefaults

`func NewModelStatWithDefaults() *ModelStat`

NewModelStatWithDefaults instantiates a new ModelStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *ModelStat) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *ModelStat) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *ModelStat) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *ModelStat) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *ModelStat) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ModelStat) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ModelStat) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ModelStat) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostPct

`func (o *ModelStat) GetCostPct() float64`

GetCostPct returns the CostPct field if non-nil, zero value otherwise.

### GetCostPctOk

`func (o *ModelStat) GetCostPctOk() (*float64, bool)`

GetCostPctOk returns a tuple with the CostPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPct

`func (o *ModelStat) SetCostPct(v float64)`

SetCostPct sets CostPct field to given value.

### HasCostPct

`func (o *ModelStat) HasCostPct() bool`

HasCostPct returns a boolean if a field has been set.

### GetErrorRate

`func (o *ModelStat) GetErrorRate() float64`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *ModelStat) GetErrorRateOk() (*float64, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *ModelStat) SetErrorRate(v float64)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *ModelStat) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *ModelStat) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ModelStat) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ModelStat) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ModelStat) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetModel

`func (o *ModelStat) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelStat) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelStat) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModelStat) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelCount

`func (o *ModelStat) GetModelCount() int64`

GetModelCount returns the ModelCount field if non-nil, zero value otherwise.

### GetModelCountOk

`func (o *ModelStat) GetModelCountOk() (*int64, bool)`

GetModelCountOk returns a tuple with the ModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCount

`func (o *ModelStat) SetModelCount(v int64)`

SetModelCount sets ModelCount field to given value.

### HasModelCount

`func (o *ModelStat) HasModelCount() bool`

HasModelCount returns a boolean if a field has been set.

### GetP50Ms

`func (o *ModelStat) GetP50Ms() float64`

GetP50Ms returns the P50Ms field if non-nil, zero value otherwise.

### GetP50MsOk

`func (o *ModelStat) GetP50MsOk() (*float64, bool)`

GetP50MsOk returns a tuple with the P50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50Ms

`func (o *ModelStat) SetP50Ms(v float64)`

SetP50Ms sets P50Ms field to given value.

### HasP50Ms

`func (o *ModelStat) HasP50Ms() bool`

HasP50Ms returns a boolean if a field has been set.

### GetP95Ms

`func (o *ModelStat) GetP95Ms() float64`

GetP95Ms returns the P95Ms field if non-nil, zero value otherwise.

### GetP95MsOk

`func (o *ModelStat) GetP95MsOk() (*float64, bool)`

GetP95MsOk returns a tuple with the P95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95Ms

`func (o *ModelStat) SetP95Ms(v float64)`

SetP95Ms sets P95Ms field to given value.

### HasP95Ms

`func (o *ModelStat) HasP95Ms() bool`

HasP95Ms returns a boolean if a field has been set.

### GetP99Ms

`func (o *ModelStat) GetP99Ms() float64`

GetP99Ms returns the P99Ms field if non-nil, zero value otherwise.

### GetP99MsOk

`func (o *ModelStat) GetP99MsOk() (*float64, bool)`

GetP99MsOk returns a tuple with the P99Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99Ms

`func (o *ModelStat) SetP99Ms(v float64)`

SetP99Ms sets P99Ms field to given value.

### HasP99Ms

`func (o *ModelStat) HasP99Ms() bool`

HasP99Ms returns a boolean if a field has been set.

### GetPromptTokens

`func (o *ModelStat) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *ModelStat) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *ModelStat) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *ModelStat) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetProvider

`func (o *ModelStat) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelStat) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelStat) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelStat) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *ModelStat) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ModelStat) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ModelStat) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ModelStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTotalTokens

`func (o *ModelStat) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *ModelStat) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *ModelStat) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *ModelStat) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


