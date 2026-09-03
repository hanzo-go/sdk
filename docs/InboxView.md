# InboxView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the lowercased external id of the org&#39;s connected account on that transport: the Discord guild id, the Slack team id, the Teams AAD tenant id, or the bound Telegram chat id. Informational only — the gate keys on (org, channel), never on the account. | [optional] 
**Channel** | Pointer to **string** | Channel is the transport this message arrived on — discord, slack, teams or telegram — and the &#x60;:channel&#x60; segment to reply through. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is Unix SECONDS, stamped by the ingest goroutine when the message was accepted — not the transport&#39;s own send time. Rows are dropped 30 days after it. | [optional] 
**Id** | Pointer to **int64** | ID is the store&#39;s row id, assigned on insert — SERVER-SET, and the cursor: pass a page&#39;s last id back as &#x60;since&#x60;. It rises with arrival order but is not contiguous, because one sequence is shared by every org in the store and a caller reads only its own rows. | [optional] 
**ReplyTo** | Pointer to **string** | ReplyTo is the transport&#39;s reply target for this message: Slack&#39;s thread_ts, or the Telegram message id it arrived as. Send it back as the body&#39;s &#x60;replyTo&#x60; to answer in the SAME thread. Empty means the transport reported none — a top-level Slack message, and every Discord and Teams message, since neither carries one — and a reply then lands at the top level of the room. | [optional] 
**RoomId** | Pointer to **string** | RoomID is the conversation on the ORIGINATING transport, and the value to send back as &#x60;room.id&#x60;: a Discord channel snowflake, a Slack conversation id (D… IM, C… public channel, G… private or mpim), a Teams conversation id (19:…@thread.… for a channel or group chat, a:… for a personal chat), or a Telegram chat id in decimal (negative for a group, positive for a DM). It is stable for the life of the room, so every message from one conversation carries the same value. | [optional] 
**RoomKind** | Pointer to **string** | RoomKind is how ingest classified the room: \&quot;dm\&quot;, \&quot;group\&quot; or \&quot;thread\&quot;. It decides which policy gated the message — dmPolicy for \&quot;dm\&quot;, groupPolicy for BOTH \&quot;group\&quot; and \&quot;thread\&quot;. Only Slack ever reports \&quot;thread\&quot;; Telegram&#39;s reply-to id becomes ReplyTo instead, and Discord&#39;s ingress is guild-scoped so its rooms are always \&quot;group\&quot;. | [optional] 
**Sender** | Pointer to **string** | Sender is the TRANSPORT-NATIVE user id of whoever wrote the message — a Discord member.user.id, a Slack U… user id, a Teams aadObjectId (falling back to from.id), a Telegram from.id in decimal. Stable per person per transport, and the identity the gate keys on: an allow entry, an access-group member and a pairing approval all name exactly this value. | [optional] 
**SenderUser** | Pointer to **string** | SenderUser is the HANZO account subject that chat identity is linked to, resolved at ingest through the org&#39;s user link. Best-effort and omitted when absent: a person who never linked their chat account — or a link store that could not be read — leaves it empty and is never blocked for it. | [optional] 
**Text** | Pointer to **string** | Text is the body as the transport delivered it, with the bot mention already stripped by the ingress adapter (on Discord it is the /hanzo prompt argument, since that ingress is slash commands only), truncated to 8 KiB on store. Inbound attachments are not stored — this is the whole of what was said. | [optional] 

## Methods

### NewInboxView

`func NewInboxView() *InboxView`

NewInboxView instantiates a new InboxView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxViewWithDefaults

`func NewInboxViewWithDefaults() *InboxView`

NewInboxViewWithDefaults instantiates a new InboxView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *InboxView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InboxView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InboxView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InboxView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChannel

`func (o *InboxView) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InboxView) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InboxView) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InboxView) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InboxView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InboxView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InboxView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InboxView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *InboxView) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboxView) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboxView) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InboxView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReplyTo

`func (o *InboxView) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *InboxView) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *InboxView) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *InboxView) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetRoomId

`func (o *InboxView) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *InboxView) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *InboxView) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *InboxView) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetRoomKind

`func (o *InboxView) GetRoomKind() string`

GetRoomKind returns the RoomKind field if non-nil, zero value otherwise.

### GetRoomKindOk

`func (o *InboxView) GetRoomKindOk() (*string, bool)`

GetRoomKindOk returns a tuple with the RoomKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomKind

`func (o *InboxView) SetRoomKind(v string)`

SetRoomKind sets RoomKind field to given value.

### HasRoomKind

`func (o *InboxView) HasRoomKind() bool`

HasRoomKind returns a boolean if a field has been set.

### GetSender

`func (o *InboxView) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *InboxView) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *InboxView) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *InboxView) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSenderUser

`func (o *InboxView) GetSenderUser() string`

GetSenderUser returns the SenderUser field if non-nil, zero value otherwise.

### GetSenderUserOk

`func (o *InboxView) GetSenderUserOk() (*string, bool)`

GetSenderUserOk returns a tuple with the SenderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUser

`func (o *InboxView) SetSenderUser(v string)`

SetSenderUser sets SenderUser field to given value.

### HasSenderUser

`func (o *InboxView) HasSenderUser() bool`

HasSenderUser returns a boolean if a field has been set.

### GetText

`func (o *InboxView) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InboxView) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InboxView) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *InboxView) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


