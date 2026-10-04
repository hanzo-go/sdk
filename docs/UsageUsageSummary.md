# UsageUsageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**UsageAccounts**](UsageAccounts.md) | Accounts is the caller&#39;s own linked provider accounts beside the org&#39;s Hanzo-routed usage, labelled row by row and never summed together. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s exclusive end, RFC3339 UTC. | [optional] 
**Interval** | Pointer to **string** | Interval is the bucket width the spend series is gap-filled at. | [optional] 
**Llm** | Pointer to [**UsageLLM**](UsageLLM.md) | LLM is the org&#39;s Hanzo-routed inference totals from the warehouse. | [optional] 
**Range** | Pointer to **string** | Range is the window label that was served. | [optional] 
**Scope** | Pointer to [**UsageUsageScope**](UsageUsageScope.md) | Scope is the tenant and subject the roll-up was answered for. | [optional] 
**Sources** | Pointer to [**UsageSources**](UsageSources.md) | Sources says which upstreams actually answered, so a zero can be read as \&quot;no data yet\&quot; rather than as a measurement. | [optional] 
**Spend** | Pointer to [**UsageSpend**](UsageSpend.md) | Spend is the categorized cost roll-up from the billing ledger. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start, RFC3339 UTC. | [optional] 

## Methods

### NewUsageUsageSummary

`func NewUsageUsageSummary() *UsageUsageSummary`

NewUsageUsageSummary instantiates a new UsageUsageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageUsageSummaryWithDefaults

`func NewUsageUsageSummaryWithDefaults() *UsageUsageSummary`

NewUsageUsageSummaryWithDefaults instantiates a new UsageUsageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *UsageUsageSummary) GetAccounts() UsageAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UsageUsageSummary) GetAccountsOk() (*UsageAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UsageUsageSummary) SetAccounts(v UsageAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UsageUsageSummary) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEnd

`func (o *UsageUsageSummary) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *UsageUsageSummary) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *UsageUsageSummary) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *UsageUsageSummary) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *UsageUsageSummary) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UsageUsageSummary) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UsageUsageSummary) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UsageUsageSummary) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLlm

`func (o *UsageUsageSummary) GetLlm() UsageLLM`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *UsageUsageSummary) GetLlmOk() (*UsageLLM, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *UsageUsageSummary) SetLlm(v UsageLLM)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *UsageUsageSummary) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetRange

`func (o *UsageUsageSummary) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *UsageUsageSummary) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *UsageUsageSummary) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *UsageUsageSummary) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *UsageUsageSummary) GetScope() UsageUsageScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UsageUsageSummary) GetScopeOk() (*UsageUsageScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UsageUsageSummary) SetScope(v UsageUsageScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UsageUsageSummary) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSources

`func (o *UsageUsageSummary) GetSources() UsageSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *UsageUsageSummary) GetSourcesOk() (*UsageSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *UsageUsageSummary) SetSources(v UsageSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *UsageUsageSummary) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSpend

`func (o *UsageUsageSummary) GetSpend() UsageSpend`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *UsageUsageSummary) GetSpendOk() (*UsageSpend, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *UsageUsageSummary) SetSpend(v UsageSpend)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *UsageUsageSummary) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStart

`func (o *UsageUsageSummary) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UsageUsageSummary) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UsageUsageSummary) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *UsageUsageSummary) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


