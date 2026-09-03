# SelfRank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int64** | CostCents is the caller&#39;s own spend in whole US cents. Always populated — your own spend is never withheld from you — so here 0 really does mean zero. | [optional] 
**Handle** | Pointer to **string** | Handle is how the caller appears on this board: their chosen handle, falling back to their username, on a user board; their org id on the global board. Present even when unlisted — this is the caller looking at themselves. | [optional] 
**Listed** | Pointer to **bool** | Listed says whether the caller is publicly visible on this board: opted in on a user board, org opted in (or the viewer is a platform admin) on the global one. False is the prompt to offer the opt-in, and explains an unranked global self. | [optional] 
**Metric** | Pointer to **int64** | Metric is whichever of the three values above the board was ranked by, so a client can compare the caller against the rows without re-reading the request. Metric &lt;&#x3D; 0 is exactly the case that leaves Ranked false. | [optional] 
**OfTotal** | Pointer to **int64** | OfTotal is the size of the universe Rank is out of — \&quot;rank N of OfTotal\&quot;. On a user board that is the org&#39;s users with any usage in the window; on the global board it is every active org for a platform admin, and the count of opted-in orgs for everyone else. | [optional] 
**Rank** | Pointer to **int64** | Rank is the caller&#39;s 1-based standing, computed as (subjects whose windowed metric strictly exceeds the caller&#39;s) + 1. It is exact against the whole ranked universe, not just the returned page, so it can far exceed len(rows). Read it only when Ranked. | [optional] 
**Ranked** | Pointer to **bool** | Ranked is false when the caller holds no position: they had no usage in the window, or (on the global board) their org has not opted into public listing and so is not ranked against a set it never joined. Rank is then 0 and means nothing. | [optional] 
**Requests** | Pointer to **int64** | Requests is the caller&#39;s own request count in the window, 0 if they were idle. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is the caller&#39;s own prompt+completion tokens in the window. | [optional] 

## Methods

### NewSelfRank

`func NewSelfRank() *SelfRank`

NewSelfRank instantiates a new SelfRank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfRankWithDefaults

`func NewSelfRankWithDefaults() *SelfRank`

NewSelfRankWithDefaults instantiates a new SelfRank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *SelfRank) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *SelfRank) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *SelfRank) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *SelfRank) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetHandle

`func (o *SelfRank) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *SelfRank) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *SelfRank) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *SelfRank) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetListed

`func (o *SelfRank) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *SelfRank) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *SelfRank) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *SelfRank) HasListed() bool`

HasListed returns a boolean if a field has been set.

### GetMetric

`func (o *SelfRank) GetMetric() int64`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SelfRank) GetMetricOk() (*int64, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SelfRank) SetMetric(v int64)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SelfRank) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetOfTotal

`func (o *SelfRank) GetOfTotal() int64`

GetOfTotal returns the OfTotal field if non-nil, zero value otherwise.

### GetOfTotalOk

`func (o *SelfRank) GetOfTotalOk() (*int64, bool)`

GetOfTotalOk returns a tuple with the OfTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfTotal

`func (o *SelfRank) SetOfTotal(v int64)`

SetOfTotal sets OfTotal field to given value.

### HasOfTotal

`func (o *SelfRank) HasOfTotal() bool`

HasOfTotal returns a boolean if a field has been set.

### GetRank

`func (o *SelfRank) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SelfRank) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SelfRank) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SelfRank) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRanked

`func (o *SelfRank) GetRanked() bool`

GetRanked returns the Ranked field if non-nil, zero value otherwise.

### GetRankedOk

`func (o *SelfRank) GetRankedOk() (*bool, bool)`

GetRankedOk returns a tuple with the Ranked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanked

`func (o *SelfRank) SetRanked(v bool)`

SetRanked sets Ranked field to given value.

### HasRanked

`func (o *SelfRank) HasRanked() bool`

HasRanked returns a boolean if a field has been set.

### GetRequests

`func (o *SelfRank) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *SelfRank) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *SelfRank) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *SelfRank) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *SelfRank) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *SelfRank) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *SelfRank) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *SelfRank) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


