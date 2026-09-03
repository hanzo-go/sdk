# ActivityTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDays** | Pointer to **int64** | ActiveDays counts the days with any usage at all — the streak/consistency number. Compare it against len(days) for the share of days the subject showed up. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is the window&#39;s spend in whole US cents, the sum of Days[].CostCents. | [optional] 
**MaxRequests** | Pointer to **int64** | MaxRequests is the same ceiling for a request-based heatmap — the busiest single day&#39;s request count, 0 for an idle window. | [optional] 
**MaxTokens** | Pointer to **int64** | MaxTokens is the busiest single day&#39;s token count: the ceiling to normalize a token heatmap against, so the darkest cell is that day. 0 for an idle window, which a client must not divide by. | [optional] 
**Requests** | Pointer to **int64** | Requests is the sum of Days[].Requests over the whole window. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is the sum of Days[].Tokens over the whole window. | [optional] 

## Methods

### NewActivityTotals

`func NewActivityTotals() *ActivityTotals`

NewActivityTotals instantiates a new ActivityTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTotalsWithDefaults

`func NewActivityTotalsWithDefaults() *ActivityTotals`

NewActivityTotalsWithDefaults instantiates a new ActivityTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDays

`func (o *ActivityTotals) GetActiveDays() int64`

GetActiveDays returns the ActiveDays field if non-nil, zero value otherwise.

### GetActiveDaysOk

`func (o *ActivityTotals) GetActiveDaysOk() (*int64, bool)`

GetActiveDaysOk returns a tuple with the ActiveDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDays

`func (o *ActivityTotals) SetActiveDays(v int64)`

SetActiveDays sets ActiveDays field to given value.

### HasActiveDays

`func (o *ActivityTotals) HasActiveDays() bool`

HasActiveDays returns a boolean if a field has been set.

### GetCostCents

`func (o *ActivityTotals) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ActivityTotals) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ActivityTotals) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ActivityTotals) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetMaxRequests

`func (o *ActivityTotals) GetMaxRequests() int64`

GetMaxRequests returns the MaxRequests field if non-nil, zero value otherwise.

### GetMaxRequestsOk

`func (o *ActivityTotals) GetMaxRequestsOk() (*int64, bool)`

GetMaxRequestsOk returns a tuple with the MaxRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequests

`func (o *ActivityTotals) SetMaxRequests(v int64)`

SetMaxRequests sets MaxRequests field to given value.

### HasMaxRequests

`func (o *ActivityTotals) HasMaxRequests() bool`

HasMaxRequests returns a boolean if a field has been set.

### GetMaxTokens

`func (o *ActivityTotals) GetMaxTokens() int64`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ActivityTotals) GetMaxTokensOk() (*int64, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ActivityTotals) SetMaxTokens(v int64)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ActivityTotals) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetRequests

`func (o *ActivityTotals) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ActivityTotals) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ActivityTotals) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ActivityTotals) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *ActivityTotals) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ActivityTotals) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ActivityTotals) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ActivityTotals) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


