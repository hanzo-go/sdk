# LeaderboardUserOptinView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSet** | Pointer to **bool** | CanSet is false when the caller&#39;s ledger identity cannot be resolved (no user name on the principal). Writing the preference would fail, so hide the control. | [optional] 
**Handle** | Pointer to **string** | Handle is the display name on the caller&#39;s listed row. Empty when they never chose one; opting in without a handle sets it to their username, so a listed row is never blank. | [optional] 
**Listed** | Pointer to **bool** | Listed is true when the caller&#39;s board row is published under Handle to other viewers. False — the default for anyone who never opted in — anonymizes the row; the metric still counts, only the name is withheld. | [optional] 

## Methods

### NewLeaderboardUserOptinView

`func NewLeaderboardUserOptinView() *LeaderboardUserOptinView`

NewLeaderboardUserOptinView instantiates a new LeaderboardUserOptinView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardUserOptinViewWithDefaults

`func NewLeaderboardUserOptinViewWithDefaults() *LeaderboardUserOptinView`

NewLeaderboardUserOptinViewWithDefaults instantiates a new LeaderboardUserOptinView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSet

`func (o *LeaderboardUserOptinView) GetCanSet() bool`

GetCanSet returns the CanSet field if non-nil, zero value otherwise.

### GetCanSetOk

`func (o *LeaderboardUserOptinView) GetCanSetOk() (*bool, bool)`

GetCanSetOk returns a tuple with the CanSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSet

`func (o *LeaderboardUserOptinView) SetCanSet(v bool)`

SetCanSet sets CanSet field to given value.

### HasCanSet

`func (o *LeaderboardUserOptinView) HasCanSet() bool`

HasCanSet returns a boolean if a field has been set.

### GetHandle

`func (o *LeaderboardUserOptinView) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LeaderboardUserOptinView) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LeaderboardUserOptinView) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *LeaderboardUserOptinView) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetListed

`func (o *LeaderboardUserOptinView) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *LeaderboardUserOptinView) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *LeaderboardUserOptinView) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *LeaderboardUserOptinView) HasListed() bool`

HasListed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


