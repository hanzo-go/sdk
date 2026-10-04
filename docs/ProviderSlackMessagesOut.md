# ProviderSlackMessagesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the resolved channel ID, reusable for further reads and replies. | [optional] 
**HasMore** | Pointer to **bool** | HasMore indicates that Slack has additional messages beyond this page. | [optional] 
**Messages** | Pointer to [**[]ProviderSlackMessage**](ProviderSlackMessage.md) | Messages is one page, newest first for history and oldest first for threads. | [optional] 
**NextCursor** | Pointer to **string** | NextCursor continues this query; never assume one page is the whole history. | [optional] 

## Methods

### NewProviderSlackMessagesOut

`func NewProviderSlackMessagesOut() *ProviderSlackMessagesOut`

NewProviderSlackMessagesOut instantiates a new ProviderSlackMessagesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackMessagesOutWithDefaults

`func NewProviderSlackMessagesOutWithDefaults() *ProviderSlackMessagesOut`

NewProviderSlackMessagesOutWithDefaults instantiates a new ProviderSlackMessagesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ProviderSlackMessagesOut) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProviderSlackMessagesOut) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProviderSlackMessagesOut) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProviderSlackMessagesOut) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetHasMore

`func (o *ProviderSlackMessagesOut) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ProviderSlackMessagesOut) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ProviderSlackMessagesOut) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ProviderSlackMessagesOut) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetMessages

`func (o *ProviderSlackMessagesOut) GetMessages() []ProviderSlackMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ProviderSlackMessagesOut) GetMessagesOk() (*[]ProviderSlackMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ProviderSlackMessagesOut) SetMessages(v []ProviderSlackMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ProviderSlackMessagesOut) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNextCursor

`func (o *ProviderSlackMessagesOut) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ProviderSlackMessagesOut) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ProviderSlackMessagesOut) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ProviderSlackMessagesOut) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


