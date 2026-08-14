# ModelStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int32** | tokens it answered with | [optional] 
**CostCents** | Pointer to **int32** | what this model cost, in cents | [optional] 
**CostPct** | Pointer to **float32** | share of total spend, 0..100 | [optional] 
**ErrorRate** | Pointer to **float32** | share of its calls that failed, 0..1 | [optional] 
**Errors** | Pointer to **int32** | calls to it that did not succeed | [optional] 
**Model** | Pointer to **string** | the model this row is about, or \&quot;other\&quot; for the fold | [optional] 
**ModelCount** | Pointer to **int32** | &gt;0 only on the \&quot;other\&quot; fold | [optional] 
**P50Ms** | Pointer to **float32** | median latency, null when no spans carry it | [optional] 
**P95Ms** | Pointer to **float32** | 95th-percentile latency, null when unknown | [optional] 
**P99Ms** | Pointer to **float32** | 99th-percentile latency, null when unknown | [optional] 
**PromptTokens** | Pointer to **int32** | tokens sent to it | [optional] 
**Provider** | Pointer to **string** | who serves it | [optional] 
**Requests** | Pointer to **int32** | calls to this model in the window | [optional] 
**TotalTokens** | Pointer to **int32** | prompt plus completion | [optional] 

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

`func (o *ModelStat) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *ModelStat) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *ModelStat) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *ModelStat) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *ModelStat) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ModelStat) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ModelStat) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ModelStat) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCostPct

`func (o *ModelStat) GetCostPct() float32`

GetCostPct returns the CostPct field if non-nil, zero value otherwise.

### GetCostPctOk

`func (o *ModelStat) GetCostPctOk() (*float32, bool)`

GetCostPctOk returns a tuple with the CostPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPct

`func (o *ModelStat) SetCostPct(v float32)`

SetCostPct sets CostPct field to given value.

### HasCostPct

`func (o *ModelStat) HasCostPct() bool`

HasCostPct returns a boolean if a field has been set.

### GetErrorRate

`func (o *ModelStat) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *ModelStat) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *ModelStat) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *ModelStat) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *ModelStat) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ModelStat) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ModelStat) SetErrors(v int32)`

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

`func (o *ModelStat) GetModelCount() int32`

GetModelCount returns the ModelCount field if non-nil, zero value otherwise.

### GetModelCountOk

`func (o *ModelStat) GetModelCountOk() (*int32, bool)`

GetModelCountOk returns a tuple with the ModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCount

`func (o *ModelStat) SetModelCount(v int32)`

SetModelCount sets ModelCount field to given value.

### HasModelCount

`func (o *ModelStat) HasModelCount() bool`

HasModelCount returns a boolean if a field has been set.

### GetP50Ms

`func (o *ModelStat) GetP50Ms() float32`

GetP50Ms returns the P50Ms field if non-nil, zero value otherwise.

### GetP50MsOk

`func (o *ModelStat) GetP50MsOk() (*float32, bool)`

GetP50MsOk returns a tuple with the P50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50Ms

`func (o *ModelStat) SetP50Ms(v float32)`

SetP50Ms sets P50Ms field to given value.

### HasP50Ms

`func (o *ModelStat) HasP50Ms() bool`

HasP50Ms returns a boolean if a field has been set.

### GetP95Ms

`func (o *ModelStat) GetP95Ms() float32`

GetP95Ms returns the P95Ms field if non-nil, zero value otherwise.

### GetP95MsOk

`func (o *ModelStat) GetP95MsOk() (*float32, bool)`

GetP95MsOk returns a tuple with the P95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95Ms

`func (o *ModelStat) SetP95Ms(v float32)`

SetP95Ms sets P95Ms field to given value.

### HasP95Ms

`func (o *ModelStat) HasP95Ms() bool`

HasP95Ms returns a boolean if a field has been set.

### GetP99Ms

`func (o *ModelStat) GetP99Ms() float32`

GetP99Ms returns the P99Ms field if non-nil, zero value otherwise.

### GetP99MsOk

`func (o *ModelStat) GetP99MsOk() (*float32, bool)`

GetP99MsOk returns a tuple with the P99Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99Ms

`func (o *ModelStat) SetP99Ms(v float32)`

SetP99Ms sets P99Ms field to given value.

### HasP99Ms

`func (o *ModelStat) HasP99Ms() bool`

HasP99Ms returns a boolean if a field has been set.

### GetPromptTokens

`func (o *ModelStat) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *ModelStat) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *ModelStat) SetPromptTokens(v int32)`

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

`func (o *ModelStat) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ModelStat) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ModelStat) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ModelStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTotalTokens

`func (o *ModelStat) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *ModelStat) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *ModelStat) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *ModelStat) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


