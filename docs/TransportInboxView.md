# TransportInboxView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the lowercased external id of the org&#39;s connected account on that transport: the Discord guild id, the Slack team id, the Teams AAD tenant id, or the bound Telegram chat id. Informational only — the gate keys on (org, transport), never on the account. | [optional] 
**ChannelId** | Pointer to **string** | ChannelID is the conversation on the ORIGINATING transport, and the value to send back as &#x60;channel.id&#x60;: a Discord channel snowflake, a Slack conversation id (D… IM, C… public channel, G… private or mpim), a Teams conversation id (19:…@thread.… for a Teams channel or group chat, a:… for a personal chat), or a Telegram chat id in decimal (negative for a group, positive for a DM). It is stable for the life of the channel, so every message from one conversation carries the same value. | [optional] 
**ChannelKind** | Pointer to **string** | ChannelKind is how ingest classified the channel: \&quot;dm\&quot;, \&quot;group\&quot; or \&quot;thread\&quot;. It decides which policy gated the message — dmPolicy for \&quot;dm\&quot;, groupPolicy for BOTH \&quot;group\&quot; and \&quot;thread\&quot;. Only Slack ever reports \&quot;thread\&quot;; Telegram&#39;s reply-to id becomes ReplyTo instead, and Discord&#39;s ingress is guild-scoped so its channels are always \&quot;group\&quot;. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is Unix SECONDS, stamped by the ingest goroutine when the message was accepted — not the transport&#39;s own send time. Rows are dropped 30 days after it. | [optional] 
**Id** | Pointer to **int64** | ID is the store&#39;s row id, assigned on insert — SERVER-SET, and the cursor: pass a page&#39;s last id back as &#x60;since&#x60;. It rises with arrival order but is not contiguous, because one sequence is shared by every org in the store and a caller reads only its own rows. | [optional] 
**ReplyTo** | Pointer to **string** | ReplyTo is the transport&#39;s reply target for this message: Slack&#39;s thread_ts, or the Telegram message id it arrived as. Send it back as the body&#39;s &#x60;replyTo&#x60; to answer in the SAME thread. Empty means the transport reported none — a top-level Slack message, and every Discord and Teams message, since neither carries one — and a reply then lands at the top level of the channel. | [optional] 
**Sender** | Pointer to **string** | Sender is the TRANSPORT-NATIVE user id of whoever wrote the message — a Discord member.user.id, a Slack U… user id, a Teams aadObjectId (falling back to from.id), a Telegram from.id in decimal. Stable per person per transport, and the identity the gate keys on: an allow entry, an access-group member and a pairing approval all name exactly this value. | [optional] 
**SenderUser** | Pointer to **string** | SenderUser is the HANZO account subject that chat identity is linked to, resolved at ingest through the org&#39;s user link. Best-effort and omitted when absent: a person who never linked their chat account — or a link store that could not be read — leaves it empty and is never blocked for it. | [optional] 
**Text** | Pointer to **string** | Text is the body as the transport delivered it, with the bot mention already stripped by the ingress adapter (on Discord it is the /hanzo prompt argument, since that ingress is slash commands only), truncated to 8 KiB on store. Inbound attachments are not stored — this is the whole of what was said. | [optional] 
**Transport** | Pointer to **string** | Transport is the transport this message arrived on — discord, slack, teams or telegram — and the &#x60;:transport&#x60; segment to reply through. | [optional] 

## Methods

### NewTransportInboxView

`func NewTransportInboxView() *TransportInboxView`

NewTransportInboxView instantiates a new TransportInboxView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportInboxViewWithDefaults

`func NewTransportInboxViewWithDefaults() *TransportInboxView`

NewTransportInboxViewWithDefaults instantiates a new TransportInboxView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TransportInboxView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransportInboxView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransportInboxView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TransportInboxView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChannelId

`func (o *TransportInboxView) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *TransportInboxView) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *TransportInboxView) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *TransportInboxView) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannelKind

`func (o *TransportInboxView) GetChannelKind() string`

GetChannelKind returns the ChannelKind field if non-nil, zero value otherwise.

### GetChannelKindOk

`func (o *TransportInboxView) GetChannelKindOk() (*string, bool)`

GetChannelKindOk returns a tuple with the ChannelKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelKind

`func (o *TransportInboxView) SetChannelKind(v string)`

SetChannelKind sets ChannelKind field to given value.

### HasChannelKind

`func (o *TransportInboxView) HasChannelKind() bool`

HasChannelKind returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransportInboxView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransportInboxView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransportInboxView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransportInboxView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *TransportInboxView) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransportInboxView) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransportInboxView) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TransportInboxView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReplyTo

`func (o *TransportInboxView) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *TransportInboxView) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *TransportInboxView) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *TransportInboxView) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSender

`func (o *TransportInboxView) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TransportInboxView) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TransportInboxView) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *TransportInboxView) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSenderUser

`func (o *TransportInboxView) GetSenderUser() string`

GetSenderUser returns the SenderUser field if non-nil, zero value otherwise.

### GetSenderUserOk

`func (o *TransportInboxView) GetSenderUserOk() (*string, bool)`

GetSenderUserOk returns a tuple with the SenderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUser

`func (o *TransportInboxView) SetSenderUser(v string)`

SetSenderUser sets SenderUser field to given value.

### HasSenderUser

`func (o *TransportInboxView) HasSenderUser() bool`

HasSenderUser returns a boolean if a field has been set.

### GetText

`func (o *TransportInboxView) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TransportInboxView) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TransportInboxView) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TransportInboxView) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTransport

`func (o *TransportInboxView) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *TransportInboxView) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *TransportInboxView) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *TransportInboxView) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


