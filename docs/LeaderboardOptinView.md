# LeaderboardOptinView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to [**LeaderboardOrgOptinView**](LeaderboardOrgOptinView.md) | Org is the caller&#39;s org&#39;s listing preference on the cross-org board, and whether this caller is allowed to change it. It is read for every caller — a member sees where their org stands even though only an admin may edit it. | [optional] 
**User** | Pointer to [**LeaderboardUserOptinView**](LeaderboardUserOptinView.md) | User is the caller&#39;s OWN listing preference, and whether they may change it. | [optional] 

## Methods

### NewLeaderboardOptinView

`func NewLeaderboardOptinView() *LeaderboardOptinView`

NewLeaderboardOptinView instantiates a new LeaderboardOptinView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardOptinViewWithDefaults

`func NewLeaderboardOptinViewWithDefaults() *LeaderboardOptinView`

NewLeaderboardOptinViewWithDefaults instantiates a new LeaderboardOptinView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *LeaderboardOptinView) GetOrg() LeaderboardOrgOptinView`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *LeaderboardOptinView) GetOrgOk() (*LeaderboardOrgOptinView, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *LeaderboardOptinView) SetOrg(v LeaderboardOrgOptinView)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *LeaderboardOptinView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetUser

`func (o *LeaderboardOptinView) GetUser() LeaderboardUserOptinView`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LeaderboardOptinView) GetUserOk() (*LeaderboardUserOptinView, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LeaderboardOptinView) SetUser(v LeaderboardUserOptinView)`

SetUser sets User field to given value.

### HasUser

`func (o *LeaderboardOptinView) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


