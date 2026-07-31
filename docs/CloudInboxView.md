# CloudInboxView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ReplyTo** | Pointer to **string** |  | [optional] 
**RoomId** | Pointer to **string** |  | [optional] 
**RoomKind** | Pointer to **string** |  | [optional] 
**Sender** | Pointer to **string** |  | [optional] 
**SenderUser** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudInboxView

`func NewCloudInboxView() *CloudInboxView`

NewCloudInboxView instantiates a new CloudInboxView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInboxViewWithDefaults

`func NewCloudInboxViewWithDefaults() *CloudInboxView`

NewCloudInboxViewWithDefaults instantiates a new CloudInboxView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudInboxView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudInboxView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudInboxView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudInboxView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChannel

`func (o *CloudInboxView) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CloudInboxView) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CloudInboxView) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CloudInboxView) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudInboxView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudInboxView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudInboxView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudInboxView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudInboxView) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudInboxView) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudInboxView) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CloudInboxView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReplyTo

`func (o *CloudInboxView) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *CloudInboxView) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *CloudInboxView) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *CloudInboxView) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetRoomId

`func (o *CloudInboxView) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *CloudInboxView) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *CloudInboxView) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *CloudInboxView) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetRoomKind

`func (o *CloudInboxView) GetRoomKind() string`

GetRoomKind returns the RoomKind field if non-nil, zero value otherwise.

### GetRoomKindOk

`func (o *CloudInboxView) GetRoomKindOk() (*string, bool)`

GetRoomKindOk returns a tuple with the RoomKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomKind

`func (o *CloudInboxView) SetRoomKind(v string)`

SetRoomKind sets RoomKind field to given value.

### HasRoomKind

`func (o *CloudInboxView) HasRoomKind() bool`

HasRoomKind returns a boolean if a field has been set.

### GetSender

`func (o *CloudInboxView) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *CloudInboxView) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *CloudInboxView) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *CloudInboxView) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSenderUser

`func (o *CloudInboxView) GetSenderUser() string`

GetSenderUser returns the SenderUser field if non-nil, zero value otherwise.

### GetSenderUserOk

`func (o *CloudInboxView) GetSenderUserOk() (*string, bool)`

GetSenderUserOk returns a tuple with the SenderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUser

`func (o *CloudInboxView) SetSenderUser(v string)`

SetSenderUser sets SenderUser field to given value.

### HasSenderUser

`func (o *CloudInboxView) HasSenderUser() bool`

HasSenderUser returns a boolean if a field has been set.

### GetText

`func (o *CloudInboxView) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CloudInboxView) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CloudInboxView) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CloudInboxView) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


