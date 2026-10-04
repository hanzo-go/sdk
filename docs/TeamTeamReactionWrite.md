# TeamTeamReactionWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emoji** | Pointer to **string** | Emoji is the reaction, from the path (percent-encoded on the wire). | [optional] 
**Id** | Pointer to **string** | ID is the message, from the path. | [optional] 
**Space** | Pointer to **string** | Space names the space holding the message. Body-only. | [optional] 

## Methods

### NewTeamTeamReactionWrite

`func NewTeamTeamReactionWrite() *TeamTeamReactionWrite`

NewTeamTeamReactionWrite instantiates a new TeamTeamReactionWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamReactionWriteWithDefaults

`func NewTeamTeamReactionWriteWithDefaults() *TeamTeamReactionWrite`

NewTeamTeamReactionWriteWithDefaults instantiates a new TeamTeamReactionWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmoji

`func (o *TeamTeamReactionWrite) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *TeamTeamReactionWrite) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *TeamTeamReactionWrite) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *TeamTeamReactionWrite) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamReactionWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamReactionWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamReactionWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamReactionWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamReactionWrite) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamReactionWrite) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamReactionWrite) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamReactionWrite) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


