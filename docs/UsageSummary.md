# UsageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**Accounts**](Accounts.md) | Accounts is the caller&#39;s own linked provider accounts beside the org&#39;s Hanzo-routed usage, labelled row by row and never summed together. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s exclusive end, RFC3339 UTC. | [optional] 
**Interval** | Pointer to **string** | Interval is the bucket width the spend series is gap-filled at. | [optional] 
**Llm** | Pointer to [**LLM**](LLM.md) | LLM is the org&#39;s Hanzo-routed inference totals from the warehouse. | [optional] 
**Range** | Pointer to **string** | Range is the window label that was served. | [optional] 
**Scope** | Pointer to [**UsageScope**](UsageScope.md) | Scope is the tenant and subject the roll-up was answered for. | [optional] 
**Sources** | Pointer to [**Sources**](Sources.md) | Sources says which upstreams actually answered, so a zero can be read as \&quot;no data yet\&quot; rather than as a measurement. | [optional] 
**Spend** | Pointer to [**Spend**](Spend.md) | Spend is the categorized cost roll-up from the billing ledger. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start, RFC3339 UTC. | [optional] 

## Methods

### NewUsageSummary

`func NewUsageSummary() *UsageSummary`

NewUsageSummary instantiates a new UsageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSummaryWithDefaults

`func NewUsageSummaryWithDefaults() *UsageSummary`

NewUsageSummaryWithDefaults instantiates a new UsageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *UsageSummary) GetAccounts() Accounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UsageSummary) GetAccountsOk() (*Accounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UsageSummary) SetAccounts(v Accounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UsageSummary) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEnd

`func (o *UsageSummary) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *UsageSummary) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *UsageSummary) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *UsageSummary) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *UsageSummary) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UsageSummary) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UsageSummary) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UsageSummary) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLlm

`func (o *UsageSummary) GetLlm() LLM`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *UsageSummary) GetLlmOk() (*LLM, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *UsageSummary) SetLlm(v LLM)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *UsageSummary) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetRange

`func (o *UsageSummary) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *UsageSummary) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *UsageSummary) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *UsageSummary) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *UsageSummary) GetScope() UsageScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UsageSummary) GetScopeOk() (*UsageScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UsageSummary) SetScope(v UsageScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UsageSummary) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSources

`func (o *UsageSummary) GetSources() Sources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *UsageSummary) GetSourcesOk() (*Sources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *UsageSummary) SetSources(v Sources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *UsageSummary) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSpend

`func (o *UsageSummary) GetSpend() Spend`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *UsageSummary) GetSpendOk() (*Spend, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *UsageSummary) SetSpend(v Spend)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *UsageSummary) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStart

`func (o *UsageSummary) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UsageSummary) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UsageSummary) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *UsageSummary) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


