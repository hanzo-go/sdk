# TeamTeamInboxItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | Archived reports that the caller archived it. | [optional] 
**CreatedOn** | Pointer to **int64** | CreatedOn is when it was filed, unix milliseconds. | [optional] 
**Doc** | Pointer to **string** | Doc is the document it is about, when it is about a document. | [optional] 
**Id** | Pointer to **string** | ID is the notification&#39;s own id — what read and archive address. | [optional] 
**Message** | Pointer to [**TeamTeamMessage**](TeamTeamMessage.md) | Message is the message that caused it, as the message ops answer it. Absent when that message has since been deleted, or is in a room the caller can no longer see. | [optional] 
**Read** | Pointer to **bool** | Read reports that the caller has seen it. | [optional] 
**Reason** | Pointer to **string** | Reason is why it was filed: \&quot;mention\&quot;, \&quot;dm\&quot; (a direct message), \&quot;reply\&quot; (in a thread the caller is in), \&quot;comment\&quot; (on something the caller follows), \&quot;assigned\&quot;, or \&quot;other\&quot;. | [optional] 
**Room** | Pointer to **string** | Room is the room it is about, when it is about a room. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid it was filed in. | [optional] 
**Thread** | Pointer to **string** | Thread is the message whose thread it is about, for a reply. | [optional] 

## Methods

### NewTeamTeamInboxItem

`func NewTeamTeamInboxItem() *TeamTeamInboxItem`

NewTeamTeamInboxItem instantiates a new TeamTeamInboxItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamInboxItemWithDefaults

`func NewTeamTeamInboxItemWithDefaults() *TeamTeamInboxItem`

NewTeamTeamInboxItemWithDefaults instantiates a new TeamTeamInboxItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TeamTeamInboxItem) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TeamTeamInboxItem) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TeamTeamInboxItem) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *TeamTeamInboxItem) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCreatedOn

`func (o *TeamTeamInboxItem) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TeamTeamInboxItem) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TeamTeamInboxItem) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TeamTeamInboxItem) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDoc

`func (o *TeamTeamInboxItem) GetDoc() string`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *TeamTeamInboxItem) GetDocOk() (*string, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *TeamTeamInboxItem) SetDoc(v string)`

SetDoc sets Doc field to given value.

### HasDoc

`func (o *TeamTeamInboxItem) HasDoc() bool`

HasDoc returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamInboxItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamInboxItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamInboxItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamInboxItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *TeamTeamInboxItem) GetMessage() TeamTeamMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TeamTeamInboxItem) GetMessageOk() (*TeamTeamMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TeamTeamInboxItem) SetMessage(v TeamTeamMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TeamTeamInboxItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRead

`func (o *TeamTeamInboxItem) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *TeamTeamInboxItem) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *TeamTeamInboxItem) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *TeamTeamInboxItem) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetReason

`func (o *TeamTeamInboxItem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TeamTeamInboxItem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TeamTeamInboxItem) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TeamTeamInboxItem) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRoom

`func (o *TeamTeamInboxItem) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *TeamTeamInboxItem) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *TeamTeamInboxItem) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *TeamTeamInboxItem) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamInboxItem) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamInboxItem) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamInboxItem) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamInboxItem) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetThread

`func (o *TeamTeamInboxItem) GetThread() string`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *TeamTeamInboxItem) GetThreadOk() (*string, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *TeamTeamInboxItem) SetThread(v string)`

SetThread sets Thread field to given value.

### HasThread

`func (o *TeamTeamInboxItem) HasThread() bool`

HasThread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


