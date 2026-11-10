# TeamTeamChannelEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | Archived closes the channel (true) or reopens it (false). An archived channel keeps its history and takes no new messages. | [optional] 
**Id** | Pointer to **string** | ID is the channel, from the path. | [optional] 
**Name** | Pointer to **string** | Name renames a channel. A direct message has no name. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid holding it. Body-only. | [optional] 
**Topic** | Pointer to **string** | Topic sets the channel&#39;s one-line subject; \&quot;\&quot; clears it. | [optional] 

## Methods

### NewTeamTeamChannelEdit

`func NewTeamTeamChannelEdit() *TeamTeamChannelEdit`

NewTeamTeamChannelEdit instantiates a new TeamTeamChannelEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamChannelEditWithDefaults

`func NewTeamTeamChannelEditWithDefaults() *TeamTeamChannelEdit`

NewTeamTeamChannelEditWithDefaults instantiates a new TeamTeamChannelEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TeamTeamChannelEdit) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TeamTeamChannelEdit) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TeamTeamChannelEdit) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *TeamTeamChannelEdit) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamChannelEdit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamChannelEdit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamChannelEdit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamChannelEdit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamChannelEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamChannelEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamChannelEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamChannelEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamChannelEdit) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamChannelEdit) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamChannelEdit) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamChannelEdit) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTopic

`func (o *TeamTeamChannelEdit) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TeamTeamChannelEdit) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TeamTeamChannelEdit) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TeamTeamChannelEdit) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


