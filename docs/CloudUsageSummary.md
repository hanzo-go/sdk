# CloudUsageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**CloudAccounts**](CloudAccounts.md) | Accounts is the caller&#39;s own linked provider accounts beside the org&#39;s Hanzo-routed usage, labelled row by row and never summed together. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s exclusive end, RFC3339 UTC. | [optional] 
**Interval** | Pointer to **string** | Interval is the bucket width the spend series is gap-filled at. | [optional] 
**Llm** | Pointer to [**CloudLLM**](CloudLLM.md) | LLM is the org&#39;s Hanzo-routed inference totals from the warehouse. | [optional] 
**Range** | Pointer to **string** | Range is the window label that was served. | [optional] 
**Scope** | Pointer to [**CloudUsageScope**](CloudUsageScope.md) | Scope is the tenant and subject the roll-up was answered for. | [optional] 
**Sources** | Pointer to [**CloudSources**](CloudSources.md) | Sources says which upstreams actually answered, so a zero can be read as \&quot;no data yet\&quot; rather than as a measurement. | [optional] 
**Spend** | Pointer to [**CloudSpend**](CloudSpend.md) | Spend is the categorized cost roll-up from the billing ledger. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start, RFC3339 UTC. | [optional] 

## Methods

### NewCloudUsageSummary

`func NewCloudUsageSummary() *CloudUsageSummary`

NewCloudUsageSummary instantiates a new CloudUsageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUsageSummaryWithDefaults

`func NewCloudUsageSummaryWithDefaults() *CloudUsageSummary`

NewCloudUsageSummaryWithDefaults instantiates a new CloudUsageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *CloudUsageSummary) GetAccounts() CloudAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *CloudUsageSummary) GetAccountsOk() (*CloudAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *CloudUsageSummary) SetAccounts(v CloudAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *CloudUsageSummary) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEnd

`func (o *CloudUsageSummary) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudUsageSummary) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudUsageSummary) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudUsageSummary) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *CloudUsageSummary) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudUsageSummary) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudUsageSummary) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudUsageSummary) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLlm

`func (o *CloudUsageSummary) GetLlm() CloudLLM`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *CloudUsageSummary) GetLlmOk() (*CloudLLM, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *CloudUsageSummary) SetLlm(v CloudLLM)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *CloudUsageSummary) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetRange

`func (o *CloudUsageSummary) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudUsageSummary) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudUsageSummary) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudUsageSummary) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *CloudUsageSummary) GetScope() CloudUsageScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudUsageSummary) GetScopeOk() (*CloudUsageScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudUsageSummary) SetScope(v CloudUsageScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudUsageSummary) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSources

`func (o *CloudUsageSummary) GetSources() CloudSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudUsageSummary) GetSourcesOk() (*CloudSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudUsageSummary) SetSources(v CloudSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudUsageSummary) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSpend

`func (o *CloudUsageSummary) GetSpend() CloudSpend`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *CloudUsageSummary) GetSpendOk() (*CloudSpend, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *CloudUsageSummary) SetSpend(v CloudSpend)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *CloudUsageSummary) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStart

`func (o *CloudUsageSummary) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudUsageSummary) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudUsageSummary) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudUsageSummary) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


