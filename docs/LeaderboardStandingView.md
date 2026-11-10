# LeaderboardStandingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the warehouse is not connected: Rows is then empty because nothing could be read, not because nobody has used anything. | [optional] 
**Locked** | Pointer to **int64** | Locked is how many ranks lie above the first row, which the reader has not earned a view of. They are counted, never sent. | [optional] 
**Rows** | Pointer to [**[]LeaderboardStandingRow**](LeaderboardStandingRow.md) | Rows are the ranks the reader may see, best first. Always a list. | [optional] 
**Scored** | Pointer to **int64** | Scored is how many of them hold any point; the rest are at the bottom. | [optional] 
**Self** | Pointer to [**LeaderboardStandingSelf**](LeaderboardStandingSelf.md) | Self is the reader&#39;s own place; absent when not signed in. | [optional] 
**Total** | Pointer to **int64** | Total is how many orgs are ranked: every org IAM holds that is not a held one. | [optional] 

## Methods

### NewLeaderboardStandingView

`func NewLeaderboardStandingView() *LeaderboardStandingView`

NewLeaderboardStandingView instantiates a new LeaderboardStandingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardStandingViewWithDefaults

`func NewLeaderboardStandingViewWithDefaults() *LeaderboardStandingView`

NewLeaderboardStandingViewWithDefaults instantiates a new LeaderboardStandingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LeaderboardStandingView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LeaderboardStandingView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LeaderboardStandingView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LeaderboardStandingView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetLocked

`func (o *LeaderboardStandingView) GetLocked() int64`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *LeaderboardStandingView) GetLockedOk() (*int64, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *LeaderboardStandingView) SetLocked(v int64)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *LeaderboardStandingView) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetRows

`func (o *LeaderboardStandingView) GetRows() []LeaderboardStandingRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *LeaderboardStandingView) GetRowsOk() (*[]LeaderboardStandingRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *LeaderboardStandingView) SetRows(v []LeaderboardStandingRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *LeaderboardStandingView) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetScored

`func (o *LeaderboardStandingView) GetScored() int64`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *LeaderboardStandingView) GetScoredOk() (*int64, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *LeaderboardStandingView) SetScored(v int64)`

SetScored sets Scored field to given value.

### HasScored

`func (o *LeaderboardStandingView) HasScored() bool`

HasScored returns a boolean if a field has been set.

### GetSelf

`func (o *LeaderboardStandingView) GetSelf() LeaderboardStandingSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LeaderboardStandingView) GetSelfOk() (*LeaderboardStandingSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LeaderboardStandingView) SetSelf(v LeaderboardStandingSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *LeaderboardStandingView) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTotal

`func (o *LeaderboardStandingView) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LeaderboardStandingView) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LeaderboardStandingView) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LeaderboardStandingView) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


