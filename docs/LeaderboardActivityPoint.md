# LeaderboardActivityPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int64** | CostCents is the day&#39;s spend in whole US cents. A series is only ever returned for a subject the caller is authorized to see, so this is never withheld: 0 means no spend that day. | [optional] 
**Day** | Pointer to **string** | Day is the UTC calendar day this point covers, \&quot;2006-01-02\&quot;. | [optional] 
**Requests** | Pointer to **int64** | Requests is the subject&#39;s request count on this day. 0 is a real, quiet day: the series is gap-filled, so every day in the range is present whether or not anything happened. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is prompt+completion tokens on this day — normally the heatmap&#39;s intensity, scaled against ActivityTotals.MaxTokens. | [optional] 

## Methods

### NewLeaderboardActivityPoint

`func NewLeaderboardActivityPoint() *LeaderboardActivityPoint`

NewLeaderboardActivityPoint instantiates a new LeaderboardActivityPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardActivityPointWithDefaults

`func NewLeaderboardActivityPointWithDefaults() *LeaderboardActivityPoint`

NewLeaderboardActivityPointWithDefaults instantiates a new LeaderboardActivityPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *LeaderboardActivityPoint) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LeaderboardActivityPoint) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LeaderboardActivityPoint) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LeaderboardActivityPoint) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetDay

`func (o *LeaderboardActivityPoint) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *LeaderboardActivityPoint) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *LeaderboardActivityPoint) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *LeaderboardActivityPoint) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetRequests

`func (o *LeaderboardActivityPoint) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LeaderboardActivityPoint) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LeaderboardActivityPoint) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LeaderboardActivityPoint) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *LeaderboardActivityPoint) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LeaderboardActivityPoint) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LeaderboardActivityPoint) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LeaderboardActivityPoint) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


