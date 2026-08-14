# BotRuns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bots** | Pointer to [**[]BotRun**](BotRun.md) | Bots is the org&#39;s live runs. Always an array, never null. | [optional] 

## Methods

### NewBotRuns

`func NewBotRuns() *BotRuns`

NewBotRuns instantiates a new BotRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotRunsWithDefaults

`func NewBotRunsWithDefaults() *BotRuns`

NewBotRunsWithDefaults instantiates a new BotRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBots

`func (o *BotRuns) GetBots() []BotRun`

GetBots returns the Bots field if non-nil, zero value otherwise.

### GetBotsOk

`func (o *BotRuns) GetBotsOk() (*[]BotRun, bool)`

GetBotsOk returns a tuple with the Bots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBots

`func (o *BotRuns) SetBots(v []BotRun)`

SetBots sets Bots field to given value.

### HasBots

`func (o *BotRuns) HasBots() bool`

HasBots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


