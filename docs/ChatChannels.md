# ChatChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]ChannelView**](ChannelView.md) | Channels is every chat transport this deployment supports, in a fixed order, each carrying whether the org has connected it, the account behind the connection, what the transport can do, the org&#39;s DM/group access policies for it, and how many pairing requests are waiting. | [optional] 

## Methods

### NewChatChannels

`func NewChatChannels() *ChatChannels`

NewChatChannels instantiates a new ChatChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatChannelsWithDefaults

`func NewChatChannelsWithDefaults() *ChatChannels`

NewChatChannelsWithDefaults instantiates a new ChatChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ChatChannels) GetChannels() []ChannelView`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ChatChannels) GetChannelsOk() (*[]ChannelView, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ChatChannels) SetChannels(v []ChannelView)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ChatChannels) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


