# CloudProviderRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int32** | CostCents is what they cost the org, in US cents. | [optional] 
**Provider** | Pointer to **string** | Provider is the upstream the requests were routed to, e.g. anthropic. | [optional] 
**Requests** | Pointer to **int32** | Requests is how many completions the org made against that provider. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is the total tokens those completions consumed, prompt plus completion. | [optional] 

## Methods

### NewCloudProviderRow

`func NewCloudProviderRow() *CloudProviderRow`

NewCloudProviderRow instantiates a new CloudProviderRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderRowWithDefaults

`func NewCloudProviderRowWithDefaults() *CloudProviderRow`

NewCloudProviderRowWithDefaults instantiates a new CloudProviderRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *CloudProviderRow) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *CloudProviderRow) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *CloudProviderRow) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *CloudProviderRow) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetProvider

`func (o *CloudProviderRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudProviderRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudProviderRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudProviderRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *CloudProviderRow) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudProviderRow) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudProviderRow) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudProviderRow) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *CloudProviderRow) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *CloudProviderRow) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *CloudProviderRow) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *CloudProviderRow) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


