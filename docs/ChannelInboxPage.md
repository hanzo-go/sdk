# ChannelInboxPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **int64** | Cursor is the row id to pass back as &#x60;since&#x60; for the next page. It is the last message&#39;s id, or the requested cursor when the page is empty. | [optional] 
**Messages** | Pointer to [**[]ChannelInboxView**](ChannelInboxView.md) | Messages are the inbound messages, oldest first. | [optional] 

## Methods

### NewChannelInboxPage

`func NewChannelInboxPage() *ChannelInboxPage`

NewChannelInboxPage instantiates a new ChannelInboxPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelInboxPageWithDefaults

`func NewChannelInboxPageWithDefaults() *ChannelInboxPage`

NewChannelInboxPageWithDefaults instantiates a new ChannelInboxPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ChannelInboxPage) GetCursor() int64`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ChannelInboxPage) GetCursorOk() (*int64, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ChannelInboxPage) SetCursor(v int64)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ChannelInboxPage) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetMessages

`func (o *ChannelInboxPage) GetMessages() []ChannelInboxView`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChannelInboxPage) GetMessagesOk() (*[]ChannelInboxView, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChannelInboxPage) SetMessages(v []ChannelInboxView)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ChannelInboxPage) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


