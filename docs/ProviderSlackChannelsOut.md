# ProviderSlackChannelsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]ProviderSlackConversation**](ProviderSlackConversation.md) | Channels contains one page of conversations visible to the bot. | [optional] 
**NextCursor** | Pointer to **string** | NextCursor is empty when there are no more pages. | [optional] 

## Methods

### NewProviderSlackChannelsOut

`func NewProviderSlackChannelsOut() *ProviderSlackChannelsOut`

NewProviderSlackChannelsOut instantiates a new ProviderSlackChannelsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackChannelsOutWithDefaults

`func NewProviderSlackChannelsOutWithDefaults() *ProviderSlackChannelsOut`

NewProviderSlackChannelsOutWithDefaults instantiates a new ProviderSlackChannelsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ProviderSlackChannelsOut) GetChannels() []ProviderSlackConversation`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ProviderSlackChannelsOut) GetChannelsOk() (*[]ProviderSlackConversation, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ProviderSlackChannelsOut) SetChannels(v []ProviderSlackConversation)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ProviderSlackChannelsOut) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetNextCursor

`func (o *ProviderSlackChannelsOut) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ProviderSlackChannelsOut) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ProviderSlackChannelsOut) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ProviderSlackChannelsOut) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


