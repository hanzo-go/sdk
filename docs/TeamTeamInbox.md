# TeamTeamInbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]TeamTeamInboxItem**](TeamTeamInboxItem.md) | Items are the notifications, newest first, at most 200. | [optional] 
**Unread** | Pointer to **int64** | Unread is how many of the caller&#39;s live (unarchived) notifications are unread — the number a badge shows, counted over all of them rather than the page. | [optional] 

## Methods

### NewTeamTeamInbox

`func NewTeamTeamInbox() *TeamTeamInbox`

NewTeamTeamInbox instantiates a new TeamTeamInbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamInboxWithDefaults

`func NewTeamTeamInboxWithDefaults() *TeamTeamInbox`

NewTeamTeamInboxWithDefaults instantiates a new TeamTeamInbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TeamTeamInbox) GetItems() []TeamTeamInboxItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TeamTeamInbox) GetItemsOk() (*[]TeamTeamInboxItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TeamTeamInbox) SetItems(v []TeamTeamInboxItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *TeamTeamInbox) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetUnread

`func (o *TeamTeamInbox) GetUnread() int64`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *TeamTeamInbox) GetUnreadOk() (*int64, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *TeamTeamInbox) SetUnread(v int64)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *TeamTeamInbox) HasUnread() bool`

HasUnread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


