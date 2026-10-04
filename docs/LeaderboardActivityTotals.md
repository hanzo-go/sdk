# LeaderboardActivityTotals

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

### NewLeaderboardActivityTotals

`func NewLeaderboardActivityTotals() *LeaderboardActivityTotals`

NewLeaderboardActivityTotals instantiates a new LeaderboardActivityTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardActivityTotalsWithDefaults

`func NewLeaderboardActivityTotalsWithDefaults() *LeaderboardActivityTotals`

NewLeaderboardActivityTotalsWithDefaults instantiates a new LeaderboardActivityTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDays

`func (o *LeaderboardActivityTotals) GetActiveDays() int64`

GetActiveDays returns the ActiveDays field if non-nil, zero value otherwise.

### GetActiveDaysOk

`func (o *LeaderboardActivityTotals) GetActiveDaysOk() (*int64, bool)`

GetActiveDaysOk returns a tuple with the ActiveDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDays

`func (o *LeaderboardActivityTotals) SetActiveDays(v int64)`

SetActiveDays sets ActiveDays field to given value.

### HasActiveDays

`func (o *LeaderboardActivityTotals) HasActiveDays() bool`

HasActiveDays returns a boolean if a field has been set.

### GetCostCents

`func (o *LeaderboardActivityTotals) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LeaderboardActivityTotals) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LeaderboardActivityTotals) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LeaderboardActivityTotals) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetMaxRequests

`func (o *LeaderboardActivityTotals) GetMaxRequests() int64`

GetMaxRequests returns the MaxRequests field if non-nil, zero value otherwise.

### GetMaxRequestsOk

`func (o *LeaderboardActivityTotals) GetMaxRequestsOk() (*int64, bool)`

GetMaxRequestsOk returns a tuple with the MaxRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequests

`func (o *LeaderboardActivityTotals) SetMaxRequests(v int64)`

SetMaxRequests sets MaxRequests field to given value.

### HasMaxRequests

`func (o *LeaderboardActivityTotals) HasMaxRequests() bool`

HasMaxRequests returns a boolean if a field has been set.

### GetMaxTokens

`func (o *LeaderboardActivityTotals) GetMaxTokens() int64`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *LeaderboardActivityTotals) GetMaxTokensOk() (*int64, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *LeaderboardActivityTotals) SetMaxTokens(v int64)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *LeaderboardActivityTotals) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetRequests

`func (o *LeaderboardActivityTotals) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LeaderboardActivityTotals) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LeaderboardActivityTotals) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LeaderboardActivityTotals) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *LeaderboardActivityTotals) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LeaderboardActivityTotals) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LeaderboardActivityTotals) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LeaderboardActivityTotals) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


