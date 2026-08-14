# OptinView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to [**OrgOptinView**](OrgOptinView.md) | Org is the caller&#39;s org&#39;s listing preference on the cross-org board, and whether this caller is allowed to change it. It is read for every caller — a member sees where their org stands even though only an admin may edit it. | [optional] 
**User** | Pointer to [**UserOptinView**](UserOptinView.md) | User is the caller&#39;s OWN listing preference, and whether they may change it. | [optional] 

## Methods

### NewOptinView

`func NewOptinView() *OptinView`

NewOptinView instantiates a new OptinView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptinViewWithDefaults

`func NewOptinViewWithDefaults() *OptinView`

NewOptinViewWithDefaults instantiates a new OptinView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *OptinView) GetOrg() OrgOptinView`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *OptinView) GetOrgOk() (*OrgOptinView, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *OptinView) SetOrg(v OrgOptinView)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *OptinView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetUser

`func (o *OptinView) GetUser() UserOptinView`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OptinView) GetUserOk() (*UserOptinView, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OptinView) SetUser(v UserOptinView)`

SetUser sets User field to given value.

### HasUser

`func (o *OptinView) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


