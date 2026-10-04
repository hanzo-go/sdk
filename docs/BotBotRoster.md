# BotBotRoster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bots** | Pointer to [**[]BotBotMember**](BotBotMember.md) | Bots is one entry per bot, each carrying the member account uuid and the Person reference the space roster addresses it by. Empty means the org has no bots — not that the roster could not be read, which is an error. | [optional] 

## Methods

### NewBotBotRoster

`func NewBotBotRoster() *BotBotRoster`

NewBotBotRoster instantiates a new BotBotRoster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotBotRosterWithDefaults

`func NewBotBotRosterWithDefaults() *BotBotRoster`

NewBotBotRosterWithDefaults instantiates a new BotBotRoster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBots

`func (o *BotBotRoster) GetBots() []BotBotMember`

GetBots returns the Bots field if non-nil, zero value otherwise.

### GetBotsOk

`func (o *BotBotRoster) GetBotsOk() (*[]BotBotMember, bool)`

GetBotsOk returns a tuple with the Bots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBots

`func (o *BotBotRoster) SetBots(v []BotBotMember)`

SetBots sets Bots field to given value.

### HasBots

`func (o *BotBotRoster) HasBots() bool`

HasBots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


