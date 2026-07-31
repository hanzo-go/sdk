# CloudBotRoster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bots** | Pointer to [**[]CloudBotMember**](CloudBotMember.md) | Bots is every agent of the caller&#39;s org, projected as a workspace member. | [optional] 

## Methods

### NewCloudBotRoster

`func NewCloudBotRoster() *CloudBotRoster`

NewCloudBotRoster instantiates a new CloudBotRoster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBotRosterWithDefaults

`func NewCloudBotRosterWithDefaults() *CloudBotRoster`

NewCloudBotRosterWithDefaults instantiates a new CloudBotRoster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBots

`func (o *CloudBotRoster) GetBots() []CloudBotMember`

GetBots returns the Bots field if non-nil, zero value otherwise.

### GetBotsOk

`func (o *CloudBotRoster) GetBotsOk() (*[]CloudBotMember, bool)`

GetBotsOk returns a tuple with the Bots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBots

`func (o *CloudBotRoster) SetBots(v []CloudBotMember)`

SetBots sets Bots field to given value.

### HasBots

`func (o *CloudBotRoster) HasBots() bool`

HasBots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


