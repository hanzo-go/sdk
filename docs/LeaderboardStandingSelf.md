# LeaderboardStandingSelf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is the reader&#39;s own full handle. | [optional] 
**Next** | Pointer to [**LeaderboardStandingNext**](LeaderboardStandingNext.md) | Next is the next rung to unlock; absent when nothing above is locked. | [optional] 
**Points** | Pointer to **int64** | Points is the reader&#39;s score. | [optional] 
**Rank** | Pointer to **int64** | Rank is the reader&#39;s place on the whole board. | [optional] 
**Ranked** | Pointer to **bool** | Ranked is false for a reader with no usage and no award yet. Rank is then 0 and the window is the one a newcomer sees. | [optional] 

## Methods

### NewLeaderboardStandingSelf

`func NewLeaderboardStandingSelf() *LeaderboardStandingSelf`

NewLeaderboardStandingSelf instantiates a new LeaderboardStandingSelf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardStandingSelfWithDefaults

`func NewLeaderboardStandingSelfWithDefaults() *LeaderboardStandingSelf`

NewLeaderboardStandingSelfWithDefaults instantiates a new LeaderboardStandingSelf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *LeaderboardStandingSelf) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LeaderboardStandingSelf) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LeaderboardStandingSelf) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *LeaderboardStandingSelf) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetNext

`func (o *LeaderboardStandingSelf) GetNext() LeaderboardStandingNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *LeaderboardStandingSelf) GetNextOk() (*LeaderboardStandingNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *LeaderboardStandingSelf) SetNext(v LeaderboardStandingNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *LeaderboardStandingSelf) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPoints

`func (o *LeaderboardStandingSelf) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *LeaderboardStandingSelf) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *LeaderboardStandingSelf) SetPoints(v int64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *LeaderboardStandingSelf) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetRank

`func (o *LeaderboardStandingSelf) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *LeaderboardStandingSelf) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *LeaderboardStandingSelf) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *LeaderboardStandingSelf) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRanked

`func (o *LeaderboardStandingSelf) GetRanked() bool`

GetRanked returns the Ranked field if non-nil, zero value otherwise.

### GetRankedOk

`func (o *LeaderboardStandingSelf) GetRankedOk() (*bool, bool)`

GetRankedOk returns a tuple with the Ranked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanked

`func (o *LeaderboardStandingSelf) SetRanked(v bool)`

SetRanked sets Ranked field to given value.

### HasRanked

`func (o *LeaderboardStandingSelf) HasRanked() bool`

HasRanked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


