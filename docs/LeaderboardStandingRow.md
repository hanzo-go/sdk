# LeaderboardStandingRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is the name to show: masked, or full for the reader&#39;s own row and for an org that listed. Empty for an org that withdrew; show it as \&quot;A builder\&quot;. | [optional] 
**Points** | Pointer to **int64** | Points is the score the board ranks by. | [optional] 
**Rank** | Pointer to **int64** | Rank is the 1-based place on the whole board, not the window. | [optional] 
**Self** | Pointer to **bool** | Self marks the reader&#39;s own row. | [optional] 

## Methods

### NewLeaderboardStandingRow

`func NewLeaderboardStandingRow() *LeaderboardStandingRow`

NewLeaderboardStandingRow instantiates a new LeaderboardStandingRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardStandingRowWithDefaults

`func NewLeaderboardStandingRowWithDefaults() *LeaderboardStandingRow`

NewLeaderboardStandingRowWithDefaults instantiates a new LeaderboardStandingRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *LeaderboardStandingRow) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LeaderboardStandingRow) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LeaderboardStandingRow) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *LeaderboardStandingRow) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetPoints

`func (o *LeaderboardStandingRow) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *LeaderboardStandingRow) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *LeaderboardStandingRow) SetPoints(v int64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *LeaderboardStandingRow) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetRank

`func (o *LeaderboardStandingRow) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *LeaderboardStandingRow) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *LeaderboardStandingRow) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *LeaderboardStandingRow) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetSelf

`func (o *LeaderboardStandingRow) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LeaderboardStandingRow) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LeaderboardStandingRow) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *LeaderboardStandingRow) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


