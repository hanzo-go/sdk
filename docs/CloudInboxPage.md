# CloudInboxPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **int32** | Cursor is the row id to pass back as &#x60;since&#x60; for the next page. It is the last message&#39;s id, or the requested cursor when the page is empty. | [optional] 
**Messages** | Pointer to [**[]CloudInboxView**](CloudInboxView.md) | Messages are the inbound messages, oldest first. | [optional] 

## Methods

### NewCloudInboxPage

`func NewCloudInboxPage() *CloudInboxPage`

NewCloudInboxPage instantiates a new CloudInboxPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInboxPageWithDefaults

`func NewCloudInboxPageWithDefaults() *CloudInboxPage`

NewCloudInboxPageWithDefaults instantiates a new CloudInboxPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *CloudInboxPage) GetCursor() int32`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *CloudInboxPage) GetCursorOk() (*int32, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *CloudInboxPage) SetCursor(v int32)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *CloudInboxPage) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetMessages

`func (o *CloudInboxPage) GetMessages() []CloudInboxView`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CloudInboxPage) GetMessagesOk() (*[]CloudInboxView, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CloudInboxPage) SetMessages(v []CloudInboxView)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CloudInboxPage) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


