# LeaderboardView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the usage warehouse is not connected or its rollup is not ready. Rows is then empty because nothing could be read — not because nobody used anything. Show that difference; never render an unavailable board as a real one. | [optional] 
**End** | Pointer to **string** | End is the EXCLUSIVE upper bound of the window, \&quot;2006-01-02\&quot; — the day after the last one counted. A board through today reports tomorrow&#39;s date here. | [optional] 
**Metric** | Pointer to **string** | Metric echoes the value ranked: tokens|requests|cost. | [optional] 
**Period** | Pointer to **string** | Period is the window&#39;s canonical label: day|week|month|all. The server resolves aliases (7d, 30d, today, …) to these, so this may differ from what was sent. | [optional] 
**Rows** | Pointer to [**[]LeaderboardRow**](LeaderboardRow.md) | Rows are the ranked subjects, best first, at most the requested limit of them. Always a list, never null: an empty one means nothing was read, not an error. | [optional] 
**Scope** | Pointer to **string** | Scope echoes the board that was served: personal|org|global. | [optional] 
**Self** | Pointer to [**SelfRank**](SelfRank.md) | Self is the caller&#39;s own standing, reported even when they fall outside Rows. Absent when the caller&#39;s ledger identity cannot be resolved, or when the query behind it failed — never faked to keep the shape tidy. | [optional] 
**Source** | Pointer to **string** | Source names the table these numbers were aggregated from (the derived daily rollup, hanzo.usage_rollup_daily), so an operator can tell exactly what was read. | [optional] 
**Start** | Pointer to **string** | Start is the first day counted, \&quot;2006-01-02\&quot; inclusive. Empty for period&#x3D;all, which has no lower bound at all. | [optional] 
**Subject** | Pointer to **string** | Subject is what the rows stand for — \&quot;user\&quot; on a personal or org board, \&quot;org\&quot; on the global one. It tells a client whether Handle names a person or a company. | [optional] 
**Total** | Pointer to **int64** | Total is how many subjects were ranked in the window — the org&#39;s active users, or the active/opted-in orgs on the global board. It is the universe the ranks are out of, so it is normally larger than len(rows). | [optional] 

## Methods

### NewLeaderboardView

`func NewLeaderboardView() *LeaderboardView`

NewLeaderboardView instantiates a new LeaderboardView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardViewWithDefaults

`func NewLeaderboardViewWithDefaults() *LeaderboardView`

NewLeaderboardViewWithDefaults instantiates a new LeaderboardView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LeaderboardView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LeaderboardView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LeaderboardView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LeaderboardView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetEnd

`func (o *LeaderboardView) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LeaderboardView) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LeaderboardView) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *LeaderboardView) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetMetric

`func (o *LeaderboardView) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *LeaderboardView) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *LeaderboardView) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *LeaderboardView) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPeriod

`func (o *LeaderboardView) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LeaderboardView) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LeaderboardView) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LeaderboardView) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRows

`func (o *LeaderboardView) GetRows() []LeaderboardRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *LeaderboardView) GetRowsOk() (*[]LeaderboardRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *LeaderboardView) SetRows(v []LeaderboardRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *LeaderboardView) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetScope

`func (o *LeaderboardView) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LeaderboardView) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LeaderboardView) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *LeaderboardView) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *LeaderboardView) GetSelf() SelfRank`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LeaderboardView) GetSelfOk() (*SelfRank, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LeaderboardView) SetSelf(v SelfRank)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *LeaderboardView) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSource

`func (o *LeaderboardView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LeaderboardView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LeaderboardView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LeaderboardView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStart

`func (o *LeaderboardView) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LeaderboardView) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LeaderboardView) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *LeaderboardView) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSubject

`func (o *LeaderboardView) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LeaderboardView) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LeaderboardView) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *LeaderboardView) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTotal

`func (o *LeaderboardView) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LeaderboardView) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LeaderboardView) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LeaderboardView) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


