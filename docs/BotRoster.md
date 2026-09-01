# BotRoster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bots** | Pointer to [**[]BotMember**](BotMember.md) | Bots is every agent of the caller&#39;s org, projected as a space member. | [optional] 

## Methods

### NewBotRoster

`func NewBotRoster() *BotRoster`

NewBotRoster instantiates a new BotRoster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotRosterWithDefaults

`func NewBotRosterWithDefaults() *BotRoster`

NewBotRosterWithDefaults instantiates a new BotRoster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBots

`func (o *BotRoster) GetBots() []BotMember`

GetBots returns the Bots field if non-nil, zero value otherwise.

### GetBotsOk

`func (o *BotRoster) GetBotsOk() (*[]BotMember, bool)`

GetBotsOk returns a tuple with the Bots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBots

`func (o *BotRoster) SetBots(v []BotMember)`

SetBots sets Bots field to given value.

### HasBots

`func (o *BotRoster) HasBots() bool`

HasBots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


