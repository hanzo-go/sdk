# LeaderboardLeaderboardRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anonymous** | Pointer to **bool** | Anonymous is true when this subject&#39;s identity was withheld and Handle is the \&quot;Anonymous\&quot; placeholder: the metric is real, the name is not. Render it as an unnamed row, never as someone actually called Anonymous. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is this subject&#39;s spend in whole US cents (not dollars, not millicents). It is 0 unless the viewer is entitled to this subject&#39;s spend — their own row, an admin looking at their own org, an explicitly cost-ranked board, or a platform admin on the global board — so 0 means \&quot;withheld or zero\&quot;, and the two are deliberately indistinguishable. | [optional] 
**Handle** | Pointer to **string** | Handle is the display identity to render: the peer&#39;s chosen handle if they opted into public listing, their username if the viewer is an admin of their org, the org&#39;s chosen display name (or its id) on an org board, and the literal \&quot;Anonymous\&quot; when the identity is withheld. Never an email or a raw ledger id. | [optional] 
**Metric** | Pointer to **int64** | Metric is the value the board was ranked by, copied from Requests, Tokens or CostCents according to the request&#39;s metric. It is what a client sizes bars against without having to know which metric was asked for. | [optional] 
**Rank** | Pointer to **int64** | Rank is this subject&#39;s 1-based standing in the window, 1 being the top. Rows are ordered by the ranked metric descending (ties broken by request count), and a board is always the top of the list — there is no offset paging — so the first row is always rank 1 and rank is also the row&#39;s index + 1. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many AI requests this subject made in the window. Volume is not sensitive, so it is reported for every row including anonymized ones. | [optional] 
**Self** | Pointer to **bool** | Self marks the caller&#39;s own row so a client can highlight it in place. At most one row carries it, and it is never set on an org board — org rows carry no user identity, so the caller&#39;s own org standing arrives in LeaderboardView.Self. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is prompt+completion tokens this subject spent in the window. Like Requests it is reported for every row. | [optional] 

## Methods

### NewLeaderboardLeaderboardRow

`func NewLeaderboardLeaderboardRow() *LeaderboardLeaderboardRow`

NewLeaderboardLeaderboardRow instantiates a new LeaderboardLeaderboardRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardLeaderboardRowWithDefaults

`func NewLeaderboardLeaderboardRowWithDefaults() *LeaderboardLeaderboardRow`

NewLeaderboardLeaderboardRowWithDefaults instantiates a new LeaderboardLeaderboardRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymous

`func (o *LeaderboardLeaderboardRow) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *LeaderboardLeaderboardRow) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *LeaderboardLeaderboardRow) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *LeaderboardLeaderboardRow) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetCostCents

`func (o *LeaderboardLeaderboardRow) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LeaderboardLeaderboardRow) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LeaderboardLeaderboardRow) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LeaderboardLeaderboardRow) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetHandle

`func (o *LeaderboardLeaderboardRow) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LeaderboardLeaderboardRow) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LeaderboardLeaderboardRow) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *LeaderboardLeaderboardRow) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetMetric

`func (o *LeaderboardLeaderboardRow) GetMetric() int64`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *LeaderboardLeaderboardRow) GetMetricOk() (*int64, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *LeaderboardLeaderboardRow) SetMetric(v int64)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *LeaderboardLeaderboardRow) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetRank

`func (o *LeaderboardLeaderboardRow) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *LeaderboardLeaderboardRow) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *LeaderboardLeaderboardRow) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *LeaderboardLeaderboardRow) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRequests

`func (o *LeaderboardLeaderboardRow) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LeaderboardLeaderboardRow) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LeaderboardLeaderboardRow) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LeaderboardLeaderboardRow) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSelf

`func (o *LeaderboardLeaderboardRow) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LeaderboardLeaderboardRow) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LeaderboardLeaderboardRow) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *LeaderboardLeaderboardRow) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTokens

`func (o *LeaderboardLeaderboardRow) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LeaderboardLeaderboardRow) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LeaderboardLeaderboardRow) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LeaderboardLeaderboardRow) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


