# TeamMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]TeamMessage**](TeamMessage.md) | Messages are the room&#39;s, oldest first, at most &#x60;messageMax&#x60; of them. | [optional] 

## Methods

### NewTeamMessages

`func NewTeamMessages() *TeamMessages`

NewTeamMessages instantiates a new TeamMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMessagesWithDefaults

`func NewTeamMessagesWithDefaults() *TeamMessages`

NewTeamMessagesWithDefaults instantiates a new TeamMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *TeamMessages) GetMessages() []TeamMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *TeamMessages) GetMessagesOk() (*[]TeamMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *TeamMessages) SetMessages(v []TeamMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *TeamMessages) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


