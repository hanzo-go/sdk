# UsageProviderRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int64** | CostCents is what they cost the org, in US cents. | [optional] 
**Provider** | Pointer to **string** | Provider is the upstream the requests were routed to, e.g. anthropic. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many completions the org made against that provider. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is the total tokens those completions consumed, prompt plus completion. | [optional] 

## Methods

### NewUsageProviderRow

`func NewUsageProviderRow() *UsageProviderRow`

NewUsageProviderRow instantiates a new UsageProviderRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageProviderRowWithDefaults

`func NewUsageProviderRowWithDefaults() *UsageProviderRow`

NewUsageProviderRowWithDefaults instantiates a new UsageProviderRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *UsageProviderRow) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *UsageProviderRow) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *UsageProviderRow) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *UsageProviderRow) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetProvider

`func (o *UsageProviderRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UsageProviderRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UsageProviderRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UsageProviderRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *UsageProviderRow) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageProviderRow) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageProviderRow) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *UsageProviderRow) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *UsageProviderRow) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UsageProviderRow) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UsageProviderRow) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *UsageProviderRow) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


