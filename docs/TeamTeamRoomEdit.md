# TeamTeamRoomEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | Archived closes the room (true) or reopens it (false). An archived room keeps its history and takes no new messages. | [optional] 
**Id** | Pointer to **string** | ID is the room, from the path. | [optional] 
**Name** | Pointer to **string** | Name renames a channel. A direct message has no name. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid holding it. Body-only. | [optional] 
**Topic** | Pointer to **string** | Topic sets the channel&#39;s one-line subject; \&quot;\&quot; clears it. | [optional] 

## Methods

### NewTeamTeamRoomEdit

`func NewTeamTeamRoomEdit() *TeamTeamRoomEdit`

NewTeamTeamRoomEdit instantiates a new TeamTeamRoomEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamRoomEditWithDefaults

`func NewTeamTeamRoomEditWithDefaults() *TeamTeamRoomEdit`

NewTeamTeamRoomEditWithDefaults instantiates a new TeamTeamRoomEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TeamTeamRoomEdit) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TeamTeamRoomEdit) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TeamTeamRoomEdit) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *TeamTeamRoomEdit) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamRoomEdit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamRoomEdit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamRoomEdit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamRoomEdit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamRoomEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamRoomEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamRoomEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamRoomEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamRoomEdit) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamRoomEdit) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamRoomEdit) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamRoomEdit) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTopic

`func (o *TeamTeamRoomEdit) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TeamTeamRoomEdit) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TeamTeamRoomEdit) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TeamTeamRoomEdit) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


