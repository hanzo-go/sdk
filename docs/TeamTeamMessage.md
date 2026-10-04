# TeamTeamMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is the team account uuid that wrote it. It is an ACCOUNT and not a display name: what to call somebody is the roster&#39;s answer (GET /v1/team/members), and copying it onto every message is how the two come to disagree. An agent&#39;s messages carry the account derived from its id, so the same field answers for both. | [optional] 
**CreatedOn** | Pointer to **int64** | CreatedOn is unix MILLIseconds, which is what the platform stamps. | [optional] 
**Doc** | Pointer to **string** | Doc is the document this message is a comment on. Empty for a message in a room. | [optional] 
**EditedOn** | Pointer to **int64** | EditedOn is when the author last edited the message, unix milliseconds. Absent for a message never edited. | [optional] 
**Files** | Pointer to [**[]TeamTeamFile**](TeamTeamFile.md) | Files are the files attached to the message. | [optional] 
**Id** | Pointer to **string** | ID is the message document&#39;s own id. | [optional] 
**LastReply** | Pointer to **int64** | LastReply is when the thread was last answered, unix milliseconds. Absent for a message with no replies. | [optional] 
**Mentions** | Pointer to **[]string** | Mentions are the account uuids the message @-mentions — exactly the people notify.go told they were mentioned. | [optional] 
**Reactions** | Pointer to [**[]TeamTeamReaction**](TeamTeamReaction.md) | Reactions are the emoji people reacted with, one entry per emoji, in the order each emoji was first used. | [optional] 
**Replies** | Pointer to **int64** | Replies is how many replies the message&#39;s thread holds. | [optional] 
**Room** | Pointer to **string** | Room is the room it was said in — the same id the room listing answers with, so a caller holding a message can name its room without a second read. A reply names the room its thread is in; a comment on a document has none. | [optional] 
**Text** | Pointer to **string** | Text is the message as PLAIN TEXT. The document stores markup; this is the same &#x60;plainText&#x60; reduction the agent responder reads a prompt with, so a caller never has to parse the client&#39;s markup to know what was said. A mention reads as \&quot;@Name\&quot; here and is listed by id in Mentions. | [optional] 
**Thread** | Pointer to **string** | Thread is the message this one replies to. Empty for a message that starts a conversation rather than answering one. | [optional] 

## Methods

### NewTeamTeamMessage

`func NewTeamTeamMessage() *TeamTeamMessage`

NewTeamTeamMessage instantiates a new TeamTeamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamMessageWithDefaults

`func NewTeamTeamMessageWithDefaults() *TeamTeamMessage`

NewTeamTeamMessageWithDefaults instantiates a new TeamTeamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *TeamTeamMessage) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TeamTeamMessage) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TeamTeamMessage) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *TeamTeamMessage) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreatedOn

`func (o *TeamTeamMessage) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TeamTeamMessage) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TeamTeamMessage) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TeamTeamMessage) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDoc

`func (o *TeamTeamMessage) GetDoc() string`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *TeamTeamMessage) GetDocOk() (*string, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *TeamTeamMessage) SetDoc(v string)`

SetDoc sets Doc field to given value.

### HasDoc

`func (o *TeamTeamMessage) HasDoc() bool`

HasDoc returns a boolean if a field has been set.

### GetEditedOn

`func (o *TeamTeamMessage) GetEditedOn() int64`

GetEditedOn returns the EditedOn field if non-nil, zero value otherwise.

### GetEditedOnOk

`func (o *TeamTeamMessage) GetEditedOnOk() (*int64, bool)`

GetEditedOnOk returns a tuple with the EditedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedOn

`func (o *TeamTeamMessage) SetEditedOn(v int64)`

SetEditedOn sets EditedOn field to given value.

### HasEditedOn

`func (o *TeamTeamMessage) HasEditedOn() bool`

HasEditedOn returns a boolean if a field has been set.

### GetFiles

`func (o *TeamTeamMessage) GetFiles() []TeamTeamFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TeamTeamMessage) GetFilesOk() (*[]TeamTeamFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TeamTeamMessage) SetFiles(v []TeamTeamFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TeamTeamMessage) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastReply

`func (o *TeamTeamMessage) GetLastReply() int64`

GetLastReply returns the LastReply field if non-nil, zero value otherwise.

### GetLastReplyOk

`func (o *TeamTeamMessage) GetLastReplyOk() (*int64, bool)`

GetLastReplyOk returns a tuple with the LastReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReply

`func (o *TeamTeamMessage) SetLastReply(v int64)`

SetLastReply sets LastReply field to given value.

### HasLastReply

`func (o *TeamTeamMessage) HasLastReply() bool`

HasLastReply returns a boolean if a field has been set.

### GetMentions

`func (o *TeamTeamMessage) GetMentions() []string`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *TeamTeamMessage) GetMentionsOk() (*[]string, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *TeamTeamMessage) SetMentions(v []string)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *TeamTeamMessage) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetReactions

`func (o *TeamTeamMessage) GetReactions() []TeamTeamReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *TeamTeamMessage) GetReactionsOk() (*[]TeamTeamReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *TeamTeamMessage) SetReactions(v []TeamTeamReaction)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *TeamTeamMessage) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetReplies

`func (o *TeamTeamMessage) GetReplies() int64`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *TeamTeamMessage) GetRepliesOk() (*int64, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *TeamTeamMessage) SetReplies(v int64)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *TeamTeamMessage) HasReplies() bool`

HasReplies returns a boolean if a field has been set.

### GetRoom

`func (o *TeamTeamMessage) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *TeamTeamMessage) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *TeamTeamMessage) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *TeamTeamMessage) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetText

`func (o *TeamTeamMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TeamTeamMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TeamTeamMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TeamTeamMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThread

`func (o *TeamTeamMessage) GetThread() string`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *TeamTeamMessage) GetThreadOk() (*string, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *TeamTeamMessage) SetThread(v string)`

SetThread sets Thread field to given value.

### HasThread

`func (o *TeamTeamMessage) HasThread() bool`

HasThread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


