# LeaderboardOrgOptinView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanManage** | Pointer to **bool** | CanManage is true only for an admin of this org (or a platform SuperAdmin) — the callers whose write of the org preference will be accepted. | [optional] 
**Display** | Pointer to **string** | Display is the name shown for the org on that board. Empty when none was chosen; opting in without one defaults it to the org id. | [optional] 
**Listed** | Pointer to **bool** | Listed is true when the org has opted onto the cross-org global board. False — the default — keeps the org off it entirely; the org&#39;s own members still see their own board. Listing consents to publishing usage VOLUME, never spend. | [optional] 

## Methods

### NewLeaderboardOrgOptinView

`func NewLeaderboardOrgOptinView() *LeaderboardOrgOptinView`

NewLeaderboardOrgOptinView instantiates a new LeaderboardOrgOptinView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardOrgOptinViewWithDefaults

`func NewLeaderboardOrgOptinViewWithDefaults() *LeaderboardOrgOptinView`

NewLeaderboardOrgOptinViewWithDefaults instantiates a new LeaderboardOrgOptinView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanManage

`func (o *LeaderboardOrgOptinView) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *LeaderboardOrgOptinView) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *LeaderboardOrgOptinView) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *LeaderboardOrgOptinView) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetDisplay

`func (o *LeaderboardOrgOptinView) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *LeaderboardOrgOptinView) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *LeaderboardOrgOptinView) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *LeaderboardOrgOptinView) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetListed

`func (o *LeaderboardOrgOptinView) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *LeaderboardOrgOptinView) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *LeaderboardOrgOptinView) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *LeaderboardOrgOptinView) HasListed() bool`

HasListed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


